// Package dataflow provides access to the Google Dataflow API.
//
// See https://cloud.google.com/dataflow
//
// Usage example:
//
//   import "google.golang.org/api/dataflow/v1b3"
//   ...
//   dataflowService, err := dataflow.New(oauthHttpClient)
package dataflow

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

const apiId = "dataflow:v1b3"
const apiName = "dataflow"
const apiVersion = "v1b3"
const basePath = "https://dataflow.googleapis.com/"

// OAuth2 scopes used by this API.
const (
	// View and manage your data across Google Cloud Platform services
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"

	// View your email address
	UserinfoEmailScope = "https://www.googleapis.com/auth/userinfo.email"
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
	client   *http.Client
	BasePath string // API endpoint base URL

	Projects *ProjectsService
}

func NewProjectsService(s *Service) *ProjectsService {
	rs := &ProjectsService{s: s}
	rs.Jobs = NewProjectsJobsService(s)
	return rs
}

type ProjectsService struct {
	s *Service

	Jobs *ProjectsJobsService
}

func NewProjectsJobsService(s *Service) *ProjectsJobsService {
	rs := &ProjectsJobsService{s: s}
	rs.Messages = NewProjectsJobsMessagesService(s)
	rs.WorkItems = NewProjectsJobsWorkItemsService(s)
	return rs
}

type ProjectsJobsService struct {
	s *Service

	Messages *ProjectsJobsMessagesService

	WorkItems *ProjectsJobsWorkItemsService
}

func NewProjectsJobsMessagesService(s *Service) *ProjectsJobsMessagesService {
	rs := &ProjectsJobsMessagesService{s: s}
	return rs
}

type ProjectsJobsMessagesService struct {
	s *Service
}

func NewProjectsJobsWorkItemsService(s *Service) *ProjectsJobsWorkItemsService {
	rs := &ProjectsJobsWorkItemsService{s: s}
	return rs
}

type ProjectsJobsWorkItemsService struct {
	s *Service
}

type ApproximateProgress struct {
	// PercentComplete: Obsolete.
	PercentComplete float64 `json:"percentComplete,omitempty"`

	// Position: Obsolete.
	Position *Position `json:"position,omitempty"`

	// RemainingTime: Obsolete.
	RemainingTime string `json:"remainingTime,omitempty"`
}

type ApproximateReportedProgress struct {
	// FractionConsumed: Completion as fraction of the input consumed, from
	// 0.0 (beginning, nothing
	// consumed), to 1.0 (end of the input, entire
	// input consumed).
	FractionConsumed float64 `json:"fractionConsumed,omitempty"`

	// Position: A Position within the work to represent a progress.
	Position *Position `json:"position,omitempty"`
}

type ApproximateSplitRequest struct {
	// FractionConsumed: A fraction at which to split the work item, from
	// 0.0 (beginning of
	// the input) to 1.0 (end of the input).
	FractionConsumed float64 `json:"fractionConsumed,omitempty"`

	// Position: A Position at which to split the work item.
	Position *Position `json:"position,omitempty"`
}

type AutoscalingSettings struct {
	// Algorithm: The algorithm to use for autoscaling.
	Algorithm string `json:"algorithm,omitempty"`

	// MaxNumWorkers: The maximum number of workers to cap scaling at.
	MaxNumWorkers int64 `json:"maxNumWorkers,omitempty"`
}

type ComputationTopology struct {
	// ComputationId: The ID of the computation.
	ComputationId string `json:"computationId,omitempty"`

	// Inputs: The inputs to the computation.
	Inputs []*StreamLocation `json:"inputs,omitempty"`

	// KeyRanges: The key ranges processed by the computation.
	KeyRanges []*KeyRangeLocation `json:"keyRanges,omitempty"`

	// Outputs: The outputs from the computation.
	Outputs []*StreamLocation `json:"outputs,omitempty"`

	// StateFamilies: The state family values.
	StateFamilies []*StateFamilyConfig `json:"stateFamilies,omitempty"`

	// SystemStageName: The system stage name.
	SystemStageName string `json:"systemStageName,omitempty"`

	// UserStageName: The user stage name.
	UserStageName string `json:"userStageName,omitempty"`
}

type ConcatPosition struct {
	// Index: Index of the inner source.
	Index int64 `json:"index,omitempty"`

	// Position: Position within the inner source.
	Position *Position `json:"position,omitempty"`
}

type CustomSourceLocation struct {
	// Stateful: Whether this source is stateful.
	Stateful bool `json:"stateful,omitempty"`
}

type DataDiskAssignment struct {
	// DataDisks: Mounted data disks. The order is important a data disk's
	// 0-based index in
	// this list defines which persistent directory the
	// disk is mounted to, for
	// example the list of {
	// "myproject-1014-104817-4c2-harness-0-disk-0" },
	// {
	// "myproject-1014-104817-4c2-harness-0-disk-1" }.
	DataDisks []string `json:"dataDisks,omitempty"`

	// VmInstance: VM instance name the data disks mounted to, for
	// example
	// "myproject-1014-104817-4c2-harness-0".
	VmInstance string `json:"vmInstance,omitempty"`
}

type DerivedSource struct {
	// DerivationMode: What source to base the produced source on (if any).
	DerivationMode string `json:"derivationMode,omitempty"`

	// Source: Specification of the source.
	Source *Source `json:"source,omitempty"`
}

type Disk struct {
	// DiskType: Disk storage type, as defined by Google Compute Engine.
	// This
	// must be a disk type appropriate to the project and zone in
	// which
	// the workers will run.  If unknown or unspecified, the
	// service
	// will attempt to choose a reasonable default.
	//
	// For example,
	// the standard persistent disk type is a resource name
	// typically ending
	// in "pd-standard".  If SSD persistent disks are
	// available, the
	// resource name typically ends with "pd-ssd".  The
	// actual valid values
	// are defined the Google Compute Engine API,
	// not by the Dataflow API;
	// consult the Google Compute Engine
	// documentation for more information
	// about determining the set of
	// available disk types for a particular
	// project and zone.
	//
	// Google Compute Engine Disk types are local to a
	// particular
	// project in a particular zone, and so the resource name
	// will
	// typically look something like this:
	//
	//
	// compute.googleapis.com/projects/<project-id>/zones/<zone>/diskTypes/pd
	// -standard
	DiskType string `json:"diskType,omitempty"`

	// MountPoint: Directory in a VM where disk is mounted.
	MountPoint string `json:"mountPoint,omitempty"`

	// SizeGb: Size of disk in GB.  If zero or unspecified, the service
	// will
	// attempt to choose a reasonable default.
	SizeGb int64 `json:"sizeGb,omitempty"`
}

type DynamicSourceSplit struct {
	// Primary: Primary part (continued to be processed by
	// worker).
	// Specified relative to the previously-current source.
	// Becomes
	// current.
	Primary *DerivedSource `json:"primary,omitempty"`

	// Residual: Residual part (returned to the pool of work).
	// Specified
	// relative to the previously-current source.
	Residual *DerivedSource `json:"residual,omitempty"`
}

type Environment struct {
	// ClusterManagerApiService: The type of cluster manager API to use.  If
	// unknown or
	// unspecified, the service will attempt to choose a
	// reasonable
	// default.  This should be in the form of the API service
	// name,
	// e.g. "compute.googleapis.com".
	ClusterManagerApiService string `json:"clusterManagerApiService,omitempty"`

	// Dataset: The dataset for the current project where various
	// workflow
	// related tables are stored.
	//
	// The supported resource type
	// is:
	//
	// Google BigQuery:
	//   bigquery.googleapis.com/{dataset}
	Dataset string `json:"dataset,omitempty"`

	// Experiments: The list of experiments to enable.
	Experiments []string `json:"experiments,omitempty"`

	// InternalExperiments: Experimental settings.
	InternalExperiments *EnvironmentInternalExperiments `json:"internalExperiments,omitempty"`

	// SdkPipelineOptions: The Dataflow SDK pipeline options specified by
	// the user. These options
	// are passed through the service and are used
	// to recreate the SDK
	// pipeline options on the worker in a language
	// agnostic and platform
	// independent way.
	SdkPipelineOptions *EnvironmentSdkPipelineOptions `json:"sdkPipelineOptions,omitempty"`

	// TempStoragePrefix: The prefix of the resources the system should use
	// for temporary
	// storage.  The system will append the suffix
	// "/temp-{JOBNAME} to
	// this resource prefix, where {JOBNAME} is the
	// value of the
	// job_name field.  The resulting bucket and object prefix
	// is used
	// as the prefix of the resources used to store temporary
	// data
	// needed during the job execution.  NOTE: This will override
	// the
	// value in taskrunner_settings.
	// The supported resource type
	// is:
	//
	// Google Cloud Storage:
	//
	// storage.googleapis.com/{bucket}/{object}
	//
	// bucket.storage.googleapis.com/{object}
	TempStoragePrefix string `json:"tempStoragePrefix,omitempty"`

	// UserAgent: A description of the process that generated the request.
	UserAgent *EnvironmentUserAgent `json:"userAgent,omitempty"`

	// Version: A structure describing which components and their versions
	// of the service
	// are required in order to run the job.
	Version *EnvironmentVersion `json:"version,omitempty"`

	// WorkerPools: Worker pools.  At least one "harness" worker pool must
	// be
	// specified in order for the job to have workers.
	WorkerPools []*WorkerPool `json:"workerPools,omitempty"`
}

type EnvironmentInternalExperiments struct {
}

type EnvironmentSdkPipelineOptions struct {
}

type EnvironmentUserAgent struct {
}

type EnvironmentVersion struct {
}

type FlattenInstruction struct {
	// Inputs: Describes the inputs to the flatten instruction.
	Inputs []*InstructionInput `json:"inputs,omitempty"`
}

