// Package container provides access to the Google Container Engine API.
//
// See https://cloud.google.com/container-engine/
//
// Usage example:
//
//   import "github.com/skelterjohn/gcloud_apis/clients/container/v1"
//   ...
//   containerService, err := container.New(oauthHttpClient)
package container

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	context "golang.org/x/net/context"
	ctxhttp "golang.org/x/net/context/ctxhttp"
	gensupport "google.golang.org/api/gensupport"
	googleapi "google.golang.org/api/googleapi"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = gensupport.MarshalJSON
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Canceled
var _ = ctxhttp.Do

const apiId = "container:v1"
const apiName = "container"
const apiVersion = "v1"
const basePath = "https://container.googleapis.com/"

// OAuth2 scopes used by this API.
const (
	// View and manage your data across Google Cloud Platform services
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"

	// View your email address
	UserinfoEmailScope = "https://www.googleapis.com/auth/userinfo.email"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.MasterProjects = NewMasterProjectsService(s)
	s.Projects = NewProjectsService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	MasterProjects *MasterProjectsService

	Projects *ProjectsService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewMasterProjectsService(s *Service) *MasterProjectsService {
	rs := &MasterProjectsService{s: s}
	rs.Zones = NewMasterProjectsZonesService(s)
	return rs
}

type MasterProjectsService struct {
	s *Service

	Zones *MasterProjectsZonesService
}

func NewMasterProjectsZonesService(s *Service) *MasterProjectsZonesService {
	rs := &MasterProjectsZonesService{s: s}
	rs.SignedUrls = NewMasterProjectsZonesSignedUrlsService(s)
	rs.Tokens = NewMasterProjectsZonesTokensService(s)
	return rs
}

type MasterProjectsZonesService struct {
	s *Service

	SignedUrls *MasterProjectsZonesSignedUrlsService

	Tokens *MasterProjectsZonesTokensService
}

func NewMasterProjectsZonesSignedUrlsService(s *Service) *MasterProjectsZonesSignedUrlsService {
	rs := &MasterProjectsZonesSignedUrlsService{s: s}
	return rs
}

type MasterProjectsZonesSignedUrlsService struct {
	s *Service
}

func NewMasterProjectsZonesTokensService(s *Service) *MasterProjectsZonesTokensService {
	rs := &MasterProjectsZonesTokensService{s: s}
	return rs
}

type MasterProjectsZonesTokensService struct {
	s *Service
}

func NewProjectsService(s *Service) *ProjectsService {
	rs := &ProjectsService{s: s}
	rs.Zones = NewProjectsZonesService(s)
	return rs
}

type ProjectsService struct {
	s *Service

	Zones *ProjectsZonesService
}

func NewProjectsZonesService(s *Service) *ProjectsZonesService {
	rs := &ProjectsZonesService{s: s}
	rs.Clusters = NewProjectsZonesClustersService(s)
	rs.Operations = NewProjectsZonesOperationsService(s)
	return rs
}

type ProjectsZonesService struct {
	s *Service

	Clusters *ProjectsZonesClustersService

	Operations *ProjectsZonesOperationsService
}

func NewProjectsZonesClustersService(s *Service) *ProjectsZonesClustersService {
	rs := &ProjectsZonesClustersService{s: s}
	rs.NodePools = NewProjectsZonesClustersNodePoolsService(s)
	return rs
}

type ProjectsZonesClustersService struct {
	s *Service

	NodePools *ProjectsZonesClustersNodePoolsService
}

func NewProjectsZonesClustersNodePoolsService(s *Service) *ProjectsZonesClustersNodePoolsService {
	rs := &ProjectsZonesClustersNodePoolsService{s: s}
	return rs
}

type ProjectsZonesClustersNodePoolsService struct {
	s *Service
}

func NewProjectsZonesOperationsService(s *Service) *ProjectsZonesOperationsService {
	rs := &ProjectsZonesOperationsService{s: s}
	return rs
}

type ProjectsZonesOperationsService struct {
	s *Service
}

// AddonsConfig: Configuration for the addons that can be automatically
// spun up in the
// cluster, enabling additional functionality.
type AddonsConfig struct {
	// HorizontalPodAutoscaling: Configuration for the horizontal pod
	// autoscaling feature, which
	// increases or decreases the number of replica pods a replication
	// controller
	// has based on the resource usage of the existing pods.
	HorizontalPodAutoscaling *HorizontalPodAutoscaling `json:"horizontalPodAutoscaling,omitempty"`

	// HttpLoadBalancing: Configuration for the HTTP (L7) load balancing
	// controller addon, which
	// makes it easy to set up HTTP load balancers for services in a
	// cluster.
	HttpLoadBalancing *HttpLoadBalancing `json:"httpLoadBalancing,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "HorizontalPodAutoscaling") to unconditionally include in API
	// requests. By default, fields with empty values are omitted from API
	// requests. However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *AddonsConfig) MarshalJSON() ([]byte, error) {
	type noMethod AddonsConfig
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// AuthenticateRequest: A request to authenticate a user based on a
// provided OAuth2 token.
//
// This should look very close to the TokenReview struct
// in
// http://github.com/kubernetes/kubernetes/blob/master/pkg/apis/authen
// tication.k8s.io/v1beta1/types.go.
// This message has 4 GKE-specific fields that get mapped from the path,
// but the
// other fields (the expected JSON payload) must match TokenReview.
type AuthenticateRequest struct {
	// ApiVersion: The api version of the TokenReview object.
	ApiVersion string `json:"apiVersion,omitempty"`

	// Kind: Fields from
	// "pkg/apis/authentication.k8s.io/v1beta1".TokenReview:
	//
	// The "kind" of the TokenReview object.
	Kind string `json:"kind,omitempty"`

	// Metadata: "pkg/api/types".ObjectMeta
	// TODO (b/30563544): Remove these unused fields.
	// If cl/129007035 is rolled out to the prod envelope before b/28297888
	// is
	// fixed, we can change this to a google.protobuf.Any. Until then, only
	// the
	// fields that are acutally populated by the Token Webhook Authenticator
	Metadata *ObjectMeta `json:"metadata,omitempty"`

	// Spec: The information about the request being evaluated. It contains
	// the token
	// that the server should authenticate.
	Spec *TokenReviewSpec `json:"spec,omitempty"`

	// Status: The response for the provided request. (this won't be filled
	// in for an
	// AuthenticateRequest, but it is part of the struct, so we need it here
	// to be
	// safe).
	Status *TokenReviewStatus `json:"status,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ApiVersion") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *AuthenticateRequest) MarshalJSON() ([]byte, error) {
	type noMethod AuthenticateRequest
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// AuthenticateResponse: A response with the authenticated
// identity.
// This should match exactly with the TokenReview struct
// from
// http://github.com/kubernetes/kubernetes/blob/master/pkg/apis/auth
// entication.k8s.io/types.go.
type AuthenticateResponse struct {
	// ApiVersion: The api version of the TokenReview object.
	ApiVersion string `json:"apiVersion,omitempty"`

	// Kind: The "kind" of the TokenReview object.
	Kind string `json:"kind,omitempty"`

	// Spec: The information about the request that was evaluated.
	// This field (along with kind & api_version) are returned unchanged
	// from the
	// AuthenticateRequest. The caller probably doesn't care, but it would
	// allow
	// the caller to verify the question the server thought it was
	// answering.
	Spec *TokenReviewSpec `json:"spec,omitempty"`

	// Status: The response for the provided request.
	Status *TokenReviewStatus `json:"status,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "ApiVersion") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *AuthenticateResponse) MarshalJSON() ([]byte, error) {
	type noMethod AuthenticateResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// AuthorizeRequest: A request to authorize a user action. The request
// contains the attributes of
// the action the user is attempting. These attributes are mapped to a
// GKE IAM
// permission and policy to check.
//
// This should look very close to the SubjectAccessReview struct
// in
// http://github.com/kubernetes/kubernetes/blob/master/pkg/apis/author
// ization/v1beta1/types.go.
// This message has 4 GKE-specific fields that get mapped from the path,
// but the
// other fields (the expected JSON payload) must match
// SubjectAccessReview.
type AuthorizeRequest struct {
	// ApiVersion: The api version of the SubjectAccessReview object.
	ApiVersion string `json:"apiVersion,omitempty"`

	// Kind: Fields from
	// "pkg/apis/authorization/v1beta1".SubjectAccessReview:
	//
	// The "kind" of the SubjectAccessReview object.
	Kind string `json:"kind,omitempty"`

	// Metadata: Fields from "pkg/api/types".ObjectMeta
	// TODO (b/30563544): Remove these unused fields.
	// If cl/129007035 is rolled out to the prod envelope before b/28297888
	// is
	// fixed, we can change this to a google.protobuf.Any. Until then, only
	// the
	// fields that are acutally populated by the Token Webhook Authenticator
	Metadata *ObjectMeta `json:"metadata,omitempty"`

	// Spec: The information about the user action being evaluated.
	Spec *SubjectAccessReviewSpec `json:"spec,omitempty"`

	// Status: The response for the provided request (this won't be filled
	// in for an
	// AuthorizeRequest, but it is part of the struct, so we need it here to
	// be
	// safe).
	Status *SubjectAccessReviewStatus `json:"status,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ApiVersion") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *AuthorizeRequest) MarshalJSON() ([]byte, error) {
	type noMethod AuthorizeRequest
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// AuthorizeResponse: A response to a request for authorization.
// This should match exactly with the SubjectAccessReview struct
// from
// http://github.com/kubernetes/kubernetes/blob/master/pkg/apis/v1be
// ta1/authorization/types.go.
type AuthorizeResponse struct {
	// ApiVersion: The api version of the SubjectAccessReview object.
	ApiVersion string `json:"apiVersion,omitempty"`

	// Kind: The "kind" of the SubjectAccessReview object.
	Kind string `json:"kind,omitempty"`

	// Spec: The information about the request that was evaluated.
	// This field (along with kind & api_version) are returned unchanged
	// from the
	// AuthorizeRequest. The caller probably doesn't care, but it would
	// allow the
	// caller to verify the question the server thought it was answering.
	Spec *SubjectAccessReviewSpec `json:"spec,omitempty"`

	// Status: The response for the provided request.
	Status *SubjectAccessReviewStatus `json:"status,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "ApiVersion") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *AuthorizeResponse) MarshalJSON() ([]byte, error) {
	type noMethod AuthorizeResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// Cluster: A Google Container Engine cluster.
type Cluster struct {
	// AddonsConfig: Configurations for the various addons available to run
	// in the cluster.
	AddonsConfig *AddonsConfig `json:"addonsConfig,omitempty"`

	// ClusterIpv4Cidr: The IP address range of the container pods in this
	// cluster,
	// in
	// [CIDR](http://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing)
	//
	// notation (e.g. `10.96.0.0/14`). Leave blank to have
	// one automatically chosen or specify a `/14` block in `10.0.0.0/8`.
	ClusterIpv4Cidr string `json:"clusterIpv4Cidr,omitempty"`

	// CreateTime: [Output only] The time the cluster was created,
	// in
	// [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text format.
	CreateTime string `json:"createTime,omitempty"`

	// CurrentMasterVersion: [Output only] The current software version of
	// the master endpoint.
	CurrentMasterVersion string `json:"currentMasterVersion,omitempty"`

	// CurrentNodeCount: [Output only] The number of nodes currently in the
	// cluster.
	CurrentNodeCount int64 `json:"currentNodeCount,omitempty"`

	// CurrentNodeVersion: [Output only] The current version of the node
	// software components.
	// If they are currently at multiple versions because they're in the
	// process
	// of being upgraded, this reflects the minimum version of all nodes.
	CurrentNodeVersion string `json:"currentNodeVersion,omitempty"`

	// Description: An optional description of this cluster.
	Description string `json:"description,omitempty"`

	// EnableKubernetesAlpha: Kubernetes alpha features are enabled on this
	// cluster. This includes alpha
	// API groups (e.g. v1alpha1) and features that may not be production
	// ready in
	// the kubernetes version of the master and nodes.
	// The cluster has no SLA for uptime and master/node upgrades are
	// disabled.
	// Alpha enabled clusters are automatically deleted two weeks after
	// creation.
	EnableKubernetesAlpha bool `json:"enableKubernetesAlpha,omitempty"`

	// Endpoint: [Output only] The IP address of this cluster's master
	// endpoint.
	// The endpoint can be accessed from the internet
	// at
	// `https://username:password@endpoint/`.
	//
	// See the `masterAuth` property of this resource for username
	// and
	// password information.
	Endpoint string `json:"endpoint,omitempty"`

	// InitialClusterVersion: [Output only] The software version of the
	// master endpoint and kubelets used
	// in the cluster when it was first created. The version can be upgraded
	// over
	// time.
	InitialClusterVersion string `json:"initialClusterVersion,omitempty"`

	// InitialNodeCount: The number of nodes to create in this cluster. You
	// must ensure that your
	// Compute Engine <a href="/compute/docs/resource-quotas">resource
	// quota</a>
	// is sufficient for this number of instances. You must also have
	// available
	// firewall and routes quota.
	// For requests, this field should only be used in lieu of a
	// "node_pool" object, since this configuration (along with
	// the
	// "node_config") will be used to create a "NodePool" object with
	// an
	// auto-generated name. Do not use this and a node_pool at the same
	// time.
	InitialNodeCount int64 `json:"initialNodeCount,omitempty"`

	// InstanceGroupUrls: [Output only] The resource URLs of
	// [instance
	// groups](/compute/docs/instance-groups/) associated with this
	// cluster.
	InstanceGroupUrls []string `json:"instanceGroupUrls,omitempty"`

	// Locations: The list of Google Compute
	// Engine
	// [locations](/compute/docs/zones#available) in which the cluster's
	// nodes
	// should be located.
	Locations []string `json:"locations,omitempty"`

	// LoggingService: The logging service the cluster should use to write
	// logs.
	// Currently available options:
	//
	// * `logging.googleapis.com` - the Google Cloud Logging service.
	// * `none` - no logs will be exported from the cluster.
	// * if left as an empty string,`logging.googleapis.com` will be used.
	LoggingService string `json:"loggingService,omitempty"`

	// MasterAuth: The authentication information for accessing the master
	// endpoint.
	MasterAuth *MasterAuth `json:"masterAuth,omitempty"`

	// MonitoringService: The monitoring service the cluster should use to
	// write metrics.
	// Currently available options:
	//
	// * `monitoring.googleapis.com` - the Google Cloud Monitoring
	// service.
	// * `none` - no metrics will be exported from the cluster.
	// * if left as an empty string, `monitoring.googleapis.com` will be
	// used.
	MonitoringService string `json:"monitoringService,omitempty"`

	// Name: The name of this cluster. The name must be unique within this
	// project
	// and zone, and can be up to 40 characters with the following
	// restrictions:
	//
	// * Lowercase letters, numbers, and hyphens only.
	// * Must start with a letter.
	// * Must end with a number or a letter.
	Name string `json:"name,omitempty"`

	// Network: The name of the Google Compute
	// Engine
	// [network](/compute/docs/networks-and-firewalls#networks) to which
	// the
	// cluster is connected. If left unspecified, the `default` network
	// will be used.
	Network string `json:"network,omitempty"`

	// NodeConfig: Parameters used in creating the cluster's nodes.
	// See `nodeConfig` for the description of its properties.
	// For requests, this field should only be used in lieu of a
	// "node_pool" object, since this configuration (along with
	// the
	// "initial_node_count") will be used to create a "NodePool" object with
	// an
	// auto-generated name. Do not use this and a node_pool at the same
	// time.
	// For responses, this field will be populated with the node
	// configuration of
	// the first node pool.
	//
	// If unspecified, the defaults are used.
	NodeConfig *NodeConfig `json:"nodeConfig,omitempty"`

	// NodeIpv4CidrSize: [Output only] The size of the address space on each
	// node for hosting
	// containers. This is provisioned from within the
	// `container_ipv4_cidr`
	// range.
	NodeIpv4CidrSize int64 `json:"nodeIpv4CidrSize,omitempty"`

	// NodePools: The node pools associated with this cluster.
	// This field should not be set if "node_config" or "initial_node_count"
	// are
	// specified.
	NodePools []*NodePool `json:"nodePools,omitempty"`

	// SelfLink: [Output only] Server-defined URL for the resource.
	SelfLink string `json:"selfLink,omitempty"`

	// ServicesIpv4Cidr: [Output only] The IP address range of the
	// Kubernetes services in
	// this cluster,
	// in
	// [CIDR](http://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing)
	//
	// notation (e.g. `1.2.3.4/29`). Service addresses are
	// typically put in the last `/16` from the container CIDR.
	ServicesIpv4Cidr string `json:"servicesIpv4Cidr,omitempty"`

	// Status: [Output only] The current status of this cluster.
	//
	// Possible values:
	//   "STATUS_UNSPECIFIED" - Not set.
	//   "PROVISIONING" - The PROVISIONING state indicates the cluster is
	// being created.
	//   "RUNNING" - The RUNNING state indicates the cluster has been
	// created and is fully
	// usable.
	//   "RECONCILING" - The RECONCILING state indicates that some work is
	// actively being done on
	// the cluster, such as upgrading the master or node software. Details
	// can
	// be found in the `statusMessage` field.
	//   "STOPPING" - The STOPPING state indicates the cluster is being
	// deleted.
	//   "ERROR" - The ERROR state indicates the cluster may be unusable.
	// Details
	// can be found in the `statusMessage` field.
	Status string `json:"status,omitempty"`

	// StatusMessage: [Output only] Additional information about the current
	// status of this
	// cluster, if available.
	StatusMessage string `json:"statusMessage,omitempty"`

	// Subnetwork: The name of the Google Compute
	// Engine
	// [subnetwork](/compute/docs/subnetworks) to which the
	// cluster is connected.
	Subnetwork string `json:"subnetwork,omitempty"`

	// Zone: [Output only] The name of the Google Compute
	// Engine
	// [zone](/compute/docs/zones#available) in which the cluster
	// resides.
	Zone string `json:"zone,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "AddonsConfig") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Cluster) MarshalJSON() ([]byte, error) {
	type noMethod Cluster
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// ClusterUpdate: ClusterUpdate describes an update to the cluster.
// Exactly one update can
// be applied to a cluster with each request, so at most one field can
// be
// provided.
type ClusterUpdate struct {
	// DesiredAddonsConfig: Configurations for the various addons available
	// to run in the cluster.
	DesiredAddonsConfig *AddonsConfig `json:"desiredAddonsConfig,omitempty"`

	// DesiredImageType: The desired image type for the node pool.
	// NOTE: Set the "desired_node_pool" field as well.
	DesiredImageType string `json:"desiredImageType,omitempty"`

	// DesiredLocations: The desired list of Google Compute
	// Engine
	// [locations](/compute/docs/zones#available) in which the cluster's
	// nodes
	// should be located. Changing the locations a cluster is in will
	// result
	// in nodes being either created or removed from the cluster, depending
	// on
	// whether locations are being added or removed.
	//
	// This list must always include the cluster's primary zone.
	DesiredLocations []string `json:"desiredLocations,omitempty"`

	// DesiredMasterMachineType: The name of a Google Compute Engine
	// [machine
	// type](/compute/docs/machine-types)
	// (e.g. `n1-standard-8`) to change the master to.
	DesiredMasterMachineType string `json:"desiredMasterMachineType,omitempty"`

	// DesiredMasterVersion:
	// Whitelisted and internal users can change the master to any
	// version.
	//
	// The Kubernetes version to change the master to. The only valid value
	// is the
	// latest supported version. Use "-" to have the server automatically
	// select
	// the latest version.
	DesiredMasterVersion string `json:"desiredMasterVersion,omitempty"`

	// DesiredMonitoringService: The monitoring service the cluster should
	// use to write metrics.
	// Currently available options:
	//
	// * "monitoring.googleapis.com" - the Google Cloud Monitoring service
	// * "none" - no metrics will be exported from the cluster
	DesiredMonitoringService string `json:"desiredMonitoringService,omitempty"`

	// DesiredNodePoolAutoscaling: Autoscaler configuration for the node
	// pool specified in
	// desired_node_pool_id. If there is only one pool in the
	// cluster and desired_node_pool_id is not provided then
	// the change applies to that single node pool.
	DesiredNodePoolAutoscaling *NodePoolAutoscaling `json:"desiredNodePoolAutoscaling,omitempty"`

	// DesiredNodePoolId: The node pool to be upgraded. This field is
	// mandatory if
	// "desired_node_version", "desired_image_family"
	// or
	// "desired_node_pool_autoscaling" is specified and there is more than
	// one
	// node pool on the cluster.
	DesiredNodePoolId string `json:"desiredNodePoolId,omitempty"`

	// DesiredNodeVersion: The Kubernetes version to change the nodes to
	// (typically an
	// upgrade). Use `-` to upgrade to the latest version supported by
	// the server.
	DesiredNodeVersion string `json:"desiredNodeVersion,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DesiredAddonsConfig")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *ClusterUpdate) MarshalJSON() ([]byte, error) {
	type noMethod ClusterUpdate
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// CreateClusterRequest: CreateClusterRequest creates a cluster.
type CreateClusterRequest struct {
	// Cluster: A
	// [cluster
	// resource](/container-engine/reference/rest/v1/projects.zones.
	// clusters)
	Cluster *Cluster `json:"cluster,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Cluster") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *CreateClusterRequest) MarshalJSON() ([]byte, error) {
	type noMethod CreateClusterRequest
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// CreateNodePoolRequest: CreateNodePoolRequest creates a node pool for
// a cluster.
type CreateNodePoolRequest struct {
	// NodePool: The node pool to create.
	NodePool *NodePool `json:"nodePool,omitempty"`

	// ForceSendFields is a list of field names (e.g. "NodePool") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *CreateNodePoolRequest) MarshalJSON() ([]byte, error) {
	type noMethod CreateNodePoolRequest
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// CreateSignedUrlsRequest: A request for signed URLs that allow for
// writing a file to a private GCS
// bucket for storing backups of hosted master data.
type CreateSignedUrlsRequest struct {
	// ClusterId: The name of this master's cluster.
	ClusterId string `json:"clusterId,omitempty"`

	// Filenames: The names of the files for which a signed URLs are being
	// requested.
	Filenames []string `json:"filenames,omitempty"`

	// ProjectNumber: The project number for which the signed URLs are being
	// requested.  This is
	// the project in which this master's cluster resides.
	//
	// Note that this must be a project number, not a project ID.
	ProjectNumber int64 `json:"projectNumber,omitempty,string"`

	// ForceSendFields is a list of field names (e.g. "ClusterId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *CreateSignedUrlsRequest) MarshalJSON() ([]byte, error) {
	type noMethod CreateSignedUrlsRequest
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// CreateTokenRequest: A request for a compute-read-write
// (https://www.googleapis.com/auth/compute)
// scoped OAuth2 access token for <project_number>, to allow hosted
// masters to
// make modifications to a user's project.
type CreateTokenRequest struct {
	// ClusterId: The name of this master's cluster.
	ClusterId string `json:"clusterId,omitempty"`

	// ProjectNumber: The project number for which the access is being
	// requested.  This is the
	// project in which this master's cluster resides.
	//
	// Note that this must be a project number, not a project ID.
	ProjectNumber int64 `json:"projectNumber,omitempty,string"`

	// ForceSendFields is a list of field names (e.g. "ClusterId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *CreateTokenRequest) MarshalJSON() ([]byte, error) {
	type noMethod CreateTokenRequest
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// HorizontalPodAutoscaling: Configuration options for the horizontal
// pod autoscaling feature, which
// increases or decreases the number of replica pods a replication
// controller
// has based on the resource usage of the existing pods.
type HorizontalPodAutoscaling struct {
	// Disabled: Whether the Horizontal Pod Autoscaling feature is enabled
	// in the cluster.
	// When enabled, it ensures that a Heapster pod is running in the
	// cluster,
	// which is also used by the Cloud Monitoring service.
	Disabled bool `json:"disabled,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Disabled") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *HorizontalPodAutoscaling) MarshalJSON() ([]byte, error) {
	type noMethod HorizontalPodAutoscaling
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// HttpLoadBalancing: Configuration options for the HTTP (L7) load
// balancing controller addon,
// which makes it easy to set up HTTP load balancers for services in a
// cluster.
type HttpLoadBalancing struct {
	// Disabled: Whether the HTTP Load Balancing controller is enabled in
	// the cluster.
	// When enabled, it runs a small pod in the cluster that manages the
	// load
	// balancers.
	Disabled bool `json:"disabled,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Disabled") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *HttpLoadBalancing) MarshalJSON() ([]byte, error) {
	type noMethod HttpLoadBalancing
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// ImageReviewContainerSpec: ImageReviewContainerSpec is a description
// of a container within the creation request.
type ImageReviewContainerSpec struct {
	// Image: This can be in the form image:tag or
	// image@SHA:012345679abcdef.
	Image string `json:"image,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Image") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *ImageReviewContainerSpec) MarshalJSON() ([]byte, error) {
	type noMethod ImageReviewContainerSpec
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// ImageReviewRequest: A request to verify an image. The request
// contains the attributes of the
// container to create. These are passed to BCID for verification.
//
// This should look very close to the ImageReview struct
// in
// http://github.com/kubernetes/kubernetes/blob/master/pkg/apis/imagep
// olicy/v1beta1/types.go.
// This message has 4 GKE-specific fields that get mapped from the path,
// but the
// other fields (the expected JSON payload) must match ImageReview.
type ImageReviewRequest struct {
	// ApiVersion: The api version of the SubjectAccessReview object.
	ApiVersion string `json:"apiVersion,omitempty"`

	// Kind: Fields from
	// "pkg/apis/authorization/v1beta1".SubjectAccessReview:
	//
	// The "kind" of the SubjectAccessReview object.
	Kind string `json:"kind,omitempty"`

	// Metadata: Fields from "pkg/api/types".ObjectMeta
	// TODO (b/30563544): Remove these unused fields.
	// If cl/129007035 is rolled out to the prod envelope before b/28297888
	// is
	// fixed, we can change this to a google.protobuf.Any. Until then, only
	// the
	// fields that are acutally populated by the Token Webhook Authenticator
	Metadata *ObjectMeta `json:"metadata,omitempty"`

	// Spec: The information about the user action being evaluated.
	Spec *ImageReviewSpec `json:"spec,omitempty"`

	// Status: The response for the provided request (this won't be filled
	// in for an
	// AuthorizeRequest, but it is part of the struct, so we need it here to
	// be
	// safe).
	Status *ImageReviewStatus `json:"status,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ApiVersion") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *ImageReviewRequest) MarshalJSON() ([]byte, error) {
	type noMethod ImageReviewRequest
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// ImageReviewResponse: A response to a request for image
// verification.
// This should match exactly with the ImageReview struct
// from
// http://github.com/kubernetes/kubernetes/blob/master/pkg/apis/v1be
// ta1/authorization/types.go.
type ImageReviewResponse struct {
	// ApiVersion: The api version of the ImageReview object.
	ApiVersion string `json:"apiVersion,omitempty"`

	// Kind: The "kind" of the ImageReview object.
	Kind string `json:"kind,omitempty"`

	// Spec: The information about the request that was evaluated.
	// This field (along with kind & api_version) are returned unchanged
	// from the
	// ImageReviewRequest. The caller probably doesn't care, but it would
	// allow the
	// caller to verify the question the server thought it was answering.
	Spec *ImageReviewSpec `json:"spec,omitempty"`

	// Status: The response for the provided request.
	Status *ImageReviewStatus `json:"status,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "ApiVersion") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *ImageReviewResponse) MarshalJSON() ([]byte, error) {
	type noMethod ImageReviewResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// ImageReviewSpec: ImageReviewSpec is a description of the pod creation
// request.
type ImageReviewSpec struct {
	// Annotations: Annotations is a list of key-value pairs extracted from
	// the Pod's annotations.
	// It only includes keys which match the pattern
	// `*.image-policy.k8s.io/*`.
	// It is up to each webhook backend to determine how to interpret these
	// annotations, if at all.
	Annotations map[string]string `json:"annotations,omitempty"`

	// Containers: Container is a subset of the information in the container
	// being created.
	Containers []*ImageReviewContainerSpec `json:"containers,omitempty"`

	// Namespace: Namespace is the namespace the pod is being created in.
	Namespace string `json:"namespace,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Annotations") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *ImageReviewSpec) MarshalJSON() ([]byte, error) {
	type noMethod ImageReviewSpec
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// ImageReviewStatus: ImageReviewStatus is the result of the token
// authentication request.
type ImageReviewStatus struct {
	// Allowed: Allowed indicates that the image is allowed to run.
	Allowed bool `json:"allowed,omitempty"`

	// Reason: Reason should be empty unless Allowed is false in which case
	// it
	// may contain a short description of what is wrong.  Kubernetes
	// may truncate excessively long errors when displaying to the user.
	Reason string `json:"reason,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Allowed") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *ImageReviewStatus) MarshalJSON() ([]byte, error) {
	type noMethod ImageReviewStatus
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// ListClustersResponse: ListClustersResponse is the result of
// ListClustersRequest.
type ListClustersResponse struct {
	// Clusters: A list of clusters in the project in the specified zone,
	// or
	// across all ones.
	Clusters []*Cluster `json:"clusters,omitempty"`

	// MissingZones: If any zones are listed here, the list of clusters
	// returned
	// may be missing those zones.
	MissingZones []string `json:"missingZones,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Clusters") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *ListClustersResponse) MarshalJSON() ([]byte, error) {
	type noMethod ListClustersResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// ListNodePoolsResponse: ListNodePoolsResponse is the result of
// ListNodePoolsRequest.
type ListNodePoolsResponse struct {
	// NodePools: A list of node pools for a cluster.
	NodePools []*NodePool `json:"nodePools,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "NodePools") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *ListNodePoolsResponse) MarshalJSON() ([]byte, error) {
	type noMethod ListNodePoolsResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// ListOperationsResponse: ListOperationsResponse is the result of
// ListOperationsRequest.
type ListOperationsResponse struct {
	// MissingZones: If any zones are listed here, the list of operations
	// returned
	// may be missing the operations from those zones.
	MissingZones []string `json:"missingZones,omitempty"`

	// Operations: A list of operations in the project in the specified
	// zone.
	Operations []*Operation `json:"operations,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "MissingZones") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *ListOperationsResponse) MarshalJSON() ([]byte, error) {
	type noMethod ListOperationsResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// MasterAuth: The authentication information for accessing the master
// endpoint.
// Authentication can be done using HTTP basic auth or using
// client
// certificates.
type MasterAuth struct {
	// ClientCertificate: [Output only] Base64-encoded public certificate
	// used by clients to
	// authenticate to the cluster endpoint.
	ClientCertificate string `json:"clientCertificate,omitempty"`

	// ClientKey: [Output only] Base64-encoded private key used by clients
	// to authenticate
	// to the cluster endpoint.
	ClientKey string `json:"clientKey,omitempty"`

	// ClusterCaCertificate: [Output only] Base64-encoded public certificate
	// that is the root of
	// trust for the cluster.
	ClusterCaCertificate string `json:"clusterCaCertificate,omitempty"`

	// Password: The password to use for HTTP basic authentication to the
	// master endpoint.
	// Because the master endpoint is open to the Internet, you should
	// create a
	// strong password.
	Password string `json:"password,omitempty"`

	// Username: The username to use for HTTP basic authentication to the
	// master endpoint.
	Username string `json:"username,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ClientCertificate")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *MasterAuth) MarshalJSON() ([]byte, error) {
	type noMethod MasterAuth
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// NodeConfig: Parameters that describe the nodes in a cluster.
type NodeConfig struct {
	// DiskSizeGb: Size of the disk attached to each node, specified in
	// GB.
	// The smallest allowed disk size is 10GB.
	//
	// If unspecified, the default disk size is 100GB.
	DiskSizeGb int64 `json:"diskSizeGb,omitempty"`

	// ImageType: The image type to use for this node. Note that for a given
	// image type,
	// the latest version of it will be used.
	ImageType string `json:"imageType,omitempty"`

	// Labels: The map of Kubernetes labels (key/value pairs) to be applied
	// to each node.
	// These will added in addition to any default label(s) that
	// Kubernetes may apply to the node.
	// In case of conflict in label keys, the applied set may differ
	// depending on
	// the Kubernetes version -- it's best to assume the behavior is
	// undefined
	// and conflicts should be avoided.
	// For more information, including usage and the valid values,
	// see:
	// http://kubernetes.io/v1.1/docs/user-guide/labels.html
	Labels map[string]string `json:"labels,omitempty"`

	// LocalSsdCount: The number of local SSD disks to be attached to the
	// node.
	//
	// The limit for this value is dependant upon the maximum number
	// of
	// disks available on a machine per zone.
	// See:
	// https://cloud.google.com/compute/docs/disks/local-ssd#local_ssd_l
	// imits
	// for more information.
	LocalSsdCount int64 `json:"localSsdCount,omitempty"`

	// MachineType: The name of a Google Compute Engine
	// [machine
	// type](/compute/docs/machine-types) (e.g.
	// `n1-standard-1`).
	//
	// If unspecified, the default machine type is
	// `n1-standard-1`.
	MachineType string `json:"machineType,omitempty"`

	// Metadata: The metadata key/value pairs assigned to instances in the
	// cluster.
	//
	// Keys must conform to the regexp [a-zA-Z0-9-_]+ and be less than 128
	// bytes
	// in length. These are reflected as part of a URL in the metadata
	// server.
	// Additionally, to avoid ambiguity, keys must not conflict with any
	// other
	// metadata keys for the project or be one of the four reserved
	// keys:
	// "instance-template", "kube-env", "startup-script", and
	// "user-data"
	//
	// Values are free-form strings, and only have meaning as interpreted
	// by
	// the image running in the instance. The only restriction placed on
	// them is
	// that each value's size must be less than or equal to 32 KB.
	//
	// The total size of all keys and values must be less than 512 KB.
	Metadata map[string]string `json:"metadata,omitempty"`

	// OauthScopes: The set of Google API scopes to be made available on all
	// of the
	// node VMs under the "default" service account.
	//
	// The following scopes are recommended, but not required, and by
	// default are
	// not included:
	//
	// * `https://www.googleapis.com/auth/compute` is required for
	// mounting
	// persistent storage on your nodes.
	// * `https://www.googleapis.com/auth/devstorage.read_only` is required
	// for
	// communicating with **gcr.io**
	// (the [Google Container Registry](/container-registry/)).
	//
	// If unspecified, no scopes are added, unless Cloud Logging or
	// Cloud
	// Monitoring are enabled, in which case their required scopes will be
	// added.
	OauthScopes []string `json:"oauthScopes,omitempty"`

	// ServiceAccount: The Google Cloud Platform Service Account to be used
	// by the node VMs. If
	// no Service Account is specified, the "default" service account is
	// used.
	ServiceAccount string `json:"serviceAccount,omitempty"`

	// Tags: The list of instance tags applied to all nodes. Tags are used
	// to identify
	// valid sources or targets for network firewalls and are specified
	// by
	// the client during cluster or node pool creation. Each tag within the
	// list
	// must comply with RFC1035.
	Tags []string `json:"tags,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DiskSizeGb") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *NodeConfig) MarshalJSON() ([]byte, error) {
	type noMethod NodeConfig
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// NodePool: NodePool contains the name and configuration for a
// cluster's node pool.
// Node pools are a set of nodes (i.e. VM's), with a common
// configuration and
// specification, under the control of the cluster master. They may have
// a set
// of Kubernetes labels applied to them, which may be used to reference
// them
// during pod scheduling. They may also be resized up or down, to
// accommodate
// the workload.
type NodePool struct {
	// Autoscaling: Autoscaler configuration for this NodePool. Autoscaler
	// is enabled
	// only if a valid configuration is present.
	Autoscaling *NodePoolAutoscaling `json:"autoscaling,omitempty"`

	// Config: The node configuration of the pool.
	Config *NodeConfig `json:"config,omitempty"`

	// InitialNodeCount: The initial node count for the pool. You must
	// ensure that your
	// Compute Engine <a href="/compute/docs/resource-quotas">resource
	// quota</a>
	// is sufficient for this number of instances. You must also have
	// available
	// firewall and routes quota.
	InitialNodeCount int64 `json:"initialNodeCount,omitempty"`

	// InstanceGroupUrls: [Output only] The resource URLs of
	// [instance
	// groups](/compute/docs/instance-groups/) associated with this
	// node pool.
	InstanceGroupUrls []string `json:"instanceGroupUrls,omitempty"`

	// Name: The name of the node pool.
	Name string `json:"name,omitempty"`

	// SelfLink: [Output only] Server-defined URL for the resource.
	SelfLink string `json:"selfLink,omitempty"`

	// Status: [Output only] The status of the nodes in this pool instance.
	//
	// Possible values:
	//   "STATUS_UNSPECIFIED" - Not set.
	//   "PROVISIONING" - The PROVISIONING state indicates the node pool is
	// being created.
	//   "RUNNING" - The RUNNING state indicates the node pool has been
	// created
	// and is fully usable.
	//   "RUNNING_WITH_ERROR" - The RUNNING_WITH_ERROR state indicates the
	// node pool has been created
	// and is partially usable. Some error state has occurred and
	// some
	// functionality may be impaired. Customer may need to reissue a
	// request
	// or trigger a new update.
	//   "RECONCILING" - The RECONCILING state indicates that some work is
	// actively being done on
	// the node pool, such as upgrading node software. Details can
	// be found in the `statusMessage` field.
	//   "STOPPING" - The STOPPING state indicates the node pool is being
	// deleted.
	//   "ERROR" - The ERROR state indicates the node pool may be unusable.
	// Details
	// can be found in the `statusMessage` field.
	Status string `json:"status,omitempty"`

	// StatusMessage: [Output only] Additional information about the current
	// status of this
	// node pool instance, if available.
	StatusMessage string `json:"statusMessage,omitempty"`

	// Version: [Output only] The version of the Kubernetes of this node.
	Version string `json:"version,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Autoscaling") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *NodePool) MarshalJSON() ([]byte, error) {
	type noMethod NodePool
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// NodePoolAutoscaling: NodePoolAutoscaling contains information
// required by cluster autoscaler to
// adjust the size of the node pool to the current cluster usage.
type NodePoolAutoscaling struct {
	// Enabled: Is autoscaling enabled for this node pool.
	Enabled bool `json:"enabled,omitempty"`

	// MaxNodeCount: Maximum number of nodes in the NodePool. Must be >=
	// min_node_count. There
	// has to enough quota to scale up the cluster.
	MaxNodeCount int64 `json:"maxNodeCount,omitempty"`

	// MinNodeCount: Minimum number of nodes in the NodePool. Must be >= 1
	// and <=
	// max_node_count.
	MinNodeCount int64 `json:"minNodeCount,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Enabled") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *NodePoolAutoscaling) MarshalJSON() ([]byte, error) {
	type noMethod NodePoolAutoscaling
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// NonResourceAttributes: The authorization attributes of a non-resource
// request.
// This should match the NonResourceAttributes struct
// in
// pkg/apis/authorization/v1beta1/types.go.
type NonResourceAttributes struct {
	// Path: The URL path of the request.
	Path string `json:"path,omitempty"`

	// Verb: The verb of the request
	Verb string `json:"verb,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Path") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *NonResourceAttributes) MarshalJSON() ([]byte, error) {
	type noMethod NonResourceAttributes
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// ObjectMeta: The fields from "pkg/api/types".ObjectMeta that actually
// get filled in on
// Authenticate/Authorize requests.
type ObjectMeta struct {
	// CreationTimestamp: Timestamp representing the server time when this
	// object was created.
	CreationTimestamp string `json:"creationTimestamp,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CreationTimestamp")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *ObjectMeta) MarshalJSON() ([]byte, error) {
	type noMethod ObjectMeta
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// Operation: This operation resource represents operations that may
// have happened or are
// happening on the cluster. All fields are output only.
type Operation struct {
	// Detail: Detailed operation progress, if available.
	Detail string `json:"detail,omitempty"`

	// Name: The server-assigned ID for the operation.
	Name string `json:"name,omitempty"`

	// OperationType: The operation type.
	//
	// Possible values:
	//   "TYPE_UNSPECIFIED" - Not set.
	//   "CREATE_CLUSTER" - Cluster create.
	//   "DELETE_CLUSTER" - Cluster delete.
	//   "UPGRADE_MASTER" - A master upgrade.
	//   "UPGRADE_NODES" - A node upgrade.
	//   "REPAIR_CLUSTER" - Cluster repair.
	//   "UPDATE_CLUSTER" - Cluster update.
	//   "CREATE_NODE_POOL" - Node pool create.
	//   "DELETE_NODE_POOL" - Node pool delete.
	OperationType string `json:"operationType,omitempty"`

	// SelfLink: Server-defined URL for the resource.
	SelfLink string `json:"selfLink,omitempty"`

	// Status: The current status of the operation.
	//
	// Possible values:
	//   "STATUS_UNSPECIFIED" - Not set.
	//   "PENDING" - The operation has been created.
	//   "RUNNING" - The operation is currently running.
	//   "DONE" - The operation is done, either cancelled or completed.
	Status string `json:"status,omitempty"`

	// StatusMessage: If an error has occurred, a textual description of the
	// error.
	StatusMessage string `json:"statusMessage,omitempty"`

	// TargetLink: Server-defined URL for the target of the operation.
	TargetLink string `json:"targetLink,omitempty"`

	// Zone: The name of the Google Compute
	// Engine
	// [zone](/compute/docs/zones#available) in which the operation
	// is taking place.
	Zone string `json:"zone,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Detail") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Operation) MarshalJSON() ([]byte, error) {
	type noMethod Operation
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// ResourceAttributes: The authorization attributes of the resource
// request.
// This should match the ResourceAttributes struct
// in
// pkg/apis/authorization/v1beta1/types.go.
type ResourceAttributes struct {
	// Group: The API group of the resource.
	Group string `json:"group,omitempty"`

	// Name: The name of the resource in the request.
	Name string `json:"name,omitempty"`

	// Namespace: The namespace of the request.
	Namespace string `json:"namespace,omitempty"`

	// Resource: The type of the resource in the request.
	Resource string `json:"resource,omitempty"`

	// Subresource: The type of the subresource in the request.
	Subresource string `json:"subresource,omitempty"`

	// Verb: The Kubernetes verb of the request (e.g. get, create, list,
	// etc.)
	Verb string `json:"verb,omitempty"`

	// Version: The API version of the resource.
	Version string `json:"version,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Group") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *ResourceAttributes) MarshalJSON() ([]byte, error) {
	type noMethod ResourceAttributes
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// ServerConfig: Container Engine service configuration.
type ServerConfig struct {
	// BuildClientInfo: apiserver build BuildData::ClientInfo()
	BuildClientInfo string `json:"buildClientInfo,omitempty"`

	// DefaultClusterVersion: Version of Kubernetes the service deploys by
	// default.
	DefaultClusterVersion string `json:"defaultClusterVersion,omitempty"`

	// DefaultImageType: Default image type.
	DefaultImageType string `json:"defaultImageType,omitempty"`

	// ValidImageTypes: List of valid image types.
	ValidImageTypes []string `json:"validImageTypes,omitempty"`

	// ValidNodeVersions: List of valid node upgrade target versions.
	ValidNodeVersions []string `json:"validNodeVersions,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "BuildClientInfo") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *ServerConfig) MarshalJSON() ([]byte, error) {
	type noMethod ServerConfig
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// SignedUrls: Signed URLs that allow for writing a file to a private
// GCS bucket for
// storing backups of hosted master data.
type SignedUrls struct {
	// SignedUrls: The signed URLs for writing the request files, in the
	// same order as the
	// filenames in the request.
	SignedUrls []string `json:"signedUrls,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "SignedUrls") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *SignedUrls) MarshalJSON() ([]byte, error) {
	type noMethod SignedUrls
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// SubjectAccessReviewSpec: The description of the request for
// authorization.
// This should match the SubjectAccessReviewSpec struct
// in
// pkg/apis/authorization/v1beta1/types.go
type SubjectAccessReviewSpec struct {
	// Groups: Any groups this user may be a part of (this is not used for
	// GKE IAM).
	Groups []string `json:"groups,omitempty"`

	// NonResourceAttributes: The attributes of the request for a
	// non-resource request. If this field is
	// set, ResourceAttributes should not be set (and will be ignored).
	NonResourceAttributes *NonResourceAttributes `json:"nonResourceAttributes,omitempty"`

	// ResourceAttributes: The attributes of the request for a resource
	// request. If this field is set,
	// NonResourceAttributes should not be set (and will be ignored).
	ResourceAttributes *ResourceAttributes `json:"resourceAttributes,omitempty"`

	// User: The user making the request.
	User string `json:"user,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Groups") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *SubjectAccessReviewSpec) MarshalJSON() ([]byte, error) {
	type noMethod SubjectAccessReviewSpec
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// SubjectAccessReviewStatus: The result of the request for
// authorization.
// This should match the SubjectAccessReviewStatus struct
// in
// pkg/apis/authorization/v1beta1/types.go.
type SubjectAccessReviewStatus struct {
	// Allowed: Is the action authorized.
	Allowed bool `json:"allowed,omitempty"`

	// Reason: Error message if unauthorized.
	Reason string `json:"reason,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Allowed") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *SubjectAccessReviewStatus) MarshalJSON() ([]byte, error) {
	type noMethod SubjectAccessReviewStatus
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// Token: A compute-read-write (https://www.googleapis.com/auth/compute)
// scoped OAuth2
// access token, to allow hosted masters to make modifications to a
// user's
// project.
type Token struct {
	// AccessToken: The OAuth2 access token
	AccessToken string `json:"accessToken,omitempty"`

	// ExpireTime: The expiration time of the token.
	ExpireTime string `json:"expireTime,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "AccessToken") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Token) MarshalJSON() ([]byte, error) {
	type noMethod Token
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// TokenReviewSpec: The description of the request for review of an
// OAuth2 token for
// purposes of authentication.
// This should match the TokenReviewSpec struct
// in
// pkg/apis/authentication.k8s.io/v1beta1/types.go
type TokenReviewSpec struct {
	// Token: The token for the server to authenticate.
	Token string `json:"token,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Token") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *TokenReviewSpec) MarshalJSON() ([]byte, error) {
	type noMethod TokenReviewSpec
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// TokenReviewStatus: The result of the request for authentication.
// This should match the TokenReviewStatus struct
// in
// pkg/apis/authentication.k8s.io/v1beta1/types.go
type TokenReviewStatus struct {
	// Authenticated: Whether the server was able to authenticate the token.
	Authenticated bool `json:"authenticated,omitempty"`

	// User: The authenticated user associated with the token.
	User *UserInfo `json:"user,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Authenticated") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *TokenReviewStatus) MarshalJSON() ([]byte, error) {
	type noMethod TokenReviewStatus
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// UpdateClusterRequest: UpdateClusterRequest updates the settings of a
// cluster.
type UpdateClusterRequest struct {
	// Update: A description of the update.
	Update *ClusterUpdate `json:"update,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Update") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *UpdateClusterRequest) MarshalJSON() ([]byte, error) {
	type noMethod UpdateClusterRequest
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// UserInfo: The attributes of an authenticated user.
// This should match the UserInfo struct
// in
// pkg/apis/authentication.k8s.io/v1beta1/types.go
type UserInfo struct {
	// Groups: Groups that this user is a part of. This is not currently
	// filled in for
	// GKE.
	Groups []string `json:"groups,omitempty"`

	// Uid: A unique identifier (across time) for the user. This is not
	// currently
	// filled in for GKE.
	Uid string `json:"uid,omitempty"`

	// Username: The name of the user. This should be the email address
	// associated with the
	// GAIA identity of the user.
	Username string `json:"username,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Groups") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *UserInfo) MarshalJSON() ([]byte, error) {
	type noMethod UserInfo
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// method id "container.masterProjects.zones.authenticate":

type MasterProjectsZonesAuthenticateCall struct {
	s                   *Service
	masterProjectId     string
	zone                string
	projectNumber       int64
	clusterId           string
	authenticaterequest *AuthenticateRequest
	urlParams_          gensupport.URLParams
	ctx_                context.Context
}

// Authenticate: Processes a request to authenticate a token. If it is
// able to authenticate
// the token, the email for the authorized user is also
// returned.
// AuthenticateResponse also contains fields from the
// AuthenticateRequest. The
// server is expected to only fill in the AuthenticateResponse.Status.
// This is
// due to how the Authentication types are defined for the Kubernetes
// webhook
// authenticator:
// https://github.com/kubernetes/kubernetes/blob/m
// aster/pkg/apis/authentication.k8s.io/v1beta1/types.go.
func (r *MasterProjectsZonesService) Authenticate(masterProjectId string, zone string, projectNumber int64, clusterId string, authenticaterequest *AuthenticateRequest) *MasterProjectsZonesAuthenticateCall {
	c := &MasterProjectsZonesAuthenticateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.masterProjectId = masterProjectId
	c.zone = zone
	c.projectNumber = projectNumber
	c.clusterId = clusterId
	c.authenticaterequest = authenticaterequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MasterProjectsZonesAuthenticateCall) Fields(s ...googleapi.Field) *MasterProjectsZonesAuthenticateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *MasterProjectsZonesAuthenticateCall) Context(ctx context.Context) *MasterProjectsZonesAuthenticateCall {
	c.ctx_ = ctx
	return c
}

func (c *MasterProjectsZonesAuthenticateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.authenticaterequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/masterProjects/{masterProjectId}/zones/{zone}/{projectNumber}/{clusterId}/authenticate")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"masterProjectId": c.masterProjectId,
		"zone":            c.zone,
		"projectNumber":   strconv.FormatInt(c.projectNumber, 10),
		"clusterId":       c.clusterId,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "container.masterProjects.zones.authenticate" call.
// Exactly one of *AuthenticateResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *AuthenticateResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *MasterProjectsZonesAuthenticateCall) Do(opts ...googleapi.CallOption) (*AuthenticateResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &AuthenticateResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Processes a request to authenticate a token. If it is able to authenticate\nthe token, the email for the authorized user is also returned.\nAuthenticateResponse also contains fields from the AuthenticateRequest. The\nserver is expected to only fill in the AuthenticateResponse.Status. This is\ndue to how the Authentication types are defined for the Kubernetes webhook\nauthenticator:\nhttps://github.com/kubernetes/kubernetes/blob/master/pkg/apis/authentication.k8s.io/v1beta1/types.go.",
	//   "flatPath": "v1/masterProjects/{masterProjectId}/zones/{zone}/{projectNumber}/{clusterId}/authenticate",
	//   "httpMethod": "POST",
	//   "id": "container.masterProjects.zones.authenticate",
	//   "parameterOrder": [
	//     "masterProjectId",
	//     "zone",
	//     "projectNumber",
	//     "clusterId"
	//   ],
	//   "parameters": {
	//     "clusterId": {
	//       "description": "The name of this master's cluster.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "masterProjectId": {
	//       "description": "The hosted master project in which this master resides.  This can be\neither a [project ID or project\nnumber](https://support.google.com/cloud/answer/6158840).",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectNumber": {
	//       "description": "The project number for which the signed URLs are being requested.  This is\nthe project in which this master's cluster resides.\n\nNote that this must be a project number, not a project ID.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The zone of this master's cluster.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/masterProjects/{masterProjectId}/zones/{zone}/{projectNumber}/{clusterId}/authenticate",
	//   "request": {
	//     "$ref": "AuthenticateRequest"
	//   },
	//   "response": {
	//     "$ref": "AuthenticateResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/userinfo.email"
	//   ]
	// }

}

// method id "container.masterProjects.zones.authorize":

type MasterProjectsZonesAuthorizeCall struct {
	s                *Service
	masterProjectId  string
	zone             string
	projectNumber    int64
	clusterId        string
	authorizerequest *AuthorizeRequest
	urlParams_       gensupport.URLParams
	ctx_             context.Context
}

// Authorize: Processes the attributes of a user request and determines
// whether or not
// to authorize the request. If unauthorized, a reason is also provided.
// The
// AuthorizeResponse also contains fields from the AuthorizeRequest.
// The
// server is expected to only fill in the AuthorizeResponse.Status. This
// is
// due to how the Authorization types are defined for the Kubernetes
// webhook
// authorizer:
// https://github.com/kubernetes/kubernetes/blob/mast
// er/pkg/apis/authorization/v1beta1/types.go.
func (r *MasterProjectsZonesService) Authorize(masterProjectId string, zone string, projectNumber int64, clusterId string, authorizerequest *AuthorizeRequest) *MasterProjectsZonesAuthorizeCall {
	c := &MasterProjectsZonesAuthorizeCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.masterProjectId = masterProjectId
	c.zone = zone
	c.projectNumber = projectNumber
	c.clusterId = clusterId
	c.authorizerequest = authorizerequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MasterProjectsZonesAuthorizeCall) Fields(s ...googleapi.Field) *MasterProjectsZonesAuthorizeCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *MasterProjectsZonesAuthorizeCall) Context(ctx context.Context) *MasterProjectsZonesAuthorizeCall {
	c.ctx_ = ctx
	return c
}

func (c *MasterProjectsZonesAuthorizeCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.authorizerequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/masterProjects/{masterProjectId}/zones/{zone}/{projectNumber}/{clusterId}/authorize")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"masterProjectId": c.masterProjectId,
		"zone":            c.zone,
		"projectNumber":   strconv.FormatInt(c.projectNumber, 10),
		"clusterId":       c.clusterId,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "container.masterProjects.zones.authorize" call.
// Exactly one of *AuthorizeResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *AuthorizeResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *MasterProjectsZonesAuthorizeCall) Do(opts ...googleapi.CallOption) (*AuthorizeResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &AuthorizeResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Processes the attributes of a user request and determines whether or not\nto authorize the request. If unauthorized, a reason is also provided. The\nAuthorizeResponse also contains fields from the AuthorizeRequest. The\nserver is expected to only fill in the AuthorizeResponse.Status. This is\ndue to how the Authorization types are defined for the Kubernetes webhook\nauthorizer:\nhttps://github.com/kubernetes/kubernetes/blob/master/pkg/apis/authorization/v1beta1/types.go.",
	//   "flatPath": "v1/masterProjects/{masterProjectId}/zones/{zone}/{projectNumber}/{clusterId}/authorize",
	//   "httpMethod": "POST",
	//   "id": "container.masterProjects.zones.authorize",
	//   "parameterOrder": [
	//     "masterProjectId",
	//     "zone",
	//     "projectNumber",
	//     "clusterId"
	//   ],
	//   "parameters": {
	//     "clusterId": {
	//       "description": "The name of this master's cluster.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "masterProjectId": {
	//       "description": "The hosted master project in which this master resides.  This can be\neither a [project ID or project\nnumber](https://support.google.com/cloud/answer/6158840).",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectNumber": {
	//       "description": "The project number for which the request is being authorized.  This is\nthe project in which this master's cluster resides.\n\nThis is an int64, so it must be a project number, not a project ID.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The zone of this master's cluster.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/masterProjects/{masterProjectId}/zones/{zone}/{projectNumber}/{clusterId}/authorize",
	//   "request": {
	//     "$ref": "AuthorizeRequest"
	//   },
	//   "response": {
	//     "$ref": "AuthorizeResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/userinfo.email"
	//   ]
	// }

}

// method id "container.masterProjects.zones.imagereview":

type MasterProjectsZonesImagereviewCall struct {
	s                  *Service
	masterProjectId    string
	zone               string
	projectNumber      int64
	clusterId          string
	imagereviewrequest *ImageReviewRequest
	urlParams_         gensupport.URLParams
	ctx_               context.Context
}

// Imagereview: Processes a request to verify the container image. If
// unverified, a reason
// is also provided. The ImageReviewResponse also contains fields from
// the
// ImageReviewRequest. The server is expected to only fill in
// the
// ImageReviewResponse.Status. This is due to how the ImageReview types
// are
// defined for the Kubernetes webhook image
// review:
// https://github.com/kubernetes/kubernetes/blob/master/pkg/apis/
// imagepolicy/v1beta1/types.go.
func (r *MasterProjectsZonesService) Imagereview(masterProjectId string, zone string, projectNumber int64, clusterId string, imagereviewrequest *ImageReviewRequest) *MasterProjectsZonesImagereviewCall {
	c := &MasterProjectsZonesImagereviewCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.masterProjectId = masterProjectId
	c.zone = zone
	c.projectNumber = projectNumber
	c.clusterId = clusterId
	c.imagereviewrequest = imagereviewrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MasterProjectsZonesImagereviewCall) Fields(s ...googleapi.Field) *MasterProjectsZonesImagereviewCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *MasterProjectsZonesImagereviewCall) Context(ctx context.Context) *MasterProjectsZonesImagereviewCall {
	c.ctx_ = ctx
	return c
}

func (c *MasterProjectsZonesImagereviewCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.imagereviewrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/masterProjects/{masterProjectId}/zones/{zone}/{projectNumber}/{clusterId}/imagereview")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"masterProjectId": c.masterProjectId,
		"zone":            c.zone,
		"projectNumber":   strconv.FormatInt(c.projectNumber, 10),
		"clusterId":       c.clusterId,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "container.masterProjects.zones.imagereview" call.
// Exactly one of *ImageReviewResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ImageReviewResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *MasterProjectsZonesImagereviewCall) Do(opts ...googleapi.CallOption) (*ImageReviewResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ImageReviewResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Processes a request to verify the container image. If unverified, a reason\nis also provided. The ImageReviewResponse also contains fields from the\nImageReviewRequest. The server is expected to only fill in the\nImageReviewResponse.Status. This is due to how the ImageReview types are\ndefined for the Kubernetes webhook image review:\nhttps://github.com/kubernetes/kubernetes/blob/master/pkg/apis/imagepolicy/v1beta1/types.go.",
	//   "flatPath": "v1/masterProjects/{masterProjectId}/zones/{zone}/{projectNumber}/{clusterId}/imagereview",
	//   "httpMethod": "POST",
	//   "id": "container.masterProjects.zones.imagereview",
	//   "parameterOrder": [
	//     "masterProjectId",
	//     "zone",
	//     "projectNumber",
	//     "clusterId"
	//   ],
	//   "parameters": {
	//     "clusterId": {
	//       "description": "The name of this master's cluster.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "masterProjectId": {
	//       "description": "The hosted master project in which this master resides.  This can be\neither a [project ID or project\nnumber](https://support.google.com/cloud/answer/6158840).",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectNumber": {
	//       "description": "The project number for which the request is being authorized.  This is\nthe project in which this master's cluster resides.\n\nThis is an int64, so it must be a project number, not a project ID.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The zone of this master's cluster.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/masterProjects/{masterProjectId}/zones/{zone}/{projectNumber}/{clusterId}/imagereview",
	//   "request": {
	//     "$ref": "ImageReviewRequest"
	//   },
	//   "response": {
	//     "$ref": "ImageReviewResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/userinfo.email"
	//   ]
	// }

}

// method id "container.masterProjects.zones.signedUrls.create":

type MasterProjectsZonesSignedUrlsCreateCall struct {
	s                       *Service
	masterProjectId         string
	zone                    string
	createsignedurlsrequest *CreateSignedUrlsRequest
	urlParams_              gensupport.URLParams
	ctx_                    context.Context
}

// Create: Creates signed URLs that allow for writing a file to a
// private GCS bucket
// for storing backups of hosted master data. Signed URLs are explained
// here:
// https://cloud.google.com/storage/docs/access-control#Signed-URLs
func (r *MasterProjectsZonesSignedUrlsService) Create(masterProjectId string, zone string, createsignedurlsrequest *CreateSignedUrlsRequest) *MasterProjectsZonesSignedUrlsCreateCall {
	c := &MasterProjectsZonesSignedUrlsCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.masterProjectId = masterProjectId
	c.zone = zone
	c.createsignedurlsrequest = createsignedurlsrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MasterProjectsZonesSignedUrlsCreateCall) Fields(s ...googleapi.Field) *MasterProjectsZonesSignedUrlsCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *MasterProjectsZonesSignedUrlsCreateCall) Context(ctx context.Context) *MasterProjectsZonesSignedUrlsCreateCall {
	c.ctx_ = ctx
	return c
}

func (c *MasterProjectsZonesSignedUrlsCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.createsignedurlsrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/masterProjects/{masterProjectId}/zones/{zone}/signedUrls")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"masterProjectId": c.masterProjectId,
		"zone":            c.zone,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "container.masterProjects.zones.signedUrls.create" call.
// Exactly one of *SignedUrls or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *SignedUrls.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *MasterProjectsZonesSignedUrlsCreateCall) Do(opts ...googleapi.CallOption) (*SignedUrls, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &SignedUrls{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates signed URLs that allow for writing a file to a private GCS bucket\nfor storing backups of hosted master data. Signed URLs are explained here:\nhttps://cloud.google.com/storage/docs/access-control#Signed-URLs",
	//   "flatPath": "v1/masterProjects/{masterProjectId}/zones/{zone}/signedUrls",
	//   "httpMethod": "POST",
	//   "id": "container.masterProjects.zones.signedUrls.create",
	//   "parameterOrder": [
	//     "masterProjectId",
	//     "zone"
	//   ],
	//   "parameters": {
	//     "masterProjectId": {
	//       "description": "The hosted master project in which this master resides.  This can be\neither a [project ID or project\nnumber](https://support.google.com/cloud/answer/6158840).",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The zone of this master's cluster.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/masterProjects/{masterProjectId}/zones/{zone}/signedUrls",
	//   "request": {
	//     "$ref": "CreateSignedUrlsRequest"
	//   },
	//   "response": {
	//     "$ref": "SignedUrls"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/userinfo.email"
	//   ]
	// }

}

// method id "container.masterProjects.zones.tokens.create":

type MasterProjectsZonesTokensCreateCall struct {
	s                  *Service
	masterProjectId    string
	zone               string
	createtokenrequest *CreateTokenRequest
	urlParams_         gensupport.URLParams
	ctx_               context.Context
}

// Create: Creates a compute-read-write
// (https://www.googleapis.com/auth/compute)
// scoped OAuth2 access token for <project_number>, to allow a hosted
// master
// to make modifications to its user's project.
func (r *MasterProjectsZonesTokensService) Create(masterProjectId string, zone string, createtokenrequest *CreateTokenRequest) *MasterProjectsZonesTokensCreateCall {
	c := &MasterProjectsZonesTokensCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.masterProjectId = masterProjectId
	c.zone = zone
	c.createtokenrequest = createtokenrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MasterProjectsZonesTokensCreateCall) Fields(s ...googleapi.Field) *MasterProjectsZonesTokensCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *MasterProjectsZonesTokensCreateCall) Context(ctx context.Context) *MasterProjectsZonesTokensCreateCall {
	c.ctx_ = ctx
	return c
}

func (c *MasterProjectsZonesTokensCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.createtokenrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/masterProjects/{masterProjectId}/zones/{zone}/tokens")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"masterProjectId": c.masterProjectId,
		"zone":            c.zone,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "container.masterProjects.zones.tokens.create" call.
// Exactly one of *Token or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Token.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *MasterProjectsZonesTokensCreateCall) Do(opts ...googleapi.CallOption) (*Token, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Token{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a compute-read-write (https://www.googleapis.com/auth/compute)\nscoped OAuth2 access token for \u003cproject_number\u003e, to allow a hosted master\nto make modifications to its user's project.",
	//   "flatPath": "v1/masterProjects/{masterProjectId}/zones/{zone}/tokens",
	//   "httpMethod": "POST",
	//   "id": "container.masterProjects.zones.tokens.create",
	//   "parameterOrder": [
	//     "masterProjectId",
	//     "zone"
	//   ],
	//   "parameters": {
	//     "masterProjectId": {
	//       "description": "The hosted master project in which this master resides.  This can be\neither a [project ID or project\nnumber](https://support.google.com/cloud/answer/6158840).",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The zone of this master's cluster.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/masterProjects/{masterProjectId}/zones/{zone}/tokens",
	//   "request": {
	//     "$ref": "CreateTokenRequest"
	//   },
	//   "response": {
	//     "$ref": "Token"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/userinfo.email"
	//   ]
	// }

}

