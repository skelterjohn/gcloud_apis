// Package pubsub provides access to the Google Cloud Pub/Sub API.
//
// See https://cloud.google.com/pubsub/docs
//
// Usage example:
//
//   import "google.golang.org/api/pubsub/v1"
//   ...
//   pubsubService, err := pubsub.New(oauthHttpClient)
package pubsub

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

const apiId = "pubsub:v1"
const apiName = "pubsub"
const apiVersion = "v1"
const basePath = "https://pubsub.googleapis.com/"

// OAuth2 scopes used by this API.
const (
	// View and manage your data across Google Cloud Platform services
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"

	// View and manage Pub/Sub topics and subscriptions
	PubsubScope = "https://www.googleapis.com/auth/pubsub"
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
	rs.Subscriptions = NewProjectsSubscriptionsService(s)
	rs.Topics = NewProjectsTopicsService(s)
	return rs
}

type ProjectsService struct {
	s *Service

	Subscriptions *ProjectsSubscriptionsService

	Topics *ProjectsTopicsService
}

func NewProjectsSubscriptionsService(s *Service) *ProjectsSubscriptionsService {
	rs := &ProjectsSubscriptionsService{s: s}
	return rs
}

type ProjectsSubscriptionsService struct {
	s *Service
}

func NewProjectsTopicsService(s *Service) *ProjectsTopicsService {
	rs := &ProjectsTopicsService{s: s}
	rs.Subscriptions = NewProjectsTopicsSubscriptionsService(s)
	return rs
}

type ProjectsTopicsService struct {
	s *Service

	Subscriptions *ProjectsTopicsSubscriptionsService
}

func NewProjectsTopicsSubscriptionsService(s *Service) *ProjectsTopicsSubscriptionsService {
	rs := &ProjectsTopicsSubscriptionsService{s: s}
	return rs
}

type ProjectsTopicsSubscriptionsService struct {
	s *Service
}

type AcknowledgeRequest struct {
	// AckIds: The acknowledgment ID for the messages being acknowledged
	// that was returned
	// by the Pub/Sub system in the `Pull` response. Must
	// not be empty.
	AckIds []string `json:"ackIds,omitempty"`
}

