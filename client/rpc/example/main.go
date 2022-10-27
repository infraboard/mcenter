package main

import (
	"fmt"
	"os"

	"github.com/infraboard/mcenter/client/rpc"
)

func main() {
	// 提前加载好 mcenter客户端, resolver需要使用
	err := rpc.LoadClientFromEnv()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
