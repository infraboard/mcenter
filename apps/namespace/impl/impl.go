package impl

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/mcube/ioc"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"

	"github.com/infraboard/mcenter/apps/counter"
	"github.com/infraboard/mcenter/apps/namespace"
	"github.com/infraboard/mcenter/apps/policy"
	"github.com/infraboard/mcenter/apps/role"
	"github.com/infraboard/mcube/ioc/config/logger"
	ioc_mongo "github.com/infraboard/mcube/ioc/config/mongo"
)

func init() {
	ioc.Controller().Registry(&impl{})
}

type impl struct {
	col *mongo.Collection
	log *zerolog.Logger
	namespace.UnimplementedRPCServer
	ioc.ObjectImpl

	counter counter.Service
	role    role.Service
	policy  policy.Service
}

func (i *impl) Init() error {
	i.col = ioc_mongo.DB().Collection(i.Name())
	i.log = logger.Sub(i.Name())

	i.role = ioc.GetController(role.AppName).(role.Service)
	i.policy = ioc.GetController(policy.AppName).(policy.Service)
	i.counter = ioc.GetController(counter.AppName).(counter.Service)
	return nil
}

func (i *impl) Name() string {
	return namespace.AppName
}

func (i *impl) Registry(server *grpc.Server) {
	namespace.RegisterRPCServer(server, i)
}
