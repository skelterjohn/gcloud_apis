// Package cloudresourcemanager provides access to the Google Cloud Resource Manager API.
//
// See https://cloud.google.com/resource-manager
//
// Usage example:
//
//   import "google.golang.org/api/cloudresourcemanager/v1beta1"
//   ...
//   cloudresourcemanagerService, err := cloudresourcemanager.New(oauthHttpClient)
package cloudresourcemanager

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

const apiId = "cloudresourcemanager:v1beta1"
const apiName = "cloudresourcemanager"
const apiVersion = "v1beta1"
const basePath = "https://cloudresourcemanager.googleapis.com/"

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
	s.Organizations = NewOrganizationsService(s)
	s.Projects = NewProjectsService(s)
	return s, nil
}

type Service struct {
	client   *http.Client
	BasePath string // API endpoint base URL

	Organizations *OrganizationsService

	Projects *ProjectsService
}

func NewOrganizationsService(s *Service) *OrganizationsService {
	rs := &OrganizationsService{s: s}
	return rs
}

type OrganizationsService struct {
	s *Service
}

func NewProjectsService(s *Service) *ProjectsService {
	rs := &ProjectsService{s: s}
	return rs
}

type ProjectsService struct {
	s *Service
}

type Binding struct {
	// Members: Specifies the identities requesting access for a Cloud
	// Platform resource.
	// `members` can have the following formats:
	//
	// *
	// `allUsers`: A special identifier that represents anyone who is
	//    on
	// the internet; with or without a Google account.
	//
	// *
	// `allAuthenticatedUsers`: A special identifier that represents anyone
	//
	//   who is authenticated with a Google account or a service account.
	//
	// *
	// `user:{emailid}`: An email address that represents a specific Google
	//
	//   account. For example, `alice@gmail.com` or `joe@example.com`.
	//
	// *
	// `serviceAccount:{emailid}`: An email address that represents a
	// service
	//    account. For example,
	// `my-other-app@appspot.gserviceaccount.com`.
	//
	// * `group:{emailid}`: An
	// email address that represents a Google group.
	//    For example,
	// `admins@example.com`.
	//
	// * `domain:{domain}`: A Google Apps domain name
	// that represents all the
	//    users of that domain. For example,
	// `google.com` or `example.com`.
	//
	Members []string `json:"members,omitempty"`

	// Role: Role that is assigned to `members`.
	// For example,
	// `roles/viewer`, `roles/editor`, or `roles/owner`.
	// Required
	Role string `json:"role,omitempty"`
}

type Empty struct {
}

type GetIamPolicyRequest struct {
}

type ListOrganizationsResponse struct {
	// NextPageToken: A pagination token to be used to retrieve the next
	// page of results. If the
	// result is too large to fit within the page
	// size specified in the request,
	// this field will be set with a token
	// that can be used to fetch the next page
	// of results. If this field is
	// empty, it indicates that this response
	// contains the last page of
	// results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Organizations: The list of Organizations that matched the list query,
	// possibly paginated.
	Organizations []*Organization `json:"organizations,omitempty"`
}

type ListProjectsResponse struct {
	// NextPageToken: Pagination token.
	//
	// If the result set is too large to
	// fit in a single response, this token
	// is returned. It encodes the
	// position of the current result cursor.
	// Feeding this value into a new
	// list request with the `page_token` parameter
	// gives the next page of
	// the results.
	//
	// When `next_page_token` is not filled in, there is no
	// next page and
	// the list returned is the last page in the result
	// set.
	//
	// Pagination tokens have a limited lifetime.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Projects: The list of projects that matched the list filter. This
	// list can
	// be paginated.
	Projects []*Project `json:"projects,omitempty"`
}

type Organization struct {
	// CreationTime: Timestamp when the Organization was created. Assigned
	// by the server.
	// @OutputOnly
	CreationTime string `json:"creationTime,omitempty"`

	// DisplayName: A friendly string to be used to refer to the
	// Organization in the UI.
	// This field is required.
	DisplayName string `json:"displayName,omitempty"`

	// OrganizationId: An immutable id for the Organization that is assigned
	// on creation. This
	// should be omitted when creating a new
	// Organization.
	// This field is read-only.
	OrganizationId string `json:"organizationId,omitempty"`

	// Owner: The owner of this Organization. The owner should be specified
	// upon
	// creation. Once set, it cannot be changed.
	// This field is
	// required.
	Owner *OrganizationOwner `json:"owner,omitempty"`
}

