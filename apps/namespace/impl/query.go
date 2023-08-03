package impl

import (
	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/infraboard/mcenter/apps/namespace"
)

func newPaggingQuery(req *namespace.QueryNamespaceRequest) *queryNamespaceRequest {
	return &queryNamespaceRequest{
		QueryNamespaceRequest: req,
	}
}

type queryNamespaceRequest struct {
	*namespace.QueryNamespaceRequest
	namespaces []string
}

func (r *queryNamespaceRequest) AddNamespace(ns []string) {
	// 如果是*表示无需过滤
	for _, v := range ns {
		if v == "*" {
			return
		}
	}

	r.namespaces = ns
}

func (r *queryNamespaceRequest) FindOptions() *options.FindOptions {
	pageSize := int64(r.Page.PageSize)
	skip := int64(r.Page.PageSize) * int64(r.Page.PageNumber-1)

	opt := &options.FindOptions{
		Sort:  bson.D{{Key: "create_at", Value: -1}},
		Limit: &pageSize,
		Skip:  &skip,
	}

	return opt
}

func (r *queryNamespaceRequest) FindFilter() bson.M {
	filter := bson.M{}

	if len(r.namespaces) > 0 {
		filter["name"] = bson.M{"$in": r.namespaces}
	}

	return filter
}

func newDescribeQuery(req *namespace.DescriptNamespaceRequest) (*describeNamespaceRequest, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	return &describeNamespaceRequest{req}, nil
}

type describeNamespaceRequest struct {
	*namespace.DescriptNamespaceRequest
}

func (r *describeNamespaceRequest) FindFilter() bson.M {
	filter := bson.M{}

	switch r.DescribeBy {
	case namespace.DESCRIBE_BY_NAME:
		filter["domain"] = r.Domain
		filter["name"] = r.Name
	default:
		filter["_id"] = r.Id
	}

	return filter
}

func newDeleteRequest(req *namespace.DeleteNamespaceRequest) (*deleteNamespaceRequest, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	return &deleteNamespaceRequest{
		DeleteNamespaceRequest: req,
	}, nil
}

type deleteNamespaceRequest struct {
	*namespace.DeleteNamespaceRequest
}

func (r *deleteNamespaceRequest) FindFilter() bson.M {
	filter := bson.M{
		"domain": r.Domain,
		"_id":    r.Name,
	}

	return filter
}
