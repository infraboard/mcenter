package instance

import (
	"crypto/md5"
	"fmt"
	"hash/fnv"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcenter/apps/service"
	"github.com/infraboard/mcube/v2/http/request"
	"github.com/infraboard/mcube/v2/tools/pretty"
)

const (
	AppName = "instances"
)

const (
	DEFAULT_PROVIDER = "default"
	DEFAULT_REGION   = "default"
	DEFAULT_ENV      = "default"
	DEFAULT_CLUSTER  = "default"
	DEFAULT_GROUP    = "default"
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
	req.Cluster = qs.Get("cluster")
	req.Group = qs.Get("group")
	req.Labels = ParseStrLable(qs.Get("labels"))
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

func (req *SearchRequest) ToJSON() string {
	return pretty.ToJSON(req)
}

func NewDefaultInstance() *Instance {
	req := NewRegistryRequest()
	svc := service.NewDefaultService()
	return &Instance{
		Namespace:    svc.Spec.Namespace,
		ServiceName:  svc.Spec.Name,
		RegistryInfo: req,
		Status:       NewDefaultStatus(),
		Config:       NewDefaultConfig(),
	}
}

func NewInstance(req *RegistryRequest, svc *service.Service) (*Instance, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	ins := &Instance{
		Domain:       svc.Spec.Domain,
		Namespace:    svc.Spec.Namespace,
		ServiceName:  svc.Spec.Name,
		RegistryInfo: req,
		Status:       NewDefaultStatus(),
		Config:       NewDefaultConfig(),
	}

	ins.Id = "ins-" + ins.FullNameHash()
	ins.Config.Weight = req.Weight
	return ins, nil
}

func NewDefaultConfig() *Config {
	return &Config{
		Enabled:   true,
		Weight:    1,
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
		Region:      DEFAULT_REGION,
		Environment: DEFAULT_ENV,
		Cluster:     DEFAULT_CLUSTER,
		Group:       DEFAULT_GROUP,
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

func (s *InstanceSet) RegistryInfoHash() string {
	h := md5.New()
	for i := range s.Items {
		ins := s.Items[i]
		h.Write([]byte(ins.RegistryInfo.String()))
	}
	return fmt.Sprintf("%x", h.Sum(nil))
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

// k=v,k=v
func ParseStrLable(str string) map[string]string {
	m := map[string]string{}
	str = strings.TrimSpace(str)
	if str != "" {
		kvs := strings.Split(str, ",")
		for i := range kvs {
			kv := kvs[i]
			item := strings.Split(kv, "=")
			key := item[0]
			if len(item) > 1 {
				m[key] = strings.Join(item[1:], "")
			}
		}
	}
	return m
}
