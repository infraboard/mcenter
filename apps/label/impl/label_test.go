package impl_test

import (
	"testing"

	"github.com/infraboard/mcenter/apps/label"
)

func TestCreateLabel(t *testing.T) {
	req := label.NewCreateLabelRequest()
	r, err := impl.CreateLabel(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}
