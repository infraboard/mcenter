package refresh

import (
	"context"

	"github.com/infraboard/mcenter/apps/token"
	"github.com/infraboard/mcenter/apps/token/provider"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

type issuer struct {
	token token.Service

	log logger.Logger
}

func (i *issuer) Init() error {
	i.token = app.GetInternalApp(token.AppName).(token.Service)
	i.log = zap.L().Named("issuer.refresh")
	return nil
}

func (i *issuer) GrantType() token.GRANT_TYPE {
	return token.GRANT_TYPE_REFRESH
}

func (i *issuer) IssueToken(ctx context.Context, req *token.IssueTokenRequest) (*token.Token, error) {
	return nil, nil
}

func init() {
	provider.Registe(&issuer{})
}
