package impl

import (
	"fmt"

	"github.com/infraboard/mcube/v2/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/infraboard/mcenter/apps/policy"
	"github.com/infraboard/mcenter/apps/token"
)

func newDescribePolicyRequest(req *policy.DescribePolicyRequest) (*describePolicyRequest, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}
	return &describePolicyRequest{req}, nil
}

type describePolicyRequest struct {
	*policy.DescribePolicyRequest
}

func (req *describePolicyRequest) String() string {
	return fmt.Sprintf("policy: %s", req.Id)
}

func (req *describePolicyRequest) FindFilter() bson.M {
	filter := bson.M{}
	if req.Id != "" {
		filter["_id"] = req.Id
	}
	return filter
}

func newQueryPolicyRequest(req *policy.QueryPolicyRequest) (*queryPolicyRequest, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	return &queryPolicyRequest{
		QueryPolicyRequest: req,
	}, nil
}

type queryPolicyRequest struct {
	*policy.QueryPolicyRequest
}

func (r *queryPolicyRequest) FindOptions() *options.FindOptions {
	pageSize := int64(r.Page.PageSize)
	skip := int64(r.Page.PageSize) * int64(r.Page.PageNumber-1)

	opt := &options.FindOptions{
		Sort:  bson.D{{Key: "create_at", Value: -1}},
		Limit: &pageSize,
		Skip:  &skip,
	}

	return opt
}

func (r *queryPolicyRequest) FindFilter() bson.M {
	filter := bson.M{}
	if r.Domain != "" {
		filter["domain"] = r.Domain
	}

	if r.Namespace != "" {
		filter["namespace"] = r.Namespace
	}
	if r.RoleId != "" {
		filter["role_id"] = r.RoleId
	}
	if r.UserId != "" {
		filter["user_id"] = r.UserId
	}

	return filter
}

func newDeletePolicyRequest(req *policy.DeletePolicyRequest) (*deletePolicyRequest, error) {
	return &deletePolicyRequest{
		DeletePolicyRequest: req,
	}, nil
}

type deletePolicyRequest struct {
	*policy.DeletePolicyRequest
}

func (r *deletePolicyRequest) FindFilter() bson.M {
	filter := bson.M{}
	token.MakeMongoFilter(filter, r.Scope)
	filter["_id"] = r.Id
	return filter
}
