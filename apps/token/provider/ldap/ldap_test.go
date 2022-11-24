package ldap_test

import (
	"testing"

	"github.com/infraboard/mcube/logger/zap"
	"github.com/stretchr/testify/assert"

	"github.com/infraboard/mcenter/apps/token/provider/ldap"
)

func TestCheckConnect(t *testing.T) {
	should := assert.New(t)

	conf := ldap.NewDefaultConfig()
	conf.URL = "ldap://127.0.0.1:389"
	conf.User = "cn=admin,dc=example,dc=org"
	conf.Password = "admin"
	conf.BaseDN = "dc=example,dc=org"
	conf.UsersFilter = "(uid={input})"

	p := ldap.NewProvider(conf)
	err := p.CheckConnect()
	should.NoError(err)
}

func TestCheckUserPassword(t *testing.T) {
	should := assert.New(t)

	conf := ldap.NewDefaultConfig()
	conf.URL = "ldap://127.0.0.1:389"
	conf.User = "cn=admin,dc=example,dc=org"
	conf.Password = "admin"
	conf.BaseDN = "dc=example,dc=org"
	conf.UsersFilter = "(uid={input})"

	p := ldap.NewProvider(conf)
	ok, err := p.CheckUserPassword("oldfish", "123456")
	if should.NoError(err) {
		should.True(ok)
	}
}

func TestUserDetail(t *testing.T) {
	should := assert.New(t)

	conf := ldap.NewDefaultConfig()
	conf.URL = "ldap://127.0.0.1:389"
	conf.User = "cn=admin,dc=example,dc=org"
	conf.Password = "admin"
	conf.BaseDN = "dc=example,dc=org"

	p := ldap.NewProvider(conf)
	ud, err := p.GetDetails("oldfish")
	if should.NoError(err) {
		t.Log(ud)
	}

}

func TestGetBaseDNFromUser(t *testing.T) {
	should := assert.New(t)

	conf := ldap.NewDefaultConfig()
	conf.User = "cn=admin,dc=example,dc=org"
	baseDN := conf.GetBaseDNFromUser()

	should.Equal("dc=example,dc=org", baseDN)
}

func init() {
	zap.DevelopmentSetup()
}
