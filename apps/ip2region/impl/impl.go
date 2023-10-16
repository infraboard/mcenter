package impl

import (
	"sync"

	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/mcenter/apps/ip2region"
	"github.com/infraboard/mcenter/apps/storage"
)

func init() {
	ioc.RegistryController(&service{
		bucketName: "ip2region",
		dbFileName: "ip2region.db",
	})
}

type service struct {
	storage    storage.Service
	log        logger.Logger
	bucketName string
	dbFileName string
	dbReader   *ip2region.IPReader
	sync.Mutex
	ioc.ObjectImpl
}

func (s *service) Init() error {
	s.storage = ioc.GetController(storage.AppName).(storage.Service)

	s.log = zap.L().Named("IP2Region")
	return nil
}

func (s *service) Name() string {
	return ip2region.AppName
}
