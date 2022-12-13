package provider

import (
	"context"
	"fmt"

	"github.com/infraboard/mcenter/apps/code"
	"github.com/infraboard/mcenter/apps/token"
)

var (
	// m is a map from scheme to issuer.
	m = make(map[token.GRANT_TYPE]Issuer)
)

// 颁发器, 可以颁发Token或者验证码
type Issuer interface {
	Init() error
	GrantType() token.GRANT_TYPE
	TokenIssuer
	CodeIssuer
}

// 访问令牌颁发器
type TokenIssuer interface {
	IssueToken(context.Context, *token.IssueTokenRequest) (*token.Token, error)
}

// 验证码颁发器
type CodeIssuer interface {
	IssueCode(context.Context, *code.IssueCodeRequest) (*code.Code, error)
}

// 注册令牌颁发器
func Registe(i Issuer) {
	m[i.GrantType()] = i
}

func GetTokenIssuer(gt token.GRANT_TYPE) TokenIssuer {
	if v, ok := m[gt]; ok {
		return v
	}

	return nil
}

func GetCodeIssuer(gt token.GRANT_TYPE) CodeIssuer {
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
