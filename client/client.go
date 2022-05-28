package client

import (
	kc "github.com/infraboard/keyauth/client"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/infraboard/mcenter/apps/instance"
)

var (
	client *ClientSet
)

// SetGlobal todo
func SetGlobal(cli *ClientSet) {
	client = cli
}

// C Global
func C() *ClientSet {
	return client
}

// NewClient todo
func NewClient(conf *kc.Config) (*ClientSet, error) {
	zap.DevelopmentSetup()
	log := zap.L()

	conn, err := grpc.Dial(
		conf.Address(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithPerRPCCredentials(conf.Authentication),
	)
	if err != nil {
		return nil, err
	}

	return &ClientSet{
		conn: conn,
		log:  log,
	}, nil
}

// Client 客户端
type ClientSet struct {
	conn *grpc.ClientConn
	log  logger.Logger
}

// Instance服务的SDK
func (c *ClientSet) Instance() instance.ServiceClient {
	return instance.NewServiceClient(c.conn)
}
