package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/restful/response"
	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mcube/ioc/config/logger"
	"github.com/rs/zerolog"

	"github.com/infraboard/mcenter/apps/token"
	"github.com/infraboard/mcenter/apps/user"
)

func init() {
	ioc.Api().Registry(&tokenHandler{})
}

type tokenHandler struct {
	service token.Service
	log     *zerolog.Logger
	ioc.ObjectImpl
}

func (h *tokenHandler) Init() error {
	h.log = logger.Sub(token.AppName)
	h.service = ioc.GetController(token.AppName).(token.Service)
	return nil
}

func (h *tokenHandler) Name() string {
	return token.AppName
}

func (h *tokenHandler) Version() string {
	return "v1"
}

func (h *tokenHandler) Registry(ws *restful.WebService) {
	tags := []string{"登录"}

	ws.Route(ws.POST("/").To(h.IssueToken).
		Doc("颁发令牌").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Auth, label.Disable).
		Metadata(label.PERMISSION_MODE, label.PERMISSION_MODE_ACL.Value()).
		Reads(token.IssueTokenRequest{}).
		Writes(token.Token{}).
		Returns(200, "OK", token.Token{}))

	ws.Route(ws.DELETE("/").To(h.RevolkToken).
		Doc("撤销令牌").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Auth, label.Enable).
		Metadata(label.PERMISSION_MODE, label.PERMISSION_MODE_ACL.Value()).
		Metadata(label.Allow, label.AllowAll()).
		Writes(token.Token{}).
		Returns(200, "OK", token.Token{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.PATCH("/").To(h.ChangeNamespace).
		Doc("切换空间").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Auth, label.Disable).
		Metadata(label.PERMISSION_MODE, label.PERMISSION_MODE_ACL.Value()).
		Reads(token.ChangeNamespaceRequest{}))

	ws.Route(ws.POST("/validate").To(h.ValidateToken).
		Doc("验证令牌").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Auth, label.Disable).
		Metadata(label.PERMISSION_MODE, label.PERMISSION_MODE_ACL.Value()).
		Reads(token.ValidateTokenRequest{}).
		Writes(token.Token{}).
		Returns(200, "OK", token.Token{}))

	// 只有主账号才能查询
	ws.Route(ws.GET("/").To(h.QueryToken).
		Doc("令牌颁发记录").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Auth, label.Enable).
		Metadata(label.PERMISSION_MODE, label.PERMISSION_MODE_ACL.Value()).
		Metadata(label.Allow, user.TypeToString(user.TYPE_PRIMARY)).
		Reads(token.QueryTokenRequest{}).
		Writes(token.TokenSet{}).
		Returns(200, "OK", token.TokenSet{}))
}

func (h *tokenHandler) IssueToken(r *restful.Request, w *restful.Response) {
	req := token.NewIssueTokenRequest()
	if err := r.ReadEntity(req); err != nil {
		response.Failed(w, err)
		return
	}

	// 补充用户的登录时的位置信息
	req.Location = token.NewNewLocationFromHttp(r.Request)
	tk, err := h.service.IssueToken(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	tk.SetCookie(w)
	response.Success(w, tk)
}

func (h *tokenHandler) RevolkToken(r *restful.Request, w *restful.Response) {
	qs := r.Request.URL.Query()
	req := token.NewRevolkTokenRequest("", "")
	req.AccessToken = token.GetAccessTokenFromHTTP(r.Request)
	req.RefreshToken = qs.Get("refresh_token")

	ins, err := h.service.RevolkToken(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *tokenHandler) ChangeNamespace(r *restful.Request, w *restful.Response) {
	req := token.NewChangeNamespaceRequest()
	if err := r.ReadEntity(req); err != nil {
		response.Failed(w, err)
		return
	}

	tk := token.GetAccessTokenFromHTTP(r.Request)
	req.Token = tk

	set, err := h.service.ChangeNamespace(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *tokenHandler) ValidateToken(r *restful.Request, w *restful.Response) {
	tk := token.GetAccessTokenFromHTTP(r.Request)
	req := token.NewValidateTokenRequest(tk)

	resp, err := h.service.ValidateToken(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, resp)
}

func (h *tokenHandler) QueryToken(r *restful.Request, w *restful.Response) {
	req := token.NewQueryTokenRequestFromHttp(r)

	resp, err := h.service.QueryToken(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, resp)
}
