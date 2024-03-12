package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/infraboard/mcube/v2/http/label"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/gorestful"
	"github.com/infraboard/mcube/v2/ioc/config/log"
	"github.com/rs/zerolog"

	"github.com/infraboard/mcenter/apps/domain"
	"github.com/infraboard/mcenter/apps/user"
)

func init() {
	ioc.Api().Registry(&handler{})
}

type handler struct {
	service domain.Service
	log     *zerolog.Logger
	ioc.ObjectImpl
}

func (h *handler) Init() error {
	h.log = log.Sub(domain.AppName)
	h.service = ioc.Controller().Get(domain.AppName).(domain.Service)
	h.Registry()
	return nil
}

func (h *handler) Name() string {
	return domain.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry() {
	tags := []string{"域管理"}

	ws := gorestful.NewGoRestfulApiService(h)
	ws.Route(ws.POST("/").To(h.CreateDomain).
		Doc("创建域").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Auth, label.Enable).
		Metadata(label.PERMISSION_MODE, label.PERMISSION_MODE_ACL.Value()).
		Metadata(label.Allow, user.TypeToString(user.TYPE_SUPPER)).
		Reads(domain.CreateDomainRequest{}).
		Writes(domain.Domain{}))

	ws.Route(ws.GET("/{name}").To(h.DescribeDomain).
		Doc("查询域").
		Param(ws.PathParameter("id", "identifier of the domain").DataType("integer").DefaultValue("1")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Auth, label.Enable).
		Metadata(label.PERMISSION_MODE, label.PERMISSION_MODE_ACL.Value()).
		Metadata(label.Allow, user.TypeToString(user.TYPE_SUPPER, user.TYPE_PRIMARY)).
		Writes(domain.Domain{}).
		Returns(200, "OK", domain.Domain{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.PUT("/{id}").To(h.PutDomain).
		Doc("更新域").
		Param(ws.PathParameter("id", "identifier of the domain").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Auth, label.Enable).
		Metadata(label.PERMISSION_MODE, label.PERMISSION_MODE_ACL.Value()).
		Metadata(label.Allow, user.TypeToString(user.TYPE_SUPPER, user.TYPE_PRIMARY)).
		Reads(domain.CreateDomainRequest{}))

	ws.Route(ws.PATCH("/{id}").To(h.PatchDomain).
		Doc("更新域").
		Param(ws.PathParameter("id", "identifier of the domain").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Auth, label.Enable).
		Metadata(label.PERMISSION_MODE, label.PERMISSION_MODE_ACL.Value()).
		Metadata(label.Allow, user.TypeToString(user.TYPE_SUPPER, user.TYPE_PRIMARY)).
		Reads(domain.CreateDomainRequest{}))
}
