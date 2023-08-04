package impl_test

import (
	"testing"

	"github.com/infraboard/mcenter/apps/domain"
	"github.com/infraboard/mcenter/apps/namespace"
	"github.com/infraboard/mcenter/apps/policy"
	"github.com/infraboard/mcube/pb/resource"
)

func TestCreatePolicy(t *testing.T) {
	req := policy.NewCreatePolicyRequest()
	req.UserId = "test@default"
	req.RoleId = "bac61744"
	req.Domain = domain.DEFAULT_DOMAIN
	req.Namespace = namespace.DEFAULT_NAMESPACE
	req.AddScope(resource.NewLabelRequirement("env", "test", "prod"))
	req.CreateBy = "admin"
	r, err := impl.CreatePolicy(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r.ToJson())
}

func TestQueryPolicy(t *testing.T) {
	req := policy.NewQueryPolicyRequest()
	req.WithRole = true
	// 查询test用户在默认空间的策略
	req.Username = "test"
	req.Domain = domain.DEFAULT_DOMAIN
	req.Namespace = namespace.DEFAULT_NAMESPACE

	r, err := impl.QueryPolicy(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r.ToJson())
}
