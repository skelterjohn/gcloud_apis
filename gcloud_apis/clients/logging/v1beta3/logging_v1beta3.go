// Package logging provides access to the Google Cloud Logging API.
//
// See https://cloud.google.com/logging/docs/
//
// Usage example:
//
//   import "google.golang.org/api/logging/v1beta3"
//   ...
//   loggingService, err := logging.New(oauthHttpClient)
package logging

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

const apiId = "logging:v1beta3"
const apiName = "logging"
const apiVersion = "v1beta3"
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
	rs.LogEntries = NewProjectsLogEntriesService(s)
	rs.LogServices = NewProjectsLogServicesService(s)
	rs.Logs = NewProjectsLogsService(s)
	rs.Metrics = NewProjectsMetricsService(s)
	rs.Sinks = NewProjectsSinksService(s)
	return rs
}

type ProjectsService struct {
	s *Service

	LogEntries *ProjectsLogEntriesService

	LogServices *ProjectsLogServicesService

	Logs *ProjectsLogsService

	Metrics *ProjectsMetricsService

	Sinks *ProjectsSinksService
}

func NewProjectsLogEntriesService(s *Service) *ProjectsLogEntriesService {
	rs := &ProjectsLogEntriesService{s: s}
	return rs
}

type ProjectsLogEntriesService struct {
	s *Service
}

func NewProjectsLogServicesService(s *Service) *ProjectsLogServicesService {
	rs := &ProjectsLogServicesService{s: s}
	rs.Indexes = NewProjectsLogServicesIndexesService(s)
	rs.Sinks = NewProjectsLogServicesSinksService(s)
	return rs
}

type ProjectsLogServicesService struct {
	s *Service

	Indexes *ProjectsLogServicesIndexesService

	Sinks *ProjectsLogServicesSinksService
}

func NewProjectsLogServicesIndexesService(s *Service) *ProjectsLogServicesIndexesService {
	rs := &ProjectsLogServicesIndexesService{s: s}
	return rs
}

type ProjectsLogServicesIndexesService struct {
	s *Service
}

func NewProjectsLogServicesSinksService(s *Service) *ProjectsLogServicesSinksService {
	rs := &ProjectsLogServicesSinksService{s: s}
	return rs
}

type ProjectsLogServicesSinksService struct {
	s *Service
}

func NewProjectsLogsService(s *Service) *ProjectsLogsService {
	rs := &ProjectsLogsService{s: s}
	rs.Entries = NewProjectsLogsEntriesService(s)
	rs.Sinks = NewProjectsLogsSinksService(s)
	return rs
}

type ProjectsLogsService struct {
	s *Service

	Entries *ProjectsLogsEntriesService

	Sinks *ProjectsLogsSinksService
}

func NewProjectsLogsEntriesService(s *Service) *ProjectsLogsEntriesService {
	rs := &ProjectsLogsEntriesService{s: s}
	return rs
}

type ProjectsLogsEntriesService struct {
	s *Service
}

func NewProjectsLogsSinksService(s *Service) *ProjectsLogsSinksService {
	rs := &ProjectsLogsSinksService{s: s}
	return rs
}

type ProjectsLogsSinksService struct {
	s *Service
}

func NewProjectsMetricsService(s *Service) *ProjectsMetricsService {
	rs := &ProjectsMetricsService{s: s}
	return rs
}

type ProjectsMetricsService struct {
	s *Service
}

func NewProjectsSinksService(s *Service) *ProjectsSinksService {
	rs := &ProjectsSinksService{s: s}
	return rs
}

type ProjectsSinksService struct {
	s *Service
}

type Empty struct {
}

type HttpRequest struct {
	// CacheHit: Whether or not an entity was served from cache
	// (with or
	// without validation).
	CacheHit bool `json:"cacheHit,omitempty"`

	// Referer: Referer (a.k.a. referrer) URL of request, as defined
	// in
	// http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html.
	Referer string `json:"referer,omitempty"`

	// RemoteIp: IP address of the client who issues the HTTP request. Could
	// be either IPv4
	// or IPv6.
	RemoteIp string `json:"remoteIp,omitempty"`

	// RequestMethod: Request method, such as `GET`, `HEAD`, `PUT` or
	// `POST`.
	RequestMethod string `json:"requestMethod,omitempty"`

	// RequestSize: Size of the HTTP request message in bytes, including
	// request headers and
	// the request body.
	RequestSize int64 `json:"requestSize,omitempty,string"`

	// RequestUrl: Contains the scheme (http|https), the host name, the path
	// and the query
	// portion of the URL that was requested.
	RequestUrl string `json:"requestUrl,omitempty"`

	// ResponseSize: Size of the HTTP response message in bytes sent back to
	// the client,
	// including response headers and response body.
	ResponseSize int64 `json:"responseSize,omitempty,string"`

	// Status: A response code indicates the status of response, e.g., 200.
	Status int64 `json:"status,omitempty"`

	// UserAgent: User agent sent by the client, e.g., "Mozilla/4.0
	// (compatible; MSIE 6.0;
	// Windows 98; Q312461; .NET CLR 1.0.3705)".
	UserAgent string `json:"userAgent,omitempty"`

	// ValidatedWithOriginServer: Whether or not the response was validated
	// with the origin server before
	// being served from cache. This field is
	// only meaningful if cache_hit is
	// True.
	ValidatedWithOriginServer bool `json:"validatedWithOriginServer,omitempty"`
}

type ListLogEntriesResponse struct {
	// Entries: A list of log entries.  Fewer than `pageSize` entries may
	// be
	// returned, but that is not an indication that there are no more
	// entries.
	Entries []*LogEntry `json:"entries,omitempty"`

	// NextPageToken: If there are more results, then `nextPageToken` is
	// returned in the
	// response.  To get the next batch of entries, use the
	// value of
	// `nextPageToken` as `pageToken` in the next call of
	// `ListLogEntries`.
	// If `nextPageToken` is empty, then there are no more
	// results.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type ListLogMetricsResponse struct {
	// Metrics: The list of metrics that was requested.
	Metrics []*LogMetric `json:"metrics,omitempty"`

	// NextPageToken: If there are more results, then `nextPageToken` is
	// returned in the
	// response.  To get the next batch of entries, use the
	// value of
	// `nextPageToken` as `pageToken` in the next call of
	// `ListLogMetrics`.
	// If `nextPageToken` is empty, then there are no more
	// results.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type ListLogServiceIndexesResponse struct {
	// NextPageToken: If there are more results, then `nextPageToken` is
	// returned in the
	// response.  To get the next batch of indexes, use the
	// value of
	// `nextPageToken` as `pageToken` in the next call of
	// `ListLogServiceIndexes`.
	// If `nextPageToken` is empty, then there are
	// no more results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServiceIndexPrefixes: A list of log service index values.
	// Each index
	// value has the form `"/value1/value2/..."`,
	// where `value1` is a value
	// in the primary index, `value2` is
	// a value in the secondary index, and
	// so forth.
	ServiceIndexPrefixes []string `json:"serviceIndexPrefixes,omitempty"`
}

type ListLogServiceSinksResponse struct {
	// Sinks: The requested log service sinks. If a returned `LogSink`
	// object has
	// an empty `destination` field, the client can retrieve the
	// complete
	// `LogSink` object by calling `logServices.sinks.get`.
	Sinks []*LogSink `json:"sinks,omitempty"`
}

type ListLogServicesResponse struct {
	// LogServices: A list of log services.
	LogServices []*LogService `json:"logServices,omitempty"`

	// NextPageToken: If there are more results, then `nextPageToken` is
	// returned in the
	// response.  To get the next batch of services, use the
	// value of
	// `nextPageToken` as `pageToken` in the next call of
	// `ListLogServices`.
	// If `nextPageToken` is empty, then there are no
	// more results.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type ListLogSinksResponse struct {
	// Sinks: The requested log sinks. If a returned `LogSink` object has
	// an
	// empty `destination` field, the client can retrieve the
	// complete
	// `LogSink` object by calling `log.sinks.get`.
	Sinks []*LogSink `json:"sinks,omitempty"`
}

type ListLogsResponse struct {
	// Logs: A list of log descriptions matching the criteria.
	Logs []*Log `json:"logs,omitempty"`

	// NextPageToken: If there are more results, then `nextPageToken` is
	// returned in the
	// response.  To get the next batch of logs, use the
	// value of
	// `nextPageToken` as `pageToken` in the next call of
	// `ListLogs`.
	// If `nextPageToken` is empty, then there are no more
	// results.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type ListSinksResponse struct {
	// Sinks: The requested sinks.  If a returned `LogSink` object has
	// an
	// empty `destination` field, the client can retrieve the
	// complete
	// `LogSink` object by calling `projects.sinks.get`.
	Sinks []*LogSink `json:"sinks,omitempty"`
}

