package impl

import (
	"context"

	"github.com/infraboard/mcube/ioc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"

	"github.com/infraboard/mcenter/apps/domain"
	"github.com/infraboard/mcenter/conf"
)

func init() {
	ioc.Controller().Registry(&service{})
}

type service struct {
	col *mongo.Collection
	domain.UnimplementedRPCServer
	ioc.ObjectImpl
}

func (s *service) Init() error {
	db, err := conf.C().Mongo.GetDB()
	if err != nil {
		return err
	}

	dc := db.Collection(s.Name())
	indexs := []mongo.IndexModel{
		{
			Keys: bson.D{{Key: "create_at", Value: -1}},
		},
		{
			Keys: bson.D{
				{Key: "name", Value: -1},
			},
			Options: options.Index().SetUnique(true),
		},
	}

	_, err = dc.Indexes().CreateMany(context.Background(), indexs)
	if err != nil {
		return err
	}

	s.col = dc

	return nil
}

func (s *service) Name() string {
	return domain.AppName
}

func (s *service) Registry(server *grpc.Server) {
	domain.RegisterRPCServer(server, s)
}
