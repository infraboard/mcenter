package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/restful/response"
	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mcube/ioc/config/logger"
	"github.com/rs/zerolog"

	"github.com/infraboard/mcenter/apps/domain"
	"github.com/infraboard/mcenter/apps/token"
)

func init() {
	ioc.Api().Registry(&oath2Handler{})
}

type oath2Handler struct {
	service token.Service
	log     *zerolog.Logger
	ioc.ObjectImpl
}

func (h *oath2Handler) Init() error {
	h.log = logger.Sub(token.AppName)
	h.service = ioc.GetController(token.AppName).(token.Service)
	return nil
}

func (h *oath2Handler) Name() string {
	return "oauth2"
}

func (h *oath2Handler) Version() string {
	return "v1"
}

func (h *oath2Handler) Registry(ws *restful.WebService) {
	tags := []string{"第三方登陆"}

	code := ws.QueryParameter("code", "oauth2 auth code").DataType("string").Required(true)
	stat := ws.QueryParameter("state", "oauth2 state").DataType("string").Required(false)
	dom := ws.QueryParameter("domain", "auth domain").DataType("string").DefaultValue(domain.DEFAULT_DOMAIN)

	ws.Route(ws.GET("/feishu").To(h.FeishuOauth2Auth).
		Doc("飞书登陆").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Auth, label.Disable).
		Metadata(label.PERMISSION_MODE, label.PERMISSION_MODE_ACL.Value()).
		Param(code).
		Param(stat).
		Param(dom).
		Writes(token.Token{}).
		Returns(200, "OK", token.Token{}))

	ws.Route(ws.GET("/dingding").To(h.DingDingOauth2Auth).
		Doc("钉钉登陆").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Auth, label.Disable).
		Metadata(label.PERMISSION_MODE, label.PERMISSION_MODE_ACL.Value()).
		Param(code).
		Param(stat).
		Param(dom).
		Writes(token.Token{}).
		Returns(200, "OK", token.Token{}))

	ws.Route(ws.GET("/wechat_work").To(h.WechatWorkOauth2Auth).
		Doc("企业微信登陆").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Auth, label.Disable).
		Metadata(label.PERMISSION_MODE, label.PERMISSION_MODE_ACL.Value()).
		Param(code).
		Param(stat).
		Param(dom).
		Writes(token.Token{}).
		Returns(200, "OK", token.Token{}))
}

func (h *oath2Handler) FeishuOauth2Auth(r *restful.Request, w *restful.Response) {
	req := token.NewFeishuAuthCodeIssueTokenRequest(
		r.QueryParameter("code"),
		r.QueryParameter("state"),
		r.QueryParameter("domain"),
	)

	// 补充用户的登录时的位置信息
	req.Location = token.NewNewLocationFromHttp(r.Request)

	// 颁发Token
	tk, err := h.service.IssueToken(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	tk.SetCookie(w)
	response.Success(w, tk)
}

func (h *oath2Handler) DingDingOauth2Auth(r *restful.Request, w *restful.Response) {
	req := token.NewDingDingAuthCodeIssueTokenRequest(
		r.QueryParameter("authCode"),
		r.QueryParameter("state"),
		r.QueryParameter("domain"),
	)

	// 补充用户的登录时的位置信息
	req.Location = token.NewNewLocationFromHttp(r.Request)

	// 颁发Token
	tk, err := h.service.IssueToken(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	tk.SetCookie(w)
	response.Success(w, tk)
}

func (h *oath2Handler) WechatWorkOauth2Auth(r *restful.Request, w *restful.Response) {
	req := token.NewDingDingAuthCodeIssueTokenRequest(
		r.QueryParameter("authCode"),
		r.QueryParameter("state"),
		r.QueryParameter("domain"),
	)

	// 补充用户的登录时的位置信息
	req.Location = token.NewNewLocationFromHttp(r.Request)

	// 颁发Token
	tk, err := h.service.IssueToken(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	tk.SetCookie(w)
	response.Success(w, tk)
}
