package role

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcenter/apps/endpoint"
	"github.com/infraboard/mcenter/common/format"
	request "github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/infraboard/mcube/pb/resource"
	"github.com/infraboard/mcube/tools/hash"
)

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

// NewQueryRoleRequestFromHTTP 列表查询请求
func NewQueryRoleRequestFromHTTP(r *http.Request) *QueryRoleRequest {
	page := request.NewPageRequestFromHTTP(r)

	req := NewQueryRoleRequest()
	req.Page = page
	return req
}

// NewQueryRoleRequest 列表查询请求
func NewQueryRoleRequest() *QueryRoleRequest {
	return &QueryRoleRequest{
		Page: request.NewDefaultPageRequest(),
	}
}

// Validate todo
func (req *QueryRoleRequest) Validate() error {
	return nil
}

// NewDescribeRoleRequestWithID todo
func NewDescribeRoleRequestWithID(id string) *DescribeRoleRequest {
	return &DescribeRoleRequest{
		Id: id,
	}
}

// NewDescribeRoleRequestWithName todo
func NewDescribeRoleRequestWithName(name string) *DescribeRoleRequest {
	return &DescribeRoleRequest{
		Name: name,
	}
}

// Validate todo
func (req *DescribeRoleRequest) Validate() error {
	if req.Id == "" && req.Name == "" {
		return fmt.Errorf("id or name required")
	}

	return nil
}

// NewDeleteRoleWithID todo
func NewDeleteRoleWithID(id string) *DeleteRoleRequest {
	return &DeleteRoleRequest{
		Id: id,
	}
}

// New 新创建一个Role
func New(req *CreateRoleRequest) (*Role, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	r := &Role{
		Meta:        resource.NewMeta(),
		Spec:        req,
		Permissions: []*Permission{},
	}
	r.Meta.Id = hash.FnvHash(r.FullName())
	return r, nil
}

// NewDefaultRole 默认实例
func NewDefaultRole() *Role {
	spec := NewCreateRoleRequest()
	return &Role{
		Spec: spec,
	}
}

// NewCreateRoleRequest 实例化请求
func NewCreateRoleRequest() *CreateRoleRequest {
	return &CreateRoleRequest{
		Lables: map[string]string{},
		Specs:  []*Spec{},
	}
}

// Validate 请求校验
func (req *CreateRoleRequest) Validate() error {
	return validate.Struct(req)
}

func (r *Role) ToJson() string {
	return format.Prettify(r)
}

func (r *Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		*resource.Meta
		*CreateRoleRequest
		Permissions []*Permission `json:"permissions"`
	}{r.Meta, r.Spec, r.Permissions})
}

func (r *Role) FullName() string {
	return fmt.Sprintf("%s@%s", r.Spec.Name, r.Spec.Domain)
}

// HasPermission 权限判断
func (r *Role) HasPermission(ep *endpoint.Endpoint) (*Permission, bool, error) {
	var (
		rok, lok bool
	)
	for i := range r.Permissions {
		perm := r.Permissions[i]

		rok = perm.MatchResource(ep.ServiceId, ep.Entry.Resource)
		lok = perm.MatchLabel(ep.Entry.Labels)
		zap.L().Debugf("resource match: service_id: %s[target: %s] resource: %s[target: %s], result: %v",
			ep.ServiceId, perm.Spec.ServiceId, ep.Entry.Resource, perm.Spec.ResourceName, rok)
		zap.L().Debugf("label match: %v from [key: %v, value: %v, result: %v",
			ep.Entry.Labels, perm.Spec.LabelKey, perm.Spec.LabelValues, lok)
		if rok && lok {
			return perm, true, nil
		}
	}
	return nil, false, nil
}

// NewRoleSet 实例化make
func NewRoleSet() *RoleSet {
	return &RoleSet{
		Items: []*Role{},
	}
}

// Add todo
func (s *RoleSet) Add(item *Role) {
	s.Total++
	s.Items = append(s.Items, item)
}

func (s *RoleSet) ToJson() string {
	return format.Prettify(s)
}

// HasPermission todo
func (s *RoleSet) HasPermission(ep *endpoint.Endpoint) (*Permission, bool, error) {
	for i := range s.Items {
		role := s.Items[i]
		p, ok, err := role.HasPermission(ep)
		if err != nil {
			return nil, false, err
		}

		if ok {
			return p, ok, nil
		}
	}

	return nil, false, nil
}

func (s *RoleSet) Permissions() *PermissionSet {
	ps := NewPermissionSet()

	for i := range s.Items {
		ps.Add(s.Items[i].Permissions...)
	}

	return ps
}

func (s *RoleSet) Len() int {
	return len(s.Items)
}

func (s *RoleSet) RoleNames() []string {
	set := []string{}
	for i := range s.Items {
		set = append(set, s.Items[i].Spec.Name)

	}

	return set
}