type InstructionInput struct {
	// OutputNum: The output index (origin zero) within the producer.
	OutputNum int64 `json:"outputNum,omitempty"`

	// ProducerInstructionIndex: The index (origin zero) of the parallel
	// instruction that produces
	// the output to be consumed by this input.
	// This index is relative
	// to the list of instructions in this input's
	// instruction's
	// containing MapTask.
	ProducerInstructionIndex int64 `json:"producerInstructionIndex,omitempty"`
}

type InstructionOutput struct {
	// Codec: The codec to use to encode data being written via this output.
	Codec *InstructionOutputCodec `json:"codec,omitempty"`

	// Name: The user-provided name of this output.
	Name string `json:"name,omitempty"`

	// SystemName: System-defined name of this output.
	// Unique across the
	// workflow.
	SystemName string `json:"systemName,omitempty"`
}

type InstructionOutputCodec struct {
}

type Job struct {
	// ClientRequestId: Client's unique identifier of the job, re-used by
	// SDK across
	// retried attempts.
	// If this field is set, the service will
	// ensure its uniqueness. That is,
	// the request to create a job will fail
	// if the service has knowledge of a
	// previously submitted job with the
	// same client's id and job name.
	// The caller may, for example, use this
	// field to ensure idempotence of job
	// creation across retried attempts
	// to create a job.
	// By default, the field is empty and, in that case,
	// the service ignores it.
	ClientRequestId string `json:"clientRequestId,omitempty"`

	// CreateTime: Timestamp when job was initially created. Immutable, set
	// by the Dataflow
	// service.
	CreateTime string `json:"createTime,omitempty"`

	// CurrentState: The current state of the job.
	//
	// Jobs are created in the
	// JOB_STATE_STOPPED state unless otherwise specified.
	//
	// A job in the
	// JOB_STATE_RUNNING state may asynchronously enter a
	// terminal state.
	// Once a job has reached a terminal state, no
	// further state updates may
	// be made.
	//
	// This field may be mutated by the Dataflow service; callers
	// cannot mutate it.
	CurrentState string `json:"currentState,omitempty"`

	// CurrentStateTime: The timestamp associated with the current state.
	CurrentStateTime string `json:"currentStateTime,omitempty"`

	// Environment: Environment for the job.
	Environment *Environment `json:"environment,omitempty"`

	// ExecutionInfo: Information about how the Dataflow service will
	// actually run the job.
	ExecutionInfo *JobExecutionInfo `json:"executionInfo,omitempty"`

	// Id: The unique ID of this job.
	//
	// This field is set by the Dataflow
	// service when the Job is
	// created, and is immutable for the life of the
	// Job.
	Id string `json:"id,omitempty"`

	// Name: The user-specified Dataflow job name.
	//
	// Only one Job with a
	// given name may exist in a project at any
	// given time.  If a caller
	// attempts to create a Job with the same
	// name as an already-existing
	// Job, the attempt will return the
	// existing Job.
	//
	// The name must match
	// the regular expression [a-z]([-a-z0-9]{0,38}[a-z0-9])?
	Name string `json:"name,omitempty"`

	// ProjectId: The project which owns the job.
	ProjectId string `json:"projectId,omitempty"`

	// ReplaceJobId: If this job is an update of an existing job, this field
	// will be the ID
	// of the job it replaced.
	//
	// When sending a
	// CreateJobRequest, you can update a job by specifying it
	// here. The job
	// named here will be stopped, and its intermediate state
	// transferred to
	// this job.
	ReplaceJobId string `json:"replaceJobId,omitempty"`

	// ReplacedByJobId: If another job is an update of this job (and thus,
	// this job is in
	// JOB_STATE_UPDATED), this field will contain the ID of
	// that job.
	ReplacedByJobId string `json:"replacedByJobId,omitempty"`

	// RequestedState: The job's requested state.
	//
	// UpdateJob may be used to
	// switch between the JOB_STATE_STOPPED and
	// JOB_STATE_RUNNING states, by
	// setting requested_state.  UpdateJob may also
	// be used to directly set
	// a job's requested state to JOB_STATE_CANCELLED or
	// JOB_STATE_DONE,
	// irrevocably terminating the job if it has not already
	// reached a
	// terminal state.
	RequestedState string `json:"requestedState,omitempty"`

	// Steps: The top-level steps that constitute the entire job.
	Steps []*Step `json:"steps,omitempty"`

	// TempFiles: A set of files the system should be aware of that are
	// used
	// for temporary storage. These temporary files will be
	// removed on
	// job completion.
	// No duplicates are allowed.
	// No file patterns are
	// supported.
	//
	// The supported files are:
	//
	// Google Cloud Storage:
	//
	// storage.googleapis.com/{bucket}/{object}
	//
	// bucket.storage.googleapis.com/{object}
	TempFiles []string `json:"tempFiles,omitempty"`

	// TransformNameMapping: Map of transform name prefixes of the job to be
	// replaced to the
	// corresponding name prefixes of the new job.
	TransformNameMapping map[string]string `json:"transformNameMapping,omitempty"`

	// Type: The type of dataflow job.
	Type string `json:"type,omitempty"`
}

type JobExecutionInfo struct {
	// Stages: A mapping from each stage to the information about that
	// stage.
	Stages map[string]JobExecutionStageInfo `json:"stages,omitempty"`
}

type JobExecutionStageInfo struct {
	// StepName: The steps associated with the execution stage.
	// Note that
	// stages may have several steps, and that a given step
	// might be run by
	// more than one stage.
	StepName []string `json:"stepName,omitempty"`
}

type JobMessage struct {
	// Id: Identifies the message.  This is automatically generated by
	// the
	// service; the caller should treat it as an opaque string.
	Id string `json:"id,omitempty"`

	// MessageImportance: Importance level of the message.
	MessageImportance string `json:"messageImportance,omitempty"`

	// MessageText: The text of the message.
	MessageText string `json:"messageText,omitempty"`

	// Time: The timestamp of the message.
	Time string `json:"time,omitempty"`
}

type JobMetrics struct {
	// MetricTime: Timestamp as of which metric values are current.
	MetricTime string `json:"metricTime,omitempty"`

	// Metrics: All metrics for this job.
	Metrics []*MetricUpdate `json:"metrics,omitempty"`
}

type KeyRangeDataDiskAssignment struct {
	// DataDisk: The name of the data disk where data for this range is
	// stored.
	// This name is local to the Google Cloud Platform project and
	// uniquely
	// identifies the disk within that project, for
	// example
	// "myproject-1014-104817-4c2-harness-0-disk-1".
	DataDisk string `json:"dataDisk,omitempty"`

	// End: The end (exclusive) of the key range.
	End string `json:"end,omitempty"`

	// Start: The start (inclusive) of the key range.
	Start string `json:"start,omitempty"`
}

type KeyRangeLocation struct {
	// DataDisk: The name of the data disk where data for this range is
	// stored.
	// This name is local to the Google Cloud Platform project and
	// uniquely
	// identifies the disk within that project, for
	// example
	// "myproject-1014-104817-4c2-harness-0-disk-1".
	DataDisk string `json:"dataDisk,omitempty"`

	// DeliveryEndpoint: The physical location of this range assignment to
	// be used for
	// streaming computation cross-worker message delivery.
	DeliveryEndpoint string `json:"deliveryEndpoint,omitempty"`

	// End: The end (exclusive) of the key range.
	End string `json:"end,omitempty"`

	// PersistentDirectory: The location of the persistent state for this
	// range, as a
	// persistent directory in the worker local filesystem.
	PersistentDirectory string `json:"persistentDirectory,omitempty"`

	// Start: The start (inclusive) of the key range.
	Start string `json:"start,omitempty"`
}

type LeaseWorkItemRequest struct {
	// CurrentWorkerTime: The current timestamp at the worker.
	CurrentWorkerTime string `json:"currentWorkerTime,omitempty"`

	// RequestedLeaseDuration: The initial lease period.
	RequestedLeaseDuration string `json:"requestedLeaseDuration,omitempty"`

	// WorkItemTypes: Filter for WorkItem type.
	WorkItemTypes []string `json:"workItemTypes,omitempty"`

	// WorkerCapabilities: Worker capabilities. WorkItems might be limited
	// to workers with specific
	// capabilities.
	WorkerCapabilities []string `json:"workerCapabilities,omitempty"`

	// WorkerId: Identifies the worker leasing work -- typically the ID of
	// the
	// virtual machine running the worker.
	WorkerId string `json:"workerId,omitempty"`
}

type LeaseWorkItemResponse struct {
	// WorkItems: A list of the leased WorkItems.
	WorkItems []*WorkItem `json:"workItems,omitempty"`
}

type ListJobMessagesResponse struct {
	// JobMessages: Messages in ascending timestamp order.
	JobMessages []*JobMessage `json:"jobMessages,omitempty"`

	// NextPageToken: The token to obtain the next page of results if there
	// are more.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type ListJobsResponse struct {
	// Jobs: A subset of the requested job information.
	Jobs []*Job `json:"jobs,omitempty"`

	// NextPageToken: Set if there may be more results than fit in this
	// response.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type MapTask struct {
	// Instructions: The instructions in the MapTask.
	Instructions []*ParallelInstruction `json:"instructions,omitempty"`

	// StageName: System-defined name of the stage containing this
	// MapTask.
	// Unique across the workflow.
	StageName string `json:"stageName,omitempty"`

	// SystemName: System-defined name of this MapTask.
	// Unique across the
	// workflow.
	SystemName string `json:"systemName,omitempty"`
}

