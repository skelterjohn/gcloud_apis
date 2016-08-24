// Package cloudfunctions provides access to the Google Cloud Functions API.
//
// See https://cloud.google.com/functions
//
// Usage example:
//
//   import "github.com/skelterjohn/gcloud_apis/clients/cloudfunctions/v1beta1"
//   ...
//   cloudfunctionsService, err := cloudfunctions.New(oauthHttpClient)
package cloudfunctions

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

const apiId = "cloudfunctions:v1beta1"
const apiName = "cloudfunctions"
const apiVersion = "v1beta1"
const basePath = "https://cloudfunctions.googleapis.com/"

// OAuth2 scopes used by this API.
const (
	// View and manage your data across Google Cloud Platform services
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"
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
	rs.Locations = NewProjectsLocationsService(s)
	rs.Regions = NewProjectsRegionsService(s)
	return rs
}

type ProjectsService struct {
	s *Service

	Locations *ProjectsLocationsService

	Regions *ProjectsRegionsService
}

func NewProjectsLocationsService(s *Service) *ProjectsLocationsService {
	rs := &ProjectsLocationsService{s: s}
	return rs
}

type ProjectsLocationsService struct {
	s *Service
}

func NewProjectsRegionsService(s *Service) *ProjectsRegionsService {
	rs := &ProjectsRegionsService{s: s}
	rs.Functions = NewProjectsRegionsFunctionsService(s)
	return rs
}

type ProjectsRegionsService struct {
	s *Service

	Functions *ProjectsRegionsFunctionsService
}

func NewProjectsRegionsFunctionsService(s *Service) *ProjectsRegionsFunctionsService {
	rs := &ProjectsRegionsFunctionsService{s: s}
	return rs
}

type ProjectsRegionsFunctionsService struct {
	s *Service
}

