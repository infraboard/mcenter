package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/http/response"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/mcenter/apps/code"
)

var (
	h = &handler{}
)

type handler struct {
	service code.Service
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(code.AppName)
	h.service = app.GetInternalApp(code.AppName).(code.Service)
	return nil
}

func (h *handler) Name() string {
	return code.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{h.Name()}

	ws.Route(ws.POST("/").To(h.IssueCode).
		Doc("issue verify code").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(code.IssueCodeRequest{}).
		Writes(response.NewData(code.Code{})))

	// ws.Route(ws.GET("/").To(h.QueryBook).
	// 	Doc("get all books").
	// 	Metadata(restfulspec.KeyOpenAPITags, tags).
	// 	Metadata("action", "list").
	// 	Reads(book.CreateBookRequest{}).
	// 	Writes(response.NewData(book.BookSet{})).
	// 	Returns(200, "OK", book.BookSet{}))

	// ws.Route(ws.GET("/{id}").To(h.DescribeBook).
	// 	Doc("get a book").
	// 	Param(ws.PathParameter("id", "identifier of the book").DataType("integer").DefaultValue("1")).
	// 	Metadata(restfulspec.KeyOpenAPITags, tags).
	// 	Writes(response.NewData(book.Book{})).
	// 	Returns(200, "OK", response.NewData(book.Book{})).
	// 	Returns(404, "Not Found", nil))

	// ws.Route(ws.PUT("/{id}").To(h.UpdateBook).
	// 	Doc("update a book").
	// 	Param(ws.PathParameter("id", "identifier of the book").DataType("string")).
	// 	Metadata(restfulspec.KeyOpenAPITags, tags).
	// 	Reads(book.CreateBookRequest{}))

	// ws.Route(ws.PATCH("/{id}").To(h.PatchBook).
	// 	Doc("patch a book").
	// 	Param(ws.PathParameter("id", "identifier of the book").DataType("string")).
	// 	Metadata(restfulspec.KeyOpenAPITags, tags).
	// 	Reads(book.CreateBookRequest{}))

	// ws.Route(ws.DELETE("/{id}").To(h.DeleteBook).
	// 	Doc("delete a book").
	// 	Metadata(restfulspec.KeyOpenAPITags, tags).
	// 	Param(ws.PathParameter("id", "identifier of the book").DataType("string")))
}

func init() {
	app.RegistryRESTfulApp(h)
}
