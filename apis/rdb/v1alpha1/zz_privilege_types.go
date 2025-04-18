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

type PrivilegeInitParameters struct {

	// Name of the database (e.g. my-db-name).
	// Database name
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name,omitempty"`

	// UUID of the Database Instance.
	// Instance on which the database is created
	// +crossplane:generate:reference:type=github.com/scaleway/crossplane-provider-scaleway/apis/rdb/v1alpha1.Instance
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Reference to a Instance in rdb to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// Selector for a Instance in rdb to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// Permission to set. Valid values are readonly, readwrite, all, custom and none.
	// Privilege
	Permission *string `json:"permission,omitempty" tf:"permission,omitempty"`

	// (Defaults to provider region) The region in which the resource exists.
	// The region you want to attach the resource to
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Name of the user (e.g. my-db-user).
	// User name
	UserName *string `json:"userName,omitempty" tf:"user_name,omitempty"`
}

type PrivilegeObservation struct {

	// Name of the database (e.g. my-db-name).
	// Database name
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name,omitempty"`

	// The ID of the user privileges, which is of the form {region}/{instance_id}/{database_name}/{user_name}, e.g. fr-par/11111111-1111-1111-1111-111111111111/database_name/foo
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// UUID of the Database Instance.
	// Instance on which the database is created
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Permission to set. Valid values are readonly, readwrite, all, custom and none.
	// Privilege
	Permission *string `json:"permission,omitempty" tf:"permission,omitempty"`

	// (Defaults to provider region) The region in which the resource exists.
	// The region you want to attach the resource to
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Name of the user (e.g. my-db-user).
	// User name
	UserName *string `json:"userName,omitempty" tf:"user_name,omitempty"`
}

type PrivilegeParameters struct {

	// Name of the database (e.g. my-db-name).
	// Database name
	// +kubebuilder:validation:Optional
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name,omitempty"`

	// UUID of the Database Instance.
	// Instance on which the database is created
	// +crossplane:generate:reference:type=github.com/scaleway/crossplane-provider-scaleway/apis/rdb/v1alpha1.Instance
	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Reference to a Instance in rdb to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// Selector for a Instance in rdb to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// Permission to set. Valid values are readonly, readwrite, all, custom and none.
	// Privilege
	// +kubebuilder:validation:Optional
	Permission *string `json:"permission,omitempty" tf:"permission,omitempty"`

	// (Defaults to provider region) The region in which the resource exists.
	// The region you want to attach the resource to
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Name of the user (e.g. my-db-user).
	// User name
	// +kubebuilder:validation:Optional
	UserName *string `json:"userName,omitempty" tf:"user_name,omitempty"`
}

// PrivilegeSpec defines the desired state of Privilege
type PrivilegeSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PrivilegeParameters `json:"forProvider"`
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
	InitProvider PrivilegeInitParameters `json:"initProvider,omitempty"`
}

// PrivilegeStatus defines the observed state of Privilege.
type PrivilegeStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PrivilegeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Privilege is the Schema for the Privileges API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,scaleway}
type Privilege struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.databaseName) || (has(self.initProvider) && has(self.initProvider.databaseName))",message="spec.forProvider.databaseName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.permission) || (has(self.initProvider) && has(self.initProvider.permission))",message="spec.forProvider.permission is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.userName) || (has(self.initProvider) && has(self.initProvider.userName))",message="spec.forProvider.userName is a required parameter"
	Spec   PrivilegeSpec   `json:"spec"`
	Status PrivilegeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PrivilegeList contains a list of Privileges
type PrivilegeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Privilege `json:"items"`
}

// Repository type metadata.
var (
	Privilege_Kind             = "Privilege"
	Privilege_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Privilege_Kind}.String()
	Privilege_KindAPIVersion   = Privilege_Kind + "." + CRDGroupVersion.String()
	Privilege_GroupVersionKind = CRDGroupVersion.WithKind(Privilege_Kind)
)

func init() {
	SchemeBuilder.Register(&Privilege{}, &PrivilegeList{})
}