// method id "container.projects.zones.getServerconfig":

type ProjectsZonesGetServerconfigCall struct {
	s            *Service
	projectId    string
	zone         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// GetServerconfig: Returns configuration info about the Container
// Engine service.
func (r *ProjectsZonesService) GetServerconfig(projectId string, zone string) *ProjectsZonesGetServerconfigCall {
	c := &ProjectsZonesGetServerconfigCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.zone = zone
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsZonesGetServerconfigCall) Fields(s ...googleapi.Field) *ProjectsZonesGetServerconfigCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsZonesGetServerconfigCall) IfNoneMatch(entityTag string) *ProjectsZonesGetServerconfigCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsZonesGetServerconfigCall) Context(ctx context.Context) *ProjectsZonesGetServerconfigCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsZonesGetServerconfigCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/zones/{zone}/serverconfig")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"zone":      c.zone,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "container.projects.zones.getServerconfig" call.
// Exactly one of *ServerConfig or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *ServerConfig.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsZonesGetServerconfigCall) Do(opts ...googleapi.CallOption) (*ServerConfig, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ServerConfig{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns configuration info about the Container Engine service.",
	//   "flatPath": "v1/projects/{projectId}/zones/{zone}/serverconfig",
	//   "httpMethod": "GET",
	//   "id": "container.projects.zones.getServerconfig",
	//   "parameterOrder": [
	//     "projectId",
	//     "zone"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "The Google Developers Console [project ID or project\nnumber](https://support.google.com/cloud/answer/6158840).",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The name of the Google Compute Engine [zone](/compute/docs/zones#available)\nto return operations for.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/zones/{zone}/serverconfig",
	//   "response": {
	//     "$ref": "ServerConfig"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "container.projects.zones.clusters.create":

