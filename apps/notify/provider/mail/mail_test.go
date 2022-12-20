package mail_test

import (
	"context"
	"testing"

	"github.com/caarlos0/env/v6"
	"github.com/infraboard/mcenter/apps/notify/provider/mail"
)

var (
	sender *mail.Sender
	ctx    = context.Background()
)

func TestSend(t *testing.T) {
	m := &mail.SendMailRequest{
		To:      []string{"719118794@qq.com"},
		Title:   "验证码",
		Content: "610101",
	}

	err := sender.Send(ctx, m)
	if err != nil {
		panic(err)
	}
}

func init() {
	conf := &mail.Config{}
	if err := env.Parse(conf); err != nil {
		panic(err)
	}
	sender = mail.NewSender(conf)
}
