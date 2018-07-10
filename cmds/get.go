package cmds

import (
	"github.com/spf13/cobra"
	"k8s.io/kubernetes/pkg/kubectl/genericclioptions"
)

func NewCmdGet(clientGetter genericclioptions.RESTClientGetter) *cobra.Command {
	cmd := &cobra.Command{
		Use:               "get",
		Short:             `Get stuff`,
		DisableAutoGenTag: true,
	}
	cmd.AddCommand(NewCmdGetCACert())
	cmd.AddCommand(NewCmdGetKubeCA(clientGetter))
	return cmd
}