// CallFunctionRequest: Request for the CallFunction method.
type CallFunctionRequest struct {
	// Data: Input to be passed to the function.
	Data string `json:"data,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Data") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *CallFunctionRequest) MarshalJSON() ([]byte, error) {
	type noMethod CallFunctionRequest
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// CallFunctionResponse: Response of CallFunction method.
type CallFunctionResponse struct {
	// Error: Either system or user-function generated error. Set if
	// execution
	// was not successful.
	Error string `json:"error,omitempty"`

	// ExecutionId: Execution id of function invocation.
	ExecutionId string `json:"executionId,omitempty"`

	// Result: Result populated for successful execution of synchronous
	// function. Will
	// not be populated if function does not return a result through
	// context.
	Result string `json:"result,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Error") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *CallFunctionResponse) MarshalJSON() ([]byte, error) {
	type noMethod CallFunctionResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// FunctionTrigger: Describes binding of computation to the event
// source.
type FunctionTrigger struct {
	// GsUri: Google Cloud Storage resource whose changes trigger the
	// events.
	// Currently, it must have the form gs://<bucket>/ (that is, it must
	// refer
	// to a bucket, rather than an object).
	GsUri string `json:"gsUri,omitempty"`

	// PubsubTopic: A pub/sub type of source.
	PubsubTopic string `json:"pubsubTopic,omitempty"`

	// WebTrigger: A web endpoint (e.g. HTTP) type of source that can be
	// trigger via URL.
	WebTrigger *WebTrigger `json:"webTrigger,omitempty"`

	// ForceSendFields is a list of field names (e.g. "GsUri") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *FunctionTrigger) MarshalJSON() ([]byte, error) {
	type noMethod FunctionTrigger
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// HostedFunction: Describes a cloud function that contains user
// computation executed in
// response to an event. It encapsulate function and triggers
// configurations.
type HostedFunction struct {
	// EntryPoint: The name of the function (as defined in source code) that
	// will be
	// executed. Defaults to the resource name suffix, if not specified.
	// For
	// backward compatibility, if function with given name is not found,
	// then the
	// system will try to use function named 'function'.
	// For Node.js this is name of a function exported by the module
	// specified
	// in source_location.
	EntryPoint string `json:"entryPoint,omitempty"`

	// GcsUrl: GCS URL pointing to the zip archive which contains the
	// function.
	GcsUrl string `json:"gcsUrl,omitempty"`

	// LatestOperation: [Output only] Name of the most recent operation
	// modifying the function. If
	// the function status is DEPLOYING or DELETING, then it points to the
	// active
	// operation.
	LatestOperation string `json:"latestOperation,omitempty"`

	// Name: A user-defined name of the function. Function names must be
	// unique
	// globally and match pattern: projects/*/regions/*/functions/*
	Name string `json:"name,omitempty"`

	// OauthScopes: The set of Google API scopes to be made available to the
	// function while
	// it is being executed. Values should be in the format of
	// scope
	// developer codes, for example:
	// "https://www.googleapis.com/auth/compute".
	OauthScopes []string `json:"oauthScopes,omitempty"`

	// SourceRepository: The hosted repository where the function is
	// defined.
	SourceRepository *SourceRepository `json:"sourceRepository,omitempty"`

	// Status: [Output only] Status of the function deployment.
	//
	// Possible values:
	//   "STATUS_UNSPECIFIED" - Status not specified.
	//   "READY" - Successfully deployed.
	//   "FAILED" - Not deployed correctly - behavior is undefined. The item
	// should be updated
	// or deleted to move it out of this state.
	//   "DEPLOYING" - Creation or update in progress.
	//   "DELETING" - Deletion in progress.
	Status string `json:"status,omitempty"`

	// Timeout: The cloud function execution timeout. Execution is
	// considered failed and
	// can be terminated if the function is not completed at the end of
	// the
	// timeout period. Defaults to 60 seconds.
	Timeout string `json:"timeout,omitempty"`

	// Triggers: List of triggers.
	Triggers []*FunctionTrigger `json:"triggers,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "EntryPoint") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *HostedFunction) MarshalJSON() ([]byte, error) {
	type noMethod HostedFunction
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// ListFunctionsResponse: Response for the ListFunctions method.
type ListFunctionsResponse struct {
	// Functions: The functions that match the request.
	Functions []*HostedFunction `json:"functions,omitempty"`

	// NextPageToken: If not empty, indicates that there may be more
	// functions that match
	// the request; this value should be passed in a new
	// ListFunctionsRequest
	// to get more functions.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Functions") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *ListFunctionsResponse) MarshalJSON() ([]byte, error) {
	type noMethod ListFunctionsResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// ListLocationsResponse: The response message for
// LocationService.ListLocations.
type ListLocationsResponse struct {
	// Locations: A list of locations that matches the specified filter in
	// the request.
	Locations []*Location `json:"locations,omitempty"`

	// NextPageToken: The standard List next-page token.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Locations") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *ListLocationsResponse) MarshalJSON() ([]byte, error) {
	type noMethod ListLocationsResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// Location: A resource that represents Google Cloud Platform location.
type Location struct {
	// Labels: Cross-service attributes for the location. For example
	//
	//     {"cloud.googleapis.com/region": "us-east1"}
	Labels map[string]string `json:"labels,omitempty"`

	// LocationId: The canonical id for this location. For example:
	// "us-east1".
	LocationId string `json:"locationId,omitempty"`

	// Metadata: Service-specific metadata. For example the available
	// capacity at the given
	// location.
	Metadata LocationMetadata `json:"metadata,omitempty"`

	// Name: Resource name for the location, which may vary between
	// implementations.
	// For example: "projects/example-project/locations/us-east1"
	Name string `json:"name,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Labels") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Location) MarshalJSON() ([]byte, error) {
	type noMethod Location
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

type LocationMetadata interface{}

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

// OperationMetadata1: Metadata describing an Operation
type OperationMetadata1 struct {
	// Request: The original request that started the operation.
	Request OperationMetadataRequest `json:"request,omitempty"`

	// Target: Target of the operation - for
	// example
	// projects/project-1/regions/region-1/functions/function-1
	Target string `json:"target,omitempty"`

	// Type: Type of operation.
	//
	// Possible values:
	//   "OPERATION_UNSPECIFIED" - Unknown operation type.
	//   "CREATE_FUNCTION" - Triggered by CreateFunction call
	//   "UPDATE_FUNCTION" - Triggered by UpdateFunction call
	//   "DELETE_FUNCTION" - Triggered by DeleteFunction call.
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Request") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *OperationMetadata1) MarshalJSON() ([]byte, error) {
	type noMethod OperationMetadata1
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

type OperationMetadataRequest interface{}

// SourceRepository: Describes the location of the function source in a
// remote repository.
type SourceRepository struct {
	// Branch: The name of the branch from which the function should be
	// fetched.
	Branch string `json:"branch,omitempty"`

	// DeployedRevision: [Output only] The id of the revision that was
	// resolved at the moment of
	// function creation or update. For example when a user deployed from
	// a
	// branch, it will be the revision id of the latest change on this
	// branch at
	// that time. If user deployed from revision then this value will be
	// always
	// equal to the revision specified by the user.
	DeployedRevision string `json:"deployedRevision,omitempty"`

	// RepositoryUrl: URL to the hosted repository where the function is
	// defined. Only paths in
	// https://source.developers.google.com domain are supported. The path
	// should
	// contain the name of the repository.
	RepositoryUrl string `json:"repositoryUrl,omitempty"`

	// Revision: The id of the revision that captures the state of the
	// repository from
	// which the function should be fetched.
	Revision string `json:"revision,omitempty"`

	// SourcePath: The path within the repository where the function is
	// defined. The path
	// should point to the directory where cloud functions files are
	// located. Use
	// '/' if the function is defined directly in the root directory of
	// a
	// repository.
	SourcePath string `json:"sourcePath,omitempty"`

	// Tag: The name of the tag that captures the state of the repository
	// from
	// which the function should be fetched.
	Tag string `json:"tag,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Branch") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *SourceRepository) MarshalJSON() ([]byte, error) {
	type noMethod SourceRepository
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

// WebTrigger: Describes WebTrigger, could be used to connect web hooks
// to function.
type WebTrigger struct {
	// Protocol: Protocol accepted by WebTrigger.
	//
	// Possible values:
	//   "HTTP" - HTTP protocol
	Protocol string `json:"protocol,omitempty"`

	// Url: [Output only] The deployed url for the function.
	Url string `json:"url,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Protocol") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *WebTrigger) MarshalJSON() ([]byte, error) {
	type noMethod WebTrigger
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// method id "cloudfunctions.operations.get":

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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+name}")
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

// Do executes the "cloudfunctions.operations.get" call.
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
	//   "flatPath": "v1beta1/operations/{operationsId}",
	//   "httpMethod": "GET",
	//   "id": "cloudfunctions.operations.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The name of the operation resource.",
	//       "location": "path",
	//       "pattern": "^operations/[^/]*$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+name}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudfunctions.projects.locations.list":

type ProjectsLocationsListCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// List: Lists information about the supported locations for this
// service.
func (r *ProjectsLocationsService) List(name string) *ProjectsLocationsListCall {
	c := &ProjectsLocationsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Filter sets the optional parameter "filter": The standard list
// filter.
func (c *ProjectsLocationsListCall) Filter(filter string) *ProjectsLocationsListCall {
	c.urlParams_.Set("filter", filter)
	return c
}

// PageSize sets the optional parameter "pageSize": The standard list
// page size.
func (c *ProjectsLocationsListCall) PageSize(pageSize int64) *ProjectsLocationsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The standard list
// page token.
func (c *ProjectsLocationsListCall) PageToken(pageToken string) *ProjectsLocationsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsLocationsListCall) Fields(s ...googleapi.Field) *ProjectsLocationsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsLocationsListCall) IfNoneMatch(entityTag string) *ProjectsLocationsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsLocationsListCall) Context(ctx context.Context) *ProjectsLocationsListCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsLocationsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+name}/locations")
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

