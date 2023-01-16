package github

import (
	"context"

	"github.com/google/go-github/v45/github"
	"github.com/infraboard/mcube/client/negotiator"
	"github.com/infraboard/mcube/client/rest"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"golang.org/x/oauth2"
)

func NewGithub(conf *Config) *Github {
	rc := rest.NewRESTClient()
	rc.SetBaseURL(conf.Endpoint)
	rc.SetHeader("Accept", "application/vnd.github+json")

	g := &Github{
		conf: conf,
		log:  zap.L().Named("scm.github"),
		rest: rc,
	}

	switch conf.AuthType {
	case AUTH_TYPE_OAUTH2:
	case AUTH_TYPE_PERSONAL_ACCESS_TOKEN:
		g.SetToken(conf.PersonalAccessToken)
	}

	return g
}

type Github struct {
	conf   *Config
	client *github.Client
	log    logger.Logger
	rest   *rest.RESTClient
}

func (g *Github) AuthCodeURL() string {
	return g.conf.AuthCodeURL("")
}

// 参考文档: https://docs.github.com/zh/developers/apps/building-github-apps/identifying-and-authorizing-users-for-github-apps#2-users-are-redirected-back-to-your-site-by-github
func (g *Github) Exchange(ctx context.Context, code string) error {
	tk := NewTokenReponse()
	err := g.rest.
		Post("login/oauth/access_token").
		Header(rest.CONTENT_TYPE_HEADER, string(negotiator.MIME_POST_FORM)).
		Body(g.conf.Oauth2Config.ExchangeTokenRequeset(code)).
		Do(ctx).
		Into(tk)
	if err == nil {
		err = tk.Error()
	}

	if err != nil {
		return err
	}
	g.log.Debug(tk)

	g.SetToken(tk.AccessToken)
	return nil
}

func (g *Github) SetToken(tk string) {
	// 设置sdk token
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: tk},
	)
	tc := oauth2.NewClient(context.Background(), ts)
	g.client = github.NewClient(tc)

	// 设置rest client tk
	g.rest.SetBearerTokenAuth(tk)
}
