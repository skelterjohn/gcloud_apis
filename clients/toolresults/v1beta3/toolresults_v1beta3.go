// Package toolresults provides access to the Cloud Tool Results API.
//
// See http://code.google.com/apis/cloud/toolresults/v1/using_rest.html
//
// Usage example:
//
//   import "google.golang.org/api/toolresults/v1beta3"
//   ...
//   toolresultsService, err := toolresults.New(oauthHttpClient)
package toolresults

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

const apiId = "toolresults:v1beta3"
const apiName = "toolresults"
const apiVersion = "v1beta3"
const basePath = "https://www.googleapis.com/toolresults/v1beta3/projects/"

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
	return s, nil
}

type Service struct {
	client   *http.Client
	BasePath string // API endpoint base URL

	Projects *ProjectsService
}

func NewProjectsService(s *Service) *ProjectsService {
	rs := &ProjectsService{s: s}
	rs.Histories = NewProjectsHistoriesService(s)
	return rs
}

type ProjectsService struct {
	s *Service

	Histories *ProjectsHistoriesService
}

func NewProjectsHistoriesService(s *Service) *ProjectsHistoriesService {
	rs := &ProjectsHistoriesService{s: s}
	rs.Executions = NewProjectsHistoriesExecutionsService(s)
	return rs
}

type ProjectsHistoriesService struct {
	s *Service

	Executions *ProjectsHistoriesExecutionsService
}

func NewProjectsHistoriesExecutionsService(s *Service) *ProjectsHistoriesExecutionsService {
	rs := &ProjectsHistoriesExecutionsService{s: s}
	rs.Steps = NewProjectsHistoriesExecutionsStepsService(s)
	return rs
}

type ProjectsHistoriesExecutionsService struct {
	s *Service

	Steps *ProjectsHistoriesExecutionsStepsService
}

func NewProjectsHistoriesExecutionsStepsService(s *Service) *ProjectsHistoriesExecutionsStepsService {
	rs := &ProjectsHistoriesExecutionsStepsService{s: s}
	rs.Thumbnails = NewProjectsHistoriesExecutionsStepsThumbnailsService(s)
	return rs
}

type ProjectsHistoriesExecutionsStepsService struct {
	s *Service

	Thumbnails *ProjectsHistoriesExecutionsStepsThumbnailsService
}

func NewProjectsHistoriesExecutionsStepsThumbnailsService(s *Service) *ProjectsHistoriesExecutionsStepsThumbnailsService {
	rs := &ProjectsHistoriesExecutionsStepsThumbnailsService{s: s}
	return rs
}

type ProjectsHistoriesExecutionsStepsThumbnailsService struct {
	s *Service
}

type Any struct {
	// TypeUrl: A URL/resource name whose content describes the type of the
	// serialized message.
	//
	// For URLs which use the schema `http`, `https`,
	// or no schema, the following restrictions and interpretations
	// apply:
	//
	// * If no schema is provided, `https` is assumed. * The last
	// segment of the URL's path must represent the fully qualified name of
	// the type (as in `path/google.protobuf.Duration`). * An HTTP GET on
	// the URL must yield a [google.protobuf.Type][] value in binary format,
	// or produce an error. * Applications are allowed to cache lookup
	// results based on the URL, or have them precompiled into a binary to
	// avoid any lookup. Therefore, binary compatibility needs to be
	// preserved on changes to types. (Use versioned type names to manage
	// breaking changes.)
	//
	// Schemas other than `http`, `https` (or the empty
	// schema) might be used with implementation specific semantics.
	TypeUrl string `json:"typeUrl,omitempty"`

	// Value: Must be valid serialized data of the above specified type.
	Value string `json:"value,omitempty"`
}

type Duration struct {
	// Nanos: Signed fractions of a second at nanosecond resolution of the
	// span of time. Durations less than one second are represented with a 0
	// `seconds` field and a positive or negative `nanos` field. For
	// durations of one second or more, a non-zero value for the `nanos`
	// field must be of the same sign as the `seconds` field. Must be from
	// -999,999,999 to +999,999,999 inclusive.
	Nanos int64 `json:"nanos,omitempty"`

	// Seconds: Signed seconds of the span of time. Must be from
	// -315,576,000,000 to +315,576,000,000 inclusive.
	Seconds int64 `json:"seconds,omitempty,string"`
}

type Execution struct {
	// CompletionTime: The time when the Execution status transitioned to
	// COMPLETE.
	//
	// This value will be set automatically when state
	// transitions to COMPLETE.
	//
	// - In response: set if the execution state
	// is COMPLETE. - In create/update request: never set
	CompletionTime *Timestamp `json:"completionTime,omitempty"`

	// CreationTime: The time when the Execution was created.
	//
	// This value
	// will be set automatically when CreateExecution is called.
	//
	// - In
	// response: always set - In create/update request: never set
	CreationTime *Timestamp `json:"creationTime,omitempty"`

	// ExecutionId: A unique identifier within a History for this
	// Execution.
	//
	// Returns INVALID_ARGUMENT if this field is set or
	// overwritten by the caller.
	//
	// - In response always set - In
	// create/update request: never set
	ExecutionId string `json:"executionId,omitempty"`

	// Outcome: Classify the result, for example into SUCCESS or FAILURE
	//
	// -
	// In response: present if set by create/update request - In
	// create/update request: optional
	Outcome *Outcome `json:"outcome,omitempty"`

	// State: The initial state is IN_PROGRESS.
	//
	// The only legal state
	// transitions is from IN_PROGRESS to COMPLETE.
	//
	// A PRECONDITION_FAILED
	// will be returned if an invalid transition is requested.
	//
	// The state
	// can only be set to COMPLETE once. A FAILED_PRECONDITION will be
	// returned if the state is set to COMPLETE multiple times.
	//
	// If the
	// state is set to COMPLETE, all the in-progress steps within the
	// execution will be set as COMPLETE. If the outcome of the step is not
	// set, the outcome will be set to INCONCLUSIVE.
	//
	// - In response always
	// set - In create/update request: optional
	State string `json:"state,omitempty"`

	// TestExecutionMatrixId: TestExecution Matrix ID that the Test Service
	// uses.
	//
	// - In response: present if set by create - In create: optional
	// - In update: never set
	TestExecutionMatrixId string `json:"testExecutionMatrixId,omitempty"`
}

type FailureDetail struct {
	// Crashed: If the failure was severe because the system under test
	// crashed.
	Crashed bool `json:"crashed,omitempty"`

	// NotInstalled: If an app is not installed and thus no test can be run
	// with the app. This might be caused by trying to run a test on an
	// unsupported platform.
	NotInstalled bool `json:"notInstalled,omitempty"`

	// OtherNativeCrash: If a native process other than the app crashed.
	OtherNativeCrash bool `json:"otherNativeCrash,omitempty"`

	// TimedOut: If the test overran some time limit, and that is why it
	// failed.
	TimedOut bool `json:"timedOut,omitempty"`
}

type FileReference struct {
	// FileUri: The URI of a file stored in Google Cloud Storage.
	//
	// For
	// example: http://storage.googleapis.com/mybucket/path/to/test.xml or
	// in gsutil format: gs://mybucket/path/to/test.xml with
	// version-specific info,
	// gs://mybucket/path/to/test.xml#1360383693690000
	//
	// An INVALID_ARGUMENT
	// error will be returned if the URI format is not supported.
	//
	// - In
	// response: always set - In create/update request: always set
	FileUri string `json:"fileUri,omitempty"`
}

