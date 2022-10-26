package resolver_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/infraboard/mcenter/client/rpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/infraboard/mcenter/client/rpc/resolver"
)

func TestResolver(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// 连接到服务
	conn, err := grpc.DialContext(
		ctx,
		// Dial to "mcenter://maudit"
		fmt.Sprintf("%s://%s", resolver.Scheme, "maudit"),
		// 认证
		grpc.WithPerRPCCredentials(rpc.NewAuthentication(os.Getenv("MCENTER_CLINET_ID"), os.Getenv("MCENTER_CLIENT_SECRET"))),
		// 不开启TLS
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		// gprc 支持的负载均衡策略: https://github.com/grpc/grpc/blob/master/doc/load-balancing.md
		grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"round_robin":{}}]}`),
		// 直到建立连接
		grpc.WithBlock(),
	)
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
}

func init() {
	// 提前加载好 mcenter客户端, resolver需要使用
	err := rpc.LoadClientFromEnv()
	if err != nil {
		panic(err)
	}
}