type OrganizationOwner struct {
	// DirectoryCustomerId: The Google for Work customer id used in the
	// Directory API.
	DirectoryCustomerId string `json:"directoryCustomerId,omitempty"`
}

type Policy struct {
	// Bindings: Associates a list of `members` to a `role`.
	// Multiple
	// `bindings` must not be specified for the same `role`.
	// `bindings` with
	// no members will result in an error.
	Bindings []*Binding `json:"bindings,omitempty"`

	// Etag: The etag is used for optimistic concurrency control as a way to
	// help
	// prevent simultaneous updates of a policy from overwriting each
	// other.
	// It is strongly suggested that systems make use of the etag in
	// the
	// read-modify-write cycle to perform policy updates in order to
	// avoid race
	// conditions: Etags are returned in the response to
	// GetIamPolicy, and
	// systems are expected to put that etag in the
	// request to SetIamPolicy to
	// ensure that their change will be applied
	// to the same version of the policy.
	//
	// If no etag is provided in the
	// call to SetIamPolicy, then the existing policy
	// is overwritten
	// blindly.
	Etag string `json:"etag,omitempty"`

	// Version: Version of the `Policy`. The default version is 0.
	// 0 =
	// resourcemanager_projects only support legacy roles.
	// 1 = supports
	// non-legacy roles
	// 2 = supports AuditConfig
	Version int64 `json:"version,omitempty"`
}

type Project struct {
	// CreateTime: Creation time.
	//
	// Read-only.
	CreateTime string `json:"createTime,omitempty"`

	// Labels: The labels associated with this project.
	//
	// Label keys must be
	// between 1 and 63 characters long and must conform
	// to the following
	// regular expression: \[a-z\](\[-a-z0-9\]*\[a-z0-9\])?.
	//
	// Label values
	// must be between 0 and 63 characters long and must conform
	// to the
	// regular expression (\[a-z\](\[-a-z0-9\]*\[a-z0-9\])?)?.
	//
	// No more than
	// 256 labels can be associated with a given resource.
	//
	// Clients should
	// store labels in a representation such as JSON that does not
	// depend on
	// specific characters being disallowed.
	//
	// Example: <code>"environment" :
	// "dev"</code>
	//
	// Read-write.
	Labels map[string]string `json:"labels,omitempty"`

	// LifecycleState: The project lifecycle state.
	//
	// Read-only.
	LifecycleState string `json:"lifecycleState,omitempty"`

	// Name: The user-assigned name of the project.
	// It must be 4 to 30
	// characters.
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
	// The only
	// supported parent type is "organization". Once set, the parent
	// cannot
	// be modified.
	//
	// Read-write.
	Parent *ResourceId `json:"parent,omitempty"`

	// ProjectId: The unique, user-assigned ID of the project.
	// It must be 6
	// to 30 lowercase letters, digits, or hyphens.
	// It must start with a
	// letter.
	// Trailing hyphens are prohibited.
	//
	// Example:
	// <code>tokyo-rain-123</code>
	//
	// Read-only after creation.
	ProjectId string `json:"projectId,omitempty"`

	// ProjectNumber: The number uniquely identifying the project.
	//
	// Example:
	// <code>415104041262</code>
	//
	// Read-only.
	ProjectNumber int64 `json:"projectNumber,omitempty,string"`
}

type ResourceId struct {
	// Id: Required field for the type-specific id. This should correspond
	// to the id
	// used in the type-specific API's.
	Id string `json:"id,omitempty"`

	// Type: Required field representing the resource type this id is
	// for.
	// At present, the only valid type is "organization".
	Type string `json:"type,omitempty"`
}

type SetIamPolicyRequest struct {
	// Policy: REQUIRED: The complete policy to be applied to the
	// `resource`. The size of
	// the policy is limited to a few 10s of KB. An
	// empty policy is a
	// valid policy but certain Cloud Platform services
	// (such as Projects)
	// might reject them.
	Policy *Policy `json:"policy,omitempty"`
}

type TestIamPermissionsRequest struct {
	// Permissions: The set of permissions to check for the `resource`.
	// Permissions with
	// wildcards (such as '*' or 'storage.*') are not
	// allowed.
	Permissions []string `json:"permissions,omitempty"`
}

type TestIamPermissionsResponse struct {
	// Permissions: A subset of `TestPermissionsRequest.permissions` that
	// the caller is
	// allowed.
	Permissions []string `json:"permissions,omitempty"`
}

// method id "cloudresourcemanager.organizations.get":