type AuditConfig struct {
	// ExemptedMembers: Specifies the identities that are exempted from
	// "data access" audit
	// logging for the `service` specified
	// above.
	// Follows the same format of Binding.members.
	ExemptedMembers []string `json:"exemptedMembers,omitempty"`

	// Service: Specifies a service that will be enabled for "data access"
	// audit
	// logging.
	// For example, `resourcemanager`, `storage`,
	// `compute`.
	// `allServices` is a special value that covers all services.
	Service string `json:"service,omitempty"`
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

type CloudAuditOptions struct {
}

type Condition struct {
	// Iam: Trusted attributes supplied by the IAM system.
	Iam string `json:"iam,omitempty"`

	// Op: An operator to apply the subject with.
	Op string `json:"op,omitempty"`

	// Svc: Trusted attributes discharged by the service.
	Svc string `json:"svc,omitempty"`

	// Sys: Trusted attributes supplied by any service that owns resources
	// and uses
	// the IAM system for access control.
	Sys string `json:"sys,omitempty"`

	// Value: The object of the condition. Exactly one of these must be set.
	Value string `json:"value,omitempty"`

	// Values: The objects of the condition. This is mutually exclusive with
	// 'value'.
	Values []string `json:"values,omitempty"`
}

type CounterOptions struct {
	// Field: The field value to attribute.
	Field string `json:"field,omitempty"`

	// Metric: The metric to update.
	Metric string `json:"metric,omitempty"`
}

type DataAccessOptions struct {
}

type Empty struct {
}

type ListSubscriptionsResponse struct {
	// NextPageToken: If not empty, indicates that there may be more
	// subscriptions that match
	// the request; this value should be passed in
	// a new
	// `ListSubscriptionsRequest` to get more subscriptions.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Subscriptions: The subscriptions that match the request.
	Subscriptions []*Subscription `json:"subscriptions,omitempty"`
}

type ListTopicSubscriptionsResponse struct {
	// NextPageToken: If not empty, indicates that there may be more
	// subscriptions that match
	// the request; this value should be passed in
	// a new
	// `ListTopicSubscriptionsRequest` to get more subscriptions.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Subscriptions: The names of the subscriptions that match the request.
	Subscriptions []string `json:"subscriptions,omitempty"`
}

type ListTopicsResponse struct {
	// NextPageToken: If not empty, indicates that there may be more topics
	// that match the
	// request; this value should be passed in a new
	// `ListTopicsRequest`.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Topics: The resulting topics.
	Topics []*Topic `json:"topics,omitempty"`
}

type LogConfig struct {
	// CloudAudit: Cloud audit options.
	CloudAudit *CloudAuditOptions `json:"cloudAudit,omitempty"`

	// Counter: Counter options.
	Counter *CounterOptions `json:"counter,omitempty"`

	// DataAccess: Data access options.
	DataAccess *DataAccessOptions `json:"dataAccess,omitempty"`
}

type ModifyAckDeadlineRequest struct {
	// AckDeadlineSeconds: The new ack deadline with respect to the time
	// this request was sent to
	// the Pub/Sub system. Must be >= 0. For
	// example, if the value is 10, the new
	// ack deadline will expire 10
	// seconds after the `ModifyAckDeadline` call
	// was made. Specifying zero
	// may immediately make the message available for
	// another pull request.
	AckDeadlineSeconds int64 `json:"ackDeadlineSeconds,omitempty"`

	// AckIds: List of acknowledgment IDs.
	AckIds []string `json:"ackIds,omitempty"`
}

type ModifyPushConfigRequest struct {
	// PushConfig: The push configuration for future deliveries.
	//
	// An empty
	// `pushConfig` indicates that the Pub/Sub system should
	// stop pushing
	// messages from the given subscription and allow
	// messages to be pulled
	// and acknowledged - effectively pausing
	// the subscription if `Pull` is
	// not called.
	PushConfig *PushConfig `json:"pushConfig,omitempty"`
}

type Policy struct {
	// AuditConfigs: Specifies audit logging configs for "data
	// access".
	// "data access": generally refers to data reads/writes and
	// admin reads.
	// "admin activity": generally refers to admin
	// writes.
	//
	// Note: `AuditConfig` doesn't apply to "admin activity", which
	// always
	// enables audit logging.
	AuditConfigs []*AuditConfig `json:"auditConfigs,omitempty"`

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

	IamOwned bool `json:"iamOwned,omitempty"`

	Rules []*Rule `json:"rules,omitempty"`

	// Version: Version of the `Policy`. The default version is 0.
	// 0 =
	// resourcemanager_projects only support legacy roles.
	// 1 = supports
	// non-legacy roles
	// 2 = supports AuditConfig
	Version int64 `json:"version,omitempty"`
}

type PublishRequest struct {
	// Messages: The messages to publish.
	Messages []*PubsubMessage `json:"messages,omitempty"`
}

type PublishResponse struct {
	// MessageIds: The server-assigned ID of each published message, in the
	// same order as
	// the messages in the request. IDs are guaranteed to be
	// unique within
	// the topic.
	MessageIds []string `json:"messageIds,omitempty"`
}

type PubsubMessage struct {
	// Attributes: Optional attributes for this message.
	Attributes map[string]string `json:"attributes,omitempty"`

	// Data: The message payload. For JSON requests, the value of this field
	// must be
	// base64-encoded.
	Data string `json:"data,omitempty"`

	// MessageId: ID of this message, assigned by the server when the
	// message is published.
	// Guaranteed to be unique within the topic. This
	// value may be read by a
	// subscriber that receives a `PubsubMessage` via
	// a `Pull` call or a push
	// delivery. It must not be populated by the
	// publisher in a `Publish` call.
	MessageId string `json:"messageId,omitempty"`

	// PublishTime: The time at which the message was published, populated
	// by the server when
	// it receives the `Publish` call. It must not be
	// populated by the
	// publisher in a `Publish` call.
	PublishTime string `json:"publishTime,omitempty"`
}

type PullRequest struct {
	// MaxMessages: The maximum number of messages returned for this
	// request. The Pub/Sub
	// system may return fewer than the number
	// specified.
	MaxMessages int64 `json:"maxMessages,omitempty"`

	// ReturnImmediately: If this is specified as true the system will
	// respond immediately even if
	// it is not able to return a message in the
	// `Pull` response. Otherwise the
	// system is allowed to wait until at
	// least one message is available rather
	// than returning no messages. The
	// client may cancel the request if it does
	// not wish to wait any longer
	// for the response.
	ReturnImmediately bool `json:"returnImmediately,omitempty"`
}

type PullResponse struct {
	// ReceivedMessages: Received Pub/Sub messages. The Pub/Sub system will
	// return zero messages if
	// there are no more available in the backlog.
	// The Pub/Sub system may return
	// fewer than the `maxMessages` requested
	// even if there are more messages
	// available in the backlog.
	ReceivedMessages []*ReceivedMessage `json:"receivedMessages,omitempty"`
}

type PushConfig struct {
	// Attributes: Endpoint configuration attributes.
	//
	// Every endpoint has a
	// set of API supported attributes that can be used to
	// control different
	// aspects of the message delivery.
	//
	// The currently supported attribute
	// is `x-goog-version`, which you can
	// use to change the format of the
	// push message. This attribute
	// indicates the version of the data
	// expected by the endpoint. This
	// controls the shape of the envelope
	// (i.e. its fields and metadata).
	// The endpoint version is based on the
	// version of the Pub/Sub
	// API.
	//
	// If not present during the
	// `CreateSubscription` call, it will default to
	// the version of the API
	// used to make such call. If not present during a
	// `ModifyPushConfig`
	// call, its value will not be changed. `GetSubscription`
	// calls will
	// always return a valid version, even if the subscription was
	// created
	// without this attribute.
	//
	// The possible values for this attribute
	// are:
	//
	// * `v1beta1`: uses the push format defined in the v1beta1
	// Pub/Sub API.
	// * `v1` or `v1beta2`: uses the push format defined in the
	// v1 Pub/Sub API.
	Attributes map[string]string `json:"attributes,omitempty"`

	// PushEndpoint: A URL locating the endpoint to which messages should be
	// pushed.
	// For example, a Webhook endpoint might use
	// "https://example.com/push".
	PushEndpoint string `json:"pushEndpoint,omitempty"`
}

type ReceivedMessage struct {
	// AckId: This ID can be used to acknowledge the received message.
	AckId string `json:"ackId,omitempty"`

	// Message: The message.
	Message *PubsubMessage `json:"message,omitempty"`
}

type Rule struct {
	// Action: Required
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
	// The format for in and not_in entries is the same
	// as for members in a
	// Binding (see google/iam/v1/policy.proto).
	NotIn []string `json:"notIn,omitempty"`

	// Permissions: A permission is a string of form '<service>.<resource
	// type>.<verb>'
	// (e.g., 'storage.buckets.list'). A value of '*' matches
	// all permissions,
	// and a verb part of '*' (e.g., 'storage.buckets.*')
	// matches all verbs.
	Permissions []string `json:"permissions,omitempty"`
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

type Subscription struct {
	// AckDeadlineSeconds: This value is the maximum time after a subscriber
	// receives a message
	// before the subscriber should acknowledge the
	// message. After message
	// delivery but before the ack deadline expires
	// and before the message is
	// acknowledged, it is an outstanding message
	// and will not be delivered
	// again during that time (on a best-effort
	// basis).
	//
	// For pull subscriptions, this value is used as the initial
	// value for the ack
	// deadline. To override this value for a given
	// message, call
	// `ModifyAckDeadline` with the corresponding `ack_id` if
	// using
	// pull.
	//
	// For push delivery, this value is also used to set the
	// request timeout for
	// the call to the push endpoint.
	//
	// If the subscriber
	// never acknowledges the message, the Pub/Sub
	// system will eventually
	// redeliver the message.
	//
	// If this parameter is not set, the default
	// value of 10 seconds is used.
	AckDeadlineSeconds int64 `json:"ackDeadlineSeconds,omitempty"`

	// Name: The name of the subscription. It must have the
	// format
	// `"projects/{project}/subscriptions/{subscription}"`.
	// `{subscription}` must
	// start with a letter, and contain only letters
	// (`[A-Za-z]`), numbers
	// (`[0-9]`), dashes (`-`), underscores (`_`),
	// periods (`.`), tildes (`~`),
	// plus (`+`) or percent signs (`%`). It
	// must be between 3 and 255 characters
	// in length, and it must not start
	// with `"goog"`.
	Name string `json:"name,omitempty"`

	// PushConfig: If push delivery is used with this subscription, this
	// field is
	// used to configure it. An empty `pushConfig` signifies that
	// the subscriber
	// will pull and ack messages using API methods.
	PushConfig *PushConfig `json:"pushConfig,omitempty"`

	// Topic: The name of the topic from which this subscription is
	// receiving messages.
	// The value of this field will be `_deleted-topic_`
	// if the topic has been
	// deleted.
	Topic string `json:"topic,omitempty"`
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

type Topic struct {
	// Name: The name of the topic. It must have the
	// format
	// `"projects/{project}/topics/{topic}"`. `{topic}` must start
	// with a letter,
	// and contain only letters (`[A-Za-z]`), numbers
	// (`[0-9]`), dashes (`-`),
	// underscores (`_`), periods (`.`), tildes
	// (`~`), plus (`+`) or percent
	// signs (`%`). It must be between 3 and
	// 255 characters in length, and it
	// must not start with `"goog"`.
	Name string `json:"name,omitempty"`
}

// method id "pubsub.projects.subscriptions.acknowledge":

type ProjectsSubscriptionsAcknowledgeCall struct {
	s                  *Service
	subscription       string
	acknowledgerequest *AcknowledgeRequest
	opt_               map[string]interface{}
}

// Acknowledge: Acknowledges the messages associated with the `ack_ids`
// in the
// `AcknowledgeRequest`. The Pub/Sub system can remove the
// relevant messages
// from the subscription.
//
// Acknowledging a message
// whose ack deadline has expired may succeed,
// but such a message may be
// redelivered later. Acknowledging a message more
// than once will not
// result in an error.
func (r *ProjectsSubscriptionsService) Acknowledge(subscription string, acknowledgerequest *AcknowledgeRequest) *ProjectsSubscriptionsAcknowledgeCall {
	c := &ProjectsSubscriptionsAcknowledgeCall{s: r.s, opt_: make(map[string]interface{})}
	c.subscription = subscription
	c.acknowledgerequest = acknowledgerequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsSubscriptionsAcknowledgeCall) Fields(s ...googleapi.Field) *ProjectsSubscriptionsAcknowledgeCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsSubscriptionsAcknowledgeCall) Do() (*Empty, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.acknowledgerequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+subscription}:acknowledge")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"subscription": c.subscription,
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
	//   "description": "Acknowledges the messages associated with the `ack_ids` in the\n`AcknowledgeRequest`. The Pub/Sub system can remove the relevant messages\nfrom the subscription.\n\nAcknowledging a message whose ack deadline has expired may succeed,\nbut such a message may be redelivered later. Acknowledging a message more\nthan once will not result in an error.",
	//   "flatPath": "v1/projects/{projectsId}/subscriptions/{subscriptionsId}:acknowledge",
	//   "httpMethod": "POST",
	//   "id": "pubsub.projects.subscriptions.acknowledge",
	//   "parameterOrder": [
	//     "subscription"
	//   ],
	//   "parameters": {
	//     "subscription": {
	//       "description": "The subscription whose message is being acknowledged.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]*/subscriptions/[^/]*$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+subscription}:acknowledge",
	//   "request": {
	//     "$ref": "AcknowledgeRequest"
	//   },
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/pubsub"
	//   ]
	// }

}

// method id "pubsub.projects.subscriptions.create":

type ProjectsSubscriptionsCreateCall struct {
	s            *Service
	name         string
	subscription *Subscription
	opt_         map[string]interface{}
}

// Create: Creates a subscription to a given topic for a given
// subscriber.
// If the subscription already exists, returns
// `ALREADY_EXISTS`.
// If the corresponding topic doesn't exist, returns
// `NOT_FOUND`.
//
// If the name is not provided in the request, the server
// will assign a random
// name for this subscription on the same project
// as the topic.
func (r *ProjectsSubscriptionsService) Create(name string, subscription *Subscription) *ProjectsSubscriptionsCreateCall {
	c := &ProjectsSubscriptionsCreateCall{s: r.s, opt_: make(map[string]interface{})}
	c.name = name
	c.subscription = subscription
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsSubscriptionsCreateCall) Fields(s ...googleapi.Field) *ProjectsSubscriptionsCreateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsSubscriptionsCreateCall) Do() (*Subscription, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.subscription)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
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
	var ret *Subscription
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a subscription to a given topic for a given subscriber.\nIf the subscription already exists, returns `ALREADY_EXISTS`.\nIf the corresponding topic doesn't exist, returns `NOT_FOUND`.\n\nIf the name is not provided in the request, the server will assign a random\nname for this subscription on the same project as the topic.",
	//   "flatPath": "v1/projects/{projectsId}/subscriptions/{subscriptionsId}",
	//   "httpMethod": "PUT",
	//   "id": "pubsub.projects.subscriptions.create",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The name of the subscription. It must have the format\n`\"projects/{project}/subscriptions/{subscription}\"`. `{subscription}` must\nstart with a letter, and contain only letters (`[A-Za-z]`), numbers\n(`[0-9]`), dashes (`-`), underscores (`_`), periods (`.`), tildes (`~`),\nplus (`+`) or percent signs (`%`). It must be between 3 and 255 characters\nin length, and it must not start with `\"goog\"`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]*/subscriptions/[^/]*$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}",
	//   "request": {
	//     "$ref": "Subscription"
	//   },
	//   "response": {
	//     "$ref": "Subscription"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/pubsub"
	//   ]
	// }

}

