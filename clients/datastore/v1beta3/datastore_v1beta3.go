// Package datastore provides access to the Google Cloud Datastore API - NEW.
//
// See https://cloud.google.com/datastore/
//
// Usage example:
//
//   import "github.com/skelterjohn/gcloud_apis/clients/datastore/v1beta3"
//   ...
//   datastoreService, err := datastore.New(oauthHttpClient)
package datastore

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
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Projects *ProjectsService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
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

// AllocateIdsRequest: The request for
// google.datastore.v1beta3.Datastore.AllocateIds.
type AllocateIdsRequest struct {
	// DatabaseId: If not empty, the ID of the database against which to
	// make the request.
	DatabaseId string `json:"databaseId,omitempty"`

	// Keys: A list of keys with incomplete key paths for which to allocate
	// IDs.
	// No key may be reserved/read-only.
	Keys []*Key `json:"keys,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DatabaseId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *AllocateIdsRequest) MarshalJSON() ([]byte, error) {
	type noMethod AllocateIdsRequest
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// AllocateIdsResponse: The response for
// google.datastore.v1beta3.Datastore.AllocateIds.
type AllocateIdsResponse struct {
	// Keys: The keys specified in the request (in the same order), each
	// with
	// its key path completed with a newly allocated ID.
	Keys []*Key `json:"keys,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Keys") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *AllocateIdsResponse) MarshalJSON() ([]byte, error) {
	type noMethod AllocateIdsResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// ArrayValue: An array value.
type ArrayValue struct {
	// Values: Values in the array.
	// The order of this array may not be preserved if it contains a mix
	// of
	// indexed and unindexed values.
	Values []*Value `json:"values,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Values") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *ArrayValue) MarshalJSON() ([]byte, error) {
	type noMethod ArrayValue
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// BeginTransactionRequest: The request for
// google.datastore.v1beta3.Datastore.BeginTransaction.
type BeginTransactionRequest struct {
	// DatabaseId: If not empty, the ID of the database against which to
	// make the request.
	DatabaseId string `json:"databaseId,omitempty"`

	// TransactionOptions: Options for a new transaction.
	TransactionOptions *TransactionOptions `json:"transactionOptions,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DatabaseId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *BeginTransactionRequest) MarshalJSON() ([]byte, error) {
	type noMethod BeginTransactionRequest
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// BeginTransactionResponse: The response for
// google.datastore.v1beta3.Datastore.BeginTransaction.
type BeginTransactionResponse struct {
	// Transaction: The transaction identifier (always present).
	Transaction string `json:"transaction,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Transaction") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *BeginTransactionResponse) MarshalJSON() ([]byte, error) {
	type noMethod BeginTransactionResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// Change: A change to a set of entities being watched.
// Usually a change to an individual entity,
// but sometimes a change to the set of entities being watched.
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
	//
	// Possible values:
	//   "NULL_VALUE" - Null value.
	NoChange string `json:"noChange,omitempty"`

	// ResumeToken: A token that provides a compact representation of all
	// the changes that
	// have been received by the caller up to this point (including this
	// one)
	// that can be used to resume the stream. May not be set on every
	// change.
	// Only valid for the targets specified in `target_ids`.
	ResumeToken string `json:"resumeToken,omitempty"`

	// TargetChange: The targets associate with this stream have been
	// modified.
	// The affected targets are listed in `target_ids`.
	//
	// Possible values:
	//   "TARGET_CHANGE_UNSPECIFIED"
	//   "TARGET_ADDED"
	//   "TARGET_REMOVED"
	//   "TARGET_REJECTED"
	TargetChange string `json:"targetChange,omitempty"`

	// TargetIds: The set of targets to which this change applies. When
	// empty, the change
	// applies to all targets.
	TargetIds []int64 `json:"targetIds,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Cause") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Change) MarshalJSON() ([]byte, error) {
	type noMethod Change
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// ChangeBatch: The streamed response for
// google.datastore.v1beta3.DatastoreWatcher.Watch
// and google.datastore.v1beta3.DatastoreWatcher.MultiWatch.
type ChangeBatch struct {
	Changes []*Change `json:"changes,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Changes") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *ChangeBatch) MarshalJSON() ([]byte, error) {
	type noMethod ChangeBatch
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// Circle: A "circle" on the surface of the Earth, defined as the area
// within a
// certain geographical distance of a point.
type Circle struct {
	// Center: The center of the circle.
	Center *LatLng `json:"center,omitempty"`

	// RadiusMeters: The "radius" of the circle, in meters.
	// Must be greater or equal to zero.
	RadiusMeters float64 `json:"radiusMeters,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Center") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Circle) MarshalJSON() ([]byte, error) {
	type noMethod Circle
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// CommitRequest: The request for
// google.datastore.v1beta3.Datastore.Commit.
type CommitRequest struct {
	// DatabaseId: If not empty, the ID of the database against which to
	// make the request.
	DatabaseId string `json:"databaseId,omitempty"`

	// Mode: The type of commit to perform. Defaults to `TRANSACTIONAL`.
	//
	// Possible values:
	//   "MODE_UNSPECIFIED" - Unspecified. This value must not be used.
	//   "TRANSACTIONAL" - Transactional: The mutations are either all
	// applied, or none are applied.
	// Learn about transactions
	// [here](https://cloud.google.com/datastore/docs/concepts/transactions).
	//   "NON_TRANSACTIONAL" - Non-transactional: The mutations may not
	// apply as all or none.
	Mode string `json:"mode,omitempty"`

	// Mutations: The mutations to perform.
	//
	// When mode is `TRANSACTIONAL`, mutations affecting a single entity
	// are
	// applied in order. The following sequences of mutations affecting a
	// single
	// entity are not permitted in a single `Commit` request:
	//
	// - `insert` followed by `insert`
	// - `update` followed by `insert`
	// - `upsert` followed by `insert`
	// - `delete` followed by `update`
	//
	// When mode is `NON_TRANSACTIONAL`, no two mutations may affect a
	// single
	// entity.
	Mutations []*Mutation `json:"mutations,omitempty"`

	// SingleUseTransaction: Options for beginning a new transaction for
	// this request.
	// The transaction is committed when the request completes. If
	// specified,
	// google.datastore.v1beta3.TransactionOptions.mode must
	// be
	// google.datastore.v1beta3.TransactionOptions.ReadWrite.
	SingleUseTransaction *TransactionOptions `json:"singleUseTransaction,omitempty"`

	// Transaction: The identifier of the transaction associated with the
	// commit. A
	// transaction identifier is returned by a call to
	// BeginTransaction.
	Transaction string `json:"transaction,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DatabaseId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *CommitRequest) MarshalJSON() ([]byte, error) {
	type noMethod CommitRequest
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// CommitResponse: The response for
// google.datastore.v1beta3.Datastore.Commit.
type CommitResponse struct {
	// IndexUpdates: The number of index entries updated during the commit,
	// or zero if none were
	// updated.
	IndexUpdates int64 `json:"indexUpdates,omitempty"`

	// MutationResults: The result of performing the mutations.
	// The i-th mutation result corresponds to the i-th mutation in the
	// request.
	MutationResults []*MutationResult `json:"mutationResults,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "IndexUpdates") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *CommitResponse) MarshalJSON() ([]byte, error) {
	type noMethod CommitResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// CommonMetadata: Metadata common to all Datastore Admin Operations.
type CommonMetadata struct {
	// Cancelling: True if the Operation is currently being cancelled.  Will
	// only be set after
	// the cancellation request was received, but before the Operation is
	// done.
	// So the state transitions are:
	// 1) done=false, cancelling=false
	// 2) done=false, cancelling=true
	// 3) done=true, cancelling=false
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
	//
	// Possible values:
	//   "UNSPECIFIED_TYPE" - Unspecified.
	//   "EXPORT" - Export.
	//   "IMPORT" - Import.
	//   "BUILD_INDEX" - Build an index.
	//   "CLEAR_INDEX" - Clear an index.
	OperationType string `json:"operationType,omitempty"`

	// StartTime: The time that work began on the Operation.
	StartTime string `json:"startTime,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Cancelling") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *CommonMetadata) MarshalJSON() ([]byte, error) {
	type noMethod CommonMetadata
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// CompositeFilter: A filter that merges multiple other filters using
// the given operator.
type CompositeFilter struct {
	// Filters: The list of filters to combine.
	// Must contain at least one filter.
	Filters []*Filter `json:"filters,omitempty"`

	// Op: The operator for combining multiple filters.
	//
	// Possible values:
	//   "OPERATOR_UNSPECIFIED" - Unspecified. This value must not be used.
	//   "AND" - The results are required to satisfy each of the combined
	// filters.
	Op string `json:"op,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Filters") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *CompositeFilter) MarshalJSON() ([]byte, error) {
	type noMethod CompositeFilter
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

// Entity: A Datastore data object.
//
// An entity is limited to 1 megabyte when stored. That
// _roughly_
// corresponds to a limit of 1 megabyte for the serialized form of
// this
// message.
type Entity struct {
	// Key: The entity's key.
	//
	// An entity must have a key, unless otherwise documented (for
	// example,
	// an entity in `Value.entity_value` may have no key).
	// An entity's kind is its key path's last element's kind,
	// or null if it has no key.
	Key *Key `json:"key,omitempty"`

	// Properties: The entity's properties.
	// The map's keys are property names.
	// A property name matching regex `__.*__` is reserved.
	// A reserved property name is forbidden in certain documented
	// contexts.
	// The name must not contain more than 500 characters.
	// The name cannot be "".
	Properties map[string]Value `json:"properties,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Key") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Entity) MarshalJSON() ([]byte, error) {
	type noMethod Entity
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// EntityFilter: Identifies a subset of Entities in a Project.  This is
// specified as
// combinations of Kind + Namespace (either or both of which may be all
// as
// described in the following examples).
// Example usage:
//
// Entire Project:
//   kinds=[], namespace_ids=[]
//
// Kinds Foo and Bar in all Namespaces:
//   kinds=['Foo', 'Bar'], namespace_ids=[]
//
// Kinds Foo and Bar only in the Default Namespace:
//   kinds=['Foo', 'Bar'], namespace_ids=['']
//
// Kinds Foo and Bar in both the Default and Baz Namespaces:
//   kinds=['Foo', 'Bar'], namespace_ids=['', 'Baz']
//
// The entire Baz Namespace:
//   kinds=[], namespace_ids=['Baz']
type EntityFilter struct {
	// Kinds: If empty, then this represents all Kinds.
	Kinds []string `json:"kinds,omitempty"`

	// NamespaceIds: An empty list represents all Namespaces.  This is the
	// preferred
	// usage for Projects that don't use Namespaces.
	//
	// An empty string element represents the Default Namespace.  This
	// should be
	// used if the Project has data in non-Default Namespaces, but doesn't
	// want to
	// include them.
	// Each Namespace in this list must be unique.
	NamespaceIds []string `json:"namespaceIds,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Kinds") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *EntityFilter) MarshalJSON() ([]byte, error) {
	type noMethod EntityFilter
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// EntityResult: The result of fetching an entity from Datastore.
type EntityResult struct {
	// Cursor: A cursor that points to the position after the result
	// entity.
	// Set only when the `EntityResult` is part of a `QueryResultBatch`
	// message.
	Cursor string `json:"cursor,omitempty"`

	// Entity: The resulting entity.
	Entity *Entity `json:"entity,omitempty"`

	// Version: The version of the entity, a strictly positive number that
	// monotonically
	// increases with changes to the entity.
	//
	// This field is set for `FULL` entity results.
	// For missing entities in
	// `LookupResponse`, this is the version of the snapshot that was used
	// to look
	// up the entity, and it is always set except for eventually consistent
	// reads.
	Version int64 `json:"version,omitempty,string"`

	// ForceSendFields is a list of field names (e.g. "Cursor") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *EntityResult) MarshalJSON() ([]byte, error) {
	type noMethod EntityResult
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// ExistenceFilter: An existence filter that must be applied and
// verified locally to resolve
// possible delete mutations.
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
	//
	// Possible values:
	//   "STRATEGY_NOT_SPECIFIED"
	//   "MURMUR128_MITZ_64"
	Strategy string `json:"strategy,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Bits") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *ExistenceFilter) MarshalJSON() ([]byte, error) {
	type noMethod ExistenceFilter
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// ExportMetadata: Metadata for Export Operations.
type ExportMetadata struct {
	// BytesProgress: An estimate of the number of bytes processed.
	BytesProgress *Progress `json:"bytesProgress,omitempty"`

	// Common: Metadata common to all Datastore Admin Operations.
	Common *CommonMetadata `json:"common,omitempty"`

	// Destination: Location for the export metadata and data files.  This
	// will be the same
	// value as the google.datastore.v1beta3.ExportRequest.destination
	// field.
	// Actual files will be nested deeper than this.  The final output
	// location
	// is provided in google.datastore.v1beta3.ExportResponse.data_location.
	Destination string `json:"destination,omitempty"`

	// EntitiesProgress: An estimate of the number of Entities processed.
	EntitiesProgress *Progress `json:"entitiesProgress,omitempty"`

	// EntityFilter: Description of which Entities are being exported.
	EntityFilter *EntityFilter `json:"entityFilter,omitempty"`

	// ForceSendFields is a list of field names (e.g. "BytesProgress") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *ExportMetadata) MarshalJSON() ([]byte, error) {
	type noMethod ExportMetadata
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// ExportRequest: The request for
// google.datastore.v1beta3.DatastoreAdmin.Export.
type ExportRequest struct {
	// DatabaseId: Database ID against which to make the request.
	// Unset indicates the default database.
	// Not currently supported.
	DatabaseId string `json:"databaseId,omitempty"`

	// Destination: Location for the export metadata and data
	// files.
	//
	// Specified as the full resource name of the external storage
	// location.
	// Currently, only Google Cloud Storage is supported.  So the
	// destination
	// should be of the form: //storage.googleapis.com/buckets/bucket-name
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
	// By nesting the data files deeper, the same destination can be used
	// in
	// multiple Export Operations without conflict.
	Destination string `json:"destination,omitempty"`

	// EntityFilter: Description of what data from the Project is included
	// in the export.
	//
	EntityFilter *EntityFilter `json:"entityFilter,omitempty"`

	// Labels: Client-assigned labels.
	Labels map[string]string `json:"labels,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DatabaseId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *ExportRequest) MarshalJSON() ([]byte, error) {
	type noMethod ExportRequest
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// ExportResponse: The response for
// google.datastore.v1beta3.DatastoreAdmin.Export.
type ExportResponse struct {
	// DataLocation: Location of the output files. This can be used to begin
	// an import into
	// Cloud Datastore (this Project or another Project).
	// See
	// google.datastore.v1beta3.ImportRequest.data_location.  Only present
	// if
	// the Operation completed successfully.
	//
	// This location will be nested deeper than the
	// initial
	// google.datastore.v1beta3.ExportRequest.destination in order to
	// support
	// multiple Exports to the same destination without conflict.
	DataLocation string `json:"dataLocation,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DataLocation") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *ExportResponse) MarshalJSON() ([]byte, error) {
	type noMethod ExportResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// Filter: A holder for any type of filter.
type Filter struct {
	// CompositeFilter: A composite filter.
	CompositeFilter *CompositeFilter `json:"compositeFilter,omitempty"`

	// PropertyFilter: A filter on a property.
	PropertyFilter *PropertyFilter `json:"propertyFilter,omitempty"`

	// StContainsFilter: A filter that selects geo points within a region.
	StContainsFilter *StContainsFilter `json:"stContainsFilter,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CompositeFilter") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Filter) MarshalJSON() ([]byte, error) {
	type noMethod Filter
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// GeoRegion: A region on the surface of the Earth.
type GeoRegion struct {
	// Circle: A circular region.
	Circle *Circle `json:"circle,omitempty"`

	// Rectangle: A rectangular region.
	Rectangle *Rectangle `json:"rectangle,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Circle") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *GeoRegion) MarshalJSON() ([]byte, error) {
	type noMethod GeoRegion
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// GqlQuery: A [GQL
// query](https://cloud.google.com/datastore/docs/apis/gql/gql_reference)
// .
type GqlQuery struct {
	// AllowLiterals: When false, the query string must not contain any
	// literals and instead
	// must bind all values. For example,
	// `SELECT * FROM Kind WHERE a = 'string literal'` is not allowed,
	// while
	// `SELECT * FROM Kind WHERE a = @value` is.
	AllowLiterals bool `json:"allowLiterals,omitempty"`

	// NamedBindings: For each non-reserved named binding site in the query
	// string,
	// there must be a named parameter with that name,
	// but not necessarily the inverse.
	// Key must match regex `A-Za-z_$*`, must not match regex
	// `__.*__`, and must not be "".
	NamedBindings map[string]GqlQueryParameter `json:"namedBindings,omitempty"`

	// PositionalBindings: Numbered binding site @1 references the first
	// numbered parameter,
	// effectively using 1-based indexing, rather than the usual 0.
	// For each binding site numbered i in `query_string`,
	// there must be an i-th numbered parameter.
	// The inverse must also be true.
	PositionalBindings []*GqlQueryParameter `json:"positionalBindings,omitempty"`

	// QueryString: A string of the format
	// described
	// [here](https://cloud.google.com/datastore/docs/apis/gql/gql_
	// reference).
	QueryString string `json:"queryString,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AllowLiterals") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *GqlQuery) MarshalJSON() ([]byte, error) {
	type noMethod GqlQuery
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// GqlQueryParameter: A binding parameter for a GQL query.
type GqlQueryParameter struct {
	// Cursor: A query cursor. Query cursors are returned in query
	// result batches.
	Cursor string `json:"cursor,omitempty"`

	// GeoRegion: Geographical region.
	GeoRegion *GeoRegion `json:"geoRegion,omitempty"`

	// Value: A value parameter.
	Value *Value `json:"value,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Cursor") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *GqlQueryParameter) MarshalJSON() ([]byte, error) {
	type noMethod GqlQueryParameter
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// ImportMetadata: Metadata for Import operations.
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

	// ForceSendFields is a list of field names (e.g. "BytesProgress") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *ImportMetadata) MarshalJSON() ([]byte, error) {
	type noMethod ImportMetadata
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// ImportRequest: The request for
// google.datastore.v1beta3.DatastoreAdmin.Import.
type ImportRequest struct {
	// DataLocation: The full resource name of the external storage
	// location.  Currently, only
	// Google Cloud Storage is supported.  So the data_location should be of
	// the
	// form:
	// //storage.googleapis.com/buckets/bucket-name/objects/object-path.
	//
	// See
	//  google.datastore.v1beta3.ExportResponse.data_location
	DataLocation string `json:"dataLocation,omitempty"`

	// DatabaseId: Database ID against which to make the request.
	// Unset indicates the default database.
	// Not currently supported.
	DatabaseId string `json:"databaseId,omitempty"`

	// EntityFilter: Optionally specify which Kinds/Namespaces are to be
	// imported. If provided,
	// the list must be a subset of the EntityFilter used in creating the
	// export,
	// otherwise a FAILED_PRECONDITION error will be returned. If no filter
	// is
	// specified then all Entities from the export are imported.
	EntityFilter *EntityFilter `json:"entityFilter,omitempty"`

	// Labels: Client-assigned labels.
	Labels map[string]string `json:"labels,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DataLocation") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *ImportRequest) MarshalJSON() ([]byte, error) {
	type noMethod ImportRequest
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// Index: An index.
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
	//
	// Possible values:
	//   "STATE_UNSPECIFIED" - The state is unspecified.
	//   "OFF" - No index data exists.
	//   "SERVING" - The index is updated when writing an entity.
	// The index is fully populated from all relevant stored entities.
	//   "BUILDING" - The index is being built: transitioning from OFF to
	// SERVING.
	// There is an active long-running operation for the index.
	// The index is updated when writing an entity.
	// Some index data may exist.
	//   "CLEARING" - The index is being cleared: transitioning from SERVING
	// to OFF.
	// There is an active long-running operation for the index.
	// The index is not updated when writing an entity.
	// Some index data may exist.
	//   "ERROR" - The index was being built or cleared, but something went
	// wrong.
	// There is no active long-running operation for the index,
	// and the most recently finished long-running operation failed.
	// The index is not updated when writing an entity.
	// Some index data may exist.
	State string `json:"state,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "DatabaseId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Index) MarshalJSON() ([]byte, error) {
	type noMethod Index
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// IndexDefinition: An index definition.
type IndexDefinition struct {
	// Kind: The kind of entity to index.
	// This field shares the constraints of field Key.PathElement.kind.
	Kind string `json:"kind,omitempty"`

	// Properties: A sequence of indexed property definitions.
	Properties []*IndexedPropertyDefinition `json:"properties,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Kind") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *IndexDefinition) MarshalJSON() ([]byte, error) {
	type noMethod IndexDefinition
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// IndexedPropertyDefinition: An indexed property definition.
type IndexedPropertyDefinition struct {
	// Mode: The indexed property's mode.
	//
	// Possible values:
	//   "MODE_UNSPECIFIED" - The mode is unspecified.
	//   "UNORDERED" - The property's values are indexed so as to support
	// query by equality.
	//   "ASCENDING" - The property's values are indexed so as to support
	// sequencing in
	// ascending order and also query by <, >, <=, >=, and =.
	//   "DESCENDING" - The property's values are indexed so as to support
	// sequencing in
	// descending order and also query by <, >, <=, >=, and =.
	//   "KEY_ANCESTRY" - The property's key values are indexed by ancestor
	// key,
	// supporting query by key ancestry relationship.
	// Currently supported only for property "__key__".
	// For example, for an entity with key path
	// /Blog:7/Post:11/Comment:5
	// the indexed ancestor key paths are /Blog:7, /Blog:7/Post:11,
	// and
	// /Blog:7/Post:11/Comment:5.
	// (The partition ID of all the ancestor keys are the same.)
	//   "GEO" - The property's values are indexed by geographic
	// location,
	// supporting query by location.
	// (Only geo point values are indexed; others are ignored.)
	Mode string `json:"mode,omitempty"`

	// Name: The indexed property's name.
	// This field shares the constraints of the keys of the map in
	// field Entity.properties.
	// If name includes "."s, it may be interpreted as a property name
	// path.
	// Required.
	Name string `json:"name,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Mode") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *IndexedPropertyDefinition) MarshalJSON() ([]byte, error) {
	type noMethod IndexedPropertyDefinition
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// Key: A unique identifier for an entity.
// If a key's partition ID or any of its path kinds or names
// are
// reserved/read-only, the key is reserved/read-only.
// A reserved/read-only key is forbidden in certain documented contexts.
type Key struct {
	// PartitionId: Entities are partitioned into subsets, currently
	// identified by a dataset
	// (usually implicitly specified by the project) and namespace
	// ID.
	// Queries are scoped to a single partition.
	PartitionId *PartitionId `json:"partitionId,omitempty"`

	// Path: The entity path.
	// An entity path consists of one or more elements composed of a kind
	// and a
	// string or numerical identifier, which identify entities. The
	// first
	// element identifies a _root entity_, the second element identifies
	// a _child_ of the root entity, the third element identifies a child of
	// the
	// second entity, and so forth. The entities identified by all prefixes
	// of
	// the path are called the element's _ancestors_.
	//
	// An entity path is always fully complete: *all* of the entity's
	// ancestors
	// are required to be in the path along with the entity identifier
	// itself.
	// The only exception is that in some documented cases, the identifier
	// in the
	// last path element (for the entity) itself may be omitted. For
	// example,
	// an entity in `Value.entity_value` may have no key.
	//
	// A path can never be empty, and a path can have at most 100 elements.
	Path []*PathElement `json:"path,omitempty"`

	// ForceSendFields is a list of field names (e.g. "PartitionId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Key) MarshalJSON() ([]byte, error) {
	type noMethod Key
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// KindExpression: A representation of a kind.
type KindExpression struct {
	// Name: The name of the kind.
	Name string `json:"name,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Name") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *KindExpression) MarshalJSON() ([]byte, error) {
	type noMethod KindExpression
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// LatLng: An object representing a latitude/longitude pair. This is
// expressed as a pair
// of doubles representing degrees latitude and degrees longitude.
// Unless
// specified otherwise, this must conform to the
// <a
// href="http://www.unoosa.org/pdf/icg/2012/template/WGS_84.pdf">WGS84
// st
// andard</a>. Values must be within normalized ranges.
//
// Example of normalization code in Python:
//
//     def NormalizeLongitude(longitude):
//       """Wraps decimal degrees longitude to [-180.0, 180.0]."""
//       q, r = divmod(longitude, 360.0)
//       if r > 180.0 or (r == 180.0 and q <= -1.0):
//         return r - 360.0
//       return r
//
//     def NormalizeLatLng(latitude, longitude):
//       """Wraps decimal degrees latitude and longitude to
//       [-180.0, 180.0] and [-90.0, 90.0], respectively."""
//       r = latitude % 360.0
//       if r <= 90.0:
//         return r, NormalizeLongitude(longitude)
//       elif r >= 270.0:
//         return r - 360, NormalizeLongitude(longitude)
//       else:
//         return 180 - r, NormalizeLongitude(longitude + 180.0)
//
//     assert 180.0 == NormalizeLongitude(180.0)
//     assert -180.0 == NormalizeLongitude(-180.0)
//     assert -179.0 == NormalizeLongitude(181.0)
//     assert (0.0, 0.0) == NormalizeLatLng(360.0, 0.0)
//     assert (0.0, 0.0) == NormalizeLatLng(-360.0, 0.0)
//     assert (85.0, 180.0) == NormalizeLatLng(95.0, 0.0)
//     assert (-85.0, -170.0) == NormalizeLatLng(-95.0, 10.0)
//     assert (90.0, 10.0) == NormalizeLatLng(90.0, 10.0)
//     assert (-90.0, -10.0) == NormalizeLatLng(-90.0, -10.0)
//     assert (0.0, -170.0) == NormalizeLatLng(-180.0, 10.0)
//     assert (0.0, -170.0) == NormalizeLatLng(180.0, 10.0)
//     assert (-90.0, 10.0) == NormalizeLatLng(270.0, 10.0)
//     assert (90.0, 10.0) == NormalizeLatLng(-270.0, 10.0)
type LatLng struct {
	// Latitude: The latitude in degrees. It must be in the range [-90.0,
	// +90.0].
	Latitude float64 `json:"latitude,omitempty"`

	// Longitude: The longitude in degrees. It must be in the range [-180.0,
	// +180.0].
	Longitude float64 `json:"longitude,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Latitude") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *LatLng) MarshalJSON() ([]byte, error) {
	type noMethod LatLng
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// ListIndexesResponse: The response for
// google.datastore.v1beta3.DatastoreAdmin.ListIndexes.
type ListIndexesResponse struct {
	// Indexes: The indexes.
	Indexes []*Index `json:"indexes,omitempty"`

	// NextPageToken: The standard List next-page token.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Indexes") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *ListIndexesResponse) MarshalJSON() ([]byte, error) {
	type noMethod ListIndexesResponse
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

// LookupIndexRequest: The request for
// google.datastore.v1beta3.DatastoreAdmin.LookupIndex.
type LookupIndexRequest struct {
	// DatabaseId: The database id.
	DatabaseId string `json:"databaseId,omitempty"`

	// IndexDefinition: The index definition.
	IndexDefinition *IndexDefinition `json:"indexDefinition,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DatabaseId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *LookupIndexRequest) MarshalJSON() ([]byte, error) {
	type noMethod LookupIndexRequest
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// LookupRequest: The request for
// google.datastore.v1beta3.Datastore.Lookup.
type LookupRequest struct {
	// DatabaseId: If not empty, the ID of the database against which to
	// make the request.
	DatabaseId string `json:"databaseId,omitempty"`

	// Keys: Keys of entities to look up.
	Keys []*Key `json:"keys,omitempty"`

	// PropertyMask: The properties to return. Defaults to returning all
	// properties.
	// If this field is set and an entity has a property not referenced in
	// the
	// mask, it will not be included
	// in
	// google.datastore.v1beta3.LookupResponse.found.entity.properties.
	// Th
	// e entity's key is always returned.
	// If an google.datastore.v1beta3.Value.entity_value is
	// returned, its key is returned as well.
	//
	// The paths in the mask are property paths: dot (`.`) separated
	// property
	// names that can be used to reference properties nested
	// in
	// google.datastore.v1beta3.Value.entity_value.
	// A property must not reference a value inside
	// an
	// google.datastore.v1beta3.Value.array_value.
	// None of these property paths may contain the `__key__` property name.
	PropertyMask string `json:"propertyMask,omitempty"`

	// ReadOptions: The options for this lookup request.
	ReadOptions *ReadOptions `json:"readOptions,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DatabaseId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *LookupRequest) MarshalJSON() ([]byte, error) {
	type noMethod LookupRequest
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// LookupResponse: The response for
// google.datastore.v1beta3.Datastore.Lookup.
type LookupResponse struct {
	// Deferred: A list of keys that were not looked up due to resource
	// constraints. The
	// order of results in this field is undefined and has no relation to
	// the
	// order of the keys in the input.
	Deferred []*Key `json:"deferred,omitempty"`

	// Found: Entities found as `ResultType.FULL` entities. The order of
	// results in this
	// field is undefined and has no relation to the order of the keys in
	// the
	// input.
	Found []*EntityResult `json:"found,omitempty"`

	// Missing: Entities not found as `ResultType.KEY_ONLY` entities. The
	// order of results
	// in this field is undefined and has no relation to the order of the
	// keys
	// in the input.
	Missing []*EntityResult `json:"missing,omitempty"`

	// Transaction: The transaction that was started as part of this Lookup
	// request.
	// Set only when
	// google.datastore.v1beta3.ReadOptions.begin_transaction
	// was set in google.datastore.v1beta3.LookupRequest.read_options.
	Transaction string `json:"transaction,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Deferred") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *LookupResponse) MarshalJSON() ([]byte, error) {
	type noMethod LookupResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// MultiWatchRequest: A request for
// google.datastore.v1beta3.DatastoreWatcher.MultiWatch
type MultiWatchRequest struct {
	// AddTargets: The set of watch targets to add to this stream.
	// changes will be returned with server
	// assigned `target_ids` in the same order as targets are specified
	// here.
	AddTargets []*WatchTarget `json:"addTargets,omitempty"`

	// DatabaseId: Database ID against which to make the request.
	DatabaseId string `json:"databaseId,omitempty"`

	// RemoveTargets: The IDs of watch targets to remove from this stream.
	RemoveTargets []int64 `json:"removeTargets,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AddTargets") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *MultiWatchRequest) MarshalJSON() ([]byte, error) {
	type noMethod MultiWatchRequest
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// Mutation: A mutation to apply to an entity.
type Mutation struct {
	// BaseVersion: The version of the entity that this mutation is being
	// applied to. If this
	// does not match the current version on the server, the mutation
	// conflicts.
	BaseVersion int64 `json:"baseVersion,omitempty,string"`

	// Delete: The key of the entity to delete. The entity may or may not
	// already exist.
	// Must have a complete key path and must not be reserved/read-only.
	Delete *Key `json:"delete,omitempty"`

	// Insert: The entity to insert. The entity must not already exist.
	// The entity key's final path element may be incomplete.
	Insert *Entity `json:"insert,omitempty"`

	// PropertyMask: The properties to write in this mutation.
	// This field is ignored for `delete`.
	//
	// If the entity already exists, only properties referenced in the mask
	// are
	// updated, others are left untouched.
	// Properties referenced in the mask but not in the entity are
	// deleted.
	// Properties not referenced in the mask may not be set in the
	// entity.
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
	// Must have a complete key path.
	Update *Entity `json:"update,omitempty"`

	// Upsert: The entity to upsert. The entity may or may not already
	// exist.
	// The entity key's final path element may be incomplete.
	Upsert *Entity `json:"upsert,omitempty"`

	// ForceSendFields is a list of field names (e.g. "BaseVersion") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Mutation) MarshalJSON() ([]byte, error) {
	type noMethod Mutation
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// MutationResult: The result of applying a mutation.
type MutationResult struct {
	// ConflictDetected: Whether a conflict was detected for this mutation.
	// Always false when a
	// conflict detection strategy field is not set in the mutation.
	ConflictDetected bool `json:"conflictDetected,omitempty"`

	// Key: The automatically allocated key.
	// Set only when the mutation allocated a key.
	Key *Key `json:"key,omitempty"`

	// Version: The version of the entity on the server after processing the
	// mutation. If
	// the mutation doesn't change anything on the server, then the version
	// will
	// be the version of the current entity or, if no entity is present, a
	// version
	// that is strictly greater than the version of any previous entity and
	// less
	// than the version of any possible future entity.
	Version int64 `json:"version,omitempty,string"`

	// ForceSendFields is a list of field names (e.g. "ConflictDetected") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *MutationResult) MarshalJSON() ([]byte, error) {
	type noMethod MutationResult
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

	// Error: The error result of the operation in case of failure.
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
	// originally returns it. If you use the default HTTP mapping above,
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

// PartitionId: A partition ID identifies a grouping of entities. The
// grouping is always
// by project and namespace, however the namespace ID may be empty.
//
// A partition ID contains several dimensions:
// project ID and namespace ID.
//
// Partition dimensions:
//
// - May be "".
// - Must be valid UTF-8 bytes.
// - Must have values that match regex `[A-Za-z\d\.\-_]{1,100}`
// If the value of any dimension matches regex `__.*__`, the partition
// is
// reserved/read-only.
// A reserved/read-only partition ID is forbidden in certain
// documented
// contexts.
//
// Foreign partition IDs (in which the project ID does
// not match the context project ID ) are discouraged.
// Reads and writes of foreign partition IDs may fail if the project is
// not in an active state.
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

	// ForceSendFields is a list of field names (e.g. "DatabaseId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *PartitionId) MarshalJSON() ([]byte, error) {
	type noMethod PartitionId
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// PathElement: A (kind, ID/name) pair used to construct a key path.
//
// If either name or ID is set, the element is complete.
// If neither is set, the element is incomplete.
type PathElement struct {
	// Id: The auto-allocated ID of the entity.
	// Never equal to zero. Values less than zero are discouraged and may
	// not
	// be supported in the future.
	Id int64 `json:"id,omitempty,string"`

	// Kind: The kind of the entity.
	// A kind matching regex `__.*__` is reserved/read-only.
	// A kind must not contain more than 1500 bytes when UTF-8
	// encoded.
	// Cannot be "".
	Kind string `json:"kind,omitempty"`

	// Name: The name of the entity.
	// A name matching regex `__.*__` is reserved/read-only.
	// A name must not be more than 1500 bytes when UTF-8 encoded.
	// Cannot be "".
	Name string `json:"name,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Id") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *PathElement) MarshalJSON() ([]byte, error) {
	type noMethod PathElement
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// Progress: Measures the progress of a particular metric.
type Progress struct {
	// WorkCompleted: Note that this may be greater than work_estimated.
	WorkCompleted int64 `json:"workCompleted,omitempty,string"`

	// WorkEstimated: An estimate of how much work needs to be performed.
	// May be zero if the
	// work estimate is unavailable.
	WorkEstimated int64 `json:"workEstimated,omitempty,string"`

	// ForceSendFields is a list of field names (e.g. "WorkCompleted") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Progress) MarshalJSON() ([]byte, error) {
	type noMethod Progress
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// Projection: A representation of a property in a projection.
type Projection struct {
	// Property: The property to project.
	Property *PropertyReference `json:"property,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Property") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Projection) MarshalJSON() ([]byte, error) {
	type noMethod Projection
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// PropertyFilter: A filter on a specific property.
type PropertyFilter struct {
	// Op: The operator to filter by.
	//
	// Possible values:
	//   "OPERATOR_UNSPECIFIED" - Unspecified. This value must not be used.
	//   "LESS_THAN" - Less than.
	//   "LESS_THAN_OR_EQUAL" - Less than or equal.
	//   "GREATER_THAN" - Greater than.
	//   "GREATER_THAN_OR_EQUAL" - Greater than or equal.
	//   "EQUAL" - Equal.
	//   "HAS_ANCESTOR" - Has ancestor.
	Op string `json:"op,omitempty"`

	// Property: The property to filter by.
	Property *PropertyReference `json:"property,omitempty"`

	// Value: The value to compare the property to.
	Value *Value `json:"value,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Op") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *PropertyFilter) MarshalJSON() ([]byte, error) {
	type noMethod PropertyFilter
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// PropertyOrder: The desired order for a specific property.
type PropertyOrder struct {
	// Direction: The direction to order by. Defaults to `ASCENDING`.
	//
	// Possible values:
	//   "DIRECTION_UNSPECIFIED" - Unspecified. This value must not be used.
	//   "ASCENDING" - Ascending.
	//   "DESCENDING" - Descending.
	Direction string `json:"direction,omitempty"`

	// Property: The property to order by.
	Property *PropertyReference `json:"property,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Direction") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *PropertyOrder) MarshalJSON() ([]byte, error) {
	type noMethod PropertyOrder
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// PropertyReference: A reference to a property relative to the kind
// expressions.
type PropertyReference struct {
	// Name: The name of the property.
	Name string `json:"name,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Name") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *PropertyReference) MarshalJSON() ([]byte, error) {
	type noMethod PropertyReference
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// Query: A query for entities.
type Query struct {
	// DistinctOn: The properties to make distinct. The query results will
	// contain the first
	// result for each distinct combination of values for the given
	// properties
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
	// Unspecified is interpreted as no limit.
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

	// ForceSendFields is a list of field names (e.g. "DistinctOn") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Query) MarshalJSON() ([]byte, error) {
	type noMethod Query
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// QueryResultBatch: A batch of results produced by a query.
type QueryResultBatch struct {
	// EndCursor: A cursor that points to the position after the last result
	// in the batch.
	EndCursor string `json:"endCursor,omitempty"`

	// EntityResultType: The result type for every entity in
	// `entity_results`.
	//
	// Possible values:
	//   "RESULT_TYPE_UNSPECIFIED" - Unspecified. This value is never used.
	//   "FULL" - The key and properties.
	//   "PROJECTION" - A projected subset of properties. The entity may
	// have no key.
	//   "KEY_ONLY" - Only the key.
	EntityResultType string `json:"entityResultType,omitempty"`

	// EntityResults: The results for this batch.
	EntityResults []*EntityResult `json:"entityResults,omitempty"`

	// MoreResults: The state of the query after the current batch.
	//
	// Possible values:
	//   "MORE_RESULTS_TYPE_UNSPECIFIED" - Unspecified. This value is never
	// used.
	//   "NOT_FINISHED" - There may be additional batches to fetch from this
	// query.
	//   "MORE_RESULTS_AFTER_LIMIT" - The query is finished, but there may
	// be more results after the limit.
	//   "MORE_RESULTS_AFTER_CURSOR" - The query is finished, but there may
	// be more results after the end cursor.
	//   "NO_MORE_RESULTS" - The query has been exhausted.
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
	// This applies to the range of results from the query's `start_cursor`
	// (or
	// the beginning of the query if no cursor was given) to this
	// batch's
	// `end_cursor` (not the query's `end_cursor`).
	//
	// In a single transaction, subsequent query result batches for the same
	// query
	// can have a greater snapshot version number. Each batch's snapshot
	// version
	// is valid for all preceding batches.
	SnapshotVersion int64 `json:"snapshotVersion,omitempty,string"`

	// ForceSendFields is a list of field names (e.g. "EndCursor") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *QueryResultBatch) MarshalJSON() ([]byte, error) {
	type noMethod QueryResultBatch
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// ReadOnly: Options specific to read-only transactions.
type ReadOnly struct {
}

// ReadOptions: The options shared by read requests.
type ReadOptions struct {
	// NewTransaction: Options for beginning a new transaction for this
	// request.
	// The new transaction handle will be returned in the corresponding
	// response
	// as either google.datastore.v1beta3.LookupResponse.transaction
	// or
	// google.datastore.v1beta3.RunQueryResponse.transaction.
	NewTransaction *TransactionOptions `json:"newTransaction,omitempty"`

	// ReadConsistency: The non-transactional read consistency to
	// use.
	// Cannot be set to `STRONG` for global queries.
	//
	// Possible values:
	//   "READ_CONSISTENCY_UNSPECIFIED" - Unspecified. This value must not
	// be used.
	//   "STRONG" - Strong consistency.
	//   "EVENTUAL" - Eventual consistency.
	ReadConsistency string `json:"readConsistency,omitempty"`

	// Transaction: The transaction in which to read.
	Transaction string `json:"transaction,omitempty"`

	// ForceSendFields is a list of field names (e.g. "NewTransaction") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *ReadOptions) MarshalJSON() ([]byte, error) {
	type noMethod ReadOptions
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// ReadWrite: Options specific to read / write transactions.
type ReadWrite struct {
}

// Rectangle: A "rectangle" on the surface of the Earth, defined as the
// area between two
// meridians and two parallels.
type Rectangle struct {
	// Northeast: The northeast point of the rectangle.
	// Its latitude must be not less than that of the southwest point.
	Northeast *LatLng `json:"northeast,omitempty"`

	// Southwest: The southwest point of the rectangle.
	Southwest *LatLng `json:"southwest,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Northeast") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Rectangle) MarshalJSON() ([]byte, error) {
	type noMethod Rectangle
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// RollbackRequest: The request for
// google.datastore.v1beta3.Datastore.Rollback.
type RollbackRequest struct {
	// DatabaseId: If not empty, the ID of the database against which to
	// make the request.
	DatabaseId string `json:"databaseId,omitempty"`

	// Transaction: The transaction identifier, returned by a call
	// to
	// google.datastore.v1beta3.Datastore.BeginTransaction.
	Transaction string `json:"transaction,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DatabaseId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *RollbackRequest) MarshalJSON() ([]byte, error) {
	type noMethod RollbackRequest
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// RollbackResponse: The response for
// google.datastore.v1beta3.Datastore.Rollback
// (an empty message).
type RollbackResponse struct {
	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`
}

// RunQueryRequest: The request for
// google.datastore.v1beta3.Datastore.RunQuery.
type RunQueryRequest struct {
	// DatabaseId: If not empty, the ID of the database against which to
	// make the request.
	DatabaseId string `json:"databaseId,omitempty"`

	// GqlQuery: The GQL query to run.
	GqlQuery *GqlQuery `json:"gqlQuery,omitempty"`

	// PartitionId: Entities are partitioned into subsets, identified by a
	// partition ID.
	// Queries are scoped to a single partition.
	// This partition ID is normalized with the standard default
	// context
	// partition ID.
	PartitionId *PartitionId `json:"partitionId,omitempty"`

	// PropertyMask: The properties to return.
	// This field must not be set for a projection query.
	//
	// See google.datastore.v1beta3.LookupRequest.property_mask.
	PropertyMask string `json:"propertyMask,omitempty"`

	// Query: The query to run.
	Query *Query `json:"query,omitempty"`

	// ReadOptions: The options for this query.
	ReadOptions *ReadOptions `json:"readOptions,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DatabaseId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *RunQueryRequest) MarshalJSON() ([]byte, error) {
	type noMethod RunQueryRequest
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// RunQueryResponse: The response for
// google.datastore.v1beta3.Datastore.RunQuery.
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
	// was set in google.datastore.v1beta3.RunQueryRequest.read_options.
	Transaction string `json:"transaction,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Batch") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *RunQueryResponse) MarshalJSON() ([]byte, error) {
	type noMethod RunQueryResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// StContainsFilter: A filter that selects geo point values that are
// within a given region.
type StContainsFilter struct {
	// ContainedIn: A region within which the property's value should be
	// contained.
	ContainedIn *GeoRegion `json:"containedIn,omitempty"`

	// Property: The property to filter by.
	Property *PropertyReference `json:"property,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ContainedIn") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *StContainsFilter) MarshalJSON() ([]byte, error) {
	type noMethod StContainsFilter
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

// TransactionOptions: Options for beginning a new
// transaction.
// Transactions can be created explicitly with calls
// to
// google.datastore.v1beta3.Datastore.BeginTransaction or implicitly
// by
// setting google.datastore.v1beta3.ReadOptions.new_transaction in
// read
// requests.
type TransactionOptions struct {
	// ReadOnly: The transaction should only allow reads.
	ReadOnly *ReadOnly `json:"readOnly,omitempty"`

	// ReadWrite: The transaction should allow both reads and writes.
	ReadWrite *ReadWrite `json:"readWrite,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ReadOnly") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *TransactionOptions) MarshalJSON() ([]byte, error) {
	type noMethod TransactionOptions
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// UpdateIndexMetadata: Metadata for build index operations.
type UpdateIndexMetadata struct {
	// Common: Metadata common to all Datastore Admin Operations.
	Common *CommonMetadata `json:"common,omitempty"`

	// IndexDefinition: The index definition.  Handy for filtering.
	IndexDefinition *IndexDefinition `json:"indexDefinition,omitempty"`

	// IndexId: The format matches that of Index.index_id.
	IndexId string `json:"indexId,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Common") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *UpdateIndexMetadata) MarshalJSON() ([]byte, error) {
	type noMethod UpdateIndexMetadata
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// Value: A message that can hold any of the supported value types and
// associated
// metadata.
type Value struct {
	// ArrayValue: An array value.
	// Cannot contain another array value.
	// A `Value` instance that sets field `array_value` must not set
	// fields
	// `meaning` or `exclude_from_indexes`.
	ArrayValue *ArrayValue `json:"arrayValue,omitempty"`

	// BlobValue: A blob value.
	// May have at most 1,000,000 bytes.
	// When `exclude_from_indexes` is false, may have at most 1500 bytes.
	// In JSON requests, must be base64-encoded.
	BlobValue string `json:"blobValue,omitempty"`

	// BooleanValue: A boolean value.
	BooleanValue bool `json:"booleanValue,omitempty"`

	// DoubleValue: A double value.
	DoubleValue float64 `json:"doubleValue,omitempty"`

	// EntityValue: An entity value.
	//
	// - May have no key.
	// - May have a key with an incomplete key path.
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
	//
	// Possible values:
	//   "NULL_VALUE" - Null value.
	NullValue string `json:"nullValue,omitempty"`

	// StringValue: A UTF-8 encoded string value.
	// When `exclude_from_indexes` is false (it is indexed) , may have at
	// most 1500 bytes.
	// Otherwise, may be set to at least 1,000,000 bytes.
	StringValue string `json:"stringValue,omitempty"`

	// TimestampValue: A timestamp value.
	// When stored in the Datastore, precise only to microseconds;
	// any additional precision is rounded down.
	TimestampValue string `json:"timestampValue,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ArrayValue") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Value) MarshalJSON() ([]byte, error) {
	type noMethod Value
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// WatchRequest: The request for
// google.datastore.v1beta3.DatastoreWatcher.Watch
type WatchRequest struct {
	// DatabaseId: Database ID against which to make the request.
	DatabaseId string `json:"databaseId,omitempty"`

	// Targets: The set of watcher targets to include in the stream.
	// changes will be returned with server
	// assigned `target_ids` in the same order as the targets are specified
	// here.
	Targets []*WatchTarget `json:"targets,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DatabaseId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *WatchRequest) MarshalJSON() ([]byte, error) {
	type noMethod WatchRequest
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// WatchTarget: A specification of a set of entities to watch.
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

	// ForceSendFields is a list of field names (e.g. "GqlQuery") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *WatchTarget) MarshalJSON() ([]byte, error) {
	type noMethod WatchTarget
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// method id "datastore.projects.allocateIds":

type ProjectsAllocateIdsCall struct {
	s                  *Service
	projectId          string
	allocateidsrequest *AllocateIdsRequest
	urlParams_         gensupport.URLParams
	ctx_               context.Context
}

// AllocateIds: Allocates IDs for the given keys, which is useful for
// referencing an entity
// before it is inserted.
func (r *ProjectsService) AllocateIds(projectId string, allocateidsrequest *AllocateIdsRequest) *ProjectsAllocateIdsCall {
	c := &ProjectsAllocateIdsCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.allocateidsrequest = allocateidsrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAllocateIdsCall) Fields(s ...googleapi.Field) *ProjectsAllocateIdsCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsAllocateIdsCall) Context(ctx context.Context) *ProjectsAllocateIdsCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsAllocateIdsCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.allocateidsrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectId}:allocateIds")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "datastore.projects.allocateIds" call.
// Exactly one of *AllocateIdsResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *AllocateIdsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsAllocateIdsCall) Do() (*AllocateIdsResponse, error) {
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
	ret := &AllocateIdsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
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
	urlParams_              gensupport.URLParams
	ctx_                    context.Context
}

// BeginTransaction: Begins a new transaction.
func (r *ProjectsService) BeginTransaction(projectId string, begintransactionrequest *BeginTransactionRequest) *ProjectsBeginTransactionCall {
	c := &ProjectsBeginTransactionCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.begintransactionrequest = begintransactionrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsBeginTransactionCall) Fields(s ...googleapi.Field) *ProjectsBeginTransactionCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsBeginTransactionCall) Context(ctx context.Context) *ProjectsBeginTransactionCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsBeginTransactionCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.begintransactionrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectId}:beginTransaction")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "datastore.projects.beginTransaction" call.
// Exactly one of *BeginTransactionResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *BeginTransactionResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsBeginTransactionCall) Do() (*BeginTransactionResponse, error) {
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
	ret := &BeginTransactionResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
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
	urlParams_    gensupport.URLParams
	ctx_          context.Context
}

// Commit: Commits a transaction, optionally creating, deleting or
// modifying some
// entities.
func (r *ProjectsService) Commit(projectId string, commitrequest *CommitRequest) *ProjectsCommitCall {
	c := &ProjectsCommitCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.commitrequest = commitrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsCommitCall) Fields(s ...googleapi.Field) *ProjectsCommitCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsCommitCall) Context(ctx context.Context) *ProjectsCommitCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsCommitCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.commitrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectId}:commit")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "datastore.projects.commit" call.
// Exactly one of *CommitResponse or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *CommitResponse.ServerResponse.Header or (if a response was returned
// at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsCommitCall) Do() (*CommitResponse, error) {
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
	ret := &CommitResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
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
	urlParams_    gensupport.URLParams
	ctx_          context.Context
}

// Export: Exports a copy of all or a subset of Entities from a Google
// Cloud Datastore
// Project to another storage system, such as Google Cloud Storage.
// Recent
// updates to Entities may not be reflected in the export. The export
// occurs
// in the background and its progress can be monitored and managed via
// the
// Operation resource that is created.  The output of an export may only
// be
// used once the associated Operation is done. If an export Operation
// is
// cancelled before completion it may leave partial data behind in
// Google
// Cloud Storage.
func (r *ProjectsService) Export(projectId string, exportrequest *ExportRequest) *ProjectsExportCall {
	c := &ProjectsExportCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.exportrequest = exportrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsExportCall) Fields(s ...googleapi.Field) *ProjectsExportCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsExportCall) Context(ctx context.Context) *ProjectsExportCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsExportCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.exportrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectId}:export")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "datastore.projects.export" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsExportCall) Do() (*Operation, error) {
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
	urlParams_    gensupport.URLParams
	ctx_          context.Context
}

