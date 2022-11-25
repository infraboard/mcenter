package ldap_test

import (
	"testing"

	"github.com/infraboard/mcube/logger/zap"
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

	conf := domain.NewDefaultConfig()
	conf.AdminUsername = "cn=admin,dc=example,dc=org"
	baseDN := conf.GetBaseDNFromUser()

	should.Equal("dc=example,dc=org", baseDN)
}

func TestBaseDnToSuffix(t *testing.T) {
	should := assert.New(t)

	conf := domain.NewDefaultConfig()
	conf.BaseDn = "dc=example,dc=org"
	baseDN := conf.BaseDnToSuffix()

	should.Equal("example.org", baseDN)
}

func init() {
	zap.DevelopmentSetup()

	conf := domain.NewDefaultConfig()
	conf.Url = "ldap://127.0.0.1:389"
	conf.AdminUsername = "cn=admin,dc=example,dc=org"
	conf.AdminPassword = "admin"
	conf.BaseDn = "dc=example,dc=org"
	conf.UsersFilter = "(uid={input})"

	p = ldap.NewProvider(conf)
	err := p.CheckConnect()
	if err != nil {
		panic(err)
	}
}
