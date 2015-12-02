// Package datastore provides access to the Google Cloud Datastore API - NEW.
//
// See https://cloud.google.com/datastore/
//
// Usage example:
//
//   import "google.golang.org/api/datastore/v1beta3"
//   ...
//   datastoreService, err := datastore.New(oauthHttpClient)
package datastore

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

const apiId = "datastore:v1beta3"
const apiName = "datastore"
const apiVersion = "v1beta3"
const basePath = "https://datastore.googleapis.com/"

// OAuth2 scopes used by this API.
const (
	// View and manage your data across Google Cloud Platform services
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"

	// View and manage your Google Cloud Datastore data
	DatastoreScope = "https://www.googleapis.com/auth/datastore"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Projects = NewProjectsService(s)
	return s, nil
}

type Service struct {
	client   *http.Client
	BasePath string // API endpoint base URL

	Projects *ProjectsService
}

func NewProjectsService(s *Service) *ProjectsService {
	rs := &ProjectsService{s: s}
	rs.Indexes = NewProjectsIndexesService(s)
	rs.Operations = NewProjectsOperationsService(s)
	return rs
}

type ProjectsService struct {
	s *Service

	Indexes *ProjectsIndexesService

	Operations *ProjectsOperationsService
}

func NewProjectsIndexesService(s *Service) *ProjectsIndexesService {
	rs := &ProjectsIndexesService{s: s}
	return rs
}

type ProjectsIndexesService struct {
	s *Service
}

func NewProjectsOperationsService(s *Service) *ProjectsOperationsService {
	rs := &ProjectsOperationsService{s: s}
	return rs
}

type ProjectsOperationsService struct {
	s *Service
}

type AllocateIdsRequest struct {
	// DatabaseId: If not empty, the ID of the database against which to
	// make the request.
	DatabaseId string `json:"databaseId,omitempty"`

	// Keys: A list of keys with incomplete key paths for which to allocate
	// IDs.
	// No key may be reserved/read-only.
	Keys []*Key `json:"keys,omitempty"`
}

type AllocateIdsResponse struct {
	// Keys: The keys specified in the request (in the same order), each
	// with
	// its key path completed with a newly allocated ID.
	Keys []*Key `json:"keys,omitempty"`
}

type ArrayValue struct {
	// Values: Values in the array.
	// The order of this array may not be
	// preserved if it contains a mix of
	// indexed and unindexed values.
	Values []*Value `json:"values,omitempty"`
}

type BeginTransactionRequest struct {
	// DatabaseId: If not empty, the ID of the database against which to
	// make the request.
	DatabaseId string `json:"databaseId,omitempty"`

	// TransactionOptions: Options for a new transaction.
	TransactionOptions *TransactionOptions `json:"transactionOptions,omitempty"`
}

type BeginTransactionResponse struct {
	// Transaction: The transaction identifier (always present).
	Transaction string `json:"transaction,omitempty"`
}

type Change struct {
	// Cause: The error that resulted in this change, if applicable.
	Cause *Status `json:"cause,omitempty"`

	// Continued: If true, more changes are needed to construct a consistent
	// snapshot.
	Continued bool `json:"continued,omitempty"`

	// ExistenceFilter: A filter representing delete mutations. Applies to
	// the *set*
	// of entities previously returned for the given `target_ids`.
	ExistenceFilter *ExistenceFilter `json:"existenceFilter,omitempty"`

	// Mutation: A mutation to a watched entity.
	Mutation *Mutation `json:"mutation,omitempty"`

	// NoChange: No change has occurred. Usually returned to provide an
	// updated
	// `resume_token`.
	NoChange string `json:"noChange,omitempty"`

	// ResumeToken: A token that provides a compact representation of all
	// the changes that
	// have been received by the caller up to this point
	// (including this one)
	// that can be used to resume the stream. May not
	// be set on every change.
	// Only valid for the targets specified in
	// `target_ids`.
	ResumeToken string `json:"resumeToken,omitempty"`

	// TargetChange: The targets associate with this stream have been
	// modified.
	// The affected targets are listed in `target_ids`.
	TargetChange string `json:"targetChange,omitempty"`

	// TargetIds: The set of targets to which this change applies. When
	// empty, the change
	// applies to all targets.
	TargetIds []int64 `json:"targetIds,omitempty"`
}

type ChangeBatch struct {
	Changes []*Change `json:"changes,omitempty"`
}

type Circle struct {
	// Center: The center of the circle.
	Center *LatLng `json:"center,omitempty"`

	// RadiusMeters: The "radius" of the circle, in meters.
	// Must be greater
	// or equal to zero.
	RadiusMeters float64 `json:"radiusMeters,omitempty"`
}

type CommitRequest struct {
	// DatabaseId: If not empty, the ID of the database against which to
	// make the request.
	DatabaseId string `json:"databaseId,omitempty"`

	// Mode: The type of commit to perform. Defaults to `TRANSACTIONAL`.
	Mode string `json:"mode,omitempty"`

	// Mutations: The mutations to perform.
	//
	// When mode is `TRANSACTIONAL`,
	// mutations affecting a single entity are
	// applied in order. The
	// following sequences of mutations affecting a single
	// entity are not
	// permitted in a single `Commit` request:
	//
	// - `insert` followed by
	// `insert`
	// - `update` followed by `insert`
	// - `upsert` followed by
	// `insert`
	// - `delete` followed by `update`
	//
	// When mode is
	// `NON_TRANSACTIONAL`, no two mutations may affect a single
	// entity.
	Mutations []*Mutation `json:"mutations,omitempty"`

	// SingleUseTransaction: Options for beginning a new transaction for
	// this request.
	// The transaction is committed when the request
	// completes. If
	// specified,
	// google.datastore.v1beta3.TransactionOptions.mode must
	// be
	// google.datastore.v1beta3.TransactionOptions.ReadWrite.
	SingleUseTransaction *TransactionOptions `json:"singleUseTransaction,omitempty"`

	// Transaction: The identifier of the transaction associated with the
	// commit. A
	// transaction identifier is returned by a call
	// to
	// BeginTransaction.
	Transaction string `json:"transaction,omitempty"`
}

type CommitResponse struct {
	// IndexUpdates: The number of index entries updated during the commit,
	// or zero if none were
	// updated.
	IndexUpdates int64 `json:"indexUpdates,omitempty"`

	// MutationResults: The result of performing the mutations.
	// The i-th
	// mutation result corresponds to the i-th mutation in the request.
	MutationResults []*MutationResult `json:"mutationResults,omitempty"`
}

type CommonMetadata struct {
	// Cancelling: True if the Operation is currently being cancelled.  Will
	// only be set after
	// the cancellation request was received, but before
	// the Operation is done.
	// So the state transitions are:
	// 1) done=false,
	// cancelling=false
	// 2) done=false, cancelling=true
	// 3) done=true,
	// cancelling=false
	Cancelling bool `json:"cancelling,omitempty"`

	// EndTime: The time the Operation ended, either successfully or
	// otherwise.
	EndTime string `json:"endTime,omitempty"`

	// Labels: The client-assigned labels which were provided when the
	// Operation was
	// created.  May also include additional labels.
	Labels map[string]string `json:"labels,omitempty"`

	// OperationType: The type of the operation.  Can be used as a filter
	// in
	// ListOperationsRequest.
	OperationType string `json:"operationType,omitempty"`

	// StartTime: The time that work began on the Operation.
	StartTime string `json:"startTime,omitempty"`
}

type CompositeFilter struct {
	// Filters: The list of filters to combine.
	// Must contain at least one
	// filter.
	Filters []*Filter `json:"filters,omitempty"`

	// Op: The operator for combining multiple filters.
	Op string `json:"op,omitempty"`
}

type Empty struct {
}

type Entity struct {
	// Key: The entity's key.
	//
	// An entity must have a key, unless otherwise
	// documented (for example,
	// an entity in `Value.entity_value` may have
	// no key).
	// An entity's kind is its key path's last element's kind,
	// or
	// null if it has no key.
	Key *Key `json:"key,omitempty"`

	// Properties: The entity's properties.
	// The map's keys are property
	// names.
	// A property name matching regex `__.*__` is reserved.
	// A
	// reserved property name is forbidden in certain documented
	// contexts.
	// The name must not contain more than 500 characters.
	// The
	// name cannot be `""`.
	Properties map[string]Value `json:"properties,omitempty"`
}