// Do executes the "cloudfunctions.projects.locations.list" call.
// Exactly one of *ListLocationsResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListLocationsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsLocationsListCall) Do(opts ...googleapi.CallOption) (*ListLocationsResponse, error) {
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
	ret := &ListLocationsResponse{
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
	//   "description": "Lists information about the supported locations for this service.",
	//   "flatPath": "v1beta1/projects/{projectsId}/locations",
	//   "httpMethod": "GET",
	//   "id": "cloudfunctions.projects.locations.list",
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
	//       "description": "The resource that owns the locations collection, if applicable.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]*$",
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
	//   "path": "v1beta1/{+name}/locations",
	//   "response": {
	//     "$ref": "ListLocationsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsLocationsListCall) Pages(ctx context.Context, f func(*ListLocationsResponse) error) error {
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

// method id "cloudfunctions.projects.regions.functions.call":

type ProjectsRegionsFunctionsCallCall struct {
	s                   *Service
	name                string
	callfunctionrequest *CallFunctionRequest
	urlParams_          gensupport.URLParams
	ctx_                context.Context
}

// Call: Invokes synchronously deployed function. To be used for
// testing, very
// limited traffic allowed.
func (r *ProjectsRegionsFunctionsService) Call(name string, callfunctionrequest *CallFunctionRequest) *ProjectsRegionsFunctionsCallCall {
	c := &ProjectsRegionsFunctionsCallCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.callfunctionrequest = callfunctionrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsRegionsFunctionsCallCall) Fields(s ...googleapi.Field) *ProjectsRegionsFunctionsCallCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsRegionsFunctionsCallCall) Context(ctx context.Context) *ProjectsRegionsFunctionsCallCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsRegionsFunctionsCallCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.callfunctionrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+name}:call")
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

