package impl

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/mcube/v2/ioc"
	"github.com/rs/zerolog"

	"github.com/infraboard/mcenter/apps/endpoint"
	"github.com/infraboard/mcenter/apps/namespace"
	"github.com/infraboard/mcenter/apps/policy"
	"github.com/infraboard/mcenter/apps/role"
	"github.com/infraboard/mcenter/apps/user"
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
	policy.UnimplementedRPCServer
	ioc.ObjectImpl

	user      user.Service
	role      role.Service
	namespace namespace.Service
	endpoint  endpoint.Service
}

func (i *impl) Init() error {
	i.col = ioc_mongo.DB().Collection(i.Name())
	i.log = log.Sub(i.Name())

	i.user = ioc.Controller().Get(user.AppName).(user.Service)
	i.role = ioc.Controller().Get(role.AppName).(role.Service)
	i.namespace = ioc.Controller().Get(namespace.AppName).(namespace.Service)
	i.endpoint = ioc.Controller().Get(endpoint.AppName).(endpoint.Service)

	policy.RegisterRPCServer(grpc.Get().Server(), i)
	return nil
}

func (i *impl) Name() string {
	return policy.AppName
}
