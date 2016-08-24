// Package ml provides access to the Google Cloud Machine Learning.
//
// See https://cloud.google.com/ml/
//
// Usage example:
//
//   import "github.com/skelterjohn/gcloud_apis/clients/ml/v1alpha3"
//   ...
//   mlService, err := ml.New(oauthHttpClient)
package ml

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

const apiId = "ml:v1alpha3"
const apiName = "ml"
const apiVersion = "v1alpha3"
const basePath = "https://ml.googleapis.com/"

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
	rs.Models = NewProjectsModelsService(s)
	rs.Operations = NewProjectsOperationsService(s)
	return rs
}

type ProjectsService struct {
	s *Service

	Models *ProjectsModelsService

	Operations *ProjectsOperationsService
}

func NewProjectsModelsService(s *Service) *ProjectsModelsService {
	rs := &ProjectsModelsService{s: s}
	rs.Versions = NewProjectsModelsVersionsService(s)
	return rs
}

type ProjectsModelsService struct {
	s *Service

	Versions *ProjectsModelsVersionsService
}

func NewProjectsModelsVersionsService(s *Service) *ProjectsModelsVersionsService {
	rs := &ProjectsModelsVersionsService{s: s}
	return rs
}

type ProjectsModelsVersionsService struct {
	s *Service
}

func NewProjectsOperationsService(s *Service) *ProjectsOperationsService {
	rs := &ProjectsOperationsService{s: s}
	return rs
}

type ProjectsOperationsService struct {
	s *Service
}

