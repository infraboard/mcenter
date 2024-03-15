package api

import (
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/log"
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
	h.log = log.Sub(label.AppName)
	h.service = ioc.Controller().Get(label.AppName).(label.Service)
	h.Registry()
	return nil
}

func (h *handler) Name() string {
	return label.AppName
}

func (h *handler) Version() string {
	return "v1"
}
