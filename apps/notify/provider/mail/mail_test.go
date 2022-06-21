package mail_test

import (
	"testing"

	"github.com/caarlos0/env/v6"
	"github.com/infraboard/mcenter/apps/notify"
	"github.com/infraboard/mcenter/apps/notify/provider/mail"
)

var sender *mail.Sender

func TestSend(t *testing.T) {
	m := &notify.SendMailRequest{
		To:      []string{"719118794@qq.com"},
		Title:   "验证码",
		Content: "610101",
	}

	err := sender.Send(m)
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
