/*
Copyright 2023.

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

// FulldeploymentSpec defines the desired state of Fulldeployment
type FulldeploymentSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Fulldeployment. Edit fulldeployment_types.go to remove/update
	ContainerImage      string `json:"containerImage,omitempty"`
	ContainerTag        string `json:"containerTag,omitempty"`
	ContainerEntrypoint string `json:"containerEntrypoint,omitempty"`
	Appurl string `json:"appurl,omitempty"`
	Databaseurl string `json:"databaseurl,omitempty"`

}

// FulldeploymentStatus defines the observed state of Fulldeployment
type FulldeploymentStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	ClientStatus string `json:"clientStatus,omitempty"`
	LastPodName  string `json:"lastPodName,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Fulldeployment is the Schema for the fulldeployments API
type Fulldeployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FulldeploymentSpec   `json:"spec,omitempty"`
	Status FulldeploymentStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// FulldeploymentList contains a list of Fulldeployment
type FulldeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Fulldeployment `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Fulldeployment{}, &FulldeploymentList{})
}
