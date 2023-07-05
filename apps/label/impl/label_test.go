package impl_test

import (
	"testing"

	"github.com/infraboard/mcenter/apps/label"
	"github.com/infraboard/mcube/pb/resource"
)

func TestCreateLabel(t *testing.T) {
	req := label.NewCreateLabelRequest()
	req.Visiable = resource.VISIABLE_GLOBAL
	req.Key = "环境"
	req.Required = true
	req.ValueType = label.VALUE_TYPE_ENUM
	req.Visiable = resource.VISIABLE_GLOBAL
	req.DefaultValue = "开发"
	req.AddEnumOption(
		label.NewEnumOption("生产"),
		label.NewEnumOption("测试"),
		label.NewEnumOption("开发"),
	)

	r, err := impl.CreateLabel(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r.ToJSON())
}

func TestQueryLabel(t *testing.T) {
	req := label.NewQueryLabelRequest()
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
