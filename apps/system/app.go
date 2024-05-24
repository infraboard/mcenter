package system

import (
	"net/http"
	"strings"

	"github.com/infraboard/mcenter/apps/endpoint"
	"github.com/infraboard/mcube/v2/exception"
	request "github.com/infraboard/mcube/v2/http/request"
)

const (
	AppName = "system"
)

// UpdatePath todo
func (r *Resource) UpdatePath(path string) {
	for _, p := range r.Paths {
		if p == path {
			return
		}
	}

	r.Paths = append(r.Paths, path)
}

// UpdateMethod todo
func (r *Resource) UpdateMethod(mothod string) {
	for _, p := range r.Methods {
		if p == mothod {
			return
		}
	}

	r.Methods = append(r.Methods, mothod)
}

// UpdateFunction todo
func (r *Resource) UpdateFunction(fuction string) {
	for _, p := range r.Functions {
		if p == fuction {
			return
		}
	}

	r.Functions = append(r.Functions, fuction)
}

// UpdateAction todo
func (r *Resource) UpdateAction(action string) {
	for _, p := range r.Actions {
		if p == action {
			return
		}
	}

	r.Actions = append(r.Actions, action)
}

// NewResourceSet todo
func NewResourceSet() *ResourceSet {
	return &ResourceSet{
		Items: []*Resource{},
	}
}

// AddEndpointSet todo
func (s *ResourceSet) AddEndpointSet(eps *endpoint.EndpointSet) {
	for i := range eps.Items {
		s.addEndpint(eps.Items[i])
	}
}

func (s *ResourceSet) addEndpint(ep *endpoint.Endpoint) {
	if ep.Entry == nil || ep.Entry.Resource == "" {
		return
	}

	rs := s.getOrCreateResource(ep.ServiceId, ep.Entry.Resource)
	rs.UpdateMethod(ep.Entry.Method)
	rs.UpdatePath(ep.Entry.Path)
	rs.UpdateFunction(ep.Entry.FunctionName)
	if v, ok := ep.Entry.Labels["action"]; ok {
		rs.UpdateAction(v)
	}
}

func (s *ResourceSet) getOrCreateResource(serviceID, name string) *Resource {
	var rs *Resource

	for i := range s.Items {
		rs = s.Items[i]
		if rs.ServiceId == serviceID && rs.Name == name {
			return rs
		}
	}

	// 添加新resource
	rs = &Resource{
		ServiceId: serviceID,
		Name:      name,
	}
	s.Items = append(s.Items, rs)
	return rs
}

// NewQueryResourceRequestFromHTTP 列表查询请求
func NewQueryResourceRequestFromHTTP(r *http.Request) *QueryResourceRequest {
	page := request.NewPageRequestFromHTTP(r)
	qs := r.URL.Query()

	query := &QueryResourceRequest{
		Page: page,
	}

	pe := qs.Get("permission_enable")
	if pe != "" {
		query.EnablePermission(pe == "true")
	}

	strIds := qs.Get("service_ids")
	if strIds != "" {
		query.ServiceIds = strings.Split(strIds, ",")
	}
	rs := qs.Get("resources")
	if rs != "" {
		query.Resources = strings.Split(rs, ",")
	}

	return query
}

func (req *QueryResourceRequest) EnablePermission(v bool) {
	req.PermissionEnable = &v
}

// Validate todo
func (req *QueryResourceRequest) Validate() error {
	if len(req.ServiceIds) == 0 {
		return exception.NewBadRequest("service_ids required, but \"\"")
	}

	return nil
}
