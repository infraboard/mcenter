package ldap

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/infraboard/mcenter/apps/token"
	"github.com/infraboard/mcenter/apps/token/provider"
	"github.com/infraboard/mcube/exception"
)

type issuer struct {
}

func (i *issuer) GrantType() token.GRANT_TYPE {
	return token.GRANT_TYPE_LDAP
}

func (i *issuer) IssueToken(ctx context.Context, req *token.IssueTokenRequest) (*token.Token, error) {
	userName, dn, err := i.genBaseDN(req.Username)
	if err != nil {
		return nil, err
	}
	fmt.Println(userName, dn)

	return nil, nil
}

var (
	emailRE = regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9\.]+)\.([a-zA-Z0-9]+)`)
)

func (i *issuer) genBaseDN(username string) (string, string, error) {
	match := emailRE.FindAllStringSubmatch(username, -1)
	if len(match) == 0 {
		return "", "", exception.NewBadRequest("ldap user name must like username@company.com")
	}

	sub := match[0]
	if len(sub) < 4 {
		return "", "", exception.NewBadRequest("ldap user name must like username@company.com")
	}

	dns := []string{}
	for _, dn := range sub[2:] {
		dns = append(dns, "dc="+dn)
	}

	return sub[1], strings.Join(dns, ","), nil
}

func init() {
	provider.Registe(&issuer{})
}
