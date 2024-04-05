// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type SSHKeyInitParameters struct {

	// The SSH key status
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// The name of the SSH key.
	// The name of the iam SSH key
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Defaults to provider project_id) The ID of the project the SSH key is associated with.
	// The project_id you want to attach the resource to
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// The public SSH key to be added.
	// The public SSH key
	PublicKey *string `json:"publicKey,omitempty" tf:"public_key,omitempty"`
}

type SSHKeyObservation struct {

	// The date and time of the creation of the iam SSH Key
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// The SSH key status
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// The fingerprint of the iam SSH key
	Fingerprint *string `json:"fingerprint,omitempty" tf:"fingerprint,omitempty"`

	// The ID of the SSH key (UUID format).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the SSH key.
	// The name of the iam SSH key
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The organization ID the SSH key is associated with.
	// The organization_id you want to attach the resource to
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	// (Defaults to provider project_id) The ID of the project the SSH key is associated with.
	// The project_id you want to attach the resource to
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// The public SSH key to be added.
	// The public SSH key
	PublicKey *string `json:"publicKey,omitempty" tf:"public_key,omitempty"`

	// The date and time of the last update of the iam SSH Key
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`
}

type SSHKeyParameters struct {

	// The SSH key status
	// +kubebuilder:validation:Optional
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// The name of the SSH key.
	// The name of the iam SSH key
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Defaults to provider project_id) The ID of the project the SSH key is associated with.
	// The project_id you want to attach the resource to
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// The public SSH key to be added.
	// The public SSH key
	// +kubebuilder:validation:Optional
	PublicKey *string `json:"publicKey,omitempty" tf:"public_key,omitempty"`
}

// SSHKeySpec defines the desired state of SSHKey
type SSHKeySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SSHKeyParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider SSHKeyInitParameters `json:"initProvider,omitempty"`
}

// SSHKeyStatus defines the observed state of SSHKey.
type SSHKeyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SSHKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SSHKey is the Schema for the SSHKeys API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,scaleway}
type SSHKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.publicKey) || (has(self.initProvider) && has(self.initProvider.publicKey))",message="spec.forProvider.publicKey is a required parameter"
	Spec   SSHKeySpec   `json:"spec"`
	Status SSHKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SSHKeyList contains a list of SSHKeys
type SSHKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SSHKey `json:"items"`
}

// Repository type metadata.
var (
	SSHKey_Kind             = "SSHKey"
	SSHKey_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SSHKey_Kind}.String()
	SSHKey_KindAPIVersion   = SSHKey_Kind + "." + CRDGroupVersion.String()
	SSHKey_GroupVersionKind = CRDGroupVersion.WithKind(SSHKey_Kind)
)

func init() {
	SchemeBuilder.Register(&SSHKey{}, &SSHKeyList{})
}
