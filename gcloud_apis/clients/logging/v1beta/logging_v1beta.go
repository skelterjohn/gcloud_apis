// Package logging provides access to the Logging API.
//
// See http://developers.google.com
//
// Usage example:
//
//   import "google.golang.org/api/logging/v1beta"
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

const apiId = "logging:v1beta"
const apiName = "logging"
const apiVersion = "v1beta"
const basePath = "https://www.googleapis.com/logging/v1beta/"

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
	rs.Logs = NewProjectsLogsService(s)
	return rs
}

type ProjectsService struct {
	s *Service

	Logs *ProjectsLogsService
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

type Any struct {
	// TypeUrl: A URL that resolves to a google.protobuf.Type value.
	TypeUrl string `json:"typeUrl,omitempty"`

	// Value: Serialized data of the specified type.
	//
	// Note: the field is
	// optional since if we are converting between formats and the type
	// could not be resolved, we store the original data pre-conversion in
	// other fields (TBD).
	Value string `json:"value,omitempty"`
}

type Label struct {
	// Key: The key of a label is a syntactically valid URL (as per RFC
	// 1738) with the "scheme" and initial slashes omitted and with the
	// additional restrictions noted below. Each key should be globally
	// unique. The "host" portion is called the "namespace" and is not
	// necessarily resolvable to a network endpoint. Instead, the namespace
	// indicates what system or entity defines the semantics of the label.
	// Namespaces do not restrict the set of objects to which a label may be
	// associated.
	//
	// Keys are defined by the following grammar:
	//
	// key =
	// hostname "/" kpath kpath = ksegment *[ "/" ksegment ] ksegment =
	// alphadigit | *[ alphadigit | "-" | "_" | "." ]
	//
	// where "hostname" and
	// "alphadigit" are defined as in RFC 1738.
	//
	// Example key:
	// spanner.google.com/universe
	Key string `json:"key,omitempty"`

	// NumValue: An integer value.
	NumValue int64 `json:"numValue,omitempty,string"`

	// StrValue: A string value.
	StrValue string `json:"strValue,omitempty"`
}

type ListLogSinksResponse struct {
	// Sinks: These may be partial sinks where only name is populated.
	Sinks []*LogSink `json:"sinks,omitempty"`
}

type ListLogsResponse struct {
	// Logs: A list of log resources.
	Logs []*Log `json:"logs,omitempty"`

	// NextPageToken: Pagination: If there are more results, retrieve them
	// by invoking ListLogs again with the same arguments and this
	// next_page_token.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type Log struct {
	// DisplayName: Name to be used when displaying the log to the user
	// (e.g., in a UI)
	DisplayName string `json:"displayName,omitempty"`

	// Name: REQUIRED: log name
	Name string `json:"name,omitempty"`

	// PayloadType: Type URL describing the expected payload type for the
	// log.
	PayloadType string `json:"payloadType,omitempty"`
}

type LogEntry struct {
	// InsertId: Unique ID used to deduplicate log entries. If provided, the
	// logging service will attempt to reject multiple log entries on the
	// same log with the same insert_id that are sent within the previous N
	// minutes. Deduplication may occur asynchronously, so the client may
	// not receive an error if the entry is recognized as a duplicate.
	InsertId string `json:"insertId,omitempty"`

	// Log: OUTPUT-ONLY: The log resource that this entry belongs to.
	Log string `json:"log,omitempty"`

	// Metadata: Metadata for the log entry.
	Metadata *LogEntryMetadata `json:"metadata,omitempty"`

	// ProtoPayload: Contains a structured (protocol buffer) log entry. If
	// the data type in proto_payload is known by the logging system, JSON
	// clients will receive it in JSON; otherwise, the client will receive
	// it as a serialized proto and must decode it. See google.protobuf.Any
	// for more details.
	ProtoPayload *Any `json:"protoPayload,omitempty"`

	// TextPayload: Contains a text representation of the log entry.
	TextPayload string `json:"textPayload,omitempty"`
}

type LogEntryMetadata struct {
	// Labels: Labels are not yet fully supported (as of May 2014). Callers
	// are expected to populate one of the following sets of labels
	// describing the source of the log entry, so they can be transformed to
	// a cloud.CloudTask proto:
	//
	// App Engine: service:
	// "appengine.googleapis.com" labels: appengine.googleapis.com/module_id
	// appengine.googleapis.com/version_id and one of:
	// appengine.googleapis.com/replica_index
	// appengine.googleapis.com/clone_id or the Compute Engine labels
	// (resource_type/resource_id) below
	//
	// Compute Engine: service:
	// "compute.googleapis.com" labels: compute.googleapis.com/resource_type
	// = "instance" compute.googleapis.com/resource_id
	//
	// Other labels may be
	// supplied, but are currently ignored and not stored with the log
	// entry.
	Labels []*Label `json:"labels,omitempty"`

	// ProjectId: If the log entry is from a Google Cloud Platform source,
	// this must be the project ID of the service that generated the entry.
	// Otherwise, the caller may populate this with whatever project name or
	// ID is meaningful, or leave it unset.
	ProjectId string `json:"projectId,omitempty"`

	// Region: If the log entry is from a Google Cloud Platform source, this
	// must be the region of the source (e.g., us-central1) if it has one,
	// or unset if "region" isn't meaningful for the service.
	//
	// For
	// non-Google Cloud Platform sources, the caller may populate this with
	// whatever location name is meaningful, or leave it unset.
	Region string `json:"region,omitempty"`

	// ServiceName: If the log entry is from a Google Cloud Platform source,
	// this must be the official API name of the service (e.g.,
	// compute.googleapis.com). Otherwise, the caller may populate this with
	// whatever service name is meaningful, or leave it unset.
	ServiceName string `json:"serviceName,omitempty"`

	// Severity: The severity of the log entry.
	Severity string `json:"severity,omitempty"`

	// TimeNanos: REQUIRED: The time the event described by the log entry
	// occurred, in nanoseconds since the Unix epoch.
	TimeNanos int64 `json:"timeNanos,omitempty,string"`

	// Timestamp: REQUIRED: The time the event described by the log entry
	// occurred.
	Timestamp *Timestamp `json:"timestamp,omitempty"`

	// UserId: If the log entry applies to an action taken by an
	// authenticated user, this field must contain the user identifier (a
	// fully qualified email address) of the user that requested or
	// performed the action. May be empty if the event described by the log
	// entry doesn't have an associated user.
	UserId string `json:"userId,omitempty"`

	// Zone: If the log entry is from a Google Cloud Platform source, this
	// must be the zone of the source (e.g., us-central1-a) if it has one,
	// or unset if "zone" isn't meaningful for the service.
	//
	// For non-Google
	// Cloud Platform sources, the caller may populate this with whatever
	// location name is meaningful, or leave it unset.
	Zone string `json:"zone,omitempty"`
}

type LogError struct {
	// Resource: The resource associated with the error. It may be different
	// from the sink destination. E.g. the sink may point to a BigQuery
	// dataset, but the error may refer to a table resource inside the
	// dataset.
	Resource string `json:"resource,omitempty"`

	// Status: The description of the last error observed.
	Status *Status `json:"status,omitempty"`

	// TimeNanos: The last time the error was observed, in nanoseconds since
	// the Unix epoch.
	TimeNanos int64 `json:"timeNanos,omitempty,string"`
}

type LogSink struct {
	// Destination: The resource to send log entries to. The supported sink
	// resource types are:
	//
	// Google Cloud Storage:
	// storage.googleapis.com/{bucket} bucket.storage.googleapis.com/ Google
	// BigQuery:
	// bigquery.googleapis.com/projects/{projectId}/datasets/{datasetId}
	//
	// Cur
	// rently the logging service supports at most one sink of each type per
	// log resource.
	Destination string `json:"destination,omitempty"`

	// Errors: All active errors found for this sink. [output only]
	Errors []*LogError `json:"errors,omitempty"`

	// Name: The name of this sink. This is a client-assigned identifier for
	// the resource.
	Name string `json:"name,omitempty"`
}

type Status struct {
	// Code: This is the status code. For Google APIs, this code must be the
	// Canonical Error Code.
	Code int64 `json:"code,omitempty"`

	// Details: This is a list of messages that carry the error details.
	// There will be a common set of message types for APIs to use.
	Details []*Any `json:"details,omitempty"`

	// Message: This is a human readable error string.
	Message string `json:"message,omitempty"`
}

type Timestamp struct {
	// Nanos: Positive fractions of a second at nanosecond resolution.
	// Negative second values with fractions may still have positive nanos
	// values that count forward in time. Must be from 0 to 999,999,999
	// inclusive.
	Nanos int64 `json:"nanos,omitempty"`

	// Seconds: Positive or negative seconds of UTC time since Unix epoch
	// 1970-01-01T00:00:00Z. Must be from from 0001-01-01T00:00:00Z to
	// 9999-12-31T23:59:59Z inclusive.
	Seconds int64 `json:"seconds,omitempty,string"`
}

type WriteLogEntriesRequest struct {
	// CommonLabels: Labels that apply to all entries in this request. If a
	// conflicting label key is present in the per-entry
	// LogEntryMetadata.label list, it overrides the value specified
	// here.
	//
	// See the documentation for LogEntryMetadata.labels for
	// additional notes.
	CommonLabels []*Label `json:"commonLabels,omitempty"`

	// Entries: Log entries to insert.
	Entries []*LogEntry `json:"entries,omitempty"`
}

// method id "logging.projects.logs.list":

type ProjectsLogsListCall struct {
	s       *Service
	project string
	opt_    map[string]interface{}
}

// List: Lists log resources belonging to the specified
// project.
//
// Requires https://www.googleapis.com/auth/logging.read
// scope.
func (r *ProjectsLogsService) List(project string) *ProjectsLogsListCall {
	c := &ProjectsLogsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	return c
}

// PageSize sets the optional parameter "pageSize": Pagination: maximum
// number of results to return per page
func (c *ProjectsLogsListCall) PageSize(pageSize int64) *ProjectsLogsListCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken": Pagination: a
// next_page_token previously returned from ListLogs
func (c *ProjectsLogsListCall) PageToken(pageToken string) *ProjectsLogsListCall {
	c.opt_["pageToken"] = pageToken
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
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{+project}/logs")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project": c.project,
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
	//   "description": "Lists log resources belonging to the specified project.\n\nRequires https://www.googleapis.com/auth/logging.read scope.",
	//   "httpMethod": "GET",
	//   "id": "logging.projects.logs.list",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "Pagination: maximum number of results to return per page",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Pagination: a next_page_token previously returned from ListLogs",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "The project ID for which to list the log resources.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{+project}/logs",
	//   "response": {
	//     "$ref": "ListLogsResponse"
	//   }
	// }

}

// method id "logging.projects.logs.entries.write":

type ProjectsLogsEntriesWriteCall struct {
	s                      *Service
	log                    string
	writelogentriesrequest *WriteLogEntriesRequest
	opt_                   map[string]interface{}
}

// Write: Creates several log entries in a log.
//
// Requires
// https://www.googleapis.com/auth/logging.write scope.
func (r *ProjectsLogsEntriesService) Write(log string, writelogentriesrequest *WriteLogEntriesRequest) *ProjectsLogsEntriesWriteCall {
	c := &ProjectsLogsEntriesWriteCall{s: r.s, opt_: make(map[string]interface{})}
	c.log = log
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

func (c *ProjectsLogsEntriesWriteCall) Do() error {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.writelogentriesrequest)
	if err != nil {
		return err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{+log}/entries:write")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"log": c.log,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return err
	}
	return nil
	// {
	//   "description": "Creates several log entries in a log.\n\nRequires https://www.googleapis.com/auth/logging.write scope.",
	//   "httpMethod": "POST",
	//   "id": "logging.projects.logs.entries.write",
	//   "parameterOrder": [
	//     "log"
	//   ],
	//   "parameters": {
	//     "log": {
	//       "description": "The log resource into which to insert the log entries.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{+log}/entries:write",
	//   "request": {
	//     "$ref": "WriteLogEntriesRequest"
	//   }
	// }

}

// method id "logging.projects.logs.sinks.delete":

type ProjectsLogsSinksDeleteCall struct {
	s    *Service
	sink string
	opt_ map[string]interface{}
}

// Delete: Deletes the specified log sink. Requires
// https://www.googleapis.com/auth/logging.admin scope.
func (r *ProjectsLogsSinksService) Delete(sink string) *ProjectsLogsSinksDeleteCall {
	c := &ProjectsLogsSinksDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.sink = sink
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogsSinksDeleteCall) Fields(s ...googleapi.Field) *ProjectsLogsSinksDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsLogsSinksDeleteCall) Do() error {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{+sink}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"sink": c.sink,
	})
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return err
	}
	return nil
	// {
	//   "description": "Deletes the specified log sink. Requires https://www.googleapis.com/auth/logging.admin scope.",
	//   "httpMethod": "DELETE",
	//   "id": "logging.projects.logs.sinks.delete",
	//   "parameterOrder": [
	//     "sink"
	//   ],
	//   "parameters": {
	//     "sink": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{+sink}"
	// }

}

