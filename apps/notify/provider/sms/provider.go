package sms

import (
	"context"
	"strings"

	"github.com/infraboard/mcenter/common/validate"
)

type SmsNotifyer interface {
	Send(context.Context, *SendSMSRequest) error
}

type SendSMSRequest struct {
	// 短信模版的Id
	TemplateId string `json:"template_id"`
	// 模版参数
	TemplateParams []string `json:"template_params"`
	// 电话号码
	PhoneNumbers []string `json:"phone_numbers"`
}

// NewSendSMSRequest todo
func NewSendSMSRequest() *SendSMSRequest {
	return &SendSMSRequest{}
}

// AddParams todo
func (req *SendSMSRequest) PhoneNumbersString() string {
	return strings.Join(req.PhoneNumbers, ",")
}

// AddParams todo
func (req *SendSMSRequest) AddParams(params ...string) {
	req.TemplateParams = append(req.TemplateParams, params...)
}

// AddPhone todo
func (req *SendSMSRequest) AddPhone(params ...string) {
	req.PhoneNumbers = append(req.PhoneNumbers, params...)
}

// Validate todo
func (req *SendSMSRequest) Validate() error {
	return validate.Validate(req)
}

// InjectDefaultIsoCode todo
func (req *SendSMSRequest) InjectDefaultIsoCode() {
	for i, number := range req.PhoneNumbers {
		if !strings.HasPrefix(number, "+") {
			req.PhoneNumbers[i] = "+86" + number
		}
	}
}
