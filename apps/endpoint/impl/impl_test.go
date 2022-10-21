package impl_test

import (
	"context"
	"testing"

	"github.com/infraboard/mcenter/apps/endpoint"
	"github.com/infraboard/mcenter/test/tools"
	"github.com/infraboard/mcube/app"
)

var (
	impl endpoint.Service
	ctx  = context.Background()
)

func TestCreateRole(t *testing.T) {
	req := endpoint.NewRegistryRequest("v0.0.1", []*endpoint.Entry{})
	r, err := impl.RegistryEndpoint(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}

func init() {
	tools.DevelopmentSetup()
	impl = app.GetInternalApp(endpoint.AppName).(endpoint.Service)
}
