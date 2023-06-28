package api

import (
	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/mcenter/apps/label"
)

func init() {
	ioc.RegistryApi(&handler{})
}

type handler struct {
	service label.Service
	log     logger.Logger
	ioc.IocObjectImpl
}

func (h *handler) Init() error {
	h.log = zap.L().Named(label.AppName)
	h.service = ioc.GetController(label.AppName).(label.Service)
	return nil
}

func (h *handler) Name() string {
	return label.AppName
}

func (h *handler) Version() string {
	return "v1"
}
