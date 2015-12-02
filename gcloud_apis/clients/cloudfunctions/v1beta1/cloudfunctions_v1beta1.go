// Package cloudfunctions provides access to the Google Cloud Functions API.
//
// See https://docs.google.com/document/d/16HutoQJLbqIhXYoz9IcyJ36b4tpD8g2E2isIwK1sbFY/view
//
// Usage example:
//
//   import "google.golang.org/api/cloudfunctions/v1beta1"
//   ...
//   cloudfunctionsService, err := cloudfunctions.New(oauthHttpClient)
package cloudfunctions

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
	client   *http.Client
	BasePath string // API endpoint base URL

	Operations *OperationsService

	Projects *ProjectsService
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
	rs.Regions = NewProjectsRegionsService(s)
	return rs
}

type ProjectsService struct {
	s *Service

	Regions *ProjectsRegionsService
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

type CallFunctionResponse struct {
	// Error: Either system or user-function generated error. Set if
	// execution
	// was not successful.
	Error string `json:"error,omitempty"`

	// ExecutionId: Execution id of function invocation.
	ExecutionId string `json:"executionId,omitempty"`

	// Result: Result populated for successful execution of synchronous
	// function. Will
	// not be populated if function does not return a result
	// through context.
	Result string `json:"result,omitempty"`
}

type FunctionTrigger struct {
	// GsUri: Google Cloud Storage resource whose changes trigger the
	// events.
	// Currently, it must have the form gs://<bucket>/ (that is, it
	// must refer
	// to a bucket, rather than an object).
	GsUri string `json:"gsUri,omitempty"`

	// PubsubTopic: A pub/sub type of source.
	PubsubTopic string `json:"pubsubTopic,omitempty"`

	// WebTrigger: A web endpoint (e.g. HTTP) type of source that can be
	// trigger via URL.
	WebTrigger *WebTrigger `json:"webTrigger,omitempty"`
}

type HostedFunction struct {
	// EntryPoint: The name of the function (as defined in source code) that
	// will be
	// executed. Defaults to the resource name suffix, if not
	// specified. For
	// backward compatibility, if function with given name is
	// not found, then the
	// system will try to use function named
	// 'function'.
	// For Node.js this is name of a function exported by the
	// module specified
	// in source_location.
	EntryPoint string `json:"entryPoint,omitempty"`

	// GcsUrl: GCS URL pointing to the zip archive which contains the
	// function.
	GcsUrl string `json:"gcsUrl,omitempty"`

	// LatestOperation: [Output only] Name of the most recent operation
	// modifying the function. If
	// the function status is DEPLOYING or
	// DELETING, then it points to the active
	// operation.
	LatestOperation string `json:"latestOperation,omitempty"`

	// Name: A user-defined name of the function. Function names must be
	// unique
	// globally and match pattern: projects/*/regions/*/functions/*
	Name string `json:"name,omitempty"`

	// OauthScopes: The set of Google API scopes to be made available to the
	// function while
	// it is being executed. Values should be in the format
	// of scope
	// developer codes, for example:
	// "https://www.googleapis.com/auth/compute".
	OauthScopes []string `json:"oauthScopes,omitempty"`

	// SourceRepository: The hosted repository where the function is
	// defined.
	SourceRepository *SourceRepository `json:"sourceRepository,omitempty"`

	// Status: [Output only] Status of the function deployment.
	Status string `json:"status,omitempty"`

	// Triggers: List of triggers.
	Triggers []*FunctionTrigger `json:"triggers,omitempty"`
}

type ListFunctionsResponse struct {
	// Functions: The functions that match the request.
	Functions []*HostedFunction `json:"functions,omitempty"`

	// NextPageToken: If not empty, indicates that there may be more
	// functions that match
	// the request; this value should be passed in a
	// new ListFunctionsRequest
	// to get more functions.
	NextPageToken string `json:"nextPageToken,omitempty"`
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

type OperationMetadata1 struct {
	// Request: The original request that started the operation.
	Request *OperationMetadataRequest `json:"request,omitempty"`

	// Target: Target of the operation - for
	// example
	// projects/project-1/regions/region-1/functions/function-1
	Target string `json:"target,omitempty"`

	// Type: Type of operation.
	Type string `json:"type,omitempty"`
}

type OperationMetadataRequest struct {
}

type SourceRepository struct {
	// Branch: The name of the branch from which the function should be
	// fetched.
	Branch string `json:"branch,omitempty"`

	// DeployedRevision: [Output only] The id of the revision that was
	// resolved at the moment of
	// function creation or update. For example
	// when a user deployed from a
	// branch, it will be the revision id of the
	// latest change on this branch at
	// that time. If user deployed from
	// revision then this value will be always
	// equal to the revision
	// specified by the user.
	DeployedRevision string `json:"deployedRevision,omitempty"`

	// RepositoryUrl: URL to the hosted repository where the function is
	// defined. Only paths in
	// https://source.developers.google.com domain
	// are supported. The path should
	// contain the name of the repository.
	RepositoryUrl string `json:"repositoryUrl,omitempty"`

	// Revision: The id of the revision that captures the state of the
	// repository from
	// which the function should be fetched.
	Revision string `json:"revision,omitempty"`

	// SourcePath: The path within the repository where the function is
	// defined. The path
	// should point to the directory where cloud functions
	// files are located. Use
	// '/' if the function is defined directly in the
	// root directory of a
	// repository.
	SourcePath string `json:"sourcePath,omitempty"`

	SourceUrl string `json:"sourceUrl,omitempty"`

	// Tag: The name of the tag that captures the state of the repository
	// from
	// which the function should be fetched.
	Tag string `json:"tag,omitempty"`
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

type WebTrigger struct {
	// Protocol: Protocol accepted by WebTrigger.
	Protocol string `json:"protocol,omitempty"`

	// Url: [Output only] The deployed url for the function.
	Url string `json:"url,omitempty"`
}

// method id "cloudfunctions.operations.get":

type OperationsGetCall struct {
	s    *Service
	name string
	opt_ map[string]interface{}
}

// Get: Gets the latest state of a long-running operation.  Clients can
// use this
// method to poll the operation result at intervals as
// recommended by the API
// service.
func (r *OperationsService) Get(name string) *OperationsGetCall {
	c := &OperationsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OperationsGetCall) Fields(s ...googleapi.Field) *OperationsGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *OperationsGetCall) Do() (*Operation, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+name}")
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

// method id "cloudfunctions.projects.regions.functions.call":

type ProjectsRegionsFunctionsCallCall struct {
	s    *Service
	name string
	opt_ map[string]interface{}
}

// Call: Invokes synchronously deployed function. To be used for
// testing, very
// limited traffic allowed.
func (r *ProjectsRegionsFunctionsService) Call(name string) *ProjectsRegionsFunctionsCallCall {
	c := &ProjectsRegionsFunctionsCallCall{s: r.s, opt_: make(map[string]interface{})}
	c.name = name
	return c
}

// Data sets the optional parameter "data": Input to be passed to the
// function.
func (c *ProjectsRegionsFunctionsCallCall) Data(data string) *ProjectsRegionsFunctionsCallCall {
	c.opt_["data"] = data
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsRegionsFunctionsCallCall) Fields(s ...googleapi.Field) *ProjectsRegionsFunctionsCallCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsRegionsFunctionsCallCall) Do() (*CallFunctionResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["data"]; ok {
		params.Set("data", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+name}:call")
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
	var ret *CallFunctionResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
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
	//     "data": {
	//       "description": "Input to be passed to the function.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "name": {
	//       "description": "The name of the function to be called.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]*/regions/[^/]*/functions/[^/]*$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/{+name}:call",
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
	opt_           map[string]interface{}
}

// Create: Creates a new function. If a function with the given name
// already exists in
// the specified project, it will return
// ALREADY_EXISTS error.
//
func (r *ProjectsRegionsFunctionsService) Create(location string, hostedfunction *HostedFunction) *ProjectsRegionsFunctionsCreateCall {
	c := &ProjectsRegionsFunctionsCreateCall{s: r.s, opt_: make(map[string]interface{})}
	c.location = location
	c.hostedfunction = hostedfunction
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsRegionsFunctionsCreateCall) Fields(s ...googleapi.Field) *ProjectsRegionsFunctionsCreateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsRegionsFunctionsCreateCall) Do() (*Operation, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.hostedfunction)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+location}/functions")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"location": c.location,
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
	//   "description": "Creates a new function. If a function with the given name already exists in\nthe specified project, it will return ALREADY_EXISTS error.\n",
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
	s    *Service
	name string
	opt_ map[string]interface{}
}

