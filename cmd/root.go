package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/infraboard/mcenter/cmd/initial"
	"github.com/infraboard/mcenter/cmd/start"
	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mcube/ioc/config/application"
)

var (
	// pusher service config option
	confType string
	confFile string
	confETCD string
)

var vers bool

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "mcenter",
	Short: "基于Etcd的微服务注册中心",
	Long:  "基于Etcd的微服务注册中心",
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
	}
	err := ioc.ConfigIocObject(req)
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
	RootCmd.PersistentFlags().StringVarP(&confType, "config-type", "t", "file", "the service config type [file/env/etcd]")
	RootCmd.PersistentFlags().StringVarP(&confFile, "config-file", "f", "etc/config.toml", "the service config from file")
	RootCmd.PersistentFlags().StringVarP(&confETCD, "config-etcd", "e", "127.0.0.1:2379", "the service config from etcd")
	RootCmd.PersistentFlags().BoolVarP(&vers, "version", "v", false, "the mcenter version")
}
