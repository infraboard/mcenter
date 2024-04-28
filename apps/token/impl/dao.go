package impl

import (
	"context"
	"fmt"
	"time"

	"github.com/infraboard/mcenter/apps/token"
	"github.com/infraboard/mcube/v2/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *service) save(ctx context.Context, tk *token.Token) error {
	if _, err := s.col.InsertOne(ctx, tk); err != nil {
		return exception.NewInternalServerError("inserted token(%s) document error, %s",
			tk.AccessToken, err)
	}

	return nil
}

func (s *service) get(ctx context.Context, id string) (*token.Token, error) {
	filter := bson.M{"_id": id}

	ins := token.NewToken(token.NewIssueTokenRequest())
	if err := s.col.FindOne(ctx, filter).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("token %s not found", id)
		}

		return nil, exception.NewInternalServerError("find token %s error, %s", id, err)
	}

	return ins, nil
}

func (s *service) update(ctx context.Context, tk *token.Token) error {
	_, err := s.col.UpdateOne(ctx, bson.M{"_id": tk.AccessToken}, bson.M{"$set": tk})
	if err != nil {
		return exception.NewInternalServerError("update token(%s) error, %s", tk.AccessToken, err)
	}

	return nil
}

func (s *service) blockOtherWebToken(ctx context.Context, tk *token.Token) error {
	// 共享账号不关闭之前的Token
	if tk.SharedUser {
		return nil
	}

	// 如果不是web登陆, 不需要关闭之前的登录令牌
	if !tk.Platform.Equal(token.PLATFORM_WEB) {
		return nil
	}

	now := time.Now()
	status := token.NewStatus()
	status.IsBlock = true
	status.BlockAt = now.UnixMilli()
	status.BlockReason = fmt.Sprintf("你于 %s 从其他地方通过 %s 登录", now.Format(time.RFC3339), tk.GrantType)
	status.BlockType = token.BLOCK_TYPE_OTHER_PLACE_LOGGED_IN

	rs, err := s.col.UpdateMany(
		ctx,
		bson.M{
			"platform":        token.PLATFORM_WEB,
			"user_id":         tk.UserId,
			"issue_at":        bson.M{"$lt": tk.IssueAt},
			"status.is_block": false,
		},
		bson.M{"$set": bson.M{"status": status}},
	)
	if err != nil {
		return err
	}
	s.log.Debug().Msgf("block %d tokens", rs.ModifiedCount)
	return nil
}

func (s *service) delete(ctx context.Context, ins *token.Token) error {
	if ins == nil || ins.AccessToken == "" {
		return fmt.Errorf("access tpken is nil")
	}

	result, err := s.col.DeleteOne(ctx, bson.M{"_id": ins.AccessToken})
	if err != nil {
		return exception.NewInternalServerError("delete token(%s) error, %s", ins.AccessToken, err)
	}

	if result.DeletedCount == 0 {
		return exception.NewNotFound("book %s not found", ins.AccessToken)
	}

	return nil
}

func newQueryRequest(req *token.QueryTokenRequest) *queryRequest {
	return &queryRequest{req}
}

type queryRequest struct {
	*token.QueryTokenRequest
}

func (r *queryRequest) FindOptions() *options.FindOptions {
	pageSize := int64(r.Page.PageSize)
	skip := int64(r.Page.PageSize) * int64(r.Page.PageNumber-1)

	opt := &options.FindOptions{
		Sort:  bson.D{{Key: "issue_at", Value: -1}},
		Limit: &pageSize,
		Skip:  &skip,
	}

	return opt
}

func (r *queryRequest) FindFilter() bson.M {
	filter := bson.M{}
	if r.GrantType != nil {
		filter["grant_type"] = r.GrantType
	}
	if r.Username != "" {
		filter["username"] = r.Username
	}
	return filter
}

func (s *service) deleteCode(ctx context.Context, ins *token.Code) error {
	result, err := s.col.DeleteOne(ctx, bson.M{"_id": ins.Id})
	if err != nil {
		return exception.NewInternalServerError("delete verify code(%s) error, %s", ins.Code, err)
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("delete verify code %s not found", ins.Code)
	}

	return nil
}
