package tencent

import "github.com/caarlos0/env/v6"

func DefaultConfig() *Config {
	return &Config{
		Endpoint:   "vms.tencentcloudapi.com",
		Region:     "ap-guangzhou",
		SignMethod: "TC3-HMAC-SHA256",
		ReqMethod:  "POST",
	}
}

func LoadConfigFromEnv() (*Config, error) {
	conf := DefaultConfig()
	err := env.Parse(conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}

type Config struct {
	/* SDK 会自动指定域名，通常无需指定域名，但访问金融区的服务时必须手动指定域名
	 * 例如 VMS 的上海金融区域名为 vms.ap-shanghai-fsi.tencentcloudapi.com */
	Endpoint string `bson:"endpoint" json:"endpoint" env:"VOICE_TENCENT_ENDPOINT"`
	// 腾讯云凭证Id
	SecretId string `bson:"secret_id" json:"secret_id" validate:"required" env:"VOICE_TENCENT_SECRET_ID"`
	// 腾讯云凭证Key
	SecretKey string `bson:"secret_key" json:"secret_key" validate:"required" env:"VOICE_TENCENT_SECRET_KEY"`
	// 腾讯云控制台: https://console.cloud.tencent.com/vms/app 查看appId
	AppId string `bson:"app_id" json:"app_id" validate:"required" env:"VOICE_TENCENT_APPID"`
	/* 实例化 VMS 的 client 对象
	* 第二个参数是地域信息，可以直接填写字符串 ap-guangzhou，或者引用预设的常量 */
	Region string `bson:"region" json:"region" env:"VOICE_TENCENT_REGION"`
	/* SDK 默认用 TC3-HMAC-SHA256 进行签名，非必要请不要修改该字段 */
	SignMethod string `bson:"sign_method" json:"sign_method" env:"VOICE_TENCENT_SIGN_METHOD"`
	/* SDK 默认使用 POST 方法
	 * 如需使用 GET 方法，可以在此处设置，但 GET 方法无法处理较大的请求 */
	ReqMethod string `bson:"req_method" json:"req_method" env:"VOICE_TENCENT_REQ_METHOD"`
}
