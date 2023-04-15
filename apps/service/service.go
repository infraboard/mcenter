package service

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/infraboard/mcenter/apps/service/provider/gitlab"
	"github.com/infraboard/mcenter/common/format"
	"github.com/infraboard/mcube/pb/resource"
)

func (s *ServiceSet) Len() int {
	return len(s.Items)
}

func (s *ServiceSet) ToJson() string {
	return format.Prettify(s)
}

func (s *ServiceSet) UpdateFromGitProject(p *gitlab.Project, tk string) {
	svc := s.GetServiceByGitSshUrl(p.GitSshUrl)
	if svc == nil {
		svc = NewServiceFromProject(p)
		svc.Spec.CodeRepository.Token = tk
		s.Add(svc)
	}
}

func (s *ServiceSet) UpdateScope(domain, namespace string) {
	for i := range s.Items {
		item := s.Items[i]
		item.Spec.Domain = domain
		item.Spec.Namespace = namespace
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
	spec.Logo = p.AvatarURL
	spec.Description = p.Description
	spec.CodeRepository.ProjectId = p.IdToString()
	spec.CodeRepository.SshUrl = p.GitSshUrl
	spec.CodeRepository.HttpUrl = p.GitHttpUrl
	spec.CodeRepository.Namespace = p.NamespacePath
	spec.CodeRepository.WebUrl = p.WebURL
	spec.CodeRepository.CreatedAt = p.CreatedAt.Unix()
	spec.CodeRepository.EnableHook = true
	spec.CodeRepository.HookConfig = gitlab.NewGitLabWebHook(
		"自动填充服务的Id",
	).ToJson()
	return svc
}

func (s *Service) GetRepositorySshUrl() string {
	if s.Spec.CodeRepository != nil {
		return s.Spec.CodeRepository.SshUrl
	}

	return ""
}

func (s *Service) ToJson() string {
	return format.Prettify(s)
}

func (s *Service) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		*resource.Meta
		*CreateServiceRequest
		Credential *Credential `json:"status"`
		Security   *Security   `json:"security"`
	}{s.Meta, s.Spec, s.Credential, s.Security})
}

func NewImageRepository() *ImageRepository {
	return &ImageRepository{
		Version: "latest",
	}
}

func NewCodeRepository() *CodeRepository {
	return &CodeRepository{
		EnableHook: true,
	}
}

func (r *CodeRepository) SetLanguage(v LANGUAGE) {
	r.Language = &v
}

func (r *CodeRepository) ProjectIdToInt64() int64 {
	pid, _ := strconv.ParseInt(r.ProjectId, 10, 64)
	return pid
}

func (r *CodeRepository) Validate() error {
	return validate.Struct(r)
}

func (r *CodeRepository) HookIdToInt64() int64 {
	pid, _ := strconv.ParseInt(r.HookId, 10, 64)
	return pid
}

func (r *CodeRepository) MakeGitlabConfig() (*gitlab.Config, error) {
	conf := gitlab.NewDefaultConfig()
	addr, err := r.HostAddress()
	if err != nil {
		return nil, err
	}
	conf.Address = addr
	conf.PrivateToken = r.Token
	return conf, nil
}

func (r *CodeRepository) HostAddress() (string, error) {
	u, err := url.Parse(r.WebUrl)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s://%s", u.Scheme, u.Host), nil
}
