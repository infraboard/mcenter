package middleware

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/response"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/mcenter/apps/endpoint"
	"github.com/infraboard/mcenter/apps/token"
	"github.com/infraboard/mcenter/client/rest"
)

// RestfulServerInterceptor go-restful认证中间件
func RestfulServerInterceptor() restful.FilterFunction {
	return newhttpAuther().GoRestfulAuthFunc
}

// 给服务端提供的RESTful接口的 认证与鉴权中间件
func newhttpAuther() *httpAuther {
	return &httpAuther{
		log:    zap.L().Named("auther.http"),
		client: rest.C(),
	}
}

type httpAuther struct {
	log logger.Logger
	// 基于rest客户端进行封装
	client *rest.ClientSet
}

// 是否开启权限的控制, 交给中间件使用方去觉得
func (a *httpAuther) GoRestfulAuthFunc(req *restful.Request, resp *restful.Response, next *restful.FilterChain) {
	// 请求拦截
	entry := endpoint.NewEntryFromRestRequest(req)

	// 有认证标签,并且开启了认证
	if entry.AuthEnable {
		// 获取用户Token, Token放在Heander Authorization
		ak := token.GetTokenFromHTTPHeader(req.Request)

		// 调用GRPC 校验用户Token合法性
		tk, err := a.client.Token().ValidateToken(req.Request.Context(), token.NewValidateTokenRequest(ak))
		if err != nil {
			response.Failed(resp.ResponseWriter, err)
			return
		}

		// 是不是需要返回用户的认证信息: 那个人, 那个空间下面， token本身的信息
		req.SetAttribute("token", tk)

		// 判断用户权限
		if entry.PermissionEnable {
			a.log.Debugf("perm not impl ...")
		}
	}

	// next flow
	next.ProcessFilter(req, resp)
}
