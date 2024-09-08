package start

import (
	"github.com/infraboard/mcube/v2/ioc/server"
	"github.com/spf13/cobra"

	// 注册所有服务
	_ "github.com/infraboard/mcenter/apps"
	_ "github.com/infraboard/mcenter/middlewares"

	// 非功能性模块
	_ "github.com/infraboard/mcube/v2/ioc/apps/apidoc/restful"
	_ "github.com/infraboard/mcube/v2/ioc/apps/metric/restful"
	_ "github.com/infraboard/mcube/v2/ioc/config/cors/gorestful"
)

// startCmd represents the start command
var Cmd = &cobra.Command{
	Use:   "start",
	Short: "mcenter API服务",
	Long:  "mcenter API服务",
	Run: func(cmd *cobra.Command, args []string) {
		cobra.CheckErr(server.Run(cmd.Context()))
	},
}
