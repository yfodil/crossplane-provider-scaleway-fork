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

type DefaultRetentionInitParameters struct {

	// The number of days you want to specify for the default retention period.
	// The number of days that you want to specify for the default retention period.
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// The default object lock retention mode you want to apply to new objects placed in the specified bucket. Valid values are GOVERNANCE or COMPLIANCE. Refer to the dedicated documentation for more information on retention modes.
	// The default Object Lock retention mode you want to apply to new objects placed in the specified bucket.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The number of years you want to specify for the default retention period.
	// The number of years that you want to specify for the default retention period.
	Years *float64 `json:"years,omitempty" tf:"years,omitempty"`
}

type DefaultRetentionObservation struct {

	// The number of days you want to specify for the default retention period.
	// The number of days that you want to specify for the default retention period.
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// The default object lock retention mode you want to apply to new objects placed in the specified bucket. Valid values are GOVERNANCE or COMPLIANCE. Refer to the dedicated documentation for more information on retention modes.
	// The default Object Lock retention mode you want to apply to new objects placed in the specified bucket.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The number of years you want to specify for the default retention period.
	// The number of years that you want to specify for the default retention period.
	Years *float64 `json:"years,omitempty" tf:"years,omitempty"`
}

type DefaultRetentionParameters struct {

	// The number of days you want to specify for the default retention period.
	// The number of days that you want to specify for the default retention period.
	// +kubebuilder:validation:Optional
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// The default object lock retention mode you want to apply to new objects placed in the specified bucket. Valid values are GOVERNANCE or COMPLIANCE. Refer to the dedicated documentation for more information on retention modes.
	// The default Object Lock retention mode you want to apply to new objects placed in the specified bucket.
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode" tf:"mode,omitempty"`

	// The number of years you want to specify for the default retention period.
	// The number of years that you want to specify for the default retention period.
	// +kubebuilder:validation:Optional
	Years *float64 `json:"years,omitempty" tf:"years,omitempty"`
}

type LockConfigurationInitParameters struct {

	// The bucket's name or regional ID.
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// (Defaults to provider project_id) The ID of the project the bucket is associated with.
	// The project_id you want to attach the resource to
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// The region you want to attach the resource to
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the object lock rule for the specified object.
	// Specifies the Object Lock rule for the specified object.
	Rule []RuleInitParameters `json:"rule,omitempty" tf:"rule,omitempty"`
}

type LockConfigurationObservation struct {

	// The bucket's name or regional ID.
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// The unique identifier of the Object bucket lock configuration.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Defaults to provider project_id) The ID of the project the bucket is associated with.
	// The project_id you want to attach the resource to
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// The region you want to attach the resource to
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the object lock rule for the specified object.
	// Specifies the Object Lock rule for the specified object.
	Rule []RuleObservation `json:"rule,omitempty" tf:"rule,omitempty"`
}

type LockConfigurationParameters struct {

	// The bucket's name or regional ID.
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// (Defaults to provider project_id) The ID of the project the bucket is associated with.
	// The project_id you want to attach the resource to
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// The region you want to attach the resource to
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the object lock rule for the specified object.
	// Specifies the Object Lock rule for the specified object.
	// +kubebuilder:validation:Optional
	Rule []RuleParameters `json:"rule,omitempty" tf:"rule,omitempty"`
}

type RuleInitParameters struct {

	// The default retention for the lock.
	DefaultRetention []DefaultRetentionInitParameters `json:"defaultRetention,omitempty" tf:"default_retention,omitempty"`
}

type RuleObservation struct {

	// The default retention for the lock.
	DefaultRetention []DefaultRetentionObservation `json:"defaultRetention,omitempty" tf:"default_retention,omitempty"`
}

type RuleParameters struct {

	// The default retention for the lock.
	// +kubebuilder:validation:Optional
	DefaultRetention []DefaultRetentionParameters `json:"defaultRetention" tf:"default_retention,omitempty"`
}

// LockConfigurationSpec defines the desired state of LockConfiguration
type LockConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LockConfigurationParameters `json:"forProvider"`
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
	InitProvider LockConfigurationInitParameters `json:"initProvider,omitempty"`
}

// LockConfigurationStatus defines the observed state of LockConfiguration.
type LockConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LockConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// LockConfiguration is the Schema for the LockConfigurations API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,scaleway}
type LockConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.bucket) || (has(self.initProvider) && has(self.initProvider.bucket))",message="spec.forProvider.bucket is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.rule) || (has(self.initProvider) && has(self.initProvider.rule))",message="spec.forProvider.rule is a required parameter"
	Spec   LockConfigurationSpec   `json:"spec"`
	Status LockConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LockConfigurationList contains a list of LockConfigurations
type LockConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LockConfiguration `json:"items"`
}

// Repository type metadata.
var (
	LockConfiguration_Kind             = "LockConfiguration"
	LockConfiguration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LockConfiguration_Kind}.String()
	LockConfiguration_KindAPIVersion   = LockConfiguration_Kind + "." + CRDGroupVersion.String()
	LockConfiguration_GroupVersionKind = CRDGroupVersion.WithKind(LockConfiguration_Kind)
)

func init() {
	SchemeBuilder.Register(&LockConfiguration{}, &LockConfigurationList{})
}
