package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/http/response"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/mcenter/apps/setting"
)

var (
	h = &handler{}
)

type handler struct {
	service setting.Service
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(setting.AppName)
	h.service = app.GetInternalApp(setting.AppName).(setting.Service)
	return nil
}

func (h *handler) Name() string {
	return setting.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{h.Name()}

	ws.Route(ws.POST("/").To(h.UpdateSetting).
		Doc("update a setting").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(setting.Setting{}).
		Writes(response.NewData(setting.Setting{})))

	ws.Route(ws.GET("/").To(h.GetSetting).
		Doc("get all setting").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata("action", "get").
		Reads(setting.Setting{}).
		Writes(response.NewData(setting.Setting{})).
		Returns(200, "OK", setting.Setting{}))
}

func init() {
	app.RegistryRESTfulApp(h)
}
