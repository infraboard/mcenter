package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/response"

	"github.com/infraboard/mcenter/apps/domain"
)

func (h *handler) CreateDomain(r *restful.Request, w *restful.Response) {
	req := domain.NewCreateDomainRequest()

	if err := r.ReadEntity(req); err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}

	set, err := h.service.CreateDomain(r.Request.Context(), req)
	if err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}

	response.Success(w.ResponseWriter, set)
}

func (u *handler) DescribeDomain(r *restful.Request, w *restful.Response) {
	req := domain.NewDescribeDomainRequestById(r.PathParameter("id"))
	ins, err := h.service.DescribeDomain(r.Request.Context(), req)
	if err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}

	response.Success(w.ResponseWriter, ins)
}

func (u *handler) PutDomain(r *restful.Request, w *restful.Response) {
	req := domain.NewPutDomainRequest(r.PathParameter("id"))

	if err := r.ReadEntity(req.Spec); err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}

	set, err := h.service.UpdateDomain(r.Request.Context(), req)
	if err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}
	response.Success(w.ResponseWriter, set)
}

func (u *handler) PatchDomain(r *restful.Request, w *restful.Response) {
	req := domain.NewPatchDomainRequestById(r.PathParameter("id"))

	if err := r.ReadEntity(req.Spec); err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}

	set, err := h.service.UpdateDomain(r.Request.Context(), req)
	if err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}
	response.Success(w.ResponseWriter, set)
}
