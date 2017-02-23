// Package containeranalysis provides access to the Container Analysis API.
//
// See https://cloud.google.com/container-analysis/api/reference/rest/
//
// Usage example:
//
//   import "github.com/skelterjohn/gcloud_apis/clients/containeranalysis/v1alpha1"
//   ...
//   containeranalysisService, err := containeranalysis.New(oauthHttpClient)
package containeranalysis

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

const apiId = "containeranalysis:v1alpha1"
const apiName = "containeranalysis"
const apiVersion = "v1alpha1"
const basePath = "https://containeranalysis.googleapis.com/"

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
	s.Projects = NewProjectsService(s)
	s.Providers = NewProvidersService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Projects *ProjectsService

	Providers *ProvidersService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewProjectsService(s *Service) *ProjectsService {
	rs := &ProjectsService{s: s}
	rs.Notes = NewProjectsNotesService(s)
	rs.Occurrences = NewProjectsOccurrencesService(s)
	return rs
}

type ProjectsService struct {
	s *Service

	Notes *ProjectsNotesService

	Occurrences *ProjectsOccurrencesService
}

func NewProjectsNotesService(s *Service) *ProjectsNotesService {
	rs := &ProjectsNotesService{s: s}
	rs.Occurrences = NewProjectsNotesOccurrencesService(s)
	return rs
}

type ProjectsNotesService struct {
	s *Service

	Occurrences *ProjectsNotesOccurrencesService
}

func NewProjectsNotesOccurrencesService(s *Service) *ProjectsNotesOccurrencesService {
	rs := &ProjectsNotesOccurrencesService{s: s}
	return rs
}

type ProjectsNotesOccurrencesService struct {
	s *Service
}

func NewProjectsOccurrencesService(s *Service) *ProjectsOccurrencesService {
	rs := &ProjectsOccurrencesService{s: s}
	return rs
}

type ProjectsOccurrencesService struct {
	s *Service
}

func NewProvidersService(s *Service) *ProvidersService {
	rs := &ProvidersService{s: s}
	rs.Notes = NewProvidersNotesService(s)
	return rs
}

type ProvidersService struct {
	s *Service

	Notes *ProvidersNotesService
}

func NewProvidersNotesService(s *Service) *ProvidersNotesService {
	rs := &ProvidersNotesService{s: s}
	rs.Occurrences = NewProvidersNotesOccurrencesService(s)
	return rs
}

type ProvidersNotesService struct {
	s *Service

	Occurrences *ProvidersNotesOccurrencesService
}

func NewProvidersNotesOccurrencesService(s *Service) *ProvidersNotesOccurrencesService {
	rs := &ProvidersNotesOccurrencesService{s: s}
	return rs
}

type ProvidersNotesOccurrencesService struct {
	s *Service
}

