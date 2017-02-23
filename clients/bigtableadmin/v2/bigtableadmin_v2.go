// Package bigtableadmin provides access to the Cloud Bigtable Admin API.
//
// See https://cloud.google.com/bigtable/
//
// Usage example:
//
//   import "github.com/skelterjohn/gcloud_apis/clients/bigtableadmin/v2"
//   ...
//   bigtableadminService, err := bigtableadmin.New(oauthHttpClient)
package bigtableadmin

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

const apiId = "bigtableadmin:v2"
const apiName = "bigtableadmin"
const apiVersion = "v2"
const basePath = "https://bigtableadmin.googleapis.com/"

// OAuth2 scopes used by this API.
const (
	// View and manage your data across Google Cloud Platform services
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"

	// View your data across Google Cloud Platform services
	CloudPlatformReadOnlyScope = "https://www.googleapis.com/auth/cloud-platform.read-only"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Operations = NewOperationsService(s)
	s.Projects = NewProjectsService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Operations *OperationsService

	Projects *ProjectsService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewOperationsService(s *Service) *OperationsService {
	rs := &OperationsService{s: s}
	return rs
}

type OperationsService struct {
	s *Service
}

func NewProjectsService(s *Service) *ProjectsService {
	rs := &ProjectsService{s: s}
	rs.Instances = NewProjectsInstancesService(s)
	return rs
}

type ProjectsService struct {
	s *Service

	Instances *ProjectsInstancesService
}

func NewProjectsInstancesService(s *Service) *ProjectsInstancesService {
	rs := &ProjectsInstancesService{s: s}
	rs.Clusters = NewProjectsInstancesClustersService(s)
	rs.Tables = NewProjectsInstancesTablesService(s)
	return rs
}

type ProjectsInstancesService struct {
	s *Service

	Clusters *ProjectsInstancesClustersService

	Tables *ProjectsInstancesTablesService
}

func NewProjectsInstancesClustersService(s *Service) *ProjectsInstancesClustersService {
	rs := &ProjectsInstancesClustersService{s: s}
	return rs
}

type ProjectsInstancesClustersService struct {
	s *Service
}

func NewProjectsInstancesTablesService(s *Service) *ProjectsInstancesTablesService {
	rs := &ProjectsInstancesTablesService{s: s}
	return rs
}

type ProjectsInstancesTablesService struct {
	s *Service
}

// Cluster: A resizable group of nodes in a particular cloud location,
// capable
// of serving all Tables in the parent
// Instance.
type Cluster struct {
	// DefaultStorageType: (`CreationOnly`)
	// The type of storage used by this cluster to serve its
	// parent instance's tables, unless explicitly overridden.
	//
	// Possible values:
	//   "STORAGE_TYPE_UNSPECIFIED" - The user did not specify a storage
	// type.
	//   "SSD" - Flash (SSD) storage should be used.
	//   "HDD" - Magnetic drive (HDD) storage should be used.
	DefaultStorageType string `json:"defaultStorageType,omitempty"`

	// Location: (`CreationOnly`)
	// The location where this cluster's nodes and storage reside. For
	// best
	// performance, clients should be located as close as possible to this
	// cluster.
	// Currently only zones are supported, so values should be of the
	// form
	// `projects/<project>/locations/<zone>`.
	Location string `json:"location,omitempty"`

	// Name: (`OutputOnly`)
	// The unique name of the cluster. Values are of the
	// form
	// `projects/<project>/instances/<instance>/clusters/a-z*`.
	Name string `json:"name,omitempty"`

	// ServeNodes: The number of nodes allocated to this cluster. More nodes
	// enable higher
	// throughput and more consistent performance.
	ServeNodes int64 `json:"serveNodes,omitempty"`

	// State: (`OutputOnly`)
	// The current state of the cluster.
	//
	// Possible values:
	//   "STATE_NOT_KNOWN" - The state of the cluster could not be
	// determined.
	//   "READY" - The cluster has been successfully created and is ready to
	// serve requests.
	//   "CREATING" - The cluster is currently being created, and may be
	// destroyed
	// if the creation process encounters an error.
	// A cluster may not be able to serve requests while being created.
	//   "RESIZING" - The cluster is currently being resized, and may revert
	// to its previous
	// node count if the process encounters an error.
	// A cluster is still capable of serving requests while being
	// resized,
	// but may exhibit performance as if its number of allocated nodes
	// is
	// between the starting and requested states.
	//   "DISABLED" - The cluster has no backing nodes. The data (tables)
	// still
	// exist, but no operations can be performed on the cluster.
	State string `json:"state,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "DefaultStorageType")
	// to unconditionally include in API requests. By default, fields with
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

// ColumnFamily: A set of columns within a table which share a common
// configuration.
type ColumnFamily struct {
	// GcRule: Garbage collection rule specified as a protobuf.
	// Must serialize to at most 500 bytes.
	//
	// NOTE: Garbage collection executes opportunistically in the
	// background, and
	// so it's possible for reads to return a cell even if it matches the
	// active
	// GC expression for its family.
	GcRule *GcRule `json:"gcRule,omitempty"`

	// ForceSendFields is a list of field names (e.g. "GcRule") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *ColumnFamily) MarshalJSON() ([]byte, error) {
	type noMethod ColumnFamily
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// CreateInstanceMetadata: The metadata for the Operation returned by
// CreateInstance.
type CreateInstanceMetadata struct {
	// FinishTime: The time at which the operation failed or was completed
	// successfully.
	FinishTime string `json:"finishTime,omitempty"`

	// OriginalRequest: The request that prompted the initiation of this
	// CreateInstance operation.
	OriginalRequest *CreateInstanceRequest `json:"originalRequest,omitempty"`

	// RequestTime: The time at which the original request was received.
	RequestTime string `json:"requestTime,omitempty"`

	// ForceSendFields is a list of field names (e.g. "FinishTime") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *CreateInstanceMetadata) MarshalJSON() ([]byte, error) {
	type noMethod CreateInstanceMetadata
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// CreateInstanceRequest: Request message for
// BigtableInstanceAdmin.CreateInstance.
type CreateInstanceRequest struct {
	// Clusters: The clusters to be created within the instance, mapped by
	// desired
	// cluster ID, e.g., just `mycluster` rather
	// than
	// `projects/myproject/instances/myinstance/clusters/mycluster`.
	// Fie
	// lds marked `OutputOnly` must be left blank.
	// Currently exactly one cluster must be specified.
	Clusters map[string]Cluster `json:"clusters,omitempty"`

	// Instance: The instance to create.
	// Fields marked `OutputOnly` must be left blank.
	Instance *Instance `json:"instance,omitempty"`

	// InstanceId: The ID to be used when referring to the new instance
	// within its project,
	// e.g., just `myinstance` rather
	// than
	// `projects/myproject/instances/myinstance`.
	InstanceId string `json:"instanceId,omitempty"`

	// Parent: The unique name of the project in which to create the new
	// instance.
	// Values are of the form `projects/<project>`.
	Parent string `json:"parent,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Clusters") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *CreateInstanceRequest) MarshalJSON() ([]byte, error) {
	type noMethod CreateInstanceRequest
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// CreateTableRequest: Request message
// for
// google.bigtable.admin.v2.BigtableTableAdmin.CreateTable
type CreateTableRequest struct {
	// InitialSplits: The optional list of row keys that will be used to
	// initially split the
	// table into several tablets (tablets are similar to HBase
	// regions).
	// Given two split keys, `s1` and `s2`, three tablets will be
	// created,
	// spanning the key ranges: `[, s1), [s1, s2), [s2, )`.
	//
	// Example:
	//
	// * Row keys := `["a", "apple", "custom", "customer_1", "customer_2",`
	//                "other", "zz"]`
	// * initial_split_keys := `["apple", "customer_1", "customer_2",
	// "other"]`
	// * Key assignment:
	//     - Tablet 1 `[, apple)                => {"a"}.`
	//     - Tablet 2 `[apple, customer_1)      => {"apple", "custom"}.`
	//     - Tablet 3 `[customer_1, customer_2) => {"customer_1"}.`
	//     - Tablet 4 `[customer_2, other)      => {"customer_2"}.`
	//     - Tablet 5 `[other, )                => {"other", "zz"}.`
	InitialSplits []*Split `json:"initialSplits,omitempty"`

	// Table: The Table to create.
	Table *Table `json:"table,omitempty"`

	// TableId: The name by which the new table should be referred to within
	// the parent
	// instance, e.g., `foobar` rather than `<parent>/tables/foobar`.
	TableId string `json:"tableId,omitempty"`

	// ForceSendFields is a list of field names (e.g. "InitialSplits") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *CreateTableRequest) MarshalJSON() ([]byte, error) {
	type noMethod CreateTableRequest
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// DropRowRangeRequest: Request message
// for
// google.bigtable.admin.v2.BigtableTableAdmin.DropRowRange
type DropRowRangeRequest struct {
	// DeleteAllDataFromTable: Delete all rows in the table. Setting this to
	// false is a no-op.
	DeleteAllDataFromTable bool `json:"deleteAllDataFromTable,omitempty"`

	// RowKeyPrefix: Delete all rows that start with this row key prefix.
	// Prefix cannot be
	// zero length.
	RowKeyPrefix string `json:"rowKeyPrefix,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "DeleteAllDataFromTable") to unconditionally include in API requests.
	// By default, fields with empty values are omitted from API requests.
	// However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *DropRowRangeRequest) MarshalJSON() ([]byte, error) {
	type noMethod DropRowRangeRequest
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// Empty: A generic empty message that you can re-use to avoid defining
// duplicated
// empty messages in your APIs. A typical example is to use it as the
// request
// or the response type of an API method. For instance:
//
//     service Foo {
//       rpc Bar(google.protobuf.Empty) returns
// (google.protobuf.Empty);
//     }
//
// The JSON representation for `Empty` is empty JSON object `{}`.
type Empty struct {
	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`
}

// GcRule: Rule for determining which cells to delete during garbage
// collection.
type GcRule struct {
	// Intersection: Delete cells that would be deleted by every nested
	// rule.
	Intersection *Intersection `json:"intersection,omitempty"`

	// MaxAge: Delete cells in a column older than the given age.
	// Values must be at least one millisecond, and will be truncated
	// to
	// microsecond granularity.
	MaxAge string `json:"maxAge,omitempty"`

	// MaxNumVersions: Delete all cells in a column except the most recent
	// N.
	MaxNumVersions int64 `json:"maxNumVersions,omitempty"`

	// Union: Delete cells that would be deleted by any nested rule.
	Union *Union `json:"union,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Intersection") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *GcRule) MarshalJSON() ([]byte, error) {
	type noMethod GcRule
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// Instance: A collection of Bigtable Tables and
// the resources that serve them.
// All tables in an instance are served from a single
// Cluster.
type Instance struct {
	// DisplayName: The descriptive name for this instance as it appears in
	// UIs.
	// Can be changed at any time, but should be kept globally unique
	// to avoid confusion.
	DisplayName string `json:"displayName,omitempty"`

	// Name: (`OutputOnly`)
	// The unique name of the instance. Values are of the
	// form
	// `projects/<project>/instances/a-z+[a-z0-9]`.
	Name string `json:"name,omitempty"`

	// State: (`OutputOnly`)
	// The current state of the instance.
	//
	// Possible values:
	//   "STATE_NOT_KNOWN" - The state of the instance could not be
	// determined.
	//   "READY" - The instance has been successfully created and can serve
	// requests
	// to its tables.
	//   "CREATING" - The instance is currently being created, and may be
	// destroyed
	// if the creation process encounters an error.
	State string `json:"state,omitempty"`

	// Type: The type of the instance. Defaults to `PRODUCTION`.
	//
	// Possible values:
	//   "TYPE_UNSPECIFIED" - The type of the instance is unspecified. If
	// set when creating an
	// instance, a `PRODUCTION` instance will be created. If set when
	// updating
	// an instance, the type will be left unchanged.
	//   "PRODUCTION" - An instance meant for production use. `serve_nodes`
	// must be set
	// on the cluster.
	//   "DEVELOPMENT" - The instance is meant for development and testing
	// purposes only; it has
	// no performance or uptime guarantees and is not covered by SLA.
	// After a development instance is created, it can be upgraded
	// by
	// updating the instance to type `PRODUCTION`. An instance created
	// as a production instance cannot be changed to a development
	// instance.
	// When creating a development instance, `serve_nodes` on the cluster
	// must
	// not be set.
	Type string `json:"type,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "DisplayName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Instance) MarshalJSON() ([]byte, error) {
	type noMethod Instance
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// Intersection: A GcRule which deletes cells matching all of the given
// rules.
type Intersection struct {
	// Rules: Only delete cells which would be deleted by every element of
	// `rules`.
	Rules []*GcRule `json:"rules,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Rules") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Intersection) MarshalJSON() ([]byte, error) {
	type noMethod Intersection
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// ListClustersResponse: Response message for
// BigtableInstanceAdmin.ListClusters.
type ListClustersResponse struct {
	// Clusters: The list of requested clusters.
	Clusters []*Cluster `json:"clusters,omitempty"`

	// FailedLocations: Locations from which Cluster information could not
	// be retrieved,
	// due to an outage or some other transient condition.
	// Clusters from these locations may be missing from `clusters`,
	// or may only have partial information returned.
	FailedLocations []string `json:"failedLocations,omitempty"`

	// NextPageToken: Set if not all clusters could be returned in a single
	// response.
	// Pass this value to `page_token` in another request to get the
	// next
	// page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`

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

// ListInstancesResponse: Response message for
// BigtableInstanceAdmin.ListInstances.
type ListInstancesResponse struct {
	// FailedLocations: Locations from which Instance information could not
	// be retrieved,
	// due to an outage or some other transient condition.
	// Instances whose Clusters are all in one of the failed locations
	// may be missing from `instances`, and Instances with at least
	// one
	// Cluster in a failed location may only have partial information
	// returned.
	FailedLocations []string `json:"failedLocations,omitempty"`

	// Instances: The list of requested instances.
	Instances []*Instance `json:"instances,omitempty"`

	// NextPageToken: Set if not all instances could be returned in a single
	// response.
	// Pass this value to `page_token` in another request to get the
	// next
	// page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "FailedLocations") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *ListInstancesResponse) MarshalJSON() ([]byte, error) {
	type noMethod ListInstancesResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// ListOperationsResponse: The response message for
// Operations.ListOperations.
type ListOperationsResponse struct {
	// NextPageToken: The standard List next-page token.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Operations: A list of operations that matches the specified filter in
	// the request.
	Operations []*Operation `json:"operations,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "NextPageToken") to
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

// ListTablesResponse: Response message
// for
// google.bigtable.admin.v2.BigtableTableAdmin.ListTables
type ListTablesResponse struct {
	// NextPageToken: Set if not all tables could be returned in a single
	// response.
	// Pass this value to `page_token` in another request to get the
	// next
	// page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Tables: The tables present in the requested instance.
	Tables []*Table `json:"tables,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "NextPageToken") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *ListTablesResponse) MarshalJSON() ([]byte, error) {
	type noMethod ListTablesResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// Modification: A create, update, or delete of a particular column
// family.
type Modification struct {
	// Create: Create a new column family with the specified schema, or fail
	// if
	// one already exists with the given ID.
	Create *ColumnFamily `json:"create,omitempty"`

	// Drop: Drop (delete) the column family with the given ID, or fail if
	// no such
	// family exists.
	Drop bool `json:"drop,omitempty"`

	// Id: The ID of the column family to be modified.
	Id string `json:"id,omitempty"`

	// Update: Update an existing column family to the specified schema, or
	// fail
	// if no column family exists with the given ID.
	Update *ColumnFamily `json:"update,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Create") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Modification) MarshalJSON() ([]byte, error) {
	type noMethod Modification
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// ModifyColumnFamiliesRequest: Request message
// for
// google.bigtable.admin.v2.BigtableTableAdmin.ModifyColumnFamilies
type ModifyColumnFamiliesRequest struct {
	// Modifications: Modifications to be atomically applied to the
	// specified table's families.
	// Entries are applied in order, meaning that earlier modifications can
	// be
	// masked by later ones (in the case of repeated updates to the same
	// family,
	// for example).
	Modifications []*Modification `json:"modifications,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Modifications") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *ModifyColumnFamiliesRequest) MarshalJSON() ([]byte, error) {
	type noMethod ModifyColumnFamiliesRequest
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// Operation: This resource represents a long-running operation that is
// the result of a
// network API call.
type Operation struct {
	// Done: If the value is `false`, it means the operation is still in
	// progress.
	// If true, the operation is completed, and either `error` or `response`
	// is
	// available.
	Done bool `json:"done,omitempty"`

	// Error: The error result of the operation in case of failure or
	// cancellation.
	Error *Status `json:"error,omitempty"`

	// Metadata: Service-specific metadata associated with the operation.
	// It typically
	// contains progress information and common metadata such as create
	// time.
	// Some services might not provide such metadata.  Any method that
	// returns a
	// long-running operation should document the metadata type, if any.
	Metadata OperationMetadata `json:"metadata,omitempty"`

	// Name: The server-assigned name, which is only unique within the same
	// service that
	// originally returns it. If you use the default HTTP mapping,
	// the
	// `name` should have the format of `operations/some/unique/name`.
	Name string `json:"name,omitempty"`

	// Response: The normal response of the operation in case of success.
	// If the original
	// method returns no data on success, such as `Delete`, the response
	// is
	// `google.protobuf.Empty`.  If the original method is
	// standard
	// `Get`/`Create`/`Update`, the response should be the resource.  For
	// other
	// methods, the response should have the type `XxxResponse`, where
	// `Xxx`
	// is the original method name.  For example, if the original method
	// name
	// is `TakeSnapshot()`, the inferred response type
	// is
	// `TakeSnapshotResponse`.
	Response OperationResponse `json:"response,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Done") to
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

type OperationMetadata interface{}

type OperationResponse interface{}

// Split: An initial split point for a newly created table.
type Split struct {
	// Key: Row key to use as an initial tablet boundary.
	Key string `json:"key,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Key") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Split) MarshalJSON() ([]byte, error) {
	type noMethod Split
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// Status: The `Status` type defines a logical error model that is
// suitable for different
// programming environments, including REST APIs and RPC APIs. It is
// used by
// [gRPC](https://github.com/grpc). The error model is designed to
// be:
//
// - Simple to use and understand for most users
// - Flexible enough to meet unexpected needs
//
// # Overview
//
// The `Status` message contains three pieces of data: error code, error
// message,
// and error details. The error code should be an enum value
// of
// google.rpc.Code, but it may accept additional error codes if needed.
// The
// error message should be a developer-facing English message that
// helps
// developers *understand* and *resolve* the error. If a localized
// user-facing
// error message is needed, put the localized message in the error
// details or
// localize it in the client. The optional error details may contain
// arbitrary
// information about the error. There is a predefined set of error
// detail types
// in the package `google.rpc` which can be used for common error
// conditions.
//
// # Language mapping
//
// The `Status` message is the logical representation of the error
// model, but it
// is not necessarily the actual wire format. When the `Status` message
// is
// exposed in different client libraries and different wire protocols,
// it can be
// mapped differently. For example, it will likely be mapped to some
// exceptions
// in Java, but more likely mapped to some error codes in C.
//
// # Other uses
//
// The error model and the `Status` message can be used in a variety
// of
// environments, either with or without APIs, to provide a
// consistent developer experience across different
// environments.
//
// Example uses of this error model include:
//
// - Partial errors. If a service needs to return partial errors to the
// client,
//     it may embed the `Status` in the normal response to indicate the
// partial
//     errors.
//
// - Workflow errors. A typical workflow has multiple steps. Each step
// may
//     have a `Status` message for error reporting purpose.
//
// - Batch operations. If a client uses batch request and batch
// response, the
//     `Status` message should be used directly inside batch response,
// one for
//     each error sub-response.
//
// - Asynchronous operations. If an API call embeds asynchronous
// operation
//     results in its response, the status of those operations should
// be
//     represented directly using the `Status` message.
//
// - Logging. If some API errors are stored in logs, the message
// `Status` could
//     be used directly after any stripping needed for security/privacy
// reasons.
type Status struct {
	// Code: The status code, which should be an enum value of
	// google.rpc.Code.
	Code int64 `json:"code,omitempty"`

	// Details: A list of messages that carry the error details.  There will
	// be a
	// common set of message types for APIs to use.
	Details []StatusDetails `json:"details,omitempty"`

	// Message: A developer-facing error message, which should be in
	// English. Any
	// user-facing error message should be localized and sent in
	// the
	// google.rpc.Status.details field, or localized by the client.
	Message string `json:"message,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Code") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Status) MarshalJSON() ([]byte, error) {
	type noMethod Status
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

type StatusDetails interface{}

// Table: A collection of user data indexed by row, column, and
// timestamp.
// Each table is served using the resources of its parent cluster.
type Table struct {
	// ColumnFamilies: (`CreationOnly`)
	// The column families configured for this table, mapped by column
	// family ID.
	// Views: `SCHEMA_VIEW`, `FULL`
	ColumnFamilies map[string]ColumnFamily `json:"columnFamilies,omitempty"`

	// Granularity: (`CreationOnly`)
	// The granularity (e.g. `MILLIS`, `MICROS`) at which timestamps are
	// stored in
	// this table. Timestamps not matching the granularity will be
	// rejected.
	// If unspecified at creation time, the value will be set to
	// `MILLIS`.
	// Views: `SCHEMA_VIEW`, `FULL`
	//
	// Possible values:
	//   "TIMESTAMP_GRANULARITY_UNSPECIFIED" - The user did not specify a
	// granularity. Should not be returned.
	// When specified during table creation, MILLIS will be used.
	//   "MILLIS" - The table keeps data versioned at a granularity of 1ms.
	Granularity string `json:"granularity,omitempty"`

	// Name: (`OutputOnly`)
	// The unique name of the table. Values are of the
	// form
	// `projects/<project>/instances/<instance>/tables/_a-zA-Z0-9*`.
	// Vie
	// ws: `NAME_ONLY`, `SCHEMA_VIEW`, `FULL`
	Name string `json:"name,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "ColumnFamilies") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Table) MarshalJSON() ([]byte, error) {
	type noMethod Table
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// Union: A GcRule which deletes cells matching any of the given rules.
type Union struct {
	// Rules: Delete cells which would be deleted by any element of `rules`.
	Rules []*GcRule `json:"rules,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Rules") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Union) MarshalJSON() ([]byte, error) {
	type noMethod Union
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// UpdateClusterMetadata: The metadata for the Operation returned by
// UpdateCluster.
type UpdateClusterMetadata struct {
	// FinishTime: The time at which the operation failed or was completed
	// successfully.
	FinishTime string `json:"finishTime,omitempty"`

	// OriginalRequest: The request that prompted the initiation of this
	// UpdateCluster operation.
	OriginalRequest *Cluster `json:"originalRequest,omitempty"`

	// RequestTime: The time at which the original request was received.
	RequestTime string `json:"requestTime,omitempty"`

	// ForceSendFields is a list of field names (e.g. "FinishTime") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *UpdateClusterMetadata) MarshalJSON() ([]byte, error) {
	type noMethod UpdateClusterMetadata
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// method id "bigtableadmin.operations.cancel":

type OperationsCancelCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
}

// Cancel: Starts asynchronous cancellation on a long-running operation.
//  The server
// makes a best effort to cancel the operation, but success is
// not
// guaranteed.  If the server doesn't support this method, it
// returns
// `google.rpc.Code.UNIMPLEMENTED`.  Clients can
// use
// Operations.GetOperation or
// other methods to check whether the cancellation succeeded or whether
// the
// operation completed despite cancellation. On successful
// cancellation,
// the operation is not deleted; instead, it becomes an operation
// with
// an Operation.error value with a google.rpc.Status.code of
// 1,
// corresponding to `Code.CANCELLED`.
func (r *OperationsService) Cancel(name string) *OperationsCancelCall {
	c := &OperationsCancelCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OperationsCancelCall) Fields(s ...googleapi.Field) *OperationsCancelCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *OperationsCancelCall) Context(ctx context.Context) *OperationsCancelCall {
	c.ctx_ = ctx
	return c
}

func (c *OperationsCancelCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/{+name}:cancel")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "bigtableadmin.operations.cancel" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *OperationsCancelCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
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
	ret := &Empty{
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
	//   "description": "Starts asynchronous cancellation on a long-running operation.  The server\nmakes a best effort to cancel the operation, but success is not\nguaranteed.  If the server doesn't support this method, it returns\n`google.rpc.Code.UNIMPLEMENTED`.  Clients can use\nOperations.GetOperation or\nother methods to check whether the cancellation succeeded or whether the\noperation completed despite cancellation. On successful cancellation,\nthe operation is not deleted; instead, it becomes an operation with\nan Operation.error value with a google.rpc.Status.code of 1,\ncorresponding to `Code.CANCELLED`.",
	//   "flatPath": "v2/operations/{operationsId}:cancel",
	//   "httpMethod": "POST",
	//   "id": "bigtableadmin.operations.cancel",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The name of the operation resource to be cancelled.",
	//       "location": "path",
	//       "pattern": "^operations/.+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2/{+name}:cancel",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "bigtableadmin.operations.delete":

type OperationsDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
}

// Delete: Deletes a long-running operation. This method indicates that
// the client is
// no longer interested in the operation result. It does not cancel
// the
// operation. If the server doesn't support this method, it
// returns
// `google.rpc.Code.UNIMPLEMENTED`.
func (r *OperationsService) Delete(name string) *OperationsDeleteCall {
	c := &OperationsDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OperationsDeleteCall) Fields(s ...googleapi.Field) *OperationsDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *OperationsDeleteCall) Context(ctx context.Context) *OperationsDeleteCall {
	c.ctx_ = ctx
	return c
}

func (c *OperationsDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "bigtableadmin.operations.delete" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *OperationsDeleteCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
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
	ret := &Empty{
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
	//   "description": "Deletes a long-running operation. This method indicates that the client is\nno longer interested in the operation result. It does not cancel the\noperation. If the server doesn't support this method, it returns\n`google.rpc.Code.UNIMPLEMENTED`.",
	//   "flatPath": "v2/operations/{operationsId}",
	//   "httpMethod": "DELETE",
	//   "id": "bigtableadmin.operations.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The name of the operation resource to be deleted.",
	//       "location": "path",
	//       "pattern": "^operations/.+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2/{+name}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "bigtableadmin.operations.get":

type OperationsGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// Get: Gets the latest state of a long-running operation.  Clients can
// use this
// method to poll the operation result at intervals as recommended by
// the API
// service.
func (r *OperationsService) Get(name string) *OperationsGetCall {
	c := &OperationsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OperationsGetCall) Fields(s ...googleapi.Field) *OperationsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *OperationsGetCall) IfNoneMatch(entityTag string) *OperationsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *OperationsGetCall) Context(ctx context.Context) *OperationsGetCall {
	c.ctx_ = ctx
	return c
}

func (c *OperationsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "bigtableadmin.operations.get" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *OperationsGetCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
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
	//   "description": "Gets the latest state of a long-running operation.  Clients can use this\nmethod to poll the operation result at intervals as recommended by the API\nservice.",
	//   "flatPath": "v2/operations/{operationsId}",
	//   "httpMethod": "GET",
	//   "id": "bigtableadmin.operations.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The name of the operation resource.",
	//       "location": "path",
	//       "pattern": "^operations/.+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2/{+name}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only"
	//   ]
	// }

}

// method id "bigtableadmin.operations.list":

type OperationsListCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// List: Lists operations that match the specified filter in the
// request. If the
// server doesn't support this method, it returns
// `UNIMPLEMENTED`.
//
// NOTE: the `name` binding below allows API services to override the
// binding
// to use different resource name schemes, such as `users/*/operations`.
func (r *OperationsService) List(name string) *OperationsListCall {
	c := &OperationsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Filter sets the optional parameter "filter": The standard list
// filter.
func (c *OperationsListCall) Filter(filter string) *OperationsListCall {
	c.urlParams_.Set("filter", filter)
	return c
}

// PageSize sets the optional parameter "pageSize": The standard list
// page size.
func (c *OperationsListCall) PageSize(pageSize int64) *OperationsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The standard list
// page token.
func (c *OperationsListCall) PageToken(pageToken string) *OperationsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OperationsListCall) Fields(s ...googleapi.Field) *OperationsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *OperationsListCall) IfNoneMatch(entityTag string) *OperationsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *OperationsListCall) Context(ctx context.Context) *OperationsListCall {
	c.ctx_ = ctx
	return c
}

func (c *OperationsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "bigtableadmin.operations.list" call.
// Exactly one of *ListOperationsResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListOperationsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *OperationsListCall) Do(opts ...googleapi.CallOption) (*ListOperationsResponse, error) {
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
	//   "description": "Lists operations that match the specified filter in the request. If the\nserver doesn't support this method, it returns `UNIMPLEMENTED`.\n\nNOTE: the `name` binding below allows API services to override the binding\nto use different resource name schemes, such as `users/*/operations`.",
	//   "flatPath": "v2/operations",
	//   "httpMethod": "GET",
	//   "id": "bigtableadmin.operations.list",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "The standard list filter.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "name": {
	//       "description": "The name of the operation collection.",
	//       "location": "path",
	//       "pattern": "^operations$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "The standard list page size.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The standard list page token.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2/{+name}",
	//   "response": {
	//     "$ref": "ListOperationsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *OperationsListCall) Pages(ctx context.Context, f func(*ListOperationsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "bigtableadmin.projects.instances.create":

type ProjectsInstancesCreateCall struct {
	s                     *Service
	parent                string
	createinstancerequest *CreateInstanceRequest
	urlParams_            gensupport.URLParams
	ctx_                  context.Context
}

// Create: Create an instance within a project.
func (r *ProjectsInstancesService) Create(parent string, createinstancerequest *CreateInstanceRequest) *ProjectsInstancesCreateCall {
	c := &ProjectsInstancesCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.createinstancerequest = createinstancerequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsInstancesCreateCall) Fields(s ...googleapi.Field) *ProjectsInstancesCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsInstancesCreateCall) Context(ctx context.Context) *ProjectsInstancesCreateCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsInstancesCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.createinstancerequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/{+parent}/instances")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "bigtableadmin.projects.instances.create" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsInstancesCreateCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
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
	//   "description": "Create an instance within a project.",
	//   "flatPath": "v2/projects/{projectsId}/instances",
	//   "httpMethod": "POST",
	//   "id": "bigtableadmin.projects.instances.create",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "The unique name of the project in which to create the new instance.\nValues are of the form `projects/\u003cproject\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2/{+parent}/instances",
	//   "request": {
	//     "$ref": "CreateInstanceRequest"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "bigtableadmin.projects.instances.delete":

type ProjectsInstancesDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
}

// Delete: Delete an instance from a project.
func (r *ProjectsInstancesService) Delete(name string) *ProjectsInstancesDeleteCall {
	c := &ProjectsInstancesDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsInstancesDeleteCall) Fields(s ...googleapi.Field) *ProjectsInstancesDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsInstancesDeleteCall) Context(ctx context.Context) *ProjectsInstancesDeleteCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsInstancesDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "bigtableadmin.projects.instances.delete" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsInstancesDeleteCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
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
	ret := &Empty{
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
	//   "description": "Delete an instance from a project.",
	//   "flatPath": "v2/projects/{projectsId}/instances/{instancesId}",
	//   "httpMethod": "DELETE",
	//   "id": "bigtableadmin.projects.instances.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The unique name of the instance to be deleted.\nValues are of the form `projects/\u003cproject\u003e/instances/\u003cinstance\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/instances/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2/{+name}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "bigtableadmin.projects.instances.get":

type ProjectsInstancesGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// Get: Gets information about an instance.
func (r *ProjectsInstancesService) Get(name string) *ProjectsInstancesGetCall {
	c := &ProjectsInstancesGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsInstancesGetCall) Fields(s ...googleapi.Field) *ProjectsInstancesGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsInstancesGetCall) IfNoneMatch(entityTag string) *ProjectsInstancesGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsInstancesGetCall) Context(ctx context.Context) *ProjectsInstancesGetCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsInstancesGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "bigtableadmin.projects.instances.get" call.
// Exactly one of *Instance or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Instance.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsInstancesGetCall) Do(opts ...googleapi.CallOption) (*Instance, error) {
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
	ret := &Instance{
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
	//   "description": "Gets information about an instance.",
	//   "flatPath": "v2/projects/{projectsId}/instances/{instancesId}",
	//   "httpMethod": "GET",
	//   "id": "bigtableadmin.projects.instances.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The unique name of the requested instance. Values are of the form\n`projects/\u003cproject\u003e/instances/\u003cinstance\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/instances/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2/{+name}",
	//   "response": {
	//     "$ref": "Instance"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only"
	//   ]
	// }

}

// method id "bigtableadmin.projects.instances.list":

type ProjectsInstancesListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// List: Lists information about instances in a project.
func (r *ProjectsInstancesService) List(parent string) *ProjectsInstancesListCall {
	c := &ProjectsInstancesListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// PageToken sets the optional parameter "pageToken": The value of
// `next_page_token` returned by a previous call.
func (c *ProjectsInstancesListCall) PageToken(pageToken string) *ProjectsInstancesListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsInstancesListCall) Fields(s ...googleapi.Field) *ProjectsInstancesListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsInstancesListCall) IfNoneMatch(entityTag string) *ProjectsInstancesListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsInstancesListCall) Context(ctx context.Context) *ProjectsInstancesListCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsInstancesListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/{+parent}/instances")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "bigtableadmin.projects.instances.list" call.
// Exactly one of *ListInstancesResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListInstancesResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsInstancesListCall) Do(opts ...googleapi.CallOption) (*ListInstancesResponse, error) {
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
	ret := &ListInstancesResponse{
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
	//   "description": "Lists information about instances in a project.",
	//   "flatPath": "v2/projects/{projectsId}/instances",
	//   "httpMethod": "GET",
	//   "id": "bigtableadmin.projects.instances.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "pageToken": {
	//       "description": "The value of `next_page_token` returned by a previous call.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "The unique name of the project for which a list of instances is requested.\nValues are of the form `projects/\u003cproject\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2/{+parent}/instances",
	//   "response": {
	//     "$ref": "ListInstancesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsInstancesListCall) Pages(ctx context.Context, f func(*ListInstancesResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "bigtableadmin.projects.instances.update":

type ProjectsInstancesUpdateCall struct {
	s          *Service
	name       string
	instance   *Instance
	urlParams_ gensupport.URLParams
	ctx_       context.Context
}

// Update: Updates an instance within a project.
func (r *ProjectsInstancesService) Update(name string, instance *Instance) *ProjectsInstancesUpdateCall {
	c := &ProjectsInstancesUpdateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.instance = instance
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsInstancesUpdateCall) Fields(s ...googleapi.Field) *ProjectsInstancesUpdateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsInstancesUpdateCall) Context(ctx context.Context) *ProjectsInstancesUpdateCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsInstancesUpdateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.instance)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "bigtableadmin.projects.instances.update" call.
// Exactly one of *Instance or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Instance.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsInstancesUpdateCall) Do(opts ...googleapi.CallOption) (*Instance, error) {
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
	ret := &Instance{
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
	//   "description": "Updates an instance within a project.",
	//   "flatPath": "v2/projects/{projectsId}/instances/{instancesId}",
	//   "httpMethod": "PUT",
	//   "id": "bigtableadmin.projects.instances.update",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "(`OutputOnly`)\nThe unique name of the instance. Values are of the form\n`projects/\u003cproject\u003e/instances/a-z+[a-z0-9]`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/instances/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2/{+name}",
	//   "request": {
	//     "$ref": "Instance"
	//   },
	//   "response": {
	//     "$ref": "Instance"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "bigtableadmin.projects.instances.clusters.create":

type ProjectsInstancesClustersCreateCall struct {
	s          *Service
	parent     string
	cluster    *Cluster
	urlParams_ gensupport.URLParams
	ctx_       context.Context
}

// Create: Creates a cluster within an instance.
func (r *ProjectsInstancesClustersService) Create(parent string, cluster *Cluster) *ProjectsInstancesClustersCreateCall {
	c := &ProjectsInstancesClustersCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.cluster = cluster
	return c
}

// ClusterId sets the optional parameter "clusterId": The ID to be used
// when referring to the new cluster within its instance,
// e.g., just `mycluster` rather
// than
// `projects/myproject/instances/myinstance/clusters/mycluster`.
func (c *ProjectsInstancesClustersCreateCall) ClusterId(clusterId string) *ProjectsInstancesClustersCreateCall {
	c.urlParams_.Set("clusterId", clusterId)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsInstancesClustersCreateCall) Fields(s ...googleapi.Field) *ProjectsInstancesClustersCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsInstancesClustersCreateCall) Context(ctx context.Context) *ProjectsInstancesClustersCreateCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsInstancesClustersCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.cluster)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/{+parent}/clusters")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "bigtableadmin.projects.instances.clusters.create" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsInstancesClustersCreateCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
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
	//   "description": "Creates a cluster within an instance.",
	//   "flatPath": "v2/projects/{projectsId}/instances/{instancesId}/clusters",
	//   "httpMethod": "POST",
	//   "id": "bigtableadmin.projects.instances.clusters.create",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "clusterId": {
	//       "description": "The ID to be used when referring to the new cluster within its instance,\ne.g., just `mycluster` rather than\n`projects/myproject/instances/myinstance/clusters/mycluster`.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "The unique name of the instance in which to create the new cluster.\nValues are of the form\n`projects/\u003cproject\u003e/instances/\u003cinstance\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/instances/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2/{+parent}/clusters",
	//   "request": {
	//     "$ref": "Cluster"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "bigtableadmin.projects.instances.clusters.delete":

type ProjectsInstancesClustersDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
}