// Delete: Deletes a function with the given name from the specified
// project. If the
// given function is used by some trigger, the trigger
// will be updated to
// remove this function.
//
func (r *ProjectsRegionsFunctionsService) Delete(name string) *ProjectsRegionsFunctionsDeleteCall {
	c := &ProjectsRegionsFunctionsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsRegionsFunctionsDeleteCall) Fields(s ...googleapi.Field) *ProjectsRegionsFunctionsDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsRegionsFunctionsDeleteCall) Do() (*Operation, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+name}")
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
	var ret *Operation
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes a function with the given name from the specified project. If the\ngiven function is used by some trigger, the trigger will be updated to\nremove this function.\n",
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
	s    *Service
	name string
	opt_ map[string]interface{}
}

// Get: Returns a function with the given name from the requested
// project.
func (r *ProjectsRegionsFunctionsService) Get(name string) *ProjectsRegionsFunctionsGetCall {
	c := &ProjectsRegionsFunctionsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsRegionsFunctionsGetCall) Fields(s ...googleapi.Field) *ProjectsRegionsFunctionsGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsRegionsFunctionsGetCall) Do() (*HostedFunction, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+name}")
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
	var ret *HostedFunction
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
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
	s        *Service
	location string
	opt_     map[string]interface{}
}

// List: Returns a list of all functions that belong to the requested
// project.
func (r *ProjectsRegionsFunctionsService) List(location string) *ProjectsRegionsFunctionsListCall {
	c := &ProjectsRegionsFunctionsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.location = location
	return c
}

