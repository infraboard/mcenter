package impl

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/infraboard/mcenter/apps/role"
	"github.com/infraboard/mcenter/apps/token"
	"github.com/infraboard/mcube/v2/pb/resource"
)

func newDescribeRoleRequest(req *role.DescribeRoleRequest) (*describeRoleRequest, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	return &describeRoleRequest{req}, nil
}

type describeRoleRequest struct {
	*role.DescribeRoleRequest
}

func (req *describeRoleRequest) String() string {
	return fmt.Sprintf("role: %s", req.Name)
}

func (req *describeRoleRequest) FindFilter() bson.M {
	filter := bson.M{}

	if req.Id != "" {
		filter["_id"] = req.Id
	}

	if req.Name != "" {
		filter["name"] = req.Name
	}

	return filter
}

// FindOptions todo
func (req *describeRoleRequest) FindOptions() *options.FindOneOptions {
	opt := &options.FindOneOptions{}

	return opt
}

func newQueryRoleRequest(req *role.QueryRoleRequest) (*queryRoleRequest, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	return &queryRoleRequest{
		QueryRoleRequest: req}, nil
}

type queryRoleRequest struct {
	*role.QueryRoleRequest
}

func (r *queryRoleRequest) FindOptions() *options.FindOptions {
	pageSize := int64(r.Page.PageSize)
	skip := int64(r.Page.PageSize) * int64(r.Page.PageNumber-1)

	opt := &options.FindOptions{
		Sort:  bson.D{{Key: "create_at", Value: -1}},
		Limit: &pageSize,
		Skip:  &skip,
	}

	return opt
}

func (r *queryRoleRequest) FindFilter() bson.M {
	filter := bson.M{}

	if r.Scope != nil {
		token.MakeMongoFilter(filter, r.Scope)

		// 全局共享
		orConditions := bson.A{
			filter,
			bson.M{"visiable": resource.VISIABLE_GLOBAL},
		}

		// 域内公开
		if r.Scope.Domain != "" {
			orConditions = append(orConditions,
				bson.M{"visiable": resource.VISIABLE_DOMAIN, "domain": r.Scope.Domain})
		}

		filter = bson.M{"$or": orConditions}
	}

	return filter
}
