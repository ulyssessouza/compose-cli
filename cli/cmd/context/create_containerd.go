/*
   Copyright 2020 Docker Compose CLI authors

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

package context

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/docker/compose-cli/api/context/store"
)

func init() {
	extraCommands = append(extraCommands, createContainerdCommand)
	extraHelp = append(extraHelp, `
Create a Containerd context:
$ docker context create containerd CONTEXT [flags]
(see docker context create containerd --help)
`)
}

func createContainerdCommand() *cobra.Command {
	var opts descriptionCreateOpts
	cmd := &cobra.Command{
		Use:   "containerd CONTEXT [flags]",
		Short: "Create context to target Containerd",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return createContainerdContext(cmd.Context(), args[0], store.ContainerdContextType, opts.description, store.ContainerdContext{})
		},
	}
	addDescriptionFlag(cmd, &opts.description)
	return cmd
}

func createContainerdContext(ctx context.Context, name string, contextType string, description string, data interface{}) error {
	s := store.ContextStore(ctx)
	result := s.Create(
		name,
		contextType,
		description,
		data,
	)
	fmt.Printf("Successfully created %s context %q\n", contextType, name)
	return result
}
