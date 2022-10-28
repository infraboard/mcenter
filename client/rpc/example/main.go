package main

import (
	"context"
	"fmt"

	"github.com/infraboard/mcenter/client/rpc"
)

func main() {
	// 提前加载好 mcenter客户端
	conf := rpc.NewDefaultConfig()
	conf.Address = "mcenter grpc address"
	conf.WithCredentials("mcenter client_id", "mcenter client_secret")
	err := rpc.LoadClientFromConfig(conf)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	// 调用rpc方法
	ci, err := rpc.C().ClientInfo(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(ci)
}
