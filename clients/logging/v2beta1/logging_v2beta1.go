// Package logging provides access to the Google Cloud Logging API.
//
// See https://cloud.google.com/logging/docs/
//
// Usage example:
//
//   import "github.com/skelterjohn/gcloud_apis/clients/logging/v2beta1"
//   ...
//   loggingService, err := logging.New(oauthHttpClient)
package logging

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

const apiId = "logging:v2beta1"
const apiName = "logging"
const apiVersion = "v2beta1"
const basePath = "https://logging.googleapis.com/"

// OAuth2 scopes used by this API.
const (
	// View and manage your data across Google Cloud Platform services
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"

	// View your data across Google Cloud Platform services
	CloudPlatformReadOnlyScope = "https://www.googleapis.com/auth/cloud-platform.read-only"

	// Administrate log data for your projects
	LoggingAdminScope = "https://www.googleapis.com/auth/logging.admin"

	// View log data for your projects
	LoggingReadScope = "https://www.googleapis.com/auth/logging.read"

	// Submit log data for your projects
	LoggingWriteScope = "https://www.googleapis.com/auth/logging.write"

	// View and write monitoring data for all of your Google and third-party
	// Cloud and API projects
	MonitoringScope = "https://www.googleapis.com/auth/monitoring"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.BillingAccounts = NewBillingAccountsService(s)
	s.Entries = NewEntriesService(s)
	s.MonitoredResourceDescriptors = NewMonitoredResourceDescriptorsService(s)
	s.Organizations = NewOrganizationsService(s)
	s.Projects = NewProjectsService(s)
	s.V2beta1 = NewV2beta1Service(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	BillingAccounts *BillingAccountsService

	Entries *EntriesService

	MonitoredResourceDescriptors *MonitoredResourceDescriptorsService

	Organizations *OrganizationsService

	Projects *ProjectsService

	V2beta1 *V2beta1Service
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewBillingAccountsService(s *Service) *BillingAccountsService {
	rs := &BillingAccountsService{s: s}
	rs.Logs = NewBillingAccountsLogsService(s)
	rs.ResourceKeys = NewBillingAccountsResourceKeysService(s)
	rs.ResourceTypes = NewBillingAccountsResourceTypesService(s)
	return rs
}

type BillingAccountsService struct {
	s *Service

	Logs *BillingAccountsLogsService

	ResourceKeys *BillingAccountsResourceKeysService

	ResourceTypes *BillingAccountsResourceTypesService
}

func NewBillingAccountsLogsService(s *Service) *BillingAccountsLogsService {
	rs := &BillingAccountsLogsService{s: s}
	return rs
}

type BillingAccountsLogsService struct {
	s *Service
}

func NewBillingAccountsResourceKeysService(s *Service) *BillingAccountsResourceKeysService {
	rs := &BillingAccountsResourceKeysService{s: s}
	return rs
}

type BillingAccountsResourceKeysService struct {
	s *Service
}

func NewBillingAccountsResourceTypesService(s *Service) *BillingAccountsResourceTypesService {
	rs := &BillingAccountsResourceTypesService{s: s}
	rs.Values = NewBillingAccountsResourceTypesValuesService(s)
	return rs
}

type BillingAccountsResourceTypesService struct {
	s *Service

	Values *BillingAccountsResourceTypesValuesService
}

func NewBillingAccountsResourceTypesValuesService(s *Service) *BillingAccountsResourceTypesValuesService {
	rs := &BillingAccountsResourceTypesValuesService{s: s}
	return rs
}

type BillingAccountsResourceTypesValuesService struct {
	s *Service
}

func NewEntriesService(s *Service) *EntriesService {
	rs := &EntriesService{s: s}
	return rs
}

type EntriesService struct {
	s *Service
}

func NewMonitoredResourceDescriptorsService(s *Service) *MonitoredResourceDescriptorsService {
	rs := &MonitoredResourceDescriptorsService{s: s}
	return rs
}

type MonitoredResourceDescriptorsService struct {
	s *Service
}

func NewOrganizationsService(s *Service) *OrganizationsService {
	rs := &OrganizationsService{s: s}
	rs.Logs = NewOrganizationsLogsService(s)
	rs.ResourceKeys = NewOrganizationsResourceKeysService(s)
	rs.ResourceTypes = NewOrganizationsResourceTypesService(s)
	return rs
}

type OrganizationsService struct {
	s *Service

	Logs *OrganizationsLogsService

	ResourceKeys *OrganizationsResourceKeysService

	ResourceTypes *OrganizationsResourceTypesService
}

func NewOrganizationsLogsService(s *Service) *OrganizationsLogsService {
	rs := &OrganizationsLogsService{s: s}
	return rs
}

type OrganizationsLogsService struct {
	s *Service
}

func NewOrganizationsResourceKeysService(s *Service) *OrganizationsResourceKeysService {
	rs := &OrganizationsResourceKeysService{s: s}
	return rs
}

type OrganizationsResourceKeysService struct {
	s *Service
}

func NewOrganizationsResourceTypesService(s *Service) *OrganizationsResourceTypesService {
	rs := &OrganizationsResourceTypesService{s: s}
	rs.Values = NewOrganizationsResourceTypesValuesService(s)
	return rs
}

type OrganizationsResourceTypesService struct {
	s *Service

	Values *OrganizationsResourceTypesValuesService
}

func NewOrganizationsResourceTypesValuesService(s *Service) *OrganizationsResourceTypesValuesService {
	rs := &OrganizationsResourceTypesValuesService{s: s}
	return rs
}

type OrganizationsResourceTypesValuesService struct {
	s *Service
}

func NewProjectsService(s *Service) *ProjectsService {
	rs := &ProjectsService{s: s}
	rs.Logs = NewProjectsLogsService(s)
	rs.Metrics = NewProjectsMetricsService(s)
	rs.ResourceKeys = NewProjectsResourceKeysService(s)
	rs.ResourceTypes = NewProjectsResourceTypesService(s)
	rs.Sinks = NewProjectsSinksService(s)
	return rs
}

type ProjectsService struct {
	s *Service

	Logs *ProjectsLogsService

	Metrics *ProjectsMetricsService

	ResourceKeys *ProjectsResourceKeysService

	ResourceTypes *ProjectsResourceTypesService

	Sinks *ProjectsSinksService
}

func NewProjectsLogsService(s *Service) *ProjectsLogsService {
	rs := &ProjectsLogsService{s: s}
	return rs
}

type ProjectsLogsService struct {
	s *Service
}

func NewProjectsMetricsService(s *Service) *ProjectsMetricsService {
	rs := &ProjectsMetricsService{s: s}
	return rs
}

type ProjectsMetricsService struct {
	s *Service
}

func NewProjectsResourceKeysService(s *Service) *ProjectsResourceKeysService {
	rs := &ProjectsResourceKeysService{s: s}
	return rs
}

type ProjectsResourceKeysService struct {
	s *Service
}

func NewProjectsResourceTypesService(s *Service) *ProjectsResourceTypesService {
	rs := &ProjectsResourceTypesService{s: s}
	rs.Values = NewProjectsResourceTypesValuesService(s)
	return rs
}

type ProjectsResourceTypesService struct {
	s *Service

	Values *ProjectsResourceTypesValuesService
}

func NewProjectsResourceTypesValuesService(s *Service) *ProjectsResourceTypesValuesService {
	rs := &ProjectsResourceTypesValuesService{s: s}
	return rs
}

type ProjectsResourceTypesValuesService struct {
	s *Service
}

func NewProjectsSinksService(s *Service) *ProjectsSinksService {
	rs := &ProjectsSinksService{s: s}
	return rs
}

type ProjectsSinksService struct {
	s *Service
}

func NewV2beta1Service(s *Service) *V2beta1Service {
	rs := &V2beta1Service{s: s}
	return rs
}

type V2beta1Service struct {
	s *Service
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

// GetLogsUsageResponse: The response from `GetLogsUsage`.
type GetLogsUsageResponse struct {
	// Usage: A collection of ranges that describes logs usage and allowed
	// quota
	// over requested time period.
	// Ranges are aligned to the full hour, and are guaranteed to be
	// contiguous.
	Usage []*Usage `json:"usage,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Usage") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *GetLogsUsageResponse) MarshalJSON() ([]byte, error) {
	type noMethod GetLogsUsageResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// HttpRequest: A common proto for logging HTTP requests. Only contains
// semantics
// defined by the HTTP specification. Product-specific
// logging
// information MUST be defined in a separate message.
type HttpRequest struct {
	// CacheFillBytes: The number of HTTP response bytes inserted into
	// cache. Set only when a
	// cache fill was attempted.
	CacheFillBytes int64 `json:"cacheFillBytes,omitempty,string"`

	// CacheHit: Whether or not an entity was served from cache
	// (with or without validation).
	CacheHit bool `json:"cacheHit,omitempty"`

	// CacheLookup: Whether or not a cache lookup was attempted.
	CacheLookup bool `json:"cacheLookup,omitempty"`

	// CacheValidatedWithOriginServer: Whether or not the response was
	// validated with the origin server before
	// being served from cache. This field is only meaningful if `cache_hit`
	// is
	// True.
	CacheValidatedWithOriginServer bool `json:"cacheValidatedWithOriginServer,omitempty"`

	// Referer: The referer URL of the request, as defined in
	// [HTTP/1.1 Header Field
	// Definitions](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html).
	Referer string `json:"referer,omitempty"`

	// RemoteIp: The IP address (IPv4 or IPv6) of the client that issued the
	// HTTP
	// request. Examples: "192.168.1.1", "FE80::0202:B3FF:FE1E:8329".
	RemoteIp string `json:"remoteIp,omitempty"`

	// RequestMethod: The request method. Examples: "GET", "HEAD",
	// "PUT", "POST".
	RequestMethod string `json:"requestMethod,omitempty"`

	// RequestSize: The size of the HTTP request message in bytes, including
	// the request
	// headers and the request body.
	RequestSize int64 `json:"requestSize,omitempty,string"`

	// RequestUrl: The scheme (http, https), the host name, the path and the
	// query
	// portion of the URL that was requested.
	// Example: "http://example.com/some/info?color=red".
	RequestUrl string `json:"requestUrl,omitempty"`

	// ResponseSize: The size of the HTTP response message sent back to the
	// client, in bytes,
	// including the response headers and the response body.
	ResponseSize int64 `json:"responseSize,omitempty,string"`

	// ServerIp: The IP address (IPv4 or IPv6) of the origin server that the
	// request was
	// sent to.
	ServerIp string `json:"serverIp,omitempty"`

	// Status: The response code indicating the status of
	// response.
	// Examples: 200, 404.
	Status int64 `json:"status,omitempty"`

	// UserAgent: The user agent sent by the client. Example:
	// "Mozilla/4.0 (compatible; MSIE 6.0; Windows 98; Q312461; .NET CLR
	// 1.0.3705)".
	UserAgent string `json:"userAgent,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CacheFillBytes") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *HttpRequest) MarshalJSON() ([]byte, error) {
	type noMethod HttpRequest
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

type InternalEntityId struct {
	// BillingAccountId: Unique identifier of a billing account
	BillingAccountId string `json:"billingAccountId,omitempty"`

	// FolderNumber: Gaia Id of a folder
	FolderNumber int64 `json:"folderNumber,omitempty,string"`

	// OrganizationNumber: Gaia Id of an organization
	OrganizationNumber int64 `json:"organizationNumber,omitempty,string"`

	// ProjectNumber: Gaia Id of a project
	ProjectNumber int64 `json:"projectNumber,omitempty,string"`

	// ForceSendFields is a list of field names (e.g. "BillingAccountId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *InternalEntityId) MarshalJSON() ([]byte, error) {
	type noMethod InternalEntityId
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// LabelDescriptor: A description of a label.
type LabelDescriptor struct {
	// Description: A human-readable description for the label.
	Description string `json:"description,omitempty"`

	// Key: The label key.
	Key string `json:"key,omitempty"`

	// ValueType: The type of data that can be assigned to the label.
	//
	// Possible values:
	//   "STRING" - A variable-length string. This is the default.
	//   "BOOL" - Boolean; true or false.
	//   "INT64" - A 64-bit signed integer.
	ValueType string `json:"valueType,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Description") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *LabelDescriptor) MarshalJSON() ([]byte, error) {
	type noMethod LabelDescriptor
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// ListLogEntriesRequest: The parameters to `ListLogEntries`.
type ListLogEntriesRequest struct {
	// Filter: Optional. A filter that chooses which log entries to return.
	// See [Advanced
	// Logs Filters](/logging/docs/view/advanced_filters).  Only log entries
	// that
	// match the filter are returned.  An empty filter matches all log
	// entries.
	Filter string `json:"filter,omitempty"`

	IsV1Request bool `json:"isV1Request,omitempty"`

	// OrderBy: Optional. How the results should be sorted.  Presently, the
	// only permitted
	// values are "timestamp asc" (default) and "timestamp desc". The
	// first
	// option returns entries in order of increasing values
	// of
	// `LogEntry.timestamp` (oldest first), and the second option returns
	// entries
	// in order of decreasing timestamps (newest first).  Entries with
	// equal
	// timestamps are returned in order of `LogEntry.insertId`.
	OrderBy string `json:"orderBy,omitempty"`

	// PageSize: Optional. The maximum number of results to return from this
	// request.
	// Non-positive values are ignored.  The presence of `nextPageToken` in
	// the
	// response indicates that more results might be available.
	PageSize int64 `json:"pageSize,omitempty"`

	// PageToken: Optional. If present, then retrieve the next batch of
	// results from the
	// preceding call to this method.  `pageToken` must be the value
	// of
	// `nextPageToken` from the previous response.  The values of other
	// method
	// parameters should be identical to those in the previous call.
	PageToken string `json:"pageToken,omitempty"`

	// ProjectIds: Deprecated. One or more project identifiers or project
	// numbers from which
	// to retrieve log entries.  Examples: "my-project-1A",
	// "1234567890". If
	// present, these project identifiers are converted to resource format
	// and
	// added to the list of resources in `resourceNames`. Callers should
	// use
	// `resourceNames` rather than this parameter.
	ProjectIds []string `json:"projectIds,omitempty"`

	// ResourceNames: Optional. One or more cloud resources from which to
	// retrieve log entries.
	// Example: "projects/my-project-1A", "projects/1234567890".
	// Projects
	// listed in `projectIds` are added to this list.
	ResourceNames []string `json:"resourceNames,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Filter") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *ListLogEntriesRequest) MarshalJSON() ([]byte, error) {
	type noMethod ListLogEntriesRequest
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// ListLogEntriesResponse: Result returned from `ListLogEntries`.
type ListLogEntriesResponse struct {
	// Entries: A list of log entries.
	Entries []*LogEntry `json:"entries,omitempty"`

	// LastObservedEntryTimestamp: The timestamp of the last log entry that
	// was examined before returning this
	// response. This can be used to observe progress between successive
	// queries,
	// in particular when only a page token is returned. Deprecated:
	// use
	// searched_through_timestamp.
	LastObservedEntryTimestamp string `json:"lastObservedEntryTimestamp,omitempty"`

	// NextPageToken: If there might be more results than appear in this
	// response, then
	// `nextPageToken` is included.  To get the next set of results, call
	// this
	// method again using the value of `nextPageToken` as `pageToken`.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// SearchedThroughTimestamp: The furthest point in time through which
	// the search has progressed. All
	// future entries returned using `nextPageToken` are guaranteed to have
	// a
	// timestamp at or past this point in time in the direction of the
	// search.
	// This value can be used to observe progress between successive
	// queries, in
	// particular when `nextPageToken` is returned without any log
	// entries.
	// If `nextPageToken` is not present in this response, then this
	// field
	// is left empty.
	SearchedThroughTimestamp string `json:"searchedThroughTimestamp,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Entries") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *ListLogEntriesResponse) MarshalJSON() ([]byte, error) {
	type noMethod ListLogEntriesResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// ListLogMetricsResponse: Result returned from ListLogMetrics.
type ListLogMetricsResponse struct {
	// Metrics: A list of logs-based metrics.
	Metrics []*LogMetric `json:"metrics,omitempty"`

	// NextPageToken: If there might be more results than appear in this
	// response, then
	// `nextPageToken` is included.  To get the next set of results, call
	// this
	// method again using the value of `nextPageToken` as `pageToken`.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Metrics") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *ListLogMetricsResponse) MarshalJSON() ([]byte, error) {
	type noMethod ListLogMetricsResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// ListLogsResponse: Result returned from ListLogs.
type ListLogsResponse struct {
	// LogIds: A list of log ids matching the criteria.
	LogIds []string `json:"logIds,omitempty"`

	// NextPageToken: If there might be more results than appear in this
	// response, then
	// `nextPageToken` is included.  To get the next set of results, call
	// this
	// method again using the value of `nextPageToken` as `pageToken`.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "LogIds") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *ListLogsResponse) MarshalJSON() ([]byte, error) {
	type noMethod ListLogsResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// ListMonitoredResourceDescriptorsResponse: Result returned from
// ListMonitoredResourceDescriptors.
type ListMonitoredResourceDescriptorsResponse struct {
	// NextPageToken: If there might be more results than appear in this
	// response, then
	// `nextPageToken` is included.  To get the next set of results, call
	// this
	// method again using the value of `nextPageToken` as `pageToken`.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ResourceDescriptors: A list of resource descriptors.
	ResourceDescriptors []*MonitoredResourceDescriptor `json:"resourceDescriptors,omitempty"`

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

func (s *ListMonitoredResourceDescriptorsResponse) MarshalJSON() ([]byte, error) {
	type noMethod ListMonitoredResourceDescriptorsResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// ListResourceKeysResponse: Result returned from
// `ListResourceKeysRequest`.
type ListResourceKeysResponse struct {
	// NextPageToken: If there might be more results than appear in this
	// response, then
	// `nextPageToken` is included.  To get the next set of results, call
	// this
	// method again using the value of `nextPageToken` as `pageToken`.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ResourceKeys: A list of log resource keys.
	ResourceKeys []*ResourceKeys `json:"resourceKeys,omitempty"`

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

func (s *ListResourceKeysResponse) MarshalJSON() ([]byte, error) {
	type noMethod ListResourceKeysResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// ListResourceValuesResponse: Result returned from ListResourceValues.
type ListResourceValuesResponse struct {
	// NextPageToken: If there might be more results than appear in this
	// response, then
	// `nextPageToken` is included.  To get the next set of results, call
	// this
	// method again using the value of `nextPageToken` as `pageToken`.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ResourceValuePrefixes: A list of log resource type index values.
	// Each index value has the form "/value1/value2/...",
	// where `value1` is a value in the primary index, `value2` is
	// a value in the secondary index, and so forth.
	ResourceValuePrefixes []string `json:"resourceValuePrefixes,omitempty"`

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

func (s *ListResourceValuesResponse) MarshalJSON() ([]byte, error) {
	type noMethod ListResourceValuesResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// ListSinksResponse: Result returned from `ListSinks`.
type ListSinksResponse struct {
	// NextPageToken: If there might be more results than appear in this
	// response, then
	// `nextPageToken` is included.  To get the next set of results, call
	// the same
	// method again using the value of `nextPageToken` as `pageToken`.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Sinks: A list of sinks.
	Sinks []*LogSink `json:"sinks,omitempty"`

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

func (s *ListSinksResponse) MarshalJSON() ([]byte, error) {
	type noMethod ListSinksResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// LogEntry: An individual entry in a log.
type LogEntry struct {
	// HttpRequest: Optional. Information about the HTTP request associated
	// with this log entry,
	// if applicable.
	HttpRequest *HttpRequest `json:"httpRequest,omitempty"`

	// InsertId: Optional. A unique ID for the log entry. If you provide
	// this
	// field, the logging service considers other log entries in the
	// same project with the same ID as duplicates which can be removed.
	// If
	// omitted, Stackdriver Logging will generate a unique ID for this
	// log entry.
	InsertId string `json:"insertId,omitempty"`

	InternalId *InternalEntityId `json:"internalId,omitempty"`

	// JsonPayload: The log entry payload, represented as a structure
	// that
	// is expressed as a JSON object.
	JsonPayload LogEntryJsonPayload `json:"jsonPayload,omitempty"`

	// Labels: Optional. A set of user-defined (key, value) data that
	// provides additional
	// information about the log entry.
	Labels map[string]string `json:"labels,omitempty"`

	// LogName: Required. The resource name of the log to which this log
	// entry
	// belongs. The format of the name
	// is
	// "projects/<project-id>/logs/<log-id>".
	// Examples:
	// "projects/my-projectid/logs/syslog",
	// "projects/my-project
	// id/logs/library.googleapis.com%2Fbook_log".
	//
	// The log ID part of resource name must be less than 512
	// characters
	// long and can only include the following characters: upper and
	// lower case alphanumeric characters: [A-Za-z0-9]; and
	// punctuation
	// characters: forward-slash, underscore, hyphen, and
	// period.
	// Forward-slash (`/`) characters in the log ID must be URL-encoded.
	LogName string `json:"logName,omitempty"`

	// Operation: Optional. Information about an operation associated with
	// the log entry, if
	// applicable.
	Operation *LogEntryOperation `json:"operation,omitempty"`

	// ProtoPayload: The log entry payload, represented as a protocol
	// buffer.  Some
	// Google Cloud Platform services use this field for their log
	// entry payloads.
	ProtoPayload LogEntryProtoPayload `json:"protoPayload,omitempty"`

	// Resource: Required. The monitored resource associated with this log
	// entry.
	// Example: a log entry that reports a database error would
	// be
	// associated with the monitored resource designating the
	// particular
	// database that reported the error.
	Resource *MonitoredResource `json:"resource,omitempty"`

	// Severity: Optional. The severity of the log entry. The default value
	// is
	// `LogSeverity.DEFAULT`.
	//
	// Possible values:
	//   "DEFAULT" - The log entry has no assigned severity level.
	//   "DEBUG" - Debug or trace information.
	//   "INFO" - Routine information, such as ongoing status or
	// performance.
	//   "NOTICE" - Normal but significant events, such as start up, shut
	// down, or
	// configuration.
	//   "WARNING" - Warning events might cause problems.
	//   "ERROR" - Error events are likely to cause problems.
	//   "CRITICAL" - Critical events cause more severe problems or brief
	// outages.
	//   "ALERT" - A person must take an action immediately.
	//   "EMERGENCY" - One or more systems are unusable.
	Severity string `json:"severity,omitempty"`

	// TextPayload: The log entry payload, represented as a Unicode string
	// (UTF-8).
	TextPayload string `json:"textPayload,omitempty"`

	// Timestamp: Optional. The time the event described by the log entry
	// occurred.  If
	// omitted, Stackdriver Logging will use the time the log entry is
	// received.
	Timestamp string `json:"timestamp,omitempty"`

	WriterEmailAddress string `json:"writerEmailAddress,omitempty"`

	// ForceSendFields is a list of field names (e.g. "HttpRequest") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *LogEntry) MarshalJSON() ([]byte, error) {
	type noMethod LogEntry
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

type LogEntryJsonPayload interface{}

type LogEntryProtoPayload interface{}

// LogEntryOperation: Additional information about a potentially
// long-running operation with which
// a log entry is associated.
type LogEntryOperation struct {
	// First: Optional. Set this to True if this is the first log entry in
	// the operation.
	First bool `json:"first,omitempty"`

	// Id: Required. An arbitrary operation identifier. Log entries with
	// the
	// same identifier are assumed to be part of the same operation.
	Id string `json:"id,omitempty"`

	// Last: Optional. Set this to True if this is the last log entry in the
	// operation.
	Last bool `json:"last,omitempty"`

	// Producer: Required. An arbitrary producer identifier. The combination
	// of
	// `id` and `producer` must be globally unique.  Examples for
	// `producer`:
	// "MyDivision.MyBigCompany.com",
	// "github.com/MyProject/MyApplication".
	Producer string `json:"producer,omitempty"`

	// ForceSendFields is a list of field names (e.g. "First") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *LogEntryOperation) MarshalJSON() ([]byte, error) {
	type noMethod LogEntryOperation
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// LogError: Describes a problem with a logging resource or operation.
type LogError struct {
	// Resource: A resource name associated with this error. For
	// example,
	// the name of a Cloud Storage bucket that has insufficient
	// permissions
	// to be a destination for log entries.
	Resource string `json:"resource,omitempty"`

	// Status: The error description, including a classification code,
	// an error message, and other details.
	Status *Status `json:"status,omitempty"`

	// TimeNanos: The time the error was observed, in nanoseconds since the
	// Unix epoch.
	TimeNanos int64 `json:"timeNanos,omitempty,string"`

	// ForceSendFields is a list of field names (e.g. "Resource") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *LogError) MarshalJSON() ([]byte, error) {
	type noMethod LogError
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// LogLine: Application log line emitted while processing a request.
type LogLine struct {
	// LogMessage: App-provided log message.
	LogMessage string `json:"logMessage,omitempty"`

	// Severity: Severity of this log entry.
	//
	// Possible values:
	//   "DEFAULT" - The log entry has no assigned severity level.
	//   "DEBUG" - Debug or trace information.
	//   "INFO" - Routine information, such as ongoing status or
	// performance.
	//   "NOTICE" - Normal but significant events, such as start up, shut
	// down, or
	// configuration.
	//   "WARNING" - Warning events might cause problems.
	//   "ERROR" - Error events are likely to cause problems.
	//   "CRITICAL" - Critical events cause more severe problems or brief
	// outages.
	//   "ALERT" - A person must take an action immediately.
	//   "EMERGENCY" - One or more systems are unusable.
	Severity string `json:"severity,omitempty"`

	// SourceLocation: Where in the source code this log message was
	// written.
	SourceLocation *SourceLocation `json:"sourceLocation,omitempty"`

	// Time: Approximate time when this log entry was made.
	Time string `json:"time,omitempty"`

	// ForceSendFields is a list of field names (e.g. "LogMessage") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *LogLine) MarshalJSON() ([]byte, error) {
	type noMethod LogLine
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// LogMetric: Describes a logs-based metric.  The value of the metric is
// the
// number of log entries that match a logs filter.
type LogMetric struct {
	// Description: A description of this metric, which is used in
	// documentation.
	Description string `json:"description,omitempty"`

	// Filter: An [advanced logs
	// filter](/logging/docs/view/advanced_filters).
	// Example: "logName:syslog AND severity>=ERROR".
	Filter string `json:"filter,omitempty"`

	// Name: Required. The client-assigned metric identifier.
	// Example:
	// "severe_errors".  Metric identifiers are limited to 100
	// characters and can include only the following characters:
	// `A-Z`,
	// `a-z`, `0-9`, and the special characters `_-.,+!*',()%/`.
	// The
	// forward-slash character (`/`) denotes a hierarchy of name pieces,
	// and it cannot be the first character of the name.  The '%'
	// character
	// is used to URL encode unsafe and reserved characters and must
	// be
	// followed by two hexadecimal digits according to RFC 1738.
	Name string `json:"name,omitempty"`

	// Version: The API version that created or updated this metric.
	//
	// Possible values:
	//   "V2" - Cloud Logging API V2.
	//   "V1" - Cloud Logging API V1.
	Version string `json:"version,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Description") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *LogMetric) MarshalJSON() ([]byte, error) {
	type noMethod LogMetric
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// LogSink: Describes a sink used to export log entries outside
// Stackdriver Logging.
type LogSink struct {
	// Destination: Required. The export destination. See
	// [Exporting Logs With
	// Sinks](/logging/docs/api/tasks/exporting-logs).
	// Examples:
	//
	//     "storage.googleapis.com/my-gcs-bucket"
	//
	// "bigquery.googleapis.com/projects/my-project-id/datasets/my-dataset"
	//     "pubsub.googleapis.com/projects/my-project/topics/my-topic"
	Destination string `json:"destination,omitempty"`

	// Errors: Output only. All active errors found for this sink.
	Errors []*LogError `json:"errors,omitempty"`

	// Filter: Optional. An [advanced logs
	// filter](/logging/docs/view/advanced_filters).
	// Only log entries matching the filter are exported. The filter
	// must be consistent with the log entry format specified by
	// the
	// `outputVersionFormat` parameter, regardless of the format of the
	// log entry that was originally written to Stackdriver Logging.
	// Example filter (V2 format):
	//
	//     logName=projects/my-projectid/logs/syslog AND severity>=ERROR
	Filter string `json:"filter,omitempty"`

	// FormatChange: When the format was changed.
	FormatChange string `json:"formatChange,omitempty"`

	// Name: Required. The client-assigned sink identifier, unique within
	// the
	// project. Example: "my-syslog-errors-to-pubsub".  Sink identifiers
	// are
	// limited to 1000 characters and can include only the following
	// characters:
	// `A-Z`, `a-z`, `0-9`, and the special characters `_-.`.  The maximum
	// length
	// of the name is 100 characters.
	Name string `json:"name,omitempty"`

	// OutputVersionFormat: Optional. The log entry version to use for this
	// sink's exported log
	// entries.  This version does not have to correspond to the version of
	// the
	// log entry that was written to Stackdriver Logging. If omitted, the
	// v2 format is used.
	//
	// Possible values:
	//   "VERSION_FORMAT_UNSPECIFIED" - An unspecified version format will
	// default to V2.
	//   "V2" - `LogEntry` version 2 format.
	//   "V1" - `LogEntry` version 1 format.
	OutputVersionFormat string `json:"outputVersionFormat,omitempty"`

	// WriterIdentity: Output only. The iam identity to which the
	// destination needs to grant write
	// access.  This may be a service account or a group.
	// Examples (Do not assume these specific values):
	//    "serviceAccount:cloud-logs@system.gserviceaccount.com"
	//    "group:cloud-logs@google.com"
	//
	//   For GCS destinations, the role "roles/owner" is required on the
	// bucket
	//   For Cloud Pubsub destinations, the role "roles/pubsub.publisher"
	// is
	//     required on the topic
	//   For BigQuery, the role "roles/editor" is required on the dataset
	WriterIdentity string `json:"writerIdentity,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Destination") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *LogSink) MarshalJSON() ([]byte, error) {
	type noMethod LogSink
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// MonitoredResource: An object representing a resource that can be used
// for monitoring, logging,
// billing, or other purposes. Examples include virtual machine
// instances,
// databases, and storage devices such as disks. The `type` field
// identifies a
// MonitoredResourceDescriptor object that describes the
// resource's
// schema. Information in the `labels` field identifies the actual
// resource and
// its attributes according to the schema. For example, a particular
// Compute
// Engine VM instance could be represented by the following object,
// because the
// MonitoredResourceDescriptor for "gce_instance" has
// labels
// "instance_id" and "zone":
//
//     { "type": "gce_instance",
//       "labels": { "instance_id": "12345678901234",
//                   "zone": "us-central1-a" }}
type MonitoredResource struct {
	// Labels: Required. Values for all of the labels listed in the
	// associated monitored
	// resource descriptor. For example, Cloud SQL databases use the
	// labels
	// "database_id" and "zone".
	Labels map[string]string `json:"labels,omitempty"`

	// Type: Required. The monitored resource type. This field must
	// match
	// the `type` field of a MonitoredResourceDescriptor object.
	// For
	// example, the type of a Cloud SQL database is "cloudsql_database".
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Labels") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *MonitoredResource) MarshalJSON() ([]byte, error) {
	type noMethod MonitoredResource
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// MonitoredResourceDescriptor: An object that describes the schema of a
// MonitoredResource object using a
// type name and a set of labels.  For example, the monitored
// resource
// descriptor for Google Compute Engine VM instances has a type
// of
// "gce_instance" and specifies the use of the labels "instance_id"
// and
// "zone" to identify particular VM instances.
//
// Different APIs can support different monitored resource types. APIs
// generally
// provide a `list` method that returns the monitored resource
// descriptors used
// by the API.
type MonitoredResourceDescriptor struct {
	// Description: Optional. A detailed description of the monitored
	// resource type that might
	// be used in documentation.
	Description string `json:"description,omitempty"`

	// DisplayName: Optional. A concise name for the monitored resource type
	// that might be
	// displayed in user interfaces. It should be a Title Cased Noun
	// Phrase,
	// without any article or other determiners. For example,
	// "Google Cloud SQL Database".
	DisplayName string `json:"displayName,omitempty"`

	// Labels: Required. A set of labels used to describe instances of this
	// monitored
	// resource type. For example, an individual Google Cloud SQL database
	// is
	// identified by values for the labels "database_id" and "zone".
	Labels []*LabelDescriptor `json:"labels,omitempty"`

	// Name: Optional. The resource name of the monitored resource
	// descriptor:
	// "projects/{project_id}/monitoredResourceDescriptors/{type
	// }" where
	// {type} is the value of the `type` field in this object
	// and
	// {project_id} is a project ID that provides API-specific context
	// for
	// accessing the type.  APIs that do not use project information can use
	// the
	// resource name format "monitoredResourceDescriptors/{type}".
	Name string `json:"name,omitempty"`

	// Type: Required. The monitored resource type. For example, the
	// type
	// "cloudsql_database" represents databases in Google Cloud SQL.
	// The maximum length of this value is 256 characters.
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Description") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *MonitoredResourceDescriptor) MarshalJSON() ([]byte, error) {
	type noMethod MonitoredResourceDescriptor
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// ReadLogEntriesRequest: The parameters to `ReadLogEntries`.
// There are two different use cases for streaming:
//
// 1.  To return a very large result set. The request eventually
//     completes when all entries have been returned.
// 2.  To "tail" a log stream, returning new entries as they arrive.
//     In this case, the request never completes.
//
// Only the first use case is supported.
type ReadLogEntriesRequest struct {
	// BatchSize: Optional. The maximum number of log entries to return in
	// each streamed
	// response.  Non-positive values are ignored and the default is 10.
	BatchSize int64 `json:"batchSize,omitempty"`

	// Filter: Optional. An [advanced logs
	// filter](/logging/docs/view/advanced_filters).
	// The response includes only entries that match the filter.
	// If `filter` is empty, then all entries in all logs are retrieved.
	Filter string `json:"filter,omitempty"`

	// MaxResponseInterval: Optional. The maximum time between successive
	// streamed responses.
	// The default value is `1s`, one second.  A response will be sent
	// at
	// least this often, even if it contains no log entries, provided
	// that there has been some progress as evidenced by a change to
	// the
	// `resumeToken` provided in the response.
	MaxResponseInterval string `json:"maxResponseInterval,omitempty"`

	// OrderBy: Optional. How the results should be sorted.  Presently, the
	// only permitted
	// values are "timestamp asc" (default) and "timestamp desc". The
	// first
	// option returns entries in order of increasing values
	// of
	// `LogEntry.timestamp` (oldest first), and the second option returns
	// entries
	// in order of decreasing timestamps (newest first).  Entries with
	// equal
	// timestamps will be returned in order of `LogEntry.insertId`.
	OrderBy string `json:"orderBy,omitempty"`

	// ProjectIds: Required. A list of project ids from which to retrieve
	// log entries.
	// Example: "my-project-id".
	ProjectIds []string `json:"projectIds,omitempty"`

	// ResumeToken: Optional. If the `resumeToken` parameter is supplied,
	// then read streaming
	// begins at the location represented by this `resumeToken` from a
	// previous
	// response. This value might be useful if streaming is stops
	// prematurely.
	ResumeToken string `json:"resumeToken,omitempty"`

	// ForceSendFields is a list of field names (e.g. "BatchSize") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *ReadLogEntriesRequest) MarshalJSON() ([]byte, error) {
	type noMethod ReadLogEntriesRequest
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// ReadLogEntriesResponse: Result returned from `ReadLogEntries`.
type ReadLogEntriesResponse struct {
	// Entries: A list of log entries - may be empty.
	Entries []*LogEntry `json:"entries,omitempty"`

	// LastObservedEntryTimestamp: The timestamp of the last log entry that
	// was examined before returning this
	// response. This can be used to observe progress between successive
	// queries,
	// in particular when only a resume token is returned.
	// Deprecated: use searched_through_timestamp.
	LastObservedEntryTimestamp string `json:"lastObservedEntryTimestamp,omitempty"`

	// ResumeToken: A token to use to resume from this position of the
	// stream.
	// Note that even if there are no entries, it might still be possible
	// to continue from this point at some later time.
	ResumeToken string `json:"resumeToken,omitempty"`

	// SearchedThroughTimestamp: The furthest point in time through which
	// the search has progressed. All
	// future entries returned using resume_token are guaranteed to have
	// a
	// timestamp at or past this point in time in the direction of the
	// search.
	// This can be used to observe progress between successive queries,
	// in
	// particular when only a page token is returned.
	SearchedThroughTimestamp string `json:"searchedThroughTimestamp,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Entries") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *ReadLogEntriesResponse) MarshalJSON() ([]byte, error) {
	type noMethod ReadLogEntriesResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// RequestLog: Complete log information about a single HTTP request to
// an App Engine
// application.
type RequestLog struct {
	// AppEngineRelease: App Engine release version.
	AppEngineRelease string `json:"appEngineRelease,omitempty"`

	// AppId: Application that handled this request.
	AppId string `json:"appId,omitempty"`

	// Cost: An indication of the relative cost of serving this request.
	Cost float64 `json:"cost,omitempty"`

	// EndTime: Time when the request finished.
	EndTime string `json:"endTime,omitempty"`

	// Finished: Whether this request is finished or active.
	Finished bool `json:"finished,omitempty"`

	// First: Whether this is the first RequestLog entry for this request.
	// If an active
	// request has several RequestLog entries written to Cloud Logging, this
	// field
	// will be set for one of them.
	First bool `json:"first,omitempty"`

	// Host: Internet host and port number of the resource being requested.
	Host string `json:"host,omitempty"`

	// HttpVersion: HTTP version of request. Example: "HTTP/1.1".
	HttpVersion string `json:"httpVersion,omitempty"`

	// InstanceId: An identifier for the instance that handled the request.
	InstanceId string `json:"instanceId,omitempty"`

	// InstanceIndex: If the instance processing this request belongs to a
	// manually scaled
	// module, then this is the 0-based index of the instance. Otherwise,
	// this
	// value is -1.
	InstanceIndex int64 `json:"instanceIndex,omitempty"`

	// Ip: Origin IP address.
	Ip string `json:"ip,omitempty"`

	// Latency: Latency of the request.
	Latency string `json:"latency,omitempty"`

	// Line: A list of log lines emitted by the application while serving
	// this request.
	Line []*LogLine `json:"line,omitempty"`

	// MegaCycles: Number of CPU megacycles used to process request.
	MegaCycles int64 `json:"megaCycles,omitempty,string"`

	// Method: Request method. Example: "GET", "HEAD", "PUT",
	// "POST", "DELETE".
	Method string `json:"method,omitempty"`

	// ModuleId: Module of the application that handled this request.
	ModuleId string `json:"moduleId,omitempty"`

	// Nickname: The logged-in user who made the request.
	//
	// Most likely, this is the part of the user's email before the `@`
	// sign.  The
	// field value is the same for different requests from the same user,
	// but
	// different users can have similar names.  This information is
	// also
	// available to the application via the App Engine Users API.
	//
	// This field will be populated starting with App Engine 1.9.21.
	Nickname string `json:"nickname,omitempty"`

	// PendingTime: Time this request spent in the pending request queue.
	PendingTime string `json:"pendingTime,omitempty"`

	// Referrer: Referrer URL of request.
	Referrer string `json:"referrer,omitempty"`

	// RequestId: Globally unique identifier for a request, which is based
	// on the request
	// start time.  Request IDs for requests which started later will
	// compare
	// greater as strings than those for requests which started earlier.
	RequestId string `json:"requestId,omitempty"`

	// Resource: Contains the path and query portion of the URL that was
	// requested. For
	// example, if the URL was "http://example.com/app?name=val", the
	// resource
	// would be "/app?name=val".  The fragment identifier, which is
	// identified by
	// the `#` character, is not included.
	Resource string `json:"resource,omitempty"`

	// ResponseSize: Size in bytes sent back to client by request.
	ResponseSize int64 `json:"responseSize,omitempty,string"`

	// SourceReference: Source code for the application that handled this
	// request. There can be
	// more than one source reference per deployed application if source
	// code is
	// distributed among multiple repositories.
	SourceReference []*SourceReference `json:"sourceReference,omitempty"`

	// StartTime: Time when the request started.
	StartTime string `json:"startTime,omitempty"`

	// Status: HTTP response status code. Example: 200, 404.
	Status int64 `json:"status,omitempty"`

	// TaskName: Task name of the request, in the case of an offline
	// request.
	TaskName string `json:"taskName,omitempty"`

	// TaskQueueName: Queue name of the request, in the case of an offline
	// request.
	TaskQueueName string `json:"taskQueueName,omitempty"`

	// TraceId: Cloud Trace identifier for this request.
	TraceId string `json:"traceId,omitempty"`

	// UrlMapEntry: File or class that handled the request.
	UrlMapEntry string `json:"urlMapEntry,omitempty"`

	// UserAgent: User agent that made the request.
	UserAgent string `json:"userAgent,omitempty"`

	// VersionId: Version of the application that handled this request.
	VersionId string `json:"versionId,omitempty"`

	// WasLoadingRequest: Whether this was a loading request for the
	// instance.
	WasLoadingRequest bool `json:"wasLoadingRequest,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AppEngineRelease") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *RequestLog) MarshalJSON() ([]byte, error) {
	type noMethod RequestLog
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// ResourceKeys: Output only. Describes resource keys for log entries.
type ResourceKeys struct {
	// DisplayName: Displayable name for this type that can be presented in
	// a UI.
	DisplayName string `json:"displayName,omitempty"`

	// Keys: A list of the names of the keys used to index and
	// label
	// individual log entries associated with this resource type.
	//
	//     [ "module_id",
	//       "version_id" ]
	Keys []string `json:"keys,omitempty"`

	// Type: The type of the resource - e.g. "gce_instance"
	// This value can appear in the `LogEntry.resource.type` field of
	// log entries
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DisplayName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *ResourceKeys) MarshalJSON() ([]byte, error) {
	type noMethod ResourceKeys
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// SourceLocation: Specifies a location in a source code file.
type SourceLocation struct {
	// File: Source file name. Depending on the runtime environment, this
	// might be a
	// simple name or a fully-qualified name.
	File string `json:"file,omitempty"`

	// FunctionName: Human-readable name of the function or method being
	// invoked, with optional
	// context such as the class or package name. This information is used
	// in
	// contexts such as the logs viewer, where a file and line number are
	// less
	// meaningful. The format can vary by language. For
	// example:
	// `qual.if.ied.Class.method` (Java), `dir/package.func` (Go),
	// `function`
	// (Python).
	FunctionName string `json:"functionName,omitempty"`

	// Line: Line within the source file.
	Line int64 `json:"line,omitempty,string"`

	// ForceSendFields is a list of field names (e.g. "File") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *SourceLocation) MarshalJSON() ([]byte, error) {
	type noMethod SourceLocation
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// SourceReference: A reference to a particular snapshot of the source
// tree used to build and
// deploy an application.
type SourceReference struct {
	// Repository: Optional. A URI string identifying the
	// repository.
	// Example: "https://github.com/GoogleCloudPlatform/kubernetes.git"
	Repository string `json:"repository,omitempty"`

	// RevisionId: The canonical and persistent identifier of the deployed
	// revision.
	// Example (git): "0035781c50ec7aa23385dc841529ce8a4b70db1b"
	RevisionId string `json:"revisionId,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Repository") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *SourceReference) MarshalJSON() ([]byte, error) {
	type noMethod SourceReference
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

// Usage: Describes logs usage over a period of time.
type Usage struct {
	// ByteCount: The volume of ingested logs, in bytes.
	ByteCount int64 `json:"byteCount,omitempty,string"`

	// ByteQuota: The allowed free quota, also in bytes. Note that the quota
	// for Free Tier
	// is monthly, while for Premium Tier, it's calculated hourly.
	ByteQuota int64 `json:"byteQuota,omitempty,string"`

	// EndTime: Exclusive. End time of the usage interval.
	EndTime string `json:"endTime,omitempty"`

	// StartTime: Inclusive. Start time of the usage interval.
	StartTime string `json:"startTime,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ByteCount") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Usage) MarshalJSON() ([]byte, error) {
	type noMethod Usage
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// WriteLogEntriesRequest: The parameters to WriteLogEntries.
type WriteLogEntriesRequest struct {
	// Entries: Required. The log entries to write. The log entries must
	// have values for
	// all required fields.
	//
	// To improve throughput and to avoid exceeding the
	// [quota limit](/logging/quota-policy) for calls to
	// `entries.write`,
	// use this field to write multiple log entries at once rather
	// than
	// calling this method for each log entry.
	Entries []*LogEntry `json:"entries,omitempty"`

	// Labels: Optional. User-defined `key:value` items that are added
	// to
	// the `labels` field of each log entry in `entries`, except when a
	// log
	// entry specifies its own `key:value` item with the same key.
	// Example: `{ "size": "large", "color":"red" }`
	Labels map[string]string `json:"labels,omitempty"`

	// LogName: Optional. A default log resource name for those log entries
	// in `entries`
	// that do not specify their own `logName`.
	// Example:
	// "projects/my-project/logs/syslog".  See
	// LogEntry.
	LogName string `json:"logName,omitempty"`

	// PartialSuccess: Optional. Whether valid entries should be written
	// even if some other
	// entries fail due to INVALID_ARGUMENT or PERMISSION_DENIED errors. If
	// any
	// entry is not written, the response status will be the error
	// associated
	// with one of the failed entries and include error details in the form
	// of
	// WriteLogEntriesPartialErrors.
	PartialSuccess bool `json:"partialSuccess,omitempty"`

	// Resource: Optional. A default monitored resource for those log
	// entries in `entries`
	// that do not specify their own `resource`.
	Resource *MonitoredResource `json:"resource,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Entries") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *WriteLogEntriesRequest) MarshalJSON() ([]byte, error) {
	type noMethod WriteLogEntriesRequest
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// WriteLogEntriesResponse: Result returned from WriteLogEntries.
// empty
type WriteLogEntriesResponse struct {
	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`
}

// method id "logging.billingAccounts.logs.delete":

type BillingAccountsLogsDeleteCall struct {
	s                 *Service
	billingAccountsId string
	logsId            string
	urlParams_        gensupport.URLParams
	ctx_              context.Context
}

// Delete: Deletes a log and all its log entries.
// The log will reappear if it receives new entries.
func (r *BillingAccountsLogsService) Delete(billingAccountsId string, logsId string) *BillingAccountsLogsDeleteCall {
	c := &BillingAccountsLogsDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.billingAccountsId = billingAccountsId
	c.logsId = logsId
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BillingAccountsLogsDeleteCall) Fields(s ...googleapi.Field) *BillingAccountsLogsDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *BillingAccountsLogsDeleteCall) Context(ctx context.Context) *BillingAccountsLogsDeleteCall {
	c.ctx_ = ctx
	return c
}

func (c *BillingAccountsLogsDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/billingAccounts/{billingAccountsId}/logs/{logsId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"billingAccountsId": c.billingAccountsId,
		"logsId":            c.logsId,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "logging.billingAccounts.logs.delete" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *BillingAccountsLogsDeleteCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
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
	//   "description": "Deletes a log and all its log entries.\nThe log will reappear if it receives new entries.",
	//   "flatPath": "v2beta1/billingAccounts/{billingAccountsId}/logs/{logsId}",
	//   "httpMethod": "DELETE",
	//   "id": "logging.billingAccounts.logs.delete",
	//   "parameterOrder": [
	//     "billingAccountsId",
	//     "logsId"
	//   ],
	//   "parameters": {
	//     "billingAccountsId": {
	//       "description": "Part of `logName`. Required. The resource name of the log to delete.  Example:\n`\"projects/my-project/logs/syslog\"`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "logsId": {
	//       "description": "Part of `logName`. See documentation of `billingAccountsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/billingAccounts/{billingAccountsId}/logs/{logsId}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/logging.admin"
	//   ]
	// }

}

// method id "logging.billingAccounts.logs.list":

type BillingAccountsLogsListCall struct {
	s                 *Service
	billingAccountsId string
	urlParams_        gensupport.URLParams
	ifNoneMatch_      string
	ctx_              context.Context
}

// List: Lists the logs in the project.
// Only logs that have entries are listed.
func (r *BillingAccountsLogsService) List(billingAccountsId string) *BillingAccountsLogsListCall {
	c := &BillingAccountsLogsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.billingAccountsId = billingAccountsId
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of results to return from this request.
// Non-positive values are ignored.  The presence of `nextPageToken` in
// the
// response indicates that more results might be available.
func (c *BillingAccountsLogsListCall) PageSize(pageSize int64) *BillingAccountsLogsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": If present, then
// retrieve the next batch of results from the
// preceding call to this method.  `pageToken` must be the value
// of
// `nextPageToken` from the previous response.  The values of other
// method
// parameters should be identical to those in the previous call.
func (c *BillingAccountsLogsListCall) PageToken(pageToken string) *BillingAccountsLogsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// ResourceIndexPrefix sets the optional parameter
// "resourceIndexPrefix": The purpose of this field is to restrict the
// listed logs to those
// with entries of a certain kind.
// If `resource_type` is the name of a resource type,
// then this field may contain values for the log resource type's
// indexes.
// Only logs that have entries whose indexes include the values are
// listed.
// The format for this field is "/val1/val2.../valN", where `val1` is
// a
// value for the first index, `val2` for the second index, etc.
// An empty value (a single slash) for an index matches all values,
// and you can omit values for later indexes entirely.
//  The maximum number of results to return from this request.
func (c *BillingAccountsLogsListCall) ResourceIndexPrefix(resourceIndexPrefix string) *BillingAccountsLogsListCall {
	c.urlParams_.Set("resourceIndexPrefix", resourceIndexPrefix)
	return c
}

// ResourceType sets the optional parameter "resourceType": If not
// empty, this field must be a resource type such as
// "gce_instance`. Only logs associated with that resource type are
// listed.
func (c *BillingAccountsLogsListCall) ResourceType(resourceType string) *BillingAccountsLogsListCall {
	c.urlParams_.Set("resourceType", resourceType)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BillingAccountsLogsListCall) Fields(s ...googleapi.Field) *BillingAccountsLogsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *BillingAccountsLogsListCall) IfNoneMatch(entityTag string) *BillingAccountsLogsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *BillingAccountsLogsListCall) Context(ctx context.Context) *BillingAccountsLogsListCall {
	c.ctx_ = ctx
	return c
}

func (c *BillingAccountsLogsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/billingAccounts/{billingAccountsId}/logs")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"billingAccountsId": c.billingAccountsId,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "logging.billingAccounts.logs.list" call.
// Exactly one of *ListLogsResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListLogsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *BillingAccountsLogsListCall) Do(opts ...googleapi.CallOption) (*ListLogsResponse, error) {
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
	ret := &ListLogsResponse{
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
	//   "description": "Lists the logs in the project.\nOnly logs that have entries are listed.",
	//   "flatPath": "v2beta1/billingAccounts/{billingAccountsId}/logs",
	//   "httpMethod": "GET",
	//   "id": "logging.billingAccounts.logs.list",
	//   "parameterOrder": [
	//     "billingAccountsId"
	//   ],
	//   "parameters": {
	//     "billingAccountsId": {
	//       "description": "Part of `parent`. The resource name of the entity whose logs are requested.\nIf both `resource_type` and `resourceIndexPrefix` are empty, then\nall logs with entries in this entity are listed.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Optional. The maximum number of results to return from this request.\nNon-positive values are ignored.  The presence of `nextPageToken` in the\nresponse indicates that more results might be available.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. If present, then retrieve the next batch of results from the\npreceding call to this method.  `pageToken` must be the value of\n`nextPageToken` from the previous response.  The values of other method\nparameters should be identical to those in the previous call.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "resourceIndexPrefix": {
	//       "description": "The purpose of this field is to restrict the listed logs to those\nwith entries of a certain kind.\nIf `resource_type` is the name of a resource type,\nthen this field may contain values for the log resource type's indexes.\nOnly logs that have entries whose indexes include the values are listed.\nThe format for this field is `\"/val1/val2.../valN\"`, where `val1` is a\nvalue for the first index, `val2` for the second index, etc.\nAn empty value (a single slash) for an index matches all values,\nand you can omit values for later indexes entirely.\nOptional. The maximum number of results to return from this request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "resourceType": {
	//       "description": "If not empty, this field must be a resource type such as\n`\"gce_instance`. Only logs associated with that resource type are listed.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/billingAccounts/{billingAccountsId}/logs",
	//   "response": {
	//     "$ref": "ListLogsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only",
	//     "https://www.googleapis.com/auth/logging.admin",
	//     "https://www.googleapis.com/auth/logging.read"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *BillingAccountsLogsListCall) Pages(ctx context.Context, f func(*ListLogsResponse) error) error {
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

// method id "logging.billingAccounts.resourceKeys.list":

type BillingAccountsResourceKeysListCall struct {
	s                 *Service
	billingAccountsId string
	urlParams_        gensupport.URLParams
	ifNoneMatch_      string
	ctx_              context.Context
}

// List: Lists the resource keys that have log entries in this project.
func (r *BillingAccountsResourceKeysService) List(billingAccountsId string) *BillingAccountsResourceKeysListCall {
	c := &BillingAccountsResourceKeysListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.billingAccountsId = billingAccountsId
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of results to return from this request.
// Non-positive values are ignored.  The presence of `nextPageToken` in
// the
// response indicates that more results might be available.
func (c *BillingAccountsResourceKeysListCall) PageSize(pageSize int64) *BillingAccountsResourceKeysListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": If present, then
// retrieve the next batch of results from the
// preceding call to this method.  `pageToken` must be the value
// of
// `nextPageToken` from the previous response.  The values of other
// method
// parameters should be identical to those in the previous call.
func (c *BillingAccountsResourceKeysListCall) PageToken(pageToken string) *BillingAccountsResourceKeysListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BillingAccountsResourceKeysListCall) Fields(s ...googleapi.Field) *BillingAccountsResourceKeysListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *BillingAccountsResourceKeysListCall) IfNoneMatch(entityTag string) *BillingAccountsResourceKeysListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *BillingAccountsResourceKeysListCall) Context(ctx context.Context) *BillingAccountsResourceKeysListCall {
	c.ctx_ = ctx
	return c
}

func (c *BillingAccountsResourceKeysListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/billingAccounts/{billingAccountsId}/resourceKeys")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"billingAccountsId": c.billingAccountsId,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "logging.billingAccounts.resourceKeys.list" call.
// Exactly one of *ListResourceKeysResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *ListResourceKeysResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *BillingAccountsResourceKeysListCall) Do(opts ...googleapi.CallOption) (*ListResourceKeysResponse, error) {
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
	ret := &ListResourceKeysResponse{
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
	//   "description": "Lists the resource keys that have log entries in this project.",
	//   "flatPath": "v2beta1/billingAccounts/{billingAccountsId}/resourceKeys",
	//   "httpMethod": "GET",
	//   "id": "logging.billingAccounts.resourceKeys.list",
	//   "parameterOrder": [
	//     "billingAccountsId"
	//   ],
	//   "parameters": {
	//     "billingAccountsId": {
	//       "description": "Part of `parent`. The resource name of the entity whose reource keys are to be listed.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Optional. The maximum number of results to return from this request.\nNon-positive values are ignored.  The presence of `nextPageToken` in the\nresponse indicates that more results might be available.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. If present, then retrieve the next batch of results from the\npreceding call to this method.  `pageToken` must be the value of\n`nextPageToken` from the previous response.  The values of other method\nparameters should be identical to those in the previous call.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/billingAccounts/{billingAccountsId}/resourceKeys",
	//   "response": {
	//     "$ref": "ListResourceKeysResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only",
	//     "https://www.googleapis.com/auth/logging.admin",
	//     "https://www.googleapis.com/auth/logging.read"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *BillingAccountsResourceKeysListCall) Pages(ctx context.Context, f func(*ListResourceKeysResponse) error) error {
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

// method id "logging.billingAccounts.resourceTypes.values.list":

type BillingAccountsResourceTypesValuesListCall struct {
	s                 *Service
	billingAccountsId string
	resourceTypesId   string
	urlParams_        gensupport.URLParams
	ifNoneMatch_      string
	ctx_              context.Context
}

// List: Lists the current index values for a log resource type.
func (r *BillingAccountsResourceTypesValuesService) List(billingAccountsId string, resourceTypesId string) *BillingAccountsResourceTypesValuesListCall {
	c := &BillingAccountsResourceTypesValuesListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.billingAccountsId = billingAccountsId
	c.resourceTypesId = resourceTypesId
	return c
}

// Depth sets the optional parameter "depth": A non-negative integer
// that limits the number of levels of the index
// hierarchy that are returned.
// If `depth` is 1 (default), only the first index key value is
// returned.
// If `depth` is 2, both primary and secondary key values are
// returned.
// If `depth` is 0, the depth is the number of slash-separators in
// the
// `indexPrefix` field, not counting a slash appearing as the last
// character
// of the prefix.
// If the `indexPrefix` field is empty, the default depth is 1.
// It is an error for `depth` to be any positive value
// less than the number of components in `indexPrefix`.
func (c *BillingAccountsResourceTypesValuesListCall) Depth(depth int64) *BillingAccountsResourceTypesValuesListCall {
	c.urlParams_.Set("depth", fmt.Sprint(depth))
	return c
}

// IndexPrefix sets the optional parameter "indexPrefix": Restricts the
// index values returned to be those with a specified prefix
// for each index key.
// This field has the form "/prefix1/prefix2/...", in order
// corresponding
// to the `ResourceKeys indexKeys`.
// Non-empty prefixes must begin with `/`.
// For example, App Engine's two keys are the module ID and the version
// ID.
// Following is the effect of using various values for `indexPrefix`:
//
// +  "/Mod/" retrieves `/Mod/10` and `/Mod/11` but not `/ModA/10`.
// +  "/Mod` retrieves `/Mod/10`, `/Mod/11` and `/ModA/10` but not
// `/XXX/33`.
// +  "/Mod/1" retrieves `/Mod/10` and `/Mod/11` but not `/ModA/10`.
// +  "/Mod/10/" retrieves `/Mod/10` only.
// +  An empty prefix or "/" retrieves all values.
func (c *BillingAccountsResourceTypesValuesListCall) IndexPrefix(indexPrefix string) *BillingAccountsResourceTypesValuesListCall {
	c.urlParams_.Set("indexPrefix", indexPrefix)
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of results to return from this request.
// Non-positive values are ignored.  The presence of `nextPageToken` in
// the
// response indicates that more results might be available.
func (c *BillingAccountsResourceTypesValuesListCall) PageSize(pageSize int64) *BillingAccountsResourceTypesValuesListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": If present, then
// retrieve the next batch of results from the
// preceding call to this method.  `pageToken` must be the value
// of
// `nextPageToken` from the previous response.  The values of other
// method
// parameters should be identical to those in the previous call.
func (c *BillingAccountsResourceTypesValuesListCall) PageToken(pageToken string) *BillingAccountsResourceTypesValuesListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BillingAccountsResourceTypesValuesListCall) Fields(s ...googleapi.Field) *BillingAccountsResourceTypesValuesListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *BillingAccountsResourceTypesValuesListCall) IfNoneMatch(entityTag string) *BillingAccountsResourceTypesValuesListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *BillingAccountsResourceTypesValuesListCall) Context(ctx context.Context) *BillingAccountsResourceTypesValuesListCall {
	c.ctx_ = ctx
	return c
}

func (c *BillingAccountsResourceTypesValuesListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/billingAccounts/{billingAccountsId}/resourceTypes/{resourceTypesId}/values")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"billingAccountsId": c.billingAccountsId,
		"resourceTypesId":   c.resourceTypesId,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "logging.billingAccounts.resourceTypes.values.list" call.
// Exactly one of *ListResourceValuesResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *ListResourceValuesResponse.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *BillingAccountsResourceTypesValuesListCall) Do(opts ...googleapi.CallOption) (*ListResourceValuesResponse, error) {
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
	ret := &ListResourceValuesResponse{
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
	//   "description": "Lists the current index values for a log resource type.",
	//   "flatPath": "v2beta1/billingAccounts/{billingAccountsId}/resourceTypes/{resourceTypesId}/values",
	//   "httpMethod": "GET",
	//   "id": "logging.billingAccounts.resourceTypes.values.list",
	//   "parameterOrder": [
	//     "billingAccountsId",
	//     "resourceTypesId"
	//   ],
	//   "parameters": {
	//     "billingAccountsId": {
	//       "description": "Part of `parent`. The resource name of a resource type whose indexes are requested.\nExample:\n`\"projects/my-project-id/resourceTypes/gae_app\"`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "depth": {
	//       "description": "A non-negative integer that limits the number of levels of the index\nhierarchy that are returned.\nIf `depth` is 1 (default), only the first index key value is returned.\nIf `depth` is 2, both primary and secondary key values are returned.\nIf `depth` is 0, the depth is the number of slash-separators in the\n`indexPrefix` field, not counting a slash appearing as the last character\nof the prefix.\nIf the `indexPrefix` field is empty, the default depth is 1.\nIt is an error for `depth` to be any positive value\nless than the number of components in `indexPrefix`.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "indexPrefix": {
	//       "description": "Restricts the index values returned to be those with a specified prefix\nfor each index key.\nThis field has the form `\"/prefix1/prefix2/...\"`, in order corresponding\nto the `ResourceKeys indexKeys`.\nNon-empty prefixes must begin with `/`.\nFor example, App Engine's two keys are the module ID and the version ID.\nFollowing is the effect of using various values for `indexPrefix`:\n\n+  `\"/Mod/\"` retrieves `/Mod/10` and `/Mod/11` but not `/ModA/10`.\n+  `\"/Mod` retrieves `/Mod/10`, `/Mod/11` and `/ModA/10` but not `/XXX/33`.\n+  `\"/Mod/1\"` retrieves `/Mod/10` and `/Mod/11` but not `/ModA/10`.\n+  `\"/Mod/10/\"` retrieves `/Mod/10` only.\n+  An empty prefix or `\"/\"` retrieves all values.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Optional. The maximum number of results to return from this request.\nNon-positive values are ignored.  The presence of `nextPageToken` in the\nresponse indicates that more results might be available.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. If present, then retrieve the next batch of results from the\npreceding call to this method.  `pageToken` must be the value of\n`nextPageToken` from the previous response.  The values of other method\nparameters should be identical to those in the previous call.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "resourceTypesId": {
	//       "description": "Part of `parent`. See documentation of `billingAccountsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/billingAccounts/{billingAccountsId}/resourceTypes/{resourceTypesId}/values",
	//   "response": {
	//     "$ref": "ListResourceValuesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only",
	//     "https://www.googleapis.com/auth/logging.admin",
	//     "https://www.googleapis.com/auth/logging.read"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *BillingAccountsResourceTypesValuesListCall) Pages(ctx context.Context, f func(*ListResourceValuesResponse) error) error {
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

// method id "logging.entries.list":

type EntriesListCall struct {
	s                     *Service
	listlogentriesrequest *ListLogEntriesRequest
	urlParams_            gensupport.URLParams
	ctx_                  context.Context
}

// List: Lists log entries.  Use this method to retrieve log entries
// from Cloud
// Logging.  For ways to export log entries, see
// [Exporting Logs](/logging/docs/export).
func (r *EntriesService) List(listlogentriesrequest *ListLogEntriesRequest) *EntriesListCall {
	c := &EntriesListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.listlogentriesrequest = listlogentriesrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *EntriesListCall) Fields(s ...googleapi.Field) *EntriesListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *EntriesListCall) Context(ctx context.Context) *EntriesListCall {
	c.ctx_ = ctx
	return c
}

func (c *EntriesListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.listlogentriesrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/entries:list")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.SetOpaque(req.URL)
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "logging.entries.list" call.
// Exactly one of *ListLogEntriesResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListLogEntriesResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *EntriesListCall) Do(opts ...googleapi.CallOption) (*ListLogEntriesResponse, error) {
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
	ret := &ListLogEntriesResponse{
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
	//   "description": "Lists log entries.  Use this method to retrieve log entries from Cloud\nLogging.  For ways to export log entries, see\n[Exporting Logs](/logging/docs/export).",
	//   "flatPath": "v2beta1/entries:list",
	//   "httpMethod": "POST",
	//   "id": "logging.entries.list",
	//   "parameterOrder": [],
	//   "parameters": {},
	//   "path": "v2beta1/entries:list",
	//   "request": {
	//     "$ref": "ListLogEntriesRequest"
	//   },
	//   "response": {
	//     "$ref": "ListLogEntriesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only",
	//     "https://www.googleapis.com/auth/logging.admin",
	//     "https://www.googleapis.com/auth/logging.read"
	//   ]
	// }

}

// method id "logging.entries.read":

type EntriesReadCall struct {
	s                     *Service
	readlogentriesrequest *ReadLogEntriesRequest
	urlParams_            gensupport.URLParams
	ctx_                  context.Context
}

// Read: Streaming read of log entries.  Similar to `List`, this method
// is intended
// for a large volume of log entries.
func (r *EntriesService) Read(readlogentriesrequest *ReadLogEntriesRequest) *EntriesReadCall {
	c := &EntriesReadCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.readlogentriesrequest = readlogentriesrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *EntriesReadCall) Fields(s ...googleapi.Field) *EntriesReadCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *EntriesReadCall) Context(ctx context.Context) *EntriesReadCall {
	c.ctx_ = ctx
	return c
}

func (c *EntriesReadCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.readlogentriesrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/entries:read")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.SetOpaque(req.URL)
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "logging.entries.read" call.
// Exactly one of *ReadLogEntriesResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ReadLogEntriesResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *EntriesReadCall) Do(opts ...googleapi.CallOption) (*ReadLogEntriesResponse, error) {
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
	ret := &ReadLogEntriesResponse{
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
	//   "description": "Streaming read of log entries.  Similar to `List`, this method is intended\nfor a large volume of log entries.",
	//   "flatPath": "v2beta1/entries:read",
	//   "httpMethod": "POST",
	//   "id": "logging.entries.read",
	//   "parameterOrder": [],
	//   "parameters": {},
	//   "path": "v2beta1/entries:read",
	//   "request": {
	//     "$ref": "ReadLogEntriesRequest"
	//   },
	//   "response": {
	//     "$ref": "ReadLogEntriesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only",
	//     "https://www.googleapis.com/auth/logging.admin",
	//     "https://www.googleapis.com/auth/logging.read"
	//   ]
	// }

}

// method id "logging.entries.write":

type EntriesWriteCall struct {
	s                      *Service
	writelogentriesrequest *WriteLogEntriesRequest
	urlParams_             gensupport.URLParams
	ctx_                   context.Context
}

// Write: Writes log entries to Stackdriver Logging.  All log entries
// are
// written by this method.
func (r *EntriesService) Write(writelogentriesrequest *WriteLogEntriesRequest) *EntriesWriteCall {
	c := &EntriesWriteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.writelogentriesrequest = writelogentriesrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *EntriesWriteCall) Fields(s ...googleapi.Field) *EntriesWriteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *EntriesWriteCall) Context(ctx context.Context) *EntriesWriteCall {
	c.ctx_ = ctx
	return c
}

func (c *EntriesWriteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.writelogentriesrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/entries:write")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.SetOpaque(req.URL)
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "logging.entries.write" call.
// Exactly one of *WriteLogEntriesResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *WriteLogEntriesResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *EntriesWriteCall) Do(opts ...googleapi.CallOption) (*WriteLogEntriesResponse, error) {
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
	ret := &WriteLogEntriesResponse{
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
	//   "description": "Writes log entries to Stackdriver Logging.  All log entries are\nwritten by this method.",
	//   "flatPath": "v2beta1/entries:write",
	//   "httpMethod": "POST",
	//   "id": "logging.entries.write",
	//   "parameterOrder": [],
	//   "parameters": {},
	//   "path": "v2beta1/entries:write",
	//   "request": {
	//     "$ref": "WriteLogEntriesRequest"
	//   },
	//   "response": {
	//     "$ref": "WriteLogEntriesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/logging.admin",
	//     "https://www.googleapis.com/auth/logging.write"
	//   ]
	// }

}

// method id "logging.monitoredResourceDescriptors.list":

type MonitoredResourceDescriptorsListCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// List: Lists the monitored resource descriptors used by Stackdriver
// Logging.
func (r *MonitoredResourceDescriptorsService) List() *MonitoredResourceDescriptorsListCall {
	c := &MonitoredResourceDescriptorsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of results to return from this request.
// Non-positive values are ignored.  The presence of `nextPageToken` in
// the
// response indicates that more results might be available.
func (c *MonitoredResourceDescriptorsListCall) PageSize(pageSize int64) *MonitoredResourceDescriptorsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": If present, then
// retrieve the next batch of results from the
// preceding call to this method.  `pageToken` must be the value
// of
// `nextPageToken` from the previous response.  The values of other
// method
// parameters should be identical to those in the previous call.
func (c *MonitoredResourceDescriptorsListCall) PageToken(pageToken string) *MonitoredResourceDescriptorsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *MonitoredResourceDescriptorsListCall) Fields(s ...googleapi.Field) *MonitoredResourceDescriptorsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *MonitoredResourceDescriptorsListCall) IfNoneMatch(entityTag string) *MonitoredResourceDescriptorsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *MonitoredResourceDescriptorsListCall) Context(ctx context.Context) *MonitoredResourceDescriptorsListCall {
	c.ctx_ = ctx
	return c
}

func (c *MonitoredResourceDescriptorsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/monitoredResourceDescriptors")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.SetOpaque(req.URL)
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "logging.monitoredResourceDescriptors.list" call.
// Exactly one of *ListMonitoredResourceDescriptorsResponse or error
// will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *ListMonitoredResourceDescriptorsResponse.ServerResponse.Header or
// (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *MonitoredResourceDescriptorsListCall) Do(opts ...googleapi.CallOption) (*ListMonitoredResourceDescriptorsResponse, error) {
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
	ret := &ListMonitoredResourceDescriptorsResponse{
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
	//   "description": "Lists the monitored resource descriptors used by Stackdriver Logging.",
	//   "flatPath": "v2beta1/monitoredResourceDescriptors",
	//   "httpMethod": "GET",
	//   "id": "logging.monitoredResourceDescriptors.list",
	//   "parameterOrder": [],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "Optional. The maximum number of results to return from this request.\nNon-positive values are ignored.  The presence of `nextPageToken` in the\nresponse indicates that more results might be available.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. If present, then retrieve the next batch of results from the\npreceding call to this method.  `pageToken` must be the value of\n`nextPageToken` from the previous response.  The values of other method\nparameters should be identical to those in the previous call.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/monitoredResourceDescriptors",
	//   "response": {
	//     "$ref": "ListMonitoredResourceDescriptorsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only",
	//     "https://www.googleapis.com/auth/logging.admin",
	//     "https://www.googleapis.com/auth/logging.read"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *MonitoredResourceDescriptorsListCall) Pages(ctx context.Context, f func(*ListMonitoredResourceDescriptorsResponse) error) error {
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

// method id "logging.organizations.logs.delete":

type OrganizationsLogsDeleteCall struct {
	s               *Service
	organizationsId string
	logsId          string
	urlParams_      gensupport.URLParams
	ctx_            context.Context
}

// Delete: Deletes a log and all its log entries.
// The log will reappear if it receives new entries.
func (r *OrganizationsLogsService) Delete(organizationsId string, logsId string) *OrganizationsLogsDeleteCall {
	c := &OrganizationsLogsDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.organizationsId = organizationsId
	c.logsId = logsId
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OrganizationsLogsDeleteCall) Fields(s ...googleapi.Field) *OrganizationsLogsDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *OrganizationsLogsDeleteCall) Context(ctx context.Context) *OrganizationsLogsDeleteCall {
	c.ctx_ = ctx
	return c
}

func (c *OrganizationsLogsDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/organizations/{organizationsId}/logs/{logsId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"organizationsId": c.organizationsId,
		"logsId":          c.logsId,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "logging.organizations.logs.delete" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *OrganizationsLogsDeleteCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
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
	//   "description": "Deletes a log and all its log entries.\nThe log will reappear if it receives new entries.",
	//   "flatPath": "v2beta1/organizations/{organizationsId}/logs/{logsId}",
	//   "httpMethod": "DELETE",
	//   "id": "logging.organizations.logs.delete",
	//   "parameterOrder": [
	//     "organizationsId",
	//     "logsId"
	//   ],
	//   "parameters": {
	//     "logsId": {
	//       "description": "Part of `logName`. See documentation of `organizationsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "organizationsId": {
	//       "description": "Part of `logName`. Required. The resource name of the log to delete.  Example:\n`\"projects/my-project/logs/syslog\"`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/organizations/{organizationsId}/logs/{logsId}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/logging.admin"
	//   ]
	// }

}

// method id "logging.organizations.logs.list":

type OrganizationsLogsListCall struct {
	s               *Service
	organizationsId string
	urlParams_      gensupport.URLParams
	ifNoneMatch_    string
	ctx_            context.Context
}

// List: Lists the logs in the project.
// Only logs that have entries are listed.
func (r *OrganizationsLogsService) List(organizationsId string) *OrganizationsLogsListCall {
	c := &OrganizationsLogsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.organizationsId = organizationsId
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of results to return from this request.
// Non-positive values are ignored.  The presence of `nextPageToken` in
// the
// response indicates that more results might be available.
func (c *OrganizationsLogsListCall) PageSize(pageSize int64) *OrganizationsLogsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": If present, then
// retrieve the next batch of results from the
// preceding call to this method.  `pageToken` must be the value
// of
// `nextPageToken` from the previous response.  The values of other
// method
// parameters should be identical to those in the previous call.
func (c *OrganizationsLogsListCall) PageToken(pageToken string) *OrganizationsLogsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// ResourceIndexPrefix sets the optional parameter
// "resourceIndexPrefix": The purpose of this field is to restrict the
// listed logs to those
// with entries of a certain kind.
// If `resource_type` is the name of a resource type,
// then this field may contain values for the log resource type's
// indexes.
// Only logs that have entries whose indexes include the values are
// listed.
// The format for this field is "/val1/val2.../valN", where `val1` is
// a
// value for the first index, `val2` for the second index, etc.
// An empty value (a single slash) for an index matches all values,
// and you can omit values for later indexes entirely.
//  The maximum number of results to return from this request.
func (c *OrganizationsLogsListCall) ResourceIndexPrefix(resourceIndexPrefix string) *OrganizationsLogsListCall {
	c.urlParams_.Set("resourceIndexPrefix", resourceIndexPrefix)
	return c
}

// ResourceType sets the optional parameter "resourceType": If not
// empty, this field must be a resource type such as
// "gce_instance`. Only logs associated with that resource type are
// listed.
func (c *OrganizationsLogsListCall) ResourceType(resourceType string) *OrganizationsLogsListCall {
	c.urlParams_.Set("resourceType", resourceType)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OrganizationsLogsListCall) Fields(s ...googleapi.Field) *OrganizationsLogsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *OrganizationsLogsListCall) IfNoneMatch(entityTag string) *OrganizationsLogsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *OrganizationsLogsListCall) Context(ctx context.Context) *OrganizationsLogsListCall {
	c.ctx_ = ctx
	return c
}

func (c *OrganizationsLogsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/organizations/{organizationsId}/logs")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"organizationsId": c.organizationsId,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "logging.organizations.logs.list" call.
// Exactly one of *ListLogsResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListLogsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *OrganizationsLogsListCall) Do(opts ...googleapi.CallOption) (*ListLogsResponse, error) {
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
	ret := &ListLogsResponse{
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
	//   "description": "Lists the logs in the project.\nOnly logs that have entries are listed.",
	//   "flatPath": "v2beta1/organizations/{organizationsId}/logs",
	//   "httpMethod": "GET",
	//   "id": "logging.organizations.logs.list",
	//   "parameterOrder": [
	//     "organizationsId"
	//   ],
	//   "parameters": {
	//     "organizationsId": {
	//       "description": "Part of `parent`. The resource name of the entity whose logs are requested.\nIf both `resource_type` and `resourceIndexPrefix` are empty, then\nall logs with entries in this entity are listed.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Optional. The maximum number of results to return from this request.\nNon-positive values are ignored.  The presence of `nextPageToken` in the\nresponse indicates that more results might be available.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. If present, then retrieve the next batch of results from the\npreceding call to this method.  `pageToken` must be the value of\n`nextPageToken` from the previous response.  The values of other method\nparameters should be identical to those in the previous call.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "resourceIndexPrefix": {
	//       "description": "The purpose of this field is to restrict the listed logs to those\nwith entries of a certain kind.\nIf `resource_type` is the name of a resource type,\nthen this field may contain values for the log resource type's indexes.\nOnly logs that have entries whose indexes include the values are listed.\nThe format for this field is `\"/val1/val2.../valN\"`, where `val1` is a\nvalue for the first index, `val2` for the second index, etc.\nAn empty value (a single slash) for an index matches all values,\nand you can omit values for later indexes entirely.\nOptional. The maximum number of results to return from this request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "resourceType": {
	//       "description": "If not empty, this field must be a resource type such as\n`\"gce_instance`. Only logs associated with that resource type are listed.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/organizations/{organizationsId}/logs",
	//   "response": {
	//     "$ref": "ListLogsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only",
	//     "https://www.googleapis.com/auth/logging.admin",
	//     "https://www.googleapis.com/auth/logging.read"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *OrganizationsLogsListCall) Pages(ctx context.Context, f func(*ListLogsResponse) error) error {
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

// method id "logging.organizations.resourceKeys.list":

type OrganizationsResourceKeysListCall struct {
	s               *Service
	organizationsId string
	urlParams_      gensupport.URLParams
	ifNoneMatch_    string
	ctx_            context.Context
}

// List: Lists the resource keys that have log entries in this project.
func (r *OrganizationsResourceKeysService) List(organizationsId string) *OrganizationsResourceKeysListCall {
	c := &OrganizationsResourceKeysListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.organizationsId = organizationsId
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of results to return from this request.
// Non-positive values are ignored.  The presence of `nextPageToken` in
// the
// response indicates that more results might be available.
func (c *OrganizationsResourceKeysListCall) PageSize(pageSize int64) *OrganizationsResourceKeysListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": If present, then
// retrieve the next batch of results from the
// preceding call to this method.  `pageToken` must be the value
// of
// `nextPageToken` from the previous response.  The values of other
// method
// parameters should be identical to those in the previous call.
func (c *OrganizationsResourceKeysListCall) PageToken(pageToken string) *OrganizationsResourceKeysListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OrganizationsResourceKeysListCall) Fields(s ...googleapi.Field) *OrganizationsResourceKeysListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *OrganizationsResourceKeysListCall) IfNoneMatch(entityTag string) *OrganizationsResourceKeysListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *OrganizationsResourceKeysListCall) Context(ctx context.Context) *OrganizationsResourceKeysListCall {
	c.ctx_ = ctx
	return c
}

func (c *OrganizationsResourceKeysListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/organizations/{organizationsId}/resourceKeys")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"organizationsId": c.organizationsId,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "logging.organizations.resourceKeys.list" call.
// Exactly one of *ListResourceKeysResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *ListResourceKeysResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *OrganizationsResourceKeysListCall) Do(opts ...googleapi.CallOption) (*ListResourceKeysResponse, error) {
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
	ret := &ListResourceKeysResponse{
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
	//   "description": "Lists the resource keys that have log entries in this project.",
	//   "flatPath": "v2beta1/organizations/{organizationsId}/resourceKeys",
	//   "httpMethod": "GET",
	//   "id": "logging.organizations.resourceKeys.list",
	//   "parameterOrder": [
	//     "organizationsId"
	//   ],
	//   "parameters": {
	//     "organizationsId": {
	//       "description": "Part of `parent`. The resource name of the entity whose reource keys are to be listed.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Optional. The maximum number of results to return from this request.\nNon-positive values are ignored.  The presence of `nextPageToken` in the\nresponse indicates that more results might be available.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. If present, then retrieve the next batch of results from the\npreceding call to this method.  `pageToken` must be the value of\n`nextPageToken` from the previous response.  The values of other method\nparameters should be identical to those in the previous call.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/organizations/{organizationsId}/resourceKeys",
	//   "response": {
	//     "$ref": "ListResourceKeysResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only",
	//     "https://www.googleapis.com/auth/logging.admin",
	//     "https://www.googleapis.com/auth/logging.read"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *OrganizationsResourceKeysListCall) Pages(ctx context.Context, f func(*ListResourceKeysResponse) error) error {
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

// method id "logging.organizations.resourceTypes.values.list":

type OrganizationsResourceTypesValuesListCall struct {
	s               *Service
	organizationsId string
	resourceTypesId string
	urlParams_      gensupport.URLParams
	ifNoneMatch_    string
	ctx_            context.Context
}

// List: Lists the current index values for a log resource type.
func (r *OrganizationsResourceTypesValuesService) List(organizationsId string, resourceTypesId string) *OrganizationsResourceTypesValuesListCall {
	c := &OrganizationsResourceTypesValuesListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.organizationsId = organizationsId
	c.resourceTypesId = resourceTypesId
	return c
}

// Depth sets the optional parameter "depth": A non-negative integer
// that limits the number of levels of the index
// hierarchy that are returned.
// If `depth` is 1 (default), only the first index key value is
// returned.
// If `depth` is 2, both primary and secondary key values are
// returned.
// If `depth` is 0, the depth is the number of slash-separators in
// the
// `indexPrefix` field, not counting a slash appearing as the last
// character
// of the prefix.
// If the `indexPrefix` field is empty, the default depth is 1.
// It is an error for `depth` to be any positive value
// less than the number of components in `indexPrefix`.
func (c *OrganizationsResourceTypesValuesListCall) Depth(depth int64) *OrganizationsResourceTypesValuesListCall {
	c.urlParams_.Set("depth", fmt.Sprint(depth))
	return c
}

// IndexPrefix sets the optional parameter "indexPrefix": Restricts the
// index values returned to be those with a specified prefix
// for each index key.
// This field has the form "/prefix1/prefix2/...", in order
// corresponding
// to the `ResourceKeys indexKeys`.
// Non-empty prefixes must begin with `/`.
// For example, App Engine's two keys are the module ID and the version
// ID.
// Following is the effect of using various values for `indexPrefix`:
//
// +  "/Mod/" retrieves `/Mod/10` and `/Mod/11` but not `/ModA/10`.
// +  "/Mod` retrieves `/Mod/10`, `/Mod/11` and `/ModA/10` but not
// `/XXX/33`.
// +  "/Mod/1" retrieves `/Mod/10` and `/Mod/11` but not `/ModA/10`.
// +  "/Mod/10/" retrieves `/Mod/10` only.
// +  An empty prefix or "/" retrieves all values.
func (c *OrganizationsResourceTypesValuesListCall) IndexPrefix(indexPrefix string) *OrganizationsResourceTypesValuesListCall {
	c.urlParams_.Set("indexPrefix", indexPrefix)
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of results to return from this request.
// Non-positive values are ignored.  The presence of `nextPageToken` in
// the
// response indicates that more results might be available.
func (c *OrganizationsResourceTypesValuesListCall) PageSize(pageSize int64) *OrganizationsResourceTypesValuesListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": If present, then
// retrieve the next batch of results from the
// preceding call to this method.  `pageToken` must be the value
// of
// `nextPageToken` from the previous response.  The values of other
// method
// parameters should be identical to those in the previous call.
func (c *OrganizationsResourceTypesValuesListCall) PageToken(pageToken string) *OrganizationsResourceTypesValuesListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OrganizationsResourceTypesValuesListCall) Fields(s ...googleapi.Field) *OrganizationsResourceTypesValuesListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *OrganizationsResourceTypesValuesListCall) IfNoneMatch(entityTag string) *OrganizationsResourceTypesValuesListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *OrganizationsResourceTypesValuesListCall) Context(ctx context.Context) *OrganizationsResourceTypesValuesListCall {
	c.ctx_ = ctx
	return c
}

func (c *OrganizationsResourceTypesValuesListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/organizations/{organizationsId}/resourceTypes/{resourceTypesId}/values")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"organizationsId": c.organizationsId,
		"resourceTypesId": c.resourceTypesId,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "logging.organizations.resourceTypes.values.list" call.
// Exactly one of *ListResourceValuesResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *ListResourceValuesResponse.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *OrganizationsResourceTypesValuesListCall) Do(opts ...googleapi.CallOption) (*ListResourceValuesResponse, error) {
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
	ret := &ListResourceValuesResponse{
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
	//   "description": "Lists the current index values for a log resource type.",
	//   "flatPath": "v2beta1/organizations/{organizationsId}/resourceTypes/{resourceTypesId}/values",
	//   "httpMethod": "GET",
	//   "id": "logging.organizations.resourceTypes.values.list",
	//   "parameterOrder": [
	//     "organizationsId",
	//     "resourceTypesId"
	//   ],
	//   "parameters": {
	//     "depth": {
	//       "description": "A non-negative integer that limits the number of levels of the index\nhierarchy that are returned.\nIf `depth` is 1 (default), only the first index key value is returned.\nIf `depth` is 2, both primary and secondary key values are returned.\nIf `depth` is 0, the depth is the number of slash-separators in the\n`indexPrefix` field, not counting a slash appearing as the last character\nof the prefix.\nIf the `indexPrefix` field is empty, the default depth is 1.\nIt is an error for `depth` to be any positive value\nless than the number of components in `indexPrefix`.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "indexPrefix": {
	//       "description": "Restricts the index values returned to be those with a specified prefix\nfor each index key.\nThis field has the form `\"/prefix1/prefix2/...\"`, in order corresponding\nto the `ResourceKeys indexKeys`.\nNon-empty prefixes must begin with `/`.\nFor example, App Engine's two keys are the module ID and the version ID.\nFollowing is the effect of using various values for `indexPrefix`:\n\n+  `\"/Mod/\"` retrieves `/Mod/10` and `/Mod/11` but not `/ModA/10`.\n+  `\"/Mod` retrieves `/Mod/10`, `/Mod/11` and `/ModA/10` but not `/XXX/33`.\n+  `\"/Mod/1\"` retrieves `/Mod/10` and `/Mod/11` but not `/ModA/10`.\n+  `\"/Mod/10/\"` retrieves `/Mod/10` only.\n+  An empty prefix or `\"/\"` retrieves all values.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "organizationsId": {
	//       "description": "Part of `parent`. The resource name of a resource type whose indexes are requested.\nExample:\n`\"projects/my-project-id/resourceTypes/gae_app\"`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Optional. The maximum number of results to return from this request.\nNon-positive values are ignored.  The presence of `nextPageToken` in the\nresponse indicates that more results might be available.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. If present, then retrieve the next batch of results from the\npreceding call to this method.  `pageToken` must be the value of\n`nextPageToken` from the previous response.  The values of other method\nparameters should be identical to those in the previous call.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "resourceTypesId": {
	//       "description": "Part of `parent`. See documentation of `organizationsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/organizations/{organizationsId}/resourceTypes/{resourceTypesId}/values",
	//   "response": {
	//     "$ref": "ListResourceValuesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only",
	//     "https://www.googleapis.com/auth/logging.admin",
	//     "https://www.googleapis.com/auth/logging.read"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *OrganizationsResourceTypesValuesListCall) Pages(ctx context.Context, f func(*ListResourceValuesResponse) error) error {
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

// method id "logging.projects.logs.delete":

type ProjectsLogsDeleteCall struct {
	s          *Service
	projectsId string
	logsId     string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
}

// Delete: Deletes a log and all its log entries.
// The log will reappear if it receives new entries.
func (r *ProjectsLogsService) Delete(projectsId string, logsId string) *ProjectsLogsDeleteCall {
	c := &ProjectsLogsDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectsId = projectsId
	c.logsId = logsId
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogsDeleteCall) Fields(s ...googleapi.Field) *ProjectsLogsDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsLogsDeleteCall) Context(ctx context.Context) *ProjectsLogsDeleteCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsLogsDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/projects/{projectsId}/logs/{logsId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectsId": c.projectsId,
		"logsId":     c.logsId,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "logging.projects.logs.delete" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsLogsDeleteCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
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
	//   "description": "Deletes a log and all its log entries.\nThe log will reappear if it receives new entries.",
	//   "flatPath": "v2beta1/projects/{projectsId}/logs/{logsId}",
	//   "httpMethod": "DELETE",
	//   "id": "logging.projects.logs.delete",
	//   "parameterOrder": [
	//     "projectsId",
	//     "logsId"
	//   ],
	//   "parameters": {
	//     "logsId": {
	//       "description": "Part of `logName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `logName`. Required. The resource name of the log to delete.  Example:\n`\"projects/my-project/logs/syslog\"`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/projects/{projectsId}/logs/{logsId}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/logging.admin"
	//   ]
	// }

}

// method id "logging.projects.logs.list":

type ProjectsLogsListCall struct {
	s            *Service
	projectsId   string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// List: Lists the logs in the project.
// Only logs that have entries are listed.
func (r *ProjectsLogsService) List(projectsId string) *ProjectsLogsListCall {
	c := &ProjectsLogsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectsId = projectsId
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of results to return from this request.
// Non-positive values are ignored.  The presence of `nextPageToken` in
// the
// response indicates that more results might be available.
func (c *ProjectsLogsListCall) PageSize(pageSize int64) *ProjectsLogsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": If present, then
// retrieve the next batch of results from the
// preceding call to this method.  `pageToken` must be the value
// of
// `nextPageToken` from the previous response.  The values of other
// method
// parameters should be identical to those in the previous call.
func (c *ProjectsLogsListCall) PageToken(pageToken string) *ProjectsLogsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// ResourceIndexPrefix sets the optional parameter
// "resourceIndexPrefix": The purpose of this field is to restrict the
// listed logs to those
// with entries of a certain kind.
// If `resource_type` is the name of a resource type,
// then this field may contain values for the log resource type's
// indexes.
// Only logs that have entries whose indexes include the values are
// listed.
// The format for this field is "/val1/val2.../valN", where `val1` is
// a
// value for the first index, `val2` for the second index, etc.
// An empty value (a single slash) for an index matches all values,
// and you can omit values for later indexes entirely.
//  The maximum number of results to return from this request.
func (c *ProjectsLogsListCall) ResourceIndexPrefix(resourceIndexPrefix string) *ProjectsLogsListCall {
	c.urlParams_.Set("resourceIndexPrefix", resourceIndexPrefix)
	return c
}

// ResourceType sets the optional parameter "resourceType": If not
// empty, this field must be a resource type such as
// "gce_instance`. Only logs associated with that resource type are
// listed.
func (c *ProjectsLogsListCall) ResourceType(resourceType string) *ProjectsLogsListCall {
	c.urlParams_.Set("resourceType", resourceType)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogsListCall) Fields(s ...googleapi.Field) *ProjectsLogsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsLogsListCall) IfNoneMatch(entityTag string) *ProjectsLogsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsLogsListCall) Context(ctx context.Context) *ProjectsLogsListCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsLogsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/projects/{projectsId}/logs")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectsId": c.projectsId,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "logging.projects.logs.list" call.
// Exactly one of *ListLogsResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListLogsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsLogsListCall) Do(opts ...googleapi.CallOption) (*ListLogsResponse, error) {
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
	ret := &ListLogsResponse{
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
	//   "description": "Lists the logs in the project.\nOnly logs that have entries are listed.",
	//   "flatPath": "v2beta1/projects/{projectsId}/logs",
	//   "httpMethod": "GET",
	//   "id": "logging.projects.logs.list",
	//   "parameterOrder": [
	//     "projectsId"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "Optional. The maximum number of results to return from this request.\nNon-positive values are ignored.  The presence of `nextPageToken` in the\nresponse indicates that more results might be available.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. If present, then retrieve the next batch of results from the\npreceding call to this method.  `pageToken` must be the value of\n`nextPageToken` from the previous response.  The values of other method\nparameters should be identical to those in the previous call.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `parent`. The resource name of the entity whose logs are requested.\nIf both `resource_type` and `resourceIndexPrefix` are empty, then\nall logs with entries in this entity are listed.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "resourceIndexPrefix": {
	//       "description": "The purpose of this field is to restrict the listed logs to those\nwith entries of a certain kind.\nIf `resource_type` is the name of a resource type,\nthen this field may contain values for the log resource type's indexes.\nOnly logs that have entries whose indexes include the values are listed.\nThe format for this field is `\"/val1/val2.../valN\"`, where `val1` is a\nvalue for the first index, `val2` for the second index, etc.\nAn empty value (a single slash) for an index matches all values,\nand you can omit values for later indexes entirely.\nOptional. The maximum number of results to return from this request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "resourceType": {
	//       "description": "If not empty, this field must be a resource type such as\n`\"gce_instance`. Only logs associated with that resource type are listed.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/projects/{projectsId}/logs",
	//   "response": {
	//     "$ref": "ListLogsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only",
	//     "https://www.googleapis.com/auth/logging.admin",
	//     "https://www.googleapis.com/auth/logging.read"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsLogsListCall) Pages(ctx context.Context, f func(*ListLogsResponse) error) error {
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

// method id "logging.projects.metrics.create":

type ProjectsMetricsCreateCall struct {
	s          *Service
	projectsId string
	logmetric  *LogMetric
	urlParams_ gensupport.URLParams
	ctx_       context.Context
}

// Create: Creates a logs-based metric.
func (r *ProjectsMetricsService) Create(projectsId string, logmetric *LogMetric) *ProjectsMetricsCreateCall {
	c := &ProjectsMetricsCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectsId = projectsId
	c.logmetric = logmetric
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsMetricsCreateCall) Fields(s ...googleapi.Field) *ProjectsMetricsCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsMetricsCreateCall) Context(ctx context.Context) *ProjectsMetricsCreateCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsMetricsCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.logmetric)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/projects/{projectsId}/metrics")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectsId": c.projectsId,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "logging.projects.metrics.create" call.
// Exactly one of *LogMetric or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *LogMetric.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsMetricsCreateCall) Do(opts ...googleapi.CallOption) (*LogMetric, error) {
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
	ret := &LogMetric{
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
	//   "description": "Creates a logs-based metric.",
	//   "flatPath": "v2beta1/projects/{projectsId}/metrics",
	//   "httpMethod": "POST",
	//   "id": "logging.projects.metrics.create",
	//   "parameterOrder": [
	//     "projectsId"
	//   ],
	//   "parameters": {
	//     "projectsId": {
	//       "description": "Part of `parent`. The resource name of the project in which to create the metric.\nExample: `\"projects/my-project-id\"`.\n\nThe new metric must be provided in the request.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/projects/{projectsId}/metrics",
	//   "request": {
	//     "$ref": "LogMetric"
	//   },
	//   "response": {
	//     "$ref": "LogMetric"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/logging.admin",
	//     "https://www.googleapis.com/auth/logging.write"
	//   ]
	// }

}

// method id "logging.projects.metrics.delete":

type ProjectsMetricsDeleteCall struct {
	s          *Service
	projectsId string
	metricsId  string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
}

// Delete: Deletes a logs-based metric.
func (r *ProjectsMetricsService) Delete(projectsId string, metricsId string) *ProjectsMetricsDeleteCall {
	c := &ProjectsMetricsDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectsId = projectsId
	c.metricsId = metricsId
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsMetricsDeleteCall) Fields(s ...googleapi.Field) *ProjectsMetricsDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsMetricsDeleteCall) Context(ctx context.Context) *ProjectsMetricsDeleteCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsMetricsDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/projects/{projectsId}/metrics/{metricsId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectsId": c.projectsId,
		"metricsId":  c.metricsId,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "logging.projects.metrics.delete" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsMetricsDeleteCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
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
	//   "description": "Deletes a logs-based metric.",
	//   "flatPath": "v2beta1/projects/{projectsId}/metrics/{metricsId}",
	//   "httpMethod": "DELETE",
	//   "id": "logging.projects.metrics.delete",
	//   "parameterOrder": [
	//     "projectsId",
	//     "metricsId"
	//   ],
	//   "parameters": {
	//     "metricsId": {
	//       "description": "Part of `metricName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `metricName`. The resource name of the metric to delete.\nExample: `\"projects/my-project-id/metrics/my-metric-id\"`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/projects/{projectsId}/metrics/{metricsId}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/logging.admin",
	//     "https://www.googleapis.com/auth/logging.write"
	//   ]
	// }

}

// method id "logging.projects.metrics.get":

type ProjectsMetricsGetCall struct {
	s            *Service
	projectsId   string
	metricsId    string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// Get: Gets a logs-based metric.
func (r *ProjectsMetricsService) Get(projectsId string, metricsId string) *ProjectsMetricsGetCall {
	c := &ProjectsMetricsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectsId = projectsId
	c.metricsId = metricsId
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsMetricsGetCall) Fields(s ...googleapi.Field) *ProjectsMetricsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsMetricsGetCall) IfNoneMatch(entityTag string) *ProjectsMetricsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsMetricsGetCall) Context(ctx context.Context) *ProjectsMetricsGetCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsMetricsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/projects/{projectsId}/metrics/{metricsId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectsId": c.projectsId,
		"metricsId":  c.metricsId,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "logging.projects.metrics.get" call.
// Exactly one of *LogMetric or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *LogMetric.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsMetricsGetCall) Do(opts ...googleapi.CallOption) (*LogMetric, error) {
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
	ret := &LogMetric{
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
	//   "description": "Gets a logs-based metric.",
	//   "flatPath": "v2beta1/projects/{projectsId}/metrics/{metricsId}",
	//   "httpMethod": "GET",
	//   "id": "logging.projects.metrics.get",
	//   "parameterOrder": [
	//     "projectsId",
	//     "metricsId"
	//   ],
	//   "parameters": {
	//     "metricsId": {
	//       "description": "Part of `metricName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `metricName`. The resource name of the desired metric.\nExample: `\"projects/my-project-id/metrics/my-metric-id\"`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/projects/{projectsId}/metrics/{metricsId}",
	//   "response": {
	//     "$ref": "LogMetric"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only",
	//     "https://www.googleapis.com/auth/logging.admin",
	//     "https://www.googleapis.com/auth/logging.read"
	//   ]
	// }

}

// method id "logging.projects.metrics.list":

type ProjectsMetricsListCall struct {
	s            *Service
	projectsId   string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// List: Lists logs-based metrics.
func (r *ProjectsMetricsService) List(projectsId string) *ProjectsMetricsListCall {
	c := &ProjectsMetricsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectsId = projectsId
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of results to return from this request.
// Non-positive values are ignored.  The presence of `nextPageToken` in
// the
// response indicates that more results might be available.
func (c *ProjectsMetricsListCall) PageSize(pageSize int64) *ProjectsMetricsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": If present, then
// retrieve the next batch of results from the
// preceding call to this method.  `pageToken` must be the value
// of
// `nextPageToken` from the previous response.  The values of other
// method
// parameters should be identical to those in the previous call.
func (c *ProjectsMetricsListCall) PageToken(pageToken string) *ProjectsMetricsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsMetricsListCall) Fields(s ...googleapi.Field) *ProjectsMetricsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsMetricsListCall) IfNoneMatch(entityTag string) *ProjectsMetricsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsMetricsListCall) Context(ctx context.Context) *ProjectsMetricsListCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsMetricsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/projects/{projectsId}/metrics")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectsId": c.projectsId,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "logging.projects.metrics.list" call.
// Exactly one of *ListLogMetricsResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListLogMetricsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsMetricsListCall) Do(opts ...googleapi.CallOption) (*ListLogMetricsResponse, error) {
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
	ret := &ListLogMetricsResponse{
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
	//   "description": "Lists logs-based metrics.",
	//   "flatPath": "v2beta1/projects/{projectsId}/metrics",
	//   "httpMethod": "GET",
	//   "id": "logging.projects.metrics.list",
	//   "parameterOrder": [
	//     "projectsId"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "Optional. The maximum number of results to return from this request.\nNon-positive values are ignored.  The presence of `nextPageToken` in the\nresponse indicates that more results might be available.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. If present, then retrieve the next batch of results from the\npreceding call to this method.  `pageToken` must be the value of\n`nextPageToken` from the previous response.  The values of other method\nparameters should be identical to those in the previous call.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `parent`. Required. The resource name containing the metrics.\nExample: `\"projects/my-project-id\"`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/projects/{projectsId}/metrics",
	//   "response": {
	//     "$ref": "ListLogMetricsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only",
	//     "https://www.googleapis.com/auth/logging.admin",
	//     "https://www.googleapis.com/auth/logging.read"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsMetricsListCall) Pages(ctx context.Context, f func(*ListLogMetricsResponse) error) error {
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

// method id "logging.projects.metrics.update":

type ProjectsMetricsUpdateCall struct {
	s          *Service
	projectsId string
	metricsId  string
	logmetric  *LogMetric
	urlParams_ gensupport.URLParams
	ctx_       context.Context
}

// Update: Creates or updates a logs-based metric.
func (r *ProjectsMetricsService) Update(projectsId string, metricsId string, logmetric *LogMetric) *ProjectsMetricsUpdateCall {
	c := &ProjectsMetricsUpdateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectsId = projectsId
	c.metricsId = metricsId
	c.logmetric = logmetric
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsMetricsUpdateCall) Fields(s ...googleapi.Field) *ProjectsMetricsUpdateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsMetricsUpdateCall) Context(ctx context.Context) *ProjectsMetricsUpdateCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsMetricsUpdateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.logmetric)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/projects/{projectsId}/metrics/{metricsId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectsId": c.projectsId,
		"metricsId":  c.metricsId,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "logging.projects.metrics.update" call.
// Exactly one of *LogMetric or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *LogMetric.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsMetricsUpdateCall) Do(opts ...googleapi.CallOption) (*LogMetric, error) {
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
	ret := &LogMetric{
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
	//   "description": "Creates or updates a logs-based metric.",
	//   "flatPath": "v2beta1/projects/{projectsId}/metrics/{metricsId}",
	//   "httpMethod": "PUT",
	//   "id": "logging.projects.metrics.update",
	//   "parameterOrder": [
	//     "projectsId",
	//     "metricsId"
	//   ],
	//   "parameters": {
	//     "metricsId": {
	//       "description": "Part of `metricName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `metricName`. The resource name of the metric to update.\nExample: `\"projects/my-project-id/metrics/my-metric-id\"`.\n\nThe updated metric must be provided in the request and have the\nsame identifier that is specified in `metricName`.\nIf the metric does not exist, it is created.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/projects/{projectsId}/metrics/{metricsId}",
	//   "request": {
	//     "$ref": "LogMetric"
	//   },
	//   "response": {
	//     "$ref": "LogMetric"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/logging.admin",
	//     "https://www.googleapis.com/auth/logging.write"
	//   ]
	// }

}

// method id "logging.projects.resourceKeys.list":

type ProjectsResourceKeysListCall struct {
	s            *Service
	projectsId   string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// List: Lists the resource keys that have log entries in this project.
func (r *ProjectsResourceKeysService) List(projectsId string) *ProjectsResourceKeysListCall {
	c := &ProjectsResourceKeysListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectsId = projectsId
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of results to return from this request.
// Non-positive values are ignored.  The presence of `nextPageToken` in
// the
// response indicates that more results might be available.
func (c *ProjectsResourceKeysListCall) PageSize(pageSize int64) *ProjectsResourceKeysListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": If present, then
// retrieve the next batch of results from the
// preceding call to this method.  `pageToken` must be the value
// of
// `nextPageToken` from the previous response.  The values of other
// method
// parameters should be identical to those in the previous call.
func (c *ProjectsResourceKeysListCall) PageToken(pageToken string) *ProjectsResourceKeysListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsResourceKeysListCall) Fields(s ...googleapi.Field) *ProjectsResourceKeysListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsResourceKeysListCall) IfNoneMatch(entityTag string) *ProjectsResourceKeysListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsResourceKeysListCall) Context(ctx context.Context) *ProjectsResourceKeysListCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsResourceKeysListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/projects/{projectsId}/resourceKeys")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectsId": c.projectsId,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "logging.projects.resourceKeys.list" call.
// Exactly one of *ListResourceKeysResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *ListResourceKeysResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsResourceKeysListCall) Do(opts ...googleapi.CallOption) (*ListResourceKeysResponse, error) {
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
	ret := &ListResourceKeysResponse{
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
	//   "description": "Lists the resource keys that have log entries in this project.",
	//   "flatPath": "v2beta1/projects/{projectsId}/resourceKeys",
	//   "httpMethod": "GET",
	//   "id": "logging.projects.resourceKeys.list",
	//   "parameterOrder": [
	//     "projectsId"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "Optional. The maximum number of results to return from this request.\nNon-positive values are ignored.  The presence of `nextPageToken` in the\nresponse indicates that more results might be available.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. If present, then retrieve the next batch of results from the\npreceding call to this method.  `pageToken` must be the value of\n`nextPageToken` from the previous response.  The values of other method\nparameters should be identical to those in the previous call.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `parent`. The resource name of the entity whose reource keys are to be listed.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/projects/{projectsId}/resourceKeys",
	//   "response": {
	//     "$ref": "ListResourceKeysResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only",
	//     "https://www.googleapis.com/auth/logging.admin",
	//     "https://www.googleapis.com/auth/logging.read"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsResourceKeysListCall) Pages(ctx context.Context, f func(*ListResourceKeysResponse) error) error {
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

// method id "logging.projects.resourceTypes.values.list":

type ProjectsResourceTypesValuesListCall struct {
	s               *Service
	projectsId      string
	resourceTypesId string
	urlParams_      gensupport.URLParams
	ifNoneMatch_    string
	ctx_            context.Context
}

// List: Lists the current index values for a log resource type.
func (r *ProjectsResourceTypesValuesService) List(projectsId string, resourceTypesId string) *ProjectsResourceTypesValuesListCall {
	c := &ProjectsResourceTypesValuesListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectsId = projectsId
	c.resourceTypesId = resourceTypesId
	return c
}

// Depth sets the optional parameter "depth": A non-negative integer
// that limits the number of levels of the index
// hierarchy that are returned.
// If `depth` is 1 (default), only the first index key value is
// returned.
// If `depth` is 2, both primary and secondary key values are
// returned.
// If `depth` is 0, the depth is the number of slash-separators in
// the
// `indexPrefix` field, not counting a slash appearing as the last
// character
// of the prefix.
// If the `indexPrefix` field is empty, the default depth is 1.
// It is an error for `depth` to be any positive value
// less than the number of components in `indexPrefix`.
func (c *ProjectsResourceTypesValuesListCall) Depth(depth int64) *ProjectsResourceTypesValuesListCall {
	c.urlParams_.Set("depth", fmt.Sprint(depth))
	return c
}

// IndexPrefix sets the optional parameter "indexPrefix": Restricts the
// index values returned to be those with a specified prefix
// for each index key.
// This field has the form "/prefix1/prefix2/...", in order
// corresponding
// to the `ResourceKeys indexKeys`.
// Non-empty prefixes must begin with `/`.
// For example, App Engine's two keys are the module ID and the version
// ID.
// Following is the effect of using various values for `indexPrefix`:
//
// +  "/Mod/" retrieves `/Mod/10` and `/Mod/11` but not `/ModA/10`.
// +  "/Mod` retrieves `/Mod/10`, `/Mod/11` and `/ModA/10` but not
// `/XXX/33`.
// +  "/Mod/1" retrieves `/Mod/10` and `/Mod/11` but not `/ModA/10`.
// +  "/Mod/10/" retrieves `/Mod/10` only.
// +  An empty prefix or "/" retrieves all values.
func (c *ProjectsResourceTypesValuesListCall) IndexPrefix(indexPrefix string) *ProjectsResourceTypesValuesListCall {
	c.urlParams_.Set("indexPrefix", indexPrefix)
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of results to return from this request.
// Non-positive values are ignored.  The presence of `nextPageToken` in
// the
// response indicates that more results might be available.
func (c *ProjectsResourceTypesValuesListCall) PageSize(pageSize int64) *ProjectsResourceTypesValuesListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": If present, then
// retrieve the next batch of results from the
// preceding call to this method.  `pageToken` must be the value
// of
// `nextPageToken` from the previous response.  The values of other
// method
// parameters should be identical to those in the previous call.
func (c *ProjectsResourceTypesValuesListCall) PageToken(pageToken string) *ProjectsResourceTypesValuesListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsResourceTypesValuesListCall) Fields(s ...googleapi.Field) *ProjectsResourceTypesValuesListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsResourceTypesValuesListCall) IfNoneMatch(entityTag string) *ProjectsResourceTypesValuesListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsResourceTypesValuesListCall) Context(ctx context.Context) *ProjectsResourceTypesValuesListCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsResourceTypesValuesListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/projects/{projectsId}/resourceTypes/{resourceTypesId}/values")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectsId":      c.projectsId,
		"resourceTypesId": c.resourceTypesId,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "logging.projects.resourceTypes.values.list" call.
// Exactly one of *ListResourceValuesResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *ListResourceValuesResponse.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsResourceTypesValuesListCall) Do(opts ...googleapi.CallOption) (*ListResourceValuesResponse, error) {
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
	ret := &ListResourceValuesResponse{
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
	//   "description": "Lists the current index values for a log resource type.",
	//   "flatPath": "v2beta1/projects/{projectsId}/resourceTypes/{resourceTypesId}/values",
	//   "httpMethod": "GET",
	//   "id": "logging.projects.resourceTypes.values.list",
	//   "parameterOrder": [
	//     "projectsId",
	//     "resourceTypesId"
	//   ],
	//   "parameters": {
	//     "depth": {
	//       "description": "A non-negative integer that limits the number of levels of the index\nhierarchy that are returned.\nIf `depth` is 1 (default), only the first index key value is returned.\nIf `depth` is 2, both primary and secondary key values are returned.\nIf `depth` is 0, the depth is the number of slash-separators in the\n`indexPrefix` field, not counting a slash appearing as the last character\nof the prefix.\nIf the `indexPrefix` field is empty, the default depth is 1.\nIt is an error for `depth` to be any positive value\nless than the number of components in `indexPrefix`.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "indexPrefix": {
	//       "description": "Restricts the index values returned to be those with a specified prefix\nfor each index key.\nThis field has the form `\"/prefix1/prefix2/...\"`, in order corresponding\nto the `ResourceKeys indexKeys`.\nNon-empty prefixes must begin with `/`.\nFor example, App Engine's two keys are the module ID and the version ID.\nFollowing is the effect of using various values for `indexPrefix`:\n\n+  `\"/Mod/\"` retrieves `/Mod/10` and `/Mod/11` but not `/ModA/10`.\n+  `\"/Mod` retrieves `/Mod/10`, `/Mod/11` and `/ModA/10` but not `/XXX/33`.\n+  `\"/Mod/1\"` retrieves `/Mod/10` and `/Mod/11` but not `/ModA/10`.\n+  `\"/Mod/10/\"` retrieves `/Mod/10` only.\n+  An empty prefix or `\"/\"` retrieves all values.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Optional. The maximum number of results to return from this request.\nNon-positive values are ignored.  The presence of `nextPageToken` in the\nresponse indicates that more results might be available.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. If present, then retrieve the next batch of results from the\npreceding call to this method.  `pageToken` must be the value of\n`nextPageToken` from the previous response.  The values of other method\nparameters should be identical to those in the previous call.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `parent`. The resource name of a resource type whose indexes are requested.\nExample:\n`\"projects/my-project-id/resourceTypes/gae_app\"`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "resourceTypesId": {
	//       "description": "Part of `parent`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/projects/{projectsId}/resourceTypes/{resourceTypesId}/values",
	//   "response": {
	//     "$ref": "ListResourceValuesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only",
	//     "https://www.googleapis.com/auth/logging.admin",
	//     "https://www.googleapis.com/auth/logging.read"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsResourceTypesValuesListCall) Pages(ctx context.Context, f func(*ListResourceValuesResponse) error) error {
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

// method id "logging.projects.sinks.create":

type ProjectsSinksCreateCall struct {
	s          *Service
	projectsId string
	logsink    *LogSink
	urlParams_ gensupport.URLParams
	ctx_       context.Context
}

// Create: Creates a sink.
func (r *ProjectsSinksService) Create(projectsId string, logsink *LogSink) *ProjectsSinksCreateCall {
	c := &ProjectsSinksCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectsId = projectsId
	c.logsink = logsink
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsSinksCreateCall) Fields(s ...googleapi.Field) *ProjectsSinksCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsSinksCreateCall) Context(ctx context.Context) *ProjectsSinksCreateCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsSinksCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.logsink)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/projects/{projectsId}/sinks")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectsId": c.projectsId,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "logging.projects.sinks.create" call.
// Exactly one of *LogSink or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *LogSink.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsSinksCreateCall) Do(opts ...googleapi.CallOption) (*LogSink, error) {
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
	ret := &LogSink{
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
	//   "description": "Creates a sink.",
	//   "flatPath": "v2beta1/projects/{projectsId}/sinks",
	//   "httpMethod": "POST",
	//   "id": "logging.projects.sinks.create",
	//   "parameterOrder": [
	//     "projectsId"
	//   ],
	//   "parameters": {
	//     "projectsId": {
	//       "description": "Part of `parent`. Required. The resource in which to create the sink.\nExample: `\"projects/my-project-id\"`.\nThe new sink must be provided in the request.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/projects/{projectsId}/sinks",
	//   "request": {
	//     "$ref": "LogSink"
	//   },
	//   "response": {
	//     "$ref": "LogSink"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/logging.admin"
	//   ]
	// }

}

// method id "logging.projects.sinks.delete":

type ProjectsSinksDeleteCall struct {
	s          *Service
	projectsId string
	sinksId    string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
}

// Delete: Deletes a sink.
func (r *ProjectsSinksService) Delete(projectsId string, sinksId string) *ProjectsSinksDeleteCall {
	c := &ProjectsSinksDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectsId = projectsId
	c.sinksId = sinksId
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsSinksDeleteCall) Fields(s ...googleapi.Field) *ProjectsSinksDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsSinksDeleteCall) Context(ctx context.Context) *ProjectsSinksDeleteCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsSinksDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/projects/{projectsId}/sinks/{sinksId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectsId": c.projectsId,
		"sinksId":    c.sinksId,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "logging.projects.sinks.delete" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsSinksDeleteCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
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
	//   "description": "Deletes a sink.",
	//   "flatPath": "v2beta1/projects/{projectsId}/sinks/{sinksId}",
	//   "httpMethod": "DELETE",
	//   "id": "logging.projects.sinks.delete",
	//   "parameterOrder": [
	//     "projectsId",
	//     "sinksId"
	//   ],
	//   "parameters": {
	//     "projectsId": {
	//       "description": "Part of `sinkName`. Required. The resource name of the sink to delete, including the parent\nresource and the sink identifier.  Example:\n`\"projects/my-project-id/sinks/my-sink-id\"`.  It is an error if the sink\ndoes not exist.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "sinksId": {
	//       "description": "Part of `sinkName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/projects/{projectsId}/sinks/{sinksId}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/logging.admin"
	//   ]
	// }

}

// method id "logging.projects.sinks.get":

type ProjectsSinksGetCall struct {
	s            *Service
	projectsId   string
	sinksId      string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// Get: Gets a sink.
func (r *ProjectsSinksService) Get(projectsId string, sinksId string) *ProjectsSinksGetCall {
	c := &ProjectsSinksGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectsId = projectsId
	c.sinksId = sinksId
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsSinksGetCall) Fields(s ...googleapi.Field) *ProjectsSinksGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsSinksGetCall) IfNoneMatch(entityTag string) *ProjectsSinksGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsSinksGetCall) Context(ctx context.Context) *ProjectsSinksGetCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsSinksGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/projects/{projectsId}/sinks/{sinksId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectsId": c.projectsId,
		"sinksId":    c.sinksId,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "logging.projects.sinks.get" call.
// Exactly one of *LogSink or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *LogSink.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsSinksGetCall) Do(opts ...googleapi.CallOption) (*LogSink, error) {
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
	ret := &LogSink{
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
	//   "description": "Gets a sink.",
	//   "flatPath": "v2beta1/projects/{projectsId}/sinks/{sinksId}",
	//   "httpMethod": "GET",
	//   "id": "logging.projects.sinks.get",
	//   "parameterOrder": [
	//     "projectsId",
	//     "sinksId"
	//   ],
	//   "parameters": {
	//     "projectsId": {
	//       "description": "Part of `sinkName`. Required. The resource name of the sink to return.\nExample: `\"projects/my-project-id/sinks/my-sink-id\"`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "sinksId": {
	//       "description": "Part of `sinkName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/projects/{projectsId}/sinks/{sinksId}",
	//   "response": {
	//     "$ref": "LogSink"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only",
	//     "https://www.googleapis.com/auth/logging.admin",
	//     "https://www.googleapis.com/auth/logging.read"
	//   ]
	// }

}

// method id "logging.projects.sinks.list":

type ProjectsSinksListCall struct {
	s            *Service
	projectsId   string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// List: Lists sinks.
func (r *ProjectsSinksService) List(projectsId string) *ProjectsSinksListCall {
	c := &ProjectsSinksListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectsId = projectsId
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of results to return from this request.
// Non-positive values are ignored.  The presence of `nextPageToken` in
// the
// response indicates that more results might be available.
func (c *ProjectsSinksListCall) PageSize(pageSize int64) *ProjectsSinksListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": If present, then
// retrieve the next batch of results from the
// preceding call to this method.  `pageToken` must be the value
// of
// `nextPageToken` from the previous response.  The values of other
// method
// parameters should be identical to those in the previous call.
func (c *ProjectsSinksListCall) PageToken(pageToken string) *ProjectsSinksListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsSinksListCall) Fields(s ...googleapi.Field) *ProjectsSinksListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsSinksListCall) IfNoneMatch(entityTag string) *ProjectsSinksListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsSinksListCall) Context(ctx context.Context) *ProjectsSinksListCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsSinksListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/projects/{projectsId}/sinks")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectsId": c.projectsId,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "logging.projects.sinks.list" call.
// Exactly one of *ListSinksResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListSinksResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsSinksListCall) Do(opts ...googleapi.CallOption) (*ListSinksResponse, error) {
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
	ret := &ListSinksResponse{
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
	//   "description": "Lists sinks.",
	//   "flatPath": "v2beta1/projects/{projectsId}/sinks",
	//   "httpMethod": "GET",
	//   "id": "logging.projects.sinks.list",
	//   "parameterOrder": [
	//     "projectsId"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "Optional. The maximum number of results to return from this request.\nNon-positive values are ignored.  The presence of `nextPageToken` in the\nresponse indicates that more results might be available.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. If present, then retrieve the next batch of results from the\npreceding call to this method.  `pageToken` must be the value of\n`nextPageToken` from the previous response.  The values of other method\nparameters should be identical to those in the previous call.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `parent`. Required. The cloud resource containing the sinks.\nExample: `\"projects/my-logging-project\"`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/projects/{projectsId}/sinks",
	//   "response": {
	//     "$ref": "ListSinksResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only",
	//     "https://www.googleapis.com/auth/logging.admin",
	//     "https://www.googleapis.com/auth/logging.read"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsSinksListCall) Pages(ctx context.Context, f func(*ListSinksResponse) error) error {
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

// method id "logging.projects.sinks.update":

type ProjectsSinksUpdateCall struct {
	s          *Service
	projectsId string
	sinksId    string
	logsink    *LogSink
	urlParams_ gensupport.URLParams
	ctx_       context.Context
}

// Update: Updates or creates a sink.
func (r *ProjectsSinksService) Update(projectsId string, sinksId string, logsink *LogSink) *ProjectsSinksUpdateCall {
	c := &ProjectsSinksUpdateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectsId = projectsId
	c.sinksId = sinksId
	c.logsink = logsink
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsSinksUpdateCall) Fields(s ...googleapi.Field) *ProjectsSinksUpdateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsSinksUpdateCall) Context(ctx context.Context) *ProjectsSinksUpdateCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsSinksUpdateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.logsink)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/projects/{projectsId}/sinks/{sinksId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectsId": c.projectsId,
		"sinksId":    c.sinksId,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "logging.projects.sinks.update" call.
// Exactly one of *LogSink or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *LogSink.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsSinksUpdateCall) Do(opts ...googleapi.CallOption) (*LogSink, error) {
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
	ret := &LogSink{
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
	//   "description": "Updates or creates a sink.",
	//   "flatPath": "v2beta1/projects/{projectsId}/sinks/{sinksId}",
	//   "httpMethod": "PUT",
	//   "id": "logging.projects.sinks.update",
	//   "parameterOrder": [
	//     "projectsId",
	//     "sinksId"
	//   ],
	//   "parameters": {
	//     "projectsId": {
	//       "description": "Part of `sinkName`. Required. The resource name of the sink to update, including the parent\nresource and the sink identifier.  If the sink does not exist, this method\ncreates the sink.  Example: `\"projects/my-project-id/sinks/my-sink-id\"`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "sinksId": {
	//       "description": "Part of `sinkName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/projects/{projectsId}/sinks/{sinksId}",
	//   "request": {
	//     "$ref": "LogSink"
	//   },
	//   "response": {
	//     "$ref": "LogSink"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/logging.admin"
	//   ]
	// }

}

// method id "logging.getLogs_usage":

type V2beta1GetLogsUsageCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// GetLogsUsage: Gets logs usage.
func (r *V2beta1Service) GetLogsUsage() *V2beta1GetLogsUsageCall {
	c := &V2beta1GetLogsUsageCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// EndTime sets the optional parameter "endTime": Exclusive. Retrieve
// logs usage before this timestamp.
// The request time period cannot be less than an hour, as that's the
// minimum
// data granularity we provide.
// If not provided, current time will be used as end_time.
func (c *V2beta1GetLogsUsageCall) EndTime(endTime string) *V2beta1GetLogsUsageCall {
	c.urlParams_.Set("endTime", endTime)
	return c
}

// ResourceName sets the optional parameter "resourceName": Required.
// Resource to retrieve logs usage for.
// Examples are "projects/my-project-id",
// "organizations/google",
// "billingaccounts/ABC1234"
func (c *V2beta1GetLogsUsageCall) ResourceName(resourceName string) *V2beta1GetLogsUsageCall {
	c.urlParams_.Set("resourceName", resourceName)
	return c
}

// ResourceTier sets the optional parameter "resourceTier": Required.
// The Stackdriver tier to retrieve logs usage for.
//
// Possible values:
//   "TIER_UNSPECIFIED"
//   "FREE"
//   "PREMIUM"
func (c *V2beta1GetLogsUsageCall) ResourceTier(resourceTier string) *V2beta1GetLogsUsageCall {
	c.urlParams_.Set("resourceTier", resourceTier)
	return c
}

// StartTime sets the optional parameter "startTime": Required.
// Inclusive. Retrieve logs usage at or after this timestamp.
func (c *V2beta1GetLogsUsageCall) StartTime(startTime string) *V2beta1GetLogsUsageCall {
	c.urlParams_.Set("startTime", startTime)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *V2beta1GetLogsUsageCall) Fields(s ...googleapi.Field) *V2beta1GetLogsUsageCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *V2beta1GetLogsUsageCall) IfNoneMatch(entityTag string) *V2beta1GetLogsUsageCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *V2beta1GetLogsUsageCall) Context(ctx context.Context) *V2beta1GetLogsUsageCall {
	c.ctx_ = ctx
	return c
}