// Delete: Deletes a cluster from an instance.
func (r *ProjectsInstancesClustersService) Delete(name string) *ProjectsInstancesClustersDeleteCall {
	c := &ProjectsInstancesClustersDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsInstancesClustersDeleteCall) Fields(s ...googleapi.Field) *ProjectsInstancesClustersDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsInstancesClustersDeleteCall) Context(ctx context.Context) *ProjectsInstancesClustersDeleteCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsInstancesClustersDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "bigtableadmin.projects.instances.clusters.delete" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsInstancesClustersDeleteCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
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
	ret := &Empty{
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
	//   "description": "Deletes a cluster from an instance.",
	//   "flatPath": "v2/projects/{projectsId}/instances/{instancesId}/clusters/{clustersId}",
	//   "httpMethod": "DELETE",
	//   "id": "bigtableadmin.projects.instances.clusters.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The unique name of the cluster to be deleted. Values are of the form\n`projects/\u003cproject\u003e/instances/\u003cinstance\u003e/clusters/\u003ccluster\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/instances/[^/]+/clusters/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2/{+name}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "bigtableadmin.projects.instances.clusters.get":

type ProjectsInstancesClustersGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// Get: Gets information about a cluster.
func (r *ProjectsInstancesClustersService) Get(name string) *ProjectsInstancesClustersGetCall {
	c := &ProjectsInstancesClustersGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsInstancesClustersGetCall) Fields(s ...googleapi.Field) *ProjectsInstancesClustersGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsInstancesClustersGetCall) IfNoneMatch(entityTag string) *ProjectsInstancesClustersGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsInstancesClustersGetCall) Context(ctx context.Context) *ProjectsInstancesClustersGetCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsInstancesClustersGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "bigtableadmin.projects.instances.clusters.get" call.
