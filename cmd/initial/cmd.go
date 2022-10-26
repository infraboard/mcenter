package initial

import (
	"github.com/spf13/cobra"

	// 注册所有服务
	_ "github.com/infraboard/mcenter/apps"
	meta "github.com/infraboard/mcenter/apps/service"
)

// initCmd represents the start command
var Cmd = &cobra.Command{
	Use:   "init",
	Short: "mcenter 服务初始化",
	Long:  "mcenter 服务初始化",
	RunE: func(cmd *cobra.Command, args []string) error {

		exec := newExcutor()
		// 初始化默认域
		if err := exec.InitDomain(cmd.Context()); err != nil {
			return err
		}

		// 初始化内置角色
		if err := exec.InitRole(cmd.Context()); err != nil {
			return err
		}

		// 初始化默认空间
		if err := exec.InitNamespace(cmd.Context()); err != nil {
			return err
		}

		// 初始化内置服务
		if err := exec.InitService(cmd.Context()); err != nil {
			return err
		}

		// 初始化管理员用户

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