type EntityFilter struct {
	// Kinds: If empty, then this represents all Kinds.
	Kinds []string `json:"kinds,omitempty"`

	// NamespaceIds: An empty list represents all Namespaces.  This is the
	// preferred
	// usage for Projects that don't use Namespaces.
	//
	// An empty
	// string element represents the Default Namespace.  This should be
	// used
	// if the Project has data in non-Default Namespaces, but doesn't want
	// to
	// include them.
	// Each Namespace in this list must be unique.
	NamespaceIds []string `json:"namespaceIds,omitempty"`
}

type EntityResult struct {
	// Cursor: A cursor that points to the position after the result
	// entity.
	// Set only when the `EntityResult` is part of a
	// `QueryResultBatch` message.
	Cursor string `json:"cursor,omitempty"`

	// Entity: The resulting entity.
	Entity *Entity `json:"entity,omitempty"`

	// Version: The version of the entity, a strictly positive number that
	// monotonically
	// increases with changes to the entity.
	//
	// This field is
	// set for `FULL` entity results.
	// For missing entities
	// in
	// `LookupResponse`, this is the version of the snapshot that was
	// used to look
	// up the entity, and it is always set except for
	// eventually consistent reads.
	Version int64 `json:"version,omitempty,string"`
}

type ExistenceFilter struct {
	// Bits: The filter bits.
	Bits string `json:"bits,omitempty"`

	// Count: The total number of keys represented in this filter. Used to
	// detect
	// the presence of false positives.
	Count int64 `json:"count,omitempty"`

	// HashCount: The number of hashes used in `bits`.
	HashCount int64 `json:"hashCount,omitempty"`

	// Strategy: The strategy used to map a key to the filter bits.
	Strategy string `json:"strategy,omitempty"`
}

type ExportMetadata struct {
	// BytesProgress: An estimate of the number of bytes processed.
	BytesProgress *Progress `json:"bytesProgress,omitempty"`

	// Common: Metadata common to all Datastore Admin Operations.
	Common *CommonMetadata `json:"common,omitempty"`

	// Destination: Location for the export metadata and data files.  This
	// will be the same
	// value as the
	// google.datastore.v1beta3.ExportRequest.destination field.
	// Actual
	// files will be nested deeper than this.  The final output location
	// is
	// provided in google.datastore.v1beta3.ExportResponse.data_location.
	Destination string `json:"destination,omitempty"`

	// EntitiesProgress: An estimate of the number of Entities processed.
	EntitiesProgress *Progress `json:"entitiesProgress,omitempty"`

	// EntityFilter: Description of which Entities are being exported.
	EntityFilter *EntityFilter `json:"entityFilter,omitempty"`
}

type ExportRequest struct {
	// DatabaseId: Database ID against which to make the request.
	// Unset
	// indicates the default database.
	// Not currently supported.
	DatabaseId string `json:"databaseId,omitempty"`

	// Destination: Location for the export metadata and data
	// files.
	//
	// Specified as the full resource name of the external storage
	// location.
	// Currently, only Google Cloud Storage is supported.  So the
	// destination
	// should be of the form:
	// //storage.googleapis.com/buckets/bucket-name
	// or
	// //storage.googleapis.com/buckets/bucket-name/objects/object-path.
	//
	//
	// The resulting files will be nested deeper than the specified
	// destination.
	// The final data location will be provided in
	// the
	// google.datastore.v1beta3.ExportResponse.data_location field.
	// That
	// value should be used for subsequent Import Operations.
	//
	// By
	// nesting the data files deeper, the same destination can be used
	// in
	// multiple Export Operations without conflict.
	Destination string `json:"destination,omitempty"`

	// EntityFilter: Description of what data from the Project is included
	// in the export.
	//
	EntityFilter *EntityFilter `json:"entityFilter,omitempty"`

	// Labels: Client-assigned labels.
	Labels map[string]string `json:"labels,omitempty"`
}

type ExportResponse struct {
	// DataLocation: Location of the output files. This can be used to begin
	// an import into
	// Cloud Datastore (this Project or another Project).
	// See
	// google.datastore.v1beta3.ImportRequest.data_location.  Only
	// present if
	// the Operation completed successfully.
	//
	// This location will
	// be nested deeper than the
	// initial
	// google.datastore.v1beta3.ExportRequest.destination in order
	// to support
	// multiple Exports to the same destination without conflict.
	DataLocation string `json:"dataLocation,omitempty"`
}

type Filter struct {
	// CompositeFilter: A composite filter.
	CompositeFilter *CompositeFilter `json:"compositeFilter,omitempty"`

	// PropertyFilter: A filter on a property.
	PropertyFilter *PropertyFilter `json:"propertyFilter,omitempty"`

	// StContainsFilter: A filter that selects geo points within a region.
	StContainsFilter *StContainsFilter `json:"stContainsFilter,omitempty"`
}

type GeoRegion struct {
	// Circle: A circular region.
	Circle *Circle `json:"circle,omitempty"`

	// Rectangle: A rectangular region.
	Rectangle *Rectangle `json:"rectangle,omitempty"`
}

type GqlQuery struct {
	// AllowLiterals: When false, the query string must not contain any
	// literals and instead
	// must bind all values. For example,
	// `SELECT *
	// FROM Kind WHERE a = 'string literal'` is not allowed, while
	// `SELECT *
	// FROM Kind WHERE a = @value` is.
	AllowLiterals bool `json:"allowLiterals,omitempty"`

	// NamedBindings: For each non-reserved named binding site in the query
	// string,
	// there must be a named parameter with that name,
	// but not
	// necessarily the inverse.
	// Key must match regex `A-Za-z_$*`, must not
	// match regex
	// `__.*__`, and must not be `""`.
	NamedBindings map[string]GqlQueryParameter `json:"namedBindings,omitempty"`

	// PositionalBindings: Numbered binding site @1 references the first
	// numbered parameter,
	// effectively using 1-based indexing, rather than
	// the usual 0.
	// For each binding site numbered i in
	// `query_string`,
	// there must be an i-th numbered parameter.
	// The inverse
	// must also be true.
	PositionalBindings []*GqlQueryParameter `json:"positionalBindings,omitempty"`

	// QueryString: A string of the format
	// described
	// [here](https://cloud.google.com/datastore/docs/apis/gql/gql_
	// reference).
	QueryString string `json:"queryString,omitempty"`
}

type GqlQueryParameter struct {
	// Cursor: A query cursor. Query cursors are returned in query
	// result
	// batches.
	Cursor string `json:"cursor,omitempty"`

	// GeoRegion: Geographical region.
	GeoRegion *GeoRegion `json:"geoRegion,omitempty"`

	// Value: A value parameter.
	Value *Value `json:"value,omitempty"`
}

type ImportMetadata struct {
	// BytesProgress: An estimate of the number of bytes processed.
	BytesProgress *Progress `json:"bytesProgress,omitempty"`

	// Common: Metadata common to all Datastore Admin Operations.
	Common *CommonMetadata `json:"common,omitempty"`

	// DataLocation: The location that we are importing from.
	DataLocation string `json:"dataLocation,omitempty"`

	// EntitiesProgress: An estimate of the number of Entities processed.
	EntitiesProgress *Progress `json:"entitiesProgress,omitempty"`

	// EntityFilter: Description of which Entities are being imported.
	EntityFilter *EntityFilter `json:"entityFilter,omitempty"`
}

type ImportRequest struct {
	// DataLocation: The full resource name of the external storage
	// location.  Currently, only
	// Google Cloud Storage is supported.  So the
	// data_location should be of the
	// form:
	// //storage.googleapis.com/buckets/bucket-name/objects/object-path.
	//
	// See
	//  google.datastore.v1beta3.ExportResponse.data_location
	DataLocation string `json:"dataLocation,omitempty"`

	// DatabaseId: Database ID against which to make the request.
	// Unset
	// indicates the default database.
	// Not currently supported.
	DatabaseId string `json:"databaseId,omitempty"`

	// EntityFilter: Optionally specify which Kinds/Namespaces are to be
	// imported. If provided,
	// the list must be a subset of the EntityFilter
	// used in creating the export,
	// otherwise a FAILED_PRECONDITION error
	// will be returned. If no filter is
	// specified then all Entities from
	// the export are imported.
	EntityFilter *EntityFilter `json:"entityFilter,omitempty"`

	// Labels: Client-assigned labels.
	Labels map[string]string `json:"labels,omitempty"`
}