type Log struct {
	// DisplayName: _Optional._ The common name of the log.  Example:
	// `"request_log"`.
	DisplayName string `json:"displayName,omitempty"`

	// Name: The resource name of the log.
	// Example:
	// `"/projects/my-gcp-project-id/logs/LOG_NAME"`, where
	// `LOG_NAME`
	// is the URL-encoded given name of the log.  The log
	// includes those
	// log entries whose `LogEntry.log` field contains this
	// given name.
	// To avoid name collisions, it is a best practice to prefix
	// the
	// given log name with the service name, but this is not
	// required.
	// Examples of log given
	// names:
	// `"appengine.googleapis.com/request_log"`, `"apache-access"`.
	Name string `json:"name,omitempty"`

	// PayloadType: _Optional_. A URI representing the expected payload type
	// for log entries.
	PayloadType string `json:"payloadType,omitempty"`
}

type LogEntry struct {
	// HttpRequest: Information about the HTTP request associated with this
	// log entry,
	// if applicable.
	HttpRequest *HttpRequest `json:"httpRequest,omitempty"`

	// InsertId: A unique ID for the log entry. If you provide this field,
	// the logging
	// service considers other log entries in the same log with
	// the same ID as
	// duplicates which can be removed.
	InsertId string `json:"insertId,omitempty"`

	// Log: The log to which this entry belongs. When a log entry is
	// ingested,
	// the value of this field is set by the logging system.
	Log string `json:"log,omitempty"`

	// Metadata: Information about the log entry.
	Metadata *LogEntryMetadata `json:"metadata,omitempty"`

	Operation *LogEntryOperation `json:"operation,omitempty"`

	// ProtoPayload: The log entry payload, represented as a protocol buffer
	// that is
	// expressed as a JSON object. You can only pass
	// `protoPayload`
	// values that belong to a set of approved types.
	ProtoPayload *LogEntryProtoPayload `json:"protoPayload,omitempty"`

	// StructPayload: The log entry payload, represented as a structure
	// that
	// is expressed as a JSON object.
	StructPayload *LogEntryStructPayload `json:"structPayload,omitempty"`

	// TextPayload: The log entry payload, represented as a Unicode string
	// (UTF-8).
	TextPayload string `json:"textPayload,omitempty"`

	WriterEmailAddress string `json:"writerEmailAddress,omitempty"`
}

type LogEntryProtoPayload struct {
}

type LogEntryStructPayload struct {
}

type LogEntryMetadata struct {
	// Labels: A set of (key, value) data that provides additional
	// information about
	// the log entry.
	// If the log entry is from one of the
	// Google Cloud Platform sources listed
	// below, the indicated (key,
	// value) information must be provided:
	//
	// Google App Engine, service_name
	// `appengine.googleapis.com`:
	//
	//
	// "appengine.googleapis.com/module_id", <module ID>
	//
	// "appengine.googleapis.com/version_id", <version ID>
	//           and one
	// of:
	//       "appengine.googleapis.com/replica_index", <instance index>
	//
	//      "appengine.googleapis.com/clone_id", <instance ID>
	//
	//     or else
	// provide the following Compute Engine labels:
	//
	// Google Compute Engine,
	// service_name `compute.googleapis.com`:
	//
	//
	// "compute.googleapis.com/resource_type", "instance"
	//
	// "compute.googleapis.com/resource_id", <instance ID>
	Labels map[string]string `json:"labels,omitempty"`

	// ProjectId: The project ID of the Google Cloud Platform service that
	// created the log
	// entry.
	ProjectId string `json:"projectId,omitempty"`

	// ProjectNumber: This field is populated by the API at ingestion time.
	ProjectNumber int64 `json:"projectNumber,omitempty,string"`

	// Region: The region name of the Google Cloud Platform service that
	// created the log
	// entry.  For example, `"us-central1"`.
	Region string `json:"region,omitempty"`

	// ServiceName: The API name of the Google Cloud Platform service that
	// created the log
	// entry.  For example, `"compute.googleapis.com"`.
	ServiceName string `json:"serviceName,omitempty"`

	// Severity: The severity of the log entry.
	Severity string `json:"severity,omitempty"`

	// Timestamp: The time the event described by the log entry
	// occurred.
	// Timestamps must be later than January 1, 1970.
	Timestamp string `json:"timestamp,omitempty"`

	// UserId: The fully-qualified email address of the authenticated user
	// that performed
	// or requested the action represented by the log entry.
	// If the log entry does
	// not apply to an action taken by an
	// authenticated user, then the field
	// should be empty.
	UserId string `json:"userId,omitempty"`

	// Zone: The zone of the Google Cloud Platform service that created the
	// log entry.
	// For example, `"us-central1-a"`.
	Zone string `json:"zone,omitempty"`
}

type LogEntryOperation struct {
	// First: True for the first entry associated with `id`.
	First bool `json:"first,omitempty"`

	// Id: An opaque identifier. A producer of log entries should ensure
	// that `id` is
	// only reused for entries related to one operation.
	//
	Id string `json:"id,omitempty"`

	// Last: True for the last entry associated with `id`.
	Last bool `json:"last,omitempty"`

	// Producer: Ensures the operation can be uniquely identified. The
	// combination of `id`
	// and `producer` should be made globally unique by
	// filling `producer` with a
	// value that disambiguates the service that
	// created `id`.
	//
	Producer string `json:"producer,omitempty"`
}

type LogError struct {
	// Resource: A resource name associated with this error. For
	// example,
	// the name of a Cloud Storage bucket that has insufficient
	// permissions
	// to be a destination for log entries.
	Resource string `json:"resource,omitempty"`

	// Status: The error description, including a classification code,
	// an
	// error message, and other details.
	Status *Status `json:"status,omitempty"`

	// TimeNanos: The time the error was observed, in nanoseconds since the
	// Unix epoch.
	TimeNanos int64 `json:"timeNanos,omitempty,string"`
}

type LogLine struct {
	// LogMessage: App provided log message.
	LogMessage string `json:"logMessage,omitempty"`

	// Severity: Severity of log.
	Severity string `json:"severity,omitempty"`

	// SourceLocation: Line of code that generated this log message.
	SourceLocation *SourceLocation `json:"sourceLocation,omitempty"`

	// Time: Time when log entry was made.  May be inaccurate.
	Time string `json:"time,omitempty"`
}

type LogMetric struct {
	// Description: A description of this metric.
	Description string `json:"description,omitempty"`

	// Filter: An [advanced logs
	// filter](/logging/docs/view/advanced_filters).
	// Example: `"log:syslog
	// AND metadata.severity>=ERROR"`.
	Filter string `json:"filter,omitempty"`

	// Name: The client-assigned name for this metric, such
	// as
	// `"severe_errors"`.  Metric names are limited to 1000
	// characters
	// and can include only the following characters: `A-Z`,
	// `a-z`,
	// `0-9`, and the special characters `_-.,+!*',()%/\`.  The
	// slash
	// character (`/`) denotes a hierarchy of name pieces, and it
	// cannot
	// be the first character of the name.
	Name string `json:"name,omitempty"`
}

type LogService struct {
	// IndexKeys: A list of the names of the keys used to index and
	// label
	// individual log entries from this service.  The first two keys
	// are
	// used as the primary and secondary index, respectively.
	// Additional
	// keys may be used to label the entries.  For example,
	// App Engine
	// indexes its entries by module and by version, so its
	// `indexKeys`
	// field is the following:
	//
	//     [ "appengine.googleapis.com/module_id",
	//
	//      "appengine.googleapis.com/version_id" ]
	IndexKeys []string `json:"indexKeys,omitempty"`

	// Name: The service's name. Example: `"appengine.googleapis.com"`.
	// Log
	// names beginning with this string are reserved for this service.
	// This
	// value can appear in the `LogEntry.metadata.serviceName` field of
	// log
	// entries associated with this log service.
	Name string `json:"name,omitempty"`
}

type LogSink struct {
	// Destination: The resource name of the destination.
	// Cloud Logging
	// writes designated log entries to this destination.
	// For example,
	// `"storage.googleapis.com/my-output-bucket"`.
	Destination string `json:"destination,omitempty"`

	// Errors: _Output only._ If any errors occur when invoking a sink
	// method,
	// then this field contains descriptions of the errors.
	Errors []*LogError `json:"errors,omitempty"`

	// Filter: An advanced logs filter. If present, only log entries
	// matching the filter
	// are written.  Only project sinks use this field;
	// log sinks and log service
	// sinks must not include a filter.
	Filter string `json:"filter,omitempty"`

	// Name: The client-assigned name of this sink. For
	// example,
	// `"my-syslog-sink"`.  The name must be unique among the sinks
	// of a
	// similar kind in the project.
	Name string `json:"name,omitempty"`
}

