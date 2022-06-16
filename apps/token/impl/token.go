package impl

import (
	"context"
	"fmt"

	"github.com/infraboard/mcenter/apps/token"
	"github.com/infraboard/mcenter/apps/user"
	"github.com/infraboard/mcube/exception"
)

func (s *service) IssueToken(ctx context.Context, req *token.IssueTokenRequest) (
	*token.Token, error) {
	switch req.GrantType {
	case token.GRANT_TYPE_PASSWORD:
		if req.Username == "" || req.Password == "" {
			return nil, exception.NewBadRequest("username and password required")
		}
		u, err := s.user.DescribeUser(ctx, user.NewDescriptUserRequestWithName(req.GetDomainWithDefault(), req.Username))
		if err != nil {
			return nil, err
		}
		fmt.Println(u)
	default:
		return nil, exception.NewBadRequest("grant type %s not implemented", req.GrantType)
	}
	return nil, nil
}

// 撤销Token
func (s *service) RevolkToken(ctx context.Context, req *token.RevolkTokenRequest) (
	*token.Token, error) {
	return nil, nil
}

// 切换Token空间
func (s *service) ChangeNamespace(ctx context.Context, req *token.ChangeNamespaceRequest) (
	*token.Token, error) {
	return nil, nil
}

// 校验Token
func (s *service) ValidateToken(ctx context.Context, req *token.ValidateTokenRequest) (
	*token.Token, error) {
	return nil, nil
}
