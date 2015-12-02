// Package dataproc provides access to the Google Cloud Dataproc API.
//
// See https://cloud.google.com/dataproc/
//
// Usage example:
//
//   import "google.golang.org/api/dataproc/v1beta1"
//   ...
//   dataprocService, err := dataproc.New(oauthHttpClient)
package dataproc

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

const apiId = "dataproc:v1beta1"
const apiName = "dataproc"
const apiVersion = "v1beta1"
const basePath = "https://dataproc.googleapis.com/"

// OAuth2 scopes used by this API.
const (
	// View and manage your data across Google Cloud Platform services
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"

	// Administrate log data for your projects
	LoggingAdminScope = "https://www.googleapis.com/auth/logging.admin"

	// View log data for your projects
	LoggingReadScope = "https://www.googleapis.com/auth/logging.read"

	// Submit log data for your projects
	LoggingWriteScope = "https://www.googleapis.com/auth/logging.write"
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
	rs.Clusters = NewProjectsClustersService(s)
	rs.Jobs = NewProjectsJobsService(s)
	return rs
}

type ProjectsService struct {
	s *Service

	Clusters *ProjectsClustersService

	Jobs *ProjectsJobsService
}

func NewProjectsClustersService(s *Service) *ProjectsClustersService {
	rs := &ProjectsClustersService{s: s}
	rs.Agents = NewProjectsClustersAgentsService(s)
	rs.Tasks = NewProjectsClustersTasksService(s)
	return rs
}

type ProjectsClustersService struct {
	s *Service

	Agents *ProjectsClustersAgentsService

	Tasks *ProjectsClustersTasksService
}

func NewProjectsClustersAgentsService(s *Service) *ProjectsClustersAgentsService {
	rs := &ProjectsClustersAgentsService{s: s}
	return rs
}

type ProjectsClustersAgentsService struct {
	s *Service
}

func NewProjectsClustersTasksService(s *Service) *ProjectsClustersTasksService {
	rs := &ProjectsClustersTasksService{s: s}
	return rs
}

type ProjectsClustersTasksService struct {
	s *Service
}

func NewProjectsJobsService(s *Service) *ProjectsJobsService {
	rs := &ProjectsJobsService{s: s}
	return rs
}

type ProjectsJobsService struct {
	s *Service
}

type Agent struct {
	// AgentId: [Required] 64 characters matching the regular expression:
	// [a-z0-9.-]{1,64}
	// An agent chosen ID. This should typically be the
	// hostname of the GCE
	// virtual machine on which the agent is currently
	// running.
	AgentId string `json:"agentId,omitempty"`

	// AgentVersion: The version of this agent in HTTP User-Agent Header
	// value format (RFC 2616
	// section 14.43), e.g., "Dataproc-Agent/1.2".
	AgentVersion string `json:"agentVersion,omitempty"`

	// Diagnostic: [Optional] The list of diagnostic records since last
	// agent updage.
	Diagnostic []*DiagnosticRecord `json:"diagnostic,omitempty"`

	// LastAgentUpdateTime: [Out] the last time this agent checked-in with
	// Dataproc.
	LastAgentUpdateTime string `json:"lastAgentUpdateTime,omitempty"`

	// Status: Agent status.
	Status string `json:"status,omitempty"`
}

type CancelJobRequest struct {
}

type CancelOperationRequest struct {
}

type Cluster struct {
	// ClusterName: [Required] The cluster name. Cluster names within a
	// project must be
	// unique. Names from deleted clusters can be reused.
	ClusterName string `json:"clusterName,omitempty"`

	// ClusterUuid: [Output-only] A cluster UUID (Unique Universal
	// Identifier). Cloud Dataproc
	// generates this value when it creates the
	// cluster.
	ClusterUuid string `json:"clusterUuid,omitempty"`

	// Configuration: [Required] The cluster configuration. Note that Cloud
	// Dataproc may set
	// default values, and values may change when clusters
	// are updated.
	Configuration *ClusterConfiguration `json:"configuration,omitempty"`

	// ProjectId: [Required] The Google Cloud Platform project ID that the
	// cluster belongs to.
	ProjectId string `json:"projectId,omitempty"`

	// Status: [Output-only] Cluster status.
	Status *ClusterStatus `json:"status,omitempty"`

	// StatusHistory: [Output-only] Previous cluster statuses.
	StatusHistory []*ClusterStatus `json:"statusHistory,omitempty"`
}

type ClusterConfiguration struct {
	// ConfigurationBucket: [Optional] A Google Cloud Storage staging bucket
	// used for sharing generated
	// SSH keys and configuration. If you do not
	// specify a staging bucket, Cloud
	// Dataproc will determine an
	// appropriate Cloud Storage location (US,
	// ASIA, or EU) for your
	// cluster's staging bucket according to the Google
	// Compute Engine zone
	// where your cluster is deployed, and then it will create
	// and manage
	// this project-level, per-location bucket for you.
	ConfigurationBucket string `json:"configurationBucket,omitempty"`

	// GceClusterConfiguration: [Optional] The shared Google Compute Engine
	// configuration settings for
	// all instances in a cluster.
	GceClusterConfiguration *GceClusterConfiguration `json:"gceClusterConfiguration,omitempty"`

	// InitializationActions: [Optional] Commands to execute on each node
	// after configuration is
	// completed. By default, executables are run on
	// master and all worker nodes.
	// You can test a node's <code>role</code>
	// metadata to run an executable on
	// a master or worker node, as shown
	// below:
	//
	//     ROLE=$(/usr/share/google/get_metadata_value
	// attributes/role)
	//     if [[ "${ROLE}" == 'Master' ]]; then
	//       ...
	// master specific actions ...
	//     else
	//       ... worker specific
	// actions ...
	//     fi
	InitializationActions []*NodeInitializationAction `json:"initializationActions,omitempty"`

	// MasterConfiguration: [Optional] The Google Compute Engine
	// configuration settings for
	// the master instance in a cluster.
	MasterConfiguration *InstanceGroupConfiguration `json:"masterConfiguration,omitempty"`

	// SecondaryWorkerConfiguration: [Optional] The Google Compute Engine
	// configuration settings for
	// additional worker instances in a cluster.
	SecondaryWorkerConfiguration *InstanceGroupConfiguration `json:"secondaryWorkerConfiguration,omitempty"`

	// SoftwareConfiguration: [Optional] The configuration settings for
	// software inside the cluster.
	SoftwareConfiguration *SoftwareConfiguration `json:"softwareConfiguration,omitempty"`

	// WorkerConfiguration: [Optional] The Google Compute Engine
	// configuration settings for
	// worker instances in a cluster.
	WorkerConfiguration *InstanceGroupConfiguration `json:"workerConfiguration,omitempty"`
}

type ClusterStatus struct {
	// Detail: Optional details of cluster's state.
	Detail string `json:"detail,omitempty"`

	// State: The cluster's state.
	State string `json:"state,omitempty"`

	// StateStartTime: Time when this state was entered.
	StateStartTime string `json:"stateStartTime,omitempty"`
}

type DiagnoseClusterOutputLocation struct {
	// OutputUri: [Output-only] The Google Cloud Storage URI of the
	// diagnostic output.
	// This will be a plain text file with summary of
	// collected diagnostics.
	OutputUri string `json:"outputUri,omitempty"`
}

type DiagnoseClusterRequest struct {
}

type DiagnosticRecord struct {
	// ErrorSample: A complete stack trace serialized to string from an
	// exception
	// encountered by the agent.
	ErrorSample string `json:"errorSample,omitempty"`
}

type DiskConfiguration struct {
	// BootDiskSizeGb: [Optional] Size in GB of the boot disk (default is
	// 500GB).
	BootDiskSizeGb int64 `json:"bootDiskSizeGb,omitempty"`

	// NumLocalSsds: [Optional] Number of attached SSDs, from 0 to 4
	// (default is 0).
	// If SSDs are not attached, the boot disk is used to
	// store runtime logs and
	// HDFS data. If one or more SSDs are attached,
	// this runtime bulk
	// data is spread across them, and the boot disk
	// contains only basic
	// configuration and installed binaries.
	NumLocalSsds int64 `json:"numLocalSsds,omitempty"`
}

type Empty struct {
}

type GceClusterConfiguration struct {
	// NetworkUri: The Google Compute Engine network to be used for machine
	// communications.
	// Inbound SSH connections are necessary to complete
	// cluster configuration.
	// Example:
	// `compute.googleapis.com/projects/[project_id]/zones/us-east1-a/default
	// `.
	// This should follow the instructions for full resource
	// names found
	// here:
	// https://engdoc.corp.google.com/eng/doc/ti-apis/
	// style/resource_names.s
	// html?cl=head
	NetworkUri string `json:"networkUri,omitempty"`

	// ServiceAccountScopes: The URIs of service account scopes to be
	// included in Google Compute Engine
	// instances. The following base set
	// of scopes is always included:
	// -
	// https://www.googleapis.com/auth/cloud.useraccounts.readonly
	// -
	// https://www.googleapis.com/auth/devstorage.read_write
	// -
	// https://www.googleapis.com/auth/logging.write
	// If no scopes are
	// specfied, the following defaults are also provided:
	// -
	// https://www.googleapis.com/auth/bigquery
	// -
	// https://www.googleapis.com/auth/bigtable.admin.table
	// -
	// https://www.googleapis.com/auth/bigtable.data
	// -
	// https://www.googleapis.com/auth/devstorage.full_control
	ServiceAccountScopes []string `json:"serviceAccountScopes,omitempty"`

	// ZoneUri: [Required] The zone where the Google Compute Engine cluster
	// will be located.
	// Example:
	// `compute.googleapis.com/projects/[project_id]/zones/us-east1-a`.
	// This
	// should follow the instructions for full resource
	// names found here:
	// https://engdoc.corp.google.com/eng/doc/ti-apis/
	// style/resource_names.s
	// html?cl=head
	ZoneUri string `json:"zoneUri,omitempty"`
}

