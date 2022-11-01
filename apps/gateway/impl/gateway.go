package impl

import (
	"context"

	"github.com/infraboard/mcenter/apps/gateway"
	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *impl) CreateGateway(ctx context.Context, req *gateway.CreateGatewayRequest) (
	*gateway.Gateway, error) {
	d, err := gateway.New(req)
	if err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}
	if _, err := s.col.InsertOne(ctx, d); err != nil {
		return nil, exception.NewInternalServerError("inserted a gateway document error, %s", err)
	}

	return d, nil
}

func (s *impl) QueryGateway(ctx context.Context, req *gateway.QueryGatewayRequest) (
	*gateway.GatewaySet, error) {
	r, err := newQueryGatewayRequest(req)
	if err != nil {
		return nil, err
	}

	resp, err := s.col.Find(ctx, r.FindFilter(), r.FindOptions())
	if err != nil {
		return nil, exception.NewInternalServerError("find gateway error, error is %s", err)
	}

	set := gateway.NewGatewaySet()
	// 循环
	if !req.SkipItems {
		for resp.Next(ctx) {
			ins := gateway.NewDefaultGateway()
			if err := resp.Decode(ins); err != nil {
				return nil, exception.NewInternalServerError("decode gateway error, error is %s", err)
			}
			set.Add(ins)
		}
	}

	// count
	count, err := s.col.CountDocuments(ctx, r.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get gateway count error, error is %s", err)
	}
	set.Total = count
	return set, nil
}

func (s *impl) DescribeGateway(ctx context.Context, req *gateway.DescribeGatewayRequest) (
	*gateway.Gateway, error) {
	query, err := newDescribeGatewayRequest(req)
	if err != nil {
		return nil, err
	}

	ins := gateway.NewDefaultGateway()
	if err := s.col.FindOne(ctx, query.FindFilter()).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("gateway %s not found", req)
		}

		return nil, exception.NewInternalServerError("find gateway %s error, %s", req, err)
	}

	return ins, nil
}
