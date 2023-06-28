package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/mcenter/apps/code"
)

func init() {
	ioc.RegistryApi(&handler{})
}

type handler struct {
	service code.Service
	log     logger.Logger
	ioc.IocObjectImpl
}

func (h *handler) Init() error {
	h.log = zap.L().Named(code.AppName)
	h.service = ioc.GetController(code.AppName).(code.Service)
	return nil
}

func (h *handler) Name() string {
	return code.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{"验证码管理"}

	ws.Route(ws.POST("/").To(h.IssueCode).
		Doc("颁发验证码").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Auth, label.Enable).
		Metadata(label.PERMISSION_MODE, label.PERMISSION_MODE_ACL.Value()).
		Reads(code.IssueCodeRequest{}).
		Writes(code.Code{}))
}