// Do executes the "cloudfunctions.projects.regions.functions.call" call.
// Exactly one of *CallFunctionResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *CallFunctionResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsRegionsFunctionsCallCall) Do(opts ...googleapi.CallOption) (*CallFunctionResponse, error) {
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
	ret := &CallFunctionResponse{
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
	//   "description": "Invokes synchronously deployed function. To be used for testing, very\nlimited traffic allowed.",
	//   "flatPath": "v1beta1/projects/{projectsId}/regions/{regionsId}/functions/{functionsId}:call",
	//   "httpMethod": "POST",
	//   "id": "cloudfunctions.projects.regions.functions.call",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The name of the function to be called.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]*/regions/[^/]*/functions/[^/]*$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+name}:call",
	//   "request": {
	//     "$ref": "CallFunctionRequest"
	//   },
	//   "response": {
	//     "$ref": "CallFunctionResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudfunctions.projects.regions.functions.create":

type ProjectsRegionsFunctionsCreateCall struct {
	s              *Service
	location       string
	hostedfunction *HostedFunction
	urlParams_     gensupport.URLParams
	ctx_           context.Context
}

// Create: Creates a new function. If a function with the given name
// already exists in
// the specified project, it will return ALREADY_EXISTS error.
func (r *ProjectsRegionsFunctionsService) Create(location string, hostedfunction *HostedFunction) *ProjectsRegionsFunctionsCreateCall {
	c := &ProjectsRegionsFunctionsCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.location = location
	c.hostedfunction = hostedfunction
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsRegionsFunctionsCreateCall) Fields(s ...googleapi.Field) *ProjectsRegionsFunctionsCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsRegionsFunctionsCreateCall) Context(ctx context.Context) *ProjectsRegionsFunctionsCreateCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsRegionsFunctionsCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.hostedfunction)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+location}/functions")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"location": c.location,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "cloudfunctions.projects.regions.functions.create" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsRegionsFunctionsCreateCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
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
	//   "description": "Creates a new function. If a function with the given name already exists in\nthe specified project, it will return ALREADY_EXISTS error.",
	//   "flatPath": "v1beta1/projects/{projectsId}/regions/{regionsId}/functions",
	//   "httpMethod": "POST",
	//   "id": "cloudfunctions.projects.regions.functions.create",
	//   "parameterOrder": [
	//     "location"
	//   ],
	//   "parameters": {
	//     "location": {
	//       "description": "The project and region in which the function should be created, specified\nin the format: projects/*/regions/*",
	//       "location": "path",
	//       "pattern": "^projects/[^/]*/regions/[^/]*$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+location}/functions",
	//   "request": {
	//     "$ref": "HostedFunction"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudfunctions.projects.regions.functions.delete":

type ProjectsRegionsFunctionsDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
}

// Delete: Deletes a function with the given name from the specified
// project. If the
// given function is used by some trigger, the trigger will be updated
// to
// remove this function.
func (r *ProjectsRegionsFunctionsService) Delete(name string) *ProjectsRegionsFunctionsDeleteCall {
	c := &ProjectsRegionsFunctionsDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsRegionsFunctionsDeleteCall) Fields(s ...googleapi.Field) *ProjectsRegionsFunctionsDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsRegionsFunctionsDeleteCall) Context(ctx context.Context) *ProjectsRegionsFunctionsDeleteCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsRegionsFunctionsDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+name}")
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

// Do executes the "cloudfunctions.projects.regions.functions.delete" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsRegionsFunctionsDeleteCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
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
	//   "description": "Deletes a function with the given name from the specified project. If the\ngiven function is used by some trigger, the trigger will be updated to\nremove this function.",
	//   "flatPath": "v1beta1/projects/{projectsId}/regions/{regionsId}/functions/{functionsId}",
	//   "httpMethod": "DELETE",
	//   "id": "cloudfunctions.projects.regions.functions.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The name of the function which should be deleted.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]*/regions/[^/]*/functions/[^/]*$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+name}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudfunctions.projects.regions.functions.get":

type ProjectsRegionsFunctionsGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// Get: Returns a function with the given name from the requested
// project.
func (r *ProjectsRegionsFunctionsService) Get(name string) *ProjectsRegionsFunctionsGetCall {
	c := &ProjectsRegionsFunctionsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsRegionsFunctionsGetCall) Fields(s ...googleapi.Field) *ProjectsRegionsFunctionsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsRegionsFunctionsGetCall) IfNoneMatch(entityTag string) *ProjectsRegionsFunctionsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsRegionsFunctionsGetCall) Context(ctx context.Context) *ProjectsRegionsFunctionsGetCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsRegionsFunctionsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+name}")
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

