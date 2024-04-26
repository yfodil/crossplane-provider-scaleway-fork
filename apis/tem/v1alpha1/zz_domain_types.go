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

	// Acceptation of the Term of Service.
	// ~> Important:  This attribute must be set to true.
	// Accept the Scaleway Terms of Service
	AcceptTos *bool `json:"acceptTos,omitempty" tf:"accept_tos,omitempty"`

	// The domain name, must not be used in another Transactional Email Domain.
	// ~> Important: Updates to name will recreate the domain.
	// The domain name used when sending emails
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Defaults to provider project_id) The ID of the project the domain is associated with.
	// The project_id you want to attach the resource to
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// (Defaults to provider region). The region in which the domain should be created.
	// The region you want to attach the resource to
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type DomainObservation struct {

	// Acceptation of the Term of Service.
	// ~> Important:  This attribute must be set to true.
	// Accept the Scaleway Terms of Service
	AcceptTos *bool `json:"acceptTos,omitempty" tf:"accept_tos,omitempty"`

	// The date and time of the Transaction Email Domain's creation (RFC 3339 format).
	// Date and time of domain's creation (RFC 3339 format)
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// The DKIM public key, as should be recorded in the DNS zone.
	// DKIM public key, as should be recorded in the DNS zone
	DKIMConfig *string `json:"dkimConfig,omitempty" tf:"dkim_config,omitempty"`

	// The ID of the Transaction Email Domain.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The error message if the last check failed.
	// Error message if the last check failed
	LastError *string `json:"lastError,omitempty" tf:"last_error,omitempty"`

	// The date and time the domain was last found to be valid (RFC 3339 format).
	// Date and time the domain was last found to be valid (RFC 3339 format)
	LastValidAt *string `json:"lastValidAt,omitempty" tf:"last_valid_at,omitempty"`

	// The Scaleway's blackhole MX server to use if you do not have one.
	// The Scaleway's blackhole MX server to use
	MxBlackhole *string `json:"mxBlackhole,omitempty" tf:"mx_blackhole,omitempty"`

	// The domain name, must not be used in another Transactional Email Domain.
	// ~> Important: Updates to name will recreate the domain.
	// The domain name used when sending emails
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The date and time of the next scheduled check (RFC 3339 format).
	// Date and time of the next scheduled check (RFC 3339 format)
	NextCheckAt *string `json:"nextCheckAt,omitempty" tf:"next_check_at,omitempty"`

	// (Defaults to provider project_id) The ID of the project the domain is associated with.
	// The project_id you want to attach the resource to
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// (Defaults to provider region). The region in which the domain should be created.
	// The region you want to attach the resource to
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The domain's reputation.
	// The domain's reputation
	Reputation []ReputationObservation `json:"reputation,omitempty" tf:"reputation,omitempty"`

	// The date and time of the revocation of the domain (RFC 3339 format).
	// Date and time of the revocation of the domain (RFC 3339 format)
	RevokedAt *string `json:"revokedAt,omitempty" tf:"revoked_at,omitempty"`

	// The SMTP host to use to send emails.
	// SMTP host to use to send emails
	SMTPHost *string `json:"smtpHost,omitempty" tf:"smtp_host,omitempty"`

	// The SMTP port to use to send emails over TLS.
	// SMTP port to use to send emails over TLS. (Port 587)
	SMTPPort *float64 `json:"smtpPort,omitempty" tf:"smtp_port,omitempty"`

	// The SMTP port to use to send emails over TLS.
	// SMTP port to use to send emails over TLS. (Port 2587)
	SMTPPortAlternative *float64 `json:"smtpPortAlternative,omitempty" tf:"smtp_port_alternative,omitempty"`

	// The SMTP port to use to send emails.
	// SMTP port to use to send emails. (Port 25)
	SMTPPortUnsecure *float64 `json:"smtpPortUnsecure,omitempty" tf:"smtp_port_unsecure,omitempty"`

	// The SMTPS port to use to send emails over TLS Wrapper.
	// SMTPS port to use to send emails over TLS Wrapper. (Port 465)
	SmtpsPort *float64 `json:"smtpsPort,omitempty" tf:"smtps_port,omitempty"`

	// The SMTPS port to use to send emails over TLS Wrapper.
	// SMTPS port to use to send emails over TLS Wrapper. (Port 2465)
	SmtpsPortAlternative *float64 `json:"smtpsPortAlternative,omitempty" tf:"smtps_port_alternative,omitempty"`

	// The snippet of the SPF record that should be registered in the DNS zone.
	// Snippet of the SPF record that should be registered in the DNS zone
	SpfConfig *string `json:"spfConfig,omitempty" tf:"spf_config,omitempty"`

	// The status of the Transaction Email Domain.
	// Status of the domain
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type DomainParameters struct {

	// Acceptation of the Term of Service.
	// ~> Important:  This attribute must be set to true.
	// Accept the Scaleway Terms of Service
	// +kubebuilder:validation:Optional
	AcceptTos *bool `json:"acceptTos,omitempty" tf:"accept_tos,omitempty"`

	// The domain name, must not be used in another Transactional Email Domain.
	// ~> Important: Updates to name will recreate the domain.
	// The domain name used when sending emails
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Defaults to provider project_id) The ID of the project the domain is associated with.
	// The project_id you want to attach the resource to
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// (Defaults to provider region). The region in which the domain should be created.
	// The region you want to attach the resource to
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type ReputationInitParameters struct {
}

type ReputationObservation struct {

	// The previously-calculated domain's reputation score.
	PreviousScore *float64 `json:"previousScore,omitempty" tf:"previous_score,omitempty"`

	// The time and date the previous reputation score was calculated.
	PreviousScoredAt *string `json:"previousScoredAt,omitempty" tf:"previous_scored_at,omitempty"`

	// A range from 0 to 100 that determines your domain's reputation score.
	Score *float64 `json:"score,omitempty" tf:"score,omitempty"`

	// The time and date the score was calculated.
	ScoredAt *string `json:"scoredAt,omitempty" tf:"scored_at,omitempty"`

	// The status of the Transaction Email Domain.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type ReputationParameters struct {
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
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,scaleway}
type Domain struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.acceptTos) || (has(self.initProvider) && has(self.initProvider.acceptTos))",message="spec.forProvider.acceptTos is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
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
