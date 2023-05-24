package impl

import (
	"context"

	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/mcenter/apps/endpoint"
	"github.com/infraboard/mcenter/apps/service"
)

func (s *impl) DescribeEndpoint(ctx context.Context, req *endpoint.DescribeEndpointRequest) (
	*endpoint.Endpoint, error) {
	r, err := newDescribeEndpointRequest(req)
	if err != nil {
		return nil, err
	}

	ins := endpoint.NewDefaultEndpoint()
	if err := s.col.FindOne(ctx, r.FindFilter()).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("endpoint %s not found", req)
		}

		return nil, exception.NewInternalServerError("find endpoint %s error, %s", req.Id, err)
	}

	return ins, nil
}

func (s *impl) QueryEndpoints(ctx context.Context, req *endpoint.QueryEndpointRequest) (
	*endpoint.EndpointSet, error) {
	r := newQueryEndpointRequest(req)
	resp, err := s.col.Find(ctx, r.FindFilter(), r.FindOptions())

	if err != nil {
		return nil, exception.NewInternalServerError("find endpoint error, error is %s", err)
	}

	set := endpoint.NewEndpointSet()
	// 循环
	for resp.Next(ctx) {
		app := endpoint.NewDefaultEndpoint()
		if err := resp.Decode(app); err != nil {
			return nil, exception.NewInternalServerError("decode domain error, error is %s", err)
		}

		set.Add(app)
	}

	// count
	count, err := s.col.CountDocuments(ctx, r.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get device count error, error is %s", err)
	}
	set.Total = count

	return set, nil
}

func (s *impl) RegistryEndpoint(ctx context.Context, req *endpoint.RegistryRequest) (*endpoint.RegistryResponse, error) {
	if req.ClientId == "" && req.ClientSecret == "" {
		req.ClientId, req.ClientSecret = service.GetClientCredential(ctx)
	}

	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	// 查询该服务
	svr, err := s.app.DescribeService(ctx, service.NewDescribeServiceRequestByClientId(req.ClientId))
	if err != nil {
		return nil, err
	}
	s.log.Debugf("service %s registry endpoints", svr.Spec.Name)

	if err := svr.Credential.Validate(req.ClientSecret); err != nil {
		return nil, err
	}

	// 生成该服务的Endpoint
	endpoints := req.Endpoints(svr.Meta.Id)
	s.log.Debugf("registry endpoints: %s", endpoints)

	// 更新已有的记录
	news := make([]interface{}, 0, len(endpoints))
	for i := range endpoints {
		if err := s.col.FindOneAndReplace(ctx, bson.M{"_id": endpoints[i].Id}, endpoints[i]).Err(); err != nil {
			if err == mongo.ErrNoDocuments {
				news = append(news, endpoints[i])
			} else {
				return nil, err
			}
		}
	}

	// 插入新增记录
	if len(news) > 0 {
		if _, err := s.col.InsertMany(ctx, news); err != nil {
			return nil, exception.NewInternalServerError("inserted a service document error, %s", err)
		}
	}

	return endpoint.NewRegistryResponse("ok"), nil
}

func (s *impl) DeleteEndpoint(ctx context.Context, req *endpoint.DeleteEndpointRequest) (*endpoint.Endpoint, error) {
	result, err := s.col.DeleteOne(ctx, bson.M{"service_id": req.ServiceId})
	if err != nil {
		return nil, exception.NewInternalServerError("delete service(%s) endpoint error, %s", req.ServiceId, err)
	}

	s.log.Infof("delete service %s endpoint success, total count: %d", req.ServiceId, result.DeletedCount)
	return nil, nil
}
