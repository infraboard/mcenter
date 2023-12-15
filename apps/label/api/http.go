package api

import (
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/logger"
	"github.com/rs/zerolog"

	"github.com/infraboard/mcenter/apps/label"
)

func init() {
	ioc.Api().Registry(&handler{})
}

type handler struct {
	service label.Service
	log     *zerolog.Logger
	ioc.ObjectImpl
}

func (h *handler) Init() error {
	h.log = logger.Sub(label.AppName)
	h.service = ioc.Controller().Get(label.AppName).(label.Service)
	return nil
}

func (h *handler) Name() string {
	return label.AppName
}

func (h *handler) Version() string {
	return "v1"
}
