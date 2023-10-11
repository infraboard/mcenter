package api

import (
	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/mcenter/apps/role"
)

func init() {
	ioc.Api().Registry(&handler{})
}

type handler struct {
	service role.Service
	log     logger.Logger
	ioc.ObjectImpl
}

func (h *handler) Init() error {
	h.log = zap.L().Named(role.AppName)
	h.service = ioc.GetController(role.AppName).(role.Service)
	return nil
}

func (h *handler) Name() string {
	return role.AppName
}

func (h *handler) Version() string {
	return "v1"
}
