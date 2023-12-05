package impl

import (
	"context"

	"github.com/infraboard/mcenter/apps/endpoint"
	"github.com/infraboard/mcenter/apps/resource"
	"github.com/infraboard/mcube/v2/http/request"
)

const (
	// MaxQueryEndpoints todo
	MaxQueryEndpoints = 1000
)

func (s *impl) QueryResources(ctx context.Context, req *resource.QueryResourceRequest) (
	*resource.ResourceSet, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	rs := resource.NewResourceSet()
	queryE := endpoint.NewQueryEndpointRequest(request.NewPageRequest(MaxQueryEndpoints, 1))
	queryE.PermissionEnable = req.PermissionEnable
	queryE.Resources = req.Resources
	queryE.ServiceIds = req.ServiceIds
	eps, err := s.ep.QueryEndpoints(ctx, queryE)
	if err != nil {
		return nil, err
	}
	if eps.Total > MaxQueryEndpoints {
		s.log.Warn().Msgf("service %s total endpoints > %d", req.ServiceIds, eps.Total)
	}

	rs.AddEndpointSet(eps)
	return rs, nil
}