type HadoopJob struct {
	// ArchiveUris: [Optional] HCFS URIs of archives to be extracted in the
	// working directory of
	// Hadoop drivers and tasks. Supported file
	// types:
	// .jar, .tar, .tar.gz, .tgz, or .zip.
	ArchiveUris []string `json:"archiveUris,omitempty"`

	// Args: [Optional] The arguments to pass to the driver. Do not
	// include
	// arguments, such as `-libjars` or `-Dfoo=bar`, that can be set as
	// job
	// properties, since a collision may occur that causes an incorrect
	// job
	// submission.
	Args []string `json:"args,omitempty"`

	// FileUris: [Optional] HCFS URIs of files to be copied to the
	// working
	// directory of Hadoop drivers and distributed tasks. Useful
	// for
	// naively parallel tasks.
	FileUris []string `json:"fileUris,omitempty"`

	// JarFileUris: [Optional] Jar file URIs to add to the CLASSPATHs of
	// the
	// Hadoop driver and tasks.
	JarFileUris []string `json:"jarFileUris,omitempty"`

	// LoggingConfiguration: [Optional] The runtime log configuration for
	// job execution.
	LoggingConfiguration *LoggingConfiguration `json:"loggingConfiguration,omitempty"`

	// MainClass: The name of the driver's main class. The jar file
	// containing the class
	// must be in the default CLASSPATH or specified in
	// `jar_file_uris`.
	MainClass string `json:"mainClass,omitempty"`

	// MainJarFileUri: The Hadoop Compatible Filesystem (HCFS) URI of the
	// jar file containing
	// the main class.
	// Examples:
	//
	// 'gs://foo-bucket/analytics-binaries/extract-useful-metrics-mr.jar'
	//
	//  'hdfs:/tmp/test-samples/custom-wordcount.jar'
	//
	// 'file:///home/usr/lib/hadoop-mapreduce/hadoop-mapreduce-examples.jar'
	MainJarFileUri string `json:"mainJarFileUri,omitempty"`

	// Properties: [Optional] A mapping of property names to values, used to
	// configure Hadoop.
	// Properties that conflict with values set by the
	// Cloud Dataproc API may be
	// overwritten. Can include properties set in
	// /etc/hadoop/conf/*-site and
	// classes in user code.
	Properties map[string]string `json:"properties,omitempty"`
}

type HiveJob struct {
	// ContinueOnFailure: [Optional] Whether to continue executing queries
	// if a query fails.
	// The default value is `false`. Setting to `true` can
	// be useful when executing
	// independent parallel queries.
	ContinueOnFailure bool `json:"continueOnFailure,omitempty"`

	// JarFileUris: [Optional] HCFS URIs of jar files to add to the
	// CLASSPATH of the
	// Hive server and Hadoop MapReduce (MR) tasks. Can
	// contain Hive SerDes
	// and UDFs.
	JarFileUris []string `json:"jarFileUris,omitempty"`

	// Properties: [Optional] A mapping of property names and values, used
	// to configure Hive.
	// Properties that conflict with values set by the
	// Cloud Dataproc API may be
	// overwritten. Can include properties set in
	// /etc/hadoop/conf/*-site.xml,
	// /etc/hive/conf/hive-site.xml, and
	// classes in user code.
	Properties map[string]string `json:"properties,omitempty"`

	// QueryFileUri: The HCFS URI of the script that contains Hive queries.
	QueryFileUri string `json:"queryFileUri,omitempty"`

	// QueryList: A list of queries.
	QueryList *QueryList `json:"queryList,omitempty"`

	// ScriptVariables: [Optional] Mapping of query variable names to values
	// (equivalent to the
	// Hive command: `SET name="value";`).
	ScriptVariables map[string]string `json:"scriptVariables,omitempty"`
}

type InstanceGroupConfiguration struct {
	// DiskConfiguration: Disk option configuration settings.
	DiskConfiguration *DiskConfiguration `json:"diskConfiguration,omitempty"`

	// ImageUri: [Output-only] The Google Compute Engine image resource used
	// for cluster
	// instances. Inferred from
	// `SoftwareConfiguration.image_version`.
	// Can actually be specified for
	// development purposes
	// Example:
	// `compute.googleapis.com/projects/debian-cloud/global/images/backports-
	// debian-7-wheezy-v20140904`.
	// This should follow the instructions for
	// full resource
	// names found here:
	// https://engdoc.corp.google.com/eng/doc/ti-apis/
	// style/resource_names.s
	// html?cl=head
	ImageUri string `json:"imageUri,omitempty"`

	// InstanceNames: The list of instance names. Dataproc derives the names
	// from
	// `cluster_name`, `num_instances`, and the instance group if not
	// set by user
	// (recommended practice is to let Dataproc derive the
	// name).
	InstanceNames []string `json:"instanceNames,omitempty"`

	// IsPreemptible: Specifies that this instance group contains
	// Preemptible Instances.
	IsPreemptible bool `json:"isPreemptible,omitempty"`

	// MachineTypeUri: The Google Compute Engine machine type used for
	// cluster instances.
	// Example:
	// `compute.googleapis.com/projects/[project_id]/zones/us-east1-a/machine
	// Types/n1-standard-2`.
	// This should follow the instructions for full
	// resource
	// names found here:
	// https://engdoc.corp.google.com/eng/doc/ti-apis/
	// style/resource_names.s
	// html?cl=head
	MachineTypeUri string `json:"machineTypeUri,omitempty"`

	// ManagedGroupConfiguration: [Output-only] The configuration for Google
	// Compute Engine Instance Group
	// Manager that manages this group.
	// This
	// is only used for preemptible instance groups.
	ManagedGroupConfiguration *ManagedGroupConfiguration `json:"managedGroupConfiguration,omitempty"`

	// NumInstances: The number of VM instances in the instance group.
	// For
	// master instance groups, must be set to 1.
	// eventually support high
	// availability with multiple masters.
	NumInstances int64 `json:"numInstances,omitempty"`
}

type Job struct {
	// DriverControlFilesUri: [Output-only] If present, the location of
	// miscellaneous control files
	// which may be used as part of job setup
	// and handling. If not present,
	// control files may be placed in the same
	// location as `driver_output_uri`.
	DriverControlFilesUri string `json:"driverControlFilesUri,omitempty"`

	// DriverInputResourceUri: [Output-only] A URI pointing to the location
	// of the stdin of the job's
	// driver program, only set if the job is
	// interactive.
	DriverInputResourceUri string `json:"driverInputResourceUri,omitempty"`

	// DriverOutputResourceUri: [Output-only] A URI pointing to the location
	// of the stdout of the job's
	// driver program.
	DriverOutputResourceUri string `json:"driverOutputResourceUri,omitempty"`

	// HadoopJob: Job is a Hadoop job.
	HadoopJob *HadoopJob `json:"hadoopJob,omitempty"`

	// HiveJob: Job is a Hive job.
	HiveJob *HiveJob `json:"hiveJob,omitempty"`

	// Interactive: [Optional] If set to `true`, the driver's stdin will be
	// kept open and
	// `driver_input_uri` will be set to provide a path at
	// which additional input
	// can be sent to the driver.
	Interactive bool `json:"interactive,omitempty"`

	// PigJob: Job is a Pig job.
	PigJob *PigJob `json:"pigJob,omitempty"`

	// Placement: [Required] Job information, including how, when, and where
	// to
	// run the job.
	Placement *JobPlacement `json:"placement,omitempty"`

	// PysparkJob: Job is a Pyspark job.
	PysparkJob *PySparkJob `json:"pysparkJob,omitempty"`

	// Reference: [Optional] The fully qualified reference to the job, which
	// can be used to
	// obtain the equivalent REST path of the job resource.
	// If this property
	// is not specified when a job is created, the server
	// generates a
	// <code>job_id</code>.
	Reference *JobReference `json:"reference,omitempty"`

	// SparkJob: Job is a Spark job.
	SparkJob *SparkJob `json:"sparkJob,omitempty"`

	// SparkSqlJob: Job is a SparkSql job.
	SparkSqlJob *SparkSqlJob `json:"sparkSqlJob,omitempty"`

	// Status: [Output-only] The job status. Additional
	// application-specific
	// status information may be contained in the
	// <code>type_job</code>
	// and <code>yarn_applications</code> fields.
	Status *JobStatus `json:"status,omitempty"`

	// StatusHistory: [Output-only] The previous job status.
	StatusHistory []*JobStatus `json:"statusHistory,omitempty"`

	// SubmittedBy: [Output-only] The email address of the user submitting
	// the job. For jobs
	// submitted on the cluster, the address is
	// <code>username@hostname</code>.
	SubmittedBy string `json:"submittedBy,omitempty"`

	// YarnApplications: [Output-only] The collection of YARN applications
	// spun up by this job.
	YarnApplications []*YarnApplication `json:"yarnApplications,omitempty"`
}

type JobPlacement struct {
	// ClusterName: [Required] The name of the cluster where the job will be
	// submitted.
	ClusterName string `json:"clusterName,omitempty"`

	// ClusterUuid: [Output-only] A cluster UUID generated by the Dataproc
	// service when the job
	// is submitted.
	ClusterUuid string `json:"clusterUuid,omitempty"`
}

