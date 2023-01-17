package instance

import (
	"fmt"
	"hash/fnv"
	"net/http"
	"sort"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcenter/apps/service"
	"github.com/infraboard/mcenter/apps/token"
	"github.com/infraboard/mcube/http/request"
)

const (
	AppName = "instances"
)

const (
	DefaultRegion      = "default"
	DefaultEnvironment = "default"
	DefaultGroup       = "default"
)

var (
	validate = validator.New()
)

type Service interface {
	RPCServer
}

func NewSearchRequestFromHttp(r *http.Request) *SearchRequest {
	qs := r.URL.Query()
	req := NewSearchRequest()
	req.Page = request.NewPageRequestFromHTTP(r)
	req.Domain = qs.Get("domain")
	req.Namespace = qs.Get("namespace")
	req.ServiceName = qs.Get("service_name")
	req.Region = qs.Get("region")
	req.Environment = qs.Get("env")
	req.Group = qs.Get("group")
	return req
}

func NewSearchRequest() *SearchRequest {
	return &SearchRequest{
		Page:   request.NewDefaultPageRequest(),
		Labels: map[string]string{},
	}
}

func (req *RegistryRequest) Validate() error {
	return validate.Struct(req)
}

func NewDefaultInstance() *Instance {
	req := NewRegistryRequest()
	app := service.NewDefaultService()
	return &Instance{
		Namespace:    app.Spec.Namespace,
		ServiceName:  app.Spec.Name,
		RegistryInfo: req,
		Status:       NewDefaultStatus(),
		Config:       NewDefaultConfig(),
	}
}

func NewInstance(req *RegistryRequest, app *service.Service) (*Instance, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	if req.Name == "" {
		req.Name = "ins-" + token.MakeBearer(8)
	}

	ins := &Instance{
		Domain:       app.Spec.Domain,
		Namespace:    app.Spec.Namespace,
		ServiceName:  app.Spec.Name,
		RegistryInfo: req,
		Status:       NewDefaultStatus(),
		Config:       NewDefaultConfig(),
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
	return fmt.Sprintf("%s.%s-%s", i.Namespace, i.ServiceName,
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
		Labels:      map[string]string{},
		Meta:        map[string]string{},
		Build:       &Build{},
		Provider:    PROVIDER_SDK,
		Protocal:    PROTOCAL_GRPC,
		Weight:      1,
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

func (s *InstanceSet) GetGroupInstance(group string) (items []*Instance) {
	for i := range s.Items {
		if s.Items[i].RegistryInfo.Group == group {
			items = append(items, s.Items[i])
		}
	}

	return
}

func (s *InstanceSet) GetOldestGroup() (items []*Instance) {
	sort.Sort(s)
	if s.Len() > 0 {
		og := s.Items[0].RegistryInfo.Group
		return s.GetGroupInstance(og)
	}

	return nil
}

func (s *InstanceSet) Len() int {
	return len(s.Items)
}

func (s *InstanceSet) Less(i, j int) bool {
	return s.Items[i].Status.Online < s.Items[j].Status.Online
}

func (s *InstanceSet) Swap(i, j int) {
	s.Items[i], s.Items[j] = s.Items[j], s.Items[i]
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
