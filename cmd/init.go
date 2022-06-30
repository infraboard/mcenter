package cmd

import (
	"context"

	"github.com/spf13/cobra"

	// 注册所有服务
	_ "github.com/infraboard/mcenter/apps/all"
	meta "github.com/infraboard/mcenter/apps/service"
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
		apps.Add("cmdb", "资源中心")
		log := zap.L().Named("init")

		impl := app.GetInternalApp(meta.AppName).(meta.MetaService)

		for _, req := range apps.items {
			app, err := impl.CreateService(context.Background(), req)
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
		items: []*meta.CreateServiceRequest{},
	}
}

type InitApps struct {
	items []*meta.CreateServiceRequest
}

func (i *InitApps) Add(name, descrption string) {
	req := meta.NewCreateServiceRequest()
	req.Name = name
	req.Description = descrption
	req.Owner = "admin"
	i.items = append(i.items, req)
}

func init() {
	RootCmd.AddCommand(initCmd)
}
