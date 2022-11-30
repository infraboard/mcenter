package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/mcenter/apps/service"
)

var (
	h = &handler{}
)

type handler struct {
	service service.MetaService
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(service.AppName)
	h.service = app.GetInternalApp(service.AppName).(service.MetaService)
	return nil
}

func (h *handler) Name() string {
	return service.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{"服务管理"}

	ws.Route(ws.POST("/").To(h.CreateService).
		Doc("创建服务").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(service.CreateServiceRequest{}).
		Writes(service.Service{}))

	ws.Route(ws.GET("/").To(h.QueryService).
		Doc("查询服务列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata("action", "list").
		Reads(service.QueryServiceRequest{}).
		Writes(service.ServiceSet{}).
		Returns(200, "OK", service.ServiceSet{}))

	ws.Route(ws.GET("/{id}").To(h.DescribeService).
		Doc("查询服务详情").
		Param(ws.PathParameter("id", "identifier of the service").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(service.Service{}).
		Returns(200, "OK", service.Service{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.PUT("/{id}").To(h.UpdateService).
		Doc("更新服务").
		Param(ws.PathParameter("id", "identifier of the service").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(service.CreateServiceRequest{}))

	ws.Route(ws.PATCH("/{id}").To(h.PatchService).
		Doc("更新服务").
		Param(ws.PathParameter("id", "identifier of the service").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(service.CreateServiceRequest{}))

	ws.Route(ws.DELETE("/{id}").To(h.DeleteService).
		Doc("删除服务").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("id", "identifier of the service").DataType("string")))
}

func init() {
	app.RegistryRESTfulApp(h)
}
