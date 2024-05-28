package role

import (
	"encoding/json"
	"fmt"

	"github.com/emicklei/go-restful/v3"
	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcenter/apps/endpoint"
	"github.com/infraboard/mcenter/apps/token"
	"github.com/infraboard/mcube/v2/exception"
	request "github.com/infraboard/mcube/v2/http/request"
	"github.com/infraboard/mcube/v2/ioc/config/log"
	"github.com/infraboard/mcube/v2/pb/resource"
	"github.com/infraboard/mcube/v2/tools/hash"
	"github.com/infraboard/mcube/v2/tools/pretty"
)

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

// NewQueryRoleRequestFromHTTP 列表查询请求
func NewQueryRoleRequestFromHTTP(r *restful.Request) *QueryRoleRequest {
	req := NewQueryRoleRequest()
	req.Page = request.NewPageRequestFromHTTP(r.Request)
	req.Scope = token.GetTokenFromRequest(r).GenScope()
	req.WithPermission = r.QueryParameter("with_permission") == "true"
	return req
}

// NewQueryRoleRequest 列表查询请求
func NewQueryRoleRequest() *QueryRoleRequest {
	return &QueryRoleRequest{
		Page:    request.NewDefaultPageRequest(),
		RoleIds: []string{},
	}
}

// Validate todo
func (req *QueryRoleRequest) AddRoleId(rid ...string) *QueryRoleRequest {
	req.RoleIds = append(req.RoleIds, rid...)
	return req
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
		Permissions: []*PermissionSpec{},
	}
	r.Meta.Id = hash.FnvHash(r.FullName())
	return r, nil
}

// NewDefaultRole 默认实例
func NewDefaultRole() *Role {
	return &Role{
		Meta:        resource.NewMeta(),
		Spec:        NewCreateRoleRequest(),
		Permissions: []*PermissionSpec{},
	}
}

// NewCreateRoleRequest 实例化请求
func NewCreateRoleRequest() *CreateRoleRequest {
	return &CreateRoleRequest{
		Lables:  map[string]string{},
		Specs:   []*PermissionSpec{},
		Enabled: true,
	}
}

// Validate 请求校验
func (req *CreateRoleRequest) Validate() error {
	return validate.Struct(req)
}

func (req *CreateRoleRequest) UpdateFromToken(tk *token.Token) {
	if tk == nil {
		return
	}
	req.CreateBy = tk.UserId
	req.Domain = tk.Domain
	req.Namespace = tk.Namespace
}

func (r *Role) ToJson() string {
	return pretty.ToJSON(r)
}

func (r *Role) CheckScope(s *resource.Scope) error {
	if s == nil {
		return nil
	}

	if !(r.Spec.Domain == s.Domain && r.Spec.Namespace == s.Namespace) {
		return exception.NewPermissionDeny("资源不属于当前空间")
	}
	return nil
}

func (r *Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		*resource.Meta
		*CreateRoleRequest
		Permissions []*PermissionSpec `json:"permissions"`
	}{r.Meta, r.Spec, r.Permissions})
}

func (r *Role) FullName() string {
	return fmt.Sprintf("%s@%s", r.Spec.Name, r.Spec.Domain)
}

// HasPermission 权限判断
func (r *Role) HasPermission(ep *endpoint.Endpoint) (*PermissionSpec, bool, error) {
	var (
		rok, lok bool
	)
	for i := range r.Permissions {
		perm := r.Permissions[i]

		rok = perm.MatchResource(ep.ServiceId, ep.Entry.Resource)
		log.L().Debug().Msgf("resource match: service_id: %s[target: %s] resource: %s[target: %s], result: %v",
			ep.ServiceId, perm.ServiceId, ep.Entry.Resource, perm.ResourceName, rok)

		lok = perm.MatchLabel(ep.Entry.Labels)
		log.L().Debug().Msgf("label match: %v from [key: %v, value: %v, result: %v",
			ep.Entry.Labels, perm.LabelKey, perm.LabelValues, lok)
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
	return pretty.ToJSON(s)
}

// HasPermission todo
func (s *RoleSet) HasPermission(ep *endpoint.Endpoint) (*Permission, bool, error) {
	for i := range s.Items {
		r := s.Items[i]
		spec, ok, err := r.HasPermission(ep)
		if err != nil {
			return nil, false, err
		}

		if ok {
			p := NewDeaultPermission()
			p.RoleId = r.Meta.Id
			p.Spec = spec
			return p, ok, nil
		}
	}

	return nil, false, nil
}

func (s *RoleSet) Permissions() *PermissionSet {
	ps := NewPermissionSet()

	for i := range s.Items {
		r := s.Items[i]
		ps.Add(NewPermission(r.Meta.Id, r.Permissions...)...)
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