type Index struct {
	// DatabaseId: If not empty, the ID of the database to which the index
	// belongs.
	DatabaseId string `json:"databaseId,omitempty"`

	// IndexDefinition: The definition of the index.  No two indexes have
	// the same definition.
	IndexDefinition *IndexDefinition `json:"indexDefinition,omitempty"`

	// IndexId: The index's ID within the indexes collection resource.
	IndexId string `json:"indexId,omitempty"`

	// ProjectId: The ID of the project to which the index belongs.
	ProjectId string `json:"projectId,omitempty"`

	// State: Required.
	State string `json:"state,omitempty"`
}

type IndexDefinition struct {
	// Kind: The kind of entity to index.
	// This field shares the constraints
	// of field Key.PathElement.kind.
	Kind string `json:"kind,omitempty"`

	// Properties: A sequence of indexed property definitions.
	Properties []*IndexedPropertyDefinition `json:"properties,omitempty"`
}

type IndexedPropertyDefinition struct {
	// Mode: The indexed property's mode.
	Mode string `json:"mode,omitempty"`

	// Name: The indexed property's name.
	// This field shares the constraints
	// of the keys of the map in
	// field Entity.properties.
	// If name includes
	// "."s, it may be interpreted as a property name path.
	// Required.
	Name string `json:"name,omitempty"`
}

type Key struct {
	// PartitionId: Entities are partitioned into subsets, currently
	// identified by a dataset
	// (usually implicitly specified by the project)
	// and namespace ID.
	// Queries are scoped to a single partition.
	PartitionId *PartitionId `json:"partitionId,omitempty"`

	// Path: The entity path.
	// An entity path consists of one or more
	// elements composed of a kind and a
	// string or numerical identifier,
	// which identify entities. The first
	// element identifies a _root
	// entity_, the second element identifies
	// a _child_ of the root entity,
	// the third element identifies a child of the
	// second entity, and so
	// forth. The entities identified by all prefixes of
	// the path are called
	// the element's _ancestors_.
	//
	// An entity path is always fully complete:
	// *all* of the entity's ancestors
	// are required to be in the path along
	// with the entity identifier itself.
	// The only exception is that in some
	// documented cases, the identifier in the
	// last path element (for the
	// entity) itself may be omitted. For example,
	// an entity in
	// `Value.entity_value` may have no key.
	//
	// A path can never be empty, and
	// a path can have at most 100 elements.
	Path []*PathElement `json:"path,omitempty"`
}

type KindExpression struct {
	// Name: The name of the kind.
	Name string `json:"name,omitempty"`
}

type LatLng struct {
	// Latitude: The latitude in degrees. It must be in the range [-90.0,
	// +90.0].
	Latitude float64 `json:"latitude,omitempty"`

	// Longitude: The longitude in degrees. It must be in the range [-180.0,
	// +180.0].
	Longitude float64 `json:"longitude,omitempty"`
}

type ListIndexesResponse struct {
	// Indexes: The indexes.
	Indexes []*Index `json:"indexes,omitempty"`

	// NextPageToken: The standard List next-page token.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type ListOperationsResponse struct {
	// NextPageToken: The standard List next-page token.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Operations: A list of operations that matches the specified filter in
	// the request.
	Operations []*Operation `json:"operations,omitempty"`
}

type LookupIndexRequest struct {
	// DatabaseId: The database id.
	DatabaseId string `json:"databaseId,omitempty"`

	// IndexDefinition: The index definition.
	IndexDefinition *IndexDefinition `json:"indexDefinition,omitempty"`
}

type LookupRequest struct {
	// DatabaseId: If not empty, the ID of the database against which to
	// make the request.
	DatabaseId string `json:"databaseId,omitempty"`

	// Keys: Keys of entities to look up.
	Keys []*Key `json:"keys,omitempty"`

	// PropertyMask: The properties to return. Defaults to returning all
	// properties.
	// If this field is set and an entity has a property not
	// referenced in the
	// mask, it will not be included
	// in
	// google.datastore.v1beta3.LookupResponse.found.entity.properties.
	// Th
	// e entity's key is always returned.
	// If an
	// google.datastore.v1beta3.Value.entity_value is
	// returned, its key is
	// returned as well.
	//
	// The paths in the mask are property paths: dot
	// (`.`) separated property
	// names that can be used to reference
	// properties nested in
	// google.datastore.v1beta3.Value.entity_value.
	// A
	// property must not reference a value inside
	// an
	// google.datastore.v1beta3.Value.array_value.
	// None of these property
	// paths may contain the `__key__` property name.
	PropertyMask string `json:"propertyMask,omitempty"`

	// ReadOptions: The options for this lookup request.
	ReadOptions *ReadOptions `json:"readOptions,omitempty"`
}

type LookupResponse struct {
	// Deferred: A list of keys that were not looked up due to resource
	// constraints. The
	// order of results in this field is undefined and has
	// no relation to the
	// order of the keys in the input.
	Deferred []*Key `json:"deferred,omitempty"`

	// Found: Entities found as `ResultType.FULL` entities. The order of
	// results in this
	// field is undefined and has no relation to the order
	// of the keys in the
	// input.
	Found []*EntityResult `json:"found,omitempty"`

	// Missing: Entities not found as `ResultType.KEY_ONLY` entities. The
	// order of results
	// in this field is undefined and has no relation to
	// the order of the keys
	// in the input.
	Missing []*EntityResult `json:"missing,omitempty"`

	// Transaction: The transaction that was started as part of this Lookup
	// request.
	// Set only when
	// google.datastore.v1beta3.ReadOptions.begin_transaction
	// was set in
	// google.datastore.v1beta3.LookupRequest.read_options.
	Transaction string `json:"transaction,omitempty"`
}

type MultiWatchRequest struct {
	// AddTargets: The set of watch targets to add to this stream.
	// changes
	// will be returned with server
	// assigned `target_ids` in the same order
	// as targets are specified here.
	AddTargets []*WatchTarget `json:"addTargets,omitempty"`

	// DatabaseId: Database ID against which to make the request.
	DatabaseId string `json:"databaseId,omitempty"`

	// RemoveTargets: The IDs of watch targets to remove from this stream.
	RemoveTargets []int64 `json:"removeTargets,omitempty"`
}

type Mutation struct {
	// BaseVersion: The version of the entity that this mutation is being
	// applied to. If this
	// does not match the current version on the server,
	// the mutation conflicts.
	BaseVersion int64 `json:"baseVersion,omitempty,string"`

	// Delete: The key of the entity to delete. The entity may or may not
	// already exist.
	// Must have a complete key path and must not be
	// reserved/read-only.
	Delete *Key `json:"delete,omitempty"`

	// Insert: The entity to insert. The entity must not already exist.
	// The
	// entity key's final path element may be incomplete.
	Insert *Entity `json:"insert,omitempty"`

	// PropertyMask: The properties to write in this mutation.
	// This field is
	// ignored for `delete`.
	//
	// If the entity already exists, only properties
	// referenced in the mask are
	// updated, others are left
	// untouched.
	// Properties referenced in the mask but not in the entity
	// are deleted.
	// Properties not referenced in the mask may not be set in
	// the entity.
	//
	// The paths in the mask follow the same rules as specified
	// in
	// google.datastore.v1beta3.LookupRequest.property_mask.
	// Additionally,
	//  none of these property paths may contain
	// a
	// google.datastore.v1beta3.Entity.properties.
	PropertyMask string `json:"propertyMask,omitempty"`

	// Update: The entity to update. The entity must already exist.
	// Must
	// have a complete key path.
	Update *Entity `json:"update,omitempty"`

	// Upsert: The entity to upsert. The entity may or may not already
	// exist.
	// The entity key's final path element may be incomplete.
	Upsert *Entity `json:"upsert,omitempty"`
}

