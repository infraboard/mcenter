package ldap_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/infraboard/mcenter/apps/domain"
	"github.com/infraboard/mcenter/apps/token/provider/ldap"
)

var (
	p *ldap.Provider
)

func TestCheckUserPassword(t *testing.T) {
	should := assert.New(t)

	u, err := p.CheckUserPassword("oldfish", "123456")
	if should.NoError(err) {
		t.Log(u)
	}
}

func TestUserDetail(t *testing.T) {
	should := assert.New(t)

	ud, err := p.GetDetails("oldfish")
	if should.NoError(err) {
		t.Log(ud)
	}

}

func TestGetBaseDNFromUser(t *testing.T) {
	should := assert.New(t)

	conf := domain.NewDefaultLDAPConfig()
	conf.BindDn = "cn=admin,dc=example,dc=org"
	baseDN := conf.GetBaseDNFromUser()

	should.Equal("dc=example,dc=org", baseDN)
}

func TestBaseDnToSuffix(t *testing.T) {
	should := assert.New(t)

	conf := domain.NewDefaultLDAPConfig()
	conf.BaseDn = "dc=example,dc=org"
	baseDN := conf.BaseDnToSuffix()

	should.Equal("example.org", baseDN)
}

func init() {
	conf := domain.NewDefaultLDAPConfig()
	conf.Url = "ldap://127.0.0.1:389"
	conf.BindDn = "cn=admin,dc=example,dc=org"
	conf.BindPassword = "admin"
	conf.BaseDn = "dc=example,dc=org"
	conf.UserFilter = "(uid={input})"

	p = ldap.NewProvider(conf)
	err := p.CheckConnect()
	if err != nil {
		panic(err)
	}
}
