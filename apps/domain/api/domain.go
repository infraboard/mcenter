package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/restful/response"

	"github.com/infraboard/mcenter/apps/domain"
	"github.com/infraboard/mcenter/apps/token"
)

func (h *handler) CreateDomain(r *restful.Request, w *restful.Response) {
	req := domain.NewCreateDomainRequest()

	if err := r.ReadEntity(req); err != nil {
		response.Failed(w, err)
		return
	}

	set, err := h.service.CreateDomain(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
}

func (h *handler) DescribeDomain(r *restful.Request, w *restful.Response) {
	req := domain.NewDescribeDomainRequestByName(r.PathParameter("name"))
	tk := token.GetTokenFromRequest(r)
	if req.Name != tk.Domain {
		response.Failed(w, exception.NewPermissionDeny("you not in this domain"))
		return
	}

	ins, err := h.service.DescribeDomain(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, ins)
}

func (h *handler) PutDomain(r *restful.Request, w *restful.Response) {
	req := domain.NewPutDomainRequestByName(r.PathParameter("name"))

	if err := r.ReadEntity(req.Spec); err != nil {
		response.Failed(w, err)
		return
	}

	tk := token.GetTokenFromRequest(r)
	if req.Name != tk.Domain {
		response.Failed(w, exception.NewPermissionDeny("you not in this domain"))
		return
	}

	set, err := h.service.UpdateDomain(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *handler) PatchDomain(r *restful.Request, w *restful.Response) {
	req := domain.NewPatchDomainRequestByName(r.PathParameter("name"))

	if err := r.ReadEntity(req.Spec); err != nil {
		response.Failed(w, err)
		return
	}

	tk := token.GetTokenFromRequest(r)
	if req.Name != tk.Domain {
		response.Failed(w, exception.NewPermissionDeny("you not in this domain"))
		return
	}

	set, err := h.service.UpdateDomain(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}
