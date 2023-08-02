package label_test

import (
	"testing"

	"github.com/infraboard/mcenter/apps/label"
)

func TestBuildLabelValue(t *testing.T) {
	req := label.NewCreateLabelRequest()
	req.Key = "UserGroup"
	req.AddEnumOption(
		label.NewEnumOption("a").
			AddChildren(label.NewEnumOption("b"), label.NewEnumOption("c")),
		label.NewEnumOption("x").
			AddChildren(label.NewEnumOption("y"), label.NewEnumOption("z")),
	)
	req.BuildValue()
	t.Log(req)
}
