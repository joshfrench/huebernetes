/*
Copyright 2021 Josh French

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type Metadata struct {
	//+kubebuilder:validation:MinLength=1
	//+kubebuilder:validation:MaxLength=32
	Name string `json:"name"`
}

type On struct {
	On bool `json:"on"`
}

// LightSpec defines the desired state of Light
type LightSpec struct {
	//+kubebuilder:validation:Pattern=`^[0-9a-f]{8}-([0-9a-f]{4}-){3}[0-9a-f]{12}$`
	Id       string `json:"id"`
	Metadata `json:"metadata"`
	On       `json:"on"`
}

// LightStatus defines the observed state of Light
type LightStatus struct {
	On `json:"on"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Light is the Schema for the lights API
type Light struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LightSpec   `json:"spec,omitempty"`
	Status LightStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// LightList contains a list of Light
type LightList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Light `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Light{}, &LightList{})
}
