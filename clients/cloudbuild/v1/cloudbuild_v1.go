// Package cloudbuild provides access to the Google Cloud Container Builder API.
//
// Usage example:
//
//   import "google.golang.org/api/cloudbuild/v1"
//   ...
//   cloudbuildService, err := cloudbuild.New(oauthHttpClient)
package cloudbuild

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

const apiId = "cloudbuild:v1"
const apiName = "cloudbuild"
const apiVersion = "v1"
const basePath = "https://cloudbuild.googleapis.com/"

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
	rs.Builds = NewProjectsBuildsService(s)
	return rs
}

type ProjectsService struct {
	s *Service

	Builds *ProjectsBuildsService
}

func NewProjectsBuildsService(s *Service) *ProjectsBuildsService {
	rs := &ProjectsBuildsService{s: s}
	return rs
}

type ProjectsBuildsService struct {
	s *Service
}

type AliasContext struct {
	// Kind: The alias kind.
	Kind string `json:"kind,omitempty"`

	// Name: The alias name.
	Name string `json:"name,omitempty"`
}

type Build struct {
	// CreateTime: The time that the build was created.
	// @OutputOnly
	CreateTime string `json:"createTime,omitempty"`

	// FinishTime: The time that execution of the build was
	// finished.
	// @OutputOnly
	FinishTime string `json:"finishTime,omitempty"`

	// Id: The unique identifier of the build.
	// @OutputOnly
	Id string `json:"id,omitempty"`

	// Images: The list of images expected to be built and pushed to GCR. If
	// an image
	// is listed here, and if the image is not produced by one of
	// the build
	// steps, the build will fail. If all the images are present
	// when the build
	// steps are complete, they will all be pushed.
	Images []string `json:"images,omitempty"`

	// ProjectId: The ID of the project.
	// @OutputOnly.
	ProjectId string `json:"projectId,omitempty"`

	// Source: Describes where to find source files to build.
	Source *Source `json:"source,omitempty"`

	// StartTime: The time that execution of the build was
	// started.
	// @OutputOnly
	StartTime string `json:"startTime,omitempty"`

	// Status: The status of the build.
	// @OutputOnly
	Status string `json:"status,omitempty"`

	// Steps: Describes the operations to be performed on the workspace.
	Steps []*BuildStep `json:"steps,omitempty"`

	// Timeout: The amount of time that this build should be allowed to run,
	// to second
	// granularity. If this amount of time elapses, work on the
	// build will cease
	// and the build status will be TIMEOUT.
	//
	// By default,
	// this will be ten minutes.
	Timeout string `json:"timeout,omitempty"`
}

type BuildStep struct {
	// Args: Command-line arguments to use when running this step's
	// container.
	Args []string `json:"args,omitempty"`

	// Dir: Working directory (relative to project source root) to use when
	// running
	// this operation's container.
	Dir string `json:"dir,omitempty"`

	// Env: Additional environment variables to set for this step's
	// container.
	Env []string `json:"env,omitempty"`

	// Name: The name of the container image to use for creating this stage
	// in the
	// pipeline, as presented to 'docker pull'.
	Name string `json:"name,omitempty"`
}

type CancelOperationRequest struct {
}

type CloudRepoSourceContext struct {
	// AliasContext: An alias, which may be a branch or tag.
	AliasContext *AliasContext `json:"aliasContext,omitempty"`

	// AliasName: The name of an alias (branch, tag, etc.).
	AliasName string `json:"aliasName,omitempty"`

	// RepoId: The ID of the repo.
	RepoId *RepoId `json:"repoId,omitempty"`

	// RevisionId: A revision ID.
	RevisionId string `json:"revisionId,omitempty"`
}

type CloudWorkspaceId struct {
	// Name: The unique name of the workspace within the repo.  This is the
	// name
	// chosen by the client in the Source API's CreateWorkspace method.
	Name string `json:"name,omitempty"`

	// RepoId: The ID of the repo containing the workspace.
	RepoId *RepoId `json:"repoId,omitempty"`
}

type CloudWorkspaceSourceContext struct {
	// SnapshotId: The ID of the snapshot.
	// An empty snapshot_id refers to
	// the most recent snapshot.
	SnapshotId string `json:"snapshotId,omitempty"`

	// WorkspaceId: The ID of the workspace.
	WorkspaceId *CloudWorkspaceId `json:"workspaceId,omitempty"`
}

