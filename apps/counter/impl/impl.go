package impl

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/infraboard/mcenter/apps/counter"
	"github.com/infraboard/mcube/v2/ioc"
	ioc_mongo "github.com/infraboard/mcube/v2/ioc/config/mongo"
)

func init() {
	ioc.Controller().Registry(&service{})
}

type service struct {
	col *mongo.Collection
	ioc.ObjectImpl
}

func (s *service) Init() error {
	s.col = ioc_mongo.DB().Collection("counter")
	return nil
}

func (s *service) GetNextSequenceValue(sequenceName string) (*counter.Count, error) {
	result := s.col.FindOneAndUpdate(
		context.Background(),
		bson.M{"_id": sequenceName},
		bson.M{"$inc": bson.M{"value": 1}},
		options.FindOneAndUpdate().SetUpsert(true),
	)

	count := counter.NewCount()
	err := result.Decode(count)
	if err != nil && err != mongo.ErrNoDocuments {
		return nil, fmt.Errorf("counter decode error, %s", err)
	}

	count.Value++
	return count, nil
}

func (s *service) Name() string {
	return counter.AppName
}
