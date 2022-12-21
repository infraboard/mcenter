package voice

import (
	"context"

	"github.com/infraboard/mcenter/common/validate"
)

type VoiceNotifyer interface {
	Call(context.Context, *SendVoiceRequest) (*SendVoiceResponse, error)
}

func NewSendVoiceRequest(number, templateId string, templateParamSet []string) *SendVoiceRequest {
	return &SendVoiceRequest{
		TemplateId:     templateId,
		TemplateParams: templateParamSet,
		PhoneNumber:    number,
		PlayTimes:      2,
	}
}

type SendVoiceRequest struct {
	// 短信模版的Id
	TemplateId string `json:"template_id"`
	// 模版参数
	TemplateParams []string `json:"template_params"`
	// 电话号码
	PhoneNumber string `json:"phone_number"`
	// 播放次数，可选，最多3次，默认2次
	PlayTimes uint64 `json:"play_times"`
	// 用户的 session 内容，腾讯 server 回包中会原样返回
	SessionContext string `json:"session_context"`
}

func (req *SendVoiceRequest) Validate() error {
	return validate.Validate(req)
}

type SendVoiceResponse struct {
	// 呼叫Id
	CallId string `json:"call_id" `
	// 用户的 session 内容，腾讯 server 回包中会原样返回
	SessionContext string `json:"session_context"`
}
