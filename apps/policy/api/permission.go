package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/http/restful/response"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/mcenter/apps/domain"
	"github.com/infraboard/mcenter/apps/policy"
	"github.com/infraboard/mcenter/apps/role"
)

var (
	h = &handler{}
)

type handler struct {
	service policy.Service
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(domain.AppName)
	h.service = app.GetInternalApp(policy.AppName).(policy.Service)
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

func (h *handler) CheckPermission(r *restful.Request, w *restful.Response) {
	req := policy.NewCheckPermissionRequest()

	if err := r.ReadEntity(req); err != nil {
		response.Failed(w, err)
		return
	}

	perm, err := h.service.CheckPermission(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, perm)
}

func init() {
	app.RegistryRESTfulApp(h)
}
