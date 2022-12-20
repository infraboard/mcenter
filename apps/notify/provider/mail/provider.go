package mail

import "context"

type MailNotifyer interface {
	Send(context.Context, *SendMailRequest) error
}

type SendMailRequest struct {
	// 用户
	To []string `json:"to"  validate:"required"`
	// 告警时标题
	Title string `json:"title"  validate:"required"`
	// 告警内容
	Content string `json:"content"`
}

// NewSendMailRequest todo
func NewSendMailRequest(to []string, title, content string) *SendMailRequest {
	return &SendMailRequest{
		To:      to,
		Title:   title,
		Content: content,
	}
}
