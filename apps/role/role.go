package role

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	request "github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/types/ftime"
	"github.com/rs/xid"
)

const (
	AppName = "role"
)

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

// NewQueryRoleRequestFromHTTP 列表查询请求
func NewQueryRoleRequestFromHTTP(r *http.Request) *QueryRoleRequest {
	page := request.NewPageRequestFromHTTP(r)

	req := NewQueryRoleRequest(page)
	return req
}

// NewQueryRoleRequest 列表查询请求
func NewQueryRoleRequest(pageReq *request.PageRequest) *QueryRoleRequest {
	return &QueryRoleRequest{
		Page: pageReq,
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
		Id:       xid.New().String(),
		CreateAt: ftime.Now().Timestamp(),
		UpdateAt: ftime.Now().Timestamp(),
		Spec:     req,
	}
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
		Type: RoleType_CUSTOM,
		Meta: map[string]string{},
	}
}

// IsCumstomType todo
func (req *CreateRoleRequest) IsCumstomType() bool {
	return req.Type.Equal(RoleType_CUSTOM)
}

// Validate 请求校验
func (req *CreateRoleRequest) Validate() error {
	return validate.Struct(req)
}

// CheckPermission 检测该角色是否具有该权限
func (r *Role) CheckPermission() error {
	return nil
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
