package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/mcenter/apps/scm"
)

var (
	h = &handler{}
)

type handler struct {
	log logger.Logger
	scm scm.Service
}

func (h *handler) Config() error {
	h.scm = app.GetInternalApp(scm.AppName).(scm.Service)
	h.log = zap.L().Named(scm.AppName)
	return nil
}

func (h *handler) Name() string {
	return scm.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{"源代码管理"}
	ws.Route(ws.GET("projects").To(h.QueryProject).
		Doc("项目列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.QueryParameter("address", "oauth2 auth code").DataType("string").Required(true)).
		Param(ws.QueryParameter("token", "oauth2 state").DataType("string").Required(false)).
		Writes(scm.ProjectSet{}).
		Returns(200, "OK", scm.ProjectSet{}))

	ws.Route(ws.GET("events/gitlab").To(h.HandleGitlabEvent).
		Doc("处理Gitlab Webhook事件").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(scm.ProjectSet{}).
		Returns(200, "OK", scm.ProjectSet{}))
}

func init() {
	app.RegistryRESTfulApp(h)
}