// Import: Imports Entities into a Google Cloud Datastore Project.
// Existing Entities
// with the same key are overwritten. The import occurs in the
// background and
// its progress can be monitored and managed via the Operation resource
// that
// is created.  If an Import Operation is cancelled, it is possible that
// a
// subset of the data has already been imported to the Datastore.
func (r *ProjectsService) Import(projectId string, importrequest *ImportRequest) *ProjectsImportCall {
	c := &ProjectsImportCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.importrequest = importrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsImportCall) Fields(s ...googleapi.Field) *ProjectsImportCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsImportCall) Context(ctx context.Context) *ProjectsImportCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsImportCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.importrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectId}:import")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "datastore.projects.import" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsImportCall) Do() (*Operation, error) {
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
	urlParams_    gensupport.URLParams
	ctx_          context.Context
}

// Lookup: Looks up entities by key.
func (r *ProjectsService) Lookup(projectId string, lookuprequest *LookupRequest) *ProjectsLookupCall {
	c := &ProjectsLookupCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.lookuprequest = lookuprequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLookupCall) Fields(s ...googleapi.Field) *ProjectsLookupCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsLookupCall) Context(ctx context.Context) *ProjectsLookupCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsLookupCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.lookuprequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectId}:lookup")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "datastore.projects.lookup" call.
