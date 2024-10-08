package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/infraboard/mcenter/cmd/initial"
	"github.com/infraboard/mcenter/cmd/start"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/application"
	"github.com/infraboard/mcube/v2/ioc/server"
)

var (
	// pusher service config option
	confType string
	confFile string
)

var vers bool

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "mcenter",
	Short: "微服务公共能力中心",
	Long:  "微服务公共能力中心",
	RunE: func(cmd *cobra.Command, args []string) error {
		if vers {
			fmt.Println(application.FullVersion())
			return nil
		}
		return cmd.Help()
	},
}

// 初始化Ioc
func initail() {
	req := ioc.NewLoadConfigRequest()
	switch confType {
	case "file":
		req.ConfigFile.Enabled = true
		req.ConfigFile.Path = confFile
	default:
		req.ConfigEnv.Enabled = true
		// 调整日志级别
		os.Setenv("LOG_LEVEL", "info")
	}

	err := ioc.ConfigIocObject(req)
	server.DefaultConfig = req
	cobra.CheckErr(err)
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// 初始化设置
	cobra.OnInitialize(initail)
	RootCmd.AddCommand(start.Cmd)
	RootCmd.AddCommand(initial.Cmd)
	err := RootCmd.Execute()
	cobra.CheckErr(err)
}

func init() {
	RootCmd.PersistentFlags().StringVarP(&confType, "config-type", "t", "file", "the service config type [file/env]")
	RootCmd.PersistentFlags().StringVarP(&confFile, "config-file", "f", "etc/config.toml", "the service config from file")
	RootCmd.PersistentFlags().BoolVarP(&vers, "version", "v", false, "the service version")
}
