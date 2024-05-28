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

type ImportInitParameters struct {

	// Bucket name containing qcow2 to import
	// Bucket containing qcow
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Key of the object to import
	// Key of the qcow file in the specified bucket
	Key *string `json:"key,omitempty" tf:"key,omitempty"`
}

type ImportObservation struct {

	// Bucket name containing qcow2 to import
	// Bucket containing qcow
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Key of the object to import
	// Key of the qcow file in the specified bucket
	Key *string `json:"key,omitempty" tf:"key,omitempty"`
}

type ImportParameters struct {

	// Bucket name containing qcow2 to import
	// Bucket containing qcow
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket" tf:"bucket,omitempty"`

	// Key of the object to import
	// Key of the qcow file in the specified bucket
	// +kubebuilder:validation:Optional
	Key *string `json:"key" tf:"key,omitempty"`
}

type SnapshotInitParameters struct {

	// Import a snapshot from a qcow2 file located in a bucket
	// Import snapshot from a qcow
	Import []ImportInitParameters `json:"import,omitempty" tf:"import,omitempty"`

	// The name of the snapshot. If not provided it will be randomly generated.
	// The name of the snapshot
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Defaults to provider project_id) The ID of the project the snapshot is
	// associated with.
	// The project_id you want to attach the resource to
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// A list of tags to apply to the snapshot.
	// The tags associated with the snapshot
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The snapshot's volume type.  The possible values are: b_ssd (Block SSD), l_ssd (Local SSD) and unified.
	// Updates to this field will recreate a new resource.
	// The snapshot's volume type
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The ID of the volume to take a snapshot from.
	// ID of the volume to take a snapshot from
	// +crossplane:generate:reference:type=Volume
	VolumeID *string `json:"volumeId,omitempty" tf:"volume_id,omitempty"`

	// Reference to a Volume to populate volumeId.
	// +kubebuilder:validation:Optional
	VolumeIDRef *v1.Reference `json:"volumeIdRef,omitempty" tf:"-"`

	// Selector for a Volume to populate volumeId.
	// +kubebuilder:validation:Optional
	VolumeIDSelector *v1.Selector `json:"volumeIdSelector,omitempty" tf:"-"`

	// (Defaults to provider zone) The zone in which
	// the snapshot should be created.
	// The zone you want to attach the resource to
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type SnapshotObservation struct {

	// The snapshot creation time.
	// The date and time of the creation of the snapshot
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// The ID of the snapshot.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Import a snapshot from a qcow2 file located in a bucket
	// Import snapshot from a qcow
	Import []ImportObservation `json:"import,omitempty" tf:"import,omitempty"`

	// The name of the snapshot. If not provided it will be randomly generated.
	// The name of the snapshot
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The organization ID the snapshot is associated with.
	// The organization_id you want to attach the resource to
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	// (Defaults to provider project_id) The ID of the project the snapshot is
	// associated with.
	// The project_id you want to attach the resource to
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// The size of the snapshot.
	// The size of the snapshot in gigabyte
	SizeInGb *float64 `json:"sizeInGb,omitempty" tf:"size_in_gb,omitempty"`

	// A list of tags to apply to the snapshot.
	// The tags associated with the snapshot
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The snapshot's volume type.  The possible values are: b_ssd (Block SSD), l_ssd (Local SSD) and unified.
	// Updates to this field will recreate a new resource.
	// The snapshot's volume type
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The ID of the volume to take a snapshot from.
	// ID of the volume to take a snapshot from
	VolumeID *string `json:"volumeId,omitempty" tf:"volume_id,omitempty"`

	// (Defaults to provider zone) The zone in which
	// the snapshot should be created.
	// The zone you want to attach the resource to
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type SnapshotParameters struct {

	// Import a snapshot from a qcow2 file located in a bucket
	// Import snapshot from a qcow
	// +kubebuilder:validation:Optional
	Import []ImportParameters `json:"import,omitempty" tf:"import,omitempty"`

	// The name of the snapshot. If not provided it will be randomly generated.
	// The name of the snapshot
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Defaults to provider project_id) The ID of the project the snapshot is
	// associated with.
	// The project_id you want to attach the resource to
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// A list of tags to apply to the snapshot.
	// The tags associated with the snapshot
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The snapshot's volume type.  The possible values are: b_ssd (Block SSD), l_ssd (Local SSD) and unified.
	// Updates to this field will recreate a new resource.
	// The snapshot's volume type
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The ID of the volume to take a snapshot from.
	// ID of the volume to take a snapshot from
	// +crossplane:generate:reference:type=Volume
	// +kubebuilder:validation:Optional
	VolumeID *string `json:"volumeId,omitempty" tf:"volume_id,omitempty"`

	// Reference to a Volume to populate volumeId.
	// +kubebuilder:validation:Optional
	VolumeIDRef *v1.Reference `json:"volumeIdRef,omitempty" tf:"-"`

	// Selector for a Volume to populate volumeId.
	// +kubebuilder:validation:Optional
	VolumeIDSelector *v1.Selector `json:"volumeIdSelector,omitempty" tf:"-"`

	// (Defaults to provider zone) The zone in which
	// the snapshot should be created.
	// The zone you want to attach the resource to
	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

// SnapshotSpec defines the desired state of Snapshot
type SnapshotSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SnapshotParameters `json:"forProvider"`
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
	InitProvider SnapshotInitParameters `json:"initProvider,omitempty"`
}

// SnapshotStatus defines the observed state of Snapshot.
type SnapshotStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SnapshotObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Snapshot is the Schema for the Snapshots API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,scaleway}
type Snapshot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SnapshotSpec   `json:"spec"`
	Status            SnapshotStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SnapshotList contains a list of Snapshots
type SnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Snapshot `json:"items"`
}

// Repository type metadata.
var (
	Snapshot_Kind             = "Snapshot"
	Snapshot_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Snapshot_Kind}.String()
	Snapshot_KindAPIVersion   = Snapshot_Kind + "." + CRDGroupVersion.String()
	Snapshot_GroupVersionKind = CRDGroupVersion.WithKind(Snapshot_Kind)
)

func init() {
	SchemeBuilder.Register(&Snapshot{}, &SnapshotList{})
}
