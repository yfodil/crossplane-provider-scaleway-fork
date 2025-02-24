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

type DatabaseInitParameters struct {

	// The database name (e.g. measurements).
	// The database name
	Dbname *string `json:"dbname,omitempty" tf:"dbname,omitempty"`

	// The database hostname. Can be an IP or a FQDN.
	// The database hostname
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// The database password.
	// The database password
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// The database port (e.g. 5432)
	// The database port
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The SQL query that will be executed when receiving a message ($TOPIC and $PAYLOAD variables are available, see documentation, e.g. INSERT INTO mytable(date, topic, value) VALUES (NOW(), $TOPIC, $PAYLOAD)).
	// SQL query to be executed ($TOPIC and $PAYLOAD variables are available, see documentation)
	Query *string `json:"query,omitempty" tf:"query,omitempty"`

	// The database username.
	// The database username
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type DatabaseObservation struct {

	// The database name (e.g. measurements).
	// The database name
	Dbname *string `json:"dbname,omitempty" tf:"dbname,omitempty"`

	// The database hostname. Can be an IP or a FQDN.
	// The database hostname
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// The database port (e.g. 5432)
	// The database port
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The SQL query that will be executed when receiving a message ($TOPIC and $PAYLOAD variables are available, see documentation, e.g. INSERT INTO mytable(date, topic, value) VALUES (NOW(), $TOPIC, $PAYLOAD)).
	// SQL query to be executed ($TOPIC and $PAYLOAD variables are available, see documentation)
	Query *string `json:"query,omitempty" tf:"query,omitempty"`

	// The database username.
	// The database username
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type DatabaseParameters struct {

	// The database name (e.g. measurements).
	// The database name
	// +kubebuilder:validation:Optional
	Dbname *string `json:"dbname" tf:"dbname,omitempty"`

	// The database hostname. Can be an IP or a FQDN.
	// The database hostname
	// +kubebuilder:validation:Optional
	Host *string `json:"host" tf:"host,omitempty"`

	// The database password.
	// The database password
	// +kubebuilder:validation:Optional
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// The database port (e.g. 5432)
	// The database port
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port" tf:"port,omitempty"`

	// The SQL query that will be executed when receiving a message ($TOPIC and $PAYLOAD variables are available, see documentation, e.g. INSERT INTO mytable(date, topic, value) VALUES (NOW(), $TOPIC, $PAYLOAD)).
	// SQL query to be executed ($TOPIC and $PAYLOAD variables are available, see documentation)
	// +kubebuilder:validation:Optional
	Query *string `json:"query" tf:"query,omitempty"`

	// The database username.
	// The database username
	// +kubebuilder:validation:Optional
	Username *string `json:"username" tf:"username,omitempty"`
}

type RestInitParameters struct {

	// a map of the extra headers to send with the HTTP call (e.g. X-Header = Value).
	// The HTTP call extra headers
	// +mapType=granular
	Headers map[string]*string `json:"headers,omitempty" tf:"headers,omitempty"`

	// The URI of the Rest endpoint (e.g. https://internal.mycompany.com/ingest/mqttdata).
	// The URI of the REST endpoint
	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`

	// The HTTP Verb used to call Rest URI (e.g. post).
	// The HTTP Verb used to call REST URI
	Verb *string `json:"verb,omitempty" tf:"verb,omitempty"`
}

type RestObservation struct {

	// a map of the extra headers to send with the HTTP call (e.g. X-Header = Value).
	// The HTTP call extra headers
	// +mapType=granular
	Headers map[string]*string `json:"headers,omitempty" tf:"headers,omitempty"`

	// The URI of the Rest endpoint (e.g. https://internal.mycompany.com/ingest/mqttdata).
	// The URI of the REST endpoint
	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`

	// The HTTP Verb used to call Rest URI (e.g. post).
	// The HTTP Verb used to call REST URI
	Verb *string `json:"verb,omitempty" tf:"verb,omitempty"`
}

type RestParameters struct {

	// a map of the extra headers to send with the HTTP call (e.g. X-Header = Value).
	// The HTTP call extra headers
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Headers map[string]*string `json:"headers" tf:"headers,omitempty"`

	// The URI of the Rest endpoint (e.g. https://internal.mycompany.com/ingest/mqttdata).
	// The URI of the REST endpoint
	// +kubebuilder:validation:Optional
	URI *string `json:"uri" tf:"uri,omitempty"`

	// The HTTP Verb used to call Rest URI (e.g. post).
	// The HTTP Verb used to call REST URI
	// +kubebuilder:validation:Optional
	Verb *string `json:"verb" tf:"verb,omitempty"`
}

type RouteInitParameters struct {

	// Configuration block for the database routes. See  product documentation for a better understanding of the parameters.
	// Database Route parameters
	Database []DatabaseInitParameters `json:"database,omitempty" tf:"database,omitempty"`

	// The hub ID to which the Route will be attached to.
	// The ID of the route's hub
	// +crossplane:generate:reference:type=github.com/scaleway/crossplane-provider-scaleway/apis/iot/v1alpha1.Hub
	HubID *string `json:"hubId,omitempty" tf:"hub_id,omitempty"`

	// Reference to a Hub in iot to populate hubId.
	// +kubebuilder:validation:Optional
	HubIDRef *v1.Reference `json:"hubIdRef,omitempty" tf:"-"`

	// Selector for a Hub in iot to populate hubId.
	// +kubebuilder:validation:Optional
	HubIDSelector *v1.Selector `json:"hubIdSelector,omitempty" tf:"-"`

	// The name of the IoT Route you want to create (e.g. my-route).
	// The name of the route
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Defaults to provider region) The region in which the Route is attached to.
	// The region you want to attach the resource to
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Configuration block for the rest routes. See product documentation for a better understanding of the parameters.
	// Rest Route parameters
	Rest []RestInitParameters `json:"rest,omitempty" tf:"rest,omitempty"`

	// Configuration block for the S3 routes. See product documentation for a better understanding of the parameters.
	// S3 Route parameters
	S3 []S3InitParameters `json:"s3,omitempty" tf:"s3,omitempty"`

	// The topic the Route subscribes to, wildcards allowed (e.g. thelab/+/temperature/#).
	// The Topic the route subscribes to (wildcards allowed)
	Topic *string `json:"topic,omitempty" tf:"topic,omitempty"`
}

type RouteObservation struct {

	// The date and time the Route was created.
	// The date and time of the creation of the IoT Route
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Configuration block for the database routes. See  product documentation for a better understanding of the parameters.
	// Database Route parameters
	Database []DatabaseObservation `json:"database,omitempty" tf:"database,omitempty"`

	// The hub ID to which the Route will be attached to.
	// The ID of the route's hub
	HubID *string `json:"hubId,omitempty" tf:"hub_id,omitempty"`

	// The ID of the Route.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the IoT Route you want to create (e.g. my-route).
	// The name of the route
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Defaults to provider region) The region in which the Route is attached to.
	// The region you want to attach the resource to
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Configuration block for the rest routes. See product documentation for a better understanding of the parameters.
	// Rest Route parameters
	Rest []RestObservation `json:"rest,omitempty" tf:"rest,omitempty"`

	// Configuration block for the S3 routes. See product documentation for a better understanding of the parameters.
	// S3 Route parameters
	S3 []S3Observation `json:"s3,omitempty" tf:"s3,omitempty"`

	// The topic the Route subscribes to, wildcards allowed (e.g. thelab/+/temperature/#).
	// The Topic the route subscribes to (wildcards allowed)
	Topic *string `json:"topic,omitempty" tf:"topic,omitempty"`
}

