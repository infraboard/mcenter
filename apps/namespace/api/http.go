package api

import (
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/logger"
	"github.com/rs/zerolog"

	"github.com/infraboard/mcenter/apps/namespace"
)

func init() {
	ioc.Api().Registry(&handler{})
}

type handler struct {
	service namespace.Service
	log     *zerolog.Logger
	ioc.ObjectImpl
}

func (h *handler) Init() error {
	h.log = logger.Sub(namespace.AppName)
	h.service = ioc.GetController(namespace.AppName).(namespace.Service)
	return nil
}

func (h *handler) Name() string {
	return namespace.AppName
}

func (h *handler) Version() string {
	return "v1"
}
