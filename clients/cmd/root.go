package cmd

import (
	"fmt"
	"os"

	"github.com/infraboard/mcenter/clients/rest"
	"github.com/infraboard/mcube/ioc/config/application"
	"github.com/spf13/cobra"
)

var vers bool

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "mctl",
	Short: "Mcenter cli",
	Long:  "Mcenter cli",
	RunE: func(cmd *cobra.Command, args []string) error {
		if vers {
			fmt.Println(application.FullVersion())
			return nil
		}
		return cmd.Help()
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// 加载RPC客户端
	if err := rest.LoadClientFromEnv(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	RootCmd.PersistentFlags().BoolVarP(&vers, "version", "v", false, "the mcenter version")
}
