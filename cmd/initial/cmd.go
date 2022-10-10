package initial

import (
	"context"

	"github.com/spf13/cobra"

	// 注册所有服务
	_ "github.com/infraboard/mcenter/apps"
	meta "github.com/infraboard/mcenter/apps/service"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger/zap"
)

// initCmd represents the start command
var Cmd = &cobra.Command{
	Use:   "init",
	Short: "mcenter 服务初始化",
	Long:  "mcenter 服务初始化",
	RunE: func(cmd *cobra.Command, args []string) error {
		apps := NewInitApps()
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
