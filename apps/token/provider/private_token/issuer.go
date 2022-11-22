package private_token

import (
	"context"

	"github.com/infraboard/mcenter/apps/token"
	"github.com/infraboard/mcenter/apps/token/provider"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

type issuer struct {
	token token.Service

	log logger.Logger
}

func (i *issuer) Init() error {
	i.token = app.GetInternalApp(token.AppName).(token.Service)
	i.log = zap.L().Named("issuer.privat_token")
	return nil
}

func (i *issuer) GrantType() token.GRANT_TYPE {
	return token.GRANT_TYPE_PRIVATE_TOKEN
}

func (i *issuer) IssueToken(ctx context.Context, req *token.IssueTokenRequest) (*token.Token, error) {
	if !req.GrantType.Equal(token.GRANT_TYPE_PRIVATE_TOKEN) {
		return nil, exception.NewBadRequest("private token issuer is only for %s", token.GRANT_TYPE_PRIVATE_TOKEN)
	}

	if req.AccessToken == "" {
		return nil, exception.NewUnauthorized("access token required")
	}

	// 判断凭证合法性
	tk, err := i.token.ValidateToken(ctx, token.NewValidateTokenRequest(req.AccessToken))
	if err != nil {
		return nil, err
	}

	// 3. 颁发Token
	newTk := token.NewToken(req)
	newTk.Domain = tk.Domain
	newTk.Username = tk.Username
	newTk.UserType = tk.UserType
	newTk.UserId = tk.UserId

	return newTk, nil
}

func init() {
	provider.Registe(&issuer{})
}
