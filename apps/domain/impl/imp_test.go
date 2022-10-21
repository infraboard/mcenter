package impl_test

import (
	"context"
	"testing"

	"github.com/infraboard/mcenter/apps/domain"
	"github.com/infraboard/mcenter/test/tools"
	"github.com/infraboard/mcube/app"
)

var (
	impl domain.Service
	ctx  = context.Background()
)

func TestCreateRole(t *testing.T) {
	req := domain.NewCreateDomainRequest()
	r, err := impl.CreateDomain(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}

func init() {
	tools.DevelopmentSetup()
	impl = app.GetInternalApp(domain.AppName).(domain.Service)
}