type OrganizationsGetCall struct {
	s              *Service
	organizationId string
	opt_           map[string]interface{}
}

// Get: Fetches an Organization resource by id.
func (r *OrganizationsService) Get(organizationId string) *OrganizationsGetCall {
	c := &OrganizationsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.organizationId = organizationId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OrganizationsGetCall) Fields(s ...googleapi.Field) *OrganizationsGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *OrganizationsGetCall) Do() (*Organization, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/organizations/{organizationId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"organizationId": c.organizationId,
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
	var ret *Organization
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Fetches an Organization resource by id.",
	//   "flatPath": "v1beta1/organizations/{organizationId}",
	//   "httpMethod": "GET",
	//   "id": "cloudresourcemanager.organizations.get",
	//   "parameterOrder": [
	//     "organizationId"
	//   ],
	//   "parameters": {
	//     "organizationId": {
	//       "description": "The id of the Organization resource to fetch.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/organizations/{organizationId}",
	//   "response": {
	//     "$ref": "Organization"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only"
	//   ]
	// }

}

// method id "cloudresourcemanager.organizations.getIamPolicy":

type OrganizationsGetIamPolicyCall struct {
	s                   *Service
	resource            string
	getiampolicyrequest *GetIamPolicyRequest
	opt_                map[string]interface{}
}

// GetIamPolicy: Gets the access control policy for a Organization
// resource. May be empty if
// no such policy or resource exists.
func (r *OrganizationsService) GetIamPolicy(resource string, getiampolicyrequest *GetIamPolicyRequest) *OrganizationsGetIamPolicyCall {
	c := &OrganizationsGetIamPolicyCall{s: r.s, opt_: make(map[string]interface{})}
	c.resource = resource
	c.getiampolicyrequest = getiampolicyrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OrganizationsGetIamPolicyCall) Fields(s ...googleapi.Field) *OrganizationsGetIamPolicyCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *OrganizationsGetIamPolicyCall) Do() (*Policy, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.getiampolicyrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/organizations/{resource}:getIamPolicy")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
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
	var ret *Policy
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets the access control policy for a Organization resource. May be empty if\nno such policy or resource exists.",
	//   "flatPath": "v1beta1/organizations/{resource}:getIamPolicy",
	//   "httpMethod": "POST",
	//   "id": "cloudresourcemanager.organizations.getIamPolicy",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "REQUIRED: The resource for which policy is being requested.\n`resource` is usually specified as a path, such as,\n`projects/{project}/zones/{zone}/disks/{disk}`.\n\nThe format for the path specified in this value is resource specific and\nis specified in the documentation for the respective GetIamPolicy rpc.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/organizations/{resource}:getIamPolicy",
	//   "request": {
	//     "$ref": "GetIamPolicyRequest"
	//   },
	//   "response": {
	//     "$ref": "Policy"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only"
	//   ]
	// }

}

// method id "cloudresourcemanager.organizations.list":

type OrganizationsListCall struct {
	s    *Service
	opt_ map[string]interface{}
}

// List: Query Organization resources.
func (r *OrganizationsService) List() *OrganizationsListCall {
	c := &OrganizationsListCall{s: r.s, opt_: make(map[string]interface{})}
	return c
}

// Filter sets the optional parameter "filter": An optional query string
// used to filter the Organizations to be return in
// the response. Filter
// rules are case-insensitive.
//
//
// Organizations may be filtered by
// `owner.directoryCustomerId` or by
// `domain`, where the domain is a
// Google for Work domain, for
// example:
//
// |Filter|Description|
// |------|-----------|
// |owner.directorycu
// stomerid:123456789|Organizations with `owner.directory_customer_id`
// equal to `123456789`.|
// |domain:google.com|Organizations corresponding
// to the domain `google.com`.|
//
// This field is optional.
func (c *OrganizationsListCall) Filter(filter string) *OrganizationsListCall {
	c.opt_["filter"] = filter
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of Organizations to return in the response.
// This field is optional.
func (c *OrganizationsListCall) PageSize(pageSize int64) *OrganizationsListCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken": A pagination token
// returned from a previous call to ListOrganizations that
// indicates
// from where listing should continue.
// This field is optional.
func (c *OrganizationsListCall) PageToken(pageToken string) *OrganizationsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OrganizationsListCall) Fields(s ...googleapi.Field) *OrganizationsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *OrganizationsListCall) Do() (*ListOrganizationsResponse, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/organizations")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ListOrganizationsResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Query Organization resources.",
	//   "flatPath": "v1beta1/organizations",
	//   "httpMethod": "GET",
	//   "id": "cloudresourcemanager.organizations.list",
	//   "parameterOrder": [],
	//   "parameters": {
	//     "filter": {
	//       "description": "An optional query string used to filter the Organizations to be return in\nthe response. Filter rules are case-insensitive.\n\n\nOrganizations may be filtered by `owner.directoryCustomerId` or by\n`domain`, where the domain is a Google for Work domain, for example:\n\n|Filter|Description|\n|------|-----------|\n|owner.directorycustomerid:123456789|Organizations with `owner.directory_customer_id` equal to `123456789`.|\n|domain:google.com|Organizations corresponding to the domain `google.com`.|\n\nThis field is optional.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "The maximum number of Organizations to return in the response.\nThis field is optional.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A pagination token returned from a previous call to ListOrganizations that\nindicates from where listing should continue.\nThis field is optional.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/organizations",
	//   "response": {
	//     "$ref": "ListOrganizationsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only"
	//   ]
	// }

}

