package github

import (
	"bytes"
	"fmt"
	"net/url"
	"strings"

	"github.com/caarlos0/env/v6"

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
		Endpoint:     "https://github.com",
		Oauth2Config: &Oauth2Config{},
	}
}

type Config struct {
	AuthType            AUTH_TYPE
	Endpoint            string        `json:"endpoint" env:"GITHUB_ENDPOINT"`
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

// 参考文档: https://docs.github.com/zh/rest/guides/basics-of-authentication?apiVersion=2022-11-28
func (c *Config) AuthURL() string {
	p, _ := url.JoinPath(c.Endpoint, "login/oauth/authorize")
	return p
}

// AuthCodeURL returns a URL to OAuth 2.0 provider's consent page
// that asks for permissions for the required scopes explicitly.
//
// State is a token to protect the user from CSRF attacks. You must
// always provide a non-empty string and validate that it matches the
// the state query parameter on your redirect callback.
// See http://tools.ietf.org/html/rfc6749#section-10.12 for more info.
//
// Opts may include AccessTypeOnline or AccessTypeOffline, as well
// as ApprovalForce.
// It can also be used to pass the PKCE challenge.
// See https://www.oauth.com/oauth2-servers/pkce/ for more info.
func (c *Config) AuthCodeURL(state string) string {
	oc := c.Oauth2Config
	if oc == nil {
		return ""
	}

	var buf bytes.Buffer
	buf.WriteString(c.AuthURL())
	v := url.Values{
		"response_type": {"code"},
		"client_id":     {oc.ClientId},
	}
	if oc.RedirectURL != "" {
		v.Set("redirect_uri", oc.RedirectURL)
	}
	if len(oc.Scopes) > 0 {
		v.Set("scope", strings.Join(oc.Scopes, " "))
	}
	if state != "" {
		// TODO(light): Docs say never to omit state; don't allow empty.
		v.Set("state", state)
	}
	if strings.Contains(c.AuthURL(), "?") {
		buf.WriteByte('&')
	} else {
		buf.WriteByte('?')
	}
	buf.WriteString(v.Encode())
	return buf.String()
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

func (o *Oauth2Config) ExchangeTokenRequeset(code string) url.Values {
	req := make(url.Values)
	req.Add("client_id", o.ClientId)
	req.Add("client_secret", o.ClientSecret)
	req.Add("code", code)
	req.Add("redirect_uri", o.RedirectURL)
	return req
}
