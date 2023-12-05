package main

import (
	"context"
	"fmt"

	"github.com/infraboard/mcenter/clients/rpc"
	"github.com/infraboard/mcube/v2/ioc"
)

func main() {
	// 加载好Ioc对象配置
	err := ioc.ConfigIocObject(ioc.NewLoadConfigRequest())
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
