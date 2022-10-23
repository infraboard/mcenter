package impl

import (
	"context"
	"fmt"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/request"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/mcenter/apps/policy"
	"github.com/infraboard/mcenter/apps/role"
)

func (s *impl) CreateRole(ctx context.Context, req *role.CreateRoleRequest) (*role.Role, error) {
	r, err := role.New(req)
	if err != nil {
		return nil, err
	}

	// 保存角色
	if _, err := s.role.InsertOne(ctx, r); err != nil {
		return nil, exception.NewInternalServerError("inserted role(%s) document error, %s",
			r.Spec.Name, err)
	}

	// 保存角色的权限
	addReq := role.NewAddPermissionToRoleRequest()
	addReq.CreateBy = req.CreateBy
	addReq.Permissions = req.Specs
	addReq.RoleId = r.Id
	perms, err := s.AddPermissionToRole(ctx, addReq)
	if err != nil {
		return nil, err
	}
	r.Permissions = perms.Items
	return r, nil
}

func (s *impl) QueryRole(ctx context.Context, req *role.QueryRoleRequest) (*role.RoleSet, error) {
	query, err := newQueryRoleRequest(req)
	if err != nil {
		return nil, err
	}

	s.log.Debugf("query role filter: %s", query.FindFilter())
	resp, err := s.role.Find(ctx, query.FindFilter(), query.FindOptions())
	if err != nil {
		return nil, exception.NewInternalServerError("find role error, error is %s", err)
	}

	set := role.NewRoleSet()
	// 循环
	for resp.Next(ctx) {
		ins := role.NewDefaultRole()
		if err := resp.Decode(ins); err != nil {
			return nil, exception.NewInternalServerError("decode role error, error is %s", err)
		}
		// 补充权限
		if req.WithPermission {
			pReq := role.NewQueryPermissionRequest()
			pReq.RoleId = ins.Id
			pReq.Page = request.NewPageRequest(role.RoleMaxPermission, 1)
			ps, err := s.QueryPermission(ctx, pReq)
			if err != nil {
				return nil, err
			}
			ins.Permissions = ps.Items
		}
		set.Add(ins)
	}

	// count
	count, err := s.role.CountDocuments(ctx, query.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get token count error, error is %s", err)
	}
	set.Total = count

	return set, nil
}

func (s *impl) DescribeRole(ctx context.Context, req *role.DescribeRoleRequest) (*role.Role, error) {
	query, err := newDescribeRoleRequest(req)
	if err != nil {
		return nil, err
	}

	ins := role.NewDefaultRole()
	if err := s.role.FindOne(ctx, query.FindFilter(), query.FindOptions()).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("role %s not found", req)
		}

		return nil, exception.NewInternalServerError("find role %s error, %s", req, err)
	}

	// 补充权限
	if req.WithPermission {
		pReq := role.NewQueryPermissionRequest()
		pReq.RoleId = ins.Id
		pReq.Page = request.NewPageRequest(role.RoleMaxPermission, 1)
		ps, err := s.QueryPermission(ctx, pReq)
		if err != nil {
			return nil, err
		}
		ins.Permissions = ps.Items
	}

	return ins, nil
}

func (s *impl) DeleteRole(ctx context.Context, req *role.DeleteRoleRequest) (*role.Role, error) {
	r, err := s.DescribeRole(ctx, role.NewDescribeRoleRequestWithID(req.Id))
	if err != nil {
		return nil, err
	}

	if r.Spec.Type.Equal(role.RoleType_BUILDIN) {
		return nil, fmt.Errorf("build_in role can't be delete")
	}

	if !req.DeletePolicy {
		queryReq := policy.NewQueryPolicyRequest()
		queryReq.Page = request.NewPageRequest(20, 1)
		queryReq.RoleId = req.Id
		ps, err := s.policy.QueryPolicy(ctx, queryReq)
		if err != nil {
			return nil, err
		}
		if ps.Total > 0 {
			return nil, exception.NewBadRequest("该角色还关联得有策略, 请先删除关联策略")
		}
	}

	resp, err := s.role.DeleteOne(ctx, bson.M{"_id": req.Id})
	if err != nil {
		return nil, exception.NewInternalServerError("delete role(%s) error, %s", req.Id, err)
	}

	if resp.DeletedCount == 0 {
		return nil, exception.NewNotFound("role(%s) not found", req.Id)
	}

	// 清除角色关联的权限
	permReq := role.NewRemovePermissionFromRoleRequest()
	permReq.RoleId = req.Id
	permReq.RemoveAll = true
	_, err = s.RemovePermissionFromRole(ctx, permReq)
	if err != nil {
		s.log.Errorf("delete role permission error, %s", err)
	}

	// 清除角色关联的策略
	_, err = s.policy.DeletePolicy(ctx, policy.NewDeletePolicyRequestWithID(req.Id))
	if err != nil {
		s.log.Errorf("delete role policy error, %s", err)
	}

	return r, nil
}
