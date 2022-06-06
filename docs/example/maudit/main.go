package main

import (
	"context"
	"fmt"
	"time"

	"github.com/infraboard/mcenter/client/rpc"
	"github.com/infraboard/mcenter/client/rpc/auth"
	"github.com/infraboard/mcenter/client/rpc/resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// demo客户端配置
	conf := rpc.NewDefaultConfig()
	conf.Address = "127.0.0.1:18010"
	conf.ClientID = "LRmqB9tQ0VLf0v1lpwzJnypX"
	conf.ClientSecret = "YRBi946B4wqqDHrfSWa4CecIDTzlBQ6N"

	// 提前加载好 mcenter客户端
	err := rpc.LoadClientFromConfig(conf)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// 连接到服务
	conn, err := grpc.DialContext(
		ctx,
		fmt.Sprintf("%s://%s", resolver.Scheme, "keyauth"), // Dial to "mcenter://keyauth"
		grpc.WithPerRPCCredentials(auth.NewAuthentication("abc", "1123")),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
		grpc.WithBlock(),
	)
	if err != nil {
		panic(err)
	}

	defer conn.Close()
}
