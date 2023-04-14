package user

import context "context"

const (
	AppName = "user"
)

type Service interface {
	// 创建用户
	CreateUser(context.Context, *CreateUserRequest) (*User, error)
	// 删除用户
	DeleteUser(context.Context, *DeleteUserRequest) (*UserSet, error)
	// 修改用户信息
	UpdateUser(context.Context, *UpdateUserRequest) (*User, error)
	// 修改用户密码, 用户需要知道原先密码
	UpdatePassword(context.Context, *UpdatePasswordRequest) (*Password, error)
	// 重置密码, 无需知道原先密码, 主账号执行
	ResetPassword(context.Context, *ResetPasswordRequest) (*Password, error)
	// RPC服务
	RPCServer
}
