package provider

import (
	"context"
	"fmt"

	"github.com/infraboard/mcenter/apps/token"
)

var (
	// m is a map from scheme to issuer.
	m = make(map[token.GRANT_TYPE]Issuer)
)

// 令牌颁发器
type Issuer interface {
	Init() error
	GrantType() token.GRANT_TYPE
	TokenIssuer
}

type TokenIssuer interface {
	IssueToken(context.Context, *token.IssueTokenRequest) (*token.Token, error)
}

// 注册令牌颁发器
func Registe(i Issuer) {
	m[i.GrantType()] = i
}

func Get(gt token.GRANT_TYPE) TokenIssuer {
	if v, ok := m[gt]; ok {
		return v
	}

	return nil
}

func Init() error {
	for k, v := range m {
		if err := v.Init(); err != nil {
			return fmt.Errorf("init %s issuer error", k)
		}
	}

	return nil
}
