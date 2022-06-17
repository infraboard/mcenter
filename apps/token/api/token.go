package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/response"

	"github.com/infraboard/mcenter/apps/token"
)

func (h *handler) IssueToken(r *restful.Request, w *restful.Response) {
	req := token.NewIssueTokenRequest()

	if err := r.ReadEntity(req); err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}

	tk, err := h.service.IssueToken(r.Request.Context(), req)
	if err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}

	response.Success(w.ResponseWriter, tk)
}

func (u *handler) RevolkToken(r *restful.Request, w *restful.Response) {
	req := token.NewRevolkTokenRequest()
	if err := r.ReadEntity(req); err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}

	ins, err := h.service.RevolkToken(r.Request.Context(), req)
	if err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}

	response.Success(w.ResponseWriter, ins)
}

func (u *handler) ChangeNamespace(r *restful.Request, w *restful.Response) {
	req := token.NewChangeNamespaceRequest()
	if err := r.ReadEntity(req); err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}

	set, err := h.service.ChangeNamespace(r.Request.Context(), req)
	if err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}
	response.Success(w.ResponseWriter, set)
}
