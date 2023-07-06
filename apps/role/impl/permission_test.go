package impl_test

import (
	"testing"

	"github.com/infraboard/mcenter/apps/role"
)

func TestAddPermissionToRole(t *testing.T) {
	req := role.NewAddPermissionToRoleRequest()
	req.RoleId = ""
	r, err := impl.AddPermissionToRole(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r.ToJson())
}
