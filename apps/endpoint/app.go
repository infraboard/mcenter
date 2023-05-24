package endpoint

import (
	"fmt"
	"hash/fnv"
	"net/http"
	"strings"

	"github.com/emicklei/go-restful/v3"
	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/http/label"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/types/ftime"
)

const (
	AppName = "endpoint"
)

// NewDefaultEndpoint todo
func NewDefaultEndpoint() *Endpoint {
	return &Endpoint{}
}

// NewEndpoint todo
func NewEndpoint(serviceID, version string, entry *Entry) *Endpoint {
	return &Endpoint{
		Id:        GenHashID(serviceID, entry.Path),
		CreateAt:  ftime.Now().Timestamp(),
		UpdateAt:  ftime.Now().Timestamp(),
		ServiceId: serviceID,
		Version:   version,
		Entry:     entry,
	}
}

// GenHashID hash id
func GenHashID(service, grpcPath string) string {
	hashedStr := fmt.Sprintf("%s-%s", service, grpcPath)
	h := fnv.New32a()
	h.Write([]byte(hashedStr))
	return fmt.Sprintf("%x", h.Sum32())
}

// LabelsToStr 扁平化标签  action:get;action:list;action-list-echo
func (e *Endpoint) LabelsToStr() string {
	labels := make([]string, 0, len(e.Entry.Labels))
	for k, v := range e.Entry.Labels {
		labels = append(labels, fmt.Sprintf("%s:%s;", k, v))
	}
	return strings.Join(labels, "")
}

// ParseLabels 解析Str格式的label
func (e *Endpoint) ParseLabels(labels string) error {
	kvs := strings.Split(strings.TrimSuffix(labels, ";"), ";")
	for _, kv := range kvs {
		kvItem := strings.Split(kv, ":")
		if len(kvItem) != 2 {
			return fmt.Errorf("labels format error, format: k:v;k:v;")
		}
		e.Entry.Labels[kvItem[0]] = kvItem[1]
	}
	return nil
}

// NewEndpointSet 实例化
func NewEndpointSet() *EndpointSet {
	return &EndpointSet{
		Items: []*Endpoint{},
	}
}

// Add 添加
func (s *EndpointSet) Add(e *Endpoint) {
	s.Items = append(s.Items, e)
}

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

// NewRegistryRequest 注册请求
func NewRegistryRequest(version string, entries []*Entry) *RegistryRequest {
	return &RegistryRequest{
		Version: version,
		Entries: entries,
	}
}

// NewDefaultRegistryRequest todo
func NewDefaultRegistryRequest() *RegistryRequest {
	return &RegistryRequest{
		Entries: []*Entry{},
	}
}

// Validate 校验注册请求合法性
func (req *RegistryRequest) Validate() error {
	if len(req.Entries) == 0 {
		return fmt.Errorf("must require *router.Entry")
	}

	return validate.Struct(req)
}

// Endpoints 功能列表
func (req *RegistryRequest) Endpoints(serviceID string) []*Endpoint {
	eps := make([]*Endpoint, 0, len(req.Entries))
	for i := range req.Entries {
		ep := &Endpoint{
			Id:        GenHashID(serviceID, req.Entries[i].UniquePath()),
			CreateAt:  ftime.Now().Timestamp(),
			UpdateAt:  ftime.Now().Timestamp(),
			ServiceId: serviceID,
			Version:   req.Version,
			Entry:     req.Entries[i],
		}
		eps = append(eps, ep)
	}
	return eps
}

// NewRegistryResponse todo
func NewRegistryResponse(message string) *RegistryResponse {
	return &RegistryResponse{Message: message}
}

// NewQueryEndpointRequestFromHTTP 列表查询请求
func NewQueryEndpointRequestFromHTTP(r *http.Request) *QueryEndpointRequest {
	page := request.NewPageRequestFromHTTP(r)
	qs := r.URL.Query()

	query := &QueryEndpointRequest{
		Page:         page,
		Path:         qs.Get("path"),
		Method:       qs.Get("method"),
		FunctionName: qs.Get("function_name"),
	}

	sids := qs.Get("service_ids")
	if sids != "" {
		query.ServiceIds = strings.Split(sids, ",")
	}
	rs := qs.Get("resources")
	if rs != "" {
		query.Resources = strings.Split(rs, ",")
	}

	return query
}

// NewQueryEndpointRequest 列表查询请求
func NewQueryEndpointRequest(pageReq *request.PageRequest) *QueryEndpointRequest {
	return &QueryEndpointRequest{
		Page: pageReq,
	}
}

// NewDescribeEndpointRequestWithID todo
func NewDescribeEndpointRequestWithID(id string) *DescribeEndpointRequest {
	return &DescribeEndpointRequest{Id: id}
}

// Validate 校验
func (req *DescribeEndpointRequest) Validate() error {
	if req.Id == "" {
		return fmt.Errorf("endpoint id is required")
	}

	return nil
}

// NewDeleteEndpointRequestWithServiceID todo
func NewDeleteEndpointRequestWithServiceID(id string) *DeleteEndpointRequest {
	return &DeleteEndpointRequest{ServiceId: id}
}

// 用于Route转换成Entry
func TransferRoutesToEntry(routes []restful.Route) (entries []*Entry) {
	for _, r := range routes {
		entries = append(entries, NewEntryFromRestRoute(r))
	}
	return
}

func NewDefaultEntry() *Entry {
	return &Entry{
		Allow:     []string{},
		Labels:    map[string]string{},
		Extension: map[string]string{},
	}
}

func (e *Entry) LoadMeta(meta map[string]interface{}) {
	if meta != nil {
		if v, ok := meta[label.Resource]; ok {
			e.Resource, _ = v.(string)
		}
		if v, ok := meta[label.Auth]; ok {
			e.AuthEnable, _ = v.(bool)
		}
		if v, ok := meta[label.Audit]; ok {
			e.AuditLog, _ = v.(bool)
		}
		if v, ok := meta[label.Permission]; ok {
			e.PermissionEnable, _ = v.(bool)
		}
		if v, ok := meta[label.Code]; ok {
			e.CodeEnable, _ = v.(bool)
		}

		if v, ok := meta[label.Action]; ok {
			e.Labels["action"], _ = v.(string)
		}
	}
}

// UniquePath todo
func (e *Entry) UniquePath() string {
	return fmt.Sprintf("%s.%s", e.Method, e.Path)
}

func (e *Entry) IsAllow(target fmt.Stringer) bool {
	for i := range e.Allow {
		if e.Allow[i] == "*" {
			return true
		}

		if e.Allow[i] == target.String() {
			return true
		}
	}

	return false
}

func NewEntryFromRestRequest(req *restful.Request) *Entry {
	entry := NewDefaultEntry()

	// 请求拦截
	route := req.SelectedRoute()
	if route == nil {
		return nil
	}

	entry.FunctionName = route.Operation()
	entry.Method = route.Method()
	entry.LoadMeta(route.Metadata())
	entry.Path = route.Path()

	entry.Path = entry.UniquePath()
	return entry
}

func NewEntryFromRestRoute(route restful.Route) *Entry {
	entry := NewDefaultEntry()
	entry.FunctionName = route.Operation
	entry.Method = route.Method
	entry.LoadMeta(route.Metadata)
	entry.Path = route.Path

	entry.Path = entry.UniquePath()
	return entry
}
