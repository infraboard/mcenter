package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/http/response"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/mcenter/apps/resource"
)

var (
	h = &handler{}
)

type handler struct {
	service resource.Service
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(h.Name())
	h.service = app.GetInternalApp(resource.AppName).(resource.Service)
	return nil
}

func (h *handler) Name() string {
	return resource.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{h.Name()}

	ws.Route(ws.GET("/").To(h.QueryResources).
		Doc("get all resource").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata("action", "list").
		Reads(resource.QueryResourceRequest{}).
		Writes(response.NewData(resource.ResourceSet{})).
		Returns(200, "OK", response.NewData(resource.ResourceSet{})))
}

func init() {
	app.RegistryRESTfulApp(h)
}
