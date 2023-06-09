package mail

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/go-gomail/gomail"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

const (
	ERR_BROKEN_CONN = "write: broken pipe"
)

func NewSender(conf *Config) MailNotifyer {
	s := &Sender{
		Config: conf,
		log:    zap.L().Named("Mail Sender"),
	}
	return s
}

type Sender struct {
	*Config

	log    logger.Logger
	sender gomail.SendCloser
	lock   sync.Mutex
}

func (s *Sender) init() error {
	d := gomail.NewDialer(s.Host, s.Port, s.Username, s.Password)
	sender, err := d.Dial()
	if err != nil {
		return err
	}
	s.sender = sender
	return nil
}

func (s *Sender) Send(ctx context.Context, req *SendMailRequest) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.sender == nil {
		if err := s.init(); err != nil {
			return err
		}
	}

	m := gomail.NewMessage()
	m.SetHeader("From", s.Username)
	m.SetHeader("To", req.To...)
	m.SetHeader("Subject", req.Title)
	m.SetBody("text/html", req.Content)

	err := gomail.Send(s.sender, m)
	if err != nil {
		// 如果sender链接异常, 尝试重链
		if strings.Contains(err.Error(), ERR_BROKEN_CONN) {
			if err := s.init(); err != nil {
				return fmt.Errorf("Sender 初始化异常, %s", err)
			}

			return gomail.Send(s.sender, m)
		}

		return err
	}

	return nil
}
