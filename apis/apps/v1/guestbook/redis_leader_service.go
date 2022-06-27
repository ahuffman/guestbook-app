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
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	appsv1 "guestbook-app/apis/apps/v1"
)

// +kubebuilder:rbac:groups=core,resources=services,verbs=get;list;watch;create;update;patch;delete

const ServiceNamespaceRedisLeader = "redis-leader"

// CreateServiceNamespaceRedisLeader creates the redis-leader Service resource.
func CreateServiceNamespaceRedisLeader(
	parent *appsv1.Guestbook,
) ([]client.Object, error) {

	resourceObjs := []client.Object{}
	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// SOURCE: https://cloud.google.com/kubernetes-engine/docs/tutorials/guestbook
			"apiVersion": "v1",
			"kind":       "Service",
			"metadata": map[string]interface{}{
				"name": "redis-leader",
				"labels": map[string]interface{}{
					"app.kubernetes.io/app":  "redis",
					"app.kubernetes.io/role": "leader",
					"app.kubernetes.io/tier": "backend",
				},
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
			},
			"spec": map[string]interface{}{
				"ports": []interface{}{
					map[string]interface{}{
						"port":       parent.Spec.RedisLeaderServicePort,   //  controlled by field: redisLeaderServicePort
						"targetPort": parent.Spec.RedisLeaderContainerPort, //  controlled by field: redisLeaderContainerPort
					},
				},
				"selector": map[string]interface{}{
					"app.kubernetes.io/app":  "redis",
					"app.kubernetes.io/role": "leader",
					"app.kubernetes.io/tier": "backend",
				},
			},
		},
	}

	resourceObjs = append(resourceObjs, resourceObj)

	return resourceObjs, nil
}
