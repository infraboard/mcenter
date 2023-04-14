package impl_test

import (
	"testing"

	"github.com/infraboard/mcenter/apps/code"
)

func TestIssueCode(t *testing.T) {
	req := code.NewIssueCodeRequest()
	req.Username = "admin"
	req.Password = "123456"

	r, err := impl.IssueCode(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}

func TestVerifyCode(t *testing.T) {
	req := code.NewVerifyCodeRequest("admin", "816241")
	r, err := impl.VerifyCode(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}