type ProjectsZonesClustersCreateCall struct {
	s                    *Service
	projectId            string
	zone                 string
	createclusterrequest *CreateClusterRequest
	urlParams_           gensupport.URLParams
	ctx_                 context.Context
}

// Create: Creates a cluster, consisting of the specified number and
// type of Google
// Compute Engine instances.
//
// By default, the cluster is created in the project's
// [default
// network](/compute/docs/networks-and-firewalls#networks).
//
// One firewall is added for the cluster. After cluster creation,
// the cluster creates routes for each node to allow the containers
// on that node to communicate with all other instances in
// the
// cluster.
//
// Finally, an entry is added to the project's global metadata
// indicating
// which CIDR range is being used by the cluster.
func (r *ProjectsZonesClustersService) Create(projectId string, zone string, createclusterrequest *CreateClusterRequest) *ProjectsZonesClustersCreateCall {
	c := &ProjectsZonesClustersCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.zone = zone
	c.createclusterrequest = createclusterrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsZonesClustersCreateCall) Fields(s ...googleapi.Field) *ProjectsZonesClustersCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsZonesClustersCreateCall) Context(ctx context.Context) *ProjectsZonesClustersCreateCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsZonesClustersCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.createclusterrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/zones/{zone}/clusters")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"zone":      c.zone,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "container.projects.zones.clusters.create" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsZonesClustersCreateCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Operation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a cluster, consisting of the specified number and type of Google\nCompute Engine instances.\n\nBy default, the cluster is created in the project's\n[default network](/compute/docs/networks-and-firewalls#networks).\n\nOne firewall is added for the cluster. After cluster creation,\nthe cluster creates routes for each node to allow the containers\non that node to communicate with all other instances in the\ncluster.\n\nFinally, an entry is added to the project's global metadata indicating\nwhich CIDR range is being used by the cluster.",
	//   "flatPath": "v1/projects/{projectId}/zones/{zone}/clusters",
	//   "httpMethod": "POST",
	//   "id": "container.projects.zones.clusters.create",
	//   "parameterOrder": [
	//     "projectId",
	//     "zone"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "The Google Developers Console [project ID or project\nnumber](https://support.google.com/cloud/answer/6158840).",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The name of the Google Compute Engine\n[zone](/compute/docs/zones#available) in which the cluster\nresides.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/zones/{zone}/clusters",
	//   "request": {
	//     "$ref": "CreateClusterRequest"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "container.projects.zones.clusters.delete":

