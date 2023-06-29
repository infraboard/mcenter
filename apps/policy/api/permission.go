package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcenter/apps/policy"
	"github.com/infraboard/mcube/http/restful/response"
)

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
