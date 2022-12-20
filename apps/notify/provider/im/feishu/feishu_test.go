package feishu_test

import (
	"context"
	"testing"

	"github.com/caarlos0/env/v6"
	"github.com/infraboard/mcenter/apps/domain"
	"github.com/infraboard/mcenter/apps/notify/provider/im/feishu"
)

var (
	impl *feishu.Feishu
	ctx  = context.Background()
)

func TestXxx(t *testing.T) {
	impl.Send(ctx)
}

func init() {
	conf := domain.NewDefaultFeishuConfig()
	if err := env.Parse(conf); err != nil {
		panic(err)
	}
	impl = feishu.NewFeishuNotifyer(conf)
}