type ProjectsZonesClustersDeleteCall struct {
	s          *Service
	projectId  string
	zone       string
	clusterId  string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
}

// Delete: Deletes the cluster, including the Kubernetes endpoint and
// all worker
// nodes.
//
// Firewalls and routes that were configured during cluster creation
// are also deleted.
//
// Other Google Compute Engine resources that might be in use by the
// cluster
// (e.g. load balancer resources) will not be deleted if they weren't
// present
// at the initial create time.
func (r *ProjectsZonesClustersService) Delete(projectId string, zone string, clusterId string) *ProjectsZonesClustersDeleteCall {
	c := &ProjectsZonesClustersDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.zone = zone
	c.clusterId = clusterId
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsZonesClustersDeleteCall) Fields(s ...googleapi.Field) *ProjectsZonesClustersDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsZonesClustersDeleteCall) Context(ctx context.Context) *ProjectsZonesClustersDeleteCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsZonesClustersDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/zones/{zone}/clusters/{clusterId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"zone":      c.zone,
		"clusterId": c.clusterId,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "container.projects.zones.clusters.delete" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsZonesClustersDeleteCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Operation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes the cluster, including the Kubernetes endpoint and all worker\nnodes.\n\nFirewalls and routes that were configured during cluster creation\nare also deleted.\n\nOther Google Compute Engine resources that might be in use by the cluster\n(e.g. load balancer resources) will not be deleted if they weren't present\nat the initial create time.",
	//   "flatPath": "v1/projects/{projectId}/zones/{zone}/clusters/{clusterId}",
	//   "httpMethod": "DELETE",
	//   "id": "container.projects.zones.clusters.delete",
	//   "parameterOrder": [
	//     "projectId",
	//     "zone",
	//     "clusterId"
	//   ],
	//   "parameters": {
	//     "clusterId": {
	//       "description": "The name of the cluster to delete.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "The Google Developers Console [project ID or project\nnumber](https://support.google.com/cloud/answer/6158840).",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The name of the Google Compute Engine\n[zone](/compute/docs/zones#available) in which the cluster\nresides.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/zones/{zone}/clusters/{clusterId}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "container.projects.zones.clusters.get":

type ProjectsZonesClustersGetCall struct {
	s            *Service
	projectId    string
	zone         string
	clusterId    string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// Get: Gets the details of a specific cluster.
func (r *ProjectsZonesClustersService) Get(projectId string, zone string, clusterId string) *ProjectsZonesClustersGetCall {
	c := &ProjectsZonesClustersGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.zone = zone
	c.clusterId = clusterId
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsZonesClustersGetCall) Fields(s ...googleapi.Field) *ProjectsZonesClustersGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsZonesClustersGetCall) IfNoneMatch(entityTag string) *ProjectsZonesClustersGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsZonesClustersGetCall) Context(ctx context.Context) *ProjectsZonesClustersGetCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsZonesClustersGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/zones/{zone}/clusters/{clusterId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"zone":      c.zone,
		"clusterId": c.clusterId,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "container.projects.zones.clusters.get" call.
// Exactly one of *Cluster or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Cluster.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsZonesClustersGetCall) Do(opts ...googleapi.CallOption) (*Cluster, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Cluster{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets the details of a specific cluster.",
	//   "flatPath": "v1/projects/{projectId}/zones/{zone}/clusters/{clusterId}",
	//   "httpMethod": "GET",
	//   "id": "container.projects.zones.clusters.get",
	//   "parameterOrder": [
	//     "projectId",
	//     "zone",
	//     "clusterId"
	//   ],
	//   "parameters": {
	//     "clusterId": {
	//       "description": "The name of the cluster to retrieve.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "The Google Developers Console [project ID or project\nnumber](https://support.google.com/cloud/answer/6158840).",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The name of the Google Compute Engine\n[zone](/compute/docs/zones#available) in which the cluster\nresides.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/zones/{zone}/clusters/{clusterId}",
	//   "response": {
	//     "$ref": "Cluster"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "container.projects.zones.clusters.list":

type ProjectsZonesClustersListCall struct {
	s            *Service
	projectId    string
	zone         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// List: Lists all clusters owned by a project in either the specified
// zone or all
// zones.
func (r *ProjectsZonesClustersService) List(projectId string, zone string) *ProjectsZonesClustersListCall {
	c := &ProjectsZonesClustersListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.zone = zone
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsZonesClustersListCall) Fields(s ...googleapi.Field) *ProjectsZonesClustersListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsZonesClustersListCall) IfNoneMatch(entityTag string) *ProjectsZonesClustersListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsZonesClustersListCall) Context(ctx context.Context) *ProjectsZonesClustersListCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsZonesClustersListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/zones/{zone}/clusters")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"zone":      c.zone,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "container.projects.zones.clusters.list" call.
// Exactly one of *ListClustersResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListClustersResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsZonesClustersListCall) Do(opts ...googleapi.CallOption) (*ListClustersResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListClustersResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists all clusters owned by a project in either the specified zone or all\nzones.",
	//   "flatPath": "v1/projects/{projectId}/zones/{zone}/clusters",
	//   "httpMethod": "GET",
	//   "id": "container.projects.zones.clusters.list",
	//   "parameterOrder": [
	//     "projectId",
	//     "zone"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "The Google Developers Console [project ID or project\nnumber](https://support.google.com/cloud/answer/6158840).",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The name of the Google Compute Engine\n[zone](/compute/docs/zones#available) in which the cluster\nresides, or \"-\" for all zones.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/zones/{zone}/clusters",
	//   "response": {
	//     "$ref": "ListClustersResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "container.projects.zones.clusters.update":

