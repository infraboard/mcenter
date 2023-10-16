package impl

import (
	"context"

	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mcube/ioc/config/logger"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"

	"github.com/infraboard/mcenter/apps/domain"
	"github.com/infraboard/mcenter/apps/user"
	ioc_mongo "github.com/infraboard/mcube/ioc/config/mongo"
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
	s.log = logger.Sub(user.AppName)
	s.domain = ioc.GetController(domain.AppName).(domain.Service)
	return nil
}

func (s *service) Name() string {
	return user.AppName
}

func (s *service) Registry(server *grpc.Server) {
	user.RegisterRPCServer(server, s)
}
