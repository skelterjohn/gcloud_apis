// Package apikeys provides access to the Google API Keys API.
//
// Usage example:
//
//   import "github.com/skelterjohn/gcloud_apis/clients/apikeys/v1"
//   ...
//   apikeysService, err := apikeys.New(oauthHttpClient)
package apikeys

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

const apiId = "apikeys:v1"
const apiName = "apikeys"
const apiVersion = "v1"
const basePath = "https://apikeys.googleapis.com/"

// OAuth2 scopes used by this API.
const (
	// View and manage your data across Google Cloud Platform services
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"

	// Manage your Google API service configuration
	ServiceManagementScope = "https://www.googleapis.com/auth/service.management"
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
	rs.ApiKeys = NewProjectsApiKeysService(s)
	rs.DeletedApiKeys = NewProjectsDeletedApiKeysService(s)
	return rs
}

type ProjectsService struct {
	s *Service

	ApiKeys *ProjectsApiKeysService

	DeletedApiKeys *ProjectsDeletedApiKeysService
}

func NewProjectsApiKeysService(s *Service) *ProjectsApiKeysService {
	rs := &ProjectsApiKeysService{s: s}
	return rs
}

type ProjectsApiKeysService struct {
	s *Service
}

func NewProjectsDeletedApiKeysService(s *Service) *ProjectsDeletedApiKeysService {
	rs := &ProjectsDeletedApiKeysService{s: s}
	return rs
}

type ProjectsDeletedApiKeysService struct {
	s *Service
}

