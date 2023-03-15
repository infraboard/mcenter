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

	ins := &Github{
		conf: conf,
		log:  zap.L().Named("scm.github"),
		rest: rc,
	}

	switch conf.AuthType {
	case AUTH_TYPE_OAUTH2:
	case AUTH_TYPE_PERSONAL_ACCESS_TOKEN:
		ctx := context.Background()
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: conf.PersonalAccessToken},
		)
		tc := oauth2.NewClient(ctx, ts)
		ins.client = github.NewClient(tc)
	}

	return ins
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
	return nil
}
