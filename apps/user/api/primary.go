package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/http/restful/response"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/mcenter/apps/user"
)

// 主账号用户管理接口

type primary struct {
	service user.Service
	log     logger.Logger
}

func (h *primary) Config() error {
	h.log = zap.L().Named(user.AppName)
	h.service = app.GetInternalApp(user.AppName).(user.Service)
	return nil
}

func (h *primary) Name() string {
	return "user/primary"
}

func (h *primary) Version() string {
	return "v1"
}

func (h *primary) Registry(ws *restful.WebService) {
	tags := []string{"子账号管理"}

	ws.Route(ws.POST("/").To(h.CreateUser).
		Doc("创建子账号").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(user.CreateUserRequest{}).
		Returns(0, "创建成功", &user.User{}))

	ws.Route(ws.GET("/").To(h.QueryUser).
		Doc("查询子账号列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(user.CreateUserRequest{}).
		Writes(user.User{}))

	ws.Route(ws.GET("/{id}").To(h.DescribeUser).
		Doc("查询子账号详情").
		Param(ws.PathParameter("id", "identifier of the user").DataType("integer").DefaultValue("1")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(user.User{}).
		Returns(200, "OK", user.User{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.PUT("/{id}").To(h.PutUser).
		Doc("修改子账号").
		Param(ws.PathParameter("id", "identifier of the user").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(user.CreateUserRequest{}))

	ws.Route(ws.PATCH("/{id}").To(h.PatchUser).
		Doc("修改子账号").
		Param(ws.PathParameter("id", "identifier of the user").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(user.CreateUserRequest{}))

	ws.Route(ws.DELETE("/{id}").To(h.DeleteUser).
		Doc("删除子账号").
		Param(ws.PathParameter("id", "identifier of the user").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags))

	ws.Route(ws.POST("/{id}/password").To(h.ResetPassword).
		Doc("重置子账号密码").
		Param(ws.PathParameter("id", "identifier of the user").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags))
}

func (h *primary) CreateUser(r *restful.Request, w *restful.Response) {
	req := user.NewCreateUserRequest()

	if err := r.ReadEntity(req); err != nil {
		response.Failed(w, err)
		return
	}

	set, err := h.service.CreateUser(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
}

func (h *primary) PutUser(r *restful.Request, w *restful.Response) {
	req := user.NewPutUserRequest(r.PathParameter("id"))
	if err := r.ReadEntity(req.Profile); err != nil {
		response.Failed(w, err)
		return
	}

	set, err := h.service.UpdateUser(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *primary) PatchUser(r *restful.Request, w *restful.Response) {
	req := user.NewPatchUserRequest(r.PathParameter("id"))
	if err := r.ReadEntity(req.Profile); err != nil {
		response.Failed(w, err)
		return
	}

	set, err := h.service.UpdateUser(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *primary) ResetPassword(r *restful.Request, w *restful.Response) {
	req := user.NewResetPasswordRequest()
	if err := r.ReadEntity(req); err != nil {
		response.Failed(w, err)
		return
	}
	req.UserId = r.PathParameter("id")

	set, err := h.service.ResetPassword(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *primary) DeleteUser(r *restful.Request, w *restful.Response) {
	req := user.NewDeleteUserRequest()
	req.UserIds = append(req.UserIds, r.PathParameter("id"))

	set, err := h.service.DeleteUser(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *primary) QueryUser(r *restful.Request, w *restful.Response) {
	req := user.NewQueryUserRequestFromHTTP(r.Request)
	ins, err := h.service.QueryUser(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *primary) DescribeUser(r *restful.Request, w *restful.Response) {
	req := user.NewDescriptUserRequestWithId(r.PathParameter("id"))
	ins, err := h.service.DescribeUser(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func init() {
	app.RegistryRESTfulApp(&primary{})
}