// Exactly one of *LookupResponse or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *LookupResponse.ServerResponse.Header or (if a response was returned
// at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsLookupCall) Do() (*LookupResponse, error) {
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
	ret := &LookupResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
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
	urlParams_        gensupport.URLParams
	ctx_              context.Context
}

// MultiWatch: Watch changes to the results of a dynamically changeable
// set of queries.
func (r *ProjectsService) MultiWatch(projectId string, multiwatchrequest *MultiWatchRequest) *ProjectsMultiWatchCall {
	c := &ProjectsMultiWatchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.multiwatchrequest = multiwatchrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsMultiWatchCall) Fields(s ...googleapi.Field) *ProjectsMultiWatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsMultiWatchCall) Context(ctx context.Context) *ProjectsMultiWatchCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsMultiWatchCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.multiwatchrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectId}:multiWatch")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "datastore.projects.multiWatch" call.
// Exactly one of *ChangeBatch or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *ChangeBatch.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsMultiWatchCall) Do() (*ChangeBatch, error) {
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
	ret := &ChangeBatch{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
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
	urlParams_      gensupport.URLParams
	ctx_            context.Context
}

// Rollback: Rolls back a transaction.
func (r *ProjectsService) Rollback(projectId string, rollbackrequest *RollbackRequest) *ProjectsRollbackCall {
	c := &ProjectsRollbackCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.rollbackrequest = rollbackrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsRollbackCall) Fields(s ...googleapi.Field) *ProjectsRollbackCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsRollbackCall) Context(ctx context.Context) *ProjectsRollbackCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsRollbackCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.rollbackrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectId}:rollback")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "datastore.projects.rollback" call.
