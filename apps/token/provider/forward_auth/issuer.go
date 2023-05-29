package forward_auth

import (
	"context"

	"github.com/infraboard/mcenter/apps/code"
	"github.com/infraboard/mcenter/apps/domain"
	"github.com/infraboard/mcenter/apps/token"
	"github.com/infraboard/mcenter/apps/token/provider"
	"github.com/infraboard/mcenter/apps/user"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

var (
	AUTH_FAILED = exception.NewUnauthorized("user or password not connrect")
)

type issuer struct {
	user   user.Service
	domain domain.Service

	log logger.Logger
}

func (i *issuer) Init() error {
	i.user = app.GetInternalApp(user.AppName).(user.Service)
	i.domain = app.GetInternalApp(domain.AppName).(domain.Service)
	i.log = zap.L().Named("issuer.forward_auth")
	return nil
}

func (i *issuer) GrantType() token.GRANT_TYPE {
	return token.GRANT_TYPE_FORWARD_AUTH
}

func (i *issuer) validate(ctx context.Context, username, pass string) (*user.User, error) {
	return nil, nil
}

func (i *issuer) IssueToken(ctx context.Context, req *token.IssueTokenRequest) (*token.Token, error) {
	u, err := i.validate(ctx, req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	// 3. 颁发Token
	tk := token.NewToken(req)
	tk.Domain = u.Spec.Domain
	tk.Username = u.Spec.Username
	tk.UserType = u.Spec.Type
	tk.UserId = u.Meta.Id
	return tk, nil
}

func (i *issuer) IssueCode(ctx context.Context, req *code.IssueCodeRequest) (*code.Code, error) {
	_, err := i.validate(ctx, req.Username, req.Password)
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
