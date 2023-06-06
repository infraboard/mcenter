package refresh

import (
	"context"
	"fmt"

	"github.com/infraboard/mcenter/apps/code"
	"github.com/infraboard/mcenter/apps/token"
	"github.com/infraboard/mcenter/apps/token/provider"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

type issuer struct {
	token token.Service

	log logger.Logger
}

func (i *issuer) Init() error {
	i.token = ioc.GetController(token.AppName).(token.Service)
	i.log = zap.L().Named("issuer.refresh")
	return nil
}

func (i *issuer) GrantType() token.GRANT_TYPE {
	return token.GRANT_TYPE_REFRESH
}

func (i *issuer) IssueToken(ctx context.Context, req *token.IssueTokenRequest) (*token.Token, error) {
	if !req.GrantType.Equal(token.GRANT_TYPE_REFRESH) {
		return nil, exception.NewBadRequest("refresh issuer is only for %s", token.GRANT_TYPE_REFRESH)
	}

	if req.AccessToken == "" || req.RefreshToken == "" {
		return nil, exception.NewUnauthorized("access token and refresh token required")
	}

	// 判断颁发凭证合法性
	tk, err := i.token.DescribeToken(ctx, token.NewDescribeTokenRequest(req.AccessToken))
	if err != nil {
		return nil, err
	}

	if tk.RefreshToken != req.RefreshToken {
		return nil, fmt.Errorf("refresh token not correct")
	}
	if tk.CheckRefreshIsExpired() {
		return nil, fmt.Errorf("refresh token is expired")
	}

	// 撤销之前的Token
	revolkReq := token.NewRevolkTokenRequest(req.AccessToken, req.RefreshToken)
	_, err = i.token.RevolkToken(ctx, revolkReq)
	if err != nil {
		return nil, err
	}

	// 3. 颁发Token
	newTk := token.NewToken(req)
	// 继承之前的授权类型
	newTk.GrantType = tk.GrantType
	newTk.Domain = tk.Domain
	newTk.Username = tk.Username
	newTk.UserType = tk.UserType
	newTk.UserId = tk.UserId

	return newTk, nil
}

func (i *issuer) IssueCode(ctx context.Context, req *code.IssueCodeRequest) (*code.Code, error) {
	return nil, fmt.Errorf("refresh not support issue code")
}

func init() {
	provider.Registe(&issuer{})
}