// method id "pubsub.projects.subscriptions.delete":

type ProjectsSubscriptionsDeleteCall struct {
	s            *Service
	subscription string
	opt_         map[string]interface{}
}

// Delete: Deletes an existing subscription. All pending messages in the
// subscription
// are immediately dropped. Calls to `Pull` after deletion
// will return
// `NOT_FOUND`. After a subscription is deleted, a new one
// may be created with
// the same name, but the new one has no association
// with the old
// subscription, or its topic unless the same topic is
// specified.
func (r *ProjectsSubscriptionsService) Delete(subscription string) *ProjectsSubscriptionsDeleteCall {
	c := &ProjectsSubscriptionsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.subscription = subscription
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsSubscriptionsDeleteCall) Fields(s ...googleapi.Field) *ProjectsSubscriptionsDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsSubscriptionsDeleteCall) Do() (*Empty, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+subscription}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"subscription": c.subscription,
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
	//   "description": "Deletes an existing subscription. All pending messages in the subscription\nare immediately dropped. Calls to `Pull` after deletion will return\n`NOT_FOUND`. After a subscription is deleted, a new one may be created with\nthe same name, but the new one has no association with the old\nsubscription, or its topic unless the same topic is specified.",
	//   "flatPath": "v1/projects/{projectsId}/subscriptions/{subscriptionsId}",
	//   "httpMethod": "DELETE",
	//   "id": "pubsub.projects.subscriptions.delete",
	//   "parameterOrder": [
	//     "subscription"
	//   ],
	//   "parameters": {
	//     "subscription": {
	//       "description": "The subscription to delete.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]*/subscriptions/[^/]*$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+subscription}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/pubsub"
	//   ]
	// }

}

