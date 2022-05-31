package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/response"

	"github.com/infraboard/mcenter/apps/application"
)

func (h *handler) CreateApplication(r *restful.Request, w *restful.Response) {
	req := application.NewCreateApplicationRequest()

	if err := r.ReadEntity(req); err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}

	set, err := h.service.CreateApplication(r.Request.Context(), req)
	if err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}

	response.Success(w.ResponseWriter, set)
}

func (u *handler) QueryApplication(r *restful.Request, w *restful.Response) {
	req := application.NewQueryApplicationRequestFromHTTP(r.Request)
	set, err := h.service.QueryApplication(r.Request.Context(), req)
	if err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}
	response.Success(w.ResponseWriter, set)
}

func (u *handler) DescribeApplication(r *restful.Request, w *restful.Response) {
	req := application.NewDescribeApplicationRequest(r.PathParameter("id"))
	ins, err := h.service.DescribeApplication(r.Request.Context(), req)
	if err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}

	response.Success(w.ResponseWriter, ins)
}

func (u *handler) UpdateApplication(r *restful.Request, w *restful.Response) {
	req := application.NewPutApplicationRequest(r.PathParameter("id"))

	if err := r.ReadEntity(req.Spec); err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}

	set, err := h.service.UpdateApplication(r.Request.Context(), req)
	if err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}
	response.Success(w.ResponseWriter, set)
}

func (u *handler) PatchApplication(r *restful.Request, w *restful.Response) {
	req := application.NewPatchApplicationRequest(r.PathParameter("id"))

	if err := r.ReadEntity(req.Spec); err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}

	set, err := h.service.UpdateApplication(r.Request.Context(), req)
	if err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}
	response.Success(w.ResponseWriter, set)
}

func (u *handler) DeleteApplication(r *restful.Request, w *restful.Response) {
	req := application.NewDeleteApplicationRequestWithID(r.PathParameter("id"))
	set, err := h.service.DeleteApplication(r.Request.Context(), req)
	if err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}
	response.Success(w.ResponseWriter, set)
}
