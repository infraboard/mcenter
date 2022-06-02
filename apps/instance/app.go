package instance

import (
	"fmt"
	"hash/fnv"

	"github.com/go-playground/validator/v10"
)

const (
	AppName = "instance"
)

const (
	DefaultRegion      = "default"
	DefaultEnvironment = "default"
	DefaultGroup       = "default"
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

func (i *Instance) FullName() string {
	return fmt.Sprintf("%s.%s.%s-%s", i.Domain, i.Namespace, i.ApplicationName,
		i.RegistryInfo.InstanceFullName())
}

func (i *Instance) FullNameHash() string {
	hash := fnv.New32a()
	hash.Write([]byte(i.FullName()))
	return fmt.Sprintf("%x", hash.Sum32())
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
		Region:      DefaultRegion,
		Environment: DefaultEnvironment,
		Group:       DefaultGroup,
		Tags:        map[string]string{},
		Build:       &Build{},
	}
}

func (r *RegistryRequest) InstanceFullName() string {
	return fmt.Sprintf("%s.%s.%s.%s.%s", r.Region, r.Environment, r.Group, r.Name, r.Protocal)
}

func NewInstanceSet() *InstanceSet {
	return &InstanceSet{
		Items: []*Instance{},
	}
}

func (s *InstanceSet) Add(item *Instance) {
	s.Items = append(s.Items, item)
}
