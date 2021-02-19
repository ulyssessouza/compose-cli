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
	"strings"

	"github.com/compose-spec/compose-go/types"
	"github.com/containerd/containerd"
	"github.com/containerd/containerd/namespaces"
	"github.com/distribution/distribution/v3/reference"
	"golang.org/x/sync/errgroup"

	"github.com/docker/compose-cli/api/progress"
)

func (s *composeService) Pull(ctx context.Context, project *types.Project) error {
	ctx = namespaces.WithNamespace(ctx, "default") // FIXME: HARDCODED!
	w := progress.ContextWriter(ctx)
	eg, ctx := errgroup.WithContext(ctx)

	for _, srv := range project.Services {
		service := srv
		if service.Image == "" {
			w.Event(progress.Event{
				ID:     service.Name,
				Status: progress.Done,
				Text:   "Skipped",
			})
			continue
		}
		eg.Go(func() error {
			ref, err := reference.ParseNormalizedNamed(service.Image)
			if err != nil {
				return err
			}
			fullyQualifiedImageName := ref.String()
			if !strings.Contains(fullyQualifiedImageName, ":") {
				fullyQualifiedImageName = fullyQualifiedImageName + ":latest"
			}
			w.Event(progress.Event{
				ID:     service.Name,
				Status: progress.Working,
				Text:   "Pulling " + fullyQualifiedImageName,
			})
			image, err := s.Client.Pull(ctx, fullyQualifiedImageName,
				containerd.WithPullUnpack,
				containerd.WithMaxConcurrentDownloads(3))
			if err != nil {
				w.Event(progress.Event{
					ID:     service.Name,
					Status: progress.Error,
					Text:   "Error",
				})
				return err
			}
			w.Event(progress.Event{
				ID:     service.Name,
				Status: progress.Done,
				Text:   "Pulled " + image.Name(),
			})
			return nil
		})
	}

	return eg.Wait()
}
