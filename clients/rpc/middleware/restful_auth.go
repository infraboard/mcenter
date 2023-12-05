package middleware

import (
	"context"
	"strings"
	"sync"

	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/v2/exception"
	"github.com/infraboard/mcube/v2/http/restful/response"
	"github.com/infraboard/mcube/v2/ioc/config/cache"
	"github.com/infraboard/mcube/v2/ioc/config/logger"
	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel/trace"

	"github.com/infraboard/mcenter/apps/endpoint"
	"github.com/infraboard/mcenter/apps/policy"
	"github.com/infraboard/mcenter/apps/token"
	"github.com/infraboard/mcenter/apps/user"
	"github.com/infraboard/mcenter/clients/rpc"
)

var (
	httpAuther *HttpAuther
	lock       sync.Mutex
)

// RestfulServerInterceptor go-restful认证中间件
func RestfulServerInterceptor() restful.FilterFunction {
	return GetHttpAuther().GoRestfulAuthFunc
}

func GetHttpAuther() *HttpAuther {
	lock.Lock()
	defer lock.Unlock()

	if httpAuther == nil {
		httpAuther = NewhttpAuther()
	}

	return httpAuther
}

// 给服务端提供的RESTful接口的 认证与鉴权中间件
func NewhttpAuther() *HttpAuther {
	return &HttpAuther{
		log:              logger.Sub("auther.http"),
		client:           rpc.C(),
		cache:            cache.C(),
		codeCheckSilence: 30 * 60,
	}
}

type HttpAuther struct {
	log *zerolog.Logger
	// 基于rpc客户端进行封装
	client *rpc.ClientSet
	// 缓存
	cache cache.Cache
	// 校验码检查静默时长, 默认值30分钟, 30分钟之内只检查一次
	codeCheckSilence int64
}

// 设置静默时长
func (a *HttpAuther) SetCodeCheckSilenceTime(t int64) {
	a.codeCheckSilence = t
}

// 是否开启权限的控制, 交给中间件使用方去觉得
func (a *HttpAuther) GoRestfulAuthFunc(req *restful.Request, resp *restful.Response, next *restful.FilterChain) {
	// 请求拦截, 权限检查
	entry := endpoint.NewEntryFromRestRequest(req)
	if err := a.PermissionCheck(req, resp, entry); err != nil {
		response.Failed(resp, err)
		return
	}

	// 补充TraceId
	span := trace.SpanFromContext(req.Request.Context())
	resp.AddHeader(response.TraceHeaderKey, span.SpanContext().TraceID().String())

	// next flow
	next.ProcessFilter(req, resp)
}

func (a *HttpAuther) PermissionCheck(req *restful.Request, resp *restful.Response, entry *endpoint.Entry) error {
	if entry != nil && entry.AuthEnable {
		// 访问令牌校验
		tk, err := a.checkAccessToken(req)
		if err != nil {
			return err
		}

		// 接口调用权限校验
		if entry.PermissionEnable {
			err := a.checkPermission(req, tk, entry)
			if err != nil {
				return err
			}
		}

		// 验证码校验(双因子认证)
		if !a.isCodeCheckSilence(req.Request.Context(), tk.Username) && entry.CodeEnable {
			_, err := a.checkCode(req, tk)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (a *HttpAuther) checkAccessToken(req *restful.Request) (*token.Token, error) {
	// 获取用户Token, Token放在Heander Authorization
	ak := token.GetAccessTokenFromHTTP(req.Request)

	if ak == "" {
		return nil, token.ErrTokenUnauthorized
	}

	// 调用GRPC 校验用户Token合法性
	tk, err := a.client.Token().ValidateToken(req.Request.Context(), token.NewValidateTokenRequest(ak))
	if err != nil {
		return nil, err
	}

	// 是不是需要返回用户的认证信息: 那个人, 那个空间下面， token本身的信息
	req.SetAttribute(token.TOKEN_ATTRIBUTE_NAME, tk)
	return tk, nil
}

func (a *HttpAuther) checkCode(req *restful.Request, tk *token.Token) (*token.Code, error) {
	// 获取用户Code, Code放在Heander X-MCENTER-CODE
	cdStr := token.GetCodeFromHTTP(req.Request)
	if cdStr == "" {
		return nil, token.ErrCodeUnauthorized
	}

	// 调用GRPC 校验用户Code合法性
	cd, err := a.client.Token().VerifyCode(req.Request.Context(), token.NewVerifyCodeRequest(tk.Username, cdStr))
	if err != nil {
		return nil, err
	}

	// 保存返回的Code信息
	req.SetAttribute(token.CODE_ATTRIBUTE_NAME, cd)
	// 加入静默池中
	a.setCodeCheckSilence(req.Request.Context(), cd)
	return cd, nil
}

func (a *HttpAuther) setCodeCheckSilence(ctx context.Context, c *token.Code) {
	err := a.cache.Set(ctx, c.Key(), c.Code, cache.WithExpiration(a.codeCheckSilence))
	if err != nil {
		a.log.Error().Msgf("set code Silence to cache error, %s", err)
	}
}

func (a *HttpAuther) isCodeCheckSilence(ctx context.Context, username string) bool {
	err := a.cache.Exist(ctx, token.NewCodeKey(username))
	if err != nil {
		a.log.Error().Msgf("check error, %s", err)
		return false
	}
	return true
}

func (a *HttpAuther) checkPermission(req *restful.Request, tk *token.Token, e *endpoint.Entry) error {
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
		return a.validatePermissionByACL(req, tk, e)
	default:
		return a.validatePermissionByPRBAC(req, tk, e)
	}
}

func (a *HttpAuther) validatePermissionByACL(req *restful.Request, tk *token.Token, e *endpoint.Entry) error {
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

func (a *HttpAuther) validatePermissionByPRBAC(r *restful.Request, tk *token.Token, e *endpoint.Entry) error {
	ci, err := a.client.ClientInfo(r.Request.Context())
	if err != nil {
		return err
	}

	req := policy.NewCheckPermissionRequest()
	req.Domain = tk.Domain
	req.Namespace = tk.Namespace
	req.Username = tk.Username
	req.ServiceId = ci.Meta.Id
	req.Path = e.UniquePath()
	a.log.Debug().Msgf("permission check request: %s", req.ToJSON())

	perm, err := a.client.Policy().CheckPermission(r.Request.Context(), req)
	if err != nil {
		return exception.NewPermissionDeny(err.Error())
	}
	a.log.Debug().Msgf("[%s] permission check passed", tk.Username)

	// 保存访问访问信息
	r.SetAttribute(policy.SCOPE_ATTRIBUTE_NAME, perm.Scope)
	return nil
}