// Exactly one of *RollbackResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *RollbackResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsRollbackCall) Do() (*RollbackResponse, error) {
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
	ret := &RollbackResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
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
	urlParams_      gensupport.URLParams
	ctx_            context.Context
}

// RunQuery: Queries for entities.
func (r *ProjectsService) RunQuery(projectId string, runqueryrequest *RunQueryRequest) *ProjectsRunQueryCall {
	c := &ProjectsRunQueryCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.runqueryrequest = runqueryrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsRunQueryCall) Fields(s ...googleapi.Field) *ProjectsRunQueryCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsRunQueryCall) Context(ctx context.Context) *ProjectsRunQueryCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsRunQueryCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.runqueryrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectId}:runQuery")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "datastore.projects.runQuery" call.
// Exactly one of *RunQueryResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *RunQueryResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsRunQueryCall) Do() (*RunQueryResponse, error) {
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
	ret := &RunQueryResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
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
	urlParams_   gensupport.URLParams
	ctx_         context.Context
}

// Watch: Watch changes to the results of a given set of queries.
func (r *ProjectsService) Watch(projectId string, watchrequest *WatchRequest) *ProjectsWatchCall {
	c := &ProjectsWatchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.watchrequest = watchrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsWatchCall) Fields(s ...googleapi.Field) *ProjectsWatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsWatchCall) Context(ctx context.Context) *ProjectsWatchCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsWatchCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.watchrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectId}:watch")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "datastore.projects.watch" call.
