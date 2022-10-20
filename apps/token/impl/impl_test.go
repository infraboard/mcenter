package impl_test

import (
	"context"
	"testing"

	"github.com/infraboard/mcenter/apps/token"
	"github.com/infraboard/mcenter/test/tools"
	"github.com/infraboard/mcube/app"
)

var (
	impl token.Service
	ctx  = context.Background()
)

func TestIssueToken(t *testing.T) {
	req := token.NewIssueTokenRequest()
	tk, err := impl.IssueToken(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tk)
}

func init() {
	tools.DevelopmentSetup()
	impl = app.GetInternalApp(token.AppName).(token.Service)
}
