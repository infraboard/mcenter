package github

import (
	"context"

	"github.com/google/go-github/v45/github"
	"github.com/infraboard/mcenter/apps/scm"
)

func (s *Github) ListProjects(ctx context.Context) (*scm.ProjectSet, error) {
	s.client.Repositories.List(ctx, "", &github.RepositoryListOptions{})
	// r, resp, err := s.client.Repositories.List(ctx, "", &github.RepositoryListOptions{})
	// s.client.Repositories.CreateHook()
	return nil, nil
}
