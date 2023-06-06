package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/ioc"
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
	ioc.IocObjectImpl
}

func (h *handler) Init() error {
	h.log = zap.L().Named(h.Name())
	h.service = ioc.GetController(resource.AppName).(resource.Service)
	return nil
}

func (h *handler) Name() string {
	return resource.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{"服务功能"}

	ws.Route(ws.GET("/").To(h.QueryResources).
		Doc("查询服务资源列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata("action", "list").
		Reads(resource.QueryResourceRequest{}).
		Writes(resource.ResourceSet{}).
		Returns(200, "OK", resource.ResourceSet{}))
}

func init() {
	ioc.RegistryApi(h)
}
