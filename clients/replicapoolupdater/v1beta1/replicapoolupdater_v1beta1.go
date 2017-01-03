// Package replicapoolupdater provides access to the Google Compute Engine Instance Group Updater API.
//
// See https://cloud.google.com/compute/docs/instance-groups/manager/#applying_rolling_updates_using_the_updater_service
//
// Usage example:
//
//   import "github.com/skelterjohn/gcloud_apis/clients/replicapoolupdater/v1beta1"
//   ...
//   replicapoolupdaterService, err := replicapoolupdater.New(oauthHttpClient)
package replicapoolupdater

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

const apiId = "replicapoolupdater:v1beta1"
const apiName = "replicapoolupdater"
const apiVersion = "v1beta1"
const basePath = "https://www.googleapis.com/replicapoolupdater/v1beta1/projects/"

// OAuth2 scopes used by this API.
const (
	// View and manage your data across Google Cloud Platform services
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"

	// View your data across Google Cloud Platform services
	CloudPlatformReadOnlyScope = "https://www.googleapis.com/auth/cloud-platform.read-only"

	// View and manage replica pools
	ReplicapoolScope = "https://www.googleapis.com/auth/replicapool"

	// View replica pools
	ReplicapoolReadonlyScope = "https://www.googleapis.com/auth/replicapool.readonly"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.RollingUpdates = NewRollingUpdatesService(s)
	s.Rollout = NewRolloutService(s)
	s.ZoneOperations = NewZoneOperationsService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	RollingUpdates *RollingUpdatesService

	Rollout *RolloutService

	ZoneOperations *ZoneOperationsService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewRollingUpdatesService(s *Service) *RollingUpdatesService {
	rs := &RollingUpdatesService{s: s}
	return rs
}

type RollingUpdatesService struct {
	s *Service
}

func NewRolloutService(s *Service) *RolloutService {
	rs := &RolloutService{s: s}
	return rs
}

type RolloutService struct {
	s *Service
}

func NewZoneOperationsService(s *Service) *ZoneOperationsService {
	rs := &ZoneOperationsService{s: s}
	return rs
}

type ZoneOperationsService struct {
	s *Service
}

// FixedOrPercent: Used to specify an amount of instances within an
// instance group. Only one of fixed and percentage can be specified.
type FixedOrPercent struct {
	// Fixed: Specify a fixed amount of instances
	Fixed int64 `json:"fixed,omitempty"`

	// Percent: Specify a percentage of the total amount of instances within
	// an instance group manager
	Percent int64 `json:"percent,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Fixed") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *FixedOrPercent) MarshalJSON() ([]byte, error) {
	type noMethod FixedOrPercent
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// InstanceUpdate: Update of a single instance.
type InstanceUpdate struct {
	// Error: Errors that occurred during the instance update.
	Error *InstanceUpdateError `json:"error,omitempty"`

	// Instance: Fully-qualified URL of the instance being updated.
	Instance string `json:"instance,omitempty"`

	// Status: Status of the instance update. Possible values are:
	// - "PENDING": The instance update is pending execution.
	// - "ROLLING_FORWARD": The instance update is going forward.
	// - "ROLLING_BACK": The instance update is being rolled back.
	// - "PAUSED": The instance update is temporarily paused (inactive).
	// - "ROLLED_OUT": The instance update is finished, the instance is
	// running the new template.
	// - "ROLLED_BACK": The instance update is finished, the instance has
	// been reverted to the previous template.
	// - "CANCELLED": The instance update is paused and no longer can be
	// resumed, undefined in which template the instance is running.
	Status string `json:"status,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Error") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *InstanceUpdate) MarshalJSON() ([]byte, error) {
	type noMethod InstanceUpdate
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// InstanceUpdateError: Errors that occurred during the instance update.
type InstanceUpdateError struct {
	// Errors: [Output Only] The array of errors encountered while
	// processing this operation.
	Errors []*InstanceUpdateErrorErrors `json:"errors,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Errors") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *InstanceUpdateError) MarshalJSON() ([]byte, error) {
	type noMethod InstanceUpdateError
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

type InstanceUpdateErrorErrors struct {
	// Code: [Output Only] The error type identifier for this error.
	Code string `json:"code,omitempty"`

	// Location: [Output Only] Indicates the field in the request that
	// caused the error. This property is optional.
	Location string `json:"location,omitempty"`

	// Message: [Output Only] An optional, human-readable error message.
	Message string `json:"message,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Code") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *InstanceUpdateErrorErrors) MarshalJSON() ([]byte, error) {
	type noMethod InstanceUpdateErrorErrors
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// InstanceUpdateList: Response returned by ListInstanceUpdates method.
type InstanceUpdateList struct {
	// Items: Collection of requested instance updates.
	Items []*InstanceUpdate `json:"items,omitempty"`

	// Kind: [Output Only] Type of the resource.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: A token used to continue a truncated list request.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// SelfLink: [Output Only] The fully qualified URL for the resource.
	SelfLink string `json:"selfLink,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Items") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *InstanceUpdateList) MarshalJSON() ([]byte, error) {
	type noMethod InstanceUpdateList
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

type ListRolloutResponse struct {
	Resources []*Rollout `json:"resources,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Resources") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *ListRolloutResponse) MarshalJSON() ([]byte, error) {
	type noMethod ListRolloutResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// Operation: An operation resource, used to manage asynchronous API
// requests.
type Operation struct {
	ClientOperationId string `json:"clientOperationId,omitempty"`

	// CreationTimestamp: [Output Only] Creation timestamp in RFC3339 text
	// format.
	CreationTimestamp string `json:"creationTimestamp,omitempty"`

	EndTime string `json:"endTime,omitempty"`

	// Error: [Output Only] If errors occurred during processing of this
	// operation, this field will be populated.
	Error *OperationError `json:"error,omitempty"`

	HttpErrorMessage string `json:"httpErrorMessage,omitempty"`

	HttpErrorStatusCode int64 `json:"httpErrorStatusCode,omitempty"`

	// Id: [Output Only] Unique identifier for the resource; defined by the
	// server.
	Id uint64 `json:"id,omitempty,string"`

	// InsertTime: [Output Only] The time that this operation was requested.
	// This is in RFC 3339 format.
	InsertTime string `json:"insertTime,omitempty"`

	// Kind: [Output Only] Type of the resource. Always
	// replicapoolupdater#operation for Operation resources.
	Kind string `json:"kind,omitempty"`

	// Name: [Output Only] Name of the resource.
	Name string `json:"name,omitempty"`

	OperationType string `json:"operationType,omitempty"`

	Progress int64 `json:"progress,omitempty"`

	// Region: [Output Only] URL of the region where the operation resides.
	Region string `json:"region,omitempty"`

	// SelfLink: [Output Only] The fully qualified URL for the resource.
	SelfLink string `json:"selfLink,omitempty"`

	// StartTime: [Output Only] The time that this operation was started by
	// the server. This is in RFC 3339 format.
	StartTime string `json:"startTime,omitempty"`

	// Status: [Output Only] Status of the operation. Can be one of the
	// following: "PENDING", "RUNNING", or "DONE".
	Status string `json:"status,omitempty"`

	// StatusMessage: [Output Only] An optional textual description of the
	// current status of the operation.
	StatusMessage string `json:"statusMessage,omitempty"`

	// TargetId: [Output Only] Unique target id which identifies a
	// particular incarnation of the target.
	TargetId uint64 `json:"targetId,omitempty,string"`

	// TargetLink: [Output Only] URL of the resource the operation is
	// mutating.
	TargetLink string `json:"targetLink,omitempty"`

	User string `json:"user,omitempty"`

	Warnings []*OperationWarnings `json:"warnings,omitempty"`

	// Zone: [Output Only] URL of the zone where the operation resides.
	Zone string `json:"zone,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "ClientOperationId")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Operation) MarshalJSON() ([]byte, error) {
	type noMethod Operation
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// OperationError: [Output Only] If errors occurred during processing of
// this operation, this field will be populated.
type OperationError struct {
	// Errors: [Output Only] The array of errors encountered while
	// processing this operation.
	Errors []*OperationErrorErrors `json:"errors,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Errors") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *OperationError) MarshalJSON() ([]byte, error) {
	type noMethod OperationError
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

type OperationErrorErrors struct {
	// Code: [Output Only] The error type identifier for this error.
	Code string `json:"code,omitempty"`

	// Location: [Output Only] Indicates the field in the request that
	// caused the error. This property is optional.
	Location string `json:"location,omitempty"`

	// Message: [Output Only] An optional, human-readable error message.
	Message string `json:"message,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Code") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *OperationErrorErrors) MarshalJSON() ([]byte, error) {
	type noMethod OperationErrorErrors
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

type OperationWarnings struct {
	// Code: [Output only] The warning type identifier for this warning.
	Code string `json:"code,omitempty"`

	// Data: [Output only] Metadata for this warning in key:value format.
	Data []*OperationWarningsData `json:"data,omitempty"`

	// Message: [Output only] Optional human-readable details for this
	// warning.
	Message string `json:"message,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Code") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *OperationWarnings) MarshalJSON() ([]byte, error) {
	type noMethod OperationWarnings
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

type OperationWarningsData struct {
	// Key: [Output Only] Metadata key for this warning.
	Key string `json:"key,omitempty"`

	// Value: [Output Only] Metadata value for this warning.
	Value string `json:"value,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Key") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *OperationWarningsData) MarshalJSON() ([]byte, error) {
	type noMethod OperationWarningsData
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// OperationList: Contains a list of Operation resources.
type OperationList struct {
	// Id: [Output Only] Unique identifier for the resource; defined by the
	// server.
	Id string `json:"id,omitempty"`

	// Items: [Output Only] The Operation resources.
	Items []*Operation `json:"items,omitempty"`

	// Kind: [Output Only] Type of resource. Always
	// replicapoolupdater#operationList for OperationList resources.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: [Output Only] A token used to continue a truncate.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// SelfLink: [Output Only] The fully qualified URL for the resource.
	SelfLink string `json:"selfLink,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Id") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *OperationList) MarshalJSON() ([]byte, error) {
	type noMethod OperationList
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

type RampUpRolloutRequest struct {
	// CanarySize: The new amount of instances in the IGM to update
	// instances to.
	CanarySize *FixedOrPercent `json:"canarySize,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CanarySize") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *RampUpRolloutRequest) MarshalJSON() ([]byte, error) {
	type noMethod RampUpRolloutRequest
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// RollingUpdate: The following represents a resource describing a
// single update (rollout) of a group of instances to the given
// template.
type RollingUpdate struct {
	// ActionType: Specifies the action to take for each instance within the
	// instance group. This can be RECREATE which will recreate each
	// instance and is only available for managed instance groups. It can
	// also be REBOOT which performs a soft reboot for each instance and is
	// only available for regular (non-managed) instance groups.
	ActionType string `json:"actionType,omitempty"`

	// CreationTimestamp: [Output Only] Creation timestamp in RFC3339 text
	// format.
	CreationTimestamp string `json:"creationTimestamp,omitempty"`

	// Description: An optional textual description of the resource;
	// provided by the client when the resource is created.
	Description string `json:"description,omitempty"`

	// Error: [Output Only] Errors that occurred during the rolling update.
	Error *RollingUpdateError `json:"error,omitempty"`

	// Id: [Output Only] Unique identifier for the resource; defined by the
	// server.
	Id string `json:"id,omitempty"`

	// InstanceGroup: Fully-qualified URL of an instance group being
	// updated. Exactly one of instanceGroupManager and instanceGroup must
	// be set.
	InstanceGroup string `json:"instanceGroup,omitempty"`

	// InstanceGroupManager: Fully-qualified URL of an instance group
	// manager being updated. Exactly one of instanceGroupManager and
	// instanceGroup must be set.
	InstanceGroupManager string `json:"instanceGroupManager,omitempty"`

	// InstanceTemplate: Fully-qualified URL of an instance template to
	// apply.
	InstanceTemplate string `json:"instanceTemplate,omitempty"`

	// Kind: [Output Only] Type of the resource.
	Kind string `json:"kind,omitempty"`

	// OldInstanceTemplate: Fully-qualified URL of the instance template
	// encountered while starting the update.
	OldInstanceTemplate string `json:"oldInstanceTemplate,omitempty"`

	// Policy: Parameters of the update process.
	Policy *RollingUpdatePolicy `json:"policy,omitempty"`

	// Progress: [Output Only] An optional progress indicator that ranges
	// from 0 to 100. There is no requirement that this be linear or support
	// any granularity of operations. This should not be used to guess at
	// when the update will be complete. This number should be monotonically
	// increasing as the update progresses.
	Progress int64 `json:"progress,omitempty"`

	// SelfLink: [Output Only] The fully qualified URL for the resource.
	SelfLink string `json:"selfLink,omitempty"`

	// Status: [Output Only] Status of the update. Possible values are:
	// - "ROLLING_FORWARD": The update is going forward.
	// - "ROLLING_BACK": The update is being rolled back.
	// - "PAUSED": The update is temporarily paused (inactive).
	// - "ROLLED_OUT": The update is finished, all instances have been
	// updated successfully.
	// - "ROLLED_BACK": The update is finished, all instances have been
	// reverted to the previous template.
	// - "CANCELLED": The update is paused and no longer can be resumed,
	// undefined how many instances are running in which template.
	Status string `json:"status,omitempty"`

	// StatusMessage: [Output Only] An optional textual description of the
	// current status of the update.
	StatusMessage string `json:"statusMessage,omitempty"`

	// User: [Output Only] User who requested the update, for example:
	// user@example.com.
	User string `json:"user,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "ActionType") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *RollingUpdate) MarshalJSON() ([]byte, error) {
	type noMethod RollingUpdate
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// RollingUpdateError: [Output Only] Errors that occurred during the
// rolling update.
type RollingUpdateError struct {
	// Errors: [Output Only] The array of errors encountered while
	// processing this operation.
	Errors []*RollingUpdateErrorErrors `json:"errors,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Errors") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *RollingUpdateError) MarshalJSON() ([]byte, error) {
	type noMethod RollingUpdateError
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

type RollingUpdateErrorErrors struct {
	// Code: [Output Only] The error type identifier for this error.
	Code string `json:"code,omitempty"`

	// Location: [Output Only] Indicates the field in the request that
	// caused the error. This property is optional.
	Location string `json:"location,omitempty"`

	// Message: [Output Only] An optional, human-readable error message.
	Message string `json:"message,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Code") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *RollingUpdateErrorErrors) MarshalJSON() ([]byte, error) {
	type noMethod RollingUpdateErrorErrors
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// RollingUpdatePolicy: Parameters of the update process.
type RollingUpdatePolicy struct {
	// AutoPauseAfterInstances: Number of instances to update before the
	// updater pauses the rolling update.
	AutoPauseAfterInstances int64 `json:"autoPauseAfterInstances,omitempty"`

	// InstanceStartupTimeoutSec: The maximum amount of time that the
	// updater waits for a HEALTHY state after all of the update steps are
	// complete. If the HEALTHY state is not received before the deadline,
	// the instance update is considered a failure.
	InstanceStartupTimeoutSec int64 `json:"instanceStartupTimeoutSec,omitempty"`

	// MaxNumConcurrentInstances: The maximum number of instances that can
	// be updated simultaneously. An instance update is considered complete
	// only after the instance is restarted and initialized.
	MaxNumConcurrentInstances int64 `json:"maxNumConcurrentInstances,omitempty"`

	// MaxNumFailedInstances: The maximum number of instance updates that
	// can fail before the group update is considered a failure. An instance
	// update is considered failed if any of its update actions (e.g. Stop
	// call on Instance resource in Rolling Reboot) failed with permanent
	// failure, or if the instance is in an UNHEALTHY state after it
	// finishes all of the update actions.
	MaxNumFailedInstances int64 `json:"maxNumFailedInstances,omitempty"`

	// MinInstanceUpdateTimeSec: The minimum amount of time that the updater
	// spends to update each instance. Update time is the time it takes to
	// complete all update actions (e.g. Stop call on Instance resource in
	// Rolling Reboot), reboot, and initialize. If the instance update
	// finishes early, the updater pauses for the remainder of the time
	// before it starts the next instance update.
	MinInstanceUpdateTimeSec int64 `json:"minInstanceUpdateTimeSec,omitempty"`

	// ForceSendFields is a list of field names (e.g.
	// "AutoPauseAfterInstances") to unconditionally include in API
	// requests. By default, fields with empty values are omitted from API
	// requests. However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *RollingUpdatePolicy) MarshalJSON() ([]byte, error) {
	type noMethod RollingUpdatePolicy
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// RollingUpdateList: Response returned by List method.
type RollingUpdateList struct {
	// Items: Collection of requested updates.
	Items []*RollingUpdate `json:"items,omitempty"`

	// Kind: [Output Only] Type of the resource.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: A token used to continue a truncated list request.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// SelfLink: [Output Only] The fully qualified URL for the resource.
	SelfLink string `json:"selfLink,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Items") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *RollingUpdateList) MarshalJSON() ([]byte, error) {
	type noMethod RollingUpdateList
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

type Rollout struct {
	// CanarySize: The amount of instances within the instance group manager
	// to be updated to the new instance template.
	CanarySize *FixedOrPercent `json:"canarySize,omitempty"`

	// CanarySizeStages: A list of the amount of instances within the
	// instance group manager to be updated to the new instance template.
	// Each target size is updated sequentially, in the order supplied.
	CanarySizeStages []*FixedOrPercent `json:"canarySizeStages,omitempty"`

	// CreationTimestamp: [Output Only] Creation timestamp in RFC3339 text
	// format.
	CreationTimestamp string `json:"creationTimestamp,omitempty"`

	// HealthCheckDeadlineSec: How long to wait for the health checks to
	// return positive before assuming the update has failed or succeeded.
	HealthCheckDeadlineSec int64 `json:"healthCheckDeadlineSec,omitempty,string"`

	// Id: [Output Only] Unique identifier for the resource; defined by the
	// server.
	Id string `json:"id,omitempty"`

	// InstanceGroupManager: Fully-qualified URL of an instance group
	// manager being updated.
	InstanceGroupManager string `json:"instanceGroupManager,omitempty"`

	// InstanceTemplate: Fully-qualified URL of an instance template to
	// apply.
	InstanceTemplate string `json:"instanceTemplate,omitempty"`

	// InstanceTemplateToRollback: Fully-qualified URL of the instance
	// template to rollback to if the rollout is cancelled.
	InstanceTemplateToRollback string `json:"instanceTemplateToRollback,omitempty"`

	// Kind: [Output Only] Type of the resource.
	Kind string `json:"kind,omitempty"`

	// ParentRollout: The parent rollout to inherit unspecified fields from.
	ParentRollout string `json:"parentRollout,omitempty"`

	// SelfLink: [Output Only] The fully qualified URL for the resource.
	SelfLink string `json:"selfLink,omitempty"`

	// State: [Output Only] The current state of the rollout.
	State string `json:"state,omitempty"`

	// UpdatePolicy: Parameters of the update process.
	UpdatePolicy *UpdatePolicy `json:"updatePolicy,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "CanarySize") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *Rollout) MarshalJSON() ([]byte, error) {
	type noMethod Rollout
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

type UpdatePolicy struct {
	// MaxSurge: Maximum number of instances that can be created above the
	// InstanceGroupManager.targetSize during the update process. By
	// default, a fixed value of 1 is used. Using maxSurge > 0 will cause
	// instance names to change during the update process. At least one of {
	// maxSurge, maxUnavailable } must be greater than 0.
	MaxSurge *FixedOrPercent `json:"maxSurge,omitempty"`

	// MaxUnavailable: Maximum number of instances that can be unavailable
	// during the update process. The instance is considered available if
	// all of the following conditions are satisfied: 1. instance's status
	// is RUNNING 2. instance's liveness health check result was observed to
	// be HEALTHY at least once By default, a fixed value of 1 is used. At
	// least one of { maxSurge, maxUnavailable } must be greater than 0.
	MaxUnavailable *FixedOrPercent `json:"maxUnavailable,omitempty"`

	// MinReadySec: Minimum number of seconds to wait for after a newly
	// created instance becomes available. This value must be from range [0,
	// 3600].
	MinReadySec int64 `json:"minReadySec,omitempty"`

	// MinimalAction: Minimal action to be taken on an instance. The order
	// of action types is: RESTART < REPLACE.
	MinimalAction string `json:"minimalAction,omitempty"`

	Type string `json:"type,omitempty"`

	// UpdateType: The type of update
	UpdateType string `json:"updateType,omitempty"`

	// ForceSendFields is a list of field names (e.g. "MaxSurge") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`
}

func (s *UpdatePolicy) MarshalJSON() ([]byte, error) {
	type noMethod UpdatePolicy
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields)
}

// method id "replicapoolupdater.rollingUpdates.cancel":

type RollingUpdatesCancelCall struct {
	s             *Service
	project       string
	zone          string
	rollingUpdate string
	urlParams_    gensupport.URLParams
	ctx_          context.Context
}

// Cancel: Cancels an update. The update must be PAUSED before it can be
// cancelled. This has no effect if the update is already CANCELLED.
// For details, see https://cloud.google.com/compute/docs/instance-groups/manager/#cancelrollingupdate
func (r *RollingUpdatesService) Cancel(project string, zone string, rollingUpdate string) *RollingUpdatesCancelCall {
	c := &RollingUpdatesCancelCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.project = project
	c.zone = zone
	c.rollingUpdate = rollingUpdate
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RollingUpdatesCancelCall) Fields(s ...googleapi.Field) *RollingUpdatesCancelCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *RollingUpdatesCancelCall) Context(ctx context.Context) *RollingUpdatesCancelCall {
	c.ctx_ = ctx
	return c
}

func (c *RollingUpdatesCancelCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/zones/{zone}/rollingUpdates/{rollingUpdate}/cancel")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"project":       c.project,
		"zone":          c.zone,
		"rollingUpdate": c.rollingUpdate,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "replicapoolupdater.rollingUpdates.cancel" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *RollingUpdatesCancelCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
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
	ret := &Operation{
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
	//   "description": "Cancels an update. The update must be PAUSED before it can be cancelled. This has no effect if the update is already CANCELLED.",
	//   "httpMethod": "POST",
	//   "id": "replicapoolupdater.rollingUpdates.cancel",
	//   "parameterOrder": [
	//     "project",
	//     "zone",
	//     "rollingUpdate"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "The Google Developers Console project name.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "rollingUpdate": {
	//       "description": "The name of the update.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The name of the zone in which the update's target resides.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/rollingUpdates/{rollingUpdate}/cancel",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/replicapool"
	//   ]
	// }

}

// method id "replicapoolupdater.rollingUpdates.get":

type RollingUpdatesGetCall struct {
	s             *Service
	project       string
	zone          string
	rollingUpdate string
	urlParams_    gensupport.URLParams
	ifNoneMatch_  string
	ctx_          context.Context
}

// Get: Returns information about an update.
// For details, see https://cloud.google.com/compute/docs/instance-groups/manager/#getlistrollingupdate
func (r *RollingUpdatesService) Get(project string, zone string, rollingUpdate string) *RollingUpdatesGetCall {
	c := &RollingUpdatesGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.project = project
	c.zone = zone
	c.rollingUpdate = rollingUpdate
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RollingUpdatesGetCall) Fields(s ...googleapi.Field) *RollingUpdatesGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *RollingUpdatesGetCall) IfNoneMatch(entityTag string) *RollingUpdatesGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *RollingUpdatesGetCall) Context(ctx context.Context) *RollingUpdatesGetCall {
	c.ctx_ = ctx
	return c
}

func (c *RollingUpdatesGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/zones/{zone}/rollingUpdates/{rollingUpdate}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"project":       c.project,
		"zone":          c.zone,
		"rollingUpdate": c.rollingUpdate,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "replicapoolupdater.rollingUpdates.get" call.
// Exactly one of *RollingUpdate or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *RollingUpdate.ServerResponse.Header or (if a response was returned
// at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *RollingUpdatesGetCall) Do(opts ...googleapi.CallOption) (*RollingUpdate, error) {
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
	ret := &RollingUpdate{
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
	//   "description": "Returns information about an update.",
	//   "httpMethod": "GET",
	//   "id": "replicapoolupdater.rollingUpdates.get",
	//   "parameterOrder": [
	//     "project",
	//     "zone",
	//     "rollingUpdate"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "The Google Developers Console project name.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "rollingUpdate": {
	//       "description": "The name of the update.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The name of the zone in which the update's target resides.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/rollingUpdates/{rollingUpdate}",
	//   "response": {
	//     "$ref": "RollingUpdate"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only",
	//     "https://www.googleapis.com/auth/replicapool",
	//     "https://www.googleapis.com/auth/replicapool.readonly"
	//   ]
	// }

}

// method id "replicapoolupdater.rollingUpdates.insert":

type RollingUpdatesInsertCall struct {
	s             *Service
	project       string
	zone          string
	rollingupdate *RollingUpdate
	urlParams_    gensupport.URLParams
	ctx_          context.Context
}

// Insert: Inserts and starts a new update.
// For details, see https://cloud.google.com/compute/docs/instance-groups/manager/#starting_an_update
func (r *RollingUpdatesService) Insert(project string, zone string, rollingupdate *RollingUpdate) *RollingUpdatesInsertCall {
	c := &RollingUpdatesInsertCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.project = project
	c.zone = zone
	c.rollingupdate = rollingupdate
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RollingUpdatesInsertCall) Fields(s ...googleapi.Field) *RollingUpdatesInsertCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *RollingUpdatesInsertCall) Context(ctx context.Context) *RollingUpdatesInsertCall {
	c.ctx_ = ctx
	return c
}

func (c *RollingUpdatesInsertCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.rollingupdate)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/zones/{zone}/rollingUpdates")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"project": c.project,
		"zone":    c.zone,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "replicapoolupdater.rollingUpdates.insert" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *RollingUpdatesInsertCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
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
	ret := &Operation{
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
	//   "description": "Inserts and starts a new update.",
	//   "httpMethod": "POST",
	//   "id": "replicapoolupdater.rollingUpdates.insert",
	//   "parameterOrder": [
	//     "project",
	//     "zone"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "The Google Developers Console project name.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The name of the zone in which the update's target resides.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/rollingUpdates",
	//   "request": {
	//     "$ref": "RollingUpdate"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/replicapool"
	//   ]
	// }

}

// method id "replicapoolupdater.rollingUpdates.list":

type RollingUpdatesListCall struct {
	s            *Service
	project      string
	zone         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// List: Lists recent updates for a given managed instance group, in
// reverse chronological order and paginated format.
// For details, see https://cloud.google.com/compute/docs/instance-groups/manager/#getlistrollingupdate
func (r *RollingUpdatesService) List(project string, zone string) *RollingUpdatesListCall {
	c := &RollingUpdatesListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.project = project
	c.zone = zone
	return c
}

// Filter sets the optional parameter "filter": Filter expression for
// filtering listed resources.
func (c *RollingUpdatesListCall) Filter(filter string) *RollingUpdatesListCall {
	c.urlParams_.Set("filter", filter)
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum count of
// results to be returned. Maximum value is 500 and default value is
// 500.
func (c *RollingUpdatesListCall) MaxResults(maxResults int64) *RollingUpdatesListCall {
	c.urlParams_.Set("maxResults", fmt.Sprint(maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Tag returned by a
// previous list request truncated by maxResults. Used to continue a
// previous list request.
func (c *RollingUpdatesListCall) PageToken(pageToken string) *RollingUpdatesListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RollingUpdatesListCall) Fields(s ...googleapi.Field) *RollingUpdatesListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *RollingUpdatesListCall) IfNoneMatch(entityTag string) *RollingUpdatesListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *RollingUpdatesListCall) Context(ctx context.Context) *RollingUpdatesListCall {
	c.ctx_ = ctx
	return c
}

func (c *RollingUpdatesListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/zones/{zone}/rollingUpdates")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"project": c.project,
		"zone":    c.zone,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "replicapoolupdater.rollingUpdates.list" call.
// Exactly one of *RollingUpdateList or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *RollingUpdateList.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *RollingUpdatesListCall) Do(opts ...googleapi.CallOption) (*RollingUpdateList, error) {
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
	ret := &RollingUpdateList{
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
	//   "description": "Lists recent updates for a given managed instance group, in reverse chronological order and paginated format.",
	//   "httpMethod": "GET",
	//   "id": "replicapoolupdater.rollingUpdates.list",
	//   "parameterOrder": [
	//     "project",
	//     "zone"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "Optional. Filter expression for filtering listed resources.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "500",
	//       "description": "Optional. Maximum count of results to be returned. Maximum value is 500 and default value is 500.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "500",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. Tag returned by a previous list request truncated by maxResults. Used to continue a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "The Google Developers Console project name.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The name of the zone in which the update's target resides.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/rollingUpdates",
	//   "response": {
	//     "$ref": "RollingUpdateList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only",
	//     "https://www.googleapis.com/auth/replicapool",
	//     "https://www.googleapis.com/auth/replicapool.readonly"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *RollingUpdatesListCall) Pages(ctx context.Context, f func(*RollingUpdateList) error) error {
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

// method id "replicapoolupdater.rollingUpdates.listInstanceUpdates":

type RollingUpdatesListInstanceUpdatesCall struct {
	s             *Service
	project       string
	zone          string
	rollingUpdate string
	urlParams_    gensupport.URLParams
	ifNoneMatch_  string
	ctx_          context.Context
}

// ListInstanceUpdates: Lists the current status for each instance
// within a given update.
// For details, see https://cloud.google.com/compute/docs/instance-groups/manager/#getlistrollingupdate
func (r *RollingUpdatesService) ListInstanceUpdates(project string, zone string, rollingUpdate string) *RollingUpdatesListInstanceUpdatesCall {
	c := &RollingUpdatesListInstanceUpdatesCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.project = project
	c.zone = zone
	c.rollingUpdate = rollingUpdate
	return c
}

// Filter sets the optional parameter "filter": Filter expression for
// filtering listed resources.
func (c *RollingUpdatesListInstanceUpdatesCall) Filter(filter string) *RollingUpdatesListInstanceUpdatesCall {
	c.urlParams_.Set("filter", filter)
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum count of
// results to be returned. Maximum value is 500 and default value is
// 500.
func (c *RollingUpdatesListInstanceUpdatesCall) MaxResults(maxResults int64) *RollingUpdatesListInstanceUpdatesCall {
	c.urlParams_.Set("maxResults", fmt.Sprint(maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Tag returned by a
// previous list request truncated by maxResults. Used to continue a
// previous list request.
func (c *RollingUpdatesListInstanceUpdatesCall) PageToken(pageToken string) *RollingUpdatesListInstanceUpdatesCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RollingUpdatesListInstanceUpdatesCall) Fields(s ...googleapi.Field) *RollingUpdatesListInstanceUpdatesCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *RollingUpdatesListInstanceUpdatesCall) IfNoneMatch(entityTag string) *RollingUpdatesListInstanceUpdatesCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *RollingUpdatesListInstanceUpdatesCall) Context(ctx context.Context) *RollingUpdatesListInstanceUpdatesCall {
	c.ctx_ = ctx
	return c
}

func (c *RollingUpdatesListInstanceUpdatesCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/zones/{zone}/rollingUpdates/{rollingUpdate}/instanceUpdates")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"project":       c.project,
		"zone":          c.zone,
		"rollingUpdate": c.rollingUpdate,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "replicapoolupdater.rollingUpdates.listInstanceUpdates" call.
// Exactly one of *InstanceUpdateList or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *InstanceUpdateList.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *RollingUpdatesListInstanceUpdatesCall) Do(opts ...googleapi.CallOption) (*InstanceUpdateList, error) {
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
	ret := &InstanceUpdateList{
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
	//   "description": "Lists the current status for each instance within a given update.",
	//   "httpMethod": "GET",
	//   "id": "replicapoolupdater.rollingUpdates.listInstanceUpdates",
	//   "parameterOrder": [
	//     "project",
	//     "zone",
	//     "rollingUpdate"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "Optional. Filter expression for filtering listed resources.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "500",
	//       "description": "Optional. Maximum count of results to be returned. Maximum value is 500 and default value is 500.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "500",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. Tag returned by a previous list request truncated by maxResults. Used to continue a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "The Google Developers Console project name.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "rollingUpdate": {
	//       "description": "The name of the update.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The name of the zone in which the update's target resides.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/rollingUpdates/{rollingUpdate}/instanceUpdates",
	//   "response": {
	//     "$ref": "InstanceUpdateList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only",
	//     "https://www.googleapis.com/auth/replicapool",
	//     "https://www.googleapis.com/auth/replicapool.readonly"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *RollingUpdatesListInstanceUpdatesCall) Pages(ctx context.Context, f func(*InstanceUpdateList) error) error {
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

// method id "replicapoolupdater.rollingUpdates.pause":

type RollingUpdatesPauseCall struct {
	s             *Service
	project       string
	zone          string
	rollingUpdate string
	urlParams_    gensupport.URLParams
	ctx_          context.Context
}

// Pause: Pauses the update in state from ROLLING_FORWARD or
// ROLLING_BACK. Has no effect if invoked when the state of the update
// is PAUSED.
// For details, see https://cloud.google.com/compute/docs/instance-groups/manager/#pausing_a_rolling_update
func (r *RollingUpdatesService) Pause(project string, zone string, rollingUpdate string) *RollingUpdatesPauseCall {
	c := &RollingUpdatesPauseCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.project = project
	c.zone = zone
	c.rollingUpdate = rollingUpdate
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RollingUpdatesPauseCall) Fields(s ...googleapi.Field) *RollingUpdatesPauseCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *RollingUpdatesPauseCall) Context(ctx context.Context) *RollingUpdatesPauseCall {
	c.ctx_ = ctx
	return c
}

func (c *RollingUpdatesPauseCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/zones/{zone}/rollingUpdates/{rollingUpdate}/pause")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"project":       c.project,
		"zone":          c.zone,
		"rollingUpdate": c.rollingUpdate,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "replicapoolupdater.rollingUpdates.pause" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *RollingUpdatesPauseCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
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
	ret := &Operation{
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
	//   "description": "Pauses the update in state from ROLLING_FORWARD or ROLLING_BACK. Has no effect if invoked when the state of the update is PAUSED.",
	//   "httpMethod": "POST",
	//   "id": "replicapoolupdater.rollingUpdates.pause",
	//   "parameterOrder": [
	//     "project",
	//     "zone",
	//     "rollingUpdate"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "The Google Developers Console project name.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "rollingUpdate": {
	//       "description": "The name of the update.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The name of the zone in which the update's target resides.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/rollingUpdates/{rollingUpdate}/pause",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/replicapool"
	//   ]
	// }

}

// method id "replicapoolupdater.rollingUpdates.resume":

type RollingUpdatesResumeCall struct {
	s             *Service
	project       string
	zone          string
	rollingUpdate string
	urlParams_    gensupport.URLParams
	ctx_          context.Context
}

// Resume: Continues an update in PAUSED state. Has no effect if invoked
// when the state of the update is ROLLED_OUT.
// For details, see https://cloud.google.com/compute/docs/instance-groups/manager/#continuerollingupdate
func (r *RollingUpdatesService) Resume(project string, zone string, rollingUpdate string) *RollingUpdatesResumeCall {
	c := &RollingUpdatesResumeCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.project = project
	c.zone = zone
	c.rollingUpdate = rollingUpdate
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RollingUpdatesResumeCall) Fields(s ...googleapi.Field) *RollingUpdatesResumeCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *RollingUpdatesResumeCall) Context(ctx context.Context) *RollingUpdatesResumeCall {
	c.ctx_ = ctx
	return c
}

func (c *RollingUpdatesResumeCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/zones/{zone}/rollingUpdates/{rollingUpdate}/resume")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"project":       c.project,
		"zone":          c.zone,
		"rollingUpdate": c.rollingUpdate,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "replicapoolupdater.rollingUpdates.resume" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *RollingUpdatesResumeCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
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
	ret := &Operation{
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
	//   "description": "Continues an update in PAUSED state. Has no effect if invoked when the state of the update is ROLLED_OUT.",
	//   "httpMethod": "POST",
	//   "id": "replicapoolupdater.rollingUpdates.resume",
	//   "parameterOrder": [
	//     "project",
	//     "zone",
	//     "rollingUpdate"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "The Google Developers Console project name.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "rollingUpdate": {
	//       "description": "The name of the update.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The name of the zone in which the update's target resides.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/rollingUpdates/{rollingUpdate}/resume",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/replicapool"
	//   ]
	// }

}

// method id "replicapoolupdater.rollingUpdates.rollback":

type RollingUpdatesRollbackCall struct {
	s             *Service
	project       string
	zone          string
	rollingUpdate string
	urlParams_    gensupport.URLParams
	ctx_          context.Context
}

// Rollback: Rolls back the update in state from ROLLING_FORWARD or
// PAUSED. Has no effect if invoked when the state of the update is
// ROLLED_BACK.
// For details, see https://cloud.google.com/compute/docs/instance-groups/manager/#rollingbackupdate
func (r *RollingUpdatesService) Rollback(project string, zone string, rollingUpdate string) *RollingUpdatesRollbackCall {
	c := &RollingUpdatesRollbackCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.project = project
	c.zone = zone
	c.rollingUpdate = rollingUpdate
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RollingUpdatesRollbackCall) Fields(s ...googleapi.Field) *RollingUpdatesRollbackCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *RollingUpdatesRollbackCall) Context(ctx context.Context) *RollingUpdatesRollbackCall {
	c.ctx_ = ctx
	return c
}

func (c *RollingUpdatesRollbackCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/zones/{zone}/rollingUpdates/{rollingUpdate}/rollback")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"project":       c.project,
		"zone":          c.zone,
		"rollingUpdate": c.rollingUpdate,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "replicapoolupdater.rollingUpdates.rollback" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *RollingUpdatesRollbackCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
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
	ret := &Operation{
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
	//   "description": "Rolls back the update in state from ROLLING_FORWARD or PAUSED. Has no effect if invoked when the state of the update is ROLLED_BACK.",
	//   "httpMethod": "POST",
	//   "id": "replicapoolupdater.rollingUpdates.rollback",
	//   "parameterOrder": [
	//     "project",
	//     "zone",
	//     "rollingUpdate"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "The Google Developers Console project name.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "rollingUpdate": {
	//       "description": "The name of the update.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The name of the zone in which the update's target resides.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/rollingUpdates/{rollingUpdate}/rollback",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/replicapool"
	//   ]
	// }

}

// method id "replicapoolupdater.rollout.abandon":

type RolloutAbandonCall struct {
	s          *Service
	project    string
	zone       string
	rollout    string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
}

// Abandon: Abandon a rollout, leaving the IGM in the state it is
// already configured. This allows you to apply a new rollout to the
// IGM.
func (r *RolloutService) Abandon(project string, zone string, rollout string) *RolloutAbandonCall {
	c := &RolloutAbandonCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.project = project
	c.zone = zone
	c.rollout = rollout
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RolloutAbandonCall) Fields(s ...googleapi.Field) *RolloutAbandonCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *RolloutAbandonCall) Context(ctx context.Context) *RolloutAbandonCall {
	c.ctx_ = ctx
	return c
}

func (c *RolloutAbandonCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/zones/{zone}/rollouts/{rollout}/abandon")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"project": c.project,
		"zone":    c.zone,
		"rollout": c.rollout,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "replicapoolupdater.rollout.abandon" call.
// Exactly one of *Rollout or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Rollout.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *RolloutAbandonCall) Do(opts ...googleapi.CallOption) (*Rollout, error) {
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
	ret := &Rollout{
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
	//   "description": "Abandon a rollout, leaving the IGM in the state it is already configured. This allows you to apply a new rollout to the IGM.",
	//   "httpMethod": "POST",
	//   "id": "replicapoolupdater.rollout.abandon",
	//   "parameterOrder": [
	//     "project",
	//     "zone",
	//     "rollout"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "The Google Developers Console project name.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "rollout": {
	//       "description": "The ID of the update.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The name of the zone in which the update's target resides.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/rollouts/{rollout}/abandon",
	//   "response": {
	//     "$ref": "Rollout"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/replicapool"
	//   ]
	// }

}

// method id "replicapoolupdater.rollout.commit":

type RolloutCommitCall struct {
	s          *Service
	project    string
	zone       string
	rollout    string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
}

// Commit: Commits a rollout, so that it is final and can not be rolled
// back
func (r *RolloutService) Commit(project string, zone string, rollout string) *RolloutCommitCall {
	c := &RolloutCommitCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.project = project
	c.zone = zone
	c.rollout = rollout
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RolloutCommitCall) Fields(s ...googleapi.Field) *RolloutCommitCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *RolloutCommitCall) Context(ctx context.Context) *RolloutCommitCall {
	c.ctx_ = ctx
	return c
}

func (c *RolloutCommitCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/zones/{zone}/rollouts/{rollout}/commit")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"project": c.project,
		"zone":    c.zone,
		"rollout": c.rollout,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "replicapoolupdater.rollout.commit" call.
// Exactly one of *Rollout or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Rollout.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *RolloutCommitCall) Do(opts ...googleapi.CallOption) (*Rollout, error) {
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
	ret := &Rollout{
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
	//   "description": "Commits a rollout, so that it is final and can not be rolled back",
	//   "httpMethod": "POST",
	//   "id": "replicapoolupdater.rollout.commit",
	//   "parameterOrder": [
	//     "project",
	//     "zone",
	//     "rollout"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "The Google Developers Console project name.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "rollout": {
	//       "description": "The ID of the update.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The name of the zone in which the update's target resides.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/rollouts/{rollout}/commit",
	//   "response": {
	//     "$ref": "Rollout"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/replicapool"
	//   ]
	// }

}

// method id "replicapoolupdater.rollout.get":

type RolloutGetCall struct {
	s            *Service
	project      string
	zone         string
	rollout      string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// Get: Get the details of a rollout
func (r *RolloutService) Get(project string, zone string, rollout string) *RolloutGetCall {
	c := &RolloutGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.project = project
	c.zone = zone
	c.rollout = rollout
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RolloutGetCall) Fields(s ...googleapi.Field) *RolloutGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *RolloutGetCall) IfNoneMatch(entityTag string) *RolloutGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *RolloutGetCall) Context(ctx context.Context) *RolloutGetCall {
	c.ctx_ = ctx
	return c
}

func (c *RolloutGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/zones/{zone}/rollouts/{rollout}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"project": c.project,
		"zone":    c.zone,
		"rollout": c.rollout,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "replicapoolupdater.rollout.get" call.
// Exactly one of *Rollout or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Rollout.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *RolloutGetCall) Do(opts ...googleapi.CallOption) (*Rollout, error) {
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
	ret := &Rollout{
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
	//   "description": "Get the details of a rollout",
	//   "httpMethod": "GET",
	//   "id": "replicapoolupdater.rollout.get",
	//   "parameterOrder": [
	//     "project",
	//     "zone",
	//     "rollout"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "The Google Developers Console project name.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "rollout": {
	//       "description": "The ID of the update.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The name of the zone in which the update's target resides.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/rollouts/{rollout}",
	//   "response": {
	//     "$ref": "Rollout"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only",
	//     "https://www.googleapis.com/auth/replicapool",
	//     "https://www.googleapis.com/auth/replicapool.readonly"
	//   ]
	// }

}

// method id "replicapoolupdater.rollout.insert":

type RolloutInsertCall struct {
	s          *Service
	project    string
	zone       string
	rollout    *Rollout
	urlParams_ gensupport.URLParams
	ctx_       context.Context
}

// Insert: Inserts and starts a new rollout
func (r *RolloutService) Insert(project string, zone string, rollout *Rollout) *RolloutInsertCall {
	c := &RolloutInsertCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.project = project
	c.zone = zone
	c.rollout = rollout
	return c
}

// UpdatePolicyInitialisationMethod sets the optional parameter
// "updatePolicyInitialisationMethod": How the update policy should be
// initialised.
//
// Possible values:
//   "FROM_IGM"
//   "FROM_PARENT"
func (c *RolloutInsertCall) UpdatePolicyInitialisationMethod(updatePolicyInitialisationMethod string) *RolloutInsertCall {
	c.urlParams_.Set("updatePolicyInitialisationMethod", updatePolicyInitialisationMethod)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RolloutInsertCall) Fields(s ...googleapi.Field) *RolloutInsertCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *RolloutInsertCall) Context(ctx context.Context) *RolloutInsertCall {
	c.ctx_ = ctx
	return c
}

func (c *RolloutInsertCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.rollout)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/zones/{zone}/rollouts")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"project": c.project,
		"zone":    c.zone,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "replicapoolupdater.rollout.insert" call.
// Exactly one of *Rollout or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Rollout.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *RolloutInsertCall) Do(opts ...googleapi.CallOption) (*Rollout, error) {
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
	ret := &Rollout{
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
	//   "description": "Inserts and starts a new rollout",
	//   "httpMethod": "POST",
	//   "id": "replicapoolupdater.rollout.insert",
	//   "parameterOrder": [
	//     "project",
	//     "zone"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "The Google Developers Console project name.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "updatePolicyInitialisationMethod": {
	//       "description": "How the update policy should be initialised.",
	//       "enum": [
	//         "FROM_IGM",
	//         "FROM_PARENT"
	//       ],
	//       "enumDescriptions": [
	//         "",
	//         ""
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The name of the zone in which the update's target resides.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/rollouts",
	//   "request": {
	//     "$ref": "Rollout"
	//   },
	//   "response": {
	//     "$ref": "Rollout"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/replicapool"
	//   ]
	// }

}

// method id "replicapoolupdater.rollout.list":

type RolloutListCall struct {
	s            *Service
	project      string
	zone         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// List: Resume a rollout. This lets the rollout continue updating
// instances after a pause.
func (r *RolloutService) List(project string, zone string) *RolloutListCall {
	c := &RolloutListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.project = project
	c.zone = zone
	return c
}

// Filter sets the optional parameter "filter": Filter expression for
// filtering listed resources.
func (c *RolloutListCall) Filter(filter string) *RolloutListCall {
	c.urlParams_.Set("filter", filter)
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum count of
// results to be returned. Maximum value is 500 and default value is
// 500.
func (c *RolloutListCall) MaxResults(maxResults int64) *RolloutListCall {
	c.urlParams_.Set("maxResults", fmt.Sprint(maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Tag returned by a
// previous list request truncated by maxResults. Used to continue a
// previous list request.
func (c *RolloutListCall) PageToken(pageToken string) *RolloutListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RolloutListCall) Fields(s ...googleapi.Field) *RolloutListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *RolloutListCall) IfNoneMatch(entityTag string) *RolloutListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *RolloutListCall) Context(ctx context.Context) *RolloutListCall {
	c.ctx_ = ctx
	return c
}

func (c *RolloutListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/zones/{zone}/rollouts")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"project": c.project,
		"zone":    c.zone,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "replicapoolupdater.rollout.list" call.
// Exactly one of *ListRolloutResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListRolloutResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *RolloutListCall) Do(opts ...googleapi.CallOption) (*ListRolloutResponse, error) {
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
	ret := &ListRolloutResponse{
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
	//   "description": "Resume a rollout. This lets the rollout continue updating instances after a pause.",
	//   "httpMethod": "GET",
	//   "id": "replicapoolupdater.rollout.list",
	//   "parameterOrder": [
	//     "project",
	//     "zone"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "Optional. Filter expression for filtering listed resources.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "500",
	//       "description": "Optional. Maximum count of results to be returned. Maximum value is 500 and default value is 500.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "500",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. Tag returned by a previous list request truncated by maxResults. Used to continue a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "The Google Developers Console project name.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The name of the zone in which the update's target resides.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/rollouts",
	//   "response": {
	//     "$ref": "ListRolloutResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only",
	//     "https://www.googleapis.com/auth/replicapool",
	//     "https://www.googleapis.com/auth/replicapool.readonly"
	//   ]
	// }

}

// method id "replicapoolupdater.rollout.pause":

type RolloutPauseCall struct {
	s          *Service
	project    string
	zone       string
	rollout    string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
}

// Pause: Pause the application of a rollout. This stops the update, and
// the instances managed by the instance group manager do not change
// their instance templates.
func (r *RolloutService) Pause(project string, zone string, rollout string) *RolloutPauseCall {
	c := &RolloutPauseCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.project = project
	c.zone = zone
	c.rollout = rollout
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RolloutPauseCall) Fields(s ...googleapi.Field) *RolloutPauseCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *RolloutPauseCall) Context(ctx context.Context) *RolloutPauseCall {
	c.ctx_ = ctx
	return c
}

func (c *RolloutPauseCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/zones/{zone}/rollouts/{rollout}/pause")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"project": c.project,
		"zone":    c.zone,
		"rollout": c.rollout,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "replicapoolupdater.rollout.pause" call.
// Exactly one of *Rollout or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Rollout.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *RolloutPauseCall) Do(opts ...googleapi.CallOption) (*Rollout, error) {
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
	ret := &Rollout{
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
	//   "description": "Pause the application of a rollout. This stops the update, and the instances managed by the instance group manager do not change their instance templates.",
	//   "httpMethod": "POST",
	//   "id": "replicapoolupdater.rollout.pause",
	//   "parameterOrder": [
	//     "project",
	//     "zone",
	//     "rollout"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "The Google Developers Console project name.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "rollout": {
	//       "description": "The ID of the update.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The name of the zone in which the update's target resides.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/rollouts/{rollout}/pause",
	//   "response": {
	//     "$ref": "Rollout"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/replicapool"
	//   ]
	// }

}

// method id "replicapoolupdater.rollout.rampup":

type RolloutRampupCall struct {
	s                    *Service
	project              string
	zone                 string
	rollout              string
	rampuprolloutrequest *RampUpRolloutRequest
	urlParams_           gensupport.URLParams
	ctx_                 context.Context
}

// Rampup: Change the amount of instances within an IGM that should be
// updated to the new instance template
func (r *RolloutService) Rampup(project string, zone string, rollout string, rampuprolloutrequest *RampUpRolloutRequest) *RolloutRampupCall {
	c := &RolloutRampupCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.project = project
	c.zone = zone
	c.rollout = rollout
	c.rampuprolloutrequest = rampuprolloutrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RolloutRampupCall) Fields(s ...googleapi.Field) *RolloutRampupCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *RolloutRampupCall) Context(ctx context.Context) *RolloutRampupCall {
	c.ctx_ = ctx
	return c
}

func (c *RolloutRampupCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.rampuprolloutrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/zones/{zone}/rollouts/{rollout}/rampup")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"project": c.project,
		"zone":    c.zone,
		"rollout": c.rollout,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "replicapoolupdater.rollout.rampup" call.
// Exactly one of *Rollout or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Rollout.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *RolloutRampupCall) Do(opts ...googleapi.CallOption) (*Rollout, error) {
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
	ret := &Rollout{
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
	//   "description": "Change the amount of instances within an IGM that should be updated to the new instance template",
	//   "httpMethod": "POST",
	//   "id": "replicapoolupdater.rollout.rampup",
	//   "parameterOrder": [
	//     "project",
	//     "zone",
	//     "rollout"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "The Google Developers Console project name.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "rollout": {
	//       "description": "The ID of the update.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The name of the zone in which the update's target resides.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/rollouts/{rollout}/rampup",
	//   "request": {
	//     "$ref": "RampUpRolloutRequest"
	//   },
	//   "response": {
	//     "$ref": "Rollout"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/replicapool"
	//   ]
	// }

}

// method id "replicapoolupdater.rollout.resume":

type RolloutResumeCall struct {
	s          *Service
	project    string
	zone       string
	rollout    string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
}

// Resume: Resume a rollout. This lets the rollout continue updating
// instances after a pause.
func (r *RolloutService) Resume(project string, zone string, rollout string) *RolloutResumeCall {
	c := &RolloutResumeCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.project = project
	c.zone = zone
	c.rollout = rollout
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RolloutResumeCall) Fields(s ...googleapi.Field) *RolloutResumeCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *RolloutResumeCall) Context(ctx context.Context) *RolloutResumeCall {
	c.ctx_ = ctx
	return c
}

func (c *RolloutResumeCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/zones/{zone}/rollouts/{rollout}/resume")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"project": c.project,
		"zone":    c.zone,
		"rollout": c.rollout,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "replicapoolupdater.rollout.resume" call.
// Exactly one of *Rollout or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Rollout.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *RolloutResumeCall) Do(opts ...googleapi.CallOption) (*Rollout, error) {
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
	ret := &Rollout{
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
	//   "description": "Resume a rollout. This lets the rollout continue updating instances after a pause.",
	//   "httpMethod": "POST",
	//   "id": "replicapoolupdater.rollout.resume",
	//   "parameterOrder": [
	//     "project",
	//     "zone",
	//     "rollout"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "The Google Developers Console project name.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "rollout": {
	//       "description": "The ID of the update.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The name of the zone in which the update's target resides.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/rollouts/{rollout}/resume",
	//   "response": {
	//     "$ref": "Rollout"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/replicapool"
	//   ]
	// }

}

// method id "replicapoolupdater.rollout.rollback":

type RolloutRollbackCall struct {
	s          *Service
	project    string
	zone       string
	rollout    string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
}

// Rollback: Rollback a rollout, cancelling the update and changing all
// instances with the updated version to have the
// instanceTemplateToRollback template.
func (r *RolloutService) Rollback(project string, zone string, rollout string) *RolloutRollbackCall {
	c := &RolloutRollbackCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.project = project
	c.zone = zone
	c.rollout = rollout
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *RolloutRollbackCall) Fields(s ...googleapi.Field) *RolloutRollbackCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *RolloutRollbackCall) Context(ctx context.Context) *RolloutRollbackCall {
	c.ctx_ = ctx
	return c
}

func (c *RolloutRollbackCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/zones/{zone}/rollouts/{rollout}/rollback")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"project": c.project,
		"zone":    c.zone,
		"rollout": c.rollout,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "replicapoolupdater.rollout.rollback" call.
// Exactly one of *Rollout or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Rollout.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *RolloutRollbackCall) Do(opts ...googleapi.CallOption) (*Rollout, error) {
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
	ret := &Rollout{
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
	//   "description": "Rollback a rollout, cancelling the update and changing all instances with the updated version to have the instanceTemplateToRollback template.",
	//   "httpMethod": "POST",
	//   "id": "replicapoolupdater.rollout.rollback",
	//   "parameterOrder": [
	//     "project",
	//     "zone",
	//     "rollout"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "The Google Developers Console project name.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "rollout": {
	//       "description": "The ID of the update.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "The name of the zone in which the update's target resides.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/rollouts/{rollout}/rollback",
	//   "response": {
	//     "$ref": "Rollout"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/replicapool"
	//   ]
	// }

}

// method id "replicapoolupdater.zoneOperations.get":

type ZoneOperationsGetCall struct {
	s            *Service
	project      string
	zone         string
	operation    string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// Get: Retrieves the specified zone-specific operation resource.
func (r *ZoneOperationsService) Get(project string, zone string, operation string) *ZoneOperationsGetCall {
	c := &ZoneOperationsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.project = project
	c.zone = zone
	c.operation = operation
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ZoneOperationsGetCall) Fields(s ...googleapi.Field) *ZoneOperationsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ZoneOperationsGetCall) IfNoneMatch(entityTag string) *ZoneOperationsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ZoneOperationsGetCall) Context(ctx context.Context) *ZoneOperationsGetCall {
	c.ctx_ = ctx
	return c
}

func (c *ZoneOperationsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/zones/{zone}/operations/{operation}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"project":   c.project,
		"zone":      c.zone,
		"operation": c.operation,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "replicapoolupdater.zoneOperations.get" call.
// Exactly one of *Operation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Operation.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ZoneOperationsGetCall) Do(opts ...googleapi.CallOption) (*Operation, error) {
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
	ret := &Operation{
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
	//   "description": "Retrieves the specified zone-specific operation resource.",
	//   "httpMethod": "GET",
	//   "id": "replicapoolupdater.zoneOperations.get",
	//   "parameterOrder": [
	//     "project",
	//     "zone",
	//     "operation"
	//   ],
	//   "parameters": {
	//     "operation": {
	//       "description": "Name of the operation resource to return.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Name of the project scoping this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "Name of the zone scoping this request.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/operations/{operation}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/replicapool"
	//   ]
	// }

}

// method id "replicapoolupdater.zoneOperations.list":

type ZoneOperationsListCall struct {
	s            *Service
	project      string
	zone         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
}

// List: Retrieves the list of Operation resources contained within the
// specified zone.
func (r *ZoneOperationsService) List(project string, zone string) *ZoneOperationsListCall {
	c := &ZoneOperationsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.project = project
	c.zone = zone
	return c
}

// Filter sets the optional parameter "filter": Filter expression for
// filtering listed resources.
func (c *ZoneOperationsListCall) Filter(filter string) *ZoneOperationsListCall {
	c.urlParams_.Set("filter", filter)
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum count of
// results to be returned. Maximum value is 500 and default value is
// 500.
func (c *ZoneOperationsListCall) MaxResults(maxResults int64) *ZoneOperationsListCall {
	c.urlParams_.Set("maxResults", fmt.Sprint(maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Tag returned by a
// previous list request truncated by maxResults. Used to continue a
// previous list request.
func (c *ZoneOperationsListCall) PageToken(pageToken string) *ZoneOperationsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ZoneOperationsListCall) Fields(s ...googleapi.Field) *ZoneOperationsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ZoneOperationsListCall) IfNoneMatch(entityTag string) *ZoneOperationsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ZoneOperationsListCall) Context(ctx context.Context) *ZoneOperationsListCall {
	c.ctx_ = ctx
	return c
}

func (c *ZoneOperationsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "{project}/zones/{zone}/operations")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"project": c.project,
		"zone":    c.zone,
	})
	if c.ctx_ != nil {
		return ctxhttp.Do(c.ctx_, c.s.client, req)
	}
	return c.s.client.Do(req)
}

// Do executes the "replicapoolupdater.zoneOperations.list" call.
// Exactly one of *OperationList or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *OperationList.ServerResponse.Header or (if a response was returned
// at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ZoneOperationsListCall) Do(opts ...googleapi.CallOption) (*OperationList, error) {
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
	ret := &OperationList{
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
	//   "description": "Retrieves the list of Operation resources contained within the specified zone.",
	//   "httpMethod": "GET",
	//   "id": "replicapoolupdater.zoneOperations.list",
	//   "parameterOrder": [
	//     "project",
	//     "zone"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "description": "Optional. Filter expression for filtering listed resources.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "default": "500",
	//       "description": "Optional. Maximum count of results to be returned. Maximum value is 500 and default value is 500.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "500",
	//       "minimum": "0",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Optional. Tag returned by a previous list request truncated by maxResults. Used to continue a previous list request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Name of the project scoping this request.",
	//       "location": "path",
	//       "pattern": "(?:(?:[-a-z0-9]{1,63}\\.)*(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?):)?(?:[0-9]{1,19}|(?:[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?))",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "zone": {
	//       "description": "Name of the zone scoping this request.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{project}/zones/{zone}/operations",
	//   "response": {
	//     "$ref": "OperationList"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/replicapool"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *ZoneOperationsListCall) Pages(ctx context.Context, f func(*OperationList) error) error {
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
