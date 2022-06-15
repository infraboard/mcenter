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