type ProjectsZonesClustersUpdateCall struct {
	s                    *Service
	projectId            string
	zone                 string
	clusterId            string
	updateclusterrequest *UpdateClusterRequest
	urlParams_           gensupport.URLParams
	ctx_                 context.Context
}

// Update: Updates the settings of a specific cluster.
func (r *ProjectsZonesClustersService) Update(projectId string, zone string, clusterId string, updateclusterrequest *UpdateClusterRequest) *ProjectsZonesClustersUpdateCall {
	c := &ProjectsZonesClustersUpdateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.zone = zone
	c.clusterId = clusterId
	c.updateclusterrequest = updateclusterrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsZonesClustersUpdateCall) Fields(s ...googleapi.Field) *ProjectsZonesClustersUpdateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsZonesClustersUpdateCall) Context(ctx context.Context) *ProjectsZonesClustersUpdateCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsZonesClustersUpdateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.updateclusterrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/zones/{zone}/clusters/{clusterId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"zone":      c.zone,
		"clusterId": c.clusterId,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "container.projects.zones.clusters.update" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsZonesClustersUpdateCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Operation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates the settings of a specific cluster.",
	//   "flatPath": "v1/projects/{projectId}/zones/{zone}/clusters/{clusterId}",
	//   "httpMethod": "PUT",
	//   "id": "container.projects.zones.clusters.update",
	//   "parameterOrder": [
	//     "projectId",
	//     "zone",
	//     "clusterId"
	//   ],
	//   "parameters": {
	//     "clusterId": {
	//       "description": "The name of the cluster to upgrade.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "The Google Developers Console [project ID or project\nnumber](https://support.google.com/cloud/answer/6158840).",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The name of the Google Compute Engine\n[zone](/compute/docs/zones#available) in which the cluster\nresides.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/zones/{zone}/clusters/{clusterId}",
	//   "request": {
	//     "$ref": "UpdateClusterRequest"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "container.projects.zones.clusters.nodePools.create":

type ProjectsZonesClustersNodePoolsCreateCall struct {
	s                     *Service
	projectId             string
	zone                  string
	clusterId             string
	createnodepoolrequest *CreateNodePoolRequest
	urlParams_            gensupport.URLParams
	ctx_                  context.Context
}

// Create: Creates a node pool for a cluster.
func (r *ProjectsZonesClustersNodePoolsService) Create(projectId string, zone string, clusterId string, createnodepoolrequest *CreateNodePoolRequest) *ProjectsZonesClustersNodePoolsCreateCall {
	c := &ProjectsZonesClustersNodePoolsCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.zone = zone
	c.clusterId = clusterId
	c.createnodepoolrequest = createnodepoolrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsZonesClustersNodePoolsCreateCall) Fields(s ...googleapi.Field) *ProjectsZonesClustersNodePoolsCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsZonesClustersNodePoolsCreateCall) Context(ctx context.Context) *ProjectsZonesClustersNodePoolsCreateCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsZonesClustersNodePoolsCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.createnodepoolrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/zones/{zone}/clusters/{clusterId}/nodePools")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"zone":      c.zone,
		"clusterId": c.clusterId,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "container.projects.zones.clusters.nodePools.create" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsZonesClustersNodePoolsCreateCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Operation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a node pool for a cluster.",
	//   "flatPath": "v1/projects/{projectId}/zones/{zone}/clusters/{clusterId}/nodePools",
	//   "httpMethod": "POST",
	//   "id": "container.projects.zones.clusters.nodePools.create",
	//   "parameterOrder": [
	//     "projectId",
	//     "zone",
	//     "clusterId"
	//   ],
	//   "parameters": {
	//     "clusterId": {
	//       "description": "The name of the cluster.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "The Google Developers Console [project ID or project\nnumber](https://developers.google.com/console/help/new/#projectnumber).",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The name of the Google Compute Engine\n[zone](/compute/docs/zones#available) in which the cluster\nresides.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/zones/{zone}/clusters/{clusterId}/nodePools",
	//   "request": {
	//     "$ref": "CreateNodePoolRequest"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "container.projects.zones.clusters.nodePools.delete":

