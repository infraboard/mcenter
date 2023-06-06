package impl

import (
	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/mcenter/apps/storage"
	"github.com/infraboard/mcenter/conf"
)

var (
	// Service 服务实例
	svr = &service{}
)

type service struct {
	log logger.Logger
	db  *mongo.Database
	ioc.IocObjectImpl
}

func (s *service) Init() error {
	db, err := conf.C().Mongo.GetDB()
	if err != nil {
		return err
	}

	s.db = db
	s.log = zap.L().Named("Storage")
	return nil
}

func (s *service) Name() string {
	return storage.AppName
}

func init() {
	ioc.RegistryController(svr)
}
