package tencent

// 使用腾讯的语音消息提供语言通知的能力
// 控制台入库 https://console.cloud.tencent.com/vms
// 开发文档参考: https://cloud.tencent.com/document/product/1128/37343
// Go SDK文档: https://cloud.tencent.com/document/product/1128/51621
// 语音消息的调用地址: vms.tencentcloudapi.com

import (
	"context"
	"fmt"

	"github.com/infraboard/mcenter/apps/notify"
	"github.com/infraboard/mcenter/apps/notify/provider/voice"
	"github.com/infraboard/mcube/v2/ioc/config/log"
	"github.com/infraboard/mcube/v2/ioc/config/validator"
	"github.com/rs/zerolog"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	vms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/vms/v20200902"
)

func NewQcloudVoice(conf *notify.TencentVoiceConfig) (voice.VoiceNotifyer, error) {
	ins := &TencentVoiceNotifyer{
		TencentVoiceConfig: conf,
		log:                log.Sub("voice.tencent"),
	}
	if err := ins.validate(); err != nil {
		return nil, err
	}
	return ins, nil
}

type TencentVoiceNotifyer struct {
	*notify.TencentVoiceConfig
	log *zerolog.Logger
}

func (q *TencentVoiceNotifyer) validate() error {
	return validator.Validate(q)
}

/* 基本类型的设置:
* SDK 采用的是指针风格指定参数，即使对于基本类型也需要用指针来对参数赋值。
* SDK 提供对基本类型的指针引用封装函数
* 帮助链接：
* 语音消息控制台：https://console.cloud.tencent.com/vms
* vms helper：https://cloud.tencent.com/document/product/1128/37720
 */
func (v *TencentVoiceNotifyer) genVMSRequest(req *voice.SendVoiceRequest) *vms.SendTtsVoiceRequest {
	/* 实例化一个请求对象，根据调用的接口和实际情况，可以进一步设置请求参数
	 * 您可以直接查询 SDK 源码确定接口有哪些属性可以设置
	 * 属性可能是基本类型，也可能引用了另一个数据结构
	 * 推荐使用 IDE 进行开发，可以方便地跳转查阅各个接口和数据结构的文档说明 */
	request := vms.NewSendTtsVoiceRequest()
	request.TemplateId = common.StringPtr(req.TemplateId)
	request.TemplateParamSet = common.StringPtrs(req.TemplateParams)
	request.CalledNumber = common.StringPtr(req.PhoneNumber)
	request.VoiceSdkAppid = common.StringPtr(v.AppId)
	request.PlayTimes = common.Uint64Ptr(req.PlayTimes)
	request.SessionContext = common.StringPtr(req.SessionContext)
	return request
}

/* 必要步骤：
* 实例化一个认证对象，入参需要传入腾讯云账户密钥对 secretId 和 secretKey
* 本示例采用从环境变量读取的方式，需要预先在环境变量中设置这两个值
* 您也可以直接在代码中写入密钥对，但需谨防泄露，不要将代码复制、上传或者分享给他人
* CAM 密匙查询: https://console.cloud.tencent.com/cam/capi
 */
func (v *TencentVoiceNotifyer) Call(ctx context.Context, req *voice.SendVoiceRequest) (*notify.VoiceResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, fmt.Errorf("validate PhoneCallRequest error, %s", err)
	}
	credential := common.NewCredential(v.SecretId, v.SecretKey)
	/* 非必要步骤:
	* 实例化一个客户端配置对象，可以指定超时时间等配置 */
	cpf := profile.NewClientProfile()

	cpf.HttpProfile.ReqMethod = v.ReqMethod

	/* SDK 有默认的超时时间，非必要请不要进行调整
	* 如有需要请在代码中查阅以获取最新的默认值 */
	//cpf.HttpProfile.ReqTimeout = 5
	cpf.HttpProfile.Endpoint = v.Endpoint

	cpf.SignMethod = v.SignMethod

	client, _ := vms.NewClient(credential, v.Region, cpf)
	request := v.genVMSRequest(req)

	// 通过 client 对象调用想要访问的接口，需要传入请求对象
	response, err := client.SendTtsVoice(request)

	// 处理异常
	if err != nil {
		return nil, fmt.Errorf("an api error has returned: %s", err)
	}

	// 打印返回的 JSON 字符串
	v.log.Debug().Msgf("response: %s", response.ToJsonString())
	return &notify.VoiceResponse{
		CallId:         *response.Response.SendStatus.CallId,
		SessionContext: *response.Response.SendStatus.SessionContext,
	}, nil
}
