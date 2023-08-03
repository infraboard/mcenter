package label

import (
	context "context"
	"strings"

	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcenter/apps/token"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/validator"
)

const (
	AppName = "label"
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
		Label:    value,
		Input:    value,
		Children: []*EnumOption{},
	}
}

func (e *EnumOption) SetLabel(label string) *EnumOption {
	e.Label = label
	return e
}

func (e *EnumOption) HasChildren() bool {
	return len(e.Children) > 0
}

func (r *EnumOption) AddChildren(opts ...*EnumOption) *EnumOption {
	r.Children = append(r.Children, opts...)
	return r
}

func NewRequiredCondition() *RequiredCondition {
	return &RequiredCondition{
		Resources: []string{},
	}
}

func (r *RequiredCondition) AddResource(resource string) {
	r.Resources = append(r.Resources, resource)
}

func NewHttpEnumConfig() *HttpEnumConfig {
	return &HttpEnumConfig{}
}

func NewQueryLabelRequest() *QueryLabelRequest {
	return &QueryLabelRequest{
		Page: request.NewDefaultPageRequest(),
		Keys: []string{},
	}
}

func NewQueryLabelRequestFromHTTP(r *restful.Request) *QueryLabelRequest {
	req := NewQueryLabelRequest()
	req.Page = request.NewPageRequestFromHTTP(r.Request)
	req.Scope = token.GetTokenFromRequest(r).GenScope()

	keys := r.QueryParameter("keys")
	if keys != "" {
		req.Keys = strings.Split(keys, ",")
	}
	return req
}

func (r *QueryLabelRequest) AddKey(keys ...string) {
	r.Keys = append(r.Keys, keys...)
}

func NewDescribeLabelRequest(id string) *DescribeLabelRequest {
	return &DescribeLabelRequest{
		Id: id,
	}
}

func NewDeleteLabelRequest(id string) *DeleteLabelRequest {
	return &DeleteLabelRequest{
		Id: id,
	}
}

func (r *DescribeLabelRequest) Validate() error {
	return validator.Validate(r)
}

func NewUpdateLabelRequest() *UpdateLabelRequest {
	return &UpdateLabelRequest{
		Spec: NewCreateLabelRequest(),
	}
}
