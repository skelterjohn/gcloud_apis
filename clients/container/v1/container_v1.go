// Package container provides access to the Google Container Engine API.
//
// See https://cloud.google.com/container-engine/
//
// Usage example:
//
//   import "google.golang.org/api/container/v1"
//   ...
//   containerService, err := container.New(oauthHttpClient)
package container

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"google.golang.org/api/googleapi"
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
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace

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
	client   *http.Client
	BasePath string // API endpoint base URL

	MasterProjects *MasterProjectsService

	Projects *ProjectsService
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
	return rs
}

type ProjectsZonesClustersService struct {
	s *Service
}

func NewProjectsZonesOperationsService(s *Service) *ProjectsZonesOperationsService {
	rs := &ProjectsZonesOperationsService{s: s}
	return rs
}

type ProjectsZonesOperationsService struct {
	s *Service
}

type AddonsConfig struct {
	// HorizontalPodAutoscaling: Configuration for the horizontal pod
	// autoscaling feature, which increases
	// or decreases the number of
	// replicas a replication controller has based on
	// the resource usage of
	// the existing replicas.
	HorizontalPodAutoscaling *HorizontalPodAutoscaling `json:"horizontalPodAutoscaling,omitempty"`

	// HttpLoadBalancing: Configuration for the HTTP (L7) load balancing
	// controller addon, which
	// makes it easy to set up HTTP load balancers
	// for services in a cluster.
	HttpLoadBalancing *HttpLoadBalancing `json:"httpLoadBalancing,omitempty"`
}

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
	// one
	// automatically chosen or specify a `/14` block in `10.0.0.0/8`.
	ClusterIpv4Cidr string `json:"clusterIpv4Cidr,omitempty"`

	// CreateTime: [Output only] The time the cluster was created,
	// in
	// [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text format.
	CreateTime string `json:"createTime,omitempty"`

	// CurrentMasterVersion: [Output only] The current software version of
	// the master endpoint.
	CurrentMasterVersion string `json:"currentMasterVersion,omitempty"`

	// CurrentNodeCount: [Output only] The number of nodes currently
	// recognized by the cluster.
	CurrentNodeCount int64 `json:"currentNodeCount,omitempty"`

	// CurrentNodeVersion: [Output only] The current version of the node
	// software components.
	// If they are currently at different versions
	// because they're in the process
	// of being upgraded, this reflects the
	// minimum version of any of them.
	CurrentNodeVersion string `json:"currentNodeVersion,omitempty"`

	// Description: An optional description of this cluster.
	Description string `json:"description,omitempty"`

	// Endpoint: [Output only] The IP address of this cluster's master
	// endpoint.
	// The endpoint can be accessed from the internet
	// at
	// `https://username:password@endpoint/`.
	//
	// See the `masterAuth`
	// property of this resource for username and
	// password information.
	Endpoint string `json:"endpoint,omitempty"`

	// InitialClusterVersion: [Output only] The software version of the
	// master and kubelets used
	// in the cluster when it was first created.
	// The version can be upgraded over
	// time.
	//
	InitialClusterVersion string `json:"initialClusterVersion,omitempty"`

	// InitialNodeCount: The number of nodes to create in this cluster. You
	// must ensure that your
	// Compute Engine <a
	// href="/compute/docs/resource-quotas">resource quota</a>
	// is sufficient
	// for this number of instances. You must also have available
	// firewall
	// and routes quota.
	InitialNodeCount int64 `json:"initialNodeCount,omitempty"`

	// InstanceGroupUrls: [Output only] The resource URLs of
	// [instance
	// groups](/compute/docs/instance-groups/) associated with
	// this
	// cluster.
	InstanceGroupUrls []string `json:"instanceGroupUrls,omitempty"`

	// LoggingService: The logging service that the cluster should write
	// logs to.
	// Currently available options:
	//
	// * `logging.googleapis.com` -
	// the Google Cloud Logging service.
	// * `none` - no logs will be exported
	// from the cluster.
	// * "" - default value: the default is
	// `logging.googleapis.com`.
	LoggingService string `json:"loggingService,omitempty"`

	// MasterAuth: The authentication information for accessing the master.
	MasterAuth *MasterAuth `json:"masterAuth,omitempty"`

	// MonitoringService: The monitoring service that the cluster should
	// write metrics to.
	// Currently available options:
	//
	// *
	// `monitoring.googleapis.com` - the Google Cloud Monitoring service.
	// *
	// `none` - no metrics will be exported from the cluster.
	// * "" - default
	// value: the default is `monitoring.googleapis.com`.
	MonitoringService string `json:"monitoringService,omitempty"`

	// Name: The name of this cluster. The name must be unique within this
	// project
	// and zone, and can be up to 40 characters with the following
	// restrictions:
	//
	// * Lowercase letters, numbers, and hyphens only.
	// * Must
	// start with a letter.
	// * Must end with a number or a letter.
	Name string `json:"name,omitempty"`

	// Network: The name of the Google Compute
	// Engine
	// [network](/compute/docs/networks-and-firewalls#networks) to
	// which the
	// cluster is connected. If left unspecified, the `default`
	// network
	// will be used.
	Network string `json:"network,omitempty"`

	// NodeConfig: Parameters used in creating the cluster's nodes.
	// See the
	// descriptions of the child properties of `nodeConfig`.
	//
	// If
	// unspecified, the defaults for all child properties are used.
	NodeConfig *NodeConfig `json:"nodeConfig,omitempty"`

	// NodeIpv4CidrSize: [Output only] The size of the address space on each
	// node for hosting
	// containers. This is provisioned from within the
	// `container_ipv4_cidr` range.
	NodeIpv4CidrSize int64 `json:"nodeIpv4CidrSize,omitempty"`

	// SelfLink: [Output only] Server-defined URL for the resource.
	SelfLink string `json:"selfLink,omitempty"`

	// ServicesIpv4Cidr: [Output only] The IP address range of the
	// Kubernetes services in
	// this cluster,
	// in
	// [CIDR](http://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing)
	//
	// notation (e.g. `1.2.3.4/29`). Service addresses are
	// typically put in
	// the last `/16` from the container CIDR.
	ServicesIpv4Cidr string `json:"servicesIpv4Cidr,omitempty"`

	// Status: [Output only] The current status of this cluster.
	Status string `json:"status,omitempty"`

	// StatusMessage: [Output only] Additional information about the current
	// status of this
	// cluster, if available.
	StatusMessage string `json:"statusMessage,omitempty"`

	// Subnetwork: The name of the Google Compute
	// Engine
	// [subnetwork](/compute/docs/subnetworks) to which the
	// cluster
	// is connected. If specified, the cluster's network must be a
	// "custom
	// subnet" network. Specification of subnetworks is an alpha
	// feature,
	// and require that the Google Compute Engine alpha API be
	// enabled.
	Subnetwork string `json:"subnetwork,omitempty"`

	// Zone: [Output only] The name of the Google Compute
	// Engine
	// [zone](/compute/docs/zones#available) in which the
	// cluster
	// resides.
	Zone string `json:"zone,omitempty"`
}

