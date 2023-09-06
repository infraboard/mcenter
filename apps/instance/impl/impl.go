package impl

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"

	"github.com/infraboard/mcenter/apps/instance"
	"github.com/infraboard/mcenter/apps/service"
	"github.com/infraboard/mcenter/conf"
)

func init() {
	ioc.Controller().Registry(&impl{})
}

type impl struct {
	col *mongo.Collection
	log logger.Logger
	instance.UnimplementedRPCServer
	ioc.IocObjectImpl

	svc service.MetaService
}

func (i *impl) Init() error {
	db, err := conf.C().Mongo.GetDB()
	if err != nil {
		return err
	}
	i.col = db.Collection(i.Name())
	i.log = zap.L().Named(i.Name())

	i.svc = ioc.GetController(service.AppName).(service.MetaService)
	return nil
}

func (i *impl) Name() string {
	return instance.AppName
}

func (i *impl) Registry(server *grpc.Server) {
	instance.RegisterRPCServer(server, i)
}
