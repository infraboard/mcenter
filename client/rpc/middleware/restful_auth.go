package middleware

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/response"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/mcenter/apps/token"
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

type httpAuther struct {
	log logger.Logger
	// 基于rpc客户端进行封装
	client *rpc.ClientSet
}

// 是否开启权限的控制, 交给中间件使用方去觉得
func (a *httpAuther) GoRestfulAuthFunc(req *restful.Request, resp *restful.Response, next *restful.FilterChain) {
	// 请求拦截
	meta := req.SelectedRoute().Metadata()
	a.log.Debugf("route meta: %s", meta)

	isAuth, ok := meta[label.Auth]
	// 有认证标签,并且开启了认证
	if ok && isAuth.(bool) {
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
		isPerm, ok := meta[label.Permission]
		if ok && isPerm.(bool) {
			a.log.Debugf("perm not impl ...")
		}
	}

	// next flow
	next.ProcessFilter(req, resp)
}
