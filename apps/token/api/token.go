package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/restful/response"
	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/mcenter/apps/token"
)

func init() {
	ioc.RegistryApi(&tokenHandler{})
}

type tokenHandler struct {
	service token.Service
	log     logger.Logger
	ioc.IocObjectImpl
}

func (h *tokenHandler) Init() error {
	h.log = zap.L().Named(token.AppName)
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
		Reads(token.IssueTokenRequest{}).
		Writes(token.Token{}).
		Returns(200, "OK", token.Token{}))

	ws.Route(ws.DELETE("/").To(h.RevolkToken).
		Doc("撤销令牌").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(token.Token{}).
		Returns(200, "OK", token.Token{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.PATCH("/").To(h.ChangeNamespace).
		Doc("切换空间").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(token.ChangeNamespaceRequest{}))

	ws.Route(ws.GET("/").To(h.ValidateToken).
		Doc("验证令牌").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(token.ValidateTokenRequest{}).
		Writes(token.Token{}).
		Returns(200, "OK", token.Token{}))
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

	set, err := h.service.ChangeNamespace(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *tokenHandler) ValidateToken(r *restful.Request, w *restful.Response) {
	tk := r.Request.Header.Get(token.VALIDATE_TOKEN_HEADER_KEY)
	req := token.NewValidateTokenRequest(tk)

	resp, err := h.service.ValidateToken(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, resp)
}
