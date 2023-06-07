package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/mcenter/apps/setting"
)

func init() {
	ioc.RegistryApi(&handler{})
}

type handler struct {
	service setting.Service
	log     logger.Logger
	ioc.IocObjectImpl
}

func (h *handler) Init() error {
	h.log = zap.L().Named(setting.AppName)
	h.service = ioc.GetController(setting.AppName).(setting.Service)
	return nil
}

func (h *handler) Name() string {
	return setting.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{"系统设置"}

	ws.Route(ws.POST("/").To(h.UpdateSetting).
		Doc("更新系统设置").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(setting.Setting{}).
		Writes((setting.Setting{})))

	ws.Route(ws.GET("/").To(h.GetSetting).
		Doc("查询系统设置").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata("action", "get").
		Reads(setting.Setting{}).
		Writes((setting.Setting{})).
		Returns(200, "OK", setting.Setting{}))
}
