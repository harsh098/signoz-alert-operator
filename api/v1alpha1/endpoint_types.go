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
)

// SecretRefSpec defines the reference to the
// kubernetes secret and the key needed to
// provide API Key for Signoz Cloud to the
// Operator.
type SecretKeyRefSpec struct {
	// Name of Secret containing the API Key
	// +kubebuilder:validation:MinLength=1
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="API Key Secret Name",xDescriptors={"urn:alm:descriptor:io.kubernetes:Secret"}
	Name string `json:"name"`

	// Key within the Secret's data storing the required
	// API Key
	// +kubebuilder:validation:MinLength=1
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="API Key Secret Key"
	Key string `json:"key"`
}

// EndpointSpec defines the desired state of Endpoint.
type EndpointSpec struct {
	// InstanceURL is the URL Endpoint starting with http:// or https:// for
	// the self-hosted or Signoz Cloud Instance.
	// +kubebuilder:validation:Required
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="SigNoz Instance URL",xDescriptors={"urn:alm:descriptor:org.w3:link"}
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:XValidation:rule="isURL(self) && (url(self).getScheme() == 'http' || url(self).getScheme() == 'https')",message="instanceURL must be a valid http or https URL"
	InstanceURL string `json:"instanceURL"`
	// SecretRef is the reference to the secret containing the API Key for Signoz.
	// Required if using Signoz Cloud.
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,displayName="API Key Secret Reference"
	SecretKeyRef *SecretKeyRefSpec `json:"secretKeyRef,omitempty"`
}

// EndpointStatus defines the observed state of Endpoint.
type EndpointStatus struct{}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Endpoint is the Schema for the endpoints API.
type Endpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EndpointSpec   `json:"spec,omitempty"`
	Status EndpointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EndpointList contains a list of Endpoint.
type EndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Endpoint `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Endpoint{}, &EndpointList{})
}