type ProjectsZonesClustersNodePoolsDeleteCall struct {
	s          *Service
	projectId  string
	zone       string
	clusterId  string
	nodePoolId string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
}

// Delete: Deletes a node pool from a cluster.
func (r *ProjectsZonesClustersNodePoolsService) Delete(projectId string, zone string, clusterId string, nodePoolId string) *ProjectsZonesClustersNodePoolsDeleteCall {
	c := &ProjectsZonesClustersNodePoolsDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.zone = zone
	c.clusterId = clusterId
	c.nodePoolId = nodePoolId
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsZonesClustersNodePoolsDeleteCall) Fields(s ...googleapi.Field) *ProjectsZonesClustersNodePoolsDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsZonesClustersNodePoolsDeleteCall) Context(ctx context.Context) *ProjectsZonesClustersNodePoolsDeleteCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsZonesClustersNodePoolsDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/zones/{zone}/clusters/{clusterId}/nodePools/{nodePoolId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectId":  c.projectId,
		"zone":       c.zone,
		"clusterId":  c.clusterId,
		"nodePoolId": c.nodePoolId,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "container.projects.zones.clusters.nodePools.delete" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsZonesClustersNodePoolsDeleteCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Operation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes a node pool from a cluster.",
	//   "flatPath": "v1/projects/{projectId}/zones/{zone}/clusters/{clusterId}/nodePools/{nodePoolId}",
	//   "httpMethod": "DELETE",
	//   "id": "container.projects.zones.clusters.nodePools.delete",
	//   "parameterOrder": [
	//     "projectId",
	//     "zone",
	//     "clusterId",
	//     "nodePoolId"
	//   ],
	//   "parameters": {
	//     "clusterId": {
	//       "description": "The name of the cluster.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "nodePoolId": {
	//       "description": "The name of the node pool to delete.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "The Google Developers Console [project ID or project\nnumber](https://developers.google.com/console/help/new/#projectnumber).",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The name of the Google Compute Engine\n[zone](/compute/docs/zones#available) in which the cluster\nresides.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/zones/{zone}/clusters/{clusterId}/nodePools/{nodePoolId}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "container.projects.zones.clusters.nodePools.get":

type ProjectsZonesClustersNodePoolsGetCall struct {
	s            *Service
	projectId    string
	zone         string
	clusterId    string
	nodePoolId   string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// Get: Retrieves the node pool requested.
func (r *ProjectsZonesClustersNodePoolsService) Get(projectId string, zone string, clusterId string, nodePoolId string) *ProjectsZonesClustersNodePoolsGetCall {
	c := &ProjectsZonesClustersNodePoolsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.zone = zone
	c.clusterId = clusterId
	c.nodePoolId = nodePoolId
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsZonesClustersNodePoolsGetCall) Fields(s ...googleapi.Field) *ProjectsZonesClustersNodePoolsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsZonesClustersNodePoolsGetCall) IfNoneMatch(entityTag string) *ProjectsZonesClustersNodePoolsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsZonesClustersNodePoolsGetCall) Context(ctx context.Context) *ProjectsZonesClustersNodePoolsGetCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsZonesClustersNodePoolsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/zones/{zone}/clusters/{clusterId}/nodePools/{nodePoolId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectId":  c.projectId,
		"zone":       c.zone,
		"clusterId":  c.clusterId,
		"nodePoolId": c.nodePoolId,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "container.projects.zones.clusters.nodePools.get" call.
// Exactly one of *NodePool or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *NodePool.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsZonesClustersNodePoolsGetCall) Do(opts ...googleapi.CallOption) (*NodePool, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &NodePool{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves the node pool requested.",
	//   "flatPath": "v1/projects/{projectId}/zones/{zone}/clusters/{clusterId}/nodePools/{nodePoolId}",
	//   "httpMethod": "GET",
	//   "id": "container.projects.zones.clusters.nodePools.get",
	//   "parameterOrder": [
	//     "projectId",
	//     "zone",
	//     "clusterId",
	//     "nodePoolId"
	//   ],
	//   "parameters": {
	//     "clusterId": {
	//       "description": "The name of the cluster.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "nodePoolId": {
	//       "description": "The name of the node pool.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "The Google Developers Console [project ID or project\nnumber](https://developers.google.com/console/help/new/#projectnumber).",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The name of the Google Compute Engine\n[zone](/compute/docs/zones#available) in which the cluster\nresides.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/zones/{zone}/clusters/{clusterId}/nodePools/{nodePoolId}",
	//   "response": {
	//     "$ref": "NodePool"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "container.projects.zones.clusters.nodePools.list":

type ProjectsZonesClustersNodePoolsListCall struct {
	s            *Service
	projectId    string
	zone         string
	clusterId    string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// List: Lists the node pools for a cluster.
func (r *ProjectsZonesClustersNodePoolsService) List(projectId string, zone string, clusterId string) *ProjectsZonesClustersNodePoolsListCall {
	c := &ProjectsZonesClustersNodePoolsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.zone = zone
	c.clusterId = clusterId
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsZonesClustersNodePoolsListCall) Fields(s ...googleapi.Field) *ProjectsZonesClustersNodePoolsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsZonesClustersNodePoolsListCall) IfNoneMatch(entityTag string) *ProjectsZonesClustersNodePoolsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsZonesClustersNodePoolsListCall) Context(ctx context.Context) *ProjectsZonesClustersNodePoolsListCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsZonesClustersNodePoolsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/zones/{zone}/clusters/{clusterId}/nodePools")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"zone":      c.zone,
		"clusterId": c.clusterId,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "container.projects.zones.clusters.nodePools.list" call.