// method id "pubsub.projects.subscriptions.get":

type ProjectsSubscriptionsGetCall struct {
	s            *Service
	subscription string
	opt_         map[string]interface{}
}

// Get: Gets the configuration details of a subscription.
func (r *ProjectsSubscriptionsService) Get(subscription string) *ProjectsSubscriptionsGetCall {
	c := &ProjectsSubscriptionsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.subscription = subscription
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsSubscriptionsGetCall) Fields(s ...googleapi.Field) *ProjectsSubscriptionsGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsSubscriptionsGetCall) Do() (*Subscription, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+subscription}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"subscription": c.subscription,
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
	var ret *Subscription
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets the configuration details of a subscription.",
	//   "flatPath": "v1/projects/{projectsId}/subscriptions/{subscriptionsId}",
	//   "httpMethod": "GET",
	//   "id": "pubsub.projects.subscriptions.get",
	//   "parameterOrder": [
	//     "subscription"
	//   ],
	//   "parameters": {
	//     "subscription": {
	//       "description": "The name of the subscription to get.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]*/subscriptions/[^/]*$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+subscription}",
	//   "response": {
	//     "$ref": "Subscription"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/pubsub"
	//   ]
	// }

}

// method id "pubsub.projects.subscriptions.getIamPolicy":

type ProjectsSubscriptionsGetIamPolicyCall struct {
	s        *Service
	resource string
	opt_     map[string]interface{}
}

// GetIamPolicy: Gets the access control policy for a `resource`. Is
// empty if the
// policy or the resource does not exist.
func (r *ProjectsSubscriptionsService) GetIamPolicy(resource string) *ProjectsSubscriptionsGetIamPolicyCall {
	c := &ProjectsSubscriptionsGetIamPolicyCall{s: r.s, opt_: make(map[string]interface{})}
	c.resource = resource
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsSubscriptionsGetIamPolicyCall) Fields(s ...googleapi.Field) *ProjectsSubscriptionsGetIamPolicyCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsSubscriptionsGetIamPolicyCall) Do() (*Policy, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+resource}:getIamPolicy")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
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
	var ret *Policy
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets the access control policy for a `resource`. Is empty if the\npolicy or the resource does not exist.",
	//   "flatPath": "v1/projects/{projectsId}/subscriptions/{subscriptionsId}:getIamPolicy",
	//   "httpMethod": "GET",
	//   "id": "pubsub.projects.subscriptions.getIamPolicy",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "REQUIRED: The resource for which policy is being requested.\n`resource` is usually specified as a path, such as,\n`projects/{project}/zones/{zone}/disks/{disk}`.\n\nThe format for the path specified in this value is resource specific and\nis specified in the documentation for the respective GetIamPolicy rpc.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]*/subscriptions/[^/]*$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+resource}:getIamPolicy",
	//   "response": {
	//     "$ref": "Policy"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/pubsub"
	//   ]
	// }

}

// method id "pubsub.projects.subscriptions.list":

type ProjectsSubscriptionsListCall struct {
	s       *Service
	project string
	opt_    map[string]interface{}
}

// List: Lists matching subscriptions.
func (r *ProjectsSubscriptionsService) List(project string) *ProjectsSubscriptionsListCall {
	c := &ProjectsSubscriptionsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	return c
}

// PageSize sets the optional parameter "pageSize": Maximum number of
// subscriptions to return.
func (c *ProjectsSubscriptionsListCall) PageSize(pageSize int64) *ProjectsSubscriptionsListCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken": The value returned
// by the last `ListSubscriptionsResponse`; indicates that
// this is a
// continuation of a prior `ListSubscriptions` call, and that the
// system
// should return the next page of data.
func (c *ProjectsSubscriptionsListCall) PageToken(pageToken string) *ProjectsSubscriptionsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsSubscriptionsListCall) Fields(s ...googleapi.Field) *ProjectsSubscriptionsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsSubscriptionsListCall) Do() (*ListSubscriptionsResponse, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+project}/subscriptions")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project": c.project,
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
	var ret *ListSubscriptionsResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists matching subscriptions.",
	//   "flatPath": "v1/projects/{projectsId}/subscriptions",
	//   "httpMethod": "GET",
	//   "id": "pubsub.projects.subscriptions.list",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "Maximum number of subscriptions to return.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The value returned by the last `ListSubscriptionsResponse`; indicates that\nthis is a continuation of a prior `ListSubscriptions` call, and that the\nsystem should return the next page of data.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "The name of the cloud project that subscriptions belong to.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]*$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+project}/subscriptions",
	//   "response": {
	//     "$ref": "ListSubscriptionsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/pubsub"
	//   ]
	// }

}

