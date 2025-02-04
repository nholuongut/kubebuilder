/*
Copyright 2020 The Nho Luong DevOps.

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

package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func (c CLI) newVersionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "version",
		Short:   fmt.Sprintf("Print the %s version", c.commandName),
		Long:    fmt.Sprintf("Print the %s version", c.commandName),
		Example: fmt.Sprintf("%s version", c.commandName),
		RunE: func(_ *cobra.Command, _ []string) error {
			fmt.Println(c.version)
			return nil
		},
	}
	return cmd
}
