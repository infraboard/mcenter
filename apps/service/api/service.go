package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/restful/response"
	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/mcenter/apps/service"
	"github.com/infraboard/mcenter/apps/token"
)

func init() {
	ioc.RegistryApi(&serviceHandler{})
}

type serviceHandler struct {
	service service.MetaService
	log     logger.Logger
	ioc.IocObjectImpl
}

func (h *serviceHandler) Init() error {
	h.log = zap.L().Named(service.AppName)
	h.service = ioc.GetController(service.AppName).(service.MetaService)
	return nil
}

func (h *serviceHandler) Name() string {
	return service.AppName
}

func (h *serviceHandler) Version() string {
	return "v1"
}

func (h *serviceHandler) Registry(ws *restful.WebService) {
	tags := []string{"服务管理"}

	ws.Route(ws.POST("/").To(h.CreateService).
		Doc("创建服务").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Create.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(service.CreateServiceRequest{}).
		Writes(service.Service{}))

	ws.Route(ws.GET("/").To(h.QueryService).
		Doc("查询服务列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.List.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Disable).
		Reads(service.QueryServiceRequest{}).
		Writes(service.ServiceSet{}).
		Returns(200, "OK", service.ServiceSet{}))

	ws.Route(ws.GET("/{id}").To(h.DescribeService).
		Doc("查询服务详情").
		Param(ws.PathParameter("id", "identifier of the service").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Get.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Disable).
		Writes(service.Service{}).
		Returns(200, "OK", service.Service{}).
		Returns(404, "Not Found", nil))

	ws.Route(ws.PUT("/{id}").To(h.UpdateService).
		Doc("更新服务").
		Param(ws.PathParameter("id", "identifier of the service").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Update.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(service.CreateServiceRequest{}))

	ws.Route(ws.PATCH("/{id}").To(h.PatchService).
		Doc("更新服务").
		Param(ws.PathParameter("id", "identifier of the service").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Update.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Reads(service.CreateServiceRequest{}))

	ws.Route(ws.DELETE("/{id}").To(h.DeleteService).
		Doc("删除服务").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata(label.Resource, h.Name()).
		Metadata(label.Action, label.Delete.Value()).
		Metadata(label.Auth, label.Enable).
		Metadata(label.Permission, label.Enable).
		Param(ws.PathParameter("id", "identifier of the service").DataType("string")))
}

func (h *serviceHandler) CreateService(r *restful.Request, w *restful.Response) {
	req := service.NewCreateServiceRequest()

	if err := r.ReadEntity(req); err != nil {
		response.Failed(w, err)
		return
	}

	// 补充用户信息
	tk := token.GetTokenFromRequest(r)
	req.Domain = tk.Domain
	req.Namespace = tk.Namespace
	req.Owner = tk.UserId

	set, err := h.service.CreateService(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, set)
}

func (h *serviceHandler) QueryService(r *restful.Request, w *restful.Response) {
	req := service.NewQueryServiceRequestFromHTTP(r)
	set, err := h.service.QueryService(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *serviceHandler) DescribeService(r *restful.Request, w *restful.Response) {
	req := service.NewDescribeServiceRequest(r.PathParameter("id"))
	ins, err := h.service.DescribeService(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *serviceHandler) UpdateService(r *restful.Request, w *restful.Response) {
	req := service.NewPutServiceRequest(r.PathParameter("id"))

	if err := r.ReadEntity(req.Spec); err != nil {
		response.Failed(w, err)
		return
	}

	set, err := h.service.UpdateService(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *serviceHandler) PatchService(r *restful.Request, w *restful.Response) {
	req := service.NewPatchServiceRequest(r.PathParameter("id"))

	if err := r.ReadEntity(req.Spec); err != nil {
		response.Failed(w, err)
		return
	}

	set, err := h.service.UpdateService(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *serviceHandler) DeleteService(r *restful.Request, w *restful.Response) {
	req := service.NewDeleteServiceRequestWithID(r.PathParameter("id"))
	set, err := h.service.DeleteService(r.Request.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}
