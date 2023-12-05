package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcenter/apps/namespace"
	"github.com/infraboard/mcenter/apps/token"
	"github.com/infraboard/mcenter/apps/user"
	"github.com/infraboard/mcube/v2/http/label"
	"github.com/infraboard/mcube/v2/http/restful/response"
)

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{"空间管理"}

	ws.Route(ws.POST("/").To(h.CreateNamespace).
		Doc("创建空间").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Auth, label.Enable).
		Metadata(label.PERMISSION_MODE, label.PERMISSION_MODE_ACL.Value()).
		Metadata(label.Allow, user.TypeToString(user.TYPE_PRIMARY)).
		Reads(namespace.CreateNamespaceRequest{}).
		Writes(namespace.Namespace{}))

	ws.Route(ws.GET("/").To(h.QueryNamespace).
		Doc("查询空间列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.List.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Writes(namespace.NamespaceSet{}).
		Returns(200, "OK", namespace.NamespaceSet{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.GET("/{id}").To(h.DescribeNamespace).
		Doc("查询空间详情").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.List.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Disable).
		Writes(namespace.Namespace{}).
		Returns(200, "OK", namespace.Namespace{}).
		Returns(404, "Not Found", nil))
}

func (h *handler) CreateNamespace(r *restful.Request, w *restful.Response) {
	req := namespace.NewCreateNamespaceRequest()
	if err := r.ReadEntity(req); err != nil {
		response.Failed(w, err)
		return
	}
	req.UpdateOwner(token.GetTokenFromRequest(r))

	set, err := h.service.CreateNamespace(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
}

func (h *handler) QueryNamespace(r *restful.Request, w *restful.Response) {
	req := namespace.NewQueryNamespaceRequestFromHTTP(r)
	set, err := h.service.QueryNamespace(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
}

func (h *handler) DescribeNamespace(r *restful.Request, w *restful.Response) {
	req := namespace.NewDescriptNamespaceRequestById(r.PathParameter("id"))
	set, err := h.service.DescribeNamespace(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
}
