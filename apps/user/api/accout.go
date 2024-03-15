package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/v2/http/label"
	"github.com/infraboard/mcube/v2/http/restful/response"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/gorestful"
	"github.com/infraboard/mcube/v2/ioc/config/log"
	"github.com/rs/zerolog"

	"github.com/infraboard/mcenter/apps/user"
)

func init() {
	ioc.Api().Registry(&sub{})
}

// 子账号用户管理接口
type sub struct {
	service user.Service
	log     *zerolog.Logger
	ioc.ObjectImpl
}

func (h *sub) Init() error {
	h.log = log.Sub(user.AppName)
	h.service = ioc.Controller().Get(user.AppName).(user.Service)
	h.Registry()
	return nil
}

func (h *sub) Name() string {
	return "account"
}

func (h *sub) Version() string {
	return "v1"
}

func (h *sub) Registry() {
	tags := []string{"账号管理"}

	ws := gorestful.ObjectRouter(h)
	ws.Route(ws.POST("/password").To(h.UpdatePassword).
		Doc("子账号修改自己密码").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Auth, label.Enable).
		Metadata(label.PERMISSION_MODE, label.PERMISSION_MODE_ACL.Value()).
		Metadata(label.Allow, user.TypeToString(user.TYPE_SUB)).
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
