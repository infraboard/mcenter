package github

import (
	"fmt"

	"github.com/caarlos0/env/v6"
	"golang.org/x/oauth2"

	"github.com/infraboard/mcenter/common/validate"
)

func LoadConfigFromEnv() (*Config, error) {
	conf := NewDefaultConfig()
	err := env.Parse(conf)
	if err != nil {
		return nil, err
	}
	if err := conf.Validate(); err != nil {
		return nil, err
	}
	return conf, nil
}

func NewDefaultConfig() *Config {
	return &Config{
		AuthType:     AUTH_TYPE_OAUTH2,
		Oauth2Config: &Oauth2Config{},
	}
}

type Config struct {
	AuthType            AUTH_TYPE
	PersonalAccessToken string        `json:"personal_access_token" env:"GITHUB_PERSONAL_ACCESS_TOKEN"`
	Oauth2Config        *Oauth2Config `json:"oauth2_config"`
}

func (c *Config) Validate() error {
	switch c.AuthType {
	case AUTH_TYPE_PERSONAL_ACCESS_TOKEN:
		if c.PersonalAccessToken == "" {
			return fmt.Errorf("personal access token required")
		}
	case AUTH_TYPE_OAUTH2:
		return validate.Validate(c.Oauth2Config)
	default:
		return fmt.Errorf("unknown auth type %s", c.AuthType)
	}

	return nil
}

type AUTH_TYPE string

const (
	// OAuth2.0认证
	AUTH_TYPE_OAUTH2 AUTH_TYPE = "oauth2"
	// 个人Token认证
	AUTH_TYPE_PERSONAL_ACCESS_TOKEN = "personal_token"
)

type Oauth2Config struct {
	AppId        string   `json:"app_id" env:"GITHUB_APP_ID"`
	ClientId     string   `json:"client_id" env:"GITHUB_CLIENT_ID"`
	ClientSecret string   `json:"client_secret" env:"GITHUB_CLIENT_SECRET"`
	RedirectURL  string   `json:"redirect_url" env:"GITHUB_REDIRECT_URL"`
	Scopes       []string `json:"scopes" env:"GITHUB_SCOPES"`
}

// 参考文档: https://docs.github.com/zh/rest/guides/basics-of-authentication?apiVersion=2022-11-28
func (c *Oauth2Config) OauthConf() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     c.ClientId,
		ClientSecret: c.ClientSecret,
		RedirectURL:  c.RedirectURL,
		Scopes:       c.Scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:   "https://github.com/login/oauth/authorize",
			TokenURL:  "https://github.com/login/oauth/access_token",
			AuthStyle: oauth2.AuthStyleInParams,
		},
	}
}
