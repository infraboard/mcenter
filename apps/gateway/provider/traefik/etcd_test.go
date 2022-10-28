package traefik_test

import (
	"context"
	"testing"

	"github.com/infraboard/mcenter/apps/gateway"
	"github.com/infraboard/mcenter/apps/gateway/provider/traefik"
)

func TestEtcdOperator(t *testing.T) {
	conf := gateway.NewDefaultTraefikConfig()
	s, err := traefik.NewEtcdOperator(conf)
	if err != nil {
		t.Fatal(err)
	}

	keys, err := s.ListKeys(context.Background(), "/")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s", keys)
}
