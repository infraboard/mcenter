package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcenter/apps/endpoint"
	"github.com/infraboard/mcube/v2/http/request"
	"github.com/infraboard/mcube/v2/http/restful/response"
)

// RegistryEndpoint 注册Endpint
func (h *handler) RegistryEndpoint(r *restful.Request, w *restful.Response) {
	req := endpoint.NewDefaultRegistryRequest()
	if err := request.GetDataFromRequest(r.Request, req); err != nil {
		response.Failed(w, err)
		return
	}

	_, err := h.service.RegistryEndpoint(
		r.Request.Context(),
		req,
	)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, req)
}

func (h *handler) QueryEndpoints(r *restful.Request, w *restful.Response) {
	req := endpoint.NewQueryEndpointRequestFromHTTP(r.Request)

	set, err := h.service.QueryEndpoints(
		r.Request.Context(),
		req,
	)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
}

func (h *handler) DescribeEndpoint(r *restful.Request, w *restful.Response) {
	req := endpoint.NewDescribeEndpointRequestWithID(r.PathParameter("id"))

	d, err := h.service.DescribeEndpoint(
		r.Request.Context(),
		req,
	)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, d)
}
