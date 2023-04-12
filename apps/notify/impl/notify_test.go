package impl_test

import (
	"testing"

	"github.com/infraboard/mcenter/apps/domain"
	"github.com/infraboard/mcenter/apps/namespace"
	"github.com/infraboard/mcenter/apps/notify"
	"github.com/infraboard/mcenter/test/tools"
)

func TestSendNotify(t *testing.T) {
	req := notify.NewSendNotifyRequest()
	req.Domain = domain.DEFAULT_DOMAIN
	req.Namespace = namespace.DEFAULT_NAMESPACE
	req.NotifyTye = notify.NOTIFY_TYPE_IM
	req.AddUser("admin")
	req.Title = "test"
	req.Content = "test content2"
	set, err := impl.SendNotify(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(set))
}

func TestQueryRecord(t *testing.T) {
	req := notify.NewQueryRecordRequest()
	set, err := impl.QueryRecord(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tools.MustToJson(set))
}
