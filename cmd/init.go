package cmd

import (
	"context"

	"github.com/spf13/cobra"

	// 注册所有服务
	_ "github.com/infraboard/mcenter/apps/all"
	"github.com/infraboard/mcenter/apps/application"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger/zap"
)

// initCmd represents the start command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "mcenter 服务初始化",
	Long:  "mcenter 服务初始化",
	RunE: func(cmd *cobra.Command, args []string) error {
		// 初始化全局变量
		if err := loadGlobalConfig(confType); err != nil {
			return err
		}

		// 初始化全局日志配置
		if err := loadGlobalLogger(); err != nil {
			return err
		}

		// 初始化全局app
		if err := app.InitAllApp(); err != nil {
			return err
		}

		apps := NewInitApps()
		apps.Add("keyauth", "用户中心")
		apps.Add("maudit", "审计中心")
		apps.Add("demo", "测试样例")
		log := zap.L().Named("init")

		impl := app.GetGrpcApp(application.AppName).(application.ServiceServer)

		for _, req := range apps.items {
			app, err := impl.CreateApplication(context.Background(), req)
			if err != nil {
				log.Errorf("init app %s error, %s", req.Name, err)
				continue
			}
			log.Infof("init app %s success, client_id: %s, client_secret: %s",
				req.Name,
				app.Credential.ClientId,
				app.Credential.ClientSecret,
			)
		}

		return nil
	},
}

func NewInitApps() *InitApps {
	return &InitApps{
		items: []*application.CreateApplicationRequest{},
	}
}

type InitApps struct {
	items []*application.CreateApplicationRequest
}

func (i *InitApps) Add(name, descrption string) {
	req := application.NewCreateApplicationRequest()
	req.Name = name
	req.Description = descrption
	req.Owner = "admin"
	i.items = append(i.items, req)
}

func init() {
	RootCmd.AddCommand(initCmd)
}