type MetricStructuredName struct {
	// Context: Zero or more labeled fields which identify the part of the
	// job this
	// metric is associated with, such as the name of a step or
	// collection.
	//
	// For example, built-in counters associated with steps
	// will have
	// context['step'] = <step-name>. Counters associated with
	// PCollections
	// in the SDK will have context['pcollection'] =
	// <pcollection-name>.
	Context map[string]string `json:"context,omitempty"`

	// Name: Worker-defined metric name.
	Name string `json:"name,omitempty"`

	// Origin: Origin (namespace) of metric name. May be blank for
	// user-define metrics;
	// will be "dataflow" for metrics defined by the
	// Dataflow service or SDK.
	Origin string `json:"origin,omitempty"`
}

type MetricUpdate struct {
	// Cumulative: True if this metric is reported as the total cumulative
	// aggregate
	// value accumulated since the worker started working on this
	// WorkItem.
	// By default this is false, indicating that this metric is
	// reported
	// as a delta that is not associated with any WorkItem.
	Cumulative bool `json:"cumulative,omitempty"`

	// Internal: Worker-computed aggregate value for internal use by the
	// Dataflow
	// service.
	Internal interface{} `json:"internal,omitempty"`

	// Kind: Metric aggregation kind.  The possible metric aggregation kinds
	// are
	// "Sum", "Max", "Min", "Mean", "Set", "And", and "Or".
	// The
	// specified aggregation kind is case-insensitive.
	//
	// If omitted, this is
	// not an aggregated value but instead
	// a single metric sample value.
	Kind string `json:"kind,omitempty"`

	// MeanCount: Worker-computed aggregate value for the "Mean" aggregation
	// kind.
	// This holds the count of the aggregated values and is used in
	// combination
	// with mean_sum above to obtain the actual mean aggregate
	// value.
	// The only possible value type is Long.
	MeanCount interface{} `json:"meanCount,omitempty"`

	// MeanSum: Worker-computed aggregate value for the "Mean" aggregation
	// kind.
	// This holds the sum of the aggregated values and is used in
	// combination
	// with mean_count below to obtain the actual mean aggregate
	// value.
	// The only possible value types are Long and Double.
	MeanSum interface{} `json:"meanSum,omitempty"`

	// Name: Name of the metric.
	Name *MetricStructuredName `json:"name,omitempty"`

	// Scalar: Worker-computed aggregate value for aggregation kinds "Sum",
	// "Max", "Min",
	// "And", and "Or".  The possible value types are Long,
	// Double, and Boolean.
	Scalar interface{} `json:"scalar,omitempty"`

	// Set: Worker-computed aggregate value for the "Set" aggregation kind.
	// The only
	// possible value type is a list of Values whose type can be
	// Long, Double,
	// or String, according to the metric's type.  All Values
	// in the list must
	// be of the same type.
	Set interface{} `json:"set,omitempty"`

	// UpdateTime: Timestamp associated with the metric value. Optional when
	// workers are
	// reporting work progress; it will be filled in responses
	// from the
	// metrics API.
	UpdateTime string `json:"updateTime,omitempty"`
}

type MountedDataDisk struct {
	// DataDisk: The name of the data disk.
	// This name is local to the Google
	// Cloud Platform project and uniquely
	// identifies the disk within that
	// project, for example
	// "myproject-1014-104817-4c2-harness-0-disk-1".
	DataDisk string `json:"dataDisk,omitempty"`
}

type MultiOutputInfo struct {
	// Tag: The id of the tag the user code will emit to this output by;
	// this
	// should correspond to the tag of some SideInputInfo.
	Tag string `json:"tag,omitempty"`
}

type Package struct {
	// Location: The resource to read the package from.  The supported
	// resource type is:
	//
	// Google Cloud Storage:
	//
	// storage.googleapis.com/{bucket}
	//   bucket.storage.googleapis.com/
	Location string `json:"location,omitempty"`

	// Name: The name of the package.
	Name string `json:"name,omitempty"`
}

type ParDoInstruction struct {
	// Input: The input.
	Input *InstructionInput `json:"input,omitempty"`

	// MultiOutputInfos: Information about each of the outputs, if user_fn
	// is a  MultiDoFn.
	MultiOutputInfos []*MultiOutputInfo `json:"multiOutputInfos,omitempty"`

	// NumOutputs: The number of outputs.
	NumOutputs int64 `json:"numOutputs,omitempty"`

	// SideInputs: Zero or more side inputs.
	SideInputs []*SideInputInfo `json:"sideInputs,omitempty"`

	// UserFn: The user function to invoke.
	UserFn *ParDoInstructionUserFn `json:"userFn,omitempty"`
}

type ParDoInstructionUserFn struct {
}

type ParallelInstruction struct {
	// Flatten: Additional information for Flatten instructions.
	Flatten *FlattenInstruction `json:"flatten,omitempty"`

	// Name: User-provided name of this operation.
	Name string `json:"name,omitempty"`

	// Outputs: Describes the outputs of the instruction.
	Outputs []*InstructionOutput `json:"outputs,omitempty"`

	// ParDo: Additional information for ParDo instructions.
	ParDo *ParDoInstruction `json:"parDo,omitempty"`

	// PartialGroupByKey: Additional information for PartialGroupByKey
	// instructions.
	PartialGroupByKey *PartialGroupByKeyInstruction `json:"partialGroupByKey,omitempty"`

	// Read: Additional information for Read instructions.
	Read *ReadInstruction `json:"read,omitempty"`

	// SystemName: System-defined name of this operation.
	// Unique across the
	// workflow.
	SystemName string `json:"systemName,omitempty"`

	// Write: Additional information for Write instructions.
	Write *WriteInstruction `json:"write,omitempty"`
}

type PartialGroupByKeyInstruction struct {
	// Input: Describes the input to the partial group-by-key instruction.
	Input *InstructionInput `json:"input,omitempty"`

	// InputElementCodec: The codec to use for interpreting an element in
	// the input PTable.
	InputElementCodec *PartialGroupByKeyInstructionInputElementCodec `json:"inputElementCodec,omitempty"`

	// ValueCombiningFn: The value combining function to invoke.
	ValueCombiningFn *PartialGroupByKeyInstructionValueCombiningFn `json:"valueCombiningFn,omitempty"`
}

type PartialGroupByKeyInstructionInputElementCodec struct {
}

type PartialGroupByKeyInstructionValueCombiningFn struct {
}

type Position struct {
	// ByteOffset: Position is a byte offset.
	ByteOffset int64 `json:"byteOffset,omitempty,string"`

	// ConcatPosition: CloudPosition is a concat position.
	ConcatPosition *ConcatPosition `json:"concatPosition,omitempty"`

	// End: Position is past all other positions. Also useful for the
	// end
	// position of an unbounded range.
	End bool `json:"end,omitempty"`

	// Key: Position is a string key, ordered lexicographically.
	Key string `json:"key,omitempty"`

	// RecordIndex: Position is a record index.
	RecordIndex int64 `json:"recordIndex,omitempty,string"`

	// ShufflePosition: CloudPosition is a base64 encoded
	// BatchShufflePosition (with FIXED
	// sharding).
	ShufflePosition string `json:"shufflePosition,omitempty"`
}

type PubsubLocation struct {
	// DropLateData: Indicates whether the pipeline allows late-arriving
	// data.
	DropLateData bool `json:"dropLateData,omitempty"`

	// IdLabel: If set, contains a pubsub label from which to extract record
	// ids.
	// If left empty, record deduplication will be strictly best
	// effort.
	IdLabel string `json:"idLabel,omitempty"`

	// Subscription: A pubsub subscription, in the form
	// of
	// "pubsub.googleapis.com/subscriptions/<project-id>/<subscription-nam
	// e>"
	Subscription string `json:"subscription,omitempty"`

	// TimestampLabel: If set, contains a pubsub label from which to extract
	// record timestamps.
	// If left empty, record timestamps will be generated
	// upon arrival.
	TimestampLabel string `json:"timestampLabel,omitempty"`

	// Topic: A pubsub topic, in the form
	// of
	// "pubsub.googleapis.com/topics/<project-id>/<topic-name>"
	Topic string `json:"topic,omitempty"`

	// TrackingSubscription: If set, specifies the pubsub subscription that
	// will be used for tracking
	// custom time timestamps for watermark
	// estimation.
	TrackingSubscription string `json:"trackingSubscription,omitempty"`
}

type ReadInstruction struct {
	// Source: The source to read from.
	Source *Source `json:"source,omitempty"`
}

type ReportWorkItemStatusRequest struct {
	// CurrentWorkerTime: The current timestamp at the worker.
	CurrentWorkerTime string `json:"currentWorkerTime,omitempty"`

	// WorkItemStatuses: The order is unimportant, except that the order of
	// the
	// WorkItemServiceState messages in the
	// ReportWorkItemStatusResponse
	// corresponds to the order of
	// WorkItemStatus messages here.
	WorkItemStatuses []*WorkItemStatus `json:"workItemStatuses,omitempty"`

	// WorkerId: The ID of the worker reporting the WorkItem status.  If
	// this
	// does not match the ID of the worker which the Dataflow
	// service
	// believes currently has the lease on the WorkItem, the
	// report
	// will be dropped (with an error response).
	WorkerId string `json:"workerId,omitempty"`
}

type ReportWorkItemStatusResponse struct {
	// WorkItemServiceStates: A set of messages indicating the service-side
	// state for each
	// WorkItem whose status was reported, in the same order
	// as the
	// WorkItemStatus messages in the ReportWorkItemStatusRequest
	// which
	// resulting in this response.
	WorkItemServiceStates []*WorkItemServiceState `json:"workItemServiceStates,omitempty"`
}

type SendWorkerMessagesRequest struct {
	// WorkerMessages: The WorkerMessages to send.
	WorkerMessages []*WorkerMessage `json:"workerMessages,omitempty"`
}

type SendWorkerMessagesResponse struct {
	// WorkerMessageResponses: The servers response to the worker messages.
	WorkerMessageResponses []*WorkerMessageResponse `json:"workerMessageResponses,omitempty"`
}

