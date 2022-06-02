package application

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/imdario/mergo"
	request "github.com/infraboard/mcube/http/request"
	pb_request "github.com/infraboard/mcube/pb/request"
	"github.com/rs/xid"
)

const (
	AppName = "application"
)

var (
	validate = validator.New()
)

func NewCreateApplicationRequest() *CreateApplicationRequest {
	return &CreateApplicationRequest{}
}

func NewApplication(req *CreateApplicationRequest) (*Application, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	return &Application{
		Id:       xid.New().String(),
		CreateAt: time.Now().UnixMilli(),
		Spec:     req,
	}, nil
}

func NewValidateCredentialRequest(clientId, clientSercet string) *ValidateCredentialRequest {
	return &ValidateCredentialRequest{
		ClientId:     clientId,
		ClientSecret: clientSercet,
	}
}

func (req *CreateApplicationRequest) Validate() error {
	return validate.Struct(req)
}

func NewApplicationApplicationSet() *ApplicationSet {
	return &ApplicationSet{
		Items: []*Application{},
	}
}

func (s *ApplicationSet) Add(item *Application) {
	s.Items = append(s.Items, item)
}

func NewDefaultApplication() *Application {
	return &Application{
		Spec: &CreateApplicationRequest{},
	}
}

func NewDescribeApplicationRequest(id string) *DescribeApplicationRequest {
	return &DescribeApplicationRequest{
		Id: id,
	}
}

func NewQueryApplicationRequestFromHTTP(r *http.Request) *QueryApplicationRequest {
	return &QueryApplicationRequest{
		Page: request.NewPageRequestFromHTTP(r),
	}
}

func NewDeleteApplicationRequestWithID(id string) *DeleteApplicationRequest {
	return &DeleteApplicationRequest{
		Id: id,
	}
}

func (i *Application) Update(req *UpdateApplicationRequest) {
	i.UpdateAt = time.Now().UnixMilli()
	i.UpdateBy = req.UpdateBy
	i.Spec = req.Spec
}

func (i *Application) Patch(req *UpdateApplicationRequest) error {
	i.UpdateAt = time.Now().UnixMicro()
	i.UpdateBy = req.UpdateBy
	return mergo.MergeWithOverwrite(i.Spec, req.Spec)
}

func NewPutApplicationRequest(id string) *UpdateApplicationRequest {
	return &UpdateApplicationRequest{
		Id:         id,
		UpdateMode: pb_request.UpdateMode_PUT,
		UpdateAt:   time.Now().UnixMilli(),
		Spec:       NewCreateApplicationRequest(),
	}
}

func NewPatchApplicationRequest(id string) *UpdateApplicationRequest {
	return &UpdateApplicationRequest{
		Id:         id,
		UpdateMode: pb_request.UpdateMode_PATCH,
		UpdateAt:   time.Now().UnixMilli(),
		Spec:       NewCreateApplicationRequest(),
	}
}

func NewDescribeApplicationRequestByClientId(clientId string) *DescribeApplicationRequest {
	return &DescribeApplicationRequest{
		DescribeBy: DescribeBy_APP_CLIENT_ID,
		ClientId:   clientId,
	}
}

func (c *Credential) Validate(clientSecret string) error {
	if c.ClientSecret != clientSecret {
		return fmt.Errorf("client_id or client_secret is not conrrect")
	}

	return nil
}