// AliasContext: An alias to a repo revision.
type AliasContext struct {
	// Kind: The alias kind.
	//
	// Possible values:
	//   "ANY" - Do not use.
	//   "FIXED" - Git tag
	//   "MOVABLE" - Git branch
	//   "OTHER" - OTHER is used to specify non-standard aliases, those not
	// of the kinds
	// above. For example, if a Git repo has a ref named "refs/foo/bar",
	// it
	// is considered to be of kind OTHER.
	Kind string `json:"kind,omitempty"`

	// Name: The alias name.
	Name string `json:"name,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Kind") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *AliasContext) MarshalJSON() ([]byte, error) {
	type noMethod AliasContext
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// Artifact: Artifact destribes a build product.
type Artifact struct {
	// Checksum: Hash or checksum value of a binary, or Docker Registry 2.0
	// digest of a
	// container.
	Checksum string `json:"checksum,omitempty"`

	// Id: Artifact ID, if any; for container images, this will be a URL by
	// digest
	// like gcr.io/projectID/imagename@sha256:123456
	Id string `json:"id,omitempty"`

	// Name: Name of the artifact. This may be the path to a binary or jar
	// file, or in
	// the case of a container build, the name used to push the container
	// image to
	// Google Container Registry, as presented to `docker push`.
	Name string `json:"name,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Checksum") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Artifact) MarshalJSON() ([]byte, error) {
	type noMethod Artifact
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// AuditConfig: Specifies the audit configuration for a service.
// It consists of which permission types are logged, and what
// identities, if
// any, are exempted from logging.
// An AuditConifg must have one or more AuditLogConfigs.
//
// If there are AuditConfigs for both `allServices` and a specific
// service,
// the union of the two AuditConfigs is used for that service: the
// log_types
// specified in each AuditConfig are enabled, and the exempted_members
// in each
// AuditConfig are exempted.
// Example Policy with multiple AuditConfigs:
// {
//   "audit_configs": [
//     {
//       "service": "allServices"
//       "audit_log_configs": [
//         {
//           "log_type": "DATA_READ",
//           "exempted_members": [
//             "user:foo@gmail.com"
//           ]
//         },
//         {
//           "log_type": "DATA_WRITE",
//         },
//         {
//           "log_type": "ADMIN_READ",
//         }
//       ]
//     },
//     {
//       "service": "fooservice@googleapis.com"
//       "audit_log_configs": [
//         {
//           "log_type": "DATA_READ",
//         },
//         {
//           "log_type": "DATA_WRITE",
//           "exempted_members": [
//             "user:bar@gmail.com"
//           ]
//         }
//       ]
//     }
//   ]
// }
// For fooservice, this policy enables DATA_READ, DATA_WRITE and
// ADMIN_READ
// logging. It also exempts foo@gmail.com from DATA_READ logging,
// and
// bar@gmail.com from DATA_WRITE logging.
type AuditConfig struct {
	// AuditLogConfigs: The configuration for logging of each type of
	// permission.
	// Next ID: 4
	AuditLogConfigs []*AuditLogConfig `json:"auditLogConfigs,omitempty"`

	ExemptedMembers []string `json:"exemptedMembers,omitempty"`

	// Service: Specifies a service that will be enabled for audit
	// logging.
	// For example, `resourcemanager`, `storage`, `compute`.
	// `allServices` is a special value that covers all services.
	Service string `json:"service,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AuditLogConfigs") to
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

// AuditLogConfig: Provides the configuration for logging a type of
// permissions.
// Example:
//
//     {
//       "audit_log_configs": [
//         {
//           "log_type": "DATA_READ",
//           "exempted_members": [
//             "user:foo@gmail.com"
//           ]
//         },
//         {
//           "log_type": "DATA_WRITE",
//         }
//       ]
//     }
//
// This enables 'DATA_READ' and 'DATA_WRITE' logging, while
// exempting
// foo@gmail.com from DATA_READ logging.
type AuditLogConfig struct {
	// ExemptedMembers: Specifies the identities that do not cause logging
	// for this type of
	// permission.
	// Follows the same format of Binding.members.
	ExemptedMembers []string `json:"exemptedMembers,omitempty"`

	// LogType: The log type that this config enables.
	//
	// Possible values:
	//   "LOG_TYPE_UNSPECIFIED" - Default case. Should never be this.
	//   "ADMIN_READ" - Admin reads. Example: CloudIAM getIamPolicy
	//   "DATA_WRITE" - Data writes. Example: CloudSQL Users create
	//   "DATA_READ" - Data reads. Example: CloudSQL Users list
	LogType string `json:"logType,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ExemptedMembers") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *AuditLogConfig) MarshalJSON() ([]byte, error) {
	type noMethod AuditLogConfig
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// Basis: Basis describes the base image portion (Note) of the
// DockerImage
// relationship.  Linked occurrences are derived from this or
// an
// equivalent image via:
//   FROM <Basis.resource_url>
// Or an equivalent reference, e.g. a tag of the resource_url.
type Basis struct {
	// Fingerprint: The fingerprint of the base image
	Fingerprint *Fingerprint `json:"fingerprint,omitempty"`

	// ResourceUrl: The resource_url for the resource representing the basis
	// of
	// associated occurrence images.
	ResourceUrl string `json:"resourceUrl,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Fingerprint") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Basis) MarshalJSON() ([]byte, error) {
	type noMethod Basis
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
	//    users of that domain. For example, `google.com` or
	// `example.com`.
	//
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

// BuildDetails: Message encapsulating build provenance details
type BuildDetails struct {
	// Provenance: The actual provenance
	Provenance *BuildProvenance `json:"provenance,omitempty"`

	// ProvenanceBytes: Serialized json representation of the provenance,
	// used in generating the
	// BuildSignature in the corresponding Result. After verifying the
	// signature,
	// provenance_bytes can be unmarshalled and compared to the provenance
	// to
	// confirm that it is unchanged. A base64-encoded string representation
	// of the
	// provenance bytes is used for the signature in order to interoperate
	// with
	// openssl which expects this format for signature verification.
	//
	// The serialized form is captured both to avoid ambiguity in how
	// the
	// provenance is marshalled to json as well to prevent incompatibilities
	// with
	// future changes.
	ProvenanceBytes string `json:"provenanceBytes,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Provenance") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *BuildDetails) MarshalJSON() ([]byte, error) {
	type noMethod BuildDetails
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// BuildProvenance: Provenance of a build. Contains all information
// needed to verify the full
// details about the build from source to completion.
type BuildProvenance struct {
	// BuildOptions: Special options applied to this build. This is a
	// catch-all field where
	// build providers can enter any desired additional details.
	BuildOptions map[string]string `json:"buildOptions,omitempty"`

	// BuilderVersion: Version string of the builder at the time this build
	// was executed.
	BuilderVersion string `json:"builderVersion,omitempty"`

	// BuiltArtifacts: Output of the build.
	BuiltArtifacts []*Artifact `json:"builtArtifacts,omitempty"`

	// Commands: Commands requested by the build.
	Commands []*Command `json:"commands,omitempty"`

	// CreateTime: Time at which the build was created.
	CreateTime string `json:"createTime,omitempty"`

	// Creator: E-mail address of the user who initiated this build. Note
	// that this was the
	// user's e-mail address at the time the build was initiated; this
	// address may
	// not represent the same end-user for all time.
	Creator string `json:"creator,omitempty"`

	// FinishTime: Time at whihc execution of the build was finished.
	FinishTime string `json:"finishTime,omitempty"`

	// Id: Unique identifier of the build.
	Id string `json:"id,omitempty"`

	// LogsBucket: Google Cloud Storage bucket where logs were written.
	LogsBucket string `json:"logsBucket,omitempty"`

	// ProjectId: ID of the project.
	ProjectId string `json:"projectId,omitempty"`

	// SourceProvenance: Details of the Source input to the build.
	SourceProvenance *Source `json:"sourceProvenance,omitempty"`

	// StartTime: Time at which execution of the build was started.
	StartTime string `json:"startTime,omitempty"`

	// TriggerId: Trigger identifier if the build was triggered
	// automatically; empty if not.
	TriggerId string `json:"triggerId,omitempty"`

	// ForceSendFields is a list of field names (e.g. "BuildOptions") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *BuildProvenance) MarshalJSON() ([]byte, error) {
	type noMethod BuildProvenance
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// BuildSignature: Message encapsulating signature of the verified build
type BuildSignature struct {
	// PublicKey: Public key of the builder which can be used to verify that
	// related
	// Findings are valid and unchanged. To verify the signature, place
	// the
	// contents of this field into a file (public.pem). The signature field
	// is
	// base64-decoded into its binary representation in signature.bin, and
	// the
	// provenance bytes from BuildDetails are base64-decoded into a
	// binary
	// representation in signed.bin. Openssl can then verify the
	// signature:
	// 'openssl sha256 -verify public.pem -signature signature.bin
	// signed.bin'
	PublicKey string `json:"publicKey,omitempty"`

	// Signature: Signature of the related BuildProvenance, encoded in a
	// base64 string.
	Signature string `json:"signature,omitempty"`

	// ForceSendFields is a list of field names (e.g. "PublicKey") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *BuildSignature) MarshalJSON() ([]byte, error) {
	type noMethod BuildSignature
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// BuildType: Note holding the version of the provider's builder and the
// signature of
// the provenance message in linked BuildDetails.
type BuildType struct {
	// BuilderVersion: Version of the builder which produced this Note.
	BuilderVersion string `json:"builderVersion,omitempty"`

	// Signature: Signature of the build in Occurrences pointing to the Note
	// containing this
	// BuilderDetails.
	Signature *BuildSignature `json:"signature,omitempty"`

	// ForceSendFields is a list of field names (e.g. "BuilderVersion") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *BuildType) MarshalJSON() ([]byte, error) {
	type noMethod BuildType
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// CloudAuditOptions: Write a Cloud Audit log
type CloudAuditOptions struct {
}

// CloudRepoSourceContext: A CloudRepoSourceContext denotes a particular
// revision in a cloud
// repo (a repo hosted by the Google Cloud Platform).
type CloudRepoSourceContext struct {
	// AliasContext: An alias, which may be a branch or tag.
	AliasContext *AliasContext `json:"aliasContext,omitempty"`

	// AliasName: The name of an alias (branch, tag, etc.).
	AliasName string `json:"aliasName,omitempty"`

	// RepoId: The ID of the repo.
	RepoId *RepoId `json:"repoId,omitempty"`

	// RevisionId: A revision ID.
	RevisionId string `json:"revisionId,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AliasContext") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *CloudRepoSourceContext) MarshalJSON() ([]byte, error) {
	type noMethod CloudRepoSourceContext
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// CloudWorkspaceId: A CloudWorkspaceId is a unique identifier for a
// cloud workspace.
// A cloud workspace is a place associated with a repo where modified
// files
// can be stored before they are committed.
type CloudWorkspaceId struct {
	// Name: The unique name of the workspace within the repo.  This is the
	// name
	// chosen by the client in the Source API's CreateWorkspace method.
	Name string `json:"name,omitempty"`

	// RepoId: The ID of the repo containing the workspace.
	RepoId *RepoId `json:"repoId,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Name") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *CloudWorkspaceId) MarshalJSON() ([]byte, error) {
	type noMethod CloudWorkspaceId
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// CloudWorkspaceSourceContext: A CloudWorkspaceSourceContext denotes a
// workspace at a particular snapshot.
type CloudWorkspaceSourceContext struct {
	// SnapshotId: The ID of the snapshot.
	// An empty snapshot_id refers to the most recent snapshot.
	SnapshotId string `json:"snapshotId,omitempty"`

	// WorkspaceId: The ID of the workspace.
	WorkspaceId *CloudWorkspaceId `json:"workspaceId,omitempty"`

	// ForceSendFields is a list of field names (e.g. "SnapshotId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *CloudWorkspaceSourceContext) MarshalJSON() ([]byte, error) {
	type noMethod CloudWorkspaceSourceContext
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// Command: Command describes a step performed as part of the build
// pipeline.
type Command struct {
	// Args: Command-line arguments used when executing this Command.
	Args []string `json:"args,omitempty"`

	// Dir: Working directory (relative to project source root) used when
	// running
	// this Command.
	Dir string `json:"dir,omitempty"`

	// Env: Environment variables set before running this Command.
	Env []string `json:"env,omitempty"`

	// Id: Optional unique identifier for this Command, used in wait_for to
	// reference
	// this Command as a dependency.
	Id string `json:"id,omitempty"`

	// Name: Name of the command, as presented on the command line, or if
	// the command is
	// packaged as a Docker container, as presented to `docker pull`.
	Name string `json:"name,omitempty"`

	// WaitFor: The ID(s) of the Command(s) that this Command depends on.
	WaitFor []string `json:"waitFor,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Args") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Command) MarshalJSON() ([]byte, error) {
	type noMethod Command
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// Condition: A condition to be met.
type Condition struct {
	// Iam: Trusted attributes supplied by the IAM system.
	//
	// Possible values:
	//   "NO_ATTR" - Default non-attribute.
	//   "AUTHORITY" - Either principal or (if present) authority selector.
	//   "ATTRIBUTION" - The principal (even if an authority selector is
	// present), which
	// must only be used for attribution, not authorization.
	//   "APPROVER" - An approver (distinct from the requester) that has
	// authorized this
	// request.
	// When used with IN, the condition indicates that one of the
	// approvers
	// associated with the request matches the specified principal, or is
	// a
	// member of the specified group. Approvers can only grant
	// additional
	// access, and are thus only used in a strictly positive context
	// (e.g. ALLOW/IN or DENY/NOT_IN).
	// See: go/rpc-security-policy-dynamicauth.
	Iam string `json:"iam,omitempty"`

	// Op: An operator to apply the subject with.
	//
	// Possible values:
	//   "NO_OP" - Default no-op.
	//   "EQUALS" - DEPRECATED. Use IN instead.
	//   "NOT_EQUALS" - DEPRECATED. Use NOT_IN instead.
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

	// Value: DEPRECATED. Use 'values' instead.
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

// DataAccessOptions: Write a Data Access (Gin) log
type DataAccessOptions struct {
}

// Derived: Derived describes the derived image portion (Occurrence) of
// the
// DockerImage relationship.  This image would be produced from a
// Dockerfile
// with FROM <DockerImage.Basis in attached Note>.
type Derived struct {
	// BaseResourceUrl: This contains the base image url for the derived
	// image Occurrence
	// @OutputOnly
	BaseResourceUrl string `json:"baseResourceUrl,omitempty"`

	// Distance: The number of layers by which this image differs from
	// the associated image basis.
	// @OutputOnly
	Distance int64 `json:"distance,omitempty"`

	// Fingerprint: The fingerprint of the derived image
	Fingerprint *Fingerprint `json:"fingerprint,omitempty"`

	// LayerInfo: This contains layer-specific metadata, if populated it
	// has length “distance” and is ordered with [distance] being
	// the
	// layer immediately following the base image and [1]
	// being the final layer.
	LayerInfo []*Layer `json:"layerInfo,omitempty"`

	// ForceSendFields is a list of field names (e.g. "BaseResourceUrl") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Derived) MarshalJSON() ([]byte, error) {
	type noMethod Derived
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// Detail: Identifies all occurences of this vulnerability in the
// package for a
// specific distro/location
// For example: glibc in cpe:/o:debian:debian_linux:8 for versions 2.1 -
// 2.2
type Detail struct {
	// CpeUri: The cpe_uri in [cpe format]
	// (https://cpe.mitre.org/specification/) in
	// which the vulnerability manifests.  Examples include distro or
	// storage
	// location for vulnerable jar.
	// This field can be used as a filter in list requests.
	CpeUri string `json:"cpeUri,omitempty"`

	// Description: A vendor-specific description of this note.
	Description string `json:"description,omitempty"`

	// FixedLocation: The fix for this specific package version.
	FixedLocation *VulnerabilityLocation `json:"fixedLocation,omitempty"`

	// MaxAffectedVersion: The max version of the package in which the
	// vulnerability exists.
	// This field can be used as a filter in list requests.
	MaxAffectedVersion *Version `json:"maxAffectedVersion,omitempty"`

	// MinAffectedVersion: The min version of the package in which the
	// vulnerability exists.
	MinAffectedVersion *Version `json:"minAffectedVersion,omitempty"`

	// Package: The name of the package where the vulnerability was
	// found.
	// This field can be used as a filter in list requests.
	Package string `json:"package,omitempty"`

	// SeverityName: The severity (eg: distro assigned severity) for this
	// vulnerability.
	SeverityName string `json:"severityName,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CpeUri") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Detail) MarshalJSON() ([]byte, error) {
	type noMethod Detail
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// Distribution: This represents a particular channel of distribution
// for a given package.
// e.g. Debian's jessie-backports dpkg mirror
type Distribution struct {
	// Architecture: The CPU architecture for which packages in this
	// distribution
	// channel were built
	//
	// Possible values:
	//   "UNKNOWN" - Unknown architecture
	//   "X86" - X86 architecture
	//   "X64" - x64 architecture
	Architecture string `json:"architecture,omitempty"`

	// CpeUri: The cpe_uri in [cpe
	// format](https://cpe.mitre.org/specification/)
	// denoting the package manager version distributing a package.
	CpeUri string `json:"cpeUri,omitempty"`

	// Description: The distribution channel-specific description of this
	// package.
	Description string `json:"description,omitempty"`

	// LatestVersion: The latest available version of this package in
	// this distribution channel.
	LatestVersion *Version `json:"latestVersion,omitempty"`

	// Maintainer: A freeform string denoting the maintainer of this
	// package.
	Maintainer string `json:"maintainer,omitempty"`

	// Url: The distribution channel-specific homepage for this package.
	Url string `json:"url,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Architecture") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Distribution) MarshalJSON() ([]byte, error) {
	type noMethod Distribution
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

// ExtendedSourceContext: An ExtendedSourceContext is a SourceContext
// combined with additional
// details describing the context.
type ExtendedSourceContext struct {
	// Context: Any source context.
	Context *SourceContext `json:"context,omitempty"`

	// Labels: Labels with user defined metadata.
	Labels map[string]string `json:"labels,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Context") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *ExtendedSourceContext) MarshalJSON() ([]byte, error) {
	type noMethod ExtendedSourceContext
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// FileHashes: Container message for hashes of byte content of files,
// used in Source
// messages to verify integrity of source input to the build.
type FileHashes struct {
	// FileHash: Collection of file hashes.
	FileHash []*Hash `json:"fileHash,omitempty"`

	// ForceSendFields is a list of field names (e.g. "FileHash") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *FileHashes) MarshalJSON() ([]byte, error) {
	type noMethod FileHashes
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// Fingerprint: A set of properties that uniquely identify a given
// Docker image.
type Fingerprint struct {
	// V1Name: The layer-id of the final layer in the Docker image’s
	// v1
	// representation.
	// This field can be used as a filter in list requests.
	V1Name string `json:"v1Name,omitempty"`

	// V2Blob: The ordered list of v2 blobs that represent a given image.
	V2Blob []string `json:"v2Blob,omitempty"`

	// V2Name: The name of the image’s v2 blobs computed via:
	//   [bottom] := v2_blobbottom := sha256(v2_blob[N] + “ ” +
	// v2_name[N+1])
	// Only the name of the final blob is kept.
	// This field can be used as a filter in list requests.
	// @OutputOnly
	V2Name string `json:"v2Name,omitempty"`

	// ForceSendFields is a list of field names (e.g. "V1Name") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Fingerprint) MarshalJSON() ([]byte, error) {
	type noMethod Fingerprint
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// GerritSourceContext: A SourceContext referring to a Gerrit project.
type GerritSourceContext struct {
	// AliasContext: An alias, which may be a branch or tag.
	AliasContext *AliasContext `json:"aliasContext,omitempty"`

	// AliasName: The name of an alias (branch, tag, etc.).
	AliasName string `json:"aliasName,omitempty"`

	// GerritProject: The full project name within the host. Projects may be
	// nested, so
	// "project/subproject" is a valid project name.
	// The "repo name" is hostURI/project.
	GerritProject string `json:"gerritProject,omitempty"`

	// HostUri: The URI of a running Gerrit instance.
	HostUri string `json:"hostUri,omitempty"`

	// RevisionId: A revision (commit) ID.
	RevisionId string `json:"revisionId,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AliasContext") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *GerritSourceContext) MarshalJSON() ([]byte, error) {
	type noMethod GerritSourceContext
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// GetIamPolicyRequest: Request message for `GetIamPolicy` method.
type GetIamPolicyRequest struct {
}

// GitSourceContext: A GitSourceContext denotes a particular revision in
// a third party Git
// repository (e.g. GitHub).
type GitSourceContext struct {
	// RevisionId: Git commit hash.
	// required.
	RevisionId string `json:"revisionId,omitempty"`

	// Url: Git repository URL.
	Url string `json:"url,omitempty"`

	// ForceSendFields is a list of field names (e.g. "RevisionId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *GitSourceContext) MarshalJSON() ([]byte, error) {
	type noMethod GitSourceContext
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// Hash: Container message for hash values.
type Hash struct {
	// Type: The type of hash that was performed.
	//
	// Possible values:
	//   "NONE" - No hash requested.
	//   "SHA256" - A sha256 hash.
	Type string `json:"type,omitempty"`

	// Value: The hash value.
	Value string `json:"value,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Type") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Hash) MarshalJSON() ([]byte, error) {
	type noMethod Hash
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// Installation: This represents how a particular software package may
// be installed on
// a system.
type Installation struct {
	// Location: All of the places within the filesystem versions of this
	// package
	// have been found.
	Location []*Location `json:"location,omitempty"`

	// Name: The name of the installed package.
	// @OutputOnly
	Name string `json:"name,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Location") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Installation) MarshalJSON() ([]byte, error) {
	type noMethod Installation
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// Layer: Layer holds metadata specific to a layer of a Docker image.
type Layer struct {
	// Arguments: The recovered arguments to the Dockerfile directive.
	Arguments string `json:"arguments,omitempty"`

	// Directive: The recovered Dockerfile directive used to construct this
	// layer.
	//
	// Possible values:
	//   "UNKNOWN_DIRECTIVE" - Default value for unsupported/missing
	// directive
	//   "MAINTAINER" -
	// https://docs.docker.com/reference/builder/#maintainer
	//   "RUN" - https://docs.docker.com/reference/builder/#run
	//   "CMD" - https://docs.docker.com/reference/builder/#cmd
	//   "LABEL" - https://docs.docker.com/reference/builder/#label
	//   "EXPOSE" - https://docs.docker.com/reference/builder/#expose
	//   "ENV" - https://docs.docker.com/reference/builder/#env
	//   "ADD" - https://docs.docker.com/reference/builder/#add
	//   "COPY" - https://docs.docker.com/reference/builder/#copy
	//   "ENTRYPOINT" -
	// https://docs.docker.com/reference/builder/#entrypoint
	//   "VOLUME" - https://docs.docker.com/reference/builder/#volume
	//   "USER" - https://docs.docker.com/reference/builder/#user
	//   "WORKDIR" - https://docs.docker.com/reference/builder/#workdir
	//   "ARG" - https://docs.docker.com/reference/builder/#arg
	//   "ONBUILD" - https://docs.docker.com/reference/builder/#onbuild
	//   "STOPSIGNAL" -
	// https://docs.docker.com/reference/builder/#stopsignal
	//   "HEALTHCHECK" -
	// https://docs.docker.com/reference/builder/#healthcheck
	//   "SHELL" - https://docs.docker.com/reference/builder/#shell
	Directive string `json:"directive,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Arguments") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Layer) MarshalJSON() ([]byte, error) {
	type noMethod Layer
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// ListNoteOccurrencesResponse: Response including listed occurrences
// for a note.
type ListNoteOccurrencesResponse struct {
	// NextPageToken: Token to receive the next page of notes.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Occurrences: The names of the Occurrences linked to the specified
	// Note for example:
	//   "projects/{project_id}/occurrences/{occurrence_id}"
	Occurrences []string `json:"occurrences,omitempty"`

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

func (s *ListNoteOccurrencesResponse) MarshalJSON() ([]byte, error) {
	type noMethod ListNoteOccurrencesResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// ListNotesResponse: Response including listed notes.
type ListNotesResponse struct {
	// NextPageToken: The next pagination token in the List response. It
	// should be used as
	// page_token for the following request. An empty value means no more
	// result.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Notes: The occurrences requested
	Notes []*Note `json:"notes,omitempty"`

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

func (s *ListNotesResponse) MarshalJSON() ([]byte, error) {
	type noMethod ListNotesResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// ListOccurrencesResponse: Response including listed occurrences.
type ListOccurrencesResponse struct {
	// NextPageToken: The next pagination token in the List response. It
	// should be used as
	// page_token for the following request. An empty value means no more
	// result.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Occurrences: The occurrences requested
	Occurrences []*Occurrence `json:"occurrences,omitempty"`

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

func (s *ListOccurrencesResponse) MarshalJSON() ([]byte, error) {
	type noMethod ListOccurrencesResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// Location: An occurrence of a particular package installation found
// within a
// system's filesystem.
// e.g. glibc was found in /var/lib/dpkg/status
type Location struct {
	// CpeUri: The cpe_uri in [cpe
	// format](https://cpe.mitre.org/specification/)
	// denoting the package manager version distributing a package.
	CpeUri string `json:"cpeUri,omitempty"`

	// Path: The path from which we gathered that this package/version is
	// installed.
	Path string `json:"path,omitempty"`

	// Version: The version installed at this location.
	Version *Version `json:"version,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CpeUri") to
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
// At present the only supported field names are
//    - "iam_principal", corresponding to IAMContext.principal;
//    - "" (empty string), resulting in one aggretated counter with no
// field.
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

// Note: Note provides a detailed description of a note using
// information
// from the provider of the note.
type Note struct {
	// BaseImage: A note describing a base image.
	BaseImage *Basis `json:"baseImage,omitempty"`

	// BuildType: Build provenance type for a verifiable build.
	BuildType *BuildType `json:"buildType,omitempty"`

	// CreateTime: The time this note was created.
	// This field can be used as a filter in list requests.
	// @OutputOnly
	CreateTime string `json:"createTime,omitempty"`

	// ExpirationTime: Time of expiration for this Note, null if Note
	// currently does not
	// expire.
	ExpirationTime string `json:"expirationTime,omitempty"`

	// Kind: This explicitly denotes which kind of note is specified.
	// This field can be used as a filter in list requests.
	// @OutputOnly
	//
	// Possible values:
	//   "UNKNOWN" - Unknown
	//   "PACKAGE_VULNERABILITY" - The note and occurrence represent a
	// package vulnerability.
	//   "BUILD_DETAILS" - The note and occurrence  assert build provenance.
	//   "IMAGE_BASIS" - This represents an image basis relationship.
	//   "PACKAGE_MANAGER" - This represents a package installed via a
	// package manager.
	Kind string `json:"kind,omitempty"`

	// LongDescription: A detailed description of this note
	LongDescription string `json:"longDescription,omitempty"`

	// Name: The name of the note in the
	// form
	// "providers/{provider_id}/notes/{note_id}"
	Name string `json:"name,omitempty"`

	// Package: A note describing a package hosted by various package
	// managers.
	Package *Package `json:"package,omitempty"`

	// RelatedUrl: Urls associated with this note
	RelatedUrl []*RelatedUrl `json:"relatedUrl,omitempty"`

	// ShortDescription: A one sentence description of this note
	ShortDescription string `json:"shortDescription,omitempty"`

	// UpdateTime: The time this note was last updated.
	// This field can be used as a filter in list requests.
	// @OutputOnly
	UpdateTime string `json:"updateTime,omitempty"`

	// VulnerabilityType: A package vulnerability type of note.
	VulnerabilityType *VulnerabilityType `json:"vulnerabilityType,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "BaseImage") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Note) MarshalJSON() ([]byte, error) {
	type noMethod Note
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// Occurrence: Occurrence includes information about analysis
// occurrences for an image.
// ``
type Occurrence struct {
	// BuildDetails: Build details for a verifiable build.
	BuildDetails *BuildDetails `json:"buildDetails,omitempty"`

	// CreateTime: The time this occurrence was created.
	// @OutputOnly
	CreateTime string `json:"createTime,omitempty"`

	// DerivedImage: Describes how this resource derives from the basis
	// in the associated note.
	DerivedImage *Derived `json:"derivedImage,omitempty"`

	// Installation: Describes the installation of a package on the linked
	// resource.
	Installation *Installation `json:"installation,omitempty"`

	// Kind: This explicitly denotes which of the occurrence details is
	// specified.
	// This field can be used as a filter in list requests.
	// @OutputOnly
	//
	// Possible values:
	//   "UNKNOWN" - Unknown
	//   "PACKAGE_VULNERABILITY" - The note and occurrence represent a
	// package vulnerability.
	//   "BUILD_DETAILS" - The note and occurrence  assert build provenance.
	//   "IMAGE_BASIS" - This represents an image basis relationship.
	//   "PACKAGE_MANAGER" - This represents a package installed via a
	// package manager.
	Kind string `json:"kind,omitempty"`

	// Name: The name of the occurrence in the
	// form
	// "projects/{project_id}/occurrences/{occurrence_id}"
	// @OutputOnly
	Name string `json:"name,omitempty"`

	// NoteName: An analysis note associated with this image, in the
	// form
	// "providers/{provider_id}/notes/{note_id}"
	// This field can be used as a filter in list requests.
	NoteName string `json:"noteName,omitempty"`

	// Remediation: A description of actions that can be taken to remedy the
	// note
	Remediation string `json:"remediation,omitempty"`

	// ResourceUrl: The unique url of the image or container for which the
	// occurrence applies.
	// Example: https://gcr.io/project/image@sha256:foo
	// This field can be used as a filter in list requests.
	ResourceUrl string `json:"resourceUrl,omitempty"`

	// UpdateTime: The time this occurrence was last updated.
	// @OutputOnly
	UpdateTime string `json:"updateTime,omitempty"`

	// VulnerabilityDetails: Details of a security vulnerability note.
	VulnerabilityDetails *VulnerabilityDetails `json:"vulnerabilityDetails,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "BuildDetails") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Occurrence) MarshalJSON() ([]byte, error) {
	type noMethod Occurrence
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// Package: This represents a particular package that is distributed
// over
// various channels.
// e.g. glibc (aka libc6) is distributed by many, at various versions.
type Package struct {
	// Distribution: The various channels by which a package is distributed.
	Distribution []*Distribution `json:"distribution,omitempty"`

	// Name: The name of the package.
	Name string `json:"name,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Distribution") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Package) MarshalJSON() ([]byte, error) {
	type noMethod Package
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// PackageIssue: This message wraps a location affected by a
// vulnerability and its
// associated fix (if one is available).
type PackageIssue struct {
	// AffectedLocation: The location of the vulnerability.
	AffectedLocation *VulnerabilityLocation `json:"affectedLocation,omitempty"`

	// FixedLocation: The location of the available fix for vulnerability.
	FixedLocation *VulnerabilityLocation `json:"fixedLocation,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AffectedLocation") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *PackageIssue) MarshalJSON() ([]byte, error) {
	type noMethod PackageIssue
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
//       "bindings": [
//         {
//           "role": "roles/owner",
//           "members": [
//             "user:mike@example.com",
//             "group:admins@example.com",
//             "domain:google.com",
//
// "serviceAccount:my-other-app@appspot.gserviceaccount.com",
//           ]
//         },
//         {
//           "role": "roles/viewer",
//           "members": ["user:sean@example.com"]
//         }
//       ]
//     }
//
// For a description of IAM and its features, see the
// [IAM developer's guide](https://cloud.google.com/iam).
type Policy struct {
	// AuditConfigs: Specifies cloud audit logging configuration for this
	// policy.
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

	// Rules: If more than one rule is specified, the rules are applied in
	// the following
	// manner:
	// - All matching LOG rules are always applied.
	// - If any DENY/DENY_WITH_LOG rule matches, permission is denied.
	//   Logging will be applied if one or more matching rule requires
	// logging.
	// - Otherwise, if any ALLOW/ALLOW_WITH_LOG rule matches, permission is
	//   granted.
	//   Logging will be applied if one or more matching rule requires
	// logging.
	// - Otherwise, if no rule applies, permission is denied.
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

// ProjectRepoId: Selects a repo using a Google Cloud Platform project
// ID
// (e.g. winged-cargo-31) and a repo name within that project.
type ProjectRepoId struct {
	// ProjectId: The ID of the project.
	ProjectId string `json:"projectId,omitempty"`

	// RepoName: The name of the repo. Leave empty for the default repo.
	RepoName string `json:"repoName,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ProjectId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *ProjectRepoId) MarshalJSON() ([]byte, error) {
	type noMethod ProjectRepoId
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// RelatedUrl: Metadata for any related url information
type RelatedUrl struct {
	// Label: Label to describe usage of the url
	Label string `json:"label,omitempty"`

	// Url: Specific url to associate with the note
	Url string `json:"url,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Label") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *RelatedUrl) MarshalJSON() ([]byte, error) {
	type noMethod RelatedUrl
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// RepoId: A unique identifier for a cloud repo.
type RepoId struct {
	// ProjectRepoId: A combination of a project ID and a repo name.
	ProjectRepoId *ProjectRepoId `json:"projectRepoId,omitempty"`

	// Uid: A server-assigned, globally unique identifier.
	Uid string `json:"uid,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ProjectRepoId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *RepoId) MarshalJSON() ([]byte, error) {
	type noMethod RepoId
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// RepoSource: RepoSource describes the location of the source in a
// Google Cloud Source
// Repository.
type RepoSource struct {
	// BranchName: Name of the branch to build.
	BranchName string `json:"branchName,omitempty"`

	// CommitSha: Explicit commit SHA to build.
	CommitSha string `json:"commitSha,omitempty"`

	// ProjectId: ID of the project that owns the repo.
	ProjectId string `json:"projectId,omitempty"`

	// RepoName: Name of the repo.
	RepoName string `json:"repoName,omitempty"`

	// TagName: Name of the tag to build.
	TagName string `json:"tagName,omitempty"`

	// ForceSendFields is a list of field names (e.g. "BranchName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *RepoSource) MarshalJSON() ([]byte, error) {
	type noMethod RepoSource
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

	// In: If one or more 'in' clauses are specified, the rule matches
	// if
	// the PRINCIPAL/AUTHORITY_SELECTOR is in at least one of these entries.
	In []string `json:"in,omitempty"`

	// LogConfig: The config returned to callers of tech.iam.IAM.CheckPolicy
	// for any entries
	// that match the LOG action.
	LogConfig []*LogConfig `json:"logConfig,omitempty"`

	// NotIn: If one or more 'not_in' clauses are specified, the rule
	// matches
	// if the PRINCIPAL/AUTHORITY_SELECTOR is in none of the entries.
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

// SetIamPolicyRequest: Request message for `SetIamPolicy` method.
type SetIamPolicyRequest struct {
	// Policy: REQUIRED: The complete policy to be applied to the
	// `resource`. The size of
	// the policy is limited to a few 10s of KB. An empty policy is a
	// valid policy but certain Cloud Platform services (such as
	// Projects)
	// might reject them.
	Policy *Policy `json:"policy,omitempty"`

	// UpdateMask: OPTIONAL: A FieldMask specifying which fields of the
	// policy to modify. Only
	// the fields in the mask will be modified. If no mask is provided, a
	// default
	// mask is used:
	// paths: "bindings, etag"
	// This field is only used by Cloud IAM.
	UpdateMask string `json:"updateMask,omitempty"`

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

// Source: Source describes the location of the source used for the
// build.
type Source struct {
	// AdditionalSourceContexts: If provided, some of the source code used
	// for the build may be found in
	// these locations, in the case where the source repository had
	// multiple
	// remotes or submodules. This list will not include the context
	// specified in
	// the source_context field.
	AdditionalSourceContexts []*ExtendedSourceContext `json:"additionalSourceContexts,omitempty"`

	// ArtifactStorageSource: If provided, the input binary artifacts for
	// the build came from this
	// location.
	ArtifactStorageSource *StorageSource `json:"artifactStorageSource,omitempty"`

	// FileHashes: Hash(es) of the build source, which can be used to verify
	// that the original
	// source integrity was maintained in the build.
	//
	// The keys to this map are file paths used as build source and the
	// values
	// contain the hash values for those files.
	//
	// If the build source came in a single package such as a gzipped
	// tarfile
	// (.tar.gz), the FileHash will be for the single path to that file.
	FileHashes map[string]FileHashes `json:"fileHashes,omitempty"`

	// RepoSource: If provided, get source from this location in a Cloud
	// Repo.
	RepoSource *RepoSource `json:"repoSource,omitempty"`

	// SourceContext: If provided, the source code used for the build came
	// from this location.
	SourceContext *ExtendedSourceContext `json:"sourceContext,omitempty"`

	// StorageSource: If provided, get the source from this location in in
	// Google Cloud
	// Storage.
	StorageSource *StorageSource `json:"storageSource,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "AdditionalSourceContexts") to unconditionally include in API
	// requests. By default, fields with empty values are omitted from API
	// requests. However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Source) MarshalJSON() ([]byte, error) {
	type noMethod Source
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// SourceContext: A SourceContext is a reference to a tree of files. A
// SourceContext together
// with a path point to a unique revision of a single file or directory.
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

	// ForceSendFields is a list of field names (e.g. "CloudRepo") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *SourceContext) MarshalJSON() ([]byte, error) {
	type noMethod SourceContext
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// StorageSource: StorageSource describes the location of the source in
// an archive file in
// Google Cloud Storage.
type StorageSource struct {
	// Bucket: Google Cloud Storage bucket containing source (see [Bucket
	// Name
	// Requirements]
	// (https://cloud.google.com/storage/docs/bucket-namin
	// g#requirements)).
	Bucket string `json:"bucket,omitempty"`

	// Generation: Google Cloud Storage generation for the object.
	Generation int64 `json:"generation,omitempty,string"`

	// Object: Google Cloud Storage object containing source.
	Object string `json:"object,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Bucket") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *StorageSource) MarshalJSON() ([]byte, error) {
	type noMethod StorageSource
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
	// [IAM
	// Overview](https://cloud.google.com/iam/docs/overview#permissions).
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

// Version: Version contains structured information about the version of
// the package.
// For a discussion of this in
// Debian/Ubuntu:
// http://serverfault.com/questions/604541/debian-packages
// -version-convention
// For a discussion of this in
// Redhat/Fedora/Centos:
// http://blog.jasonantman.com/2014/07/how-yum-and-
// rpm-compare-versions/
type Version struct {
	// Epoch: Used to correct mistakes in the version numbering scheme.
	Epoch int64 `json:"epoch,omitempty"`

	// Kind: Distinguish between sentinel MIN/MAX versions and normal
	// versions.
	// If kind is not NORMAL, then the other fields are ignored.
	//
	// Possible values:
	//   "NORMAL" - A standard package version, defined by the other fields.
	//   "MINIMUM" - A special version representing negative infinity,
	// other fields are ignored.
	//   "MAXIMUM" - A special version representing positive infinity,
	// other fields are ignored.
	Kind string `json:"kind,omitempty"`

	// Name: The main part of the version name.
	Name string `json:"name,omitempty"`

	// Revision: The iteration of the package build from the above version.
	Revision string `json:"revision,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Epoch") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Version) MarshalJSON() ([]byte, error) {
	type noMethod Version
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// VulnerabilityDetails: Used by Occurrence to point to where the
// vulnerability exists and how
// to fix it.
type VulnerabilityDetails struct {
	// CvssScore: The CVSS score of this vulnerability. CVSS score is on a
	// scale of 0-10
	// where 0 indicates low severity and 10 indicates high
	// severity.
	// @OutputOnly
	CvssScore float64 `json:"cvssScore,omitempty"`

	// PackageIssue: The set of affected locations and their fixes (if
	// available) within
	// the associated resource.
	PackageIssue []*PackageIssue `json:"packageIssue,omitempty"`

	// Severity: The note provider assigned Severity of the
	// vulnerability.
	// @OutputOnly
	//
	// Possible values:
	//   "UNKNOWN" - Unknown Impact
	//   "MINIMAL" - Minimal Impact
	//   "LOW" - Low Impact
	//   "MEDIUM" - Medium Impact
	//   "HIGH" - High Impact
	//   "CRITICAL" - Critical Impact
	Severity string `json:"severity,omitempty"`

	// Type: The type of package; whether native or non native(ruby
	// gems,
	// node.js packages etc)
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CvssScore") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *VulnerabilityDetails) MarshalJSON() ([]byte, error) {
	type noMethod VulnerabilityDetails
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// VulnerabilityLocation: The location of the vulnerability
type VulnerabilityLocation struct {
	// CpeUri: The cpe_uri in [cpe format]
	// (https://cpe.mitre.org/specification/)
	// format. Examples include distro or storage location for vulnerable
	// jar.
	// This field can be used as a filter in list requests.
	CpeUri string `json:"cpeUri,omitempty"`

	// Package: The package being described.
	Package string `json:"package,omitempty"`

	// Version: The version of the package being described.
	// This field can be used as a filter in list requests.
	Version *Version `json:"version,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CpeUri") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *VulnerabilityLocation) MarshalJSON() ([]byte, error) {
	type noMethod VulnerabilityLocation
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// VulnerabilityType: VulnerabilityType provides metadata about a
// security vulnerability.
type VulnerabilityType struct {
	// CvssScore: The CVSS score for this Vulnerability.
	CvssScore float64 `json:"cvssScore,omitempty"`

	// Details: All information about the package to specifically identify
	// this
	// vulnerability. One entry per (version range and cpe_uri) the
	// package vulnerability has manifested in.
	Details []*Detail `json:"details,omitempty"`

	// Severity: Note provider assigned impact of the vulnerability
	//
	// Possible values:
	//   "UNKNOWN" - Unknown Impact
	//   "MINIMAL" - Minimal Impact
	//   "LOW" - Low Impact
	//   "MEDIUM" - Medium Impact
	//   "HIGH" - High Impact
	//   "CRITICAL" - Critical Impact
	Severity string `json:"severity,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CvssScore") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *VulnerabilityType) MarshalJSON() ([]byte, error) {
	type noMethod VulnerabilityType
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// method id "containeranalysis.projects.notes.create":

type ProjectsNotesCreateCall struct {
	s          *Service
	parent     string
	note       *Note
	urlParams_ gensupport.URLParams
	ctx_       context.Context
}

// Create: Creates a new note.
func (r *ProjectsNotesService) Create(parent string, note *Note) *ProjectsNotesCreateCall {
	c := &ProjectsNotesCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.note = note
	return c
}

// Name sets the optional parameter "name": The name of the
// project.
// Should be of the form "providers/{provider_id}".
// @deprecated
func (c *ProjectsNotesCreateCall) Name(name string) *ProjectsNotesCreateCall {
	c.urlParams_.Set("name", name)
	return c
}

// NoteId sets the optional parameter "noteId": The ID to use for this
// note.
func (c *ProjectsNotesCreateCall) NoteId(noteId string) *ProjectsNotesCreateCall {
	c.urlParams_.Set("noteId", noteId)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsNotesCreateCall) Fields(s ...googleapi.Field) *ProjectsNotesCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsNotesCreateCall) Context(ctx context.Context) *ProjectsNotesCreateCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsNotesCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.note)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha1/{+parent}/notes")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "containeranalysis.projects.notes.create" call.
// Exactly one of *Note or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Note.ServerResponse.Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsNotesCreateCall) Do(opts ...googleapi.CallOption) (*Note, error) {
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
	ret := &Note{
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
	//   "description": "Creates a new note.",
	//   "flatPath": "v1alpha1/projects/{projectsId}/notes",
	//   "httpMethod": "POST",
	//   "id": "containeranalysis.projects.notes.create",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The name of the project.\nShould be of the form \"providers/{provider_id}\".\n@deprecated",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "noteId": {
	//       "description": "The ID to use for this note.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "The parent field will contain the projectId for example:\n\"project/{project_id}",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1alpha1/{+parent}/notes",
	//   "request": {
	//     "$ref": "Note"
	//   },
	//   "response": {
	//     "$ref": "Note"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "containeranalysis.projects.notes.delete":

type ProjectsNotesDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
}

// Delete: Deletes the given note from the system.
func (r *ProjectsNotesService) Delete(name string) *ProjectsNotesDeleteCall {
	c := &ProjectsNotesDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsNotesDeleteCall) Fields(s ...googleapi.Field) *ProjectsNotesDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsNotesDeleteCall) Context(ctx context.Context) *ProjectsNotesDeleteCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsNotesDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha1/{+name}")
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

// Do executes the "containeranalysis.projects.notes.delete" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsNotesDeleteCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
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
	//   "description": "Deletes the given note from the system.",
	//   "flatPath": "v1alpha1/projects/{projectsId}/notes/{notesId}",
	//   "httpMethod": "DELETE",
	//   "id": "containeranalysis.projects.notes.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The name of the note in the form\n\"providers/{provider_id}/notes/{note_id}\"",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/notes/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1alpha1/{+name}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "containeranalysis.projects.notes.get":

type ProjectsNotesGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// Get: Returns the requested occurrence
func (r *ProjectsNotesService) Get(name string) *ProjectsNotesGetCall {
	c := &ProjectsNotesGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsNotesGetCall) Fields(s ...googleapi.Field) *ProjectsNotesGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsNotesGetCall) IfNoneMatch(entityTag string) *ProjectsNotesGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsNotesGetCall) Context(ctx context.Context) *ProjectsNotesGetCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsNotesGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha1/{+name}")
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

// Do executes the "containeranalysis.projects.notes.get" call.
// Exactly one of *Note or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Note.ServerResponse.Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsNotesGetCall) Do(opts ...googleapi.CallOption) (*Note, error) {
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
	ret := &Note{
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
	//   "description": "Returns the requested occurrence",
	//   "flatPath": "v1alpha1/projects/{projectsId}/notes/{notesId}",
	//   "httpMethod": "GET",
	//   "id": "containeranalysis.projects.notes.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The name of the occurrence in the form\n\"providers/{provider_id}/notes/{note_id}\"",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/notes/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1alpha1/{+name}",
	//   "response": {
	//     "$ref": "Note"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "containeranalysis.projects.notes.getIamPolicy":

type ProjectsNotesGetIamPolicyCall struct {
	s                   *Service
	resource            string
	getiampolicyrequest *GetIamPolicyRequest
	urlParams_          gensupport.URLParams
	ctx_                context.Context
}

// GetIamPolicy: Gets the access control policy for a note or occurrence
// resource.
// Requires "containeranalysis.notes.setIamPolicy"
// or
// "containeranalysis.occurrences.setIamPolicy" permission if the
// resource is
// a note or occurrence, respectively.
// Attempting this RPC on a resource without the needed permission will
// note
// in a PERMISSION_DENIED error.
// Attempting this RPC on a non-existent resource will note in a
// NOT_FOUND
// error if the user has list permission on the project,
// or a PERMISSION_DENIED error otherwise.
func (r *ProjectsNotesService) GetIamPolicy(resource string, getiampolicyrequest *GetIamPolicyRequest) *ProjectsNotesGetIamPolicyCall {
	c := &ProjectsNotesGetIamPolicyCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.resource = resource
	c.getiampolicyrequest = getiampolicyrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsNotesGetIamPolicyCall) Fields(s ...googleapi.Field) *ProjectsNotesGetIamPolicyCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsNotesGetIamPolicyCall) Context(ctx context.Context) *ProjectsNotesGetIamPolicyCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsNotesGetIamPolicyCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.getiampolicyrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha1/{+resource}:getIamPolicy")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "containeranalysis.projects.notes.getIamPolicy" call.
// Exactly one of *Policy or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Policy.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsNotesGetIamPolicyCall) Do(opts ...googleapi.CallOption) (*Policy, error) {
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
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets the access control policy for a note or occurrence resource.\nRequires \"containeranalysis.notes.setIamPolicy\" or\n\"containeranalysis.occurrences.setIamPolicy\" permission if the resource is\na note or occurrence, respectively.\nAttempting this RPC on a resource without the needed permission will note\nin a PERMISSION_DENIED error.\nAttempting this RPC on a non-existent resource will note in a NOT_FOUND\nerror if the user has list permission on the project,\nor a PERMISSION_DENIED error otherwise.",
	//   "flatPath": "v1alpha1/projects/{projectsId}/notes/{notesId}:getIamPolicy",
	//   "httpMethod": "POST",
	//   "id": "containeranalysis.projects.notes.getIamPolicy",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "REQUIRED: The resource for which the policy is being requested.\nSee the operation documentation for the appropriate value for this field.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/notes/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1alpha1/{+resource}:getIamPolicy",
	//   "request": {
	//     "$ref": "GetIamPolicyRequest"
	//   },
	//   "response": {
	//     "$ref": "Policy"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "containeranalysis.projects.notes.list":

type ProjectsNotesListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// List: Lists all notes for a given project.  Filters can be used on
// this
// field to list all notes with a specific parameter.
func (r *ProjectsNotesService) List(parent string) *ProjectsNotesListCall {
	c := &ProjectsNotesListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// Filter sets the optional parameter "filter": The filter expression.
func (c *ProjectsNotesListCall) Filter(filter string) *ProjectsNotesListCall {
	c.urlParams_.Set("filter", filter)
	return c
}

// Name sets the optional parameter "name": The name field will contain
// the projectId for example:
// "providers/{provider_id}
// @deprecated
func (c *ProjectsNotesListCall) Name(name string) *ProjectsNotesListCall {
	c.urlParams_.Set("name", name)
	return c
}

// PageSize sets the optional parameter "pageSize": Number of notes to
// return in the list.
func (c *ProjectsNotesListCall) PageSize(pageSize int64) *ProjectsNotesListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": Token to provide
// to skip to a particular spot in the list.
func (c *ProjectsNotesListCall) PageToken(pageToken string) *ProjectsNotesListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsNotesListCall) Fields(s ...googleapi.Field) *ProjectsNotesListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsNotesListCall) IfNoneMatch(entityTag string) *ProjectsNotesListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsNotesListCall) Context(ctx context.Context) *ProjectsNotesListCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsNotesListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha1/{+parent}/notes")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "containeranalysis.projects.notes.list" call.
// Exactly one of *ListNotesResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListNotesResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsNotesListCall) Do(opts ...googleapi.CallOption) (*ListNotesResponse, error) {
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
	ret := &ListNotesResponse{
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
	//   "description": "Lists all notes for a given project.  Filters can be used on this\nfield to list all notes with a specific parameter.",
	//   "flatPath": "v1alpha1/projects/{projectsId}/notes",
	//   "httpMethod": "GET",
	//   "id": "containeranalysis.projects.notes.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "The filter expression.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "name": {
	//       "description": "The name field will contain the projectId for example:\n\"providers/{provider_id}\n@deprecated",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Number of notes to return in the list.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Token to provide to skip to a particular spot in the list.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "The parent field will contain the projectId for example:\n\"project/{project_id}",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1alpha1/{+parent}/notes",
	//   "response": {
	//     "$ref": "ListNotesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsNotesListCall) Pages(ctx context.Context, f func(*ListNotesResponse) error) error {
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

// method id "containeranalysis.projects.notes.setIamPolicy":

type ProjectsNotesSetIamPolicyCall struct {
	s                   *Service
	resource            string
	setiampolicyrequest *SetIamPolicyRequest
	urlParams_          gensupport.URLParams
	ctx_                context.Context
}

// SetIamPolicy: Sets the access control policy on the specified note or
// occurrence
// resource.
// Requires "containeranalysis.notes.setIamPolicy"
// or
// "containeranalysis.occurrences.setIamPolicy" permission if the
// resource is
// a note or occurrence, respectively.
// Attempting this RPC on a resource without the needed permission will
// note
// in a PERMISSION_DENIED error.
// Attempting this RPC on a non-existent resource will note in a
// NOT_FOUND
// error if the user has list permission on the project, or
// a
// PERMISSION_DENIED error otherwise.
func (r *ProjectsNotesService) SetIamPolicy(resource string, setiampolicyrequest *SetIamPolicyRequest) *ProjectsNotesSetIamPolicyCall {
	c := &ProjectsNotesSetIamPolicyCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.resource = resource
	c.setiampolicyrequest = setiampolicyrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsNotesSetIamPolicyCall) Fields(s ...googleapi.Field) *ProjectsNotesSetIamPolicyCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsNotesSetIamPolicyCall) Context(ctx context.Context) *ProjectsNotesSetIamPolicyCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsNotesSetIamPolicyCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.setiampolicyrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha1/{+resource}:setIamPolicy")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "containeranalysis.projects.notes.setIamPolicy" call.
// Exactly one of *Policy or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Policy.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsNotesSetIamPolicyCall) Do(opts ...googleapi.CallOption) (*Policy, error) {
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
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Sets the access control policy on the specified note or occurrence\nresource.\nRequires \"containeranalysis.notes.setIamPolicy\" or\n\"containeranalysis.occurrences.setIamPolicy\" permission if the resource is\na note or occurrence, respectively.\nAttempting this RPC on a resource without the needed permission will note\nin a PERMISSION_DENIED error.\nAttempting this RPC on a non-existent resource will note in a NOT_FOUND\nerror if the user has list permission on the project, or a\nPERMISSION_DENIED error otherwise.",
	//   "flatPath": "v1alpha1/projects/{projectsId}/notes/{notesId}:setIamPolicy",
	//   "httpMethod": "POST",
	//   "id": "containeranalysis.projects.notes.setIamPolicy",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "REQUIRED: The resource for which the policy is being specified.\nSee the operation documentation for the appropriate value for this field.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/notes/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1alpha1/{+resource}:setIamPolicy",
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

// method id "containeranalysis.projects.notes.testIamPermissions":

type ProjectsNotesTestIamPermissionsCall struct {
	s                         *Service
	resource                  string
	testiampermissionsrequest *TestIamPermissionsRequest
	urlParams_                gensupport.URLParams
	ctx_                      context.Context
}

// TestIamPermissions: Returns permissions that a caller has on the
// specified note or occurrence
// resource.
// Requires list permission on the project (e.g., "storage.objects.list"
// on
// the containing bucket for testing permission of an
// object).
// Attempting this RPC on a non-existent resource will note in a
// NOT_FOUND
// error if the user has list permission on the project,
// or a PERMISSION_DENIED error otherwise.
func (r *ProjectsNotesService) TestIamPermissions(resource string, testiampermissionsrequest *TestIamPermissionsRequest) *ProjectsNotesTestIamPermissionsCall {
	c := &ProjectsNotesTestIamPermissionsCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.resource = resource
	c.testiampermissionsrequest = testiampermissionsrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsNotesTestIamPermissionsCall) Fields(s ...googleapi.Field) *ProjectsNotesTestIamPermissionsCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsNotesTestIamPermissionsCall) Context(ctx context.Context) *ProjectsNotesTestIamPermissionsCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsNotesTestIamPermissionsCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.testiampermissionsrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha1/{+resource}:testIamPermissions")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "containeranalysis.projects.notes.testIamPermissions" call.
// Exactly one of *TestIamPermissionsResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *TestIamPermissionsResponse.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsNotesTestIamPermissionsCall) Do(opts ...googleapi.CallOption) (*TestIamPermissionsResponse, error) {
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
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns permissions that a caller has on the specified note or occurrence\nresource.\nRequires list permission on the project (e.g., \"storage.objects.list\" on\nthe containing bucket for testing permission of an object).\nAttempting this RPC on a non-existent resource will note in a NOT_FOUND\nerror if the user has list permission on the project,\nor a PERMISSION_DENIED error otherwise.",
	//   "flatPath": "v1alpha1/projects/{projectsId}/notes/{notesId}:testIamPermissions",
	//   "httpMethod": "POST",
	//   "id": "containeranalysis.projects.notes.testIamPermissions",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "REQUIRED: The resource for which the policy detail is being requested.\nSee the operation documentation for the appropriate value for this field.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/notes/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1alpha1/{+resource}:testIamPermissions",
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

// method id "containeranalysis.projects.notes.update":

type ProjectsNotesUpdateCall struct {
	s          *Service
	name       string
	note       *Note
	urlParams_ gensupport.URLParams
	ctx_       context.Context
}

// Update: Updates an existing note.
func (r *ProjectsNotesService) Update(name string, note *Note) *ProjectsNotesUpdateCall {
	c := &ProjectsNotesUpdateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.note = note
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsNotesUpdateCall) Fields(s ...googleapi.Field) *ProjectsNotesUpdateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsNotesUpdateCall) Context(ctx context.Context) *ProjectsNotesUpdateCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsNotesUpdateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.note)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha1/{+name}")
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

// Do executes the "containeranalysis.projects.notes.update" call.
// Exactly one of *Note or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Note.ServerResponse.Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsNotesUpdateCall) Do(opts ...googleapi.CallOption) (*Note, error) {
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
	ret := &Note{
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
	//   "description": "Updates an existing note.",
	//   "flatPath": "v1alpha1/projects/{projectsId}/notes/{notesId}",
	//   "httpMethod": "PUT",
	//   "id": "containeranalysis.projects.notes.update",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The name of the note.\nShould be of the form \"projects/{provider_id}/notes/{note_id}\".",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/notes/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1alpha1/{+name}",
	//   "request": {
	//     "$ref": "Note"
	//   },
	//   "response": {
	//     "$ref": "Note"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "containeranalysis.projects.notes.occurrences.list":

type ProjectsNotesOccurrencesListCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// List: Lists the names of Occurrences linked to a particular Note.
func (r *ProjectsNotesOccurrencesService) List(name string) *ProjectsNotesOccurrencesListCall {
	c := &ProjectsNotesOccurrencesListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Filter sets the optional parameter "filter": The filter expression.
func (c *ProjectsNotesOccurrencesListCall) Filter(filter string) *ProjectsNotesOccurrencesListCall {
	c.urlParams_.Set("filter", filter)
	return c
}

// PageSize sets the optional parameter "pageSize": Number of notes to
// return in the list.
func (c *ProjectsNotesOccurrencesListCall) PageSize(pageSize int64) *ProjectsNotesOccurrencesListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": Token to provide
// to skip to a particular spot in the list.
func (c *ProjectsNotesOccurrencesListCall) PageToken(pageToken string) *ProjectsNotesOccurrencesListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsNotesOccurrencesListCall) Fields(s ...googleapi.Field) *ProjectsNotesOccurrencesListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsNotesOccurrencesListCall) IfNoneMatch(entityTag string) *ProjectsNotesOccurrencesListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsNotesOccurrencesListCall) Context(ctx context.Context) *ProjectsNotesOccurrencesListCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsNotesOccurrencesListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha1/{+name}/occurrences")
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

// Do executes the "containeranalysis.projects.notes.occurrences.list" call.
// Exactly one of *ListNoteOccurrencesResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *ListNoteOccurrencesResponse.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsNotesOccurrencesListCall) Do(opts ...googleapi.CallOption) (*ListNoteOccurrencesResponse, error) {
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
	ret := &ListNoteOccurrencesResponse{
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
	//   "description": "Lists the names of Occurrences linked to a particular Note.",
	//   "flatPath": "v1alpha1/projects/{projectsId}/notes/{notesId}/occurrences",
	//   "httpMethod": "GET",
	//   "id": "containeranalysis.projects.notes.occurrences.list",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "The filter expression.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "name": {
	//       "description": "The name field will contain the note name for example:\n  \"provider/{provider_id}/notes/{note_id}\"",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/notes/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Number of notes to return in the list.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Token to provide to skip to a particular spot in the list.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1alpha1/{+name}/occurrences",
	//   "response": {
	//     "$ref": "ListNoteOccurrencesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsNotesOccurrencesListCall) Pages(ctx context.Context, f func(*ListNoteOccurrencesResponse) error) error {
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

// method id "containeranalysis.projects.occurrences.create":

type ProjectsOccurrencesCreateCall struct {
	s          *Service
	parent     string
	occurrence *Occurrence
	urlParams_ gensupport.URLParams
	ctx_       context.Context
}

// Create: Creates a new occurrence.
func (r *ProjectsOccurrencesService) Create(parent string, occurrence *Occurrence) *ProjectsOccurrencesCreateCall {
	c := &ProjectsOccurrencesCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	c.occurrence = occurrence
	return c
}

// Name sets the optional parameter "name": The name of the project.
// Should be of the form "projects/{project_id}".
// @deprecated
func (c *ProjectsOccurrencesCreateCall) Name(name string) *ProjectsOccurrencesCreateCall {
	c.urlParams_.Set("name", name)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsOccurrencesCreateCall) Fields(s ...googleapi.Field) *ProjectsOccurrencesCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsOccurrencesCreateCall) Context(ctx context.Context) *ProjectsOccurrencesCreateCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsOccurrencesCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.occurrence)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha1/{+parent}/occurrences")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "containeranalysis.projects.occurrences.create" call.
// Exactly one of *Occurrence or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Occurrence.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsOccurrencesCreateCall) Do(opts ...googleapi.CallOption) (*Occurrence, error) {
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
	ret := &Occurrence{
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
	//   "description": "Creates a new occurrence.",
	//   "flatPath": "v1alpha1/projects/{projectsId}/occurrences",
	//   "httpMethod": "POST",
	//   "id": "containeranalysis.projects.occurrences.create",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The name of the project.  Should be of the form \"projects/{project_id}\".\n@deprecated",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "The parent field will contain the projectId for example:\n\"project/{project_id}",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1alpha1/{+parent}/occurrences",
	//   "request": {
	//     "$ref": "Occurrence"
	//   },
	//   "response": {
	//     "$ref": "Occurrence"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "containeranalysis.projects.occurrences.delete":

type ProjectsOccurrencesDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
}

// Delete: Deletes the given occurrence from the system.
func (r *ProjectsOccurrencesService) Delete(name string) *ProjectsOccurrencesDeleteCall {
	c := &ProjectsOccurrencesDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsOccurrencesDeleteCall) Fields(s ...googleapi.Field) *ProjectsOccurrencesDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsOccurrencesDeleteCall) Context(ctx context.Context) *ProjectsOccurrencesDeleteCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsOccurrencesDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha1/{+name}")
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

// Do executes the "containeranalysis.projects.occurrences.delete" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsOccurrencesDeleteCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
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
	//   "description": "Deletes the given occurrence from the system.",
	//   "flatPath": "v1alpha1/projects/{projectsId}/occurrences/{occurrencesId}",
	//   "httpMethod": "DELETE",
	//   "id": "containeranalysis.projects.occurrences.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The name of the occurrence in the form\n\"projects/{project_id}/occurrences/{occurrence_id}\"",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/occurrences/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1alpha1/{+name}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "containeranalysis.projects.occurrences.get":

type ProjectsOccurrencesGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// Get: Returns the requested occurrence
func (r *ProjectsOccurrencesService) Get(name string) *ProjectsOccurrencesGetCall {
	c := &ProjectsOccurrencesGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsOccurrencesGetCall) Fields(s ...googleapi.Field) *ProjectsOccurrencesGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsOccurrencesGetCall) IfNoneMatch(entityTag string) *ProjectsOccurrencesGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsOccurrencesGetCall) Context(ctx context.Context) *ProjectsOccurrencesGetCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsOccurrencesGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha1/{+name}")
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

// Do executes the "containeranalysis.projects.occurrences.get" call.
// Exactly one of *Occurrence or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Occurrence.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsOccurrencesGetCall) Do(opts ...googleapi.CallOption) (*Occurrence, error) {
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
	ret := &Occurrence{
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
	//   "description": "Returns the requested occurrence",
	//   "flatPath": "v1alpha1/projects/{projectsId}/occurrences/{occurrencesId}",
	//   "httpMethod": "GET",
	//   "id": "containeranalysis.projects.occurrences.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The name of the occurrence in the form\n\"projects/{project_id}/occurrences/{occurrence_id}\"",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/occurrences/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1alpha1/{+name}",
	//   "response": {
	//     "$ref": "Occurrence"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "containeranalysis.projects.occurrences.getIamPolicy":

type ProjectsOccurrencesGetIamPolicyCall struct {
	s                   *Service
	resource            string
	getiampolicyrequest *GetIamPolicyRequest
	urlParams_          gensupport.URLParams
	ctx_                context.Context
}

// GetIamPolicy: Gets the access control policy for a note or occurrence
// resource.
// Requires "containeranalysis.notes.setIamPolicy"
// or
// "containeranalysis.occurrences.setIamPolicy" permission if the
// resource is
// a note or occurrence, respectively.
// Attempting this RPC on a resource without the needed permission will
// note
// in a PERMISSION_DENIED error.
// Attempting this RPC on a non-existent resource will note in a
// NOT_FOUND
// error if the user has list permission on the project,
// or a PERMISSION_DENIED error otherwise.
func (r *ProjectsOccurrencesService) GetIamPolicy(resource string, getiampolicyrequest *GetIamPolicyRequest) *ProjectsOccurrencesGetIamPolicyCall {
	c := &ProjectsOccurrencesGetIamPolicyCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.resource = resource
	c.getiampolicyrequest = getiampolicyrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsOccurrencesGetIamPolicyCall) Fields(s ...googleapi.Field) *ProjectsOccurrencesGetIamPolicyCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsOccurrencesGetIamPolicyCall) Context(ctx context.Context) *ProjectsOccurrencesGetIamPolicyCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsOccurrencesGetIamPolicyCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.getiampolicyrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha1/{+resource}:getIamPolicy")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "containeranalysis.projects.occurrences.getIamPolicy" call.
// Exactly one of *Policy or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Policy.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsOccurrencesGetIamPolicyCall) Do(opts ...googleapi.CallOption) (*Policy, error) {
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
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets the access control policy for a note or occurrence resource.\nRequires \"containeranalysis.notes.setIamPolicy\" or\n\"containeranalysis.occurrences.setIamPolicy\" permission if the resource is\na note or occurrence, respectively.\nAttempting this RPC on a resource without the needed permission will note\nin a PERMISSION_DENIED error.\nAttempting this RPC on a non-existent resource will note in a NOT_FOUND\nerror if the user has list permission on the project,\nor a PERMISSION_DENIED error otherwise.",
	//   "flatPath": "v1alpha1/projects/{projectsId}/occurrences/{occurrencesId}:getIamPolicy",
	//   "httpMethod": "POST",
	//   "id": "containeranalysis.projects.occurrences.getIamPolicy",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "REQUIRED: The resource for which the policy is being requested.\nSee the operation documentation for the appropriate value for this field.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/occurrences/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1alpha1/{+resource}:getIamPolicy",
	//   "request": {
	//     "$ref": "GetIamPolicyRequest"
	//   },
	//   "response": {
	//     "$ref": "Policy"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "containeranalysis.projects.occurrences.getNotes":

type ProjectsOccurrencesGetNotesCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// GetNotes: Gets the note that this occurrence is attached to.
func (r *ProjectsOccurrencesService) GetNotes(name string) *ProjectsOccurrencesGetNotesCall {
	c := &ProjectsOccurrencesGetNotesCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsOccurrencesGetNotesCall) Fields(s ...googleapi.Field) *ProjectsOccurrencesGetNotesCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsOccurrencesGetNotesCall) IfNoneMatch(entityTag string) *ProjectsOccurrencesGetNotesCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsOccurrencesGetNotesCall) Context(ctx context.Context) *ProjectsOccurrencesGetNotesCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsOccurrencesGetNotesCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha1/{+name}/notes")
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

// Do executes the "containeranalysis.projects.occurrences.getNotes" call.
// Exactly one of *Note or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Note.ServerResponse.Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsOccurrencesGetNotesCall) Do(opts ...googleapi.CallOption) (*Note, error) {
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
	ret := &Note{
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
	//   "description": "Gets the note that this occurrence is attached to.",
	//   "flatPath": "v1alpha1/projects/{projectsId}/occurrences/{occurrencesId}/notes",
	//   "httpMethod": "GET",
	//   "id": "containeranalysis.projects.occurrences.getNotes",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The name of the occurrence in the form\n\"projects/{project_id}/occurrences/{occurrence_id}\"",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/occurrences/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1alpha1/{+name}/notes",
	//   "response": {
	//     "$ref": "Note"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "containeranalysis.projects.occurrences.list":

type ProjectsOccurrencesListCall struct {
	s            *Service
	parent       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// List: Lists all occurrences for a given project/Digest.  Filters can
// be used on
// this field to list all digests containing a specific occurrence in
// a
// project.
func (r *ProjectsOccurrencesService) List(parent string) *ProjectsOccurrencesListCall {
	c := &ProjectsOccurrencesListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.parent = parent
	return c
}

// Filter sets the optional parameter "filter": The filter expression.
func (c *ProjectsOccurrencesListCall) Filter(filter string) *ProjectsOccurrencesListCall {
	c.urlParams_.Set("filter", filter)
	return c
}

// Name sets the optional parameter "name": The name field will contain
// the projectId for example:
// "projects/{project_id}
// @deprecated
func (c *ProjectsOccurrencesListCall) Name(name string) *ProjectsOccurrencesListCall {
	c.urlParams_.Set("name", name)
	return c
}

// PageSize sets the optional parameter "pageSize": Number of notes to
// return in the list.
func (c *ProjectsOccurrencesListCall) PageSize(pageSize int64) *ProjectsOccurrencesListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": Token to provide
// to skip to a particular spot in the list.
func (c *ProjectsOccurrencesListCall) PageToken(pageToken string) *ProjectsOccurrencesListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsOccurrencesListCall) Fields(s ...googleapi.Field) *ProjectsOccurrencesListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsOccurrencesListCall) IfNoneMatch(entityTag string) *ProjectsOccurrencesListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsOccurrencesListCall) Context(ctx context.Context) *ProjectsOccurrencesListCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsOccurrencesListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha1/{+parent}/occurrences")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"parent": c.parent,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "containeranalysis.projects.occurrences.list" call.
// Exactly one of *ListOccurrencesResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListOccurrencesResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsOccurrencesListCall) Do(opts ...googleapi.CallOption) (*ListOccurrencesResponse, error) {
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
	ret := &ListOccurrencesResponse{
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
	//   "description": "Lists all occurrences for a given project/Digest.  Filters can be used on\nthis field to list all digests containing a specific occurrence in a\nproject.",
	//   "flatPath": "v1alpha1/projects/{projectsId}/occurrences",
	//   "httpMethod": "GET",
	//   "id": "containeranalysis.projects.occurrences.list",
	//   "parameterOrder": [
	//     "parent"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "The filter expression.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "name": {
	//       "description": "The name field will contain the projectId for example:\n\"projects/{project_id}\n@deprecated",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Number of notes to return in the list.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Token to provide to skip to a particular spot in the list.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "The parent field will contain the projectId for example:\n\"project/{project_id}",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1alpha1/{+parent}/occurrences",
	//   "response": {
	//     "$ref": "ListOccurrencesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsOccurrencesListCall) Pages(ctx context.Context, f func(*ListOccurrencesResponse) error) error {
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

// method id "containeranalysis.projects.occurrences.setIamPolicy":

type ProjectsOccurrencesSetIamPolicyCall struct {
	s                   *Service
	resource            string
	setiampolicyrequest *SetIamPolicyRequest
	urlParams_          gensupport.URLParams
	ctx_                context.Context
}

// SetIamPolicy: Sets the access control policy on the specified note or
// occurrence
// resource.
// Requires "containeranalysis.notes.setIamPolicy"
// or
// "containeranalysis.occurrences.setIamPolicy" permission if the
// resource is
// a note or occurrence, respectively.
// Attempting this RPC on a resource without the needed permission will
// note
// in a PERMISSION_DENIED error.
// Attempting this RPC on a non-existent resource will note in a
// NOT_FOUND
// error if the user has list permission on the project, or
// a
// PERMISSION_DENIED error otherwise.
func (r *ProjectsOccurrencesService) SetIamPolicy(resource string, setiampolicyrequest *SetIamPolicyRequest) *ProjectsOccurrencesSetIamPolicyCall {
	c := &ProjectsOccurrencesSetIamPolicyCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.resource = resource
	c.setiampolicyrequest = setiampolicyrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsOccurrencesSetIamPolicyCall) Fields(s ...googleapi.Field) *ProjectsOccurrencesSetIamPolicyCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsOccurrencesSetIamPolicyCall) Context(ctx context.Context) *ProjectsOccurrencesSetIamPolicyCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsOccurrencesSetIamPolicyCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.setiampolicyrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha1/{+resource}:setIamPolicy")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "containeranalysis.projects.occurrences.setIamPolicy" call.
// Exactly one of *Policy or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Policy.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProjectsOccurrencesSetIamPolicyCall) Do(opts ...googleapi.CallOption) (*Policy, error) {
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
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Sets the access control policy on the specified note or occurrence\nresource.\nRequires \"containeranalysis.notes.setIamPolicy\" or\n\"containeranalysis.occurrences.setIamPolicy\" permission if the resource is\na note or occurrence, respectively.\nAttempting this RPC on a resource without the needed permission will note\nin a PERMISSION_DENIED error.\nAttempting this RPC on a non-existent resource will note in a NOT_FOUND\nerror if the user has list permission on the project, or a\nPERMISSION_DENIED error otherwise.",
	//   "flatPath": "v1alpha1/projects/{projectsId}/occurrences/{occurrencesId}:setIamPolicy",
	//   "httpMethod": "POST",
	//   "id": "containeranalysis.projects.occurrences.setIamPolicy",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "REQUIRED: The resource for which the policy is being specified.\nSee the operation documentation for the appropriate value for this field.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/occurrences/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1alpha1/{+resource}:setIamPolicy",
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

// method id "containeranalysis.projects.occurrences.testIamPermissions":

type ProjectsOccurrencesTestIamPermissionsCall struct {
	s                         *Service
	resource                  string
	testiampermissionsrequest *TestIamPermissionsRequest
	urlParams_                gensupport.URLParams
	ctx_                      context.Context
}

// TestIamPermissions: Returns permissions that a caller has on the
// specified note or occurrence
// resource.
// Requires list permission on the project (e.g., "storage.objects.list"
// on
// the containing bucket for testing permission of an
// object).
// Attempting this RPC on a non-existent resource will note in a
// NOT_FOUND
// error if the user has list permission on the project,
// or a PERMISSION_DENIED error otherwise.
func (r *ProjectsOccurrencesService) TestIamPermissions(resource string, testiampermissionsrequest *TestIamPermissionsRequest) *ProjectsOccurrencesTestIamPermissionsCall {
	c := &ProjectsOccurrencesTestIamPermissionsCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.resource = resource
	c.testiampermissionsrequest = testiampermissionsrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsOccurrencesTestIamPermissionsCall) Fields(s ...googleapi.Field) *ProjectsOccurrencesTestIamPermissionsCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsOccurrencesTestIamPermissionsCall) Context(ctx context.Context) *ProjectsOccurrencesTestIamPermissionsCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsOccurrencesTestIamPermissionsCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.testiampermissionsrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha1/{+resource}:testIamPermissions")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "containeranalysis.projects.occurrences.testIamPermissions" call.
// Exactly one of *TestIamPermissionsResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *TestIamPermissionsResponse.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsOccurrencesTestIamPermissionsCall) Do(opts ...googleapi.CallOption) (*TestIamPermissionsResponse, error) {
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
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns permissions that a caller has on the specified note or occurrence\nresource.\nRequires list permission on the project (e.g., \"storage.objects.list\" on\nthe containing bucket for testing permission of an object).\nAttempting this RPC on a non-existent resource will note in a NOT_FOUND\nerror if the user has list permission on the project,\nor a PERMISSION_DENIED error otherwise.",
	//   "flatPath": "v1alpha1/projects/{projectsId}/occurrences/{occurrencesId}:testIamPermissions",
	//   "httpMethod": "POST",
	//   "id": "containeranalysis.projects.occurrences.testIamPermissions",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "REQUIRED: The resource for which the policy detail is being requested.\nSee the operation documentation for the appropriate value for this field.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/occurrences/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1alpha1/{+resource}:testIamPermissions",
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

// method id "containeranalysis.projects.occurrences.update":

type ProjectsOccurrencesUpdateCall struct {
	s          *Service
	name       string
	occurrence *Occurrence
	urlParams_ gensupport.URLParams
	ctx_       context.Context
}

// Update: Updates an existing occurrence.
func (r *ProjectsOccurrencesService) Update(name string, occurrence *Occurrence) *ProjectsOccurrencesUpdateCall {
	c := &ProjectsOccurrencesUpdateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.occurrence = occurrence
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsOccurrencesUpdateCall) Fields(s ...googleapi.Field) *ProjectsOccurrencesUpdateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsOccurrencesUpdateCall) Context(ctx context.Context) *ProjectsOccurrencesUpdateCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsOccurrencesUpdateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.occurrence)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha1/{+name}")
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

// Do executes the "containeranalysis.projects.occurrences.update" call.
// Exactly one of *Occurrence or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Occurrence.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ProjectsOccurrencesUpdateCall) Do(opts ...googleapi.CallOption) (*Occurrence, error) {
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
	ret := &Occurrence{
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
	//   "description": "Updates an existing occurrence.",
	//   "flatPath": "v1alpha1/projects/{projectsId}/occurrences/{occurrencesId}",
	//   "httpMethod": "PUT",
	//   "id": "containeranalysis.projects.occurrences.update",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The name of the occurrence.\nShould be of the form \"projects/{project_id}/occurrences/{occurrence_id}\".",
	//       "location": "path",
	//       "pattern": "^projects/[^/]+/occurrences/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1alpha1/{+name}",
	//   "request": {
	//     "$ref": "Occurrence"
	//   },
	//   "response": {
	//     "$ref": "Occurrence"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "containeranalysis.providers.notes.create":

type ProvidersNotesCreateCall struct {
	s          *Service
	name       string
	note       *Note
	urlParams_ gensupport.URLParams
	ctx_       context.Context
}

// Create: Creates a new note.
func (r *ProvidersNotesService) Create(name string, note *Note) *ProvidersNotesCreateCall {
	c := &ProvidersNotesCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.note = note
	return c
}

// NoteId sets the optional parameter "noteId": The ID to use for this
// note.
func (c *ProvidersNotesCreateCall) NoteId(noteId string) *ProvidersNotesCreateCall {
	c.urlParams_.Set("noteId", noteId)
	return c
}

// Parent sets the optional parameter "parent": The parent field will
// contain the projectId for example:
// "project/{project_id}
func (c *ProvidersNotesCreateCall) Parent(parent string) *ProvidersNotesCreateCall {
	c.urlParams_.Set("parent", parent)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProvidersNotesCreateCall) Fields(s ...googleapi.Field) *ProvidersNotesCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProvidersNotesCreateCall) Context(ctx context.Context) *ProvidersNotesCreateCall {
	c.ctx_ = ctx
	return c
}

func (c *ProvidersNotesCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.note)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha1/{+name}/notes")
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

// Do executes the "containeranalysis.providers.notes.create" call.
// Exactly one of *Note or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Note.ServerResponse.Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProvidersNotesCreateCall) Do(opts ...googleapi.CallOption) (*Note, error) {
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
	ret := &Note{
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
	//   "description": "Creates a new note.",
	//   "flatPath": "v1alpha1/providers/{providersId}/notes",
	//   "httpMethod": "POST",
	//   "id": "containeranalysis.providers.notes.create",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The name of the project.\nShould be of the form \"providers/{provider_id}\".\n@deprecated",
	//       "location": "path",
	//       "pattern": "^providers/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "noteId": {
	//       "description": "The ID to use for this note.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "The parent field will contain the projectId for example:\n\"project/{project_id}",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1alpha1/{+name}/notes",
	//   "request": {
	//     "$ref": "Note"
	//   },
	//   "response": {
	//     "$ref": "Note"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "containeranalysis.providers.notes.delete":

type ProvidersNotesDeleteCall struct {
	s          *Service
	name       string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
}

// Delete: Deletes the given note from the system.
func (r *ProvidersNotesService) Delete(name string) *ProvidersNotesDeleteCall {
	c := &ProvidersNotesDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProvidersNotesDeleteCall) Fields(s ...googleapi.Field) *ProvidersNotesDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProvidersNotesDeleteCall) Context(ctx context.Context) *ProvidersNotesDeleteCall {
	c.ctx_ = ctx
	return c
}

func (c *ProvidersNotesDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha1/{+name}")
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

// Do executes the "containeranalysis.providers.notes.delete" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProvidersNotesDeleteCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
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
	//   "description": "Deletes the given note from the system.",
	//   "flatPath": "v1alpha1/providers/{providersId}/notes/{notesId}",
	//   "httpMethod": "DELETE",
	//   "id": "containeranalysis.providers.notes.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The name of the note in the form\n\"providers/{provider_id}/notes/{note_id}\"",
	//       "location": "path",
	//       "pattern": "^providers/[^/]+/notes/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1alpha1/{+name}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "containeranalysis.providers.notes.get":

type ProvidersNotesGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// Get: Returns the requested occurrence
func (r *ProvidersNotesService) Get(name string) *ProvidersNotesGetCall {
	c := &ProvidersNotesGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProvidersNotesGetCall) Fields(s ...googleapi.Field) *ProvidersNotesGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProvidersNotesGetCall) IfNoneMatch(entityTag string) *ProvidersNotesGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProvidersNotesGetCall) Context(ctx context.Context) *ProvidersNotesGetCall {
	c.ctx_ = ctx
	return c
}

func (c *ProvidersNotesGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha1/{+name}")
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

// Do executes the "containeranalysis.providers.notes.get" call.
// Exactly one of *Note or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Note.ServerResponse.Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProvidersNotesGetCall) Do(opts ...googleapi.CallOption) (*Note, error) {
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
	ret := &Note{
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
	//   "description": "Returns the requested occurrence",
	//   "flatPath": "v1alpha1/providers/{providersId}/notes/{notesId}",
	//   "httpMethod": "GET",
	//   "id": "containeranalysis.providers.notes.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The name of the occurrence in the form\n\"providers/{provider_id}/notes/{note_id}\"",
	//       "location": "path",
	//       "pattern": "^providers/[^/]+/notes/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1alpha1/{+name}",
	//   "response": {
	//     "$ref": "Note"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "containeranalysis.providers.notes.getIamPolicy":

type ProvidersNotesGetIamPolicyCall struct {
	s                   *Service
	resource            string
	getiampolicyrequest *GetIamPolicyRequest
	urlParams_          gensupport.URLParams
	ctx_                context.Context
}

// GetIamPolicy: Gets the access control policy for a note or occurrence
// resource.
// Requires "containeranalysis.notes.setIamPolicy"
// or
// "containeranalysis.occurrences.setIamPolicy" permission if the
// resource is
// a note or occurrence, respectively.
// Attempting this RPC on a resource without the needed permission will
// note
// in a PERMISSION_DENIED error.
// Attempting this RPC on a non-existent resource will note in a
// NOT_FOUND
// error if the user has list permission on the project,
// or a PERMISSION_DENIED error otherwise.
func (r *ProvidersNotesService) GetIamPolicy(resource string, getiampolicyrequest *GetIamPolicyRequest) *ProvidersNotesGetIamPolicyCall {
	c := &ProvidersNotesGetIamPolicyCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.resource = resource
	c.getiampolicyrequest = getiampolicyrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProvidersNotesGetIamPolicyCall) Fields(s ...googleapi.Field) *ProvidersNotesGetIamPolicyCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProvidersNotesGetIamPolicyCall) Context(ctx context.Context) *ProvidersNotesGetIamPolicyCall {
	c.ctx_ = ctx
	return c
}

func (c *ProvidersNotesGetIamPolicyCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.getiampolicyrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha1/{+resource}:getIamPolicy")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "containeranalysis.providers.notes.getIamPolicy" call.
// Exactly one of *Policy or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Policy.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProvidersNotesGetIamPolicyCall) Do(opts ...googleapi.CallOption) (*Policy, error) {
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
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets the access control policy for a note or occurrence resource.\nRequires \"containeranalysis.notes.setIamPolicy\" or\n\"containeranalysis.occurrences.setIamPolicy\" permission if the resource is\na note or occurrence, respectively.\nAttempting this RPC on a resource without the needed permission will note\nin a PERMISSION_DENIED error.\nAttempting this RPC on a non-existent resource will note in a NOT_FOUND\nerror if the user has list permission on the project,\nor a PERMISSION_DENIED error otherwise.",
	//   "flatPath": "v1alpha1/providers/{providersId}/notes/{notesId}:getIamPolicy",
	//   "httpMethod": "POST",
	//   "id": "containeranalysis.providers.notes.getIamPolicy",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "REQUIRED: The resource for which the policy is being requested.\nSee the operation documentation for the appropriate value for this field.",
	//       "location": "path",
	//       "pattern": "^providers/[^/]+/notes/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1alpha1/{+resource}:getIamPolicy",
	//   "request": {
	//     "$ref": "GetIamPolicyRequest"
	//   },
	//   "response": {
	//     "$ref": "Policy"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "containeranalysis.providers.notes.list":

type ProvidersNotesListCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// List: Lists all notes for a given project.  Filters can be used on
// this
// field to list all notes with a specific parameter.
func (r *ProvidersNotesService) List(name string) *ProvidersNotesListCall {
	c := &ProvidersNotesListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Filter sets the optional parameter "filter": The filter expression.
func (c *ProvidersNotesListCall) Filter(filter string) *ProvidersNotesListCall {
	c.urlParams_.Set("filter", filter)
	return c
}

// PageSize sets the optional parameter "pageSize": Number of notes to
// return in the list.
func (c *ProvidersNotesListCall) PageSize(pageSize int64) *ProvidersNotesListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": Token to provide
// to skip to a particular spot in the list.
func (c *ProvidersNotesListCall) PageToken(pageToken string) *ProvidersNotesListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Parent sets the optional parameter "parent": The parent field will
// contain the projectId for example:
// "project/{project_id}
func (c *ProvidersNotesListCall) Parent(parent string) *ProvidersNotesListCall {
	c.urlParams_.Set("parent", parent)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProvidersNotesListCall) Fields(s ...googleapi.Field) *ProvidersNotesListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProvidersNotesListCall) IfNoneMatch(entityTag string) *ProvidersNotesListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProvidersNotesListCall) Context(ctx context.Context) *ProvidersNotesListCall {
	c.ctx_ = ctx
	return c
}

func (c *ProvidersNotesListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha1/{+name}/notes")
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

// Do executes the "containeranalysis.providers.notes.list" call.
// Exactly one of *ListNotesResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListNotesResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProvidersNotesListCall) Do(opts ...googleapi.CallOption) (*ListNotesResponse, error) {
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
	ret := &ListNotesResponse{
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
	//   "description": "Lists all notes for a given project.  Filters can be used on this\nfield to list all notes with a specific parameter.",
	//   "flatPath": "v1alpha1/providers/{providersId}/notes",
	//   "httpMethod": "GET",
	//   "id": "containeranalysis.providers.notes.list",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "The filter expression.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "name": {
	//       "description": "The name field will contain the projectId for example:\n\"providers/{provider_id}\n@deprecated",
	//       "location": "path",
	//       "pattern": "^providers/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Number of notes to return in the list.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Token to provide to skip to a particular spot in the list.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "parent": {
	//       "description": "The parent field will contain the projectId for example:\n\"project/{project_id}",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1alpha1/{+name}/notes",
	//   "response": {
	//     "$ref": "ListNotesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProvidersNotesListCall) Pages(ctx context.Context, f func(*ListNotesResponse) error) error {
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

// method id "containeranalysis.providers.notes.setIamPolicy":

type ProvidersNotesSetIamPolicyCall struct {
	s                   *Service
	resource            string
	setiampolicyrequest *SetIamPolicyRequest
	urlParams_          gensupport.URLParams
	ctx_                context.Context
}

// SetIamPolicy: Sets the access control policy on the specified note or
// occurrence
// resource.
// Requires "containeranalysis.notes.setIamPolicy"
// or
// "containeranalysis.occurrences.setIamPolicy" permission if the
// resource is
// a note or occurrence, respectively.
// Attempting this RPC on a resource without the needed permission will
// note
// in a PERMISSION_DENIED error.
// Attempting this RPC on a non-existent resource will note in a
// NOT_FOUND
// error if the user has list permission on the project, or
// a
// PERMISSION_DENIED error otherwise.
func (r *ProvidersNotesService) SetIamPolicy(resource string, setiampolicyrequest *SetIamPolicyRequest) *ProvidersNotesSetIamPolicyCall {
	c := &ProvidersNotesSetIamPolicyCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.resource = resource
	c.setiampolicyrequest = setiampolicyrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProvidersNotesSetIamPolicyCall) Fields(s ...googleapi.Field) *ProvidersNotesSetIamPolicyCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProvidersNotesSetIamPolicyCall) Context(ctx context.Context) *ProvidersNotesSetIamPolicyCall {
	c.ctx_ = ctx
	return c
}

func (c *ProvidersNotesSetIamPolicyCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.setiampolicyrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha1/{+resource}:setIamPolicy")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "containeranalysis.providers.notes.setIamPolicy" call.
// Exactly one of *Policy or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Policy.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *ProvidersNotesSetIamPolicyCall) Do(opts ...googleapi.CallOption) (*Policy, error) {
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
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Sets the access control policy on the specified note or occurrence\nresource.\nRequires \"containeranalysis.notes.setIamPolicy\" or\n\"containeranalysis.occurrences.setIamPolicy\" permission if the resource is\na note or occurrence, respectively.\nAttempting this RPC on a resource without the needed permission will note\nin a PERMISSION_DENIED error.\nAttempting this RPC on a non-existent resource will note in a NOT_FOUND\nerror if the user has list permission on the project, or a\nPERMISSION_DENIED error otherwise.",
	//   "flatPath": "v1alpha1/providers/{providersId}/notes/{notesId}:setIamPolicy",
	//   "httpMethod": "POST",
	//   "id": "containeranalysis.providers.notes.setIamPolicy",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "REQUIRED: The resource for which the policy is being specified.\nSee the operation documentation for the appropriate value for this field.",
	//       "location": "path",
	//       "pattern": "^providers/[^/]+/notes/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1alpha1/{+resource}:setIamPolicy",
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

// method id "containeranalysis.providers.notes.testIamPermissions":

type ProvidersNotesTestIamPermissionsCall struct {
	s                         *Service
	resource                  string
	testiampermissionsrequest *TestIamPermissionsRequest
	urlParams_                gensupport.URLParams
	ctx_                      context.Context
}

// TestIamPermissions: Returns permissions that a caller has on the
// specified note or occurrence
// resource.
// Requires list permission on the project (e.g., "storage.objects.list"
// on
// the containing bucket for testing permission of an
// object).
// Attempting this RPC on a non-existent resource will note in a
// NOT_FOUND
// error if the user has list permission on the project,
// or a PERMISSION_DENIED error otherwise.
func (r *ProvidersNotesService) TestIamPermissions(resource string, testiampermissionsrequest *TestIamPermissionsRequest) *ProvidersNotesTestIamPermissionsCall {
	c := &ProvidersNotesTestIamPermissionsCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.resource = resource
	c.testiampermissionsrequest = testiampermissionsrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProvidersNotesTestIamPermissionsCall) Fields(s ...googleapi.Field) *ProvidersNotesTestIamPermissionsCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProvidersNotesTestIamPermissionsCall) Context(ctx context.Context) *ProvidersNotesTestIamPermissionsCall {
	c.ctx_ = ctx
	return c
}

func (c *ProvidersNotesTestIamPermissionsCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.testiampermissionsrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha1/{+resource}:testIamPermissions")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "containeranalysis.providers.notes.testIamPermissions" call.
// Exactly one of *TestIamPermissionsResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *TestIamPermissionsResponse.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProvidersNotesTestIamPermissionsCall) Do(opts ...googleapi.CallOption) (*TestIamPermissionsResponse, error) {
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
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns permissions that a caller has on the specified note or occurrence\nresource.\nRequires list permission on the project (e.g., \"storage.objects.list\" on\nthe containing bucket for testing permission of an object).\nAttempting this RPC on a non-existent resource will note in a NOT_FOUND\nerror if the user has list permission on the project,\nor a PERMISSION_DENIED error otherwise.",
	//   "flatPath": "v1alpha1/providers/{providersId}/notes/{notesId}:testIamPermissions",
	//   "httpMethod": "POST",
	//   "id": "containeranalysis.providers.notes.testIamPermissions",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "REQUIRED: The resource for which the policy detail is being requested.\nSee the operation documentation for the appropriate value for this field.",
	//       "location": "path",
	//       "pattern": "^providers/[^/]+/notes/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1alpha1/{+resource}:testIamPermissions",
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

// method id "containeranalysis.providers.notes.update":

type ProvidersNotesUpdateCall struct {
	s          *Service
	name       string
	note       *Note
	urlParams_ gensupport.URLParams
	ctx_       context.Context
}

// Update: Updates an existing note.
func (r *ProvidersNotesService) Update(name string, note *Note) *ProvidersNotesUpdateCall {
	c := &ProvidersNotesUpdateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.note = note
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProvidersNotesUpdateCall) Fields(s ...googleapi.Field) *ProvidersNotesUpdateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProvidersNotesUpdateCall) Context(ctx context.Context) *ProvidersNotesUpdateCall {
	c.ctx_ = ctx
	return c
}

func (c *ProvidersNotesUpdateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.note)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha1/{+name}")
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

// Do executes the "containeranalysis.providers.notes.update" call.
// Exactly one of *Note or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Note.ServerResponse.Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProvidersNotesUpdateCall) Do(opts ...googleapi.CallOption) (*Note, error) {
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
	ret := &Note{
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
	//   "description": "Updates an existing note.",
	//   "flatPath": "v1alpha1/providers/{providersId}/notes/{notesId}",
	//   "httpMethod": "PUT",
	//   "id": "containeranalysis.providers.notes.update",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The name of the note.\nShould be of the form \"projects/{provider_id}/notes/{note_id}\".",
	//       "location": "path",
	//       "pattern": "^providers/[^/]+/notes/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1alpha1/{+name}",
	//   "request": {
	//     "$ref": "Note"
	//   },
	//   "response": {
	//     "$ref": "Note"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "containeranalysis.providers.notes.occurrences.list":

type ProvidersNotesOccurrencesListCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// List: Lists the names of Occurrences linked to a particular Note.
func (r *ProvidersNotesOccurrencesService) List(name string) *ProvidersNotesOccurrencesListCall {
	c := &ProvidersNotesOccurrencesListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Filter sets the optional parameter "filter": The filter expression.
func (c *ProvidersNotesOccurrencesListCall) Filter(filter string) *ProvidersNotesOccurrencesListCall {
	c.urlParams_.Set("filter", filter)
	return c
}

// PageSize sets the optional parameter "pageSize": Number of notes to
// return in the list.
func (c *ProvidersNotesOccurrencesListCall) PageSize(pageSize int64) *ProvidersNotesOccurrencesListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": Token to provide
// to skip to a particular spot in the list.
func (c *ProvidersNotesOccurrencesListCall) PageToken(pageToken string) *ProvidersNotesOccurrencesListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProvidersNotesOccurrencesListCall) Fields(s ...googleapi.Field) *ProvidersNotesOccurrencesListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProvidersNotesOccurrencesListCall) IfNoneMatch(entityTag string) *ProvidersNotesOccurrencesListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProvidersNotesOccurrencesListCall) Context(ctx context.Context) *ProvidersNotesOccurrencesListCall {
	c.ctx_ = ctx
	return c
}

func (c *ProvidersNotesOccurrencesListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha1/{+name}/occurrences")
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

// Do executes the "containeranalysis.providers.notes.occurrences.list" call.
// Exactly one of *ListNoteOccurrencesResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *ListNoteOccurrencesResponse.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProvidersNotesOccurrencesListCall) Do(opts ...googleapi.CallOption) (*ListNoteOccurrencesResponse, error) {
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
	ret := &ListNoteOccurrencesResponse{
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
	//   "description": "Lists the names of Occurrences linked to a particular Note.",
	//   "flatPath": "v1alpha1/providers/{providersId}/notes/{notesId}/occurrences",
	//   "httpMethod": "GET",
	//   "id": "containeranalysis.providers.notes.occurrences.list",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "The filter expression.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "name": {
	//       "description": "The name field will contain the note name for example:\n  \"provider/{provider_id}/notes/{note_id}\"",
	//       "location": "path",
	//       "pattern": "^providers/[^/]+/notes/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Number of notes to return in the list.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Token to provide to skip to a particular spot in the list.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1alpha1/{+name}/occurrences",
	//   "response": {
	//     "$ref": "ListNoteOccurrencesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProvidersNotesOccurrencesListCall) Pages(ctx context.Context, f func(*ListNoteOccurrencesResponse) error) error {
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