// method id "cloudresourcemanager.organizations.setIamPolicy":

type OrganizationsSetIamPolicyCall struct {
	s                   *Service
	resource            string
	setiampolicyrequest *SetIamPolicyRequest
	opt_                map[string]interface{}
}

// SetIamPolicy: Sets the access control policy on a Organization
// resource. Replaces any
// existing policy.
func (r *OrganizationsService) SetIamPolicy(resource string, setiampolicyrequest *SetIamPolicyRequest) *OrganizationsSetIamPolicyCall {
	c := &OrganizationsSetIamPolicyCall{s: r.s, opt_: make(map[string]interface{})}
	c.resource = resource
	c.setiampolicyrequest = setiampolicyrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OrganizationsSetIamPolicyCall) Fields(s ...googleapi.Field) *OrganizationsSetIamPolicyCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *OrganizationsSetIamPolicyCall) Do() (*Policy, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.setiampolicyrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/organizations/{resource}:setIamPolicy")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
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
	var ret *Policy
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Sets the access control policy on a Organization resource. Replaces any\nexisting policy.",
	//   "flatPath": "v1beta1/organizations/{resource}:setIamPolicy",
	//   "httpMethod": "POST",
	//   "id": "cloudresourcemanager.organizations.setIamPolicy",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "REQUIRED: The resource for which policy is being specified.\n`resource` is usually specified as a path, such as,\n`projects/{project}/zones/{zone}/disks/{disk}`.\n\nThe format for the path specified in this value is resource specific and\nis specified in the documentation for the respective SetIamPolicy rpc.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/organizations/{resource}:setIamPolicy",
	//   "request": {
	//     "$ref": "SetIamPolicyRequest"
	//   },
	//   "response": {
	//     "$ref": "Policy"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudresourcemanager.organizations.testIamPermissions":

type OrganizationsTestIamPermissionsCall struct {
	s                         *Service
	resource                  string
	testiampermissionsrequest *TestIamPermissionsRequest
	opt_                      map[string]interface{}
}

// TestIamPermissions: Returns permissions that a caller has on the
// specified Organization.
func (r *OrganizationsService) TestIamPermissions(resource string, testiampermissionsrequest *TestIamPermissionsRequest) *OrganizationsTestIamPermissionsCall {
	c := &OrganizationsTestIamPermissionsCall{s: r.s, opt_: make(map[string]interface{})}
	c.resource = resource
	c.testiampermissionsrequest = testiampermissionsrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OrganizationsTestIamPermissionsCall) Fields(s ...googleapi.Field) *OrganizationsTestIamPermissionsCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *OrganizationsTestIamPermissionsCall) Do() (*TestIamPermissionsResponse, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.testiampermissionsrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/organizations/{resource}:testIamPermissions")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
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
	var ret *TestIamPermissionsResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns permissions that a caller has on the specified Organization.",
	//   "flatPath": "v1beta1/organizations/{resource}:testIamPermissions",
	//   "httpMethod": "POST",
	//   "id": "cloudresourcemanager.organizations.testIamPermissions",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "REQUIRED: The resource for which policy detail is being requested.\n`resource` is usually specified as a path, such as,\n`projects/{project}/zones/{zone}/disks/{disk}`.\n\nThe format for the path specified in this value is resource specific and\nis specified in the documentation for the respective TestIamPermissions\nrpc.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/organizations/{resource}:testIamPermissions",
	//   "request": {
	//     "$ref": "TestIamPermissionsRequest"
	//   },
	//   "response": {
	//     "$ref": "TestIamPermissionsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only"
	//   ]
	// }

}

