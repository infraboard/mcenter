package middleware

import (
	"context"
	"fmt"

	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/restful/response"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/mcenter/apps/endpoint"
	"github.com/infraboard/mcenter/apps/permission"
	"github.com/infraboard/mcenter/apps/token"
	"github.com/infraboard/mcenter/apps/user"
	"github.com/infraboard/mcenter/client/rpc"
)

// RestfulServerInterceptor go-restful认证中间件
func RestfulServerInterceptor() restful.FilterFunction {
	return newhttpAuther().GoRestfulAuthFunc
}

// 给服务端提供的RESTful接口的 认证与鉴权中间件
func newhttpAuther() *httpAuther {
	return &httpAuther{
		log:    zap.L().Named("auther.http"),
		client: rpc.C(),
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
}

// 是否开启权限的控制, 交给中间件使用方去觉得
func (a *httpAuther) GoRestfulAuthFunc(req *restful.Request, resp *restful.Response, next *restful.FilterChain) {
	// 请求拦截
	entry := endpoint.NewEntryFromRestRequest(req)

	if entry != nil && entry.AuthEnable {
		// 获取用户Token, Token放在Heander Authorization
		ak := token.GetTokenFromHTTPHeader(req.Request)

		// 调用GRPC 校验用户Token合法性
		tk, err := a.client.Token().ValidateToken(req.Request.Context(), token.NewValidateTokenRequest(ak))
		if err != nil {
			response.Failed(resp, err)
			return
		}

		// 是不是需要返回用户的认证信息: 那个人, 那个空间下面， token本身的信息
		req.SetAttribute(token.TOKEN_ATTRIBUTE_NAME, tk)

		if entry.PermissionEnable {
			// 权限检查
			err := a.CheckPermission(req.Request.Context(), tk, entry)
			if err != nil {
				response.Failed(resp, err)
				return
			}
		}
	}

	// next flow
	next.ProcessFilter(req, resp)
}

func (a *httpAuther) CheckPermission(ctx context.Context, tk *token.Token, e *endpoint.Entry) error {
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
		return a.ValidatePermissionByACL(ctx, tk, e)
	case PRBAC_MODE:
		return a.ValidatePermissionByPRBAC(ctx, tk, e)
	default:
		return fmt.Errorf("only support acl and prbac")
	}
}

func (a *httpAuther) ValidatePermissionByACL(ctx context.Context, tk *token.Token, e *endpoint.Entry) error {
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

func (a *httpAuther) ValidatePermissionByPRBAC(ctx context.Context, tk *token.Token, e *endpoint.Entry) error {
	ci, err := a.client.ClientInfo(ctx)
	if err != nil {
		return err
	}

	req := permission.NewCheckPermissionRequest()
	req.Username = tk.Username
	req.Namespace = tk.Namespace
	req.ServiceId = ci.Id
	req.Path = e.UniquePath()
	_, err = a.client.Permission().CheckPermission(ctx, req)
	if err != nil {
		return exception.NewPermissionDeny(err.Error())
	}
	a.log.Debugf("[%s] permission check passed", tk.Username)
	return nil
}
