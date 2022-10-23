package impl_test

import (
	"context"
	"testing"

	"github.com/infraboard/mcenter/apps/domain"
	"github.com/infraboard/mcenter/apps/role"
	"github.com/infraboard/mcenter/test/tools"
	"github.com/infraboard/mcube/app"
)

var (
	impl role.Service
	ctx  = context.Background()
)

func TestCreateRole(t *testing.T) {
	req := role.NewCreateRoleRequest()
	req.CreateBy = "test"
	req.Domain = domain.DEFAULT_DOMAIN
	req.Type = role.RoleType_GLOBAL
	req.Name = "test_role"
	req.Description = "测试"
	req.Specs = []*role.Spec{
		{
			Desc:         "测试",
			Effect:       role.EffectType_ALLOW,
			ServiceId:    "xxx",
			ResourceName: "ECS",
			LabelKey:     "action",
			LabelValues:  []string{"create", "list"},
		},
	}
	r, err := impl.CreateRole(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}

func TestCreateAdminRole(t *testing.T) {
	req := role.NewCreateRoleRequest()
	req.CreateBy = "admin"
	req.Domain = domain.DEFAULT_DOMAIN
	req.Type = role.RoleType_GLOBAL
	req.Name = role.ADMIN_ROLE_NAME
	req.Description = "管理员角色"
	req.Specs = []*role.Spec{
		{
			Desc:         "管理员具备所有权限",
			Effect:       role.EffectType_ALLOW,
			ServiceId:    "*",
			ResourceName: "*",
			MatchAll:     true,
		},
	}
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

func init() {
	tools.DevelopmentSetup()
	impl = app.GetInternalApp(role.AppName).(role.Service)
}
