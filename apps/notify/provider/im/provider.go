package im

import "context"

// IM消息通知器
type ImNotifyer interface {
	SendMessage(context.Context, *SendMessageRequest) error
}

func NewSendMessageRequest(title, content string, uid string) *SendMessageRequest {
	return &SendMessageRequest{
		Uid:     uid,
		Title:   title,
		Content: content,
	}
}

type SendMessageRequest struct {
	// 飞书用户Id
	Uid string `json:"users" validate:"required"`
	// 告警时标题
	Title string `json:"title" validate:"required"`
	// 告警内容
	Content string `json:"content"`
}
