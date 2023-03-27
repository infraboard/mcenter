package gitlab

import (
	"encoding/json"
	"fmt"
	"net/url"
	"sort"
	"time"

	"github.com/infraboard/mcenter/common/format"
)

func NewProjectSet() *ProjectSet {
	return &ProjectSet{
		Items: []*Project{},
	}
}

type ProjectSet struct {
	Total int64
	Items []*Project
}

func (s *ProjectSet) String() string {
	return format.Prettify(s)
}

func (s *ProjectSet) Len() int {
	return len(s.Items)
}

func (s *ProjectSet) GitSshUrls() (urls []string) {
	for i := range s.Items {
		item := s.Items[i]
		if item.GitSshUrl != "" {
			urls = append(urls, item.GitSshUrl)
		}
	}
	return
}

type Project struct {
	// 项目id
	Id int64 `json:"id"`
	// 描述
	Description string `json:"description"`
	// 项目创建时间
	CreatedAt time.Time `json:"created_at"`
	// 名称
	Name string `json:"name"`
	// 项目的Web访问地址
	WebURL string `json:"web_url"`
	// 项目的Logo地址
	AvatarURL string `json:"avatar_url"`
	// ssh 地址
	GitSshUrl string `json:"ssh_url_to_repo"`
	// http 地址
	GitHttpUrl string `json:"http_url_to_repo"`
	// namespace
	NamespacePath string `json:"path_with_namespace"`
}

func (s *Project) IdToString() string {
	return fmt.Sprintf("%d", s.Id)
}

func NewGitLabWebHook(token string) *GitLabWebHook {
	return &GitLabWebHook{
		PushEvents:          true,
		IssuesEvents:        true,
		MergeRequestsEvents: true,
		TagPushEvents:       true,
		NoteEvents:          true,
		ReleasesEvents:      true,
		Token:               token,
		Url:                 "http://api.mdevcloud.cn/mpaas/api/v1/triggers/gitlab",
	}
}

func ParseGitLabWebHookFromString(conf string) (*GitLabWebHook, error) {
	hook := NewGitLabWebHook("")
	if conf != "" {
		err := json.Unmarshal([]byte(conf), hook)
		if err != nil {
			return nil, err
		}
	}

	return hook, nil
}

type GitLabWebHook struct {
	PushEventsBranchFilter   string `json:"push_events_branch_filter"`
	PushEvents               bool   `json:"push_events"`
	IssuesEvents             bool   `json:"issues_events"`
	ConfidentialIssuesEvents bool   `json:"confidential_issues_events"`
	MergeRequestsEvents      bool   `json:"merge_requests_events"`
	TagPushEvents            bool   `json:"tag_push_events"`
	NoteEvents               bool   `json:"note_events"`
	ConfidentialNoteEvents   bool   `json:"confidential_note_events"`
	WikiPageEvents           bool   `json:"wiki_page_events"`
	ReleasesEvents           bool   `json:"releases_events"`
	Token                    string `json:"token"`
	Url                      string `json:"url"`
}

func (h *GitLabWebHook) ToJson() string {
	return format.Prettify(h)
}

func (req *GitLabWebHook) FormValue() url.Values {
	val := make(url.Values)
	val.Set("push_events", fmt.Sprintf("%t", req.PushEvents))
	val.Set("tag_push_events", fmt.Sprintf("%t", req.TagPushEvents))
	val.Set("merge_requests_events", fmt.Sprintf("%t", req.MergeRequestsEvents))
	val.Set("token", req.Token)
	val.Set("url", req.Url)
	return val
}

func NewAddProjectHookRequest(projectID int64, webhook *GitLabWebHook) *AddProjectHookRequest {
	return &AddProjectHookRequest{
		ProjectID: projectID,
		WebHook:   webhook,
	}
}

type AddProjectHookRequest struct {
	// 项目Id
	ProjectID int64 `json:"project_id"`
	// Gitlab WebHook配置
	WebHook *GitLabWebHook `json:"webhook"`
}

func NewAddProjectHookResponse() *AddProjectHookResponse {
	return &AddProjectHookResponse{
		GitLabWebHook: &GitLabWebHook{},
	}
}

type AddProjectHookResponse struct {
	ID int64 `json:"id"`
	*GitLabWebHook
}

func (r *AddProjectHookResponse) IDToString() string {
	return fmt.Sprintf("%d", r.ID)
}

func NewDeleteProjectHookReqeust(projectID, hookID int64) *DeleteProjectHookReqeust {
	return &DeleteProjectHookReqeust{
		ProjectID: projectID,
		HookID:    hookID,
	}
}

type DeleteProjectHookReqeust struct {
	ProjectID int64 `json:"project_id"`
	HookID    int64 `json:"hook_id"`
}

func NewListProjectRequest() *ListProjectRequest {
	return &ListProjectRequest{
		Owned:     true,
		Simple:    true,
		PageNumer: 1,
		PageSize:  20,
	}
}

type ListProjectRequest struct {
	Owned     bool
	Simple    bool
	PageSize  int64
	PageNumer int64
}

func (r *ListProjectRequest) PageSizeToString() string {
	return fmt.Sprintf("%d", r.PageSize)
}

func (r *ListProjectRequest) PageNumerToString() string {
	return fmt.Sprintf("%d", r.PageNumer)
}

func NewProjectLanguageSet(percentage map[string]float64) *ProjectLanguageSet {
	set := &ProjectLanguageSet{
		Items: []*ProjectLanguage{},
	}
	for k, v := range percentage {
		set.Add(&ProjectLanguage{
			Language:   k,
			Percentage: v,
		})
	}
	sort.Sort(set)
	return set
}

type ProjectLanguageSet struct {
	Items []*ProjectLanguage `json:"items"`
}

type ProjectLanguage struct {
	Language   string  `json:"language"`
	Percentage float64 `json:"percentage"`
}

func (p *ProjectLanguageSet) String() string {
	return format.Prettify(p)
}

func (p *ProjectLanguageSet) Add(item *ProjectLanguage) {
	p.Items = append(p.Items, item)
}

func (p *ProjectLanguageSet) Len() int {
	return len(p.Items)
}
func (p *ProjectLanguageSet) Less(i, j int) bool {
	return p.Items[i].Percentage > p.Items[j].Percentage
}
func (p *ProjectLanguageSet) Swap(i, j int) {
	p.Items[i], p.Items[j] = p.Items[j], p.Items[i]
}

// 主语言
func (p *ProjectLanguageSet) Primary() string {
	if p.Len() == 0 {
		return ""
	}
	return p.Items[0].Language
}
