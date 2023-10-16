package api

import (
	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mcube/ioc/config/logger"
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
	h.log = logger.Sub(role.AppName)
	h.service = ioc.GetController(role.AppName).(role.Service)
	return nil
}

func (h *handler) Name() string {
	return role.AppName
}

func (h *handler) Version() string {
	return "v1"
}
