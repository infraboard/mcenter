package rpc_test

import (
	"context"
	"os"
	"testing"

	"github.com/infraboard/mcenter/apps/endpoint"
	"github.com/infraboard/mcenter/apps/instance"
	"github.com/infraboard/mcenter/apps/token"
	"github.com/infraboard/mcenter/clients/rpc"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/ioc/health"
)

var (
	c   *rpc.ClientSet
	ctx = context.Background()
)

func TestValidateToken(t *testing.T) {
	req := token.NewValidateTokenRequest("t9EE4ov1J9hap3YE6fDNxE5m")
	tk, err := c.Token().ValidateToken(ctx, req)
	if err != nil {
		if e, ok := err.(exception.APIException); ok {
			t.Fatal(e.ToJson())
		} else {
			t.Fatal(err)
		}
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
	req.ServiceId = os.Getenv("MCENTER_CLINET_ID")
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
	req.ServiceName = "cmdb"
	resp, err := c.Instance().Search(ctx, req)

	if err != nil {
		t.Fatal(err)
	}

	for i := range resp.Items {
		t.Log(resp.Items[i])
	}
}

func TestHealthCheck(t *testing.T) {
	req := health.NewHealthCheckRequest()
	resp, err := c.Health().Check(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}

func init() {
	err := rpc.LoadClientFromEnv()
	if err != nil {
		panic(err)
	}
	c = rpc.C()
}