// Exactly one of *ChangeBatch or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *ChangeBatch.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsWatchCall) Do() (*ChangeBatch, error) {
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
	ret := &ChangeBatch{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
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
	s            *Service
	projectId    string
	indexId      string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// Get: Gets an index.
func (r *ProjectsIndexesService) Get(projectId string, indexId string) *ProjectsIndexesGetCall {
	c := &ProjectsIndexesGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.indexId = indexId
	return c
}

// DatabaseId sets the optional parameter "databaseId": The database id.
func (c *ProjectsIndexesGetCall) DatabaseId(databaseId string) *ProjectsIndexesGetCall {
	c.urlParams_.Set("databaseId", databaseId)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsIndexesGetCall) Fields(s ...googleapi.Field) *ProjectsIndexesGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsIndexesGetCall) IfNoneMatch(entityTag string) *ProjectsIndexesGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsIndexesGetCall) Context(ctx context.Context) *ProjectsIndexesGetCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsIndexesGetCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectId}/indexes/{+indexId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"indexId":   c.indexId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		req.Header.Set("If-None-Match", c.ifNoneMatch_)
	}
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "datastore.projects.indexes.get" call.
// Exactly one of *Index or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Index.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsIndexesGetCall) Do() (*Index, error) {
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
	ret := &Index{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
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
	s            *Service
	projectId    string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// List: Lists the indexes that match the specified filters.
// Only lists indexes that are not in their initial state.
func (r *ProjectsIndexesService) List(projectId string) *ProjectsIndexesListCall {
	c := &ProjectsIndexesListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	return c
}

// DatabaseId sets the optional parameter "databaseId": The database id.
func (c *ProjectsIndexesListCall) DatabaseId(databaseId string) *ProjectsIndexesListCall {
	c.urlParams_.Set("databaseId", databaseId)
	return c
}

// Filter sets the optional parameter "filter": The standard List
// filter.
// The filter expression may filter only on definition.kind,
// eg
// '(definition.kind == "User") or (definition.kind == "Widget")'.
func (c *ProjectsIndexesListCall) Filter(filter string) *ProjectsIndexesListCall {
	c.urlParams_.Set("filter", filter)
	return c
}

// PageSize sets the optional parameter "pageSize": The standard List
// page size.
func (c *ProjectsIndexesListCall) PageSize(pageSize int64) *ProjectsIndexesListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The standard List
// page token.
func (c *ProjectsIndexesListCall) PageToken(pageToken string) *ProjectsIndexesListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsIndexesListCall) Fields(s ...googleapi.Field) *ProjectsIndexesListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsIndexesListCall) IfNoneMatch(entityTag string) *ProjectsIndexesListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsIndexesListCall) Context(ctx context.Context) *ProjectsIndexesListCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsIndexesListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectId}/indexes")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		req.Header.Set("If-None-Match", c.ifNoneMatch_)
	}
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "datastore.projects.indexes.list" call.
// Exactly one of *ListIndexesResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListIndexesResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsIndexesListCall) Do() (*ListIndexesResponse, error) {
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
	ret := &ListIndexesResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
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
	urlParams_         gensupport.URLParams
	ctx_               context.Context
}

