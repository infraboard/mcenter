package impl_test

import (
	"testing"

	"github.com/infraboard/mcenter/apps/domain"
	"github.com/infraboard/mcenter/apps/role"
)

func TestCreateRole(t *testing.T) {
	req := role.NewCreateRoleRequest()
	req.CreateBy = "admin"
	req.Domain = domain.DEFAULT_DOMAIN
	req.Global = true
	req.Name = "developer"
	req.Description = "开发"
	req.Specs = []*role.Spec{
		{
			Desc:         "mpaas只读权限",
			Effect:       role.EffectType_ALLOW,
			ServiceId:    "cd08fc9c",
			ResourceName: "*",
			LabelKey:     "action",
			LabelValues:  []string{"list", "get"},
		},
		{
			Desc:         "构建配置管理权限",
			Effect:       role.EffectType_ALLOW,
			ServiceId:    "cd08fc9c",
			ResourceName: "builds",
			LabelKey:     "action",
			MatchAll:     true,
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
	req.Domain = domain.DEFAULT_DOMAIN
	req.WithPermission = true
	r, err := impl.QueryRole(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
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
	req := role.NewDescribeRoleRequestWithID("cd9ncsmv9mc17sg8rr90")
	r, err := impl.DescribeRole(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}
