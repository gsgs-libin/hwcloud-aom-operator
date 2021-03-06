/*
Copyright 2022 dgsfor@gmail.com.

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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type AomGroupRuleSpec struct {
	MaxInstances int `json:"max_instances"`
	MinInstances int `json:"min_instances"`
	CooldownTime int `json:"cooldown_time"`
}

// AomGroupSpec defines the desired state of AomGroup
type AomGroupSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of AomGroup. Edit aomgroup_types.go to remove/update
	Deployment     string           `json:"deployment"`
	Id             string           `json:"id,omitempty"`
	GroupAttribute AomGroupRuleSpec `json:"group_attribute,omitempty"`
}

type AomGroupPhase string

var (
	AomGroupRunning AomGroupPhase = "Running"
	AomGroupFailed  AomGroupPhase = "Failed"
)

type AomGroupCondition struct {
	Ready                 bool        `json:"ready"`
	Message               string      `json:"message"`
	LastedTranslationTime metav1.Time `json:"lastedTranslationTime"`
}

// AomGroupStatus defines the observed state of AomGroup
type AomGroupStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Conditions []AomGroupCondition `json:"conditions"`
	Phase      AomGroupPhase       `json:"phase"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// AomGroup is the Schema for the aomgroups API
type AomGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AomGroupSpec   `json:"spec,omitempty"`
	Status AomGroupStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// AomGroupList contains a list of AomGroup
type AomGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AomGroup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AomGroup{}, &AomGroupList{})
}
