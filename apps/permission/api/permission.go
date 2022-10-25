package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcenter/apps/permission"
	"github.com/infraboard/mcube/http/response"
)

func (h *handler) CheckPermission(r *restful.Request, w *restful.Response) {
	req := permission.NewCheckPermissionRequest()

	if err := r.ReadEntity(req); err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}

	perm, err := h.service.CheckPermission(r.Request.Context(), req)
	if err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}

	response.Success(w.ResponseWriter, perm)
}
