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

package compose

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/docker/compose-cli/api/client"
	"github.com/docker/compose-cli/api/compose"

	"github.com/spf13/cobra"
)

type eventsOpts struct {
	*composeOptions
	json bool
}

func eventsCommand(p *projectOptions) *cobra.Command {
	opts := eventsOpts{
		composeOptions: &composeOptions{
			projectOptions: p,
		},
	}
	cmd := &cobra.Command{
		Use:   "events [options] [--] [SERVICE...]",
		Short: "Receive real time events from containers.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runEvents(cmd.Context(), opts, args)
		},
	}

	cmd.Flags().BoolVar(&opts.json, "json", false, "Output events as a stream of json objects")
	return cmd
}

func runEvents(ctx context.Context, opts eventsOpts, services []string) error {
	c, err := client.New(ctx)
	if err != nil {
		return err
	}

	project, err := opts.toProjectName()
	if err != nil {
		return err
	}

	return c.ComposeService().Events(ctx, project, compose.EventsOptions{
		Services: services,
		Consumer: func(event compose.Event) error {
			if opts.json {
				marshal, err := json.Marshal(map[string]interface{}{
					"time":       event.Timestamp,
					"type":       "container",
					"service":    event.Service,
					"id":         event.Container,
					"action":     event.Status,
					"attributes": event.Attributes,
				})
				if err != nil {
					return err
				}
				fmt.Println(string(marshal))
			} else {
				fmt.Println(event)
			}
			return nil
		},
	})
}