type RouteParameters struct {

	// Configuration block for the database routes. See  product documentation for a better understanding of the parameters.
	// Database Route parameters
	// +kubebuilder:validation:Optional
	Database []DatabaseParameters `json:"database,omitempty" tf:"database,omitempty"`

	// The hub ID to which the Route will be attached to.
	// The ID of the route's hub
	// +crossplane:generate:reference:type=github.com/scaleway/crossplane-provider-scaleway/apis/iot/v1alpha1.Hub
	// +kubebuilder:validation:Optional
	HubID *string `json:"hubId,omitempty" tf:"hub_id,omitempty"`

	// Reference to a Hub in iot to populate hubId.
	// +kubebuilder:validation:Optional
	HubIDRef *v1.Reference `json:"hubIdRef,omitempty" tf:"-"`

	// Selector for a Hub in iot to populate hubId.
	// +kubebuilder:validation:Optional
	HubIDSelector *v1.Selector `json:"hubIdSelector,omitempty" tf:"-"`

	// The name of the IoT Route you want to create (e.g. my-route).
	// The name of the route
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Defaults to provider region) The region in which the Route is attached to.
	// The region you want to attach the resource to
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Configuration block for the rest routes. See product documentation for a better understanding of the parameters.
	// Rest Route parameters
	// +kubebuilder:validation:Optional
	Rest []RestParameters `json:"rest,omitempty" tf:"rest,omitempty"`

	// Configuration block for the S3 routes. See product documentation for a better understanding of the parameters.
	// S3 Route parameters
	// +kubebuilder:validation:Optional
	S3 []S3Parameters `json:"s3,omitempty" tf:"s3,omitempty"`

	// The topic the Route subscribes to, wildcards allowed (e.g. thelab/+/temperature/#).
	// The Topic the route subscribes to (wildcards allowed)
	// +kubebuilder:validation:Optional
	Topic *string `json:"topic,omitempty" tf:"topic,omitempty"`
}