type JobReference struct {
	// JobId: [Required] The job ID, which must be unique within the
	// project. The job ID
	// is generated by the server upon job submission or
	// provided by the user as a
	// means to perform retries without creating
	// duplicate jobs. The ID must
	// contain only letters (a-z, A-Z), numbers
	// (0-9), underscores (_), or
	// hyphens (-). The maximum length is 512
	// characters.
	JobId string `json:"jobId,omitempty"`

	// ProjectId: [Required] The ID of the Google Cloud Platform project
	// that the job
	// belongs to.
	ProjectId string `json:"projectId,omitempty"`
}

type JobStatus struct {
	// Details: [Optional] Job state details, such as an error
	// description
	// if the state is <code>ERROR</code>.
	Details string `json:"details,omitempty"`

	// State: [Required] A state message specifying the overall job state.
	State string `json:"state,omitempty"`

	// StateStartTime: [Output-only] The time when this state was entered.
	StateStartTime string `json:"stateStartTime,omitempty"`
}

type LeaseTasksRequest struct {
	// AgentId: The agent's id.
	AgentId string `json:"agentId,omitempty"`

	// CurrentAgentTime: The current timestamp at the worker.
	CurrentAgentTime string `json:"currentAgentTime,omitempty"`

	// RequestedLeaseDuration: The requested initial lease period.
	RequestedLeaseDuration string `json:"requestedLeaseDuration,omitempty"`
}

type LeaseTasksResponse struct {
	// LeaseExpirationTime: The worker-local lease expiration time.
	LeaseExpirationTime string `json:"leaseExpirationTime,omitempty"`

	// ReportStatusInterval: The interval at which status should be
	// reported.
	ReportStatusInterval string `json:"reportStatusInterval,omitempty"`

	// Tasks: A list of tasks that have been leased.
	Tasks []*Task `json:"tasks,omitempty"`
}

type ListAgentsResponse struct {
	// Agents: A list of agents.
	Agents []*Agent `json:"agents,omitempty"`

	// NextPageToken: The token to send to ListAgents to acquire any
	// following pages.
	// Will be empty for last page.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type ListClustersResponse struct {
	// Clusters: [Output-only] The clusters in the project.
	Clusters []*Cluster `json:"clusters,omitempty"`

	// NextPageToken: The standard List next-page token.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type ListJobsResponse struct {
	// Jobs: [Output-only] Jobs list.
	Jobs []*Job `json:"jobs,omitempty"`

	// NextPageToken: [Optional] This token is included in the response if
	// there are more results
	// to fetch. To fetch additional results, provide
	// this value as the
	// `page_token` in a subsequent
	// <code>ListJobsRequest</code>.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type ListOperationsResponse struct {
	// NextPageToken: The standard List next-page token.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Operations: A list of operations that matches the specified filter in
	// the request.
	Operations []*Operation `json:"operations,omitempty"`
}

type LoggingConfiguration struct {
	// DriverLogLevels: The per-package log levels for the driver. This may
	// include
	// "root" package name to configure rootLogger.
	// Examples:
	//
	// 'com.google = FATAL', 'root = INFO', 'org.apache = DEBUG'
	DriverLogLevels map[string]string `json:"driverLogLevels,omitempty"`
}

type ManagedGroupConfiguration struct {
	// InstanceGroupManagerName: [Output-only] The name of the Instance
	// Group Manager for this group.
	InstanceGroupManagerName string `json:"instanceGroupManagerName,omitempty"`

	// InstanceTemplateName: [Output-only] The name of the Instance Template
	// used for the Managed
	// Instance Group.
	InstanceTemplateName string `json:"instanceTemplateName,omitempty"`
}

type NodeInitializationAction struct {
	// ExecutableFile: [Required] Google Cloud Storage URI of executable
	// file.
	ExecutableFile string `json:"executableFile,omitempty"`

	// ExecutionTimeout: [Optional] Amount of time executable has to
	// complete. Default is
	// 10 minutes. Cluster creation fails with an
	// explanatory error message (the
	// name of the executable that caused the
	// error and the exceeded timeout
	// period) if the executable is not
	// completed at end of the timeout period.
	ExecutionTimeout string `json:"executionTimeout,omitempty"`
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
	// ClusterName: Name of the cluster for the operation.
	ClusterName string `json:"clusterName,omitempty"`

	// ClusterUuid: Cluster UUId for the operation.
	ClusterUuid string `json:"clusterUuid,omitempty"`

	// Details: A message containing any operation metadata details.
	Details string `json:"details,omitempty"`

	// EndTime: The time that the operation completed.
	EndTime string `json:"endTime,omitempty"`

	// InnerState: A message containing the detailed operation state.
	InnerState string `json:"innerState,omitempty"`

	// InsertTime: The time that the operation was requested.
	InsertTime string `json:"insertTime,omitempty"`

	// StartTime: The time that the operation was started by the server.
	StartTime string `json:"startTime,omitempty"`

	// State: A message containing the operation state.
	State string `json:"state,omitempty"`

	// Status: [Output-only] Current operation status.
	Status *OperationStatus `json:"status,omitempty"`

	// StatusHistory: [Output-only] Previous operation status.
	StatusHistory []*OperationStatus `json:"statusHistory,omitempty"`
}

type OperationStatus struct {
	// Details: A message containing any operation metadata details.
	Details string `json:"details,omitempty"`

	// InnerState: A message containing the detailed operation state.
	InnerState string `json:"innerState,omitempty"`

	// State: A message containing the operation state.
	State string `json:"state,omitempty"`

	// StateStartTime: The time this state was entered.
	StateStartTime string `json:"stateStartTime,omitempty"`
}

type PigJob struct {
	// ContinueOnFailure: [Optional] Whether to continue executing queries
	// if a query fails.
	// The default value is `false`. Setting to `true` can
	// be useful when executing
	// independent parallel queries.
	ContinueOnFailure bool `json:"continueOnFailure,omitempty"`

	// JarFileUris: [Optional] HCFS URIs of jar files to add to the
	// CLASSPATH of
	// the Pig Client and Hadoop MapReduce (MR) tasks. Can
	// contain Pig UDFs.
	JarFileUris []string `json:"jarFileUris,omitempty"`

	// LoggingConfiguration: [Optional] The runtime log configuration for
	// job execution.
	LoggingConfiguration *LoggingConfiguration `json:"loggingConfiguration,omitempty"`

	// Properties: [Optional] A mapping of property names to values, used to
	// configure Pig.
	// Properties that conflict with values set by the Cloud
	// Dataproc API may be
	// overwritten. Can include properties set in
	// /etc/hadoop/conf/*-site.xml,
	// /etc/pig/conf/pig.properties, and
	// classes in user code.
	Properties map[string]string `json:"properties,omitempty"`

	// QueryFileUri: The HCFS URI of the script that contains the Pig
	// queries.
	QueryFileUri string `json:"queryFileUri,omitempty"`

	// QueryList: A list of queries.
	QueryList *QueryList `json:"queryList,omitempty"`

	// ScriptVariables: [Optional] Mapping of query variable names to values
	// (equivalent to the Pig
	// command: `name=[value]`).
	ScriptVariables map[string]string `json:"scriptVariables,omitempty"`
}

type PySparkJob struct {
	// ArchiveUris: [Optional] HCFS URIs of archives to be extracted in the
	// working directory of
	// .jar, .tar, .tar.gz, .tgz, and .zip.
	ArchiveUris []string `json:"archiveUris,omitempty"`

	// Args: [Optional] The arguments to pass to the driver.  Do not include
	// arguments,
	// such as `--conf`, that can be set as job properties, since
	// a collision may
	// occur that causes an incorrect job submission.
	Args []string `json:"args,omitempty"`

	// FileUris: [Optional] HCFS URIs of files to be copied to the working
	// directory of
	// Python drivers and distributed tasks. Useful for naively
	// parallel tasks.
	FileUris []string `json:"fileUris,omitempty"`

	// JarFileUris: [Optional] HCFS URIs of jar files to add to the
	// CLASSPATHs of the
	// Python driver and tasks.
	JarFileUris []string `json:"jarFileUris,omitempty"`

	// LoggingConfiguration: [Optional] The runtime log configuration for
	// job execution.
	LoggingConfiguration *LoggingConfiguration `json:"loggingConfiguration,omitempty"`

	// MainPythonFileUri: [Required] The Hadoop Compatible Filesystem (HCFS)
	// URI of the main
	// Python file to use as the driver. Must be a .py file.
	MainPythonFileUri string `json:"mainPythonFileUri,omitempty"`

	// Properties: [Optional] A mapping of property names to values, used to
	// configure PySpark.
	// Properties that conflict with values set by the
	// Cloud Dataproc API may be
	// overwritten. Can include properties set
	// in
	// /etc/spark/conf/spark-defaults.conf and classes in user code.
	Properties map[string]string `json:"properties,omitempty"`

	// PythonFileUris: [Optional] HCFS file URIs of Python files to pass to
	// the PySpark
	// framework. Supported file types: .py, .egg, and .zip.
	PythonFileUris []string `json:"pythonFileUris,omitempty"`
}

type QueryList struct {
	// Queries: [Required] The queries to execute. You do not need to
	// terminate a query
	// with a semicolon. Multiple queries can be specified
	// in one string
	// by separating each with a semicolon. Here is an example
	// of an Cloud
	// Dataproc API snippet that uses a QueryList to specify a
	// HiveJob:
	//
	//     "hiveJob": {
	//       "queryList": {
	//         "queries": [
	//
	//          "query1",
	//           "query2",
	//           "query3;query4",
	//
	//     ]
	//       }
	//     }
	Queries []string `json:"queries,omitempty"`
}

