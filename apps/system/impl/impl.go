package impl

import (
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"

	"github.com/infraboard/mcenter/apps/endpoint"
	"github.com/infraboard/mcenter/apps/system"
	"github.com/infraboard/mcube/v2/ioc/config/log"
	ioc_mongo "github.com/infraboard/mcube/v2/ioc/config/mongo"
)

func init() {
	ioc.Controller().Registry(&impl{})
}

type impl struct {
	system.UnimplementedRPCServer
	ioc.ObjectImpl

	ep  endpoint.Service
	col *mongo.Collection
	log *zerolog.Logger
}

func (i *impl) Init() error {
	i.col = ioc_mongo.DB().Collection(i.Name())
	i.log = log.Sub(i.Name())

	i.ep = ioc.Controller().Get(endpoint.AppName).(endpoint.Service)
	return nil
}

func (i *impl) Name() string {
	return system.AppName
}

func (i *impl) Registry(server *grpc.Server) {
	system.RegisterRPCServer(server, i)
}