type History struct {
	// DisplayName: A short human-readable (plain text) name to display in
	// the UI. Maximum of 100 characters.
	//
	// - In response: present if set
	// during create. - In create request: optional
	DisplayName string `json:"displayName,omitempty"`

	// HistoryId: A unique identifier within a project for this
	// History.
	//
	// Returns INVALID_ARGUMENT if this field is set or
	// overwritten by the caller.
	//
	// - In response always set - In create
	// request: never set
	HistoryId string `json:"historyId,omitempty"`

	// Name: A name to uniquely identify a history within a project. Maximum
	// of 100 characters.
	//
	// - In response always set - In create request:
	// always set
	Name string `json:"name,omitempty"`
}

type Image struct {
	// Error: An error explaining why the thumbnail could not be rendered.
	Error *Status `json:"error,omitempty"`

	// SourceImage: A reference to the full-size, original image.
	//
	// This is
	// the same as the tool_outputs entry for the image under its
	// Step.
	//
	// Always set.
	SourceImage *ToolOutputReference `json:"sourceImage,omitempty"`

	// StepId: The step to which the image is attached.
	//
	// Always set.
	StepId string `json:"stepId,omitempty"`

	// Thumbnail: The thumbnail.
	Thumbnail *Thumbnail `json:"thumbnail,omitempty"`
}

type InconclusiveDetail struct {
	// AbortedByUser: If the end user aborted the test execution before a
	// pass or fail could be determined. For example, the user pressed
	// ctrl-c which sent a kill signal to the test runner while the test was
	// running.
	AbortedByUser bool `json:"abortedByUser,omitempty"`

	// InfrastructureFailure: If the test runner could not determine success
	// or failure because the test depends on a component other than the
	// system under test which failed.
	//
	// For example, a mobile test requires
	// provisioning a device where the test executes, and that provisioning
	// can fail.
	InfrastructureFailure bool `json:"infrastructureFailure,omitempty"`

	// NativeCrash: A native process crashed on the device, producing a
	// tombstone. It is unclear whether the crash was related to the app
	// under test.
	//
	// For example, OpenGL crashed, but it is unclear if the
	// app is responsible. TODO(yinfu): Remove after all reference from
	// TestService are deleted.
	NativeCrash bool `json:"nativeCrash,omitempty"`
}

type ListExecutionsResponse struct {
	// Executions: Executions.
	//
	// Always set.
	Executions []*Execution `json:"executions,omitempty"`

	// NextPageToken: A continuation token to resume the query at the next
	// item.
	//
	// Will only be set if there are more Executions to fetch.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type ListHistoriesResponse struct {
	// Histories: Histories.
	Histories []*History `json:"histories,omitempty"`

	// NextPageToken: A continuation token to resume the query at the next
	// item.
	//
	// Will only be set if there are more histories to fetch.
	//
	// Tokens
	// are valid for up to one hour from the time of the first list request.
	// For instance, if you make a list request at 1PM and use the token
	// from this first request 10 minutes later, the token from this second
	// response will only be valid for 50 minutes.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type ListStepThumbnailsResponse struct {
	// NextPageToken: A continuation token to resume the query at the next
	// item.
	//
	// If set, indicates that there are more thumbnails to read, by
	// calling list again with this value in the page_token field.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Thumbnails: A list of image data.
	//
	// Images are returned in a
	// deterministic order; they are ordered by these factors, in order of
	// importance: * First, by their associated test case. Images without a
	// test case are considered greater than images with one. * Second, by
	// their creation time. Images without a creation time are greater than
	// images with one. * Third, by the order in which they were added to
	// the step (by calls to CreateStep or UpdateStep).
	Thumbnails []*Image `json:"thumbnails,omitempty"`
}

type ListStepsResponse struct {
	// NextPageToken: A continuation token to resume the query at the next
	// item.
	//
	// If set, indicates that there are more steps to read, by
	// calling list again with this value in the page_token field.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Steps: Steps.
	Steps []*Step `json:"steps,omitempty"`
}

type Outcome struct {
	// FailureDetail: More information about a FAILURE outcome.
	//
	// Returns
	// INVALID_ARGUMENT if this field is set but the summary is not
	// FAILURE.
	//
	// Optional
	FailureDetail *FailureDetail `json:"failureDetail,omitempty"`

	// InconclusiveDetail: More information about an INCONCLUSIVE
	// outcome.
	//
	// Returns INVALID_ARGUMENT if this field is set but the
	// summary is not INCONCLUSIVE.
	//
	// Optional
	InconclusiveDetail *InconclusiveDetail `json:"inconclusiveDetail,omitempty"`

	// SkippedDetail: More information about a SKIPPED outcome.
	//
	// Returns
	// INVALID_ARGUMENT if this field is set but the summary is not
	// SKIPPED.
	//
	// Optional
	SkippedDetail *SkippedDetail `json:"skippedDetail,omitempty"`

	// SuccessDetail: More information about a SUCCESS outcome.
	//
	// Returns
	// INVALID_ARGUMENT if this field is set but the summary is not
	// SUCCESS.
	//
	// Optional
	SuccessDetail *SuccessDetail `json:"successDetail,omitempty"`

	// Summary: The simplest way to interpret a result.
	//
	// Required
	Summary string `json:"summary,omitempty"`
}

type ProjectSettings struct {
	// DefaultBucket: The name of the Google Cloud Storage bucket to which
	// results are written.
	//
	// By default, this is unset.
	//
	// In update request:
	// optional In response: optional
	DefaultBucket string `json:"defaultBucket,omitempty"`

	// Name: The name of the project's settings.
	//
	// Always of the form:
	// projects/{project-id}/settings
	//
	// In update request: never set In
	// response: always set
	Name string `json:"name,omitempty"`
}

type PublishXunitXmlFilesRequest struct {
	// XunitXmlFiles: URI of the Xunit XML files to publish.
	//
	// The maximum
	// size of the file this reference is pointing to is 50MB.
	//
	// Required.
	XunitXmlFiles []*FileReference `json:"xunitXmlFiles,omitempty"`
}

type SkippedDetail struct {
	// IncompatibleAppVersion: If the App doesn't support the specific API
	// level.
	IncompatibleAppVersion bool `json:"incompatibleAppVersion,omitempty"`

	// IncompatibleArchitecture: If the App doesn't run on the specific
	// architecture, for example, x86.
	IncompatibleArchitecture bool `json:"incompatibleArchitecture,omitempty"`

	// IncompatibleDevice: If the requested OS version doesn't run on the
	// specific device model.
	IncompatibleDevice bool `json:"incompatibleDevice,omitempty"`
}

type Status struct {
	// Code: The status code, which should be an enum value of
	// [google.rpc.Code][].
	Code int64 `json:"code,omitempty"`

	// Details: A list of messages that carry the error details. There will
	// be a common set of message types for APIs to use.
	Details []*Any `json:"details,omitempty"`

	// Message: A developer-facing error message, which should be in
	// English. Any user-facing error message should be localized and sent
	// in the [google.rpc.Status.details][] field, or localized by the
	// client.
	Message string `json:"message,omitempty"`
}

