package impl

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/infraboard/mcube/ioc"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"

	"github.com/infraboard/mcenter/apps/service"
	"github.com/infraboard/mcube/ioc/config/logger"
	ioc_mongo "github.com/infraboard/mcube/ioc/config/mongo"
)

func init() {
	ioc.Controller().Registry(&impl{})
}

type impl struct {
	col *mongo.Collection
	log *zerolog.Logger
	service.UnimplementedRPCServer
	ioc.ObjectImpl
}

func (i *impl) Init() error {
	i.col = ioc_mongo.DB().Collection(i.Name())

	// 创建索引
	indexs := []mongo.IndexModel{
		{
			Keys: bson.D{{Key: "create_at", Value: -1}},
		},
		{
			Keys: bson.D{
				{Key: "domain", Value: -1},
				{Key: "namespace", Value: -1},
				{Key: "name", Value: -1},
			},
			Options: options.Index().SetUnique(true),
		},
	}
	_, err := i.col.Indexes().CreateMany(context.Background(), indexs)
	if err != nil {
		return err
	}

	i.log = logger.Sub(i.Name())
	return nil
}

func (i *impl) Name() string {
	return service.AppName
}

func (i *impl) Registry(server *grpc.Server) {
	service.RegisterRPCServer(server, i)
}
