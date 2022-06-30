package impl_test

import (
	"context"
	"os"
	"testing"

	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc/metadata"

	"github.com/infraboard/mcenter/apps/instance"
	"github.com/infraboard/mcenter/client/rpc/auth"
	"github.com/infraboard/mcenter/conf"

	// 注册所有服务
	_ "github.com/infraboard/mcenter/apps/all"
)

var (
	impl instance.Service
)

func TestSearch(t *testing.T) {
	req := instance.NewSearchRequest()
	set, err := impl.Search(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}

func TestRegistry(t *testing.T) {
	req := instance.NewRegistryRequest()
	req.Name = "keyauth-001"
	req.Address = "127.0.0.1:18050"
	md := metadata.Pairs(auth.ClientHeaderKey, os.Getenv("MCENTER_CLINET_ID"))
	ctx := metadata.NewIncomingContext(context.Background(), md)
	ins, err := impl.RegistryInstance(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func init() {
	zap.DevelopmentSetup()

	if err := conf.LoadConfigFromEnv(); err != nil {
		panic(err)
	}

	if err := app.InitAllApp(); err != nil {
		panic(err)
	}

	impl = app.GetInternalApp(instance.AppName).(instance.Service)
}
