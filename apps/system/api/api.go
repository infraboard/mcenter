package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/infraboard/mcube/v2/http/label"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/gorestful"
	"github.com/infraboard/mcube/v2/ioc/config/log"
	"github.com/rs/zerolog"

	"github.com/infraboard/mcenter/apps/service"
	"github.com/infraboard/mcenter/apps/system"
)

func init() {
	ioc.Api().Registry(&handler{})
}

type handler struct {
	ioc.ObjectImpl

	log      *zerolog.Logger
	resource system.Service
	service  service.MetaService
}

func (h *handler) Init() error {
	h.log = log.Sub(h.Name())
	h.resource = ioc.Controller().Get(system.AppName).(system.Service)
	h.service = ioc.Controller().Get(service.AppName).(service.MetaService)
	h.Registry()
	return nil
}

func (h *handler) Name() string {
	return system.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry() {
	tags := []string{"服务资源列表"}

	ws := gorestful.ObjectRouter(h)
	ws.Route(ws.GET("/service").To(h.QueryService).
		Doc("查询服务列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.List.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Disable).
		Reads(service.QueryServiceRequest{}).
		Writes(service.ServiceSet{}).
		Returns(200, "OK", service.ServiceSet{}))
	ws.Route(ws.GET("/resource").To(h.QueryResource).
		Doc("查询服务资源列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.List.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Disable).
		Reads(system.QueryResourceRequest{}).
		Writes(system.ResourceSet{}).
		Returns(200, "OK", system.ResourceSet{}))
}