// Exactly one of *ListNodePoolsResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListNodePoolsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsZonesClustersNodePoolsListCall) Do(opts ...googleapi.CallOption) (*ListNodePoolsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListNodePoolsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists the node pools for a cluster.",
	//   "flatPath": "v1/projects/{projectId}/zones/{zone}/clusters/{clusterId}/nodePools",
	//   "httpMethod": "GET",
	//   "id": "container.projects.zones.clusters.nodePools.list",
	//   "parameterOrder": [
	//     "projectId",
	//     "zone",
	//     "clusterId"
	//   ],
	//   "parameters": {
	//     "clusterId": {
	//       "description": "The name of the cluster.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "The Google Developers Console [project ID or project\nnumber](https://developers.google.com/console/help/new/#projectnumber).",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The name of the Google Compute Engine\n[zone](/compute/docs/zones#available) in which the cluster\nresides.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/zones/{zone}/clusters/{clusterId}/nodePools",
	//   "response": {
	//     "$ref": "ListNodePoolsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "container.projects.zones.operations.get":

type ProjectsZonesOperationsGetCall struct {
	s            *Service
	projectId    string
	zone         string
	operationId  string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// Get: Gets the specified operation.
func (r *ProjectsZonesOperationsService) Get(projectId string, zone string, operationId string) *ProjectsZonesOperationsGetCall {
	c := &ProjectsZonesOperationsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.zone = zone
	c.operationId = operationId
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsZonesOperationsGetCall) Fields(s ...googleapi.Field) *ProjectsZonesOperationsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsZonesOperationsGetCall) IfNoneMatch(entityTag string) *ProjectsZonesOperationsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsZonesOperationsGetCall) Context(ctx context.Context) *ProjectsZonesOperationsGetCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsZonesOperationsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/zones/{zone}/operations/{operationId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectId":   c.projectId,
		"zone":        c.zone,
		"operationId": c.operationId,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "container.projects.zones.operations.get" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsZonesOperationsGetCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Operation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets the specified operation.",
	//   "flatPath": "v1/projects/{projectId}/zones/{zone}/operations/{operationId}",
	//   "httpMethod": "GET",
	//   "id": "container.projects.zones.operations.get",
	//   "parameterOrder": [
	//     "projectId",
	//     "zone",
	//     "operationId"
	//   ],
	//   "parameters": {
	//     "operationId": {
	//       "description": "The server-assigned `name` of the operation.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "The Google Developers Console [project ID or project\nnumber](https://support.google.com/cloud/answer/6158840).",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The name of the Google Compute Engine\n[zone](/compute/docs/zones#available) in which the cluster\nresides.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/zones/{zone}/operations/{operationId}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "container.projects.zones.operations.list":

type ProjectsZonesOperationsListCall struct {
	s            *Service
	projectId    string
	zone         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// List: Lists all operations in a project in a specific zone or all
// zones.
func (r *ProjectsZonesOperationsService) List(projectId string, zone string) *ProjectsZonesOperationsListCall {
	c := &ProjectsZonesOperationsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.zone = zone
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsZonesOperationsListCall) Fields(s ...googleapi.Field) *ProjectsZonesOperationsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsZonesOperationsListCall) IfNoneMatch(entityTag string) *ProjectsZonesOperationsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsZonesOperationsListCall) Context(ctx context.Context) *ProjectsZonesOperationsListCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsZonesOperationsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/zones/{zone}/operations")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"zone":      c.zone,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "container.projects.zones.operations.list" call.
// Exactly one of *ListOperationsResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListOperationsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsZonesOperationsListCall) Do(opts ...googleapi.CallOption) (*ListOperationsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListOperationsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists all operations in a project in a specific zone or all zones.",
	//   "flatPath": "v1/projects/{projectId}/zones/{zone}/operations",
	//   "httpMethod": "GET",
	//   "id": "container.projects.zones.operations.list",
	//   "parameterOrder": [
	//     "projectId",
	//     "zone"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "The Google Developers Console [project ID or project\nnumber](https://support.google.com/cloud/answer/6158840).",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The name of the Google Compute Engine [zone](/compute/docs/zones#available)\nto return operations for, or `-` for all zones.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/zones/{zone}/operations",
	//   "response": {
	//     "$ref": "ListOperationsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}