type ReportTaskStatusRequest struct {
	// AgentId: The id of the agent reporting task status.
	AgentId string `json:"agentId,omitempty"`

	// CurrentWorkerTime: The current timestamp at the worker.
	CurrentWorkerTime string `json:"currentWorkerTime,omitempty"`

	// Status: Status for a single task.
	Status *TaskStatus `json:"status,omitempty"`
}

type ReportTaskStatusResponse struct {
	// LeaseExpirationTime: New task lease expiration timestamp in
	// worker-local time.
	LeaseExpirationTime string `json:"leaseExpirationTime,omitempty"`

	// ReportStatusInterval: The interval at which status should be
	// reported.
	ReportStatusInterval string `json:"reportStatusInterval,omitempty"`
}

type SoftwareConfiguration struct {
	// ImageVersion: [Optional] The version of software inside the cluster.
	// It must match the
	// regular expression `[0-9]+\.[0-9]+`. If
	// unspecified, it defaults to the
	// latest version (see [Cloud Dataproc
	// Versioning](/dataproc/versioning)).
	ImageVersion string `json:"imageVersion,omitempty"`
}

type SparkJob struct {
	// ArchiveUris: [Optional] HCFS URIs of archives to be extracted in the
	// working directory of
	// Spark drivers and tasks. Supported file
	// types:
	// .jar, .tar, .tar.gz, .tgz, and .zip.
	ArchiveUris []string `json:"archiveUris,omitempty"`

	// Args: [Optional] The arguments to pass to the driver. Do not include
	// arguments,
	// such as `--conf`, that can be set as job properties, since
	// a collision may
	// occur that causes an incorrect job submission.
	Args []string `json:"args,omitempty"`

	// FileUris: [Optional] HCFS URIs of files to be copied to the working
	// directory of
	// Spark drivers and distributed tasks. Useful for naively
	// parallel tasks.
	FileUris []string `json:"fileUris,omitempty"`

	// JarFileUris: [Optional] HCFS URIs of jar files to add to the
	// CLASSPATHs of the
	// Spark driver and tasks.
	JarFileUris []string `json:"jarFileUris,omitempty"`

	// LoggingConfiguration: [Optional] The runtime log configuration for
	// job execution.
	LoggingConfiguration *LoggingConfiguration `json:"loggingConfiguration,omitempty"`

	// MainClass: The name of the driver's main class. The jar file that
	// contains the class
	// must be in the default CLASSPATH or specified in
	// `jar_file_uris`.
	MainClass string `json:"mainClass,omitempty"`

	// MainJarFileUri: The Hadoop Compatible Filesystem (HCFS) URI of the
	// jar file that
	// contains the main class.
	MainJarFileUri string `json:"mainJarFileUri,omitempty"`

	// Properties: [Optional] A mapping of property names to values, used to
	// configure Spark.
	// Properties that conflict with values set by the
	// Cloud Dataproc API may be
	// overwritten. Can include properties set
	// in
	// /etc/spark/conf/spark-defaults.conf and classes in user code.
	Properties map[string]string `json:"properties,omitempty"`
}

type SparkSqlJob struct {
	// JarFileUris: [Optional] HCFS URIs of jar files to be added to the
	// Spark CLASSPATH.
	JarFileUris []string `json:"jarFileUris,omitempty"`

	// LoggingConfiguration: [Optional] The runtime log configuration for
	// job execution.
	LoggingConfiguration *LoggingConfiguration `json:"loggingConfiguration,omitempty"`

	// Properties: [Optional] A mapping of property names to values, used to
	// configure
	// Spark SQL's SparkConf. Properties that conflict with values
	// set by the
	// Cloud Dataproc API may be overwritten.
	Properties map[string]string `json:"properties,omitempty"`

	// QueryFileUri: The HCFS URI of the script that contains SQL queries.
	QueryFileUri string `json:"queryFileUri,omitempty"`

	// QueryList: A list of queries.
	QueryList *QueryList `json:"queryList,omitempty"`

	// ScriptVariables: [Optional] Mapping of query variable names to values
	// (equivalent to the
	// Spark SQL command: SET `name="value";`).
	ScriptVariables map[string]string `json:"scriptVariables,omitempty"`
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

type SubmitJobRequest struct {
	// Job: [Required] The job resource.
	Job *Job `json:"job,omitempty"`
}

type SystemTaskStatus struct {
	// State: The outcome of reconfiguration.
	State string `json:"state,omitempty"`
}

type Task struct {
	// Configuration: Configuration for this task.
	Configuration *TaskConfiguration `json:"configuration,omitempty"`

	// Status: The status of a task.
	Status *TaskStatus `json:"status,omitempty"`

	// TaskId: System defined task id.
	TaskId string `json:"taskId,omitempty"`
}

type TaskClusterConfiguration struct {
	// AddMembers: New nodes to register with cluster.
	AddMembers []string `json:"addMembers,omitempty"`

	// RemoveMembers: Existing nodes to decommission.
	RemoveMembers []string `json:"removeMembers,omitempty"`

	// Type: Type of configuration change.
	Type string `json:"type,omitempty"`
}

type TaskConfiguration struct {
	// ClusterConfiguration: Cluster reconfiguration task.
	ClusterConfiguration *TaskClusterConfiguration `json:"clusterConfiguration,omitempty"`

	// JobConfiguration: Configuration of a Job-based task.
	JobConfiguration *TaskJobConfiguration `json:"jobConfiguration,omitempty"`

	// MaintenanceCommand: Execute cluster maintenance command.
	MaintenanceCommand *TaskMaintenanceCommand `json:"maintenanceCommand,omitempty"`
}

type TaskJobConfiguration struct {
	// ArchiveUris: Required archives for the driver program or distributed
	// program.
	// Used by Hadoop, Spark, and PySpark jobs.
	ArchiveUris []string `json:"archiveUris,omitempty"`

	// Args: Arguments for the driver program. Used by Hadoop, Spark, and
	// PySpark jobs.
	Args []string `json:"args,omitempty"`

	// ContinueOnFailure: True to continue processing pig or hive queries if
	// an earlier
	// query fails.
	ContinueOnFailure bool `json:"continueOnFailure,omitempty"`

	// DriverControlFilesUri: [Output-only] If present, the location of
	// miscellaneous control files
	// which may be used as part of job setup
	// and handling. If not present,
	// control files may be placed in the same
	// location as driver_output_uri.
	DriverControlFilesUri string `json:"driverControlFilesUri,omitempty"`

	// DriverInputUri: [Output-only] A URI pointing to the location of the
	// stdin of the job's
	// driver program, only set if the job is
	// interactive.
	DriverInputUri string `json:"driverInputUri,omitempty"`

	// DriverOutputUri: Output URI for driver output.
	DriverOutputUri string `json:"driverOutputUri,omitempty"`

	// FileUris: Required files for the driver program or distributed
	// program.
	// Used by Hadoop, Spark, and PySpark jobs.
	FileUris []string `json:"fileUris,omitempty"`

	// Interactive: [Optional] If set to true, then the driver's stdin will
	// be kept open and
	// driver_input_uri will be set to provide a path at
	// which additional input
	// can be sent to the driver.
	Interactive bool `json:"interactive,omitempty"`

	// JarFileUris: JAR files that are required by the job.
	JarFileUris []string `json:"jarFileUris,omitempty"`

	// LoggingConfiguration: Logging configuration for the job.
	LoggingConfiguration *TaskLoggingConfiguration `json:"loggingConfiguration,omitempty"`

	// MainClass: A class name that is contained either in core Hadoop or
	// Spark
	// libraries or within a JAR specified within jar_file_uris.
	MainClass string `json:"mainClass,omitempty"`

	// MainJarFileUri: A JAR containing the main driver and containing a
	// METADATA entry
	// for a main class contained within the jar.
	MainJarFileUri string `json:"mainJarFileUri,omitempty"`

	// MainPythonFileUri: The main Python file for a PySpark application.
	MainPythonFileUri string `json:"mainPythonFileUri,omitempty"`

	// Properties: Properties for the submitted job.
	Properties map[string]string `json:"properties,omitempty"`

	// PythonFileUris: URIs of files required by the PySpark application
	PythonFileUris []string `json:"pythonFileUris,omitempty"`

	// QueryFileUri: A URI of a file containing queries
	QueryFileUri string `json:"queryFileUri,omitempty"`

	// QueryList: A list of queries specified within the API.
	QueryList *TaskQueryList `json:"queryList,omitempty"`

	// ScriptVariables: Variables to be substituted in Pig and Hive scripts.
	ScriptVariables map[string]string `json:"scriptVariables,omitempty"`

	// SubmittedBy: The user that the job should be attributed to in Hadoop
	// as a posix-style
	// username. If the user is not a member of the
	// system's user-database,
	// the task will be started as a system account.
	SubmittedBy string `json:"submittedBy,omitempty"`

	// Type: The type of the job.
	Type string `json:"type,omitempty"`
}

type TaskJobStatus struct {
	// DriverExitCode: If the driver has exited, its exit code.
	DriverExitCode int64 `json:"driverExitCode,omitempty"`

	// DriverState: The state of the driver.
	DriverState string `json:"driverState,omitempty"`

	// YarnApplications: A list of YARN applications that have been launched
	// for this task.
	YarnApplications []*TaskYarnApplication `json:"yarnApplications,omitempty"`
}

type TaskLoggingConfiguration struct {
	// LogLevels: Map of logger name to log4j log level.
	LogLevels map[string]string `json:"logLevels,omitempty"`
}

type TaskMaintenanceCommand struct {
	// Args: Arguments to pass to the script.
	Args []string `json:"args,omitempty"`

	// Environment: The environment variables.
	Environment map[string]string `json:"environment,omitempty"`

	// GcsUri: The executable is stored on GCS.
	GcsUri string `json:"gcsUri,omitempty"`

	// LocalPath: The executable is a file on agent.
	LocalPath string `json:"localPath,omitempty"`

	// ScriptOutputUri: The GCS URI where executable output will be stored.
	ScriptOutputUri string `json:"scriptOutputUri,omitempty"`
}

type TaskQueryList struct {
	// Queries: The queries to execute. The format of the queries is
	// task-type dependent,
	// but in each case each query should be executed
	// within its own
	// invocation of the interpreter for that task type.
	Queries []string `json:"queries,omitempty"`
}

type TaskStatus struct {
	// JobStatus: The status of the Job.
	JobStatus *TaskJobStatus `json:"jobStatus,omitempty"`

	// SystemTaskStatus: The status of the SystemTask.
	SystemTaskStatus *SystemTaskStatus `json:"systemTaskStatus,omitempty"`
}

type TaskYarnApplication struct {
	// Id: YARN application id.
	Id int64 `json:"id,omitempty"`

	// Name: YARN application name.
	Name string `json:"name,omitempty"`

	// Progress: The progress of the YARN application.
	Progress float64 `json:"progress,omitempty"`

	// State: The state of the YARN application.
	State string `json:"state,omitempty"`

	// TrackingUrl: The tracking URL for the YARN application. This URL may
	// or may not be
	// accessible from outside the cluster.
	TrackingUrl string `json:"trackingUrl,omitempty"`
}

type YarnApplication struct {
	// Name: [Required] The application name.
	Name string `json:"name,omitempty"`

	// Progress: [Required] The numerical progress of the application, from
	// 1 to 100.
	Progress float64 `json:"progress,omitempty"`

	// State: [Required] The application state.
	State string `json:"state,omitempty"`

	// TrackingUrl: [Optional] The HTTP URL of the ApplicationMaster,
	// HistoryServer, or
	// TimelineServer that provides application-specific
	// information. The URL uses
	// the internal hostname, and requires a proxy
	// server for resolution and,
	// possibly, access.
	TrackingUrl string `json:"trackingUrl,omitempty"`
}

// method id "dataproc.operations.cancel":

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
// [operations.get](/dataproc/reference/rest/v1beta1/operations/get)
// or
// other methods to check whether the cancellation succeeded or
// whether the
// operation completed despite cancellation.
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/{+name}:cancel")
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
	//   "description": "Starts asynchronous cancellation on a long-running operation.  The server\nmakes a best effort to cancel the operation, but success is not\nguaranteed.  If the server doesn't support this method, it returns\n`google.rpc.Code.UNIMPLEMENTED`.  Clients can use\n[operations.get](/dataproc/reference/rest/v1beta1/operations/get) or\nother methods to check whether the cancellation succeeded or whether the\noperation completed despite cancellation.",
	//   "flatPath": "v1beta1/operations/{operationsId}:cancel",
	//   "httpMethod": "POST",
	//   "id": "dataproc.operations.cancel",
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
	//   "path": "v1beta1/{+name}:cancel",
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

// method id "dataproc.operations.delete":

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
	var ret *Empty
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes a long-running operation. This method indicates that the client is\nno longer interested in the operation result. It does not cancel the\noperation. If the server doesn't support this method, it returns\n`google.rpc.Code.UNIMPLEMENTED`.",
	//   "flatPath": "v1beta1/operations/{operationsId}",
	//   "httpMethod": "DELETE",
	//   "id": "dataproc.operations.delete",
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
	//   "path": "v1beta1/{+name}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dataproc.operations.get":

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
	//   "id": "dataproc.operations.get",
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
	//   "path": "v1beta1/{+name}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dataproc.operations.list":

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
	var ret *ListOperationsResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists operations that match the specified filter in the request. If the\nserver doesn't support this method, it returns `UNIMPLEMENTED`.\n\nNOTE: the `name` binding below allows API services to override the binding\nto use different resource name schemes, such as `users/*/operations`.",
	//   "flatPath": "v1beta1/operations",
	//   "httpMethod": "GET",
	//   "id": "dataproc.operations.list",
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
	//   "path": "v1beta1/{+name}",
	//   "response": {
	//     "$ref": "ListOperationsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dataproc.projects.clusters.create":

type ProjectsClustersCreateCall struct {
	s         *Service
	projectId string
	cluster   *Cluster
	opt_      map[string]interface{}
}

// Create: Creates a cluster in a project.
func (r *ProjectsClustersService) Create(projectId string, cluster *Cluster) *ProjectsClustersCreateCall {
	c := &ProjectsClustersCreateCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.cluster = cluster
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsClustersCreateCall) Fields(s ...googleapi.Field) *ProjectsClustersCreateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsClustersCreateCall) Do() (*Operation, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.cluster)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/projects/{projectId}/clusters")
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
	//   "description": "Creates a cluster in a project.",
	//   "flatPath": "v1beta1/projects/{projectId}/clusters",
	//   "httpMethod": "POST",
	//   "id": "dataproc.projects.clusters.create",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "[Required] The ID of the Google Cloud Platform project that the cluster belongs to.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/projects/{projectId}/clusters",
	//   "request": {
	//     "$ref": "Cluster"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dataproc.projects.clusters.delete":

