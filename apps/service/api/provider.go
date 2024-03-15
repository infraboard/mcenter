package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcenter/apps/service"
	"github.com/infraboard/mcube/v2/http/label"
	"github.com/infraboard/mcube/v2/http/restful/response"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/gorestful"
	"github.com/infraboard/mcube/v2/ioc/config/log"
	"github.com/rs/zerolog"
)

func init() {
	ioc.Api().Registry(&providerHandler{})
}

type providerHandler struct {
	log *zerolog.Logger

	service service.MetaService
	ioc.ObjectImpl
}

func (h *providerHandler) Init() error {
	h.log = log.Sub(service.AppName)
	h.service = ioc.Controller().Get(service.AppName).(service.MetaService)
	h.Registry()
	return nil
}

func (h *providerHandler) Name() string {
	return "providers"
}

func (h *providerHandler) Version() string {
	return "v1"
}

func (h *providerHandler) Registry() {
	tags := []string{"服务提供商"}

	ws := gorestful.ObjectRouter(h)
	ws.Route(ws.GET("/gitlab/projects").To(h.QueryGitlabProject).
		Doc("查询Gitlab项目列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.List.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Disable).
		Reads(service.QueryServiceRequest{}).
		Writes(service.ServiceSet{}).
		Returns(200, "OK", service.ServiceSet{}))
}

func (h *providerHandler) QueryGitlabProject(r *restful.Request, w *restful.Response) {
	req := service.NewQueryGitlabProjectRequestFromHTTP(r)
	set, err := h.service.QueryGitlabProject(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
}
