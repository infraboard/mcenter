package provider

import (
	"context"

	"github.com/infraboard/mcenter/apps/token"
)

var (
	// m is a map from scheme to issuer.
	m = make(map[token.GRANT_TYPE]Issuer)
)

// 注册令牌颁发器
func Registe(i Issuer) {
	m[i.GrantType()] = i
}

func Get(gt token.GRANT_TYPE) Issuer {
	if v, ok := m[gt]; ok {
		return v
	}
	return nil
}

// 令牌颁发器
type Issuer interface {
	GrantType() token.GRANT_TYPE
	IssueToken(context.Context, *token.IssueTokenRequest) (*token.Token, error)
}