type ProjectsClustersDeleteCall struct {
	s           *Service
	projectId   string
	clusterName string
	opt_        map[string]interface{}
}

// Delete: Deletes a cluster in a project.
func (r *ProjectsClustersService) Delete(projectId string, clusterName string) *ProjectsClustersDeleteCall {
	c := &ProjectsClustersDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.clusterName = clusterName
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsClustersDeleteCall) Fields(s ...googleapi.Field) *ProjectsClustersDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsClustersDeleteCall) Do() (*Operation, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/projects/{projectId}/clusters/{clusterName}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId":   c.projectId,
		"clusterName": c.clusterName,
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
	//   "description": "Deletes a cluster in a project.",
	//   "flatPath": "v1beta1/projects/{projectId}/clusters/{clusterName}",
	//   "httpMethod": "DELETE",
	//   "id": "dataproc.projects.clusters.delete",
	//   "parameterOrder": [
	//     "projectId",
	//     "clusterName"
	//   ],
	//   "parameters": {
	//     "clusterName": {
	//       "description": "[Required] The cluster name.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "[Required] The ID of the Google Cloud Platform project that the cluster belongs to.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/projects/{projectId}/clusters/{clusterName}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dataproc.projects.clusters.diagnose":

type ProjectsClustersDiagnoseCall struct {
	s                      *Service
	projectId              string
	clusterName            string
	diagnoseclusterrequest *DiagnoseClusterRequest
	opt_                   map[string]interface{}
}

// Diagnose: Gets cluster diagnostic information.
// After the operation
// completes, the Operation.response field
// contains
// `DiagnoseClusterOutputLocation`.
func (r *ProjectsClustersService) Diagnose(projectId string, clusterName string, diagnoseclusterrequest *DiagnoseClusterRequest) *ProjectsClustersDiagnoseCall {
	c := &ProjectsClustersDiagnoseCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.clusterName = clusterName
	c.diagnoseclusterrequest = diagnoseclusterrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsClustersDiagnoseCall) Fields(s ...googleapi.Field) *ProjectsClustersDiagnoseCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsClustersDiagnoseCall) Do() (*Operation, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.diagnoseclusterrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/projects/{projectId}/clusters/{clusterName}:diagnose")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId":   c.projectId,
		"clusterName": c.clusterName,
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
	//   "description": "Gets cluster diagnostic information.\nAfter the operation completes, the Operation.response field\ncontains `DiagnoseClusterOutputLocation`.",
	//   "flatPath": "v1beta1/projects/{projectId}/clusters/{clusterName}:diagnose",
	//   "httpMethod": "POST",
	//   "id": "dataproc.projects.clusters.diagnose",
	//   "parameterOrder": [
	//     "projectId",
	//     "clusterName"
	//   ],
	//   "parameters": {
	//     "clusterName": {
	//       "description": "[Required] The cluster name.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "[Required] The ID of the Google Cloud Platform project that the cluster\nbelongs to.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/projects/{projectId}/clusters/{clusterName}:diagnose",
	//   "request": {
	//     "$ref": "DiagnoseClusterRequest"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dataproc.projects.clusters.get":

type ProjectsClustersGetCall struct {
	s           *Service
	projectId   string
	clusterName string
	opt_        map[string]interface{}
}

// Get: Gets the resource representation for a cluster in a project.
func (r *ProjectsClustersService) Get(projectId string, clusterName string) *ProjectsClustersGetCall {
	c := &ProjectsClustersGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.clusterName = clusterName
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsClustersGetCall) Fields(s ...googleapi.Field) *ProjectsClustersGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsClustersGetCall) Do() (*Cluster, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/projects/{projectId}/clusters/{clusterName}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId":   c.projectId,
		"clusterName": c.clusterName,
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
	var ret *Cluster
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets the resource representation for a cluster in a project.",
	//   "flatPath": "v1beta1/projects/{projectId}/clusters/{clusterName}",
	//   "httpMethod": "GET",
	//   "id": "dataproc.projects.clusters.get",
	//   "parameterOrder": [
	//     "projectId",
	//     "clusterName"
	//   ],
	//   "parameters": {
	//     "clusterName": {
	//       "description": "[Required] The cluster name.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "[Required] The ID of the Google Cloud Platform project that the cluster belongs to.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/projects/{projectId}/clusters/{clusterName}",
	//   "response": {
	//     "$ref": "Cluster"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dataproc.projects.clusters.list":

type ProjectsClustersListCall struct {
	s         *Service
	projectId string
	opt_      map[string]interface{}
}

// List: Lists all clusters in a project.
func (r *ProjectsClustersService) List(projectId string) *ProjectsClustersListCall {
	c := &ProjectsClustersListCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	return c
}

// PageSize sets the optional parameter "pageSize": The standard List
// page size.
func (c *ProjectsClustersListCall) PageSize(pageSize int64) *ProjectsClustersListCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken": The standard List
// page token.
func (c *ProjectsClustersListCall) PageToken(pageToken string) *ProjectsClustersListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsClustersListCall) Fields(s ...googleapi.Field) *ProjectsClustersListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsClustersListCall) Do() (*ListClustersResponse, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/projects/{projectId}/clusters")
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
	var ret *ListClustersResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists all clusters in a project.",
	//   "flatPath": "v1beta1/projects/{projectId}/clusters",
	//   "httpMethod": "GET",
	//   "id": "dataproc.projects.clusters.list",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "The standard List page size.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The standard List page token.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "[Required] The ID of the Google Cloud Platform project that the cluster belongs to.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/projects/{projectId}/clusters",
	//   "response": {
	//     "$ref": "ListClustersResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dataproc.projects.clusters.patch":

type ProjectsClustersPatchCall struct {
	s           *Service
	projectId   string
	clusterName string
	cluster     *Cluster
	opt_        map[string]interface{}
}

// Patch: Updates a cluster in a project.
func (r *ProjectsClustersService) Patch(projectId string, clusterName string, cluster *Cluster) *ProjectsClustersPatchCall {
	c := &ProjectsClustersPatchCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.clusterName = clusterName
	c.cluster = cluster
	return c
}

// UpdateMask sets the optional parameter "updateMask": [Required]
// Specifies the path, relative to <code>Cluster</code>, of
// the field to
// update. For example, to change the number of workers
// in a cluster to
// 5, the <code>update_mask</code> parameter would be
// specified as
// <code>configuration.worker_configuration.num_instances</code>,
// and
// the `PATCH` request body would specify the new value, as follows:
//
//
//  {
//       "configuration":{
//         "workerConfiguration":{
//
// "numInstances":"5"
//         }
//       }
//     }
// <strong>Note:</strong>
// Currently,
// <code>configuration.worker_configuration.num_instances</code>
// is the
// only field that can be updated.
func (c *ProjectsClustersPatchCall) UpdateMask(updateMask string) *ProjectsClustersPatchCall {
	c.opt_["updateMask"] = updateMask
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsClustersPatchCall) Fields(s ...googleapi.Field) *ProjectsClustersPatchCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsClustersPatchCall) Do() (*Operation, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.cluster)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["updateMask"]; ok {
		params.Set("updateMask", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/projects/{projectId}/clusters/{clusterName}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId":   c.projectId,
		"clusterName": c.clusterName,
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
	//   "description": "Updates a cluster in a project.",
	//   "flatPath": "v1beta1/projects/{projectId}/clusters/{clusterName}",
	//   "httpMethod": "PATCH",
	//   "id": "dataproc.projects.clusters.patch",
	//   "parameterOrder": [
	//     "projectId",
	//     "clusterName"
	//   ],
	//   "parameters": {
	//     "clusterName": {
	//       "description": "[Required] The cluster name.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "[Required] The ID of the Google Cloud Platform project the\ncluster belongs to.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "updateMask": {
	//       "description": "[Required] Specifies the path, relative to \u003ccode\u003eCluster\u003c/code\u003e, of\nthe field to update. For example, to change the number of workers\nin a cluster to 5, the \u003ccode\u003eupdate_mask\u003c/code\u003e parameter would be\nspecified as \u003ccode\u003econfiguration.worker_configuration.num_instances\u003c/code\u003e,\nand the `PATCH` request body would specify the new value, as follows:\n\n    {\n      \"configuration\":{\n        \"workerConfiguration\":{\n          \"numInstances\":\"5\"\n        }\n      }\n    }\n\u003cstrong\u003eNote:\u003c/strong\u003e Currently, \u003ccode\u003econfiguration.worker_configuration.num_instances\u003c/code\u003e\nis the only field that can be updated.",
	//       "format": "google-fieldmask",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/projects/{projectId}/clusters/{clusterName}",
	//   "request": {
	//     "$ref": "Cluster"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dataproc.projects.clusters.agents.create":

type ProjectsClustersAgentsCreateCall struct {
	s           *Service
	projectId   string
	clusterUuid string
	agentId     string
	agent       *Agent
	opt_        map[string]interface{}
}

// Create: Add a new agent to Dataproc's view of a cluster. This is the
// first
// Dataproc method that an agent should invoke after starting. If
// an
// agent has already been created with the given agent_id within
// the
// same cluster, this method will return a Conflict status code
// and the
// agent is expected to call GetAgent to retrieve the
// last registration
// and subsequently call UpdateAgent, if required.
func (r *ProjectsClustersAgentsService) Create(projectId string, clusterUuid string, agentId string, agent *Agent) *ProjectsClustersAgentsCreateCall {
	c := &ProjectsClustersAgentsCreateCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.clusterUuid = clusterUuid
	c.agentId = agentId
	c.agent = agent
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsClustersAgentsCreateCall) Fields(s ...googleapi.Field) *ProjectsClustersAgentsCreateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsClustersAgentsCreateCall) Do() (*Agent, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.agent)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/projects/{projectId}/clusters/{clusterUuid}/agents/{agentId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId":   c.projectId,
		"clusterUuid": c.clusterUuid,
		"agentId":     c.agentId,
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
	var ret *Agent
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Add a new agent to Dataproc's view of a cluster. This is the first\nDataproc method that an agent should invoke after starting. If an\nagent has already been created with the given agent_id within\nthe same cluster, this method will return a Conflict status code\nand the agent is expected to call GetAgent to retrieve the\nlast registration and subsequently call UpdateAgent, if required.",
	//   "flatPath": "v1beta1/projects/{projectId}/clusters/{clusterUuid}/agents/{agentId}",
	//   "httpMethod": "PUT",
	//   "id": "dataproc.projects.clusters.agents.create",
	//   "parameterOrder": [
	//     "projectId",
	//     "clusterUuid",
	//     "agentId"
	//   ],
	//   "parameters": {
	//     "agentId": {
	//       "description": "[Required] Agent ID being registered.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "clusterUuid": {
	//       "description": "Cluster that this agent is associated with",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "Project ID that this agent is associated with",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/projects/{projectId}/clusters/{clusterUuid}/agents/{agentId}",
	//   "request": {
	//     "$ref": "Agent"
	//   },
	//   "response": {
	//     "$ref": "Agent"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/logging.admin",
	//     "https://www.googleapis.com/auth/logging.read",
	//     "https://www.googleapis.com/auth/logging.write"
	//   ]
	// }

}

// method id "dataproc.projects.clusters.agents.delete":

type ProjectsClustersAgentsDeleteCall struct {
	s           *Service
	projectId   string
	clusterUuid string
	agentId     string
	opt_        map[string]interface{}
}

// Delete: Delete an agent from Dataproc's view of a cluster. Deleting
// an
// agent is not required, but could be used in a shutdown sequence
// to
// indicate to Dataproc that the agent is to be considered dead
// and all
// agent-owned resources and tasks are free to be re-distributed.
func (r *ProjectsClustersAgentsService) Delete(projectId string, clusterUuid string, agentId string) *ProjectsClustersAgentsDeleteCall {
	c := &ProjectsClustersAgentsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.clusterUuid = clusterUuid
	c.agentId = agentId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsClustersAgentsDeleteCall) Fields(s ...googleapi.Field) *ProjectsClustersAgentsDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsClustersAgentsDeleteCall) Do() (*Empty, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/projects/{projectId}/clusters/{clusterUuid}/agents/{agentId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId":   c.projectId,
		"clusterUuid": c.clusterUuid,
		"agentId":     c.agentId,
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
	//   "description": "Delete an agent from Dataproc's view of a cluster. Deleting an\nagent is not required, but could be used in a shutdown sequence\nto indicate to Dataproc that the agent is to be considered dead\nand all agent-owned resources and tasks are free to be re-distributed.",
	//   "flatPath": "v1beta1/projects/{projectId}/clusters/{clusterUuid}/agents/{agentId}",
	//   "httpMethod": "DELETE",
	//   "id": "dataproc.projects.clusters.agents.delete",
	//   "parameterOrder": [
	//     "projectId",
	//     "clusterUuid",
	//     "agentId"
	//   ],
	//   "parameters": {
	//     "agentId": {
	//       "description": "The agent.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "clusterUuid": {
	//       "description": "The agent's cluster.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "The agent's project.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/projects/{projectId}/clusters/{clusterUuid}/agents/{agentId}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/logging.admin",
	//     "https://www.googleapis.com/auth/logging.read",
	//     "https://www.googleapis.com/auth/logging.write"
	//   ]
	// }

}

// method id "dataproc.projects.clusters.agents.get":

type ProjectsClustersAgentsGetCall struct {
	s           *Service
	projectId   string
	clusterUuid string
	agentId     string
	opt_        map[string]interface{}
}

// Get: Retrieve an agent.
func (r *ProjectsClustersAgentsService) Get(projectId string, clusterUuid string, agentId string) *ProjectsClustersAgentsGetCall {
	c := &ProjectsClustersAgentsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.clusterUuid = clusterUuid
	c.agentId = agentId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsClustersAgentsGetCall) Fields(s ...googleapi.Field) *ProjectsClustersAgentsGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsClustersAgentsGetCall) Do() (*Agent, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/projects/{projectId}/clusters/{clusterUuid}/agents/{agentId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId":   c.projectId,
		"clusterUuid": c.clusterUuid,
		"agentId":     c.agentId,
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
	var ret *Agent
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieve an agent.",
	//   "flatPath": "v1beta1/projects/{projectId}/clusters/{clusterUuid}/agents/{agentId}",
	//   "httpMethod": "GET",
	//   "id": "dataproc.projects.clusters.agents.get",
	//   "parameterOrder": [
	//     "projectId",
	//     "clusterUuid",
	//     "agentId"
	//   ],
	//   "parameters": {
	//     "agentId": {
	//       "description": "The agent's ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "clusterUuid": {
	//       "description": "The agent's cluster.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "The agent's project ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/projects/{projectId}/clusters/{clusterUuid}/agents/{agentId}",
	//   "response": {
	//     "$ref": "Agent"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/logging.admin",
	//     "https://www.googleapis.com/auth/logging.read",
	//     "https://www.googleapis.com/auth/logging.write"
	//   ]
	// }

}

// method id "dataproc.projects.clusters.agents.list":

type ProjectsClustersAgentsListCall struct {
	s           *Service
	projectId   string
	clusterUuid string
	opt_        map[string]interface{}
}

// List: List all agents Dataproc is aware of within a cluster.
func (r *ProjectsClustersAgentsService) List(projectId string, clusterUuid string) *ProjectsClustersAgentsListCall {
	c := &ProjectsClustersAgentsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.clusterUuid = clusterUuid
	return c
}

// PageSize sets the optional parameter "pageSize": Requested page size
// for listing.
func (c *ProjectsClustersAgentsListCall) PageSize(pageSize int64) *ProjectsClustersAgentsListCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken": Page continuation
// token.
func (c *ProjectsClustersAgentsListCall) PageToken(pageToken string) *ProjectsClustersAgentsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsClustersAgentsListCall) Fields(s ...googleapi.Field) *ProjectsClustersAgentsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsClustersAgentsListCall) Do() (*ListAgentsResponse, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/projects/{projectId}/clusters/{clusterUuid}/agents")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId":   c.projectId,
		"clusterUuid": c.clusterUuid,
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
	var ret *ListAgentsResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List all agents Dataproc is aware of within a cluster.",
	//   "flatPath": "v1beta1/projects/{projectId}/clusters/{clusterUuid}/agents",
	//   "httpMethod": "GET",
	//   "id": "dataproc.projects.clusters.agents.list",
	//   "parameterOrder": [
	//     "projectId",
	//     "clusterUuid"
	//   ],
	//   "parameters": {
	//     "clusterUuid": {
	//       "description": "The cluster from which to list agents.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Requested page size for listing.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Page continuation token.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "The project from which to list agents.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/projects/{projectId}/clusters/{clusterUuid}/agents",
	//   "response": {
	//     "$ref": "ListAgentsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/logging.admin",
	//     "https://www.googleapis.com/auth/logging.read",
	//     "https://www.googleapis.com/auth/logging.write"
	//   ]
	// }

}