// method id "cloudresourcemanager.organizations.update":

type OrganizationsUpdateCall struct {
	s              *Service
	organizationId string
	organization   *Organization
	opt_           map[string]interface{}
}

// Update: Updates an Organization resource.
func (r *OrganizationsService) Update(organizationId string, organization *Organization) *OrganizationsUpdateCall {
	c := &OrganizationsUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.organizationId = organizationId
	c.organization = organization
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OrganizationsUpdateCall) Fields(s ...googleapi.Field) *OrganizationsUpdateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *OrganizationsUpdateCall) Do() (*Organization, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.organization)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/organizations/{organizationId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"organizationId": c.organizationId,
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
	var ret *Organization
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates an Organization resource.",
	//   "flatPath": "v1beta1/organizations/{organizationId}",
	//   "httpMethod": "PUT",
	//   "id": "cloudresourcemanager.organizations.update",
	//   "parameterOrder": [
	//     "organizationId"
	//   ],
	//   "parameters": {
	//     "organizationId": {
	//       "description": "An immutable id for the Organization that is assigned on creation. This\nshould be omitted when creating a new Organization.\nThis field is read-only.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/organizations/{organizationId}",
	//   "request": {
	//     "$ref": "Organization"
	//   },
	//   "response": {
	//     "$ref": "Organization"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudresourcemanager.projects.create":

type ProjectsCreateCall struct {
	s       *Service
	project *Project
	opt_    map[string]interface{}
}

// Create: Creates a project resource.
//
// Initially, the project resource
// is owned by its creator exclusively.
// The creator can later grant
// permission to others to read or update the
// project.
//
// Several APIs are
// activated automatically for the project, including
// Google Cloud
// Storage.
func (r *ProjectsService) Create(project *Project) *ProjectsCreateCall {
	c := &ProjectsCreateCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsCreateCall) Fields(s ...googleapi.Field) *ProjectsCreateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsCreateCall) Do() (*Project, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.project)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/projects")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
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
	var ret *Project
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a project resource.\n\nInitially, the project resource is owned by its creator exclusively.\nThe creator can later grant permission to others to read or update the\nproject.\n\nSeveral APIs are activated automatically for the project, including\nGoogle Cloud Storage.",
	//   "flatPath": "v1beta1/projects",
	//   "httpMethod": "POST",
	//   "id": "cloudresourcemanager.projects.create",
	//   "parameterOrder": [],
	//   "parameters": {},
	//   "path": "v1beta1/projects",
	//   "request": {
	//     "$ref": "Project"
	//   },
	//   "response": {
	//     "$ref": "Project"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudresourcemanager.projects.delete":

type ProjectsDeleteCall struct {
	s         *Service
	projectId string
	opt_      map[string]interface{}
}