// Exactly one of *Cluster or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Cluster.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsInstancesClustersGetCall) Do(opts ...googleapi.CallOption) (*Cluster, error) {
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
	//   "description": "Gets information about a cluster.",
	//   "flatPath": "v2/projects/{projectsId}/instances/{instancesId}/clusters/{clustersId}",
	//   "httpMethod": "GET",
	//   "id": "bigtableadmin.projects.instances.clusters.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The unique name of the requested cluster. Values are of the form\n`projects/\u003cproject\u003e/instances/\u003cinstance\u003e/clusters/\u003ccluster\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/instances/[^/]+/clusters/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2/{+name}",
	//   "response": {
	//     "$ref": "Cluster"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only"
	//   ]
	// }

}

// method id "bigtableadmin.projects.instances.clusters.list":

type ProjectsInstancesClustersListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// List: Lists information about clusters in an instance.
func (r *ProjectsInstancesClustersService) List(parent string) *ProjectsInstancesClustersListCall {
	c := &ProjectsInstancesClustersListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// PageToken sets the optional parameter "pageToken": The value of
// `next_page_token` returned by a previous call.
func (c *ProjectsInstancesClustersListCall) PageToken(pageToken string) *ProjectsInstancesClustersListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsInstancesClustersListCall) Fields(s ...googleapi.Field) *ProjectsInstancesClustersListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsInstancesClustersListCall) IfNoneMatch(entityTag string) *ProjectsInstancesClustersListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsInstancesClustersListCall) Context(ctx context.Context) *ProjectsInstancesClustersListCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsInstancesClustersListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/{+parent}/clusters")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "bigtableadmin.projects.instances.clusters.list" call.
// Exactly one of *ListClustersResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListClustersResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsInstancesClustersListCall) Do(opts ...googleapi.CallOption) (*ListClustersResponse, error) {
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
	//   "description": "Lists information about clusters in an instance.",
	//   "flatPath": "v2/projects/{projectsId}/instances/{instancesId}/clusters",
	//   "httpMethod": "GET",
	//   "id": "bigtableadmin.projects.instances.clusters.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "pageToken": {
	//       "description": "The value of `next_page_token` returned by a previous call.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "The unique name of the instance for which a list of clusters is requested.\nValues are of the form `projects/\u003cproject\u003e/instances/\u003cinstance\u003e`.\nUse `\u003cinstance\u003e = '-'` to list Clusters for all Instances in a project,\ne.g., `projects/myproject/instances/-`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/instances/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2/{+parent}/clusters",
	//   "response": {
	//     "$ref": "ListClustersResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsInstancesClustersListCall) Pages(ctx context.Context, f func(*ListClustersResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "bigtableadmin.projects.instances.clusters.update":

