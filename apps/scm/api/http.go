package api

import (
	"fmt"

	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/mcenter/apps/scm"
)

var (
	h = &handler{}
)

type handler struct {
	log logger.Logger
	scm scm.Service
}

func (h *handler) Config() error {
	h.scm = app.GetInternalApp(scm.AppName).(scm.Service)
	h.log = zap.L().Named(scm.AppName)
	return nil
}

func (h *handler) Name() string {
	return scm.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{"scm"}
	fmt.Println(tags)

}

func init() {
	app.RegistryRESTfulApp(h)
}