type Empty struct {
}

type GerritSourceContext struct {
	// AliasContext: An alias, which may be a branch or tag.
	AliasContext *AliasContext `json:"aliasContext,omitempty"`

	// AliasName: The name of an alias (branch, tag, etc.).
	AliasName string `json:"aliasName,omitempty"`

	// GerritProject: The full project name within the host. Projects may be
	// nested, so
	// "project/subproject" is a valid project name.
	// The "repo
	// name" is hostURI/project.
	GerritProject string `json:"gerritProject,omitempty"`

	// HostUri: The URI of a running Gerrit instance.
	HostUri string `json:"hostUri,omitempty"`

	// RevisionId: A revision (commit) ID.
	RevisionId string `json:"revisionId,omitempty"`
}

type GitSourceContext struct {
	// RevisionId: Git commit hash.
	// required.
	RevisionId string `json:"revisionId,omitempty"`

	// Url: Git repository URL.
	Url string `json:"url,omitempty"`
}

type ListBuildsResponse struct {
	// Builds: Builds will be sorted by create_time, descending.
	Builds []*Build `json:"builds,omitempty"`

	// NextPageToken: Token to receive the next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type ListOperationsResponse struct {
	// NextPageToken: The standard List next-page token.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Operations: A list of operations that matches the specified filter in
	// the request.
	Operations []*Operation `json:"operations,omitempty"`
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

type ProjectRepoId struct {
	// ProjectId: The ID of the project.
	ProjectId string `json:"projectId,omitempty"`

	// RepoName: The name of the repo. Leave empty for the default repo.
	RepoName string `json:"repoName,omitempty"`
}

type RepoId struct {
	// ProjectRepoId: A combination of a project ID and a repo name.
	ProjectRepoId *ProjectRepoId `json:"projectRepoId,omitempty"`

	// Uid: A server-assigned, globally unique identifier.
	Uid string `json:"uid,omitempty"`
}

type Source struct {
	// RepoSource: If provided, get source from this location in source
	// control.
	RepoSource *SourceContext `json:"repoSource,omitempty"`

	// StorageSource: If provided, get source from this location in in
	// Google Cloud Storage.
	StorageSource *StorageSource `json:"storageSource,omitempty"`
}

type SourceContext struct {
	// CloudRepo: A SourceContext referring to a revision in a cloud repo.
	CloudRepo *CloudRepoSourceContext `json:"cloudRepo,omitempty"`

	// CloudWorkspace: A SourceContext referring to a snapshot in a cloud
	// workspace.
	CloudWorkspace *CloudWorkspaceSourceContext `json:"cloudWorkspace,omitempty"`

	// Gerrit: A SourceContext referring to a Gerrit project.
	Gerrit *GerritSourceContext `json:"gerrit,omitempty"`

	// Git: A SourceContext referring to any third party Git repo (e.g.
	// GitHub).
	Git *GitSourceContext `json:"git,omitempty"`
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

type StorageSource struct {
	// Bucket: Google Cloud Storage bucket containing source (see
	// [Bucket
	// Name
	// Requirements](https://cloud.google.com/storage/docs/bucket-naming#requ
	// irements)).
	Bucket string `json:"bucket,omitempty"`

	// Object: Google Cloud Storage object containing source.
	//
	// This object
	// must be an archive file (zip, tar, tar.gz) containing source
	// to
	// build.
	Object string `json:"object,omitempty"`
}

// method id "cloudbuild.operations.cancel":

type OperationsCancelCall struct {
	s                      *Service
	name                   string
	canceloperationrequest *CancelOperationRequest
	opt_                   map[string]interface{}
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
func (r *OperationsService) Cancel(name string, canceloperationrequest *CancelOperationRequest) *OperationsCancelCall {
	c := &OperationsCancelCall{s: r.s, opt_: make(map[string]interface{})}
	c.name = name
	c.canceloperationrequest = canceloperationrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OperationsCancelCall) Fields(s ...googleapi.Field) *OperationsCancelCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *OperationsCancelCall) Do() (*Empty, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.canceloperationrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}:cancel")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
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
	var ret *Empty
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Starts asynchronous cancellation on a long-running operation.  The server\nmakes a best effort to cancel the operation, but success is not\nguaranteed.  If the server doesn't support this method, it returns\n`google.rpc.Code.UNIMPLEMENTED`.  Clients can use\nOperations.GetOperation or\nother methods to check whether the cancellation succeeded or whether the\noperation completed despite cancellation.",
	//   "flatPath": "v1/operations/{operationsId}:cancel",
	//   "httpMethod": "POST",
	//   "id": "cloudbuild.operations.cancel",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The name of the operation resource to be cancelled.",
	//       "location": "path",
	//       "pattern": "^operations/.*$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}:cancel",
	//   "request": {
	//     "$ref": "CancelOperationRequest"
	//   },
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudbuild.operations.delete":

type OperationsDeleteCall struct {
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
func (r *OperationsService) Delete(name string) *OperationsDeleteCall {
	c := &OperationsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OperationsDeleteCall) Fields(s ...googleapi.Field) *OperationsDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *OperationsDeleteCall) Do() (*Empty, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}")
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
	//   "flatPath": "v1/operations/{operationsId}",
	//   "httpMethod": "DELETE",
	//   "id": "cloudbuild.operations.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The name of the operation resource to be deleted.",
	//       "location": "path",
	//       "pattern": "^operations/.*$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudbuild.operations.get":

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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}")
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
	//   "flatPath": "v1/operations/{operationsId}",
	//   "httpMethod": "GET",
	//   "id": "cloudbuild.operations.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The name of the operation resource.",
	//       "location": "path",
	//       "pattern": "^operations/.*$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudbuild.operations.list":

type OperationsListCall struct {
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
func (r *OperationsService) List(name string) *OperationsListCall {
	c := &OperationsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.name = name
	return c
}

// Filter sets the optional parameter "filter": The standard list
// filter.
func (c *OperationsListCall) Filter(filter string) *OperationsListCall {
	c.opt_["filter"] = filter
	return c
}

// PageSize sets the optional parameter "pageSize": The standard list
// page size.
func (c *OperationsListCall) PageSize(pageSize int64) *OperationsListCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken": The standard list
// page token.
func (c *OperationsListCall) PageToken(pageToken string) *OperationsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OperationsListCall) Fields(s ...googleapi.Field) *OperationsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *OperationsListCall) Do() (*ListOperationsResponse, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}")
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
	//   "flatPath": "v1/operations",
	//   "httpMethod": "GET",
	//   "id": "cloudbuild.operations.list",
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
	//   "path": "v1/{+name}",
	//   "response": {
	//     "$ref": "ListOperationsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudbuild.projects.builds.cancel":

