package wx_test

import (
	"context"
	"os"
	"testing"

	"github.com/infraboard/mcenter/apps/token"
	"github.com/infraboard/mcenter/apps/token/provider"
	"github.com/infraboard/mcenter/test/tools"
)

var (
	impl provider.TokenIssuer
	ctx  = context.Background()
)

func TestIssueToken(t *testing.T) {
	req := token.NewWechatAuthCodeIssueTokenRequest(os.Getenv("WECHAT_AUTH_CODE"), "11111", "")
	tk, err := impl.IssueToken(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tk.JsonFormat())
}

func init() {
	tools.DevelopmentSetup()
	impl = provider.Get(token.GRANT_TYPE_WECHAT_WORK)
}
