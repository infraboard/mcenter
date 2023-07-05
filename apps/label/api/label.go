package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcenter/apps/label"
	"github.com/infraboard/mcenter/apps/token"
	http_label "github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/restful/response"
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

func (h *handler) QueryLabel(r *restful.Request, w *restful.Response) {
	req := label.NewQueryLabelRequestFromHTTP(r)
	set, err := h.service.QueryLabel(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}
