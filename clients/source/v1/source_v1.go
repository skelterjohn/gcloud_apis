// Package source provides access to the Cloud Source Repositories API.
//
// See https://cloud.google.com/eap/cloud-repositories/cloud-source-api
//
// Usage example:
//
//   import "google.golang.org/api/source/v1"
//   ...
//   sourceService, err := source.New(oauthHttpClient)
package source

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

const apiId = "source:v1"
const apiName = "source"
const apiVersion = "v1"
const basePath = "https://source.googleapis.com/"

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
	s.V1 = NewV1Service(s)
	return s, nil
}

type Service struct {
	client   *http.Client
	BasePath string // API endpoint base URL

	Projects *ProjectsService

	V1 *V1Service
}

func NewProjectsService(s *Service) *ProjectsService {
	rs := &ProjectsService{s: s}
	rs.Repos = NewProjectsReposService(s)
	return rs
}

type ProjectsService struct {
	s *Service

	Repos *ProjectsReposService
}

func NewProjectsReposService(s *Service) *ProjectsReposService {
	rs := &ProjectsReposService{s: s}
	rs.Aliases = NewProjectsReposAliasesService(s)
	rs.Files = NewProjectsReposFilesService(s)
	rs.Revisions = NewProjectsReposRevisionsService(s)
	rs.Workspaces = NewProjectsReposWorkspacesService(s)
	return rs
}

type ProjectsReposService struct {
	s *Service

	Aliases *ProjectsReposAliasesService

	Files *ProjectsReposFilesService

	Revisions *ProjectsReposRevisionsService

	Workspaces *ProjectsReposWorkspacesService
}

func NewProjectsReposAliasesService(s *Service) *ProjectsReposAliasesService {
	rs := &ProjectsReposAliasesService{s: s}
	rs.Files = NewProjectsReposAliasesFilesService(s)
	return rs
}

type ProjectsReposAliasesService struct {
	s *Service

	Files *ProjectsReposAliasesFilesService
}

func NewProjectsReposAliasesFilesService(s *Service) *ProjectsReposAliasesFilesService {
	rs := &ProjectsReposAliasesFilesService{s: s}
	return rs
}

type ProjectsReposAliasesFilesService struct {
	s *Service
}

func NewProjectsReposFilesService(s *Service) *ProjectsReposFilesService {
	rs := &ProjectsReposFilesService{s: s}
	return rs
}

type ProjectsReposFilesService struct {
	s *Service
}

func NewProjectsReposRevisionsService(s *Service) *ProjectsReposRevisionsService {
	rs := &ProjectsReposRevisionsService{s: s}
	rs.Files = NewProjectsReposRevisionsFilesService(s)
	return rs
}

type ProjectsReposRevisionsService struct {
	s *Service

	Files *ProjectsReposRevisionsFilesService
}

func NewProjectsReposRevisionsFilesService(s *Service) *ProjectsReposRevisionsFilesService {
	rs := &ProjectsReposRevisionsFilesService{s: s}
	return rs
}

type ProjectsReposRevisionsFilesService struct {
	s *Service
}

func NewProjectsReposWorkspacesService(s *Service) *ProjectsReposWorkspacesService {
	rs := &ProjectsReposWorkspacesService{s: s}
	rs.Files = NewProjectsReposWorkspacesFilesService(s)
	rs.Snapshots = NewProjectsReposWorkspacesSnapshotsService(s)
	return rs
}

type ProjectsReposWorkspacesService struct {
	s *Service

	Files *ProjectsReposWorkspacesFilesService

	Snapshots *ProjectsReposWorkspacesSnapshotsService
}

func NewProjectsReposWorkspacesFilesService(s *Service) *ProjectsReposWorkspacesFilesService {
	rs := &ProjectsReposWorkspacesFilesService{s: s}
	return rs
}

type ProjectsReposWorkspacesFilesService struct {
	s *Service
}

func NewProjectsReposWorkspacesSnapshotsService(s *Service) *ProjectsReposWorkspacesSnapshotsService {
	rs := &ProjectsReposWorkspacesSnapshotsService{s: s}
	rs.Files = NewProjectsReposWorkspacesSnapshotsFilesService(s)
	return rs
}

type ProjectsReposWorkspacesSnapshotsService struct {
	s *Service

	Files *ProjectsReposWorkspacesSnapshotsFilesService
}

func NewProjectsReposWorkspacesSnapshotsFilesService(s *Service) *ProjectsReposWorkspacesSnapshotsFilesService {
	rs := &ProjectsReposWorkspacesSnapshotsFilesService{s: s}
	return rs
}

type ProjectsReposWorkspacesSnapshotsFilesService struct {
	s *Service
}

func NewV1Service(s *Service) *V1Service {
	rs := &V1Service{s: s}
	return rs
}

type V1Service struct {
	s *Service
}

type Action struct {
	// CopyAction: A CopyAction.
	CopyAction *CopyAction `json:"copyAction,omitempty"`

	// DeleteAction: A DeleteAction.
	DeleteAction *DeleteAction `json:"deleteAction,omitempty"`

	// WriteAction: A WriteAction.
	WriteAction *WriteAction `json:"writeAction,omitempty"`
}

type Alias struct {
	// Kind: The alias kind.
	Kind string `json:"kind,omitempty"`

	// Name: The alias name.
	Name string `json:"name,omitempty"`

	// RevisionId: The revision referred to by this alias.
	// For git tags and
	// branches, this is the corresponding hash.
	RevisionId string `json:"revisionId,omitempty"`

	// WorkspaceNames: The list of workspace names whose alias is this
	// one.
	// NOT YET IMPLEMENTED (b/16943429).
	WorkspaceNames []string `json:"workspaceNames,omitempty"`
}

type AliasContext struct {
	// Kind: The alias kind.
	Kind string `json:"kind,omitempty"`

	// Name: The alias name.
	Name string `json:"name,omitempty"`
}

