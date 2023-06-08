# mcenter客户端


## RPC客户端

用于内部服务调用

```go
package main

import (
	"context"
	"fmt"

	"github.com/infraboard/mcenter/clients/rpc"
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
```


## REST客户端

用于外部服务调用


## CLI客户端

作为工具用于集成, 比如放置于容器内使用