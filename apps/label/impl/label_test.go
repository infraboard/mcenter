package impl_test

import (
	"testing"

	"github.com/infraboard/mcenter/apps/label"
)

func TestCreateLabel(t *testing.T) {
	req := label.NewCreateLabelRequest()
	req.Scope = label.SCOPE_STSTEM
	req.Key = "环境"
	req.Required = true
	req.ValueType = label.VALUE_TYPE_ENUM
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
