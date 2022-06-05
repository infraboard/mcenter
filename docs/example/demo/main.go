package main

import (
	"context"
	"fmt"
	"time"

	"github.com/infraboard/mcenter/client"
	"github.com/infraboard/mcenter/client/auth"
	"github.com/infraboard/mcenter/client/resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// demo客户端配置
	conf := client.NewDefaultConfig()
	conf.Address = "127.0.0.1:18010"
	conf.ClientID = "vfIOFJB6Ezr65FjiDYTHdXH5"
	conf.ClientSecret = "hSgAy3QkKFNlUsE2lv8b72DpDnONNf5T"

	// 提前加载好 mcenter客户端
	err := client.LoadClientFromConfig(conf)
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
		// grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
		grpc.WithBlock(),
	)
	if err != nil {
		panic(err)
	}

	defer conn.Close()
}
