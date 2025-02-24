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

type UserDataInitParameters struct {

	// Key of the user data.
	// The key of the user data to set.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// The ID of the server associated with.
	// The ID of the server
	// +crossplane:generate:reference:type=github.com/scaleway/crossplane-provider-scaleway/apis/instance/v1alpha1.Server
	ServerID *string `json:"serverId,omitempty" tf:"server_id,omitempty"`

	// Reference to a Server in instance to populate serverId.
	// +kubebuilder:validation:Optional
	ServerIDRef *v1.Reference `json:"serverIdRef,omitempty" tf:"-"`

	// Selector for a Server in instance to populate serverId.
	// +kubebuilder:validation:Optional
	ServerIDSelector *v1.Selector `json:"serverIdSelector,omitempty" tf:"-"`

	// Value associated with your key
	// The value of the user data to set.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`

	// (Defaults to provider zone) The zone in which the server should be created.
	// The zone you want to attach the resource to
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type UserDataObservation struct {

	// The ID of the instance's user data.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Key of the user data.
	// The key of the user data to set.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// The ID of the server associated with.
	// The ID of the server
	ServerID *string `json:"serverId,omitempty" tf:"server_id,omitempty"`

	// Value associated with your key
	// The value of the user data to set.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`

	// (Defaults to provider zone) The zone in which the server should be created.
	// The zone you want to attach the resource to
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type UserDataParameters struct {

	// Key of the user data.
	// The key of the user data to set.
	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// The ID of the server associated with.
	// The ID of the server
	// +crossplane:generate:reference:type=github.com/scaleway/crossplane-provider-scaleway/apis/instance/v1alpha1.Server
	// +kubebuilder:validation:Optional
	ServerID *string `json:"serverId,omitempty" tf:"server_id,omitempty"`

	// Reference to a Server in instance to populate serverId.
	// +kubebuilder:validation:Optional
	ServerIDRef *v1.Reference `json:"serverIdRef,omitempty" tf:"-"`

	// Selector for a Server in instance to populate serverId.
	// +kubebuilder:validation:Optional
	ServerIDSelector *v1.Selector `json:"serverIdSelector,omitempty" tf:"-"`

	// Value associated with your key
	// The value of the user data to set.
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`

	// (Defaults to provider zone) The zone in which the server should be created.
	// The zone you want to attach the resource to
	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

// UserDataSpec defines the desired state of UserData
type UserDataSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UserDataParameters `json:"forProvider"`
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
	InitProvider UserDataInitParameters `json:"initProvider,omitempty"`
}

// UserDataStatus defines the observed state of UserData.
type UserDataStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UserDataObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// UserData is the Schema for the UserDatas API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,scaleway}
type UserData struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.key) || (has(self.initProvider) && has(self.initProvider.key))",message="spec.forProvider.key is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.value) || (has(self.initProvider) && has(self.initProvider.value))",message="spec.forProvider.value is a required parameter"
	Spec   UserDataSpec   `json:"spec"`
	Status UserDataStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UserDataList contains a list of UserDatas
type UserDataList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UserData `json:"items"`
}

// Repository type metadata.
var (
	UserData_Kind             = "UserData"
	UserData_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: UserData_Kind}.String()
	UserData_KindAPIVersion   = UserData_Kind + "." + CRDGroupVersion.String()
	UserData_GroupVersionKind = CRDGroupVersion.WithKind(UserData_Kind)
)

func init() {
	SchemeBuilder.Register(&UserData{}, &UserDataList{})
}