// AndroidApplication: Identifier of an Android application for API key
// use.
type AndroidApplication struct {
	// PackageName: The package name of the application.
	PackageName string `json:"packageName,omitempty"`

	// Sha1Fingerprint: The 20 byte SHA1 fingerprint of the application.
	Sha1Fingerprint string `json:"sha1Fingerprint,omitempty"`

	// ForceSendFields is a list of field names (e.g. "PackageName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *AndroidApplication) MarshalJSON() ([]byte, error) {
	type noMethod AndroidApplication
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// AndroidKeyDetails: Key details that are specific to android keys.
type AndroidKeyDetails struct {
	// AllowedApplications: A list of Android applications that are allowed
	// to make API calls with
	// this key.
	AllowedApplications []*AndroidApplication `json:"allowedApplications,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AllowedApplications")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *AndroidKeyDetails) MarshalJSON() ([]byte, error) {
	type noMethod AndroidKeyDetails
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// ApiKey: The representation of an API key managed by the `ApiKeys`
// API.
// An API key is used for programmatic access to a project by a service
// account.
type ApiKey struct {
	// AndroidKeyDetails: Key details that are specific to android keys.
	AndroidKeyDetails *AndroidKeyDetails `json:"androidKeyDetails,omitempty"`

	// BrowserKeyDetails: Key details that are specific to browser keys.
	BrowserKeyDetails *BrowserKeyDetails `json:"browserKeyDetails,omitempty"`

	// CreateTime: A timestamp identifying the time this API key was
	// originally created.
	// @OutputOnly
	CreateTime string `json:"createTime,omitempty"`

	// CreatedBy: Email address of the user who originally created this API
	// key.
	// @OutputOnly
	CreatedBy string `json:"createdBy,omitempty"`

	// CurrentKey: An encrypted and signed value held by this API
	// key.
	// @OutputOnly
	CurrentKey string `json:"currentKey,omitempty"`

	// DisplayName: Human-readable display name of this API key.
	// Modifiable by user.
	DisplayName string `json:"displayName,omitempty"`

	// IosKeyDetails: Key details that are specific to iOS keys.
	IosKeyDetails *IosKeyDetails `json:"iosKeyDetails,omitempty"`

	// KeyId: Unique identifier for this ApiKey assigned by the
	// server.
	// @OutputOnly
	KeyId string `json:"keyId,omitempty"`

	// PreviousKey: The value of `current_key` before this API key was
	// regenerated.
	// @OutputOnly
	PreviousKey string `json:"previousKey,omitempty"`

	// PreviousKeyExpireTime: The expiration time for the validity of a
	// `previous_key` value after an
	// API key regeneration.
	// @OutputOnly
	PreviousKeyExpireTime string `json:"previousKeyExpireTime,omitempty"`

	// ServerKeyDetails: Key details that are specific to server keys.
	ServerKeyDetails *ServerKeyDetails `json:"serverKeyDetails,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "AndroidKeyDetails")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *ApiKey) MarshalJSON() ([]byte, error) {
	type noMethod ApiKey
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// BrowserKeyDetails: Key details that are specific to browser keys.
type BrowserKeyDetails struct {
	// AllowedReferrers: A list of regular expressions for the referrer URLs
	// that are allowed when
	// making an API call with this key.
	AllowedReferrers []string `json:"allowedReferrers,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AllowedReferrers") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *BrowserKeyDetails) MarshalJSON() ([]byte, error) {
	type noMethod BrowserKeyDetails
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// DeletedApiKey: Information about a deleted API key.
type DeletedApiKey struct {
	// ApiKey: The API key that was deleted
	// @OutputOnly
	ApiKey *ApiKey `json:"apiKey,omitempty"`

	// DeletionTime: The time at which the key was deleted
	// @OutputOnly
	DeletionTime string `json:"deletionTime,omitempty"`

	// Source: What caused the key to be deleted
	// @OutputOnly
	//
	// Possible values:
	//   "DELETION" - This API Key was deleted via a DeleteApiKey API call.
	//   "REGENERATION" - This API Key was deleted by a RegenerateApiKey API
	// call.
	Source string `json:"source,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ApiKey") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *DeletedApiKey) MarshalJSON() ([]byte, error) {
	type noMethod DeletedApiKey
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

// GetProjectForApiKeyResponse: Response message for
// `GetProjectForApiKey` method.
type GetProjectForApiKeyResponse struct {
	// ProjectNumber: The project number corresponding to the project key in
	// the requests.
	// The project number that owns the API key specified in the request.
	ProjectNumber int64 `json:"projectNumber,omitempty,string"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "ProjectNumber") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *GetProjectForApiKeyResponse) MarshalJSON() ([]byte, error) {
	type noMethod GetProjectForApiKeyResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// IosKeyDetails: Key details that are specific to iOS keys.
type IosKeyDetails struct {
	// AllowedBundleIds: A list of bundle IDs that are allowed when making
	// API calls with this key.
	AllowedBundleIds []string `json:"allowedBundleIds,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AllowedBundleIds") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *IosKeyDetails) MarshalJSON() ([]byte, error) {
	type noMethod IosKeyDetails
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// ListApiKeysResponse: Response message for `ListApiKeys` method.
type ListApiKeysResponse struct {
	// Keys: A list of API keys.
	Keys []*ApiKey `json:"keys,omitempty"`

	// NextPageToken: The pagination token for the next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`

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

func (s *ListApiKeysResponse) MarshalJSON() ([]byte, error) {
	type noMethod ListApiKeysResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// ListDeletedApiKeysResponse: Response message for `ListDeletedApiKeys`
// method.
type ListDeletedApiKeysResponse struct {
	// Keys: A list of deleted API keys.
	Keys []*DeletedApiKey `json:"keys,omitempty"`

	// NextPageToken: The pagination token for the next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`

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

func (s *ListDeletedApiKeysResponse) MarshalJSON() ([]byte, error) {
	type noMethod ListDeletedApiKeysResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// ServerKeyDetails: Key details that are specific to server keys.
type ServerKeyDetails struct {
	// AllowedIps: A list of the caller IP addresses that are allowed when
	// making an API call
	// with this key.
	AllowedIps []string `json:"allowedIps,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AllowedIps") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *ServerKeyDetails) MarshalJSON() ([]byte, error) {
	type noMethod ServerKeyDetails
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// method id "apikeys.projects.getProjectForApiKey":

type ProjectsGetProjectForApiKeyCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// GetProjectForApiKey: Get the project info about an API key.
func (r *ProjectsService) GetProjectForApiKey() *ProjectsGetProjectForApiKeyCall {
	c := &ProjectsGetProjectForApiKeyCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// ApiKey sets the optional parameter "apiKey": Finds the project that
// owns the key with this `current_key` value.
func (c *ProjectsGetProjectForApiKeyCall) ApiKey(apiKey string) *ProjectsGetProjectForApiKeyCall {
	c.urlParams_.Set("apiKey", apiKey)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsGetProjectForApiKeyCall) Fields(s ...googleapi.Field) *ProjectsGetProjectForApiKeyCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsGetProjectForApiKeyCall) IfNoneMatch(entityTag string) *ProjectsGetProjectForApiKeyCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsGetProjectForApiKeyCall) Context(ctx context.Context) *ProjectsGetProjectForApiKeyCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsGetProjectForApiKeyCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects:getProjectForApiKey")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.SetOpaque(req.URL)
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "apikeys.projects.getProjectForApiKey" call.
// Exactly one of *GetProjectForApiKeyResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *GetProjectForApiKeyResponse.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsGetProjectForApiKeyCall) Do(opts ...googleapi.CallOption) (*GetProjectForApiKeyResponse, error) {
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
	ret := &GetProjectForApiKeyResponse{
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
	//   "description": "Get the project info about an API key.",
	//   "flatPath": "v1/projects:getProjectForApiKey",
	//   "httpMethod": "GET",
	//   "id": "apikeys.projects.getProjectForApiKey",
	//   "parameterOrder": [],
	//   "parameters": {
	//     "apiKey": {
	//       "description": "Finds the project that owns the key with this `current_key` value.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects:getProjectForApiKey",
	//   "response": {
	//     "$ref": "GetProjectForApiKeyResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/service.management"
	//   ]
	// }

}

// method id "apikeys.projects.apiKeys.batchDelete":

type ProjectsApiKeysBatchDeleteCall struct {
	s          *Service
	projectId  string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
}

// BatchDelete: Bulk delete a list of API keys.
func (r *ProjectsApiKeysService) BatchDelete(projectId string) *ProjectsApiKeysBatchDeleteCall {
	c := &ProjectsApiKeysBatchDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	return c
}

// KeyIds sets the optional parameter "keyIds": The identifiers for the
// keys to be deleted.
func (c *ProjectsApiKeysBatchDeleteCall) KeyIds(keyIds ...string) *ProjectsApiKeysBatchDeleteCall {
	c.urlParams_.SetMulti("keyIds", append([]string{}, keyIds...))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsApiKeysBatchDeleteCall) Fields(s ...googleapi.Field) *ProjectsApiKeysBatchDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsApiKeysBatchDeleteCall) Context(ctx context.Context) *ProjectsApiKeysBatchDeleteCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsApiKeysBatchDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/apiKeys:batchDelete")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "apikeys.projects.apiKeys.batchDelete" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsApiKeysBatchDeleteCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
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
	//   "description": "Bulk delete a list of API keys.",
	//   "flatPath": "v1/projects/{projectId}/apiKeys:batchDelete",
	//   "httpMethod": "POST",
	//   "id": "apikeys.projects.apiKeys.batchDelete",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "keyIds": {
	//       "description": "The identifiers for the keys to be deleted.",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "The project that owns the API keys.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/apiKeys:batchDelete",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/service.management"
	//   ]
	// }

}

// method id "apikeys.projects.apiKeys.create":

type ProjectsApiKeysCreateCall struct {
	s          *Service
	projectId  string
	apikey     *ApiKey
	urlParams_ gensupport.URLParams
	ctx_       context.Context
}

// Create: Creates a new API key.
func (r *ProjectsApiKeysService) Create(projectId string, apikey *ApiKey) *ProjectsApiKeysCreateCall {
	c := &ProjectsApiKeysCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.apikey = apikey
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsApiKeysCreateCall) Fields(s ...googleapi.Field) *ProjectsApiKeysCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsApiKeysCreateCall) Context(ctx context.Context) *ProjectsApiKeysCreateCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsApiKeysCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.apikey)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/apiKeys")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "apikeys.projects.apiKeys.create" call.
// Exactly one of *ApiKey or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *ApiKey.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsApiKeysCreateCall) Do(opts ...googleapi.CallOption) (*ApiKey, error) {
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
	ret := &ApiKey{
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
	//   "description": "Creates a new API key.",
	//   "flatPath": "v1/projects/{projectId}/apiKeys",
	//   "httpMethod": "POST",
	//   "id": "apikeys.projects.apiKeys.create",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "The project for which this API key will be created.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/apiKeys",
	//   "request": {
	//     "$ref": "ApiKey"
	//   },
	//   "response": {
	//     "$ref": "ApiKey"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/service.management"
	//   ]
	// }

}

// method id "apikeys.projects.apiKeys.delete":

type ProjectsApiKeysDeleteCall struct {
	s          *Service
	projectId  string
	keyId      string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
}

// Delete: Deletes an API key.
func (r *ProjectsApiKeysService) Delete(projectId string, keyId string) *ProjectsApiKeysDeleteCall {
	c := &ProjectsApiKeysDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.keyId = keyId
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsApiKeysDeleteCall) Fields(s ...googleapi.Field) *ProjectsApiKeysDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsApiKeysDeleteCall) Context(ctx context.Context) *ProjectsApiKeysDeleteCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsApiKeysDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/apiKeys/{keyId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"keyId":     c.keyId,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "apikeys.projects.apiKeys.delete" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsApiKeysDeleteCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
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
	//   "description": "Deletes an API key.",
	//   "flatPath": "v1/projects/{projectId}/apiKeys/{keyId}",
	//   "httpMethod": "DELETE",
	//   "id": "apikeys.projects.apiKeys.delete",
	//   "parameterOrder": [
	//     "projectId",
	//     "keyId"
	//   ],
	//   "parameters": {
	//     "keyId": {
	//       "description": "The identifier for the key to be deleted.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "The project that owns the API key.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/apiKeys/{keyId}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/service.management"
	//   ]
	// }

}

// method id "apikeys.projects.apiKeys.get":

type ProjectsApiKeysGetCall struct {
	s            *Service
	projectId    string
	keyId        string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// Get: Gets the metadata for an API key.
func (r *ProjectsApiKeysService) Get(projectId string, keyId string) *ProjectsApiKeysGetCall {
	c := &ProjectsApiKeysGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.keyId = keyId
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsApiKeysGetCall) Fields(s ...googleapi.Field) *ProjectsApiKeysGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsApiKeysGetCall) IfNoneMatch(entityTag string) *ProjectsApiKeysGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsApiKeysGetCall) Context(ctx context.Context) *ProjectsApiKeysGetCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsApiKeysGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/apiKeys/{keyId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"keyId":     c.keyId,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "apikeys.projects.apiKeys.get" call.
// Exactly one of *ApiKey or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *ApiKey.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsApiKeysGetCall) Do(opts ...googleapi.CallOption) (*ApiKey, error) {
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
	ret := &ApiKey{
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
	//   "description": "Gets the metadata for an API key.",
	//   "flatPath": "v1/projects/{projectId}/apiKeys/{keyId}",
	//   "httpMethod": "GET",
	//   "id": "apikeys.projects.apiKeys.get",
	//   "parameterOrder": [
	//     "projectId",
	//     "keyId"
	//   ],
	//   "parameters": {
	//     "keyId": {
	//       "description": "The identifier for the key to be retrieved.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "The project that owns the API key.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/apiKeys/{keyId}",
	//   "response": {
	//     "$ref": "ApiKey"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/service.management"
	//   ]
	// }

}

// method id "apikeys.projects.apiKeys.list":

type ProjectsApiKeysListCall struct {
	s            *Service
	projectId    string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// List: Lists the API keys owned by a project.
func (r *ProjectsApiKeysService) List(projectId string) *ProjectsApiKeysListCall {
	c := &ProjectsApiKeysListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	return c
}

// PageSize sets the optional parameter "pageSize": Specifies the
// maximum number of results to be returned at a time.
func (c *ProjectsApiKeysListCall) PageSize(pageSize int64) *ProjectsApiKeysListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": Requests a
// specific page of results.
func (c *ProjectsApiKeysListCall) PageToken(pageToken string) *ProjectsApiKeysListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsApiKeysListCall) Fields(s ...googleapi.Field) *ProjectsApiKeysListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsApiKeysListCall) IfNoneMatch(entityTag string) *ProjectsApiKeysListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsApiKeysListCall) Context(ctx context.Context) *ProjectsApiKeysListCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsApiKeysListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/apiKeys")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "apikeys.projects.apiKeys.list" call.
// Exactly one of *ListApiKeysResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListApiKeysResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsApiKeysListCall) Do(opts ...googleapi.CallOption) (*ListApiKeysResponse, error) {
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
	ret := &ListApiKeysResponse{
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
	//   "description": "Lists the API keys owned by a project.",
	//   "flatPath": "v1/projects/{projectId}/apiKeys",
	//   "httpMethod": "GET",
	//   "id": "apikeys.projects.apiKeys.list",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "Specifies the maximum number of results to be returned at a time.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Requests a specific page of results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "Lists all API keys associated with this project.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/apiKeys",
	//   "response": {
	//     "$ref": "ListApiKeysResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/service.management"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsApiKeysListCall) Pages(ctx context.Context, f func(*ListApiKeysResponse) error) error {
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

// method id "apikeys.projects.apiKeys.patch":

type ProjectsApiKeysPatchCall struct {
	s          *Service
	projectId  string
	keyId      string
	apikey     *ApiKey
	urlParams_ gensupport.URLParams
	ctx_       context.Context
}

// Patch: Patches the modifiable fields of an API key.
func (r *ProjectsApiKeysService) Patch(projectId string, keyId string, apikey *ApiKey) *ProjectsApiKeysPatchCall {
	c := &ProjectsApiKeysPatchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.keyId = keyId
	c.apikey = apikey
	return c
}

// UpdateMask sets the optional parameter "updateMask": Field mask for
// updates.
func (c *ProjectsApiKeysPatchCall) UpdateMask(updateMask string) *ProjectsApiKeysPatchCall {
	c.urlParams_.Set("updateMask", updateMask)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsApiKeysPatchCall) Fields(s ...googleapi.Field) *ProjectsApiKeysPatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsApiKeysPatchCall) Context(ctx context.Context) *ProjectsApiKeysPatchCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsApiKeysPatchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.apikey)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/apiKeys/{keyId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"keyId":     c.keyId,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "apikeys.projects.apiKeys.patch" call.
// Exactly one of *ApiKey or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *ApiKey.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsApiKeysPatchCall) Do(opts ...googleapi.CallOption) (*ApiKey, error) {
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
	ret := &ApiKey{
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
	//   "description": "Patches the modifiable fields of an API key.",
	//   "flatPath": "v1/projects/{projectId}/apiKeys/{keyId}",
	//   "httpMethod": "PATCH",
	//   "id": "apikeys.projects.apiKeys.patch",
	//   "parameterOrder": [
	//     "projectId",
	//     "keyId"
	//   ],
	//   "parameters": {
	//     "keyId": {
	//       "description": "The identifier for the key to be modified.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "The project that owns the API key.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "updateMask": {
	//       "description": "Field mask for updates.",
	//       "format": "google-fieldmask",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/apiKeys/{keyId}",
	//   "request": {
	//     "$ref": "ApiKey"
	//   },
	//   "response": {
	//     "$ref": "ApiKey"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/service.management"
	//   ]
	// }

}

// method id "apikeys.projects.apiKeys.regenerate":

type ProjectsApiKeysRegenerateCall struct {
	s          *Service
	projectId  string
	keyId      string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
}

// Regenerate: Regenerates the key string for the specified API
// key.
// This writes a new key string to `current_key` and writes the previous
// key
// string to `previous_key`.
// Returns the updated key entry.
func (r *ProjectsApiKeysService) Regenerate(projectId string, keyId string) *ProjectsApiKeysRegenerateCall {
	c := &ProjectsApiKeysRegenerateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.keyId = keyId
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsApiKeysRegenerateCall) Fields(s ...googleapi.Field) *ProjectsApiKeysRegenerateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsApiKeysRegenerateCall) Context(ctx context.Context) *ProjectsApiKeysRegenerateCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsApiKeysRegenerateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/apiKeys/{keyId}:regenerate")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"keyId":     c.keyId,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "apikeys.projects.apiKeys.regenerate" call.
// Exactly one of *ApiKey or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *ApiKey.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsApiKeysRegenerateCall) Do(opts ...googleapi.CallOption) (*ApiKey, error) {
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
	ret := &ApiKey{
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
	//   "description": "Regenerates the key string for the specified API key.\nThis writes a new key string to `current_key` and writes the previous key\nstring to `previous_key`.\nReturns the updated key entry.",
	//   "flatPath": "v1/projects/{projectId}/apiKeys/{keyId}:regenerate",
	//   "httpMethod": "POST",
	//   "id": "apikeys.projects.apiKeys.regenerate",
	//   "parameterOrder": [
	//     "projectId",
	//     "keyId"
	//   ],
	//   "parameters": {
	//     "keyId": {
	//       "description": "The identifier for the key to be regenerated.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "The project that owns the API key.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/apiKeys/{keyId}:regenerate",
	//   "response": {
	//     "$ref": "ApiKey"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/service.management"
	//   ]
	// }

}

// method id "apikeys.projects.apiKeys.revert":

type ProjectsApiKeysRevertCall struct {
	s          *Service
	projectId  string
	keyId      string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
}

// Revert: Reverts a previous key regeneration.
// This swaps the contents of `current_key` and `previous_key`.
// Returns the updated key entry.
func (r *ProjectsApiKeysService) Revert(projectId string, keyId string) *ProjectsApiKeysRevertCall {
	c := &ProjectsApiKeysRevertCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	c.keyId = keyId
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsApiKeysRevertCall) Fields(s ...googleapi.Field) *ProjectsApiKeysRevertCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsApiKeysRevertCall) Context(ctx context.Context) *ProjectsApiKeysRevertCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsApiKeysRevertCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/apiKeys/{keyId}:revert")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"keyId":     c.keyId,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "apikeys.projects.apiKeys.revert" call.
// Exactly one of *ApiKey or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *ApiKey.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsApiKeysRevertCall) Do(opts ...googleapi.CallOption) (*ApiKey, error) {
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
	ret := &ApiKey{
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
	//   "description": "Reverts a previous key regeneration.\nThis swaps the contents of `current_key` and `previous_key`.\nReturns the updated key entry.",
	//   "flatPath": "v1/projects/{projectId}/apiKeys/{keyId}:revert",
	//   "httpMethod": "POST",
	//   "id": "apikeys.projects.apiKeys.revert",
	//   "parameterOrder": [
	//     "projectId",
	//     "keyId"
	//   ],
	//   "parameters": {
	//     "keyId": {
	//       "description": "The identifier for the key to be reverted.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "The project that owns the API key.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/apiKeys/{keyId}:revert",
	//   "response": {
	//     "$ref": "ApiKey"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/service.management"
	//   ]
	// }

}

// method id "apikeys.projects.deletedApiKeys.list":

type ProjectsDeletedApiKeysListCall struct {
	s            *Service
	projectId    string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// List: Lists the deleted API keys owned by a project.
func (r *ProjectsDeletedApiKeysService) List(projectId string) *ProjectsDeletedApiKeysListCall {
	c := &ProjectsDeletedApiKeysListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectId = projectId
	return c
}

// PageSize sets the optional parameter "pageSize": Specifies the
// maximum number of results to be returned at a time.
func (c *ProjectsDeletedApiKeysListCall) PageSize(pageSize int64) *ProjectsDeletedApiKeysListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": Requests a
// specific page of results.
func (c *ProjectsDeletedApiKeysListCall) PageToken(pageToken string) *ProjectsDeletedApiKeysListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsDeletedApiKeysListCall) Fields(s ...googleapi.Field) *ProjectsDeletedApiKeysListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsDeletedApiKeysListCall) IfNoneMatch(entityTag string) *ProjectsDeletedApiKeysListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsDeletedApiKeysListCall) Context(ctx context.Context) *ProjectsDeletedApiKeysListCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsDeletedApiKeysListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/deletedApiKeys")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "apikeys.projects.deletedApiKeys.list" call.
// Exactly one of *ListDeletedApiKeysResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *ListDeletedApiKeysResponse.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsDeletedApiKeysListCall) Do(opts ...googleapi.CallOption) (*ListDeletedApiKeysResponse, error) {
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
	ret := &ListDeletedApiKeysResponse{
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
	//   "description": "Lists the deleted API keys owned by a project.",
	//   "flatPath": "v1/projects/{projectId}/deletedApiKeys",
	//   "httpMethod": "GET",
	//   "id": "apikeys.projects.deletedApiKeys.list",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "Specifies the maximum number of results to be returned at a time.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Requests a specific page of results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "Lists all deleted API keys associated with this project.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/deletedApiKeys",
	//   "response": {
	//     "$ref": "ListDeletedApiKeysResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/service.management"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsDeletedApiKeysListCall) Pages(ctx context.Context, f func(*ListDeletedApiKeysResponse) error) error {
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
