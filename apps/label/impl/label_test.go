package impl_test

import (
	"testing"

	"github.com/infraboard/mcenter/apps/domain"
	"github.com/infraboard/mcenter/apps/label"
	"github.com/infraboard/mcenter/apps/namespace"
	"github.com/infraboard/mcube/pb/resource"
)

func TestCreateEnvLabel(t *testing.T) {
	req := label.NewCreateLabelRequest()
	req.Visiable = resource.VISIABLE_GLOBAL
	req.Key = "Env"
	req.KeyDesc = "环境"
	req.ValueDesc = "资源所属环境"
	req.Domain = domain.DEFAULT_DOMAIN
	req.Namespace = namespace.DEFAULT_NAMESPACE
	req.Required = true
	req.ValueType = label.VALUE_TYPE_ENUM
	req.Visiable = resource.VISIABLE_GLOBAL
	req.DefaultValue = "开发"
	req.AddEnumOption(
		label.NewEnumOption("开发"),
		label.NewEnumOption("测试"),
		label.NewEnumOption("生产"),
	)

	r, err := impl.CreateLabel(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r.ToJSON())
}

func TestCreateResourceGroupLabel(t *testing.T) {
	req := label.NewCreateLabelRequest()
	req.Visiable = resource.VISIABLE_GLOBAL
	req.Key = "ResourceGroup"
	req.KeyDesc = "资源组"
	req.ValueDesc = "资源所属组"
	req.Domain = domain.DEFAULT_DOMAIN
	req.Namespace = namespace.DEFAULT_NAMESPACE
	req.Required = true
	req.ValueType = label.VALUE_TYPE_ENUM
	req.Visiable = resource.VISIABLE_GLOBAL
	req.DefaultValue = "默认"
	req.AddEnumOption(
		label.NewEnumOption("默认"),
	)

	r, err := impl.CreateLabel(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r.ToJSON())
}

func TestCreateUserGroupLabel(t *testing.T) {
	req := label.NewCreateLabelRequest()
	req.Visiable = resource.VISIABLE_GLOBAL
	req.Key = "UserGroup"
	req.KeyDesc = "用户组"
	req.ValueDesc = "用户所属组"
	req.Domain = domain.DEFAULT_DOMAIN
	req.Namespace = namespace.DEFAULT_NAMESPACE
	req.Required = false
	req.ValueType = label.VALUE_TYPE_ENUM
	req.Visiable = resource.VISIABLE_GLOBAL
	req.DefaultValue = ""
	req.AddEnumOption(
		label.NewEnumOption("研发部").
			AddChildren(label.NewEnumOption("开发组01")).
			AddChildren(label.NewEnumOption("开发组02")),
		label.NewEnumOption("测试部").
			AddChildren(label.NewEnumOption("测试组01")).
			AddChildren(label.NewEnumOption("测试组02")),
	)

	r, err := impl.CreateLabel(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r.ToJSON())
}

func TestCreateGrepGroupLabel(t *testing.T) {
	req := label.NewCreateLabelRequest()
	req.Visiable = resource.VISIABLE_GLOBAL
	req.Key = "DeployGrepGroup"
	req.KeyDesc = "灰度组"
	req.ValueDesc = "部署所属灰度组"
	req.Domain = domain.DEFAULT_DOMAIN
	req.Namespace = namespace.DEFAULT_NAMESPACE
	req.Required = true
	req.ValueType = label.VALUE_TYPE_ENUM
	req.Visiable = resource.VISIABLE_GLOBAL
	req.DefaultValue = "default"
	req.AddEnumOption(
		label.NewEnumOption("default").SetLabel("默认组"),
	)

	r, err := impl.CreateLabel(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r.ToJSON())
}

func TestQueryLabel(t *testing.T) {
	req := label.NewQueryLabelRequest()
	req.AddKey("UserGroup")
	r, err := impl.QueryLabel(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r.ToJSON())
}

func TestDeleteLabel(t *testing.T) {
	req := label.NewDeleteLabelRequest("aeea340b076aa5cf")
	r, err := impl.DeleteLabel(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r.ToJSON())
}