type ProjectsInstancesClustersUpdateCall struct {
	s          *Service
	name       string
	cluster    *Cluster
	urlParams_ gensupport.URLParams
	ctx_       context.Context
}

// Update: Updates a cluster within an instance.
func (r *ProjectsInstancesClustersService) Update(name string, cluster *Cluster) *ProjectsInstancesClustersUpdateCall {
	c := &ProjectsInstancesClustersUpdateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.cluster = cluster
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsInstancesClustersUpdateCall) Fields(s ...googleapi.Field) *ProjectsInstancesClustersUpdateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsInstancesClustersUpdateCall) Context(ctx context.Context) *ProjectsInstancesClustersUpdateCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsInstancesClustersUpdateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.cluster)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "bigtableadmin.projects.instances.clusters.update" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsInstancesClustersUpdateCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
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
	//   "description": "Updates a cluster within an instance.",
	//   "flatPath": "v2/projects/{projectsId}/instances/{instancesId}/clusters/{clustersId}",
	//   "httpMethod": "PUT",
	//   "id": "bigtableadmin.projects.instances.clusters.update",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "(`OutputOnly`)\nThe unique name of the cluster. Values are of the form\n`projects/\u003cproject\u003e/instances/\u003cinstance\u003e/clusters/a-z*`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/instances/[^/]+/clusters/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2/{+name}",
	//   "request": {
	//     "$ref": "Cluster"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "bigtableadmin.projects.instances.tables.create":

