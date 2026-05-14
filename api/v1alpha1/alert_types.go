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
	signoz "github.com/harsh098/signoz-alert-operator/internal/signozclient"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EndpointReference names an Endpoint resource in the same namespace as the
// referring Alert.
type EndpointReference struct {
	// name of the referenced Endpoint.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=253
	Name string `json:"name"`
}

// AlertSpec defines the desired state of Alert.
type AlertSpec struct {
	// endpointRef references the Endpoint pointing at the SigNoz instance
	// this Alert is registered against.
	// +kubebuilder:validation:Required
	EndpointRef EndpointReference `json:"endpointRef"`

	// rule is the SigNoz rule body, transmitted verbatim to
	// `POST /api/v2/rules` (`RuletypesPostableRule`). See the SigNoz API
	// schema reference for field definitions and validation rules.
	//
	// This version of the operator currently supports the following schema
	// for alert rules:
	//
	// Visit: https://signoz.io/api-reference/v0.122.0#/operations/CreateRule
	//
	// Validated at admission time by the operator's validating webhook
	// against the generated `RuletypesPostableRule` schema; malformed
	// payloads are rejected before persistence. The field is declared
	// Schemaless in the CRD because `RuletypesPostableRule` uses oneOf
	// and nullable constructs that are not expressible in Kubernetes CRD
	// OpenAPI v3.
	//
	// MUST be treated as immutable by reconciler code: the generated
	// DeepCopy for `RuletypesPostableRule` is a shallow struct copy that
	// shares nested maps and pointers with the informer cache.
	//
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Schemaless
	// +kubebuilder:validation:Type=object
	// +kubebuilder:pruning:PreserveUnknownFields
	Rule signoz.RuletypesPostableRule `json:"rule"`
}

// AlertStatus defines the observed state of Alert.
type AlertStatus struct {
	// observedGeneration is the most recent .metadata.generation reconciled
	// by the controller.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// ruleID is the identifier SigNoz assigns to the rule created from this
	// Alert. The controller uses it to reconcile subsequent updates and
	// deletions, so once set it must not be cleared except on rule deletion.
	// +optional
	RuleID string `json:"ruleID,omitempty"`

	// conditions represent the current state of the Alert resource.
	// +listType=map
	// +listMapKey=type
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Namespaced,shortName=sza
// +kubebuilder:printcolumn:name="Endpoint",type=string,JSONPath=".spec.endpointRef.name"
// +kubebuilder:printcolumn:name="Rule ID",type=string,JSONPath=".status.ruleID"
// +kubebuilder:printcolumn:name="Ready",type=string,JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=".metadata.creationTimestamp"

// Alert is the Schema for the alerts API.
type Alert struct {
	metav1.TypeMeta `json:",inline"`

	// metadata is standard object metadata.
	// +optional
	metav1.ObjectMeta `json:"metadata,omitzero"`

	// spec is the desired state of the Alert.
	// +required
	Spec AlertSpec `json:"spec"`

	// status is the observed state of the Alert.
	// +optional
	Status AlertStatus `json:"status,omitzero"`
}

// +kubebuilder:object:root=true

// AlertList contains a list of Alerts.
type AlertList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitzero"`
	Items           []Alert `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Alert{}, &AlertList{})
}
