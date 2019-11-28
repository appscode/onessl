/*
Copyright The Kubepack Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmds

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"kmodules.xyz/client-go/tools/backup"
)

func NewCmdBackup(clientGetter genericclioptions.RESTClientGetter) *cobra.Command {
	var (
		ClusterName string
		BackupDir   string
		Sanitize    bool
	)
	cmd := &cobra.Command{
		Use:               "backup",
		Short:             "Backup cluster objects",
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {
			restConfig, err := clientGetter.ToRESTConfig()
			if err != nil {
				Fatal(err)
			}
			mgr := backup.NewBackupManager(ClusterName, restConfig, Sanitize)
			filename, err := mgr.BackupToTar(BackupDir)
			if err != nil {
				Fatal(err)
			}
			fmt.Printf("Cluster objects are stored in %s", filename)
			os.Exit(0)
		},
	}
	cmd.Flags().BoolVar(&Sanitize, "sanitize", Sanitize, " Sanitize fields in YAML")
	cmd.Flags().StringVar(&BackupDir, "backup-dir", BackupDir, "Directory where yaml files will be saved")
	return cmd
}