// Do executes the "cloudfunctions.projects.regions.functions.get" call.
// Exactly one of *HostedFunction or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *HostedFunction.ServerResponse.Header or (if a response was returned
// at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsRegionsFunctionsGetCall) Do(opts ...googleapi.CallOption) (*HostedFunction, error) {
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
	ret := &HostedFunction{
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
	//   "description": "Returns a function with the given name from the requested project.",
	//   "flatPath": "v1beta1/projects/{projectsId}/regions/{regionsId}/functions/{functionsId}",
	//   "httpMethod": "GET",
	//   "id": "cloudfunctions.projects.regions.functions.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The name of the function which details should be obtained.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]*/regions/[^/]*/functions/[^/]*$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+name}",
	//   "response": {
	//     "$ref": "HostedFunction"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudfunctions.projects.regions.functions.list":

type ProjectsRegionsFunctionsListCall struct {
	s            *Service
	location     string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// List: Returns a list of all functions that belong to the requested
// project.
func (r *ProjectsRegionsFunctionsService) List(location string) *ProjectsRegionsFunctionsListCall {
	c := &ProjectsRegionsFunctionsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.location = location
	return c
}

// PageSize sets the optional parameter "pageSize": Maximum number of
// functions to return.
func (c *ProjectsRegionsFunctionsListCall) PageSize(pageSize int64) *ProjectsRegionsFunctionsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The value returned
// by the last ListFunctionsResponse; indicates that
// this is a continuation of a prior ListFunctions call, and that
// the
// system should return the next page of data.
func (c *ProjectsRegionsFunctionsListCall) PageToken(pageToken string) *ProjectsRegionsFunctionsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsRegionsFunctionsListCall) Fields(s ...googleapi.Field) *ProjectsRegionsFunctionsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsRegionsFunctionsListCall) IfNoneMatch(entityTag string) *ProjectsRegionsFunctionsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsRegionsFunctionsListCall) Context(ctx context.Context) *ProjectsRegionsFunctionsListCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsRegionsFunctionsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+location}/functions")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"location": c.location,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "cloudfunctions.projects.regions.functions.list" call.
// Exactly one of *ListFunctionsResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListFunctionsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsRegionsFunctionsListCall) Do(opts ...googleapi.CallOption) (*ListFunctionsResponse, error) {
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
	ret := &ListFunctionsResponse{
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
	//   "description": "Returns a list of all functions that belong to the requested project.",
	//   "flatPath": "v1beta1/projects/{projectsId}/regions/{regionsId}/functions",
	//   "httpMethod": "GET",
	//   "id": "cloudfunctions.projects.regions.functions.list",
	//   "parameterOrder": [
	//     "location"
	//   ],
	//   "parameters": {
	//     "location": {
	//       "description": "The project and region from which functions should be listed, specified\nin the format: projects/*/regions/*",
	//       "location": "path",
	//       "pattern": "^projects/[^/]*/regions/[^/]*$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Maximum number of functions to return.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The value returned by the last ListFunctionsResponse; indicates that\nthis is a continuation of a prior ListFunctions call, and that the\nsystem should return the next page of data.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+location}/functions",
	//   "response": {
	//     "$ref": "ListFunctionsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsRegionsFunctionsListCall) Pages(ctx context.Context, f func(*ListFunctionsResponse) error) error {
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

// method id "cloudfunctions.projects.regions.functions.update":

type ProjectsRegionsFunctionsUpdateCall struct {
	s              *Service
	name           string
	hostedfunction *HostedFunction
	urlParams_     gensupport.URLParams
	ctx_           context.Context
}

// Update: Updates existing function.
func (r *ProjectsRegionsFunctionsService) Update(name string, hostedfunction *HostedFunction) *ProjectsRegionsFunctionsUpdateCall {
	c := &ProjectsRegionsFunctionsUpdateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.hostedfunction = hostedfunction
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsRegionsFunctionsUpdateCall) Fields(s ...googleapi.Field) *ProjectsRegionsFunctionsUpdateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsRegionsFunctionsUpdateCall) Context(ctx context.Context) *ProjectsRegionsFunctionsUpdateCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsRegionsFunctionsUpdateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.hostedfunction)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+name}")
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

// Do executes the "cloudfunctions.projects.regions.functions.update" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsRegionsFunctionsUpdateCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
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
	//   "description": "Updates existing function.",
	//   "flatPath": "v1beta1/projects/{projectsId}/regions/{regionsId}/functions/{functionsId}",
	//   "httpMethod": "PUT",
	//   "id": "cloudfunctions.projects.regions.functions.update",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The name of the function to be updated.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]*/regions/[^/]*/functions/[^/]*$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+name}",
	//   "request": {
	//     "$ref": "HostedFunction"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}
