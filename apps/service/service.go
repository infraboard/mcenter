package service

import (
	"encoding/json"
	"fmt"
	"net/url"
	"sort"
	"strconv"

	"github.com/infraboard/mcenter/apps/service/provider/gitlab"
	"github.com/infraboard/mcube/v2/ioc/config/application"
	"github.com/infraboard/mcube/v2/pb/resource"
	"github.com/infraboard/mcube/v2/tools/pretty"
	"google.golang.org/grpc/metadata"
)

func (s *ServiceSet) Len() int {
	return len(s.Items)
}

func (s *ServiceSet) ToJson() string {
	return pretty.ToJSON(s)
}

func (s *ServiceSet) UpdateFromGitProject(p *gitlab.Project, tk string) {
	svc := s.GetServiceByGitSshUrl(p.GitSshUrl)
	if svc == nil {
		// 创建新的服务
		svc = NewServiceFromProject(p)
		// 补充WebHook
		app := application.Get()
		svc.Spec.CodeRepository.EnableHook = app.IsInternalIP()
		svc.Spec.CodeRepository.Token = tk
		hc := gitlab.NewGitLabWebHook(svc.Meta.Id)
		hc.Url = fmt.Sprintf("%s/mflow/api/v1/triggers/gitlab", app.Endpoint())
		svc.Spec.CodeRepository.HookConfig = hc.ToJson()
		s.Add(svc)
	} else {
		// 更新服务
		svc.UpdateCreateAt(p.CreatedAtTime().Unix())
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

func (s *ServiceSet) GetServiceById(id string) *Service {
	for i := range s.Items {
		item := s.Items[i]
		if item.Meta.Id == id {
			return item
		}
	}

	return nil
}

func (s *ServiceSet) Sort() *ServiceSet {
	sort.Sort(s)
	return s
}

func (s *ServiceSet) Less(i, j int) bool {
	return s.Items[i].Meta.CreateAt < s.Items[j].Meta.CreateAt
}

func (s *ServiceSet) Swap(i, j int) {
	s.Items[i], s.Items[j] = s.Items[j], s.Items[i]
}

func NewServiceFromProject(p *gitlab.Project) *Service {
	svc := NewDefaultService()
	svc.Meta.CreateAt = p.CreatedAtTime().Unix()

	spec := svc.Spec
	spec.Name = p.UniName()
	spec.Description = p.Name
	spec.Logo = p.AvatarURL
	spec.Description = p.Description
	spec.CodeRepository.ProjectId = p.IdToString()
	spec.CodeRepository.SshUrl = p.GitSshUrl
	spec.CodeRepository.HttpUrl = p.GitHttpUrl
	spec.CodeRepository.Namespace = p.NameSpace()
	spec.CodeRepository.WebUrl = p.WebURL
	spec.CodeRepository.CreatedAt = p.CreatedAtTime().Unix()
	return svc
}

func (s *Service) GetRepositorySshUrl() string {
	if s.Spec.CodeRepository != nil {
		return s.Spec.CodeRepository.SshUrl
	}

	return ""
}

func (s *Service) UpdateCreateAt(ts int64) {
	if s.Meta == nil {
		s.Meta = resource.NewMeta()
	}
	s.Meta.CreateAt = ts
}

func (s *Service) InjectGrpcClientMeta(md metadata.MD) {
	md.Append(GRPC_CLIENT_SERVICE_JSON, s.ToJson())
}

func (s *Service) Desense() *Service {
	if s.Credential != nil {
		s.Credential.ClientSecret = ""
	}
	if s.Security != nil {
		s.Security.EncryptKey = ""
	}
	return s
}

func (s *Service) ToJson() string {
	return pretty.ToJSON(s)
}

func (s *Service) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		*resource.Meta
		*CreateServiceRequest
		Credential *Credential `json:"credential"`
		Security   *Security   `json:"security"`
	}{s.Meta, s.Spec, s.Credential, s.Security})
}

func (s *Service) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(
		data,
		&struct {
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
