package mail

import (
	"context"
	"strings"
)

type MailNotifyer interface {
	Send(context.Context, *SendMailRequest) error
}

type SendMailRequest struct {
	// 邮箱列表
	To []string `json:"to"  validate:"required"`
	// 告警时标题
	Title string `json:"title"  validate:"required"`
	// 告警内容
	Content string `json:"content"`
}

func (req *SendMailRequest) ToStrings() string {
	return strings.Join(req.To, ",")
}

func (req *SendMailRequest) AddTo(tos ...string) {
	req.To = append(req.To, tos...)
}

// NewSendMailRequest todo
func NewSendMailRequest(title, content string, tos ...string) *SendMailRequest {
	return &SendMailRequest{
		Title:   title,
		Content: content,
		To:      tos,
	}
}
