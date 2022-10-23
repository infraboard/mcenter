package namespace

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcenter/apps/token"
	request "github.com/infraboard/mcube/http/request"
)

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

const (
	AppName = "namespace"
)

const (
	DEFAULT_NAMESPACE = "default"
)

// NewDefaultNamespace todo
func NewDefaultNamespace() *Namespace {
	return &Namespace{
		Spec: &CreateNamespaceRequest{
			Enabled: true,
		},
	}
}

// NewCreateNamespaceRequest todo
func NewCreateNamespaceRequest() *CreateNamespaceRequest {
	return &CreateNamespaceRequest{}
}

func (req *CreateNamespaceRequest) UpdateOwner(tk *token.Token) {
	req.Owner = tk.Username
	req.Domain = tk.Domain
}

// Validate todo
func (req *CreateNamespaceRequest) Validate() error {
	return validate.Struct(req)
}

// NewNamespaceSet 实例化
func NewNamespaceSet() *NamespaceSet {
	return &NamespaceSet{
		Items: []*Namespace{},
	}
}

// Add 添加应用
func (s *NamespaceSet) Add(item *Namespace) {
	s.Items = append(s.Items, item)
}

// NewDescriptNamespaceRequest new实例
func NewDescriptNamespaceRequest(domain, name string) *DescriptNamespaceRequest {
	return &DescriptNamespaceRequest{
		Domain: domain,
		Name:   name,
	}
}

// NewQueryNamespaceRequestFromHTTP 列表查询请求
func NewQueryNamespaceRequestFromHTTP(r *http.Request) *QueryNamespaceRequest {
	qs := r.URL.Query()
	return &QueryNamespaceRequest{
		Page: request.NewPageRequestFromHTTP(r),
		Name: []string{qs.Get("name")},
	}
}

// NewQueryNamespaceRequest 列表查询请求
func NewQueryNamespaceRequest() *QueryNamespaceRequest {
	return &QueryNamespaceRequest{
		Page: request.NewDefaultPageRequest(),
	}
}

func (req *QueryNamespaceRequest) UpdateOwner(tk *token.Token) {
	req.Domain = tk.Domain
	req.Name = []string{tk.Username}
}

// Validate 校验详情查询请求
func (req *DescriptNamespaceRequest) Validate() error {
	if req.Name == "" {
		return errors.New("id  is required")
	}

	return nil
}

// NewDeleteNamespaceRequestWithID todo
func NewDeleteNamespaceRequest(name string) *DeleteNamespaceRequest {
	return &DeleteNamespaceRequest{
		Name: name,
	}
}

// Validate todo
func (req *DeleteNamespaceRequest) Validate() error {
	if req.Name == "" {
		return fmt.Errorf("name required")
	}

	return nil
}
