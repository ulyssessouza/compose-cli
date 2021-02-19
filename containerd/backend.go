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

package containerd

import (
	"context"

	"github.com/docker/compose-cli/api/backend"
	"github.com/docker/compose-cli/api/cloud"
	"github.com/docker/compose-cli/api/compose"
	"github.com/docker/compose-cli/api/containers"
	"github.com/docker/compose-cli/api/context/store"
	"github.com/docker/compose-cli/api/resources"
	"github.com/docker/compose-cli/api/secrets"
	"github.com/docker/compose-cli/api/volumes"
	containerdCompose "github.com/docker/compose-cli/containerd/compose"
)

type containerdAPIService struct {
	composeService compose.Service
}

func init() {
	backend.Register(store.ContainerdContextType, store.ContainerdContextType, service, cloud.NotImplementedCloudService)
}

func service(ctx context.Context) (backend.Service, error) {
	s, err := containerdCompose.NewComposeService()
	if err != nil {
		return nil, err
	}
	return &containerdAPIService{
		composeService: s,
	}, nil
}

func (s *containerdAPIService) ContainerService() containers.Service {
	return nil
}

func (s *containerdAPIService) ComposeService() compose.Service {
	return s.composeService
}

func (s *containerdAPIService) SecretsService() secrets.Service {
	return nil
}

func (s *containerdAPIService) VolumeService() volumes.Service {
	return nil
}

func (s *containerdAPIService) ResourceService() resources.Service {
	return nil
}
