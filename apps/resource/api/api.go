package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/infraboard/mcube/v2/http/label"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/gorestful"
	"github.com/infraboard/mcube/v2/ioc/config/log"
	"github.com/rs/zerolog"

	"github.com/infraboard/mcenter/apps/resource"
)

func init() {
	ioc.Api().Registry(&handler{})
}

type handler struct {
	service resource.Service
	log     *zerolog.Logger
	ioc.ObjectImpl
}

func (h *handler) Init() error {
	h.log = log.Sub(h.Name())
	h.service = ioc.Controller().Get(resource.AppName).(resource.Service)
	h.Registry()
	return nil
}

func (h *handler) Name() string {
	return resource.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry() {
	tags := []string{"服务功能"}

	ws := gorestful.ObjectRouter(h)
	ws.Route(ws.GET("/").To(h.QueryResources).
		Doc("查询服务资源列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.List.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Disable).
		Reads(resource.QueryResourceRequest{}).
		Writes(resource.ResourceSet{}).
		Returns(200, "OK", resource.ResourceSet{}))
}
