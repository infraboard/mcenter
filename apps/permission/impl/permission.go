package impl

import (
	"context"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/request"

	"github.com/infraboard/mcenter/apps/endpoint"
	"github.com/infraboard/mcenter/apps/namespace"
	"github.com/infraboard/mcenter/apps/permission"
	"github.com/infraboard/mcenter/apps/policy"
	"github.com/infraboard/mcenter/apps/role"
)

func (s *service) QueryPermission(ctx context.Context, req *permission.QueryPermissionRequest) (
	*role.PermissionSet, error) {

	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest("validate param error, %s", err)
	}

	// 获取用户的策略列表
	preq := policy.NewQueryPolicyRequest()
	preq.Page = request.NewPageRequest(100, 1)
	preq.Domain = req.Domain
	preq.Username = req.Username
	preq.Namespace = req.Namespace

	policySet, err := s.policy.QueryPolicy(ctx, preq)
	if err != nil {
		return nil, err
	}

	// 获取用户的角色列表
	rset, err := policySet.GetRoles(ctx, s.role, true)
	if err != nil {
		return nil, err
	}

	return rset.Permissions(), nil
}

func (s *service) QueryRole(ctx context.Context, req *permission.QueryRoleRequest) (
	*role.RoleSet, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest("validate param error, %s", err)
	}

	// 获取用户的策略列表
	preq := policy.NewQueryPolicyRequest()
	preq.Page = request.NewPageRequest(100, 1)
	preq.Username = req.Username
	preq.Domain = req.Domain
	preq.Namespace = req.Namespace

	policySet, err := s.policy.QueryPolicy(ctx, preq)
	if err != nil {
		return nil, err
	}

	return policySet.GetRoles(ctx, s.role, req.WithPermission)
}

func (s *service) CheckPermission(ctx context.Context, req *permission.CheckPermissionRequest) (
	*role.Permission, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest("validate param error, %s", err)
	}

	// 判断是否是空间所有者
	ns, err := s.namespace.DescribeNamespace(ctx, namespace.NewDescriptNamespaceRequest(req.Domain, req.Namespace))
	if err != nil {
		return nil, err
	}
	if ns.IsOwner(req.Username) {
		return role.OwnerAdminPermssion(), nil
	}

	// 空间普通用户鉴权
	roleReq := permission.NewQueryRoleRequest(req.Namespace)
	roleReq.WithPermission = true
	roleReq.Username = req.Username
	roleReq.Domain = req.Domain
	roleSet, err := s.QueryRole(ctx, roleReq)
	if err != nil {
		return nil, err
	}

	if roleSet.Len() == 0 {
		return nil, exception.NewPermissionDeny("no permission")
	}

	fn := endpoint.NewDescribeEndpointRequestWithID(endpoint.GenHashID(req.ServiceId, req.Path))
	ep, err := s.endpoint.DescribeEndpoint(ctx, fn)
	if err != nil {
		return nil, err
	}
	s.log.Debugf("check roles %s has permission access endpoint [%s]", roleSet.RoleNames(), ep.Entry)

	// 不需要鉴权
	if !ep.Entry.PermissionEnable {
		return role.NewSkipPermission("endpoint not enable permission check, allow all access"), nil
	}

	// 验证是否有权限访问该功能
	p, ok, err := roleSet.HasPermission(ep)
	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, exception.NewPermissionDeny("in namespace %s, role %s has no permission access endpoint: %s",
			req.Namespace,
			roleSet.RoleNames(),
			ep.Entry.Path,
		)
	}

	return p, nil
}
