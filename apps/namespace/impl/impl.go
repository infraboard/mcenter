package impl

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/mcube/v2/ioc"
	"github.com/rs/zerolog"

	"github.com/infraboard/mcenter/apps/counter"
	"github.com/infraboard/mcenter/apps/namespace"
	"github.com/infraboard/mcenter/apps/policy"
	"github.com/infraboard/mcenter/apps/role"
	"github.com/infraboard/mcube/v2/ioc/config/grpc"
	"github.com/infraboard/mcube/v2/ioc/config/log"
	ioc_mongo "github.com/infraboard/mcube/v2/ioc/config/mongo"
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
	i.log = log.Sub(i.Name())

	i.role = ioc.Controller().Get(role.AppName).(role.Service)
	i.policy = ioc.Controller().Get(policy.AppName).(policy.Service)
	i.counter = ioc.Controller().Get(counter.AppName).(counter.Service)

	namespace.RegisterRPCServer(grpc.Get().Server(), i)
	return nil
}

func (i *impl) Name() string {
	return namespace.AppName
}
