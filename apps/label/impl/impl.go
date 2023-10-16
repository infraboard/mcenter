package impl

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/mcube/ioc"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"

	"github.com/infraboard/mcenter/apps/label"
	"github.com/infraboard/mcube/ioc/config/logger"
	ioc_mongo "github.com/infraboard/mcube/ioc/config/mongo"
)

func init() {
	ioc.Controller().Registry(&impl{})
}

type impl struct {
	col *mongo.Collection
	log *zerolog.Logger
	label.UnimplementedRPCServer
	ioc.ObjectImpl
}

func (i *impl) Init() error {
	i.col = ioc_mongo.DB().Collection(i.Name())
	i.log = logger.Sub(i.Name())
	return nil
}

func (i *impl) Name() string {
	return label.AppName
}

func (i *impl) Registry(server *grpc.Server) {
	label.RegisterRPCServer(server, i)
}
