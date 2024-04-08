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

type CronInitParameters struct {

	// The key-value mapping to define arguments that will be passed to your container’s event object
	// during
	// Cron arguments as json object to pass through during execution.
	Args *string `json:"args,omitempty" tf:"args,omitempty"`

	// The container ID to link with your cron.
	// The Container ID to link with your trigger.
	// +crossplane:generate:reference:type=Container
	ContainerID *string `json:"containerId,omitempty" tf:"container_id,omitempty"`

	// Reference to a Container to populate containerId.
	// +kubebuilder:validation:Optional
	ContainerIDRef *v1.Reference `json:"containerIdRef,omitempty" tf:"-"`

	// Selector for a Container to populate containerId.
	// +kubebuilder:validation:Optional
	ContainerIDSelector *v1.Selector `json:"containerIdSelector,omitempty" tf:"-"`

	// (Defaults to provider region) The region
	// in where the job was created.
	// The region you want to attach the resource to
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Cron format string, e.g. @hourly, as schedule time of its jobs to be created and
	// executed.
	// Cron format string, e.g. 0 * * * * or @hourly, as schedule time of its jobs to be created and executed.
	Schedule *string `json:"schedule,omitempty" tf:"schedule,omitempty"`
}

type CronObservation struct {

	// The key-value mapping to define arguments that will be passed to your container’s event object
	// during
	// Cron arguments as json object to pass through during execution.
	Args *string `json:"args,omitempty" tf:"args,omitempty"`

	// The container ID to link with your cron.
	// The Container ID to link with your trigger.
	ContainerID *string `json:"containerId,omitempty" tf:"container_id,omitempty"`

	// The container CRON's ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Defaults to provider region) The region
	// in where the job was created.
	// The region you want to attach the resource to
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Cron format string, e.g. @hourly, as schedule time of its jobs to be created and
	// executed.
	// Cron format string, e.g. 0 * * * * or @hourly, as schedule time of its jobs to be created and executed.
	Schedule *string `json:"schedule,omitempty" tf:"schedule,omitempty"`

	// The cron status.
	// Cron job status.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type CronParameters struct {

	// The key-value mapping to define arguments that will be passed to your container’s event object
	// during
	// Cron arguments as json object to pass through during execution.
	// +kubebuilder:validation:Optional
	Args *string `json:"args,omitempty" tf:"args,omitempty"`

	// The container ID to link with your cron.
	// The Container ID to link with your trigger.
	// +crossplane:generate:reference:type=Container
	// +kubebuilder:validation:Optional
	ContainerID *string `json:"containerId,omitempty" tf:"container_id,omitempty"`

	// Reference to a Container to populate containerId.
	// +kubebuilder:validation:Optional
	ContainerIDRef *v1.Reference `json:"containerIdRef,omitempty" tf:"-"`

	// Selector for a Container to populate containerId.
	// +kubebuilder:validation:Optional
	ContainerIDSelector *v1.Selector `json:"containerIdSelector,omitempty" tf:"-"`

	// (Defaults to provider region) The region
	// in where the job was created.
	// The region you want to attach the resource to
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Cron format string, e.g. @hourly, as schedule time of its jobs to be created and
	// executed.
	// Cron format string, e.g. 0 * * * * or @hourly, as schedule time of its jobs to be created and executed.
	// +kubebuilder:validation:Optional
	Schedule *string `json:"schedule,omitempty" tf:"schedule,omitempty"`
}

// CronSpec defines the desired state of Cron
type CronSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CronParameters `json:"forProvider"`
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
	InitProvider CronInitParameters `json:"initProvider,omitempty"`
}

// CronStatus defines the observed state of Cron.
type CronStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CronObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Cron is the Schema for the Crons API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,scaleway}
type Cron struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.args) || (has(self.initProvider) && has(self.initProvider.args))",message="spec.forProvider.args is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.schedule) || (has(self.initProvider) && has(self.initProvider.schedule))",message="spec.forProvider.schedule is a required parameter"
	Spec   CronSpec   `json:"spec"`
	Status CronStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CronList contains a list of Crons
type CronList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cron `json:"items"`
}

// Repository type metadata.
var (
	Cron_Kind             = "Cron"
	Cron_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Cron_Kind}.String()
	Cron_KindAPIVersion   = Cron_Kind + "." + CRDGroupVersion.String()
	Cron_GroupVersionKind = CRDGroupVersion.WithKind(Cron_Kind)
)

func init() {
	SchemeBuilder.Register(&Cron{}, &CronList{})
}
