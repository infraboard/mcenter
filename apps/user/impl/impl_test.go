package impl_test

import (
	"context"
	"testing"

	"github.com/infraboard/mcenter/apps/user"
	"github.com/infraboard/mcenter/test/tools"
	"github.com/infraboard/mcube/app"
)

var (
	impl user.Service
	ctx  = context.Background()
)

func TestCreateRole(t *testing.T) {
	req := user.NewCreateUserRequest()
	r, err := impl.CreateUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}

func init() {
	tools.DevelopmentSetup()
	impl = app.GetInternalApp(user.AppName).(user.Service)
}