type ProjectsInstancesTablesCreateCall struct {
	s                  *Service
	parent             string
	createtablerequest *CreateTableRequest
	urlParams_         gensupport.URLParams
	ctx_               context.Context
}

// Create: Creates a new table in the specified instance.
// The table can be created with a full set of initial column
// families,
// specified in the request.
func (r *ProjectsInstancesTablesService) Create(parent string, createtablerequest *CreateTableRequest) *ProjectsInstancesTablesCreateCall {
	c := &ProjectsInstancesTablesCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.createtablerequest = createtablerequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsInstancesTablesCreateCall) Fields(s ...googleapi.Field) *ProjectsInstancesTablesCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsInstancesTablesCreateCall) Context(ctx context.Context) *ProjectsInstancesTablesCreateCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsInstancesTablesCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.createtablerequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/{+parent}/tables")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "bigtableadmin.projects.instances.tables.create" call.
// Exactly one of *Table or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Table.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsInstancesTablesCreateCall) Do(opts ...googleapi.CallOption) (*Table, error) {
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
	ret := &Table{
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
	//   "description": "Creates a new table in the specified instance.\nThe table can be created with a full set of initial column families,\nspecified in the request.",
	//   "flatPath": "v2/projects/{projectsId}/instances/{instancesId}/tables",
	//   "httpMethod": "POST",
	//   "id": "bigtableadmin.projects.instances.tables.create",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "parent": {
	//       "description": "The unique name of the instance in which to create the table.\nValues are of the form `projects/\u003cproject\u003e/instances/\u003cinstance\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/instances/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2/{+parent}/tables",
	//   "request": {
	//     "$ref": "CreateTableRequest"
	//   },
	//   "response": {
	//     "$ref": "Table"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "bigtableadmin.projects.instances.tables.delete":

type ProjectsInstancesTablesDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
}

