package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/mcenter/apps/domain"
	"github.com/infraboard/mcenter/apps/policy"
	"github.com/infraboard/mcenter/apps/role"
)

func init() {
	ioc.RegistryApi(&handler{})
}

type handler struct {
	service policy.Service
	log     logger.Logger
	ioc.IocObjectImpl
}

func (h *handler) Init() error {
	h.log = zap.L().Named(domain.AppName)
	h.service = ioc.GetController(policy.AppName).(policy.Service)
	return nil
}

func (h *handler) Name() string {
	return "permission"
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{"用户权限"}

	ws.Route(ws.POST("/").To(h.CheckPermission).
		Doc("权限校验").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(role.Permission{}).
		Returns(200, "OK", role.Permission{}).
		Returns(404, "Not Found", nil))
}
