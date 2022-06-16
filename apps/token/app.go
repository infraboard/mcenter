package token

const (
	AppName = "token"
)

// NewIssueTokenRequest 默认请求
func NewIssueTokenRequest() *IssueTokenRequest {
	return &IssueTokenRequest{}
}

// NewRevolkTokenRequest 撤销Token请求
func NewRevolkTokenRequest() *RevolkTokenRequest {
	return &RevolkTokenRequest{}
}

func NewChangeNamespaceRequest() *ChangeNamespaceRequest {
	return &ChangeNamespaceRequest{}
}
