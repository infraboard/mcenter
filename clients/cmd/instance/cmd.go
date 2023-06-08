package instance

import (
	"github.com/spf13/cobra"
)

// Cmd represents the start command
var Cmd = &cobra.Command{
	Use:   "instance",
	Short: "instance 管理",
	Long:  "instance 管理",
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.Help()
		return nil
	},
}
