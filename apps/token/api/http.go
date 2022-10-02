package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/http/response"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/mcenter/apps/domain"
	"github.com/infraboard/mcenter/apps/token"
)

var (
	h = &handler{}
)

type handler struct {
	service token.Service
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(token.AppName)
	h.service = app.GetInternalApp(token.AppName).(token.Service)
	return nil
}

func (h *handler) Name() string {
	return token.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{"登录"}

	ws.Route(ws.POST("/").To(h.IssueToken).
		Doc("颁发令牌").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(domain.CreateDomainRequest{}).
		Writes(response.NewData(domain.Domain{})))

	ws.Route(ws.DELETE("/").To(h.RevolkToken).
		Doc("撤销令牌").
		Param(ws.PathParameter("id", "identifier of the domain").DataType("integer").DefaultValue("1")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(response.NewData(domain.Domain{})).
		Returns(200, "OK", response.NewData(domain.Domain{})).
		Returns(404, "Not Found", nil))

	ws.Route(ws.PATCH("/").To(h.ChangeNamespace).
		Doc("切换空间").
		Param(ws.PathParameter("id", "identifier of the domain").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(domain.CreateDomainRequest{}))
}

func init() {
	app.RegistryRESTfulApp(h)
}