// Lookup: Looks up an index by definition.
func (r *ProjectsIndexesService) Lookup(projectId string, lookupindexrequest *LookupIndexRequest) *ProjectsIndexesLookupCall {
	c := &ProjectsIndexesLookupCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.lookupindexrequest = lookupindexrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsIndexesLookupCall) Fields(s ...googleapi.Field) *ProjectsIndexesLookupCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsIndexesLookupCall) Context(ctx context.Context) *ProjectsIndexesLookupCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsIndexesLookupCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.lookupindexrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectId}/indexes:lookup")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "datastore.projects.indexes.lookup" call.
// Exactly one of *Index or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Index.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsIndexesLookupCall) Do() (*Index, error) {
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
	ret := &Index{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
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
	s          *Service
	projectId  string
	indexId    string
	index      *Index
	urlParams_ gensupport.URLParams
	ctx_       context.Context
}

// Update: Updates an index's state.
// The input index must specify a (project_id, index_id) tuple
// or an index definition (but not both), and a new state.
// This new state must be SERVING or OFF.
// The state of the key index and the kind index cannot be updated.
// If the index is already in the requested state, does nothing and
// returns
// a successful but unnamed operation.  Otherwise:
// Returns an unfinished operation.
// - If the new state is SERVING, sets the index's state to BUILDING and
// the
//     result operation's field metadata.common.operation_type is
// BUILD_INDEX.
// - If the new state is OFF, sets the index's state to CLEARING and
// the
//    result operation's field metadata.common.operation_type is
// CLEAR_INDEX.
// Once the operation finishes,
// if it is successful the index's state is the new state,
// and otherwise the index's state is ERROR.
// The result operation's field response is of type
// google.protobuf.Empty.
// The result operation's field metadata is of type UpdateIndexMetadata.
func (r *ProjectsIndexesService) Update(projectId string, indexId string, index *Index) *ProjectsIndexesUpdateCall {
	c := &ProjectsIndexesUpdateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.indexId = indexId
	c.index = index
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsIndexesUpdateCall) Fields(s ...googleapi.Field) *ProjectsIndexesUpdateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsIndexesUpdateCall) Context(ctx context.Context) *ProjectsIndexesUpdateCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsIndexesUpdateCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.index)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectId}/indexes/{+indexId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"indexId":   c.indexId,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "datastore.projects.indexes.update" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsIndexesUpdateCall) Do() (*Operation, error) {
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
// operation completed despite cancellation.
func (r *ProjectsOperationsService) Cancel(name string) *ProjectsOperationsCancelCall {
	c := &ProjectsOperationsCancelCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsOperationsCancelCall) Fields(s ...googleapi.Field) *ProjectsOperationsCancelCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsOperationsCancelCall) Context(ctx context.Context) *ProjectsOperationsCancelCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsOperationsCancelCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/{+name}:cancel")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "datastore.projects.operations.cancel" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsOperationsCancelCall) Do() (*Empty, error) {
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
func (r *ProjectsOperationsService) Delete(name string) *ProjectsOperationsDeleteCall {
	c := &ProjectsOperationsDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsOperationsDeleteCall) Fields(s ...googleapi.Field) *ProjectsOperationsDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsOperationsDeleteCall) Context(ctx context.Context) *ProjectsOperationsDeleteCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsOperationsDeleteCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "datastore.projects.operations.delete" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsOperationsDeleteCall) Do() (*Empty, error) {
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
func (r *ProjectsOperationsService) Get(name string) *ProjectsOperationsGetCall {
	c := &ProjectsOperationsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsOperationsGetCall) Fields(s ...googleapi.Field) *ProjectsOperationsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsOperationsGetCall) IfNoneMatch(entityTag string) *ProjectsOperationsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsOperationsGetCall) Context(ctx context.Context) *ProjectsOperationsGetCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsOperationsGetCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		req.Header.Set("If-None-Match", c.ifNoneMatch_)
	}
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "datastore.projects.operations.get" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsOperationsGetCall) Do() (*Operation, error) {
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
func (r *ProjectsOperationsService) List(name string) *ProjectsOperationsListCall {
	c := &ProjectsOperationsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Filter sets the optional parameter "filter": The standard list
// filter.
func (c *ProjectsOperationsListCall) Filter(filter string) *ProjectsOperationsListCall {
	c.urlParams_.Set("filter", filter)
	return c
}

// PageSize sets the optional parameter "pageSize": The standard list
// page size.
func (c *ProjectsOperationsListCall) PageSize(pageSize int64) *ProjectsOperationsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The standard list
// page token.
func (c *ProjectsOperationsListCall) PageToken(pageToken string) *ProjectsOperationsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsOperationsListCall) Fields(s ...googleapi.Field) *ProjectsOperationsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsOperationsListCall) IfNoneMatch(entityTag string) *ProjectsOperationsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsOperationsListCall) Context(ctx context.Context) *ProjectsOperationsListCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsOperationsListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		req.Header.Set("If-None-Match", c.ifNoneMatch_)
	}
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "datastore.projects.operations.list" call.
// Exactly one of *ListOperationsResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListOperationsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsOperationsListCall) Do() (*ListOperationsResponse, error) {
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
