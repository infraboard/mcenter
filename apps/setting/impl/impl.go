package impl

import (
	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/infraboard/mcenter/apps/setting"
	"github.com/infraboard/mcenter/conf"
)

func init() {
	ioc.RegistryController(&service{})
}

type service struct {
	log logger.Logger
	col *mongo.Collection
	ioc.IocObjectImpl
}

func (s *service) Init() error {
	db, err := conf.C().Mongo.GetDB()
	if err != nil {
		return err
	}

	s.log = zap.L().Named(s.Name())
	s.col = db.Collection(s.Name())
	return nil
}

func (s *service) Name() string {
	return setting.AppName
}
