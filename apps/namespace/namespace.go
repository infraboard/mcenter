package namespace

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/emicklei/go-restful/v3"
	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcenter/apps/token"
	"github.com/infraboard/mcube/v2/exception"
	request "github.com/infraboard/mcube/v2/http/request"
	resource "github.com/infraboard/mcube/v2/pb/resource"
	"github.com/infraboard/mcube/v2/tools/hash"
	"github.com/infraboard/mcube/v2/tools/pretty"
)

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

// NewNamespace todo
func New(req *CreateNamespaceRequest) (*Namespace, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	ins := &Namespace{
		Meta: resource.NewMeta(),
		Spec: req,
	}

	ins.Meta.Id = hash.FnvHash(req.Domain, req.Name)
	return ins, nil
}

// NewDefaultNamespace todo
func NewDefaultNamespace() *Namespace {
	return &Namespace{
		Spec: NewCreateNamespaceRequest(),
	}
}

func (n *Namespace) IsManager(username string) bool {
	if n.Spec.Owner == username {
		return true
	}

	for i := range n.Spec.Assistants {
		if n.Spec.Assistants[i] == username {
			return true
		}
	}
	return false
}

func (n *Namespace) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		*resource.Meta
		*CreateNamespaceRequest
	}{n.Meta, n.Spec})
}

func (n *Namespace) ToJson() string {
	return pretty.ToJSON(n)
}

// NewCreateNamespaceRequest todo
func NewCreateNamespaceRequest() *CreateNamespaceRequest {
	return &CreateNamespaceRequest{
		Assistants:    []string{},
		Enabled:       true,
		Extension:     map[string]string{},
		Labels:        map[string]string{},
		ResourceQuota: NewResourceQuotaSet(),
	}
}

func (req *CreateNamespaceRequest) UpdateOwner(tk *token.Token) {
	if tk == nil {
		return
	}

	// 默认Owner是自己
	if req.Owner != "" {
		req.Owner = tk.Username
	}

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

func (s *NamespaceSet) ToJson() string {
	return pretty.ToJSON(s)
}

// Add 添加应用
func (s *NamespaceSet) Add(item *Namespace) {
	s.Items = append(s.Items, item)
}

// NewDescriptNamespaceRequestByName new实例
func NewDescriptNamespaceRequestByName(domain, name string) *DescriptNamespaceRequest {
	return &DescriptNamespaceRequest{
		DescribeBy: DESCRIBE_BY_NAME,
		Domain:     domain,
		Name:       name,
	}
}

// NewDescriptNamespaceRequestById new实例
func NewDescriptNamespaceRequestById(id string) *DescriptNamespaceRequest {
	return &DescriptNamespaceRequest{
		DescribeBy: DESCRIBE_BY_ID,
		Id:         id,
	}
}

// NewQueryNamespaceRequestFromHTTP 列表查询请求
func NewQueryNamespaceRequestFromHTTP(r *restful.Request) *QueryNamespaceRequest {
	return &QueryNamespaceRequest{
		Page:     request.NewPageRequestFromHTTP(r.Request),
		Name:     []string{r.QueryParameter("name")},
		Username: r.QueryParameter("username"),
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
	switch req.DescribeBy {
	case DESCRIBE_BY_NAME:
		if req.Name == "" || req.Domain == "" {
			return errors.New("name and domain is required")
		}
	default:
		if req.Id == "" {
			return errors.New("id is required")
		}
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

func NewResourceQuotaSet() *ResourceQuotaSet {
	return &ResourceQuotaSet{
		Items: []*ResourceQuota{},
	}
}

func (s *ResourceQuotaSet) Add(items ...*ResourceQuota) {
	s.Items = append(s.Items, items...)
}

func NewResourceQuota(resource string, soft, hard int64) *ResourceQuota {
	return &ResourceQuota{
		ResourceName: resource,
		Soft:         soft,
		Hard:         hard,
	}
}