// method id "dataproc.projects.clusters.agents.update":

type ProjectsClustersAgentsUpdateCall struct {
	s           *Service
	projectId   string
	clusterUuid string
	agentId     string
	agent       *Agent
	opt_        map[string]interface{}
}

// Update: Update Dataproc's view of an agent. This is currently used to
// provide a
// is_healthy bit, but is expected to be extended to include
// daemon
// information and VM metrics for inclusion in cloud metrics.
func (r *ProjectsClustersAgentsService) Update(projectId string, clusterUuid string, agentId string, agent *Agent) *ProjectsClustersAgentsUpdateCall {
	c := &ProjectsClustersAgentsUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.clusterUuid = clusterUuid
	c.agentId = agentId
	c.agent = agent
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsClustersAgentsUpdateCall) Fields(s ...googleapi.Field) *ProjectsClustersAgentsUpdateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsClustersAgentsUpdateCall) Do() (*Agent, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.agent)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/projects/{projectId}/clusters/{clusterUuid}/agents/{agentId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId":   c.projectId,
		"clusterUuid": c.clusterUuid,
		"agentId":     c.agentId,
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
	var ret *Agent
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Update Dataproc's view of an agent. This is currently used to provide a\nis_healthy bit, but is expected to be extended to include daemon\ninformation and VM metrics for inclusion in cloud metrics.",
	//   "flatPath": "v1beta1/projects/{projectId}/clusters/{clusterUuid}/agents/{agentId}",
	//   "httpMethod": "PUT",
	//   "id": "dataproc.projects.clusters.agents.update",
	//   "parameterOrder": [
	//     "projectId",
	//     "clusterUuid",
	//     "agentId"
	//   ],
	//   "parameters": {
	//     "agentId": {
	//       "description": "[Required] ID of agent sending the update.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "clusterUuid": {
	//       "description": "The cluster on which the agent is running.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "The agent's project.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/projects/{projectId}/clusters/{clusterUuid}/agents/{agentId}",
	//   "request": {
	//     "$ref": "Agent"
	//   },
	//   "response": {
	//     "$ref": "Agent"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/logging.admin",
	//     "https://www.googleapis.com/auth/logging.read",
	//     "https://www.googleapis.com/auth/logging.write"
	//   ]
	// }

}

