package token

import "github.com/infraboard/mcenter/apps/domain"

const (
	AppName = "token"
)

// NewIssueTokenRequest 默认请求
func NewIssueTokenRequest() *IssueTokenRequest {
	return &IssueTokenRequest{}
}

func (req *IssueTokenRequest) GetDomainWithDefault() string {
	if req.Domain != "" {
		return req.Domain
	}

	return domain.DEFAULT_DOMAIN
}

// NewRevolkTokenRequest 撤销Token请求
func NewRevolkTokenRequest() *RevolkTokenRequest {
	return &RevolkTokenRequest{}
}

func NewChangeNamespaceRequest() *ChangeNamespaceRequest {
	return &ChangeNamespaceRequest{}
}
