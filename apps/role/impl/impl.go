package impl

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/mcube/v2/ioc"
	"github.com/rs/zerolog"

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
	role *mongo.Collection
	perm *mongo.Collection
	log  *zerolog.Logger
	role.UnimplementedRPCServer
	ioc.ObjectImpl

	policy policy.Service
}

func (i *impl) Init() error {
	db := ioc_mongo.DB()
	i.role = db.Collection("role")
	i.perm = db.Collection("permission")

	i.policy = ioc.Controller().Get(policy.AppName).(policy.Service)
	i.log = log.Sub(i.Name())

	role.RegisterRPCServer(grpc.Get().Server(), i)
	return nil
}

func (i *impl) Name() string {
	return role.AppName
}
