package permission

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	request "github.com/infraboard/mcube/http/request"
)

const (
	AppName = "permission"
)

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

type Service interface {
	RPCServer
}

// NewQueryPermissionRequest todo
func NewQueryPermissionRequest(page *request.PageRequest) *QueryPermissionRequest {
	return &QueryPermissionRequest{
		Page: page,
	}
}

// Validate 校验请求合法
func (req *QueryPermissionRequest) Validate() error {
	if req.Namespace == "" {
		return fmt.Errorf("namespace required")
	}

	return nil
}

// NewCheckPermissionrequest todo
func NewCheckPermissionRequest() *CheckPermissionRequest {
	return &CheckPermissionRequest{
		Page: request.NewPageRequest(100, 1),
	}
}

// Validate 校验请求合法
func (req *CheckPermissionRequest) Validate() error {
	if req.Namespace == "" {
		return fmt.Errorf("namespace required")
	}

	if req.EndpointId == "" && (req.ServiceId == "" || req.Path == "") {
		return fmt.Errorf("endpoint_id or (service_id and path) required when check")
	}

	return nil
}

// NewQueryRoleRequest todo
func NewQueryRoleRequest(namespaceId string) *QueryRoleRequest {
	return &QueryRoleRequest{
		Page:      request.NewPageRequest(100, 1),
		Namespace: namespaceId,
	}
}

// Validate 校验请求合法
func (req *QueryRoleRequest) Validate() error {
	if req.Namespace == "" {
		return fmt.Errorf("namespace required")
	}

	return nil
}
