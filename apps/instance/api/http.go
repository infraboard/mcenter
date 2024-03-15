package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/infraboard/mcenter/apps/instance"
	"github.com/infraboard/mcube/v2/http/label"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/gorestful"
	"github.com/infraboard/mcube/v2/ioc/config/log"
	"github.com/rs/zerolog"
)

func init() {
	ioc.Api().Registry(&handler{})
}

type handler struct {
	service instance.Service
	log     *zerolog.Logger
	ioc.ObjectImpl
}

func (h *handler) Init() error {
	h.log = log.Sub(instance.AppName)
	h.service = ioc.Controller().Get(instance.AppName).(instance.Service)
	h.Registry()
	return nil
}

func (h *handler) Name() string {
	return instance.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry() {
	tags := []string{"服务实例管理"}

	ws := gorestful.ObjectRouter(h)
	ws.Route(ws.GET("/").To(h.SearchInstance).
		Doc("搜索实例").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Auth, label.Disable).
		Metadata(label.PERMISSION_MODE, label.PERMISSION_MODE_ACL.Value()).
		Reads(instance.RegistryRequest{}).
		Writes(instance.Instance{}))

	ws.Route(ws.POST("/").To(h.RegistryInstance).
		Doc("实例注册").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Auth, label.Disable).
		Metadata(label.PERMISSION_MODE, label.PERMISSION_MODE_ACL.Value()).
		Reads(instance.UnregistryRequest{}).
		Writes(instance.Instance{}))

	ws.Route(ws.DELETE("/{instance_id}").To(h.UnRegistryInstance).
		Metadata(label.Auth, label.Disable).
		Metadata(label.PERMISSION_MODE, label.PERMISSION_MODE_ACL.Value()).
		Doc("实例注销").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(instance.UnregistryRequest{}).
		Writes(instance.Instance{}))
}