type SeqMapTask struct {
	// Inputs: Information about each of the inputs.
	Inputs []*SideInputInfo `json:"inputs,omitempty"`

	// Name: The user-provided name of the SeqDo operation.
	Name string `json:"name,omitempty"`

	// OutputInfos: Information about each of the outputs.
	OutputInfos []*SeqMapTaskOutputInfo `json:"outputInfos,omitempty"`

	// StageName: System-defined name of the stage containing the SeqDo
	// operation.
	// Unique across the workflow.
	StageName string `json:"stageName,omitempty"`

	// SystemName: System-defined name of the SeqDo operation.
	// Unique across
	// the workflow.
	SystemName string `json:"systemName,omitempty"`

	// UserFn: The user function to invoke.
	UserFn *SeqMapTaskUserFn `json:"userFn,omitempty"`
}

type SeqMapTaskUserFn struct {
}

type SeqMapTaskOutputInfo struct {
	// Sink: The sink to write the output value to.
	Sink *Sink `json:"sink,omitempty"`

	// Tag: The id of the TupleTag the user code will tag the output value
	// by.
	Tag string `json:"tag,omitempty"`
}

type ShellTask struct {
	// Command: The shell command to run.
	Command string `json:"command,omitempty"`

	// ExitCode: Exit code for the task.
	ExitCode int64 `json:"exitCode,omitempty"`
}

type SideInputInfo struct {
	// Kind: How to interpret the source element(s) as a side input value.
	Kind *SideInputInfoKind `json:"kind,omitempty"`

	// Sources: The source(s) to read element(s) from to get the value of
	// this side input.
	// If more than one source, then the elements are taken
	// from the
	// sources, in the specified order if order matters.
	// At least
	// one source is required.
	Sources []*Source `json:"sources,omitempty"`

	// Tag: The id of the tag the user code will access this side input
	// by;
	// this should correspond to the tag of some MultiOutputInfo.
	Tag string `json:"tag,omitempty"`
}

type SideInputInfoKind struct {
}

type Sink struct {
	// Codec: The codec to use to encode data written to the sink.
	Codec *SinkCodec `json:"codec,omitempty"`

	// Spec: The sink to write to, plus its parameters.
	Spec *SinkSpec `json:"spec,omitempty"`
}

type SinkCodec struct {
}

type SinkSpec struct {
}

type Source struct {
	// BaseSpecs: While splitting, sources may specify the produced
	// bundles
	// as differences against another source, in order to save
	// backend-side
	// memory and allow bigger jobs. For details, see
	// SourceSplitRequest.
	// To support this use case, the full set of
	// parameters of the source
	// is logically obtained by taking the latest
	// explicitly specified value
	// of each parameter in the order:
	// base_specs
	// (later items win), spec (overrides anything in base_specs).
	BaseSpecs []*SourceBaseSpecs `json:"baseSpecs,omitempty"`

	// Codec: The codec to use to decode data read from the source.
	Codec *SourceCodec `json:"codec,omitempty"`

	// DoesNotNeedSplitting: Setting this value to true hints to the
	// framework that the source
	// doesn't need splitting, and using
	// SourceSplitRequest on it would
	// yield
	// SOURCE_SPLIT_OUTCOME_USE_CURRENT.
	//
	// E.g. a file splitter may set this
	// to true when splitting a single file
	// into a set of byte ranges of
	// appropriate size, and set this
	// to false when splitting a filepattern
	// into individual files.
	// However, for efficiency, a file splitter may
	// decide to produce
	// file subranges directly from the filepattern to
	// avoid a splitting
	// round-trip.
	//
	// See SourceSplitRequest for an overview
	// of the splitting process.
	//
	// This field is meaningful only in the
	// Source objects populated
	// by the user (e.g. when filling in a
	// DerivedSource).
	// Source objects supplied by the framework to the user
	// don't have
	// this field populated.
	DoesNotNeedSplitting bool `json:"doesNotNeedSplitting,omitempty"`

	// Metadata: Optionally, metadata for this source can be supplied right
	// away,
	// avoiding a SourceGetMetadataOperation roundtrip
	// (see
	// SourceOperationRequest).
	//
	// This field is meaningful only in the Source
	// objects populated
	// by the user (e.g. when filling in a
	// DerivedSource).
	// Source objects supplied by the framework to the user
	// don't have
	// this field populated.
	Metadata *SourceMetadata `json:"metadata,omitempty"`

	// Spec: The source to read from, plus its parameters.
	Spec *SourceSpec `json:"spec,omitempty"`
}

type SourceBaseSpecs struct {
}

type SourceCodec struct {
}

type SourceSpec struct {
}

type SourceFork struct {
	// Primary: DEPRECATED
	Primary *SourceSplitShard `json:"primary,omitempty"`

	// PrimarySource: DEPRECATED
	PrimarySource *DerivedSource `json:"primarySource,omitempty"`

	// Residual: DEPRECATED
	Residual *SourceSplitShard `json:"residual,omitempty"`

	// ResidualSource: DEPRECATED
	ResidualSource *DerivedSource `json:"residualSource,omitempty"`
}

type SourceGetMetadataRequest struct {
	// Source: Specification of the source whose metadata should be
	// computed.
	Source *Source `json:"source,omitempty"`
}

type SourceGetMetadataResponse struct {
	// Metadata: The computed metadata.
	Metadata *SourceMetadata `json:"metadata,omitempty"`
}

type SourceMetadata struct {
	// EstimatedSizeBytes: An estimate of the total size (in bytes) of the
	// data that would be
	// read from this source.  This estimate is in terms
	// of external storage
	// size, before any decompression or other
	// processing done by the reader.
	EstimatedSizeBytes int64 `json:"estimatedSizeBytes,omitempty,string"`

	// Infinite: Specifies that the size of this source is known to be
	// infinite
	// (this is a streaming source).
	Infinite bool `json:"infinite,omitempty"`

	// ProducesSortedKeys: Whether this source is known to produce key/value
	// pairs with
	// the (encoded) keys in lexicographically sorted order.
	ProducesSortedKeys bool `json:"producesSortedKeys,omitempty"`
}

type SourceOperationRequest struct {
	// GetMetadata: Information about a request to get metadata about a
	// source.
	GetMetadata *SourceGetMetadataRequest `json:"getMetadata,omitempty"`

	// Split: Information about a request to split a source.
	Split *SourceSplitRequest `json:"split,omitempty"`
}

type SourceOperationResponse struct {
	// GetMetadata: A response to a request to get metadata about a source.
	GetMetadata *SourceGetMetadataResponse `json:"getMetadata,omitempty"`

	// Split: A response to a request to split a source.
	Split *SourceSplitResponse `json:"split,omitempty"`
}

type SourceSplitOptions struct {
	// DesiredBundleSizeBytes: The source should be split into a set of
	// bundles where the estimated size
	// of each is approximately this many
	// bytes.
	DesiredBundleSizeBytes int64 `json:"desiredBundleSizeBytes,omitempty,string"`

	// DesiredShardSizeBytes: DEPRECATED in favor of
	// desired_bundle_size_bytes.
	DesiredShardSizeBytes int64 `json:"desiredShardSizeBytes,omitempty,string"`
}

type SourceSplitRequest struct {
	// Options: Hints for tuning the splitting process.
	Options *SourceSplitOptions `json:"options,omitempty"`

	// Source: Specification of the source to be split.
	Source *Source `json:"source,omitempty"`
}

type SourceSplitResponse struct {
	// Bundles: If outcome is SPLITTING_HAPPENED, then this is a list of
	// bundles
	// into which the source was split. Otherwise this field is
	// ignored.
	// This list can be empty, which means the source represents an
	// empty input.
	Bundles []*DerivedSource `json:"bundles,omitempty"`

	// Outcome: Indicates whether splitting happened and produced a list of
	// bundles.
	// If this is USE_CURRENT_SOURCE_AS_IS, the current source
	// should
	// be processed "as is" without splitting. "bundles" is ignored
	// in this case.
	// If this is SPLITTING_HAPPENED, then "bundles" contains
	// a list of
	// bundles into which the source was split.
	Outcome string `json:"outcome,omitempty"`

	// Shards: DEPRECATED in favor of bundles.
	Shards []*SourceSplitShard `json:"shards,omitempty"`
}

type SourceSplitShard struct {
	// DerivationMode: DEPRECATED
	DerivationMode string `json:"derivationMode,omitempty"`

	// Source: DEPRECATED
	Source *Source `json:"source,omitempty"`
}

