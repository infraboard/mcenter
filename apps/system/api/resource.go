package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcenter/apps/domain"
	"github.com/infraboard/mcenter/apps/namespace"
	"github.com/infraboard/mcenter/apps/service"
	"github.com/infraboard/mcenter/apps/system"
	"github.com/infraboard/mcube/v2/http/request"
	"github.com/infraboard/mcube/v2/http/restful/response"
)

func (h *handler) QueryResource(r *restful.Request, w *restful.Response) {
	req := system.NewQueryResourceRequestFromHTTP(r.Request)

	set, err := h.resource.QueryResource(
		r.Request.Context(),
		req,
	)
	if err != nil {
		response.Failed(w, err)
		return
	}

	if len(set.Items) == 0 {
		set.Items = []*system.Resource{}
	}

	response.Success(w, set)
}

func (h *handler) QueryService(r *restful.Request, w *restful.Response) {
	req := service.NewQueryServiceRequest()
	req.Page = request.NewPageRequestFromHTTP(r.Request)
	req.Scope.Domain = domain.DEFAULT_DOMAIN
	req.Scope.Namespace = namespace.SYSTEM_NAMESPACE
	set, err := h.service.QueryService(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
}
