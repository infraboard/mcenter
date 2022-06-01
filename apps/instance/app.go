package instance

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

const (
	AppName = "instance"
)

var (
	validate = validator.New()
)

func NewSearchRequest() *SearchRequest {
	return &SearchRequest{}
}

func (req *RegistryRequest) Validate() error {
	return validate.Struct(req)
}

func NewInstance(req *RegistryRequest) (*Instance, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	return &Instance{
		RegistryInfo: req,
		Status:       &Status{},
		Config:       &Config{},
	}, nil
}

func (i *Instance) ShortDesc() string {
	return fmt.Sprintf("%s-%s-%s-%s", i.Domain, i.Namespace, i.ApplicationName, i.RegistryInfo.Name)
}

func NewHeartbeatResponse() *HeartbeatResponse {
	return &HeartbeatResponse{}
}

func NewHeartbeatRequest(id string) *HeartbeatRequest {
	return &HeartbeatRequest{
		InstanceId: id,
	}
}

func NewRegistryRequest() *RegistryRequest {
	return &RegistryRequest{
		Tags:  map[string]string{},
		Build: &Build{},
	}
}

func NewInstanceSet() *InstanceSet {
	return &InstanceSet{
		Items: []*Instance{},
	}
}

func (s *InstanceSet) Add(item *Instance) {
	s.Items = append(s.Items, item)
}
