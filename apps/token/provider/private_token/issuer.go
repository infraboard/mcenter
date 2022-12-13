package private_token

import (
	"context"

	"github.com/infraboard/mcenter/apps/code"
	"github.com/infraboard/mcenter/apps/token"
	"github.com/infraboard/mcenter/apps/token/provider"
	"github.com/infraboard/mcenter/apps/user"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

type issuer struct {
	token token.Service
	user  user.Service

	log logger.Logger
}

func (i *issuer) Init() error {
	i.token = app.GetInternalApp(token.AppName).(token.Service)
	i.user = app.GetInternalApp(user.AppName).(user.Service)
	i.log = zap.L().Named("issuer.privat_token")
	return nil
}

func (i *issuer) GrantType() token.GRANT_TYPE {
	return token.GRANT_TYPE_PRIVATE_TOKEN
}

func (i *issuer) validate(ctx context.Context, accessToken string) (*user.User, error) {
	if accessToken == "" {
		return nil, exception.NewUnauthorized("access token required")
	}

	// 判断凭证合法性
	tk, err := i.token.ValidateToken(ctx, token.NewValidateTokenRequest(accessToken))
	if err != nil {
		return nil, err
	}

	u, err := i.user.DescribeUser(ctx, user.NewDescriptUserRequestWithId(tk.UserId))
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (i *issuer) IssueToken(ctx context.Context, req *token.IssueTokenRequest) (*token.Token, error) {
	u, err := i.validate(ctx, req.AccessToken)
	if err != nil {
		return nil, err
	}

	// 3. 颁发Token
	newTk := token.NewToken(req)
	newTk.Domain = u.Spec.Domain
	newTk.Username = u.Spec.Username
	newTk.UserType = u.Spec.Type
	newTk.UserId = u.Id

	return newTk, nil
}

func (i *issuer) IssueCode(ctx context.Context, req *code.IssueCodeRequest) (*code.Code, error) {
	_, err := i.validate(ctx, req.AccessToken)
	if err != nil {
		return nil, err
	}

	// 颁发Token
	c, err := code.NewCode(req)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func init() {
	provider.Registe(&issuer{})
}