type ChangedFileInfo struct {
	// FromPath: Related file path for copies or renames.
	//
	// For copies, the
	// type will be ADDED and the from_path will point to the
	// source of the
	// copy. For renames, the type will be ADDED, the from_path
	// will point
	// to the source of the rename, and another ChangedFileInfo record
	// with
	// that path will appear with type DELETED. In other words, a rename
	// is
	// represented as a copy plus a delete of the old path.
	FromPath string `json:"fromPath,omitempty"`

	// Hash: A hex-encoded hash for the file.
	// Not necessarily a hash of the
	// file's contents. Two paths in the same
	// revision with the same hash
	// have the same contents with high probability.
	// Empty if the operation
	// is CONFLICTED.
	Hash string `json:"hash,omitempty"`

	// Operation: The operation type for the file.
	Operation string `json:"operation,omitempty"`

	// Path: The path of the file.
	Path string `json:"path,omitempty"`
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

type CommitWorkspaceRequest struct {
	// Author: required
	Author string `json:"author,omitempty"`

	// CurrentSnapshotId: If non-empty, current_snapshot_id must refer to
	// the most recent update to
	// the workspace, or ABORTED is returned.
	CurrentSnapshotId string `json:"currentSnapshotId,omitempty"`

	// Message: required
	Message string `json:"message,omitempty"`

	// Paths: The subset of modified paths to commit. If empty, then commit
	// all
	// modified paths.
	Paths []string `json:"paths,omitempty"`

	// WorkspaceId: The ID of the workspace.
	WorkspaceId *CloudWorkspaceId `json:"workspaceId,omitempty"`
}

type CopyAction struct {
	// FromPath: The path to copy from.
	FromPath string `json:"fromPath,omitempty"`

	// FromRevisionId: The revision ID from which to copy the file.
	FromRevisionId string `json:"fromRevisionId,omitempty"`

	// FromSnapshotId: The snapshot ID from which to copy the file.
	FromSnapshotId string `json:"fromSnapshotId,omitempty"`

	// ToPath: The path to copy to.
	ToPath string `json:"toPath,omitempty"`
}

type CreateWorkspaceRequest struct {
	// Actions: An ordered sequence of actions to perform in the workspace.
	// Can be empty.
	// Specifying actions here instead of using
	// ModifyWorkspace saves one RPC.
	Actions []*Action `json:"actions,omitempty"`

	RepoId *RepoId `json:"repoId,omitempty"`

	// Workspace: The following fields of workspace, with the allowable
	// exception of
	// baseline, must be set. No other fields of workspace
	// should be set.
	//
	// id.name
	// Provides the name of the workspace and must
	// be unique within the repo.
	// Note: Do not set field id.repo_id.  The
	// repo_id is provided above as a
	// CreateWorkspaceRequest
	// field.
	//
	// alias:
	// If alias names an existing movable alias, the
	// workspace's baseline
	// is set to the alias's revision.
	//
	// If alias does
	// not name an existing movable alias, then the workspace is
	// created
	// with no baseline. When the workspace is committed, a new
	// root
	// revision is created with no parents. The new revision becomes
	// the
	// workspace's baseline and the alias name is used to create a
	// movable alias
	// referring to the revision.
	//
	// baseline:
	// A revision ID
	// (hexadecimal string) for sequencing. If non-empty, alias
	// must name an
	// existing movable alias and baseline must match the alias's
	// revision
	// ID.
	Workspace *Workspace `json:"workspace,omitempty"`
}

type DeleteAction struct {
	// Path: The path of the file or directory. If path refers to
	// a
	// directory, the directory and its contents are deleted.
	Path string `json:"path,omitempty"`
}

type DirectoryEntry struct {
	// Info: Information about the entry.
	Info *FileInfo `json:"info,omitempty"`

	// IsDir: Whether the entry is a file or directory.
	IsDir bool `json:"isDir,omitempty"`

	// LastModifiedRevisionId: ID of the revision that most recently
	// modified this file.
	LastModifiedRevisionId string `json:"lastModifiedRevisionId,omitempty"`

	// Name: Name of the entry relative to the directory.
	Name string `json:"name,omitempty"`
}

type Empty struct {
}

type ExternalReference struct {
}

type File struct {
	// Contents: The contents of the file.
	Contents string `json:"contents,omitempty"`

	// Info: Information about the file.
	Info *FileInfo `json:"info,omitempty"`

	// Path: The path to the file starting from the root of the revision.
	Path string `json:"path,omitempty"`
}

type FileInfo struct {
	// Hash: A hex-encoded cryptographic hash of the file's contents,
	// possibly with other data.
	Hash string `json:"hash,omitempty"`

	// IsText: An educated guess as to whether the file is human-readable
	// text, or
	// binary. Typically available only when file contents are
	// retrieved (since
	// the guess depends on examining a prefix of the
	// contents), but some systems
	// might store this metadata for every file.
	IsText bool `json:"isText,omitempty"`

	// Mode: The mode of the file: an executable, a symbolic link, or
	// neither.
	Mode string `json:"mode,omitempty"`

	// Size: The size of the file in bytes.
	Size int64 `json:"size,omitempty,string"`
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

type GetRevisionsResponse struct {
	// Revisions: The revisions.
	Revisions []*Revision `json:"revisions,omitempty"`
}

type GitSourceContext struct {
	// RevisionId: Git commit hash.
	// required.
	RevisionId string `json:"revisionId,omitempty"`

	// Url: Git repository URL.
	Url string `json:"url,omitempty"`
}

type ListAliasesResponse struct {
	// Aliases: The list of aliases.
	Aliases []*Alias `json:"aliases,omitempty"`

	// NextPageToken: Use as the value of page_token in the next
	// call to
	// obtain the next page of results.
	// If empty, there are no more results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// TotalAliases: The total number of aliases in the repo of the kind
	// specified in the
	// request.
	TotalAliases int64 `json:"totalAliases,omitempty"`
}

type ListChangedFilesRequest struct {
	// PageSize: The maximum number of values to return.
	PageSize int64 `json:"pageSize,omitempty"`

	// PageToken: The value of next_page_token from the previous call.
	// Omit
	// for the first page.
	PageToken string `json:"pageToken,omitempty"`

	// SourceContext1: The first source context to compare.
	SourceContext1 *SourceContext `json:"sourceContext1,omitempty"`

	// SourceContext2: The second source context to compare.
	SourceContext2 *SourceContext `json:"sourceContext2,omitempty"`
}

type ListChangedFilesResponse struct {
	// ChangedFiles: Note: ChangedFileInfo.from_path is not set here.
	// ListChangedFiles does not
	// perform rename/copy detection.
	//
	// The
	// ChangedFileInfo.Type describes the changes from source_context1
	// to
	// source_context2. Thus ADDED would mean a file is not present
	// in
	// source_context1 but is present in source_context2.
	ChangedFiles []*ChangedFileInfo `json:"changedFiles,omitempty"`

	// NextPageToken: Use as the value of page_token in the next
	// call to
	// obtain the next page of results.
	// If empty, there are no more results.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type ListFilesResponse struct {
	// Files: info.size and contents are not set.
	Files []*File `json:"files,omitempty"`

	// NextPageToken: Use as the value of page_token in the next
	// call to
	// obtain the next page of results.
	// If empty, there are no more results.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type ListReposResponse struct {
	// Repos: The listed repos.
	Repos []*Repo `json:"repos,omitempty"`
}

type ListRevisionsResponse struct {
	// NextPageToken: Use as the value of page_token in the next
	// call to
	// obtain the next page of results.
	// If empty, there are no more results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Revisions: The list of revisions.
	Revisions []*Revision `json:"revisions,omitempty"`
}

type ListSnapshotsResponse struct {
	// NextPageToken: Use as the value of page_token in the next
	// call to
	// obtain the next page of results.
	// If empty, there are no more results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Snapshots: The list of snapshots.
	Snapshots []*Snapshot `json:"snapshots,omitempty"`
}

type ListWorkspacesResponse struct {
	// Workspaces: The listed workspaces.
	Workspaces []*Workspace `json:"workspaces,omitempty"`
}

type MergeInfo struct {
	// CommonAncestorRevisionId: Revision ID of the closest common ancestor
	// of the file trees that are
	// participating in a refresh or merge.
	// During a refresh, the common
	// ancestor is the baseline of the
	// workspace.  During a merge of two
	// branches, the common ancestor is
	// derived from the workspace baseline and
	// the alias of the branch being
	// merged in.  The repository state at the
	// common ancestor provides the
	// base version for a three-way merge.
	CommonAncestorRevisionId string `json:"commonAncestorRevisionId,omitempty"`

	// IsRefresh: If true, a refresh operation is in progress.  If false, a
	// merge is in
	// progress.
	IsRefresh bool `json:"isRefresh,omitempty"`

	// OtherRevisionId: During a refresh, the ID of the revision with which
	// the workspace is being
	// refreshed. This is the revision ID to which
	// the workspace's alias refers
	// at the time of the RefreshWorkspace
	// call. During a merge, the ID of the
	// revision that's being merged into
	// the workspace's alias. This is the
	// revision_id field of the
	// MergeRequest.
	OtherRevisionId string `json:"otherRevisionId,omitempty"`

	// WorkspaceAfterSnapshotId: The workspace snapshot immediately after
	// the refresh or merge RPC
	// completes.  If a file has conflicts, this
	// snapshot contains the
	// version of the file with conflict markers.
	WorkspaceAfterSnapshotId string `json:"workspaceAfterSnapshotId,omitempty"`

	// WorkspaceBeforeSnapshotId: During a refresh, the snapshot ID of the
	// latest change to the workspace
	// before the refresh.  During a merge,
	// the workspace's baseline, which is
	// identical to the commit hash of
	// the workspace's alias before initiating
	// the merge.
	WorkspaceBeforeSnapshotId string `json:"workspaceBeforeSnapshotId,omitempty"`
}

type MergeRequest struct {
	// RevisionId: The other revision to be merged.
	RevisionId string `json:"revisionId,omitempty"`

	// WorkspaceId: The workspace to use for the merge. The revision
	// referred to
	// by the workspace's alias will be one of the revisions
	// merged.
	WorkspaceId *CloudWorkspaceId `json:"workspaceId,omitempty"`
}

type ModifyWorkspaceRequest struct {
	// Actions: An ordered sequence of actions to perform in the workspace.
	// May not be
	// empty.
	Actions []*Action `json:"actions,omitempty"`

	// CurrentSnapshotId: If non-empty, current_snapshot_id must refer to
	// the most recent update to
	// the workspace, or ABORTED is returned.
	CurrentSnapshotId string `json:"currentSnapshotId,omitempty"`

	// WorkspaceId: The ID of the workspace.
	WorkspaceId *CloudWorkspaceId `json:"workspaceId,omitempty"`
}

type ProjectRepoId struct {
	// ProjectId: The ID of the project.
	ProjectId string `json:"projectId,omitempty"`

	// RepoName: The name of the repo. Leave empty for the default repo.
	RepoName string `json:"repoName,omitempty"`
}

type ReadResponse struct {
	// Entries: Contains the directory entries if the request specifies a
	// directory.
	Entries []*DirectoryEntry `json:"entries,omitempty"`

	// ExternalReference: The read path denotes a Git submodule.
	ExternalReference *ExternalReference `json:"externalReference,omitempty"`

	// File: Contains file metadata and contents if the request specifies a
	// file.
	File *File `json:"file,omitempty"`

	// NextPageToken: Use as the value of page_token in the next
	// call to
	// obtain the next page of results.
	// If empty, there are no more results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// SourceContext: Returns the SourceContext actually used, resolving any
	// alias in the input
	// SourceContext into its revision ID and returning
	// the actual current
	// snapshot ID if the read was from a workspace with
	// an unspecified snapshot
	// ID.
	SourceContext *SourceContext `json:"sourceContext,omitempty"`
}

type RefreshWorkspaceRequest struct {
	// WorkspaceId: The ID of the workspace.
	WorkspaceId *CloudWorkspaceId `json:"workspaceId,omitempty"`
}

type Repo struct {
	// CreateTime: Timestamp when the repo was created.
	CreateTime string `json:"createTime,omitempty"`

	// Id: Randomly generated ID that uniquely identifies a repo.
	Id string `json:"id,omitempty"`

	// Name: Human-readable, user-defined name of the repository. Names must
	// be
	// alphanumeric, lowercase, begin with a letter, and be between 3 and
	// 63
	// characters long. The - character can appear in the middle
	// positions.
	// (Names must satisfy the regular
	// expression
	// a-z{1,61}[a-z0-9].)
	Name string `json:"name,omitempty"`

	// ProjectId: Immutable, globally unique, DNS-compatible textual
	// identifier.
	// Examples: user-chosen-project-id, yellow-banana-33.
	ProjectId string `json:"projectId,omitempty"`

	// RepoSyncConfig: How RepoSync is configured for this repo. If missing,
	// this
	// repo is not set up for RepoSync.
	RepoSyncConfig *RepoSyncConfig `json:"repoSyncConfig,omitempty"`

	// State: The state the repo is in.
	State string `json:"state,omitempty"`

	// Vcs: The version control system of the repo.
	Vcs string `json:"vcs,omitempty"`
}

type RepoId struct {
	// ProjectRepoId: A combination of a project ID and a repo name.
	ProjectRepoId *ProjectRepoId `json:"projectRepoId,omitempty"`

	// Uid: A server-assigned, globally unique identifier.
	Uid string `json:"uid,omitempty"`
}

type RepoSyncConfig struct {
	// ExternalRepoUrl: If this repo is enabled for RepoSync, this will be
	// the URL of the
	// external repo that this repo should sync with.
	ExternalRepoUrl string `json:"externalRepoUrl,omitempty"`

	// Status: The status of RepoSync.
	Status string `json:"status,omitempty"`
}

type ResolveFilesRequest struct {
	// ResolvedPaths: Files that should be marked as resolved in the
	// workspace.  All files in
	// resolved_paths must currently be in the
	// CONFLICTED state in
	// Workspace.changed_files.  NOTE: Changing a file's
	// contents to match the
	// contents in the workspace baseline, then
	// calling ResolveFiles on it, will
	// cause the file to be removed from
	// the changed_files list entirely.
	// If resolved_paths is empty,
	// INVALID_ARGUMENT is returned.
	// If resolved_paths contains duplicates,
	// INVALID_ARGUMENT is returned.
	// If resolved_paths contains a path that
	// was never unresolved,
	// or has already been resolved,
	// FAILED_PRECONDITION is returned.
	ResolvedPaths []string `json:"resolvedPaths,omitempty"`

	// WorkspaceId: The ID of the workspace.
	WorkspaceId *CloudWorkspaceId `json:"workspaceId,omitempty"`
}

type RevertRefreshRequest struct {
	// WorkspaceId: The ID of the workspace.
	WorkspaceId *CloudWorkspaceId `json:"workspaceId,omitempty"`
}

type Revision struct {
	// Author: The name of the user who wrote the revision. (In Git, this
	// can
	// differ from committer.)
	Author string `json:"author,omitempty"`

	// BranchName: Mercurial branch name.
	BranchName string `json:"branchName,omitempty"`

	// ChangedFiles: Files changed in this revision.
	ChangedFiles []*ChangedFileInfo `json:"changedFiles,omitempty"`

	// ChangedFilesUnknown: In some cases changed-file
	// information is
	// generated asynchronously. So there is a period
	// of time when it is not
	// available. This field encodes that fact.
	// (An empty changed_files
	// field is not sufficient, since it is
	// possible for a revision to have
	// no changed files.)
	ChangedFilesUnknown bool `json:"changedFilesUnknown,omitempty"`

	// CommitMessage: The message added by the committer.
	CommitMessage string `json:"commitMessage,omitempty"`

	// CommitTime: When the revision was committed.
	CommitTime string `json:"commitTime,omitempty"`

	// Committer: The name of the user who committed the revision.
	Committer string `json:"committer,omitempty"`

	// CreateTime: When the revision was made. This may or may not be
	// reliable, depending on
	// the version control system being used.
	CreateTime string `json:"createTime,omitempty"`

	// Id: The unique ID of the revision. For many version control systems,
	// this
	// will be string of hex digits representing a hash value.
	Id string `json:"id,omitempty"`

	// ParentIds: The revision IDs of this revision's parents.
	ParentIds []string `json:"parentIds,omitempty"`
}

type Snapshot struct {
	// ChangedFiles: The set of files modified in this snapshot, relative to
	// the workspace
	// baseline. ChangedFileInfo.from_path is not set.
	ChangedFiles []*ChangedFileInfo `json:"changedFiles,omitempty"`

	// CreateTime: Timestamp when the snapshot was created.
	CreateTime string `json:"createTime,omitempty"`

	// SnapshotId: The ID of the snapshot.
	SnapshotId string `json:"snapshotId,omitempty"`
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

type Workspace struct {
	// Alias: The alias associated with the workspace. When the workspace is
	// committed,
	// this alias will be moved to point to the new revision.
	Alias string `json:"alias,omitempty"`

	// Baseline: The revision of the workspace's alias when the workspace
	// was
	// created.
	Baseline string `json:"baseline,omitempty"`

	// ChangedFiles: The set of files modified in this workspace.
	ChangedFiles []*ChangedFileInfo `json:"changedFiles,omitempty"`

	// CurrentSnapshotId: If non-empty, current_snapshot_id refers to the
	// most recent update to the
	// workspace.
	CurrentSnapshotId string `json:"currentSnapshotId,omitempty"`

	// Id: The ID of the workspace.
	Id *CloudWorkspaceId `json:"id,omitempty"`

	// MergeInfo: Information needed to manage a refresh or merge operation.
	// Present only
	// during a merge (i.e. after a call to Merge) or a call
	// to
	// RefreshWorkspace which results in conflicts.
	MergeInfo *MergeInfo `json:"mergeInfo,omitempty"`
}

type WriteAction struct {
	// Contents: The new contents of the file.
	Contents string `json:"contents,omitempty"`

	// Mode: The new mode of the file.
	Mode string `json:"mode,omitempty"`

	// Path: The path of the file to write.
	Path string `json:"path,omitempty"`
}

// method id "source.projects.repos.create":

type ProjectsReposCreateCall struct {
	s         *Service
	projectId string
	repo      *Repo
	opt_      map[string]interface{}
}

// Create: Creates a repo in the given project. The provided repo
// message should have
// its name field set to the desired repo name. No
// other repo fields should
// be set. Omitting the name is the same as
// specifying "default"
//
// Repo names must satisfy the regular
// expression
// `a-z{1,61}[a-z0-9]`. (Note that repo names must contain
// at
// least three characters and may not contain underscores.) The
// special name
// "default" is the default repo for the project; this is
// the repo shown when
// visiting the Cloud Developers Console, and can be
// accessed via git's HTTP
// protocol at
// `https://source.developers.google.com/p/PROJECT_ID`. You may
// create
// other repos with this API and access them
// at
// `https://source.developers.google.com/p/PROJECT_ID/r/NAME`.
func (r *ProjectsReposService) Create(projectId string, repo *Repo) *ProjectsReposCreateCall {
	c := &ProjectsReposCreateCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.repo = repo
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsReposCreateCall) Fields(s ...googleapi.Field) *ProjectsReposCreateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsReposCreateCall) Do() (*Repo, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.repo)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/repos")
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
	var ret *Repo
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a repo in the given project. The provided repo message should have\nits name field set to the desired repo name. No other repo fields should\nbe set. Omitting the name is the same as specifying \"default\"\n\nRepo names must satisfy the regular expression\n`a-z{1,61}[a-z0-9]`. (Note that repo names must contain at\nleast three characters and may not contain underscores.) The special name\n\"default\" is the default repo for the project; this is the repo shown when\nvisiting the Cloud Developers Console, and can be accessed via git's HTTP\nprotocol at `https://source.developers.google.com/p/PROJECT_ID`. You may\ncreate other repos with this API and access them at\n`https://source.developers.google.com/p/PROJECT_ID/r/NAME`.",
	//   "flatPath": "v1/projects/{projectId}/repos",
	//   "httpMethod": "POST",
	//   "id": "source.projects.repos.create",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "The project in which to create the repo.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/repos",
	//   "request": {
	//     "$ref": "Repo"
	//   },
	//   "response": {
	//     "$ref": "Repo"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "source.projects.repos.delete":

type ProjectsReposDeleteCall struct {
	s         *Service
	projectId string
	repoName  string
	opt_      map[string]interface{}
}

// Delete: Deletes a repo.
func (r *ProjectsReposService) Delete(projectId string, repoName string) *ProjectsReposDeleteCall {
	c := &ProjectsReposDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.repoName = repoName
	return c
}

// Uid sets the optional parameter "uid": A server-assigned, globally
// unique identifier.
func (c *ProjectsReposDeleteCall) Uid(uid string) *ProjectsReposDeleteCall {
	c.opt_["uid"] = uid
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsReposDeleteCall) Fields(s ...googleapi.Field) *ProjectsReposDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsReposDeleteCall) Do() (*Empty, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["uid"]; ok {
		params.Set("uid", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/repos/{repoName}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"repoName":  c.repoName,
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
	//   "description": "Deletes a repo.",
	//   "flatPath": "v1/projects/{projectId}/repos/{repoName}",
	//   "httpMethod": "DELETE",
	//   "id": "source.projects.repos.delete",
	//   "parameterOrder": [
	//     "projectId",
	//     "repoName"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "The ID of the project.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "repoName": {
	//       "description": "The name of the repo. Leave empty for the default repo.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "uid": {
	//       "description": "A server-assigned, globally unique identifier.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/repos/{repoName}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "source.projects.repos.get":

type ProjectsReposGetCall struct {
	s         *Service
	projectId string
	repoName  string
	opt_      map[string]interface{}
}

// Get: Returns information about a repo.
func (r *ProjectsReposService) Get(projectId string, repoName string) *ProjectsReposGetCall {
	c := &ProjectsReposGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.repoName = repoName
	return c
}

// Uid sets the optional parameter "uid": A server-assigned, globally
// unique identifier.
func (c *ProjectsReposGetCall) Uid(uid string) *ProjectsReposGetCall {
	c.opt_["uid"] = uid
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsReposGetCall) Fields(s ...googleapi.Field) *ProjectsReposGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsReposGetCall) Do() (*Repo, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["uid"]; ok {
		params.Set("uid", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/repos/{repoName}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"repoName":  c.repoName,
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
	var ret *Repo
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns information about a repo.",
	//   "flatPath": "v1/projects/{projectId}/repos/{repoName}",
	//   "httpMethod": "GET",
	//   "id": "source.projects.repos.get",
	//   "parameterOrder": [
	//     "projectId",
	//     "repoName"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "The ID of the project.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "repoName": {
	//       "description": "The name of the repo. Leave empty for the default repo.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "uid": {
	//       "description": "A server-assigned, globally unique identifier.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/repos/{repoName}",
	//   "response": {
	//     "$ref": "Repo"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "source.projects.repos.list":

type ProjectsReposListCall struct {
	s         *Service
	projectId string
	opt_      map[string]interface{}
}

// List: Returns all repos belonging to a project, specified by its
// project ID. The
// response list is sorted by name with the default repo
// listed first.
func (r *ProjectsReposService) List(projectId string) *ProjectsReposListCall {
	c := &ProjectsReposListCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsReposListCall) Fields(s ...googleapi.Field) *ProjectsReposListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsReposListCall) Do() (*ListReposResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/repos")
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
	var ret *ListReposResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns all repos belonging to a project, specified by its project ID. The\nresponse list is sorted by name with the default repo listed first.",
	//   "flatPath": "v1/projects/{projectId}/repos",
	//   "httpMethod": "GET",
	//   "id": "source.projects.repos.list",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "The project ID whose repos should be listed.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/repos",
	//   "response": {
	//     "$ref": "ListReposResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "source.projects.repos.merge":

type ProjectsReposMergeCall struct {
	s            *Service
	projectId    string
	repoName     string
	mergerequest *MergeRequest
	opt_         map[string]interface{}
}

// Merge: Merges a revision into a movable alias, using a workspace
// associated with
// that alias to store modified files. The workspace
// must not have any
// modified files. Note that Merge neither creates the
// workspace nor commits
// it; those actions must be done separately.
// Returns ABORTED when the
// workspace is simultaneously modified by
// another client.
func (r *ProjectsReposService) Merge(projectId string, repoName string, mergerequest *MergeRequest) *ProjectsReposMergeCall {
	c := &ProjectsReposMergeCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.repoName = repoName
	c.mergerequest = mergerequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsReposMergeCall) Fields(s ...googleapi.Field) *ProjectsReposMergeCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsReposMergeCall) Do() (*Workspace, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.mergerequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/repos/{repoName}:merge")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"repoName":  c.repoName,
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
	var ret *Workspace
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Merges a revision into a movable alias, using a workspace associated with\nthat alias to store modified files. The workspace must not have any\nmodified files. Note that Merge neither creates the workspace nor commits\nit; those actions must be done separately. Returns ABORTED when the\nworkspace is simultaneously modified by another client.",
	//   "flatPath": "v1/projects/{projectId}/repos/{repoName}:merge",
	//   "httpMethod": "POST",
	//   "id": "source.projects.repos.merge",
	//   "parameterOrder": [
	//     "projectId",
	//     "repoName"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "The ID of the project.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "repoName": {
	//       "description": "The name of the repo. Leave empty for the default repo.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/repos/{repoName}:merge",
	//   "request": {
	//     "$ref": "MergeRequest"
	//   },
	//   "response": {
	//     "$ref": "Workspace"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "source.projects.repos.aliases.create":

type ProjectsReposAliasesCreateCall struct {
	s         *Service
	projectId string
	repoName  string
	alias     *Alias
	opt_      map[string]interface{}
}

// Create: Creates a new alias. It is an ALREADY_EXISTS error if an
// alias with that
// name and kind already exists.
func (r *ProjectsReposAliasesService) Create(projectId string, repoName string, alias *Alias) *ProjectsReposAliasesCreateCall {
	c := &ProjectsReposAliasesCreateCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.repoName = repoName
	c.alias = alias
	return c
}

// Uid sets the optional parameter "uid": A server-assigned, globally
// unique identifier.
func (c *ProjectsReposAliasesCreateCall) Uid(uid string) *ProjectsReposAliasesCreateCall {
	c.opt_["uid"] = uid
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsReposAliasesCreateCall) Fields(s ...googleapi.Field) *ProjectsReposAliasesCreateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsReposAliasesCreateCall) Do() (*Alias, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.alias)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["uid"]; ok {
		params.Set("uid", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/repos/{repoName}/aliases")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"repoName":  c.repoName,
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
	var ret *Alias
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a new alias. It is an ALREADY_EXISTS error if an alias with that\nname and kind already exists.",
	//   "flatPath": "v1/projects/{projectId}/repos/{repoName}/aliases",
	//   "httpMethod": "POST",
	//   "id": "source.projects.repos.aliases.create",
	//   "parameterOrder": [
	//     "projectId",
	//     "repoName"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "The ID of the project.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "repoName": {
	//       "description": "The name of the repo. Leave empty for the default repo.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "uid": {
	//       "description": "A server-assigned, globally unique identifier.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/repos/{repoName}/aliases",
	//   "request": {
	//     "$ref": "Alias"
	//   },
	//   "response": {
	//     "$ref": "Alias"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "source.projects.repos.aliases.delete":

type ProjectsReposAliasesDeleteCall struct {
	s         *Service
	projectId string
	repoName  string
	kind      string
	name      string
	opt_      map[string]interface{}
}

// Delete: Deletes the alias with the given name and kind. Kind cannot
// be ANY.  If
// the alias does not exist, NOT_FOUND is returned.  If the
// request provides
// a revision ID and the alias does not refer to that
// revision, ABORTED is
// returned.
func (r *ProjectsReposAliasesService) Delete(projectId string, repoName string, kind string, name string) *ProjectsReposAliasesDeleteCall {
	c := &ProjectsReposAliasesDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.repoName = repoName
	c.kind = kind
	c.name = name
	return c
}

// RevisionId sets the optional parameter "revisionId": If non-empty,
// must match the revision that the alias refers to.
func (c *ProjectsReposAliasesDeleteCall) RevisionId(revisionId string) *ProjectsReposAliasesDeleteCall {
	c.opt_["revisionId"] = revisionId
	return c
}

// Uid sets the optional parameter "uid": A server-assigned, globally
// unique identifier.
func (c *ProjectsReposAliasesDeleteCall) Uid(uid string) *ProjectsReposAliasesDeleteCall {
	c.opt_["uid"] = uid
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsReposAliasesDeleteCall) Fields(s ...googleapi.Field) *ProjectsReposAliasesDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsReposAliasesDeleteCall) Do() (*Empty, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["revisionId"]; ok {
		params.Set("revisionId", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["uid"]; ok {
		params.Set("uid", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/repos/{repoName}/aliases/{kind}/{name}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"repoName":  c.repoName,
		"kind":      c.kind,
		"name":      c.name,
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
	//   "description": "Deletes the alias with the given name and kind. Kind cannot be ANY.  If\nthe alias does not exist, NOT_FOUND is returned.  If the request provides\na revision ID and the alias does not refer to that revision, ABORTED is\nreturned.",
	//   "flatPath": "v1/projects/{projectId}/repos/{repoName}/aliases/{kind}/{name}",
	//   "httpMethod": "DELETE",
	//   "id": "source.projects.repos.aliases.delete",
	//   "parameterOrder": [
	//     "projectId",
	//     "repoName",
	//     "kind",
	//     "name"
	//   ],
	//   "parameters": {
	//     "kind": {
	//       "description": "The kind of the alias to delete.",
	//       "enum": [
	//         "ANY",
	//         "FIXED",
	//         "MOVABLE",
	//         "MERCURIAL_BRANCH_DEPRECATED",
	//         "OTHER",
	//         "SPECIAL_DEPRECATED"
	//       ],
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "name": {
	//       "description": "The name of the alias to delete.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "The ID of the project.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "repoName": {
	//       "description": "The name of the repo. Leave empty for the default repo.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "revisionId": {
	//       "description": "If non-empty, must match the revision that the alias refers to.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "uid": {
	//       "description": "A server-assigned, globally unique identifier.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/repos/{repoName}/aliases/{kind}/{name}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "source.projects.repos.aliases.get":

type ProjectsReposAliasesGetCall struct {
	s         *Service
	projectId string
	repoName  string
	kind      string
	name      string
	opt_      map[string]interface{}
}

// Get: Returns information about an alias. Kind ANY returns a FIXED
// or
// MOVABLE alias, in that order, and ignores all other kinds.
func (r *ProjectsReposAliasesService) Get(projectId string, repoName string, kind string, name string) *ProjectsReposAliasesGetCall {
	c := &ProjectsReposAliasesGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.repoName = repoName
	c.kind = kind
	c.name = name
	return c
}

// Uid sets the optional parameter "uid": A server-assigned, globally
// unique identifier.
func (c *ProjectsReposAliasesGetCall) Uid(uid string) *ProjectsReposAliasesGetCall {
	c.opt_["uid"] = uid
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsReposAliasesGetCall) Fields(s ...googleapi.Field) *ProjectsReposAliasesGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsReposAliasesGetCall) Do() (*Alias, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["uid"]; ok {
		params.Set("uid", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/repos/{repoName}/aliases/{kind}/{name}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"repoName":  c.repoName,
		"kind":      c.kind,
		"name":      c.name,
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
	var ret *Alias
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns information about an alias. Kind ANY returns a FIXED or\nMOVABLE alias, in that order, and ignores all other kinds.",
	//   "flatPath": "v1/projects/{projectId}/repos/{repoName}/aliases/{kind}/{name}",
	//   "httpMethod": "GET",
	//   "id": "source.projects.repos.aliases.get",
	//   "parameterOrder": [
	//     "projectId",
	//     "repoName",
	//     "kind",
	//     "name"
	//   ],
	//   "parameters": {
	//     "kind": {
	//       "description": "The kind of the alias.",
	//       "enum": [
	//         "ANY",
	//         "FIXED",
	//         "MOVABLE",
	//         "MERCURIAL_BRANCH_DEPRECATED",
	//         "OTHER",
	//         "SPECIAL_DEPRECATED"
	//       ],
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "name": {
	//       "description": "The alias name.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "The ID of the project.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "repoName": {
	//       "description": "The name of the repo. Leave empty for the default repo.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "uid": {
	//       "description": "A server-assigned, globally unique identifier.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/repos/{repoName}/aliases/{kind}/{name}",
	//   "response": {
	//     "$ref": "Alias"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "source.projects.repos.aliases.list":

type ProjectsReposAliasesListCall struct {
	s         *Service
	projectId string
	repoName  string
	opt_      map[string]interface{}
}

// List: Returns a list of aliases of the given kind. Kind ANY returns
// all aliases
// in the repo. The order in which the aliases are returned
// is undefined.
func (r *ProjectsReposAliasesService) List(projectId string, repoName string) *ProjectsReposAliasesListCall {
	c := &ProjectsReposAliasesListCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.repoName = repoName
	return c
}

// Kind sets the optional parameter "kind": Return only aliases of this
// kind.
func (c *ProjectsReposAliasesListCall) Kind(kind string) *ProjectsReposAliasesListCall {
	c.opt_["kind"] = kind
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of values to return.
func (c *ProjectsReposAliasesListCall) PageSize(pageSize int64) *ProjectsReposAliasesListCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken": The value of
// next_page_token from the previous call.
// Omit for the first page.
func (c *ProjectsReposAliasesListCall) PageToken(pageToken string) *ProjectsReposAliasesListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Uid sets the optional parameter "uid": A server-assigned, globally
// unique identifier.
func (c *ProjectsReposAliasesListCall) Uid(uid string) *ProjectsReposAliasesListCall {
	c.opt_["uid"] = uid
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsReposAliasesListCall) Fields(s ...googleapi.Field) *ProjectsReposAliasesListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsReposAliasesListCall) Do() (*ListAliasesResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["kind"]; ok {
		params.Set("kind", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageSize"]; ok {
		params.Set("pageSize", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["uid"]; ok {
		params.Set("uid", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/repos/{repoName}/aliases")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"repoName":  c.repoName,
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
	var ret *ListAliasesResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns a list of aliases of the given kind. Kind ANY returns all aliases\nin the repo. The order in which the aliases are returned is undefined.",
	//   "flatPath": "v1/projects/{projectId}/repos/{repoName}/aliases",
	//   "httpMethod": "GET",
	//   "id": "source.projects.repos.aliases.list",
	//   "parameterOrder": [
	//     "projectId",
	//     "repoName"
	//   ],
	//   "parameters": {
	//     "kind": {
	//       "description": "Return only aliases of this kind.",
	//       "enum": [
	//         "ANY",
	//         "FIXED",
	//         "MOVABLE",
	//         "MERCURIAL_BRANCH_DEPRECATED",
	//         "OTHER",
	//         "SPECIAL_DEPRECATED"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "The maximum number of values to return.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The value of next_page_token from the previous call.\nOmit for the first page.",
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
	//       "description": "The name of the repo. Leave empty for the default repo.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "uid": {
	//       "description": "A server-assigned, globally unique identifier.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/repos/{repoName}/aliases",
	//   "response": {
	//     "$ref": "ListAliasesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "source.projects.repos.aliases.listFiles":

type ProjectsReposAliasesListFilesCall struct {
	s         *Service
	projectId string
	repoName  string
	kind      string
	name      string
	opt_      map[string]interface{}
}

// ListFiles: ListFiles returns a list of all files in a SourceContext.
// The
// information about each file includes its path and its hash.
// The
// result is ordered by path. Pagination is supported.
func (r *ProjectsReposAliasesService) ListFiles(projectId string, repoName string, kind string, name string) *ProjectsReposAliasesListFilesCall {
	c := &ProjectsReposAliasesListFilesCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.repoName = repoName
	c.kind = kind
	c.name = name
	return c
}

// AliasName sets the optional parameter "aliasName": The name of an
// alias (branch, tag, etc.).
func (c *ProjectsReposAliasesListFilesCall) AliasName(aliasName string) *ProjectsReposAliasesListFilesCall {
	c.opt_["aliasName"] = aliasName
	return c
}

// CloudWorkspaceSnapshotId sets the optional parameter
// "cloudWorkspace.snapshotId": The ID of the snapshot.
// An empty
// snapshot_id refers to the most recent snapshot.
func (c *ProjectsReposAliasesListFilesCall) CloudWorkspaceSnapshotId(cloudWorkspaceSnapshotId string) *ProjectsReposAliasesListFilesCall {
	c.opt_["cloudWorkspace.snapshotId"] = cloudWorkspaceSnapshotId
	return c
}

// CloudWorkspaceWorkspaceIdName sets the optional parameter
// "cloudWorkspace.workspaceId.name": The unique name of the workspace
// within the repo.  This is the name
// chosen by the client in the Source
// API's CreateWorkspace method.
func (c *ProjectsReposAliasesListFilesCall) CloudWorkspaceWorkspaceIdName(cloudWorkspaceWorkspaceIdName string) *ProjectsReposAliasesListFilesCall {
	c.opt_["cloudWorkspace.workspaceId.name"] = cloudWorkspaceWorkspaceIdName
	return c
}

// CloudWorkspaceWorkspaceIdRepoIdProjectRepoIdProjectId sets the
// optional parameter
// "cloudWorkspace.workspaceId.repoId.projectRepoId.projectId": The ID
// of the project.
func (c *ProjectsReposAliasesListFilesCall) CloudWorkspaceWorkspaceIdRepoIdProjectRepoIdProjectId(cloudWorkspaceWorkspaceIdRepoIdProjectRepoIdProjectId string) *ProjectsReposAliasesListFilesCall {
	c.opt_["cloudWorkspace.workspaceId.repoId.projectRepoId.projectId"] = cloudWorkspaceWorkspaceIdRepoIdProjectRepoIdProjectId
	return c
}

// CloudWorkspaceWorkspaceIdRepoIdProjectRepoIdRepoName sets the
// optional parameter
// "cloudWorkspace.workspaceId.repoId.projectRepoId.repoName": The name
// of the repo. Leave empty for the default repo.
func (c *ProjectsReposAliasesListFilesCall) CloudWorkspaceWorkspaceIdRepoIdProjectRepoIdRepoName(cloudWorkspaceWorkspaceIdRepoIdProjectRepoIdRepoName string) *ProjectsReposAliasesListFilesCall {
	c.opt_["cloudWorkspace.workspaceId.repoId.projectRepoId.repoName"] = cloudWorkspaceWorkspaceIdRepoIdProjectRepoIdRepoName
	return c
}

// CloudWorkspaceWorkspaceIdRepoIdUid sets the optional parameter
// "cloudWorkspace.workspaceId.repoId.uid": A server-assigned, globally
// unique identifier.
func (c *ProjectsReposAliasesListFilesCall) CloudWorkspaceWorkspaceIdRepoIdUid(cloudWorkspaceWorkspaceIdRepoIdUid string) *ProjectsReposAliasesListFilesCall {
	c.opt_["cloudWorkspace.workspaceId.repoId.uid"] = cloudWorkspaceWorkspaceIdRepoIdUid
	return c
}

// GerritAliasContextKind sets the optional parameter
// "gerrit.aliasContext.kind": The alias kind.
func (c *ProjectsReposAliasesListFilesCall) GerritAliasContextKind(gerritAliasContextKind string) *ProjectsReposAliasesListFilesCall {
	c.opt_["gerrit.aliasContext.kind"] = gerritAliasContextKind
	return c
}

// GerritAliasContextName sets the optional parameter
// "gerrit.aliasContext.name": The alias name.
func (c *ProjectsReposAliasesListFilesCall) GerritAliasContextName(gerritAliasContextName string) *ProjectsReposAliasesListFilesCall {
	c.opt_["gerrit.aliasContext.name"] = gerritAliasContextName
	return c
}

// GerritAliasName sets the optional parameter "gerrit.aliasName": The
// name of an alias (branch, tag, etc.).
func (c *ProjectsReposAliasesListFilesCall) GerritAliasName(gerritAliasName string) *ProjectsReposAliasesListFilesCall {
	c.opt_["gerrit.aliasName"] = gerritAliasName
	return c
}

// GerritGerritProject sets the optional parameter
// "gerrit.gerritProject": The full project name within the host.
// Projects may be nested, so
// "project/subproject" is a valid project
// name.
// The "repo name" is hostURI/project.
func (c *ProjectsReposAliasesListFilesCall) GerritGerritProject(gerritGerritProject string) *ProjectsReposAliasesListFilesCall {
	c.opt_["gerrit.gerritProject"] = gerritGerritProject
	return c
}

// GerritHostUri sets the optional parameter "gerrit.hostUri": The URI
// of a running Gerrit instance.
func (c *ProjectsReposAliasesListFilesCall) GerritHostUri(gerritHostUri string) *ProjectsReposAliasesListFilesCall {
	c.opt_["gerrit.hostUri"] = gerritHostUri
	return c
}

// GerritRevisionId sets the optional parameter "gerrit.revisionId": A
// revision (commit) ID.
func (c *ProjectsReposAliasesListFilesCall) GerritRevisionId(gerritRevisionId string) *ProjectsReposAliasesListFilesCall {
	c.opt_["gerrit.revisionId"] = gerritRevisionId
	return c
}

// GitRevisionId sets the optional parameter "git.revisionId": Git
// commit hash.
// required.
func (c *ProjectsReposAliasesListFilesCall) GitRevisionId(gitRevisionId string) *ProjectsReposAliasesListFilesCall {
	c.opt_["git.revisionId"] = gitRevisionId
	return c
}

// GitUrl sets the optional parameter "git.url": Git repository URL.
func (c *ProjectsReposAliasesListFilesCall) GitUrl(gitUrl string) *ProjectsReposAliasesListFilesCall {
	c.opt_["git.url"] = gitUrl
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of values to return.
func (c *ProjectsReposAliasesListFilesCall) PageSize(pageSize int64) *ProjectsReposAliasesListFilesCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken": The value of
// next_page_token from the previous call.
// Omit for the first page.
func (c *ProjectsReposAliasesListFilesCall) PageToken(pageToken string) *ProjectsReposAliasesListFilesCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// RevisionId sets the optional parameter "revisionId": A revision ID.
func (c *ProjectsReposAliasesListFilesCall) RevisionId(revisionId string) *ProjectsReposAliasesListFilesCall {
	c.opt_["revisionId"] = revisionId
	return c
}

// Uid sets the optional parameter "uid": A server-assigned, globally
// unique identifier.
func (c *ProjectsReposAliasesListFilesCall) Uid(uid string) *ProjectsReposAliasesListFilesCall {
	c.opt_["uid"] = uid
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsReposAliasesListFilesCall) Fields(s ...googleapi.Field) *ProjectsReposAliasesListFilesCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsReposAliasesListFilesCall) Do() (*ListFilesResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["aliasName"]; ok {
		params.Set("aliasName", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["cloudWorkspace.snapshotId"]; ok {
		params.Set("cloudWorkspace.snapshotId", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["cloudWorkspace.workspaceId.name"]; ok {
		params.Set("cloudWorkspace.workspaceId.name", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["cloudWorkspace.workspaceId.repoId.projectRepoId.projectId"]; ok {
		params.Set("cloudWorkspace.workspaceId.repoId.projectRepoId.projectId", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["cloudWorkspace.workspaceId.repoId.projectRepoId.repoName"]; ok {
		params.Set("cloudWorkspace.workspaceId.repoId.projectRepoId.repoName", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["cloudWorkspace.workspaceId.repoId.uid"]; ok {
		params.Set("cloudWorkspace.workspaceId.repoId.uid", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["gerrit.aliasContext.kind"]; ok {
		params.Set("gerrit.aliasContext.kind", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["gerrit.aliasContext.name"]; ok {
		params.Set("gerrit.aliasContext.name", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["gerrit.aliasName"]; ok {
		params.Set("gerrit.aliasName", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["gerrit.gerritProject"]; ok {
		params.Set("gerrit.gerritProject", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["gerrit.hostUri"]; ok {
		params.Set("gerrit.hostUri", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["gerrit.revisionId"]; ok {
		params.Set("gerrit.revisionId", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["git.revisionId"]; ok {
		params.Set("git.revisionId", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["git.url"]; ok {
		params.Set("git.url", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageSize"]; ok {
		params.Set("pageSize", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["revisionId"]; ok {
		params.Set("revisionId", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["uid"]; ok {
		params.Set("uid", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/repos/{repoName}/aliases/{kind}/{name}:listFiles")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"repoName":  c.repoName,
		"kind":      c.kind,
		"name":      c.name,
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
	var ret *ListFilesResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "ListFiles returns a list of all files in a SourceContext. The\ninformation about each file includes its path and its hash.\nThe result is ordered by path. Pagination is supported.",
	//   "flatPath": "v1/projects/{projectId}/repos/{repoName}/aliases/{kind}/{name}:listFiles",
	//   "httpMethod": "GET",
	//   "id": "source.projects.repos.aliases.listFiles",
	//   "parameterOrder": [
	//     "projectId",
	//     "repoName",
	//     "kind",
	//     "name"
	//   ],
	//   "parameters": {
	//     "aliasName": {
	//       "description": "The name of an alias (branch, tag, etc.).",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "cloudWorkspace.snapshotId": {
	//       "description": "The ID of the snapshot.\nAn empty snapshot_id refers to the most recent snapshot.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "cloudWorkspace.workspaceId.name": {
	//       "description": "The unique name of the workspace within the repo.  This is the name\nchosen by the client in the Source API's CreateWorkspace method.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "cloudWorkspace.workspaceId.repoId.projectRepoId.projectId": {
	//       "description": "The ID of the project.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "cloudWorkspace.workspaceId.repoId.projectRepoId.repoName": {
	//       "description": "The name of the repo. Leave empty for the default repo.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "cloudWorkspace.workspaceId.repoId.uid": {
	//       "description": "A server-assigned, globally unique identifier.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "gerrit.aliasContext.kind": {
	//       "description": "The alias kind.",
	//       "enum": [
	//         "ANY",
	//         "FIXED",
	//         "MOVABLE",
	//         "OTHER"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "gerrit.aliasContext.name": {
	//       "description": "The alias name.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "gerrit.aliasName": {
	//       "description": "The name of an alias (branch, tag, etc.).",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "gerrit.gerritProject": {
	//       "description": "The full project name within the host. Projects may be nested, so\n\"project/subproject\" is a valid project name.\nThe \"repo name\" is hostURI/project.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "gerrit.hostUri": {
	//       "description": "The URI of a running Gerrit instance.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "gerrit.revisionId": {
	//       "description": "A revision (commit) ID.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "git.revisionId": {
	//       "description": "Git commit hash.\nrequired.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "git.url": {
	//       "description": "Git repository URL.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "kind": {
	//       "description": "The alias kind.",
	//       "enum": [
	//         "ANY",
	//         "FIXED",
	//         "MOVABLE",
	//         "OTHER"
	//       ],
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "name": {
	//       "description": "The alias name.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "The maximum number of values to return.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The value of next_page_token from the previous call.\nOmit for the first page.",
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
	//       "description": "The name of the repo. Leave empty for the default repo.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "revisionId": {
	//       "description": "A revision ID.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "uid": {
	//       "description": "A server-assigned, globally unique identifier.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/repos/{repoName}/aliases/{kind}/{name}:listFiles",
	//   "response": {
	//     "$ref": "ListFilesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "source.projects.repos.aliases.update":

type ProjectsReposAliasesUpdateCall struct {
	s         *Service
	projectId string
	repoName  string
	aliasesId string
	alias     *Alias
	opt_      map[string]interface{}
}

// Update: Updates the alias with the given name and kind. Kind cannot
// be ANY.  If
// the alias does not exist, NOT_FOUND is returned. If the
// request provides
// an old revision ID and the alias does not refer to
// that revision, ABORTED
// is returned.
func (r *ProjectsReposAliasesService) Update(projectId string, repoName string, aliasesId string, alias *Alias) *ProjectsReposAliasesUpdateCall {
	c := &ProjectsReposAliasesUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.repoName = repoName
	c.aliasesId = aliasesId
	c.alias = alias
	return c
}

// OldRevisionId sets the optional parameter "oldRevisionId": If
// non-empty, must match the revision that the alias refers to.
func (c *ProjectsReposAliasesUpdateCall) OldRevisionId(oldRevisionId string) *ProjectsReposAliasesUpdateCall {
	c.opt_["oldRevisionId"] = oldRevisionId
	return c
}

// Uid sets the optional parameter "uid": A server-assigned, globally
// unique identifier.
func (c *ProjectsReposAliasesUpdateCall) Uid(uid string) *ProjectsReposAliasesUpdateCall {
	c.opt_["uid"] = uid
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsReposAliasesUpdateCall) Fields(s ...googleapi.Field) *ProjectsReposAliasesUpdateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsReposAliasesUpdateCall) Do() (*Alias, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.alias)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["oldRevisionId"]; ok {
		params.Set("oldRevisionId", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["uid"]; ok {
		params.Set("uid", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/repos/{repoName}/aliases/{aliasesId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"repoName":  c.repoName,
		"aliasesId": c.aliasesId,
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
	var ret *Alias
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates the alias with the given name and kind. Kind cannot be ANY.  If\nthe alias does not exist, NOT_FOUND is returned. If the request provides\nan old revision ID and the alias does not refer to that revision, ABORTED\nis returned.",
	//   "flatPath": "v1/projects/{projectId}/repos/{repoName}/aliases/{aliasesId}",
	//   "httpMethod": "PUT",
	//   "id": "source.projects.repos.aliases.update",
	//   "parameterOrder": [
	//     "projectId",
	//     "repoName",
	//     "aliasesId"
	//   ],
	//   "parameters": {
	//     "aliasesId": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "oldRevisionId": {
	//       "description": "If non-empty, must match the revision that the alias refers to.",
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
	//       "description": "The name of the repo. Leave empty for the default repo.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "uid": {
	//       "description": "A server-assigned, globally unique identifier.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/repos/{repoName}/aliases/{aliasesId}",
	//   "request": {
	//     "$ref": "Alias"
	//   },
	//   "response": {
	//     "$ref": "Alias"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "source.projects.repos.aliases.files.get":

type ProjectsReposAliasesFilesGetCall struct {
	s         *Service
	projectId string
	repoName  string
	kind      string
	name      string
	path      string
	opt_      map[string]interface{}
}

// Get: Read is given a SourceContext and path, and returns
// file or
// directory information about that path.
func (r *ProjectsReposAliasesFilesService) Get(projectId string, repoName string, kind string, name string, path string) *ProjectsReposAliasesFilesGetCall {
	c := &ProjectsReposAliasesFilesGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.repoName = repoName
	c.kind = kind
	c.name = name
	c.path = path
	return c
}

// AliasName sets the optional parameter "aliasName": The name of an
// alias (branch, tag, etc.).
func (c *ProjectsReposAliasesFilesGetCall) AliasName(aliasName string) *ProjectsReposAliasesFilesGetCall {
	c.opt_["aliasName"] = aliasName
	return c
}

// CloudWorkspaceSnapshotId sets the optional parameter
// "cloudWorkspace.snapshotId": The ID of the snapshot.
// An empty
// snapshot_id refers to the most recent snapshot.
func (c *ProjectsReposAliasesFilesGetCall) CloudWorkspaceSnapshotId(cloudWorkspaceSnapshotId string) *ProjectsReposAliasesFilesGetCall {
	c.opt_["cloudWorkspace.snapshotId"] = cloudWorkspaceSnapshotId
	return c
}

// CloudWorkspaceWorkspaceIdName sets the optional parameter
// "cloudWorkspace.workspaceId.name": The unique name of the workspace
// within the repo.  This is the name
// chosen by the client in the Source
// API's CreateWorkspace method.
func (c *ProjectsReposAliasesFilesGetCall) CloudWorkspaceWorkspaceIdName(cloudWorkspaceWorkspaceIdName string) *ProjectsReposAliasesFilesGetCall {
	c.opt_["cloudWorkspace.workspaceId.name"] = cloudWorkspaceWorkspaceIdName
	return c
}

// CloudWorkspaceWorkspaceIdRepoIdProjectRepoIdProjectId sets the
// optional parameter
// "cloudWorkspace.workspaceId.repoId.projectRepoId.projectId": The ID
// of the project.
func (c *ProjectsReposAliasesFilesGetCall) CloudWorkspaceWorkspaceIdRepoIdProjectRepoIdProjectId(cloudWorkspaceWorkspaceIdRepoIdProjectRepoIdProjectId string) *ProjectsReposAliasesFilesGetCall {
	c.opt_["cloudWorkspace.workspaceId.repoId.projectRepoId.projectId"] = cloudWorkspaceWorkspaceIdRepoIdProjectRepoIdProjectId
	return c
}

// CloudWorkspaceWorkspaceIdRepoIdProjectRepoIdRepoName sets the
// optional parameter
// "cloudWorkspace.workspaceId.repoId.projectRepoId.repoName": The name
// of the repo. Leave empty for the default repo.
func (c *ProjectsReposAliasesFilesGetCall) CloudWorkspaceWorkspaceIdRepoIdProjectRepoIdRepoName(cloudWorkspaceWorkspaceIdRepoIdProjectRepoIdRepoName string) *ProjectsReposAliasesFilesGetCall {
	c.opt_["cloudWorkspace.workspaceId.repoId.projectRepoId.repoName"] = cloudWorkspaceWorkspaceIdRepoIdProjectRepoIdRepoName
	return c
}

// CloudWorkspaceWorkspaceIdRepoIdUid sets the optional parameter
// "cloudWorkspace.workspaceId.repoId.uid": A server-assigned, globally
// unique identifier.
func (c *ProjectsReposAliasesFilesGetCall) CloudWorkspaceWorkspaceIdRepoIdUid(cloudWorkspaceWorkspaceIdRepoIdUid string) *ProjectsReposAliasesFilesGetCall {
	c.opt_["cloudWorkspace.workspaceId.repoId.uid"] = cloudWorkspaceWorkspaceIdRepoIdUid
	return c
}

// GerritAliasContextKind sets the optional parameter
// "gerrit.aliasContext.kind": The alias kind.
func (c *ProjectsReposAliasesFilesGetCall) GerritAliasContextKind(gerritAliasContextKind string) *ProjectsReposAliasesFilesGetCall {
	c.opt_["gerrit.aliasContext.kind"] = gerritAliasContextKind
	return c
}

// GerritAliasContextName sets the optional parameter
// "gerrit.aliasContext.name": The alias name.
func (c *ProjectsReposAliasesFilesGetCall) GerritAliasContextName(gerritAliasContextName string) *ProjectsReposAliasesFilesGetCall {
	c.opt_["gerrit.aliasContext.name"] = gerritAliasContextName
	return c
}

// GerritAliasName sets the optional parameter "gerrit.aliasName": The
// name of an alias (branch, tag, etc.).
func (c *ProjectsReposAliasesFilesGetCall) GerritAliasName(gerritAliasName string) *ProjectsReposAliasesFilesGetCall {
	c.opt_["gerrit.aliasName"] = gerritAliasName
	return c
}

// GerritGerritProject sets the optional parameter
// "gerrit.gerritProject": The full project name within the host.
// Projects may be nested, so
// "project/subproject" is a valid project
// name.
// The "repo name" is hostURI/project.
func (c *ProjectsReposAliasesFilesGetCall) GerritGerritProject(gerritGerritProject string) *ProjectsReposAliasesFilesGetCall {
	c.opt_["gerrit.gerritProject"] = gerritGerritProject
	return c
}

// GerritHostUri sets the optional parameter "gerrit.hostUri": The URI
// of a running Gerrit instance.
func (c *ProjectsReposAliasesFilesGetCall) GerritHostUri(gerritHostUri string) *ProjectsReposAliasesFilesGetCall {
	c.opt_["gerrit.hostUri"] = gerritHostUri
	return c
}

// GerritRevisionId sets the optional parameter "gerrit.revisionId": A
// revision (commit) ID.
func (c *ProjectsReposAliasesFilesGetCall) GerritRevisionId(gerritRevisionId string) *ProjectsReposAliasesFilesGetCall {
	c.opt_["gerrit.revisionId"] = gerritRevisionId
	return c
}

// GitRevisionId sets the optional parameter "git.revisionId": Git
// commit hash.
// required.
func (c *ProjectsReposAliasesFilesGetCall) GitRevisionId(gitRevisionId string) *ProjectsReposAliasesFilesGetCall {
	c.opt_["git.revisionId"] = gitRevisionId
	return c
}

// GitUrl sets the optional parameter "git.url": Git repository URL.
func (c *ProjectsReposAliasesFilesGetCall) GitUrl(gitUrl string) *ProjectsReposAliasesFilesGetCall {
	c.opt_["git.url"] = gitUrl
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of values to return.
func (c *ProjectsReposAliasesFilesGetCall) PageSize(pageSize int64) *ProjectsReposAliasesFilesGetCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken": The value of
// next_page_token from the previous call.
// Omit for the first page, or
// if using start_index.
func (c *ProjectsReposAliasesFilesGetCall) PageToken(pageToken string) *ProjectsReposAliasesFilesGetCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// RevisionId sets the optional parameter "revisionId": A revision ID.
func (c *ProjectsReposAliasesFilesGetCall) RevisionId(revisionId string) *ProjectsReposAliasesFilesGetCall {
	c.opt_["revisionId"] = revisionId
	return c
}

// StartPosition sets the optional parameter "startPosition": If path
// refers to a file, the position of the first byte of its contents
// to
// return. If path refers to a directory, the position of the first
// entry
// in the listing. If page_token is specified, this field is
// ignored.
func (c *ProjectsReposAliasesFilesGetCall) StartPosition(startPosition int64) *ProjectsReposAliasesFilesGetCall {
	c.opt_["startPosition"] = startPosition
	return c
}

// Uid sets the optional parameter "uid": A server-assigned, globally
// unique identifier.
func (c *ProjectsReposAliasesFilesGetCall) Uid(uid string) *ProjectsReposAliasesFilesGetCall {
	c.opt_["uid"] = uid
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsReposAliasesFilesGetCall) Fields(s ...googleapi.Field) *ProjectsReposAliasesFilesGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsReposAliasesFilesGetCall) Do() (*ReadResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["aliasName"]; ok {
		params.Set("aliasName", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["cloudWorkspace.snapshotId"]; ok {
		params.Set("cloudWorkspace.snapshotId", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["cloudWorkspace.workspaceId.name"]; ok {
		params.Set("cloudWorkspace.workspaceId.name", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["cloudWorkspace.workspaceId.repoId.projectRepoId.projectId"]; ok {
		params.Set("cloudWorkspace.workspaceId.repoId.projectRepoId.projectId", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["cloudWorkspace.workspaceId.repoId.projectRepoId.repoName"]; ok {
		params.Set("cloudWorkspace.workspaceId.repoId.projectRepoId.repoName", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["cloudWorkspace.workspaceId.repoId.uid"]; ok {
		params.Set("cloudWorkspace.workspaceId.repoId.uid", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["gerrit.aliasContext.kind"]; ok {
		params.Set("gerrit.aliasContext.kind", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["gerrit.aliasContext.name"]; ok {
		params.Set("gerrit.aliasContext.name", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["gerrit.aliasName"]; ok {
		params.Set("gerrit.aliasName", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["gerrit.gerritProject"]; ok {
		params.Set("gerrit.gerritProject", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["gerrit.hostUri"]; ok {
		params.Set("gerrit.hostUri", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["gerrit.revisionId"]; ok {
		params.Set("gerrit.revisionId", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["git.revisionId"]; ok {
		params.Set("git.revisionId", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["git.url"]; ok {
		params.Set("git.url", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageSize"]; ok {
		params.Set("pageSize", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["revisionId"]; ok {
		params.Set("revisionId", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["startPosition"]; ok {
		params.Set("startPosition", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["uid"]; ok {
		params.Set("uid", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/repos/{repoName}/aliases/{kind}/{name}/files/{+path}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"repoName":  c.repoName,
		"kind":      c.kind,
		"name":      c.name,
		"path":      c.path,
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
	var ret *ReadResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Read is given a SourceContext and path, and returns\nfile or directory information about that path.",
	//   "flatPath": "v1/projects/{projectId}/repos/{repoName}/aliases/{kind}/{name}/files/{filesId}",
	//   "httpMethod": "GET",
	//   "id": "source.projects.repos.aliases.files.get",
	//   "parameterOrder": [
	//     "projectId",
	//     "repoName",
	//     "kind",
	//     "name",
	//     "path"
	//   ],
	//   "parameters": {
	//     "aliasName": {
	//       "description": "The name of an alias (branch, tag, etc.).",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "cloudWorkspace.snapshotId": {
	//       "description": "The ID of the snapshot.\nAn empty snapshot_id refers to the most recent snapshot.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "cloudWorkspace.workspaceId.name": {
	//       "description": "The unique name of the workspace within the repo.  This is the name\nchosen by the client in the Source API's CreateWorkspace method.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "cloudWorkspace.workspaceId.repoId.projectRepoId.projectId": {
	//       "description": "The ID of the project.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "cloudWorkspace.workspaceId.repoId.projectRepoId.repoName": {
	//       "description": "The name of the repo. Leave empty for the default repo.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "cloudWorkspace.workspaceId.repoId.uid": {
	//       "description": "A server-assigned, globally unique identifier.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "gerrit.aliasContext.kind": {
	//       "description": "The alias kind.",
	//       "enum": [
	//         "ANY",
	//         "FIXED",
	//         "MOVABLE",
	//         "OTHER"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "gerrit.aliasContext.name": {
	//       "description": "The alias name.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "gerrit.aliasName": {
	//       "description": "The name of an alias (branch, tag, etc.).",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "gerrit.gerritProject": {
	//       "description": "The full project name within the host. Projects may be nested, so\n\"project/subproject\" is a valid project name.\nThe \"repo name\" is hostURI/project.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "gerrit.hostUri": {
	//       "description": "The URI of a running Gerrit instance.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "gerrit.revisionId": {
	//       "description": "A revision (commit) ID.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "git.revisionId": {
	//       "description": "Git commit hash.\nrequired.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "git.url": {
	//       "description": "Git repository URL.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "kind": {
	//       "description": "The alias kind.",
	//       "enum": [
	//         "ANY",
	//         "FIXED",
	//         "MOVABLE",
	//         "OTHER"
	//       ],
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "name": {
	//       "description": "The alias name.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "The maximum number of values to return.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "The value of next_page_token from the previous call.\nOmit for the first page, or if using start_index.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "path": {
	//       "description": "Path to the file or directory from the root directory of the source\ncontext. It must not have leading or trailing slashes.",
	//       "location": "path",
	//       "pattern": "^.*$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "The ID of the project.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "repoName": {
	//       "description": "The name of the repo. Leave empty for the default repo.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "revisionId": {
	//       "description": "A revision ID.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "startPosition": {
	//       "description": "If path refers to a file, the position of the first byte of its contents\nto return. If path refers to a directory, the position of the first entry\nin the listing. If page_token is specified, this field is ignored.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "uid": {
	//       "description": "A server-assigned, globally unique identifier.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/repos/{repoName}/aliases/{kind}/{name}/files/{+path}",
	//   "response": {
	//     "$ref": "ReadResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "source.projects.repos.files.readFromWorkspaceOrAlias":

type ProjectsReposFilesReadFromWorkspaceOrAliasCall struct {
	s         *Service
	projectId string
	repoName  string
	path      string
	opt_      map[string]interface{}
}

// ReadFromWorkspaceOrAlias: ReadFromWorkspaceOrAlias performs a Read
// using either the most recent
// snapshot of the given workspace, if the
// workspace exists, or the
// revision referred to by the given alias if
// the workspace does not exist.
func (r *ProjectsReposFilesService) ReadFromWorkspaceOrAlias(projectId string, repoName string, path string) *ProjectsReposFilesReadFromWorkspaceOrAliasCall {
	c := &ProjectsReposFilesReadFromWorkspaceOrAliasCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.repoName = repoName
	c.path = path
	return c
}

// Alias sets the optional parameter "alias": MOVABLE alias to read
// from, if the workspace doesn't exist.
func (c *ProjectsReposFilesReadFromWorkspaceOrAliasCall) Alias(alias string) *ProjectsReposFilesReadFromWorkspaceOrAliasCall {
	c.opt_["alias"] = alias
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of values to return.
func (c *ProjectsReposFilesReadFromWorkspaceOrAliasCall) PageSize(pageSize int64) *ProjectsReposFilesReadFromWorkspaceOrAliasCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken": The value of
// next_page_token from the previous call.
// Omit for the first page.
func (c *ProjectsReposFilesReadFromWorkspaceOrAliasCall) PageToken(pageToken string) *ProjectsReposFilesReadFromWorkspaceOrAliasCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// StartPosition sets the optional parameter "startPosition": If path
// refers to a file, the position of the first byte of its contents
// to
// return. If path refers to a directory, the position of the first
// entry
// in the listing. If page_token is specified, this field is
// ignored.
func (c *ProjectsReposFilesReadFromWorkspaceOrAliasCall) StartPosition(startPosition int64) *ProjectsReposFilesReadFromWorkspaceOrAliasCall {
	c.opt_["startPosition"] = startPosition
	return c
}

// Uid sets the optional parameter "uid": A server-assigned, globally
// unique identifier.
func (c *ProjectsReposFilesReadFromWorkspaceOrAliasCall) Uid(uid string) *ProjectsReposFilesReadFromWorkspaceOrAliasCall {
	c.opt_["uid"] = uid
	return c
}

// WorkspaceName sets the optional parameter "workspaceName": Workspace
// to read from, if it exists.
func (c *ProjectsReposFilesReadFromWorkspaceOrAliasCall) WorkspaceName(workspaceName string) *ProjectsReposFilesReadFromWorkspaceOrAliasCall {
	c.opt_["workspaceName"] = workspaceName
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsReposFilesReadFromWorkspaceOrAliasCall) Fields(s ...googleapi.Field) *ProjectsReposFilesReadFromWorkspaceOrAliasCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsReposFilesReadFromWorkspaceOrAliasCall) Do() (*ReadResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["alias"]; ok {
		params.Set("alias", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageSize"]; ok {
		params.Set("pageSize", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["startPosition"]; ok {
		params.Set("startPosition", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["uid"]; ok {
		params.Set("uid", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["workspaceName"]; ok {
		params.Set("workspaceName", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/repos/{repoName}/files/{+path}:readFromWorkspaceOrAlias")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"repoName":  c.repoName,
		"path":      c.path,
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
	var ret *ReadResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "ReadFromWorkspaceOrAlias performs a Read using either the most recent\nsnapshot of the given workspace, if the workspace exists, or the\nrevision referred to by the given alias if the workspace does not exist.",
	//   "flatPath": "v1/projects/{projectId}/repos/{repoName}/files/{filesId}:readFromWorkspaceOrAlias",
	//   "httpMethod": "GET",
	//   "id": "source.projects.repos.files.readFromWorkspaceOrAlias",
	//   "parameterOrder": [
	//     "projectId",
	//     "repoName",
	//     "path"
	//   ],
	//   "parameters": {
	//     "alias": {
	//       "description": "MOVABLE alias to read from, if the workspace doesn't exist.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "The maximum number of values to return.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "The value of next_page_token from the previous call.\nOmit for the first page.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "path": {
	//       "description": "Path to the file or directory from the root directory of the source\ncontext. It must not have leading or trailing slashes.",
	//       "location": "path",
	//       "pattern": "^.*$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "The ID of the project.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "repoName": {
	//       "description": "The name of the repo. Leave empty for the default repo.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "startPosition": {
	//       "description": "If path refers to a file, the position of the first byte of its contents\nto return. If path refers to a directory, the position of the first entry\nin the listing. If page_token is specified, this field is ignored.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "uid": {
	//       "description": "A server-assigned, globally unique identifier.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "workspaceName": {
	//       "description": "Workspace to read from, if it exists.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/repos/{repoName}/files/{+path}:readFromWorkspaceOrAlias",
	//   "response": {
	//     "$ref": "ReadResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "source.projects.repos.revisions.get":

type ProjectsReposRevisionsGetCall struct {
	s          *Service
	projectId  string
	repoName   string
	revisionId string
	opt_       map[string]interface{}
}

// Get: Retrieves revision metadata for a single revision.
func (r *ProjectsReposRevisionsService) Get(projectId string, repoName string, revisionId string) *ProjectsReposRevisionsGetCall {
	c := &ProjectsReposRevisionsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.repoName = repoName
	c.revisionId = revisionId
	return c
}

// Uid sets the optional parameter "uid": A server-assigned, globally
// unique identifier.
func (c *ProjectsReposRevisionsGetCall) Uid(uid string) *ProjectsReposRevisionsGetCall {
	c.opt_["uid"] = uid
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsReposRevisionsGetCall) Fields(s ...googleapi.Field) *ProjectsReposRevisionsGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsReposRevisionsGetCall) Do() (*Revision, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["uid"]; ok {
		params.Set("uid", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/repos/{repoName}/revisions/{revisionId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId":  c.projectId,
		"repoName":   c.repoName,
		"revisionId": c.revisionId,
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
	var ret *Revision
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves revision metadata for a single revision.",
	//   "flatPath": "v1/projects/{projectId}/repos/{repoName}/revisions/{revisionId}",
	//   "httpMethod": "GET",
	//   "id": "source.projects.repos.revisions.get",
	//   "parameterOrder": [
	//     "projectId",
	//     "repoName",
	//     "revisionId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "The ID of the project.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "repoName": {
	//       "description": "The name of the repo. Leave empty for the default repo.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "revisionId": {
	//       "description": "The ID of the revision.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "uid": {
	//       "description": "A server-assigned, globally unique identifier.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/repos/{repoName}/revisions/{revisionId}",
	//   "response": {
	//     "$ref": "Revision"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "source.projects.repos.revisions.getBatchGet":

type ProjectsReposRevisionsGetBatchGetCall struct {
	s         *Service
	projectId string
	repoName  string
	opt_      map[string]interface{}
}

// GetBatchGet: Retrieves revision metadata for several revisions at
// once. It returns an
// error if any retrieval fails.
func (r *ProjectsReposRevisionsService) GetBatchGet(projectId string, repoName string) *ProjectsReposRevisionsGetBatchGetCall {
	c := &ProjectsReposRevisionsGetBatchGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.repoName = repoName
	return c
}

// RevisionIds sets the optional parameter "revisionIds": The revision
// IDs to retrieve.
func (c *ProjectsReposRevisionsGetBatchGetCall) RevisionIds(revisionIds string) *ProjectsReposRevisionsGetBatchGetCall {
	c.opt_["revisionIds"] = revisionIds
	return c
}

// Uid sets the optional parameter "uid": A server-assigned, globally
// unique identifier.
func (c *ProjectsReposRevisionsGetBatchGetCall) Uid(uid string) *ProjectsReposRevisionsGetBatchGetCall {
	c.opt_["uid"] = uid
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsReposRevisionsGetBatchGetCall) Fields(s ...googleapi.Field) *ProjectsReposRevisionsGetBatchGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsReposRevisionsGetBatchGetCall) Do() (*GetRevisionsResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["revisionIds"]; ok {
		params.Set("revisionIds", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["uid"]; ok {
		params.Set("uid", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/repos/{repoName}/revisions:batchGet")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"repoName":  c.repoName,
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
	var ret *GetRevisionsResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves revision metadata for several revisions at once. It returns an\nerror if any retrieval fails.",
	//   "flatPath": "v1/projects/{projectId}/repos/{repoName}/revisions:batchGet",
	//   "httpMethod": "GET",
	//   "id": "source.projects.repos.revisions.getBatchGet",
	//   "parameterOrder": [
	//     "projectId",
	//     "repoName"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "The ID of the project.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "repoName": {
	//       "description": "The name of the repo. Leave empty for the default repo.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "revisionIds": {
	//       "description": "The revision IDs to retrieve.",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "uid": {
	//       "description": "A server-assigned, globally unique identifier.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/repos/{repoName}/revisions:batchGet",
	//   "response": {
	//     "$ref": "GetRevisionsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "source.projects.repos.revisions.list":

type ProjectsReposRevisionsListCall struct {
	s         *Service
	projectId string
	repoName  string
	opt_      map[string]interface{}
}

// List: Retrieves all revisions topologically between the starts and
// ends.
// Uses the commit date to break ties in the topology (e.g. when a
// revision
// has two parents).
func (r *ProjectsReposRevisionsService) List(projectId string, repoName string) *ProjectsReposRevisionsListCall {
	c := &ProjectsReposRevisionsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.repoName = repoName
	return c
}

// Ends sets the optional parameter "ends": Revision IDs (hexadecimal
// strings) that specify where the listing ends. If
// this field is
// present, the listing will contain only revisions that
// are
// topologically between starts and ends, inclusive.
func (c *ProjectsReposRevisionsListCall) Ends(ends string) *ProjectsReposRevisionsListCall {
	c.opt_["ends"] = ends
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of values to return.
func (c *ProjectsReposRevisionsListCall) PageSize(pageSize int64) *ProjectsReposRevisionsListCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken": The value of
// next_page_token from the previous call.
// Omit for the first page.
func (c *ProjectsReposRevisionsListCall) PageToken(pageToken string) *ProjectsReposRevisionsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Path sets the optional parameter "path": List only those revisions
// that modify path.
func (c *ProjectsReposRevisionsListCall) Path(path string) *ProjectsReposRevisionsListCall {
	c.opt_["path"] = path
	return c
}

// Starts sets the optional parameter "starts": Revision IDs
// (hexadecimal strings) that specify where the listing
// begins. If
// empty, the repo heads (revisions with no children) are used.
func (c *ProjectsReposRevisionsListCall) Starts(starts string) *ProjectsReposRevisionsListCall {
	c.opt_["starts"] = starts
	return c
}

// Uid sets the optional parameter "uid": A server-assigned, globally
// unique identifier.
func (c *ProjectsReposRevisionsListCall) Uid(uid string) *ProjectsReposRevisionsListCall {
	c.opt_["uid"] = uid
	return c
}

// WalkDirection sets the optional parameter "walkDirection": The
// direction to walk the graph.
func (c *ProjectsReposRevisionsListCall) WalkDirection(walkDirection string) *ProjectsReposRevisionsListCall {
	c.opt_["walkDirection"] = walkDirection
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsReposRevisionsListCall) Fields(s ...googleapi.Field) *ProjectsReposRevisionsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsReposRevisionsListCall) Do() (*ListRevisionsResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["ends"]; ok {
		params.Set("ends", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageSize"]; ok {
		params.Set("pageSize", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["path"]; ok {
		params.Set("path", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["starts"]; ok {
		params.Set("starts", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["uid"]; ok {
		params.Set("uid", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["walkDirection"]; ok {
		params.Set("walkDirection", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/repos/{repoName}/revisions")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"repoName":  c.repoName,
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
	var ret *ListRevisionsResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves all revisions topologically between the starts and ends.\nUses the commit date to break ties in the topology (e.g. when a revision\nhas two parents).",
	//   "flatPath": "v1/projects/{projectId}/repos/{repoName}/revisions",
	//   "httpMethod": "GET",
	//   "id": "source.projects.repos.revisions.list",
	//   "parameterOrder": [
	//     "projectId",
	//     "repoName"
	//   ],
	//   "parameters": {
	//     "ends": {
	//       "description": "Revision IDs (hexadecimal strings) that specify where the listing ends. If\nthis field is present, the listing will contain only revisions that are\ntopologically between starts and ends, inclusive.",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "The maximum number of values to return.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The value of next_page_token from the previous call.\nOmit for the first page.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "path": {
	//       "description": "List only those revisions that modify path.",
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
	//       "description": "The name of the repo. Leave empty for the default repo.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "starts": {
	//       "description": "Revision IDs (hexadecimal strings) that specify where the listing\nbegins. If empty, the repo heads (revisions with no children) are used.",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "uid": {
	//       "description": "A server-assigned, globally unique identifier.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "walkDirection": {
	//       "description": "The direction to walk the graph.",
	//       "enum": [
	//         "BACKWARD",
	//         "FORWARD"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/repos/{repoName}/revisions",
	//   "response": {
	//     "$ref": "ListRevisionsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "source.projects.repos.revisions.listFiles":

type ProjectsReposRevisionsListFilesCall struct {
	s          *Service
	projectId  string
	repoName   string
	revisionId string
	opt_       map[string]interface{}
}

// ListFiles: ListFiles returns a list of all files in a SourceContext.
// The
// information about each file includes its path and its hash.
// The
// result is ordered by path. Pagination is supported.
func (r *ProjectsReposRevisionsService) ListFiles(projectId string, repoName string, revisionId string) *ProjectsReposRevisionsListFilesCall {
	c := &ProjectsReposRevisionsListFilesCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.repoName = repoName
	c.revisionId = revisionId
	return c
}

// AliasContextKind sets the optional parameter "aliasContext.kind": The
// alias kind.
func (c *ProjectsReposRevisionsListFilesCall) AliasContextKind(aliasContextKind string) *ProjectsReposRevisionsListFilesCall {
	c.opt_["aliasContext.kind"] = aliasContextKind
	return c
}

// AliasContextName sets the optional parameter "aliasContext.name": The
// alias name.
func (c *ProjectsReposRevisionsListFilesCall) AliasContextName(aliasContextName string) *ProjectsReposRevisionsListFilesCall {
	c.opt_["aliasContext.name"] = aliasContextName
	return c
}

// AliasName sets the optional parameter "aliasName": The name of an
// alias (branch, tag, etc.).
func (c *ProjectsReposRevisionsListFilesCall) AliasName(aliasName string) *ProjectsReposRevisionsListFilesCall {
	c.opt_["aliasName"] = aliasName
	return c
}

// CloudWorkspaceSnapshotId sets the optional parameter
// "cloudWorkspace.snapshotId": The ID of the snapshot.
// An empty
// snapshot_id refers to the most recent snapshot.
func (c *ProjectsReposRevisionsListFilesCall) CloudWorkspaceSnapshotId(cloudWorkspaceSnapshotId string) *ProjectsReposRevisionsListFilesCall {
	c.opt_["cloudWorkspace.snapshotId"] = cloudWorkspaceSnapshotId
	return c
}

// CloudWorkspaceWorkspaceIdName sets the optional parameter
// "cloudWorkspace.workspaceId.name": The unique name of the workspace
// within the repo.  This is the name
// chosen by the client in the Source
// API's CreateWorkspace method.
func (c *ProjectsReposRevisionsListFilesCall) CloudWorkspaceWorkspaceIdName(cloudWorkspaceWorkspaceIdName string) *ProjectsReposRevisionsListFilesCall {
	c.opt_["cloudWorkspace.workspaceId.name"] = cloudWorkspaceWorkspaceIdName
	return c
}

// CloudWorkspaceWorkspaceIdRepoIdProjectRepoIdProjectId sets the
// optional parameter
// "cloudWorkspace.workspaceId.repoId.projectRepoId.projectId": The ID
// of the project.
func (c *ProjectsReposRevisionsListFilesCall) CloudWorkspaceWorkspaceIdRepoIdProjectRepoIdProjectId(cloudWorkspaceWorkspaceIdRepoIdProjectRepoIdProjectId string) *ProjectsReposRevisionsListFilesCall {
	c.opt_["cloudWorkspace.workspaceId.repoId.projectRepoId.projectId"] = cloudWorkspaceWorkspaceIdRepoIdProjectRepoIdProjectId
	return c
}

// CloudWorkspaceWorkspaceIdRepoIdProjectRepoIdRepoName sets the
// optional parameter
// "cloudWorkspace.workspaceId.repoId.projectRepoId.repoName": The name
// of the repo. Leave empty for the default repo.
func (c *ProjectsReposRevisionsListFilesCall) CloudWorkspaceWorkspaceIdRepoIdProjectRepoIdRepoName(cloudWorkspaceWorkspaceIdRepoIdProjectRepoIdRepoName string) *ProjectsReposRevisionsListFilesCall {
	c.opt_["cloudWorkspace.workspaceId.repoId.projectRepoId.repoName"] = cloudWorkspaceWorkspaceIdRepoIdProjectRepoIdRepoName
	return c
}

// CloudWorkspaceWorkspaceIdRepoIdUid sets the optional parameter
// "cloudWorkspace.workspaceId.repoId.uid": A server-assigned, globally
// unique identifier.
func (c *ProjectsReposRevisionsListFilesCall) CloudWorkspaceWorkspaceIdRepoIdUid(cloudWorkspaceWorkspaceIdRepoIdUid string) *ProjectsReposRevisionsListFilesCall {
	c.opt_["cloudWorkspace.workspaceId.repoId.uid"] = cloudWorkspaceWorkspaceIdRepoIdUid
	return c
}

// GerritAliasContextKind sets the optional parameter
// "gerrit.aliasContext.kind": The alias kind.
func (c *ProjectsReposRevisionsListFilesCall) GerritAliasContextKind(gerritAliasContextKind string) *ProjectsReposRevisionsListFilesCall {
	c.opt_["gerrit.aliasContext.kind"] = gerritAliasContextKind
	return c
}

// GerritAliasContextName sets the optional parameter
// "gerrit.aliasContext.name": The alias name.
func (c *ProjectsReposRevisionsListFilesCall) GerritAliasContextName(gerritAliasContextName string) *ProjectsReposRevisionsListFilesCall {
	c.opt_["gerrit.aliasContext.name"] = gerritAliasContextName
	return c
}

// GerritAliasName sets the optional parameter "gerrit.aliasName": The
// name of an alias (branch, tag, etc.).
func (c *ProjectsReposRevisionsListFilesCall) GerritAliasName(gerritAliasName string) *ProjectsReposRevisionsListFilesCall {
	c.opt_["gerrit.aliasName"] = gerritAliasName
	return c
}

// GerritGerritProject sets the optional parameter
// "gerrit.gerritProject": The full project name within the host.
// Projects may be nested, so
// "project/subproject" is a valid project
// name.
// The "repo name" is hostURI/project.
func (c *ProjectsReposRevisionsListFilesCall) GerritGerritProject(gerritGerritProject string) *ProjectsReposRevisionsListFilesCall {
	c.opt_["gerrit.gerritProject"] = gerritGerritProject
	return c
}

// GerritHostUri sets the optional parameter "gerrit.hostUri": The URI
// of a running Gerrit instance.
func (c *ProjectsReposRevisionsListFilesCall) GerritHostUri(gerritHostUri string) *ProjectsReposRevisionsListFilesCall {
	c.opt_["gerrit.hostUri"] = gerritHostUri
	return c
}

// GerritRevisionId sets the optional parameter "gerrit.revisionId": A
// revision (commit) ID.
func (c *ProjectsReposRevisionsListFilesCall) GerritRevisionId(gerritRevisionId string) *ProjectsReposRevisionsListFilesCall {
	c.opt_["gerrit.revisionId"] = gerritRevisionId
	return c
}

// GitRevisionId sets the optional parameter "git.revisionId": Git
// commit hash.
// required.
func (c *ProjectsReposRevisionsListFilesCall) GitRevisionId(gitRevisionId string) *ProjectsReposRevisionsListFilesCall {
	c.opt_["git.revisionId"] = gitRevisionId
	return c
}

// GitUrl sets the optional parameter "git.url": Git repository URL.
func (c *ProjectsReposRevisionsListFilesCall) GitUrl(gitUrl string) *ProjectsReposRevisionsListFilesCall {
	c.opt_["git.url"] = gitUrl
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of values to return.
func (c *ProjectsReposRevisionsListFilesCall) PageSize(pageSize int64) *ProjectsReposRevisionsListFilesCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken": The value of
// next_page_token from the previous call.
// Omit for the first page.
func (c *ProjectsReposRevisionsListFilesCall) PageToken(pageToken string) *ProjectsReposRevisionsListFilesCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Uid sets the optional parameter "uid": A server-assigned, globally
// unique identifier.
func (c *ProjectsReposRevisionsListFilesCall) Uid(uid string) *ProjectsReposRevisionsListFilesCall {
	c.opt_["uid"] = uid
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsReposRevisionsListFilesCall) Fields(s ...googleapi.Field) *ProjectsReposRevisionsListFilesCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsReposRevisionsListFilesCall) Do() (*ListFilesResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["aliasContext.kind"]; ok {
		params.Set("aliasContext.kind", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["aliasContext.name"]; ok {
		params.Set("aliasContext.name", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["aliasName"]; ok {
		params.Set("aliasName", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["cloudWorkspace.snapshotId"]; ok {
		params.Set("cloudWorkspace.snapshotId", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["cloudWorkspace.workspaceId.name"]; ok {
		params.Set("cloudWorkspace.workspaceId.name", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["cloudWorkspace.workspaceId.repoId.projectRepoId.projectId"]; ok {
		params.Set("cloudWorkspace.workspaceId.repoId.projectRepoId.projectId", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["cloudWorkspace.workspaceId.repoId.projectRepoId.repoName"]; ok {
		params.Set("cloudWorkspace.workspaceId.repoId.projectRepoId.repoName", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["cloudWorkspace.workspaceId.repoId.uid"]; ok {
		params.Set("cloudWorkspace.workspaceId.repoId.uid", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["gerrit.aliasContext.kind"]; ok {
		params.Set("gerrit.aliasContext.kind", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["gerrit.aliasContext.name"]; ok {
		params.Set("gerrit.aliasContext.name", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["gerrit.aliasName"]; ok {
		params.Set("gerrit.aliasName", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["gerrit.gerritProject"]; ok {
		params.Set("gerrit.gerritProject", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["gerrit.hostUri"]; ok {
		params.Set("gerrit.hostUri", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["gerrit.revisionId"]; ok {
		params.Set("gerrit.revisionId", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["git.revisionId"]; ok {
		params.Set("git.revisionId", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["git.url"]; ok {
		params.Set("git.url", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageSize"]; ok {
		params.Set("pageSize", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["uid"]; ok {
		params.Set("uid", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/repos/{repoName}/revisions/{revisionId}:listFiles")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId":  c.projectId,
		"repoName":   c.repoName,
		"revisionId": c.revisionId,
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
	var ret *ListFilesResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "ListFiles returns a list of all files in a SourceContext. The\ninformation about each file includes its path and its hash.\nThe result is ordered by path. Pagination is supported.",
	//   "flatPath": "v1/projects/{projectId}/repos/{repoName}/revisions/{revisionId}:listFiles",
	//   "httpMethod": "GET",
	//   "id": "source.projects.repos.revisions.listFiles",
	//   "parameterOrder": [
	//     "projectId",
	//     "repoName",
	//     "revisionId"
	//   ],
	//   "parameters": {
	//     "aliasContext.kind": {
	//       "description": "The alias kind.",
	//       "enum": [
	//         "ANY",
	//         "FIXED",
	//         "MOVABLE",
	//         "OTHER"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "aliasContext.name": {
	//       "description": "The alias name.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "aliasName": {
	//       "description": "The name of an alias (branch, tag, etc.).",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "cloudWorkspace.snapshotId": {
	//       "description": "The ID of the snapshot.\nAn empty snapshot_id refers to the most recent snapshot.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "cloudWorkspace.workspaceId.name": {
	//       "description": "The unique name of the workspace within the repo.  This is the name\nchosen by the client in the Source API's CreateWorkspace method.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "cloudWorkspace.workspaceId.repoId.projectRepoId.projectId": {
	//       "description": "The ID of the project.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "cloudWorkspace.workspaceId.repoId.projectRepoId.repoName": {
	//       "description": "The name of the repo. Leave empty for the default repo.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "cloudWorkspace.workspaceId.repoId.uid": {
	//       "description": "A server-assigned, globally unique identifier.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "gerrit.aliasContext.kind": {
	//       "description": "The alias kind.",
	//       "enum": [
	//         "ANY",
	//         "FIXED",
	//         "MOVABLE",
	//         "OTHER"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "gerrit.aliasContext.name": {
	//       "description": "The alias name.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "gerrit.aliasName": {
	//       "description": "The name of an alias (branch, tag, etc.).",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "gerrit.gerritProject": {
	//       "description": "The full project name within the host. Projects may be nested, so\n\"project/subproject\" is a valid project name.\nThe \"repo name\" is hostURI/project.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "gerrit.hostUri": {
	//       "description": "The URI of a running Gerrit instance.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "gerrit.revisionId": {
	//       "description": "A revision (commit) ID.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "git.revisionId": {
	//       "description": "Git commit hash.\nrequired.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "git.url": {
	//       "description": "Git repository URL.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "The maximum number of values to return.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The value of next_page_token from the previous call.\nOmit for the first page.",
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
	//       "description": "The name of the repo. Leave empty for the default repo.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "revisionId": {
	//       "description": "A revision ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "uid": {
	//       "description": "A server-assigned, globally unique identifier.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/repos/{repoName}/revisions/{revisionId}:listFiles",
	//   "response": {
	//     "$ref": "ListFilesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "source.projects.repos.revisions.files.get":

type ProjectsReposRevisionsFilesGetCall struct {
	s          *Service
	projectId  string
	repoName   string
	revisionId string
	path       string
	opt_       map[string]interface{}
}

// Get: Read is given a SourceContext and path, and returns
// file or
// directory information about that path.
func (r *ProjectsReposRevisionsFilesService) Get(projectId string, repoName string, revisionId string, path string) *ProjectsReposRevisionsFilesGetCall {
	c := &ProjectsReposRevisionsFilesGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.repoName = repoName
	c.revisionId = revisionId
	c.path = path
	return c
}

// AliasContextKind sets the optional parameter "aliasContext.kind": The
// alias kind.
func (c *ProjectsReposRevisionsFilesGetCall) AliasContextKind(aliasContextKind string) *ProjectsReposRevisionsFilesGetCall {
	c.opt_["aliasContext.kind"] = aliasContextKind
	return c
}

// AliasContextName sets the optional parameter "aliasContext.name": The
// alias name.
func (c *ProjectsReposRevisionsFilesGetCall) AliasContextName(aliasContextName string) *ProjectsReposRevisionsFilesGetCall {
	c.opt_["aliasContext.name"] = aliasContextName
	return c
}

// AliasName sets the optional parameter "aliasName": The name of an
// alias (branch, tag, etc.).
func (c *ProjectsReposRevisionsFilesGetCall) AliasName(aliasName string) *ProjectsReposRevisionsFilesGetCall {
	c.opt_["aliasName"] = aliasName
	return c
}

// CloudWorkspaceSnapshotId sets the optional parameter
// "cloudWorkspace.snapshotId": The ID of the snapshot.
// An empty
// snapshot_id refers to the most recent snapshot.
func (c *ProjectsReposRevisionsFilesGetCall) CloudWorkspaceSnapshotId(cloudWorkspaceSnapshotId string) *ProjectsReposRevisionsFilesGetCall {
	c.opt_["cloudWorkspace.snapshotId"] = cloudWorkspaceSnapshotId
	return c
}

// CloudWorkspaceWorkspaceIdName sets the optional parameter
// "cloudWorkspace.workspaceId.name": The unique name of the workspace
// within the repo.  This is the name
// chosen by the client in the Source
// API's CreateWorkspace method.
func (c *ProjectsReposRevisionsFilesGetCall) CloudWorkspaceWorkspaceIdName(cloudWorkspaceWorkspaceIdName string) *ProjectsReposRevisionsFilesGetCall {
	c.opt_["cloudWorkspace.workspaceId.name"] = cloudWorkspaceWorkspaceIdName
	return c
}

// CloudWorkspaceWorkspaceIdRepoIdProjectRepoIdProjectId sets the
// optional parameter
// "cloudWorkspace.workspaceId.repoId.projectRepoId.projectId": The ID
// of the project.
func (c *ProjectsReposRevisionsFilesGetCall) CloudWorkspaceWorkspaceIdRepoIdProjectRepoIdProjectId(cloudWorkspaceWorkspaceIdRepoIdProjectRepoIdProjectId string) *ProjectsReposRevisionsFilesGetCall {
	c.opt_["cloudWorkspace.workspaceId.repoId.projectRepoId.projectId"] = cloudWorkspaceWorkspaceIdRepoIdProjectRepoIdProjectId
	return c
}

// CloudWorkspaceWorkspaceIdRepoIdProjectRepoIdRepoName sets the
// optional parameter
// "cloudWorkspace.workspaceId.repoId.projectRepoId.repoName": The name
// of the repo. Leave empty for the default repo.
func (c *ProjectsReposRevisionsFilesGetCall) CloudWorkspaceWorkspaceIdRepoIdProjectRepoIdRepoName(cloudWorkspaceWorkspaceIdRepoIdProjectRepoIdRepoName string) *ProjectsReposRevisionsFilesGetCall {
	c.opt_["cloudWorkspace.workspaceId.repoId.projectRepoId.repoName"] = cloudWorkspaceWorkspaceIdRepoIdProjectRepoIdRepoName
	return c
}

// CloudWorkspaceWorkspaceIdRepoIdUid sets the optional parameter
// "cloudWorkspace.workspaceId.repoId.uid": A server-assigned, globally
// unique identifier.
func (c *ProjectsReposRevisionsFilesGetCall) CloudWorkspaceWorkspaceIdRepoIdUid(cloudWorkspaceWorkspaceIdRepoIdUid string) *ProjectsReposRevisionsFilesGetCall {
	c.opt_["cloudWorkspace.workspaceId.repoId.uid"] = cloudWorkspaceWorkspaceIdRepoIdUid
	return c
}

// GerritAliasContextKind sets the optional parameter
// "gerrit.aliasContext.kind": The alias kind.
func (c *ProjectsReposRevisionsFilesGetCall) GerritAliasContextKind(gerritAliasContextKind string) *ProjectsReposRevisionsFilesGetCall {
	c.opt_["gerrit.aliasContext.kind"] = gerritAliasContextKind
	return c
}

// GerritAliasContextName sets the optional parameter
// "gerrit.aliasContext.name": The alias name.
func (c *ProjectsReposRevisionsFilesGetCall) GerritAliasContextName(gerritAliasContextName string) *ProjectsReposRevisionsFilesGetCall {
	c.opt_["gerrit.aliasContext.name"] = gerritAliasContextName
	return c
}

// GerritAliasName sets the optional parameter "gerrit.aliasName": The
// name of an alias (branch, tag, etc.).
func (c *ProjectsReposRevisionsFilesGetCall) GerritAliasName(gerritAliasName string) *ProjectsReposRevisionsFilesGetCall {
	c.opt_["gerrit.aliasName"] = gerritAliasName
	return c
}

// GerritGerritProject sets the optional parameter
// "gerrit.gerritProject": The full project name within the host.
// Projects may be nested, so
// "project/subproject" is a valid project
// name.
// The "repo name" is hostURI/project.
func (c *ProjectsReposRevisionsFilesGetCall) GerritGerritProject(gerritGerritProject string) *ProjectsReposRevisionsFilesGetCall {
	c.opt_["gerrit.gerritProject"] = gerritGerritProject
	return c
}

// GerritHostUri sets the optional parameter "gerrit.hostUri": The URI
// of a running Gerrit instance.
func (c *ProjectsReposRevisionsFilesGetCall) GerritHostUri(gerritHostUri string) *ProjectsReposRevisionsFilesGetCall {
	c.opt_["gerrit.hostUri"] = gerritHostUri
	return c
}

// GerritRevisionId sets the optional parameter "gerrit.revisionId": A
// revision (commit) ID.
func (c *ProjectsReposRevisionsFilesGetCall) GerritRevisionId(gerritRevisionId string) *ProjectsReposRevisionsFilesGetCall {
	c.opt_["gerrit.revisionId"] = gerritRevisionId
	return c
}

// GitRevisionId sets the optional parameter "git.revisionId": Git
// commit hash.
// required.
func (c *ProjectsReposRevisionsFilesGetCall) GitRevisionId(gitRevisionId string) *ProjectsReposRevisionsFilesGetCall {
	c.opt_["git.revisionId"] = gitRevisionId
	return c
}

// GitUrl sets the optional parameter "git.url": Git repository URL.
func (c *ProjectsReposRevisionsFilesGetCall) GitUrl(gitUrl string) *ProjectsReposRevisionsFilesGetCall {
	c.opt_["git.url"] = gitUrl
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of values to return.
func (c *ProjectsReposRevisionsFilesGetCall) PageSize(pageSize int64) *ProjectsReposRevisionsFilesGetCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken": The value of
// next_page_token from the previous call.
// Omit for the first page, or
// if using start_index.
func (c *ProjectsReposRevisionsFilesGetCall) PageToken(pageToken string) *ProjectsReposRevisionsFilesGetCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// StartPosition sets the optional parameter "startPosition": If path
// refers to a file, the position of the first byte of its contents
// to
// return. If path refers to a directory, the position of the first
// entry
// in the listing. If page_token is specified, this field is
// ignored.
func (c *ProjectsReposRevisionsFilesGetCall) StartPosition(startPosition int64) *ProjectsReposRevisionsFilesGetCall {
	c.opt_["startPosition"] = startPosition
	return c
}

// Uid sets the optional parameter "uid": A server-assigned, globally
// unique identifier.
func (c *ProjectsReposRevisionsFilesGetCall) Uid(uid string) *ProjectsReposRevisionsFilesGetCall {
	c.opt_["uid"] = uid
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsReposRevisionsFilesGetCall) Fields(s ...googleapi.Field) *ProjectsReposRevisionsFilesGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsReposRevisionsFilesGetCall) Do() (*ReadResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["aliasContext.kind"]; ok {
		params.Set("aliasContext.kind", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["aliasContext.name"]; ok {
		params.Set("aliasContext.name", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["aliasName"]; ok {
		params.Set("aliasName", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["cloudWorkspace.snapshotId"]; ok {
		params.Set("cloudWorkspace.snapshotId", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["cloudWorkspace.workspaceId.name"]; ok {
		params.Set("cloudWorkspace.workspaceId.name", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["cloudWorkspace.workspaceId.repoId.projectRepoId.projectId"]; ok {
		params.Set("cloudWorkspace.workspaceId.repoId.projectRepoId.projectId", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["cloudWorkspace.workspaceId.repoId.projectRepoId.repoName"]; ok {
		params.Set("cloudWorkspace.workspaceId.repoId.projectRepoId.repoName", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["cloudWorkspace.workspaceId.repoId.uid"]; ok {
		params.Set("cloudWorkspace.workspaceId.repoId.uid", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["gerrit.aliasContext.kind"]; ok {
		params.Set("gerrit.aliasContext.kind", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["gerrit.aliasContext.name"]; ok {
		params.Set("gerrit.aliasContext.name", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["gerrit.aliasName"]; ok {
		params.Set("gerrit.aliasName", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["gerrit.gerritProject"]; ok {
		params.Set("gerrit.gerritProject", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["gerrit.hostUri"]; ok {
		params.Set("gerrit.hostUri", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["gerrit.revisionId"]; ok {
		params.Set("gerrit.revisionId", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["git.revisionId"]; ok {
		params.Set("git.revisionId", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["git.url"]; ok {
		params.Set("git.url", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageSize"]; ok {
		params.Set("pageSize", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["startPosition"]; ok {
		params.Set("startPosition", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["uid"]; ok {
		params.Set("uid", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/repos/{repoName}/revisions/{revisionId}/files/{+path}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId":  c.projectId,
		"repoName":   c.repoName,
		"revisionId": c.revisionId,
		"path":       c.path,
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
	var ret *ReadResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Read is given a SourceContext and path, and returns\nfile or directory information about that path.",
	//   "flatPath": "v1/projects/{projectId}/repos/{repoName}/revisions/{revisionId}/files/{filesId}",
	//   "httpMethod": "GET",
	//   "id": "source.projects.repos.revisions.files.get",
	//   "parameterOrder": [
	//     "projectId",
	//     "repoName",
	//     "revisionId",
	//     "path"
	//   ],
	//   "parameters": {
	//     "aliasContext.kind": {
	//       "description": "The alias kind.",
	//       "enum": [
	//         "ANY",
	//         "FIXED",
	//         "MOVABLE",
	//         "OTHER"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "aliasContext.name": {
	//       "description": "The alias name.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "aliasName": {
	//       "description": "The name of an alias (branch, tag, etc.).",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "cloudWorkspace.snapshotId": {
	//       "description": "The ID of the snapshot.\nAn empty snapshot_id refers to the most recent snapshot.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "cloudWorkspace.workspaceId.name": {
	//       "description": "The unique name of the workspace within the repo.  This is the name\nchosen by the client in the Source API's CreateWorkspace method.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "cloudWorkspace.workspaceId.repoId.projectRepoId.projectId": {
	//       "description": "The ID of the project.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "cloudWorkspace.workspaceId.repoId.projectRepoId.repoName": {
	//       "description": "The name of the repo. Leave empty for the default repo.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "cloudWorkspace.workspaceId.repoId.uid": {
	//       "description": "A server-assigned, globally unique identifier.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "gerrit.aliasContext.kind": {
	//       "description": "The alias kind.",
	//       "enum": [
	//         "ANY",
	//         "FIXED",
	//         "MOVABLE",
	//         "OTHER"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "gerrit.aliasContext.name": {
	//       "description": "The alias name.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "gerrit.aliasName": {
	//       "description": "The name of an alias (branch, tag, etc.).",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "gerrit.gerritProject": {
	//       "description": "The full project name within the host. Projects may be nested, so\n\"project/subproject\" is a valid project name.\nThe \"repo name\" is hostURI/project.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "gerrit.hostUri": {
	//       "description": "The URI of a running Gerrit instance.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "gerrit.revisionId": {
	//       "description": "A revision (commit) ID.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "git.revisionId": {
	//       "description": "Git commit hash.\nrequired.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "git.url": {
	//       "description": "Git repository URL.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "The maximum number of values to return.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "The value of next_page_token from the previous call.\nOmit for the first page, or if using start_index.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "path": {
	//       "description": "Path to the file or directory from the root directory of the source\ncontext. It must not have leading or trailing slashes.",
	//       "location": "path",
	//       "pattern": "^.*$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "The ID of the project.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "repoName": {
	//       "description": "The name of the repo. Leave empty for the default repo.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "revisionId": {
	//       "description": "A revision ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "startPosition": {
	//       "description": "If path refers to a file, the position of the first byte of its contents\nto return. If path refers to a directory, the position of the first entry\nin the listing. If page_token is specified, this field is ignored.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "uid": {
	//       "description": "A server-assigned, globally unique identifier.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/repos/{repoName}/revisions/{revisionId}/files/{+path}",
	//   "response": {
	//     "$ref": "ReadResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "source.projects.repos.workspaces.commitWorkspace":

type ProjectsReposWorkspacesCommitWorkspaceCall struct {
	s                      *Service
	projectId              string
	repoName               string
	name                   string
	commitworkspacerequest *CommitWorkspaceRequest
	opt_                   map[string]interface{}
}

// CommitWorkspace: Commits some or all of the modified files in a
// workspace. This creates a
// new revision in the repo with the
// workspace's contents. Returns ABORTED if the workspace ID
// in the
// request contains a snapshot ID and it is not the same as
// the
// workspace's current snapshot ID or if the workspace is
// simultaneously
// modified by another client.
func (r *ProjectsReposWorkspacesService) CommitWorkspace(projectId string, repoName string, name string, commitworkspacerequest *CommitWorkspaceRequest) *ProjectsReposWorkspacesCommitWorkspaceCall {
	c := &ProjectsReposWorkspacesCommitWorkspaceCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.repoName = repoName
	c.name = name
	c.commitworkspacerequest = commitworkspacerequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsReposWorkspacesCommitWorkspaceCall) Fields(s ...googleapi.Field) *ProjectsReposWorkspacesCommitWorkspaceCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsReposWorkspacesCommitWorkspaceCall) Do() (*Workspace, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.commitworkspacerequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/repos/{repoName}/workspaces/{name}:commitWorkspace")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"repoName":  c.repoName,
		"name":      c.name,
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
	var ret *Workspace
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Commits some or all of the modified files in a workspace. This creates a\nnew revision in the repo with the workspace's contents. Returns ABORTED if the workspace ID\nin the request contains a snapshot ID and it is not the same as the\nworkspace's current snapshot ID or if the workspace is simultaneously\nmodified by another client.",
	//   "flatPath": "v1/projects/{projectId}/repos/{repoName}/workspaces/{name}:commitWorkspace",
	//   "httpMethod": "POST",
	//   "id": "source.projects.repos.workspaces.commitWorkspace",
	//   "parameterOrder": [
	//     "projectId",
	//     "repoName",
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The unique name of the workspace within the repo.  This is the name\nchosen by the client in the Source API's CreateWorkspace method.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "The ID of the project.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "repoName": {
	//       "description": "The name of the repo. Leave empty for the default repo.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/repos/{repoName}/workspaces/{name}:commitWorkspace",
	//   "request": {
	//     "$ref": "CommitWorkspaceRequest"
	//   },
	//   "response": {
	//     "$ref": "Workspace"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "source.projects.repos.workspaces.create":

type ProjectsReposWorkspacesCreateCall struct {
	s                      *Service
	projectId              string
	repoName               string
	createworkspacerequest *CreateWorkspaceRequest
	opt_                   map[string]interface{}
}

// Create: Creates a workspace.
func (r *ProjectsReposWorkspacesService) Create(projectId string, repoName string, createworkspacerequest *CreateWorkspaceRequest) *ProjectsReposWorkspacesCreateCall {
	c := &ProjectsReposWorkspacesCreateCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.repoName = repoName
	c.createworkspacerequest = createworkspacerequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsReposWorkspacesCreateCall) Fields(s ...googleapi.Field) *ProjectsReposWorkspacesCreateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsReposWorkspacesCreateCall) Do() (*Workspace, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.createworkspacerequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/repos/{repoName}/workspaces")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"repoName":  c.repoName,
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
	var ret *Workspace
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a workspace.",
	//   "flatPath": "v1/projects/{projectId}/repos/{repoName}/workspaces",
	//   "httpMethod": "POST",
	//   "id": "source.projects.repos.workspaces.create",
	//   "parameterOrder": [
	//     "projectId",
	//     "repoName"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "The ID of the project.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "repoName": {
	//       "description": "The name of the repo. Leave empty for the default repo.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/repos/{repoName}/workspaces",
	//   "request": {
	//     "$ref": "CreateWorkspaceRequest"
	//   },
	//   "response": {
	//     "$ref": "Workspace"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "source.projects.repos.workspaces.delete":

type ProjectsReposWorkspacesDeleteCall struct {
	s         *Service
	projectId string
	repoName  string
	name      string
	opt_      map[string]interface{}
}

// Delete: Deletes a workspace. Uncommitted changes are lost. If the
// workspace does
// not exist, NOT_FOUND is returned. Returns ABORTED when
// the workspace is
// simultaneously modified by another client.
func (r *ProjectsReposWorkspacesService) Delete(projectId string, repoName string, name string) *ProjectsReposWorkspacesDeleteCall {
	c := &ProjectsReposWorkspacesDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.repoName = repoName
	c.name = name
	return c
}

// CurrentSnapshotId sets the optional parameter "currentSnapshotId": If
// non-empty, current_snapshot_id must refer to the most recent update
// to
// the workspace, or ABORTED is returned.
func (c *ProjectsReposWorkspacesDeleteCall) CurrentSnapshotId(currentSnapshotId string) *ProjectsReposWorkspacesDeleteCall {
	c.opt_["currentSnapshotId"] = currentSnapshotId
	return c
}

// Uid sets the optional parameter "uid": A server-assigned, globally
// unique identifier.
func (c *ProjectsReposWorkspacesDeleteCall) Uid(uid string) *ProjectsReposWorkspacesDeleteCall {
	c.opt_["uid"] = uid
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsReposWorkspacesDeleteCall) Fields(s ...googleapi.Field) *ProjectsReposWorkspacesDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsReposWorkspacesDeleteCall) Do() (*Empty, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["currentSnapshotId"]; ok {
		params.Set("currentSnapshotId", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["uid"]; ok {
		params.Set("uid", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/repos/{repoName}/workspaces/{name}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"repoName":  c.repoName,
		"name":      c.name,
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
	//   "description": "Deletes a workspace. Uncommitted changes are lost. If the workspace does\nnot exist, NOT_FOUND is returned. Returns ABORTED when the workspace is\nsimultaneously modified by another client.",
	//   "flatPath": "v1/projects/{projectId}/repos/{repoName}/workspaces/{name}",
	//   "httpMethod": "DELETE",
	//   "id": "source.projects.repos.workspaces.delete",
	//   "parameterOrder": [
	//     "projectId",
	//     "repoName",
	//     "name"
	//   ],
	//   "parameters": {
	//     "currentSnapshotId": {
	//       "description": "If non-empty, current_snapshot_id must refer to the most recent update to\nthe workspace, or ABORTED is returned.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "name": {
	//       "description": "The unique name of the workspace within the repo.  This is the name\nchosen by the client in the Source API's CreateWorkspace method.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "The ID of the project.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "repoName": {
	//       "description": "The name of the repo. Leave empty for the default repo.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "uid": {
	//       "description": "A server-assigned, globally unique identifier.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/repos/{repoName}/workspaces/{name}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "source.projects.repos.workspaces.get":

type ProjectsReposWorkspacesGetCall struct {
	s         *Service
	projectId string
	repoName  string
	name      string
	opt_      map[string]interface{}
}

// Get: Returns workspace metadata.
func (r *ProjectsReposWorkspacesService) Get(projectId string, repoName string, name string) *ProjectsReposWorkspacesGetCall {
	c := &ProjectsReposWorkspacesGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.repoName = repoName
	c.name = name
	return c
}

// Uid sets the optional parameter "uid": A server-assigned, globally
// unique identifier.
func (c *ProjectsReposWorkspacesGetCall) Uid(uid string) *ProjectsReposWorkspacesGetCall {
	c.opt_["uid"] = uid
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsReposWorkspacesGetCall) Fields(s ...googleapi.Field) *ProjectsReposWorkspacesGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsReposWorkspacesGetCall) Do() (*Workspace, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["uid"]; ok {
		params.Set("uid", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/repos/{repoName}/workspaces/{name}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"repoName":  c.repoName,
		"name":      c.name,
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
	var ret *Workspace
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns workspace metadata.",
	//   "flatPath": "v1/projects/{projectId}/repos/{repoName}/workspaces/{name}",
	//   "httpMethod": "GET",
	//   "id": "source.projects.repos.workspaces.get",
	//   "parameterOrder": [
	//     "projectId",
	//     "repoName",
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The unique name of the workspace within the repo.  This is the name\nchosen by the client in the Source API's CreateWorkspace method.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "The ID of the project.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "repoName": {
	//       "description": "The name of the repo. Leave empty for the default repo.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "uid": {
	//       "description": "A server-assigned, globally unique identifier.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/repos/{repoName}/workspaces/{name}",
	//   "response": {
	//     "$ref": "Workspace"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "source.projects.repos.workspaces.list":

type ProjectsReposWorkspacesListCall struct {
	s         *Service
	projectId string
	repoName  string
	opt_      map[string]interface{}
}

// List: Returns all workspaces belonging to a repo.
func (r *ProjectsReposWorkspacesService) List(projectId string, repoName string) *ProjectsReposWorkspacesListCall {
	c := &ProjectsReposWorkspacesListCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.repoName = repoName
	return c
}

// Uid sets the optional parameter "uid": A server-assigned, globally
// unique identifier.
func (c *ProjectsReposWorkspacesListCall) Uid(uid string) *ProjectsReposWorkspacesListCall {
	c.opt_["uid"] = uid
	return c
}

// View sets the optional parameter "view": Specifies which parts of the
// Workspace resource should be returned in the
// response.
func (c *ProjectsReposWorkspacesListCall) View(view string) *ProjectsReposWorkspacesListCall {
	c.opt_["view"] = view
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsReposWorkspacesListCall) Fields(s ...googleapi.Field) *ProjectsReposWorkspacesListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsReposWorkspacesListCall) Do() (*ListWorkspacesResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["uid"]; ok {
		params.Set("uid", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["view"]; ok {
		params.Set("view", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/repos/{repoName}/workspaces")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"repoName":  c.repoName,
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
	var ret *ListWorkspacesResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns all workspaces belonging to a repo.",
	//   "flatPath": "v1/projects/{projectId}/repos/{repoName}/workspaces",
	//   "httpMethod": "GET",
	//   "id": "source.projects.repos.workspaces.list",
	//   "parameterOrder": [
	//     "projectId",
	//     "repoName"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "The ID of the project.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "repoName": {
	//       "description": "The name of the repo. Leave empty for the default repo.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "uid": {
	//       "description": "A server-assigned, globally unique identifier.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "view": {
	//       "description": "Specifies which parts of the Workspace resource should be returned in the\nresponse.",
	//       "enum": [
	//         "STANDARD",
	//         "MINIMAL",
	//         "FULL"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/repos/{repoName}/workspaces",
	//   "response": {
	//     "$ref": "ListWorkspacesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "source.projects.repos.workspaces.listFiles":

type ProjectsReposWorkspacesListFilesCall struct {
	s         *Service
	projectId string
	repoName  string
	name      string
	opt_      map[string]interface{}
}

// ListFiles: ListFiles returns a list of all files in a SourceContext.
// The
// information about each file includes its path and its hash.
// The
// result is ordered by path. Pagination is supported.
func (r *ProjectsReposWorkspacesService) ListFiles(projectId string, repoName string, name string) *ProjectsReposWorkspacesListFilesCall {
	c := &ProjectsReposWorkspacesListFilesCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.repoName = repoName
	c.name = name
	return c
}

// CloudRepoAliasContextKind sets the optional parameter
// "cloudRepo.aliasContext.kind": The alias kind.
func (c *ProjectsReposWorkspacesListFilesCall) CloudRepoAliasContextKind(cloudRepoAliasContextKind string) *ProjectsReposWorkspacesListFilesCall {
	c.opt_["cloudRepo.aliasContext.kind"] = cloudRepoAliasContextKind
	return c
}

// CloudRepoAliasContextName sets the optional parameter
// "cloudRepo.aliasContext.name": The alias name.
func (c *ProjectsReposWorkspacesListFilesCall) CloudRepoAliasContextName(cloudRepoAliasContextName string) *ProjectsReposWorkspacesListFilesCall {
	c.opt_["cloudRepo.aliasContext.name"] = cloudRepoAliasContextName
	return c
}

// CloudRepoAliasName sets the optional parameter "cloudRepo.aliasName":
// The name of an alias (branch, tag, etc.).
func (c *ProjectsReposWorkspacesListFilesCall) CloudRepoAliasName(cloudRepoAliasName string) *ProjectsReposWorkspacesListFilesCall {
	c.opt_["cloudRepo.aliasName"] = cloudRepoAliasName
	return c
}

// CloudRepoRepoIdProjectRepoIdProjectId sets the optional parameter
// "cloudRepo.repoId.projectRepoId.projectId": The ID of the project.
func (c *ProjectsReposWorkspacesListFilesCall) CloudRepoRepoIdProjectRepoIdProjectId(cloudRepoRepoIdProjectRepoIdProjectId string) *ProjectsReposWorkspacesListFilesCall {
	c.opt_["cloudRepo.repoId.projectRepoId.projectId"] = cloudRepoRepoIdProjectRepoIdProjectId
	return c
}

// CloudRepoRepoIdProjectRepoIdRepoName sets the optional parameter
// "cloudRepo.repoId.projectRepoId.repoName": The name of the repo.
// Leave empty for the default repo.
func (c *ProjectsReposWorkspacesListFilesCall) CloudRepoRepoIdProjectRepoIdRepoName(cloudRepoRepoIdProjectRepoIdRepoName string) *ProjectsReposWorkspacesListFilesCall {
	c.opt_["cloudRepo.repoId.projectRepoId.repoName"] = cloudRepoRepoIdProjectRepoIdRepoName
	return c
}

// CloudRepoRepoIdUid sets the optional parameter
// "cloudRepo.repoId.uid": A server-assigned, globally unique
// identifier.
func (c *ProjectsReposWorkspacesListFilesCall) CloudRepoRepoIdUid(cloudRepoRepoIdUid string) *ProjectsReposWorkspacesListFilesCall {
	c.opt_["cloudRepo.repoId.uid"] = cloudRepoRepoIdUid
	return c
}

// CloudRepoRevisionId sets the optional parameter
// "cloudRepo.revisionId": A revision ID.
func (c *ProjectsReposWorkspacesListFilesCall) CloudRepoRevisionId(cloudRepoRevisionId string) *ProjectsReposWorkspacesListFilesCall {
	c.opt_["cloudRepo.revisionId"] = cloudRepoRevisionId
	return c
}

// GerritAliasContextKind sets the optional parameter
// "gerrit.aliasContext.kind": The alias kind.
func (c *ProjectsReposWorkspacesListFilesCall) GerritAliasContextKind(gerritAliasContextKind string) *ProjectsReposWorkspacesListFilesCall {
	c.opt_["gerrit.aliasContext.kind"] = gerritAliasContextKind
	return c
}

// GerritAliasContextName sets the optional parameter
// "gerrit.aliasContext.name": The alias name.
func (c *ProjectsReposWorkspacesListFilesCall) GerritAliasContextName(gerritAliasContextName string) *ProjectsReposWorkspacesListFilesCall {
	c.opt_["gerrit.aliasContext.name"] = gerritAliasContextName
	return c
}

// GerritAliasName sets the optional parameter "gerrit.aliasName": The
// name of an alias (branch, tag, etc.).
func (c *ProjectsReposWorkspacesListFilesCall) GerritAliasName(gerritAliasName string) *ProjectsReposWorkspacesListFilesCall {
	c.opt_["gerrit.aliasName"] = gerritAliasName
	return c
}

// GerritGerritProject sets the optional parameter
// "gerrit.gerritProject": The full project name within the host.
// Projects may be nested, so
// "project/subproject" is a valid project
// name.
// The "repo name" is hostURI/project.
func (c *ProjectsReposWorkspacesListFilesCall) GerritGerritProject(gerritGerritProject string) *ProjectsReposWorkspacesListFilesCall {
	c.opt_["gerrit.gerritProject"] = gerritGerritProject
	return c
}

// GerritHostUri sets the optional parameter "gerrit.hostUri": The URI
// of a running Gerrit instance.
func (c *ProjectsReposWorkspacesListFilesCall) GerritHostUri(gerritHostUri string) *ProjectsReposWorkspacesListFilesCall {
	c.opt_["gerrit.hostUri"] = gerritHostUri
	return c
}

// GerritRevisionId sets the optional parameter "gerrit.revisionId": A
// revision (commit) ID.
func (c *ProjectsReposWorkspacesListFilesCall) GerritRevisionId(gerritRevisionId string) *ProjectsReposWorkspacesListFilesCall {
	c.opt_["gerrit.revisionId"] = gerritRevisionId
	return c
}

// GitRevisionId sets the optional parameter "git.revisionId": Git
// commit hash.
// required.
func (c *ProjectsReposWorkspacesListFilesCall) GitRevisionId(gitRevisionId string) *ProjectsReposWorkspacesListFilesCall {
	c.opt_["git.revisionId"] = gitRevisionId
	return c
}

// GitUrl sets the optional parameter "git.url": Git repository URL.
func (c *ProjectsReposWorkspacesListFilesCall) GitUrl(gitUrl string) *ProjectsReposWorkspacesListFilesCall {
	c.opt_["git.url"] = gitUrl
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of values to return.
func (c *ProjectsReposWorkspacesListFilesCall) PageSize(pageSize int64) *ProjectsReposWorkspacesListFilesCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken": The value of
// next_page_token from the previous call.
// Omit for the first page.
func (c *ProjectsReposWorkspacesListFilesCall) PageToken(pageToken string) *ProjectsReposWorkspacesListFilesCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// SnapshotId sets the optional parameter "snapshotId": The ID of the
// snapshot.
// An empty snapshot_id refers to the most recent snapshot.
func (c *ProjectsReposWorkspacesListFilesCall) SnapshotId(snapshotId string) *ProjectsReposWorkspacesListFilesCall {
	c.opt_["snapshotId"] = snapshotId
	return c
}

// Uid sets the optional parameter "uid": A server-assigned, globally
// unique identifier.
func (c *ProjectsReposWorkspacesListFilesCall) Uid(uid string) *ProjectsReposWorkspacesListFilesCall {
	c.opt_["uid"] = uid
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsReposWorkspacesListFilesCall) Fields(s ...googleapi.Field) *ProjectsReposWorkspacesListFilesCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsReposWorkspacesListFilesCall) Do() (*ListFilesResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["cloudRepo.aliasContext.kind"]; ok {
		params.Set("cloudRepo.aliasContext.kind", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["cloudRepo.aliasContext.name"]; ok {
		params.Set("cloudRepo.aliasContext.name", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["cloudRepo.aliasName"]; ok {
		params.Set("cloudRepo.aliasName", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["cloudRepo.repoId.projectRepoId.projectId"]; ok {
		params.Set("cloudRepo.repoId.projectRepoId.projectId", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["cloudRepo.repoId.projectRepoId.repoName"]; ok {
		params.Set("cloudRepo.repoId.projectRepoId.repoName", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["cloudRepo.repoId.uid"]; ok {
		params.Set("cloudRepo.repoId.uid", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["cloudRepo.revisionId"]; ok {
		params.Set("cloudRepo.revisionId", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["gerrit.aliasContext.kind"]; ok {
		params.Set("gerrit.aliasContext.kind", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["gerrit.aliasContext.name"]; ok {
		params.Set("gerrit.aliasContext.name", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["gerrit.aliasName"]; ok {
		params.Set("gerrit.aliasName", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["gerrit.gerritProject"]; ok {
		params.Set("gerrit.gerritProject", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["gerrit.hostUri"]; ok {
		params.Set("gerrit.hostUri", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["gerrit.revisionId"]; ok {
		params.Set("gerrit.revisionId", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["git.revisionId"]; ok {
		params.Set("git.revisionId", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["git.url"]; ok {
		params.Set("git.url", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageSize"]; ok {
		params.Set("pageSize", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["snapshotId"]; ok {
		params.Set("snapshotId", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["uid"]; ok {
		params.Set("uid", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/repos/{repoName}/workspaces/{name}:listFiles")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"repoName":  c.repoName,
		"name":      c.name,
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
	var ret *ListFilesResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "ListFiles returns a list of all files in a SourceContext. The\ninformation about each file includes its path and its hash.\nThe result is ordered by path. Pagination is supported.",
	//   "flatPath": "v1/projects/{projectId}/repos/{repoName}/workspaces/{name}:listFiles",
	//   "httpMethod": "GET",
	//   "id": "source.projects.repos.workspaces.listFiles",
	//   "parameterOrder": [
	//     "projectId",
	//     "repoName",
	//     "name"
	//   ],
	//   "parameters": {
	//     "cloudRepo.aliasContext.kind": {
	//       "description": "The alias kind.",
	//       "enum": [
	//         "ANY",
	//         "FIXED",
	//         "MOVABLE",
	//         "OTHER"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "cloudRepo.aliasContext.name": {
	//       "description": "The alias name.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "cloudRepo.aliasName": {
	//       "description": "The name of an alias (branch, tag, etc.).",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "cloudRepo.repoId.projectRepoId.projectId": {
	//       "description": "The ID of the project.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "cloudRepo.repoId.projectRepoId.repoName": {
	//       "description": "The name of the repo. Leave empty for the default repo.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "cloudRepo.repoId.uid": {
	//       "description": "A server-assigned, globally unique identifier.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "cloudRepo.revisionId": {
	//       "description": "A revision ID.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "gerrit.aliasContext.kind": {
	//       "description": "The alias kind.",
	//       "enum": [
	//         "ANY",
	//         "FIXED",
	//         "MOVABLE",
	//         "OTHER"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "gerrit.aliasContext.name": {
	//       "description": "The alias name.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "gerrit.aliasName": {
	//       "description": "The name of an alias (branch, tag, etc.).",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "gerrit.gerritProject": {
	//       "description": "The full project name within the host. Projects may be nested, so\n\"project/subproject\" is a valid project name.\nThe \"repo name\" is hostURI/project.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "gerrit.hostUri": {
	//       "description": "The URI of a running Gerrit instance.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "gerrit.revisionId": {
	//       "description": "A revision (commit) ID.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "git.revisionId": {
	//       "description": "Git commit hash.\nrequired.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "git.url": {
	//       "description": "Git repository URL.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "name": {
	//       "description": "The unique name of the workspace within the repo.  This is the name\nchosen by the client in the Source API's CreateWorkspace method.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "The maximum number of values to return.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The value of next_page_token from the previous call.\nOmit for the first page.",
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
	//       "description": "The name of the repo. Leave empty for the default repo.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "snapshotId": {
	//       "description": "The ID of the snapshot.\nAn empty snapshot_id refers to the most recent snapshot.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "uid": {
	//       "description": "A server-assigned, globally unique identifier.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/repos/{repoName}/workspaces/{name}:listFiles",
	//   "response": {
	//     "$ref": "ListFilesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "source.projects.repos.workspaces.modifyWorkspace":

type ProjectsReposWorkspacesModifyWorkspaceCall struct {
	s                      *Service
	projectId              string
	repoName               string
	name                   string
	modifyworkspacerequest *ModifyWorkspaceRequest
	opt_                   map[string]interface{}
}

// ModifyWorkspace: Applies an ordered sequence of file modification
// actions to a workspace.
// Returns ABORTED if current_snapshot_id in the
// request does not refer to
// the most recent update to the workspace or
// if the workspace is
// simultaneously modified by another client.
func (r *ProjectsReposWorkspacesService) ModifyWorkspace(projectId string, repoName string, name string, modifyworkspacerequest *ModifyWorkspaceRequest) *ProjectsReposWorkspacesModifyWorkspaceCall {
	c := &ProjectsReposWorkspacesModifyWorkspaceCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.repoName = repoName
	c.name = name
	c.modifyworkspacerequest = modifyworkspacerequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsReposWorkspacesModifyWorkspaceCall) Fields(s ...googleapi.Field) *ProjectsReposWorkspacesModifyWorkspaceCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsReposWorkspacesModifyWorkspaceCall) Do() (*Workspace, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.modifyworkspacerequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/repos/{repoName}/workspaces/{name}:modifyWorkspace")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"repoName":  c.repoName,
		"name":      c.name,
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
	var ret *Workspace
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Applies an ordered sequence of file modification actions to a workspace.\nReturns ABORTED if current_snapshot_id in the request does not refer to\nthe most recent update to the workspace or if the workspace is\nsimultaneously modified by another client.",
	//   "flatPath": "v1/projects/{projectId}/repos/{repoName}/workspaces/{name}:modifyWorkspace",
	//   "httpMethod": "POST",
	//   "id": "source.projects.repos.workspaces.modifyWorkspace",
	//   "parameterOrder": [
	//     "projectId",
	//     "repoName",
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The unique name of the workspace within the repo.  This is the name\nchosen by the client in the Source API's CreateWorkspace method.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "The ID of the project.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "repoName": {
	//       "description": "The name of the repo. Leave empty for the default repo.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/repos/{repoName}/workspaces/{name}:modifyWorkspace",
	//   "request": {
	//     "$ref": "ModifyWorkspaceRequest"
	//   },
	//   "response": {
	//     "$ref": "Workspace"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "source.projects.repos.workspaces.refreshWorkspace":

type ProjectsReposWorkspacesRefreshWorkspaceCall struct {
	s                       *Service
	projectId               string
	repoName                string
	name                    string
	refreshworkspacerequest *RefreshWorkspaceRequest
	opt_                    map[string]interface{}
}

// RefreshWorkspace: Brings a workspace up to date by merging in the
// changes made between its
// baseline and the revision to which its alias
// currently refers.
// FAILED_PRECONDITION is returned if the alias refers
// to a revision that is
// not a descendant of the workspace baseline, or
// if the workspace has no
// baseline. Returns ABORTED when the workspace
// is simultaneously modified by
// another client.
//
// A refresh may involve
// merging files in the workspace with files in the
// current alias
// revision. If this merge results in conflicts, then the
// workspace is
// in a merge state: the merge_info field of Workspace will
// be
// populated, and conflicting files in the workspace will contain
// conflict
// markers.
func (r *ProjectsReposWorkspacesService) RefreshWorkspace(projectId string, repoName string, name string, refreshworkspacerequest *RefreshWorkspaceRequest) *ProjectsReposWorkspacesRefreshWorkspaceCall {
	c := &ProjectsReposWorkspacesRefreshWorkspaceCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.repoName = repoName
	c.name = name
	c.refreshworkspacerequest = refreshworkspacerequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsReposWorkspacesRefreshWorkspaceCall) Fields(s ...googleapi.Field) *ProjectsReposWorkspacesRefreshWorkspaceCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsReposWorkspacesRefreshWorkspaceCall) Do() (*Workspace, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.refreshworkspacerequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/repos/{repoName}/workspaces/{name}:refreshWorkspace")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"repoName":  c.repoName,
		"name":      c.name,
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
	var ret *Workspace
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Brings a workspace up to date by merging in the changes made between its\nbaseline and the revision to which its alias currently refers.\nFAILED_PRECONDITION is returned if the alias refers to a revision that is\nnot a descendant of the workspace baseline, or if the workspace has no\nbaseline. Returns ABORTED when the workspace is simultaneously modified by\nanother client.\n\nA refresh may involve merging files in the workspace with files in the\ncurrent alias revision. If this merge results in conflicts, then the\nworkspace is in a merge state: the merge_info field of Workspace will be\npopulated, and conflicting files in the workspace will contain conflict\nmarkers.",
	//   "flatPath": "v1/projects/{projectId}/repos/{repoName}/workspaces/{name}:refreshWorkspace",
	//   "httpMethod": "POST",
	//   "id": "source.projects.repos.workspaces.refreshWorkspace",
	//   "parameterOrder": [
	//     "projectId",
	//     "repoName",
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The unique name of the workspace within the repo.  This is the name\nchosen by the client in the Source API's CreateWorkspace method.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "The ID of the project.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "repoName": {
	//       "description": "The name of the repo. Leave empty for the default repo.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/repos/{repoName}/workspaces/{name}:refreshWorkspace",
	//   "request": {
	//     "$ref": "RefreshWorkspaceRequest"
	//   },
	//   "response": {
	//     "$ref": "Workspace"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "source.projects.repos.workspaces.resolveFiles":

type ProjectsReposWorkspacesResolveFilesCall struct {
	s                   *Service
	projectId           string
	repoName            string
	name                string
	resolvefilesrequest *ResolveFilesRequest
	opt_                map[string]interface{}
}

// ResolveFiles: Marks files modified as part of a merge as having been
// resolved. Returns
// ABORTED when the workspace is simultaneously
// modified by another client.
func (r *ProjectsReposWorkspacesService) ResolveFiles(projectId string, repoName string, name string, resolvefilesrequest *ResolveFilesRequest) *ProjectsReposWorkspacesResolveFilesCall {
	c := &ProjectsReposWorkspacesResolveFilesCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.repoName = repoName
	c.name = name
	c.resolvefilesrequest = resolvefilesrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsReposWorkspacesResolveFilesCall) Fields(s ...googleapi.Field) *ProjectsReposWorkspacesResolveFilesCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsReposWorkspacesResolveFilesCall) Do() (*Workspace, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.resolvefilesrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/repos/{repoName}/workspaces/{name}:resolveFiles")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"repoName":  c.repoName,
		"name":      c.name,
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
	var ret *Workspace
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Marks files modified as part of a merge as having been resolved. Returns\nABORTED when the workspace is simultaneously modified by another client.",
	//   "flatPath": "v1/projects/{projectId}/repos/{repoName}/workspaces/{name}:resolveFiles",
	//   "httpMethod": "POST",
	//   "id": "source.projects.repos.workspaces.resolveFiles",
	//   "parameterOrder": [
	//     "projectId",
	//     "repoName",
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The unique name of the workspace within the repo.  This is the name\nchosen by the client in the Source API's CreateWorkspace method.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "The ID of the project.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "repoName": {
	//       "description": "The name of the repo. Leave empty for the default repo.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/repos/{repoName}/workspaces/{name}:resolveFiles",
	//   "request": {
	//     "$ref": "ResolveFilesRequest"
	//   },
	//   "response": {
	//     "$ref": "Workspace"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "source.projects.repos.workspaces.revertRefresh":

type ProjectsReposWorkspacesRevertRefreshCall struct {
	s                    *Service
	projectId            string
	repoName             string
	name                 string
	revertrefreshrequest *RevertRefreshRequest
	opt_                 map[string]interface{}
}

// RevertRefresh: If a call to RefreshWorkspace results in conflicts,
// use RevertRefresh to
// restore the workspace to the state it was in
// before the refresh.  Returns
// FAILED_PRECONDITION if not preceded by a
// call to RefreshWorkspace, or if
// there are no unresolved conflicts
// remaining. Returns ABORTED when the
// workspace is simultaneously
// modified by another client.
func (r *ProjectsReposWorkspacesService) RevertRefresh(projectId string, repoName string, name string, revertrefreshrequest *RevertRefreshRequest) *ProjectsReposWorkspacesRevertRefreshCall {
	c := &ProjectsReposWorkspacesRevertRefreshCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.repoName = repoName
	c.name = name
	c.revertrefreshrequest = revertrefreshrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsReposWorkspacesRevertRefreshCall) Fields(s ...googleapi.Field) *ProjectsReposWorkspacesRevertRefreshCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsReposWorkspacesRevertRefreshCall) Do() (*Workspace, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.revertrefreshrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/repos/{repoName}/workspaces/{name}:revertRefresh")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"repoName":  c.repoName,
		"name":      c.name,
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
	var ret *Workspace
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "If a call to RefreshWorkspace results in conflicts, use RevertRefresh to\nrestore the workspace to the state it was in before the refresh.  Returns\nFAILED_PRECONDITION if not preceded by a call to RefreshWorkspace, or if\nthere are no unresolved conflicts remaining. Returns ABORTED when the\nworkspace is simultaneously modified by another client.",
	//   "flatPath": "v1/projects/{projectId}/repos/{repoName}/workspaces/{name}:revertRefresh",
	//   "httpMethod": "POST",
	//   "id": "source.projects.repos.workspaces.revertRefresh",
	//   "parameterOrder": [
	//     "projectId",
	//     "repoName",
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The unique name of the workspace within the repo.  This is the name\nchosen by the client in the Source API's CreateWorkspace method.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "The ID of the project.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "repoName": {
	//       "description": "The name of the repo. Leave empty for the default repo.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/repos/{repoName}/workspaces/{name}:revertRefresh",
	//   "request": {
	//     "$ref": "RevertRefreshRequest"
	//   },
	//   "response": {
	//     "$ref": "Workspace"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "source.projects.repos.workspaces.files.get":

type ProjectsReposWorkspacesFilesGetCall struct {
	s         *Service
	projectId string
	repoName  string
	name      string
	path      string
	opt_      map[string]interface{}
}

// Get: Read is given a SourceContext and path, and returns
// file or
// directory information about that path.
func (r *ProjectsReposWorkspacesFilesService) Get(projectId string, repoName string, name string, path string) *ProjectsReposWorkspacesFilesGetCall {
	c := &ProjectsReposWorkspacesFilesGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.repoName = repoName
	c.name = name
	c.path = path
	return c
}

// CloudRepoAliasContextKind sets the optional parameter
// "cloudRepo.aliasContext.kind": The alias kind.
func (c *ProjectsReposWorkspacesFilesGetCall) CloudRepoAliasContextKind(cloudRepoAliasContextKind string) *ProjectsReposWorkspacesFilesGetCall {
	c.opt_["cloudRepo.aliasContext.kind"] = cloudRepoAliasContextKind
	return c
}

// CloudRepoAliasContextName sets the optional parameter
// "cloudRepo.aliasContext.name": The alias name.
func (c *ProjectsReposWorkspacesFilesGetCall) CloudRepoAliasContextName(cloudRepoAliasContextName string) *ProjectsReposWorkspacesFilesGetCall {
	c.opt_["cloudRepo.aliasContext.name"] = cloudRepoAliasContextName
	return c
}

// CloudRepoAliasName sets the optional parameter "cloudRepo.aliasName":
// The name of an alias (branch, tag, etc.).
func (c *ProjectsReposWorkspacesFilesGetCall) CloudRepoAliasName(cloudRepoAliasName string) *ProjectsReposWorkspacesFilesGetCall {
	c.opt_["cloudRepo.aliasName"] = cloudRepoAliasName
	return c
}

// CloudRepoRepoIdProjectRepoIdProjectId sets the optional parameter
// "cloudRepo.repoId.projectRepoId.projectId": The ID of the project.
func (c *ProjectsReposWorkspacesFilesGetCall) CloudRepoRepoIdProjectRepoIdProjectId(cloudRepoRepoIdProjectRepoIdProjectId string) *ProjectsReposWorkspacesFilesGetCall {
	c.opt_["cloudRepo.repoId.projectRepoId.projectId"] = cloudRepoRepoIdProjectRepoIdProjectId
	return c
}

// CloudRepoRepoIdProjectRepoIdRepoName sets the optional parameter
// "cloudRepo.repoId.projectRepoId.repoName": The name of the repo.
// Leave empty for the default repo.
func (c *ProjectsReposWorkspacesFilesGetCall) CloudRepoRepoIdProjectRepoIdRepoName(cloudRepoRepoIdProjectRepoIdRepoName string) *ProjectsReposWorkspacesFilesGetCall {
	c.opt_["cloudRepo.repoId.projectRepoId.repoName"] = cloudRepoRepoIdProjectRepoIdRepoName
	return c
}

// CloudRepoRepoIdUid sets the optional parameter
// "cloudRepo.repoId.uid": A server-assigned, globally unique
// identifier.
func (c *ProjectsReposWorkspacesFilesGetCall) CloudRepoRepoIdUid(cloudRepoRepoIdUid string) *ProjectsReposWorkspacesFilesGetCall {
	c.opt_["cloudRepo.repoId.uid"] = cloudRepoRepoIdUid
	return c
}

// CloudRepoRevisionId sets the optional parameter
// "cloudRepo.revisionId": A revision ID.
func (c *ProjectsReposWorkspacesFilesGetCall) CloudRepoRevisionId(cloudRepoRevisionId string) *ProjectsReposWorkspacesFilesGetCall {
	c.opt_["cloudRepo.revisionId"] = cloudRepoRevisionId
	return c
}

// GerritAliasContextKind sets the optional parameter
// "gerrit.aliasContext.kind": The alias kind.
func (c *ProjectsReposWorkspacesFilesGetCall) GerritAliasContextKind(gerritAliasContextKind string) *ProjectsReposWorkspacesFilesGetCall {
	c.opt_["gerrit.aliasContext.kind"] = gerritAliasContextKind
	return c
}

// GerritAliasContextName sets the optional parameter
// "gerrit.aliasContext.name": The alias name.
func (c *ProjectsReposWorkspacesFilesGetCall) GerritAliasContextName(gerritAliasContextName string) *ProjectsReposWorkspacesFilesGetCall {
	c.opt_["gerrit.aliasContext.name"] = gerritAliasContextName
	return c
}

// GerritAliasName sets the optional parameter "gerrit.aliasName": The
// name of an alias (branch, tag, etc.).
func (c *ProjectsReposWorkspacesFilesGetCall) GerritAliasName(gerritAliasName string) *ProjectsReposWorkspacesFilesGetCall {
	c.opt_["gerrit.aliasName"] = gerritAliasName
	return c
}

// GerritGerritProject sets the optional parameter
// "gerrit.gerritProject": The full project name within the host.
// Projects may be nested, so
// "project/subproject" is a valid project
// name.
// The "repo name" is hostURI/project.
func (c *ProjectsReposWorkspacesFilesGetCall) GerritGerritProject(gerritGerritProject string) *ProjectsReposWorkspacesFilesGetCall {
	c.opt_["gerrit.gerritProject"] = gerritGerritProject
	return c
}

// GerritHostUri sets the optional parameter "gerrit.hostUri": The URI
// of a running Gerrit instance.
func (c *ProjectsReposWorkspacesFilesGetCall) GerritHostUri(gerritHostUri string) *ProjectsReposWorkspacesFilesGetCall {
	c.opt_["gerrit.hostUri"] = gerritHostUri
	return c
}

// GerritRevisionId sets the optional parameter "gerrit.revisionId": A
// revision (commit) ID.
func (c *ProjectsReposWorkspacesFilesGetCall) GerritRevisionId(gerritRevisionId string) *ProjectsReposWorkspacesFilesGetCall {
	c.opt_["gerrit.revisionId"] = gerritRevisionId
	return c
}

// GitRevisionId sets the optional parameter "git.revisionId": Git
// commit hash.
// required.
func (c *ProjectsReposWorkspacesFilesGetCall) GitRevisionId(gitRevisionId string) *ProjectsReposWorkspacesFilesGetCall {
	c.opt_["git.revisionId"] = gitRevisionId
	return c
}

// GitUrl sets the optional parameter "git.url": Git repository URL.
func (c *ProjectsReposWorkspacesFilesGetCall) GitUrl(gitUrl string) *ProjectsReposWorkspacesFilesGetCall {
	c.opt_["git.url"] = gitUrl
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of values to return.
func (c *ProjectsReposWorkspacesFilesGetCall) PageSize(pageSize int64) *ProjectsReposWorkspacesFilesGetCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken": The value of
// next_page_token from the previous call.
// Omit for the first page, or
// if using start_index.
func (c *ProjectsReposWorkspacesFilesGetCall) PageToken(pageToken string) *ProjectsReposWorkspacesFilesGetCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// SnapshotId sets the optional parameter "snapshotId": The ID of the
// snapshot.
// An empty snapshot_id refers to the most recent snapshot.
func (c *ProjectsReposWorkspacesFilesGetCall) SnapshotId(snapshotId string) *ProjectsReposWorkspacesFilesGetCall {
	c.opt_["snapshotId"] = snapshotId
	return c
}

// StartPosition sets the optional parameter "startPosition": If path
// refers to a file, the position of the first byte of its contents
// to
// return. If path refers to a directory, the position of the first
// entry
// in the listing. If page_token is specified, this field is
// ignored.
func (c *ProjectsReposWorkspacesFilesGetCall) StartPosition(startPosition int64) *ProjectsReposWorkspacesFilesGetCall {
	c.opt_["startPosition"] = startPosition
	return c
}

// Uid sets the optional parameter "uid": A server-assigned, globally
// unique identifier.
func (c *ProjectsReposWorkspacesFilesGetCall) Uid(uid string) *ProjectsReposWorkspacesFilesGetCall {
	c.opt_["uid"] = uid
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsReposWorkspacesFilesGetCall) Fields(s ...googleapi.Field) *ProjectsReposWorkspacesFilesGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsReposWorkspacesFilesGetCall) Do() (*ReadResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["cloudRepo.aliasContext.kind"]; ok {
		params.Set("cloudRepo.aliasContext.kind", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["cloudRepo.aliasContext.name"]; ok {
		params.Set("cloudRepo.aliasContext.name", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["cloudRepo.aliasName"]; ok {
		params.Set("cloudRepo.aliasName", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["cloudRepo.repoId.projectRepoId.projectId"]; ok {
		params.Set("cloudRepo.repoId.projectRepoId.projectId", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["cloudRepo.repoId.projectRepoId.repoName"]; ok {
		params.Set("cloudRepo.repoId.projectRepoId.repoName", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["cloudRepo.repoId.uid"]; ok {
		params.Set("cloudRepo.repoId.uid", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["cloudRepo.revisionId"]; ok {
		params.Set("cloudRepo.revisionId", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["gerrit.aliasContext.kind"]; ok {
		params.Set("gerrit.aliasContext.kind", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["gerrit.aliasContext.name"]; ok {
		params.Set("gerrit.aliasContext.name", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["gerrit.aliasName"]; ok {
		params.Set("gerrit.aliasName", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["gerrit.gerritProject"]; ok {
		params.Set("gerrit.gerritProject", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["gerrit.hostUri"]; ok {
		params.Set("gerrit.hostUri", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["gerrit.revisionId"]; ok {
		params.Set("gerrit.revisionId", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["git.revisionId"]; ok {
		params.Set("git.revisionId", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["git.url"]; ok {
		params.Set("git.url", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageSize"]; ok {
		params.Set("pageSize", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["snapshotId"]; ok {
		params.Set("snapshotId", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["startPosition"]; ok {
		params.Set("startPosition", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["uid"]; ok {
		params.Set("uid", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/repos/{repoName}/workspaces/{name}/files/{+path}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"repoName":  c.repoName,
		"name":      c.name,
		"path":      c.path,
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
	var ret *ReadResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Read is given a SourceContext and path, and returns\nfile or directory information about that path.",
	//   "flatPath": "v1/projects/{projectId}/repos/{repoName}/workspaces/{name}/files/{filesId}",
	//   "httpMethod": "GET",
	//   "id": "source.projects.repos.workspaces.files.get",
	//   "parameterOrder": [
	//     "projectId",
	//     "repoName",
	//     "name",
	//     "path"
	//   ],
	//   "parameters": {
	//     "cloudRepo.aliasContext.kind": {
	//       "description": "The alias kind.",
	//       "enum": [
	//         "ANY",
	//         "FIXED",
	//         "MOVABLE",
	//         "OTHER"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "cloudRepo.aliasContext.name": {
	//       "description": "The alias name.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "cloudRepo.aliasName": {
	//       "description": "The name of an alias (branch, tag, etc.).",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "cloudRepo.repoId.projectRepoId.projectId": {
	//       "description": "The ID of the project.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "cloudRepo.repoId.projectRepoId.repoName": {
	//       "description": "The name of the repo. Leave empty for the default repo.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "cloudRepo.repoId.uid": {
	//       "description": "A server-assigned, globally unique identifier.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "cloudRepo.revisionId": {
	//       "description": "A revision ID.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "gerrit.aliasContext.kind": {
	//       "description": "The alias kind.",
	//       "enum": [
	//         "ANY",
	//         "FIXED",
	//         "MOVABLE",
	//         "OTHER"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "gerrit.aliasContext.name": {
	//       "description": "The alias name.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "gerrit.aliasName": {
	//       "description": "The name of an alias (branch, tag, etc.).",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "gerrit.gerritProject": {
	//       "description": "The full project name within the host. Projects may be nested, so\n\"project/subproject\" is a valid project name.\nThe \"repo name\" is hostURI/project.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "gerrit.hostUri": {
	//       "description": "The URI of a running Gerrit instance.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "gerrit.revisionId": {
	//       "description": "A revision (commit) ID.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "git.revisionId": {
	//       "description": "Git commit hash.\nrequired.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "git.url": {
	//       "description": "Git repository URL.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "name": {
	//       "description": "The unique name of the workspace within the repo.  This is the name\nchosen by the client in the Source API's CreateWorkspace method.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "The maximum number of values to return.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "The value of next_page_token from the previous call.\nOmit for the first page, or if using start_index.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "path": {
	//       "description": "Path to the file or directory from the root directory of the source\ncontext. It must not have leading or trailing slashes.",
	//       "location": "path",
	//       "pattern": "^.*$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "The ID of the project.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "repoName": {
	//       "description": "The name of the repo. Leave empty for the default repo.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "snapshotId": {
	//       "description": "The ID of the snapshot.\nAn empty snapshot_id refers to the most recent snapshot.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "startPosition": {
	//       "description": "If path refers to a file, the position of the first byte of its contents\nto return. If path refers to a directory, the position of the first entry\nin the listing. If page_token is specified, this field is ignored.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "uid": {
	//       "description": "A server-assigned, globally unique identifier.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/repos/{repoName}/workspaces/{name}/files/{+path}",
	//   "response": {
	//     "$ref": "ReadResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "source.projects.repos.workspaces.snapshots.get":

type ProjectsReposWorkspacesSnapshotsGetCall struct {
	s          *Service
	projectId  string
	repoName   string
	name       string
	snapshotId string
	opt_       map[string]interface{}
}

// Get: Gets a workspace snapshot.
func (r *ProjectsReposWorkspacesSnapshotsService) Get(projectId string, repoName string, name string, snapshotId string) *ProjectsReposWorkspacesSnapshotsGetCall {
	c := &ProjectsReposWorkspacesSnapshotsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.repoName = repoName
	c.name = name
	c.snapshotId = snapshotId
	return c
}

// Uid sets the optional parameter "uid": A server-assigned, globally
// unique identifier.
func (c *ProjectsReposWorkspacesSnapshotsGetCall) Uid(uid string) *ProjectsReposWorkspacesSnapshotsGetCall {
	c.opt_["uid"] = uid
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsReposWorkspacesSnapshotsGetCall) Fields(s ...googleapi.Field) *ProjectsReposWorkspacesSnapshotsGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsReposWorkspacesSnapshotsGetCall) Do() (*Snapshot, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["uid"]; ok {
		params.Set("uid", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/repos/{repoName}/workspaces/{name}/snapshots/{snapshotId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId":  c.projectId,
		"repoName":   c.repoName,
		"name":       c.name,
		"snapshotId": c.snapshotId,
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
	var ret *Snapshot
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets a workspace snapshot.",
	//   "flatPath": "v1/projects/{projectId}/repos/{repoName}/workspaces/{name}/snapshots/{snapshotId}",
	//   "httpMethod": "GET",
	//   "id": "source.projects.repos.workspaces.snapshots.get",
	//   "parameterOrder": [
	//     "projectId",
	//     "repoName",
	//     "name",
	//     "snapshotId"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The unique name of the workspace within the repo.  This is the name\nchosen by the client in the Source API's CreateWorkspace method.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "The ID of the project.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "repoName": {
	//       "description": "The name of the repo. Leave empty for the default repo.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "snapshotId": {
	//       "description": "The ID of the snapshot to get. If empty, the most recent snapshot is\nretrieved.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "uid": {
	//       "description": "A server-assigned, globally unique identifier.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/repos/{repoName}/workspaces/{name}/snapshots/{snapshotId}",
	//   "response": {
	//     "$ref": "Snapshot"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "source.projects.repos.workspaces.snapshots.list":

type ProjectsReposWorkspacesSnapshotsListCall struct {
	s         *Service
	projectId string
	repoName  string
	name      string
	opt_      map[string]interface{}
}

// List: Lists all the snapshots made to a workspace, sorted from most
// recent to
// least recent.
func (r *ProjectsReposWorkspacesSnapshotsService) List(projectId string, repoName string, name string) *ProjectsReposWorkspacesSnapshotsListCall {
	c := &ProjectsReposWorkspacesSnapshotsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.repoName = repoName
	c.name = name
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of values to return.
func (c *ProjectsReposWorkspacesSnapshotsListCall) PageSize(pageSize int64) *ProjectsReposWorkspacesSnapshotsListCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken": The value of
// next_page_token from the previous call.
// Omit for the first page.
func (c *ProjectsReposWorkspacesSnapshotsListCall) PageToken(pageToken string) *ProjectsReposWorkspacesSnapshotsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Uid sets the optional parameter "uid": A server-assigned, globally
// unique identifier.
func (c *ProjectsReposWorkspacesSnapshotsListCall) Uid(uid string) *ProjectsReposWorkspacesSnapshotsListCall {
	c.opt_["uid"] = uid
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsReposWorkspacesSnapshotsListCall) Fields(s ...googleapi.Field) *ProjectsReposWorkspacesSnapshotsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsReposWorkspacesSnapshotsListCall) Do() (*ListSnapshotsResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["pageSize"]; ok {
		params.Set("pageSize", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["uid"]; ok {
		params.Set("uid", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/repos/{repoName}/workspaces/{name}/snapshots")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"repoName":  c.repoName,
		"name":      c.name,
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
	var ret *ListSnapshotsResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists all the snapshots made to a workspace, sorted from most recent to\nleast recent.",
	//   "flatPath": "v1/projects/{projectId}/repos/{repoName}/workspaces/{name}/snapshots",
	//   "httpMethod": "GET",
	//   "id": "source.projects.repos.workspaces.snapshots.list",
	//   "parameterOrder": [
	//     "projectId",
	//     "repoName",
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The unique name of the workspace within the repo.  This is the name\nchosen by the client in the Source API's CreateWorkspace method.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "The maximum number of values to return.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The value of next_page_token from the previous call.\nOmit for the first page.",
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
	//       "description": "The name of the repo. Leave empty for the default repo.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "uid": {
	//       "description": "A server-assigned, globally unique identifier.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/repos/{repoName}/workspaces/{name}/snapshots",
	//   "response": {
	//     "$ref": "ListSnapshotsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "source.projects.repos.workspaces.snapshots.listFiles":

type ProjectsReposWorkspacesSnapshotsListFilesCall struct {
	s          *Service
	projectId  string
	repoName   string
	name       string
	snapshotId string
	opt_       map[string]interface{}
}

// ListFiles: ListFiles returns a list of all files in a SourceContext.
// The
// information about each file includes its path and its hash.
// The
// result is ordered by path. Pagination is supported.
func (r *ProjectsReposWorkspacesSnapshotsService) ListFiles(projectId string, repoName string, name string, snapshotId string) *ProjectsReposWorkspacesSnapshotsListFilesCall {
	c := &ProjectsReposWorkspacesSnapshotsListFilesCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.repoName = repoName
	c.name = name
	c.snapshotId = snapshotId
	return c
}

// CloudRepoAliasContextKind sets the optional parameter
// "cloudRepo.aliasContext.kind": The alias kind.
func (c *ProjectsReposWorkspacesSnapshotsListFilesCall) CloudRepoAliasContextKind(cloudRepoAliasContextKind string) *ProjectsReposWorkspacesSnapshotsListFilesCall {
	c.opt_["cloudRepo.aliasContext.kind"] = cloudRepoAliasContextKind
	return c
}

// CloudRepoAliasContextName sets the optional parameter
// "cloudRepo.aliasContext.name": The alias name.
func (c *ProjectsReposWorkspacesSnapshotsListFilesCall) CloudRepoAliasContextName(cloudRepoAliasContextName string) *ProjectsReposWorkspacesSnapshotsListFilesCall {
	c.opt_["cloudRepo.aliasContext.name"] = cloudRepoAliasContextName
	return c
}

// CloudRepoAliasName sets the optional parameter "cloudRepo.aliasName":
// The name of an alias (branch, tag, etc.).
func (c *ProjectsReposWorkspacesSnapshotsListFilesCall) CloudRepoAliasName(cloudRepoAliasName string) *ProjectsReposWorkspacesSnapshotsListFilesCall {
	c.opt_["cloudRepo.aliasName"] = cloudRepoAliasName
	return c
}

// CloudRepoRepoIdProjectRepoIdProjectId sets the optional parameter
// "cloudRepo.repoId.projectRepoId.projectId": The ID of the project.
func (c *ProjectsReposWorkspacesSnapshotsListFilesCall) CloudRepoRepoIdProjectRepoIdProjectId(cloudRepoRepoIdProjectRepoIdProjectId string) *ProjectsReposWorkspacesSnapshotsListFilesCall {
	c.opt_["cloudRepo.repoId.projectRepoId.projectId"] = cloudRepoRepoIdProjectRepoIdProjectId
	return c
}

// CloudRepoRepoIdProjectRepoIdRepoName sets the optional parameter
// "cloudRepo.repoId.projectRepoId.repoName": The name of the repo.
// Leave empty for the default repo.
func (c *ProjectsReposWorkspacesSnapshotsListFilesCall) CloudRepoRepoIdProjectRepoIdRepoName(cloudRepoRepoIdProjectRepoIdRepoName string) *ProjectsReposWorkspacesSnapshotsListFilesCall {
	c.opt_["cloudRepo.repoId.projectRepoId.repoName"] = cloudRepoRepoIdProjectRepoIdRepoName
	return c
}

// CloudRepoRepoIdUid sets the optional parameter
// "cloudRepo.repoId.uid": A server-assigned, globally unique
// identifier.
func (c *ProjectsReposWorkspacesSnapshotsListFilesCall) CloudRepoRepoIdUid(cloudRepoRepoIdUid string) *ProjectsReposWorkspacesSnapshotsListFilesCall {
	c.opt_["cloudRepo.repoId.uid"] = cloudRepoRepoIdUid
	return c
}

// CloudRepoRevisionId sets the optional parameter
// "cloudRepo.revisionId": A revision ID.
func (c *ProjectsReposWorkspacesSnapshotsListFilesCall) CloudRepoRevisionId(cloudRepoRevisionId string) *ProjectsReposWorkspacesSnapshotsListFilesCall {
	c.opt_["cloudRepo.revisionId"] = cloudRepoRevisionId
	return c
}

// GerritAliasContextKind sets the optional parameter
// "gerrit.aliasContext.kind": The alias kind.
func (c *ProjectsReposWorkspacesSnapshotsListFilesCall) GerritAliasContextKind(gerritAliasContextKind string) *ProjectsReposWorkspacesSnapshotsListFilesCall {
	c.opt_["gerrit.aliasContext.kind"] = gerritAliasContextKind
	return c
}

// GerritAliasContextName sets the optional parameter
// "gerrit.aliasContext.name": The alias name.
func (c *ProjectsReposWorkspacesSnapshotsListFilesCall) GerritAliasContextName(gerritAliasContextName string) *ProjectsReposWorkspacesSnapshotsListFilesCall {
	c.opt_["gerrit.aliasContext.name"] = gerritAliasContextName
	return c
}

// GerritAliasName sets the optional parameter "gerrit.aliasName": The
// name of an alias (branch, tag, etc.).
func (c *ProjectsReposWorkspacesSnapshotsListFilesCall) GerritAliasName(gerritAliasName string) *ProjectsReposWorkspacesSnapshotsListFilesCall {
	c.opt_["gerrit.aliasName"] = gerritAliasName
	return c
}

// GerritGerritProject sets the optional parameter
// "gerrit.gerritProject": The full project name within the host.
// Projects may be nested, so
// "project/subproject" is a valid project
// name.
// The "repo name" is hostURI/project.
func (c *ProjectsReposWorkspacesSnapshotsListFilesCall) GerritGerritProject(gerritGerritProject string) *ProjectsReposWorkspacesSnapshotsListFilesCall {
	c.opt_["gerrit.gerritProject"] = gerritGerritProject
	return c
}

// GerritHostUri sets the optional parameter "gerrit.hostUri": The URI
// of a running Gerrit instance.
func (c *ProjectsReposWorkspacesSnapshotsListFilesCall) GerritHostUri(gerritHostUri string) *ProjectsReposWorkspacesSnapshotsListFilesCall {
	c.opt_["gerrit.hostUri"] = gerritHostUri
	return c
}

// GerritRevisionId sets the optional parameter "gerrit.revisionId": A
// revision (commit) ID.
func (c *ProjectsReposWorkspacesSnapshotsListFilesCall) GerritRevisionId(gerritRevisionId string) *ProjectsReposWorkspacesSnapshotsListFilesCall {
	c.opt_["gerrit.revisionId"] = gerritRevisionId
	return c
}

// GitRevisionId sets the optional parameter "git.revisionId": Git
// commit hash.
// required.
func (c *ProjectsReposWorkspacesSnapshotsListFilesCall) GitRevisionId(gitRevisionId string) *ProjectsReposWorkspacesSnapshotsListFilesCall {
	c.opt_["git.revisionId"] = gitRevisionId
	return c
}

// GitUrl sets the optional parameter "git.url": Git repository URL.
func (c *ProjectsReposWorkspacesSnapshotsListFilesCall) GitUrl(gitUrl string) *ProjectsReposWorkspacesSnapshotsListFilesCall {
	c.opt_["git.url"] = gitUrl
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of values to return.
func (c *ProjectsReposWorkspacesSnapshotsListFilesCall) PageSize(pageSize int64) *ProjectsReposWorkspacesSnapshotsListFilesCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken": The value of
// next_page_token from the previous call.
// Omit for the first page.
func (c *ProjectsReposWorkspacesSnapshotsListFilesCall) PageToken(pageToken string) *ProjectsReposWorkspacesSnapshotsListFilesCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Uid sets the optional parameter "uid": A server-assigned, globally
// unique identifier.
func (c *ProjectsReposWorkspacesSnapshotsListFilesCall) Uid(uid string) *ProjectsReposWorkspacesSnapshotsListFilesCall {
	c.opt_["uid"] = uid
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsReposWorkspacesSnapshotsListFilesCall) Fields(s ...googleapi.Field) *ProjectsReposWorkspacesSnapshotsListFilesCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsReposWorkspacesSnapshotsListFilesCall) Do() (*ListFilesResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["cloudRepo.aliasContext.kind"]; ok {
		params.Set("cloudRepo.aliasContext.kind", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["cloudRepo.aliasContext.name"]; ok {
		params.Set("cloudRepo.aliasContext.name", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["cloudRepo.aliasName"]; ok {
		params.Set("cloudRepo.aliasName", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["cloudRepo.repoId.projectRepoId.projectId"]; ok {
		params.Set("cloudRepo.repoId.projectRepoId.projectId", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["cloudRepo.repoId.projectRepoId.repoName"]; ok {
		params.Set("cloudRepo.repoId.projectRepoId.repoName", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["cloudRepo.repoId.uid"]; ok {
		params.Set("cloudRepo.repoId.uid", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["cloudRepo.revisionId"]; ok {
		params.Set("cloudRepo.revisionId", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["gerrit.aliasContext.kind"]; ok {
		params.Set("gerrit.aliasContext.kind", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["gerrit.aliasContext.name"]; ok {
		params.Set("gerrit.aliasContext.name", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["gerrit.aliasName"]; ok {
		params.Set("gerrit.aliasName", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["gerrit.gerritProject"]; ok {
		params.Set("gerrit.gerritProject", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["gerrit.hostUri"]; ok {
		params.Set("gerrit.hostUri", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["gerrit.revisionId"]; ok {
		params.Set("gerrit.revisionId", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["git.revisionId"]; ok {
		params.Set("git.revisionId", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["git.url"]; ok {
		params.Set("git.url", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageSize"]; ok {
		params.Set("pageSize", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["uid"]; ok {
		params.Set("uid", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/repos/{repoName}/workspaces/{name}/snapshots/{snapshotId}:listFiles")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId":  c.projectId,
		"repoName":   c.repoName,
		"name":       c.name,
		"snapshotId": c.snapshotId,
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
	var ret *ListFilesResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "ListFiles returns a list of all files in a SourceContext. The\ninformation about each file includes its path and its hash.\nThe result is ordered by path. Pagination is supported.",
	//   "flatPath": "v1/projects/{projectId}/repos/{repoName}/workspaces/{name}/snapshots/{snapshotId}:listFiles",
	//   "httpMethod": "GET",
	//   "id": "source.projects.repos.workspaces.snapshots.listFiles",
	//   "parameterOrder": [
	//     "projectId",
	//     "repoName",
	//     "name",
	//     "snapshotId"
	//   ],
	//   "parameters": {
	//     "cloudRepo.aliasContext.kind": {
	//       "description": "The alias kind.",
	//       "enum": [
	//         "ANY",
	//         "FIXED",
	//         "MOVABLE",
	//         "OTHER"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "cloudRepo.aliasContext.name": {
	//       "description": "The alias name.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "cloudRepo.aliasName": {
	//       "description": "The name of an alias (branch, tag, etc.).",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "cloudRepo.repoId.projectRepoId.projectId": {
	//       "description": "The ID of the project.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "cloudRepo.repoId.projectRepoId.repoName": {
	//       "description": "The name of the repo. Leave empty for the default repo.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "cloudRepo.repoId.uid": {
	//       "description": "A server-assigned, globally unique identifier.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "cloudRepo.revisionId": {
	//       "description": "A revision ID.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "gerrit.aliasContext.kind": {
	//       "description": "The alias kind.",
	//       "enum": [
	//         "ANY",
	//         "FIXED",
	//         "MOVABLE",
	//         "OTHER"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "gerrit.aliasContext.name": {
	//       "description": "The alias name.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "gerrit.aliasName": {
	//       "description": "The name of an alias (branch, tag, etc.).",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "gerrit.gerritProject": {
	//       "description": "The full project name within the host. Projects may be nested, so\n\"project/subproject\" is a valid project name.\nThe \"repo name\" is hostURI/project.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "gerrit.hostUri": {
	//       "description": "The URI of a running Gerrit instance.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "gerrit.revisionId": {
	//       "description": "A revision (commit) ID.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "git.revisionId": {
	//       "description": "Git commit hash.\nrequired.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "git.url": {
	//       "description": "Git repository URL.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "name": {
	//       "description": "The unique name of the workspace within the repo.  This is the name\nchosen by the client in the Source API's CreateWorkspace method.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "The maximum number of values to return.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The value of next_page_token from the previous call.\nOmit for the first page.",
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
	//       "description": "The name of the repo. Leave empty for the default repo.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "snapshotId": {
	//       "description": "The ID of the snapshot.\nAn empty snapshot_id refers to the most recent snapshot.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "uid": {
	//       "description": "A server-assigned, globally unique identifier.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/repos/{repoName}/workspaces/{name}/snapshots/{snapshotId}:listFiles",
	//   "response": {
	//     "$ref": "ListFilesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "source.projects.repos.workspaces.snapshots.files.get":

type ProjectsReposWorkspacesSnapshotsFilesGetCall struct {
	s          *Service
	projectId  string
	repoName   string
	name       string
	snapshotId string
	path       string
	opt_       map[string]interface{}
}

// Get: Read is given a SourceContext and path, and returns
// file or
// directory information about that path.
func (r *ProjectsReposWorkspacesSnapshotsFilesService) Get(projectId string, repoName string, name string, snapshotId string, path string) *ProjectsReposWorkspacesSnapshotsFilesGetCall {
	c := &ProjectsReposWorkspacesSnapshotsFilesGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.repoName = repoName
	c.name = name
	c.snapshotId = snapshotId
	c.path = path
	return c
}

// CloudRepoAliasContextKind sets the optional parameter
// "cloudRepo.aliasContext.kind": The alias kind.
func (c *ProjectsReposWorkspacesSnapshotsFilesGetCall) CloudRepoAliasContextKind(cloudRepoAliasContextKind string) *ProjectsReposWorkspacesSnapshotsFilesGetCall {
	c.opt_["cloudRepo.aliasContext.kind"] = cloudRepoAliasContextKind
	return c
}

// CloudRepoAliasContextName sets the optional parameter
// "cloudRepo.aliasContext.name": The alias name.
func (c *ProjectsReposWorkspacesSnapshotsFilesGetCall) CloudRepoAliasContextName(cloudRepoAliasContextName string) *ProjectsReposWorkspacesSnapshotsFilesGetCall {
	c.opt_["cloudRepo.aliasContext.name"] = cloudRepoAliasContextName
	return c
}

// CloudRepoAliasName sets the optional parameter "cloudRepo.aliasName":
// The name of an alias (branch, tag, etc.).
func (c *ProjectsReposWorkspacesSnapshotsFilesGetCall) CloudRepoAliasName(cloudRepoAliasName string) *ProjectsReposWorkspacesSnapshotsFilesGetCall {
	c.opt_["cloudRepo.aliasName"] = cloudRepoAliasName
	return c
}

// CloudRepoRepoIdProjectRepoIdProjectId sets the optional parameter
// "cloudRepo.repoId.projectRepoId.projectId": The ID of the project.
func (c *ProjectsReposWorkspacesSnapshotsFilesGetCall) CloudRepoRepoIdProjectRepoIdProjectId(cloudRepoRepoIdProjectRepoIdProjectId string) *ProjectsReposWorkspacesSnapshotsFilesGetCall {
	c.opt_["cloudRepo.repoId.projectRepoId.projectId"] = cloudRepoRepoIdProjectRepoIdProjectId
	return c
}

// CloudRepoRepoIdProjectRepoIdRepoName sets the optional parameter
// "cloudRepo.repoId.projectRepoId.repoName": The name of the repo.
// Leave empty for the default repo.
func (c *ProjectsReposWorkspacesSnapshotsFilesGetCall) CloudRepoRepoIdProjectRepoIdRepoName(cloudRepoRepoIdProjectRepoIdRepoName string) *ProjectsReposWorkspacesSnapshotsFilesGetCall {
	c.opt_["cloudRepo.repoId.projectRepoId.repoName"] = cloudRepoRepoIdProjectRepoIdRepoName
	return c
}

// CloudRepoRepoIdUid sets the optional parameter
// "cloudRepo.repoId.uid": A server-assigned, globally unique
// identifier.
func (c *ProjectsReposWorkspacesSnapshotsFilesGetCall) CloudRepoRepoIdUid(cloudRepoRepoIdUid string) *ProjectsReposWorkspacesSnapshotsFilesGetCall {
	c.opt_["cloudRepo.repoId.uid"] = cloudRepoRepoIdUid
	return c
}

// CloudRepoRevisionId sets the optional parameter
// "cloudRepo.revisionId": A revision ID.
func (c *ProjectsReposWorkspacesSnapshotsFilesGetCall) CloudRepoRevisionId(cloudRepoRevisionId string) *ProjectsReposWorkspacesSnapshotsFilesGetCall {
	c.opt_["cloudRepo.revisionId"] = cloudRepoRevisionId
	return c
}

// GerritAliasContextKind sets the optional parameter
// "gerrit.aliasContext.kind": The alias kind.
func (c *ProjectsReposWorkspacesSnapshotsFilesGetCall) GerritAliasContextKind(gerritAliasContextKind string) *ProjectsReposWorkspacesSnapshotsFilesGetCall {
	c.opt_["gerrit.aliasContext.kind"] = gerritAliasContextKind
	return c
}

// GerritAliasContextName sets the optional parameter
// "gerrit.aliasContext.name": The alias name.
func (c *ProjectsReposWorkspacesSnapshotsFilesGetCall) GerritAliasContextName(gerritAliasContextName string) *ProjectsReposWorkspacesSnapshotsFilesGetCall {
	c.opt_["gerrit.aliasContext.name"] = gerritAliasContextName
	return c
}

// GerritAliasName sets the optional parameter "gerrit.aliasName": The
// name of an alias (branch, tag, etc.).
func (c *ProjectsReposWorkspacesSnapshotsFilesGetCall) GerritAliasName(gerritAliasName string) *ProjectsReposWorkspacesSnapshotsFilesGetCall {
	c.opt_["gerrit.aliasName"] = gerritAliasName
	return c
}

// GerritGerritProject sets the optional parameter
// "gerrit.gerritProject": The full project name within the host.
// Projects may be nested, so
// "project/subproject" is a valid project
// name.
// The "repo name" is hostURI/project.
func (c *ProjectsReposWorkspacesSnapshotsFilesGetCall) GerritGerritProject(gerritGerritProject string) *ProjectsReposWorkspacesSnapshotsFilesGetCall {
	c.opt_["gerrit.gerritProject"] = gerritGerritProject
	return c
}

// GerritHostUri sets the optional parameter "gerrit.hostUri": The URI
// of a running Gerrit instance.
func (c *ProjectsReposWorkspacesSnapshotsFilesGetCall) GerritHostUri(gerritHostUri string) *ProjectsReposWorkspacesSnapshotsFilesGetCall {
	c.opt_["gerrit.hostUri"] = gerritHostUri
	return c
}

// GerritRevisionId sets the optional parameter "gerrit.revisionId": A
// revision (commit) ID.
func (c *ProjectsReposWorkspacesSnapshotsFilesGetCall) GerritRevisionId(gerritRevisionId string) *ProjectsReposWorkspacesSnapshotsFilesGetCall {
	c.opt_["gerrit.revisionId"] = gerritRevisionId
	return c
}

// GitRevisionId sets the optional parameter "git.revisionId": Git
// commit hash.
// required.
func (c *ProjectsReposWorkspacesSnapshotsFilesGetCall) GitRevisionId(gitRevisionId string) *ProjectsReposWorkspacesSnapshotsFilesGetCall {
	c.opt_["git.revisionId"] = gitRevisionId
	return c
}

// GitUrl sets the optional parameter "git.url": Git repository URL.
func (c *ProjectsReposWorkspacesSnapshotsFilesGetCall) GitUrl(gitUrl string) *ProjectsReposWorkspacesSnapshotsFilesGetCall {
	c.opt_["git.url"] = gitUrl
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of values to return.
func (c *ProjectsReposWorkspacesSnapshotsFilesGetCall) PageSize(pageSize int64) *ProjectsReposWorkspacesSnapshotsFilesGetCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken": The value of
// next_page_token from the previous call.
// Omit for the first page, or
// if using start_index.
func (c *ProjectsReposWorkspacesSnapshotsFilesGetCall) PageToken(pageToken string) *ProjectsReposWorkspacesSnapshotsFilesGetCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// StartPosition sets the optional parameter "startPosition": If path
// refers to a file, the position of the first byte of its contents
// to
// return. If path refers to a directory, the position of the first
// entry
// in the listing. If page_token is specified, this field is
// ignored.
func (c *ProjectsReposWorkspacesSnapshotsFilesGetCall) StartPosition(startPosition int64) *ProjectsReposWorkspacesSnapshotsFilesGetCall {
	c.opt_["startPosition"] = startPosition
	return c
}

// Uid sets the optional parameter "uid": A server-assigned, globally
// unique identifier.
func (c *ProjectsReposWorkspacesSnapshotsFilesGetCall) Uid(uid string) *ProjectsReposWorkspacesSnapshotsFilesGetCall {
	c.opt_["uid"] = uid
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsReposWorkspacesSnapshotsFilesGetCall) Fields(s ...googleapi.Field) *ProjectsReposWorkspacesSnapshotsFilesGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsReposWorkspacesSnapshotsFilesGetCall) Do() (*ReadResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["cloudRepo.aliasContext.kind"]; ok {
		params.Set("cloudRepo.aliasContext.kind", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["cloudRepo.aliasContext.name"]; ok {
		params.Set("cloudRepo.aliasContext.name", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["cloudRepo.aliasName"]; ok {
		params.Set("cloudRepo.aliasName", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["cloudRepo.repoId.projectRepoId.projectId"]; ok {
		params.Set("cloudRepo.repoId.projectRepoId.projectId", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["cloudRepo.repoId.projectRepoId.repoName"]; ok {
		params.Set("cloudRepo.repoId.projectRepoId.repoName", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["cloudRepo.repoId.uid"]; ok {
		params.Set("cloudRepo.repoId.uid", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["cloudRepo.revisionId"]; ok {
		params.Set("cloudRepo.revisionId", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["gerrit.aliasContext.kind"]; ok {
		params.Set("gerrit.aliasContext.kind", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["gerrit.aliasContext.name"]; ok {
		params.Set("gerrit.aliasContext.name", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["gerrit.aliasName"]; ok {
		params.Set("gerrit.aliasName", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["gerrit.gerritProject"]; ok {
		params.Set("gerrit.gerritProject", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["gerrit.hostUri"]; ok {
		params.Set("gerrit.hostUri", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["gerrit.revisionId"]; ok {
		params.Set("gerrit.revisionId", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["git.revisionId"]; ok {
		params.Set("git.revisionId", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["git.url"]; ok {
		params.Set("git.url", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageSize"]; ok {
		params.Set("pageSize", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["startPosition"]; ok {
		params.Set("startPosition", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["uid"]; ok {
		params.Set("uid", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/projects/{projectId}/repos/{repoName}/workspaces/{name}/snapshots/{snapshotId}/files/{+path}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId":  c.projectId,
		"repoName":   c.repoName,
		"name":       c.name,
		"snapshotId": c.snapshotId,
		"path":       c.path,
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
	var ret *ReadResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Read is given a SourceContext and path, and returns\nfile or directory information about that path.",
	//   "flatPath": "v1/projects/{projectId}/repos/{repoName}/workspaces/{name}/snapshots/{snapshotId}/files/{filesId}",
	//   "httpMethod": "GET",
	//   "id": "source.projects.repos.workspaces.snapshots.files.get",
	//   "parameterOrder": [
	//     "projectId",
	//     "repoName",
	//     "name",
	//     "snapshotId",
	//     "path"
	//   ],
	//   "parameters": {
	//     "cloudRepo.aliasContext.kind": {
	//       "description": "The alias kind.",
	//       "enum": [
	//         "ANY",
	//         "FIXED",
	//         "MOVABLE",
	//         "OTHER"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "cloudRepo.aliasContext.name": {
	//       "description": "The alias name.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "cloudRepo.aliasName": {
	//       "description": "The name of an alias (branch, tag, etc.).",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "cloudRepo.repoId.projectRepoId.projectId": {
	//       "description": "The ID of the project.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "cloudRepo.repoId.projectRepoId.repoName": {
	//       "description": "The name of the repo. Leave empty for the default repo.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "cloudRepo.repoId.uid": {
	//       "description": "A server-assigned, globally unique identifier.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "cloudRepo.revisionId": {
	//       "description": "A revision ID.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "gerrit.aliasContext.kind": {
	//       "description": "The alias kind.",
	//       "enum": [
	//         "ANY",
	//         "FIXED",
	//         "MOVABLE",
	//         "OTHER"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "gerrit.aliasContext.name": {
	//       "description": "The alias name.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "gerrit.aliasName": {
	//       "description": "The name of an alias (branch, tag, etc.).",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "gerrit.gerritProject": {
	//       "description": "The full project name within the host. Projects may be nested, so\n\"project/subproject\" is a valid project name.\nThe \"repo name\" is hostURI/project.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "gerrit.hostUri": {
	//       "description": "The URI of a running Gerrit instance.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "gerrit.revisionId": {
	//       "description": "A revision (commit) ID.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "git.revisionId": {
	//       "description": "Git commit hash.\nrequired.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "git.url": {
	//       "description": "Git repository URL.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "name": {
	//       "description": "The unique name of the workspace within the repo.  This is the name\nchosen by the client in the Source API's CreateWorkspace method.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "The maximum number of values to return.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "The value of next_page_token from the previous call.\nOmit for the first page, or if using start_index.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "path": {
	//       "description": "Path to the file or directory from the root directory of the source\ncontext. It must not have leading or trailing slashes.",
	//       "location": "path",
	//       "pattern": "^.*$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "The ID of the project.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "repoName": {
	//       "description": "The name of the repo. Leave empty for the default repo.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "snapshotId": {
	//       "description": "The ID of the snapshot.\nAn empty snapshot_id refers to the most recent snapshot.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "startPosition": {
	//       "description": "If path refers to a file, the position of the first byte of its contents\nto return. If path refers to a directory, the position of the first entry\nin the listing. If page_token is specified, this field is ignored.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "uid": {
	//       "description": "A server-assigned, globally unique identifier.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/projects/{projectId}/repos/{repoName}/workspaces/{name}/snapshots/{snapshotId}/files/{+path}",
	//   "response": {
	//     "$ref": "ReadResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "source.listChangedFiles":

type V1ListChangedFilesCall struct {
	s                       *Service
	listchangedfilesrequest *ListChangedFilesRequest
	opt_                    map[string]interface{}
}

// ListChangedFiles: ListChangedFiles computes the files that have
// changed between two revisions
// or workspace snapshots in the same
// repo. It returns a list of
// ChangeFileInfos.
//
// ListChangedFiles does
// not perform copy/rename detection, so the from_path of
// ChangeFileInfo
// is unset. Examine the changed_files field of the Revision
// resource to
// determine copy/rename information.
//
// The result is ordered by path.
// Pagination is supported.
func (r *V1Service) ListChangedFiles(listchangedfilesrequest *ListChangedFilesRequest) *V1ListChangedFilesCall {
	c := &V1ListChangedFilesCall{s: r.s, opt_: make(map[string]interface{})}
	c.listchangedfilesrequest = listchangedfilesrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *V1ListChangedFilesCall) Fields(s ...googleapi.Field) *V1ListChangedFilesCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *V1ListChangedFilesCall) Do() (*ListChangedFilesResponse, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.listchangedfilesrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1:listChangedFiles")
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
	var ret *ListChangedFilesResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "ListChangedFiles computes the files that have changed between two revisions\nor workspace snapshots in the same repo. It returns a list of\nChangeFileInfos.\n\nListChangedFiles does not perform copy/rename detection, so the from_path of\nChangeFileInfo is unset. Examine the changed_files field of the Revision\nresource to determine copy/rename information.\n\nThe result is ordered by path. Pagination is supported.",
	//   "flatPath": "v1:listChangedFiles",
	//   "httpMethod": "POST",
	//   "id": "source.listChangedFiles",
	//   "parameterOrder": [],
	//   "parameters": {},
	//   "path": "v1:listChangedFiles",
	//   "request": {
	//     "$ref": "ListChangedFilesRequest"
	//   },
	//   "response": {
	//     "$ref": "ListChangedFilesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}
