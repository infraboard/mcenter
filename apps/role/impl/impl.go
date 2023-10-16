package impl

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"

	"github.com/infraboard/mcenter/apps/policy"
	"github.com/infraboard/mcenter/apps/role"
	ioc_mongo "github.com/infraboard/mcube/ioc/config/mongo"
)

func init() {
	ioc.Controller().Registry(&impl{})
}

type impl struct {
	role *mongo.Collection
	perm *mongo.Collection
	log  logger.Logger
	role.UnimplementedRPCServer
	ioc.ObjectImpl

	policy policy.Service
}

func (i *impl) Init() error {
	db := ioc_mongo.DB()
	i.role = db.Collection("role")
	i.perm = db.Collection("permission")

	i.policy = ioc.GetController(policy.AppName).(policy.Service)
	i.log = zap.L().Named(i.Name())
	return nil
}

func (i *impl) Name() string {
	return role.AppName
}

func (i *impl) Registry(server *grpc.Server) {
	role.RegisterRPCServer(server, i)
}
