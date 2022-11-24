package ldap

import (
	"fmt"
	"strings"
)

// NewDefaultConfig represents the default LDAP config.
func NewDefaultConfig() *LdapConfig {
	return &LdapConfig{
		MailAttribute:        "mail",
		DisplayNameAttribute: "displayName",
		GroupnameAttribute:   "cn",
		UsernameAttribute:    "uid",
		UsersFilter:          "(uid={input})",
		GroupsFilter:         "(|(member={dn})(uid={username})(uid={input}))",
	}
}

// GetBaseDNFromUser 从用户中获取BaseDN
func (c *LdapConfig) GetBaseDNFromUser() string {
	baseDN := []string{}
	for _, item := range strings.Split(c.AdminUsername, ",") {
		if !strings.HasPrefix(item, "cn=") {
			baseDN = append(baseDN, item)
		}
	}

	return strings.Join(baseDN, ",")
}

// Validate todo
func (c *LdapConfig) Validate() error {
	if c.Url == "" {
		return fmt.Errorf("url required")
	}

	if c.AdminUsername == "" || c.AdminPassword == "" {
		return fmt.Errorf("username and password required")
	}

	return nil
}

// Desensitize todo
func (c *LdapConfig) Desensitize() {
	c.AdminPassword = ""
}