type RequestLog struct {
	// AppEngineRelease: App Engine release version string.
	AppEngineRelease string `json:"appEngineRelease,omitempty"`

	// AppId: Identifies the application that handled this request.
	AppId string `json:"appId,omitempty"`

	// Cost: An indication of the relative cost of serving this request.
	Cost float64 `json:"cost,omitempty"`

	// EndTime: Time at which request was known to end processing.
	EndTime string `json:"endTime,omitempty"`

	// Finished: If true, represents a finished request.  Otherwise, the
	// request is active.
	Finished bool `json:"finished,omitempty"`

	// Host: The Internet host and port number of the resource being
	// requested.
	Host string `json:"host,omitempty"`

	// HttpVersion: HTTP version of request.
	HttpVersion string `json:"httpVersion,omitempty"`

	// InstanceId: An opaque identifier for the instance that handled the
	// request.
	InstanceId string `json:"instanceId,omitempty"`

	// InstanceIndex: If the instance that processed this request was
	// individually addressable
	// (i.e. belongs to a manually scaled module),
	// this is the index of the
	// instance.
	InstanceIndex int64 `json:"instanceIndex,omitempty"`

	// Ip: Origin IP address.
	Ip string `json:"ip,omitempty"`

	// Latency: Latency of the request.
	Latency string `json:"latency,omitempty"`

	// Line: List of log lines emitted by the application while serving this
	// request,
	// if requested.
	Line []*LogLine `json:"line,omitempty"`

	// MegaCycles: Number of CPU megacycles used to process request.
	MegaCycles int64 `json:"megaCycles,omitempty,string"`

	// Method: Request method, such as `GET`, `HEAD`, `PUT`, `POST`, or
	// `DELETE`.
	Method string `json:"method,omitempty"`

	// ModuleId: Identifies the module of the application that handled this
	// request.
	ModuleId string `json:"moduleId,omitempty"`

	// Nickname: A string that identifies a logged-in user who made this
	// request, or empty
	// if the user is not logged in.
	//
	// Most likely, this is
	// the part of the user's email before the '@' sign.  The
	// field value is
	// the same for different requests from the same user, but
	// different
	// users may have a similar name.  This information is also
	// available to
	// the application via Users API.
	//
	// This field will be populated starting
	// with App Engine 1.9.21.
	//
	Nickname string `json:"nickname,omitempty"`

	// PendingTime: Time this request spent in the pending request queue, if
	// it was pending at
	// all.
	PendingTime string `json:"pendingTime,omitempty"`

	// Referrer: Referrer URL of request.
	Referrer string `json:"referrer,omitempty"`

	// RequestId: Globally unique identifier for a request, based on request
	// start time.
	// Request IDs for requests which started later will compare
	// greater as
	// strings than those for requests which started earlier.
	RequestId string `json:"requestId,omitempty"`

	// Resource: Contains the path and query portion of the URL that was
	// requested. For
	// example, if the URL was
	// "http://example.com/app?name=val", the resource
	// would be
	// "/app?name=val". Any trailing fragment (separated by a '#'
	// character)
	// will not be included.
	Resource string `json:"resource,omitempty"`

	// ResponseSize: Size in bytes sent back to client by request.
	ResponseSize int64 `json:"responseSize,omitempty,string"`

	// SourceReference: Source code for the application that handled this
	// request. There can be
	// more than one source reference per deployed
	// application if source code is
	// distributed among multiple
	// repositories.
	SourceReference []*SourceReference `json:"sourceReference,omitempty"`

	// StartTime: Time at which request was known to have begun processing.
	StartTime string `json:"startTime,omitempty"`

	// Status: Response status of request.
	Status int64 `json:"status,omitempty"`

	// TaskName: Task name of the request (for an offline request).
	TaskName string `json:"taskName,omitempty"`

	// TaskQueueName: Queue name of the request (for an offline request).
	TaskQueueName string `json:"taskQueueName,omitempty"`

	// TraceId: Cloud Trace identifier of the trace for this request.
	TraceId string `json:"traceId,omitempty"`

	// UrlMapEntry: File or class within URL mapping used for request.
	// Useful for tracking
	// down the source code which was responsible for
	// managing request.
	// Especially for multiply mapped handlers.
	UrlMapEntry string `json:"urlMapEntry,omitempty"`

	// UserAgent: User agent used for making request.
	UserAgent string `json:"userAgent,omitempty"`

	// VersionId: Version of the application that handled this request.
	VersionId string `json:"versionId,omitempty"`

	// WasLoadingRequest: Was this request a loading request for this
	// instance?
	WasLoadingRequest bool `json:"wasLoadingRequest,omitempty"`
}

type SourceLocation struct {
	// File: Source file name. May or may not be a fully qualified name,
	// depending on
	// the runtime environment.
	File string `json:"file,omitempty"`

	// FunctionName: Human-readable name of the function or method being
	// invoked, with optional
	// context such as the class or package name, for
	// use in contexts such as the
	// logs viewer where file:line number is
	// less meaningful. This may vary by
	// language, for example:
	//   in Java:
	// qual.if.ied.Class.method
	//   in Go: dir/package.func
	//   in Python:
	// function
	// ...
	FunctionName string `json:"functionName,omitempty"`

	// Line: Line within the source file.
	Line int64 `json:"line,omitempty,string"`
}