type Step struct {
	// CompletionTime: The time when the step status was set to
	// complete.
	//
	// This value will be set automatically when state
	// transitions to COMPLETE.
	//
	// - In response: set if the execution state
	// is COMPLETE. - In create/update request: never set
	CompletionTime *Timestamp `json:"completionTime,omitempty"`

	// CreationTime: The time when the step was created.
	//
	// - In response:
	// always set - In create/update request: never set
	CreationTime *Timestamp `json:"creationTime,omitempty"`

	// Description: A description of this tool For example: mvn clean
	// package -D skipTests=true
	//
	// - In response: present if set by
	// create/update request - In create/update request: optional
	Description string `json:"description,omitempty"`

	// DimensionValue: If the execution containing this step has any
	// dimension_definition set, then this field allows the child to specify
	// the values of the dimensions.
	//
	// The keys must exactly match the
	// dimension_definition of the execution.
	//
	// For example, if the execution
	// has `dimension_definition = ['attempt', 'device']` then a step must
	// define values for those dimensions, eg. `dimension_value =
	// ['attempt': '1', 'device': 'Nexus 6']`
	//
	// If a step does not
	// participate in one dimension of the matrix, the value for that
	// dimension should be empty string. For example, if one of the tests is
	// executed by a runner which does not support retries, the step could
	// have `dimension_value = ['attempt': '', 'device': 'Nexus 6']`
	//
	// If the
	// step does not participate in any dimensions of the matrix, it may
	// leave dimension_value unset.
	//
	// A PRECONDITION_FAILED will be returned
	// if any of the keys do not exist in the dimension_definition of the
	// execution.
	//
	// A PRECONDITION_FAILED will be returned if another step in
	// this execution already has the same name and dimension_value, but
	// differs on other data fields, for example, step field is
	// different.
	//
	// A PRECONDITION_FAILED will be returned if dimension_value
	// is set, and there is a dimension_definition in the execution which is
	// not specified as one of the keys.
	//
	// - In response: present if set by
	// create - In create request: optional - In update request: never set
	DimensionValue []*StepDimensionValueEntry `json:"dimensionValue,omitempty"`

	// HasImages: Whether any of the outputs of this step are images whose
	// thumbnails can be fetched with ListThumbnails.
	//
	// - In response: always
	// set - In create/update request: never set
	HasImages bool `json:"hasImages,omitempty"`

	// Labels: Arbitrary user-supplied key/value pairs that are associated
	// with the step.
	//
	// Users are responsible for managing the key namespace
	// such that keys don't accidentally collide.
	//
	// An INVALID_ARGUMENT will
	// be returned if the number of labels exceeds 100 or if the length of
	// any of the keys or values exceeds 100 characters.
	//
	// - In response:
	// always set - In create request: optional - In update request:
	// optional; any new key/value pair will be added to the map, and any
	// new value for an existing key will update that key's value
	Labels []*StepLabelsEntry `json:"labels,omitempty"`

	// Name: A short human-readable name to display in the UI. Maximum of
	// 100 characters. For example: Clean build
	//
	// A PRECONDITION_FAILED will
	// be returned upon creating a new step if it shares its name and
	// dimension_value with an existing step. If two steps represent a
	// similar action, but have different dimension values, they should
	// share the same name. For instance, if the same set of tests is run on
	// two different platforms, the two steps should have the same name.
	//
	// -
	// In response: always set - In create request: always set - In update
	// request: never set
	Name string `json:"name,omitempty"`

	// Outcome: Classification of the result, for example into SUCCESS or
	// FAILURE
	//
	// - In response: present if set by create/update request - In
	// create/update request: optional
	Outcome *Outcome `json:"outcome,omitempty"`

	// RunDuration: How long it took for this step to run.
	//
	// If unset, this
	// is set to the difference between creation_time and completion_time
	// when the step is set to the COMPLETE state. In some cases, it is
	// appropriate to set this value separately: For instance, if a step is
	// created, but the operation it represents is queued for a few minutes
	// before it executes, it would be appropriate not to include the time
	// spent queued in its run_duration.
	//
	// PRECONDITION_FAILED will be
	// returned if one attempts to set a run_duration on a step which
	// already has this field set.
	//
	// - In response: present if previously
	// set; always present on COMPLETE step - In create request: optional -
	// In update request: optional
	RunDuration *Duration `json:"runDuration,omitempty"`

	// State: The initial state is IN_PROGRESS. The only legal state
	// transitions are * IN_PROGRESS -> COMPLETE
	//
	// A PRECONDITION_FAILED will
	// be returned if an invalid transition is requested.
	//
	// It is valid to
	// create Step with a state set to COMPLETE. The state can only be set
	// to COMPLETE once. A PRECONDITION_FAILED will be returned if the state
	// is set to COMPLETE multiple times.
	//
	// - In response: always set - In
	// create/update request: optional
	State string `json:"state,omitempty"`

	// StepId: A unique identifier within a Execution for this
	// Step.
	//
	// Returns INVALID_ARGUMENT if this field is set or overwritten
	// by the caller.
	//
	// - In response: always set - In create/update request:
	// never set
	StepId string `json:"stepId,omitempty"`

	// TestExecutionStep: An execution of a test runner.
	TestExecutionStep *TestExecutionStep `json:"testExecutionStep,omitempty"`

	// ToolExecutionStep: An execution of a tool (used for steps we don't
	// explicitly support).
	ToolExecutionStep *ToolExecutionStep `json:"toolExecutionStep,omitempty"`
}

type StepDimensionValueEntry struct {
	Key string `json:"key,omitempty"`

	Value string `json:"value,omitempty"`
}

type StepLabelsEntry struct {
	Key string `json:"key,omitempty"`

	Value string `json:"value,omitempty"`
}

type SuccessDetail struct {
	// OtherNativeCrash: If a native process other than the app crashed.
	OtherNativeCrash bool `json:"otherNativeCrash,omitempty"`
}

type TestCaseReference struct {
	// ClassName: The name of the class.
	ClassName string `json:"className,omitempty"`

	// Name: The name of the test case.
	//
	// Required.
	Name string `json:"name,omitempty"`

	// TestSuiteName: The name of the test suite to which this test case
	// belongs.
	TestSuiteName string `json:"testSuiteName,omitempty"`
}

type TestExecutionStep struct {
	// TestSuiteOverviews: List of test suite overview contents. This could
	// be parsed from xUnit XML log by server, or uploaded directly by user.
	// This references should only be called when test suites are fully
	// parsed or uploaded.
	//
	// The maximum allowed number of test suite
	// overviews per step is 1000.
	//
	// - In response: always set - In create
	// request: optional - In update request: never (use
	// publishXunitXmlFiles custom method instead)
	TestSuiteOverviews []*TestSuiteOverview `json:"testSuiteOverviews,omitempty"`

	// TestTiming: The timing break down of the test execution.
	//
	// - In
	// response: present if set by create or update - In create/update
	// request: optional
	TestTiming *TestTiming `json:"testTiming,omitempty"`

	// ToolExecution: Represents the execution of the test runner.
	//
	// The exit
	// code of this tool will be used to determine if the test passed.
	//
	// - In
	// response: always set - In create/update request: optional
	ToolExecution *ToolExecution `json:"toolExecution,omitempty"`
}

type TestSuiteOverview struct {
	// ErrorCount: Number of test cases in error, typically set by the
	// service by parsing the xml_source.
	//
	// - In create/response: always set
	// - In update request: never
	ErrorCount int64 `json:"errorCount,omitempty"`

	// FailureCount: Number of failed test cases, typically set by the
	// service by parsing the xml_source. May also be set by the user.
	//
	// - In
	// create/response: always set - In update request: never
	FailureCount int64 `json:"failureCount,omitempty"`

	// Name: The name of the test suite.
	//
	// - In create/response: always set -
	// In update request: never
	Name string `json:"name,omitempty"`

	// SkippedCount: Number of test cases not run, typically set by the
	// service by parsing the xml_source.
	//
	// - In create/response: always set
	// - In update request: never
	SkippedCount int64 `json:"skippedCount,omitempty"`

	// TotalCount: Number of test cases, typically set by the service by
	// parsing the xml_source.
	//
	// - In create/response: always set - In update
	// request: never
	TotalCount int64 `json:"totalCount,omitempty"`

	// XmlSource: If this test suite was parsed from XML, this is the URI
	// where the original XML file is stored.
	//
	// Note: Multiple test suites
	// can share the same xml_source
	//
	// Returns INVALID_ARGUMENT if the uri
	// format is not supported.
	//
	// - In create/response: optional - In update
	// request: never
	XmlSource *FileReference `json:"xmlSource,omitempty"`
}

