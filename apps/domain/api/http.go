package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/mcenter/apps/domain"
)

func init() {
	ioc.RegistryApi(&handler{})
}

type handler struct {
	service domain.Service
	log     logger.Logger
	ioc.IocObjectImpl
}

func (h *handler) Init() error {
	h.log = zap.L().Named(domain.AppName)
	h.service = ioc.GetController(domain.AppName).(domain.Service)
	return nil
}

func (h *handler) Name() string {
	return domain.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{"域管理"}

	ws.Route(ws.POST("/").To(h.CreateDomain).
		Doc("创建域").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(domain.CreateDomainRequest{}).
		Writes(domain.Domain{}))

	ws.Route(ws.GET("/{id}").To(h.DescribeDomain).
		Doc("查询域").
		Param(ws.PathParameter("id", "identifier of the domain").DataType("integer").DefaultValue("1")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(domain.Domain{}).
		Returns(200, "OK", domain.Domain{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.PUT("/{id}").To(h.PutDomain).
		Doc("更新域").
		Param(ws.PathParameter("id", "identifier of the domain").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(domain.CreateDomainRequest{}))

	ws.Route(ws.PATCH("/{id}").To(h.PatchDomain).
		Doc("更新域").
		Param(ws.PathParameter("id", "identifier of the domain").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(domain.CreateDomainRequest{}))
}