type MutationResult struct {
	// ConflictDetected: Whether a conflict was detected for this mutation.
	// Always false when a
	// conflict detection strategy field is not set in
	// the mutation.
	ConflictDetected bool `json:"conflictDetected,omitempty"`

	// Key: The automatically allocated key.
	// Set only when the mutation
	// allocated a key.
	Key *Key `json:"key,omitempty"`

	// Version: The version of the entity on the server after processing the
	// mutation. If
	// the mutation doesn't change anything on the server, then
	// the version will
	// be the version of the current entity or, if no
	// entity is present, a version
	// that is strictly greater than the
	// version of any previous entity and less
	// than the version of any
	// possible future entity.
	Version int64 `json:"version,omitempty,string"`
}

type Operation struct {
	// Done: If the value is `false`, it means the operation is still in
	// progress.
	// If true, the operation is completed, and either `error` or
	// `response` is
	// available.
	Done bool `json:"done,omitempty"`

	// Error: The error result of the operation in case of failure.
	Error *Status `json:"error,omitempty"`

	// Metadata: Service-specific metadata associated with the operation.
	// It typically
	// contains progress information and common metadata such
	// as create time.
	// Some services might not provide such metadata.  Any
	// method that returns a
	// long-running operation should document the
	// metadata type, if any.
	Metadata *OperationMetadata `json:"metadata,omitempty"`

	// Name: The server-assigned name, which is only unique within the same
	// service that
	// originally returns it. If you use the default HTTP
	// mapping above, the
	// `name` should have the format of
	// `operations/some/unique/name`.
	Name string `json:"name,omitempty"`

	// Response: The normal response of the operation in case of success.
	// If the original
	// method returns no data on success, such as `Delete`,
	// the response is
	// `google.protobuf.Empty`.  If the original method is
	// standard
	// `Get`/`Create`/`Update`, the response should be the
	// resource.  For other
	// methods, the response should have the type
	// `XxxResponse`, where `Xxx`
	// is the original method name.  For example,
	// if the original method name
	// is `TakeSnapshot()`, the inferred
	// response type is
	// `TakeSnapshotResponse`.
	Response *OperationResponse `json:"response,omitempty"`
}

type OperationMetadata struct {
}

type OperationResponse struct {
}

type PartitionId struct {
	// DatabaseId: If not empty, the ID of the database to which the
	// entities
	// belong.
	DatabaseId string `json:"databaseId,omitempty"`

	// NamespaceId: If not empty, the ID of the namespace to which the
	// entities belong.
	NamespaceId string `json:"namespaceId,omitempty"`

	// ProjectId: The ID of the project to which the entities belong.
	ProjectId string `json:"projectId,omitempty"`
}

type PathElement struct {
	// Id: The auto-allocated ID of the entity.
	// Never equal to zero. Values
	// less than zero are discouraged and may not
	// be supported in the
	// future.
	Id int64 `json:"id,omitempty,string"`

	// Kind: The kind of the entity.
	// A kind matching regex `__.*__` is
	// reserved/read-only.
	// A kind must not contain more than 1500 bytes when
	// UTF-8 encoded.
	// Cannot be `""`.
	Kind string `json:"kind,omitempty"`

	// Name: The name of the entity.
	// A name matching regex `__.*__` is
	// reserved/read-only.
	// A name must not be more than 1500 bytes when
	// UTF-8 encoded.
	// Cannot be `""`.
	Name string `json:"name,omitempty"`
}

type Progress struct {
	// WorkCompleted: Note that this may be greater than work_estimated.
	WorkCompleted int64 `json:"workCompleted,omitempty,string"`

	// WorkEstimated: An estimate of how much work needs to be performed.
	// May be zero if the
	// work estimate is unavailable.
	WorkEstimated int64 `json:"workEstimated,omitempty,string"`
}

type Projection struct {
	// Property: The property to project.
	Property *PropertyReference `json:"property,omitempty"`
}

type PropertyFilter struct {
	// Op: The operator to filter by.
	Op string `json:"op,omitempty"`

	// Property: The property to filter by.
	Property *PropertyReference `json:"property,omitempty"`

	// Value: The value to compare the property to.
	Value *Value `json:"value,omitempty"`
}

type PropertyOrder struct {
	// Direction: The direction to order by. Defaults to `ASCENDING`.
	Direction string `json:"direction,omitempty"`

	// Property: The property to order by.
	Property *PropertyReference `json:"property,omitempty"`
}

type PropertyReference struct {
	// Name: The name of the property.
	Name string `json:"name,omitempty"`
}

type Query struct {
	// DistinctOn: The properties to make distinct. The query results will
	// contain the first
	// result for each distinct combination of values for
	// the given properties
	// (if empty, all results are returned).
	DistinctOn []*PropertyReference `json:"distinctOn,omitempty"`

	// EndCursor: An ending point for the query results. Query cursors
	// are
	// returned in query result batches.
	EndCursor string `json:"endCursor,omitempty"`

	// Filter: The filter to apply.
	Filter *Filter `json:"filter,omitempty"`

	// Kind: The kinds to query (if empty, returns entities of all
	// kinds).
	// Currently at most 1 kind may be specified.
	Kind []*KindExpression `json:"kind,omitempty"`

	// Limit: The maximum number of results to return. Applies after all
	// other
	// constraints. Optional.
	// Unspecified is interpreted as no
	// limit.
	// Must be >= 0 if specified.
	Limit int64 `json:"limit,omitempty"`

	// Offset: The number of results to skip. Applies before limit, but
	// after all other
	// constraints. Optional. Must be >= 0 if specified.
	Offset int64 `json:"offset,omitempty"`

	// Order: The order to apply to the query results (if empty, order is
	// unspecified).
	Order []*PropertyOrder `json:"order,omitempty"`

	// Projection: The projection to return. Defaults to returning all
	// properties.
	Projection []*Projection `json:"projection,omitempty"`

	// StartCursor: A starting point for the query results. Query cursors
	// are
	// returned in query result batches.
	StartCursor string `json:"startCursor,omitempty"`
}

type QueryResultBatch struct {
	// EndCursor: A cursor that points to the position after the last result
	// in the batch.
	EndCursor string `json:"endCursor,omitempty"`

	// EntityResultType: The result type for every entity in
	// `entity_results`.
	EntityResultType string `json:"entityResultType,omitempty"`

	// EntityResults: The results for this batch.
	EntityResults []*EntityResult `json:"entityResults,omitempty"`

	// MoreResults: The state of the query after the current batch.
	MoreResults string `json:"moreResults,omitempty"`

	// SkippedCursor: A cursor that points to the position after the last
	// skipped result.
	// Will be set when `skipped_results` != 0.
	SkippedCursor string `json:"skippedCursor,omitempty"`

	// SkippedResults: The number of results skipped, typically because of
	// an offset.
	SkippedResults int64 `json:"skippedResults,omitempty"`

	// SnapshotVersion: The version number of the snapshot this batch was
	// returned from.
	// This applies to the range of results from the query's
	// `start_cursor` (or
	// the beginning of the query if no cursor was given)
	// to this batch's
	// `end_cursor` (not the query's `end_cursor`).
	//
	// In a
	// single transaction, subsequent query result batches for the same
	// query
	// can have a greater snapshot version number. Each batch's
	// snapshot version
	// is valid for all preceding batches.
	SnapshotVersion int64 `json:"snapshotVersion,omitempty,string"`
}

type ReadOnly struct {
}

type ReadOptions struct {
	// NewTransaction: Options for beginning a new transaction for this
	// request.
	// The new transaction handle will be returned in the
	// corresponding response
	// as either
	// google.datastore.v1beta3.LookupResponse.transaction
	// or
	// google.datastore.v1beta3.RunQueryResponse.transaction.
	NewTransaction *TransactionOptions `json:"newTransaction,omitempty"`

	// ReadConsistency: The non-transactional read consistency to
	// use.
	// Cannot be set to `STRONG` for global queries.
	ReadConsistency string `json:"readConsistency,omitempty"`

	// Transaction: The transaction in which to read.
	Transaction string `json:"transaction,omitempty"`
}

type ReadWrite struct {
}

type Rectangle struct {
	// Northeast: The northeast point of the rectangle.
	// Its latitude must be
	// not less than that of the southwest point.
	Northeast *LatLng `json:"northeast,omitempty"`

	// Southwest: The southwest point of the rectangle.
	Southwest *LatLng `json:"southwest,omitempty"`
}

