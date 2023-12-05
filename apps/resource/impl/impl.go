package impl

import (
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"

	"github.com/infraboard/mcenter/apps/endpoint"
	"github.com/infraboard/mcenter/apps/resource"
	"github.com/infraboard/mcube/v2/ioc/config/logger"
	ioc_mongo "github.com/infraboard/mcube/v2/ioc/config/mongo"
)

func init() {
	ioc.Controller().Registry(&impl{})
}

type impl struct {
	col *mongo.Collection
	log *zerolog.Logger
	resource.UnimplementedRPCServer
	ioc.ObjectImpl

	ep endpoint.Service
}

func (i *impl) Init() error {
	i.col = ioc_mongo.DB().Collection(i.Name())
	i.log = logger.Sub(i.Name())

	i.ep = ioc.GetController(endpoint.AppName).(endpoint.Service)
	return nil
}

func (i *impl) Name() string {
	return resource.AppName
}

func (i *impl) Registry(server *grpc.Server) {
	resource.RegisterRPCServer(server, i)
}
