package impl_test

import (
	"os"
	"testing"

	"github.com/infraboard/mcenter/apps/domain"
	"github.com/infraboard/mcenter/apps/user"
)

func TestCreateSupperUser(t *testing.T) {
	req := user.NewCreateUserRequest()
	req.Domain = domain.DEFAULT_DOMAIN
	req.Username = "admin"
	req.Password = "123456"
	req.Type = user.TYPE_SUPPER
	r, err := impl.CreateUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r.ToJson())
}

func TestCreateSubUser(t *testing.T) {
	req := user.NewCreateUserRequest()
	req.Domain = domain.DEFAULT_DOMAIN
	req.Username = "test"
	req.Password = "123456"
	req.Type = user.TYPE_SUB
	r, err := impl.CreateUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r.ToJson())
}

func TestQueryUser(t *testing.T) {
	req := user.NewQueryUserRequest()
	r, err := impl.QueryUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r.ToJson())
}

func TestDescribeUser(t *testing.T) {
	req := user.NewDescriptUserRequestWithId("admin@default")
	r, err := impl.DescribeUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r.ToJson())
}

func TestPatchUser(t *testing.T) {
	req := user.NewPatchUserRequest("admin@default")
	req.Profile.Phone = os.Getenv("TEST_CALL_NUMBER")
	req.Feishu.UserId = os.Getenv("FEISHU_USER_ID")
	r, err := impl.UpdateUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r.ToJson())
}
