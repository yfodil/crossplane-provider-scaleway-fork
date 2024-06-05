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

type DomainInitParameters struct {

	// The ID of the function you want to create a domain with.
	// The ID of the function
	// +crossplane:generate:reference:type=Function
	FunctionID *string `json:"functionId,omitempty" tf:"function_id,omitempty"`

	// Reference to a Function to populate functionId.
	// +kubebuilder:validation:Optional
	FunctionIDRef *v1.Reference `json:"functionIdRef,omitempty" tf:"-"`

	// Selector for a Function to populate functionId.
	// +kubebuilder:validation:Optional
	FunctionIDSelector *v1.Selector `json:"functionIdSelector,omitempty" tf:"-"`

	// The hostname that should resolve to your function id native domain.
	// You should use a CNAME domain record that point to your native function domain_name for it.
	// The hostname that should be redirected to the function
	Hostname *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	// (Defaults to provider region) The region in where the domain was created.
	// The region you want to attach the resource to
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type DomainObservation struct {

	// The ID of the function you want to create a domain with.
	// The ID of the function
	FunctionID *string `json:"functionId,omitempty" tf:"function_id,omitempty"`

	// The hostname that should resolve to your function id native domain.
	// You should use a CNAME domain record that point to your native function domain_name for it.
	// The hostname that should be redirected to the function
	Hostname *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	// The function domain's ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Defaults to provider region) The region in where the domain was created.
	// The region you want to attach the resource to
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The URL that triggers the function
	// URL to use to trigger the function
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type DomainParameters struct {

	// The ID of the function you want to create a domain with.
	// The ID of the function
	// +crossplane:generate:reference:type=Function
	// +kubebuilder:validation:Optional
	FunctionID *string `json:"functionId,omitempty" tf:"function_id,omitempty"`

	// Reference to a Function to populate functionId.
	// +kubebuilder:validation:Optional
	FunctionIDRef *v1.Reference `json:"functionIdRef,omitempty" tf:"-"`

	// Selector for a Function to populate functionId.
	// +kubebuilder:validation:Optional
	FunctionIDSelector *v1.Selector `json:"functionIdSelector,omitempty" tf:"-"`

	// The hostname that should resolve to your function id native domain.
	// You should use a CNAME domain record that point to your native function domain_name for it.
	// The hostname that should be redirected to the function
	// +kubebuilder:validation:Optional
	Hostname *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	// (Defaults to provider region) The region in where the domain was created.
	// The region you want to attach the resource to
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

// DomainSpec defines the desired state of Domain
type DomainSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DomainParameters `json:"forProvider"`
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
	InitProvider DomainInitParameters `json:"initProvider,omitempty"`
}

// DomainStatus defines the observed state of Domain.
type DomainStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DomainObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Domain is the Schema for the Domains API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,scaleway}
type Domain struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.hostname) || (has(self.initProvider) && has(self.initProvider.hostname))",message="spec.forProvider.hostname is a required parameter"
	Spec   DomainSpec   `json:"spec"`
	Status DomainStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DomainList contains a list of Domains
type DomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Domain `json:"items"`
}

// Repository type metadata.
var (
	Domain_Kind             = "Domain"
	Domain_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Domain_Kind}.String()
	Domain_KindAPIVersion   = Domain_Kind + "." + CRDGroupVersion.String()
	Domain_GroupVersionKind = CRDGroupVersion.WithKind(Domain_Kind)
)

func init() {
	SchemeBuilder.Register(&Domain{}, &DomainList{})
}
