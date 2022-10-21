package impl_test

import (
	"context"
	"testing"

	"github.com/infraboard/mcenter/apps/namespace"
	"github.com/infraboard/mcenter/test/tools"
	"github.com/infraboard/mcube/app"
)

var (
	impl namespace.Service
	ctx  = context.Background()
)

func TestCreateRole(t *testing.T) {
	req := namespace.NewCreateNamespaceRequest()
	r, err := impl.CreateNamespace(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}

func init() {
	tools.DevelopmentSetup()
	impl = app.GetInternalApp(namespace.AppName).(namespace.Service)
}
