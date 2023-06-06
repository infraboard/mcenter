package start

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/infraboard/mcube/ioc"
	"github.com/spf13/cobra"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/mcenter/conf"
	"github.com/infraboard/mcenter/protocol"

	// 注册所有服务
	_ "github.com/infraboard/mcenter/apps"
)

// startCmd represents the start command
var Cmd = &cobra.Command{
	Use:   "start",
	Short: "mcenter API服务",
	Long:  "mcenter API服务",
	RunE: func(cmd *cobra.Command, args []string) error {
		conf := conf.C()
		// 初始化服务
		svr, err := newService(conf)
		if err != nil {
			return err
		}

		// 启动服务
		svr.start()
		return nil
	},
}

func newService(cnf *conf.Config) (*service, error) {
	http := protocol.NewHTTPService()
	grpc := protocol.NewGRPCService()
	svr := &service{
		http: http,
		grpc: grpc,
		log:  zap.L().Named("CLI"),
	}

	return svr, nil
}

type service struct {
	http *protocol.HTTPService
	grpc *protocol.GRPCService

	log logger.Logger
}

func (s *service) start() {
	s.log.Infof("loaded controllers: %s", ioc.ListControllerObjectNames())
	s.log.Infof("loaded apis: %s", ioc.ListApiObjectNames())
	go s.grpc.Start()
	go s.http.Start()

	// 处理信号量
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP, syscall.SIGQUIT)
	s.waitSign(ch)
}

func (s *service) waitSign(sign chan os.Signal) {
	for sg := range sign {
		switch v := sg.(type) {
		default:
			s.log.Infof("receive signal '%v', start graceful shutdown", v.String())

			if err := s.grpc.Stop(); err != nil {
				s.log.Errorf("grpc graceful shutdown err: %s, force exit", err)
			} else {
				s.log.Info("grpc service stop complete")
			}

			if err := s.http.Stop(); err != nil {
				s.log.Errorf("http graceful shutdown err: %s, force exit", err)
			} else {
				s.log.Infof("http service stop complete")
			}
			os.Exit(0)
		}
	}
}