// Delete: Marks the project identified by the specified
// `project_id`
// (for example, `my-project-123`) for deletion.
// This method will only
// affect the project if the following criteria are met:
//
// + The project
// does not have a billing account associated with it.
// + The project has
// a lifecycle state of
// ACTIVE.
//
// This method changes the project's
// lifecycle state from
// ACTIVE
// to DELETE_REQUESTED.
// The deletion starts
// at an unspecified time,
// at which point the lifecycle state changes to
// DELETE_IN_PROGRESS.
//
// Until the deletion completes, you can check the
// lifecycle state
// checked by retrieving the project with
// GetProject,
// and the project remains visible to ListProjects.
// However,
// you cannot update the project.
//
// After the deletion completes, the
// project is not retrievable by
// the  GetProject and
// ListProjects
// methods.
//
// The caller must have modify permissions for this project.
func (r *ProjectsService) Delete(projectId string) *ProjectsDeleteCall {
	c := &ProjectsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsDeleteCall) Fields(s ...googleapi.Field) *ProjectsDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsDeleteCall) Do() (*Empty, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/projects/{projectId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
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
	var ret *Empty
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Marks the project identified by the specified\n`project_id` (for example, `my-project-123`) for deletion.\nThis method will only affect the project if the following criteria are met:\n\n+ The project does not have a billing account associated with it.\n+ The project has a lifecycle state of\nACTIVE.\n\nThis method changes the project's lifecycle state from\nACTIVE\nto DELETE_REQUESTED.\nThe deletion starts at an unspecified time,\nat which point the lifecycle state changes to DELETE_IN_PROGRESS.\n\nUntil the deletion completes, you can check the lifecycle state\nchecked by retrieving the project with GetProject,\nand the project remains visible to ListProjects.\nHowever, you cannot update the project.\n\nAfter the deletion completes, the project is not retrievable by\nthe  GetProject and\nListProjects methods.\n\nThe caller must have modify permissions for this project.",
	//   "flatPath": "v1beta1/projects/{projectId}",
	//   "httpMethod": "DELETE",
	//   "id": "cloudresourcemanager.projects.delete",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "The project ID (for example, `foo-bar-123`).\n\nRequired.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/projects/{projectId}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudresourcemanager.projects.get":

type ProjectsGetCall struct {
	s         *Service
	projectId string
	opt_      map[string]interface{}
}

// Get: Retrieves the project identified by the specified
// `project_id`
// (for example, `my-project-123`).
//
// The caller must have read
// permissions for this project.
func (r *ProjectsService) Get(projectId string) *ProjectsGetCall {
	c := &ProjectsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsGetCall) Fields(s ...googleapi.Field) *ProjectsGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsGetCall) Do() (*Project, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/projects/{projectId}")
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
	var ret *Project
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves the project identified by the specified\n`project_id` (for example, `my-project-123`).\n\nThe caller must have read permissions for this project.",
	//   "flatPath": "v1beta1/projects/{projectId}",
	//   "httpMethod": "GET",
	//   "id": "cloudresourcemanager.projects.get",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "The project ID (for example, `my-project-123`).\n\nRequired.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/projects/{projectId}",
	//   "response": {
	//     "$ref": "Project"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only"
	//   ]
	// }

}

// method id "cloudresourcemanager.projects.getIamPolicy":

type ProjectsGetIamPolicyCall struct {
	s                   *Service
	resource            string
	getiampolicyrequest *GetIamPolicyRequest
	opt_                map[string]interface{}
}

// GetIamPolicy: Returns the IAM access control policy for specified
// project.
func (r *ProjectsService) GetIamPolicy(resource string, getiampolicyrequest *GetIamPolicyRequest) *ProjectsGetIamPolicyCall {
	c := &ProjectsGetIamPolicyCall{s: r.s, opt_: make(map[string]interface{})}
	c.resource = resource
	c.getiampolicyrequest = getiampolicyrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsGetIamPolicyCall) Fields(s ...googleapi.Field) *ProjectsGetIamPolicyCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsGetIamPolicyCall) Do() (*Policy, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.getiampolicyrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/projects/{resource}:getIamPolicy")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
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
	var ret *Policy
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns the IAM access control policy for specified project.",
	//   "flatPath": "v1beta1/projects/{resource}:getIamPolicy",
	//   "httpMethod": "POST",
	//   "id": "cloudresourcemanager.projects.getIamPolicy",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "REQUIRED: The resource for which policy is being requested.\n`resource` is usually specified as a path, such as,\n`projects/{project}/zones/{zone}/disks/{disk}`.\n\nThe format for the path specified in this value is resource specific and\nis specified in the documentation for the respective GetIamPolicy rpc.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/projects/{resource}:getIamPolicy",
	//   "request": {
	//     "$ref": "GetIamPolicyRequest"
	//   },
	//   "response": {
	//     "$ref": "Policy"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only"
	//   ]
	// }

}

// method id "cloudresourcemanager.projects.list":

type ProjectsListCall struct {
	s    *Service
	opt_ map[string]interface{}
}

// List: Lists projects that are visible to the user and satisfy
// the
// specified filter. This method returns projects in an unspecified
// order.
// New projects do not necessarily appear at the end of the list.
func (r *ProjectsService) List() *ProjectsListCall {
	c := &ProjectsListCall{s: r.s, opt_: make(map[string]interface{})}
	return c
}

