package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/restful/response"

	"github.com/infraboard/mcenter/apps/code"
)

func (h *handler) IssueCode(r *restful.Request, w *restful.Response) {
	req := code.NewIssueCodeRequest()

	if err := r.ReadEntity(req); err != nil {
		response.Failed(w, err)
		return
	}

	cd, err := h.service.IssueCode(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, cd)
}
