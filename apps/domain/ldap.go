package domain

import (
	"fmt"
	"strings"

	"github.com/infraboard/mcube/tools/sense"
)

// NewDefaultConfig represents the default LDAP config.
func NewDefaultLDAPConfig() *LdapConfig {
	return &LdapConfig{
		MailAttribute:        "mail",
		DisplayNameAttribute: "displayName",
		GroupnameAttribute:   "cn",
		UsernameAttribute:    "uid",
		UserFilter:           "(uid={input})",
		GroupFilter:          "(|(member={dn})(uid={username})(uid={input}))",
	}
}

// GetBaseDNFromUser 从用户中获取BaseDN
func (c *LdapConfig) GetBaseDNFromUser() string {
	return strings.Join(c.getBaseDN(c.BindDn), ",")
}

func (c *LdapConfig) BaseDnToSuffix() string {
	values := []string{}
	dn := c.getBaseDN(c.BaseDn)
	for _, v := range dn {
		kvs := strings.Split(v, "=")
		if len(kvs) > 1 {
			values = append(values, kvs[1])
		}
	}

	return strings.Join(values, ".")
}

func (c *LdapConfig) getBaseDN(entry string) []string {
	baseDN := []string{}
	for _, item := range strings.Split(entry, ",") {
		if !strings.HasPrefix(item, "cn=") {
			baseDN = append(baseDN, item)
		}
	}
	return baseDN
}

func (c *LdapConfig) DeSense() {
	c.BindPassword = sense.ShortDesenser.DeSense(c.BindPassword)
}

// Validate todo
func (c *LdapConfig) Validate() error {
	if c.Url == "" {
		return fmt.Errorf("url required")
	}

	if c.BindDn == "" || c.BindPassword == "" {
		return fmt.Errorf("username and password required")
	}

	return nil
}

// Desensitize todo
func (c *LdapConfig) Desensitize() {
	c.BindPassword = ""
}
