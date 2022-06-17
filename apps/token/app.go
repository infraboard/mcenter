package token

import (
	"time"

	"github.com/infraboard/mcenter/common/tools"
	"github.com/infraboard/mcube/http/request"
)

const (
	AppName = "token"
)

// NewIssueTokenRequest 默认请求
func NewIssueTokenRequest() *IssueTokenRequest {
	return &IssueTokenRequest{}
}

// AbnormalUserCheckKey todo
func (req *IssueTokenRequest) AbnormalUserCheckKey() string {
	key := ""
	switch req.GrantType {
	case GRANT_TYPE_PASSWORD, GRANT_TYPE_LDAP:
		key = req.Username
	case GRANT_TYPE_PRIVATE_TOKEN:
		key = req.AccessToken
	case GRANT_TYPE_REFRESH:
		key = req.RefreshToken
	case GRANT_TYPE_AUTH_CODE:
		key = req.AuthCode
	}
	return "abnormal_" + key
}

// NewRevolkTokenRequest 撤销Token请求
func NewRevolkTokenRequest() *RevolkTokenRequest {
	return &RevolkTokenRequest{}
}

func NewChangeNamespaceRequest() *ChangeNamespaceRequest {
	return &ChangeNamespaceRequest{}
}

func NewToken(req *IssueTokenRequest) *Token {
	tk := &Token{
		AccessToken:      tools.MakeBearer(24),
		RefreshToken:     tools.MakeBearer(32),
		IssueAt:          time.Now().UnixMilli(),
		AccessExpiredAt:  req.ExpiredAt,
		RefreshExpiredAt: req.ExpiredAt * 4,
		GrantType:        req.GrantType,
		Type:             req.Type,
		Description:      req.Description,
		Status:           NewStatus(),
		Location:         req.Location,
	}
	switch req.GrantType {
	case GRANT_TYPE_PRIVATE_TOKEN:
		tk.Platform = PLATFORM_API
	default:
		tk.Platform = PLATFORM_WEB
	}
	return tk
}

func NewStatus() *Status {
	return &Status{
		IsBlock: false,
	}
}

func NewPlatform(p PLATFORM) *PLATFORM {
	return &p
}

// 查询用户最近一次登陆记录
func NewQueryUserWebLastToken(uid string) *QueryTokenRequest {
	return &QueryTokenRequest{
		Page:     request.NewPageRequest(1, 1),
		Platform: NewPlatform(PLATFORM_WEB),
		UserId:   uid,
	}
}

func (s *TokenSet) Length() int {
	return len(s.Items)
}
