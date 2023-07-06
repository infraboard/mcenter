package role

import context "context"

const (
	AppName = "role"
)

type Service interface {
	CreateRole(context.Context, *CreateRoleRequest) (*Role, error)
	DeleteRole(context.Context, *DeleteRoleRequest) (*Role, error)
	AddPermissionToRole(context.Context, *AddPermissionToRoleRequest) (*Role, error)
	RemovePermissionFromRole(context.Context, *RemovePermissionFromRoleRequest) (*Role, error)
	UpdatePermission(context.Context, *UpdatePermissionRequest) (*Role, error)
	RPCServer
}
