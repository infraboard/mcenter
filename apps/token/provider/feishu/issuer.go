package feishu

import (
	"context"

	"github.com/infraboard/mcenter/apps/domain"
	"github.com/infraboard/mcenter/apps/token"
	"github.com/infraboard/mcenter/apps/token/provider"
)

type issuer struct {
	conf *domain.FeishuConfig
}

func (i *issuer) Init() error {
	return nil
}

func (i *issuer) GrantType() token.GRANT_TYPE {
	return token.GRANT_TYPE_FEISHU
}

func (i *issuer) IssueToken(ctx context.Context, req *token.IssueTokenRequest) (*token.Token, error) {
	return nil, nil
}

func init() {
	provider.Registe(&issuer{})
}
