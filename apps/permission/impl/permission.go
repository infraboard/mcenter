package impl

import (
	"context"
	"fmt"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/request"

	"github.com/infraboard/mcenter/apps/permission"
	"github.com/infraboard/mcenter/apps/policy"
)

func (s *service) QueryPermission(ctx context.Context, req *permission.QueryPermissionRequest) (
	*permission.PermissionSet, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest("validate param error, %s", err)
	}

	// 获取用户的策略列表
	preq := policy.NewQueryPolicyRequest(request.NewPageRequest(100, 1))
	preq.Username = req.Username
	preq.Namespace = req.Namespace

	policySet, err := s.policy.QueryPolicy(ctx, preq)
	if err != nil {
		return nil, err
	}

	// 获取用户的角色列表
	rset, err := policySet.GetRoles(ctx, s.role)
	if err != nil {
		return nil, err
	}
	fmt.Println(rset)

	return s.queryPermission(ctx, req)
}
