package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcenter/apps/label"
	"github.com/infraboard/mcenter/apps/token"
	http_label "github.com/infraboard/mcube/v2/http/label"
	"github.com/infraboard/mcube/v2/http/restful/response"
	"github.com/infraboard/mcube/v2/pb/request"
)

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{"标签管理"}

	ws.Route(ws.POST("/").To(h.CreateLabel).
		Doc("创建标签").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(http_label.Resource, h.Name()).
		Metadata(http_label.Action, http_label.Create.Value()).
		Metadata(http_label.Auth, http_label.Enable).
		Metadata(http_label.Permission, http_label.Enable).
		Reads(label.CreateLabelRequest{}).
		Writes(label.Label{}))

	ws.Route(ws.GET("/").To(h.QueryLabel).
		Doc("查询标签").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(http_label.Resource, h.Name()).
		Metadata(http_label.Action, http_label.List.Value()).
		Metadata(http_label.Auth, http_label.Enable).
		Metadata(http_label.Permission, http_label.Disable).
		Reads(label.QueryLabelRequest{}).
		Writes(label.LabelSet{}).
		Returns(200, "OK", label.LabelSet{}))

	ws.Route(ws.GET("/{id}").To(h.DescribeLabel).
		Doc("标签详情").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(http_label.Resource, h.Name()).
		Metadata(http_label.Action, http_label.Delete.Value()).
		Metadata(http_label.Auth, http_label.Enable).
		Metadata(http_label.Permission, http_label.Disable).
		Reads(label.DescribeLabelRequest{}).
		Writes(label.Label{}).
		Returns(200, "OK", label.Label{}))

	ws.Route(ws.PUT("/{id}").To(h.UpdateLabel(request.UpdateMode_PUT)).
		Doc("更新标签").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(http_label.Resource, h.Name()).
		Metadata(http_label.Action, http_label.Update.Value()).
		Metadata(http_label.Auth, http_label.Enable).
		Metadata(http_label.Permission, http_label.Disable).
		Reads(label.DescribeLabelRequest{}).
		Writes(label.Label{}).
		Returns(200, "OK", label.Label{}))

	ws.Route(ws.PATCH("/{id}").To(h.UpdateLabel(request.UpdateMode_PATCH)).
		Doc("更新标签").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(http_label.Resource, h.Name()).
		Metadata(http_label.Action, http_label.Update.Value()).
		Metadata(http_label.Auth, http_label.Enable).
		Metadata(http_label.Permission, http_label.Disable).
		Reads(label.DescribeLabelRequest{}).
		Writes(label.Label{}).
		Returns(200, "OK", label.Label{}))

	ws.Route(ws.DELETE("/{id}").To(h.DeleteLabel).
		Doc("删除标签").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(http_label.Resource, h.Name()).
		Metadata(http_label.Action, http_label.Delete.Value()).
		Metadata(http_label.Auth, http_label.Enable).
		Metadata(http_label.Permission, http_label.Enable).
		Reads(label.DeleteLabelRequest{}).
		Writes(label.Label{}).
		Returns(200, "OK", label.Label{}))
}

func (h *handler) CreateLabel(r *restful.Request, w *restful.Response) {
	req := label.NewCreateLabelRequest()
	if err := r.ReadEntity(req); err != nil {
		response.Failed(w, err)
		return
	}

	// 补充用户信息
	tk := token.GetTokenFromRequest(r)
	req.CreateBy = tk.UserId
	req.Domain = tk.Domain
	req.Namespace = tk.Namespace

	set, err := h.service.CreateLabel(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
}

func (h *handler) UpdateLabel(m request.UpdateMode) restful.RouteFunction {
	return func(r *restful.Request, w *restful.Response) {
		req := label.NewUpdateLabelRequest()
		req.Id = r.PathParameter("id")
		req.UpdateMode = m
		if err := r.ReadEntity(req.Spec); err != nil {
			response.Failed(w, err)
			return
		}

		// 补充用户信息
		tk := token.GetTokenFromRequest(r)
		req.UpdateBy = tk.UserId

		set, err := h.service.UpdateLabel(r.Request.Context(), req)
		if err != nil {
			response.Failed(w, err)
			return
		}

		response.Success(w, set)
	}
}

func (h *handler) QueryLabel(r *restful.Request, w *restful.Response) {
	req := label.NewQueryLabelRequestFromHTTP(r)
	set, err := h.service.QueryLabel(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *handler) DescribeLabel(r *restful.Request, w *restful.Response) {
	req := label.NewDescribeLabelRequest(r.PathParameter("id"))
	set, err := h.service.DescribeLabel(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *handler) DeleteLabel(r *restful.Request, w *restful.Response) {
	req := label.NewDeleteLabelRequest(r.PathParameter("id"))
	set, err := h.service.DeleteLabel(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}
