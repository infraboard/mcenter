package initial

import (
	"github.com/spf13/cobra"

	"github.com/infraboard/mcenter/apps/service"
)

// initCmd represents the start command
var Cmd = &cobra.Command{
	Use:   "init",
	Short: "mcenter 服务初始化",
	Long:  "mcenter 服务初始化",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()

		exec := newExcutor()
		// 初始化默认域
		if err := exec.InitDomain(ctx); err != nil {
			return err
		}

		// 初始化内置角色
		if err := exec.InitRole(ctx); err != nil {
			return err
		}

		// 初始化默认空间
		if err := exec.InitNamespace(ctx); err != nil {
			return err
		}

		// 初始化内置服务
		if err := exec.InitService(ctx); err != nil {
			return err
		}

		// 初始化系统设置
		if err := exec.InitSystemSetting(ctx); err != nil {
			return err
		}

		// 初始化管理员用户
		if err := exec.InitAdminUser(ctx); err != nil {
			return err
		}

		return nil
	},
}

func NewInitApps() *InitApps {
	return &InitApps{
		items: []*service.CreateServiceRequest{},
	}
}

type InitApps struct {
	items []*service.CreateServiceRequest
}

func (i *InitApps) Add(name, descrption string) {
	req := service.NewCreateServiceRequest()
	req.Name = name
	req.Description = descrption
	req.Owner = "admin"
	i.items = append(i.items, req)
}
