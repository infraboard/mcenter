package domain

import context "context"

const (
	AppName = "domain"
)

type Service interface {
	// 创建域
	CreateDomain(context.Context, *CreateDomainRequest) (*Domain, error)
	// 更新域
	UpdateDomain(context.Context, *UpdateDomainRequest) (*Domain, error)
	// RPC
	RPCServer
}