type ClusterUpdate struct {
	// DesiredAddonsConfig: Configurations for the various addons available
	// to run in the cluster.
	DesiredAddonsConfig *AddonsConfig `json:"desiredAddonsConfig,omitempty"`

	// DesiredMasterMachineType: The name of a Google Compute Engine
	// [machine type](/compute/docs/machine-types)
	// (e.g. `n1-standard-8`) to
	// change the master to.
	DesiredMasterMachineType string `json:"desiredMasterMachineType,omitempty"`

	// DesiredMasterVersion: The Kubernetes version to change the master to
	// (typically an
	// upgrade). Use "-" to upgrade to the latest version
	// supported by
	// the server.
	DesiredMasterVersion string `json:"desiredMasterVersion,omitempty"`

	// DesiredMonitoringService: The monitoring service that the cluster
	// should write metrics to.
	// Currently available options:
	//
	// *
	// "monitoring.googleapis.com" - the Google Cloud Monitoring service
	// *
	// "none" - no metrics will be exported from the cluster
	DesiredMonitoringService string `json:"desiredMonitoringService,omitempty"`

	// DesiredNodeVersion: The Kubernetes version to change the nodes to
	// (typically an
	// upgrade). Use `-` to upgrade to the latest version
	// supported by
	// the server.
	DesiredNodeVersion string `json:"desiredNodeVersion,omitempty"`
}