type ProjectsBuildsCancelCall struct {
	s         *Service
	projectId string
	id        string
	opt_      map[string]interface{}
}

// Cancel: Cancels a requested build in progress.
func (r *ProjectsBuildsService) Cancel(projectId string, id string) *ProjectsBuildsCancelCall {
	c := &ProjectsBuildsCancelCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.id = id
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsBuildsCancelCall) Fields(s ...googleapi.Field) *ProjectsBuildsCancelCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsBuildsCancelCall) Do() (*Build, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/builds/{id}:cancel")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"id":        c.id,
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
	var ret *Build
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Cancels a requested build in progress.",
	//   "flatPath": "v1/projects/{projectId}/builds/{id}:cancel",
	//   "httpMethod": "POST",
	//   "id": "cloudbuild.projects.builds.cancel",
	//   "parameterOrder": [
	//     "projectId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The ID of the build.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "The ID of the project.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/builds/{id}:cancel",
	//   "response": {
	//     "$ref": "Build"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudbuild.projects.builds.create":

type ProjectsBuildsCreateCall struct {
	s         *Service
	projectId string
	build     *Build
	opt_      map[string]interface{}
}

// Create: Starts a build with the specified configuration.
//
// The
// long-running Operation returned by this method will include the ID
// of
// the build, which can be passed to GetBuild to determine its status
// (e.g.,
// success or failure).
func (r *ProjectsBuildsService) Create(projectId string, build *Build) *ProjectsBuildsCreateCall {
	c := &ProjectsBuildsCreateCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.build = build
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsBuildsCreateCall) Fields(s ...googleapi.Field) *ProjectsBuildsCreateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsBuildsCreateCall) Do() (*Operation, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.build)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/builds")
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
	//   "description": "Starts a build with the specified configuration.\n\nThe long-running Operation returned by this method will include the ID of\nthe build, which can be passed to GetBuild to determine its status (e.g.,\nsuccess or failure).",
	//   "flatPath": "v1/projects/{projectId}/builds",
	//   "httpMethod": "POST",
	//   "id": "cloudbuild.projects.builds.create",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "The ID of the project.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/builds",
	//   "request": {
	//     "$ref": "Build"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudbuild.projects.builds.get":

