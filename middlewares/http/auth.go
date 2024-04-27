package http

import (
	"context"
	"strings"

	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcenter/apps/endpoint"
	"github.com/infraboard/mcenter/apps/policy"
	"github.com/infraboard/mcenter/apps/token"
	"github.com/infraboard/mcenter/apps/user"
	"github.com/infraboard/mcube/v2/exception"
	"github.com/infraboard/mcube/v2/http/restful/response"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/application"
	"github.com/infraboard/mcube/v2/ioc/config/cache"
	"github.com/infraboard/mcube/v2/ioc/config/log"
	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel/trace"

	"github.com/infraboard/mcube/v2/ioc/config/gorestful"
)

func init() {
	ioc.Config().Registry(&httpAuther{
		codeCheckSilence: 30 * 60,
	})
}

type httpAuther struct {
	log              *zerolog.Logger
	tk               token.Service
	cache            cache.Cache
	policy           policy.Service
	codeCheckSilence int64

	ioc.ObjectImpl
}

func (a *httpAuther) Init() error {
	a.log = log.Sub("auther.http")
	a.tk = ioc.Controller().Get(token.AppName).(token.Service)
	a.policy = ioc.Controller().Get(policy.AppName).(policy.Service)
	a.cache = cache.C()

	// 注册认证中间件
	gorestful.RootRouter().Filter(a.GoRestfulAuthFunc)
	return nil
}

func (a *httpAuther) Name() string {
	return "http_auth"
}

// 设置静默时长
func (a *httpAuther) SetCodeCheckSilenceTime(t int64) {
	a.codeCheckSilence = t
}

func (a *httpAuther) GoRestfulAuthFunc(req *restful.Request, resp *restful.Response, next *restful.FilterChain) {
	// 请求拦截
	entry := endpoint.NewEntryFromRestRequest(req)

	if entry != nil && entry.AuthEnable {
		// 访问令牌校验
		tk, err := a.CheckAccessToken(req)
		if err != nil {
			response.Failed(resp, err)
			return
		}

		// 接口调用权限校验
		err = a.CheckPermission(req, tk, entry)
		if err != nil {
			response.Failed(resp, err)
			return
		}

		// 验证码校验(双因子认证)
		if !a.IsCodeCheckSilence(req.Request.Context(), tk.Username) && entry.CodeEnable {
			_, err := a.CheckCode(req, tk)
			if err != nil {
				response.Failed(resp, err)
				return
			}
		}
	}

	// 补充TraceId
	span := trace.SpanFromContext(req.Request.Context())
	resp.AddHeader(response.TraceHeaderKey, span.SpanContext().TraceID().String())
	next.ProcessFilter(req, resp)
}

func (a *httpAuther) CheckAccessToken(req *restful.Request) (*token.Token, error) {
	// 获取用户Token, Token放在Heander Authorization
	ak := token.GetAccessTokenFromHTTP(req.Request)

	if ak == "" {
		return nil, token.ErrTokenUnauthorized
	}

	// 调用GRPC 校验用户Token合法性
	tk, err := a.tk.ValidateToken(req.Request.Context(), token.NewValidateTokenRequest(ak))
	if err != nil {
		return nil, err
	}

	// 是不是需要返回用户的认证信息: 那个人, 那个空间下面， token本身的信息
	req.SetAttribute(token.TOKEN_ATTRIBUTE_NAME, tk)
	return tk, nil
}

func (a *httpAuther) CheckPermission(req *restful.Request, tk *token.Token, e *endpoint.Entry) error {
	if tk == nil {
		return exception.NewUnauthorized("validate permission need token")
	}

	// 如果是超级管理员不做权限校验, 直接放行
	if tk.UserType.IsIn(user.TYPE_SUPPER) {
		a.log.Debug().Msgf("[%s] supper admin skip permission check!", tk.Username)
		return nil
	}

	switch strings.ToUpper(e.PermissionMode) {
	case "ACL":
		return a.validatePermissionByACL(tk, e)
	default:
		return a.validatePermissionByPRBAC(req, tk, e)
	}
}

func (a *httpAuther) validatePermissionByACL(tk *token.Token, e *endpoint.Entry) error {
	// 检查是否是允许的类型
	if len(e.Allow) > 0 {
		a.log.Debug().Msgf("[%s] start check permission to mcenter ...", tk.Username)
		if !e.IsAllow(tk.UserType) {
			return exception.NewPermissionDeny("no permission, allow: %s, but current: %s", e.Allow, tk.UserType)
		}
		a.log.Debug().Msgf("[%s] permission check passed", tk.Username)
	}
	return nil
}

func (a *httpAuther) validatePermissionByPRBAC(r *restful.Request, tk *token.Token, e *endpoint.Entry) error {
	if !e.PermissionEnable {
		a.log.Debug().Msgf("permission check disabled, skip permission check")
		return nil
	}

	req := policy.NewCheckPermissionRequest()
	req.Domain = tk.Domain
	req.Namespace = tk.Namespace
	req.UserId = tk.UserId
	req.ServiceId = application.Get().AppName
	req.Path = e.UniquePath()
	a.log.Debug().Msgf("permission check request: %s", req.ToJSON())

	perm, err := a.policy.CheckPermission(r.Request.Context(), req)
	if err != nil {
		return exception.NewPermissionDeny(err.Error())
	}
	a.log.Debug().Msgf("[%s] permission check passed", tk.Username)

	// 保存访问访问信息
	r.SetAttribute(policy.SCOPE_ATTRIBUTE_NAME, perm.Scope)
	return nil
}

func (a *httpAuther) CheckCode(req *restful.Request, tk *token.Token) (*token.Code, error) {
	// 获取用户Code, Code放在Heander X-MCENTER-CODE
	cdStr := token.GetCodeFromHTTP(req.Request)
	if cdStr == "" {
		return nil, token.ErrCodeUnauthorized
	}

	// 调用GRPC 校验用户Code合法性
	cd, err := a.tk.VerifyCode(req.Request.Context(), token.NewVerifyCodeRequest(tk.Username, cdStr))
	if err != nil {
		return nil, err
	}

	// 保存返回的Code信息
	req.SetAttribute(token.CODE_ATTRIBUTE_NAME, cd)
	// 加入静默池中
	a.SetCodeCheckSilence(req.Request.Context(), cd)
	return cd, nil
}

func (a *httpAuther) SetCodeCheckSilence(ctx context.Context, c *token.Code) {
	err := a.cache.Set(ctx, c.Key(), c.Code, cache.WithExpiration(a.codeCheckSilence))
	if err != nil {
		a.log.Error().Msgf("set code Silence to cache error, %s", err)
	}
}

func (a *httpAuther) IsCodeCheckSilence(ctx context.Context, username string) bool {
	err := a.cache.Exist(ctx, token.NewCodeKey(username))
	if err != nil {
		a.log.Error().Msgf("check error, %s", err)
		return false
	}
	return true
}
