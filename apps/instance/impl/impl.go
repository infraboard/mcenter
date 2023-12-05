package impl

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/mcube/v2/ioc"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"

	"github.com/infraboard/mcenter/apps/instance"
	"github.com/infraboard/mcenter/apps/service"
	"github.com/infraboard/mcube/v2/ioc/config/logger"
	ioc_mongo "github.com/infraboard/mcube/v2/ioc/config/mongo"
)

func init() {
	ioc.Controller().Registry(&impl{})
}

type impl struct {
	col *mongo.Collection
	log *zerolog.Logger
	instance.UnimplementedRPCServer
	ioc.ObjectImpl

	svc service.MetaService
}

func (i *impl) Init() error {
	i.col = ioc_mongo.DB().Collection(i.Name())
	i.log = logger.Sub(i.Name())

	i.svc = ioc.GetController(service.AppName).(service.MetaService)
	return nil
}

func (i *impl) Name() string {
	return instance.AppName
}

func (i *impl) Registry(server *grpc.Server) {
	instance.RegisterRPCServer(server, i)
}