type SourceReference struct {
	// Repository: Optional. A URI string identifying the
	// repository.
	// Example:
	// "https://github.com/GoogleCloudPlatform/kubernetes.git"
	Repository string `json:"repository,omitempty"`

	// RevisionId: The canonical (and persistent) identifier of the deployed
	// revision.
	// Example (git): "0035781c50ec7aa23385dc841529ce8a4b70db1b"
	RevisionId string `json:"revisionId,omitempty"`
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

type WriteLogEntriesRequest struct {
	// CommonLabels: Metadata labels that apply to all log entries in this
	// request, so that you
	// don't have to repeat them in each log entry's
	// `metadata.labels` field.  If
	// any of the log entries contains a (key,
	// value) with the same key that is in
	// `commonLabels`, then the entry's
	// (key, value) overrides the one in
	// `commonLabels`.
	CommonLabels map[string]string `json:"commonLabels,omitempty"`

	// Entries: Log entries to insert.
	Entries []*LogEntry `json:"entries,omitempty"`
}

type WriteLogEntriesResponse struct {
}

// method id "logging.projects.logEntries.list":

type ProjectsLogEntriesListCall struct {
	s          *Service
	projectsId string
	opt_       map[string]interface{}
}

// List: Lists log entries in the specified project.
//
func (r *ProjectsLogEntriesService) List(projectsId string) *ProjectsLogEntriesListCall {
	c := &ProjectsLogEntriesListCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectsId = projectsId
	return c
}

// Filter sets the optional parameter "filter": An [advanced logs
// filter](/logging/docs/view/advanced_filters).
// The response includes
// only entries that match the filter.
// If `filter` is empty, then all
// entries in all logs are retrieved.
func (c *ProjectsLogEntriesListCall) Filter(filter string) *ProjectsLogEntriesListCall {
	c.opt_["filter"] = filter
	return c
}

// OrderBy sets the optional parameter "orderBy": Sort order of the
// results, consisting of a `LogEntry` field optionally
// followed by a
// space and `desc`.  Examples:
// `"metadata.timestamp"`,
// `"metadata.timestamp desc"`.  The only
// `LogEntry` field supported for
// sorting is `metadata.timestamp`. The
// default sort order is ascending (from
// older to newer entries) unless
// `desc` is appended.
func (c *ProjectsLogEntriesListCall) OrderBy(orderBy string) *ProjectsLogEntriesListCall {
	c.opt_["orderBy"] = orderBy
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of entries to return per request.  Fewer entries may be
// returned, but
// that is not an indication that there are no more entries.
func (c *ProjectsLogEntriesListCall) PageSize(pageSize int64) *ProjectsLogEntriesListCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken": An opaque token,
// returned as `nextPageToken` by a prior `ListLogEntries`
// operation. If
// a page token is specified, other request parameters must
// match the
// parameters from the request that generated the page token.
func (c *ProjectsLogEntriesListCall) PageToken(pageToken string) *ProjectsLogEntriesListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogEntriesListCall) Fields(s ...googleapi.Field) *ProjectsLogEntriesListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsLogEntriesListCall) Do() (*ListLogEntriesResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["filter"]; ok {
		params.Set("filter", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["orderBy"]; ok {
		params.Set("orderBy", fmt.Sprintf("%v", v))
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectsId}/logEntries")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectsId": c.projectsId,
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
	var ret *ListLogEntriesResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists log entries in the specified project.\n",
	//   "flatPath": "v1beta3/projects/{projectsId}/logEntries",
	//   "httpMethod": "GET",
	//   "id": "logging.projects.logEntries.list",
	//   "parameterOrder": [
	//     "projectsId"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "An [advanced logs filter](/logging/docs/view/advanced_filters).\nThe response includes only entries that match the filter.\nIf `filter` is empty, then all entries in all logs are retrieved.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "orderBy": {
	//       "description": "Sort order of the results, consisting of a `LogEntry` field optionally\nfollowed by a space and `desc`.  Examples: `\"metadata.timestamp\"`,\n`\"metadata.timestamp desc\"`.  The only `LogEntry` field supported for\nsorting is `metadata.timestamp`. The default sort order is ascending (from\nolder to newer entries) unless `desc` is appended.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "The maximum number of entries to return per request.  Fewer entries may be\nreturned, but that is not an indication that there are no more entries.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "An opaque token, returned as `nextPageToken` by a prior `ListLogEntries`\noperation. If a page token is specified, other request parameters must\nmatch the parameters from the request that generated the page token.",
	//       "format": "byte",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `projectName`. The resource name of the project from which to retrieve log entries.  The\nlog service or log containing the entries is specified in the `filter`\nparameter.  Example: `projects/my_project_id`.\n",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/logEntries",
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

// method id "logging.projects.logServices.list":

type ProjectsLogServicesListCall struct {
	s          *Service
	projectsId string
	opt_       map[string]interface{}
}

// List: Lists the log services that have log entries in this project.
//
func (r *ProjectsLogServicesService) List(projectsId string) *ProjectsLogServicesListCall {
	c := &ProjectsLogServicesListCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectsId = projectsId
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of `LogService` objects to return in one operation.
func (c *ProjectsLogServicesListCall) PageSize(pageSize int64) *ProjectsLogServicesListCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken": An opaque token,
// returned as `nextPageToken` by a prior `ListLogServices`
// operation.
// If `pageToken` is supplied, then the other fields of this
// request are
// ignored, and instead the previous `ListLogServices` operation
// is
// continued.
func (c *ProjectsLogServicesListCall) PageToken(pageToken string) *ProjectsLogServicesListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogServicesListCall) Fields(s ...googleapi.Field) *ProjectsLogServicesListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsLogServicesListCall) Do() (*ListLogServicesResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["pageSize"]; ok {
		params.Set("pageSize", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectsId}/logServices")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectsId": c.projectsId,
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
	var ret *ListLogServicesResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists the log services that have log entries in this project.\n",
	//   "flatPath": "v1beta3/projects/{projectsId}/logServices",
	//   "httpMethod": "GET",
	//   "id": "logging.projects.logServices.list",
	//   "parameterOrder": [
	//     "projectsId"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "The maximum number of `LogService` objects to return in one operation.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "An opaque token, returned as `nextPageToken` by a prior `ListLogServices`\noperation.  If `pageToken` is supplied, then the other fields of this\nrequest are ignored, and instead the previous `ListLogServices` operation\nis continued.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `projectName`. The resource name of the project whose services are to be listed.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/logServices",
	//   "response": {
	//     "$ref": "ListLogServicesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only",
	//     "https://www.googleapis.com/auth/logging.admin",
	//     "https://www.googleapis.com/auth/logging.read"
	//   ]
	// }

}

// method id "logging.projects.logServices.indexes.list":

type ProjectsLogServicesIndexesListCall struct {
	s             *Service
	projectsId    string
	logServicesId string
	opt_          map[string]interface{}
}

// List: Lists the current index values for a log service.
//
func (r *ProjectsLogServicesIndexesService) List(projectsId string, logServicesId string) *ProjectsLogServicesIndexesListCall {
	c := &ProjectsLogServicesIndexesListCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectsId = projectsId
	c.logServicesId = logServicesId
	return c
}

// Depth sets the optional parameter "depth": A non-negative integer
// that limits the number of levels of the index
// hierarchy that are
// returned.
// If `depth` is 1 (default), only the first index key value
// is returned.
// If `depth` is 2, both primary and secondary key values
// are returned.
// If `depth` is 0, the depth is the number of
// slash-separators in the
// `indexPrefix` field, not counting a slash
// appearing as the last character
// of the prefix.
// If the `indexPrefix`
// field is empty, the default depth is 1.
// It is an error for `depth` to
// be any positive value
// less than the number of components in
// `indexPrefix`.
func (c *ProjectsLogServicesIndexesListCall) Depth(depth int64) *ProjectsLogServicesIndexesListCall {
	c.opt_["depth"] = depth
	return c
}

// IndexPrefix sets the optional parameter "indexPrefix": Restricts the
// index values returned to be those with a specified prefix
// for each
// index key.
// This field has the form `"/prefix1/prefix2/..."`, in order
// corresponding
// to the `LogService indexKeys`.
// Non-empty prefixes must
// begin with `/`.
// For example, App Engine's two keys are the module ID
// and the version ID.
// Following is the effect of using various values
// for `indexPrefix`:
//
// +  `"/Mod/"` retrieves `/Mod/10` and `/Mod/11`
// but not `/ModA/10`.
// +  `"/Mod` retrieves `/Mod/10`, `/Mod/11` and
// `/ModA/10` but not `/XXX/33`.
// +  `"/Mod/1"` retrieves `/Mod/10` and
// `/Mod/11` but not `/ModA/10`.
// +  `"/Mod/10/"` retrieves `/Mod/10`
// only.
// +  An empty prefix or `"/"` retrieves all values.
func (c *ProjectsLogServicesIndexesListCall) IndexPrefix(indexPrefix string) *ProjectsLogServicesIndexesListCall {
	c.opt_["indexPrefix"] = indexPrefix
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of log service index resources to return in one
// operation.
func (c *ProjectsLogServicesIndexesListCall) PageSize(pageSize int64) *ProjectsLogServicesIndexesListCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken": An opaque token,
// returned as `nextPageToken` by a prior
// `ListLogServiceIndexes`
// operation.  If `pageToken` is supplied, then the
// other fields of this
// request are ignored, and instead the previous
// `ListLogServiceIndexes`
// operation is continued.
func (c *ProjectsLogServicesIndexesListCall) PageToken(pageToken string) *ProjectsLogServicesIndexesListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogServicesIndexesListCall) Fields(s ...googleapi.Field) *ProjectsLogServicesIndexesListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsLogServicesIndexesListCall) Do() (*ListLogServiceIndexesResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["depth"]; ok {
		params.Set("depth", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["indexPrefix"]; ok {
		params.Set("indexPrefix", fmt.Sprintf("%v", v))
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectsId}/logServices/{logServicesId}/indexes")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectsId":    c.projectsId,
		"logServicesId": c.logServicesId,
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
	var ret *ListLogServiceIndexesResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists the current index values for a log service.\n",
	//   "flatPath": "v1beta3/projects/{projectsId}/logServices/{logServicesId}/indexes",
	//   "httpMethod": "GET",
	//   "id": "logging.projects.logServices.indexes.list",
	//   "parameterOrder": [
	//     "projectsId",
	//     "logServicesId"
	//   ],
	//   "parameters": {
	//     "depth": {
	//       "description": "A non-negative integer that limits the number of levels of the index\nhierarchy that are returned.\nIf `depth` is 1 (default), only the first index key value is returned.\nIf `depth` is 2, both primary and secondary key values are returned.\nIf `depth` is 0, the depth is the number of slash-separators in the\n`indexPrefix` field, not counting a slash appearing as the last character\nof the prefix.\nIf the `indexPrefix` field is empty, the default depth is 1.\nIt is an error for `depth` to be any positive value\nless than the number of components in `indexPrefix`.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "indexPrefix": {
	//       "description": "Restricts the index values returned to be those with a specified prefix\nfor each index key.\nThis field has the form `\"/prefix1/prefix2/...\"`, in order corresponding\nto the `LogService indexKeys`.\nNon-empty prefixes must begin with `/`.\nFor example, App Engine's two keys are the module ID and the version ID.\nFollowing is the effect of using various values for `indexPrefix`:\n\n+  `\"/Mod/\"` retrieves `/Mod/10` and `/Mod/11` but not `/ModA/10`.\n+  `\"/Mod` retrieves `/Mod/10`, `/Mod/11` and `/ModA/10` but not `/XXX/33`.\n+  `\"/Mod/1\"` retrieves `/Mod/10` and `/Mod/11` but not `/ModA/10`.\n+  `\"/Mod/10/\"` retrieves `/Mod/10` only.\n+  An empty prefix or `\"/\"` retrieves all values.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "logServicesId": {
	//       "description": "Part of `serviceName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "The maximum number of log service index resources to return in one\noperation.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "An opaque token, returned as `nextPageToken` by a prior\n`ListLogServiceIndexes` operation.  If `pageToken` is supplied, then the\nother fields of this request are ignored, and instead the previous\n`ListLogServiceIndexes` operation is continued.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `serviceName`. The resource name of a log service whose service indexes are requested.\nExample:\n`\"projects/my-project-id/logServices/appengine.googleapis.com\"`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/logServices/{logServicesId}/indexes",
	//   "response": {
	//     "$ref": "ListLogServiceIndexesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only",
	//     "https://www.googleapis.com/auth/logging.admin",
	//     "https://www.googleapis.com/auth/logging.read"
	//   ]
	// }

}

// method id "logging.projects.logServices.sinks.create":

type ProjectsLogServicesSinksCreateCall struct {
	s             *Service
	projectsId    string
	logServicesId string
	logsink       *LogSink
	opt_          map[string]interface{}
}

// Create: Creates a log service sink.
// All log entries from a specified
// log service are written to the
// destination.
func (r *ProjectsLogServicesSinksService) Create(projectsId string, logServicesId string, logsink *LogSink) *ProjectsLogServicesSinksCreateCall {
	c := &ProjectsLogServicesSinksCreateCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectsId = projectsId
	c.logServicesId = logServicesId
	c.logsink = logsink
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogServicesSinksCreateCall) Fields(s ...googleapi.Field) *ProjectsLogServicesSinksCreateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsLogServicesSinksCreateCall) Do() (*LogSink, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.logsink)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectsId}/logServices/{logServicesId}/sinks")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectsId":    c.projectsId,
		"logServicesId": c.logServicesId,
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
	var ret *LogSink
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a log service sink.\nAll log entries from a specified log service are written to the\ndestination.",
	//   "flatPath": "v1beta3/projects/{projectsId}/logServices/{logServicesId}/sinks",
	//   "httpMethod": "POST",
	//   "id": "logging.projects.logServices.sinks.create",
	//   "parameterOrder": [
	//     "projectsId",
	//     "logServicesId"
	//   ],
	//   "parameters": {
	//     "logServicesId": {
	//       "description": "Part of `serviceName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `serviceName`. The resource name of the log service to which the sink is bound.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/logServices/{logServicesId}/sinks",
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

// method id "logging.projects.logServices.sinks.delete":

type ProjectsLogServicesSinksDeleteCall struct {
	s             *Service
	projectsId    string
	logServicesId string
	sinksId       string
	opt_          map[string]interface{}
}

// Delete: Deletes a log service sink.
// After deletion, no new log
// entries are written to the destination.
func (r *ProjectsLogServicesSinksService) Delete(projectsId string, logServicesId string, sinksId string) *ProjectsLogServicesSinksDeleteCall {
	c := &ProjectsLogServicesSinksDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectsId = projectsId
	c.logServicesId = logServicesId
	c.sinksId = sinksId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogServicesSinksDeleteCall) Fields(s ...googleapi.Field) *ProjectsLogServicesSinksDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsLogServicesSinksDeleteCall) Do() (*Empty, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectsId}/logServices/{logServicesId}/sinks/{sinksId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectsId":    c.projectsId,
		"logServicesId": c.logServicesId,
		"sinksId":       c.sinksId,
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
	//   "description": "Deletes a log service sink.\nAfter deletion, no new log entries are written to the destination.",
	//   "flatPath": "v1beta3/projects/{projectsId}/logServices/{logServicesId}/sinks/{sinksId}",
	//   "httpMethod": "DELETE",
	//   "id": "logging.projects.logServices.sinks.delete",
	//   "parameterOrder": [
	//     "projectsId",
	//     "logServicesId",
	//     "sinksId"
	//   ],
	//   "parameters": {
	//     "logServicesId": {
	//       "description": "Part of `sinkName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `sinkName`. The resource name of the log service sink to delete.",
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
	//   "path": "v1beta3/projects/{projectsId}/logServices/{logServicesId}/sinks/{sinksId}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/logging.admin"
	//   ]
	// }

}

// method id "logging.projects.logServices.sinks.get":

type ProjectsLogServicesSinksGetCall struct {
	s             *Service
	projectsId    string
	logServicesId string
	sinksId       string
	opt_          map[string]interface{}
}

// Get: Gets a log service sink.
func (r *ProjectsLogServicesSinksService) Get(projectsId string, logServicesId string, sinksId string) *ProjectsLogServicesSinksGetCall {
	c := &ProjectsLogServicesSinksGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectsId = projectsId
	c.logServicesId = logServicesId
	c.sinksId = sinksId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogServicesSinksGetCall) Fields(s ...googleapi.Field) *ProjectsLogServicesSinksGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsLogServicesSinksGetCall) Do() (*LogSink, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectsId}/logServices/{logServicesId}/sinks/{sinksId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectsId":    c.projectsId,
		"logServicesId": c.logServicesId,
		"sinksId":       c.sinksId,
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
	var ret *LogSink
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets a log service sink.",
	//   "flatPath": "v1beta3/projects/{projectsId}/logServices/{logServicesId}/sinks/{sinksId}",
	//   "httpMethod": "GET",
	//   "id": "logging.projects.logServices.sinks.get",
	//   "parameterOrder": [
	//     "projectsId",
	//     "logServicesId",
	//     "sinksId"
	//   ],
	//   "parameters": {
	//     "logServicesId": {
	//       "description": "Part of `sinkName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `sinkName`. The resource name of the log service sink to return.",
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
	//   "path": "v1beta3/projects/{projectsId}/logServices/{logServicesId}/sinks/{sinksId}",
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

// method id "logging.projects.logServices.sinks.list":

type ProjectsLogServicesSinksListCall struct {
	s             *Service
	projectsId    string
	logServicesId string
	opt_          map[string]interface{}
}

// List: Lists log service sinks associated with a log service.
func (r *ProjectsLogServicesSinksService) List(projectsId string, logServicesId string) *ProjectsLogServicesSinksListCall {
	c := &ProjectsLogServicesSinksListCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectsId = projectsId
	c.logServicesId = logServicesId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogServicesSinksListCall) Fields(s ...googleapi.Field) *ProjectsLogServicesSinksListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsLogServicesSinksListCall) Do() (*ListLogServiceSinksResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectsId}/logServices/{logServicesId}/sinks")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectsId":    c.projectsId,
		"logServicesId": c.logServicesId,
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
	var ret *ListLogServiceSinksResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists log service sinks associated with a log service.",
	//   "flatPath": "v1beta3/projects/{projectsId}/logServices/{logServicesId}/sinks",
	//   "httpMethod": "GET",
	//   "id": "logging.projects.logServices.sinks.list",
	//   "parameterOrder": [
	//     "projectsId",
	//     "logServicesId"
	//   ],
	//   "parameters": {
	//     "logServicesId": {
	//       "description": "Part of `serviceName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `serviceName`. The log service whose sinks are wanted.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/logServices/{logServicesId}/sinks",
	//   "response": {
	//     "$ref": "ListLogServiceSinksResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only",
	//     "https://www.googleapis.com/auth/logging.admin",
	//     "https://www.googleapis.com/auth/logging.read"
	//   ]
	// }

}

