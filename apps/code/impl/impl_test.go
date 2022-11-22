package impl_test

import (
	"context"
	"testing"

	"github.com/infraboard/mcenter/apps/code"
	"github.com/infraboard/mcenter/test/tools"
	"github.com/infraboard/mcube/app"
)

var (
	impl code.Service
	ctx  = context.Background()
)

func TestIssueCode(t *testing.T) {
	req := code.NewIssueCodeRequest()
	r, err := impl.IssueCode(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}

func init() {
	tools.DevelopmentSetup()
	impl = app.GetInternalApp(code.AppName).(code.Service)
}
