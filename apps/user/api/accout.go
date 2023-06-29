package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/restful/response"
	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/mcenter/apps/user"
)

func init() {
	ioc.RegistryApi(&sub{})
}

// 子账号用户管理接口
type sub struct {
	service user.Service
	log     logger.Logger
	ioc.IocObjectImpl
}

func (h *sub) Init() error {
	h.log = zap.L().Named(user.AppName)
	h.service = ioc.GetController(user.AppName).(user.Service)
	return nil
}

func (h *sub) Name() string {
	return "account"
}

func (h *sub) Version() string {
	return "v1"
}

func (h *sub) Registry(ws *restful.WebService) {
	tags := []string{"账号管理"}

	ws.Route(ws.POST("/password").To(h.UpdatePassword).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Allow, user.TYPE_SUB).
		Doc("子账号修改自己密码").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(user.UpdatePasswordRequest{}).
		Returns(0, "OK", &user.User{}))
}

func (h *sub) UpdatePassword(r *restful.Request, w *restful.Response) {
	req := user.NewUpdatePasswordRequest()
	if err := r.ReadEntity(req); err != nil {
		response.Failed(w, err)
		return
	}

	req.UserId = r.PathParameter("id")
	set, err := h.service.UpdatePassword(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
}