// method id "pubsub.projects.subscriptions.modifyAckDeadline":

type ProjectsSubscriptionsModifyAckDeadlineCall struct {
	s                        *Service
	subscription             string
	modifyackdeadlinerequest *ModifyAckDeadlineRequest
	opt_                     map[string]interface{}
}

// ModifyAckDeadline: Modifies the ack deadline for a specific message.
// This method is useful
// to indicate that more time is needed to process
// a message by the
// subscriber, or to make the message available for
// redelivery if the
// processing was interrupted.
func (r *ProjectsSubscriptionsService) ModifyAckDeadline(subscription string, modifyackdeadlinerequest *ModifyAckDeadlineRequest) *ProjectsSubscriptionsModifyAckDeadlineCall {
	c := &ProjectsSubscriptionsModifyAckDeadlineCall{s: r.s, opt_: make(map[string]interface{})}
	c.subscription = subscription
	c.modifyackdeadlinerequest = modifyackdeadlinerequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsSubscriptionsModifyAckDeadlineCall) Fields(s ...googleapi.Field) *ProjectsSubscriptionsModifyAckDeadlineCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsSubscriptionsModifyAckDeadlineCall) Do() (*Empty, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.modifyackdeadlinerequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+subscription}:modifyAckDeadline")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"subscription": c.subscription,
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
	//   "description": "Modifies the ack deadline for a specific message. This method is useful\nto indicate that more time is needed to process a message by the\nsubscriber, or to make the message available for redelivery if the\nprocessing was interrupted.",
	//   "flatPath": "v1/projects/{projectsId}/subscriptions/{subscriptionsId}:modifyAckDeadline",
	//   "httpMethod": "POST",
	//   "id": "pubsub.projects.subscriptions.modifyAckDeadline",
	//   "parameterOrder": [
	//     "subscription"
	//   ],
	//   "parameters": {
	//     "subscription": {
	//       "description": "The name of the subscription.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]*/subscriptions/[^/]*$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+subscription}:modifyAckDeadline",
	//   "request": {
	//     "$ref": "ModifyAckDeadlineRequest"
	//   },
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/pubsub"
	//   ]
	// }

}

// method id "pubsub.projects.subscriptions.modifyPushConfig":

type ProjectsSubscriptionsModifyPushConfigCall struct {
	s                       *Service
	subscription            string
	modifypushconfigrequest *ModifyPushConfigRequest
	opt_                    map[string]interface{}
}

// ModifyPushConfig: Modifies the `PushConfig` for a specified
// subscription.
//
// This may be used to change a push subscription to a
// pull one (signified by
// an empty `PushConfig`) or vice versa, or
// change the endpoint URL and other
// attributes of a push subscription.
// Messages will accumulate for delivery
// continuously through the call
// regardless of changes to the `PushConfig`.
func (r *ProjectsSubscriptionsService) ModifyPushConfig(subscription string, modifypushconfigrequest *ModifyPushConfigRequest) *ProjectsSubscriptionsModifyPushConfigCall {
	c := &ProjectsSubscriptionsModifyPushConfigCall{s: r.s, opt_: make(map[string]interface{})}
	c.subscription = subscription
	c.modifypushconfigrequest = modifypushconfigrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsSubscriptionsModifyPushConfigCall) Fields(s ...googleapi.Field) *ProjectsSubscriptionsModifyPushConfigCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsSubscriptionsModifyPushConfigCall) Do() (*Empty, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.modifypushconfigrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+subscription}:modifyPushConfig")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"subscription": c.subscription,
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
	//   "description": "Modifies the `PushConfig` for a specified subscription.\n\nThis may be used to change a push subscription to a pull one (signified by\nan empty `PushConfig`) or vice versa, or change the endpoint URL and other\nattributes of a push subscription. Messages will accumulate for delivery\ncontinuously through the call regardless of changes to the `PushConfig`.",
	//   "flatPath": "v1/projects/{projectsId}/subscriptions/{subscriptionsId}:modifyPushConfig",
	//   "httpMethod": "POST",
	//   "id": "pubsub.projects.subscriptions.modifyPushConfig",
	//   "parameterOrder": [
	//     "subscription"
	//   ],
	//   "parameters": {
	//     "subscription": {
	//       "description": "The name of the subscription.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]*/subscriptions/[^/]*$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+subscription}:modifyPushConfig",
	//   "request": {
	//     "$ref": "ModifyPushConfigRequest"
	//   },
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/pubsub"
	//   ]
	// }

}

// method id "pubsub.projects.subscriptions.pull":

type ProjectsSubscriptionsPullCall struct {
	s            *Service
	subscription string
	pullrequest  *PullRequest
	opt_         map[string]interface{}
}

// Pull: Pulls messages from the server. Returns an empty list if there
// are no
// messages available in the backlog. The server may return
// `UNAVAILABLE` if
// there are too many concurrent pull requests pending
// for the given
// subscription.
func (r *ProjectsSubscriptionsService) Pull(subscription string, pullrequest *PullRequest) *ProjectsSubscriptionsPullCall {
	c := &ProjectsSubscriptionsPullCall{s: r.s, opt_: make(map[string]interface{})}
	c.subscription = subscription
	c.pullrequest = pullrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsSubscriptionsPullCall) Fields(s ...googleapi.Field) *ProjectsSubscriptionsPullCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsSubscriptionsPullCall) Do() (*PullResponse, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.pullrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+subscription}:pull")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"subscription": c.subscription,
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
	var ret *PullResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Pulls messages from the server. Returns an empty list if there are no\nmessages available in the backlog. The server may return `UNAVAILABLE` if\nthere are too many concurrent pull requests pending for the given\nsubscription.",
	//   "flatPath": "v1/projects/{projectsId}/subscriptions/{subscriptionsId}:pull",
	//   "httpMethod": "POST",
	//   "id": "pubsub.projects.subscriptions.pull",
	//   "parameterOrder": [
	//     "subscription"
	//   ],
	//   "parameters": {
	//     "subscription": {
	//       "description": "The subscription from which messages should be pulled.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]*/subscriptions/[^/]*$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+subscription}:pull",
	//   "request": {
	//     "$ref": "PullRequest"
	//   },
	//   "response": {
	//     "$ref": "PullResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/pubsub"
	//   ]
	// }

}

