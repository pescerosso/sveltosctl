/*
Copyright 2022. projectsveltos.io. All rights reserved.

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

package utils

import (
	"context"

	"github.com/go-logr/logr"

	configv1alpha1 "github.com/projectsveltos/cluster-api-feature-manager/api/v1alpha1"
	"github.com/projectsveltos/sveltosctl/internal/logs"
)

// ListClusterReports returns all current ClusterProfiles
func (a *k8sAccess) ListClusterProfiles(ctx context.Context,
	logger logr.Logger) (*configv1alpha1.ClusterProfileList, error) {

	logger.V(logs.LogVerbose).Info("Get all ClusterProfiles")
	clusterProfiles := &configv1alpha1.ClusterProfileList{}
	err := a.client.List(ctx, clusterProfiles)
	return clusterProfiles, err
}
