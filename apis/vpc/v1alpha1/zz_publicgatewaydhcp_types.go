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

type PublicGatewayDHCPInitParameters struct {

	// The IP address of the public gateway DHCP config.
	// The address of the DHCP server. This will be the gateway's address in the private network. Defaults to the first address of the subnet
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// TLD given to hostnames in the Private Network. Allowed characters are a-z0-9-.. Defaults to the slugified Private Network name if created along a GatewayNetwork, or else to priv.
	// TLD given to hostnames in the Private Network. Allowed characters are `a-z0-9-.`. Defaults to the slugified Private Network name if created along a GatewayNetwork, or else to `priv`.
	DNSLocalName *string `json:"dnsLocalName,omitempty" tf:"dns_local_name,omitempty"`

	// Additional DNS search paths
	// Additional DNS search paths.
	DNSSearch []*string `json:"dnsSearch,omitempty" tf:"dns_search,omitempty"`

	// Override the DNS server list pushed to DHCP clients, instead of the gateway itself
	// Override the DNS server list pushed to DHCP clients, instead of the gateway itself.
	DNSServersOverride []*string `json:"dnsServersOverride,omitempty" tf:"dns_servers_override,omitempty"`

	// Whether to enable dynamic pooling of IPs. By turning the dynamic pool off, only pre-existing DHCP reservations will be handed out. Defaults to true.
	// Whether to enable dynamic pooling of IPs. By turning the dynamic pool off, only pre-existing DHCP reservations will be handed out. Defaults to true.
	EnableDynamic *bool `json:"enableDynamic,omitempty" tf:"enable_dynamic,omitempty"`

	// High IP (excluded) of the dynamic address pool. Defaults to the last address of the subnet.
	// High IP (included) of the dynamic address pool. Defaults to the last address of the subnet.
	PoolHigh *string `json:"poolHigh,omitempty" tf:"pool_high,omitempty"`

	// Low IP (included) of the dynamic address pool. Defaults to the second address of the subnet.
	// Low IP (included) of the dynamic address pool. Defaults to the second address of the subnet.
	PoolLow *string `json:"poolLow,omitempty" tf:"pool_low,omitempty"`

	// (Defaults to provider project_id) The ID of the project the public gateway DHCP config is associated with.
	// The project_id you want to attach the resource to
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Whether the gateway should push custom DNS servers to clients. This allows for instance hostname -> IP resolution. Defaults to true.
	// Whether the gateway should push custom DNS servers to clients. This allows for instance hostname -> IP resolution. Defaults to true.
	PushDNSServer *bool `json:"pushDnsServer,omitempty" tf:"push_dns_server,omitempty"`

	// Whether the gateway should push a default route to DHCP clients or only hand out IPs. Defaults to true.
	// Whether the gateway should push a default route to DHCP clients or only hand out IPs. Defaults to true.
	PushDefaultRoute *bool `json:"pushDefaultRoute,omitempty" tf:"push_default_route,omitempty"`

	// After how long, in seconds, a DHCP client will query for a new lease if previous renews fail. Must be 30s lower than valid_lifetime. Defaults to 51m (3060s).
	// After how long, in seconds, a DHCP client will query for a new lease if previous renews fail. Must be 30s lower than `valid_lifetime`. Defaults to 51m (3060s).
	RebindTimer *float64 `json:"rebindTimer,omitempty" tf:"rebind_timer,omitempty"`

	// After how long, in seconds, a renewal will be attempted. Must be 30s lower than rebind_timer. Defaults to 50m (3000s).
	// After how long, in seconds, a renew will be attempted. Must be 30s lower than `rebind_timer`. Defaults to 50m (3000s).
	RenewTimer *float64 `json:"renewTimer,omitempty" tf:"renew_timer,omitempty"`

	// The subnet to associate with the public gateway DHCP config.
	// Subnet for the DHCP server
	Subnet *string `json:"subnet,omitempty" tf:"subnet,omitempty"`

	// For how long, in seconds, will DHCP entries will be valid. Defaults to 1h (3600s).
	// For how long, in seconds, will DHCP entries will be valid. Defaults to 1h (3600s).
	ValidLifetime *float64 `json:"validLifetime,omitempty" tf:"valid_lifetime,omitempty"`

	// (Defaults to provider zone) The zone in which the public gateway DHCP config should be created.
	// The zone you want to attach the resource to
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type PublicGatewayDHCPObservation struct {

	// The IP address of the public gateway DHCP config.
	// The address of the DHCP server. This will be the gateway's address in the private network. Defaults to the first address of the subnet
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// The date and time of the creation of the public gateway DHCP config.
	// The date and time of the creation of the public gateway.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// TLD given to hostnames in the Private Network. Allowed characters are a-z0-9-.. Defaults to the slugified Private Network name if created along a GatewayNetwork, or else to priv.
	// TLD given to hostnames in the Private Network. Allowed characters are `a-z0-9-.`. Defaults to the slugified Private Network name if created along a GatewayNetwork, or else to `priv`.
	DNSLocalName *string `json:"dnsLocalName,omitempty" tf:"dns_local_name,omitempty"`

	// Additional DNS search paths
	// Additional DNS search paths.
	DNSSearch []*string `json:"dnsSearch,omitempty" tf:"dns_search,omitempty"`

	// Override the DNS server list pushed to DHCP clients, instead of the gateway itself
	// Override the DNS server list pushed to DHCP clients, instead of the gateway itself.
	DNSServersOverride []*string `json:"dnsServersOverride,omitempty" tf:"dns_servers_override,omitempty"`

	// Whether to enable dynamic pooling of IPs. By turning the dynamic pool off, only pre-existing DHCP reservations will be handed out. Defaults to true.
	// Whether to enable dynamic pooling of IPs. By turning the dynamic pool off, only pre-existing DHCP reservations will be handed out. Defaults to true.
	EnableDynamic *bool `json:"enableDynamic,omitempty" tf:"enable_dynamic,omitempty"`

	// The ID of the public gateway DHCP config.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The organization ID the public gateway DHCP config is associated with.
	// The organization_id you want to attach the resource to
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	// High IP (excluded) of the dynamic address pool. Defaults to the last address of the subnet.
	// High IP (included) of the dynamic address pool. Defaults to the last address of the subnet.
	PoolHigh *string `json:"poolHigh,omitempty" tf:"pool_high,omitempty"`

	// Low IP (included) of the dynamic address pool. Defaults to the second address of the subnet.
	// Low IP (included) of the dynamic address pool. Defaults to the second address of the subnet.
	PoolLow *string `json:"poolLow,omitempty" tf:"pool_low,omitempty"`

	// (Defaults to provider project_id) The ID of the project the public gateway DHCP config is associated with.
	// The project_id you want to attach the resource to
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Whether the gateway should push custom DNS servers to clients. This allows for instance hostname -> IP resolution. Defaults to true.
	// Whether the gateway should push custom DNS servers to clients. This allows for instance hostname -> IP resolution. Defaults to true.
	PushDNSServer *bool `json:"pushDnsServer,omitempty" tf:"push_dns_server,omitempty"`

	// Whether the gateway should push a default route to DHCP clients or only hand out IPs. Defaults to true.
	// Whether the gateway should push a default route to DHCP clients or only hand out IPs. Defaults to true.
	PushDefaultRoute *bool `json:"pushDefaultRoute,omitempty" tf:"push_default_route,omitempty"`

	// After how long, in seconds, a DHCP client will query for a new lease if previous renews fail. Must be 30s lower than valid_lifetime. Defaults to 51m (3060s).
	// After how long, in seconds, a DHCP client will query for a new lease if previous renews fail. Must be 30s lower than `valid_lifetime`. Defaults to 51m (3060s).
	RebindTimer *float64 `json:"rebindTimer,omitempty" tf:"rebind_timer,omitempty"`

	// After how long, in seconds, a renewal will be attempted. Must be 30s lower than rebind_timer. Defaults to 50m (3000s).
	// After how long, in seconds, a renew will be attempted. Must be 30s lower than `rebind_timer`. Defaults to 50m (3000s).
	RenewTimer *float64 `json:"renewTimer,omitempty" tf:"renew_timer,omitempty"`

	// The subnet to associate with the public gateway DHCP config.
	// Subnet for the DHCP server
	Subnet *string `json:"subnet,omitempty" tf:"subnet,omitempty"`

	// The date and time of the last update of the public gateway DHCP config.
	// The date and time of the last update of the public gateway.
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`

	// For how long, in seconds, will DHCP entries will be valid. Defaults to 1h (3600s).
	// For how long, in seconds, will DHCP entries will be valid. Defaults to 1h (3600s).
	ValidLifetime *float64 `json:"validLifetime,omitempty" tf:"valid_lifetime,omitempty"`

	// (Defaults to provider zone) The zone in which the public gateway DHCP config should be created.
	// The zone you want to attach the resource to
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type PublicGatewayDHCPParameters struct {

	// The IP address of the public gateway DHCP config.
	// The address of the DHCP server. This will be the gateway's address in the private network. Defaults to the first address of the subnet
	// +kubebuilder:validation:Optional
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// TLD given to hostnames in the Private Network. Allowed characters are a-z0-9-.. Defaults to the slugified Private Network name if created along a GatewayNetwork, or else to priv.
	// TLD given to hostnames in the Private Network. Allowed characters are `a-z0-9-.`. Defaults to the slugified Private Network name if created along a GatewayNetwork, or else to `priv`.
	// +kubebuilder:validation:Optional
	DNSLocalName *string `json:"dnsLocalName,omitempty" tf:"dns_local_name,omitempty"`

	// Additional DNS search paths
	// Additional DNS search paths.
	// +kubebuilder:validation:Optional
	DNSSearch []*string `json:"dnsSearch,omitempty" tf:"dns_search,omitempty"`

	// Override the DNS server list pushed to DHCP clients, instead of the gateway itself
	// Override the DNS server list pushed to DHCP clients, instead of the gateway itself.
	// +kubebuilder:validation:Optional
	DNSServersOverride []*string `json:"dnsServersOverride,omitempty" tf:"dns_servers_override,omitempty"`

	// Whether to enable dynamic pooling of IPs. By turning the dynamic pool off, only pre-existing DHCP reservations will be handed out. Defaults to true.
	// Whether to enable dynamic pooling of IPs. By turning the dynamic pool off, only pre-existing DHCP reservations will be handed out. Defaults to true.
	// +kubebuilder:validation:Optional
	EnableDynamic *bool `json:"enableDynamic,omitempty" tf:"enable_dynamic,omitempty"`

	// High IP (excluded) of the dynamic address pool. Defaults to the last address of the subnet.
	// High IP (included) of the dynamic address pool. Defaults to the last address of the subnet.
	// +kubebuilder:validation:Optional
	PoolHigh *string `json:"poolHigh,omitempty" tf:"pool_high,omitempty"`

	// Low IP (included) of the dynamic address pool. Defaults to the second address of the subnet.
	// Low IP (included) of the dynamic address pool. Defaults to the second address of the subnet.
	// +kubebuilder:validation:Optional
	PoolLow *string `json:"poolLow,omitempty" tf:"pool_low,omitempty"`

	// (Defaults to provider project_id) The ID of the project the public gateway DHCP config is associated with.
	// The project_id you want to attach the resource to
	// +kubebuilder:validation:Optional
	ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

	// Whether the gateway should push custom DNS servers to clients. This allows for instance hostname -> IP resolution. Defaults to true.
	// Whether the gateway should push custom DNS servers to clients. This allows for instance hostname -> IP resolution. Defaults to true.
	// +kubebuilder:validation:Optional
	PushDNSServer *bool `json:"pushDnsServer,omitempty" tf:"push_dns_server,omitempty"`

	// Whether the gateway should push a default route to DHCP clients or only hand out IPs. Defaults to true.
	// Whether the gateway should push a default route to DHCP clients or only hand out IPs. Defaults to true.
	// +kubebuilder:validation:Optional
	PushDefaultRoute *bool `json:"pushDefaultRoute,omitempty" tf:"push_default_route,omitempty"`

	// After how long, in seconds, a DHCP client will query for a new lease if previous renews fail. Must be 30s lower than valid_lifetime. Defaults to 51m (3060s).
	// After how long, in seconds, a DHCP client will query for a new lease if previous renews fail. Must be 30s lower than `valid_lifetime`. Defaults to 51m (3060s).
	// +kubebuilder:validation:Optional
	RebindTimer *float64 `json:"rebindTimer,omitempty" tf:"rebind_timer,omitempty"`

	// After how long, in seconds, a renewal will be attempted. Must be 30s lower than rebind_timer. Defaults to 50m (3000s).
	// After how long, in seconds, a renew will be attempted. Must be 30s lower than `rebind_timer`. Defaults to 50m (3000s).
	// +kubebuilder:validation:Optional
	RenewTimer *float64 `json:"renewTimer,omitempty" tf:"renew_timer,omitempty"`

	// The subnet to associate with the public gateway DHCP config.
	// Subnet for the DHCP server
	// +kubebuilder:validation:Optional
	Subnet *string `json:"subnet,omitempty" tf:"subnet,omitempty"`

	// For how long, in seconds, will DHCP entries will be valid. Defaults to 1h (3600s).
	// For how long, in seconds, will DHCP entries will be valid. Defaults to 1h (3600s).
	// +kubebuilder:validation:Optional
	ValidLifetime *float64 `json:"validLifetime,omitempty" tf:"valid_lifetime,omitempty"`

	// (Defaults to provider zone) The zone in which the public gateway DHCP config should be created.
	// The zone you want to attach the resource to
	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

// PublicGatewayDHCPSpec defines the desired state of PublicGatewayDHCP
type PublicGatewayDHCPSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PublicGatewayDHCPParameters `json:"forProvider"`
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
	InitProvider PublicGatewayDHCPInitParameters `json:"initProvider,omitempty"`
}

// PublicGatewayDHCPStatus defines the observed state of PublicGatewayDHCP.
type PublicGatewayDHCPStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PublicGatewayDHCPObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// PublicGatewayDHCP is the Schema for the PublicGatewayDHCPs API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,scaleway}
type PublicGatewayDHCP struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.subnet) || (has(self.initProvider) && has(self.initProvider.subnet))",message="spec.forProvider.subnet is a required parameter"
	Spec   PublicGatewayDHCPSpec   `json:"spec"`
	Status PublicGatewayDHCPStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PublicGatewayDHCPList contains a list of PublicGatewayDHCPs
type PublicGatewayDHCPList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PublicGatewayDHCP `json:"items"`
}

// Repository type metadata.
var (
	PublicGatewayDHCP_Kind             = "PublicGatewayDHCP"
	PublicGatewayDHCP_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PublicGatewayDHCP_Kind}.String()
	PublicGatewayDHCP_KindAPIVersion   = PublicGatewayDHCP_Kind + "." + CRDGroupVersion.String()
	PublicGatewayDHCP_GroupVersionKind = CRDGroupVersion.WithKind(PublicGatewayDHCP_Kind)
)

func init() {
	SchemeBuilder.Register(&PublicGatewayDHCP{}, &PublicGatewayDHCPList{})
}
