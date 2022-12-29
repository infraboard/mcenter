package github

import (
	"context"

	"github.com/google/go-github/v45/github"
	"github.com/infraboard/mcube/client/rest"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"golang.org/x/oauth2"
)

func NewGithub(conf *Config) *Github {
	ins := &Github{
		conf: conf,
		log:  zap.L().Named("scm.github"),
		rest: rest.NewRESTClient(),
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

func (g *Github) oauth2Config() *oauth2.Config {
	return g.conf.Oauth2Config.OauthConf()
}

func (g *Github) AuthCodeURL() string {
	return g.oauth2Config().AuthCodeURL("")
}

func (g *Github) Exchange(ctx context.Context, code string) error {
	tk, err := g.oauth2Config().Exchange(ctx, code)
	if err != nil {
		return err
	}
	g.log.Debug(tk)
	return nil
}
