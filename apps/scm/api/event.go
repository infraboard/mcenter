package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/response"

	"github.com/infraboard/mcenter/apps/scm"
)

func (h *handler) QueryProject(r *restful.Request, w *restful.Response) {
	qs := r.Request.URL.Query()
	req := scm.NewQueryProjectRequest()
	req.Address = qs.Get("address")
	req.Token = qs.Get("token")

	set, err := h.scm.QueryProject(r.Request.Context(), req)
	if err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}

	response.Success(w.ResponseWriter, set)
}