type TestTiming struct {
	// TestProcessDuration: How long it took to run the test process.
	//
	// - In
	// response: present if previously set. - In create/update request:
	// optional
	TestProcessDuration *Duration `json:"testProcessDuration,omitempty"`
}

type Thumbnail struct {
	// ContentType: The thumbnail's content type, i.e. "image/png".
	//
	// Always
	// set.
	ContentType string `json:"contentType,omitempty"`

	// Data: The thumbnail file itself.
	//
	// That is, the bytes here are
	// precisely the bytes that make up the thumbnail file; they can be
	// served as an image as-is (with the appropriate content type.)
	//
	// Always
	// set.
	Data string `json:"data,omitempty"`

	// HeightPx: The height of the thumbnail, in pixels.
	//
	// Always set.
	HeightPx int64 `json:"heightPx,omitempty"`

	// WidthPx: The width of the thumbnail, in pixels.
	//
	// Always set.
	WidthPx int64 `json:"widthPx,omitempty"`
}

type Timestamp struct {
	// Nanos: Non-negative fractions of a second at nanosecond resolution.
	// Negative second values with fractions must still have non-negative
	// nanos values that count forward in time. Must be from 0 to
	// 999,999,999 inclusive.
	Nanos int64 `json:"nanos,omitempty"`

	// Seconds: Represents seconds of UTC time since Unix epoch
	// 1970-01-01T00:00:00Z. Must be from from 0001-01-01T00:00:00Z to
	// 9999-12-31T23:59:59Z inclusive.
	Seconds int64 `json:"seconds,omitempty,string"`
}

type ToolExecution struct {
	// CommandLineArguments: The full tokenized command line including the
	// program name (equivalent to argv in a C program).
	//
	// - In response:
	// present if set by create request - In create request: optional - In
	// update request: never set
	CommandLineArguments []string `json:"commandLineArguments,omitempty"`

	// ExitCode: Tool execution exit code. This field will be set once the
	// tool has exited.
	//
	// - In response: present if set by create/update
	// request - In create request: optional - In update request: optional,
	// a FAILED_PRECONDITION error will be returned if an exit_code is
	// already set.
	ExitCode *ToolExitCode `json:"exitCode,omitempty"`

	// ToolLogs: References to any plain text logs output the tool
	// execution.
	//
	// This field can be set before the tool has exited in order
	// to be able to have access to a live view of the logs while the tool
	// is running.
	//
	// The maximum allowed number of tool logs per step is
	// 1000.
	//
	// - In response: present if set by create/update request - In
	// create request: optional - In update request: optional, any value
	// provided will be appended to the existing list
	ToolLogs []*FileReference `json:"toolLogs,omitempty"`

	// ToolOutputs: References to opaque files of any format output by the
	// tool execution.
	//
	// The maximum allowed number of tool outputs per step
	// is 1000.
	//
	// - In response: present if set by create/update request - In
	// create request: optional - In update request: optional, any value
	// provided will be appended to the existing list
	ToolOutputs []*ToolOutputReference `json:"toolOutputs,omitempty"`
}

type ToolExecutionStep struct {
	// ToolExecution: A Tool execution.
	//
	// - In response: present if set by
	// create/update request - In create/update request: optional
	ToolExecution *ToolExecution `json:"toolExecution,omitempty"`
}

type ToolExitCode struct {
	// Number: Tool execution exit code. A value of 0 means that the
	// execution was successful.
	//
	// - In response: always set - In
	// create/update request: always set
	Number int64 `json:"number,omitempty"`
}

type ToolOutputReference struct {
	// CreationTime: The creation time of the file.
	//
	// - In response: present
	// if set by create/update request - In create/update request: optional
	CreationTime *Timestamp `json:"creationTime,omitempty"`

	// Output: A FileReference to an output file.
	//
	// - In response: always set
	// - In create/update request: always set
	Output *FileReference `json:"output,omitempty"`

	// TestCase: The test case to which this output file belongs.
	//
	// - In
	// response: present if set by create/update request - In create/update
	// request: optional
	TestCase *TestCaseReference `json:"testCase,omitempty"`
}

// method id "toolresults.projects.getSettings":

type ProjectsGetSettingsCall struct {
	s         *Service
	projectId string
	opt_      map[string]interface{}
}

// GetSettings: Gets the Tool Results settings for a project.
//
// May
// return any of the following canonical error codes:
//
// -
// PERMISSION_DENIED - if the user is not authorized to read from
// project
func (r *ProjectsService) GetSettings(projectId string) *ProjectsGetSettingsCall {
	c := &ProjectsGetSettingsCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsGetSettingsCall) Fields(s ...googleapi.Field) *ProjectsGetSettingsCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsGetSettingsCall) Do() (*ProjectSettings, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{projectId}/settings")
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
	var ret *ProjectSettings
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets the Tool Results settings for a project.\n\nMay return any of the following canonical error codes:\n\n- PERMISSION_DENIED - if the user is not authorized to read from project",
	//   "httpMethod": "GET",
	//   "id": "toolresults.projects.getSettings",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "A Project id.\n\nRequired.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{projectId}/settings",
	//   "response": {
	//     "$ref": "ProjectSettings"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "toolresults.projects.initializeSettings":

type ProjectsInitializeSettingsCall struct {
	s         *Service
	projectId string
	opt_      map[string]interface{}
}

// InitializeSettings: Creates resources for settings which have not yet
// been set.
//
// Currently, this creates a single resource: a Google Cloud
// Storage bucket, to be used as the default bucket for this project.
// The bucket is created in the name of the user calling. Except in rare
// cases, calling this method in parallel from multiple clients will
// only create a single bucket. In order to avoid unnecessary storage
// charges, the bucket is configured to automatically delete objects
// older than 90 days.
//
// The bucket is created with the project-private
// ACL: All project team members are given permissions to the bucket and
// objects created within it according to their roles. Project owners
// have owners rights, and so on. The default ACL on objects created in
// the bucket is project-private as well. See Google Cloud Storage
// documentation for more details.
//
// If there is already a default bucket
// set, this call does nothing.
//
// May return any canonical error codes,
// including the following:
//
// - PERMISSION_DENIED - if the user is not
// authorized to write to project - Any error code raised by Google
// Cloud Storage
func (r *ProjectsService) InitializeSettings(projectId string) *ProjectsInitializeSettingsCall {
	c := &ProjectsInitializeSettingsCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsInitializeSettingsCall) Fields(s ...googleapi.Field) *ProjectsInitializeSettingsCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsInitializeSettingsCall) Do() (*ProjectSettings, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{projectId}:initializeSettings")
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
	var ret *ProjectSettings
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates resources for settings which have not yet been set.\n\nCurrently, this creates a single resource: a Google Cloud Storage bucket, to be used as the default bucket for this project. The bucket is created in the name of the user calling. Except in rare cases, calling this method in parallel from multiple clients will only create a single bucket. In order to avoid unnecessary storage charges, the bucket is configured to automatically delete objects older than 90 days.\n\nThe bucket is created with the project-private ACL: All project team members are given permissions to the bucket and objects created within it according to their roles. Project owners have owners rights, and so on. The default ACL on objects created in the bucket is project-private as well. See Google Cloud Storage documentation for more details.\n\nIf there is already a default bucket set, this call does nothing.\n\nMay return any canonical error codes, including the following:\n\n- PERMISSION_DENIED - if the user is not authorized to write to project - Any error code raised by Google Cloud Storage",
	//   "httpMethod": "POST",
	//   "id": "toolresults.projects.initializeSettings",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "A Project id.\n\nRequired.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{projectId}:initializeSettings",
	//   "response": {
	//     "$ref": "ProjectSettings"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "toolresults.projects.histories.create":

type ProjectsHistoriesCreateCall struct {
	s         *Service
	projectId string
	history   *History
	opt_      map[string]interface{}
}

// Create: Creates a History.
//
// The returned History will have the id
// set.
//
// May return any of the following canonical error codes:
//
// -
// PERMISSION_DENIED - if the user is not authorized to write to project
// - INVALID_ARGUMENT - if the request is malformed - NOT_FOUND - if the
// containing project does not exist
func (r *ProjectsHistoriesService) Create(projectId string, history *History) *ProjectsHistoriesCreateCall {
	c := &ProjectsHistoriesCreateCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.history = history
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsHistoriesCreateCall) Fields(s ...googleapi.Field) *ProjectsHistoriesCreateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsHistoriesCreateCall) Do() (*History, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.history)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{projectId}/histories")
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
	var ret *History
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a History.\n\nThe returned History will have the id set.\n\nMay return any of the following canonical error codes:\n\n- PERMISSION_DENIED - if the user is not authorized to write to project - INVALID_ARGUMENT - if the request is malformed - NOT_FOUND - if the containing project does not exist",
	//   "httpMethod": "POST",
	//   "id": "toolresults.projects.histories.create",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "A Project id.\n\nRequired.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{projectId}/histories",
	//   "request": {
	//     "$ref": "History"
	//   },
	//   "response": {
	//     "$ref": "History"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "toolresults.projects.histories.get":

