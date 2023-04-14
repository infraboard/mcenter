package impl_test

import (
	"testing"

	"github.com/infraboard/mcenter/apps/endpoint"
)

func TestCreateRole(t *testing.T) {
	req := endpoint.NewRegistryRequest("v0.0.1", []*endpoint.Entry{})
	r, err := impl.RegistryEndpoint(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}
