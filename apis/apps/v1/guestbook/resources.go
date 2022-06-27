/*
Copyright 2022.

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

package guestbook

import (
	"fmt"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/yaml"

	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"

	appsv1 "guestbook-app/apis/apps/v1"
)

// sampleGuestbook is a sample containing all fields
const sampleGuestbook = `apiVersion: apps.demo.com/v1
kind: Guestbook
metadata:
  name: guestbook-sample
spec:
  namespace: "guestbook-application"
  redisLeaderReplicas: 1
  redisLeaderImage: "docker.io/redis:6.0.5"
  redisLeaderContainerPort: 6379
  redisLeaderServicePort: 6379
  redisFollowerReplicas: 2
  redisFollowerImage: "gcr.io/google_samples/gb-redis-follower:v2"
  redisFollowerContainerPort: 6379
  redisFollowerServicePort: 6379
  guestBookReplicas: 3
  guestBookImage: "gcr.io/google_samples/gb-frontend:v5"
  guestBookContainerPort: 80
  guestBookServicePort: 80
`

// sampleGuestbookRequired is a sample containing only required fields
const sampleGuestbookRequired = `apiVersion: apps.demo.com/v1
kind: Guestbook
metadata:
  name: guestbook-sample
spec:
  namespace: "guestbook-application"
  redisLeaderImage: "docker.io/redis:6.0.5"
  redisFollowerImage: "gcr.io/google_samples/gb-redis-follower:v2"
  guestBookImage: "gcr.io/google_samples/gb-frontend:v5"
`

// Sample returns the sample manifest for this custom resource.
func Sample(requiredOnly bool) string {
	if requiredOnly {
		return sampleGuestbookRequired
	}

	return sampleGuestbook
}

// Generate returns the child resources that are associated with this workload given
// appropriate structured inputs.
func Generate(workloadObj appsv1.Guestbook) ([]client.Object, error) {
	resourceObjects := []client.Object{}

	for _, f := range CreateFuncs {
		resources, err := f(&workloadObj)

		if err != nil {
			return nil, err
		}

		resourceObjects = append(resourceObjects, resources...)
	}

	return resourceObjects, nil
}

// GenerateForCLI returns the child resources that are associated with this workload given
// appropriate YAML manifest files.
func GenerateForCLI(workloadFile []byte) ([]client.Object, error) {
	var workloadObj appsv1.Guestbook
	if err := yaml.Unmarshal(workloadFile, &workloadObj); err != nil {
		return nil, fmt.Errorf("failed to unmarshal yaml into workload, %w", err)
	}

	if err := workload.Validate(&workloadObj); err != nil {
		return nil, fmt.Errorf("error validating workload yaml, %w", err)
	}

	return Generate(workloadObj)
}

// CreateFuncs is an array of functions that are called to create the child resources for the controller
// in memory during the reconciliation loop prior to persisting the changes or updates to the Kubernetes
// database.
var CreateFuncs = []func(
	*appsv1.Guestbook,
) ([]client.Object, error){
	CreateNamespaceNamespace,
	CreateDeploymentNamespaceRedisLeader,
	CreateServiceNamespaceRedisLeader,
	CreateDeploymentNamespaceRedisFollower,
	CreateServiceNamespaceRedisFollower,
	CreateDeploymentNamespaceFrontend,
	CreateServiceNamespaceFrontend,
}

// InitFuncs is an array of functions that are called prior to starting the controller manager.  This is
// necessary in instances which the controller needs to "own" objects which depend on resources to
// pre-exist in the cluster. A common use case for this is the need to own a custom resource.
// If the controller needs to own a custom resource type, the CRD that defines it must
// first exist. In this case, the InitFunc will create the CRD so that the controller
// can own custom resources of that type.  Without the InitFunc the controller will
// crash loop because when it tries to own a non-existent resource type during manager
// setup, it will fail.
var InitFuncs = []func(
	*appsv1.Guestbook,
) ([]client.Object, error){}

func ConvertWorkload(component workload.Workload) (*appsv1.Guestbook, error) {
	p, ok := component.(*appsv1.Guestbook)
	if !ok {
		return nil, appsv1.ErrUnableToConvertGuestbook
	}

	return p, nil
}