// method id "pubsub.projects.subscriptions.setIamPolicy":

type ProjectsSubscriptionsSetIamPolicyCall struct {
	s                   *Service
	resource            string
	setiampolicyrequest *SetIamPolicyRequest
	opt_                map[string]interface{}
}

// SetIamPolicy: Sets the access control policy on the specified
// resource. Replaces any
// existing policy.
func (r *ProjectsSubscriptionsService) SetIamPolicy(resource string, setiampolicyrequest *SetIamPolicyRequest) *ProjectsSubscriptionsSetIamPolicyCall {
	c := &ProjectsSubscriptionsSetIamPolicyCall{s: r.s, opt_: make(map[string]interface{})}
	c.resource = resource
	c.setiampolicyrequest = setiampolicyrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsSubscriptionsSetIamPolicyCall) Fields(s ...googleapi.Field) *ProjectsSubscriptionsSetIamPolicyCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsSubscriptionsSetIamPolicyCall) Do() (*Policy, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+resource}:setIamPolicy")
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
	//   "description": "Sets the access control policy on the specified resource. Replaces any\nexisting policy.",
	//   "flatPath": "v1/projects/{projectsId}/subscriptions/{subscriptionsId}:setIamPolicy",
	//   "httpMethod": "POST",
	//   "id": "pubsub.projects.subscriptions.setIamPolicy",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "REQUIRED: The resource for which policy is being specified.\n`resource` is usually specified as a path, such as,\n`projects/{project}/zones/{zone}/disks/{disk}`.\n\nThe format for the path specified in this value is resource specific and\nis specified in the documentation for the respective SetIamPolicy rpc.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]*/subscriptions/[^/]*$",
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
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/pubsub"
	//   ]
	// }

}

// method id "pubsub.projects.subscriptions.testIamPermissions":

type ProjectsSubscriptionsTestIamPermissionsCall struct {
	s                         *Service
	resource                  string
	testiampermissionsrequest *TestIamPermissionsRequest
	opt_                      map[string]interface{}
}

// TestIamPermissions: Returns permissions that a caller has on the
// specified resource.
func (r *ProjectsSubscriptionsService) TestIamPermissions(resource string, testiampermissionsrequest *TestIamPermissionsRequest) *ProjectsSubscriptionsTestIamPermissionsCall {
	c := &ProjectsSubscriptionsTestIamPermissionsCall{s: r.s, opt_: make(map[string]interface{})}
	c.resource = resource
	c.testiampermissionsrequest = testiampermissionsrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsSubscriptionsTestIamPermissionsCall) Fields(s ...googleapi.Field) *ProjectsSubscriptionsTestIamPermissionsCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsSubscriptionsTestIamPermissionsCall) Do() (*TestIamPermissionsResponse, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+resource}:testIamPermissions")
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
	//   "description": "Returns permissions that a caller has on the specified resource.",
	//   "flatPath": "v1/projects/{projectsId}/subscriptions/{subscriptionsId}:testIamPermissions",
	//   "httpMethod": "POST",
	//   "id": "pubsub.projects.subscriptions.testIamPermissions",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "REQUIRED: The resource for which policy detail is being requested.\n`resource` is usually specified as a path, such as,\n`projects/{project}/zones/{zone}/disks/{disk}`.\n\nThe format for the path specified in this value is resource specific and\nis specified in the documentation for the respective TestIamPermissions\nrpc.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]*/subscriptions/[^/]*$",
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
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/pubsub"
	//   ]
	// }

}

// method id "pubsub.projects.topics.create":

type ProjectsTopicsCreateCall struct {
	s     *Service
	name  string
	topic *Topic
	opt_  map[string]interface{}
}

// Create: Creates the given topic with the given name.
func (r *ProjectsTopicsService) Create(name string, topic *Topic) *ProjectsTopicsCreateCall {
	c := &ProjectsTopicsCreateCall{s: r.s, opt_: make(map[string]interface{})}
	c.name = name
	c.topic = topic
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsTopicsCreateCall) Fields(s ...googleapi.Field) *ProjectsTopicsCreateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsTopicsCreateCall) Do() (*Topic, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.topic)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
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
	var ret *Topic
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates the given topic with the given name.",
	//   "flatPath": "v1/projects/{projectsId}/topics/{topicsId}",
	//   "httpMethod": "PUT",
	//   "id": "pubsub.projects.topics.create",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "The name of the topic. It must have the format\n`\"projects/{project}/topics/{topic}\"`. `{topic}` must start with a letter,\nand contain only letters (`[A-Za-z]`), numbers (`[0-9]`), dashes (`-`),\nunderscores (`_`), periods (`.`), tildes (`~`), plus (`+`) or percent\nsigns (`%`). It must be between 3 and 255 characters in length, and it\nmust not start with `\"goog\"`.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]*/topics/[^/]*$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}",
	//   "request": {
	//     "$ref": "Topic"
	//   },
	//   "response": {
	//     "$ref": "Topic"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/pubsub"
	//   ]
	// }

}

// method id "pubsub.projects.topics.delete":

type ProjectsTopicsDeleteCall struct {
	s     *Service
	topic string
	opt_  map[string]interface{}
}

// Delete: Deletes the topic with the given name. Returns `NOT_FOUND` if
// the topic
// does not exist. After a topic is deleted, a new topic may
// be created with
// the same name; this is an entirely new topic with
// none of the old
// configuration or subscriptions. Existing
// subscriptions to this topic are
// not deleted, but their `topic` field
// is set to `_deleted-topic_`.
func (r *ProjectsTopicsService) Delete(topic string) *ProjectsTopicsDeleteCall {
	c := &ProjectsTopicsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.topic = topic
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsTopicsDeleteCall) Fields(s ...googleapi.Field) *ProjectsTopicsDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsTopicsDeleteCall) Do() (*Empty, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+topic}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"topic": c.topic,
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
	//   "description": "Deletes the topic with the given name. Returns `NOT_FOUND` if the topic\ndoes not exist. After a topic is deleted, a new topic may be created with\nthe same name; this is an entirely new topic with none of the old\nconfiguration or subscriptions. Existing subscriptions to this topic are\nnot deleted, but their `topic` field is set to `_deleted-topic_`.",
	//   "flatPath": "v1/projects/{projectsId}/topics/{topicsId}",
	//   "httpMethod": "DELETE",
	//   "id": "pubsub.projects.topics.delete",
	//   "parameterOrder": [
	//     "topic"
	//   ],
	//   "parameters": {
	//     "topic": {
	//       "description": "Name of the topic to delete.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]*/topics/[^/]*$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+topic}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/pubsub"
	//   ]
	// }

}