// method id "logging.projects.logServices.sinks.update":

type ProjectsLogServicesSinksUpdateCall struct {
	s             *Service
	projectsId    string
	logServicesId string
	sinksId       string
	logsink       *LogSink
	opt_          map[string]interface{}
}

// Update: Updates a log service sink.
// If the sink does not exist, it is
// created.
func (r *ProjectsLogServicesSinksService) Update(projectsId string, logServicesId string, sinksId string, logsink *LogSink) *ProjectsLogServicesSinksUpdateCall {
	c := &ProjectsLogServicesSinksUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectsId = projectsId
	c.logServicesId = logServicesId
	c.sinksId = sinksId
	c.logsink = logsink
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogServicesSinksUpdateCall) Fields(s ...googleapi.Field) *ProjectsLogServicesSinksUpdateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsLogServicesSinksUpdateCall) Do() (*LogSink, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.logsink)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectsId}/logServices/{logServicesId}/sinks/{sinksId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectsId":    c.projectsId,
		"logServicesId": c.logServicesId,
		"sinksId":       c.sinksId,
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
	var ret *LogSink
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates a log service sink.\nIf the sink does not exist, it is created.",
	//   "flatPath": "v1beta3/projects/{projectsId}/logServices/{logServicesId}/sinks/{sinksId}",
	//   "httpMethod": "PUT",
	//   "id": "logging.projects.logServices.sinks.update",
	//   "parameterOrder": [
	//     "projectsId",
	//     "logServicesId",
	//     "sinksId"
	//   ],
	//   "parameters": {
	//     "logServicesId": {
	//       "description": "Part of `sinkName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `sinkName`. The resource name of the log service sink to update.",
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
	//   "path": "v1beta3/projects/{projectsId}/logServices/{logServicesId}/sinks/{sinksId}",
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

// method id "logging.projects.logs.delete":

type ProjectsLogsDeleteCall struct {
	s          *Service
	projectsId string
	logsId     string
	opt_       map[string]interface{}
}

// Delete: Deletes a log and all its log entries.
// The log will reappear
// if it receives new entries.
//
func (r *ProjectsLogsService) Delete(projectsId string, logsId string) *ProjectsLogsDeleteCall {
	c := &ProjectsLogsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectsId = projectsId
	c.logsId = logsId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogsDeleteCall) Fields(s ...googleapi.Field) *ProjectsLogsDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsLogsDeleteCall) Do() (*Empty, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectsId}/logs/{logsId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectsId": c.projectsId,
		"logsId":     c.logsId,
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
	//   "description": "Deletes a log and all its log entries.\nThe log will reappear if it receives new entries.\n",
	//   "flatPath": "v1beta3/projects/{projectsId}/logs/{logsId}",
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
	//       "description": "Part of `logName`. The resource name of the log to be deleted.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/logs/{logsId}",
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
	s          *Service
	projectsId string
	opt_       map[string]interface{}
}