// method id "logging.projects.logs.sinks.get":

type ProjectsLogsSinksGetCall struct {
	s    *Service
	sink string
	opt_ map[string]interface{}
}

// Get: Get the specified log sink resource. Requires
// https://www.googleapis.com/auth/logging.read scope.
func (r *ProjectsLogsSinksService) Get(sink string) *ProjectsLogsSinksGetCall {
	c := &ProjectsLogsSinksGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.sink = sink
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "{+sink}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"sink": c.sink,
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
	//   "description": "Get the specified log sink resource. Requires https://www.googleapis.com/auth/logging.read scope.",
	//   "httpMethod": "GET",
	//   "id": "logging.projects.logs.sinks.get",
	//   "parameterOrder": [
	//     "sink"
	//   ],
	//   "parameters": {
	//     "sink": {
	//       "description": "Obsolete: field numbers 1-3",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{+sink}",
	//   "response": {
	//     "$ref": "LogSink"
	//   }
	// }

}

// method id "logging.projects.logs.sinks.list":

type ProjectsLogsSinksListCall struct {
	s    *Service
	log  string
	opt_ map[string]interface{}
}

// List: List log sinks associated with the specified log. Requires
// https://www.googleapis.com/auth/logging.read scope.
func (r *ProjectsLogsSinksService) List(log string) *ProjectsLogsSinksListCall {
	c := &ProjectsLogsSinksListCall{s: r.s, opt_: make(map[string]interface{})}
	c.log = log
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "{+log}/sinks")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"log": c.log,
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
	//   "description": "List log sinks associated with the specified log. Requires https://www.googleapis.com/auth/logging.read scope.",
	//   "httpMethod": "GET",
	//   "id": "logging.projects.logs.sinks.list",
	//   "parameterOrder": [
	//     "log"
	//   ],
	//   "parameters": {
	//     "log": {
	//       "description": "Obsolete: field numbers 1 and 2.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{+log}/sinks",
	//   "response": {
	//     "$ref": "ListLogSinksResponse"
	//   }
	// }

}

