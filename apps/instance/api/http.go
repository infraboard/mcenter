package api

import (
	"fmt"

	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcenter/apps/instance"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

var (
	h = &handler{}
)

type handler struct {
	service instance.Service
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(instance.AppName)
	h.service = app.GetInternalApp(instance.AppName).(instance.Service)
	return nil
}

func (h *handler) Name() string {
	return instance.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{"instances"}
	fmt.Println(tags)
}

func init() {
	app.RegistryRESTfulApp(h)
}
