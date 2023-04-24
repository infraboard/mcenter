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
	// 消息标题
	Title string `json:"title" validate:"required"`
	// 消息内容
	Content string `json:"content"`
	// 消息格式, 可以为空
	ContentType string `json:"content_type"`
}
