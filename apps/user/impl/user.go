package impl

import (
	"context"

	"github.com/infraboard/mcenter/apps/user"
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

// 删除用户
func (s *service) DeleteUser(ctx context.Context, req *user.DeleteUserRequest) (*user.UserSet, error) {
	return nil, nil
}

// 修改用户信息
func (s *service) UpdateUser(ctx context.Context, req *user.UpdateUserRequest) (*user.User, error) {
	return nil, nil
}

// 修改用户密码, 用户需要知道原先密码
func (s *service) UpdatePassword(ctx context.Context, req *user.UpdatePasswordRequest) (*user.Password, error) {
	return nil, nil
}

// 重置密码, 无需知道原先密码, 主账号执行
func (s *service) ResetPassword(ctx context.Context, req *user.ResetPasswordRequest) (*user.Password, error) {
	return nil, nil
}
