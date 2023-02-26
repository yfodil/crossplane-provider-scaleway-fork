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

type CredentialObservation struct {

	// The credential ID (UUID format).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The protocol associated to the Credential. Possible values are nats and sqs_sns.
	// The Namespace protocol
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Credential used to connect to the SQS/SNS service. Only one of nats_credentials
	// and sqs_sns_credentials may be set.
	// The credential used to connect to the SQS/SNS service
	// +kubebuilder:validation:Optional
	SqsSnsCredentials []SqsSnsCredentialsObservation `json:"sqsSnsCredentials,omitempty" tf:"sqs_sns_credentials,omitempty"`
}

type CredentialParameters struct {

	// The credential name..
	// The name of the Credential
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The namespace containing the Credential.
	// The ID of the Namespace associated to
	// +crossplane:generate:reference:type=MNQNamespace
	// +kubebuilder:validation:Optional
	NamespaceID *string `json:"namespaceId,omitempty" tf:"namespace_id,omitempty"`

	// Reference to a MNQNamespace to populate namespaceId.
	// +kubebuilder:validation:Optional
	NamespaceIDRef *v1.Reference `json:"namespaceIdRef,omitempty" tf:"-"`

	// Selector for a MNQNamespace to populate namespaceId.
	// +kubebuilder:validation:Optional
	NamespaceIDSelector *v1.Selector `json:"namespaceIdSelector,omitempty" tf:"-"`

	// Credentials file used to connect to the NATS service. Only one of nats_credentials and sqs_sns_credentials may be set.
	// credential for NATS protocol
	// +kubebuilder:validation:Optional
	NatsCredentials []NatsCredentialsParameters `json:"natsCredentials,omitempty" tf:"nats_credentials,omitempty"`

	// (Defaults to provider region). The region
	// in which the namespace should be created.
	// The region you want to attach the resource to
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Credential used to connect to the SQS/SNS service. Only one of nats_credentials
	// and sqs_sns_credentials may be set.
	// The credential used to connect to the SQS/SNS service
	// +kubebuilder:validation:Optional
	SqsSnsCredentials []SqsSnsCredentialsParameters `json:"sqsSnsCredentials,omitempty" tf:"sqs_sns_credentials,omitempty"`
}

type NatsCredentialsObservation struct {
}

type NatsCredentialsParameters struct {
}

type PermissionsObservation struct {
}

type PermissionsParameters struct {

	// . Defines if user can manage the associated resource(s).
	// Allow manage the associated resource
	// +kubebuilder:validation:Optional
	CanManage *bool `json:"canManage,omitempty" tf:"can_manage,omitempty"`

	// . Defines if user can publish messages to the service.
	// Allow publish messages to the service
	// +kubebuilder:validation:Optional
	CanPublish *bool `json:"canPublish,omitempty" tf:"can_publish,omitempty"`

	// . Defines if user can receive messages from the service.
	// Allow receive messages from the service
	// +kubebuilder:validation:Optional
	CanReceive *bool `json:"canReceive,omitempty" tf:"can_receive,omitempty"`
}

type SqsSnsCredentialsObservation struct {

	// The ID of the key.
	// The key of the credential
	AccessKey *string `json:"accessKey,omitempty" tf:"access_key,omitempty"`
}

type SqsSnsCredentialsParameters struct {

	// List of permissions associated to this Credential. Only one of permissions may be set.
	// The permission associated to this credential.
	// +kubebuilder:validation:Optional
	Permissions []PermissionsParameters `json:"permissions,omitempty" tf:"permissions,omitempty"`
}

// CredentialSpec defines the desired state of Credential
type CredentialSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CredentialParameters `json:"forProvider"`
}

// CredentialStatus defines the observed state of Credential.
type CredentialStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CredentialObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Credential is the Schema for the Credentials API. Manages Scaleway Messaging and Queuing Credential.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,scaleway}
type Credential struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CredentialSpec   `json:"spec"`
	Status            CredentialStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CredentialList contains a list of Credentials
type CredentialList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Credential `json:"items"`
}

// Repository type metadata.
var (
	Credential_Kind             = "Credential"
	Credential_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Credential_Kind}.String()
	Credential_KindAPIVersion   = Credential_Kind + "." + CRDGroupVersion.String()
	Credential_GroupVersionKind = CRDGroupVersion.WithKind(Credential_Kind)
)

func init() {
	SchemeBuilder.Register(&Credential{}, &CredentialList{})
}