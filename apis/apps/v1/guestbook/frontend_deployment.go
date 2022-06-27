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

// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete

const DeploymentNamespaceFrontend = "frontend"

// CreateDeploymentNamespaceFrontend creates the frontend Deployment resource.
func CreateDeploymentNamespaceFrontend(
	parent *appsv1.Guestbook,
) ([]client.Object, error) {

	resourceObjs := []client.Object{}
	var resourceObj = &unstructured.Unstructured{
		Object: map[string]interface{}{
			// SOURCE: https://cloud.google.com/kubernetes-engine/docs/tutorials/guestbook
			"apiVersion": "apps/v1",
			"kind":       "Deployment",
			"metadata": map[string]interface{}{
				"name":      "frontend",
				"namespace": parent.Spec.Namespace, //  controlled by field: namespace
			},
			"spec": map[string]interface{}{
				"replicas": parent.Spec.GuestBookReplicas, //  controlled by field: guestBookReplicas
				"selector": map[string]interface{}{
					"matchLabels": map[string]interface{}{
						"app.kubernetes.io/app":  "guestbook",
						"app.kubernetes.io/tier": "frontend",
					},
				},
				"template": map[string]interface{}{
					"metadata": map[string]interface{}{
						"labels": map[string]interface{}{
							"app.kubernetes.io/app":  "guestbook",
							"app.kubernetes.io/tier": "frontend",
						},
					},
					"spec": map[string]interface{}{
						"containers": []interface{}{
							map[string]interface{}{
								"name":  "php-redis",
								"image": parent.Spec.GuestBookImage, //  controlled by field: guestBookImage
								"env": []interface{}{
									map[string]interface{}{
										"name":  "GET_HOSTS_FROM",
										"value": "dns",
									},
								},
								"resources": map[string]interface{}{
									"requests": map[string]interface{}{
										"cpu":    "100m",
										"memory": "100Mi",
									},
								},
								"ports": []interface{}{
									map[string]interface{}{
										"containerPort": parent.Spec.GuestBookContainerPort, //  controlled by field: guestBookContainerPort
									},
								},
							},
						},
					},
				},
			},
		},
	}

	resourceObjs = append(resourceObjs, resourceObj)

	return resourceObjs, nil
}