// method id "dataproc.projects.clusters.tasks.lease":

type ProjectsClustersTasksLeaseCall struct {
	s                 *Service
	projectId         string
	clusterUuid       string
	leasetasksrequest *LeaseTasksRequest
	opt_              map[string]interface{}
}

// Lease: Obtain a lease on one or more tasks. Any given task may be in
// any state
// and each agent is expected to start any non-started tasks
// and to monitor
// any YarnApplications spawned by any already running
// tasks. It's expected
// that monitoring previously launched tasks will
// be more prevalent when
// drivers are run entirely within YARN
// containers.
//
// While there's a single lease expiration time, in the
// event of multiple
// tasks being leased to the agent in a single call,
// each task has a
// unique lease and status must be reported before the
// lease times out or
// the task can be considered orphaned.
//
// The service
// will determine how many tasks to lease to agents in a single
// call.
func (r *ProjectsClustersTasksService) Lease(projectId string, clusterUuid string, leasetasksrequest *LeaseTasksRequest) *ProjectsClustersTasksLeaseCall {
	c := &ProjectsClustersTasksLeaseCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.clusterUuid = clusterUuid
	c.leasetasksrequest = leasetasksrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsClustersTasksLeaseCall) Fields(s ...googleapi.Field) *ProjectsClustersTasksLeaseCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsClustersTasksLeaseCall) Do() (*LeaseTasksResponse, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.leasetasksrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/projects/{projectId}/clusters/{clusterUuid}/tasks:lease")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId":   c.projectId,
		"clusterUuid": c.clusterUuid,
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
	var ret *LeaseTasksResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Obtain a lease on one or more tasks. Any given task may be in any state\nand each agent is expected to start any non-started tasks and to monitor\nany YarnApplications spawned by any already running tasks. It's expected\nthat monitoring previously launched tasks will be more prevalent when\ndrivers are run entirely within YARN containers.\n\nWhile there's a single lease expiration time, in the event of multiple\ntasks being leased to the agent in a single call, each task has a\nunique lease and status must be reported before the lease times out or\nthe task can be considered orphaned.\n\nThe service will determine how many tasks to lease to agents in a single\ncall.",
	//   "flatPath": "v1beta1/projects/{projectId}/clusters/{clusterUuid}/tasks:lease",
	//   "httpMethod": "POST",
	//   "id": "dataproc.projects.clusters.tasks.lease",
	//   "parameterOrder": [
	//     "projectId",
	//     "clusterUuid"
	//   ],
	//   "parameters": {
	//     "clusterUuid": {
	//       "description": "The cluster id of the agent.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "The project id of the agent.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/projects/{projectId}/clusters/{clusterUuid}/tasks:lease",
	//   "request": {
	//     "$ref": "LeaseTasksRequest"
	//   },
	//   "response": {
	//     "$ref": "LeaseTasksResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/logging.admin",
	//     "https://www.googleapis.com/auth/logging.read",
	//     "https://www.googleapis.com/auth/logging.write"
	//   ]
	// }

}

// method id "dataproc.projects.clusters.tasks.report":

type ProjectsClustersTasksReportCall struct {
	s                       *Service
	projectId               string
	clusterUuid             string
	taskId                  string
	reporttaskstatusrequest *ReportTaskStatusRequest
	opt_                    map[string]interface{}
}