// method id "logging.projects.logs.sinks.patch":

type ProjectsLogsSinksPatchCall struct {
	s       *Service
	sink    string
	logsink *LogSink
	opt_    map[string]interface{}
}

// Patch: Create or update the specified log sink resource. Requires
// https://www.googleapis.com/auth/logging.admin scope. This method
// supports patch semantics.
func (r *ProjectsLogsSinksService) Patch(sink string, logsink *LogSink) *ProjectsLogsSinksPatchCall {
	c := &ProjectsLogsSinksPatchCall{s: r.s, opt_: make(map[string]interface{})}
	c.sink = sink
	c.logsink = logsink
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLogsSinksPatchCall) Fields(s ...googleapi.Field) *ProjectsLogsSinksPatchCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsLogsSinksPatchCall) Do() (*LogSink, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "{+sink}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"sink": c.sink,
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
	//   "description": "Create or update the specified log sink resource. Requires https://www.googleapis.com/auth/logging.admin scope. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "logging.projects.logs.sinks.patch",
	//   "parameterOrder": [
	//     "sink"
	//   ],
	//   "parameters": {
	//     "sink": {
	//       "description": "Obsolete: field numbers 1-3",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{+sink}",
	//   "request": {
	//     "$ref": "LogSink"
	//   },
	//   "response": {
	//     "$ref": "LogSink"
	//   }
	// }

}

// method id "logging.projects.logs.sinks.update":

type ProjectsLogsSinksUpdateCall struct {
	s       *Service
	sink    string
	logsink *LogSink
	opt_    map[string]interface{}
}

// Update: Create or update the specified log sink resource. Requires
// https://www.googleapis.com/auth/logging.admin scope.
func (r *ProjectsLogsSinksService) Update(sink string, logsink *LogSink) *ProjectsLogsSinksUpdateCall {
	c := &ProjectsLogsSinksUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.sink = sink
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "{+sink}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"sink": c.sink,
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
	//   "description": "Create or update the specified log sink resource. Requires https://www.googleapis.com/auth/logging.admin scope.",
	//   "httpMethod": "PUT",
	//   "id": "logging.projects.logs.sinks.update",
	//   "parameterOrder": [
	//     "sink"
	//   ],
	//   "parameters": {
	//     "sink": {
	//       "description": "Obsolete: field numbers 1-3",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{+sink}",
	//   "request": {
	//     "$ref": "LogSink"
	//   },
	//   "response": {
	//     "$ref": "LogSink"
	//   }
	// }

}