// Filter sets the optional parameter "filter": An expression for
// filtering the results of the request.  Filter rules are
// case
// insensitive. The fields eligible for filtering are:
//
// + `name`
// +
// `id`
// + <code>labels.<em>key</em></code> where *key* is the name of a
// label
//
// Some examples of using labels as
// filters:
//
// |Filter|Description|
// |------|-----------|
// |name:*|The
// project has a name.|
// |name:Howl|The project's name is `Howl` or
// `howl`.|
// |name:HOWL|Equivalent to above.|
// |NAME:howl|Equivalent to
// above.|
// |labels.color:*|The project has the label
// `color`.|
// |labels.color:red|The project's label `color` has the value
// `red`.|
// |labels.color:red&nbsp;label.size:big|The project's label
// `color` has the value `red` and its label `size` has the value `big`.
func (c *ProjectsListCall) Filter(filter string) *ProjectsListCall {
	c.opt_["filter"] = filter
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of Projects to return in the response.
// The server can return fewer
// projects than requested.
// If unspecified, server picks an appropriate
// default.
func (c *ProjectsListCall) PageSize(pageSize int64) *ProjectsListCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken": A pagination token
// returned from a previous call to ListProject
// that indicates from
// where listing should continue.
func (c *ProjectsListCall) PageToken(pageToken string) *ProjectsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsListCall) Fields(s ...googleapi.Field) *ProjectsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsListCall) Do() (*ListProjectsResponse, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/projects")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", "google-api-go-client/0.5")
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *ListProjectsResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists projects that are visible to the user and satisfy the\nspecified filter. This method returns projects in an unspecified order.\nNew projects do not necessarily appear at the end of the list.",
	//   "flatPath": "v1beta1/projects",
	//   "httpMethod": "GET",
	//   "id": "cloudresourcemanager.projects.list",
	//   "parameterOrder": [],
	//   "parameters": {
	//     "filter": {
	//       "description": "An expression for filtering the results of the request.  Filter rules are\ncase insensitive. The fields eligible for filtering are:\n\n+ `name`\n+ `id`\n+ \u003ccode\u003elabels.\u003cem\u003ekey\u003c/em\u003e\u003c/code\u003e where *key* is the name of a label\n\nSome examples of using labels as filters:\n\n|Filter|Description|\n|------|-----------|\n|name:*|The project has a name.|\n|name:Howl|The project's name is `Howl` or `howl`.|\n|name:HOWL|Equivalent to above.|\n|NAME:howl|Equivalent to above.|\n|labels.color:*|The project has the label `color`.|\n|labels.color:red|The project's label `color` has the value `red`.|\n|labels.color:red\u0026nbsp;label.size:big|The project's label `color` has the value `red` and its label `size` has the value `big`.\n\nOptional.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "The maximum number of Projects to return in the response.\nThe server can return fewer projects than requested.\nIf unspecified, server picks an appropriate default.\n\nOptional.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A pagination token returned from a previous call to ListProject\nthat indicates from where listing should continue.\n\nOptional.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/projects",
	//   "response": {
	//     "$ref": "ListProjectsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only"
	//   ]
	// }

}

// method id "cloudresourcemanager.projects.setIamPolicy":

type ProjectsSetIamPolicyCall struct {
	s                   *Service
	resource            string
	setiampolicyrequest *SetIamPolicyRequest
	opt_                map[string]interface{}
}