// Delete: Permanently deletes a specified table and all of its data.
func (r *ProjectsInstancesTablesService) Delete(name string) *ProjectsInstancesTablesDeleteCall {
	c := &ProjectsInstancesTablesDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsInstancesTablesDeleteCall) Fields(s ...googleapi.Field) *ProjectsInstancesTablesDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsInstancesTablesDeleteCall) Context(ctx context.Context) *ProjectsInstancesTablesDeleteCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsInstancesTablesDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "bigtableadmin.projects.instances.tables.delete" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsInstancesTablesDeleteCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
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
	ret := &Empty{
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
	//   "description": "Permanently deletes a specified table and all of its data.",
	//   "flatPath": "v2/projects/{projectsId}/instances/{instancesId}/tables/{tablesId}",
	//   "httpMethod": "DELETE",
	//   "id": "bigtableadmin.projects.instances.tables.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The unique name of the table to be deleted.\nValues are of the form\n`projects/\u003cproject\u003e/instances/\u003cinstance\u003e/tables/\u003ctable\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/instances/[^/]+/tables/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2/{+name}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "bigtableadmin.projects.instances.tables.dropRowRange":

type ProjectsInstancesTablesDropRowRangeCall struct {
	s                   *Service
	name                string
	droprowrangerequest *DropRowRangeRequest
	urlParams_          gensupport.URLParams
	ctx_                context.Context
}

// DropRowRange: Permanently drop/delete a row range from a specified
// table. The request can
// specify whether to delete all rows in a table, or only those that
// match a
// particular prefix.
func (r *ProjectsInstancesTablesService) DropRowRange(name string, droprowrangerequest *DropRowRangeRequest) *ProjectsInstancesTablesDropRowRangeCall {
	c := &ProjectsInstancesTablesDropRowRangeCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.droprowrangerequest = droprowrangerequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsInstancesTablesDropRowRangeCall) Fields(s ...googleapi.Field) *ProjectsInstancesTablesDropRowRangeCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsInstancesTablesDropRowRangeCall) Context(ctx context.Context) *ProjectsInstancesTablesDropRowRangeCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsInstancesTablesDropRowRangeCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.droprowrangerequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/{+name}:dropRowRange")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "bigtableadmin.projects.instances.tables.dropRowRange" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsInstancesTablesDropRowRangeCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
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
	ret := &Empty{
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
	//   "description": "Permanently drop/delete a row range from a specified table. The request can\nspecify whether to delete all rows in a table, or only those that match a\nparticular prefix.",
	//   "flatPath": "v2/projects/{projectsId}/instances/{instancesId}/tables/{tablesId}:dropRowRange",
	//   "httpMethod": "POST",
	//   "id": "bigtableadmin.projects.instances.tables.dropRowRange",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The unique name of the table on which to drop a range of rows.\nValues are of the form\n`projects/\u003cproject\u003e/instances/\u003cinstance\u003e/tables/\u003ctable\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/instances/[^/]+/tables/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2/{+name}:dropRowRange",
	//   "request": {
	//     "$ref": "DropRowRangeRequest"
	//   },
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "bigtableadmin.projects.instances.tables.get":

type ProjectsInstancesTablesGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// Get: Gets metadata information about the specified table.
func (r *ProjectsInstancesTablesService) Get(name string) *ProjectsInstancesTablesGetCall {
	c := &ProjectsInstancesTablesGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// View sets the optional parameter "view": The view to be applied to
// the returned table's fields.
// Defaults to `SCHEMA_ONLY` if unspecified.
//
// Possible values:
//   "VIEW_UNSPECIFIED"
//   "NAME_ONLY"
//   "SCHEMA_VIEW"
//   "FULL"
func (c *ProjectsInstancesTablesGetCall) View(view string) *ProjectsInstancesTablesGetCall {
	c.urlParams_.Set("view", view)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsInstancesTablesGetCall) Fields(s ...googleapi.Field) *ProjectsInstancesTablesGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsInstancesTablesGetCall) IfNoneMatch(entityTag string) *ProjectsInstancesTablesGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsInstancesTablesGetCall) Context(ctx context.Context) *ProjectsInstancesTablesGetCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsInstancesTablesGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "bigtableadmin.projects.instances.tables.get" call.
