package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcenter/apps/service"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/http/restful/response"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

var (
	ph = &providerHandler{}
)

type providerHandler struct {
	log logger.Logger

	service service.MetaService
}

func (h *providerHandler) Config() error {
	h.log = zap.L().Named(service.AppName)
	h.service = app.GetInternalApp(service.AppName).(service.MetaService)
	return nil
}

func (h *providerHandler) Name() string {
	return "providers"
}

func (h *providerHandler) Version() string {
	return "v1"
}

func (h *providerHandler) Registry(ws *restful.WebService) {
	tags := []string{"服务提供商"}

	ws.Route(ws.GET("/gitlab/projects").To(h.QueryGitlabProject).
		Doc("查询Gitlab项目列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata("action", "list").
		Reads(service.QueryServiceRequest{}).
		Writes(service.ServiceSet{}).
		Returns(200, "OK", service.ServiceSet{}))
}

func (h *providerHandler) QueryGitlabProject(r *restful.Request, w *restful.Response) {
	req := service.NewQueryGitlabProjectRequestFromHTTP(r.Request)
	set, err := h.service.QueryGitlabProject(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func init() {
	app.RegistryRESTfulApp(ph)
}
