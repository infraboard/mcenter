package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/http/response"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/mcenter/apps/domain"
)

var (
	h = &handler{}
)

type handler struct {
	service domain.Service
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(domain.AppName)
	h.service = app.GetInternalApp(domain.AppName).(domain.Service)
	return nil
}

func (h *handler) Name() string {
	return domain.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{"domain"}

	ws.Route(ws.POST("/").To(h.CreateDomain).
		Doc("create a domain").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(domain.CreateDomainRequest{}).
		Writes(response.NewData(domain.Domain{})))

	ws.Route(ws.GET("/{id}").To(h.DescribeDomain).
		Doc("get a domain").
		Param(ws.PathParameter("id", "identifier of the domain").DataType("integer").DefaultValue("1")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(response.NewData(domain.Domain{})).
		Returns(200, "OK", response.NewData(domain.Domain{})).
		Returns(404, "Not Found", nil))

	ws.Route(ws.PUT("/{id}").To(h.PutDomain).
		Doc("update a domain").
		Param(ws.PathParameter("id", "identifier of the domain").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(domain.CreateDomainRequest{}))

	ws.Route(ws.PATCH("/{id}").To(h.PatchDomain).
		Doc("patch a domain").
		Param(ws.PathParameter("id", "identifier of the domain").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(domain.CreateDomainRequest{}))
}

func init() {
	app.RegistryRESTfulApp(h)
}
