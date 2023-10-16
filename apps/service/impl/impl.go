package impl

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"

	"github.com/infraboard/mcenter/apps/service"
	"github.com/infraboard/mcenter/conf"
)

func init() {
	ioc.Controller().Registry(&impl{})
}

type impl struct {
	col *mongo.Collection
	log logger.Logger
	service.UnimplementedRPCServer
	ioc.ObjectImpl
}

func (i *impl) Init() error {
	db, err := conf.C().Mongo.GetDB()
	if err != nil {
		return err
	}
	i.col = db.Collection(i.Name())

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
	_, err = i.col.Indexes().CreateMany(context.Background(), indexs)
	if err != nil {
		return err
	}

	i.log = zap.L().Named(i.Name())
	return nil
}

func (i *impl) Name() string {
	return service.AppName
}

func (i *impl) Registry(server *grpc.Server) {
	service.RegisterRPCServer(server, i)
}
