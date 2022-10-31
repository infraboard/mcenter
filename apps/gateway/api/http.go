package api

import (
	"fmt"

	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/mcenter/apps/gateway"
)

var (
	h = &handler{}
)

type handler struct {
	service gateway.Service
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(gateway.AppName)
	h.service = app.GetInternalApp(gateway.AppName).(gateway.Service)
	return nil
}

func (h *handler) Name() string {
	return gateway.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{"网关管理"}
	fmt.Println(tags)
}

func init() {
	app.RegistryRESTfulApp(h)
}
