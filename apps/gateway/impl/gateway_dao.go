package impl

import (
	"github.com/infraboard/mcenter/apps/gateway"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func newQueryGatewayRequest(req *gateway.QueryGatewayRequest) (*queryGatewayRequest, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	return &queryGatewayRequest{
		QueryGatewayRequest: req}, nil
}

type queryGatewayRequest struct {
	*gateway.QueryGatewayRequest
}

func (r *queryGatewayRequest) FindOptions() *options.FindOptions {
	pageSize := int64(r.Page.PageSize)
	skip := int64(r.Page.PageSize) * int64(r.Page.PageNumber-1)

	opt := &options.FindOptions{
		Sort:  bson.D{{Key: "create_at", Value: -1}},
		Limit: &pageSize,
		Skip:  &skip,
	}

	return opt
}

func (r *queryGatewayRequest) FindFilter() bson.M {
	filter := bson.M{}

	return filter
}

func newDescribeGatewayRequest(req *gateway.DescribeGatewayRequest) (*describeGatewayRequest, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	return &describeGatewayRequest{req}, nil
}

type describeGatewayRequest struct {
	*gateway.DescribeGatewayRequest
}

func (req *describeGatewayRequest) FindFilter() bson.M {
	filter := bson.M{}

	if req.Id != "" {
		filter["_id"] = req.Id
	}

	return filter
}
