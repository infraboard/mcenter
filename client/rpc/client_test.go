package rpc_test

import (
	"context"
	"os"
	"testing"

	"github.com/infraboard/mcenter/apps/endpoint"
	"github.com/infraboard/mcenter/apps/instance"
	"github.com/infraboard/mcenter/apps/token"
	"github.com/infraboard/mcenter/client/rpc"
)

var (
	c   *rpc.ClientSet
	ctx = context.Background()
)

func TestValidateToken(t *testing.T) {
	req := token.NewValidateTokenRequest("A2OGh9ORR8BSjY9juj898W9Z")
	tk, err := c.Token().ValidateToken(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(tk)
}

func TestEndpointRegistry(t *testing.T) {
	req := endpoint.NewRegistryRequest("0.1", []*endpoint.Entry{
		{
			FunctionName:     "test",
			Path:             "POST./maudit/api/v1/books",
			Method:           "POST",
			Resource:         "Book",
			AuthEnable:       true,
			PermissionEnable: true,
			Labels:           map[string]string{"action": "create"},
		},
	})
	req.ClientId = os.Getenv("MCENTER_CLINET_ID")
	req.ClientSecret = os.Getenv("MCENTER_CLIENT_SECRET")
	resp, err := c.Endpoint().RegistryEndpoint(ctx, req)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(resp)
}

func TestInstanceRegistry(t *testing.T) {
	req := instance.NewRegistryRequest()
	req.Name = "maudit_test_01"
	req.Address = "127.0.0.1:1100"
	resp, err := c.Instance().RegistryInstance(ctx, req)

	if err != nil {
		t.Fatal(err)
	}

	t.Log(resp)
}

func TestInstanceSearch(t *testing.T) {
	req := instance.NewSearchRequest()
	req.ServiceName = "maudit"
	resp, err := c.Instance().Search(ctx, req)

	if err != nil {
		t.Fatal(err)
	}

	for i := range resp.Items {
		t.Log(resp.Items[i])
	}
}

func init() {
	err := rpc.LoadClientFromEnv()
	if err != nil {
		panic(err)
	}
	c = rpc.C()
}
