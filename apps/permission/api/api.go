package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/http/response"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/mcenter/apps/domain"
	"github.com/infraboard/mcenter/apps/permission"
	"github.com/infraboard/mcenter/apps/role"
)

var (
	h = &handler{}
)

type handler struct {
	service permission.Service
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(domain.AppName)
	h.service = app.GetInternalApp(permission.AppName).(permission.Service)
	return nil
}

func (h *handler) Name() string {
	return permission.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{"用户权限"}

	ws.Route(ws.POST("/").To(h.CheckPermission).
		Doc("权限校验").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(response.NewData(role.Permission{})).
		Returns(200, "OK", response.NewData(role.Permission{})).
		Returns(404, "Not Found", nil))
}

func init() {
	app.RegistryRESTfulApp(h)
}
