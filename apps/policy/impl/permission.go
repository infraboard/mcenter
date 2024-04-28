package impl

import (
	"context"

	"github.com/infraboard/mcenter/apps/endpoint"
	"github.com/infraboard/mcenter/apps/namespace"
	"github.com/infraboard/mcenter/apps/policy"
	"github.com/infraboard/mcenter/apps/role"
	"github.com/infraboard/mcenter/apps/user"
	"github.com/infraboard/mcube/v2/exception"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *impl) CheckPermission(ctx context.Context, req *policy.CheckPermissionRequest) (
	*role.Permission, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest("validate param error, %s", err)
	}

	// 判断是否是空间所有者
	ns, err := s.namespace.DescribeNamespace(
		ctx,
		namespace.NewDescriptNamespaceRequestByName(req.Domain, req.Namespace),
	)
	if err != nil {
		return nil, err
	}
	// 空间管理员直接给空间内所有权限
	if ns.IsManager(req.UserId) {
		return role.NamespaceManagerPermssion(), nil
	}

	// 查询用户权限策略
	pReq := policy.NewQueryPolicyRequest()
	pReq.UserId = req.UserId
	pReq.Domain = req.Domain
	pReq.Namespace = req.Namespace
	pReq.WithRole = true
	ps, err := s.QueryPolicy(ctx, pReq)
	if err != nil {
		return nil, err
	}

	// 无用户相关权限策略设置
	if ps.Len() == 0 {
		return nil, exception.NewPermissionDeny("no permission")
	}

	// 查询用户需要鉴权的功能
	fn := endpoint.NewDescribeEndpointRequestWithID(endpoint.GenHashID(req.ServiceId, req.Path))
	ep, err := s.endpoint.DescribeEndpoint(ctx, fn)
	if err != nil {
		return nil, err
	}

	// 判断改功能是否需要鉴权
	if !req.ForceCheck && !ep.Entry.PermissionEnable {
		return role.NewSkipPermission("endpoint not enable permission check, allow all access"), nil
	}

	// 判断策略是否允许
	var perm *role.Permission
	for i := range ps.Items {
		p := ps.Items[i]
		permOk, ok, err := p.Role.HasPermission(ep)
		if err != nil {
			return nil, err
		}
		if ok {
			perm = role.NewPermissionFromSpec(p.Spec.RoleId, permOk)
			perm.Scope = p.Spec.Scope
			s.log.Debug().Msgf("check roles %s has permission access endpoint [%s]", p.Role.Spec.Name, ep.Entry)
		}
	}

	if perm == nil {
		return nil, exception.NewPermissionDeny("in namespace %s, role %s has no permission access endpoint: %s",
			req.Namespace,
			ps.RoleNames(),
			ep.Entry.Path,
		)
	}

	return perm, nil
}

// 查询用户策略允许的空间
func (i *impl) AvailableNamespace(
	ctx context.Context,
	in *policy.AvailableNamespaceRequest) (
	*namespace.NamespaceSet, error) {
	u, err := i.user.DescribeUser(ctx, user.NewDescriptUserRequestById(in.UserId))
	if err != nil {
		return nil, err
	}

	nsReq := namespace.NewQueryNamespaceRequest()
	nsReq.Domain = u.Spec.Domain

	// 如果是子账号, 只返回子账号加入的空间
	if u.Spec.Type.Equal(user.TYPE_SUB) {
		cursor, err := i.col.Aggregate(ctx, bson.A{
			bson.M{"$group": bson.M{"_id": "$namespace"}},
		})
		if err != nil {
			return nil, err
		}

		result := map[string]string{}
		for cursor.Next(ctx) {
			if err := cursor.Decode(&result); err != nil {
				return nil, err
			}
			nsReq.AddName(result["_id"])
		}
	}

	nss, err := i.namespace.QueryNamespace(ctx, nsReq)
	if err != nil {
		return nil, err
	}

	return nss, nil
}
