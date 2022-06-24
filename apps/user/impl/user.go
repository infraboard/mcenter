package impl

import (
	"context"

	"github.com/infraboard/mcenter/apps/user"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/pb/request"
)

func (s *service) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.User, error) {
	u, err := user.New(req)
	if err != nil {
		return nil, err
	}

	// 如果是管理员创建的账号需要用户自己重置密码
	if req.CreateBy.IsIn(user.CREATE_BY_ADMIN) {
		u.Password.SetNeedReset("admin created user need reset when first login")
	}

	if err := s.save(ctx, u); err != nil {
		return nil, err
	}

	u.Password = nil
	return u, nil
}

// 查询用户列表
func (s *service) QueryUser(ctx context.Context, req *user.QueryUserRequest) (*user.UserSet, error) {
	return nil, nil
}

// 查询用户详情
func (s *service) DescribeUser(ctx context.Context, req *user.DescribeUserRequest) (*user.User, error) {
	return nil, nil
}

// 修改用户信息
func (s *service) UpdateUser(ctx context.Context, req *user.UpdateUserRequest) (*user.User, error) {
	ins, err := s.DescribeUser(ctx, user.NewDescriptUserRequestWithId(req.UserId))
	if err != nil {
		return nil, err
	}

	switch req.UpdateMode {
	case request.UpdateMode_PUT:
		ins.Update(req)
	case request.UpdateMode_PATCH:
		err := ins.Patch(req)
		if err != nil {
			return nil, err
		}
	}

	if err := s.update(ctx, ins); err != nil {
		return nil, err
	}

	return nil, nil
}

// 删除用户
func (s *service) DeleteUser(ctx context.Context, req *user.DeleteUserRequest) (*user.UserSet, error) {
	// 判断这些要删除的用户是否存在
	queryReq := user.NewQueryUserRequest()
	queryReq.UserIds = req.UserIds
	set, err := s.QueryUser(ctx, queryReq)
	if err != nil {
		return nil, err
	}

	noExist := []string{}
	for _, uid := range req.UserIds {
		if !set.HasUser(uid) {
			noExist = append(noExist, uid)
		}
	}
	if len(noExist) > 0 {
		return nil, exception.NewBadRequest("user %v not found", req.UserIds)
	}

	if err := s.delete(ctx, set); err != nil {
		return nil, err
	}
	return set, nil
}

// 修改用户密码, 用户需要知道原先密码
func (s *service) UpdatePassword(ctx context.Context, req *user.UpdatePasswordRequest) (*user.Password, error) {
	return nil, nil
}

// 重置密码, 无需知道原先密码, 主账号执行
func (s *service) ResetPassword(ctx context.Context, req *user.ResetPasswordRequest) (*user.Password, error) {
	return nil, nil
}
