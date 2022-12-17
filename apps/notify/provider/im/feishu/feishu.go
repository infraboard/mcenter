package feishu

import (
	"context"
	"fmt"

	"github.com/chyroc/lark"
	"github.com/infraboard/mcube/logger"
)

var (
	ctx = context.Background()
)

type Feishu struct {
	client *lark.Lark
	log    logger.Logger
}

func (f *Feishu) Send() {
	data, resp, err := f.client.Message.Send().ToUserID("uid").SendCard(ctx, "")
	if err != nil {
		f.log.Errorf("send feishu message %s to user %s error, %s")
	}
	fmt.Println(data, resp)
}
