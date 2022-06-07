package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/http/response"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/mcenter/apps/application"
)

var (
	h = &handler{}
)

type handler struct {
	service application.ServiceServer
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(application.AppName)
	h.service = app.GetGrpcApp(application.AppName).(application.ServiceServer)
	return nil
}

func (h *handler) Name() string {
	return application.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{"applications"}

	ws.Route(ws.POST("/").To(h.CreateApplication).
		Doc("create a application").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(application.CreateApplicationRequest{}).
		Writes(response.NewData(application.Application{})))

	ws.Route(ws.GET("/").To(h.QueryApplication).
		Doc("get all application").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata("action", "list").
		Reads(application.QueryApplicationRequest{}).
		Writes(response.NewData(application.ApplicationSet{})).
		Returns(200, "OK", application.ApplicationSet{}))

	ws.Route(ws.GET("/{id}").To(h.DescribeApplication).
		Doc("get a application").
		Param(ws.PathParameter("id", "identifier of the application").DataType("integer").DefaultValue("1")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(response.NewData(application.Application{})).
		Returns(200, "OK", response.NewData(application.Application{})).
		Returns(404, "Not Found", nil))

	ws.Route(ws.PUT("/{id}").To(h.UpdateApplication).
		Doc("update a application").
		Param(ws.PathParameter("id", "identifier of the application").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(application.CreateApplicationRequest{}))

	ws.Route(ws.PATCH("/{id}").To(h.PatchApplication).
		Doc("patch a application").
		Param(ws.PathParameter("id", "identifier of the application").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(application.CreateApplicationRequest{}))

	ws.Route(ws.DELETE("/{id}").To(h.DeleteApplication).
		Doc("delete a application").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("id", "identifier of the application").DataType("string")))
}

func init() {
	app.RegistryRESTfulApp(h)
}
