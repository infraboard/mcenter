package impl

import (
	"fmt"

	"github.com/infraboard/mcenter/apps/role"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func newQueryPermissionRequest(req *role.QueryPermissionRequest) (*queryPermissionRequest, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	return &queryPermissionRequest{
		QueryPermissionRequest: req}, nil
}

type queryPermissionRequest struct {
	*role.QueryPermissionRequest
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

func newDescribePermissionRequest(req *role.DescribePermissionRequest) (*describePermissionRequest, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	return &describePermissionRequest{req}, nil
}

type describePermissionRequest struct {
	*role.DescribePermissionRequest
}

func (req *describePermissionRequest) String() string {
	return fmt.Sprintf("permission: %s", req.Id)
}

func (req *describePermissionRequest) FindFilter() bson.M {
	filter := bson.M{}

	if req.Id != "" {
		filter["_id"] = req.Id
	}

	return filter
}

// FindOptions todo
func (req *describePermissionRequest) FindOptions() *options.FindOneOptions {
	opt := &options.FindOneOptions{}

	return opt
}

func newDeletePermissionRequest(req *role.RemovePermissionFromRoleRequest) (*deletePermissionRequest, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	return &deletePermissionRequest{
		RemovePermissionFromRoleRequest: req}, nil
}

type deletePermissionRequest struct {
	*role.RemovePermissionFromRoleRequest
}

func (r *deletePermissionRequest) FindFilter() bson.M {
	filter := bson.M{}

	filter["role_id"] = r.RoleId
	if !r.RemoveAll {
		filter["_id"] = bson.M{"$in": r.PermissionId}
	}

	return filter
}

func TansferPermissionToDocs(perms []*role.Permission) []interface{} {
	docs := []interface{}{}
	for i := range perms {
		docs = append(docs, perms[i])
	}
	return docs
}
