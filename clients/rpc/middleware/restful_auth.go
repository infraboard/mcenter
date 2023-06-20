package middleware

import (
	"fmt"
	"time"

	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/cache"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/restful/response"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/mcenter/apps/code"
	"github.com/infraboard/mcenter/apps/endpoint"
	"github.com/infraboard/mcenter/apps/policy"
	"github.com/infraboard/mcenter/apps/token"
	"github.com/infraboard/mcenter/apps/user"
	"github.com/infraboard/mcenter/clients/rpc"
)

// RestfulServerInterceptor go-restful认证中间件
func RestfulServerInterceptor() restful.FilterFunction {
	return newhttpAuther().GoRestfulAuthFunc
}

// 给服务端提供的RESTful接口的 认证与鉴权中间件
func newhttpAuther() *httpAuther {
	return &httpAuther{
		log:              zap.L().Named("auther.http"),
		client:           rpc.C(),
		cache:            cache.C(),
		codeCheckSilence: 30 * time.Minute,
		mode:             PRBAC_MODE,
	}
}

type PermissionMode int

const (
	// PRBAC_MODE 基于策略的权限校验
	PRBAC_MODE PermissionMode = 1
	// ACL_MODE 基于用户类型的权限校验
	ACL_MODE PermissionMode = 2
)

type httpAuther struct {
	log logger.Logger
	// 基于rpc客户端进行封装
	client *rpc.ClientSet
	// 鉴权模式
	mode PermissionMode
	// 缓存
	cache cache.Cache
	// 校验码检查静默时长, 默认值30分钟, 30分钟之内只检查一次
	codeCheckSilence time.Duration
}

// 设置权限校验策略
func (a *httpAuther) SetPermissionMode(m PermissionMode) {
	a.mode = m
}

// 设置静默时长
func (a *httpAuther) SetCodeCheckSilenceTime(t time.Duration) {
	a.codeCheckSilence = t
}

// 是否开启权限的控制, 交给中间件使用方去觉得
func (a *httpAuther) GoRestfulAuthFunc(req *restful.Request, resp *restful.Response, next *restful.FilterChain) {
	// 权限检查
	if err := a.PermissionCheck(req, resp); err != nil {
		response.Failed(resp, err)
		return
	}

	// next flow
	next.ProcessFilter(req, resp)
}

func (a *httpAuther) PermissionCheck(req *restful.Request, resp *restful.Response) error {
	// 请求拦截
	entry := endpoint.NewEntryFromRestRequest(req)

	if entry != nil && entry.AuthEnable {
		// 访问令牌校验
		tk, err := a.CheckAccessToken(req)
		if err != nil {
			return err
		}

		// 接口调用权限校验
		if entry.PermissionEnable {
			err := a.CheckPermission(req, tk, entry)
			if err != nil {
				return err
			}
		}

		// 验证码校验(双因子认证)
		if !a.IsCodeCheckSilence(tk.Username) && entry.CodeEnable {
			_, err := a.CheckCode(req, tk)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (a *httpAuther) CheckAccessToken(req *restful.Request) (*token.Token, error) {
	// 获取用户Token, Token放在Heander Authorization
	ak := token.GetAccessTokenFromHTTP(req.Request)

	if ak == "" {
		return nil, token.ErrUnauthorized
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

func (a *httpAuther) CheckCode(req *restful.Request, tk *token.Token) (*code.Code, error) {
	// 获取用户Code, Code放在Heander X-MCENTER-CODE
	cdStr := code.GetCodeFromHTTP(req.Request)
	if cdStr == "" {
		return nil, code.ErrUnauthorized
	}

	// 调用GRPC 校验用户Code合法性
	cd, err := a.client.Code().VerifyCode(req.Request.Context(), code.NewVerifyCodeRequest(tk.Username, cdStr))
	if err != nil {
		return nil, err
	}

	// 保存返回的Code信息
	req.SetAttribute(code.CODE_ATTRIBUTE_NAME, cd)
	// 加入静默池中
	a.SetCodeCheckSilence(cd)
	return cd, nil
}

func (a *httpAuther) SetCodeCheckSilence(c *code.Code) {
	err := a.cache.PutWithTTL(c.Key(), c.Code, a.codeCheckSilence)
	if err != nil {
		a.log.Errorf("set code Silence to cache error, %s", err)
	}
}

func (a *httpAuther) IsCodeCheckSilence(username string) bool {
	return a.cache.IsExist(code.NewCodeKey(username))
}

func (a *httpAuther) CheckPermission(req *restful.Request, tk *token.Token, e *endpoint.Entry) error {
	if tk == nil {
		return exception.NewUnauthorized("validate permission need token")
	}

	// 如果是超级管理员不做权限校验, 直接放行
	if tk.UserType.IsIn(user.TYPE_SUPPER) {
		a.log.Debugf("[%s] supper admin skip permission check!", tk.Username)
		return nil
	}

	switch a.mode {
	case ACL_MODE:
		return a.ValidatePermissionByACL(req, tk, e)
	case PRBAC_MODE:
		return a.ValidatePermissionByPRBAC(req, tk, e)
	default:
		return fmt.Errorf("only support acl and prbac")
	}
}

func (a *httpAuther) ValidatePermissionByACL(req *restful.Request, tk *token.Token, e *endpoint.Entry) error {
	// 检查是否是允许的类型
	if len(e.Allow) > 0 {
		a.log.Debugf("[%s] start check permission to keyauth ...", tk.Username)
		if !e.IsAllow(tk.UserType) {
			return exception.NewPermissionDeny("no permission, allow: %s, but current: %s", e.Allow, tk.UserType)
		}
		a.log.Debugf("[%s] permission check passed", tk.Username)
	}

	return nil
}

func (a *httpAuther) ValidatePermissionByPRBAC(r *restful.Request, tk *token.Token, e *endpoint.Entry) error {
	ci, err := a.client.ClientInfo(r.Request.Context())
	if err != nil {
		return err
	}

	req := policy.NewCheckPermissionRequest()
	req.Username = tk.Username
	req.Namespace = tk.Namespace
	req.ServiceId = ci.Meta.Id
	req.Path = e.UniquePath()
	perm, err := a.client.Policy().CheckPermission(r.Request.Context(), req)
	if err != nil {
		return exception.NewPermissionDeny(err.Error())
	}
	a.log.Debugf("[%s] permission check passed", tk.Username)

	// 保存访问访问信息
	r.SetAttribute(policy.SCOPE_ATTRIBUTE_NAME, perm.Scope)
	return nil
}
