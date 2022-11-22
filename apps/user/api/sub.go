package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/http/response"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/mcenter/apps/user"
)

// 主账号用户管理接口

type sub struct {
	service user.Service
	log     logger.Logger
}

func (h *sub) Config() error {
	h.log = zap.L().Named(user.AppName)
	h.service = app.GetInternalApp(user.AppName).(user.Service)
	return nil
}

func (h *sub) Name() string {
	return "user/sub"
}

func (h *sub) Version() string {
	return "v1"
}

func (h *sub) Registry(ws *restful.WebService) {
	tags := []string{"账号管理"}

	ws.Route(ws.POST("/password").To(h.UpdatePassword).
		Doc("子账号修改自己密码").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(user.UpdatePasswordRequest{}).
		Returns(0, "OK", &user.User{}))
}

func (h *sub) UpdatePassword(r *restful.Request, w *restful.Response) {
	req := user.NewUpdatePasswordRequest()
	if err := r.ReadEntity(req); err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}

	req.UserId = r.PathParameter("id")
	set, err := h.service.UpdatePassword(r.Request.Context(), req)
	if err != nil {
		response.Failed(w.ResponseWriter, err)
		return
	}

	response.Success(w.ResponseWriter, set)
}

func init() {
	app.RegistryRESTfulApp(&sub{})
}
