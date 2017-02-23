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

// SearchResponse: Response message for `resources.search`.
type SearchResponse struct {
	// MatchedCount: The approximate total number of resources that match
	// the query.  It will
	// never be less than the number of resources returned so far, but it
	// can
	// change as additional pages of results are returned.
	MatchedCount int64 `json:"matchedCount,omitempty,string"`

	// NextPageToken: If there are more results than those appearing in this
	// response, then
	// `next_page_token` is included.  To get the next set of results, call
	// this
	// method again using the value of `next_page_token` as `page_token`.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Results: A list of resources that match the search query.
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

// SearchResult: A single Google Cloud Platform resource.
type SearchResult struct {
	// DiscoveryType: The JSON schema name listed in the discovery
	// document.
	// Example: `Project`.
	DiscoveryType string `json:"discoveryType,omitempty"`

	// DiscoveryUrl: The URL of the discovery document containing the
	// resource's JSON schema.
	// Example:
	// `https://cloudresourcemanager.googleapis.com/$discovery/rest`.
	DiscoveryUrl string `json:"discoveryUrl,omitempty"`

	// Resource: The matched resource, expressed as a JSON object.
	Resource SearchResultResource `json:"resource,omitempty"`

	// ResourceName: The RPC resource name: a scheme-less URI that includes
	// the DNS-compatible
	// API service name. The URI does not include an API version and does
	// not
	// support %-encoding.
	// Example:
	// `//cloudresourcemanager.googleapis.com/projects/my-project-123`.
	ResourceName string `json:"resourceName,omitempty"`

	// ResourceType: A domain-scoped name that describes the protocol buffer
	// message type.
	// Example:
	// `type.googleapis.com/google.cloud.resourcemanager.v1.Project`.
	ResourceType string `json:"resourceType,omitempty"`

	// ResourceUrl: The REST URL for accessing the resource. An HTTP GET
	// operation using this
	// URL returns an object equivalent to the value in the `resource`
	// field.
	// Example:
	// `https://cloudresourcemanager.googleapis.com/v1/projec
	// ts/my-project-123`.
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

// Search: Lists accessible Google Cloud Platform resources that match a
// query. A
// resource is accessible to the caller if the caller has permission
// to perform a GET operation on the resource.
func (r *ResourcesService) Search() *ResourcesSearchCall {
	c := &ResourcesSearchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// OrderBy sets the optional parameter "orderBy": A comma-separated list
// of string-valued fields for sorting the
// results.  If this field is omitted, then the order of results is
// not
// defined. You can use fields from the resource schemas as well as
// the
// built-in fields `resourceName` and `resourceType`. Field values are
// ordered
// by their UTF-8 encodings.
//
// Fields are sorted in ascending order by default. To sort a field in
// descending
// order, append " desc" to the field name. For example, the
// `order_by`
// value "resource_type desc,resource_name" sorts results by resource
// type
// in descending order; resources with the same type are returned in
// ascending
// order of their names.
func (c *ResourcesSearchCall) OrderBy(orderBy string) *ResourcesSearchCall {
	c.urlParams_.Set("orderBy", orderBy)
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of resources to return from this request.  The
// presence of `next_page_token` in the response indicates that more
// resources
// are available.  The default value of `page_size` is 20 and the
// maximum
// value is 1000.
func (c *ResourcesSearchCall) PageSize(pageSize int64) *ResourcesSearchCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": If present, then
// retrieve the next batch of results from the
// preceding call to this method.  `page_token` must be the value
// of
// `next_page_token` from the previous response.  The values of other
// method
// parameters, including the query and sort order, must be identical to
// those
// in the previous call.
func (c *ResourcesSearchCall) PageToken(pageToken string) *ResourcesSearchCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Query sets the optional parameter "query": The query string. If the
// query is missing or empty,
// all accessible resources are returned.
//
// Any field in a supported resource type's schema may be specified
// in
// the query. Additionally, every resource has a `@type` field whose
// value is
// the resource's type URL. See `SearchResult.resource_type` for
// more
// information.
//
// Example: The following query searches for accessible Compute Engine
// VM
// instances (`@type:Instance`) that have an `env` label value of `prod`
// and
// that have a machine type that starts with "n1-stand":
//
//     @type:Instance labels.env:prod machineType:n1-stand*
//
// For more information, see [Search
// Queries](/resource-search/docs/search-queries)
// and [Resource Types](/resource-search/docs/reference/Resource.Types).
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
	//   "description": "Lists accessible Google Cloud Platform resources that match a query. A\nresource is accessible to the caller if the caller has permission\nto perform a GET operation on the resource.",
	//   "flatPath": "v1/resources:search",
	//   "httpMethod": "GET",
	//   "id": "cloudresourcesearch.resources.search",
	//   "parameterOrder": [],
	//   "parameters": {
	//     "orderBy": {
	//       "description": "Optional. A comma-separated list of string-valued fields for sorting the\nresults.  If this field is omitted, then the order of results is not\ndefined. You can use fields from the resource schemas as well as the\nbuilt-in fields `resourceName` and `resourceType`. Field values are ordered\nby their UTF-8 encodings.\n\nFields are sorted in ascending order by default. To sort a field in descending\norder, append `\" desc\"` to the field name. For example, the `order_by`\nvalue `\"resource_type desc,resource_name\"` sorts results by resource type\nin descending order; resources with the same type are returned in ascending\norder of their names.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Optional. The maximum number of resources to return from this request.  The\npresence of `next_page_token` in the response indicates that more resources\nare available.  The default value of `page_size` is 20 and the maximum\nvalue is 1000.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. If present, then retrieve the next batch of results from the\npreceding call to this method.  `page_token` must be the value of\n`next_page_token` from the previous response.  The values of other method\nparameters, including the query and sort order, must be identical to those\nin the previous call.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "query": {
	//       "description": "Optional. The query string. If the query is missing or empty,\nall accessible resources are returned.\n\nAny field in a supported resource type's schema may be specified in\nthe query. Additionally, every resource has a `@type` field whose value is\nthe resource's type URL. See `SearchResult.resource_type` for more\ninformation.\n\nExample: The following query searches for accessible Compute Engine VM\ninstances (`@type:Instance`) that have an `env` label value of `prod` and\nthat have a machine type that starts with `\"n1-stand\"`:\n\n    @type:Instance labels.env:prod machineType:n1-stand*\n\nFor more information, see [Search Queries](/resource-search/docs/search-queries)\nand [Resource Types](/resource-search/docs/reference/Resource.Types).",
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