// method id "pubsub.projects.topics.get":

type ProjectsTopicsGetCall struct {
	s     *Service
	topic string
	opt_  map[string]interface{}
}

// Get: Gets the configuration of a topic.
func (r *ProjectsTopicsService) Get(topic string) *ProjectsTopicsGetCall {
	c := &ProjectsTopicsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.topic = topic
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsTopicsGetCall) Fields(s ...googleapi.Field) *ProjectsTopicsGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsTopicsGetCall) Do() (*Topic, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+topic}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"topic": c.topic,
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
	var ret *Topic
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets the configuration of a topic.",
	//   "flatPath": "v1/projects/{projectsId}/topics/{topicsId}",
	//   "httpMethod": "GET",
	//   "id": "pubsub.projects.topics.get",
	//   "parameterOrder": [
	//     "topic"
	//   ],
	//   "parameters": {
	//     "topic": {
	//       "description": "The name of the topic to get.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]*/topics/[^/]*$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+topic}",
	//   "response": {
	//     "$ref": "Topic"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/pubsub"
	//   ]
	// }

}

// method id "pubsub.projects.topics.getIamPolicy":

type ProjectsTopicsGetIamPolicyCall struct {
	s        *Service
	resource string
	opt_     map[string]interface{}
}

// GetIamPolicy: Gets the access control policy for a `resource`. Is
// empty if the
// policy or the resource does not exist.
func (r *ProjectsTopicsService) GetIamPolicy(resource string) *ProjectsTopicsGetIamPolicyCall {
	c := &ProjectsTopicsGetIamPolicyCall{s: r.s, opt_: make(map[string]interface{})}
	c.resource = resource
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsTopicsGetIamPolicyCall) Fields(s ...googleapi.Field) *ProjectsTopicsGetIamPolicyCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsTopicsGetIamPolicyCall) Do() (*Policy, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+resource}:getIamPolicy")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"resource": c.resource,
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
	var ret *Policy
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets the access control policy for a `resource`. Is empty if the\npolicy or the resource does not exist.",
	//   "flatPath": "v1/projects/{projectsId}/topics/{topicsId}:getIamPolicy",
	//   "httpMethod": "GET",
	//   "id": "pubsub.projects.topics.getIamPolicy",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "REQUIRED: The resource for which policy is being requested.\n`resource` is usually specified as a path, such as,\n`projects/{project}/zones/{zone}/disks/{disk}`.\n\nThe format for the path specified in this value is resource specific and\nis specified in the documentation for the respective GetIamPolicy rpc.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]*/topics/[^/]*$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+resource}:getIamPolicy",
	//   "response": {
	//     "$ref": "Policy"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/pubsub"
	//   ]
	// }

}

// method id "pubsub.projects.topics.list":

type ProjectsTopicsListCall struct {
	s       *Service
	project string
	opt_    map[string]interface{}
}

// List: Lists matching topics.
func (r *ProjectsTopicsService) List(project string) *ProjectsTopicsListCall {
	c := &ProjectsTopicsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	return c
}

// PageSize sets the optional parameter "pageSize": Maximum number of
// topics to return.
func (c *ProjectsTopicsListCall) PageSize(pageSize int64) *ProjectsTopicsListCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken": The value returned
// by the last `ListTopicsResponse`; indicates that this is
// a
// continuation of a prior `ListTopics` call, and that the system
// should
// return the next page of data.
func (c *ProjectsTopicsListCall) PageToken(pageToken string) *ProjectsTopicsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsTopicsListCall) Fields(s ...googleapi.Field) *ProjectsTopicsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsTopicsListCall) Do() (*ListTopicsResponse, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+project}/topics")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project": c.project,
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
	var ret *ListTopicsResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists matching topics.",
	//   "flatPath": "v1/projects/{projectsId}/topics",
	//   "httpMethod": "GET",
	//   "id": "pubsub.projects.topics.list",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "Maximum number of topics to return.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The value returned by the last `ListTopicsResponse`; indicates that this is\na continuation of a prior `ListTopics` call, and that the system should\nreturn the next page of data.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "The name of the cloud project that topics belong to.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]*$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+project}/topics",
	//   "response": {
	//     "$ref": "ListTopicsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/pubsub"
	//   ]
	// }

}

// method id "pubsub.projects.topics.publish":

type ProjectsTopicsPublishCall struct {
	s              *Service
	topic          string
	publishrequest *PublishRequest
	opt_           map[string]interface{}
}

// Publish: Adds one or more messages to the topic. Returns `NOT_FOUND`
// if the topic
// does not exist. The message payload must not be empty;
// it must contain
//  either a non-empty data field, or at least one
// attribute.
func (r *ProjectsTopicsService) Publish(topic string, publishrequest *PublishRequest) *ProjectsTopicsPublishCall {
	c := &ProjectsTopicsPublishCall{s: r.s, opt_: make(map[string]interface{})}
	c.topic = topic
	c.publishrequest = publishrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsTopicsPublishCall) Fields(s ...googleapi.Field) *ProjectsTopicsPublishCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsTopicsPublishCall) Do() (*PublishResponse, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.publishrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+topic}:publish")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"topic": c.topic,
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
	var ret *PublishResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Adds one or more messages to the topic. Returns `NOT_FOUND` if the topic\ndoes not exist. The message payload must not be empty; it must contain\n either a non-empty data field, or at least one attribute.",
	//   "flatPath": "v1/projects/{projectsId}/topics/{topicsId}:publish",
	//   "httpMethod": "POST",
	//   "id": "pubsub.projects.topics.publish",
	//   "parameterOrder": [
	//     "topic"
	//   ],
	//   "parameters": {
	//     "topic": {
	//       "description": "The messages in the request will be published on this topic.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]*/topics/[^/]*$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+topic}:publish",
	//   "request": {
	//     "$ref": "PublishRequest"
	//   },
	//   "response": {
	//     "$ref": "PublishResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/pubsub"
	//   ]
	// }

}