type CreateClusterRequest struct {
	// Cluster: A [cluster
	// resource](/container-engine/reference/rest/v1/projects.zones.clusters)
	Cluster *Cluster `json:"cluster,omitempty"`
}

type CreateSignedUrlsRequest struct {
	// ClusterId: The name of this master's cluster.
	ClusterId string `json:"clusterId,omitempty"`

	// Filenames: The names of the files for which a signed URLs are being
	// requested.
	Filenames []string `json:"filenames,omitempty"`

	// ProjectNumber: The project number for which the signed URLs are being
	// requested.  This is
	// the project in which this master's cluster
	// resides.
	//
	// Note that this must be a project number, not a project ID.
	ProjectNumber int64 `json:"projectNumber,omitempty,string"`
}

type CreateTokenRequest struct {
	// ClusterId: The name of this master's cluster.
	ClusterId string `json:"clusterId,omitempty"`

	// ProjectNumber: The project number for which the access is being
	// requested.  This is the
	// project in which this master's cluster
	// resides.
	//
	// Note that this must be a project number, not a project ID.
	ProjectNumber int64 `json:"projectNumber,omitempty,string"`
}

type HorizontalPodAutoscaling struct {
	// Disabled: Whether the Horizontal Pod Autoscaling feature is enabled
	// in the cluster.
	// When enabled, it ensures that a Heapster pod is
	// running in the cluster,
	// which is also used by the Cloud Monitoring
	// service.
	Disabled bool `json:"disabled,omitempty"`
}

type HttpLoadBalancing struct {
	// Disabled: Whether the HTTP Load Balancing controller is enabled in
	// the cluster.
	// When enabled, it runs a small pod in the cluster that
	// manages the load
	// balancers.
	Disabled bool `json:"disabled,omitempty"`
}

type ListClustersResponse struct {
	// Clusters: A list of clusters in the project in the specified zone,
	// or
	// across all ones.
	Clusters []*Cluster `json:"clusters,omitempty"`
}

type ListOperationsResponse struct {
	// Operations: A list of operations in the project in the specified
	// zone.
	Operations []*Operation `json:"operations,omitempty"`
}

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

	// Password: The password to use for HTTP basic authentication when
	// accessing the
	// Kubernetes master endpoint. Because the master endpoint
	// is open to
	// the internet, you should create a strong password.
	Password string `json:"password,omitempty"`

	// Username: The username to use for HTTP basic authentication when
	// accessing the
	// Kubernetes master endpoint.
	Username string `json:"username,omitempty"`
}

type NodeConfig struct {
	// DiskSizeGb: Size of the disk attached to each node, specified in
	// GB.
	// The smallest allowed disk size is 10GB.
	//
	// If unspecified, the
	// default disk size is 100GB.
	DiskSizeGb int64 `json:"diskSizeGb,omitempty"`

	// MachineType: The name of a Google Compute Engine [machine
	// type](/compute/docs/machine-types) (e.g.
	// `n1-standard-1`).
	//
	// If
	// unspecified, the default machine type is
	// `n1-standard-1`.
	MachineType string `json:"machineType,omitempty"`

	// Metadata: The metadata key/value pairs assigned to instances in the
	// cluster.
	//
	// Keys must conform to the regexp [a-zA-Z0-9-_]+ and be less
	// than 128 bytes
	// in length. These are reflected as part of a URL in the
	// metadata server.
	// Additionally, to avoid ambiguity, keys must not
	// conflict with any other
	// metadata keys for the project or be one of
	// the four reserved keys:
	// "instance-template", "kube-env",
	// "startup-script", and "user-data"
	//
	// Values can be free-form strings,
	// and only have meaning as interpreted by
	// the image running in the
	// instance. The only restriction placed on them is
	// that each value's
	// size must be less than or equal to 32768 bytes.
	//
	// The total size of
	// all keys and values must be less than 512 KB.
	Metadata map[string]string `json:"metadata,omitempty"`

	// OauthScopes: The set of Google API scopes to be made available on all
	// of the
	// node VMs under the "default" service account.
	//
	// The following
	// scopes are recommended, but not required, and by default are
	// not
	// included:
	//
	// * `https://www.googleapis.com/auth/compute` is required
	// for mounting
	// persistent storage on your nodes.
	// *
	// `https://www.googleapis.com/auth/devstorage.read_only` is required
	// for
	// communicating with **gcr.io**
	// (the [Google Container
	// Registry](/container-registry/).
	//
	// If unspecified, no scopes are
	// added.
	OauthScopes []string `json:"oauthScopes,omitempty"`
}