type RollbackRequest struct {
	// DatabaseId: If not empty, the ID of the database against which to
	// make the request.
	DatabaseId string `json:"databaseId,omitempty"`

	// Transaction: The transaction identifier, returned by a call
	// to
	// google.datastore.v1beta3.Datastore.BeginTransaction.
	Transaction string `json:"transaction,omitempty"`
}

type RollbackResponse struct {
}

type RunQueryRequest struct {
	// DatabaseId: If not empty, the ID of the database against which to
	// make the request.
	DatabaseId string `json:"databaseId,omitempty"`

	// GqlQuery: The GQL query to run.
	GqlQuery *GqlQuery `json:"gqlQuery,omitempty"`

	// PartitionId: Entities are partitioned into subsets, identified by a
	// partition ID.
	// Queries are scoped to a single partition.
	// This
	// partition ID is normalized with the standard default
	// context
	// partition ID.
	PartitionId *PartitionId `json:"partitionId,omitempty"`

	// PropertyMask: The properties to return.
	// This field must not be set
	// for a projection query.
	//
	// See
	// google.datastore.v1beta3.LookupRequest.property_mask.
	PropertyMask string `json:"propertyMask,omitempty"`

	// Query: The query to run.
	Query *Query `json:"query,omitempty"`

	// ReadOptions: The options for this query.
	ReadOptions *ReadOptions `json:"readOptions,omitempty"`
}

type RunQueryResponse struct {
	// Batch: A batch of query results (always present).
	Batch *QueryResultBatch `json:"batch,omitempty"`

	// Query: The parsed form of the `GqlQuery` from the request, if it was
	// set.
	Query *Query `json:"query,omitempty"`

	// Transaction: The transaction that was started as part of this
	// RunQuery request.
	// Set only when
	// google.datastore.v1beta3.ReadOptions.begin_transaction
	// was set in
	// google.datastore.v1beta3.RunQueryRequest.read_options.
	Transaction string `json:"transaction,omitempty"`
}

type StContainsFilter struct {
	// ContainedIn: A region within which the property's value should be
	// contained.
	ContainedIn *GeoRegion `json:"containedIn,omitempty"`

	// Property: The property to filter by.
	Property *PropertyReference `json:"property,omitempty"`
}

type Status struct {
	// Code: The status code, which should be an enum value of
	// google.rpc.Code.
	Code int64 `json:"code,omitempty"`

	// Details: A list of messages that carry the error details.  There will
	// be a
	// common set of message types for APIs to use.
	Details []*StatusDetails `json:"details,omitempty"`

	// Message: A developer-facing error message, which should be in
	// English. Any
	// user-facing error message should be localized and sent
	// in the
	// google.rpc.Status.details field, or localized by the client.
	Message string `json:"message,omitempty"`
}

type StatusDetails struct {
}

type TransactionOptions struct {
	// ReadOnly: The transaction should only allow reads.
	ReadOnly *ReadOnly `json:"readOnly,omitempty"`

	// ReadWrite: The transaction should allow both reads and writes.
	ReadWrite *ReadWrite `json:"readWrite,omitempty"`
}

type UpdateIndexMetadata struct {
	// Common: Metadata common to all Datastore Admin Operations.
	Common *CommonMetadata `json:"common,omitempty"`

	// IndexDefinition: The index definition.  Handy for filtering.
	IndexDefinition *IndexDefinition `json:"indexDefinition,omitempty"`

	// IndexId: The format matches that of Index.index_id.
	IndexId string `json:"indexId,omitempty"`
}

type Value struct {
	// ArrayValue: An array value.
	// Cannot contain another array value.
	// A
	// `Value` instance that sets field `array_value` must not set
	// fields
	// `meaning` or `exclude_from_indexes`.
	ArrayValue *ArrayValue `json:"arrayValue,omitempty"`

	// BlobValue: A blob value.
	// May have at most 1,000,000 bytes.
	// When
	// `exclude_from_indexes` is false, may have at most 1500 bytes.
	// In JSON
	// requests, must be base64-encoded.
	BlobValue string `json:"blobValue,omitempty"`

	// BooleanValue: A boolean value.
	BooleanValue bool `json:"booleanValue,omitempty"`

	// DoubleValue: A double value.
	DoubleValue float64 `json:"doubleValue,omitempty"`

	// EntityValue: An entity value.
	//
	// - May have no key.
	// - May have a key
	// with an incomplete key path.
	// - May have a reserved/read-only key.
	EntityValue *Entity `json:"entityValue,omitempty"`

	// ExcludeFromIndexes: If the value should be excluded from all indexes
	// including those defined
	// explicitly.
	ExcludeFromIndexes bool `json:"excludeFromIndexes,omitempty"`

	// GeoPointValue: A geo point value representing a point on the surface
	// of Earth.
	GeoPointValue *LatLng `json:"geoPointValue,omitempty"`

	// IntegerValue: An integer value.
	IntegerValue int64 `json:"integerValue,omitempty,string"`

	// KeyValue: A key value.
	KeyValue *Key `json:"keyValue,omitempty"`

	// Meaning: The `meaning` field should only be populated for backwards
	// compatibility.
	Meaning int64 `json:"meaning,omitempty"`

	// NullValue: A null value.
	NullValue string `json:"nullValue,omitempty"`

	// StringValue: A UTF-8 encoded string value.
	// When
	// `exclude_from_indexes` is false (it is indexed) , may have at most
	// 1500 bytes.
	// Otherwise, may be set to at least 1,000,000 bytes.
	StringValue string `json:"stringValue,omitempty"`

	// TimestampValue: A timestamp value.
	// When stored in the Datastore,
	// precise only to microseconds;
	// any additional precision is rounded
	// down.
	TimestampValue string `json:"timestampValue,omitempty"`
}

type WatchRequest struct {
	// DatabaseId: Database ID against which to make the request.
	DatabaseId string `json:"databaseId,omitempty"`

	// Targets: The set of watcher targets to include in the stream.
	// changes
	// will be returned with server
	// assigned `target_ids` in the same order
	// as the targets are specified here.
	Targets []*WatchTarget `json:"targets,omitempty"`
}

type WatchTarget struct {
	// GqlQuery: The GQL query to watch.
	GqlQuery *GqlQuery `json:"gqlQuery,omitempty"`

	// PartitionId: The partition id to watch.
	PartitionId *PartitionId `json:"partitionId,omitempty"`

	// Query: The query to watch.
	Query *Query `json:"query,omitempty"`

	// ResumeToken: A resume token from a stream containing an identical
	// watch target.
	ResumeToken string `json:"resumeToken,omitempty"`
}

// method id "datastore.projects.allocateIds":

type ProjectsAllocateIdsCall struct {
	s                  *Service
	projectId          string
	allocateidsrequest *AllocateIdsRequest
	opt_               map[string]interface{}
}

// AllocateIds: Allocates IDs for the given keys, which is useful for
// referencing an entity
// before it is inserted.
func (r *ProjectsService) AllocateIds(projectId string, allocateidsrequest *AllocateIdsRequest) *ProjectsAllocateIdsCall {
	c := &ProjectsAllocateIdsCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.allocateidsrequest = allocateidsrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAllocateIdsCall) Fields(s ...googleapi.Field) *ProjectsAllocateIdsCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsAllocateIdsCall) Do() (*AllocateIdsResponse, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.allocateidsrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectId}:allocateIds")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
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
	var ret *AllocateIdsResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Allocates IDs for the given keys, which is useful for referencing an entity\nbefore it is inserted.",
	//   "flatPath": "v1beta3/projects/{projectId}:allocateIds",
	//   "httpMethod": "POST",
	//   "id": "datastore.projects.allocateIds",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "The ID of the project against which to make the request.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectId}:allocateIds",
	//   "request": {
	//     "$ref": "AllocateIdsRequest"
	//   },
	//   "response": {
	//     "$ref": "AllocateIdsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/datastore"
	//   ]
	// }

}

// method id "datastore.projects.beginTransaction":

type ProjectsBeginTransactionCall struct {
	s                       *Service
	projectId               string
	begintransactionrequest *BeginTransactionRequest
	opt_                    map[string]interface{}
}