type S3InitParameters struct {

	// The name of the S3 route's destination bucket (e.g. my-object-storage).
	// The name of the S3 route's destination bucket
	BucketName *string `json:"bucketName,omitempty" tf:"bucket_name,omitempty"`

	// The region of the S3 route's destination bucket (e.g. fr-par).
	// The region of the S3 route's destination bucket
	BucketRegion *string `json:"bucketRegion,omitempty" tf:"bucket_region,omitempty"`

	// The string to prefix object names with (e.g. mykeyprefix-).
	// The string to prefix object names with
	ObjectPrefix *string `json:"objectPrefix,omitempty" tf:"object_prefix,omitempty"`

	// How the S3 route's objects will be created (e.g. per_topic). See documentation for behaviour details.
	// How the S3 route's objects will be created: one per topic or one per message
	Strategy *string `json:"strategy,omitempty" tf:"strategy,omitempty"`
}

type S3Observation struct {

	// The name of the S3 route's destination bucket (e.g. my-object-storage).
	// The name of the S3 route's destination bucket
	BucketName *string `json:"bucketName,omitempty" tf:"bucket_name,omitempty"`

	// The region of the S3 route's destination bucket (e.g. fr-par).
	// The region of the S3 route's destination bucket
	BucketRegion *string `json:"bucketRegion,omitempty" tf:"bucket_region,omitempty"`

	// The string to prefix object names with (e.g. mykeyprefix-).
	// The string to prefix object names with
	ObjectPrefix *string `json:"objectPrefix,omitempty" tf:"object_prefix,omitempty"`

	// How the S3 route's objects will be created (e.g. per_topic). See documentation for behaviour details.
	// How the S3 route's objects will be created: one per topic or one per message
	Strategy *string `json:"strategy,omitempty" tf:"strategy,omitempty"`
}

type S3Parameters struct {

	// The name of the S3 route's destination bucket (e.g. my-object-storage).
	// The name of the S3 route's destination bucket
	// +kubebuilder:validation:Optional
	BucketName *string `json:"bucketName" tf:"bucket_name,omitempty"`

	// The region of the S3 route's destination bucket (e.g. fr-par).
	// The region of the S3 route's destination bucket
	// +kubebuilder:validation:Optional
	BucketRegion *string `json:"bucketRegion" tf:"bucket_region,omitempty"`

	// The string to prefix object names with (e.g. mykeyprefix-).
	// The string to prefix object names with
	// +kubebuilder:validation:Optional
	ObjectPrefix *string `json:"objectPrefix,omitempty" tf:"object_prefix,omitempty"`

	// How the S3 route's objects will be created (e.g. per_topic). See documentation for behaviour details.
	// How the S3 route's objects will be created: one per topic or one per message
	// +kubebuilder:validation:Optional
	Strategy *string `json:"strategy" tf:"strategy,omitempty"`
}

// RouteSpec defines the desired state of Route
type RouteSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RouteParameters `json:"forProvider"`
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
	InitProvider RouteInitParameters `json:"initProvider,omitempty"`
}

// RouteStatus defines the observed state of Route.
type RouteStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RouteObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Route is the Schema for the Routes API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,scaleway}
type Route struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.topic) || (has(self.initProvider) && has(self.initProvider.topic))",message="spec.forProvider.topic is a required parameter"
	Spec   RouteSpec   `json:"spec"`
	Status RouteStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RouteList contains a list of Routes
type RouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Route `json:"items"`
}

// Repository type metadata.
var (
	Route_Kind             = "Route"
	Route_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Route_Kind}.String()
	Route_KindAPIVersion   = Route_Kind + "." + CRDGroupVersion.String()
	Route_GroupVersionKind = CRDGroupVersion.WithKind(Route_Kind)
)

func init() {
	SchemeBuilder.Register(&Route{}, &RouteList{})
}
