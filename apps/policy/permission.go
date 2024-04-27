package policy

import (
	"fmt"

	"github.com/emicklei/go-restful/v3"
	request "github.com/infraboard/mcube/v2/http/request"
	"github.com/infraboard/mcube/v2/tools/pretty"
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

func (req *CheckPermissionRequest) ToJSON() string {
	return pretty.ToJSON(req)
}

func NewAvailableNamespaceRequest() *AvailableNamespaceRequest {
	return &AvailableNamespaceRequest{
		Page: request.NewDefaultPageRequest(),
	}
}

func NewAvailableNamespaceRequestFromHTTP(r *restful.Request) *AvailableNamespaceRequest {
	page := request.NewPageRequestFromHTTP(r.Request)
	req := NewAvailableNamespaceRequest()
	req.Page = page
	return req
}

// NewQueryRoleRequestFromHTTP 列表查询请求
// func NewQueryPolicyRequestFromHTTP() *QueryPolicyRequest {
// 	page := request.NewPageRequestFromHTTP(r.Request)
// 	req := NewQueryPolicyRequest()
// 	req.Page = page

// 	tk := token.GetTokenFromRequest(r)
// 	req.Domain = tk.Domain
// 	req.Namespace = tk.Namespace
// 	req.WithRole = r.QueryParameter("with_role") == "true"
// 	req.WithNamespace = r.QueryParameter("with_namespace") == "true"
// 	return req
// }
