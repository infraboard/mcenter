package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/v2/http/label"
	"github.com/infraboard/mcube/v2/http/restful/response"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/log"
	"github.com/rs/zerolog"

	"github.com/infraboard/mcenter/apps/policy"
	"github.com/infraboard/mcenter/apps/role"
)

func init() {
	ioc.Api().Registry(&handler{})
}

type handler struct {
	service policy.Service
	log     *zerolog.Logger
	ioc.ObjectImpl
}

func (h *handler) Init() error {
	h.log = log.Sub(policy.AppName)
	h.service = ioc.Controller().Get(policy.AppName).(policy.Service)
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
		Metadata(label.Auth, label.Enable).
		Metadata(label.PERMISSION_MODE, label.PERMISSION_MODE_ACL.Value()).
		Metadata(label.Allow, label.AllowAll()).
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
