package impl_test

import (
	"context"
	"os"
	"testing"

	"github.com/infraboard/mcube/app"
	"google.golang.org/grpc/metadata"

	"github.com/infraboard/mcenter/apps/instance"
	"github.com/infraboard/mcenter/apps/service"
	"github.com/infraboard/mcenter/test/tools"
)

var (
	impl instance.Service
	ctx  = context.Background()
)

func TestSearch(t *testing.T) {
	req := instance.NewSearchRequest()
	set, err := impl.Search(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}

func TestRegistry(t *testing.T) {
	req := instance.NewRegistryRequest()
	req.Name = "keyauth-001"
	req.Address = "127.0.0.1:18050"
	md := metadata.Pairs(service.ClientHeaderKey, os.Getenv("MCENTER_CLINET_ID"))
	ctx := metadata.NewIncomingContext(ctx, md)
	ins, err := impl.RegistryInstance(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func init() {
	tools.DevelopmentSetup()
	impl = app.GetInternalApp(instance.AppName).(instance.Service)
}
