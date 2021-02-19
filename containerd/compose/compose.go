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

	"github.com/docker/compose-cli/api/compose"
	"github.com/docker/compose-cli/api/errdefs"

	"github.com/compose-spec/compose-go/types"
	"github.com/sanathkr/go-yaml"

	"github.com/containerd/containerd"
)

type composeService struct {
	*containerd.Client
}

const (
	defaultContainerdSocketPath = "/run/containerd/containerd.sock"
)

// NewComposeService create a local implementation of the compose.Service API
func NewComposeService() (compose.Service, error) {
	client, err := containerd.New(defaultContainerdSocketPath)
	if err != nil {
		return nil, err
	}
	return &composeService{
		Client: client,
	}, nil
}

func (s *composeService) Pause(ctx context.Context, project *types.Project) error {
	return errdefs.ErrNotImplemented
}

func (s *composeService) UnPause(ctx context.Context, project *types.Project) error {
	return errdefs.ErrNotImplemented
}

func (s *composeService) Start(ctx context.Context, project *types.Project, options compose.StartOptions) error {
	return errdefs.ErrNotImplemented
}

func (s *composeService) Build(ctx context.Context, project *types.Project) error {
	return errdefs.ErrNotImplemented
}

func (s *composeService) Push(ctx context.Context, project *types.Project) error {
	return errdefs.ErrNotImplemented
}

func (s *composeService) Create(ctx context.Context, project *types.Project, opts compose.CreateOptions) error {
	return errdefs.ErrNotImplemented
}

func (s *composeService) Stop(ctx context.Context, project *types.Project, options compose.StopOptions) error {
	return errdefs.ErrNotImplemented
}

func (s *composeService) Down(ctx context.Context, projectName string, options compose.DownOptions) error {
	return errdefs.ErrNotImplemented
}

func (s *composeService) Logs(ctx context.Context, projectName string, consumer compose.LogConsumer, options compose.LogOptions) error {
	return errdefs.ErrNotImplemented
}

func (s *composeService) Ps(ctx context.Context, projectName string, options compose.PsOptions) ([]compose.ContainerSummary, error) {
	return nil, errdefs.ErrNotImplemented
}

func (s *composeService) List(ctx context.Context, options compose.ListOptions) ([]compose.Stack, error) {
	return nil, errdefs.ErrNotImplemented
}

func (s *composeService) Kill(ctx context.Context, project *types.Project, options compose.KillOptions) error {
	return errdefs.ErrNotImplemented
}

func (s *composeService) RunOneOffContainer(ctx context.Context, project *types.Project, opts compose.RunOptions) (int, error) {
	return 0, errdefs.ErrNotImplemented
}

func (s *composeService) Remove(ctx context.Context, project *types.Project, options compose.RemoveOptions) ([]string, error) {
	return nil, errdefs.ErrNotImplemented
}

func (s *composeService) Exec(ctx context.Context, project *types.Project, opts compose.RunOptions) error {
	return errdefs.ErrNotImplemented
}

func (s *composeService) Up(ctx context.Context, project *types.Project, options compose.UpOptions) error {
	return errdefs.ErrNotImplemented
}

func (s *composeService) Convert(ctx context.Context, project *types.Project, options compose.ConvertOptions) ([]byte, error) {
	switch options.Format {
	case "json":
		return json.MarshalIndent(project, "", "  ")
	case "yaml":
		return yaml.Marshal(project)
	default:
		return nil, fmt.Errorf("unsupported format %q", options)
	}
}
