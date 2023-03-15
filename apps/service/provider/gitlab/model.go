package gitlab

import (
	"fmt"
	"net/url"
)

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

func NewDeleteProjectReqeust(projectID, hookID int64) *DeleteProjectReqeust {
	return &DeleteProjectReqeust{
		ProjectID: projectID,
		HookID:    hookID,
	}
}

type DeleteProjectReqeust struct {
	ProjectID int64 `json:"project_id"`
	HookID    int64 `json:"hook_id"`
}
