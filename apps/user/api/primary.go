package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/http/response"
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
	return user.AppName
}

func (h *primary) Version() string {
	return "v1"
}

func (h *primary) Registry(ws *restful.WebService) {
	tags := []string{"user"}

	ws.Route(ws.POST("/").To(h.CreateUser).
		Doc("create a user").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(user.CreateUserRequest{}).
		Writes(response.NewData(user.User{})))

	ws.Route(ws.GET("/").To(h.QueryUser).
		Doc("query a user").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(user.CreateUserRequest{}).
		Writes(response.NewData(user.User{})))

	ws.Route(ws.GET("/{id}").To(h.DescribeUser).
		Doc("get a user").
		Param(ws.PathParameter("id", "identifier of the user").DataType("integer").DefaultValue("1")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(response.NewData(user.User{})).
		Returns(200, "OK", response.NewData(user.User{})).
		Returns(404, "Not Found", nil))

	ws.Route(ws.PUT("/{id}").To(h.PutUser).
		Doc("update a user").
		Param(ws.PathParameter("id", "identifier of the user").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(user.CreateUserRequest{}))

	ws.Route(ws.PATCH("/{id}").To(h.PatchUser).
		Doc("patch a user").
		Param(ws.PathParameter("id", "identifier of the user").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(user.CreateUserRequest{}))

	ws.Route(ws.DELETE("/{id}").To(h.DeleteUser).
		Doc("delete a user").
		Param(ws.PathParameter("id", "identifier of the user").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags))
}

func (h *primary) CreateUser(r *restful.Request, w *restful.Response) {
	req := user.NewCreateUserRequest()

	if err := r.ReadEntity(req); err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}

	set, err := h.service.CreateUser(r.Request.Context(), req)
	if err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}

	response.Success(w.ResponseWriter, set)
}

func (h *primary) PutUser(r *restful.Request, w *restful.Response) {
	req := user.NewPutUserRequest(r.PathParameter("id"))
	if err := r.ReadEntity(req.Profile); err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}

	set, err := h.service.UpdateUser(r.Request.Context(), req)
	if err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}
	response.Success(w.ResponseWriter, set)
}

func (h *primary) PatchUser(r *restful.Request, w *restful.Response) {
	req := user.NewPatchUserRequest(r.PathParameter("id"))
	if err := r.ReadEntity(req.Profile); err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}

	set, err := h.service.UpdateUser(r.Request.Context(), req)
	if err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}
	response.Success(w.ResponseWriter, set)
}

func (h *primary) ResetPassword(r *restful.Request, w *restful.Response) {
	req := user.NewResetPasswordRequest()
	if err := r.ReadEntity(req); err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}
	req.UserId = r.PathParameter("id")

	set, err := h.service.ResetPassword(r.Request.Context(), req)
	if err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}
	response.Success(w.ResponseWriter, set)
}

func (h *primary) DeleteUser(r *restful.Request, w *restful.Response) {
	req := user.NewDeleteUserRequest()
	req.UserIds = append(req.UserIds, r.PathParameter("id"))

	set, err := h.service.DeleteUser(r.Request.Context(), req)
	if err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}
	response.Success(w.ResponseWriter, set)
}

func (h *primary) QueryUser(r *restful.Request, w *restful.Response) {
	req := user.NewQueryUserRequestFromHTTP(r.Request)
	ins, err := h.service.QueryUser(r.Request.Context(), req)
	if err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}

	response.Success(w.ResponseWriter, ins)
}

func (h *primary) DescribeUser(r *restful.Request, w *restful.Response) {
	req := user.NewDescriptUserRequestWithId(r.PathParameter("id"))
	ins, err := h.service.DescribeUser(r.Request.Context(), req)
	if err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}

	response.Success(w.ResponseWriter, ins)
}

func init() {
	app.RegistryRESTfulApp(&primary{})
}
