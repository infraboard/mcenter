package rest

import (
	"context"
	"fmt"

	"github.com/infraboard/mcenter/apps/token"
	"github.com/infraboard/mcube/client/rest"
	"github.com/infraboard/mcube/http/response"
)

type TokenService interface {
	// 校验Token
	ValidateToken(context.Context, *token.ValidateTokenRequest) (*token.Token, error)
}

type tokenImpl struct {
	client *rest.RESTClient
}

func (i *tokenImpl) ValidateToken(ctx context.Context, req *token.ValidateTokenRequest) (*token.Token, error) {
	ins := token.NewDefaultToken()
	resp := response.NewData(ins)

	fmt.Println("bearer " + req.AccessToken)
	err := i.client.
		Get("token").
		Header(token.VALIDATE_TOKEN_HEADER_KEY, req.AccessToken).
		Do(ctx).
		Into(resp)
	if err != nil {
		return nil, err
	}

	if resp.Error() != nil {
		return nil, err
	}

	return ins, nil
}