func (c *V2beta1GetLogsUsageCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2beta1/logs_usage")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.SetOpaque(req.URL)
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "logging.getLogs_usage" call.
// Exactly one of *GetLogsUsageResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *GetLogsUsageResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *V2beta1GetLogsUsageCall) Do(opts ...googleapi.CallOption) (*GetLogsUsageResponse, error) {
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
	ret := &GetLogsUsageResponse{
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
	//   "description": "Gets logs usage.",
	//   "flatPath": "v2beta1/logs_usage",
	//   "httpMethod": "GET",
	//   "id": "logging.getLogs_usage",
	//   "parameterOrder": [],
	//   "parameters": {
	//     "endTime": {
	//       "description": "Optional. Exclusive. Retrieve logs usage before this timestamp.\nThe request time period cannot be less than an hour, as that's the minimum\ndata granularity we provide.\nIf not provided, current time will be used as end_time.",
	//       "format": "google-datetime",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "resourceName": {
	//       "description": "Required. Resource to retrieve logs usage for.\nExamples are \"projects/my-project-id\", \"organizations/google\",\n\"billingaccounts/ABC1234\"",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "resourceTier": {
	//       "description": "Required. The Stackdriver tier to retrieve logs usage for.",
	//       "enum": [
	//         "TIER_UNSPECIFIED",
	//         "FREE",
	//         "PREMIUM"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "startTime": {
	//       "description": "Required. Inclusive. Retrieve logs usage at or after this timestamp.",
	//       "format": "google-datetime",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2beta1/logs_usage",
	//   "response": {
	//     "$ref": "GetLogsUsageResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/logging.read",
	//     "https://www.googleapis.com/auth/monitoring"
	//   ]
	// }

}