// List: Lists the logs in the project.
// Only logs that have entries are
// listed.
//
func (r *ProjectsLogsService) List(projectsId string) *ProjectsLogsListCall {
	c := &ProjectsLogsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectsId = projectsId
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of results to return.
func (c *ProjectsLogsListCall) PageSize(pageSize int64) *ProjectsLogsListCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken": An opaque token,
// returned as `nextPageToken` by a prior `ListLogs`
// operation.  If
// `pageToken` is supplied, then the other fields of this
// request are
// ignored, and instead the previous `ListLogs` operation is
// continued.
func (c *ProjectsLogsListCall) PageToken(pageToken string) *ProjectsLogsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// ServiceIndexPrefix sets the optional parameter "serviceIndexPrefix":
// The purpose of this field is to restrict the listed logs to
// those
// with entries of a certain kind.
// If `serviceName` is the name of
// a log service,
// then this field may contain values for the log
// service's indexes.
// Only logs that have entries whose indexes include
// the values are listed.
// The format for this field is
// `"/val1/val2.../valN"`, where `val1` is a
// value for the first index,
// `val2` for the second index, etc.
// An empty value (a single slash) for
// an index matches all values,
// and you can omit values for later
// indexes entirely.
func (c *ProjectsLogsListCall) ServiceIndexPrefix(serviceIndexPrefix string) *ProjectsLogsListCall {
	c.opt_["serviceIndexPrefix"] = serviceIndexPrefix
	return c
}

// ServiceName sets the optional parameter "serviceName": If not empty,
// this field must be a log service name such
// as
// `"compute.googleapis.com"`. Only logs associated with that
// that
// log service are listed.
func (c *ProjectsLogsListCall) ServiceName(serviceName string) *ProjectsLogsListCall {
	c.opt_["serviceName"] = serviceName
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogsListCall) Fields(s ...googleapi.Field) *ProjectsLogsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsLogsListCall) Do() (*ListLogsResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["pageSize"]; ok {
		params.Set("pageSize", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["serviceIndexPrefix"]; ok {
		params.Set("serviceIndexPrefix", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["serviceName"]; ok {
		params.Set("serviceName", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectsId}/logs")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectsId": c.projectsId,
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
	var ret *ListLogsResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists the logs in the project.\nOnly logs that have entries are listed.\n",
	//   "flatPath": "v1beta3/projects/{projectsId}/logs",
	//   "httpMethod": "GET",
	//   "id": "logging.projects.logs.list",
	//   "parameterOrder": [
	//     "projectsId"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "The maximum number of results to return.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "An opaque token, returned as `nextPageToken` by a prior `ListLogs`\noperation.  If `pageToken` is supplied, then the other fields of this\nrequest are ignored, and instead the previous `ListLogs` operation is\ncontinued.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `projectName`. The resource name of the project whose logs are requested.\nIf both `serviceName` and `serviceIndexPrefix` are empty, then\nall logs with entries in this project are listed.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "serviceIndexPrefix": {
	//       "description": "The purpose of this field is to restrict the listed logs to those\nwith entries of a certain kind.\nIf `serviceName` is the name of a log service,\nthen this field may contain values for the log service's indexes.\nOnly logs that have entries whose indexes include the values are listed.\nThe format for this field is `\"/val1/val2.../valN\"`, where `val1` is a\nvalue for the first index, `val2` for the second index, etc.\nAn empty value (a single slash) for an index matches all values,\nand you can omit values for later indexes entirely.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "serviceName": {
	//       "description": "If not empty, this field must be a log service name such as\n`\"compute.googleapis.com\"`. Only logs associated with that\nthat log service are listed.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/logs",
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

// method id "logging.projects.logs.entries.write":

type ProjectsLogsEntriesWriteCall struct {
	s                      *Service
	projectsId             string
	logsId                 string
	writelogentriesrequest *WriteLogEntriesRequest
	opt_                   map[string]interface{}
}

// Write: Writes log entries to Cloud Logging. Each entry consists of
// a
// `LogEntry` object.  You must fill in all the fields of the
// object,
// including one of the payload fields.  You may supply a
// map,
// `commonLabels`, that holds default (key, value) data for
// the
// `entries[].metadata.labels` map in each entry, saving you
// the
// trouble of creating identical copies for each entry.
//
//
func (r *ProjectsLogsEntriesService) Write(projectsId string, logsId string, writelogentriesrequest *WriteLogEntriesRequest) *ProjectsLogsEntriesWriteCall {
	c := &ProjectsLogsEntriesWriteCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectsId = projectsId
	c.logsId = logsId
	c.writelogentriesrequest = writelogentriesrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogsEntriesWriteCall) Fields(s ...googleapi.Field) *ProjectsLogsEntriesWriteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsLogsEntriesWriteCall) Do() (*WriteLogEntriesResponse, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.writelogentriesrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectsId}/logs/{logsId}/entries:write")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectsId": c.projectsId,
		"logsId":     c.logsId,
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
	var ret *WriteLogEntriesResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Writes log entries to Cloud Logging. Each entry consists of a\n`LogEntry` object.  You must fill in all the fields of the\nobject, including one of the payload fields.  You may supply a\nmap, `commonLabels`, that holds default (key, value) data for the\n`entries[].metadata.labels` map in each entry, saving you the\ntrouble of creating identical copies for each entry.\n\n",
	//   "flatPath": "v1beta3/projects/{projectsId}/logs/{logsId}/entries:write",
	//   "httpMethod": "POST",
	//   "id": "logging.projects.logs.entries.write",
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
	//       "description": "Part of `logName`. The resource name of the log that will receive the log entries.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/logs/{logsId}/entries:write",
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

// method id "logging.projects.logs.sinks.create":

type ProjectsLogsSinksCreateCall struct {
	s          *Service
	projectsId string
	logsId     string
	logsink    *LogSink
	opt_       map[string]interface{}
}

// Create: Creates a log sink.
// All log entries for a specified log are
// written to the destination.
func (r *ProjectsLogsSinksService) Create(projectsId string, logsId string, logsink *LogSink) *ProjectsLogsSinksCreateCall {
	c := &ProjectsLogsSinksCreateCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectsId = projectsId
	c.logsId = logsId
	c.logsink = logsink
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogsSinksCreateCall) Fields(s ...googleapi.Field) *ProjectsLogsSinksCreateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsLogsSinksCreateCall) Do() (*LogSink, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.logsink)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectsId}/logs/{logsId}/sinks")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectsId": c.projectsId,
		"logsId":     c.logsId,
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
	var ret *LogSink
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a log sink.\nAll log entries for a specified log are written to the destination.",
	//   "flatPath": "v1beta3/projects/{projectsId}/logs/{logsId}/sinks",
	//   "httpMethod": "POST",
	//   "id": "logging.projects.logs.sinks.create",
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
	//       "description": "Part of `logName`. The resource name of the log to which to the sink is bound.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/logs/{logsId}/sinks",
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

// method id "logging.projects.logs.sinks.delete":

type ProjectsLogsSinksDeleteCall struct {
	s          *Service
	projectsId string
	logsId     string
	sinksId    string
	opt_       map[string]interface{}
}

// Delete: Deletes a log sink.
// After deletion, no new log entries are
// written to the destination.
func (r *ProjectsLogsSinksService) Delete(projectsId string, logsId string, sinksId string) *ProjectsLogsSinksDeleteCall {
	c := &ProjectsLogsSinksDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectsId = projectsId
	c.logsId = logsId
	c.sinksId = sinksId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogsSinksDeleteCall) Fields(s ...googleapi.Field) *ProjectsLogsSinksDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsLogsSinksDeleteCall) Do() (*Empty, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectsId}/logs/{logsId}/sinks/{sinksId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectsId": c.projectsId,
		"logsId":     c.logsId,
		"sinksId":    c.sinksId,
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
	//   "description": "Deletes a log sink.\nAfter deletion, no new log entries are written to the destination.",
	//   "flatPath": "v1beta3/projects/{projectsId}/logs/{logsId}/sinks/{sinksId}",
	//   "httpMethod": "DELETE",
	//   "id": "logging.projects.logs.sinks.delete",
	//   "parameterOrder": [
	//     "projectsId",
	//     "logsId",
	//     "sinksId"
	//   ],
	//   "parameters": {
	//     "logsId": {
	//       "description": "Part of `sinkName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `sinkName`. The resource name of the log sink to delete.",
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
	//   "path": "v1beta3/projects/{projectsId}/logs/{logsId}/sinks/{sinksId}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/logging.admin"
	//   ]
	// }

}

// method id "logging.projects.logs.sinks.get":

type ProjectsLogsSinksGetCall struct {
	s          *Service
	projectsId string
	logsId     string
	sinksId    string
	opt_       map[string]interface{}
}

// Get: Gets a log sink.
func (r *ProjectsLogsSinksService) Get(projectsId string, logsId string, sinksId string) *ProjectsLogsSinksGetCall {
	c := &ProjectsLogsSinksGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectsId = projectsId
	c.logsId = logsId
	c.sinksId = sinksId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogsSinksGetCall) Fields(s ...googleapi.Field) *ProjectsLogsSinksGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsLogsSinksGetCall) Do() (*LogSink, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectsId}/logs/{logsId}/sinks/{sinksId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectsId": c.projectsId,
		"logsId":     c.logsId,
		"sinksId":    c.sinksId,
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
	var ret *LogSink
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets a log sink.",
	//   "flatPath": "v1beta3/projects/{projectsId}/logs/{logsId}/sinks/{sinksId}",
	//   "httpMethod": "GET",
	//   "id": "logging.projects.logs.sinks.get",
	//   "parameterOrder": [
	//     "projectsId",
	//     "logsId",
	//     "sinksId"
	//   ],
	//   "parameters": {
	//     "logsId": {
	//       "description": "Part of `sinkName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `sinkName`. The resource name of the log sink to return.",
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
	//   "path": "v1beta3/projects/{projectsId}/logs/{logsId}/sinks/{sinksId}",
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

// method id "logging.projects.logs.sinks.list":

type ProjectsLogsSinksListCall struct {
	s          *Service
	projectsId string
	logsId     string
	opt_       map[string]interface{}
}

// List: Lists log sinks associated with a log.
func (r *ProjectsLogsSinksService) List(projectsId string, logsId string) *ProjectsLogsSinksListCall {
	c := &ProjectsLogsSinksListCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectsId = projectsId
	c.logsId = logsId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogsSinksListCall) Fields(s ...googleapi.Field) *ProjectsLogsSinksListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsLogsSinksListCall) Do() (*ListLogSinksResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectsId}/logs/{logsId}/sinks")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectsId": c.projectsId,
		"logsId":     c.logsId,
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
	var ret *ListLogSinksResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists log sinks associated with a log.",
	//   "flatPath": "v1beta3/projects/{projectsId}/logs/{logsId}/sinks",
	//   "httpMethod": "GET",
	//   "id": "logging.projects.logs.sinks.list",
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
	//       "description": "Part of `logName`. The log whose sinks are wanted.\nFor example, `\"compute.google.com/syslog\"`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/logs/{logsId}/sinks",
	//   "response": {
	//     "$ref": "ListLogSinksResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only",
	//     "https://www.googleapis.com/auth/logging.admin",
	//     "https://www.googleapis.com/auth/logging.read"
	//   ]
	// }

}

