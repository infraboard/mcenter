package api

import (
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/log"
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
	h.log = log.Sub(namespace.AppName)
	h.service = ioc.Controller().Get(namespace.AppName).(namespace.Service)
	h.Registry()
	return nil
}

func (h *handler) Name() string {
	return namespace.AppName
}

func (h *handler) Version() string {
	return "v1"
}
