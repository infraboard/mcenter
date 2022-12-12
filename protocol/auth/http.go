package auth

import (
	"fmt"

	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcenter/apps/token"
	"github.com/infraboard/mcenter/apps/user"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/restful/response"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func NewHttpAuther() *httpAuther {
	return &httpAuther{
		log: zap.L().Named("auther.http"),
		tk:  app.GetInternalApp(token.AppName).(token.Service),
	}
}

type httpAuther struct {
	log logger.Logger
	tk  token.Service
}

func (a *httpAuther) GoRestfulAuthFunc(req *restful.Request, resp *restful.Response, next *restful.FilterChain) {
	// 请求拦截
	meta := req.SelectedRoute().Metadata()
	a.log.Debugf("route meta: %s", meta)

	isAuth, ok := meta[label.Auth]
	// 有认证标签,并且开启了认证
	if ok && isAuth.(bool) {
		// 获取token
		ak := token.GetAccessTokenFromHTTP(req.Request)

		if ak == "" {
			response.Failed(resp, token.ErrUnauthorized)
			return
		}

		tk, err := a.tk.ValidateToken(req.Request.Context(), token.NewValidateTokenRequest(ak))
		if err != nil {
			response.Failed(resp, err)
			return
		}

		// 判断用户权限
		v, ok := meta[label.Allow]
		if ok {
			ut := v.(user.TYPE)
			// 权限的编号来判断
			if tk.UserType < ut {
				response.Failed(resp, fmt.Errorf("permission deny: %s, required: %s", tk.UserType, ut))
				return
			}
		}
	}

	// next flow
	next.ProcessFilter(req, resp)

	// 响应拦截
	// a.log.Debugf("%s", resp)
}
