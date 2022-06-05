package instance

import (
	"fmt"
	"hash/fnv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcenter/apps/application"
	"github.com/infraboard/mcenter/common/tools"
	"github.com/infraboard/mcube/http/request"
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
	return &SearchRequest{
		Page: request.NewDefaultPageRequest(),
		Tags: map[string]string{},
	}
}

func (req *RegistryRequest) Validate() error {
	return validate.Struct(req)
}

func NewDefaultInstance() *Instance {
	req := NewRegistryRequest()
	app := application.NewDefaultApplication()
	return &Instance{
		Domain:          app.Spec.Domain,
		Namespace:       app.Spec.Namespace,
		ApplicationName: app.Spec.Name,
		RegistryInfo:    req,
		Status:          NewDefaultStatus(),
		Config:          NewDefaultConfig(),
	}
}

func NewInstance(req *RegistryRequest, app *application.Application) (*Instance, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	if req.Name == "" {
		req.Name = "ins-" + tools.MakeBearer(8)
	}

	ins := &Instance{
		Domain:          app.Spec.Domain,
		Namespace:       app.Spec.Namespace,
		ApplicationName: app.Spec.Name,
		RegistryInfo:    req,
		Status:          NewDefaultStatus(),
		Config:          NewDefaultConfig(),
	}

	ins.Id = ins.FullNameHash()
	return ins, nil
}

func NewDefaultConfig() *Config {
	return &Config{
		Heartbeat: &HeartbeatConfig{},
	}
}

func NewDefaultStatus() *Status {
	return &Status{
		Online: time.Now().UnixMilli(),
	}
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

func NewDescribeInstanceRequest(id string) *DescribeInstanceRequest {
	return &DescribeInstanceRequest{
		Id: id,
	}
}

func NewUnregistryRequest(instanceId string) *UnregistryRequest {
	return &UnregistryRequest{
		InstanceId: instanceId,
	}
}
