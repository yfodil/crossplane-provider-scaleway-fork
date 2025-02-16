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

type ObjectInitParameters struct {

	// The bucket's name or regional ID.
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// The content of the file to upload. Only one of file, content or content_base64 can be defined.
	// Content of the file to upload
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// The base64-encoded content of the file to upload. Only one of file, content or content_base64 can be defined.
	// Content of the file to upload, should be base64 encoded
	ContentBase64 *string `json:"contentBase64,omitempty" tf:"content_base64,omitempty"`

	// The name of the file to upload, defaults to an empty file.
	// Path of the file to upload, defaults to an empty file
	File *string `json:"file,omitempty" tf:"file,omitempty"`

	// Hash of the file, used to trigger the upload on file change.
	// File hash to trigger upload
	Hash *string `json:"hash,omitempty" tf:"hash,omitempty"`

	// The path to the object.
	// Key of the object
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Map of metadata used for the object (keys must be lowercase).
	// Map of object's metadata, only lower case keys are allowed
	// +mapType=granular
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// (Defaults to provider project_id) The ID of the project the bucket is associated with.
	// The project_id you want to attach the resource to
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// The Scaleway region the bucket resides in.
	// The region you want to attach the resource to
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Customer's encryption keys to encrypt data (SSE-C)
	// Customer's encryption keys to encrypt data (SSE-C)
	SseCustomerKeySecretRef *v1.SecretKeySelector `json:"sseCustomerKeySecretRef,omitempty" tf:"-"`

	// Specifies the Scaleway storage class (STANDARD, GLACIER, or ONEZONE_IA) used to store the object.
	// Specifies the Scaleway Object Storage class to which you want the object to transition
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`

	// Map of tags.
	// Map of object's tags
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Visibility of the object, public-read or private.
	// Visibility of the object, public-read or private
	Visibility *string `json:"visibility,omitempty" tf:"visibility,omitempty"`
}

type ObjectObservation struct {

	// The bucket's name or regional ID.
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// The content of the file to upload. Only one of file, content or content_base64 can be defined.
	// Content of the file to upload
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// The base64-encoded content of the file to upload. Only one of file, content or content_base64 can be defined.
	// Content of the file to upload, should be base64 encoded
	ContentBase64 *string `json:"contentBase64,omitempty" tf:"content_base64,omitempty"`

	// The name of the file to upload, defaults to an empty file.
	// Path of the file to upload, defaults to an empty file
	File *string `json:"file,omitempty" tf:"file,omitempty"`

	// Hash of the file, used to trigger the upload on file change.
	// File hash to trigger upload
	Hash *string `json:"hash,omitempty" tf:"hash,omitempty"`

	// The path of the object, including the name of the bucket.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The path to the object.
	// Key of the object
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Map of metadata used for the object (keys must be lowercase).
	// Map of object's metadata, only lower case keys are allowed
	// +mapType=granular
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// (Defaults to provider project_id) The ID of the project the bucket is associated with.
	// The project_id you want to attach the resource to
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// The Scaleway region the bucket resides in.
	// The region you want to attach the resource to
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the Scaleway storage class (STANDARD, GLACIER, or ONEZONE_IA) used to store the object.
	// Specifies the Scaleway Object Storage class to which you want the object to transition
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`

	// Map of tags.
	// Map of object's tags
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Visibility of the object, public-read or private.
	// Visibility of the object, public-read or private
	Visibility *string `json:"visibility,omitempty" tf:"visibility,omitempty"`
}

type ObjectParameters struct {

	// The bucket's name or regional ID.
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// The content of the file to upload. Only one of file, content or content_base64 can be defined.
	// Content of the file to upload
	// +kubebuilder:validation:Optional
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// The base64-encoded content of the file to upload. Only one of file, content or content_base64 can be defined.
	// Content of the file to upload, should be base64 encoded
	// +kubebuilder:validation:Optional
	ContentBase64 *string `json:"contentBase64,omitempty" tf:"content_base64,omitempty"`

	// The name of the file to upload, defaults to an empty file.
	// Path of the file to upload, defaults to an empty file
	// +kubebuilder:validation:Optional
	File *string `json:"file,omitempty" tf:"file,omitempty"`

	// Hash of the file, used to trigger the upload on file change.
	// File hash to trigger upload
	// +kubebuilder:validation:Optional
	Hash *string `json:"hash,omitempty" tf:"hash,omitempty"`

	// The path to the object.
	// Key of the object
	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Map of metadata used for the object (keys must be lowercase).
	// Map of object's metadata, only lower case keys are allowed
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// (Defaults to provider project_id) The ID of the project the bucket is associated with.
	// The project_id you want to attach the resource to
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// The Scaleway region the bucket resides in.
	// The region you want to attach the resource to
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Customer's encryption keys to encrypt data (SSE-C)
	// Customer's encryption keys to encrypt data (SSE-C)
	// +kubebuilder:validation:Optional
	SseCustomerKeySecretRef *v1.SecretKeySelector `json:"sseCustomerKeySecretRef,omitempty" tf:"-"`

	// Specifies the Scaleway storage class (STANDARD, GLACIER, or ONEZONE_IA) used to store the object.
	// Specifies the Scaleway Object Storage class to which you want the object to transition
	// +kubebuilder:validation:Optional
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`

	// Map of tags.
	// Map of object's tags
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Visibility of the object, public-read or private.
	// Visibility of the object, public-read or private
	// +kubebuilder:validation:Optional
	Visibility *string `json:"visibility,omitempty" tf:"visibility,omitempty"`
}

// ObjectSpec defines the desired state of Object
type ObjectSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ObjectParameters `json:"forProvider"`
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
	InitProvider ObjectInitParameters `json:"initProvider,omitempty"`
}

// ObjectStatus defines the observed state of Object.
type ObjectStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ObjectObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Object is the Schema for the Objects API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,scaleway}
type Object struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.bucket) || (has(self.initProvider) && has(self.initProvider.bucket))",message="spec.forProvider.bucket is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.key) || (has(self.initProvider) && has(self.initProvider.key))",message="spec.forProvider.key is a required parameter"
	Spec   ObjectSpec   `json:"spec"`
	Status ObjectStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ObjectList contains a list of Objects
type ObjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Object `json:"items"`
}

// Repository type metadata.
var (
	Object_Kind             = "Object"
	Object_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Object_Kind}.String()
	Object_KindAPIVersion   = Object_Kind + "." + CRDGroupVersion.String()
	Object_GroupVersionKind = CRDGroupVersion.WithKind(Object_Kind)
)

func init() {
	SchemeBuilder.Register(&Object{}, &ObjectList{})
}
