package impl_test

import (
	"testing"

	"github.com/infraboard/mcenter/apps/domain"
	"github.com/infraboard/mcenter/apps/namespace"
	"github.com/infraboard/mcenter/apps/role"
	"github.com/infraboard/mcube/v2/pb/resource"
)

func TestCreateRole(t *testing.T) {
	req := role.NewCreateRoleRequest()
	req.CreateBy = "admin"
	req.Domain = domain.DEFAULT_DOMAIN
	req.Namespace = namespace.SYSTEM_NAMESPACE
	req.Visiable = resource.VISIABLE_GLOBAL
	req.Name = "developer"
	req.Description = "服务"
	req.Specs = []*role.PermissionSpec{
		{
			Desc:         "mpaas只读权限",
			Effect:       role.EffectType_ALLOW,
			ServiceId:    "cd08fc9c",
			ResourceName: "*",
			LabelKey:     "action",
			LabelValues:  []string{"list", "get"},
		},
		{
			Desc:         "构建配置权限",
			Effect:       role.EffectType_ALLOW,
			ServiceId:    "cd08fc9c",
			ResourceName: "builds",
			LabelKey:     "action",
			MatchAll:     true,
		},
		{
			Desc:         "部署权限",
			Effect:       role.EffectType_ALLOW,
			ServiceId:    "cd08fc9c",
			ResourceName: "deploys",
			LabelKey:     "action",
			MatchAll:     true,
		},
		{
			Desc:         "Gitlab事件模拟",
			Effect:       role.EffectType_ALLOW,
			ServiceId:    "cd08fc9c",
			ResourceName: "triggers",
			LabelKey:     "action",
			LabelValues:  []string{"create"},
		},
	}
	r, err := impl.CreateRole(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}

func TestCreateAdminRole(t *testing.T) {
	req := role.CreateAdminRoleRequest("admin")
	r, err := impl.CreateRole(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}

func TestQueryRole(t *testing.T) {
	req := role.NewQueryRoleRequest()
	req.Scope.Domain = domain.DEFAULT_DOMAIN
	req.WithPermission = true
	r, err := impl.QueryRole(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r.ToJson())
}

func TestDescribeRoleWithName(t *testing.T) {
	req := role.NewDescribeRoleRequestWithName(role.ADMIN_ROLE_NAME)
	r, err := impl.DescribeRole(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}

func TestDescribeRoleWithId(t *testing.T) {
	req := role.NewDescribeRoleRequestWithID("cgtlcjts99bkv5rhgc10")
	r, err := impl.DescribeRole(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}

func TestDeleteRole(t *testing.T) {
	req := role.NewDeleteRoleWithID("cgtlcjts99bkv5rhgc10")
	r, err := impl.DeleteRole(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}
