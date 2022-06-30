package main

import (
	"context"
	"time"

	"github.com/infraboard/mcenter/apps/instance"
	"github.com/infraboard/mcenter/client/rpc"
)

func main() {
	// keyauth客户端配置
	conf := rpc.NewDefaultConfig()
	conf.Address = "127.0.0.1:18010"
	conf.ClientID = "8yJ3p8XSaNhwmnyNV4V413LK"
	conf.ClientSecret = "8rISq5gwdCQGZWEEHxsLa5qfx0eWYcjo"

	// 提前加载好 mcenter客户端
	err := rpc.LoadClientFromConfig(conf)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// 注册服务实例
	req := instance.NewRegistryRequest()
	req.Address = "127.0.0.1:18050"
	ins, err := rpc.C().Instance().RegistryInstance(ctx, req)
	if err != nil {
		panic(err)
	}

	// 上报实例心跳
	time.Sleep(15 * time.Second)

	// 注销实例
	rpc.C().Instance().UnRegistryInstance(context.Background(), instance.NewUnregistryRequest(ins.Id))
}