// BeginTransaction: Begins a new transaction.
func (r *ProjectsService) BeginTransaction(projectId string, begintransactionrequest *BeginTransactionRequest) *ProjectsBeginTransactionCall {
	c := &ProjectsBeginTransactionCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.begintransactionrequest = begintransactionrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsBeginTransactionCall) Fields(s ...googleapi.Field) *ProjectsBeginTransactionCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsBeginTransactionCall) Do() (*BeginTransactionResponse, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.begintransactionrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectId}:beginTransaction")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
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
	var ret *BeginTransactionResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Begins a new transaction.",
	//   "flatPath": "v1beta3/projects/{projectId}:beginTransaction",
	//   "httpMethod": "POST",
	//   "id": "datastore.projects.beginTransaction",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "The ID of the project against which to make the request.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectId}:beginTransaction",
	//   "request": {
	//     "$ref": "BeginTransactionRequest"
	//   },
	//   "response": {
	//     "$ref": "BeginTransactionResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/datastore"
	//   ]
	// }

}

// method id "datastore.projects.commit":

type ProjectsCommitCall struct {
	s             *Service
	projectId     string
	commitrequest *CommitRequest
	opt_          map[string]interface{}
}

// Commit: Commits a transaction, optionally creating, deleting or
// modifying some
// entities.
func (r *ProjectsService) Commit(projectId string, commitrequest *CommitRequest) *ProjectsCommitCall {
	c := &ProjectsCommitCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.commitrequest = commitrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsCommitCall) Fields(s ...googleapi.Field) *ProjectsCommitCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsCommitCall) Do() (*CommitResponse, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.commitrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectId}:commit")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
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
	var ret *CommitResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Commits a transaction, optionally creating, deleting or modifying some\nentities.",
	//   "flatPath": "v1beta3/projects/{projectId}:commit",
	//   "httpMethod": "POST",
	//   "id": "datastore.projects.commit",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "The ID of the project against which to make the request.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectId}:commit",
	//   "request": {
	//     "$ref": "CommitRequest"
	//   },
	//   "response": {
	//     "$ref": "CommitResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/datastore"
	//   ]
	// }

}

// method id "datastore.projects.export":

type ProjectsExportCall struct {
	s             *Service
	projectId     string
	exportrequest *ExportRequest
	opt_          map[string]interface{}
}

// Export: Exports a copy of all or a subset of Entities from a Google
// Cloud Datastore
// Project to another storage system, such as Google
// Cloud Storage. Recent
// updates to Entities may not be reflected in the
// export. The export occurs
// in the background and its progress can be
// monitored and managed via the
// Operation resource that is created.
// The output of an export may only be
// used once the associated
// Operation is done. If an export Operation is
// cancelled before
// completion it may leave partial data behind in Google
// Cloud Storage.
func (r *ProjectsService) Export(projectId string, exportrequest *ExportRequest) *ProjectsExportCall {
	c := &ProjectsExportCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.exportrequest = exportrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsExportCall) Fields(s ...googleapi.Field) *ProjectsExportCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsExportCall) Do() (*Operation, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.exportrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectId}:export")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
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
	//   "description": "Exports a copy of all or a subset of Entities from a Google Cloud Datastore\nProject to another storage system, such as Google Cloud Storage. Recent\nupdates to Entities may not be reflected in the export. The export occurs\nin the background and its progress can be monitored and managed via the\nOperation resource that is created.  The output of an export may only be\nused once the associated Operation is done. If an export Operation is\ncancelled before completion it may leave partial data behind in Google\nCloud Storage.",
	//   "flatPath": "v1beta3/projects/{projectId}:export",
	//   "httpMethod": "POST",
	//   "id": "datastore.projects.export",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "Project ID against which to make the request.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectId}:export",
	//   "request": {
	//     "$ref": "ExportRequest"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/datastore"
	//   ]
	// }

}

// method id "datastore.projects.import":

type ProjectsImportCall struct {
	s             *Service
	projectId     string
	importrequest *ImportRequest
	opt_          map[string]interface{}
}

// Import: Imports Entities into a Google Cloud Datastore Project.
// Existing Entities
// with the same key are overwritten. The import
// occurs in the background and
// its progress can be monitored and
// managed via the Operation resource that
// is created.  If an Import
// Operation is cancelled, it is possible that a
// subset of the data has
// already been imported to the Datastore.
func (r *ProjectsService) Import(projectId string, importrequest *ImportRequest) *ProjectsImportCall {
	c := &ProjectsImportCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.importrequest = importrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsImportCall) Fields(s ...googleapi.Field) *ProjectsImportCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsImportCall) Do() (*Operation, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.importrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectId}:import")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
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
	//   "description": "Imports Entities into a Google Cloud Datastore Project. Existing Entities\nwith the same key are overwritten. The import occurs in the background and\nits progress can be monitored and managed via the Operation resource that\nis created.  If an Import Operation is cancelled, it is possible that a\nsubset of the data has already been imported to the Datastore.",
	//   "flatPath": "v1beta3/projects/{projectId}:import",
	//   "httpMethod": "POST",
	//   "id": "datastore.projects.import",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "Project ID against which to make the request.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectId}:import",
	//   "request": {
	//     "$ref": "ImportRequest"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/datastore"
	//   ]
	// }

}

// method id "datastore.projects.lookup":

type ProjectsLookupCall struct {
	s             *Service
	projectId     string
	lookuprequest *LookupRequest
	opt_          map[string]interface{}
}

// Lookup: Looks up entities by key.
func (r *ProjectsService) Lookup(projectId string, lookuprequest *LookupRequest) *ProjectsLookupCall {
	c := &ProjectsLookupCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.lookuprequest = lookuprequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLookupCall) Fields(s ...googleapi.Field) *ProjectsLookupCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsLookupCall) Do() (*LookupResponse, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.lookuprequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectId}:lookup")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
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
	var ret *LookupResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Looks up entities by key.",
	//   "flatPath": "v1beta3/projects/{projectId}:lookup",
	//   "httpMethod": "POST",
	//   "id": "datastore.projects.lookup",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "The ID of the project against which to make the request.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectId}:lookup",
	//   "request": {
	//     "$ref": "LookupRequest"
	//   },
	//   "response": {
	//     "$ref": "LookupResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/datastore"
	//   ]
	// }

}

// method id "datastore.projects.multiWatch":

type ProjectsMultiWatchCall struct {
	s                 *Service
	projectId         string
	multiwatchrequest *MultiWatchRequest
	opt_              map[string]interface{}
}

// MultiWatch: Watch changes to the results of a dynamically changeable
// set of queries.
func (r *ProjectsService) MultiWatch(projectId string, multiwatchrequest *MultiWatchRequest) *ProjectsMultiWatchCall {
	c := &ProjectsMultiWatchCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.multiwatchrequest = multiwatchrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsMultiWatchCall) Fields(s ...googleapi.Field) *ProjectsMultiWatchCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsMultiWatchCall) Do() (*ChangeBatch, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.multiwatchrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectId}:multiWatch")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
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
	var ret *ChangeBatch
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Watch changes to the results of a dynamically changeable set of queries.",
	//   "flatPath": "v1beta3/projects/{projectId}:multiWatch",
	//   "httpMethod": "POST",
	//   "id": "datastore.projects.multiWatch",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "Project ID against which to make the request.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectId}:multiWatch",
	//   "request": {
	//     "$ref": "MultiWatchRequest"
	//   },
	//   "response": {
	//     "$ref": "ChangeBatch"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/datastore"
	//   ]
	// }

}

// method id "datastore.projects.rollback":

type ProjectsRollbackCall struct {
	s               *Service
	projectId       string
	rollbackrequest *RollbackRequest
	opt_            map[string]interface{}
}

// Rollback: Rolls back a transaction.
func (r *ProjectsService) Rollback(projectId string, rollbackrequest *RollbackRequest) *ProjectsRollbackCall {
	c := &ProjectsRollbackCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.rollbackrequest = rollbackrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsRollbackCall) Fields(s ...googleapi.Field) *ProjectsRollbackCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsRollbackCall) Do() (*RollbackResponse, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.rollbackrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectId}:rollback")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
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
	var ret *RollbackResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Rolls back a transaction.",
	//   "flatPath": "v1beta3/projects/{projectId}:rollback",
	//   "httpMethod": "POST",
	//   "id": "datastore.projects.rollback",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "The ID of the project against which to make the request.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectId}:rollback",
	//   "request": {
	//     "$ref": "RollbackRequest"
	//   },
	//   "response": {
	//     "$ref": "RollbackResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/datastore"
	//   ]
	// }

}

