package token

import context "context"

type Service interface {
	// 颁发Token
	IssueToken(context.Context, *IssueTokenRequest) (*Token, error)
	// 撤销Token
	RevolkToken(context.Context, *RevolkTokenRequest) (*Token, error)
	// 切换Token空间
	ChangeNamespace(context.Context, *ChangeNamespaceRequest) (*Token, error)
	// 查询Token, 用于查询Token颁发记录, 也就是登陆日志
	QueryToken(context.Context, *QueryTokenRequest) (*TokenSet, error)
	// 查询Token详情
	DescribeToken(context.Context, *DescribeTokenRequest) (*Token, error)
	// RPC
	RPCServer
}
