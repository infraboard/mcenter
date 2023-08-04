package tencent

import (
	"context"
	"fmt"
	"strings"

	"github.com/infraboard/mcenter/apps/notify"
	sms_provider "github.com/infraboard/mcenter/apps/notify/provider/sms"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20190711"
)

func NewSender(conf *notify.TencentSmsConfig) (sms_provider.SmsNotifyer, error) {
	if err := conf.Validate(); err != nil {
		return nil, fmt.Errorf("validate tencent sms config error, %s", err)
	}

	credential := common.NewCredential(
		conf.SecretId,
		conf.SecretKey,
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = conf.GetEndpoint()
	client, err := sms.NewClient(credential, "", cpf)
	if err != nil {
		return nil, err
	}

	return &Sender{
		conf: conf,
		sms:  client,
		log:  zap.L().Named("tencent.sms"),
	}, nil
}

type Sender struct {
	conf *notify.TencentSmsConfig
	sms  *sms.Client
	log  logger.Logger
}

// Send todo
func (s *Sender) Send(ctx context.Context, req *sms_provider.SendSMSRequest) error {
	// 补充默认+86
	req.InjectDefaultIsoCode()

	if err := req.Validate(); err != nil {
		return fmt.Errorf("validate send sms request error, %s", err)
	}

	request := sms.NewSendSmsRequest()

	request.PhoneNumberSet = common.StringPtrs(req.PhoneNumbers)
	request.TemplateParamSet = common.StringPtrs(req.TemplateParams)
	request.TemplateID = common.StringPtr(req.TemplateId)
	request.SmsSdkAppid = common.StringPtr(s.conf.AppId)
	request.Sign = common.StringPtr(s.conf.Sign)

	response, err := s.sms.SendSmsWithContext(ctx, request)
	if err != nil {
		return err
	}

	for i := range response.Response.SendStatusSet {
		if strings.ToUpper(*(response.Response.SendStatusSet[i].Code)) != "OK" {
			return fmt.Errorf("send sms error, response is %s", response.ToJsonString())
		}
	}

	s.log.Debugf("send sms response success: %s", response.ToJsonString())
	return nil
}
