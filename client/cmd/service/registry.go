package service

import "github.com/spf13/cobra"

// Cmd represents the start command
var registry = &cobra.Command{
	Use:   "registry",
	Short: "服务实例注册",
	Long:  "服务实例注册",
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.Help()
		return nil
	},
}

func init() {
	Cmd.AddCommand(registry)
}
