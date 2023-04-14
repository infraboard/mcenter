package impl_test

import (
	"testing"

	"github.com/infraboard/mcenter/apps/domain"
	"github.com/infraboard/mcenter/apps/namespace"
	"github.com/infraboard/mcenter/apps/policy"
)

func TestCreatePolicy(t *testing.T) {
	req := policy.NewCreatePolicyRequest()
	req.Username = "test"
	req.RoleId = "cdahojmv9mc70r5h74gg"
	req.Domain = domain.DEFAULT_DOMAIN
	req.Namespace = namespace.DEFAULT_NAMESPACE
	req.CreateBy = "admin"
	r, err := impl.CreatePolicy(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}

func TestQueryPolicy(t *testing.T) {
	req := policy.NewQueryPolicyRequest()
	req.WithRole = true
	r, err := impl.QueryPolicy(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}
