package role

import context "context"

type Service interface {
	CreateRole(context.Context, *CreateRoleRequest) (*Role, error)
	DeleteRole(context.Context, *DeleteRoleRequest) (*Role, error)
	AddPermissionToRole(context.Context, *AddPermissionToRoleRequest) (*PermissionSet, error)
	RemovePermissionFromRole(context.Context, *RemovePermissionFromRoleRequest) (*PermissionSet, error)
	UpdatePermission(context.Context, *UpdatePermissionRequest) (*Permission, error)
	RPCServer
}