type Operation struct {
	// Detail: Detailed operation progress, if available.
	Detail string `json:"detail,omitempty"`

	// Name: The server-assigned ID for the operation.
	Name string `json:"name,omitempty"`

	// OperationType: The operation type.
	OperationType string `json:"operationType,omitempty"`

	// SelfLink: Server-defined URL for the resource.
	SelfLink string `json:"selfLink,omitempty"`

	// Status: The current status of the operation.
	Status string `json:"status,omitempty"`

	// StatusMessage: If an error has occurred, a textual description of the
	// error.
	StatusMessage string `json:"statusMessage,omitempty"`

	// TargetLink: Server-defined URL for the target of the operation.
	TargetLink string `json:"targetLink,omitempty"`

	// Zone: The name of the Google Compute
	// Engine
	// [zone](/compute/docs/zones#available) in which the
	// operation
	// is taking place.
	Zone string `json:"zone,omitempty"`
}

type ServerConfig struct {
	// BuildClientInfo: apiserver build BuildData::ClientInfo()
	BuildClientInfo string `json:"buildClientInfo,omitempty"`

	// DefaultClusterVersion: What version this server deploys by default.
	DefaultClusterVersion string `json:"defaultClusterVersion,omitempty"`

	// ValidNodeVersions: List of valid node upgrade target versions.
	ValidNodeVersions []string `json:"validNodeVersions,omitempty"`
}

type SignedUrls struct {
	// SignedUrls: The signed URLs for writing the request files, in the
	// same order as the
	// filenames in the request.
	SignedUrls []string `json:"signedUrls,omitempty"`
}

type Token struct {
	// AccessToken: The OAuth2 access token
	AccessToken string `json:"accessToken,omitempty"`

	// ExpireTime: The expiration time of the token.
	ExpireTime string `json:"expireTime,omitempty"`
}

type UpdateClusterRequest struct {
	// Update: A description of the update.
	Update *ClusterUpdate `json:"update,omitempty"`
}

// method id "container.masterProjects.zones.signedUrls.create":

type MasterProjectsZonesSignedUrlsCreateCall struct {
	s                       *Service
	masterProjectId         string
	zone                    string
	createsignedurlsrequest *CreateSignedUrlsRequest
	opt_                    map[string]interface{}
}

// Create: Creates signed URLs that allow for writing a file to a
// private GCS bucket
// for storing backups of hosted master data. Signed
// URLs are explained
// here:
// https://cloud.google.com/storage/docs/access-control#Signed-URLs
func (r *MasterProjectsZonesSignedUrlsService) Create(masterProjectId string, zone string, createsignedurlsrequest *CreateSignedUrlsRequest) *MasterProjectsZonesSignedUrlsCreateCall {
	c := &MasterProjectsZonesSignedUrlsCreateCall{s: r.s, opt_: make(map[string]interface{})}
	c.masterProjectId = masterProjectId
	c.zone = zone
	c.createsignedurlsrequest = createsignedurlsrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MasterProjectsZonesSignedUrlsCreateCall) Fields(s ...googleapi.Field) *MasterProjectsZonesSignedUrlsCreateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *MasterProjectsZonesSignedUrlsCreateCall) Do() (*SignedUrls, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.createsignedurlsrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/masterProjects/{masterProjectId}/zones/{zone}/signedUrls")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"masterProjectId": c.masterProjectId,
		"zone":            c.zone,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *SignedUrls
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
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
	//       "description": "The hosted master project in which this master resides.  This can be\neither a [project ID or project\nnumber](https://developers.google.com/console/help/new/#projectnumber).",
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
	opt_               map[string]interface{}
}