// method id "logging.projects.logs.sinks.update":

type ProjectsLogsSinksUpdateCall struct {
	s          *Service
	projectsId string
	logsId     string
	sinksId    string
	logsink    *LogSink
	opt_       map[string]interface{}
}

// Update: Updates a log sink.
// If the sink does not exist, it is
// created.
func (r *ProjectsLogsSinksService) Update(projectsId string, logsId string, sinksId string, logsink *LogSink) *ProjectsLogsSinksUpdateCall {
	c := &ProjectsLogsSinksUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectsId = projectsId
	c.logsId = logsId
	c.sinksId = sinksId
	c.logsink = logsink
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogsSinksUpdateCall) Fields(s ...googleapi.Field) *ProjectsLogsSinksUpdateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsLogsSinksUpdateCall) Do() (*LogSink, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.logsink)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectsId}/logs/{logsId}/sinks/{sinksId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectsId": c.projectsId,
		"logsId":     c.logsId,
		"sinksId":    c.sinksId,
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
	var ret *LogSink
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates a log sink.\nIf the sink does not exist, it is created.",
	//   "flatPath": "v1beta3/projects/{projectsId}/logs/{logsId}/sinks/{sinksId}",
	//   "httpMethod": "PUT",
	//   "id": "logging.projects.logs.sinks.update",
	//   "parameterOrder": [
	//     "projectsId",
	//     "logsId",
	//     "sinksId"
	//   ],
	//   "parameters": {
	//     "logsId": {
	//       "description": "Part of `sinkName`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `sinkName`. The resource name of the sink to update.",
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
	//   "path": "v1beta3/projects/{projectsId}/logs/{logsId}/sinks/{sinksId}",
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

// method id "logging.projects.metrics.create":

type ProjectsMetricsCreateCall struct {
	s          *Service
	projectsId string
	logmetric  *LogMetric
	opt_       map[string]interface{}
}

// Create: Creates a logs-based metric.
func (r *ProjectsMetricsService) Create(projectsId string, logmetric *LogMetric) *ProjectsMetricsCreateCall {
	c := &ProjectsMetricsCreateCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectsId = projectsId
	c.logmetric = logmetric
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsMetricsCreateCall) Fields(s ...googleapi.Field) *ProjectsMetricsCreateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsMetricsCreateCall) Do() (*LogMetric, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.logmetric)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectsId}/metrics")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectsId": c.projectsId,
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
	var ret *LogMetric
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a logs-based metric.",
	//   "flatPath": "v1beta3/projects/{projectsId}/metrics",
	//   "httpMethod": "POST",
	//   "id": "logging.projects.metrics.create",
	//   "parameterOrder": [
	//     "projectsId"
	//   ],
	//   "parameters": {
	//     "projectsId": {
	//       "description": "Part of `projectName`. The resource name of the project in which to create the metric.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/metrics",
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
	opt_       map[string]interface{}
}

// Delete: Deletes a logs-based metric.
func (r *ProjectsMetricsService) Delete(projectsId string, metricsId string) *ProjectsMetricsDeleteCall {
	c := &ProjectsMetricsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectsId = projectsId
	c.metricsId = metricsId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsMetricsDeleteCall) Fields(s ...googleapi.Field) *ProjectsMetricsDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsMetricsDeleteCall) Do() (*Empty, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectsId}/metrics/{metricsId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectsId": c.projectsId,
		"metricsId":  c.metricsId,
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
	//   "description": "Deletes a logs-based metric.",
	//   "flatPath": "v1beta3/projects/{projectsId}/metrics/{metricsId}",
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
	//       "description": "Part of `metricName`. The resource name of the metric to delete.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/metrics/{metricsId}",
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
	s          *Service
	projectsId string
	metricsId  string
	opt_       map[string]interface{}
}

// Get: Gets a logs-based metric.
func (r *ProjectsMetricsService) Get(projectsId string, metricsId string) *ProjectsMetricsGetCall {
	c := &ProjectsMetricsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectsId = projectsId
	c.metricsId = metricsId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsMetricsGetCall) Fields(s ...googleapi.Field) *ProjectsMetricsGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsMetricsGetCall) Do() (*LogMetric, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectsId}/metrics/{metricsId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectsId": c.projectsId,
		"metricsId":  c.metricsId,
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
	var ret *LogMetric
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets a logs-based metric.",
	//   "flatPath": "v1beta3/projects/{projectsId}/metrics/{metricsId}",
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
	//       "description": "Part of `metricName`. The resource name of the desired metric.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/metrics/{metricsId}",
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
	s          *Service
	projectsId string
	opt_       map[string]interface{}
}

