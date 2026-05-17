/*
Copyright 2026.

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
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// EndpointRefSpec stores reference to the
// object of kind Endpoint in the same namespace
// as the Alert
type EndpointRefSpec struct {
	// Name refers to the name of the kubernetes
	// object of kind Endpoint in the same namespace
	// as the Alert
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Required
	Name string `json:"name"`
}

// AlertSpec defines the desired state of Alert.
type AlertSpec struct {
	// EndpointRef holds reference to the kind Endpoint
	// the points at the target Signoz Instance
	// +kubebuilder:validation:Required
	EndpointRef EndpointRefSpec `json:"endpointRef"`
	// Rule is the SigNoz alert rule body (matches RuletypesPostableRule from
	// the SigNoz API). Schema-validated server-side only at the top level —
	// contents are forwarded verbatim to SigNoz.
	// The current controller is wire compatible with
	// https://signoz.io/api-reference/v0.124.0#/operations/CreateRule
	// +kubebuilder:validation:Schemaless
	// +kubebuilder:pruning:PreserveUnknownFields
	// +kubebuilder:validation:Required
	Rule runtime.RawExtension `json:"rule"`
}

// AlertStatus defines the observed state of Alert.
type AlertStatus struct {
	// RuleID is the SigNoz-assigned rule ID. Empty until first successful create.
	// +optional
	RuleID string `json:"ruleID,omitempty"`

	// HTTPStatus is the HTTP status code from the most recent SigNoz API call.
	// +optional
	HTTPStatus int `json:"httpStatus,omitempty"`

	// Errors is the response body from the most recent non-2xx SigNoz API call. Empty on success.
	// +optional
	Errors string `json:"errors,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Alert is the Schema for the alerts API.
type Alert struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AlertSpec   `json:"spec,omitempty"`
	Status AlertStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AlertList contains a list of Alert.
type AlertList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Alert `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Alert{}, &AlertList{})
}
