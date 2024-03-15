package api

import (
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/log"
	"github.com/rs/zerolog"

	"github.com/infraboard/mcenter/apps/role"
)

func init() {
	ioc.Api().Registry(&handler{})
}

type handler struct {
	service role.Service
	log     *zerolog.Logger
	ioc.ObjectImpl
}

func (h *handler) Init() error {
	h.log = log.Sub(role.AppName)
	h.service = ioc.Controller().Get(role.AppName).(role.Service)
	h.Registry()
	return nil
}

func (h *handler) Name() string {
	return role.AppName
}

func (h *handler) Version() string {
	return "v1"
}
