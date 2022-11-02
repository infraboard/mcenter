package config_test

import (
	"testing"

	"github.com/infraboard/mcenter/apps/gateway/provider/traefik/config"
	"github.com/infraboard/mcenter/apps/gateway/provider/traefik/kv"
)

func TestXxx(t *testing.T) {
	svr := config.Service{
		LoadBalancer: &config.ServersLoadBalancer{
			Servers: []config.Server{
				{URL: "127.0.0.1:8010"},
				{URL: "127.0.0.1:8020"},
			},
		},
	}
	kv.GetElementKvs("traefik", svr)
	// for i := range kvs.Items {
	// 	item := kvs.Items[i]
	// 	t.Log(item.Key, string(item.Value))
	// }
}
