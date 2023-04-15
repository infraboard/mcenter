package initial

import (
	"github.com/spf13/cobra"
)

var (
	debug bool
)

// initCmd represents the start command
var Cmd = &cobra.Command{
	Use:   "init",
	Short: "mcenter 服务初始化",
	Long:  "mcenter 服务初始化",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		exec, err := NewExecutorFromCLI()
		cobra.CheckErr(err)

		// 初始化默认域
		err = exec.InitDomain(ctx)
		cobra.CheckErr(err)

		// 初始化管理员用户
		err = exec.InitAdminUser(ctx)
		cobra.CheckErr(err)

		// 初始化默认空间
		err = exec.InitNamespace(ctx)
		cobra.CheckErr(err)

		// 初始化内置角色
		err = exec.InitRole(ctx)
		cobra.CheckErr(err)

		// 初始化内置服务
		err = exec.InitService(ctx)
		cobra.CheckErr(err)

		// 初始化系统设置
		err = exec.InitSystemSetting(ctx)
		cobra.CheckErr(err)
	},
}

func init() {
	Cmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "show debug info")
}
