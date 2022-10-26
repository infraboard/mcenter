package instance

import (
	"fmt"

	"github.com/infraboard/mcenter/client/rest"
	"github.com/spf13/cobra"
)

// Cmd represents the start command
var registry = &cobra.Command{
	Use:   "registry",
	Short: "服务实例注册",
	Long:  "服务实例注册",
	RunE: func(cmd *cobra.Command, args []string) error {
		ins, err := rest.C().Instance().RegistryInstance(cmd.Context(), nil)
		if err != nil {
			return err
		}
		fmt.Println(ins)
		return nil
	},
}

func init() {
	Cmd.AddCommand(registry)
}
