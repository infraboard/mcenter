package tencent_test

import (
	"context"
	"os"
	"testing"

	"github.com/infraboard/mcenter/apps/notify"
	"github.com/infraboard/mcenter/apps/notify/provider/voice"
	"github.com/infraboard/mcenter/apps/notify/provider/voice/tencent"
	"github.com/infraboard/mcube/logger/zap"
)

var (
	notifyer voice.VoiceNotifyer

	ctx             = context.Background()
	voiceTemplateId = os.Getenv("VOICE_TENCENT_TEMPLATE_ID")
	// 被叫手机号码，采用 e.164 标准，格式为+[国家或地区码][用户号码], 例如：+8613711112222
	testPhoneNumber = os.Getenv("VOICE_TEST_PHONE_NUMBER")
)

func TestQcloudVoice(t *testing.T) {
	req := voice.NewSendVoiceRequest(testPhoneNumber, voiceTemplateId, []string{"测试"})
	_, err := notifyer.Call(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
}

func init() {
	zap.DevelopmentSetup()

	conf, err := notify.LoadTencentVoiceConfigFromEnv()
	if err != nil {
		panic(err)
	}
	vms, err := tencent.NewQcloudVoice(conf)
	if err != nil {
		panic(err)
	}
	notifyer = vms
}