type StateFamilyConfig struct {
	// IsRead: If true, this family corresponds to a read operation.
	IsRead bool `json:"isRead,omitempty"`

	// StateFamily: The state family value.
	StateFamily string `json:"stateFamily,omitempty"`
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

type Step struct {
	// Kind: The kind of step in the dataflow Job.
	Kind string `json:"kind,omitempty"`

	// Name: Name identifying the step. This must be unique for each
	// step
	// with respect to all other steps in the dataflow Job.
	Name string `json:"name,omitempty"`

	// Properties: Named properties associated with the step.  Each kind
	// of
	// predefined step has its own required set of properties.
	Properties *StepProperties `json:"properties,omitempty"`
}

type StepProperties struct {
}

type StreamLocation struct {
	// CustomSourceLocation: The stream is a custom source.
	CustomSourceLocation *CustomSourceLocation `json:"customSourceLocation,omitempty"`

	// PubsubLocation: The stream is a pubsub stream.
	PubsubLocation *PubsubLocation `json:"pubsubLocation,omitempty"`

	// SideInputLocation: The stream is a streaming side input.
	SideInputLocation *StreamingSideInputLocation `json:"sideInputLocation,omitempty"`

	// StreamingStageLocation: The stream is part of another computation
	// within the current
	// streaming Dataflow job.
	StreamingStageLocation *StreamingStageLocation `json:"streamingStageLocation,omitempty"`
}

type StreamingComputationRanges struct {
	// ComputationId: The ID of the computation.
	ComputationId string `json:"computationId,omitempty"`

	// RangeAssignments: Data disk assignments for ranges from this
	// computation.
	RangeAssignments []*KeyRangeDataDiskAssignment `json:"rangeAssignments,omitempty"`
}

type StreamingComputationTask struct {
	// ComputationRanges: Contains ranges of a streaming computation this
	// task should apply to.
	ComputationRanges []*StreamingComputationRanges `json:"computationRanges,omitempty"`

	// DataDisks: Describes the set of data disks this task should apply to.
	DataDisks []*MountedDataDisk `json:"dataDisks,omitempty"`

	// TaskType: A type of streaming computation task.
	TaskType string `json:"taskType,omitempty"`
}

type StreamingSetupTask struct {
	// ReceiveWorkPort: The TCP port on which the worker should listen for
	// messages from
	// other streaming computation workers.
	ReceiveWorkPort int64 `json:"receiveWorkPort,omitempty"`

	// StreamingComputationTopology: The global topology of the streaming
	// Dataflow job.
	StreamingComputationTopology *TopologyConfig `json:"streamingComputationTopology,omitempty"`

	// WorkerHarnessPort: The TCP port used by the worker to communicate
	// with the Dataflow
	// worker harness.
	WorkerHarnessPort int64 `json:"workerHarnessPort,omitempty"`
}

type StreamingSideInputLocation struct {
	// StateFamily: Identifies the state family where this side input is
	// stored.
	StateFamily string `json:"stateFamily,omitempty"`

	// Tag: Identifies the particular side input within the streaming
	// Dataflow job.
	Tag string `json:"tag,omitempty"`
}

type StreamingStageLocation struct {
	// StreamId: Identifies the particular stream within the streaming
	// Dataflow
	// job.
	StreamId string `json:"streamId,omitempty"`
}

type TaskRunnerSettings struct {
	// Alsologtostderr: Also send taskrunner log info to stderr?
	Alsologtostderr bool `json:"alsologtostderr,omitempty"`

	// BaseTaskDir: Location on the worker for task-specific subdirectories.
	BaseTaskDir string `json:"baseTaskDir,omitempty"`

	// BaseUrl: The base URL for the taskrunner to use when accessing Google
	// Cloud APIs.
	//
	// When workers access Google Cloud APIs, they logically do
	// so via
	// relative URLs.  If this field is specified, it supplies the
	// base
	// URL to use for resolving these relative URLs.  The
	// normative
	// algorithm used is defined by RFC 1808, "Relative Uniform
	// Resource
	// Locators".
	//
	// If not specified, the default value is
	// "http://www.googleapis.com/"
	BaseUrl string `json:"baseUrl,omitempty"`

	// CommandlinesFileName: Store preprocessing commands in this file.
	CommandlinesFileName string `json:"commandlinesFileName,omitempty"`

	// ContinueOnException: Do we continue taskrunner if an exception is
	// hit?
	ContinueOnException bool `json:"continueOnException,omitempty"`

	// DataflowApiVersion: API version of endpoint, e.g. "v1b3"
	DataflowApiVersion string `json:"dataflowApiVersion,omitempty"`

	// HarnessCommand: Command to launch the worker harness.
	HarnessCommand string `json:"harnessCommand,omitempty"`

	// LanguageHint: Suggested backend language.
	LanguageHint string `json:"languageHint,omitempty"`

	// LogDir: Directory on the VM to store logs.
	LogDir string `json:"logDir,omitempty"`

	// LogToSerialconsole: Send taskrunner log into to Google Compute Engine
	// VM serial console?
	LogToSerialconsole bool `json:"logToSerialconsole,omitempty"`

	// LogUploadLocation: Indicates where to put logs.  If this is not
	// specified, the logs
	// will not be uploaded.
	//
	// The supported resource
	// type is:
	//
	// Google Cloud Storage:
	//
	// storage.googleapis.com/{bucket}/{object}
	//
	// bucket.storage.googleapis.com/{object}
	LogUploadLocation string `json:"logUploadLocation,omitempty"`

	// OauthScopes: OAuth2 scopes to be requested by the taskrunner in order
	// to
	// access the dataflow API.
	OauthScopes []string `json:"oauthScopes,omitempty"`

	// ParallelWorkerSettings: Settings to pass to the parallel worker
	// harness.
	ParallelWorkerSettings *WorkerSettings `json:"parallelWorkerSettings,omitempty"`

	// StreamingWorkerMainClass: Streaming worker main class name.
	StreamingWorkerMainClass string `json:"streamingWorkerMainClass,omitempty"`

	// TaskGroup: The UNIX group ID on the worker VM to use for tasks
	// launched by
	// taskrunner; e.g. "wheel".
	TaskGroup string `json:"taskGroup,omitempty"`

	// TaskUser: The UNIX user ID on the worker VM to use for tasks launched
	// by
	// taskrunner; e.g. "root".
	TaskUser string `json:"taskUser,omitempty"`

	// TempStoragePrefix: The prefix of the resources the taskrunner should
	// use for
	// temporary storage.
	//
	// The supported resource type is:
	//
	// Google
	// Cloud Storage:
	//   storage.googleapis.com/{bucket}/{object}
	//
	// bucket.storage.googleapis.com/{object}
	TempStoragePrefix string `json:"tempStoragePrefix,omitempty"`

	// VmId: ID string of VM.
	VmId string `json:"vmId,omitempty"`

	// WorkflowFileName: Store the workflow in this file.
	WorkflowFileName string `json:"workflowFileName,omitempty"`
}

type TopologyConfig struct {
	// Computations: The computations associated with a streaming Dataflow
	// job.
	Computations []*ComputationTopology `json:"computations,omitempty"`

	// DataDiskAssignments: The disks assigned to a streaming Dataflow job.
	DataDiskAssignments []*DataDiskAssignment `json:"dataDiskAssignments,omitempty"`

	// UserStageToComputationNameMap: Maps user stage names to stable
	// computation names.
	UserStageToComputationNameMap map[string]string `json:"userStageToComputationNameMap,omitempty"`
}

type WorkItem struct {
	// Configuration: Work item-specific configuration as an opaque blob.
	Configuration string `json:"configuration,omitempty"`

	// Id: Identifies this WorkItem.
	Id int64 `json:"id,omitempty,string"`

	// InitialReportIndex: The initial index to use when reporting the
	// status of the WorkItem.
	InitialReportIndex int64 `json:"initialReportIndex,omitempty,string"`

	// JobId: Identifies the workflow job this WorkItem belongs to.
	JobId string `json:"jobId,omitempty"`

	// LeaseExpireTime: Time when the lease on this Work will expire.
	LeaseExpireTime string `json:"leaseExpireTime,omitempty"`

	// MapTask: Additional information for MapTask WorkItems.
	MapTask *MapTask `json:"mapTask,omitempty"`

	// Packages: Any required packages that need to be fetched in order to
	// execute
	// this WorkItem.
	Packages []*Package `json:"packages,omitempty"`

	// ProjectId: Identifies the cloud project this WorkItem belongs to.
	ProjectId string `json:"projectId,omitempty"`

	// ReportStatusInterval: Recommended reporting interval.
	ReportStatusInterval string `json:"reportStatusInterval,omitempty"`

	// SeqMapTask: Additional information for SeqMapTask WorkItems.
	SeqMapTask *SeqMapTask `json:"seqMapTask,omitempty"`

	// ShellTask: Additional information for ShellTask WorkItems.
	ShellTask *ShellTask `json:"shellTask,omitempty"`

	// SourceOperationTask: Additional information for source operation
	// WorkItems.
	SourceOperationTask *SourceOperationRequest `json:"sourceOperationTask,omitempty"`

	// StreamingComputationTask: Additional information for
	// StreamingComputationTask WorkItems.
	StreamingComputationTask *StreamingComputationTask `json:"streamingComputationTask,omitempty"`

	// StreamingSetupTask: Additional information for StreamingSetupTask
	// WorkItems.
	StreamingSetupTask *StreamingSetupTask `json:"streamingSetupTask,omitempty"`
}

type WorkItemServiceState struct {
	// HarnessData: Other data returned by the service, specific to the
	// particular
	// worker harness.
	HarnessData *WorkItemServiceStateHarnessData `json:"harnessData,omitempty"`

	// LeaseExpireTime: Time at which the current lease will expire.
	LeaseExpireTime string `json:"leaseExpireTime,omitempty"`

	// NextReportIndex: The index value to use for the next report sent by
	// the worker.
	// Note: If the report call fails for whatever reason, the
	// worker should
	// reuse this index for subsequent report attempts.
	NextReportIndex int64 `json:"nextReportIndex,omitempty,string"`

	// ReportStatusInterval: New recommended reporting interval.
	ReportStatusInterval string `json:"reportStatusInterval,omitempty"`

	// SplitRequest: The progress point in the WorkItem where the Dataflow
	// service
	// suggests that the worker truncate the task.
	SplitRequest *ApproximateSplitRequest `json:"splitRequest,omitempty"`

	// SuggestedStopPoint: DEPRECATED in favor of split_request.
	SuggestedStopPoint *ApproximateProgress `json:"suggestedStopPoint,omitempty"`

	// SuggestedStopPosition: Obsolete, always empty.
	SuggestedStopPosition *Position `json:"suggestedStopPosition,omitempty"`
}

type WorkItemServiceStateHarnessData struct {
}

type WorkItemStatus struct {
	// Completed: True if the WorkItem was completed (successfully or
	// unsuccessfully).
	Completed bool `json:"completed,omitempty"`

	// DynamicSourceSplit: See documentation of stop_position.
	DynamicSourceSplit *DynamicSourceSplit `json:"dynamicSourceSplit,omitempty"`

	// Errors: Specifies errors which occurred during processing.  If errors
	// are
	// provided, and completed = true, then the WorkItem is
	// considered
	// to have failed.
	Errors []*Status `json:"errors,omitempty"`

	// MetricUpdates: Worker output metrics (counters) for this WorkItem.
	MetricUpdates []*MetricUpdate `json:"metricUpdates,omitempty"`

	// Progress: DEPRECATED in favor of reported_progress.
	Progress *ApproximateProgress `json:"progress,omitempty"`

	// ReportIndex: The report index.  When a WorkItem is leased, the lease
	// will
	// contain an initial report index.  When a WorkItem's status
	// is
	// reported to the system, the report should be sent with
	// that report
	// index, and the response will contain the index the
	// worker should use
	// for the next report.  Reports received with
	// unexpected index values
	// will be rejected by the service.
	//
	// In order to preserve idempotency,
	// the worker should not alter the
	// contents of a report, even if the
	// worker must submit the same
	// report multiple times before getting back
	// a response.  The worker
	// should not submit a subsequent report until
	// the response for the
	// previous report had been received from the
	// service.
	ReportIndex int64 `json:"reportIndex,omitempty,string"`

	// ReportedProgress: The worker's progress through this WorkItem.
	ReportedProgress *ApproximateReportedProgress `json:"reportedProgress,omitempty"`

	// RequestedLeaseDuration: Amount of time the worker requests for its
	// lease.
	RequestedLeaseDuration string `json:"requestedLeaseDuration,omitempty"`

	// SourceFork: DEPRECATED in favor of dynamic_source_split.
	SourceFork *SourceFork `json:"sourceFork,omitempty"`

	// SourceOperationResponse: If the work item represented a
	// SourceOperationRequest, and the work
	// is completed, contains the
	// result of the operation.
	SourceOperationResponse *SourceOperationResponse `json:"sourceOperationResponse,omitempty"`

	// StopPosition: A worker may split an active map task in two parts,
	// "primary" and
	// "residual", continuing to process the primary part and
	// returning the
	// residual part into the pool of available work.
	// This
	// event is called a "dynamic split" and is critical to the dynamic
	// work
	// rebalancing feature. The two obtained sub-tasks are called
	// "parts" of
	// the split.
	// The parts, if concatenated, must represent the same input
	// as would
	// be read by the current task if the split did not happen.
	// The
	// exact way in which the original task is decomposed into the two
	// parts
	// is specified either as a position demarcating them
	// (stop_position),
	// or explicitly as two DerivedSources, if this
	// task consumes a
	// user-defined source type (dynamic_source_split).
	//
	// The "current" task
	// is adjusted as a result of the split: after a task
	// with range [A, B)
	// sends a stop_position update at C, its range is
	// considered to be [A,
	// C), e.g.:
	// * Progress should be interpreted relative to the new range,
	// e.g.
	//   "75% completed" means "75% of [A, C) completed"
	// * The worker
	// should interpret proposed_stop_position relative to the
	//   new range,
	// e.g. "split at 68%" should be interpreted as
	//   "split at 68% of [A,
	// C)".
	// * If the worker chooses to split again using stop_position,
	// only
	//   stop_positions in [A, C) will be accepted.
	// *
	// Etc.
	// dynamic_source_split has similar semantics: e.g., if a task
	// with
	// source S splits using dynamic_source_split into {P, R}
	// (where P
	// and R must be together equivalent to S), then subsequent
	// progress and
	// proposed_stop_position should be interpreted relative
	// to P, and in a
	// potential subsequent dynamic_source_split into {P', R'},
	// P' and R'
	// must be together equivalent to P, etc.
	StopPosition *Position `json:"stopPosition,omitempty"`

	// WorkItemId: Identifies the WorkItem.
	WorkItemId string `json:"workItemId,omitempty"`
}

type WorkerHealthReport struct {
	// ReportInterval: The interval at which the worker is sending health
	// reports.
	// The default value of 0 should be interpreted as the field is
	// not being
	// explicitly set by the worker.
	ReportInterval string `json:"reportInterval,omitempty"`

	// VmIsHealthy: Whether the VM is healthy.
	VmIsHealthy bool `json:"vmIsHealthy,omitempty"`

	// VmStartupTime: The time the VM was booted.
	VmStartupTime string `json:"vmStartupTime,omitempty"`
}

type WorkerHealthReportResponse struct {
	// ReportInterval: A positive value indicates the worker should change
	// its reporting interval
	// to the specified value.
	//
	// The default value of
	// zero means no change in report rate is requested by
	// the server.
	ReportInterval string `json:"reportInterval,omitempty"`
}

type WorkerMessage struct {
	// Labels: Labels are used to group WorkerMessages.
	// For example, a
	// worker_message about a particular container
	// might have the labels:
	// {
	// "JOB_ID": "2015-04-22",
	//   "WORKER_ID": "wordcount-vm-2015…"
	//
	// "CONTAINER_TYPE": "worker",
	//   "CONTAINER_ID": "ac1234def"}
	// Label tags
	// typically correspond to Label enum values. However, for ease
	// of
	// development other strings can be used as tags. LABEL_UNSPECIFIED
	// should
	// not be used here.
	Labels map[string]string `json:"labels,omitempty"`

	// Time: The timestamp of the worker_message.
	Time string `json:"time,omitempty"`

	// WorkerHealthReport: The health of a worker.
	WorkerHealthReport *WorkerHealthReport `json:"workerHealthReport,omitempty"`
}

type WorkerMessageResponse struct {
	// WorkerHealthReportResponse: The service's response to a worker's
	// health report.
	WorkerHealthReportResponse *WorkerHealthReportResponse `json:"workerHealthReportResponse,omitempty"`
}

type WorkerPool struct {
	// AutoscalingSettings: Settings for autoscaling of this WorkerPool.
	AutoscalingSettings *AutoscalingSettings `json:"autoscalingSettings,omitempty"`

	// DataDisks: Data disks that are used by a VM in this workflow.
	DataDisks []*Disk `json:"dataDisks,omitempty"`

	// DefaultPackageSet: The default package set to install.  This allows
	// the service to
	// select a default set of packages which are useful to
	// worker
	// harnesses written in a particular language.
	DefaultPackageSet string `json:"defaultPackageSet,omitempty"`

	// DiskSizeGb: Size of root disk for VMs, in GB.  If zero or
	// unspecified, the service will
	// attempt to choose a reasonable default.
	DiskSizeGb int64 `json:"diskSizeGb,omitempty"`

	// DiskSourceImage: Fully qualified source image for disks.
	DiskSourceImage string `json:"diskSourceImage,omitempty"`

	// DiskType: Type of root disk for VMs.  If empty or unspecified, the
	// service will
	// attempt to choose a reasonable default.
	DiskType string `json:"diskType,omitempty"`

	// Kind: The kind of the worker pool; currently only 'harness' and
	// 'shuffle'
	// are supported.
	Kind string `json:"kind,omitempty"`

	// MachineType: Machine type (e.g. "n1-standard-1").  If empty or
	// unspecified, the
	// service will attempt to choose a reasonable default.
	MachineType string `json:"machineType,omitempty"`

	// Metadata: Metadata to set on the Google Compute Engine VMs.
	Metadata map[string]string `json:"metadata,omitempty"`

	// Network: Network to which VMs will be assigned.  If empty or
	// unspecified,
	// the service will use the network "default".
	Network string `json:"network,omitempty"`

	// NumWorkers: Number of Google Compute Engine workers in this pool
	// needed to
	// execute the job.  If zero or unspecified, the service
	// will
	// attempt to choose a reasonable default.
	NumWorkers int64 `json:"numWorkers,omitempty"`

	// OnHostMaintenance: The action to take on host maintenance, as defined
	// by the Google
	// Compute Engine API.
	OnHostMaintenance string `json:"onHostMaintenance,omitempty"`

	// Packages: Packages to be installed on workers.
	Packages []*Package `json:"packages,omitempty"`

	// PoolArgs: Extra arguments for this worker pool.
	PoolArgs *WorkerPoolPoolArgs `json:"poolArgs,omitempty"`

	// TaskrunnerSettings: Settings passed through to Google Compute Engine
	// workers when
	// using the standard Dataflow task runner.  Users should
	// ignore
	// this field.
	TaskrunnerSettings *TaskRunnerSettings `json:"taskrunnerSettings,omitempty"`

	// TeardownPolicy: Sets the policy for determining when to turndown
	// worker pool.
	// Allowed values are: TEARDOWN_ALWAYS,
	// TEARDOWN_ON_SUCCESS, and
	// TEARDOWN_NEVER.
	// TEARDOWN_ALWAYS means
	// workers are always torn down regardless of whether
	// the job succeeds.
	// TEARDOWN_ON_SUCCESS means workers are torn down
	// if the job succeeds.
	// TEARDOWN_NEVER means the workers are never torn
	// down.
	//
	// If the workers
	// are not torn down by the service, they will
	// continue to run and use
	// Google Compute Engine VM resources in the
	// user's project until they
	// are explicitly terminated by the user.
	// Because of this, Google
	// recommends using the TEARDOWN_ALWAYS
	// policy except for small,
	// manually supervised test jobs.
	//
	// If unknown or unspecified, the
	// service will attempt to choose a reasonable
	// default.
	TeardownPolicy string `json:"teardownPolicy,omitempty"`

	// Zone: Zone to run the worker pools in (e.g. "us-central1-a").
	// If
	// empty or unspecified, the service will attempt to choose a
	// reasonable
	// default.
	Zone string `json:"zone,omitempty"`
}

type WorkerPoolPoolArgs struct {
}

type WorkerSettings struct {
	// BaseUrl: The base URL for accessing Google Cloud APIs.
	//
	// When workers
	// access Google Cloud APIs, they logically do so via
	// relative URLs.  If
	// this field is specified, it supplies the base
	// URL to use for
	// resolving these relative URLs.  The normative
	// algorithm used is
	// defined by RFC 1808, "Relative Uniform Resource
	// Locators".
	//
	// If not
	// specified, the default value is "http://www.googleapis.com/"
	BaseUrl string `json:"baseUrl,omitempty"`

	// ReportingEnabled: Send work progress updates to service.
	ReportingEnabled bool `json:"reportingEnabled,omitempty"`

	// ServicePath: The Dataflow service path relative to the root URL, for
	// example,
	// "dataflow/v1b3/projects".
	ServicePath string `json:"servicePath,omitempty"`

	// ShuffleServicePath: The Shuffle service path relative to the root
	// URL, for example,
	// "shuffle/v1beta1".
	ShuffleServicePath string `json:"shuffleServicePath,omitempty"`

	// TempStoragePrefix: The prefix of the resources the system should use
	// for temporary
	// storage.
	//
	// The supported resource type is:
	//
	// Google Cloud
	// Storage:
	//   storage.googleapis.com/{bucket}/{object}
	//
	// bucket.storage.googleapis.com/{object}
	TempStoragePrefix string `json:"tempStoragePrefix,omitempty"`

	// WorkerId: ID of the worker running this pipeline.
	WorkerId string `json:"workerId,omitempty"`
}

type WriteInstruction struct {
	// Input: The input.
	Input *InstructionInput `json:"input,omitempty"`

	// Sink: The sink to write to.
	Sink *Sink `json:"sink,omitempty"`
}

// method id "dataflow.projects.workerMessages":

type ProjectsWorkerMessagesCall struct {
	s                         *Service
	projectId                 string
	sendworkermessagesrequest *SendWorkerMessagesRequest
	opt_                      map[string]interface{}
}

// WorkerMessages: Send a worker_message to the service.
func (r *ProjectsService) WorkerMessages(projectId string, sendworkermessagesrequest *SendWorkerMessagesRequest) *ProjectsWorkerMessagesCall {
	c := &ProjectsWorkerMessagesCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.sendworkermessagesrequest = sendworkermessagesrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsWorkerMessagesCall) Fields(s ...googleapi.Field) *ProjectsWorkerMessagesCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsWorkerMessagesCall) Do() (*SendWorkerMessagesResponse, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.sendworkermessagesrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1b3/projects/{projectId}/WorkerMessages")
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
	var ret *SendWorkerMessagesResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Send a worker_message to the service.",
	//   "flatPath": "v1b3/projects/{projectId}/WorkerMessages",
	//   "httpMethod": "POST",
	//   "id": "dataflow.projects.workerMessages",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "The project to send the WorkerMessages to.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1b3/projects/{projectId}/WorkerMessages",
	//   "request": {
	//     "$ref": "SendWorkerMessagesRequest"
	//   },
	//   "response": {
	//     "$ref": "SendWorkerMessagesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/userinfo.email"
	//   ]
	// }

}

