package impl

import (
	"context"

	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"

	"github.com/infraboard/mcenter/apps/domain"
	"github.com/infraboard/mcenter/apps/user"
	"github.com/infraboard/mcenter/conf"
)

func init() {
	ioc.Controller().Registry(&service{})
}

type service struct {
	log    logger.Logger
	col    *mongo.Collection
	domain domain.Service

	user.UnimplementedRPCServer
	ioc.ObjectImpl
}

func (s *service) Init() error {
	db, err := conf.C().Mongo.GetDB()
	if err != nil {
		return err
	}

	uc := db.Collection("user")

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

	_, err = uc.Indexes().CreateMany(context.Background(), indexs)
	if err != nil {
		return err
	}

	s.col = uc
	s.log = zap.L().Named(user.AppName)
	s.domain = ioc.GetController(domain.AppName).(domain.Service)
	return nil
}

func (s *service) Name() string {
	return user.AppName
}

func (s *service) Registry(server *grpc.Server) {
	user.RegisterRPCServer(server, s)
}
