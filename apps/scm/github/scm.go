package github

import (
	"context"

	"github.com/google/go-github/v45/github"
	"golang.org/x/oauth2"
)

func NewScm(token string) *Scm {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	// user, resp, err := client.Users.Get(ctx, "")
	// if err != nil {
	// 	fmt.Printf("\nerror: %v\n", err)
	// 	return
	// }

	return &Scm{
		client: client,
	}
}

type Scm struct {
	client *github.Client
}