// SetIamPolicy: Sets the IAM access control policy for the specified
// project. We do not
// currently support 'domain:' prefixed members in a
// Binding of a Policy.
//
// Calling this method requires enabling the App
// Engine Admin API.
func (r *ProjectsService) SetIamPolicy(resource string, setiampolicyrequest *SetIamPolicyRequest) *ProjectsSetIamPolicyCall {
	c := &ProjectsSetIamPolicyCall{s: r.s, opt_: make(map[string]interface{})}
	c.resource = resource
	c.setiampolicyrequest = setiampolicyrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsSetIamPolicyCall) Fields(s ...googleapi.Field) *ProjectsSetIamPolicyCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsSetIamPolicyCall) Do() (*Policy, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.setiampolicyrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/projects/{resource}:setIamPolicy")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
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
	var ret *Policy
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Sets the IAM access control policy for the specified project. We do not\ncurrently support 'domain:' prefixed members in a Binding of a Policy.\n\nCalling this method requires enabling the App Engine Admin API.",
	//   "flatPath": "v1beta1/projects/{resource}:setIamPolicy",
	//   "httpMethod": "POST",
	//   "id": "cloudresourcemanager.projects.setIamPolicy",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "REQUIRED: The resource for which policy is being specified.\n`resource` is usually specified as a path, such as,\n`projects/{project}/zones/{zone}/disks/{disk}`.\n\nThe format for the path specified in this value is resource specific and\nis specified in the documentation for the respective SetIamPolicy rpc.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/projects/{resource}:setIamPolicy",
	//   "request": {
	//     "$ref": "SetIamPolicyRequest"
	//   },
	//   "response": {
	//     "$ref": "Policy"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudresourcemanager.projects.testIamPermissions":

type ProjectsTestIamPermissionsCall struct {
	s                         *Service
	resource                  string
	testiampermissionsrequest *TestIamPermissionsRequest
	opt_                      map[string]interface{}
}

// TestIamPermissions: Tests the specified permissions against the IAM
// access control policy
// for the specified project.
func (r *ProjectsService) TestIamPermissions(resource string, testiampermissionsrequest *TestIamPermissionsRequest) *ProjectsTestIamPermissionsCall {
	c := &ProjectsTestIamPermissionsCall{s: r.s, opt_: make(map[string]interface{})}
	c.resource = resource
	c.testiampermissionsrequest = testiampermissionsrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsTestIamPermissionsCall) Fields(s ...googleapi.Field) *ProjectsTestIamPermissionsCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsTestIamPermissionsCall) Do() (*TestIamPermissionsResponse, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.testiampermissionsrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/projects/{resource}:testIamPermissions")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
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
	var ret *TestIamPermissionsResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Tests the specified permissions against the IAM access control policy\nfor the specified project.",
	//   "flatPath": "v1beta1/projects/{resource}:testIamPermissions",
	//   "httpMethod": "POST",
	//   "id": "cloudresourcemanager.projects.testIamPermissions",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "REQUIRED: The resource for which policy detail is being requested.\n`resource` is usually specified as a path, such as,\n`projects/{project}/zones/{zone}/disks/{disk}`.\n\nThe format for the path specified in this value is resource specific and\nis specified in the documentation for the respective TestIamPermissions\nrpc.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/projects/{resource}:testIamPermissions",
	//   "request": {
	//     "$ref": "TestIamPermissionsRequest"
	//   },
	//   "response": {
	//     "$ref": "TestIamPermissionsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only"
	//   ]
	// }

}

// method id "cloudresourcemanager.projects.undelete":

type ProjectsUndeleteCall struct {
	s         *Service
	projectId string
	opt_      map[string]interface{}
}

// Undelete: Restores the project identified by the
// specified
// `project_id` (for example, `my-project-123`).
// You can only
// use this method for a project that has a lifecycle state
// of
// DELETE_REQUESTED.
// After deletion starts, as indicated by a
// lifecycle state of
// DELETE_IN_PROGRESS,
// the project cannot be
// restored.
//
// The caller must have modify permissions for this project.
func (r *ProjectsService) Undelete(projectId string) *ProjectsUndeleteCall {
	c := &ProjectsUndeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsUndeleteCall) Fields(s ...googleapi.Field) *ProjectsUndeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsUndeleteCall) Do() (*Empty, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/projects/{projectId}:undelete")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
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
	var ret *Empty
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Restores the project identified by the specified\n`project_id` (for example, `my-project-123`).\nYou can only use this method for a project that has a lifecycle state of\nDELETE_REQUESTED.\nAfter deletion starts, as indicated by a lifecycle state of\nDELETE_IN_PROGRESS,\nthe project cannot be restored.\n\nThe caller must have modify permissions for this project.",
	//   "flatPath": "v1beta1/projects/{projectId}:undelete",
	//   "httpMethod": "POST",
	//   "id": "cloudresourcemanager.projects.undelete",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "The project ID (for example, `foo-bar-123`).\n\nRequired.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/projects/{projectId}:undelete",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "cloudresourcemanager.projects.update":

type ProjectsUpdateCall struct {
	s         *Service
	projectId string
	project   *Project
	opt_      map[string]interface{}
}

// Update: Updates the attributes of the project identified by the
// specified
// `project_id` (for example, `my-project-123`).
//
// The caller
// must have modify permissions for this project.
func (r *ProjectsService) Update(projectId string, project *Project) *ProjectsUpdateCall {
	c := &ProjectsUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.project = project
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsUpdateCall) Fields(s ...googleapi.Field) *ProjectsUpdateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsUpdateCall) Do() (*Project, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.project)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/projects/{projectId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
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
	var ret *Project
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates the attributes of the project identified by the specified\n`project_id` (for example, `my-project-123`).\n\nThe caller must have modify permissions for this project.",
	//   "flatPath": "v1beta1/projects/{projectId}",
	//   "httpMethod": "PUT",
	//   "id": "cloudresourcemanager.projects.update",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "The project ID (for example, `my-project-123`).\n\nRequired.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/projects/{projectId}",
	//   "request": {
	//     "$ref": "Project"
	//   },
	//   "response": {
	//     "$ref": "Project"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}
