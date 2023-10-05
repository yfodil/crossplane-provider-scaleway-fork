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

type CertificateInitParameters struct {

	// Configuration block for custom certificate chain. Only one of letsencrypt and custom_certificate should be specified.
	// The custom type certificate type configuration
	CustomCertificate []CustomCertificateInitParameters `json:"customCertificate,omitempty" tf:"custom_certificate,omitempty"`

	// Configuration block for Let's Encrypt configuration. Only one of letsencrypt and custom_certificate should be specified.
	// The Let's Encrypt type certificate configuration
	Letsencrypt []LetsencryptInitParameters `json:"letsencrypt,omitempty" tf:"letsencrypt,omitempty"`

	// The name of the certificate backend.
	// The name of the load-balancer certificate
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type CertificateObservation struct {

	// Main domain of the certificate
	// The main domain name of the certificate
	CommonName *string `json:"commonName,omitempty" tf:"common_name,omitempty"`

	// Configuration block for custom certificate chain. Only one of letsencrypt and custom_certificate should be specified.
	// The custom type certificate type configuration
	CustomCertificate []CustomCertificateObservation `json:"customCertificate,omitempty" tf:"custom_certificate,omitempty"`

	// The identifier (SHA-1) of the certificate
	// The identifier (SHA-1) of the certificate
	Fingerprint *string `json:"fingerprint,omitempty" tf:"fingerprint,omitempty"`

	// The ID of the load-balancer certificate.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The load-balancer ID this certificate is attached to.
	// The load-balancer ID
	LBID *string `json:"lbId,omitempty" tf:"lb_id,omitempty"`

	// Configuration block for Let's Encrypt configuration. Only one of letsencrypt and custom_certificate should be specified.
	// The Let's Encrypt type certificate configuration
	Letsencrypt []LetsencryptObservation `json:"letsencrypt,omitempty" tf:"letsencrypt,omitempty"`

	// The name of the certificate backend.
	// The name of the load-balancer certificate
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The not valid after validity bound timestamp
	// The not valid after validity bound timestamp
	NotValidAfter *string `json:"notValidAfter,omitempty" tf:"not_valid_after,omitempty"`

	// The not valid before validity bound timestamp
	// The not valid before validity bound timestamp
	NotValidBefore *string `json:"notValidBefore,omitempty" tf:"not_valid_before,omitempty"`

	// Certificate status
	// The status of certificate
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// The alternative domain names of the certificate
	// The alternative domain names of the certificate
	SubjectAlternativeName []*string `json:"subjectAlternativeName,omitempty" tf:"subject_alternative_name,omitempty"`
}

type CertificateParameters struct {

	// Configuration block for custom certificate chain. Only one of letsencrypt and custom_certificate should be specified.
	// The custom type certificate type configuration
	// +kubebuilder:validation:Optional
	CustomCertificate []CustomCertificateParameters `json:"customCertificate,omitempty" tf:"custom_certificate,omitempty"`

	// The load-balancer ID this certificate is attached to.
	// The load-balancer ID
	// +crossplane:generate:reference:type=LB
	// +kubebuilder:validation:Optional
	LBID *string `json:"lbId,omitempty" tf:"lb_id,omitempty"`

	// Reference to a LB to populate lbId.
	// +kubebuilder:validation:Optional
	LBIDRef *v1.Reference `json:"lbIdRef,omitempty" tf:"-"`

	// Selector for a LB to populate lbId.
	// +kubebuilder:validation:Optional
	LBIDSelector *v1.Selector `json:"lbIdSelector,omitempty" tf:"-"`

	// Configuration block for Let's Encrypt configuration. Only one of letsencrypt and custom_certificate should be specified.
	// The Let's Encrypt type certificate configuration
	// +kubebuilder:validation:Optional
	Letsencrypt []LetsencryptParameters `json:"letsencrypt,omitempty" tf:"letsencrypt,omitempty"`

	// The name of the certificate backend.
	// The name of the load-balancer certificate
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type CustomCertificateInitParameters struct {

	// Full PEM-formatted certificate chain.
	// The full PEM-formatted certificate chain
	CertificateChain *string `json:"certificateChain,omitempty" tf:"certificate_chain,omitempty"`
}

type CustomCertificateObservation struct {

	// Full PEM-formatted certificate chain.
	// The full PEM-formatted certificate chain
	CertificateChain *string `json:"certificateChain,omitempty" tf:"certificate_chain,omitempty"`
}

type CustomCertificateParameters struct {

	// Full PEM-formatted certificate chain.
	// The full PEM-formatted certificate chain
	// +kubebuilder:validation:Optional
	CertificateChain *string `json:"certificateChain" tf:"certificate_chain,omitempty"`
}

type LetsencryptInitParameters struct {

	// Main domain of the certificate. A new certificate will be created if this field is changed.
	// The main domain name of the certificate
	CommonName *string `json:"commonName,omitempty" tf:"common_name,omitempty"`

	// Array of alternative domain names.  A new certificate will be created if this field is changed.
	// The alternative domain names of the certificate
	SubjectAlternativeName []*string `json:"subjectAlternativeName,omitempty" tf:"subject_alternative_name,omitempty"`
}

type LetsencryptObservation struct {

	// Main domain of the certificate. A new certificate will be created if this field is changed.
	// The main domain name of the certificate
	CommonName *string `json:"commonName,omitempty" tf:"common_name,omitempty"`

	// Array of alternative domain names.  A new certificate will be created if this field is changed.
	// The alternative domain names of the certificate
	SubjectAlternativeName []*string `json:"subjectAlternativeName,omitempty" tf:"subject_alternative_name,omitempty"`
}

type LetsencryptParameters struct {

	// Main domain of the certificate. A new certificate will be created if this field is changed.
	// The main domain name of the certificate
	// +kubebuilder:validation:Optional
	CommonName *string `json:"commonName" tf:"common_name,omitempty"`

	// Array of alternative domain names.  A new certificate will be created if this field is changed.
	// The alternative domain names of the certificate
	// +kubebuilder:validation:Optional
	SubjectAlternativeName []*string `json:"subjectAlternativeName,omitempty" tf:"subject_alternative_name,omitempty"`
}

// CertificateSpec defines the desired state of Certificate
type CertificateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CertificateParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider CertificateInitParameters `json:"initProvider,omitempty"`
}

// CertificateStatus defines the observed state of Certificate.
type CertificateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CertificateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Certificate is the Schema for the Certificates API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,scaleway}
type Certificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CertificateSpec   `json:"spec"`
	Status            CertificateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CertificateList contains a list of Certificates
type CertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Certificate `json:"items"`
}

// Repository type metadata.
var (
	Certificate_Kind             = "Certificate"
	Certificate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Certificate_Kind}.String()
	Certificate_KindAPIVersion   = Certificate_Kind + "." + CRDGroupVersion.String()
	Certificate_GroupVersionKind = CRDGroupVersion.WithKind(Certificate_Kind)
)

func init() {
	SchemeBuilder.Register(&Certificate{}, &CertificateList{})
}