// method id "dataflow.projects.jobs.create":

type ProjectsJobsCreateCall struct {
	s         *Service
	projectId string
	job       *Job
	opt_      map[string]interface{}
}

// Create: Creates a dataflow job.
func (r *ProjectsJobsService) Create(projectId string, job *Job) *ProjectsJobsCreateCall {
	c := &ProjectsJobsCreateCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.job = job
	return c
}

// ReplaceJobId sets the optional parameter "replaceJobId":
// DEPRECATED.
// This field is now on the Job message.
func (c *ProjectsJobsCreateCall) ReplaceJobId(replaceJobId string) *ProjectsJobsCreateCall {
	c.opt_["replaceJobId"] = replaceJobId
	return c
}

// View sets the optional parameter "view": Level of information
// requested in response.
func (c *ProjectsJobsCreateCall) View(view string) *ProjectsJobsCreateCall {
	c.opt_["view"] = view
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsJobsCreateCall) Fields(s ...googleapi.Field) *ProjectsJobsCreateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsJobsCreateCall) Do() (*Job, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.job)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["replaceJobId"]; ok {
		params.Set("replaceJobId", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["view"]; ok {
		params.Set("view", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1b3/projects/{projectId}/jobs")
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
	//   "description": "Creates a dataflow job.",
	//   "flatPath": "v1b3/projects/{projectId}/jobs",
	//   "httpMethod": "POST",
	//   "id": "dataflow.projects.jobs.create",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "The project which owns the job.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "replaceJobId": {
	//       "description": "DEPRECATED.\nThis field is now on the Job message.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "view": {
	//       "description": "Level of information requested in response.",
	//       "enum": [
	//         "JOB_VIEW_UNKNOWN",
	//         "JOB_VIEW_SUMMARY",
	//         "JOB_VIEW_ALL"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1b3/projects/{projectId}/jobs",
	//   "request": {
	//     "$ref": "Job"
	//   },
	//   "response": {
	//     "$ref": "Job"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/userinfo.email"
	//   ]
	// }

}

// method id "dataflow.projects.jobs.get":

type ProjectsJobsGetCall struct {
	s         *Service
	projectId string
	jobId     string
	opt_      map[string]interface{}
}

// Get: Gets the state of the specified dataflow job.
func (r *ProjectsJobsService) Get(projectId string, jobId string) *ProjectsJobsGetCall {
	c := &ProjectsJobsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.jobId = jobId
	return c
}

// View sets the optional parameter "view": Level of information
// requested in response.
func (c *ProjectsJobsGetCall) View(view string) *ProjectsJobsGetCall {
	c.opt_["view"] = view
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
	if v, ok := c.opt_["view"]; ok {
		params.Set("view", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1b3/projects/{projectId}/jobs/{jobId}")
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
	//   "description": "Gets the state of the specified dataflow job.",
	//   "flatPath": "v1b3/projects/{projectId}/jobs/{jobId}",
	//   "httpMethod": "GET",
	//   "id": "dataflow.projects.jobs.get",
	//   "parameterOrder": [
	//     "projectId",
	//     "jobId"
	//   ],
	//   "parameters": {
	//     "jobId": {
	//       "description": "Identifies a single job.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "The project which owns the job.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "view": {
	//       "description": "Level of information requested in response.",
	//       "enum": [
	//         "JOB_VIEW_UNKNOWN",
	//         "JOB_VIEW_SUMMARY",
	//         "JOB_VIEW_ALL"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1b3/projects/{projectId}/jobs/{jobId}",
	//   "response": {
	//     "$ref": "Job"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/userinfo.email"
	//   ]
	// }

}

// method id "dataflow.projects.jobs.getMetrics":

type ProjectsJobsGetMetricsCall struct {
	s         *Service
	projectId string
	jobId     string
	opt_      map[string]interface{}
}

// GetMetrics: Request the job status.
func (r *ProjectsJobsService) GetMetrics(projectId string, jobId string) *ProjectsJobsGetMetricsCall {
	c := &ProjectsJobsGetMetricsCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.jobId = jobId
	return c
}

// StartTime sets the optional parameter "startTime": Return only metric
// data that has changed since this time.
// Default is to return all
// information about all metrics for the job.
func (c *ProjectsJobsGetMetricsCall) StartTime(startTime string) *ProjectsJobsGetMetricsCall {
	c.opt_["startTime"] = startTime
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsJobsGetMetricsCall) Fields(s ...googleapi.Field) *ProjectsJobsGetMetricsCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsJobsGetMetricsCall) Do() (*JobMetrics, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["startTime"]; ok {
		params.Set("startTime", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1b3/projects/{projectId}/jobs/{jobId}/metrics")
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
	var ret *JobMetrics
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Request the job status.",
	//   "flatPath": "v1b3/projects/{projectId}/jobs/{jobId}/metrics",
	//   "httpMethod": "GET",
	//   "id": "dataflow.projects.jobs.getMetrics",
	//   "parameterOrder": [
	//     "projectId",
	//     "jobId"
	//   ],
	//   "parameters": {
	//     "jobId": {
	//       "description": "The job to get messages for.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "A project id.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "startTime": {
	//       "description": "Return only metric data that has changed since this time.\nDefault is to return all information about all metrics for the job.",
	//       "format": "google-datetime",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1b3/projects/{projectId}/jobs/{jobId}/metrics",
	//   "response": {
	//     "$ref": "JobMetrics"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/userinfo.email"
	//   ]
	// }

}

// method id "dataflow.projects.jobs.list":

type ProjectsJobsListCall struct {
	s         *Service
	projectId string
	opt_      map[string]interface{}
}

// List: List the jobs of a project
func (r *ProjectsJobsService) List(projectId string) *ProjectsJobsListCall {
	c := &ProjectsJobsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	return c
}

// PageSize sets the optional parameter "pageSize": If there are many
// jobs, limit response to at most this many.
// The actual number of jobs
// returned will be the lesser of max_responses
// and an unspecified
// server-defined limit.
func (c *ProjectsJobsListCall) PageSize(pageSize int64) *ProjectsJobsListCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken": Set this to the
// 'next_page_token' field of a previous response
// to request additional
// results in a long list.
func (c *ProjectsJobsListCall) PageToken(pageToken string) *ProjectsJobsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// View sets the optional parameter "view": Level of information
// requested in response. Default is SUMMARY.
func (c *ProjectsJobsListCall) View(view string) *ProjectsJobsListCall {
	c.opt_["view"] = view
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
	if v, ok := c.opt_["pageSize"]; ok {
		params.Set("pageSize", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["view"]; ok {
		params.Set("view", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1b3/projects/{projectId}/jobs")
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
	//   "description": "List the jobs of a project",
	//   "flatPath": "v1b3/projects/{projectId}/jobs",
	//   "httpMethod": "GET",
	//   "id": "dataflow.projects.jobs.list",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "If there are many jobs, limit response to at most this many.\nThe actual number of jobs returned will be the lesser of max_responses\nand an unspecified server-defined limit.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Set this to the 'next_page_token' field of a previous response\nto request additional results in a long list.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "The project which owns the jobs.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "view": {
	//       "description": "Level of information requested in response. Default is SUMMARY.",
	//       "enum": [
	//         "JOB_VIEW_UNKNOWN",
	//         "JOB_VIEW_SUMMARY",
	//         "JOB_VIEW_ALL"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1b3/projects/{projectId}/jobs",
	//   "response": {
	//     "$ref": "ListJobsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/userinfo.email"
	//   ]
	// }

}

// method id "dataflow.projects.jobs.update":

type ProjectsJobsUpdateCall struct {
	s         *Service
	projectId string
	jobId     string
	job       *Job
	opt_      map[string]interface{}
}

// Update: Updates the state of an existing dataflow job.
func (r *ProjectsJobsService) Update(projectId string, jobId string, job *Job) *ProjectsJobsUpdateCall {
	c := &ProjectsJobsUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.jobId = jobId
	c.job = job
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsJobsUpdateCall) Fields(s ...googleapi.Field) *ProjectsJobsUpdateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsJobsUpdateCall) Do() (*Job, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.job)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1b3/projects/{projectId}/jobs/{jobId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
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
	//   "description": "Updates the state of an existing dataflow job.",
	//   "flatPath": "v1b3/projects/{projectId}/jobs/{jobId}",
	//   "httpMethod": "PUT",
	//   "id": "dataflow.projects.jobs.update",
	//   "parameterOrder": [
	//     "projectId",
	//     "jobId"
	//   ],
	//   "parameters": {
	//     "jobId": {
	//       "description": "Identifies a single job.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "The project which owns the job.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1b3/projects/{projectId}/jobs/{jobId}",
	//   "request": {
	//     "$ref": "Job"
	//   },
	//   "response": {
	//     "$ref": "Job"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/userinfo.email"
	//   ]
	// }

}

// method id "dataflow.projects.jobs.messages.list":

type ProjectsJobsMessagesListCall struct {
	s         *Service
	projectId string
	jobId     string
	opt_      map[string]interface{}
}

// List: Request the job status.
func (r *ProjectsJobsMessagesService) List(projectId string, jobId string) *ProjectsJobsMessagesListCall {
	c := &ProjectsJobsMessagesListCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.jobId = jobId
	return c
}

// EndTime sets the optional parameter "endTime": Return only messages
// with timestamps < end_time. The default is now
// (i.e. return up to the
// latest messages available).
func (c *ProjectsJobsMessagesListCall) EndTime(endTime string) *ProjectsJobsMessagesListCall {
	c.opt_["endTime"] = endTime
	return c
}

// MinimumImportance sets the optional parameter "minimumImportance":
// Filter to only get messages with importance >= level
func (c *ProjectsJobsMessagesListCall) MinimumImportance(minimumImportance string) *ProjectsJobsMessagesListCall {
	c.opt_["minimumImportance"] = minimumImportance
	return c
}

// PageSize sets the optional parameter "pageSize": If specified,
// determines the maximum number of messages to
// return.  If unspecified,
// the service may choose an appropriate
// default, or may return an
// arbitrarily large number of results.
func (c *ProjectsJobsMessagesListCall) PageSize(pageSize int64) *ProjectsJobsMessagesListCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken": If supplied, this
// should be the value of next_page_token returned
// by an earlier call.
// This will cause the next page of results to
// be returned.
func (c *ProjectsJobsMessagesListCall) PageToken(pageToken string) *ProjectsJobsMessagesListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// StartTime sets the optional parameter "startTime": If specified,
// return only messages with timestamps >= start_time.
// The default is
// the job creation time (i.e. beginning of messages).
func (c *ProjectsJobsMessagesListCall) StartTime(startTime string) *ProjectsJobsMessagesListCall {
	c.opt_["startTime"] = startTime
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsJobsMessagesListCall) Fields(s ...googleapi.Field) *ProjectsJobsMessagesListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsJobsMessagesListCall) Do() (*ListJobMessagesResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["endTime"]; ok {
		params.Set("endTime", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["minimumImportance"]; ok {
		params.Set("minimumImportance", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageSize"]; ok {
		params.Set("pageSize", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["startTime"]; ok {
		params.Set("startTime", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1b3/projects/{projectId}/jobs/{jobId}/messages")
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
	var ret *ListJobMessagesResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Request the job status.",
	//   "flatPath": "v1b3/projects/{projectId}/jobs/{jobId}/messages",
	//   "httpMethod": "GET",
	//   "id": "dataflow.projects.jobs.messages.list",
	//   "parameterOrder": [
	//     "projectId",
	//     "jobId"
	//   ],
	//   "parameters": {
	//     "endTime": {
	//       "description": "Return only messages with timestamps \u003c end_time. The default is now\n(i.e. return up to the latest messages available).",
	//       "format": "google-datetime",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "jobId": {
	//       "description": "The job to get messages about.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "minimumImportance": {
	//       "description": "Filter to only get messages with importance \u003e= level",
	//       "enum": [
	//         "JOB_MESSAGE_IMPORTANCE_UNKNOWN",
	//         "JOB_MESSAGE_DEBUG",
	//         "JOB_MESSAGE_DETAILED",
	//         "JOB_MESSAGE_BASIC",
	//         "JOB_MESSAGE_WARNING",
	//         "JOB_MESSAGE_ERROR"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "If specified, determines the maximum number of messages to\nreturn.  If unspecified, the service may choose an appropriate\ndefault, or may return an arbitrarily large number of results.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "If supplied, this should be the value of next_page_token returned\nby an earlier call. This will cause the next page of results to\nbe returned.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "A project id.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "startTime": {
	//       "description": "If specified, return only messages with timestamps \u003e= start_time.\nThe default is the job creation time (i.e. beginning of messages).",
	//       "format": "google-datetime",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1b3/projects/{projectId}/jobs/{jobId}/messages",
	//   "response": {
	//     "$ref": "ListJobMessagesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/userinfo.email"
	//   ]
	// }

}

// method id "dataflow.projects.jobs.workItems.lease":

type ProjectsJobsWorkItemsLeaseCall struct {
	s                    *Service
	projectId            string
	jobId                string
	leaseworkitemrequest *LeaseWorkItemRequest
	opt_                 map[string]interface{}
}

// Lease: Leases a dataflow WorkItem to run.
func (r *ProjectsJobsWorkItemsService) Lease(projectId string, jobId string, leaseworkitemrequest *LeaseWorkItemRequest) *ProjectsJobsWorkItemsLeaseCall {
	c := &ProjectsJobsWorkItemsLeaseCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.jobId = jobId
	c.leaseworkitemrequest = leaseworkitemrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsJobsWorkItemsLeaseCall) Fields(s ...googleapi.Field) *ProjectsJobsWorkItemsLeaseCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsJobsWorkItemsLeaseCall) Do() (*LeaseWorkItemResponse, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.leaseworkitemrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1b3/projects/{projectId}/jobs/{jobId}/workItems:lease")
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
	var ret *LeaseWorkItemResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Leases a dataflow WorkItem to run.",
	//   "flatPath": "v1b3/projects/{projectId}/jobs/{jobId}/workItems:lease",
	//   "httpMethod": "POST",
	//   "id": "dataflow.projects.jobs.workItems.lease",
	//   "parameterOrder": [
	//     "projectId",
	//     "jobId"
	//   ],
	//   "parameters": {
	//     "jobId": {
	//       "description": "Identifies the workflow job this worker belongs to.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "Identifies the project this worker belongs to.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1b3/projects/{projectId}/jobs/{jobId}/workItems:lease",
	//   "request": {
	//     "$ref": "LeaseWorkItemRequest"
	//   },
	//   "response": {
	//     "$ref": "LeaseWorkItemResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/userinfo.email"
	//   ]
	// }

}

// method id "dataflow.projects.jobs.workItems.reportStatus":

type ProjectsJobsWorkItemsReportStatusCall struct {
	s                           *Service
	projectId                   string
	jobId                       string
	reportworkitemstatusrequest *ReportWorkItemStatusRequest
	opt_                        map[string]interface{}
}

// ReportStatus: Reports the status of dataflow WorkItems leased by a
// worker.
func (r *ProjectsJobsWorkItemsService) ReportStatus(projectId string, jobId string, reportworkitemstatusrequest *ReportWorkItemStatusRequest) *ProjectsJobsWorkItemsReportStatusCall {
	c := &ProjectsJobsWorkItemsReportStatusCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.jobId = jobId
	c.reportworkitemstatusrequest = reportworkitemstatusrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsJobsWorkItemsReportStatusCall) Fields(s ...googleapi.Field) *ProjectsJobsWorkItemsReportStatusCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsJobsWorkItemsReportStatusCall) Do() (*ReportWorkItemStatusResponse, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.reportworkitemstatusrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1b3/projects/{projectId}/jobs/{jobId}/workItems:reportStatus")
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
	var ret *ReportWorkItemStatusResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Reports the status of dataflow WorkItems leased by a worker.",
	//   "flatPath": "v1b3/projects/{projectId}/jobs/{jobId}/workItems:reportStatus",
	//   "httpMethod": "POST",
	//   "id": "dataflow.projects.jobs.workItems.reportStatus",
	//   "parameterOrder": [
	//     "projectId",
	//     "jobId"
	//   ],
	//   "parameters": {
	//     "jobId": {
	//       "description": "The job which the WorkItem is part of.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "The project which owns the WorkItem's job.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1b3/projects/{projectId}/jobs/{jobId}/workItems:reportStatus",
	//   "request": {
	//     "$ref": "ReportWorkItemStatusRequest"
	//   },
	//   "response": {
	//     "$ref": "ReportWorkItemStatusResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/userinfo.email"
	//   ]
	// }

}