// PageSize sets the optional parameter "pageSize": Maximum number of
// functions to return.
func (c *ProjectsRegionsFunctionsListCall) PageSize(pageSize int64) *ProjectsRegionsFunctionsListCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken": The value returned
// by the last ListFunctionsResponse; indicates that
// this is a
// continuation of a prior ListFunctions call, and that the
// system
// should return the next page of data.
func (c *ProjectsRegionsFunctionsListCall) PageToken(pageToken string) *ProjectsRegionsFunctionsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsRegionsFunctionsListCall) Fields(s ...googleapi.Field) *ProjectsRegionsFunctionsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsRegionsFunctionsListCall) Do() (*ListFunctionsResponse, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+location}/functions")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"location": c.location,
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
	var ret *ListFunctionsResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
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
	//       "description": "The project and region in which the function should be created, specified\nin the format: projects/*/regions/*",
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

// method id "cloudfunctions.projects.regions.functions.update":

type ProjectsRegionsFunctionsUpdateCall struct {
	s              *Service
	name           string
	hostedfunction *HostedFunction
	opt_           map[string]interface{}
}

// Update: Updates existing function.
//
func (r *ProjectsRegionsFunctionsService) Update(name string, hostedfunction *HostedFunction) *ProjectsRegionsFunctionsUpdateCall {
	c := &ProjectsRegionsFunctionsUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.name = name
	c.hostedfunction = hostedfunction
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsRegionsFunctionsUpdateCall) Fields(s ...googleapi.Field) *ProjectsRegionsFunctionsUpdateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsRegionsFunctionsUpdateCall) Do() (*Operation, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.hostedfunction)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+name}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
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
	//   "description": "Updates existing function.\n",
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