// Report: Report status for a task and extend the lease provided for
// the task.
func (r *ProjectsClustersTasksService) Report(projectId string, clusterUuid string, taskId string, reporttaskstatusrequest *ReportTaskStatusRequest) *ProjectsClustersTasksReportCall {
	c := &ProjectsClustersTasksReportCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.clusterUuid = clusterUuid
	c.taskId = taskId
	c.reporttaskstatusrequest = reporttaskstatusrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsClustersTasksReportCall) Fields(s ...googleapi.Field) *ProjectsClustersTasksReportCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsClustersTasksReportCall) Do() (*ReportTaskStatusResponse, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.reporttaskstatusrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/projects/{projectId}/clusters/{clusterUuid}/tasks/{taskId}:report")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId":   c.projectId,
		"clusterUuid": c.clusterUuid,
		"taskId":      c.taskId,
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
	var ret *ReportTaskStatusResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Report status for a task and extend the lease provided for the task.",
	//   "flatPath": "v1beta1/projects/{projectId}/clusters/{clusterUuid}/tasks/{taskId}:report",
	//   "httpMethod": "POST",
	//   "id": "dataproc.projects.clusters.tasks.report",
	//   "parameterOrder": [
	//     "projectId",
	//     "clusterUuid",
	//     "taskId"
	//   ],
	//   "parameters": {
	//     "clusterUuid": {
	//       "description": "The cluster id of the agent.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "The project id of the agent.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "taskId": {
	//       "description": "The task that is being reported on.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/projects/{projectId}/clusters/{clusterUuid}/tasks/{taskId}:report",
	//   "request": {
	//     "$ref": "ReportTaskStatusRequest"
	//   },
	//   "response": {
	//     "$ref": "ReportTaskStatusResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/logging.admin",
	//     "https://www.googleapis.com/auth/logging.read",
	//     "https://www.googleapis.com/auth/logging.write"
	//   ]
	// }

}

// method id "dataproc.projects.jobs.cancel":

type ProjectsJobsCancelCall struct {
	s                *Service
	projectId        string
	jobId            string
	canceljobrequest *CancelJobRequest
	opt_             map[string]interface{}
}

// Cancel: Starts a job cancellation request. To access the job
// resource
// after cancellation,
// call
// [jobs.list](/dataproc/reference/rest/v1beta1/projects.jobs/list)
// or
// [jobs.get](/dataproc/reference/rest/v1beta1/projects.jobs/get).
func (r *ProjectsJobsService) Cancel(projectId string, jobId string, canceljobrequest *CancelJobRequest) *ProjectsJobsCancelCall {
	c := &ProjectsJobsCancelCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.jobId = jobId
	c.canceljobrequest = canceljobrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsJobsCancelCall) Fields(s ...googleapi.Field) *ProjectsJobsCancelCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsJobsCancelCall) Do() (*Job, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.canceljobrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/projects/{projectId}/jobs/{jobId}:cancel")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"jobId":     c.jobId,
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
	var ret *Job
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Starts a job cancellation request. To access the job resource\nafter cancellation, call\n[jobs.list](/dataproc/reference/rest/v1beta1/projects.jobs/list) or\n[jobs.get](/dataproc/reference/rest/v1beta1/projects.jobs/get).",
	//   "flatPath": "v1beta1/projects/{projectId}/jobs/{jobId}:cancel",
	//   "httpMethod": "POST",
	//   "id": "dataproc.projects.jobs.cancel",
	//   "parameterOrder": [
	//     "projectId",
	//     "jobId"
	//   ],
	//   "parameters": {
	//     "jobId": {
	//       "description": "[Required] The job ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "[Required] The ID of the Google Cloud Platform project that the job\nbelongs to.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/projects/{projectId}/jobs/{jobId}:cancel",
	//   "request": {
	//     "$ref": "CancelJobRequest"
	//   },
	//   "response": {
	//     "$ref": "Job"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dataproc.projects.jobs.delete":

type ProjectsJobsDeleteCall struct {
	s         *Service
	projectId string
	jobId     string
	opt_      map[string]interface{}
}

// Delete: Deletes the job from the project. If the job is active, the
// delete fails,
// and the response returns `FAILED_PRECONDITION`.
func (r *ProjectsJobsService) Delete(projectId string, jobId string) *ProjectsJobsDeleteCall {
	c := &ProjectsJobsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.jobId = jobId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsJobsDeleteCall) Fields(s ...googleapi.Field) *ProjectsJobsDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsJobsDeleteCall) Do() (*Empty, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/projects/{projectId}/jobs/{jobId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"jobId":     c.jobId,
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
	//   "description": "Deletes the job from the project. If the job is active, the delete fails,\nand the response returns `FAILED_PRECONDITION`.",
	//   "flatPath": "v1beta1/projects/{projectId}/jobs/{jobId}",
	//   "httpMethod": "DELETE",
	//   "id": "dataproc.projects.jobs.delete",
	//   "parameterOrder": [
	//     "projectId",
	//     "jobId"
	//   ],
	//   "parameters": {
	//     "jobId": {
	//       "description": "[Required] The job ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "[Required] The ID of the Google Cloud Platform project that the job\nbelongs to.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/projects/{projectId}/jobs/{jobId}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dataproc.projects.jobs.get":

type ProjectsJobsGetCall struct {
	s         *Service
	projectId string
	jobId     string
	opt_      map[string]interface{}
}

// Get: Gets the resource representation for a job in a project.
func (r *ProjectsJobsService) Get(projectId string, jobId string) *ProjectsJobsGetCall {
	c := &ProjectsJobsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.jobId = jobId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsJobsGetCall) Fields(s ...googleapi.Field) *ProjectsJobsGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsJobsGetCall) Do() (*Job, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/projects/{projectId}/jobs/{jobId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"jobId":     c.jobId,
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
	var ret *Job
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets the resource representation for a job in a project.",
	//   "flatPath": "v1beta1/projects/{projectId}/jobs/{jobId}",
	//   "httpMethod": "GET",
	//   "id": "dataproc.projects.jobs.get",
	//   "parameterOrder": [
	//     "projectId",
	//     "jobId"
	//   ],
	//   "parameters": {
	//     "jobId": {
	//       "description": "[Required] The job ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "[Required] The ID of the Google Cloud Platform project that the job\nbelongs to.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/projects/{projectId}/jobs/{jobId}",
	//   "response": {
	//     "$ref": "Job"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dataproc.projects.jobs.list":

type ProjectsJobsListCall struct {
	s         *Service
	projectId string
	opt_      map[string]interface{}
}

// List: Lists jobs in a project.
func (r *ProjectsJobsService) List(projectId string) *ProjectsJobsListCall {
	c := &ProjectsJobsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	return c
}

// ClusterName sets the optional parameter "clusterName": [Optional] If
// set, the returned jobs list includes only jobs that were
// submitted to
// the named cluster.
func (c *ProjectsJobsListCall) ClusterName(clusterName string) *ProjectsJobsListCall {
	c.opt_["clusterName"] = clusterName
	return c
}

// JobStateMatcher sets the optional parameter "jobStateMatcher":
// [Optional] Specifies enumerated categories of jobs to list.
func (c *ProjectsJobsListCall) JobStateMatcher(jobStateMatcher string) *ProjectsJobsListCall {
	c.opt_["jobStateMatcher"] = jobStateMatcher
	return c
}

// PageSize sets the optional parameter "pageSize": [Optional] The
// number of results to return in each response.
func (c *ProjectsJobsListCall) PageSize(pageSize int64) *ProjectsJobsListCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken": [Optional] The
// page token, returned by a previous call, to request the
// next page of
// results.
func (c *ProjectsJobsListCall) PageToken(pageToken string) *ProjectsJobsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsJobsListCall) Fields(s ...googleapi.Field) *ProjectsJobsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsJobsListCall) Do() (*ListJobsResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["clusterName"]; ok {
		params.Set("clusterName", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["jobStateMatcher"]; ok {
		params.Set("jobStateMatcher", fmt.Sprintf("%v", v))
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/projects/{projectId}/jobs")
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
	var ret *ListJobsResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists jobs in a project.",
	//   "flatPath": "v1beta1/projects/{projectId}/jobs",
	//   "httpMethod": "GET",
	//   "id": "dataproc.projects.jobs.list",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "clusterName": {
	//       "description": "[Optional] If set, the returned jobs list includes only jobs that were\nsubmitted to the named cluster.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "jobStateMatcher": {
	//       "description": "[Optional] Specifies enumerated categories of jobs to list.",
	//       "enum": [
	//         "ALL",
	//         "ACTIVE",
	//         "NON_ACTIVE"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "[Optional] The number of results to return in each response.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "[Optional] The page token, returned by a previous call, to request the\nnext page of results.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "[Required] The ID of the Google Cloud Platform project that the job\nbelongs to.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/projects/{projectId}/jobs",
	//   "response": {
	//     "$ref": "ListJobsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "dataproc.projects.jobs.submit":

type ProjectsJobsSubmitCall struct {
	s                *Service
	projectId        string
	submitjobrequest *SubmitJobRequest
	opt_             map[string]interface{}
}

// Submit: Submits a job to a cluster.
func (r *ProjectsJobsService) Submit(projectId string, submitjobrequest *SubmitJobRequest) *ProjectsJobsSubmitCall {
	c := &ProjectsJobsSubmitCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.submitjobrequest = submitjobrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsJobsSubmitCall) Fields(s ...googleapi.Field) *ProjectsJobsSubmitCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsJobsSubmitCall) Do() (*Job, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.submitjobrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1beta1/projects/{projectId}/jobs:submit")
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
	var ret *Job
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Submits a job to a cluster.",
	//   "flatPath": "v1beta1/projects/{projectId}/jobs:submit",
	//   "httpMethod": "POST",
	//   "id": "dataproc.projects.jobs.submit",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "[Required] The ID of the Google Cloud Platform project that the job\nbelongs to.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1beta1/projects/{projectId}/jobs:submit",
	//   "request": {
	//     "$ref": "SubmitJobRequest"
	//   },
	//   "response": {
	//     "$ref": "Job"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}