// method id "datastore.projects.runQuery":

type ProjectsRunQueryCall struct {
	s               *Service
	projectId       string
	runqueryrequest *RunQueryRequest
	opt_            map[string]interface{}
}

// RunQuery: Queries for entities.
func (r *ProjectsService) RunQuery(projectId string, runqueryrequest *RunQueryRequest) *ProjectsRunQueryCall {
	c := &ProjectsRunQueryCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.runqueryrequest = runqueryrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsRunQueryCall) Fields(s ...googleapi.Field) *ProjectsRunQueryCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsRunQueryCall) Do() (*RunQueryResponse, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.runqueryrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectId}:runQuery")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
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
	var ret *RunQueryResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Queries for entities.",
	//   "flatPath": "v1beta3/projects/{projectId}:runQuery",
	//   "httpMethod": "POST",
	//   "id": "datastore.projects.runQuery",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "The ID of the project against which to make the request.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectId}:runQuery",
	//   "request": {
	//     "$ref": "RunQueryRequest"
	//   },
	//   "response": {
	//     "$ref": "RunQueryResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/datastore"
	//   ]
	// }

}

// method id "datastore.projects.watch":

type ProjectsWatchCall struct {
	s            *Service
	projectId    string
	watchrequest *WatchRequest
	opt_         map[string]interface{}
}

// Watch: Watch changes to the results of a given set of queries.
func (r *ProjectsService) Watch(projectId string, watchrequest *WatchRequest) *ProjectsWatchCall {
	c := &ProjectsWatchCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.watchrequest = watchrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsWatchCall) Fields(s ...googleapi.Field) *ProjectsWatchCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsWatchCall) Do() (*ChangeBatch, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.watchrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectId}:watch")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
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
	var ret *ChangeBatch
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Watch changes to the results of a given set of queries.",
	//   "flatPath": "v1beta3/projects/{projectId}:watch",
	//   "httpMethod": "POST",
	//   "id": "datastore.projects.watch",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "Project ID against which to make the request.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectId}:watch",
	//   "request": {
	//     "$ref": "WatchRequest"
	//   },
	//   "response": {
	//     "$ref": "ChangeBatch"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/datastore"
	//   ]
	// }

}

// method id "datastore.projects.indexes.get":

type ProjectsIndexesGetCall struct {
	s         *Service
	projectId string
	indexId   string
	opt_      map[string]interface{}
}

// Get: Gets an index.
func (r *ProjectsIndexesService) Get(projectId string, indexId string) *ProjectsIndexesGetCall {
	c := &ProjectsIndexesGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.indexId = indexId
	return c
}

// DatabaseId sets the optional parameter "databaseId": The database id.
func (c *ProjectsIndexesGetCall) DatabaseId(databaseId string) *ProjectsIndexesGetCall {
	c.opt_["databaseId"] = databaseId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsIndexesGetCall) Fields(s ...googleapi.Field) *ProjectsIndexesGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsIndexesGetCall) Do() (*Index, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["databaseId"]; ok {
		params.Set("databaseId", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectId}/indexes/{+indexId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"indexId":   c.indexId,
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
	var ret *Index
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets an index.",
	//   "flatPath": "v1beta3/projects/{projectId}/indexes/{indexesId}",
	//   "httpMethod": "GET",
	//   "id": "datastore.projects.indexes.get",
	//   "parameterOrder": [
	//     "projectId",
	//     "indexId"
	//   ],
	//   "parameters": {
	//     "databaseId": {
	//       "description": "The database id.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "indexId": {
	//       "description": "The format matches that of Index.index_id.",
	//       "location": "path",
	//       "pattern": "^[^/]*$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "The project id.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectId}/indexes/{+indexId}",
	//   "response": {
	//     "$ref": "Index"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/datastore"
	//   ]
	// }

}

// method id "datastore.projects.indexes.list":

type ProjectsIndexesListCall struct {
	s         *Service
	projectId string
	opt_      map[string]interface{}
}

// List: Lists the indexes that match the specified filters.
// Only lists
// indexes that are not in their initial state.
func (r *ProjectsIndexesService) List(projectId string) *ProjectsIndexesListCall {
	c := &ProjectsIndexesListCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	return c
}

// DatabaseId sets the optional parameter "databaseId": The database id.
func (c *ProjectsIndexesListCall) DatabaseId(databaseId string) *ProjectsIndexesListCall {
	c.opt_["databaseId"] = databaseId
	return c
}

// Filter sets the optional parameter "filter": The standard List
// filter.
// The filter expression may filter only on definition.kind,
// eg
// '(definition.kind == "User") or (definition.kind == "Widget")'.
func (c *ProjectsIndexesListCall) Filter(filter string) *ProjectsIndexesListCall {
	c.opt_["filter"] = filter
	return c
}

// PageSize sets the optional parameter "pageSize": The standard List
// page size.
func (c *ProjectsIndexesListCall) PageSize(pageSize int64) *ProjectsIndexesListCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken": The standard List
// page token.
func (c *ProjectsIndexesListCall) PageToken(pageToken string) *ProjectsIndexesListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsIndexesListCall) Fields(s ...googleapi.Field) *ProjectsIndexesListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsIndexesListCall) Do() (*ListIndexesResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["databaseId"]; ok {
		params.Set("databaseId", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["filter"]; ok {
		params.Set("filter", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageSize"]; ok {
		params.Set("pageSize", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectId}/indexes")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
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
	var ret *ListIndexesResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists the indexes that match the specified filters.\nOnly lists indexes that are not in their initial state.",
	//   "flatPath": "v1beta3/projects/{projectId}/indexes",
	//   "httpMethod": "GET",
	//   "id": "datastore.projects.indexes.list",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "databaseId": {
	//       "description": "The database id.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "filter": {
	//       "description": "The standard List filter.\nThe filter expression may filter only on definition.kind, eg\n'(definition.kind == \"User\") or (definition.kind == \"Widget\")'.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "The standard List page size.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The standard List page token.",
	//       "format": "byte",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "The project id.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectId}/indexes",
	//   "response": {
	//     "$ref": "ListIndexesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/datastore"
	//   ]
	// }

}

// method id "datastore.projects.indexes.lookup":

type ProjectsIndexesLookupCall struct {
	s                  *Service
	projectId          string
	lookupindexrequest *LookupIndexRequest
	opt_               map[string]interface{}
}

// Lookup: Looks up an index by definition.
func (r *ProjectsIndexesService) Lookup(projectId string, lookupindexrequest *LookupIndexRequest) *ProjectsIndexesLookupCall {
	c := &ProjectsIndexesLookupCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.lookupindexrequest = lookupindexrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsIndexesLookupCall) Fields(s ...googleapi.Field) *ProjectsIndexesLookupCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsIndexesLookupCall) Do() (*Index, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.lookupindexrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectId}/indexes:lookup")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
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
	var ret *Index
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Looks up an index by definition.",
	//   "flatPath": "v1beta3/projects/{projectId}/indexes:lookup",
	//   "httpMethod": "POST",
	//   "id": "datastore.projects.indexes.lookup",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "The project id.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectId}/indexes:lookup",
	//   "request": {
	//     "$ref": "LookupIndexRequest"
	//   },
	//   "response": {
	//     "$ref": "Index"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/datastore"
	//   ]
	// }

}

// method id "datastore.projects.indexes.update":

type ProjectsIndexesUpdateCall struct {
	s         *Service
	projectId string
	indexId   string
	index     *Index
	opt_      map[string]interface{}
}

// Update: Updates an index's state.
// The input index must specify a
// (project_id, index_id) tuple
// or an index definition (but not both),
// and a new state.
// This new state must be SERVING or OFF.
// The state of
// the key index and the kind index cannot be updated.
// If the index is
// already in the requested state, does nothing and returns
// a successful
// but unnamed operation.  Otherwise:
// Returns an unfinished operation.
// -
// If the new state is SERVING, sets the index's state to BUILDING and
// the
//     result operation's field metadata.common.operation_type is
// BUILD_INDEX.
// - If the new state is OFF, sets the index's state to
// CLEARING and the
//    result operation's field
// metadata.common.operation_type is CLEAR_INDEX.
// Once the operation
// finishes,
// if it is successful the index's state is the new state,
// and
// otherwise the index's state is ERROR.
// The result operation's field
// response is of type google.protobuf.Empty.
// The result operation's
// field metadata is of type UpdateIndexMetadata.
func (r *ProjectsIndexesService) Update(projectId string, indexId string, index *Index) *ProjectsIndexesUpdateCall {
	c := &ProjectsIndexesUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.indexId = indexId
	c.index = index
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsIndexesUpdateCall) Fields(s ...googleapi.Field) *ProjectsIndexesUpdateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsIndexesUpdateCall) Do() (*Operation, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.index)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectId}/indexes/{+indexId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"indexId":   c.indexId,
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
	//   "description": "Updates an index's state.\nThe input index must specify a (project_id, index_id) tuple\nor an index definition (but not both), and a new state.\nThis new state must be SERVING or OFF.\nThe state of the key index and the kind index cannot be updated.\nIf the index is already in the requested state, does nothing and returns\na successful but unnamed operation.  Otherwise:\nReturns an unfinished operation.\n- If the new state is SERVING, sets the index's state to BUILDING and the\n    result operation's field metadata.common.operation_type is BUILD_INDEX.\n- If the new state is OFF, sets the index's state to CLEARING and the\n   result operation's field metadata.common.operation_type is CLEAR_INDEX.\nOnce the operation finishes,\nif it is successful the index's state is the new state,\nand otherwise the index's state is ERROR.\nThe result operation's field response is of type google.protobuf.Empty.\nThe result operation's field metadata is of type UpdateIndexMetadata.",
	//   "flatPath": "v1beta3/projects/{projectId}/indexes/{indexesId}",
	//   "httpMethod": "PUT",
	//   "id": "datastore.projects.indexes.update",
	//   "parameterOrder": [
	//     "projectId",
	//     "indexId"
	//   ],
	//   "parameters": {
	//     "indexId": {
	//       "description": "The index's ID within the indexes collection resource.",
	//       "location": "path",
	//       "pattern": "^[^/]*$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "The ID of the project to which the index belongs.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectId}/indexes/{+indexId}",
	//   "request": {
	//     "$ref": "Index"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/datastore"
	//   ]
	// }

}

