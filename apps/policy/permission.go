package policy

import (
	"fmt"

	request "github.com/infraboard/mcube/http/request"
)

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

	if req.ServiceId == "" || req.Path == "" {
		return fmt.Errorf("service_id and path required when check")
	}

	return nil
}

func (req *CheckPermissionRequest) UniquePath() string {
	return fmt.Sprintf("%s.%s", req.Method, req.Path)
}