type ProjectsHistoriesGetCall struct {
	s         *Service
	projectId string
	historyId string
	opt_      map[string]interface{}
}

// Get: Gets a History.
//
// May return any of the following canonical error
// codes:
//
// - PERMISSION_DENIED - if the user is not authorized to read
// project - INVALID_ARGUMENT - if the request is malformed - NOT_FOUND
// - if the History does not exist
func (r *ProjectsHistoriesService) Get(projectId string, historyId string) *ProjectsHistoriesGetCall {
	c := &ProjectsHistoriesGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.historyId = historyId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsHistoriesGetCall) Fields(s ...googleapi.Field) *ProjectsHistoriesGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsHistoriesGetCall) Do() (*History, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{projectId}/histories/{historyId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"historyId": c.historyId,
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
	var ret *History
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets a History.\n\nMay return any of the following canonical error codes:\n\n- PERMISSION_DENIED - if the user is not authorized to read project - INVALID_ARGUMENT - if the request is malformed - NOT_FOUND - if the History does not exist",
	//   "httpMethod": "GET",
	//   "id": "toolresults.projects.histories.get",
	//   "parameterOrder": [
	//     "projectId",
	//     "historyId"
	//   ],
	//   "parameters": {
	//     "historyId": {
	//       "description": "A History id.\n\nRequired.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "A Project id.\n\nRequired.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{projectId}/histories/{historyId}",
	//   "response": {
	//     "$ref": "History"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "toolresults.projects.histories.list":

type ProjectsHistoriesListCall struct {
	s         *Service
	projectId string
	opt_      map[string]interface{}
}

// List: Lists Histories for a given Project.
//
// The histories are sorted
// by modification time in descending order. The history_id key will be
// used to order the history with the same modification time.
//
// May
// return any of the following canonical error codes:
//
// -
// PERMISSION_DENIED - if the user is not authorized to read project -
// INVALID_ARGUMENT - if the request is malformed - NOT_FOUND - if the
// History does not exist
func (r *ProjectsHistoriesService) List(projectId string) *ProjectsHistoriesListCall {
	c := &ProjectsHistoriesListCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	return c
}

// FilterByName sets the optional parameter "filterByName": If set, only
// return histories with the given name.
func (c *ProjectsHistoriesListCall) FilterByName(filterByName string) *ProjectsHistoriesListCall {
	c.opt_["filterByName"] = filterByName
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of Histories to fetch.
//
// Default value: 20. The server will use this
// default if the field is not set or has a value of 0. Any value
// greater than 100 will be treated as 100.
func (c *ProjectsHistoriesListCall) PageSize(pageSize int64) *ProjectsHistoriesListCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken": A continuation
// token to resume the query at the next item.
func (c *ProjectsHistoriesListCall) PageToken(pageToken string) *ProjectsHistoriesListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsHistoriesListCall) Fields(s ...googleapi.Field) *ProjectsHistoriesListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsHistoriesListCall) Do() (*ListHistoriesResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["filterByName"]; ok {
		params.Set("filterByName", fmt.Sprintf("%v", v))
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "{projectId}/histories")
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
	var ret *ListHistoriesResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists Histories for a given Project.\n\nThe histories are sorted by modification time in descending order. The history_id key will be used to order the history with the same modification time.\n\nMay return any of the following canonical error codes:\n\n- PERMISSION_DENIED - if the user is not authorized to read project - INVALID_ARGUMENT - if the request is malformed - NOT_FOUND - if the History does not exist",
	//   "httpMethod": "GET",
	//   "id": "toolresults.projects.histories.list",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "filterByName": {
	//       "description": "If set, only return histories with the given name.\n\nOptional.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "The maximum number of Histories to fetch.\n\nDefault value: 20. The server will use this default if the field is not set or has a value of 0. Any value greater than 100 will be treated as 100.\n\nOptional.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A continuation token to resume the query at the next item.\n\nOptional.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "A Project id.\n\nRequired.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{projectId}/histories",
	//   "response": {
	//     "$ref": "ListHistoriesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "toolresults.projects.histories.executions.create":

type ProjectsHistoriesExecutionsCreateCall struct {
	s         *Service
	projectId string
	historyId string
	execution *Execution
	opt_      map[string]interface{}
}

// Create: Creates an Execution.
//
// The returned Execution will have the
// id set.
//
// May return any of the following canonical error codes:
//
// -
// PERMISSION_DENIED - if the user is not authorized to write to project
// - INVALID_ARGUMENT - if the request is malformed - NOT_FOUND - if the
// containing History does not exist
func (r *ProjectsHistoriesExecutionsService) Create(projectId string, historyId string, execution *Execution) *ProjectsHistoriesExecutionsCreateCall {
	c := &ProjectsHistoriesExecutionsCreateCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.historyId = historyId
	c.execution = execution
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsHistoriesExecutionsCreateCall) Fields(s ...googleapi.Field) *ProjectsHistoriesExecutionsCreateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsHistoriesExecutionsCreateCall) Do() (*Execution, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.execution)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{projectId}/histories/{historyId}/executions")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"historyId": c.historyId,
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
	var ret *Execution
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates an Execution.\n\nThe returned Execution will have the id set.\n\nMay return any of the following canonical error codes:\n\n- PERMISSION_DENIED - if the user is not authorized to write to project - INVALID_ARGUMENT - if the request is malformed - NOT_FOUND - if the containing History does not exist",
	//   "httpMethod": "POST",
	//   "id": "toolresults.projects.histories.executions.create",
	//   "parameterOrder": [
	//     "projectId",
	//     "historyId"
	//   ],
	//   "parameters": {
	//     "historyId": {
	//       "description": "A History id.\n\nRequired.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "A Project id.\n\nRequired.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{projectId}/histories/{historyId}/executions",
	//   "request": {
	//     "$ref": "Execution"
	//   },
	//   "response": {
	//     "$ref": "Execution"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "toolresults.projects.histories.executions.get":

type ProjectsHistoriesExecutionsGetCall struct {
	s           *Service
	projectId   string
	historyId   string
	executionId string
	opt_        map[string]interface{}
}

// Get: Gets an Execution.
//
// May return any of the following canonical
// error codes:
//
// - PERMISSION_DENIED - if the user is not authorized to
// write to project - INVALID_ARGUMENT - if the request is malformed -
// NOT_FOUND - if the Execution does not exist
func (r *ProjectsHistoriesExecutionsService) Get(projectId string, historyId string, executionId string) *ProjectsHistoriesExecutionsGetCall {
	c := &ProjectsHistoriesExecutionsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.historyId = historyId
	c.executionId = executionId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsHistoriesExecutionsGetCall) Fields(s ...googleapi.Field) *ProjectsHistoriesExecutionsGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsHistoriesExecutionsGetCall) Do() (*Execution, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{projectId}/histories/{historyId}/executions/{executionId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId":   c.projectId,
		"historyId":   c.historyId,
		"executionId": c.executionId,
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
	var ret *Execution
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets an Execution.\n\nMay return any of the following canonical error codes:\n\n- PERMISSION_DENIED - if the user is not authorized to write to project - INVALID_ARGUMENT - if the request is malformed - NOT_FOUND - if the Execution does not exist",
	//   "httpMethod": "GET",
	//   "id": "toolresults.projects.histories.executions.get",
	//   "parameterOrder": [
	//     "projectId",
	//     "historyId",
	//     "executionId"
	//   ],
	//   "parameters": {
	//     "executionId": {
	//       "description": "An Execution id.\n\nRequired.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "historyId": {
	//       "description": "A History id.\n\nRequired.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "A Project id.\n\nRequired.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{projectId}/histories/{historyId}/executions/{executionId}",
	//   "response": {
	//     "$ref": "Execution"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "toolresults.projects.histories.executions.list":

type ProjectsHistoriesExecutionsListCall struct {
	s         *Service
	projectId string
	historyId string
	opt_      map[string]interface{}
}

// List: Lists Histories for a given Project.
//
// The executions are sorted
// by creation_time in descending order. The execution_id key will be
// used to order the executions with the same creation_time.
//
// May return
// any of the following canonical error codes:
//
// - PERMISSION_DENIED - if
// the user is not authorized to read project - INVALID_ARGUMENT - if
// the request is malformed - NOT_FOUND - if the containing History does
// not exist
func (r *ProjectsHistoriesExecutionsService) List(projectId string, historyId string) *ProjectsHistoriesExecutionsListCall {
	c := &ProjectsHistoriesExecutionsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.historyId = historyId
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of Executions to fetch.
//
// Default value: 25. The server will use this
// default if the field is not set or has a value of 0.
func (c *ProjectsHistoriesExecutionsListCall) PageSize(pageSize int64) *ProjectsHistoriesExecutionsListCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken": A continuation
// token to resume the query at the next item.
func (c *ProjectsHistoriesExecutionsListCall) PageToken(pageToken string) *ProjectsHistoriesExecutionsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsHistoriesExecutionsListCall) Fields(s ...googleapi.Field) *ProjectsHistoriesExecutionsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsHistoriesExecutionsListCall) Do() (*ListExecutionsResponse, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "{projectId}/histories/{historyId}/executions")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId": c.projectId,
		"historyId": c.historyId,
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
	var ret *ListExecutionsResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists Histories for a given Project.\n\nThe executions are sorted by creation_time in descending order. The execution_id key will be used to order the executions with the same creation_time.\n\nMay return any of the following canonical error codes:\n\n- PERMISSION_DENIED - if the user is not authorized to read project - INVALID_ARGUMENT - if the request is malformed - NOT_FOUND - if the containing History does not exist",
	//   "httpMethod": "GET",
	//   "id": "toolresults.projects.histories.executions.list",
	//   "parameterOrder": [
	//     "projectId",
	//     "historyId"
	//   ],
	//   "parameters": {
	//     "historyId": {
	//       "description": "A History id.\n\nRequired.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "The maximum number of Executions to fetch.\n\nDefault value: 25. The server will use this default if the field is not set or has a value of 0.\n\nOptional.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A continuation token to resume the query at the next item.\n\nOptional.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "A Project id.\n\nRequired.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{projectId}/histories/{historyId}/executions",
	//   "response": {
	//     "$ref": "ListExecutionsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "toolresults.projects.histories.executions.patch":

type ProjectsHistoriesExecutionsPatchCall struct {
	s           *Service
	projectId   string
	historyId   string
	executionId string
	execution   *Execution
	opt_        map[string]interface{}
}

// Patch: Updates an existing Execution with the supplied partial
// entity.
//
// May return any of the following canonical error codes:
//
// -
// PERMISSION_DENIED - if the user is not authorized to write to project
// - INVALID_ARGUMENT - if the request is malformed -
// FAILED_PRECONDITION - if the requested state transition is illegal -
// NOT_FOUND - if the containing History does not exist
func (r *ProjectsHistoriesExecutionsService) Patch(projectId string, historyId string, executionId string, execution *Execution) *ProjectsHistoriesExecutionsPatchCall {
	c := &ProjectsHistoriesExecutionsPatchCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.historyId = historyId
	c.executionId = executionId
	c.execution = execution
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsHistoriesExecutionsPatchCall) Fields(s ...googleapi.Field) *ProjectsHistoriesExecutionsPatchCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsHistoriesExecutionsPatchCall) Do() (*Execution, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.execution)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{projectId}/histories/{historyId}/executions/{executionId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId":   c.projectId,
		"historyId":   c.historyId,
		"executionId": c.executionId,
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
	var ret *Execution
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates an existing Execution with the supplied partial entity.\n\nMay return any of the following canonical error codes:\n\n- PERMISSION_DENIED - if the user is not authorized to write to project - INVALID_ARGUMENT - if the request is malformed - FAILED_PRECONDITION - if the requested state transition is illegal - NOT_FOUND - if the containing History does not exist",
	//   "httpMethod": "PATCH",
	//   "id": "toolresults.projects.histories.executions.patch",
	//   "parameterOrder": [
	//     "projectId",
	//     "historyId",
	//     "executionId"
	//   ],
	//   "parameters": {
	//     "executionId": {
	//       "description": "Required.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "historyId": {
	//       "description": "Required.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "A Project id. Required.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{projectId}/histories/{historyId}/executions/{executionId}",
	//   "request": {
	//     "$ref": "Execution"
	//   },
	//   "response": {
	//     "$ref": "Execution"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "toolresults.projects.histories.executions.steps.create":

type ProjectsHistoriesExecutionsStepsCreateCall struct {
	s           *Service
	projectId   string
	historyId   string
	executionId string
	step        *Step
	opt_        map[string]interface{}
}

// Create: Creates a Step.
//
// The returned Step will have the id set.
//
// May
// return any of the following canonical error codes:
//
// -
// PERMISSION_DENIED - if the user is not authorized to write to project
// - INVALID_ARGUMENT - if the request is malformed -
// FAILED_PRECONDITION - if the step is too large (more than 10Mib) -
// NOT_FOUND - if the containing Execution does not exist
func (r *ProjectsHistoriesExecutionsStepsService) Create(projectId string, historyId string, executionId string, step *Step) *ProjectsHistoriesExecutionsStepsCreateCall {
	c := &ProjectsHistoriesExecutionsStepsCreateCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.historyId = historyId
	c.executionId = executionId
	c.step = step
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsHistoriesExecutionsStepsCreateCall) Fields(s ...googleapi.Field) *ProjectsHistoriesExecutionsStepsCreateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsHistoriesExecutionsStepsCreateCall) Do() (*Step, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.step)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{projectId}/histories/{historyId}/executions/{executionId}/steps")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId":   c.projectId,
		"historyId":   c.historyId,
		"executionId": c.executionId,
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
	var ret *Step
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a Step.\n\nThe returned Step will have the id set.\n\nMay return any of the following canonical error codes:\n\n- PERMISSION_DENIED - if the user is not authorized to write to project - INVALID_ARGUMENT - if the request is malformed - FAILED_PRECONDITION - if the step is too large (more than 10Mib) - NOT_FOUND - if the containing Execution does not exist",
	//   "httpMethod": "POST",
	//   "id": "toolresults.projects.histories.executions.steps.create",
	//   "parameterOrder": [
	//     "projectId",
	//     "historyId",
	//     "executionId"
	//   ],
	//   "parameters": {
	//     "executionId": {
	//       "description": "A Execution id.\n\nRequired.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "historyId": {
	//       "description": "A History id.\n\nRequired.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "A Project id.\n\nRequired.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{projectId}/histories/{historyId}/executions/{executionId}/steps",
	//   "request": {
	//     "$ref": "Step"
	//   },
	//   "response": {
	//     "$ref": "Step"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "toolresults.projects.histories.executions.steps.get":

