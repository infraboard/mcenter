package role

import (
	"fmt"
	"hash/fnv"
	"net/http"
	"time"

	"github.com/infraboard/mcenter/apps/endpoint"
	"github.com/infraboard/mcenter/common/format"
	"github.com/infraboard/mcube/exception"
	request "github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/logger/zap"
)

// NewAddPermissionToRoleRequest todo
func NewAddPermissionToRoleRequest() *AddPermissionToRoleRequest {
	return &AddPermissionToRoleRequest{
		Permissions: []*Spec{},
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
		Spec: &Spec{
			Effect: EffectType_ALLOW,
			Desc:   message,
		},
	}
}

func NewPermission(roleId string, perms []*Spec) []*Permission {
	set := []*Permission{}
	for i := range perms {
		set = append(set, &Permission{
			Id:       perms[i].HashID(roleId),
			CreateAt: time.Now().Unix(),
			RoleId:   roleId,
			Spec:     perms[i],
		})
	}
	return set
}

func (req *Spec) HashID(roleId string) string {
	h := fnv.New32a()

	h.Write([]byte(roleId + req.Effect.String() + req.ServiceId + req.ResourceName))
	return fmt.Sprintf("%x", h.Sum32())
}

// NewDefaultPermission todo
func NewDefaultPermission() *Spec {
	return &Spec{
		Effect: EffectType_ALLOW,
	}
}

// Validate todo
func (p *Spec) Validate() error {
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
	return format.Prettify(p)
}

// MatchResource 检测资源是否匹配
func (p *Permission) MatchResource(serviceID, resourceName string) bool {
	// 服务匹配
	if p.Spec.ServiceId != "*" && p.Spec.ServiceId != serviceID {
		return false
	}

	// 资源匹配
	if p.Spec.ResourceName != "*" && p.Spec.ResourceName != resourceName {
		return false
	}

	return true
}

// MatchLabel 匹配Label
func (p *Permission) MatchLabel(label map[string]string) bool {
	for k, v := range label {
		// 匹配key
		if p.Spec.LabelKey == "*" || p.Spec.LabelKey == k {
			// 匹配value
			if p.isMatchAllValue() {
				return true
			}
			for i := range p.Spec.LabelValues {
				if p.Spec.LabelValues[i] == v {
					return true
				}
			}
		}
	}

	return false
}

func (p *Permission) isMatchAllValue() bool {
	if p.Spec.MatchAll {
		return true
	}

	for i := range p.Spec.LabelValues {
		if p.Spec.LabelValues[i] == "*" {
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

// HasPermission 权限判断
func (s *PermissionSet) HasPermission(ep *endpoint.Endpoint) (*Permission, bool, error) {
	var (
		rok, lok bool
	)
	for i := range s.Items {
		perm := s.Items[i]

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

// NewCheckPermissionrequest todo
func NewCheckPermissionRequest() *CheckPermissionRequest {
	return &CheckPermissionRequest{
		Page: request.NewPageRequest(100, 1),
	}
}
