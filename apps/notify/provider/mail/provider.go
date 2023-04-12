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

func (req *SendMailRequest) HasTo() bool {
	return len(req.To) > 0
}

func (req *SendMailRequest) ToStrings() string {
	return strings.Join(req.To, ",")
}

func (req *SendMailRequest) AddTo(tos ...string) {
	for _, v := range tos {
		if v != "" {
			req.To = append(req.To, v)
		}
	}
}

// NewSendMailRequest todo
func NewSendMailRequest(title, content string, tos ...string) *SendMailRequest {
	req := &SendMailRequest{
		Title:   title,
		Content: content,
	}
	req.AddTo(tos...)
	return req
}
