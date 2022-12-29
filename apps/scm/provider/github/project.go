package github

import (
	"context"

	"github.com/google/go-github/v45/github"
	"github.com/infraboard/mcenter/apps/scm"
)

func (s *Github) ListOrganizations(ctx context.Context) (*scm.ProjectSet, error) {
	s.client.Organizations.List(ctx, "", &github.ListOptions{})
	// s.client.Repositories.List()
	// r, resp, err := s.client.Repositories.List(ctx, "", &github.RepositoryListOptions{})
	// s.client.Repositories.CreateHook()
	return nil, nil
}

func (s *Github) ListProjects(ctx context.Context) (*scm.ProjectSet, error) {
	s.client.Organizations.ListProjects(ctx, "", &github.ProjectListOptions{})
	return nil, nil
}
