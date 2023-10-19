package impl

import (
	"sync"

	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mcube/ioc/apps/oss"
	"github.com/infraboard/mcube/ioc/config/logger"
	"github.com/rs/zerolog"

	"github.com/infraboard/mcenter/apps/ip2region"
)

func init() {
	ioc.RegistryController(&service{
		bucketName: "ip2region",
		dbFileName: "ip2region.db",
	})
}

type service struct {
	oss        oss.Service
	log        *zerolog.Logger
	bucketName string
	dbFileName string
	dbReader   *ip2region.IPReader
	sync.Mutex
	ioc.ObjectImpl
}

func (s *service) Init() error {
	s.oss = ioc.GetController(oss.AppName).(oss.Service)

	s.log = logger.Sub("IP2Region")
	return nil
}

func (s *service) Name() string {
	return ip2region.AppName
}
