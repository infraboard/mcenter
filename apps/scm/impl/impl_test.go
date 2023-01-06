package impl_test

import (
	"context"
	"os"
	"testing"

	"github.com/infraboard/mcenter/apps/scm"
	"github.com/infraboard/mcenter/test/tools"
	"github.com/infraboard/mcube/app"
)

var (
	impl scm.Service
	ctx  = context.Background()
)

func TestQueryProject(t *testing.T) {
	req := scm.NewQueryProjectRequest()
	req.Token = os.Getenv("GITLAB_PRIVATE_TOkEN")
	ps, err := impl.QueryProject(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ps)
}

func init() {
	tools.DevelopmentSetup()
	impl = app.GetInternalApp(scm.AppName).(scm.Service)
}
