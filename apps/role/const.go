package role

import "github.com/infraboard/mcenter/apps/domain"

const (
	// AdminRoleName 管理员角色名称, 系统初始化时创建
	ADMIN_ROLE_NAME = "admin"
	// VisitorRoleName 访客角色名称, 系统初始化时创建
	VISITOR_ROLE_NAME = "visitor"
)

const (
	// RoleMaxPermission 一个角色允许添加的最大权限数
	RoleMaxPermission = 500
)

func CreateAdminRoleRequest() *CreateRoleRequest {
	req := NewCreateRoleRequest()
	req.CreateBy = "admin"
	req.Domain = domain.DEFAULT_DOMAIN
	req.Type = RoleType_GLOBAL
	req.Name = ADMIN_ROLE_NAME
	req.Description = "管理员角色"
	req.Specs = []*Spec{
		{
			Desc:         "管理员具备所有权限",
			Effect:       EffectType_ALLOW,
			ServiceId:    "*",
			ResourceName: "*",
			MatchAll:     true,
		},
	}
	return req
}
