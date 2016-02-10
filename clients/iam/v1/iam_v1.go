// Package iam provides access to the Google Identity and Access Management API.
//
// See https://cloud.google.com/iam/
//
// Usage example:
//
//   import "github.com/skelterjohn/gcloud_apis/clients/iam/v1"
//   ...
//   iamService, err := iam.New(oauthHttpClient)
package iam

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

const apiId = "iam:v1"
const apiName = "iam"
const apiVersion = "v1"
const basePath = "https://iam.googleapis.com/"

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
	s.IamPolicies = NewIamPoliciesService(s)
	s.Projects = NewProjectsService(s)
	s.Roles = NewRolesService(s)
	s.V1 = NewV1Service(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	IamPolicies *IamPoliciesService

	Projects *ProjectsService

	Roles *RolesService

	V1 *V1Service
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewIamPoliciesService(s *Service) *IamPoliciesService {
	rs := &IamPoliciesService{s: s}
	return rs
}

type IamPoliciesService struct {
	s *Service
}

func NewProjectsService(s *Service) *ProjectsService {
	rs := &ProjectsService{s: s}
	rs.ServiceAccounts = NewProjectsServiceAccountsService(s)
	return rs
}

type ProjectsService struct {
	s *Service

	ServiceAccounts *ProjectsServiceAccountsService
}

func NewProjectsServiceAccountsService(s *Service) *ProjectsServiceAccountsService {
	rs := &ProjectsServiceAccountsService{s: s}
	rs.Keys = NewProjectsServiceAccountsKeysService(s)
	return rs
}

type ProjectsServiceAccountsService struct {
	s *Service

	Keys *ProjectsServiceAccountsKeysService
}

func NewProjectsServiceAccountsKeysService(s *Service) *ProjectsServiceAccountsKeysService {
	rs := &ProjectsServiceAccountsKeysService{s: s}
	return rs
}

type ProjectsServiceAccountsKeysService struct {
	s *Service
}

func NewRolesService(s *Service) *RolesService {
	rs := &RolesService{s: s}
	return rs
}

type RolesService struct {
	s *Service
}

func NewV1Service(s *Service) *V1Service {
	rs := &V1Service{s: s}
	return rs
}

type V1Service struct {
	s *Service
}

// AuditConfig: Enables "data access" audit logging for a service and
// specifies a list
// of members that are log-exempted.
type AuditConfig struct {
	// ExemptedMembers: Specifies the identities that are exempted from
	// "data access" audit
	// logging for the `service` specified above.
	// Follows the same format of Binding.members.
	ExemptedMembers []string `json:"exemptedMembers,omitempty"`

	// Service: Specifies a service that will be enabled for "data access"
	// audit
	// logging.
	// For example, `resourcemanager`, `storage`, `compute`.
	// `allServices` is a special value that covers all services.
	Service string `json:"service,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ExemptedMembers") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *AuditConfig) MarshalJSON() ([]byte, error) {
	type noMethod AuditConfig
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// Binding: Associates `members` with a `role`.
type Binding struct {
	// Members: Specifies the identities requesting access for a Cloud
	// Platform resource.
	// `members` can have the following values:
	//
	// * `allUsers`: A special identifier that represents anyone who is
	//    on the internet; with or without a Google account.
	//
	// * `allAuthenticatedUsers`: A special identifier that represents
	// anyone
	//    who is authenticated with a Google account or a service
	// account.
	//
	// * `user:{emailid}`: An email address that represents a specific
	// Google
	//    account. For example, `alice@gmail.com` or `joe@example.com`.
	//
	// * `serviceAccount:{emailid}`: An email address that represents a
	// service
	//    account. For example,
	// `my-other-app@appspot.gserviceaccount.com`.
	//
	// * `group:{emailid}`: An email address that represents a Google
	// group.
	//    For example, `admins@example.com`.
	//
	// * `domain:{domain}`: A Google Apps domain name that represents all
	// the
	//    users of that domain. For example, `google.com` or `example.com`.
	//
	Members []string `json:"members,omitempty"`

	// Role: Role that is assigned to `members`.
	// For example, `roles/viewer`, `roles/editor`, or
	// `roles/owner`.
	// Required
	Role string `json:"role,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Members") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Binding) MarshalJSON() ([]byte, error) {
	type noMethod Binding
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// CloudAuditOptions: Write a Cloud Audit log
type CloudAuditOptions struct {
}

// Condition: A condition to be met.
type Condition struct {
	// Iam: Trusted attributes supplied by the IAM system.
	//
	// Possible values:
	//   "NO_ATTR" - Default non-attribute.
	//   "AUTHORITY" - Either principal or (if present) authority
	//   "ATTRIBUTION" - selector
	// Always the original principal, but making clear
	Iam string `json:"iam,omitempty"`

	// Op: An operator to apply the subject with.
	//
	// Possible values:
	//   "NO_OP" - Default no-op.
	//   "EQUALS" - Equality check.
	//   "NOT_EQUALS" - Non-equality check.
	//   "IN" - Set-inclusion check.
	//   "NOT_IN" - Set-exclusion check.
	//   "DISCHARGED" - Subject is discharged
	Op string `json:"op,omitempty"`

	// Svc: Trusted attributes discharged by the service.
	Svc string `json:"svc,omitempty"`

	// Sys: Trusted attributes supplied by any service that owns resources
	// and uses
	// the IAM system for access control.
	//
	// Possible values:
	//   "NO_ATTR" - Default non-attribute type
	//   "REGION" - Region of the resource
	//   "SERVICE" - Service name
	//   "NAME" - Resource name
	//   "IP" - IP address of the caller
	Sys string `json:"sys,omitempty"`

	// Value: The object of the condition. Exactly one of these must be set.
	Value string `json:"value,omitempty"`

	// Values: The objects of the condition. This is mutually exclusive with
	// 'value'.
	Values []string `json:"values,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Iam") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Condition) MarshalJSON() ([]byte, error) {
	type noMethod Condition
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// CounterOptions: Options for counters
type CounterOptions struct {
	// Field: The field value to attribute.
	Field string `json:"field,omitempty"`

	// Metric: The metric to update.
	Metric string `json:"metric,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Field") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *CounterOptions) MarshalJSON() ([]byte, error) {
	type noMethod CounterOptions
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// CreateServiceAccountRequest: The service account create request.
type CreateServiceAccountRequest struct {
	// AccountId: Required. The account id that is used to generate the
	// service account
	// email address and a stable unique id. It is unique within a
	// project,
	// must be 1-63 characters long, and match the regular
	// expression
	// [a-z]([-a-z0-9]*[a-z0-9]) to comply with RFC1035.
	AccountId string `json:"accountId,omitempty"`

	// ServiceAccount: The ServiceAccount resource to create. Currently only
	// the
	// display_name and the description are user assignable values.
	ServiceAccount *ServiceAccount `json:"serviceAccount,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AccountId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *CreateServiceAccountRequest) MarshalJSON() ([]byte, error) {
	type noMethod CreateServiceAccountRequest
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// DataAccessOptions: Write a Data Access (Gin) log
type DataAccessOptions struct {
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

// GetPolicyDetailsRequest: The request to get the current policy and
// the policies on the inherited
// resources the user has access to.
type GetPolicyDetailsRequest struct {
	// FullResourcePath: REQUIRED: The full resource path of the current
	// policy being
	// requested, e.g., //dataflow.googleapis.com/projects/../jobs/..
	FullResourcePath string `json:"fullResourcePath,omitempty"`

	// PageSize: Limit on the number of policies to include in the
	// response.
	// Further accounts can subsequently be obtained by including
	// the
	// [GetPolicyDetailsResponse.next_page_token] in a subsequent
	// request.
	// If zero, the default page size 20 will be used.
	// Must be given a value in range [0, 100], otherwise an invalid
	// argument
	// error will be returned.
	PageSize int64 `json:"pageSize,omitempty"`

	// PageToken: Optional pagination token returned in an
	// earlier
	// [GetPolicyDetailsResponse.next_page_token].
	PageToken string `json:"pageToken,omitempty"`

	// ForceSendFields is a list of field names (e.g. "FullResourcePath") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *GetPolicyDetailsRequest) MarshalJSON() ([]byte, error) {
	type noMethod GetPolicyDetailsRequest
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// GetPolicyDetailsResponse: The response to the GetPolicyDetailsRequest
// containing the current policy and
// the policies on the inherited resources the user has access to.
type GetPolicyDetailsResponse struct {
	// NextPageToken: To retrieve the next page of results,
	// set
	// [GetPolicyDetailsRequest.page_token] to this value.
	// If this value is empty, then there are not any further policies that
	// the
	// user has access to.
	// The lifetime is 60 minutes. An "Expired pagination token" error will
	// be
	// returned if exceeded.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Policies: The current policy and all the inherited policies the user
	// has
	// access to.
	Policies []*PolicyDetail `json:"policies,omitempty"`

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

func (s *GetPolicyDetailsResponse) MarshalJSON() ([]byte, error) {
	type noMethod GetPolicyDetailsResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// ListServiceAccountKeysResponse: The service account keys list
// response.
type ListServiceAccountKeysResponse struct {
	// Keys: The public keys for the service account.
	Keys []*ServiceAccountKey `json:"keys,omitempty"`

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

func (s *ListServiceAccountKeysResponse) MarshalJSON() ([]byte, error) {
	type noMethod ListServiceAccountKeysResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// ListServiceAccountsResponse: The service account list response.
type ListServiceAccountsResponse struct {
	// Accounts: The list of matching service accounts.
	Accounts []*ServiceAccount `json:"accounts,omitempty"`

	// NextPageToken: To retrieve the next page of results,
	// set
	// [ListServiceAccountsRequest.page_token] to this value.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Accounts") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *ListServiceAccountsResponse) MarshalJSON() ([]byte, error) {
	type noMethod ListServiceAccountsResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// LogConfig: Specifies what kind of log the caller must write
// Increment a streamz counter with the specified metric and field
// names.
//
// Metric names should start with a '/', generally be
// lowercase-only,
// and end in "_count". Field names should not contain an initial
// slash.
// The actual exported metric names will have "/iam/policy"
// prepended.
//
// Field names correspond to IAM request parameters and field values
// are
// their respective values.
//
// At present only "iam_principal", corresponding to
// IAMContext.principal,
// is supported.
//
// Examples:
//   counter { metric: "/debug_access_count"  field: "iam_principal" }
//   ==> increment counter /iam/policy/backend_debug_access_count
//                         {iam_principal=[value of
// IAMContext.principal]}
//
// At this time we do not support:
// * multiple field names (though this may be supported in the future)
// * decrementing the counter
// * incrementing it by anything other than 1
type LogConfig struct {
	// CloudAudit: Cloud audit options.
	CloudAudit *CloudAuditOptions `json:"cloudAudit,omitempty"`

	// Counter: Counter options.
	Counter *CounterOptions `json:"counter,omitempty"`

	// DataAccess: Data access options.
	DataAccess *DataAccessOptions `json:"dataAccess,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CloudAudit") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *LogConfig) MarshalJSON() ([]byte, error) {
	type noMethod LogConfig
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// Policy: Defines an Identity and Access Management (IAM) policy. It is
// used to
// specify access control policies for Cloud Platform resources.
//
//
// A `Policy` consists of a list of `bindings`. A `Binding` binds a list
// of
// `members` to a `role`, where the members can be user accounts, Google
// groups,
// Google domains, and service accounts. A `role` is a named list of
// permissions
// defined by IAM.
//
// **Example**
//
//     {
//         "bindings": [
//          {
//              "role": "roles/owner",
//              "members": [
//              "user:mike@example.com",
//              "group:admins@example.com",
//              "domain:google.com",
//
// "serviceAccount:my-other-app@appspot.gserviceaccount.com"]
//          },
//          {
//              "role": "roles/viewer",
//              "members": ["user:sean@example.com"]
//          }
//          ]
//     }
//
// For a description of IAM and its features, see the
// [IAM developer's guide](https://cloud.google.com/iam).
type Policy struct {
	// AuditConfigs: Specifies audit logging configs for "data
	// access".
	// "data access": generally refers to data reads/writes and admin
	// reads.
	// "admin activity": generally refers to admin writes.
	//
	// Note: `AuditConfig` doesn't apply to "admin activity", which
	// always
	// enables audit logging.
	AuditConfigs []*AuditConfig `json:"auditConfigs,omitempty"`

	// Bindings: Associates a list of `members` to a `role`.
	// Multiple `bindings` must not be specified for the same
	// `role`.
	// `bindings` with no members will result in an error.
	Bindings []*Binding `json:"bindings,omitempty"`

	// Etag: `etag` is used for optimistic concurrency control as a way to
	// help
	// prevent simultaneous updates of a policy from overwriting each
	// other.
	// It is strongly suggested that systems make use of the `etag` in
	// the
	// read-modify-write cycle to perform policy updates in order to avoid
	// race
	// conditions: An `etag` is returned in the response to `getIamPolicy`,
	// and
	// systems are expected to put that etag in the request to
	// `setIamPolicy` to
	// ensure that their change will be applied to the same version of the
	// policy.
	//
	// If no `etag` is provided in the call to `setIamPolicy`, then the
	// existing
	// policy is overwritten blindly.
	Etag string `json:"etag,omitempty"`

	IamOwned bool `json:"iamOwned,omitempty"`

	Rules []*Rule `json:"rules,omitempty"`

	// Version: Version of the `Policy`. The default version is 0.
	Version int64 `json:"version,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "AuditConfigs") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Policy) MarshalJSON() ([]byte, error) {
	type noMethod Policy
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// PolicyDetail: A policy and its full resource path.
type PolicyDetail struct {
	// FullResourcePath: The full resource path of the policy
	// e.g., //dataflow.googleapis.com/projects/../jobs/..
	// Note that a resource and its inherited resource have
	// different
	// full_resource_path.
	FullResourcePath string `json:"fullResourcePath,omitempty"`

	// Policy: The policy of a resource/project/folder.
	Policy *Policy `json:"policy,omitempty"`

	// ForceSendFields is a list of field names (e.g. "FullResourcePath") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *PolicyDetail) MarshalJSON() ([]byte, error) {
	type noMethod PolicyDetail
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// QueryGrantableRolesRequest: The grantable role query request.
type QueryGrantableRolesRequest struct {
	// FullResourceName: Required. The full resource name to query the list
	// of grantable roles for.
	//
	// This will follow the GCP resource format.
	// For example, a GCP project of ID "my-project" will be
	// named
	// "//cloudresourcemanager.googleapis.com/projects/my-project".
	FullResourceName string `json:"fullResourceName,omitempty"`

	// ForceSendFields is a list of field names (e.g. "FullResourceName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *QueryGrantableRolesRequest) MarshalJSON() ([]byte, error) {
	type noMethod QueryGrantableRolesRequest
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// QueryGrantableRolesResponse: The grantable role query response.
type QueryGrantableRolesResponse struct {
	// Roles: The list of matching roles.
	Roles []*Role `json:"roles,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Roles") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *QueryGrantableRolesResponse) MarshalJSON() ([]byte, error) {
	type noMethod QueryGrantableRolesResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// Role: A role in the Identity and Access Management API.
type Role struct {
	// Description: Optional.  A human-readable description for the role.
	Description string `json:"description,omitempty"`

	// Name: The name of the role.
	//
	// Examples of roles names are:
	// "roles/editor", "roles/viewer" and "roles/logging.viewer"
	Name string `json:"name,omitempty"`

	// Title: Optional.  A human-readable title for the role.  Typically
	// this
	// is limited to 100 UTF-8 bytes.
	Title string `json:"title,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Description") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Role) MarshalJSON() ([]byte, error) {
	type noMethod Role
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// Rule: A rule to be applied in a Policy.
type Rule struct {
	// Action: Required
	//
	// Possible values:
	//   "NO_ACTION" - Default no action.
	//   "ALLOW" - Matching 'Entries' grant access.
	//   "ALLOW_WITH_LOG" - Matching 'Entries' grant access and the caller
	// promises to log
	// the request per the returned log_configs.
	//   "DENY" - Matching 'Entries' deny access.
	//   "DENY_WITH_LOG" - Matching 'Entries' deny access and the caller
	// promises to log
	// the request per the returned log_configs.
	//   "LOG" - Matching 'Entries' tell IAM.Check callers to generate logs.
	Action string `json:"action,omitempty"`

	// Conditions: Additional restrictions that must be met
	Conditions []*Condition `json:"conditions,omitempty"`

	// Description: Human-readable description of the rule.
	Description string `json:"description,omitempty"`

	// In: The rule matches if the PRINCIPAL/AUTHORITY_SELECTOR is in
	// this
	// set of entries.
	In []string `json:"in,omitempty"`

	// LogConfig: The config returned to callers of tech.iam.IAM.CheckPolicy
	// for any entries
	// that match the LOG action.
	LogConfig []*LogConfig `json:"logConfig,omitempty"`

	// NotIn: The rule matches if the PRINCIPAL/AUTHORITY_SELECTOR is not in
	// this
	// set of entries.
	// The format for in and not_in entries is the same as for members in
	// a
	// Binding (see google/iam/v1/policy.proto).
	NotIn []string `json:"notIn,omitempty"`

	// Permissions: A permission is a string of form '<service>.<resource
	// type>.<verb>'
	// (e.g., 'storage.buckets.list'). A value of '*' matches all
	// permissions,
	// and a verb part of '*' (e.g., 'storage.buckets.*') matches all verbs.
	Permissions []string `json:"permissions,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Action") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Rule) MarshalJSON() ([]byte, error) {
	type noMethod Rule
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// ServiceAccount: A service account in the Identity and Access
// Management API.
//
// To create a service account, you specify the project_id and
// account_id
// for the account.  The account_id is unique within the project, and
// used
// to generate the service account email address and a stable
// unique id.
//
// All other methods can identify accounts using the
// format
// "projects/{project}/serviceAccounts/{account}".
// Using '-' as a wildcard for the project, will infer the project
// from
// the account. The account value can be the email address or the
// unique_id
// of the service account.
type ServiceAccount struct {
	// Description: Optional. A user-specified opaque description of the
	// service account.
	Description string `json:"description,omitempty"`

	// DisplayName: Optional. A user-specified description of the service
	// account.  Must be
	// fewer than 100 UTF-8 bytes.
	DisplayName string `json:"displayName,omitempty"`

	// Email: @OutputOnly Email address of the service account.
	Email string `json:"email,omitempty"`

	// Etag: Used to perform a consistent read-modify-write.
	Etag string `json:"etag,omitempty"`

	// Name: The resource name of the service account in the
	// format
	// "projects/{project}/serviceAccounts/{account}".
	//
	// In requests using '-' as a wildcard for the project, will infer the
	// project
	// from the account and the account value can be the email address or
	// the
	// unique_id of the service account.
	//
	// In responses the resource name will always be in the
	// format
	// "projects/{project}/serviceAccounts/{email}".
	Name string `json:"name,omitempty"`

	// Oauth2ClientId: @OutputOnly. The OAuth2 client id for the service
	// account.
	// This is used in conjunction with the OAuth2 clientconfig API to
	// make
	// three legged OAuth2 (3LO) flows to access the data of Google users.
	Oauth2ClientId string `json:"oauth2ClientId,omitempty"`

	// ProjectId: @OutputOnly The id of the project that owns the service
	// account.
	ProjectId string `json:"projectId,omitempty"`

	// UniqueId: @OutputOnly unique and stable id of the service account.
	UniqueId string `json:"uniqueId,omitempty"`

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

func (s *ServiceAccount) MarshalJSON() ([]byte, error) {
	type noMethod ServiceAccount
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// ServiceAccountKey: Represents a service account key.
//
// A service account can have 0 or more key pairs. The private keys for
// these
// are not stored by Google.
// ServiceAccountKeys are immutable.
type ServiceAccountKey struct {
	// Name: The resource name of the service account key in the
	// format
	// "projects/{project}/serviceAccounts/{email}/keys/{key}".
	Name string `json:"name,omitempty"`

	// PrivateKeyData: The key data.
	PrivateKeyData string `json:"privateKeyData,omitempty"`

	// PrivateKeyType: The type of the private key.
	//
	// Possible values:
	//   "TYPE_UNSPECIFIED" - Unspecified.
	//   "TYPE_PKCS12_FILE" - PKCS12 format.
	// The password for the PKCS12 file is 'notasecret'.
	// For more information, see https://tools.ietf.org/html/rfc7292.
	//   "TYPE_GOOGLE_CREDENTIALS_FILE" - Google Credentials File format.
	PrivateKeyType string `json:"privateKeyType,omitempty"`

	// ValidAfterTime: The key can be used after this timestamp.
	ValidAfterTime string `json:"validAfterTime,omitempty"`

	// ValidBeforeTime: The key can be used before this timestamp.
	ValidBeforeTime string `json:"validBeforeTime,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Name") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *ServiceAccountKey) MarshalJSON() ([]byte, error) {
	type noMethod ServiceAccountKey
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// SetIamPolicyRequest: Request message for `SetIamPolicy` method.
type SetIamPolicyRequest struct {
	// Policy: REQUIRED: The complete policy to be applied to the
	// `resource`. The size of
	// the policy is limited to a few 10s of KB. An empty policy is a
	// valid policy but certain Cloud Platform services (such as
	// Projects)
	// might reject them.
	Policy *Policy `json:"policy,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Policy") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *SetIamPolicyRequest) MarshalJSON() ([]byte, error) {
	type noMethod SetIamPolicyRequest
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// SignBlobRequest: The service account sign blob request.
type SignBlobRequest struct {
	// BytesToSign: The bytes to sign
	BytesToSign string `json:"bytesToSign,omitempty"`

	// ForceSendFields is a list of field names (e.g. "BytesToSign") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *SignBlobRequest) MarshalJSON() ([]byte, error) {
	type noMethod SignBlobRequest
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// SignBlobResponse: The service account sign blob response.
type SignBlobResponse struct {
	// KeyId: The id of the key used to sign the blob.
	KeyId string `json:"keyId,omitempty"`

	// Signature: The signed blob.
	Signature string `json:"signature,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "KeyId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *SignBlobResponse) MarshalJSON() ([]byte, error) {
	type noMethod SignBlobResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// TestIamPermissionsRequest: Request message for `TestIamPermissions`
// method.
type TestIamPermissionsRequest struct {
	// Permissions: The set of permissions to check for the `resource`.
	// Permissions with
	// wildcards (such as '*' or 'storage.*') are not allowed. For
	// more
	// information see
	// IAM Overview.
	Permissions []string `json:"permissions,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Permissions") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *TestIamPermissionsRequest) MarshalJSON() ([]byte, error) {
	type noMethod TestIamPermissionsRequest
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// TestIamPermissionsResponse: Response message for `TestIamPermissions`
// method.
type TestIamPermissionsResponse struct {
	// Permissions: A subset of `TestPermissionsRequest.permissions` that
	// the caller is
	// allowed.
	Permissions []string `json:"permissions,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Permissions") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *TestIamPermissionsResponse) MarshalJSON() ([]byte, error) {
	type noMethod TestIamPermissionsResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// method id "iam.iamPolicies.getPolicyDetails":

type IamPoliciesGetPolicyDetailsCall struct {
	s                       *Service
	getpolicydetailsrequest *GetPolicyDetailsRequest
	urlParams_              gensupport.URLParams
	ctx_                    context.Context
}

// GetPolicyDetails: Returns the current policy and the policies on the
// inherited resources
// the user has access to.
func (r *IamPoliciesService) GetPolicyDetails(getpolicydetailsrequest *GetPolicyDetailsRequest) *IamPoliciesGetPolicyDetailsCall {
	c := &IamPoliciesGetPolicyDetailsCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.getpolicydetailsrequest = getpolicydetailsrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *IamPoliciesGetPolicyDetailsCall) Fields(s ...googleapi.Field) *IamPoliciesGetPolicyDetailsCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *IamPoliciesGetPolicyDetailsCall) Context(ctx context.Context) *IamPoliciesGetPolicyDetailsCall {
	c.ctx_ = ctx
	return c
}

func (c *IamPoliciesGetPolicyDetailsCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.getpolicydetailsrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/iamPolicies:getPolicyDetails")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "iam.iamPolicies.getPolicyDetails" call.
// Exactly one of *GetPolicyDetailsResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *GetPolicyDetailsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *IamPoliciesGetPolicyDetailsCall) Do(opts ...googleapi.CallOption) (*GetPolicyDetailsResponse, error) {
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
	ret := &GetPolicyDetailsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns the current policy and the policies on the inherited resources\nthe user has access to.",
	//   "flatPath": "v1/iamPolicies:getPolicyDetails",
	//   "httpMethod": "POST",
	//   "id": "iam.iamPolicies.getPolicyDetails",
	//   "parameterOrder": [],
	//   "parameters": {},
	//   "path": "v1/iamPolicies:getPolicyDetails",
	//   "request": {
	//     "$ref": "GetPolicyDetailsRequest"
	//   },
	//   "response": {
	//     "$ref": "GetPolicyDetailsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "iam.projects.serviceAccounts.create":

type ProjectsServiceAccountsCreateCall struct {
	s                           *Service
	name                        string
	createserviceaccountrequest *CreateServiceAccountRequest
	urlParams_                  gensupport.URLParams
	ctx_                        context.Context
}

// Create: Creates a service account and returns it.
func (r *ProjectsServiceAccountsService) Create(name string, createserviceaccountrequest *CreateServiceAccountRequest) *ProjectsServiceAccountsCreateCall {
	c := &ProjectsServiceAccountsCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.createserviceaccountrequest = createserviceaccountrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsServiceAccountsCreateCall) Fields(s ...googleapi.Field) *ProjectsServiceAccountsCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsServiceAccountsCreateCall) Context(ctx context.Context) *ProjectsServiceAccountsCreateCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsServiceAccountsCreateCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.createserviceaccountrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}/serviceAccounts")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "iam.projects.serviceAccounts.create" call.
// Exactly one of *ServiceAccount or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *ServiceAccount.ServerResponse.Header or (if a response was returned
// at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsServiceAccountsCreateCall) Do(opts ...googleapi.CallOption) (*ServiceAccount, error) {
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
	ret := &ServiceAccount{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a service account and returns it.",
	//   "flatPath": "v1/projects/{projectsId}/serviceAccounts",
	//   "httpMethod": "POST",
	//   "id": "iam.projects.serviceAccounts.create",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The resource name of the project associated with the service\naccounts, such as \"projects/123\"",
	//       "location": "path",
	//       "pattern": "^projects/[^/]*$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}/serviceAccounts",
	//   "request": {
	//     "$ref": "CreateServiceAccountRequest"
	//   },
	//   "response": {
	//     "$ref": "ServiceAccount"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "iam.projects.serviceAccounts.delete":

type ProjectsServiceAccountsDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
}

// Delete: Deletes a service acount.
func (r *ProjectsServiceAccountsService) Delete(name string) *ProjectsServiceAccountsDeleteCall {
	c := &ProjectsServiceAccountsDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsServiceAccountsDeleteCall) Fields(s ...googleapi.Field) *ProjectsServiceAccountsDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsServiceAccountsDeleteCall) Context(ctx context.Context) *ProjectsServiceAccountsDeleteCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsServiceAccountsDeleteCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "iam.projects.serviceAccounts.delete" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsServiceAccountsDeleteCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
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
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes a service acount.",
	//   "flatPath": "v1/projects/{projectsId}/serviceAccounts/{serviceAccountsId}",
	//   "httpMethod": "DELETE",
	//   "id": "iam.projects.serviceAccounts.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The resource name of the service account in the format\n\"projects/{project}/serviceAccounts/{account}\".\nUsing '-' as a wildcard for the project, will infer the project from\nthe account. The account value can be the email address or the unique_id\nof the service account.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]*/serviceAccounts/[^/]*$",
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

// method id "iam.projects.serviceAccounts.get":

type ProjectsServiceAccountsGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// Get: Gets a ServiceAccount
func (r *ProjectsServiceAccountsService) Get(name string) *ProjectsServiceAccountsGetCall {
	c := &ProjectsServiceAccountsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsServiceAccountsGetCall) Fields(s ...googleapi.Field) *ProjectsServiceAccountsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsServiceAccountsGetCall) IfNoneMatch(entityTag string) *ProjectsServiceAccountsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsServiceAccountsGetCall) Context(ctx context.Context) *ProjectsServiceAccountsGetCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsServiceAccountsGetCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		req.Header.Set("If-None-Match", c.ifNoneMatch_)
	}
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "iam.projects.serviceAccounts.get" call.
// Exactly one of *ServiceAccount or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *ServiceAccount.ServerResponse.Header or (if a response was returned
// at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsServiceAccountsGetCall) Do(opts ...googleapi.CallOption) (*ServiceAccount, error) {
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
	ret := &ServiceAccount{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets a ServiceAccount",
	//   "flatPath": "v1/projects/{projectsId}/serviceAccounts/{serviceAccountsId}",
	//   "httpMethod": "GET",
	//   "id": "iam.projects.serviceAccounts.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The resource name of the service account in the format\n\"projects/{project}/serviceAccounts/{account}\".\nUsing '-' as a wildcard for the project, will infer the project from\nthe account. The account value can be the email address or the unique_id\nof the service account.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]*/serviceAccounts/[^/]*$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}",
	//   "response": {
	//     "$ref": "ServiceAccount"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "iam.projects.serviceAccounts.list":

type ProjectsServiceAccountsListCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// List: Lists service accounts for a project.
func (r *ProjectsServiceAccountsService) List(name string) *ProjectsServiceAccountsListCall {
	c := &ProjectsServiceAccountsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// PageSize sets the optional parameter "pageSize": Optional limit on
// the number of service accounts to include in the
// response. Further accounts can subsequently be obtained by including
// the
// [ListServiceAccountsResponse.next_page_token] in a subsequent
// request.
func (c *ProjectsServiceAccountsListCall) PageSize(pageSize int64) *ProjectsServiceAccountsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": Optional
// pagination token returned in an
// earlier
// [ListServiceAccountsResponse.next_page_token].
func (c *ProjectsServiceAccountsListCall) PageToken(pageToken string) *ProjectsServiceAccountsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// RemoveDeletedServiceAccounts sets the optional parameter
// "removeDeletedServiceAccounts": Do not list service accounts deleted
// from Gaia.
// <b><font color="red">DO NOT INCLUDE IN EXTERNAL
// DOCUMENTATION</font></b>.
func (c *ProjectsServiceAccountsListCall) RemoveDeletedServiceAccounts(removeDeletedServiceAccounts bool) *ProjectsServiceAccountsListCall {
	c.urlParams_.Set("removeDeletedServiceAccounts", fmt.Sprint(removeDeletedServiceAccounts))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsServiceAccountsListCall) Fields(s ...googleapi.Field) *ProjectsServiceAccountsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsServiceAccountsListCall) IfNoneMatch(entityTag string) *ProjectsServiceAccountsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsServiceAccountsListCall) Context(ctx context.Context) *ProjectsServiceAccountsListCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsServiceAccountsListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}/serviceAccounts")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		req.Header.Set("If-None-Match", c.ifNoneMatch_)
	}
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "iam.projects.serviceAccounts.list" call.
// Exactly one of *ListServiceAccountsResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *ListServiceAccountsResponse.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsServiceAccountsListCall) Do(opts ...googleapi.CallOption) (*ListServiceAccountsResponse, error) {
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
	ret := &ListServiceAccountsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists service accounts for a project.",
	//   "flatPath": "v1/projects/{projectsId}/serviceAccounts",
	//   "httpMethod": "GET",
	//   "id": "iam.projects.serviceAccounts.list",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The resource name of the project associated with the service\naccounts, such as \"projects/123\"",
	//       "location": "path",
	//       "pattern": "^projects/[^/]*$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Optional limit on the number of service accounts to include in the\nresponse. Further accounts can subsequently be obtained by including the\n[ListServiceAccountsResponse.next_page_token] in a subsequent request.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional pagination token returned in an earlier\n[ListServiceAccountsResponse.next_page_token].",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "removeDeletedServiceAccounts": {
	//       "description": "Do not list service accounts deleted from Gaia.\n\u003cb\u003e\u003cfont color=\"red\"\u003eDO NOT INCLUDE IN EXTERNAL DOCUMENTATION\u003c/font\u003e\u003c/b\u003e.",
	//       "location": "query",
	//       "type": "boolean"
	//     }
	//   },
	//   "path": "v1/{+name}/serviceAccounts",
	//   "response": {
	//     "$ref": "ListServiceAccountsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsServiceAccountsListCall) Pages(ctx context.Context, f func(*ListServiceAccountsResponse) error) error {
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

// method id "iam.projects.serviceAccounts.signBlob":

type ProjectsServiceAccountsSignBlobCall struct {
	s               *Service
	name            string
	signblobrequest *SignBlobRequest
	urlParams_      gensupport.URLParams
	ctx_            context.Context
}

// SignBlob: Signs a blob using a service account.
func (r *ProjectsServiceAccountsService) SignBlob(name string, signblobrequest *SignBlobRequest) *ProjectsServiceAccountsSignBlobCall {
	c := &ProjectsServiceAccountsSignBlobCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.signblobrequest = signblobrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsServiceAccountsSignBlobCall) Fields(s ...googleapi.Field) *ProjectsServiceAccountsSignBlobCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsServiceAccountsSignBlobCall) Context(ctx context.Context) *ProjectsServiceAccountsSignBlobCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsServiceAccountsSignBlobCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.signblobrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}:signBlob")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "iam.projects.serviceAccounts.signBlob" call.
// Exactly one of *SignBlobResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *SignBlobResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsServiceAccountsSignBlobCall) Do(opts ...googleapi.CallOption) (*SignBlobResponse, error) {
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
	ret := &SignBlobResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Signs a blob using a service account.",
	//   "flatPath": "v1/projects/{projectsId}/serviceAccounts/{serviceAccountsId}:signBlob",
	//   "httpMethod": "POST",
	//   "id": "iam.projects.serviceAccounts.signBlob",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The resource name of the service account in the format\n\"projects/{project}/serviceAccounts/{account}\".\nUsing '-' as a wildcard for the project, will infer the project from\nthe account. The account value can be the email address or the unique_id\nof the service account.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]*/serviceAccounts/[^/]*$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}:signBlob",
	//   "request": {
	//     "$ref": "SignBlobRequest"
	//   },
	//   "response": {
	//     "$ref": "SignBlobResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "iam.projects.serviceAccounts.update":

type ProjectsServiceAccountsUpdateCall struct {
	s              *Service
	name           string
	serviceaccount *ServiceAccount
	urlParams_     gensupport.URLParams
	ctx_           context.Context
}

// Update: Updates a service account.
//
// Currently the only updatable fields are 'display_name' and
// 'description'.
// The 'etag' is mandatory.
func (r *ProjectsServiceAccountsService) Update(name string, serviceaccount *ServiceAccount) *ProjectsServiceAccountsUpdateCall {
	c := &ProjectsServiceAccountsUpdateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.serviceaccount = serviceaccount
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsServiceAccountsUpdateCall) Fields(s ...googleapi.Field) *ProjectsServiceAccountsUpdateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsServiceAccountsUpdateCall) Context(ctx context.Context) *ProjectsServiceAccountsUpdateCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsServiceAccountsUpdateCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.serviceaccount)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "iam.projects.serviceAccounts.update" call.
// Exactly one of *ServiceAccount or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *ServiceAccount.ServerResponse.Header or (if a response was returned
// at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsServiceAccountsUpdateCall) Do(opts ...googleapi.CallOption) (*ServiceAccount, error) {
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
	ret := &ServiceAccount{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates a service account.\n\nCurrently the only updatable fields are 'display_name' and 'description'.\nThe 'etag' is mandatory.",
	//   "flatPath": "v1/projects/{projectsId}/serviceAccounts/{serviceAccountsId}",
	//   "httpMethod": "PUT",
	//   "id": "iam.projects.serviceAccounts.update",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The resource name of the service account in the format\n\"projects/{project}/serviceAccounts/{account}\".\n\nIn requests using '-' as a wildcard for the project, will infer the project\nfrom the account and the account value can be the email address or the\nunique_id of the service account.\n\nIn responses the resource name will always be in the format\n\"projects/{project}/serviceAccounts/{email}\".",
	//       "location": "path",
	//       "pattern": "^projects/[^/]*/serviceAccounts/[^/]*$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}",
	//   "request": {
	//     "$ref": "ServiceAccount"
	//   },
	//   "response": {
	//     "$ref": "ServiceAccount"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "iam.projects.serviceAccounts.keys.create":

type ProjectsServiceAccountsKeysCreateCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
}

// Create: Creates a service account key and returns it.
func (r *ProjectsServiceAccountsKeysService) Create(name string) *ProjectsServiceAccountsKeysCreateCall {
	c := &ProjectsServiceAccountsKeysCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// PrivateKeyType sets the optional parameter "privateKeyType": The type
// of the key requested. GOOGLE_CREDENTIALS is the default key type.
//
// Possible values:
//   "TYPE_UNSPECIFIED"
//   "TYPE_PKCS12_FILE"
//   "TYPE_GOOGLE_CREDENTIALS_FILE"
func (c *ProjectsServiceAccountsKeysCreateCall) PrivateKeyType(privateKeyType string) *ProjectsServiceAccountsKeysCreateCall {
	c.urlParams_.Set("privateKeyType", privateKeyType)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsServiceAccountsKeysCreateCall) Fields(s ...googleapi.Field) *ProjectsServiceAccountsKeysCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsServiceAccountsKeysCreateCall) Context(ctx context.Context) *ProjectsServiceAccountsKeysCreateCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsServiceAccountsKeysCreateCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}/keys")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "iam.projects.serviceAccounts.keys.create" call.
// Exactly one of *ServiceAccountKey or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ServiceAccountKey.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsServiceAccountsKeysCreateCall) Do(opts ...googleapi.CallOption) (*ServiceAccountKey, error) {
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
	ret := &ServiceAccountKey{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a service account key and returns it.",
	//   "flatPath": "v1/projects/{projectsId}/serviceAccounts/{serviceAccountsId}/keys",
	//   "httpMethod": "POST",
	//   "id": "iam.projects.serviceAccounts.keys.create",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The resource name of the service account in the format\n\"projects/{project}/serviceAccounts/{account}\".\nUsing '-' as a wildcard for the project, will infer the project from\nthe account. The account value can be the email address or the unique_id\nof the service account.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]*/serviceAccounts/[^/]*$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "privateKeyType": {
	//       "description": "The type of the key requested. GOOGLE_CREDENTIALS is the default key type.",
	//       "enum": [
	//         "TYPE_UNSPECIFIED",
	//         "TYPE_PKCS12_FILE",
	//         "TYPE_GOOGLE_CREDENTIALS_FILE"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}/keys",
	//   "response": {
	//     "$ref": "ServiceAccountKey"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "iam.projects.serviceAccounts.keys.delete":

type ProjectsServiceAccountsKeysDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
}

// Delete: Deletes a service account key.
func (r *ProjectsServiceAccountsKeysService) Delete(name string) *ProjectsServiceAccountsKeysDeleteCall {
	c := &ProjectsServiceAccountsKeysDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsServiceAccountsKeysDeleteCall) Fields(s ...googleapi.Field) *ProjectsServiceAccountsKeysDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsServiceAccountsKeysDeleteCall) Context(ctx context.Context) *ProjectsServiceAccountsKeysDeleteCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsServiceAccountsKeysDeleteCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "iam.projects.serviceAccounts.keys.delete" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsServiceAccountsKeysDeleteCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
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
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes a service account key.",
	//   "flatPath": "v1/projects/{projectsId}/serviceAccounts/{serviceAccountsId}/keys/{keysId}",
	//   "httpMethod": "DELETE",
	//   "id": "iam.projects.serviceAccounts.keys.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The resource name of the service account key in the format\n\"projects/{project}/serviceAccounts/{account}/keys/{key}\".\nUsing '-' as a wildcard for the project will infer the project from\nthe account. The account value can be the email address or the unique_id\nof the service account.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]*/serviceAccounts/[^/]*/keys/[^/]*$",
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

// method id "iam.projects.serviceAccounts.keys.get":

type ProjectsServiceAccountsKeysGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// Get: Gets the ServiceAccountKey
// by key id.
func (r *ProjectsServiceAccountsKeysService) Get(name string) *ProjectsServiceAccountsKeysGetCall {
	c := &ProjectsServiceAccountsKeysGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsServiceAccountsKeysGetCall) Fields(s ...googleapi.Field) *ProjectsServiceAccountsKeysGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsServiceAccountsKeysGetCall) IfNoneMatch(entityTag string) *ProjectsServiceAccountsKeysGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsServiceAccountsKeysGetCall) Context(ctx context.Context) *ProjectsServiceAccountsKeysGetCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsServiceAccountsKeysGetCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		req.Header.Set("If-None-Match", c.ifNoneMatch_)
	}
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "iam.projects.serviceAccounts.keys.get" call.
// Exactly one of *ServiceAccountKey or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ServiceAccountKey.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsServiceAccountsKeysGetCall) Do(opts ...googleapi.CallOption) (*ServiceAccountKey, error) {
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
	ret := &ServiceAccountKey{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets the ServiceAccountKey\nby key id.",
	//   "flatPath": "v1/projects/{projectsId}/serviceAccounts/{serviceAccountsId}/keys/{keysId}",
	//   "httpMethod": "GET",
	//   "id": "iam.projects.serviceAccounts.keys.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The resource name of the service account key in the format\n\"projects/{project}/serviceAccounts/{account}/keys/{key}\".\nUsing '-' as a wildcard for the project will infer the project from\nthe account. The account value can be the email address or the unique_id\nof the service account.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]*/serviceAccounts/[^/]*/keys/[^/]*$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}",
	//   "response": {
	//     "$ref": "ServiceAccountKey"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "iam.projects.serviceAccounts.keys.list":

type ProjectsServiceAccountsKeysListCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// List: Lists service account keys
func (r *ProjectsServiceAccountsKeysService) List(name string) *ProjectsServiceAccountsKeysListCall {
	c := &ProjectsServiceAccountsKeysListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// KeyTypes sets the optional parameter "keyTypes": The type of keys the
// user wants to list. If empty, all
// key types are included in the response. Duplicate key
// types are not allowed.
//
// Possible values:
//   "KEY_TYPE_UNSPECIFIED"
//   "USER_MANAGED"
//   "SYSTEM_MANAGED"
func (c *ProjectsServiceAccountsKeysListCall) KeyTypes(keyTypes ...string) *ProjectsServiceAccountsKeysListCall {
	c.urlParams_.SetMulti("keyTypes", append([]string{}, keyTypes...))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsServiceAccountsKeysListCall) Fields(s ...googleapi.Field) *ProjectsServiceAccountsKeysListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsServiceAccountsKeysListCall) IfNoneMatch(entityTag string) *ProjectsServiceAccountsKeysListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsServiceAccountsKeysListCall) Context(ctx context.Context) *ProjectsServiceAccountsKeysListCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsServiceAccountsKeysListCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}/keys")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		req.Header.Set("If-None-Match", c.ifNoneMatch_)
	}
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "iam.projects.serviceAccounts.keys.list" call.
// Exactly one of *ListServiceAccountKeysResponse or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *ListServiceAccountKeysResponse.ServerResponse.Header or (if a
// response was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsServiceAccountsKeysListCall) Do(opts ...googleapi.CallOption) (*ListServiceAccountKeysResponse, error) {
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
	ret := &ListServiceAccountKeysResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists service account keys",
	//   "flatPath": "v1/projects/{projectsId}/serviceAccounts/{serviceAccountsId}/keys",
	//   "httpMethod": "GET",
	//   "id": "iam.projects.serviceAccounts.keys.list",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "keyTypes": {
	//       "description": "The type of keys the user wants to list. If empty, all\nkey types are included in the response. Duplicate key\ntypes are not allowed.",
	//       "enum": [
	//         "KEY_TYPE_UNSPECIFIED",
	//         "USER_MANAGED",
	//         "SYSTEM_MANAGED"
	//       ],
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "name": {
	//       "description": "The resource name of the service account in the format\n\"projects/{project}/serviceAccounts/{account}\".\nUsing '-' as a wildcard for the project, will infer the project from\nthe account. The account value can be the email address or the unique_id\nof the service account.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]*/serviceAccounts/[^/]*$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}/keys",
	//   "response": {
	//     "$ref": "ListServiceAccountKeysResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "iam.roles.queryGrantableRoles":

type RolesQueryGrantableRolesCall struct {
	s                          *Service
	querygrantablerolesrequest *QueryGrantableRolesRequest
	urlParams_                 gensupport.URLParams
	ctx_                       context.Context
}

// QueryGrantableRoles: Queries roles that can be granted on a
// particular resource.
func (r *RolesService) QueryGrantableRoles(querygrantablerolesrequest *QueryGrantableRolesRequest) *RolesQueryGrantableRolesCall {
	c := &RolesQueryGrantableRolesCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.querygrantablerolesrequest = querygrantablerolesrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RolesQueryGrantableRolesCall) Fields(s ...googleapi.Field) *RolesQueryGrantableRolesCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *RolesQueryGrantableRolesCall) Context(ctx context.Context) *RolesQueryGrantableRolesCall {
	c.ctx_ = ctx
	return c
}

func (c *RolesQueryGrantableRolesCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.querygrantablerolesrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/roles:queryGrantableRoles")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "iam.roles.queryGrantableRoles" call.
// Exactly one of *QueryGrantableRolesResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *QueryGrantableRolesResponse.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *RolesQueryGrantableRolesCall) Do(opts ...googleapi.CallOption) (*QueryGrantableRolesResponse, error) {
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
	ret := &QueryGrantableRolesResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Queries roles that can be granted on a particular resource.",
	//   "flatPath": "v1/roles:queryGrantableRoles",
	//   "httpMethod": "POST",
	//   "id": "iam.roles.queryGrantableRoles",
	//   "parameterOrder": [],
	//   "parameters": {},
	//   "path": "v1/roles:queryGrantableRoles",
	//   "request": {
	//     "$ref": "QueryGrantableRolesRequest"
	//   },
	//   "response": {
	//     "$ref": "QueryGrantableRolesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "iam.getIamPolicy":

type V1GetIamPolicyCall struct {
	s          *Service
	resource   string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
}

// GetIamPolicy: Returns the IAM access control policy for specified IAM
// resource.
func (r *V1Service) GetIamPolicy(resource string) *V1GetIamPolicyCall {
	c := &V1GetIamPolicyCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.resource = resource
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *V1GetIamPolicyCall) Fields(s ...googleapi.Field) *V1GetIamPolicyCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *V1GetIamPolicyCall) Context(ctx context.Context) *V1GetIamPolicyCall {
	c.ctx_ = ctx
	return c
}

func (c *V1GetIamPolicyCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+resource}:getIamPolicy")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "iam.getIamPolicy" call.
// Exactly one of *Policy or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Policy.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *V1GetIamPolicyCall) Do(opts ...googleapi.CallOption) (*Policy, error) {
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
	ret := &Policy{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns the IAM access control policy for specified IAM resource.",
	//   "flatPath": "v1/{v1Id}:getIamPolicy",
	//   "httpMethod": "POST",
	//   "id": "iam.getIamPolicy",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "REQUIRED: The resource for which the policy is being requested.\n`resource` is usually specified as a path, such as\n`projects/*project*/zones/*zone*/disks/*disk*`.\n\nThe format for the path specified in this value is resource specific and\nis specified in the `getIamPolicy` documentation.",
	//       "location": "path",
	//       "pattern": "^.*$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+resource}:getIamPolicy",
	//   "response": {
	//     "$ref": "Policy"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "iam.setIamPolicy":

type V1SetIamPolicyCall struct {
	s                   *Service
	resource            string
	setiampolicyrequest *SetIamPolicyRequest
	urlParams_          gensupport.URLParams
	ctx_                context.Context
}

// SetIamPolicy: Sets the IAM access control policy for the specified
// IAM resource.
func (r *V1Service) SetIamPolicy(resource string, setiampolicyrequest *SetIamPolicyRequest) *V1SetIamPolicyCall {
	c := &V1SetIamPolicyCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.resource = resource
	c.setiampolicyrequest = setiampolicyrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *V1SetIamPolicyCall) Fields(s ...googleapi.Field) *V1SetIamPolicyCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *V1SetIamPolicyCall) Context(ctx context.Context) *V1SetIamPolicyCall {
	c.ctx_ = ctx
	return c
}

func (c *V1SetIamPolicyCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.setiampolicyrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+resource}:setIamPolicy")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "iam.setIamPolicy" call.
// Exactly one of *Policy or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Policy.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *V1SetIamPolicyCall) Do(opts ...googleapi.CallOption) (*Policy, error) {
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
	ret := &Policy{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Sets the IAM access control policy for the specified IAM resource.",
	//   "flatPath": "v1/{v1Id}:setIamPolicy",
	//   "httpMethod": "POST",
	//   "id": "iam.setIamPolicy",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "REQUIRED: The resource for which the policy is being specified.\n`resource` is usually specified as a path, such as\n`projects/*project*/zones/*zone*/disks/*disk*`.\n\nThe format for the path specified in this value is resource specific and\nis specified in the `setIamPolicy` documentation.",
	//       "location": "path",
	//       "pattern": "^.*$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+resource}:setIamPolicy",
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

// method id "iam.testIamPermissions":

type V1TestIamPermissionsCall struct {
	s                         *Service
	resource                  string
	testiampermissionsrequest *TestIamPermissionsRequest
	urlParams_                gensupport.URLParams
	ctx_                      context.Context
}

// TestIamPermissions: Tests the specified permissions against the IAM
// access control policy
// for the specified IAM resource.
func (r *V1Service) TestIamPermissions(resource string, testiampermissionsrequest *TestIamPermissionsRequest) *V1TestIamPermissionsCall {
	c := &V1TestIamPermissionsCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.resource = resource
	c.testiampermissionsrequest = testiampermissionsrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *V1TestIamPermissionsCall) Fields(s ...googleapi.Field) *V1TestIamPermissionsCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *V1TestIamPermissionsCall) Context(ctx context.Context) *V1TestIamPermissionsCall {
	c.ctx_ = ctx
	return c
}

func (c *V1TestIamPermissionsCall) doRequest(alt string) (*http.Response, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.testiampermissionsrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+resource}:testIamPermissions")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "iam.testIamPermissions" call.
// Exactly one of *TestIamPermissionsResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *TestIamPermissionsResponse.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *V1TestIamPermissionsCall) Do(opts ...googleapi.CallOption) (*TestIamPermissionsResponse, error) {
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
	ret := &TestIamPermissionsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Tests the specified permissions against the IAM access control policy\nfor the specified IAM resource.",
	//   "flatPath": "v1/{v1Id}:testIamPermissions",
	//   "httpMethod": "POST",
	//   "id": "iam.testIamPermissions",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "REQUIRED: The resource for which the policy detail is being requested.\n`resource` is usually specified as a path, such as\n`projects/*project*/zones/*zone*/disks/*disk*`.\n\nThe format for the path specified in this value is resource specific and\nis specified in the `testIamPermissions` documentation.",
	//       "location": "path",
	//       "pattern": "^.*$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+resource}:testIamPermissions",
	//   "request": {
	//     "$ref": "TestIamPermissionsRequest"
	//   },
	//   "response": {
	//     "$ref": "TestIamPermissionsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}
