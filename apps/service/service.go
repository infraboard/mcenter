package service

import "github.com/infraboard/mcenter/apps/service/provider/gitlab"

func (s *ServiceSet) UpdateFromGitProject(p *gitlab.Project) {
	svc := s.GetServiceByGitSshUrl(p.GitSshUrl)
	if svc == nil {
		svc = NewServiceFromProject(p)
		s.Add(svc)
	}
}

func (s *ServiceSet) GetServiceByGitSshUrl(gitSshUrl string) *Service {
	for i := range s.Items {
		item := s.Items[i]
		if item.GetRepositorySshUrl() == gitSshUrl {
			return item
		}
	}

	return nil
}

func NewServiceFromProject(p *gitlab.Project) *Service {
	svc := NewDefaultService()
	spec := svc.Spec
	spec.Name = p.Name
	spec.Description = p.Description
	spec.Repository.ProjectId = p.IdToString()
	spec.Repository.SshUrl = p.GitSshUrl
	spec.Repository.HttpUrl = p.GitHttpUrl
	spec.Repository.Namespace = p.NamespacePath
	return svc
}

func (s *Service) GetRepositorySshUrl() string {
	if s.Spec.Repository != nil {
		return s.Spec.Repository.SshUrl
	}

	return ""
}
