package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/infraboard/mcube/v2/http/label"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/gorestful"
	"github.com/infraboard/mcube/v2/ioc/config/log"
	"github.com/rs/zerolog"

	"github.com/infraboard/mcenter/apps/endpoint"
)

func init() {
	ioc.Api().Registry(&handler{})
}

type handler struct {
	service endpoint.Service
	log     *zerolog.Logger
	ioc.ObjectImpl
}

func (h *handler) Init() error {
	h.log = log.Sub(h.Name())
	h.service = ioc.Controller().Get(endpoint.AppName).(endpoint.Service)
	h.Registry()
	return nil
}

func (h *handler) Name() string {
	return endpoint.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry() {
	tags := []string{"服务功能"}

	ws := gorestful.ObjectRouter(h)
	ws.Route(ws.POST("/").To(h.RegistryEndpoint).
		Doc("注册服务功能列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Auth, label.Disable).
		Metadata(label.PERMISSION_MODE, label.PERMISSION_MODE_ACL.Value()).
		Reads(endpoint.RegistryRequest{}).
		Writes(endpoint.EndpointSet{}))

	ws.Route(ws.GET("/").To(h.QueryEndpoints).
		Doc("查询服务功能列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Auth, label.Enable).
		Metadata(label.PERMISSION_MODE, label.PERMISSION_MODE_ACL.Value()).
		Metadata(label.Allow, label.AllowAll()).
		Reads(endpoint.QueryEndpointRequest{}).
		Writes(endpoint.EndpointSet{}).
		Returns(200, "OK", endpoint.NewEndpointSet()))

	ws.Route(ws.GET("/{id}").To(h.DescribeEndpoint).
		Doc("查询服务功能详情").
		Param(ws.PathParameter("id", "identifier of the service").DataType("integer").DefaultValue("1")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Auth, label.Enable).
		Metadata(label.PERMISSION_MODE, label.PERMISSION_MODE_ACL.Value()).
		Metadata(label.Allow, label.AllowAll()).
		Writes(endpoint.Endpoint{}).
		Returns(200, "OK", endpoint.Endpoint{}).
		Returns(404, "Not Found", nil))
}
