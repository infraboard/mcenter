package github

import (
	"context"

	"github.com/google/go-github/v45/github"
	"golang.org/x/oauth2"
)

func NewGithubFromPersonalAccessToken(token string) *Github {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	return &Github{
		client: client,
	}
}

func NewGithubFromOauthApp(conf *oauth2.Config) *Github {
	return &Github{}
}

type Github struct {
	client *github.Client
}
