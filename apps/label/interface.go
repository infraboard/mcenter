package label

import (
	context "context"

	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcenter/apps/token"
	"github.com/infraboard/mcube/http/request"
)

const (
	AppName = "labels"
)

type Service interface {
	// 创建标签
	CreateLabel(context.Context, *CreateLabelRequest) (*Label, error)
	// 修改标签
	UpdateLabel(context.Context, *UpdateLabelRequest) (*Label, error)
	// 删除标签
	DeleteLabel(context.Context, *DeleteLabelRequest) (*Label, error)
	RPCServer
}

func NewCreateLabelRequest() *CreateLabelRequest {
	return &CreateLabelRequest{
		EnumOptions:       []*EnumOption{},
		HttpEnumConfig:    NewHttpEnumConfig(),
		RequiredCondition: NewRequiredCondition(),
		Extensions:        map[string]string{},
	}
}

func (r *CreateLabelRequest) AddEnumOption(opts ...*EnumOption) {
	r.EnumOptions = append(r.EnumOptions, opts...)
}

func NewEnumOption(value string) *EnumOption {
	return &EnumOption{
		Label: value,
		Value: value,
	}
}

func NewRequiredCondition() *RequiredCondition {
	return &RequiredCondition{
		Resources: []string{},
	}
}

func (e *EnumOption) SetLabel(label string) *EnumOption {
	e.Label = label
	return e
}

func NewHttpEnumConfig() *HttpEnumConfig {
	return &HttpEnumConfig{}
}

func NewQueryLabelRequest() *QueryLabelRequest {
	return &QueryLabelRequest{
		Page: request.NewDefaultPageRequest(),
	}
}

func NewQueryLabelRequestFromHTTP(r *restful.Request) *QueryLabelRequest {
	req := NewQueryLabelRequest()
	req.Page = request.NewPageRequestFromHTTP(r.Request)
	req.Scope = token.GetTokenFromRequest(r).GenScope()
	return req
}
