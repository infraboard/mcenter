package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/http/response"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/mcenter/apps/endpoint"
)

var (
	h = &handler{}
)

type handler struct {
	service endpoint.Service
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(h.Name())
	h.service = app.GetInternalApp(endpoint.AppName).(endpoint.Service)
	return nil
}

func (h *handler) Name() string {
	return endpoint.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{h.Name()}

	// rr := router.ResourceRouter("resource")
	// rr.BasePath("resources")
	// rr.Handle("GET", "/", h.ListResource)

	ws.Route(ws.POST("/").To(h.RegistryEndpoint).
		Doc("registry a service endpoint").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(endpoint.RegistryRequest{}).
		Writes(response.NewData(endpoint.EndpointSet{})))

	ws.Route(ws.GET("/").To(h.QueryEndpoints).
		Doc("get all endpoints").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata("action", "list").
		Reads(endpoint.QueryEndpointRequest{}).
		Writes(response.NewData(endpoint.EndpointSet{})).
		Returns(200, "OK", endpoint.NewEndpointSet()))

	ws.Route(ws.GET("/{id}").To(h.DescribeEndpoint).
		Doc("get a endpoint").
		Param(ws.PathParameter("id", "identifier of the service").DataType("integer").DefaultValue("1")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(response.NewData(endpoint.Endpoint{})).
		Returns(200, "OK", response.NewData(endpoint.Endpoint{})).
		Returns(404, "Not Found", nil))

	// ws.Route(ws.PUT("/{id}").To(h.UpdateService).
	// 	Doc("update a service").
	// 	Param(ws.PathParameter("id", "identifier of the service").DataType("string")).
	// 	Metadata(restfulspec.KeyOpenAPITags, tags).
	// 	Reads(service.CreateServiceRequest{}))

	// ws.Route(ws.PATCH("/{id}").To(h.PatchService).
	// 	Doc("patch a service").
	// 	Param(ws.PathParameter("id", "identifier of the service").DataType("string")).
	// 	Metadata(restfulspec.KeyOpenAPITags, tags).
	// 	Reads(service.CreateServiceRequest{}))

	// ws.Route(ws.DELETE("/{id}").To(h.DeleteService).
	// 	Doc("delete a service").
	// 	Metadata(restfulspec.KeyOpenAPITags, tags).
	// 	Param(ws.PathParameter("id", "identifier of the service").DataType("string")))
}

func init() {
	app.RegistryRESTfulApp(h)
}
