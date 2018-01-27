package cmds

import (
	"github.com/spf13/cobra"
)

func NewCmdGet() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "get",
		Short:             `Get stuff`,
		DisableAutoGenTag: true,
	}
	cmd.AddCommand(NewCmdGetCACert())
	cmd.AddCommand(NewCmdGetKubeCA())
	return cmd
}