// List: Lists the logs-based metrics associated with a project.
func (r *ProjectsMetricsService) List(projectsId string) *ProjectsMetricsListCall {
	c := &ProjectsMetricsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectsId = projectsId
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of `LogMetric` objects to return in one operation.
func (c *ProjectsMetricsListCall) PageSize(pageSize int64) *ProjectsMetricsListCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken": An opaque token,
// returned as `nextPageToken` by a prior `ListLogMetrics`
// operation.
// If `pageToken` is supplied, then the other fields of this
// request are
// ignored, and instead the previous `ListLogMetrics` operation
// is
// continued.
func (c *ProjectsMetricsListCall) PageToken(pageToken string) *ProjectsMetricsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsMetricsListCall) Fields(s ...googleapi.Field) *ProjectsMetricsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsMetricsListCall) Do() (*ListLogMetricsResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["pageSize"]; ok {
		params.Set("pageSize", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectsId}/metrics")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectsId": c.projectsId,
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
	var ret *ListLogMetricsResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists the logs-based metrics associated with a project.",
	//   "flatPath": "v1beta3/projects/{projectsId}/metrics",
	//   "httpMethod": "GET",
	//   "id": "logging.projects.metrics.list",
	//   "parameterOrder": [
	//     "projectsId"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "The maximum number of `LogMetric` objects to return in one operation.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "An opaque token, returned as `nextPageToken` by a prior `ListLogMetrics`\noperation.  If `pageToken` is supplied, then the other fields of this\nrequest are ignored, and instead the previous `ListLogMetrics` operation is\ncontinued.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `projectName`. The resource name for the project whose metrics are wanted.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/metrics",
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

// method id "logging.projects.metrics.update":

type ProjectsMetricsUpdateCall struct {
	s          *Service
	projectsId string
	metricsId  string
	logmetric  *LogMetric
	opt_       map[string]interface{}
}

// Update: Creates or updates a logs-based metric.
func (r *ProjectsMetricsService) Update(projectsId string, metricsId string, logmetric *LogMetric) *ProjectsMetricsUpdateCall {
	c := &ProjectsMetricsUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectsId = projectsId
	c.metricsId = metricsId
	c.logmetric = logmetric
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsMetricsUpdateCall) Fields(s ...googleapi.Field) *ProjectsMetricsUpdateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsMetricsUpdateCall) Do() (*LogMetric, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.logmetric)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectsId}/metrics/{metricsId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectsId": c.projectsId,
		"metricsId":  c.metricsId,
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
	var ret *LogMetric
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates or updates a logs-based metric.",
	//   "flatPath": "v1beta3/projects/{projectsId}/metrics/{metricsId}",
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
	//       "description": "Part of `metricName`. The resource name of the metric to update.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/metrics/{metricsId}",
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

// method id "logging.projects.sinks.create":

type ProjectsSinksCreateCall struct {
	s          *Service
	projectsId string
	logsink    *LogSink
	opt_       map[string]interface{}
}

// Create: Creates a project sink.
// A logs filter determines which log
// entries are written to the destination.
func (r *ProjectsSinksService) Create(projectsId string, logsink *LogSink) *ProjectsSinksCreateCall {
	c := &ProjectsSinksCreateCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectsId = projectsId
	c.logsink = logsink
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsSinksCreateCall) Fields(s ...googleapi.Field) *ProjectsSinksCreateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsSinksCreateCall) Do() (*LogSink, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.logsink)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectsId}/sinks")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectsId": c.projectsId,
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
	var ret *LogSink
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a project sink.\nA logs filter determines which log entries are written to the destination.",
	//   "flatPath": "v1beta3/projects/{projectsId}/sinks",
	//   "httpMethod": "POST",
	//   "id": "logging.projects.sinks.create",
	//   "parameterOrder": [
	//     "projectsId"
	//   ],
	//   "parameters": {
	//     "projectsId": {
	//       "description": "Part of `projectName`. The resource name of the project to which the sink is bound.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/sinks",
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
	opt_       map[string]interface{}
}

// Delete: Deletes a project sink.
// After deletion, no new log entries
// are written to the destination.
func (r *ProjectsSinksService) Delete(projectsId string, sinksId string) *ProjectsSinksDeleteCall {
	c := &ProjectsSinksDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectsId = projectsId
	c.sinksId = sinksId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsSinksDeleteCall) Fields(s ...googleapi.Field) *ProjectsSinksDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsSinksDeleteCall) Do() (*Empty, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectsId}/sinks/{sinksId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectsId": c.projectsId,
		"sinksId":    c.sinksId,
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
	//   "description": "Deletes a project sink.\nAfter deletion, no new log entries are written to the destination.",
	//   "flatPath": "v1beta3/projects/{projectsId}/sinks/{sinksId}",
	//   "httpMethod": "DELETE",
	//   "id": "logging.projects.sinks.delete",
	//   "parameterOrder": [
	//     "projectsId",
	//     "sinksId"
	//   ],
	//   "parameters": {
	//     "projectsId": {
	//       "description": "Part of `sinkName`. The resource name of the project sink to delete.",
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
	//   "path": "v1beta3/projects/{projectsId}/sinks/{sinksId}",
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
	s          *Service
	projectsId string
	sinksId    string
	opt_       map[string]interface{}
}

// Get: Gets a project sink.
func (r *ProjectsSinksService) Get(projectsId string, sinksId string) *ProjectsSinksGetCall {
	c := &ProjectsSinksGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectsId = projectsId
	c.sinksId = sinksId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsSinksGetCall) Fields(s ...googleapi.Field) *ProjectsSinksGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsSinksGetCall) Do() (*LogSink, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectsId}/sinks/{sinksId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectsId": c.projectsId,
		"sinksId":    c.sinksId,
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
	var ret *LogSink
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets a project sink.",
	//   "flatPath": "v1beta3/projects/{projectsId}/sinks/{sinksId}",
	//   "httpMethod": "GET",
	//   "id": "logging.projects.sinks.get",
	//   "parameterOrder": [
	//     "projectsId",
	//     "sinksId"
	//   ],
	//   "parameters": {
	//     "projectsId": {
	//       "description": "Part of `sinkName`. The resource name of the project sink to return.",
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
	//   "path": "v1beta3/projects/{projectsId}/sinks/{sinksId}",
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
	s          *Service
	projectsId string
	opt_       map[string]interface{}
}

// List: Lists project sinks associated with a project.
func (r *ProjectsSinksService) List(projectsId string) *ProjectsSinksListCall {
	c := &ProjectsSinksListCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectsId = projectsId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsSinksListCall) Fields(s ...googleapi.Field) *ProjectsSinksListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsSinksListCall) Do() (*ListSinksResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectsId}/sinks")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectsId": c.projectsId,
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
	var ret *ListSinksResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists project sinks associated with a project.",
	//   "flatPath": "v1beta3/projects/{projectsId}/sinks",
	//   "httpMethod": "GET",
	//   "id": "logging.projects.sinks.list",
	//   "parameterOrder": [
	//     "projectsId"
	//   ],
	//   "parameters": {
	//     "projectsId": {
	//       "description": "Part of `projectName`. The project whose sinks are wanted.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta3/projects/{projectsId}/sinks",
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

// method id "logging.projects.sinks.update":

type ProjectsSinksUpdateCall struct {
	s          *Service
	projectsId string
	sinksId    string
	logsink    *LogSink
	opt_       map[string]interface{}
}

// Update: Updates a project sink.
// If the sink does not exist, it is
// created.
// The destination, filter, or both may be updated.
func (r *ProjectsSinksService) Update(projectsId string, sinksId string, logsink *LogSink) *ProjectsSinksUpdateCall {
	c := &ProjectsSinksUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectsId = projectsId
	c.sinksId = sinksId
	c.logsink = logsink
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsSinksUpdateCall) Fields(s ...googleapi.Field) *ProjectsSinksUpdateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsSinksUpdateCall) Do() (*LogSink, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.logsink)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta3/projects/{projectsId}/sinks/{sinksId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectsId": c.projectsId,
		"sinksId":    c.sinksId,
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
	var ret *LogSink
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates a project sink.\nIf the sink does not exist, it is created.\nThe destination, filter, or both may be updated.",
	//   "flatPath": "v1beta3/projects/{projectsId}/sinks/{sinksId}",
	//   "httpMethod": "PUT",
	//   "id": "logging.projects.sinks.update",
	//   "parameterOrder": [
	//     "projectsId",
	//     "sinksId"
	//   ],
	//   "parameters": {
	//     "projectsId": {
	//       "description": "Part of `sinkName`. The resource name of the project sink to update.",
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
	//   "path": "v1beta3/projects/{projectsId}/sinks/{sinksId}",
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
