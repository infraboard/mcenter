package impl_test

import (
	"testing"

	"github.com/infraboard/mcenter/apps/token"
	"github.com/infraboard/mcenter/test/tools"
)

func TestIssueToken(t *testing.T) {
	req := token.NewIssueTokenRequest()
	req.Username = "admin"
	req.Password = "123456"
	tk, err := impl.IssueToken(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tk.Json())
}

func TestValidateToken(t *testing.T) {
	req := token.NewValidateTokenRequest(tools.AccessToken())
	tk, err := impl.ValidateToken(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tk)
}

func TestQueryToken(t *testing.T) {
	req := token.NewQueryTokenRequest()
	set, err := impl.QueryToken(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set.JsonFormat())
}

func TestChangeNamespace(t *testing.T) {
	req := token.NewChangeNamespaceRequest()
	req.Namespace = ""
	req.Token = tools.AccessToken()
	tk, err := impl.ChangeNamespace(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tk)
}
