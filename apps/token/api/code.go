package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/v2/http/label"
	"github.com/infraboard/mcube/v2/http/restful/response"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/logger"
	"github.com/rs/zerolog"

	"github.com/infraboard/mcenter/apps/token"
)

func init() {
	ioc.Api().Registry(&codeHandler{})
}

type codeHandler struct {
	service token.Service
	log     *zerolog.Logger
	ioc.ObjectImpl
}

func (h *codeHandler) Init() error {
	h.log = logger.Sub(token.AppName)
	h.service = ioc.Controller().Get(token.AppName).(token.Service)
	return nil
}

func (h *codeHandler) Name() string {
	return "code"
}

func (h *codeHandler) Version() string {
	return "v1"
}

func (h *codeHandler) Registry(ws *restful.WebService) {
	tags := []string{"验证码管理"}

	ws.Route(ws.POST("/").To(h.IssueCode).
		Doc("颁发验证码").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Auth, label.Enable).
		Metadata(label.PERMISSION_MODE, label.PERMISSION_MODE_ACL.Value()).
		Reads(token.IssueCodeRequest{}).
		Writes(token.Code{}))
}

func (h *codeHandler) IssueCode(r *restful.Request, w *restful.Response) {
	req := token.NewIssueCodeRequest()

	if err := r.ReadEntity(req); err != nil {
		response.Failed(w, err)
		return
	}

	cd, err := h.service.IssueCode(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, cd)
}
