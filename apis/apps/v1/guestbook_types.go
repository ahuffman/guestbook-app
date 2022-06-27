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

package v1

import (
	"errors"

	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"
	"github.com/nukleros/operator-builder-tools/pkg/status"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var ErrUnableToConvertGuestbook = errors.New("unable to convert to Guestbook")

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// GuestbookSpec defines the desired state of Guestbook.
type GuestbookSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Namespace string `json:"namespace,omitempty"`

	// +kubebuilder:default=1
	// +kubebuilder:validation:Optional
	// (Default: 1)
	RedisLeaderReplicas int `json:"redisLeaderReplicas,omitempty"`

	RedisLeaderImage string `json:"redisLeaderImage,omitempty"`

	// +kubebuilder:default=6379
	// +kubebuilder:validation:Optional
	// (Default: 6379)
	RedisLeaderContainerPort int `json:"redisLeaderContainerPort,omitempty"`

	// +kubebuilder:default=6379
	// +kubebuilder:validation:Optional
	// (Default: 6379)
	RedisLeaderServicePort int `json:"redisLeaderServicePort,omitempty"`

	// +kubebuilder:default=2
	// +kubebuilder:validation:Optional
	// (Default: 2)
	RedisFollowerReplicas int `json:"redisFollowerReplicas,omitempty"`

	RedisFollowerImage string `json:"redisFollowerImage,omitempty"`

	// +kubebuilder:default=6379
	// +kubebuilder:validation:Optional
	// (Default: 6379)
	RedisFollowerContainerPort int `json:"redisFollowerContainerPort,omitempty"`

	// +kubebuilder:default=6379
	// +kubebuilder:validation:Optional
	// (Default: 6379)
	RedisFollowerServicePort int `json:"redisFollowerServicePort,omitempty"`

	// +kubebuilder:default=3
	// +kubebuilder:validation:Optional
	// (Default: 3)
	GuestBookReplicas int `json:"guestBookReplicas,omitempty"`

	GuestBookImage string `json:"guestBookImage,omitempty"`

	// +kubebuilder:default=80
	// +kubebuilder:validation:Optional
	// (Default: 80)
	GuestBookContainerPort int `json:"guestBookContainerPort,omitempty"`

	// +kubebuilder:default=80
	// +kubebuilder:validation:Optional
	// (Default: 80)
	GuestBookServicePort int `json:"guestBookServicePort,omitempty"`
}

// GuestbookStatus defines the observed state of Guestbook.
type GuestbookStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Created               bool                     `json:"created,omitempty"`
	DependenciesSatisfied bool                     `json:"dependenciesSatisfied,omitempty"`
	Conditions            []*status.PhaseCondition `json:"conditions,omitempty"`
	Resources             []*status.ChildResource  `json:"resources,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster

// Guestbook is the Schema for the guestbooks API.
type Guestbook struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GuestbookSpec   `json:"spec,omitempty"`
	Status            GuestbookStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GuestbookList contains a list of Guestbook.
type GuestbookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Guestbook `json:"items"`
}

// interface methods

// GetReadyStatus returns the ready status for a component.
func (component *Guestbook) GetReadyStatus() bool {
	return component.Status.Created
}

// SetReadyStatus sets the ready status for a component.
func (component *Guestbook) SetReadyStatus(ready bool) {
	component.Status.Created = ready
}

// GetDependencyStatus returns the dependency status for a component.
func (component *Guestbook) GetDependencyStatus() bool {
	return component.Status.DependenciesSatisfied
}

// SetDependencyStatus sets the dependency status for a component.
func (component *Guestbook) SetDependencyStatus(dependencyStatus bool) {
	component.Status.DependenciesSatisfied = dependencyStatus
}

// GetPhaseConditions returns the phase conditions for a component.
func (component *Guestbook) GetPhaseConditions() []*status.PhaseCondition {
	return component.Status.Conditions
}

// SetPhaseCondition sets the phase conditions for a component.
func (component *Guestbook) SetPhaseCondition(condition *status.PhaseCondition) {
	for i, currentCondition := range component.GetPhaseConditions() {
		if currentCondition.Phase == condition.Phase {
			component.Status.Conditions[i] = condition

			return
		}
	}

	// phase not found, lets add it to the list.
	component.Status.Conditions = append(component.Status.Conditions, condition)
}

// GetResources returns the child resource status for a component.
func (component *Guestbook) GetChildResourceConditions() []*status.ChildResource {
	return component.Status.Resources
}

// SetResources sets the phase conditions for a component.
func (component *Guestbook) SetChildResourceCondition(resource *status.ChildResource) {
	for i, currentResource := range component.GetChildResourceConditions() {
		if currentResource.Group == resource.Group && currentResource.Version == resource.Version && currentResource.Kind == resource.Kind {
			if currentResource.Name == resource.Name && currentResource.Namespace == resource.Namespace {
				component.Status.Resources[i] = resource

				return
			}
		}
	}

	// phase not found, lets add it to the collection
	component.Status.Resources = append(component.Status.Resources, resource)
}

// GetDependencies returns the dependencies for a component.
func (*Guestbook) GetDependencies() []workload.Workload {
	return []workload.Workload{}
}

// GetComponentGVK returns a GVK object for the component.
func (*Guestbook) GetWorkloadGVK() schema.GroupVersionKind {
	return GroupVersion.WithKind("Guestbook")
}

func init() {
	SchemeBuilder.Register(&Guestbook{}, &GuestbookList{})
}