type ProjectsHistoriesExecutionsStepsGetCall struct {
	s           *Service
	projectId   string
	historyId   string
	executionId string
	stepId      string
	opt_        map[string]interface{}
}

// Get: Gets a Step.
//
// May return any of the following canonical error
// codes:
//
// - PERMISSION_DENIED - if the user is not authorized to read
// project - INVALID_ARGUMENT - if the request is malformed - NOT_FOUND
// - if the Step does not exist
func (r *ProjectsHistoriesExecutionsStepsService) Get(projectId string, historyId string, executionId string, stepId string) *ProjectsHistoriesExecutionsStepsGetCall {
	c := &ProjectsHistoriesExecutionsStepsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.historyId = historyId
	c.executionId = executionId
	c.stepId = stepId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsHistoriesExecutionsStepsGetCall) Fields(s ...googleapi.Field) *ProjectsHistoriesExecutionsStepsGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsHistoriesExecutionsStepsGetCall) Do() (*Step, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{projectId}/histories/{historyId}/executions/{executionId}/steps/{stepId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId":   c.projectId,
		"historyId":   c.historyId,
		"executionId": c.executionId,
		"stepId":      c.stepId,
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
	var ret *Step
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets a Step.\n\nMay return any of the following canonical error codes:\n\n- PERMISSION_DENIED - if the user is not authorized to read project - INVALID_ARGUMENT - if the request is malformed - NOT_FOUND - if the Step does not exist",
	//   "httpMethod": "GET",
	//   "id": "toolresults.projects.histories.executions.steps.get",
	//   "parameterOrder": [
	//     "projectId",
	//     "historyId",
	//     "executionId",
	//     "stepId"
	//   ],
	//   "parameters": {
	//     "executionId": {
	//       "description": "A Execution id.\n\nRequired.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "historyId": {
	//       "description": "A History id.\n\nRequired.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "A Project id.\n\nRequired.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "stepId": {
	//       "description": "A Step id.\n\nRequired.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{projectId}/histories/{historyId}/executions/{executionId}/steps/{stepId}",
	//   "response": {
	//     "$ref": "Step"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "toolresults.projects.histories.executions.steps.list":

type ProjectsHistoriesExecutionsStepsListCall struct {
	s           *Service
	projectId   string
	historyId   string
	executionId string
	opt_        map[string]interface{}
}

// List: Lists Steps for a given Execution.
//
// The steps are sorted by
// creation_time in descending order. The step_id key will be used to
// order the steps with the same creation_time.
//
// May return any of the
// following canonical error codes:
//
// - PERMISSION_DENIED - if the user
// is not authorized to read project - INVALID_ARGUMENT - if the request
// is malformed - FAILED_PRECONDITION - if an argument in the request
// happens to be invalid; e.g. if an attempt is made to list the
// children of a nonexistent Step - NOT_FOUND - if the containing
// Execution does not exist
func (r *ProjectsHistoriesExecutionsStepsService) List(projectId string, historyId string, executionId string) *ProjectsHistoriesExecutionsStepsListCall {
	c := &ProjectsHistoriesExecutionsStepsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.historyId = historyId
	c.executionId = executionId
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of Steps to fetch.
//
// Default value: 25. The server will use this
// default if the field is not set or has a value of 0.
func (c *ProjectsHistoriesExecutionsStepsListCall) PageSize(pageSize int64) *ProjectsHistoriesExecutionsStepsListCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken": A continuation
// token to resume the query at the next item.
func (c *ProjectsHistoriesExecutionsStepsListCall) PageToken(pageToken string) *ProjectsHistoriesExecutionsStepsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsHistoriesExecutionsStepsListCall) Fields(s ...googleapi.Field) *ProjectsHistoriesExecutionsStepsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsHistoriesExecutionsStepsListCall) Do() (*ListStepsResponse, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "{projectId}/histories/{historyId}/executions/{executionId}/steps")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId":   c.projectId,
		"historyId":   c.historyId,
		"executionId": c.executionId,
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
	var ret *ListStepsResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists Steps for a given Execution.\n\nThe steps are sorted by creation_time in descending order. The step_id key will be used to order the steps with the same creation_time.\n\nMay return any of the following canonical error codes:\n\n- PERMISSION_DENIED - if the user is not authorized to read project - INVALID_ARGUMENT - if the request is malformed - FAILED_PRECONDITION - if an argument in the request happens to be invalid; e.g. if an attempt is made to list the children of a nonexistent Step - NOT_FOUND - if the containing Execution does not exist",
	//   "httpMethod": "GET",
	//   "id": "toolresults.projects.histories.executions.steps.list",
	//   "parameterOrder": [
	//     "projectId",
	//     "historyId",
	//     "executionId"
	//   ],
	//   "parameters": {
	//     "executionId": {
	//       "description": "A Execution id.\n\nRequired.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "historyId": {
	//       "description": "A History id.\n\nRequired.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "The maximum number of Steps to fetch.\n\nDefault value: 25. The server will use this default if the field is not set or has a value of 0.\n\nOptional.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A continuation token to resume the query at the next item.\n\nOptional.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "A Project id.\n\nRequired.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{projectId}/histories/{historyId}/executions/{executionId}/steps",
	//   "response": {
	//     "$ref": "ListStepsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "toolresults.projects.histories.executions.steps.patch":

type ProjectsHistoriesExecutionsStepsPatchCall struct {
	s           *Service
	projectId   string
	historyId   string
	executionId string
	stepId      string
	step        *Step
	opt_        map[string]interface{}
}

// Patch: Updates an existing Step with the supplied partial
// entity.
//
// May return any of the following canonical error codes:
//
// -
// PERMISSION_DENIED - if the user is not authorized to write project -
// INVALID_ARGUMENT - if the request is malformed - FAILED_PRECONDITION
// - if the requested state transition is illegal (e.g try to upload a
// duplicate xml file), if the updated step is too large (more than
// 10Mib) - NOT_FOUND - if the containing Execution does not exist
func (r *ProjectsHistoriesExecutionsStepsService) Patch(projectId string, historyId string, executionId string, stepId string, step *Step) *ProjectsHistoriesExecutionsStepsPatchCall {
	c := &ProjectsHistoriesExecutionsStepsPatchCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.historyId = historyId
	c.executionId = executionId
	c.stepId = stepId
	c.step = step
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsHistoriesExecutionsStepsPatchCall) Fields(s ...googleapi.Field) *ProjectsHistoriesExecutionsStepsPatchCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsHistoriesExecutionsStepsPatchCall) Do() (*Step, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.step)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{projectId}/histories/{historyId}/executions/{executionId}/steps/{stepId}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId":   c.projectId,
		"historyId":   c.historyId,
		"executionId": c.executionId,
		"stepId":      c.stepId,
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
	var ret *Step
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates an existing Step with the supplied partial entity.\n\nMay return any of the following canonical error codes:\n\n- PERMISSION_DENIED - if the user is not authorized to write project - INVALID_ARGUMENT - if the request is malformed - FAILED_PRECONDITION - if the requested state transition is illegal (e.g try to upload a duplicate xml file), if the updated step is too large (more than 10Mib) - NOT_FOUND - if the containing Execution does not exist",
	//   "httpMethod": "PATCH",
	//   "id": "toolresults.projects.histories.executions.steps.patch",
	//   "parameterOrder": [
	//     "projectId",
	//     "historyId",
	//     "executionId",
	//     "stepId"
	//   ],
	//   "parameters": {
	//     "executionId": {
	//       "description": "A Execution id.\n\nRequired.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "historyId": {
	//       "description": "A History id.\n\nRequired.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "A Project id.\n\nRequired.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "stepId": {
	//       "description": "A Step id.\n\nRequired.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{projectId}/histories/{historyId}/executions/{executionId}/steps/{stepId}",
	//   "request": {
	//     "$ref": "Step"
	//   },
	//   "response": {
	//     "$ref": "Step"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "toolresults.projects.histories.executions.steps.publishXunitXmlFiles":

