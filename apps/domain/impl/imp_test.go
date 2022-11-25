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

func TestQueryDoamin(t *testing.T) {
	req := domain.NewQueryDomainRequest()
	ins, err := impl.QueryDoamin(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestCreateDomain(t *testing.T) {
	req := domain.NewCreateDomainRequest()
	req.Name = domain.DEFAULT_DOMAIN
	ins, err := impl.CreateDomain(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestPatchUpdateDomain(t *testing.T) {
	conf := domain.NewDefaultConfig()
	conf.Url = "ldap://127.0.0.1:389"
	conf.BindDn = "cn=admin,dc=example,dc=org"
	conf.BindPassword = "admin"
	conf.BaseDn = "dc=example,dc=org"
	conf.UserFilter = "(uid={input})"

	req := domain.NewPatchDomainRequestByName(domain.DEFAULT_DOMAIN)
	req.Spec.LdapSetting = conf
	ins, err := impl.UpdateDomain(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func init() {
	tools.DevelopmentSetup()
	impl = app.GetInternalApp(domain.AppName).(domain.Service)
}
