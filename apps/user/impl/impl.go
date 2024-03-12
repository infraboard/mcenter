package impl

import (
	"context"

	"github.com/infraboard/mcube/v2/ioc"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/infraboard/mcenter/apps/domain"
	"github.com/infraboard/mcenter/apps/user"
	"github.com/infraboard/mcube/v2/ioc/config/grpc"
	"github.com/infraboard/mcube/v2/ioc/config/log"
	ioc_mongo "github.com/infraboard/mcube/v2/ioc/config/mongo"
)

func init() {
	ioc.Controller().Registry(&service{})
}

type service struct {
	log    *zerolog.Logger
	col    *mongo.Collection
	domain domain.Service

	user.UnimplementedRPCServer
	ioc.ObjectImpl
}

func (s *service) Init() error {
	uc := ioc_mongo.DB().Collection("user")

	indexs := []mongo.IndexModel{
		{
			Keys: bson.D{{Key: "create_at", Value: -1}},
		},
		{
			Keys: bson.D{
				{Key: "domain", Value: -1},
				{Key: "username", Value: -1},
			},
			Options: options.Index().SetUnique(true),
		},
	}

	_, err := uc.Indexes().CreateMany(context.Background(), indexs)
	if err != nil {
		return err
	}

	s.col = uc
	s.log = log.Sub(user.AppName)
	s.domain = ioc.Controller().Get(domain.AppName).(domain.Service)

	user.RegisterRPCServer(grpc.Get().Server(), s)
	return nil
}

func (s *service) Name() string {
	return user.AppName
}
