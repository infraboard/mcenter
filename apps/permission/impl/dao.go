package impl

import (
	"context"

	"github.com/infraboard/mcenter/apps/permission"
	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *service) queryPermission(ctx context.Context, req *permission.QueryPermissionRequest) (
	*permission.PermissionSet, error) {
	query, err := newQueryPermissionRequest(req)
	if err != nil {
		return nil, err
	}

	s.log.Debugf("query permission filter: %s", query.FindFilter())
	resp, err := s.col.Find(context.TODO(), query.FindFilter(), query.FindOptions())
	if err != nil {
		return nil, exception.NewInternalServerError("find permissionn error, error is %s", err)
	}

	// 循环
	set := permission.NewPermissionSet()
	for resp.Next(context.TODO()) {
		ins := permission.NewDeaultPermission()
		if err := resp.Decode(ins); err != nil {
			return nil, exception.NewInternalServerError("decode permission error, error is %s", err)
		}
		set.Add(ins)
	}

	// count
	count, err := s.col.CountDocuments(context.TODO(), query.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get permission count error, error is %s", err)
	}
	set.Total = count
	return set, nil
}

func newQueryPermissionRequest(req *permission.QueryPermissionRequest) (*queryPermissionRequest, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	return &queryPermissionRequest{
		QueryPermissionRequest: req}, nil
}

type queryPermissionRequest struct {
	*permission.QueryPermissionRequest
}

func (r *queryPermissionRequest) FindOptions() *options.FindOptions {
	pageSize := int64(r.Page.PageSize)
	skip := int64(r.Page.PageSize) * int64(r.Page.PageNumber-1)

	opt := &options.FindOptions{
		Sort:  bson.D{{Key: "create_at", Value: -1}},
		Limit: &pageSize,
		Skip:  &skip,
	}

	return opt
}

func (r *queryPermissionRequest) FindFilter() bson.M {
	filter := bson.M{}

	if r.RoleId != "" {
		filter["role_id"] = r.RoleId
	}

	return filter
}
