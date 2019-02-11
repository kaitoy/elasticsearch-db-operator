/*
Copyright 2019 kaitoy.

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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// TemplateSpec defines the desired state of Template
// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-templates.html
type TemplateSpec struct {
	IndexPatterns []string           `json:"index_patterns"`
	Settings      *Settings          `json:"settings,omitempty"`
	Order         int32              `json:"order,omitempty"`
	Version       int32              `json:"version,omitempty"`
	Mappings      map[string]Mapping `json:"mappings"`
}

// TemplateStatus defines the observed state of Template
type TemplateStatus struct {
	Conditions []TemplateCondition `json:"conditions"`
}

// TemplateCondition represents information about the status of a mapping
type TemplateCondition struct {
	StatusCode         int         `json:"statusCode"`
	Status             string      `json:"status"`
	LastProbeTime      metav1.Time `json:"lastProbeTime"`
	LastTransitionTime metav1.Time `json:"lastTransitionTime"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Template is the Schema for the templates API
// +k8s:openapi-gen=true
type Template struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	URL struct {
		ElasticsearchEndpoint string `json:"elasticsearchEndpoint"`
		Template              string `json:"template"`
	} `json:"url"`
	Spec   TemplateSpec   `json:"spec,omitempty"`
	Status TemplateStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TemplateList contains a list of Template
type TemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Template `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Template{}, &TemplateList{})
}
