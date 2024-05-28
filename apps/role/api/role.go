package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcenter/apps/role"
	"github.com/infraboard/mcenter/apps/token"
	"github.com/infraboard/mcube/v2/http/label"
	"github.com/infraboard/mcube/v2/http/restful/response"
	"github.com/infraboard/mcube/v2/ioc/config/gorestful"
)

func (h *handler) Registry() {
	tags := []string{"角色管理"}

	ws := gorestful.ObjectRouter(h)
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
		Writes(role.RoleSet{}).
		Returns(200, "OK", role.RoleSet{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.GET("/{id}").To(h.DescribeRole).
		Doc("查询角色详情").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Get.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Disable).
		Writes(role.Role{}).
		Returns(200, "OK", role.Role{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.DELETE("/{id}").To(h.DeleteRole).
		Doc("删除角色").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Delete.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Param(ws.PathParameter("id", "identifier of the service").DataType("string")))

	ws.Route(ws.POST("/{id}/permissions").To(h.AddPermissionToRole).
		Doc("角色添加权限").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Update.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Writes(role.PermissionSet{}).
		Returns(200, "OK", role.PermissionSet{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.PATCH("/{id}/permissions/{perm_id}").To(h.UpdatePermission).
		Doc("修改角色权限").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Update.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Writes(role.PermissionSet{}).
		Returns(200, "OK", role.PermissionSet{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.DELETE("/{id}/permissions").To(h.RemovePermissionFromRole).
		Doc("角色移除权限").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Update.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Writes(role.PermissionSet{}).
		Returns(200, "OK", role.PermissionSet{}).
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

func (h *handler) DescribeRole(r *restful.Request, w *restful.Response) {
	req := role.NewDescribeRoleRequestWithID(r.PathParameter("id"))
	set, err := h.service.DescribeRole(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
}

func (h *handler) DeleteRole(r *restful.Request, w *restful.Response) {
	req := role.NewDeleteRoleWithID(r.PathParameter("id"))
	req.Scope = token.GetTokenFromRequest(r).GenScope()
	set, err := h.service.DeleteRole(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
}

func (h *handler) AddPermissionToRole(r *restful.Request, w *restful.Response) {
	req := role.NewAddPermissionToRoleRequest()
	if err := r.ReadEntity(&req.Permissions); err != nil {
		response.Failed(w, err)
		return
	}

	req.RoleId = r.PathParameter("id")
	tk := token.GetTokenFromRequest(r)
	req.CreateBy = tk.UserId
	req.Scope = tk.GenScope()
	set, err := h.service.AddPermissionToRole(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
}

func (h *handler) RemovePermissionFromRole(r *restful.Request, w *restful.Response) {
	req := role.NewRemovePermissionFromRoleRequest()
	if err := r.ReadEntity(req); err != nil {
		response.Failed(w, err)
		return
	}

	tk := token.GetTokenFromRequest(r)
	req.Scope = tk.GenScope()
	set, err := h.service.RemovePermissionFromRole(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
}

func (h *handler) UpdatePermission(r *restful.Request, w *restful.Response) {
	req := role.NewUpdatePermissionRequest()
	if err := r.ReadEntity(req); err != nil {
		response.Failed(w, err)
		return
	}

	tk := token.GetTokenFromRequest(r)
	req.Scope = tk.GenScope()
	set, err := h.service.UpdatePermission(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
}
