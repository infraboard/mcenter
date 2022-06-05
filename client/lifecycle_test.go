package client_test

import (
	"context"
	"testing"
	"time"

	"github.com/infraboard/mcenter/apps/instance"
	"github.com/infraboard/mcenter/client"
)

func TestLifecycle(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

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
	// 注销实例
	lf.UnRegistry(context.Background())
}

func init() {
	// 提前加载好 mcenter客户端, resolver需要使用
	err := client.LoadClientFromEnv()
	if err != nil {
		panic(err)
	}
}
