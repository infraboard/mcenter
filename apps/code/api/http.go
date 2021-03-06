package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/http/response"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/mcenter/apps/code"
)

var (
	h = &handler{}
)

type handler struct {
	service code.Service
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(code.AppName)
	h.service = app.GetInternalApp(code.AppName).(code.Service)
	return nil
}

func (h *handler) Name() string {
	return code.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{h.Name()}

	ws.Route(ws.POST("/").To(h.IssueCode).
		Doc("issue verify code").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(code.IssueCodeRequest{}).
		Writes(response.NewData(code.Code{})))
}

func init() {
	app.RegistryRESTfulApp(h)
}