// Exactly one of *Table or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Table.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsInstancesTablesGetCall) Do(opts ...googleapi.CallOption) (*Table, error) {
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
	ret := &Table{
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
	//   "description": "Gets metadata information about the specified table.",
	//   "flatPath": "v2/projects/{projectsId}/instances/{instancesId}/tables/{tablesId}",
	//   "httpMethod": "GET",
	//   "id": "bigtableadmin.projects.instances.tables.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The unique name of the requested table.\nValues are of the form\n`projects/\u003cproject\u003e/instances/\u003cinstance\u003e/tables/\u003ctable\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/instances/[^/]+/tables/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "view": {
	//       "description": "The view to be applied to the returned table's fields.\nDefaults to `SCHEMA_ONLY` if unspecified.",
	//       "enum": [
	//         "VIEW_UNSPECIFIED",
	//         "NAME_ONLY",
	//         "SCHEMA_VIEW",
	//         "FULL"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2/{+name}",
	//   "response": {
	//     "$ref": "Table"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only"
	//   ]
	// }

}

// method id "bigtableadmin.projects.instances.tables.list":

type ProjectsInstancesTablesListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// List: Lists all tables served from a specified instance.
func (r *ProjectsInstancesTablesService) List(parent string) *ProjectsInstancesTablesListCall {
	c := &ProjectsInstancesTablesListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// PageToken sets the optional parameter "pageToken": The value of
// `next_page_token` returned by a previous call.
func (c *ProjectsInstancesTablesListCall) PageToken(pageToken string) *ProjectsInstancesTablesListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// View sets the optional parameter "view": The view to be applied to
// the returned tables' fields.
// Defaults to `NAME_ONLY` if unspecified; no others are currently
// supported.
//
// Possible values:
//   "VIEW_UNSPECIFIED"
//   "NAME_ONLY"
//   "SCHEMA_VIEW"
//   "FULL"
func (c *ProjectsInstancesTablesListCall) View(view string) *ProjectsInstancesTablesListCall {
	c.urlParams_.Set("view", view)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsInstancesTablesListCall) Fields(s ...googleapi.Field) *ProjectsInstancesTablesListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsInstancesTablesListCall) IfNoneMatch(entityTag string) *ProjectsInstancesTablesListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsInstancesTablesListCall) Context(ctx context.Context) *ProjectsInstancesTablesListCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsInstancesTablesListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/{+parent}/tables")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "bigtableadmin.projects.instances.tables.list" call.
