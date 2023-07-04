package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcenter/apps/role"
	"github.com/infraboard/mcenter/apps/token"
	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/response"
)

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{"角色管理"}

	ws.Route(ws.POST("/").To(h.CreateRole).
		Doc("创建角色").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Create.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(role.CreateRoleRequest{}).
		Writes(role.Role{}))

	ws.Route(ws.GET("/").To(h.QueryRole).
		Doc("查询角色列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.List.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Disable).
		Writes(role.Role{}).
		Returns(200, "OK", role.Role{}).
		Returns(404, "Not Found", nil))
}

func (h *handler) CreateRole(r *restful.Request, w *restful.Response) {
	req := role.NewCreateRoleRequest()
	if err := r.ReadEntity(req); err != nil {
		response.Failed(w, err)
		return
	}

	req.UpdateFromToken(token.GetTokenFromRequest(r))
	set, err := h.service.CreateRole(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
}

func (h *handler) QueryRole(r *restful.Request, w *restful.Response) {
	req := role.NewQueryRoleRequestFromHTTP(r)
	set, err := h.service.QueryRole(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
}
