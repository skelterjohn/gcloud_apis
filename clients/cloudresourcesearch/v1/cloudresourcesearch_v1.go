// Package cloudresourcesearch provides access to the Google Cloud Resource Search.
//
// See https://groups.google.com/forum/#!forum/resource-search-api-alpha
//
// Usage example:
//
//   import "github.com/skelterjohn/gcloud_apis/clients/cloudresourcesearch/v1"
//   ...
//   cloudresourcesearchService, err := cloudresourcesearch.New(oauthHttpClient)
package cloudresourcesearch

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

const apiId = "cloudresourcesearch:v1"
const apiName = "cloudresourcesearch"
const apiVersion = "v1"
const basePath = "https://cloudresourcesearch.googleapis.com/"

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
	s.Resources = NewResourcesService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Resources *ResourcesService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewResourcesService(s *Service) *ResourcesService {
	rs := &ResourcesService{s: s}
	return rs
}

type ResourcesService struct {
	s *Service
}

// BillingAccount: A billing account in [Google
// Cloud
// Console](https://console.cloud.google.com/). You can assign a billing
// account
// to one or more projects.
type BillingAccount struct {
	// DisplayName: The display name given to the billing account, such as
	// `My Billing
	// Account`. This name is displayed in the Google Cloud Console.
	DisplayName string `json:"displayName,omitempty"`

	// Name: The resource name of the billing account. The resource name has
	// the form
	// `billingAccounts/{billing_account_id}`. For
	// example,
	// `billingAccounts/012345-567890-ABCDEF` would be the resource name
	// for
	// billing account `012345-567890-ABCDEF`.
	Name string `json:"name,omitempty"`

	// Open: True if the billing account is open, and will therefore be
	// charged for any
	// usage on associated projects. False if the billing account is closed,
	// and
	// therefore projects associated with it will be unable to use paid
	// services.
	Open bool `json:"open,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DisplayName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *BillingAccount) MarshalJSON() ([]byte, error) {
	type noMethod BillingAccount
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// Organization: The root node in the resource hierarchy to which a
// particular entity's
// (e.g., company) resources belong.
type Organization struct {
	// CreationTime: Timestamp when the Organization was created. Assigned
	// by the server.
	// @OutputOnly
	CreationTime string `json:"creationTime,omitempty"`

	// DisplayName: A friendly string to be used to refer to the
	// Organization in the UI.
	// Assigned by the server, set to the firm name of the Google For
	// Work
	// customer that owns this organization.
	// @OutputOnly
	DisplayName string `json:"displayName,omitempty"`

	// LifecycleState: The organization's current lifecycle state. Assigned
	// by the server.
	// @OutputOnly
	//
	// Possible values:
	//   "LIFECYCLE_STATE_UNSPECIFIED" - Unspecified state.  This is only
	// useful for distinguishing unset values.
	//   "ACTIVE" - The normal and active state.
	//   "DELETE_REQUESTED" - The organization has been marked for deletion
	// by the user.
	LifecycleState string `json:"lifecycleState,omitempty"`

	// Name: Output Only. The resource name of the organization. This is
	// the
	// organization's relative path in the API. Its format
	// is
	// "organizations/[organization_id]". For example, "organizations/1234".
	Name string `json:"name,omitempty"`

	// OrganizationId: An immutable id for the Organization that is assigned
	// on creation. This
	// should be omitted when creating a new Organization.
	// This field is read-only.
	// This field is deprecated and will be removed in v1. Use name instead.
	OrganizationId string `json:"organizationId,omitempty"`

	// Owner: The owner of this Organization. The owner should be specified
	// on
	// creation. Once set, it cannot be changed.
	// This field is required.
	Owner *OrganizationOwner `json:"owner,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CreationTime") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Organization) MarshalJSON() ([]byte, error) {
	type noMethod Organization
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// OrganizationOwner: The entity that owns an Organization. The lifetime
// of the Organization and
// all of its descendants are bound to the `OrganizationOwner`. If
// the
// `OrganizationOwner` is deleted, the Organization and all its
// descendants will
// be deleted.
type OrganizationOwner struct {
	// DirectoryCustomerId: The Google for Work customer id used in the
	// Directory API.
	DirectoryCustomerId string `json:"directoryCustomerId,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DirectoryCustomerId")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *OrganizationOwner) MarshalJSON() ([]byte, error) {
	type noMethod OrganizationOwner
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// Project: A Project is a high-level Google Cloud Platform entity.  It
// is a
// container for ACLs, APIs, AppEngine Apps, VMs, and other
// Google Cloud Platform resources.
type Project struct {
	// CreateTime: Creation time.
	//
	// Read-only.
	CreateTime string `json:"createTime,omitempty"`

	// Labels: The labels associated with this Project.
	//
	// Label keys must be between 1 and 63 characters long and must
	// conform
	// to the following regular expression:
	// \[a-z\](\[-a-z0-9\]*\[a-z0-9\])?.
	//
	// Label values must be between 0 and 63 characters long and must
	// conform
	// to the regular expression (\[a-z\](\[-a-z0-9\]*\[a-z0-9\])?)?.
	//
	// No more than 256 labels can be associated with a given
	// resource.
	//
	// Clients should store labels in a representation such as JSON that
	// does not
	// depend on specific characters being disallowed.
	//
	// Example: <code>"environment" : "dev"</code>
	//
	// Read-write.
	Labels map[string]string `json:"labels,omitempty"`

	// LifecycleState: The Project lifecycle state.
	//
	// Read-only.
	//
	// Possible values:
	//   "LIFECYCLE_STATE_UNSPECIFIED" - Unspecified state.  This is only
	// used/useful for distinguishing
	// unset values.
	//   "ACTIVE" - The normal and active state.
	//   "DELETE_REQUESTED" - The project has been marked for deletion by
	// the user
	// (by invoking DeleteProject)
	// or by the system (Google Cloud Platform).
	// This can generally be reversed by invoking UndeleteProject.
	//   "DELETE_IN_PROGRESS" - This lifecycle state is no longer used and
	// is not returned by the API.
	LifecycleState string `json:"lifecycleState,omitempty"`

	// Name: The user-assigned display name of the Project.
	// It must be 4 to 30 characters.
	// Allowed characters are: lowercase and uppercase letters,
	// numbers,
	// hyphen, single-quote, double-quote, space, and exclamation
	// point.
	//
	// Example: <code>My Project</code>
	//
	// Read-write.
	Name string `json:"name,omitempty"`

	// Parent: An optional reference to a parent Resource.
	//
	// The only supported parent type is "organization". Once set, the
	// parent
	// cannot be modified. The `parent` can be set on creation or using
	// the
	// `UpdateProject` method; the end user must have
	// the
	// `resourcemanager.projects.create` permission on the
	// parent.
	//
	// Read-write.
	Parent *ResourceId `json:"parent,omitempty"`

	// ProjectId: The unique, user-assigned ID of the Project.
	// It must be 6 to 30 lowercase letters, digits, or hyphens.
	// It must start with a letter.
	// Trailing hyphens are prohibited.
	//
	// Example: <code>tokyo-rain-123</code>
	//
	// Read-only after creation.
	ProjectId string `json:"projectId,omitempty"`

	// ProjectNumber: The number uniquely identifying the project.
	//
	// Example: <code>415104041262</code>
	//
	// Read-only.
	ProjectNumber int64 `json:"projectNumber,omitempty,string"`

	// ForceSendFields is a list of field names (e.g. "CreateTime") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Project) MarshalJSON() ([]byte, error) {
	type noMethod Project
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// ResourceId: A container to reference an id for any resource type. A
// `resource` in Google
// Cloud Platform is a generic term for something you (a developer) may
// want to
// interact with through one of our API's. Some examples are an
// AppEngine app,
// a Compute Engine instance, a Cloud SQL database, and so on.
type ResourceId struct {
	// Id: Required field for the type-specific id. This should correspond
	// to the id
	// used in the type-specific API's.
	Id string `json:"id,omitempty"`

	// Type: Required field representing the resource type this id is
	// for.
	// At present, the valid types are "project" and "organization".
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Id") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *ResourceId) MarshalJSON() ([]byte, error) {
	type noMethod ResourceId
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// SearchResponse: Response message for Search().
type SearchResponse struct {
	// MatchedCount: The approximate number of documents that match the
	// query. It is greater
	// than or equal to the number of documents actually returned.
	MatchedCount int64 `json:"matchedCount,omitempty,string"`

	// NextPageToken: If there are more results, retrieve them by invoking
	// search call with the
	// same arguments and this `nextPageToken`. If there are no more
	// results, this
	// field is not set.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Results: The list of resources that match the search query.
	Results []*SearchResult `json:"results,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "MatchedCount") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *SearchResponse) MarshalJSON() ([]byte, error) {
	type noMethod SearchResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// SearchResult: A single Google Cloud Platform resource returned in
// SearchResourcesResponse.
type SearchResult struct {
	// DiscoveryType: The JSON schema name listed in the discovery
	// document.
	// Example: Project
	DiscoveryType string `json:"discoveryType,omitempty"`

	// DiscoveryUrl: The URL of the discovery document containing the
	// resource's JSON schema.
	// Example: https://cloudresourcemanager.googleapis.com/$discovery/rest
	DiscoveryUrl string `json:"discoveryUrl,omitempty"`

	// Resource: The matched resource, expressed as a JSON object.
	Resource SearchResultResource `json:"resource,omitempty"`

	// ResourceName: The RPC resource name. It is a scheme-less URI that
	// includes the DNS-
	// compatible API service name. It does not include API version, or
	// support
	// %-encoding.
	// Example:
	// //cloudresourcemanager.googleapis.com/projects/my-project-123
	ResourceName string `json:"resourceName,omitempty"`

	// ResourceType: A domain-scoped name that describes the protocol buffer
	// message type.
	// Example: type.googleapis.com/google.cloud.resourcemanager.v1.Project
	ResourceType string `json:"resourceType,omitempty"`

	// ResourceUrl: The REST URL for accessing the resource. HTTP GET on the
	// `resource_url`
	// would return a JSON object equivalent to the `resource`
	// below.
	// Example:
	// https://cloudresourcemanager.googleapis.com/v1/projects/my-project-123
	ResourceUrl string `json:"resourceUrl,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DiscoveryType") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *SearchResult) MarshalJSON() ([]byte, error) {
	type noMethod SearchResult
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

type SearchResultResource interface{}

// method id "cloudresourcesearch.resources.search":

type ResourcesSearchCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// Search: Lists accessible Google Cloud Platform resources that match
// the query. A
// resource is accessible to the caller if they have the IAM .get
// permission
// for it.
func (r *ResourcesService) Search() *ResourcesSearchCall {
	c := &ResourcesSearchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// OrderBy sets the optional parameter "orderBy": Comma-separated list
// of string fields for sorting on the search result,
// including fields from the resources and the built-in fields
// (`resourceName`
// and `resourceType`). Strings are sorted as binary strings based on
// their
// UTF-8 encoding.
//
// The default sorting order is ascending. To specify descending order
// for a
// field, a suffix " desc" should be appended to the field name.
// For
// example: `orderBy="foo desc,bar".
func (c *ResourcesSearchCall) OrderBy(orderBy string) *ResourcesSearchCall {
	c.urlParams_.Set("orderBy", orderBy)
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of search results to return per page. Searches perform
// best when the `pageSize` is kept as small as possible. If not
// specified, 20
// results are returned per page. At most 1000 results are returned per
// page.
func (c *ResourcesSearchCall) PageSize(pageSize int64) *ResourcesSearchCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A `nextPageToken`
// returned from previous SearchResources call as the
// starting point for this call.
func (c *ResourcesSearchCall) PageToken(pageToken string) *ResourcesSearchCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Query sets the optional parameter "query": The query string in search
// query syntax. If the query is missing or empty,
// all resources are returned.
//
// Any field in a supported resource type's JSON schema may be specified
// in
// the query. Additionally, every resource has a `@type` field whose
// value is
// the resource's type URL. See `SearchResult.resource_type` for
// more
// information.
//
// Example: The following query searches for all Google Compute Engine
// VM
// instances accessible to the caller. The query is further restricted
// on the
// `labels` and `machineType` fields of the resource. Only VM instances
// with
// the label `env` set to "prod" and `machineType` including a token
// phrase
// with the prefix "n1-stand" are matched.
//   @type:Instance labels.env:prod machineType:n1-stand*
func (c *ResourcesSearchCall) Query(query string) *ResourcesSearchCall {
	c.urlParams_.Set("query", query)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ResourcesSearchCall) Fields(s ...googleapi.Field) *ResourcesSearchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ResourcesSearchCall) IfNoneMatch(entityTag string) *ResourcesSearchCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ResourcesSearchCall) Context(ctx context.Context) *ResourcesSearchCall {
	c.ctx_ = ctx
	return c
}

func (c *ResourcesSearchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/resources:search")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.SetOpaque(req.URL)
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "cloudresourcesearch.resources.search" call.
// Exactly one of *SearchResponse or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *SearchResponse.ServerResponse.Header or (if a response was returned
// at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ResourcesSearchCall) Do(opts ...googleapi.CallOption) (*SearchResponse, error) {
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
	ret := &SearchResponse{
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
	//   "description": "Lists accessible Google Cloud Platform resources that match the query. A\nresource is accessible to the caller if they have the IAM .get permission\nfor it.",
	//   "flatPath": "v1/resources:search",
	//   "httpMethod": "GET",
	//   "id": "cloudresourcesearch.resources.search",
	//   "parameterOrder": [],
	//   "parameters": {
	//     "orderBy": {
	//       "description": "Comma-separated list of string fields for sorting on the search result,\nincluding fields from the resources and the built-in fields (`resourceName`\nand `resourceType`). Strings are sorted as binary strings based on their\nUTF-8 encoding.\n\nThe default sorting order is ascending. To specify descending order for a\nfield, a suffix `\" desc\"` should be appended to the field name. For\nexample: `orderBy=\"foo desc,bar\"`.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "The maximum number of search results to return per page. Searches perform\nbest when the `pageSize` is kept as small as possible. If not specified, 20\nresults are returned per page. At most 1000 results are returned per page.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A `nextPageToken` returned from previous SearchResources call as the\nstarting point for this call.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "query": {
	//       "description": "The query string in search query syntax. If the query is missing or empty,\nall resources are returned.\n\nAny field in a supported resource type's JSON schema may be specified in\nthe query. Additionally, every resource has a `@type` field whose value is\nthe resource's type URL. See `SearchResult.resource_type` for more\ninformation.\n\nExample: The following query searches for all Google Compute Engine VM\ninstances accessible to the caller. The query is further restricted on the\n`labels` and `machineType` fields of the resource. Only VM instances with\nthe label `env` set to \"prod\" and `machineType` including a token phrase\nwith the prefix \"n1-stand\" are matched.\n  @type:Instance labels.env:prod machineType:n1-stand*",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/resources:search",
	//   "response": {
	//     "$ref": "SearchResponse"
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
func (c *ResourcesSearchCall) Pages(ctx context.Context, f func(*SearchResponse) error) error {
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
