package impl_test

import (
	"testing"

	"github.com/infraboard/mcenter/apps/notify"
	"github.com/infraboard/mcenter/test/tools"
)

func TestSendNotify(t *testing.T) {
	req := notify.NewSendNotifyRequest()
	req.NotifyTye = notify.NOTIFY_TYPE_IM
	req.AddUser("admin")
	req.Title = "test"
	req.Content = "test content"
	set, err := impl.SendNotify(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(set))
}