// Create: Creates a compute-read-write
// (https://www.googleapis.com/auth/compute)
// scoped OAuth2 access token
// for <project_number>, to allow a hosted master
// to make modifications
// to its user's project.
func (r *MasterProjectsZonesTokensService) Create(masterProjectId string, zone string, createtokenrequest *CreateTokenRequest) *MasterProjectsZonesTokensCreateCall {
	c := &MasterProjectsZonesTokensCreateCall{s: r.s, opt_: make(map[string]interface{})}
	c.masterProjectId = masterProjectId
	c.zone = zone
	c.createtokenrequest = createtokenrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MasterProjectsZonesTokensCreateCall) Fields(s ...googleapi.Field) *MasterProjectsZonesTokensCreateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *MasterProjectsZonesTokensCreateCall) Do() (*Token, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.createtokenrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/masterProjects/{masterProjectId}/zones/{zone}/tokens")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"masterProjectId": c.masterProjectId,
		"zone":            c.zone,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Token
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
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
	//       "description": "The hosted master project in which this master resides.  This can be\neither a [project ID or project\nnumber](https://developers.google.com/console/help/new/#projectnumber).",
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
	s         *Service
	projectId string
	zone      string
	opt_      map[string]interface{}
}

// GetServerconfig: Returns configuration info about the Container
// Engine service.
func (r *ProjectsZonesService) GetServerconfig(projectId string, zone string) *ProjectsZonesGetServerconfigCall {
	c := &ProjectsZonesGetServerconfigCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.zone = zone
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsZonesGetServerconfigCall) Fields(s ...googleapi.Field) *ProjectsZonesGetServerconfigCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsZonesGetServerconfigCall) Do() (*ServerConfig, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/zones/{zone}/serverconfig")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"zone":      c.zone,
	})
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ServerConfig
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
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
	//       "description": "The Google Developers Console [project ID or project\nnumber](https://developers.google.com/console/help/new/#projectnumber).",
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
	opt_                 map[string]interface{}
}

// Create: Creates a cluster, consisting of the specified number and
// type of Google
// Compute Engine instances.
//
// By default, the cluster is
// created in the project's
// [default
// network](/compute/docs/networks-and-firewalls#networks).
//
// One
// firewall is added for the cluster. After cluster creation,
// the
// cluster creates routes for each node to allow the containers
// on that
// node to communicate with all other instances in
// the
// cluster.
//
// Finally, an entry is added to the project's global
// metadata indicating
// which CIDR range is being used by the cluster.
func (r *ProjectsZonesClustersService) Create(projectId string, zone string, createclusterrequest *CreateClusterRequest) *ProjectsZonesClustersCreateCall {
	c := &ProjectsZonesClustersCreateCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.zone = zone
	c.createclusterrequest = createclusterrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsZonesClustersCreateCall) Fields(s ...googleapi.Field) *ProjectsZonesClustersCreateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsZonesClustersCreateCall) Do() (*Operation, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.createclusterrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/zones/{zone}/clusters")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"zone":      c.zone,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Operation
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
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
	s         *Service
	projectId string
	zone      string
	clusterId string
	opt_      map[string]interface{}
}

// Delete: Deletes the cluster, including the Kubernetes endpoint and
// all worker
// nodes.
//
// Firewalls and routes that were configured during
// cluster creation
// are also deleted.
func (r *ProjectsZonesClustersService) Delete(projectId string, zone string, clusterId string) *ProjectsZonesClustersDeleteCall {
	c := &ProjectsZonesClustersDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.zone = zone
	c.clusterId = clusterId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsZonesClustersDeleteCall) Fields(s ...googleapi.Field) *ProjectsZonesClustersDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsZonesClustersDeleteCall) Do() (*Operation, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/zones/{zone}/clusters/{clusterId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"zone":      c.zone,
		"clusterId": c.clusterId,
	})
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Operation
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes the cluster, including the Kubernetes endpoint and all worker\nnodes.\n\nFirewalls and routes that were configured during cluster creation\nare also deleted.",
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
	s         *Service
	projectId string
	zone      string
	clusterId string
	opt_      map[string]interface{}
}

