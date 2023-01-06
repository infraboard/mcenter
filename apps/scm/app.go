package scm

import (
	"fmt"
	"path"

	"github.com/infraboard/mcenter/common/validate"
)

const (
	AppName = "scm"
)

type Service interface {
	RPCServer
}

func NewQueryProjectRequest() *QueryProjectRequest {
	return &QueryProjectRequest{}
}

func NewDefaultWebHookEvent() *WebHookEvent {
	return &WebHookEvent{
		Commits: []*Commit{},
	}
}

func NewProjectSet() *ProjectSet {
	return &ProjectSet{
		Items: []*Project{},
	}
}

func (e *WebHookEvent) ShortDesc() string {
	return fmt.Sprintf("%s %s [%s]", e.Ref, e.EventName, e.ObjectKind)
}

func (e *WebHookEvent) GetBranche() string {
	return path.Base(e.GetRef())
}

func (req *QueryProjectRequest) SetProviderFromString(provider string) error {
	p, err := ParsePROVIDERFromString(provider)
	if err != nil {
		return err
	}
	req.Provider = p
	return nil
}

func (req *QueryProjectRequest) Validate() error {
	return validate.Validate(req)
}
