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
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	version "gomodules.xyz/version"
)

func NewCmdSemver() *cobra.Command {
	var (
		minor bool
		check string
	)
	cmd := &cobra.Command{
		Use:               "semver",
		Short:             "Print sanitized semver version",
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 1 {
				Fatal(errors.Errorf("multiple version found: %v", strings.Join(args, ",")))
			}
			if len(args) == 0 {
				Fatal(errors.Errorf("missing version"))
			}
			gitVersion := args[0]

			gv, err := version.NewVersion(gitVersion)
			if err != nil {
				Fatal(errors.Wrapf(err, "invalid version %s", gitVersion))
			}
			m := gv.ToMutator().ResetMetadata().ResetPrerelease()
			if minor {
				m = m.ResetPatch()
			}
			if check == "" {
				fmt.Print(m.String())
				return
			}

			c, err := version.NewConstraint(check)
			if err != nil {
				Fatal(errors.Wrapf(err, "invalid constraint %s", gitVersion))
			}
			if !c.Check(m.Done()) {
				os.Exit(1)
			}
		},
	}

	cmd.Flags().BoolVar(&minor, "minor", minor, "print major.minor.0 version")
	cmd.Flags().StringVar(&check, "check", check, "check constraint")
	return cmd
}
