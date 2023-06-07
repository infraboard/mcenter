package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcenter/apps/resource"
	"github.com/infraboard/mcube/http/restful/response"
)

func (h *handler) QueryResources(r *restful.Request, w *restful.Response) {
	req := resource.NewQueryResourceRequestFromHTTP(r.Request)

	set, err := h.service.QueryResources(
		r.Request.Context(),
		req,
	)
	if err != nil {
		response.Failed(w, err)
		return
	}

	if len(set.Items) == 0 {
		set.Items = []*resource.Resource{}
	}

	response.Success(w, set)
}