// Exactly one of *ListTablesResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListTablesResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsInstancesTablesListCall) Do(opts ...googleapi.CallOption) (*ListTablesResponse, error) {
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
	ret := &ListTablesResponse{
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
	//   "description": "Lists all tables served from a specified instance.",
	//   "flatPath": "v2/projects/{projectsId}/instances/{instancesId}/tables",
	//   "httpMethod": "GET",
	//   "id": "bigtableadmin.projects.instances.tables.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "pageToken": {
	//       "description": "The value of `next_page_token` returned by a previous call.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "The unique name of the instance for which tables should be listed.\nValues are of the form `projects/\u003cproject\u003e/instances/\u003cinstance\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/instances/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "view": {
	//       "description": "The view to be applied to the returned tables' fields.\nDefaults to `NAME_ONLY` if unspecified; no others are currently supported.",
	//       "enum": [
	//         "VIEW_UNSPECIFIED",
	//         "NAME_ONLY",
	//         "SCHEMA_VIEW",
	//         "FULL"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2/{+parent}/tables",
	//   "response": {
	//     "$ref": "ListTablesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsInstancesTablesListCall) Pages(ctx context.Context, f func(*ListTablesResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "bigtableadmin.projects.instances.tables.modifyColumnFamilies":

type ProjectsInstancesTablesModifyColumnFamiliesCall struct {
	s                           *Service
	name                        string
	modifycolumnfamiliesrequest *ModifyColumnFamiliesRequest
	urlParams_                  gensupport.URLParams
	ctx_                        context.Context
}

// ModifyColumnFamilies: Atomically performs a series of column family
// modifications
// on the specified table.
func (r *ProjectsInstancesTablesService) ModifyColumnFamilies(name string, modifycolumnfamiliesrequest *ModifyColumnFamiliesRequest) *ProjectsInstancesTablesModifyColumnFamiliesCall {
	c := &ProjectsInstancesTablesModifyColumnFamiliesCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.modifycolumnfamiliesrequest = modifycolumnfamiliesrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsInstancesTablesModifyColumnFamiliesCall) Fields(s ...googleapi.Field) *ProjectsInstancesTablesModifyColumnFamiliesCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsInstancesTablesModifyColumnFamiliesCall) Context(ctx context.Context) *ProjectsInstancesTablesModifyColumnFamiliesCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsInstancesTablesModifyColumnFamiliesCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.modifycolumnfamiliesrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/{+name}:modifyColumnFamilies")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "bigtableadmin.projects.instances.tables.modifyColumnFamilies" call.
// Exactly one of *Table or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Table.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsInstancesTablesModifyColumnFamiliesCall) Do(opts ...googleapi.CallOption) (*Table, error) {
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
	ret := &Table{
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
	//   "description": "Atomically performs a series of column family modifications\non the specified table.",
	//   "flatPath": "v2/projects/{projectsId}/instances/{instancesId}/tables/{tablesId}:modifyColumnFamilies",
	//   "httpMethod": "POST",
	//   "id": "bigtableadmin.projects.instances.tables.modifyColumnFamilies",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The unique name of the table whose families should be modified.\nValues are of the form\n`projects/\u003cproject\u003e/instances/\u003cinstance\u003e/tables/\u003ctable\u003e`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/instances/[^/]+/tables/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2/{+name}:modifyColumnFamilies",
	//   "request": {
	//     "$ref": "ModifyColumnFamiliesRequest"
	//   },
	//   "response": {
	//     "$ref": "Table"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}