// GoogleCloudMlV1alpha3__CreateVersionRequest: Uploads the provided
// trained model version to Cloud Machine Learning.
type GoogleCloudMlV1alpha3__CreateVersionRequest struct {
	// Parent: The name of the model.
	Parent string `json:"parent,omitempty"`

	// Version: The version details.
	Version *GoogleCloudMlV1alpha3__Version `json:"version,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Parent") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *GoogleCloudMlV1alpha3__CreateVersionRequest) MarshalJSON() ([]byte, error) {
	type noMethod GoogleCloudMlV1alpha3__CreateVersionRequest
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// GoogleCloudMlV1alpha3__GetConfigResponse: Returns service
// configuration associated with a given project.
type GoogleCloudMlV1alpha3__GetConfigResponse struct {
	// ServiceAccount: The service account Cloud ML uses to access resources
	// in the project.
	ServiceAccount string `json:"serviceAccount,omitempty"`

	// ServiceAccountProject: Project number associated with
	// 'service_account'.
	ServiceAccountProject int64 `json:"serviceAccountProject,omitempty,string"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "ServiceAccount") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *GoogleCloudMlV1alpha3__GetConfigResponse) MarshalJSON() ([]byte, error) {
	type noMethod GoogleCloudMlV1alpha3__GetConfigResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// GoogleCloudMlV1alpha3__GetHyperparametersRequest: Returns the
// hyperparameters associated with the given run.
type GoogleCloudMlV1alpha3__GetHyperparametersRequest struct {
	// JobName: The job name. Must be unique within the project.
	JobName string `json:"jobName,omitempty"`

	// RunId: The run id for these results.
	RunId string `json:"runId,omitempty"`

	// ForceSendFields is a list of field names (e.g. "JobName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *GoogleCloudMlV1alpha3__GetHyperparametersRequest) MarshalJSON() ([]byte, error) {
	type noMethod GoogleCloudMlV1alpha3__GetHyperparametersRequest
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// GoogleCloudMlV1alpha3__GetHyperparametersResponse: Returns the
// hyperparameters associated with the given run.
type GoogleCloudMlV1alpha3__GetHyperparametersResponse struct {
	// Hyperparameters: The hyperparameters given to the requested run.
	Hyperparameters map[string]string `json:"hyperparameters,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Hyperparameters") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *GoogleCloudMlV1alpha3__GetHyperparametersResponse) MarshalJSON() ([]byte, error) {
	type noMethod GoogleCloudMlV1alpha3__GetHyperparametersResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// GoogleCloudMlV1alpha3__HyperparameterJobResult: Represents the result
// of a hyperparameter tuning run from a training job.
type GoogleCloudMlV1alpha3__HyperparameterJobResult struct {
	// BestObjectiveValue: The best objective value seen for this run.
	BestObjectiveValue float64 `json:"bestObjectiveValue,omitempty"`

	// BestTrainingStep: The training step associated with the best
	// objective value.
	BestTrainingStep int64 `json:"bestTrainingStep,omitempty,string"`

	// Hyperparameters: The hyperparameters given to this run.
	Hyperparameters map[string]string `json:"hyperparameters,omitempty"`

	// RunId: The run id for these results.
	RunId string `json:"runId,omitempty"`

	// ForceSendFields is a list of field names (e.g. "BestObjectiveValue")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *GoogleCloudMlV1alpha3__HyperparameterJobResult) MarshalJSON() ([]byte, error) {
	type noMethod GoogleCloudMlV1alpha3__HyperparameterJobResult
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// GoogleCloudMlV1alpha3__HyperparameterSpec: Represents a set of
// hyperparameters to optimize.
type GoogleCloudMlV1alpha3__HyperparameterSpec struct {
	// Goal: Should the evaluation metric be maximized or minimized?
	//
	// Possible values:
	//   "GOAL_TYPE_UNSPECIFIED" - Goal Type will default to maximize.
	//   "MAXIMIZE" - Maximize the goal metric.
	//   "MINIMIZE" - Minimize the goal metric.
	Goal string `json:"goal,omitempty"`

	// MaxParallelRuns: How many training runs should be run in parallel.
	// More parallelization
	// will be faster, but parallel runs only benefit from the information
	// gained
	// by previous runs. Each run will use the number of workers specified
	// by
	// the worker, parameter server, and evaluator specs. Defaults to one.
	MaxParallelRuns int64 `json:"maxParallelRuns,omitempty"`

	// MaxRuns: How many training runs should be attempted to optimize.
	// Defaults to one.
	MaxRuns int64 `json:"maxRuns,omitempty"`

	// Params: The set of parameters to tune.
	Params []*GoogleCloudMlV1alpha3__ParameterConfig `json:"params,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Goal") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *GoogleCloudMlV1alpha3__HyperparameterSpec) MarshalJSON() ([]byte, error) {
	type noMethod GoogleCloudMlV1alpha3__HyperparameterSpec
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// GoogleCloudMlV1alpha3__JobMetadata: Represents the metadata of the
// longrunning.Operation created by
// the SubmitTrainingJob or SubmitPredictionJob method.
type GoogleCloudMlV1alpha3__JobMetadata struct {
	// CreateTime: When the job was submitted.
	CreateTime string `json:"createTime,omitempty"`

	// CreateVersionRequest: The create version job request recorded.
	// The create version job request recorded.
	CreateVersionRequest *GoogleCloudMlV1alpha3__CreateVersionRequest `json:"createVersionRequest,omitempty"`

	// EndTime: When the job processing was completed.
	EndTime string `json:"endTime,omitempty"`

	// IsCancellationRequested: Whether the cancellation of this job has
	// been requested.
	IsCancellationRequested bool `json:"isCancellationRequested,omitempty"`

	// JobState: The detailed state of a job.
	//
	// Possible values:
	//   "UNKNOWN_JOB_STATE"
	//   "QUEUED"
	//   "PREPARING"
	//   "RUNNING"
	//   "SUCCEEDED"
	//   "FAILED"
	//   "CANCELLED"
	JobState string `json:"jobState,omitempty"`

	// PredictionRequest: The prediction job request recorded.
	PredictionRequest *GoogleCloudMlV1alpha3__SubmitPredictionJobRequest `json:"predictionRequest,omitempty"`

	// PredictionResult: The current prediction job result.
	PredictionResult *GoogleCloudMlV1alpha3__PredictionJobResult `json:"predictionResult,omitempty"`

	// StartTime: When the job processing was started.
	StartTime string `json:"startTime,omitempty"`

	// StatusMetricsSnapshot: Progress report. A set of key-value pairs that
	// convey the current
	// state of the job.
	StatusMetricsSnapshot map[string]string `json:"statusMetricsSnapshot,omitempty"`

	// StatusMetricsSnapshotTime: The time at which the 'status' field was
	// populated.
	StatusMetricsSnapshotTime string `json:"statusMetricsSnapshotTime,omitempty"`

	// TrainingRequest: The training job request recorded.
	TrainingRequest *GoogleCloudMlV1alpha3__SubmitTrainingJobRequest `json:"trainingRequest,omitempty"`

	// TrainingResult: The current training job result.
	TrainingResult *GoogleCloudMlV1alpha3__TrainingJobResult `json:"trainingResult,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CreateTime") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *GoogleCloudMlV1alpha3__JobMetadata) MarshalJSON() ([]byte, error) {
	type noMethod GoogleCloudMlV1alpha3__JobMetadata
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// GoogleCloudMlV1alpha3__ListModelsResponse: Response message for the
// ListModels method.
type GoogleCloudMlV1alpha3__ListModelsResponse struct {
	// Models: The list of models.
	Models []*GoogleCloudMlV1alpha3__Model `json:"models,omitempty"`

	// NextPageToken: Optional pagination token to use for retrieving the
	// next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Models") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *GoogleCloudMlV1alpha3__ListModelsResponse) MarshalJSON() ([]byte, error) {
	type noMethod GoogleCloudMlV1alpha3__ListModelsResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// GoogleCloudMlV1alpha3__ListVersionsResponse: Response message for the
// ListVersions method.
type GoogleCloudMlV1alpha3__ListVersionsResponse struct {
	// NextPageToken: Optional pagination token to use for retrieving the
	// next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Versions: The list of versions.
	Versions []*GoogleCloudMlV1alpha3__Version `json:"versions,omitempty"`

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

func (s *GoogleCloudMlV1alpha3__ListVersionsResponse) MarshalJSON() ([]byte, error) {
	type noMethod GoogleCloudMlV1alpha3__ListVersionsResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// GoogleCloudMlV1alpha3__MasterSpec: Represents a specification of
// master worker.
type GoogleCloudMlV1alpha3__MasterSpec struct {
	// ReplicaCount: The required number of master worker servers. Defaults
	// to 1.
	// If the number of replicas is greater than 1, the first master
	// worker
	// will be the one monitored for success/failure.
	ReplicaCount uint64 `json:"replicaCount,omitempty,string"`

	// Resources: Resource requirements for a master.
	// If not specified, default master configuration will be used.
	Resources *GoogleCloudMlV1alpha3__ResourceSpec `json:"resources,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ReplicaCount") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *GoogleCloudMlV1alpha3__MasterSpec) MarshalJSON() ([]byte, error) {
	type noMethod GoogleCloudMlV1alpha3__MasterSpec
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// GoogleCloudMlV1alpha3__Model: Represents a machine learning model
// resource that can be used to perform
// prediction.
type GoogleCloudMlV1alpha3__Model struct {
	// DefaultVersion: The default version of the model.
	DefaultVersion *GoogleCloudMlV1alpha3__Version `json:"defaultVersion,omitempty"`

	// Name: The user-specified name of the model.
	Name string `json:"name,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "DefaultVersion") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *GoogleCloudMlV1alpha3__Model) MarshalJSON() ([]byte, error) {
	type noMethod GoogleCloudMlV1alpha3__Model
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// GoogleCloudMlV1alpha3__ParameterConfig: Represents a single
// hyperparameter to optimize.
type GoogleCloudMlV1alpha3__ParameterConfig struct {
	// Categories: The list of possible categories. REQUIRED if
	// type==CATEGORICAL.
	Categories []string `json:"categories,omitempty"`

	// FeasiblePoints: A list of feasible points, to be provided if and only
	// if type==DISCRETE.
	// The list should be in strictly increasing order. For instance,
	// this
	// Parameter might have possible settings of 1.5, 2.5, and 4.0. This
	// list
	// shouldn't be too large - probably not more than 1,000 points.
	FeasiblePoints []float64 `json:"feasiblePoints,omitempty"`

	// MaxValue: max_value is REQUIRED if type==DOUBLE or type==INTEGER.
	// This field
	// should be unset if type==CATEGORICAL. This value should be integers
	// if
	// type==INTEGER.
	MaxValue float64 `json:"maxValue,omitempty"`

	// MinValue: min_value is REQUIRED if type==DOUBLE or type==INTEGER.
	// This field
	// should be unset if type==CATEGORICAL. This value should be integers
	// if
	// type==INTEGER.
	MinValue float64 `json:"minValue,omitempty"`

	// Name: The parameter name must be unique amongst all ParameterConfigs
	// in a
	// HyperparameterConfig. E.g., "learning_rate".
	Name string `json:"name,omitempty"`

	// ScaleType: How the parameter should be scaled to the hypercube. Leave
	// unset for
	// categorical Parameters. We strongly suggest using some kind of
	// scaling for
	// real or integral parameters (e.g., UNIT_LINEAR_SCALE).
	//
	// Possible values:
	//   "NONE" - By default, no scaling is applied.
	//   "UNIT_LINEAR_SCALE" - Scales the feasible space to (0, 1) linearly.
	//   "UNIT_LOG_SCALE" - Scales the feasible space logarithmically to (0,
	// 1). The entire feasible
	// space must be strictly positive.
	//   "UNIT_REVERSE_LOG_SCALE" - Scales the feasible space "reverse"
	// logarithmically to (0, 1). The result
	// is that values close to the top of the feasible space are spread out
	// more
	// than points near the bottom. The entire feasible space must be
	// strictly
	// positive.
	ScaleType string `json:"scaleType,omitempty"`

	// Type: The type of the parameter.
	//
	// Possible values:
	//   "PARAMETER_TYPE_UNSPECIFIED" - Parameter type must be specified.
	// Unspecified values will be treated
	// as an error.
	//   "DOUBLE" - Type for real-valued parameters.
	//   "INTEGER" - Type for integral parameters.
	//   "CATEGORICAL" - The parameter is categorical, with a value chosen
	// from the categories
	// field.
	//   "DISCRETE" - The parameter is real valued, with a fixed set of
	// feasible points. If
	// type==DISCRETE, feasible_points must be provided, and
	// {min_value, max_value} will be ignored.
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Categories") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *GoogleCloudMlV1alpha3__ParameterConfig) MarshalJSON() ([]byte, error) {
	type noMethod GoogleCloudMlV1alpha3__ParameterConfig
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// GoogleCloudMlV1alpha3__ParameterServerSpec: Represents a
// specification of parameter servers.
type GoogleCloudMlV1alpha3__ParameterServerSpec struct {
	// ReplicaCount: The required number of parameter servers. Defaults to
	// 1.
	ReplicaCount uint64 `json:"replicaCount,omitempty,string"`

	// Resources: Resource requirements for a single parameter server.
	// If not specified, default parameter server configuration will be
	// used.
	Resources *GoogleCloudMlV1alpha3__ResourceSpec `json:"resources,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ReplicaCount") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *GoogleCloudMlV1alpha3__ParameterServerSpec) MarshalJSON() ([]byte, error) {
	type noMethod GoogleCloudMlV1alpha3__ParameterServerSpec
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// GoogleCloudMlV1alpha3__PredictRequest: Request for predictions to be
// issued against a trained model.
type GoogleCloudMlV1alpha3__PredictRequest struct {
	// Instances: Input data, one string per instance. The type of the input
	// data is
	// stored in model metadata.
	Instances []string `json:"instances,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Instances") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *GoogleCloudMlV1alpha3__PredictRequest) MarshalJSON() ([]byte, error) {
	type noMethod GoogleCloudMlV1alpha3__PredictRequest
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// GoogleCloudMlV1alpha3__PredictResponse: Response with prediction
// tensors produced by a trained model.
type GoogleCloudMlV1alpha3__PredictResponse struct {
	// Predictions: Predictions are produced in the same order as the input.
	Predictions []string `json:"predictions,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Predictions") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *GoogleCloudMlV1alpha3__PredictResponse) MarshalJSON() ([]byte, error) {
	type noMethod GoogleCloudMlV1alpha3__PredictResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// GoogleCloudMlV1alpha3__PredictionJobResult: Represents the result of
// the longrunning.Operation created by the
// SubmitPredictionJob method.
type GoogleCloudMlV1alpha3__PredictionJobResult struct {
	// InstancesProcessed: The number of data instances processed
	InstancesProcessed uint64 `json:"instancesProcessed,omitempty,string"`

	// OutputUri: The URI of the output dir provided at job creation time.
	OutputUri string `json:"outputUri,omitempty"`

	// ForceSendFields is a list of field names (e.g. "InstancesProcessed")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *GoogleCloudMlV1alpha3__PredictionJobResult) MarshalJSON() ([]byte, error) {
	type noMethod GoogleCloudMlV1alpha3__PredictionJobResult
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// GoogleCloudMlV1alpha3__ReportMetricRequest: Request Metric for
// hyperparameter tuning metric reporting.
type GoogleCloudMlV1alpha3__ReportMetricRequest struct {
	// Completed: Set to true to mark this run as completed.
	Completed bool `json:"completed,omitempty"`

	// JobName: The job name. Must be unique within the project.
	JobName string `json:"jobName,omitempty"`

	// ObjectiveValue: Required.  The value of the metric being optimized.
	ObjectiveValue float64 `json:"objectiveValue,omitempty"`

	// RunId: The run to report metrics for.
	RunId string `json:"runId,omitempty"`

	// TrainingStep: Optional.  What training step this metric is for.
	TrainingStep int64 `json:"trainingStep,omitempty,string"`

	// ForceSendFields is a list of field names (e.g. "Completed") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *GoogleCloudMlV1alpha3__ReportMetricRequest) MarshalJSON() ([]byte, error) {
	type noMethod GoogleCloudMlV1alpha3__ReportMetricRequest
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// GoogleCloudMlV1alpha3__ResourceSpec: Resource specification for a
// program running inside a VM.
type GoogleCloudMlV1alpha3__ResourceSpec struct {
	// Cpus: The number of CPUs.
	Cpus uint64 `json:"cpus,omitempty,string"`

	// Disk: The required disk size (in bytes).
	Disk uint64 `json:"disk,omitempty,string"`

	// Gpus: The number of GPUs. Only allowed to be set within WorkerSpec.
	Gpus uint64 `json:"gpus,omitempty,string"`

	// Ram: The required RAM size (in bytes).
	Ram uint64 `json:"ram,omitempty,string"`

	// ForceSendFields is a list of field names (e.g. "Cpus") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *GoogleCloudMlV1alpha3__ResourceSpec) MarshalJSON() ([]byte, error) {
	type noMethod GoogleCloudMlV1alpha3__ResourceSpec
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// GoogleCloudMlV1alpha3__SetDefaultVersionRequest: Request message for
// the SetDefaultVersion request.
type GoogleCloudMlV1alpha3__SetDefaultVersionRequest struct {
}

// GoogleCloudMlV1alpha3__SubmitPredictionJobRequest: Request to submit
// a batch prediction job against a trained model.
type GoogleCloudMlV1alpha3__SubmitPredictionJobRequest struct {
	// InputUris: URIs of the TFRecord input data files. May contain
	// file-system appropriate
	// wildcards (Google Cloud Storage prefixes, for example.).
	InputUris []string `json:"inputUris,omitempty"`

	// JobName: The job name. Must be unique within the project.
	JobName string `json:"jobName,omitempty"`

	// OutputUri: The uri to which output is written, a Google Cloud Storage
	// prefix
	// for example.
	OutputUri string `json:"outputUri,omitempty"`

	// Parent: The name of the model against which prediction should be
	// performed. It is
	// of the form:
	// /projects/project_id/models/model_name/versions/version_id.
	// The version information may be omitted, in which case prediction
	// is
	// performed against the default version.
	Parent string `json:"parent,omitempty"`

	// Zone: Optional. The Google Compute Engine zone to run the training
	// job in.
	// If not specified, the system will pick a zone based on the
	// job
	// requirements. It is the user's responsibility to procure quota
	// in this zone.
	Zone string `json:"zone,omitempty"`

	// ForceSendFields is a list of field names (e.g. "InputUris") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *GoogleCloudMlV1alpha3__SubmitPredictionJobRequest) MarshalJSON() ([]byte, error) {
	type noMethod GoogleCloudMlV1alpha3__SubmitPredictionJobRequest
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// GoogleCloudMlV1alpha3__SubmitTrainingJobRequest: Request message for
// the SubmitTrainingJob method.
type GoogleCloudMlV1alpha3__SubmitTrainingJobRequest struct {
	// EvalDataPaths: The URIs of the data used for evaluation.
	EvalDataPaths []string `json:"evalDataPaths,omitempty"`

	// Hyperparameters: Optional set of Hyperparameters to tune.
	Hyperparameters *GoogleCloudMlV1alpha3__HyperparameterSpec `json:"hyperparameters,omitempty"`

	// JobArgs: Additional command line arguments to pass to the program.
	JobArgs []string `json:"jobArgs,omitempty"`

	// JobName: The job name. Must be unique within the project.
	JobName string `json:"jobName,omitempty"`

	// LogLevel: Specify minimun log levels.
	//
	// Possible values:
	//   "UNKNOWN" - No log level specified.
	//   "DEBUG" - Debug logs and above will be saved.
	//   "INFO" - Info logs and above will be saved.
	//   "WARNING" - Warning logs and above will be saved.
	//   "ERROR" - Error logs and above will be saved.
	//   "CRITICAL" - Only critical logs will be saved.
	LogLevel string `json:"logLevel,omitempty"`

	// MasterSpec: Specification of master workers. The first master
	// worker
	// will be the task monitored for job success/failure.
	MasterSpec *GoogleCloudMlV1alpha3__MasterSpec `json:"masterSpec,omitempty"`

	// MetadataPath: Optional. The URI of the file containing metadata.
	MetadataPath string `json:"metadataPath,omitempty"`

	// ModuleName: The python module name to run after installing the
	// trainer_uri packages.
	// The worker will be given a flag telling the type of the
	// process:
	// worker, parameter server or evaluator. Required if docker_image is
	// not set.
	ModuleName string `json:"moduleName,omitempty"`

	// OutputPath: The URI of the Google Cloud Storage location to which
	// output should be
	// written. This URI may have additional subdirectories appended to it,
	// e.g.,
	// during hyperparameter tuning, and the possibly modified URI is passed
	// to
	// the binary as the --output_path flag. Training binaries should
	// generally
	// respect this parameter for their output and are required to do so
	// when
	// launching from Google Cloud Dataflow or when using hyperparameter
	// tuning.
	OutputPath string `json:"outputPath,omitempty"`

	// Parent: The project name.
	Parent string `json:"parent,omitempty"`

	// PsSpec: Specification of parameter servers.
	PsSpec *GoogleCloudMlV1alpha3__ParameterServerSpec `json:"psSpec,omitempty"`

	// TensorflowVersion: The version of TensorFlow to use. If unset,
	// defaults to LATEST.
	//
	// Possible values:
	//   "LATEST" - Equivalent to the latest supported version, i.e., the
	// last version
	// in this list of enums.
	//   "V_0_9_1" - Tensorflow version 0.9.1
	TensorflowVersion string `json:"tensorflowVersion,omitempty"`

	// TrainDataPaths: The URIs of the training data.
	TrainDataPaths []string `json:"trainDataPaths,omitempty"`

	// TrainerUri: Tarball Python Packages from Google Cloud Storage with
	// the training program
	// and any additional dependencies. At least one is required if
	// docker_image
	// is not set.
	TrainerUri []string `json:"trainerUri,omitempty"`

	// WorkerSpec: Specification of workers.
	WorkerSpec *GoogleCloudMlV1alpha3__WorkerSpec `json:"workerSpec,omitempty"`

	// Zone: Optional. The Google Compute Engine zone to run the training
	// job in.
	// If not specified, the system will pick a zone based on the
	// job
	// requirements. It is the user's responsibility to procure quota
	// in this zone.
	Zone string `json:"zone,omitempty"`

	// ForceSendFields is a list of field names (e.g. "EvalDataPaths") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *GoogleCloudMlV1alpha3__SubmitTrainingJobRequest) MarshalJSON() ([]byte, error) {
	type noMethod GoogleCloudMlV1alpha3__SubmitTrainingJobRequest
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// GoogleCloudMlV1alpha3__TrainingJobResult: Represents the result of
// the longrunning.Operation created by the
// SubmitTrainingJob method.
type GoogleCloudMlV1alpha3__TrainingJobResult struct {
	// OutputPath: The URI of the output location provided at job creation
	// time.
	OutputPath string `json:"outputPath,omitempty"`

	// Runs: Results for individual Hyperparameter runs.
	Runs []*GoogleCloudMlV1alpha3__HyperparameterJobResult `json:"runs,omitempty"`

	// RunsCompleted: The number of tuning runs completed successfully.
	RunsCompleted uint64 `json:"runsCompleted,omitempty,string"`

	// StepsCompleted: The accomplished number of steps.
	StepsCompleted uint64 `json:"stepsCompleted,omitempty,string"`

	// ForceSendFields is a list of field names (e.g. "OutputPath") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *GoogleCloudMlV1alpha3__TrainingJobResult) MarshalJSON() ([]byte, error) {
	type noMethod GoogleCloudMlV1alpha3__TrainingJobResult
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// GoogleCloudMlV1alpha3__Version: Represents a version of the model.
type GoogleCloudMlV1alpha3__Version struct {
	// CreateTime: The creation time of the version.
	CreateTime string `json:"createTime,omitempty"`

	// IsDefault: Whether the version is default within the model.
	IsDefault bool `json:"isDefault,omitempty"`

	// LastUseTime: The last usage time of the version.
	LastUseTime string `json:"lastUseTime,omitempty"`

	// Name: The user-specified name of the model version.
	Name string `json:"name,omitempty"`

	// OriginUri: Google Cloud Storage location containing the model graph,
	// weights and
	// additional metadata at the moment when the version is created.
	OriginUri string `json:"originUri,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "CreateTime") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *GoogleCloudMlV1alpha3__Version) MarshalJSON() ([]byte, error) {
	type noMethod GoogleCloudMlV1alpha3__Version
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// GoogleCloudMlV1alpha3__WorkerSpec: Represents a specification of
// training workers.
type GoogleCloudMlV1alpha3__WorkerSpec struct {
	// ReplicaCount: The required number of workers. Defaults to 1.
	ReplicaCount uint64 `json:"replicaCount,omitempty,string"`

	// Resources: Resource requirements for a single worker. Optional.
	// If not specified, default worker configuration will be used.
	Resources *GoogleCloudMlV1alpha3__ResourceSpec `json:"resources,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ReplicaCount") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *GoogleCloudMlV1alpha3__WorkerSpec) MarshalJSON() ([]byte, error) {
	type noMethod GoogleCloudMlV1alpha3__WorkerSpec
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// GoogleCloudMlV1beta1__OperationMetadata: Represents the metadata of
// the longrunning.Operation.
type GoogleCloudMlV1beta1__OperationMetadata struct {
	// CreateTime: When the operation was submitted.
	CreateTime string `json:"createTime,omitempty"`

	// EndTime: When the operation processing was completed.
	EndTime string `json:"endTime,omitempty"`

	// IsCancellationRequested: Whether the cancellation of this operation
	// has been requested.
	IsCancellationRequested bool `json:"isCancellationRequested,omitempty"`

	// StartTime: When the operation processing was started.
	StartTime string `json:"startTime,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CreateTime") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *GoogleCloudMlV1beta1__OperationMetadata) MarshalJSON() ([]byte, error) {
	type noMethod GoogleCloudMlV1beta1__OperationMetadata
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// GoogleCloudMlV1beta1__Version: Represents a version of the model.
type GoogleCloudMlV1beta1__Version struct {
	// CreateTime: Output only. The creation time of the version.
	CreateTime string `json:"createTime,omitempty"`

	// DeploymentUri: Required. Google Cloud Storage object containing the
	// model graph, weights
	// and additional metadata at the moment when the version is created.
	DeploymentUri string `json:"deploymentUri,omitempty"`

	// Description: Optional. The description of the model version.
	Description string `json:"description,omitempty"`

	// IsDefault: Output only. Whether the version is default within the
	// model.
	IsDefault bool `json:"isDefault,omitempty"`

	// LastUseTime: Output only. The last usage time of the version.
	LastUseTime string `json:"lastUseTime,omitempty"`

	// Name: Required.The user-specified name of the model version.
	Name string `json:"name,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CreateTime") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *GoogleCloudMlV1beta1__Version) MarshalJSON() ([]byte, error) {
	type noMethod GoogleCloudMlV1beta1__Version
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// GoogleLongrunning__ListOperationsResponse: The response message for
// Operations.ListOperations.
type GoogleLongrunning__ListOperationsResponse struct {
	// NextPageToken: The standard List next-page token.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Operations: A list of operations that matches the specified filter in
	// the request.
	Operations []*GoogleLongrunning__Operation `json:"operations,omitempty"`

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

func (s *GoogleLongrunning__ListOperationsResponse) MarshalJSON() ([]byte, error) {
	type noMethod GoogleLongrunning__ListOperationsResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// GoogleLongrunning__Operation: This resource represents a long-running
// operation that is the result of a
// network API call.
type GoogleLongrunning__Operation struct {
	// Done: If the value is `false`, it means the operation is still in
	// progress.
	// If true, the operation is completed, and either `error` or `response`
	// is
	// available.
	Done bool `json:"done,omitempty"`

	// Error: The error result of the operation in case of failure.
	Error *GoogleRpc__Status `json:"error,omitempty"`

	// Metadata: Service-specific metadata associated with the operation.
	// It typically
	// contains progress information and common metadata such as create
	// time.
	// Some services might not provide such metadata.  Any method that
	// returns a
	// long-running operation should document the metadata type, if any.
	Metadata GoogleLongrunning__OperationMetadata `json:"metadata,omitempty"`

	// Name: The server-assigned name, which is only unique within the same
	// service that
	// originally returns it. If you use the default HTTP mapping,
	// the
	// `name` should have the format of `operations/some/unique/name`.
	Name string `json:"name,omitempty"`

	// Response: The normal response of the operation in case of success.
	// If the original
	// method returns no data on success, such as `Delete`, the response
	// is
	// `google.protobuf.Empty`.  If the original method is
	// standard
	// `Get`/`Create`/`Update`, the response should be the resource.  For
	// other
	// methods, the response should have the type `XxxResponse`, where
	// `Xxx`
	// is the original method name.  For example, if the original method
	// name
	// is `TakeSnapshot()`, the inferred response type
	// is
	// `TakeSnapshotResponse`.
	Response GoogleLongrunning__OperationResponse `json:"response,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Done") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *GoogleLongrunning__Operation) MarshalJSON() ([]byte, error) {
	type noMethod GoogleLongrunning__Operation
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

type GoogleLongrunning__OperationMetadata interface{}

type GoogleLongrunning__OperationResponse interface{}

// GoogleProtobuf__Empty: A generic empty message that you can re-use to
// avoid defining duplicated
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
type GoogleProtobuf__Empty struct {
	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`
}

// GoogleRpc__Status: The `Status` type defines a logical error model
// that is suitable for different
// programming environments, including REST APIs and RPC APIs. It is
// used by
// [gRPC](https://github.com/grpc). The error model is designed to
// be:
//
// - Simple to use and understand for most users
// - Flexible enough to meet unexpected needs
//
// # Overview
//
// The `Status` message contains three pieces of data: error code, error
// message,
// and error details. The error code should be an enum value
// of
// google.rpc.Code, but it may accept additional error codes if needed.
// The
// error message should be a developer-facing English message that
// helps
// developers *understand* and *resolve* the error. If a localized
// user-facing
// error message is needed, put the localized message in the error
// details or
// localize it in the client. The optional error details may contain
// arbitrary
// information about the error. There is a predefined set of error
// detail types
// in the package `google.rpc` which can be used for common error
// conditions.
//
// # Language mapping
//
// The `Status` message is the logical representation of the error
// model, but it
// is not necessarily the actual wire format. When the `Status` message
// is
// exposed in different client libraries and different wire protocols,
// it can be
// mapped differently. For example, it will likely be mapped to some
// exceptions
// in Java, but more likely mapped to some error codes in C.
//
// # Other uses
//
// The error model and the `Status` message can be used in a variety
// of
// environments, either with or without APIs, to provide a
// consistent developer experience across different
// environments.
//
// Example uses of this error model include:
//
// - Partial errors. If a service needs to return partial errors to the
// client,
//     it may embed the `Status` in the normal response to indicate the
// partial
//     errors.
//
// - Workflow errors. A typical workflow has multiple steps. Each step
// may
//     have a `Status` message for error reporting purpose.
//
// - Batch operations. If a client uses batch request and batch
// response, the
//     `Status` message should be used directly inside batch response,
// one for
//     each error sub-response.
//
// - Asynchronous operations. If an API call embeds asynchronous
// operation
//     results in its response, the status of those operations should
// be
//     represented directly using the `Status` message.
//
// - Logging. If some API errors are stored in logs, the message
// `Status` could
//     be used directly after any stripping needed for security/privacy
// reasons.
type GoogleRpc__Status struct {
	// Code: The status code, which should be an enum value of
	// google.rpc.Code.
	Code int64 `json:"code,omitempty"`

	// Details: A list of messages that carry the error details.  There will
	// be a
	// common set of message types for APIs to use.
	Details []GoogleRpc__StatusDetails `json:"details,omitempty"`

	// Message: A developer-facing error message, which should be in
	// English. Any
	// user-facing error message should be localized and sent in
	// the
	// google.rpc.Status.details field, or localized by the client.
	Message string `json:"message,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Code") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *GoogleRpc__Status) MarshalJSON() ([]byte, error) {
	type noMethod GoogleRpc__Status
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

type GoogleRpc__StatusDetails interface{}

// method id "ml.projects.getConfig":

type ProjectsGetConfigCall struct {
	s            *Service
	projectsId   string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// GetConfig: Get the service config associated with a given project.
func (r *ProjectsService) GetConfig(projectsId string) *ProjectsGetConfigCall {
	c := &ProjectsGetConfigCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectsId = projectsId
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsGetConfigCall) Fields(s ...googleapi.Field) *ProjectsGetConfigCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsGetConfigCall) IfNoneMatch(entityTag string) *ProjectsGetConfigCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsGetConfigCall) Context(ctx context.Context) *ProjectsGetConfigCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsGetConfigCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha3/projects/{projectsId}:getConfig")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectsId": c.projectsId,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "ml.projects.getConfig" call.
// Exactly one of *GoogleCloudMlV1alpha3__GetConfigResponse or error
// will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *GoogleCloudMlV1alpha3__GetConfigResponse.ServerResponse.Header or
// (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsGetConfigCall) Do(opts ...googleapi.CallOption) (*GoogleCloudMlV1alpha3__GetConfigResponse, error) {
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
	ret := &GoogleCloudMlV1alpha3__GetConfigResponse{
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
	//   "description": "Get the service config associated with a given project.",
	//   "flatPath": "v1alpha3/projects/{projectsId}:getConfig",
	//   "httpMethod": "GET",
	//   "id": "ml.projects.getConfig",
	//   "parameterOrder": [
	//     "projectsId"
	//   ],
	//   "parameters": {
	//     "projectsId": {
	//       "description": "Part of `name`. The project name.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1alpha3/projects/{projectsId}:getConfig",
	//   "response": {
	//     "$ref": "GoogleCloudMlV1alpha3__GetConfigResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "ml.projects.hyperparameters":

type ProjectsHyperparametersCall struct {
	s                                                *Service
	projectsId                                       string
	googlecloudmlv1alpha3__gethyperparametersrequest *GoogleCloudMlV1alpha3__GetHyperparametersRequest
	urlParams_                                       gensupport.URLParams
	ctx_                                             context.Context
}

// Hyperparameters: Get the hyperparameters assigned to the given run.
func (r *ProjectsService) Hyperparameters(projectsId string, googlecloudmlv1alpha3__gethyperparametersrequest *GoogleCloudMlV1alpha3__GetHyperparametersRequest) *ProjectsHyperparametersCall {
	c := &ProjectsHyperparametersCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectsId = projectsId
	c.googlecloudmlv1alpha3__gethyperparametersrequest = googlecloudmlv1alpha3__gethyperparametersrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsHyperparametersCall) Fields(s ...googleapi.Field) *ProjectsHyperparametersCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsHyperparametersCall) Context(ctx context.Context) *ProjectsHyperparametersCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsHyperparametersCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googlecloudmlv1alpha3__gethyperparametersrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha3/projects/{projectsId}/hyperparameters")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectsId": c.projectsId,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "ml.projects.hyperparameters" call.
// Exactly one of *GoogleCloudMlV1alpha3__GetHyperparametersResponse or
// error will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *GoogleCloudMlV1alpha3__GetHyperparametersResponse.ServerResponse.Head
// er or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsHyperparametersCall) Do(opts ...googleapi.CallOption) (*GoogleCloudMlV1alpha3__GetHyperparametersResponse, error) {
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
	ret := &GoogleCloudMlV1alpha3__GetHyperparametersResponse{
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
	//   "description": "Get the hyperparameters assigned to the given run.",
	//   "flatPath": "v1alpha3/projects/{projectsId}/hyperparameters",
	//   "httpMethod": "POST",
	//   "id": "ml.projects.hyperparameters",
	//   "parameterOrder": [
	//     "projectsId"
	//   ],
	//   "parameters": {
	//     "projectsId": {
	//       "description": "Part of `parent`. The project name.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1alpha3/projects/{projectsId}/hyperparameters",
	//   "request": {
	//     "$ref": "GoogleCloudMlV1alpha3__GetHyperparametersRequest"
	//   },
	//   "response": {
	//     "$ref": "GoogleCloudMlV1alpha3__GetHyperparametersResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "ml.projects.predict":

type ProjectsPredictCall struct {
	s                                     *Service
	projectsId                            string
	googlecloudmlv1alpha3__predictrequest *GoogleCloudMlV1alpha3__PredictRequest
	urlParams_                            gensupport.URLParams
	ctx_                                  context.Context
}

// Predict: Performs prediction on the data in the request.
func (r *ProjectsService) Predict(projectsId string, googlecloudmlv1alpha3__predictrequest *GoogleCloudMlV1alpha3__PredictRequest) *ProjectsPredictCall {
	c := &ProjectsPredictCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectsId = projectsId
	c.googlecloudmlv1alpha3__predictrequest = googlecloudmlv1alpha3__predictrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsPredictCall) Fields(s ...googleapi.Field) *ProjectsPredictCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsPredictCall) Context(ctx context.Context) *ProjectsPredictCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsPredictCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googlecloudmlv1alpha3__predictrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha3/projects/{projectsId}:predict")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectsId": c.projectsId,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "ml.projects.predict" call.
// Exactly one of *GoogleCloudMlV1alpha3__PredictResponse or error will
// be non-nil. Any non-2xx status code is an error. Response headers are
// in either
// *GoogleCloudMlV1alpha3__PredictResponse.ServerResponse.Header or (if
// a response was returned at all) in error.(*googleapi.Error).Header.
// Use googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsPredictCall) Do(opts ...googleapi.CallOption) (*GoogleCloudMlV1alpha3__PredictResponse, error) {
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
	ret := &GoogleCloudMlV1alpha3__PredictResponse{
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
	//   "description": "Performs prediction on the data in the request.",
	//   "flatPath": "v1alpha3/projects/{projectsId}:predict",
	//   "httpMethod": "POST",
	//   "id": "ml.projects.predict",
	//   "parameterOrder": [
	//     "projectsId"
	//   ],
	//   "parameters": {
	//     "projectsId": {
	//       "description": "Part of `parent`. The name of the model against which prediction should be performed. It is\nof the form: /projects/project_id/models/model_name.\nThe version information may be omitted, in which case prediction is\nperformed against the default version.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1alpha3/projects/{projectsId}:predict",
	//   "request": {
	//     "$ref": "GoogleCloudMlV1alpha3__PredictRequest"
	//   },
	//   "response": {
	//     "$ref": "GoogleCloudMlV1alpha3__PredictResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "ml.projects.reportMetric":

type ProjectsReportMetricCall struct {
	s                                          *Service
	projectsId                                 string
	googlecloudmlv1alpha3__reportmetricrequest *GoogleCloudMlV1alpha3__ReportMetricRequest
	urlParams_                                 gensupport.URLParams
	ctx_                                       context.Context
}

// ReportMetric: Report the progress of a Training Job
func (r *ProjectsService) ReportMetric(projectsId string, googlecloudmlv1alpha3__reportmetricrequest *GoogleCloudMlV1alpha3__ReportMetricRequest) *ProjectsReportMetricCall {
	c := &ProjectsReportMetricCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectsId = projectsId
	c.googlecloudmlv1alpha3__reportmetricrequest = googlecloudmlv1alpha3__reportmetricrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsReportMetricCall) Fields(s ...googleapi.Field) *ProjectsReportMetricCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsReportMetricCall) Context(ctx context.Context) *ProjectsReportMetricCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsReportMetricCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googlecloudmlv1alpha3__reportmetricrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha3/projects/{projectsId}:reportMetric")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectsId": c.projectsId,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "ml.projects.reportMetric" call.
// Exactly one of *GoogleProtobuf__Empty or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *GoogleProtobuf__Empty.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsReportMetricCall) Do(opts ...googleapi.CallOption) (*GoogleProtobuf__Empty, error) {
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
	ret := &GoogleProtobuf__Empty{
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
	//   "description": "Report the progress of a Training Job",
	//   "flatPath": "v1alpha3/projects/{projectsId}:reportMetric",
	//   "httpMethod": "POST",
	//   "id": "ml.projects.reportMetric",
	//   "parameterOrder": [
	//     "projectsId"
	//   ],
	//   "parameters": {
	//     "projectsId": {
	//       "description": "Part of `parent`. The project name.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1alpha3/projects/{projectsId}:reportMetric",
	//   "request": {
	//     "$ref": "GoogleCloudMlV1alpha3__ReportMetricRequest"
	//   },
	//   "response": {
	//     "$ref": "GoogleProtobuf__Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "ml.projects.submitPredictionJob":

type ProjectsSubmitPredictionJobCall struct {
	s                                                 *Service
	projectsId                                        string
	googlecloudmlv1alpha3__submitpredictionjobrequest *GoogleCloudMlV1alpha3__SubmitPredictionJobRequest
	urlParams_                                        gensupport.URLParams
	ctx_                                              context.Context
}

// SubmitPredictionJob: Performs batch prediction on the files specified
// in the request.
// JobMetadata and will contain PredictionJobResult when completed.
func (r *ProjectsService) SubmitPredictionJob(projectsId string, googlecloudmlv1alpha3__submitpredictionjobrequest *GoogleCloudMlV1alpha3__SubmitPredictionJobRequest) *ProjectsSubmitPredictionJobCall {
	c := &ProjectsSubmitPredictionJobCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectsId = projectsId
	c.googlecloudmlv1alpha3__submitpredictionjobrequest = googlecloudmlv1alpha3__submitpredictionjobrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsSubmitPredictionJobCall) Fields(s ...googleapi.Field) *ProjectsSubmitPredictionJobCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsSubmitPredictionJobCall) Context(ctx context.Context) *ProjectsSubmitPredictionJobCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsSubmitPredictionJobCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googlecloudmlv1alpha3__submitpredictionjobrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha3/projects/{projectsId}:submitPredictionJob")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectsId": c.projectsId,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "ml.projects.submitPredictionJob" call.
// Exactly one of *GoogleLongrunning__Operation or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *GoogleLongrunning__Operation.ServerResponse.Header or (if a
// response was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsSubmitPredictionJobCall) Do(opts ...googleapi.CallOption) (*GoogleLongrunning__Operation, error) {
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
	ret := &GoogleLongrunning__Operation{
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
	//   "description": "Performs batch prediction on the files specified in the request.\nJobMetadata and will contain PredictionJobResult when completed.",
	//   "flatPath": "v1alpha3/projects/{projectsId}:submitPredictionJob",
	//   "httpMethod": "POST",
	//   "id": "ml.projects.submitPredictionJob",
	//   "parameterOrder": [
	//     "projectsId"
	//   ],
	//   "parameters": {
	//     "projectsId": {
	//       "description": "Part of `parent`. The name of the model against which prediction should be performed. It is\nof the form: /projects/project_id/models/model_name/versions/version_id.\nThe version information may be omitted, in which case prediction is\nperformed against the default version.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1alpha3/projects/{projectsId}:submitPredictionJob",
	//   "request": {
	//     "$ref": "GoogleCloudMlV1alpha3__SubmitPredictionJobRequest"
	//   },
	//   "response": {
	//     "$ref": "GoogleLongrunning__Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "ml.projects.submitTrainingJob":

type ProjectsSubmitTrainingJobCall struct {
	s                                               *Service
	projectsId                                      string
	googlecloudmlv1alpha3__submittrainingjobrequest *GoogleCloudMlV1alpha3__SubmitTrainingJobRequest
	urlParams_                                      gensupport.URLParams
	ctx_                                            context.Context
}

// SubmitTrainingJob: Create a training job. The resulting operation
// contains
// JobMetadata and will contain JobResult when completed.
func (r *ProjectsService) SubmitTrainingJob(projectsId string, googlecloudmlv1alpha3__submittrainingjobrequest *GoogleCloudMlV1alpha3__SubmitTrainingJobRequest) *ProjectsSubmitTrainingJobCall {
	c := &ProjectsSubmitTrainingJobCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectsId = projectsId
	c.googlecloudmlv1alpha3__submittrainingjobrequest = googlecloudmlv1alpha3__submittrainingjobrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsSubmitTrainingJobCall) Fields(s ...googleapi.Field) *ProjectsSubmitTrainingJobCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsSubmitTrainingJobCall) Context(ctx context.Context) *ProjectsSubmitTrainingJobCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsSubmitTrainingJobCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googlecloudmlv1alpha3__submittrainingjobrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha3/projects/{projectsId}:submitTrainingJob")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectsId": c.projectsId,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "ml.projects.submitTrainingJob" call.
// Exactly one of *GoogleLongrunning__Operation or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *GoogleLongrunning__Operation.ServerResponse.Header or (if a
// response was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsSubmitTrainingJobCall) Do(opts ...googleapi.CallOption) (*GoogleLongrunning__Operation, error) {
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
	ret := &GoogleLongrunning__Operation{
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
	//   "description": "Create a training job. The resulting operation contains\nJobMetadata and will contain JobResult when completed.",
	//   "flatPath": "v1alpha3/projects/{projectsId}:submitTrainingJob",
	//   "httpMethod": "POST",
	//   "id": "ml.projects.submitTrainingJob",
	//   "parameterOrder": [
	//     "projectsId"
	//   ],
	//   "parameters": {
	//     "projectsId": {
	//       "description": "Part of `parent`. The project name.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1alpha3/projects/{projectsId}:submitTrainingJob",
	//   "request": {
	//     "$ref": "GoogleCloudMlV1alpha3__SubmitTrainingJobRequest"
	//   },
	//   "response": {
	//     "$ref": "GoogleLongrunning__Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "ml.projects.models.create":

type ProjectsModelsCreateCall struct {
	s                            *Service
	projectsId                   string
	googlecloudmlv1alpha3__model *GoogleCloudMlV1alpha3__Model
	urlParams_                   gensupport.URLParams
	ctx_                         context.Context
}

// Create: Create a model which will later contain a set of model
// versions.
func (r *ProjectsModelsService) Create(projectsId string, googlecloudmlv1alpha3__model *GoogleCloudMlV1alpha3__Model) *ProjectsModelsCreateCall {
	c := &ProjectsModelsCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectsId = projectsId
	c.googlecloudmlv1alpha3__model = googlecloudmlv1alpha3__model
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsModelsCreateCall) Fields(s ...googleapi.Field) *ProjectsModelsCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsModelsCreateCall) Context(ctx context.Context) *ProjectsModelsCreateCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsModelsCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googlecloudmlv1alpha3__model)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha3/projects/{projectsId}/models")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectsId": c.projectsId,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "ml.projects.models.create" call.
// Exactly one of *GoogleCloudMlV1alpha3__Model or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *GoogleCloudMlV1alpha3__Model.ServerResponse.Header or (if a
// response was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsModelsCreateCall) Do(opts ...googleapi.CallOption) (*GoogleCloudMlV1alpha3__Model, error) {
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
	ret := &GoogleCloudMlV1alpha3__Model{
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
	//   "description": "Create a model which will later contain a set of model versions.",
	//   "flatPath": "v1alpha3/projects/{projectsId}/models",
	//   "httpMethod": "POST",
	//   "id": "ml.projects.models.create",
	//   "parameterOrder": [
	//     "projectsId"
	//   ],
	//   "parameters": {
	//     "projectsId": {
	//       "description": "Part of `parent`. The project name.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1alpha3/projects/{projectsId}/models",
	//   "request": {
	//     "$ref": "GoogleCloudMlV1alpha3__Model"
	//   },
	//   "response": {
	//     "$ref": "GoogleCloudMlV1alpha3__Model"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "ml.projects.models.createVersion":

type ProjectsModelsCreateVersionCall struct {
	s                              *Service
	projectsId                     string
	modelsId                       string
	googlecloudmlv1alpha3__version *GoogleCloudMlV1alpha3__Version
	urlParams_                     gensupport.URLParams
	ctx_                           context.Context
}

// CreateVersion: Upload a trained TensorFlow model version. The result
// of the operation
// is a Version..
func (r *ProjectsModelsService) CreateVersion(projectsId string, modelsId string, googlecloudmlv1alpha3__version *GoogleCloudMlV1alpha3__Version) *ProjectsModelsCreateVersionCall {
	c := &ProjectsModelsCreateVersionCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectsId = projectsId
	c.modelsId = modelsId
	c.googlecloudmlv1alpha3__version = googlecloudmlv1alpha3__version
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsModelsCreateVersionCall) Fields(s ...googleapi.Field) *ProjectsModelsCreateVersionCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsModelsCreateVersionCall) Context(ctx context.Context) *ProjectsModelsCreateVersionCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsModelsCreateVersionCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googlecloudmlv1alpha3__version)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha3/projects/{projectsId}/models/{modelsId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectsId": c.projectsId,
		"modelsId":   c.modelsId,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "ml.projects.models.createVersion" call.
// Exactly one of *GoogleLongrunning__Operation or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *GoogleLongrunning__Operation.ServerResponse.Header or (if a
// response was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsModelsCreateVersionCall) Do(opts ...googleapi.CallOption) (*GoogleLongrunning__Operation, error) {
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
	ret := &GoogleLongrunning__Operation{
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
	//   "description": "Upload a trained TensorFlow model version. The result of the operation\nis a Version..",
	//   "flatPath": "v1alpha3/projects/{projectsId}/models/{modelsId}",
	//   "httpMethod": "POST",
	//   "id": "ml.projects.models.createVersion",
	//   "parameterOrder": [
	//     "projectsId",
	//     "modelsId"
	//   ],
	//   "parameters": {
	//     "modelsId": {
	//       "description": "Part of `parent`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `parent`. The name of the model.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1alpha3/projects/{projectsId}/models/{modelsId}",
	//   "request": {
	//     "$ref": "GoogleCloudMlV1alpha3__Version"
	//   },
	//   "response": {
	//     "$ref": "GoogleLongrunning__Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "ml.projects.models.delete":

type ProjectsModelsDeleteCall struct {
	s          *Service
	projectsId string
	modelsId   string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
}

// Delete: Delete the model and all versions in it.
func (r *ProjectsModelsService) Delete(projectsId string, modelsId string) *ProjectsModelsDeleteCall {
	c := &ProjectsModelsDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectsId = projectsId
	c.modelsId = modelsId
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsModelsDeleteCall) Fields(s ...googleapi.Field) *ProjectsModelsDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsModelsDeleteCall) Context(ctx context.Context) *ProjectsModelsDeleteCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsModelsDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha3/projects/{projectsId}/models/{modelsId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectsId": c.projectsId,
		"modelsId":   c.modelsId,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "ml.projects.models.delete" call.
// Exactly one of *GoogleProtobuf__Empty or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *GoogleProtobuf__Empty.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsModelsDeleteCall) Do(opts ...googleapi.CallOption) (*GoogleProtobuf__Empty, error) {
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
	ret := &GoogleProtobuf__Empty{
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
	//   "description": "Delete the model and all versions in it.",
	//   "flatPath": "v1alpha3/projects/{projectsId}/models/{modelsId}",
	//   "httpMethod": "DELETE",
	//   "id": "ml.projects.models.delete",
	//   "parameterOrder": [
	//     "projectsId",
	//     "modelsId"
	//   ],
	//   "parameters": {
	//     "modelsId": {
	//       "description": "Part of `name`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `name`. The name of the model.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1alpha3/projects/{projectsId}/models/{modelsId}",
	//   "response": {
	//     "$ref": "GoogleProtobuf__Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "ml.projects.models.get":

type ProjectsModelsGetCall struct {
	s            *Service
	projectsId   string
	modelsId     string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// Get: Describe a model and versions in it.
func (r *ProjectsModelsService) Get(projectsId string, modelsId string) *ProjectsModelsGetCall {
	c := &ProjectsModelsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectsId = projectsId
	c.modelsId = modelsId
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsModelsGetCall) Fields(s ...googleapi.Field) *ProjectsModelsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsModelsGetCall) IfNoneMatch(entityTag string) *ProjectsModelsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsModelsGetCall) Context(ctx context.Context) *ProjectsModelsGetCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsModelsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha3/projects/{projectsId}/models/{modelsId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectsId": c.projectsId,
		"modelsId":   c.modelsId,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "ml.projects.models.get" call.
// Exactly one of *GoogleCloudMlV1alpha3__Model or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *GoogleCloudMlV1alpha3__Model.ServerResponse.Header or (if a
// response was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsModelsGetCall) Do(opts ...googleapi.CallOption) (*GoogleCloudMlV1alpha3__Model, error) {
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
	ret := &GoogleCloudMlV1alpha3__Model{
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
	//   "description": "Describe a model and versions in it.",
	//   "flatPath": "v1alpha3/projects/{projectsId}/models/{modelsId}",
	//   "httpMethod": "GET",
	//   "id": "ml.projects.models.get",
	//   "parameterOrder": [
	//     "projectsId",
	//     "modelsId"
	//   ],
	//   "parameters": {
	//     "modelsId": {
	//       "description": "Part of `name`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `name`. The name of the model.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1alpha3/projects/{projectsId}/models/{modelsId}",
	//   "response": {
	//     "$ref": "GoogleCloudMlV1alpha3__Model"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "ml.projects.models.list":

type ProjectsModelsListCall struct {
	s            *Service
	projectsId   string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// List: List models in the project.
func (r *ProjectsModelsService) List(projectsId string) *ProjectsModelsListCall {
	c := &ProjectsModelsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectsId = projectsId
	return c
}

// Filter sets the optional parameter "filter": Specifies the subset of
// models to retrieve.
func (c *ProjectsModelsListCall) Filter(filter string) *ProjectsModelsListCall {
	c.urlParams_.Set("filter", filter)
	return c
}

// OrderBy sets the optional parameter "orderBy": Specifies the ordering
// of the models.
func (c *ProjectsModelsListCall) OrderBy(orderBy string) *ProjectsModelsListCall {
	c.urlParams_.Set("orderBy", orderBy)
	return c
}

// PageSize sets the optional parameter "pageSize": Optional page size.
// The default is 100.
func (c *ProjectsModelsListCall) PageSize(pageSize int64) *ProjectsModelsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": An optional
// pagination token, if available, for continuing the enumeration.
func (c *ProjectsModelsListCall) PageToken(pageToken string) *ProjectsModelsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsModelsListCall) Fields(s ...googleapi.Field) *ProjectsModelsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsModelsListCall) IfNoneMatch(entityTag string) *ProjectsModelsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsModelsListCall) Context(ctx context.Context) *ProjectsModelsListCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsModelsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha3/projects/{projectsId}/models")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectsId": c.projectsId,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "ml.projects.models.list" call.
// Exactly one of *GoogleCloudMlV1alpha3__ListModelsResponse or error
// will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *GoogleCloudMlV1alpha3__ListModelsResponse.ServerResponse.Header or
// (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsModelsListCall) Do(opts ...googleapi.CallOption) (*GoogleCloudMlV1alpha3__ListModelsResponse, error) {
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
	ret := &GoogleCloudMlV1alpha3__ListModelsResponse{
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
	//   "description": "List models in the project.",
	//   "flatPath": "v1alpha3/projects/{projectsId}/models",
	//   "httpMethod": "GET",
	//   "id": "ml.projects.models.list",
	//   "parameterOrder": [
	//     "projectsId"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "Specifies the subset of models to retrieve.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "orderBy": {
	//       "description": "Specifies the ordering of the models.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Optional page size. The default is 100.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "An optional pagination token, if available, for continuing the enumeration.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `parent`. The name of the project whose models are to be listed.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1alpha3/projects/{projectsId}/models",
	//   "response": {
	//     "$ref": "GoogleCloudMlV1alpha3__ListModelsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsModelsListCall) Pages(ctx context.Context, f func(*GoogleCloudMlV1alpha3__ListModelsResponse) error) error {
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

// method id "ml.projects.models.versions.delete":

type ProjectsModelsVersionsDeleteCall struct {
	s          *Service
	projectsId string
	modelsId   string
	versionsId string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
}

// Delete: Delete a version.
func (r *ProjectsModelsVersionsService) Delete(projectsId string, modelsId string, versionsId string) *ProjectsModelsVersionsDeleteCall {
	c := &ProjectsModelsVersionsDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectsId = projectsId
	c.modelsId = modelsId
	c.versionsId = versionsId
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsModelsVersionsDeleteCall) Fields(s ...googleapi.Field) *ProjectsModelsVersionsDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsModelsVersionsDeleteCall) Context(ctx context.Context) *ProjectsModelsVersionsDeleteCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsModelsVersionsDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha3/projects/{projectsId}/models/{modelsId}/versions/{versionsId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectsId": c.projectsId,
		"modelsId":   c.modelsId,
		"versionsId": c.versionsId,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "ml.projects.models.versions.delete" call.
// Exactly one of *GoogleProtobuf__Empty or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *GoogleProtobuf__Empty.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsModelsVersionsDeleteCall) Do(opts ...googleapi.CallOption) (*GoogleProtobuf__Empty, error) {
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
	ret := &GoogleProtobuf__Empty{
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
	//   "description": "Delete a version.",
	//   "flatPath": "v1alpha3/projects/{projectsId}/models/{modelsId}/versions/{versionsId}",
	//   "httpMethod": "DELETE",
	//   "id": "ml.projects.models.versions.delete",
	//   "parameterOrder": [
	//     "projectsId",
	//     "modelsId",
	//     "versionsId"
	//   ],
	//   "parameters": {
	//     "modelsId": {
	//       "description": "Part of `name`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `name`. The name of the version.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "versionsId": {
	//       "description": "Part of `name`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1alpha3/projects/{projectsId}/models/{modelsId}/versions/{versionsId}",
	//   "response": {
	//     "$ref": "GoogleProtobuf__Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "ml.projects.models.versions.get":

type ProjectsModelsVersionsGetCall struct {
	s            *Service
	projectsId   string
	modelsId     string
	versionsId   string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// Get: Get version metadata.
func (r *ProjectsModelsVersionsService) Get(projectsId string, modelsId string, versionsId string) *ProjectsModelsVersionsGetCall {
	c := &ProjectsModelsVersionsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectsId = projectsId
	c.modelsId = modelsId
	c.versionsId = versionsId
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsModelsVersionsGetCall) Fields(s ...googleapi.Field) *ProjectsModelsVersionsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsModelsVersionsGetCall) IfNoneMatch(entityTag string) *ProjectsModelsVersionsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsModelsVersionsGetCall) Context(ctx context.Context) *ProjectsModelsVersionsGetCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsModelsVersionsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha3/projects/{projectsId}/models/{modelsId}/versions/{versionsId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectsId": c.projectsId,
		"modelsId":   c.modelsId,
		"versionsId": c.versionsId,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "ml.projects.models.versions.get" call.
// Exactly one of *GoogleCloudMlV1alpha3__Version or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *GoogleCloudMlV1alpha3__Version.ServerResponse.Header or (if a
// response was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsModelsVersionsGetCall) Do(opts ...googleapi.CallOption) (*GoogleCloudMlV1alpha3__Version, error) {
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
	ret := &GoogleCloudMlV1alpha3__Version{
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
	//   "description": "Get version metadata.",
	//   "flatPath": "v1alpha3/projects/{projectsId}/models/{modelsId}/versions/{versionsId}",
	//   "httpMethod": "GET",
	//   "id": "ml.projects.models.versions.get",
	//   "parameterOrder": [
	//     "projectsId",
	//     "modelsId",
	//     "versionsId"
	//   ],
	//   "parameters": {
	//     "modelsId": {
	//       "description": "Part of `name`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `name`. The name of the version.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "versionsId": {
	//       "description": "Part of `name`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1alpha3/projects/{projectsId}/models/{modelsId}/versions/{versionsId}",
	//   "response": {
	//     "$ref": "GoogleCloudMlV1alpha3__Version"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "ml.projects.models.versions.list":

type ProjectsModelsVersionsListCall struct {
	s            *Service
	projectsId   string
	modelsId     string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// List: List versions in the model.
func (r *ProjectsModelsVersionsService) List(projectsId string, modelsId string) *ProjectsModelsVersionsListCall {
	c := &ProjectsModelsVersionsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectsId = projectsId
	c.modelsId = modelsId
	return c
}

// Filter sets the optional parameter "filter": Specifies the subset of
// versions to retrieve.
func (c *ProjectsModelsVersionsListCall) Filter(filter string) *ProjectsModelsVersionsListCall {
	c.urlParams_.Set("filter", filter)
	return c
}

// OrderBy sets the optional parameter "orderBy": Specifies the ordering
// of the versions.
func (c *ProjectsModelsVersionsListCall) OrderBy(orderBy string) *ProjectsModelsVersionsListCall {
	c.urlParams_.Set("orderBy", orderBy)
	return c
}

// PageSize sets the optional parameter "pageSize": Optional page size.
// The default is 100.
func (c *ProjectsModelsVersionsListCall) PageSize(pageSize int64) *ProjectsModelsVersionsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": An optional
// pagination token, if available, for continuing the enumeration.
func (c *ProjectsModelsVersionsListCall) PageToken(pageToken string) *ProjectsModelsVersionsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsModelsVersionsListCall) Fields(s ...googleapi.Field) *ProjectsModelsVersionsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsModelsVersionsListCall) IfNoneMatch(entityTag string) *ProjectsModelsVersionsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsModelsVersionsListCall) Context(ctx context.Context) *ProjectsModelsVersionsListCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsModelsVersionsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha3/projects/{projectsId}/models/{modelsId}/versions")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectsId": c.projectsId,
		"modelsId":   c.modelsId,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "ml.projects.models.versions.list" call.
// Exactly one of *GoogleCloudMlV1alpha3__ListVersionsResponse or error
// will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *GoogleCloudMlV1alpha3__ListVersionsResponse.ServerResponse.Header or
// (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsModelsVersionsListCall) Do(opts ...googleapi.CallOption) (*GoogleCloudMlV1alpha3__ListVersionsResponse, error) {
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
	ret := &GoogleCloudMlV1alpha3__ListVersionsResponse{
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
	//   "description": "List versions in the model.",
	//   "flatPath": "v1alpha3/projects/{projectsId}/models/{modelsId}/versions",
	//   "httpMethod": "GET",
	//   "id": "ml.projects.models.versions.list",
	//   "parameterOrder": [
	//     "projectsId",
	//     "modelsId"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "Specifies the subset of versions to retrieve.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "modelsId": {
	//       "description": "Part of `parent`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "orderBy": {
	//       "description": "Specifies the ordering of the versions.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Optional page size. The default is 100.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "An optional pagination token, if available, for continuing the enumeration.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `parent`. The name of the model whose versions are to be listed.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1alpha3/projects/{projectsId}/models/{modelsId}/versions",
	//   "response": {
	//     "$ref": "GoogleCloudMlV1alpha3__ListVersionsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsModelsVersionsListCall) Pages(ctx context.Context, f func(*GoogleCloudMlV1alpha3__ListVersionsResponse) error) error {
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

// method id "ml.projects.models.versions.setDefault":

type ProjectsModelsVersionsSetDefaultCall struct {
	s                                               *Service
	projectsId                                      string
	modelsId                                        string
	versionsId                                      string
	googlecloudmlv1alpha3__setdefaultversionrequest *GoogleCloudMlV1alpha3__SetDefaultVersionRequest
	urlParams_                                      gensupport.URLParams
	ctx_                                            context.Context
}

// SetDefault: Mark the version as default within the model.
func (r *ProjectsModelsVersionsService) SetDefault(projectsId string, modelsId string, versionsId string, googlecloudmlv1alpha3__setdefaultversionrequest *GoogleCloudMlV1alpha3__SetDefaultVersionRequest) *ProjectsModelsVersionsSetDefaultCall {
	c := &ProjectsModelsVersionsSetDefaultCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectsId = projectsId
	c.modelsId = modelsId
	c.versionsId = versionsId
	c.googlecloudmlv1alpha3__setdefaultversionrequest = googlecloudmlv1alpha3__setdefaultversionrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsModelsVersionsSetDefaultCall) Fields(s ...googleapi.Field) *ProjectsModelsVersionsSetDefaultCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsModelsVersionsSetDefaultCall) Context(ctx context.Context) *ProjectsModelsVersionsSetDefaultCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsModelsVersionsSetDefaultCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.googlecloudmlv1alpha3__setdefaultversionrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha3/projects/{projectsId}/models/{modelsId}/versions/{versionsId}:setDefault")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectsId": c.projectsId,
		"modelsId":   c.modelsId,
		"versionsId": c.versionsId,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "ml.projects.models.versions.setDefault" call.
// Exactly one of *GoogleCloudMlV1alpha3__Version or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *GoogleCloudMlV1alpha3__Version.ServerResponse.Header or (if a
// response was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsModelsVersionsSetDefaultCall) Do(opts ...googleapi.CallOption) (*GoogleCloudMlV1alpha3__Version, error) {
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
	ret := &GoogleCloudMlV1alpha3__Version{
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
	//   "description": "Mark the version as default within the model.",
	//   "flatPath": "v1alpha3/projects/{projectsId}/models/{modelsId}/versions/{versionsId}:setDefault",
	//   "httpMethod": "POST",
	//   "id": "ml.projects.models.versions.setDefault",
	//   "parameterOrder": [
	//     "projectsId",
	//     "modelsId",
	//     "versionsId"
	//   ],
	//   "parameters": {
	//     "modelsId": {
	//       "description": "Part of `name`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `name`. The version name which is being made default within the model.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "versionsId": {
	//       "description": "Part of `name`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1alpha3/projects/{projectsId}/models/{modelsId}/versions/{versionsId}:setDefault",
	//   "request": {
	//     "$ref": "GoogleCloudMlV1alpha3__SetDefaultVersionRequest"
	//   },
	//   "response": {
	//     "$ref": "GoogleCloudMlV1alpha3__Version"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "ml.projects.operations.cancel":

type ProjectsOperationsCancelCall struct {
	s            *Service
	projectsId   string
	operationsId string
	urlParams_   gensupport.URLParams
	ctx_         context.Context
}

// Cancel: Starts asynchronous cancellation on a long-running operation.
//  The server
// makes a best effort to cancel the operation, but success is
// not
// guaranteed.  If the server doesn't support this method, it
// returns
// `google.rpc.Code.UNIMPLEMENTED`.  Clients can
// use
// Operations.GetOperation or
// other methods to check whether the cancellation succeeded or whether
// the
// operation completed despite cancellation.
func (r *ProjectsOperationsService) Cancel(projectsId string, operationsId string) *ProjectsOperationsCancelCall {
	c := &ProjectsOperationsCancelCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectsId = projectsId
	c.operationsId = operationsId
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsOperationsCancelCall) Fields(s ...googleapi.Field) *ProjectsOperationsCancelCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsOperationsCancelCall) Context(ctx context.Context) *ProjectsOperationsCancelCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsOperationsCancelCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha3/projects/{projectsId}/operations/{operationsId}:cancel")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectsId":   c.projectsId,
		"operationsId": c.operationsId,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "ml.projects.operations.cancel" call.
// Exactly one of *GoogleProtobuf__Empty or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *GoogleProtobuf__Empty.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsOperationsCancelCall) Do(opts ...googleapi.CallOption) (*GoogleProtobuf__Empty, error) {
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
	ret := &GoogleProtobuf__Empty{
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
	//   "description": "Starts asynchronous cancellation on a long-running operation.  The server\nmakes a best effort to cancel the operation, but success is not\nguaranteed.  If the server doesn't support this method, it returns\n`google.rpc.Code.UNIMPLEMENTED`.  Clients can use\nOperations.GetOperation or\nother methods to check whether the cancellation succeeded or whether the\noperation completed despite cancellation.",
	//   "flatPath": "v1alpha3/projects/{projectsId}/operations/{operationsId}:cancel",
	//   "httpMethod": "POST",
	//   "id": "ml.projects.operations.cancel",
	//   "parameterOrder": [
	//     "projectsId",
	//     "operationsId"
	//   ],
	//   "parameters": {
	//     "operationsId": {
	//       "description": "Part of `name`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `name`. The name of the operation resource to be cancelled.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1alpha3/projects/{projectsId}/operations/{operationsId}:cancel",
	//   "response": {
	//     "$ref": "GoogleProtobuf__Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "ml.projects.operations.delete":

type ProjectsOperationsDeleteCall struct {
	s            *Service
	projectsId   string
	operationsId string
	urlParams_   gensupport.URLParams
	ctx_         context.Context
}

// Delete: Deletes a long-running operation. This method indicates that
// the client is
// no longer interested in the operation result. It does not cancel
// the
// operation. If the server doesn't support this method, it
// returns
// `google.rpc.Code.UNIMPLEMENTED`.
func (r *ProjectsOperationsService) Delete(projectsId string, operationsId string) *ProjectsOperationsDeleteCall {
	c := &ProjectsOperationsDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectsId = projectsId
	c.operationsId = operationsId
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsOperationsDeleteCall) Fields(s ...googleapi.Field) *ProjectsOperationsDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsOperationsDeleteCall) Context(ctx context.Context) *ProjectsOperationsDeleteCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsOperationsDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha3/projects/{projectsId}/operations/{operationsId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectsId":   c.projectsId,
		"operationsId": c.operationsId,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "ml.projects.operations.delete" call.
// Exactly one of *GoogleProtobuf__Empty or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *GoogleProtobuf__Empty.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsOperationsDeleteCall) Do(opts ...googleapi.CallOption) (*GoogleProtobuf__Empty, error) {
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
	ret := &GoogleProtobuf__Empty{
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
	//   "description": "Deletes a long-running operation. This method indicates that the client is\nno longer interested in the operation result. It does not cancel the\noperation. If the server doesn't support this method, it returns\n`google.rpc.Code.UNIMPLEMENTED`.",
	//   "flatPath": "v1alpha3/projects/{projectsId}/operations/{operationsId}",
	//   "httpMethod": "DELETE",
	//   "id": "ml.projects.operations.delete",
	//   "parameterOrder": [
	//     "projectsId",
	//     "operationsId"
	//   ],
	//   "parameters": {
	//     "operationsId": {
	//       "description": "Part of `name`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `name`. The name of the operation resource to be deleted.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1alpha3/projects/{projectsId}/operations/{operationsId}",
	//   "response": {
	//     "$ref": "GoogleProtobuf__Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "ml.projects.operations.get":

type ProjectsOperationsGetCall struct {
	s            *Service
	projectsId   string
	operationsId string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// Get: Gets the latest state of a long-running operation.  Clients can
// use this
// method to poll the operation result at intervals as recommended by
// the API
// service.
func (r *ProjectsOperationsService) Get(projectsId string, operationsId string) *ProjectsOperationsGetCall {
	c := &ProjectsOperationsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectsId = projectsId
	c.operationsId = operationsId
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsOperationsGetCall) Fields(s ...googleapi.Field) *ProjectsOperationsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsOperationsGetCall) IfNoneMatch(entityTag string) *ProjectsOperationsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsOperationsGetCall) Context(ctx context.Context) *ProjectsOperationsGetCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsOperationsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha3/projects/{projectsId}/operations/{operationsId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectsId":   c.projectsId,
		"operationsId": c.operationsId,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "ml.projects.operations.get" call.
// Exactly one of *GoogleLongrunning__Operation or error will be
// non-nil. Any non-2xx status code is an error. Response headers are in
// either *GoogleLongrunning__Operation.ServerResponse.Header or (if a
// response was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ProjectsOperationsGetCall) Do(opts ...googleapi.CallOption) (*GoogleLongrunning__Operation, error) {
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
	ret := &GoogleLongrunning__Operation{
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
	//   "description": "Gets the latest state of a long-running operation.  Clients can use this\nmethod to poll the operation result at intervals as recommended by the API\nservice.",
	//   "flatPath": "v1alpha3/projects/{projectsId}/operations/{operationsId}",
	//   "httpMethod": "GET",
	//   "id": "ml.projects.operations.get",
	//   "parameterOrder": [
	//     "projectsId",
	//     "operationsId"
	//   ],
	//   "parameters": {
	//     "operationsId": {
	//       "description": "Part of `name`. See documentation of `projectsId`.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "projectsId": {
	//       "description": "Part of `name`. The name of the operation resource.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1alpha3/projects/{projectsId}/operations/{operationsId}",
	//   "response": {
	//     "$ref": "GoogleLongrunning__Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "ml.projects.operations.list":

type ProjectsOperationsListCall struct {
	s            *Service
	projectsId   string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// List: Lists operations that match the specified filter in the
// request. If the
// server doesn't support this method, it returns
// `UNIMPLEMENTED`.
//
// NOTE: the `name` binding below allows API services to override the
// binding
// to use different resource name schemes, such as `users/*/operations`.
func (r *ProjectsOperationsService) List(projectsId string) *ProjectsOperationsListCall {
	c := &ProjectsOperationsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.projectsId = projectsId
	return c
}

// Filter sets the optional parameter "filter": The standard list
// filter.
func (c *ProjectsOperationsListCall) Filter(filter string) *ProjectsOperationsListCall {
	c.urlParams_.Set("filter", filter)
	return c
}

// PageSize sets the optional parameter "pageSize": The standard list
// page size.
func (c *ProjectsOperationsListCall) PageSize(pageSize int64) *ProjectsOperationsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": The standard list
// page token.
func (c *ProjectsOperationsListCall) PageToken(pageToken string) *ProjectsOperationsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsOperationsListCall) Fields(s ...googleapi.Field) *ProjectsOperationsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ProjectsOperationsListCall) IfNoneMatch(entityTag string) *ProjectsOperationsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ProjectsOperationsListCall) Context(ctx context.Context) *ProjectsOperationsListCall {
	c.ctx_ = ctx
	return c
}

func (c *ProjectsOperationsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1alpha3/projects/{projectsId}/operations")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"projectsId": c.projectsId,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "ml.projects.operations.list" call.
// Exactly one of *GoogleLongrunning__ListOperationsResponse or error
// will be non-nil. Any non-2xx status code is an error. Response
// headers are in either
// *GoogleLongrunning__ListOperationsResponse.ServerResponse.Header or
// (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *ProjectsOperationsListCall) Do(opts ...googleapi.CallOption) (*GoogleLongrunning__ListOperationsResponse, error) {
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
	ret := &GoogleLongrunning__ListOperationsResponse{
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
	//   "description": "Lists operations that match the specified filter in the request. If the\nserver doesn't support this method, it returns `UNIMPLEMENTED`.\n\nNOTE: the `name` binding below allows API services to override the binding\nto use different resource name schemes, such as `users/*/operations`.",
	//   "flatPath": "v1alpha3/projects/{projectsId}/operations",
	//   "httpMethod": "GET",
	//   "id": "ml.projects.operations.list",
	//   "parameterOrder": [
	//     "projectsId"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "The standard list filter.",
	//       "location": "query",
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
	//     },
	//     "projectsId": {
	//       "description": "Part of `name`. The name of the operation collection.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1alpha3/projects/{projectsId}/operations",
	//   "response": {
	//     "$ref": "GoogleLongrunning__ListOperationsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ProjectsOperationsListCall) Pages(ctx context.Context, f func(*GoogleLongrunning__ListOperationsResponse) error) error {
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
