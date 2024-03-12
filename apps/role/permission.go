package role

import (
	"fmt"
	"hash/fnv"
	"net/http"
	"time"

	"github.com/infraboard/mcenter/apps/endpoint"
	"github.com/infraboard/mcube/v2/exception"
	request "github.com/infraboard/mcube/v2/http/request"
	"github.com/infraboard/mcube/v2/ioc/config/log"
	"github.com/infraboard/mcube/v2/tools/pretty"
)

// NewAddPermissionToRoleRequest todo
func NewAddPermissionToRoleRequest() *AddPermissionToRoleRequest {
	return &AddPermissionToRoleRequest{
		Permissions: []*PermissionSpec{},
	}
}

func (req *AddPermissionToRoleRequest) Validate() error {
	return validate.Struct(req)
}

func (req *AddPermissionToRoleRequest) Length() int {
	return len(req.Permissions)
}

// NewRemovePermissionFromRoleRequest todo
func NewRemovePermissionFromRoleRequest() *RemovePermissionFromRoleRequest {
	return &RemovePermissionFromRoleRequest{
		PermissionId: []string{},
	}
}

func (req *RemovePermissionFromRoleRequest) Validate() error {
	return validate.Struct(req)
}

// NewQueryPermissionRequest todo
func NewQueryPermissionRequest() *QueryPermissionRequest {
	return &QueryPermissionRequest{
		Page: request.NewDefaultPageRequest(),
	}
}

// NewQueryPermissionRequestFromHTTP 列表查询请求
func NewQueryPermissionRequestFromHTTP(r *http.Request) *QueryPermissionRequest {
	req := NewQueryPermissionRequest()
	req.Page = request.NewPageRequestFromHTTP(r)

	return req
}

func NewDescribePermissionRequestWithID(id string) *DescribePermissionRequest {
	return &DescribePermissionRequest{Id: id}
}

func (req *DescribePermissionRequest) Validate() error {
	if req.Id == "" {
		return exception.NewBadRequest("id required")
	}
	return nil
}

func NewUpdatePermissionRequest() *UpdatePermissionRequest {
	return &UpdatePermissionRequest{}
}

func (req *UpdatePermissionRequest) Validate() error {
	if req.Id == "" {
		return exception.NewBadRequest("id required")
	}

	return nil
}

func NewDeaultPermission() *Permission {
	return &Permission{}
}

func NewSkipPermission(message string) *Permission {
	return &Permission{
		Spec: &PermissionSpec{
			Effect: EffectType_ALLOW,
			Desc:   message,
		},
	}
}

func NewPermission(roleId string, perms ...*PermissionSpec) []*Permission {
	set := []*Permission{}
	for i := range perms {
		set = append(set, NewPermissionFromSpec(roleId, perms[i]))
	}
	return set
}

func NewPermissionFromSpec(roleId string, spec *PermissionSpec) *Permission {
	return &Permission{
		Id:       spec.HashID(roleId),
		CreateAt: time.Now().Unix(),
		RoleId:   roleId,
		Spec:     spec,
	}
}

func (req *PermissionSpec) HashID(roleId string) string {
	h := fnv.New32a()

	h.Write([]byte(roleId + req.Effect.String() + req.ServiceId + req.ResourceName))
	return fmt.Sprintf("%x", h.Sum32())
}

// NewDefaultPermissionSpec todo
func NewDefaultPermissionSpec() *PermissionSpec {
	return &PermissionSpec{
		Effect: EffectType_ALLOW,
	}
}

// Validate todo
func (p *PermissionSpec) Validate() error {
	if p.ServiceId == "" || p.ResourceName == "" || p.LabelKey == "" {
		return fmt.Errorf("permisson required service_id, resource_name and label_key")
	}

	if len(p.LabelValues) == 0 {
		return fmt.Errorf("permission label_values required")
	}

	return nil
}

// ID 计算唯一ID
func (p *Permission) ID(namespace string) string {
	return namespace + "." + p.Spec.ResourceName
}

func (p *Permission) ToJson() string {
	return pretty.ToJSON(p)
}

// MatchResource 检测资源是否匹配
func (p *PermissionSpec) MatchResource(serviceID, resourceName string) bool {
	// 服务匹配
	if p.ServiceId != "*" && p.ServiceId != serviceID {
		return false
	}

	// 资源匹配
	if p.ResourceName != "*" && p.ResourceName != resourceName {
		return false
	}

	return true
}

// MatchLabel 匹配Label
func (p *PermissionSpec) MatchLabel(label map[string]string) bool {
	for k, v := range label {
		// 匹配key
		if p.LabelKey == "*" || p.LabelKey == k {
			// 匹配value
			if p.isMatchAllValue() {
				return true
			}
			for i := range p.LabelValues {
				if p.LabelValues[i] == v {
					return true
				}
			}
		}
	}

	return false
}

func (p *PermissionSpec) isMatchAllValue() bool {
	if p.MatchAll {
		return true
	}

	for i := range p.LabelValues {
		if p.LabelValues[i] == "*" {
			return true
		}
	}

	return false
}

// NewPermissionSet todo
func NewPermissionSet() *PermissionSet {
	return &PermissionSet{
		Items: []*Permission{},
	}
}

func (req *QueryPermissionRequest) Validate() error {
	return validate.Struct(req)
}

// Add todo
func (s *PermissionSet) Add(items ...*Permission) {
	s.Items = append(s.Items, items...)
}

func (s *PermissionSet) ToJSON() string {
	return pretty.ToJSON(s)
}

func (s *PermissionSet) PermissionSpecs() (items []*PermissionSpec) {
	for i := range s.Items {
		item := s.Items[i]
		items = append(items, item.Spec)
	}
	return
}

// HasPermission 权限判断
func (s *PermissionSet) HasPermission(ep *endpoint.Endpoint) (*Permission, bool, error) {
	var (
		rok, lok bool
	)
	for i := range s.Items {
		perm := s.Items[i]

		rok = perm.Spec.MatchResource(ep.ServiceId, ep.Entry.Resource)
		lok = perm.Spec.MatchLabel(ep.Entry.Labels)

		log.L().Debug().Msgf("resource match: service_id: %s[target: %s] resource: %s[target: %s], result: %v",
			ep.ServiceId, perm.Spec.ServiceId, ep.Entry.Resource, perm.Spec.ResourceName, rok)
		log.L().Debug().Msgf("label match: %v from [key: %v, value: %v, result: %v",
			ep.Entry.Labels, perm.Spec.LabelKey, perm.Spec.LabelValues, lok)
		if rok && lok {

			return perm, true, nil
		}
	}
	return nil, false, nil
}

// NewCheckPermissionrequest todo
func NewCheckPermissionRequest() *CheckPermissionRequest {
	return &CheckPermissionRequest{
		Page: request.NewPageRequest(100, 1),
	}
}
