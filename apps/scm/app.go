package scm

import (
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

func NewProjectSet() *ProjectSet {
	return &ProjectSet{
		Items: []*Project{},
	}
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
