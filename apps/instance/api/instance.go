package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcenter/apps/instance"
	"github.com/infraboard/mcube/http/restful/response"
)

func (h *handler) RegistryInstance(r *restful.Request, w *restful.Response) {
	req := instance.NewRegistryRequest()

	if err := r.ReadEntity(req); err != nil {
		response.Failed(w, err)
		return
	}

	set, err := h.service.RegistryInstance(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
}