// method id "datastore.projects.operations.cancel":

type ProjectsOperationsCancelCall struct {
	s    *Service
	name string
	opt_ map[string]interface{}
}

// Cancel: Starts asynchronous cancellation on a long-running operation.
//  The server
// makes a best effort to cancel the operation, but success
// is not
// guaranteed.  If the server doesn't support this method, it
// returns
// `google.rpc.Code.UNIMPLEMENTED`.  Clients can
// use
// Operations.GetOperation or
// other methods to check whether the
// cancellation succeeded or whether the
// operation completed despite
// cancellation.
func (r *ProjectsOperationsService) Cancel(name string) *ProjectsOperationsCancelCall {
	c := &ProjectsOperationsCancelCall{s: r.s, opt_: make(map[string]interface{})}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsOperationsCancelCall) Fields(s ...googleapi.Field) *ProjectsOperationsCancelCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsOperationsCancelCall) Do() (*Empty, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/{+name}:cancel")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
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
	var ret *Empty
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Starts asynchronous cancellation on a long-running operation.  The server\nmakes a best effort to cancel the operation, but success is not\nguaranteed.  If the server doesn't support this method, it returns\n`google.rpc.Code.UNIMPLEMENTED`.  Clients can use\nOperations.GetOperation or\nother methods to check whether the cancellation succeeded or whether the\noperation completed despite cancellation.",
	//   "flatPath": "v1beta3/projects/{projectsId}/operations/{operationsId}:cancel",
	//   "httpMethod": "POST",
	//   "id": "datastore.projects.operations.cancel",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The name of the operation resource to be cancelled.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]*/operations/[^/]*$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/{+name}:cancel",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/datastore"
	//   ]
	// }

}

// method id "datastore.projects.operations.delete":

type ProjectsOperationsDeleteCall struct {
	s    *Service
	name string
	opt_ map[string]interface{}
}

// Delete: Deletes a long-running operation. This method indicates that
// the client is
// no longer interested in the operation result. It does
// not cancel the
// operation. If the server doesn't support this method,
// it returns
// `google.rpc.Code.UNIMPLEMENTED`.
func (r *ProjectsOperationsService) Delete(name string) *ProjectsOperationsDeleteCall {
	c := &ProjectsOperationsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsOperationsDeleteCall) Fields(s ...googleapi.Field) *ProjectsOperationsDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsOperationsDeleteCall) Do() (*Empty, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/{+name}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
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
	var ret *Empty
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes a long-running operation. This method indicates that the client is\nno longer interested in the operation result. It does not cancel the\noperation. If the server doesn't support this method, it returns\n`google.rpc.Code.UNIMPLEMENTED`.",
	//   "flatPath": "v1beta3/projects/{projectsId}/operations/{operationsId}",
	//   "httpMethod": "DELETE",
	//   "id": "datastore.projects.operations.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The name of the operation resource to be deleted.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]*/operations/[^/]*$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/{+name}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/datastore"
	//   ]
	// }

}

// method id "datastore.projects.operations.get":

type ProjectsOperationsGetCall struct {
	s    *Service
	name string
	opt_ map[string]interface{}
}

// Get: Gets the latest state of a long-running operation.  Clients can
// use this
// method to poll the operation result at intervals as
// recommended by the API
// service.
func (r *ProjectsOperationsService) Get(name string) *ProjectsOperationsGetCall {
	c := &ProjectsOperationsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsOperationsGetCall) Fields(s ...googleapi.Field) *ProjectsOperationsGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsOperationsGetCall) Do() (*Operation, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/{+name}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
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
	//   "description": "Gets the latest state of a long-running operation.  Clients can use this\nmethod to poll the operation result at intervals as recommended by the API\nservice.",
	//   "flatPath": "v1beta3/projects/{projectsId}/operations/{operationsId}",
	//   "httpMethod": "GET",
	//   "id": "datastore.projects.operations.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The name of the operation resource.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]*/operations/[^/]*$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/{+name}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/datastore"
	//   ]
	// }

}

// method id "datastore.projects.operations.list":

type ProjectsOperationsListCall struct {
	s    *Service
	name string
	opt_ map[string]interface{}
}

// List: Lists operations that match the specified filter in the
// request. If the
// server doesn't support this method, it returns
// `UNIMPLEMENTED`.
//
// NOTE: the `name` binding below allows API services
// to override the binding
// to use different resource name schemes, such
// as `users/*/operations`.
func (r *ProjectsOperationsService) List(name string) *ProjectsOperationsListCall {
	c := &ProjectsOperationsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.name = name
	return c
}

// Filter sets the optional parameter "filter": The standard list
// filter.
func (c *ProjectsOperationsListCall) Filter(filter string) *ProjectsOperationsListCall {
	c.opt_["filter"] = filter
	return c
}

// PageSize sets the optional parameter "pageSize": The standard list
// page size.
func (c *ProjectsOperationsListCall) PageSize(pageSize int64) *ProjectsOperationsListCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken": The standard list
// page token.
func (c *ProjectsOperationsListCall) PageToken(pageToken string) *ProjectsOperationsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsOperationsListCall) Fields(s ...googleapi.Field) *ProjectsOperationsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsOperationsListCall) Do() (*ListOperationsResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["filter"]; ok {
		params.Set("filter", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageSize"]; ok {
		params.Set("pageSize", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/{+name}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
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
	//   "description": "Lists operations that match the specified filter in the request. If the\nserver doesn't support this method, it returns `UNIMPLEMENTED`.\n\nNOTE: the `name` binding below allows API services to override the binding\nto use different resource name schemes, such as `users/*/operations`.",
	//   "flatPath": "v1beta3/projects/{projectsId}/operations",
	//   "httpMethod": "GET",
	//   "id": "datastore.projects.operations.list",
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
	//       "pattern": "^projects/[^/]*/operations$",
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
	//   "path": "v1beta3/{+name}",
	//   "response": {
	//     "$ref": "ListOperationsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/datastore"
	//   ]
	// }

}
