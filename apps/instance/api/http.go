package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcenter/apps/instance"
	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func init() {
	ioc.RegistryApi(&handler{})
}

type handler struct {
	service instance.Service
	log     logger.Logger
	ioc.IocObjectImpl
}

func (h *handler) Init() error {
	h.log = zap.L().Named(instance.AppName)
	h.service = ioc.GetController(instance.AppName).(instance.Service)
	return nil
}

func (h *handler) Name() string {
	return instance.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{"服务实例管理"}

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
