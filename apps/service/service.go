package service

import "github.com/infraboard/mcenter/apps/service/provider/gitlab"

func (s *ServiceSet) UpdateFromGitProject(p *gitlab.Project) {
	for i := range s.Items {
		item := s.Items[i]
		if item.GetRepositorySshUrl() == p.GitSshUrl {

		}
	}
}

func (s *Service) GetRepositorySshUrl() string {
	if s.Spec.Repository != nil {
		return s.Spec.Repository.SshUrl
	}

	return ""
}
