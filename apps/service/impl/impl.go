package impl

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"

	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"

	"github.com/infraboard/mcenter/apps/service"
	"github.com/infraboard/mcenter/conf"
)

var (
	// Service 服务实例
	svr = &impl{}
)

type impl struct {
	col *mongo.Collection
	log logger.Logger
	service.UnimplementedRPCServer
	ioc.IocObjectImpl
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
			Keys: bsonx.Doc{{Key: "create_at", Value: bsonx.Int32(-1)}},
		},
		{
			Keys: bsonx.Doc{
				{Key: "domain", Value: bsonx.Int32(-1)},
				{Key: "namespace", Value: bsonx.Int32(-1)},
				{Key: "name", Value: bsonx.Int32(-1)},
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
	service.RegisterRPCServer(server, svr)
}

func init() {
	ioc.RegistryController(svr)
}