type ProjectsHistoriesExecutionsStepsPublishXunitXmlFilesCall struct {
	s                           *Service
	projectId                   string
	historyId                   string
	executionId                 string
	stepId                      string
	publishxunitxmlfilesrequest *PublishXunitXmlFilesRequest
	opt_                        map[string]interface{}
}

// PublishXunitXmlFiles: Publish xml files to an existing Step.
//
// May
// return any of the following canonical error codes:
//
// -
// PERMISSION_DENIED - if the user is not authorized to write project -
// INVALID_ARGUMENT - if the request is malformed - FAILED_PRECONDITION
// - if the requested state transition is illegal, e.g try to upload a
// duplicate xml file or a file too large. - NOT_FOUND - if the
// containing Execution does not exist
func (r *ProjectsHistoriesExecutionsStepsService) PublishXunitXmlFiles(projectId string, historyId string, executionId string, stepId string, publishxunitxmlfilesrequest *PublishXunitXmlFilesRequest) *ProjectsHistoriesExecutionsStepsPublishXunitXmlFilesCall {
	c := &ProjectsHistoriesExecutionsStepsPublishXunitXmlFilesCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.historyId = historyId
	c.executionId = executionId
	c.stepId = stepId
	c.publishxunitxmlfilesrequest = publishxunitxmlfilesrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsHistoriesExecutionsStepsPublishXunitXmlFilesCall) Fields(s ...googleapi.Field) *ProjectsHistoriesExecutionsStepsPublishXunitXmlFilesCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsHistoriesExecutionsStepsPublishXunitXmlFilesCall) Do() (*Step, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.publishxunitxmlfilesrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{projectId}/histories/{historyId}/executions/{executionId}/steps/{stepId}:publishXunitXmlFiles")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId":   c.projectId,
		"historyId":   c.historyId,
		"executionId": c.executionId,
		"stepId":      c.stepId,
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
	var ret *Step
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Publish xml files to an existing Step.\n\nMay return any of the following canonical error codes:\n\n- PERMISSION_DENIED - if the user is not authorized to write project - INVALID_ARGUMENT - if the request is malformed - FAILED_PRECONDITION - if the requested state transition is illegal, e.g try to upload a duplicate xml file or a file too large. - NOT_FOUND - if the containing Execution does not exist",
	//   "httpMethod": "POST",
	//   "id": "toolresults.projects.histories.executions.steps.publishXunitXmlFiles",
	//   "parameterOrder": [
	//     "projectId",
	//     "historyId",
	//     "executionId",
	//     "stepId"
	//   ],
	//   "parameters": {
	//     "executionId": {
	//       "description": "A Execution id.\n\nRequired.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "historyId": {
	//       "description": "A History id.\n\nRequired.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "A Project id.\n\nRequired.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "stepId": {
	//       "description": "A Step id. Note: This step must include a TestExecutionStep.\n\nRequired.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{projectId}/histories/{historyId}/executions/{executionId}/steps/{stepId}:publishXunitXmlFiles",
	//   "request": {
	//     "$ref": "PublishXunitXmlFilesRequest"
	//   },
	//   "response": {
	//     "$ref": "Step"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "toolresults.projects.histories.executions.steps.thumbnails.list":

type ProjectsHistoriesExecutionsStepsThumbnailsListCall struct {
	s           *Service
	projectId   string
	historyId   string
	executionId string
	stepId      string
	opt_        map[string]interface{}
}

// List: Lists thumbnails of images attached to a step.
//
// May return any
// of the following canonical error codes: - PERMISSION_DENIED - if the
// user is not authorized to read from the project, or from any of the
// images - INVALID_ARGUMENT - if the request is malformed - NOT_FOUND -
// if the step does not exist, or if any of the images do not exist
func (r *ProjectsHistoriesExecutionsStepsThumbnailsService) List(projectId string, historyId string, executionId string, stepId string) *ProjectsHistoriesExecutionsStepsThumbnailsListCall {
	c := &ProjectsHistoriesExecutionsStepsThumbnailsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	c.historyId = historyId
	c.executionId = executionId
	c.stepId = stepId
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of thumbnails to fetch.
//
// Default value: 50. The server will use this
// default if the field is not set or has a value of 0.
func (c *ProjectsHistoriesExecutionsStepsThumbnailsListCall) PageSize(pageSize int64) *ProjectsHistoriesExecutionsStepsThumbnailsListCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken": A continuation
// token to resume the query at the next item.
func (c *ProjectsHistoriesExecutionsStepsThumbnailsListCall) PageToken(pageToken string) *ProjectsHistoriesExecutionsStepsThumbnailsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsHistoriesExecutionsStepsThumbnailsListCall) Fields(s ...googleapi.Field) *ProjectsHistoriesExecutionsStepsThumbnailsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsHistoriesExecutionsStepsThumbnailsListCall) Do() (*ListStepThumbnailsResponse, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "{projectId}/histories/{historyId}/executions/{executionId}/steps/{stepId}/thumbnails")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"projectId":   c.projectId,
		"historyId":   c.historyId,
		"executionId": c.executionId,
		"stepId":      c.stepId,
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
	var ret *ListStepThumbnailsResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists thumbnails of images attached to a step.\n\nMay return any of the following canonical error codes: - PERMISSION_DENIED - if the user is not authorized to read from the project, or from any of the images - INVALID_ARGUMENT - if the request is malformed - NOT_FOUND - if the step does not exist, or if any of the images do not exist",
	//   "httpMethod": "GET",
	//   "id": "toolresults.projects.histories.executions.steps.thumbnails.list",
	//   "parameterOrder": [
	//     "projectId",
	//     "historyId",
	//     "executionId",
	//     "stepId"
	//   ],
	//   "parameters": {
	//     "executionId": {
	//       "description": "An Execution id.\n\nRequired.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "historyId": {
	//       "description": "A History id.\n\nRequired.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "The maximum number of thumbnails to fetch.\n\nDefault value: 50. The server will use this default if the field is not set or has a value of 0.\n\nOptional.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A continuation token to resume the query at the next item.\n\nOptional.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectId": {
	//       "description": "A Project id.\n\nRequired.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "stepId": {
	//       "description": "A Step id.\n\nRequired.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{projectId}/histories/{historyId}/executions/{executionId}/steps/{stepId}/thumbnails",
	//   "response": {
	//     "$ref": "ListStepThumbnailsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}