type ProjectsBuildsGetCall struct {
	s         *Service
	projectId string
	id        string
	opt_      map[string]interface{}
}

// Get: Returns information about a previously requested build.
//
// The
// Build that is returned includes its status (e.g., success or
// failure,
// or in-progress), and timing information.
func (r *ProjectsBuildsService) Get(projectId string, id string) *ProjectsBuildsGetCall {
	c := &ProjectsBuildsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.id = id
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsBuildsGetCall) Fields(s ...googleapi.Field) *ProjectsBuildsGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsBuildsGetCall) Do() (*Build, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/builds/{id}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"id":        c.id,
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
	var ret *Build
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns information about a previously requested build.\n\nThe Build that is returned includes its status (e.g., success or failure,\nor in-progress), and timing information.",
	//   "flatPath": "v1/projects/{projectId}/builds/{id}",
	//   "httpMethod": "GET",
	//   "id": "cloudbuild.projects.builds.get",
	//   "parameterOrder": [
	//     "projectId",
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The ID of the build.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "The ID of the project.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/builds/{id}",
	//   "response": {
	//     "$ref": "Build"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudbuild.projects.builds.list":

type ProjectsBuildsListCall struct {
	s         *Service
	projectId string
	opt_      map[string]interface{}
}

// List: Lists previously requested builds.
//
// Previously requested builds
// may still be in-progress, or may have finished
// successfully or
// unsuccessfully.
func (r *ProjectsBuildsService) List(projectId string) *ProjectsBuildsListCall {
	c := &ProjectsBuildsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	return c
}

// PageSize sets the optional parameter "pageSize": Number of results to
// return in the list.
func (c *ProjectsBuildsListCall) PageSize(pageSize int64) *ProjectsBuildsListCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken": Token to provide
// to skip to a particular spot in the list.
func (c *ProjectsBuildsListCall) PageToken(pageToken string) *ProjectsBuildsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// RepoName sets the optional parameter "repoName": If present, only
// builds from source repos with the given name will be
// returned.
func (c *ProjectsBuildsListCall) RepoName(repoName string) *ProjectsBuildsListCall {
	c.opt_["repoName"] = repoName
	return c
}

// RevisionId sets the optional parameter "revisionId": If present, only
// builds from source repos, from the given revision ID will
// be
// returned. A revision ID may be either a branch/tag name
// (e.g.,
// "master") or a Git commit SHA.
//
// If a revision ID is specified, a
// repo_name must also be specified.
func (c *ProjectsBuildsListCall) RevisionId(revisionId string) *ProjectsBuildsListCall {
	c.opt_["revisionId"] = revisionId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsBuildsListCall) Fields(s ...googleapi.Field) *ProjectsBuildsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsBuildsListCall) Do() (*ListBuildsResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["pageSize"]; ok {
		params.Set("pageSize", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["repoName"]; ok {
		params.Set("repoName", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["revisionId"]; ok {
		params.Set("revisionId", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/builds")
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
	var ret *ListBuildsResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists previously requested builds.\n\nPreviously requested builds may still be in-progress, or may have finished\nsuccessfully or unsuccessfully.",
	//   "flatPath": "v1/projects/{projectId}/builds",
	//   "httpMethod": "GET",
	//   "id": "cloudbuild.projects.builds.list",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "Number of results to return in the list.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Token to provide to skip to a particular spot in the list.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "The ID of the project.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "repoName": {
	//       "description": "If present, only builds from source repos with the given name will be\nreturned.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "revisionId": {
	//       "description": "If present, only builds from source repos, from the given revision ID will\nbe returned. A revision ID may be either a branch/tag name\n(e.g., \"master\") or a Git commit SHA.\n\nIf a revision ID is specified, a repo_name must also be specified.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/builds",
	//   "response": {
	//     "$ref": "ListBuildsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}
