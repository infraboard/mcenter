package impl_test

import (
	"context"
	"testing"

	"github.com/infraboard/mcenter/apps/gateway"
	"github.com/infraboard/mcenter/test/tools"
	"github.com/infraboard/mcube/app"
)

var (
	impl gateway.Service
	ctx  = context.Background()
)

func TestCreateGateway(t *testing.T) {
	req := gateway.NewCreateGatewayRequest()
	ins, err := impl.CreateGateway(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func init() {
	tools.DevelopmentSetup()
	impl = app.GetInternalApp(gateway.AppName).(gateway.Service)
}