// method id "pubsub.projects.topics.setIamPolicy":

type ProjectsTopicsSetIamPolicyCall struct {
	s                   *Service
	resource            string
	setiampolicyrequest *SetIamPolicyRequest
	opt_                map[string]interface{}
}

// SetIamPolicy: Sets the access control policy on the specified
// resource. Replaces any
// existing policy.
func (r *ProjectsTopicsService) SetIamPolicy(resource string, setiampolicyrequest *SetIamPolicyRequest) *ProjectsTopicsSetIamPolicyCall {
	c := &ProjectsTopicsSetIamPolicyCall{s: r.s, opt_: make(map[string]interface{})}
	c.resource = resource
	c.setiampolicyrequest = setiampolicyrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsTopicsSetIamPolicyCall) Fields(s ...googleapi.Field) *ProjectsTopicsSetIamPolicyCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsTopicsSetIamPolicyCall) Do() (*Policy, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+resource}:setIamPolicy")
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
	//   "description": "Sets the access control policy on the specified resource. Replaces any\nexisting policy.",
	//   "flatPath": "v1/projects/{projectsId}/topics/{topicsId}:setIamPolicy",
	//   "httpMethod": "POST",
	//   "id": "pubsub.projects.topics.setIamPolicy",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "REQUIRED: The resource for which policy is being specified.\n`resource` is usually specified as a path, such as,\n`projects/{project}/zones/{zone}/disks/{disk}`.\n\nThe format for the path specified in this value is resource specific and\nis specified in the documentation for the respective SetIamPolicy rpc.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]*/topics/[^/]*$",
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
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/pubsub"
	//   ]
	// }

}

// method id "pubsub.projects.topics.testIamPermissions":

type ProjectsTopicsTestIamPermissionsCall struct {
	s                         *Service
	resource                  string
	testiampermissionsrequest *TestIamPermissionsRequest
	opt_                      map[string]interface{}
}

// TestIamPermissions: Returns permissions that a caller has on the
// specified resource.
func (r *ProjectsTopicsService) TestIamPermissions(resource string, testiampermissionsrequest *TestIamPermissionsRequest) *ProjectsTopicsTestIamPermissionsCall {
	c := &ProjectsTopicsTestIamPermissionsCall{s: r.s, opt_: make(map[string]interface{})}
	c.resource = resource
	c.testiampermissionsrequest = testiampermissionsrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsTopicsTestIamPermissionsCall) Fields(s ...googleapi.Field) *ProjectsTopicsTestIamPermissionsCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsTopicsTestIamPermissionsCall) Do() (*TestIamPermissionsResponse, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+resource}:testIamPermissions")
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
	//   "description": "Returns permissions that a caller has on the specified resource.",
	//   "flatPath": "v1/projects/{projectsId}/topics/{topicsId}:testIamPermissions",
	//   "httpMethod": "POST",
	//   "id": "pubsub.projects.topics.testIamPermissions",
	//   "parameterOrder": [
	//     "resource"
	//   ],
	//   "parameters": {
	//     "resource": {
	//       "description": "REQUIRED: The resource for which policy detail is being requested.\n`resource` is usually specified as a path, such as,\n`projects/{project}/zones/{zone}/disks/{disk}`.\n\nThe format for the path specified in this value is resource specific and\nis specified in the documentation for the respective TestIamPermissions\nrpc.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]*/topics/[^/]*$",
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
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/pubsub"
	//   ]
	// }

}

// method id "pubsub.projects.topics.subscriptions.list":

type ProjectsTopicsSubscriptionsListCall struct {
	s     *Service
	topic string
	opt_  map[string]interface{}
}

// List: Lists the name of the subscriptions for this topic.
func (r *ProjectsTopicsSubscriptionsService) List(topic string) *ProjectsTopicsSubscriptionsListCall {
	c := &ProjectsTopicsSubscriptionsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.topic = topic
	return c
}

// PageSize sets the optional parameter "pageSize": Maximum number of
// subscription names to return.
func (c *ProjectsTopicsSubscriptionsListCall) PageSize(pageSize int64) *ProjectsTopicsSubscriptionsListCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken": The value returned
// by the last `ListTopicSubscriptionsResponse`; indicates
// that this is
// a continuation of a prior `ListTopicSubscriptions` call, and
// that the
// system should return the next page of data.
func (c *ProjectsTopicsSubscriptionsListCall) PageToken(pageToken string) *ProjectsTopicsSubscriptionsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsTopicsSubscriptionsListCall) Fields(s ...googleapi.Field) *ProjectsTopicsSubscriptionsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsTopicsSubscriptionsListCall) Do() (*ListTopicSubscriptionsResponse, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+topic}/subscriptions")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"topic": c.topic,
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
	var ret *ListTopicSubscriptionsResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists the name of the subscriptions for this topic.",
	//   "flatPath": "v1/projects/{projectsId}/topics/{topicsId}/subscriptions",
	//   "httpMethod": "GET",
	//   "id": "pubsub.projects.topics.subscriptions.list",
	//   "parameterOrder": [
	//     "topic"
	//   ],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "Maximum number of subscription names to return.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "The value returned by the last `ListTopicSubscriptionsResponse`; indicates\nthat this is a continuation of a prior `ListTopicSubscriptions` call, and\nthat the system should return the next page of data.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "topic": {
	//       "description": "The name of the topic that subscriptions are attached to.",
	//       "location": "path",
	//       "pattern": "^projects/[^/]*/topics/[^/]*$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+topic}/subscriptions",
	//   "response": {
	//     "$ref": "ListTopicSubscriptionsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/pubsub"
	//   ]
	// }

}