// Get: Gets a specific cluster.
func (r *ProjectsZonesClustersService) Get(projectId string, zone string, clusterId string) *ProjectsZonesClustersGetCall {
	c := &ProjectsZonesClustersGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.zone = zone
	c.clusterId = clusterId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsZonesClustersGetCall) Fields(s ...googleapi.Field) *ProjectsZonesClustersGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsZonesClustersGetCall) Do() (*Cluster, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/zones/{zone}/clusters/{clusterId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"zone":      c.zone,
		"clusterId": c.clusterId,
	})
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Cluster
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets a specific cluster.",
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
	s         *Service
	projectId string
	zone      string
	opt_      map[string]interface{}
}

// List: Lists all clusters owned by a project in either the specified
// zone or all zones.
func (r *ProjectsZonesClustersService) List(projectId string, zone string) *ProjectsZonesClustersListCall {
	c := &ProjectsZonesClustersListCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.zone = zone
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsZonesClustersListCall) Fields(s ...googleapi.Field) *ProjectsZonesClustersListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsZonesClustersListCall) Do() (*ListClustersResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/zones/{zone}/clusters")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"zone":      c.zone,
	})
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ListClustersResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists all clusters owned by a project in either the specified zone or all zones.",
	//   "flatPath": "v1/projects/{projectId}/zones/{zone}/clusters",
	//   "httpMethod": "GET",
	//   "id": "container.projects.zones.clusters.list",
	//   "parameterOrder": [
	//     "projectId",
	//     "zone"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "The Google Developers Console [project ID or project\nnumber](https://developers.google.com/console/help/new/#projectnumber).",
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
	opt_                 map[string]interface{}
}

// Update: Updates settings of a specific cluster.
func (r *ProjectsZonesClustersService) Update(projectId string, zone string, clusterId string, updateclusterrequest *UpdateClusterRequest) *ProjectsZonesClustersUpdateCall {
	c := &ProjectsZonesClustersUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.zone = zone
	c.clusterId = clusterId
	c.updateclusterrequest = updateclusterrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsZonesClustersUpdateCall) Fields(s ...googleapi.Field) *ProjectsZonesClustersUpdateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsZonesClustersUpdateCall) Do() (*Operation, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.updateclusterrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/zones/{zone}/clusters/{clusterId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"zone":      c.zone,
		"clusterId": c.clusterId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Operation
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates settings of a specific cluster.",
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

// method id "container.projects.zones.operations.get":

type ProjectsZonesOperationsGetCall struct {
	s           *Service
	projectId   string
	zone        string
	operationId string
	opt_        map[string]interface{}
}

// Get: Gets the specified operation.
func (r *ProjectsZonesOperationsService) Get(projectId string, zone string, operationId string) *ProjectsZonesOperationsGetCall {
	c := &ProjectsZonesOperationsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.zone = zone
	c.operationId = operationId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsZonesOperationsGetCall) Fields(s ...googleapi.Field) *ProjectsZonesOperationsGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsZonesOperationsGetCall) Do() (*Operation, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/zones/{zone}/operations/{operationId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId":   c.projectId,
		"zone":        c.zone,
		"operationId": c.operationId,
	})
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Operation
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
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
	s         *Service
	projectId string
	zone      string
	opt_      map[string]interface{}
}

// List: Lists all operations in a project in a specific zone or all
// zones.
func (r *ProjectsZonesOperationsService) List(projectId string, zone string) *ProjectsZonesOperationsListCall {
	c := &ProjectsZonesOperationsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.zone = zone
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsZonesOperationsListCall) Fields(s ...googleapi.Field) *ProjectsZonesOperationsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsZonesOperationsListCall) Do() (*ListOperationsResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/zones/{zone}/operations")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"zone":      c.zone,
	})
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ListOperationsResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
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
	//       "description": "The Google Developers Console [project ID or project\nnumber](https://developers.google.com/console/help/new/#projectnumber).",
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