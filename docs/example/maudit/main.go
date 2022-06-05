package main

import (
	"context"
	"time"

	"github.com/infraboard/mcenter/apps/instance"
	"github.com/infraboard/mcenter/client"
)

func main() {
	// maudit客户端配置
	conf := client.NewDefaultConfig()
	conf.Address = "127.0.0.1:18010"
	conf.ClientID = "LRmqB9tQ0VLf0v1lpwzJnypX"
	conf.ClientSecret = "YRBi946B4wqqDHrfSWa4CecIDTzlBQ6N"

	// 提前加载好 mcenter客户端
	err := client.LoadClientFromConfig(conf)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// 注册服务实例
	req := instance.NewRegistryRequest()
	req.Address = "127.0.0.1:18050"
	lf, err := client.C().Registry(ctx, req)
	if err != nil {
		panic(err)
	}

	// 上报实例心跳
	lf.Heartbeat(ctx)
	time.Sleep(15 * time.Second)

	// 注销实例
	lf.UnRegistry(context.Background())
}
