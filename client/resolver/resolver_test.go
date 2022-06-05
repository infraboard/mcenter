package resolver_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/infraboard/mcenter/apps/instance"
	"github.com/infraboard/mcenter/client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/infraboard/mcenter/client/auth"
	"github.com/infraboard/mcenter/client/resolver"
)

func TestResolver(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// 连接到服务
	conn, err := grpc.DialContext(
		ctx,
		fmt.Sprintf("%s://%s", resolver.Scheme, "keyauth"), // Dial to "mcenter://keyauth"
		grpc.WithPerRPCCredentials(auth.NewAuthentication("abc", "1123")),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// 注册服务实例
	req := instance.NewRegistryRequest()
	req.Name = "cmdb"
	req.Address = "127.0.0.1:18050"
	lf, err := client.C().Registry(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	// 上报实例心跳
	lf.Heartbeat(ctx)
	time.Sleep(15 * time.Second)
	lf.UnRegistry(context.Background())
}

func init() {
	// 提前加载好 mcenter客户端, resolver需要使用
	err := client.LoadClientFromEnv()
	if err != nil {
		panic(err)
	}
}
