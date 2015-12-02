// Package bigtableclusteradmin provides access to the Google Cloud Bigtable Cluster Admin API.
//
// Usage example:
//
//   import "google.golang.org/api/bigtableclusteradmin/v1"
//   ...
//   bigtableclusteradminService, err := bigtableclusteradmin.New(oauthHttpClient)
package bigtableclusteradmin

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

const apiId = "bigtableclusteradmin:v1"
const apiName = "bigtableclusteradmin"
const apiVersion = "v1"
const basePath = "https://bigtableclusteradmin.googleapis.com/v1/"

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
	rs.Aggregated = NewProjectsAggregatedService(s)
	rs.Zones = NewProjectsZonesService(s)
	return rs
}

type ProjectsService struct {
	s *Service

	Aggregated *ProjectsAggregatedService

	Zones *ProjectsZonesService
}

func NewProjectsAggregatedService(s *Service) *ProjectsAggregatedService {
	rs := &ProjectsAggregatedService{s: s}
	rs.Clusters = NewProjectsAggregatedClustersService(s)
	return rs
}

type ProjectsAggregatedService struct {
	s *Service

	Clusters *ProjectsAggregatedClustersService
}

func NewProjectsAggregatedClustersService(s *Service) *ProjectsAggregatedClustersService {
	rs := &ProjectsAggregatedClustersService{s: s}
	return rs
}

type ProjectsAggregatedClustersService struct {
	s *Service
}

func NewProjectsZonesService(s *Service) *ProjectsZonesService {
	rs := &ProjectsZonesService{s: s}
	rs.Clusters = NewProjectsZonesClustersService(s)
	return rs
}

type ProjectsZonesService struct {
	s *Service

	Clusters *ProjectsZonesClustersService
}

func NewProjectsZonesClustersService(s *Service) *ProjectsZonesClustersService {
	rs := &ProjectsZonesClustersService{s: s}
	return rs
}

type ProjectsZonesClustersService struct {
	s *Service
}

type CancelOperationRequest struct {
}

type Cluster struct {
	CurrentOperation *Operation `json:"currentOperation,omitempty"`

	DefaultStorageType string `json:"defaultStorageType,omitempty"`

	DisplayName string `json:"displayName,omitempty"`

	HddBytes int64 `json:"hddBytes,omitempty,string"`

	Name string `json:"name,omitempty"`

	ServeNodes int64 `json:"serveNodes,omitempty"`

	SsdBytes int64 `json:"ssdBytes,omitempty,string"`
}

type CreateClusterRequest struct {
	Cluster *Cluster `json:"cluster,omitempty"`

	ClusterId string `json:"clusterId,omitempty"`

	Name string `json:"name,omitempty"`
}

type Empty struct {
}

type ListClustersResponse struct {
	Clusters []*Cluster `json:"clusters,omitempty"`

	FailedZones []*Zone `json:"failedZones,omitempty"`
}

type ListOperationsResponse struct {
	NextPageToken string `json:"nextPageToken,omitempty"`

	Operations []*Operation `json:"operations,omitempty"`
}

type ListZonesResponse struct {
	Zones []*Zone `json:"zones,omitempty"`
}

type Operation struct {
	Done bool `json:"done,omitempty"`

	Error *Status `json:"error,omitempty"`

	Metadata *OperationMetadata `json:"metadata,omitempty"`

	Name string `json:"name,omitempty"`

	Response *OperationResponse `json:"response,omitempty"`
}

type OperationMetadata struct {
}

type OperationResponse struct {
}

type Status struct {
	Code int64 `json:"code,omitempty"`

	Details []*StatusDetails `json:"details,omitempty"`

	Message string `json:"message,omitempty"`
}

type StatusDetails struct {
}

type Zone struct {
	DisplayName string `json:"displayName,omitempty"`

	Name string `json:"name,omitempty"`

	Status string `json:"status,omitempty"`
}

// method id "bigtableclusteradmin.operations.cancel":

type OperationsCancelCall struct {
	s                      *Service
	name                   string
	canceloperationrequest *CancelOperationRequest
	opt_                   map[string]interface{}
}

// Cancel: Starts asynchronous cancellation on a long-running operation.
// The server makes a best effort to cancel the operation, but success
// is not guaranteed. If the server doesn't support this method, it
// returns `google.rpc.Code.UNIMPLEMENTED`. Clients can use
// [Operations.GetOperation][google.longrunning.Operations.GetOperation]
// or other methods to check whether the cancellation succeeded or
// whether the operation completed despite cancellation.
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "{+name}:cancel")
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
	//   "description": "Starts asynchronous cancellation on a long-running operation. The server makes a best effort to cancel the operation, but success is not guaranteed. If the server doesn't support this method, it returns `google.rpc.Code.UNIMPLEMENTED`. Clients can use [Operations.GetOperation][google.longrunning.Operations.GetOperation] or other methods to check whether the cancellation succeeded or whether the operation completed despite cancellation.",
	//   "httpMethod": "POST",
	//   "id": "bigtableclusteradmin.operations.cancel",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{+name}:cancel",
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

// method id "bigtableclusteradmin.operations.delete":

type OperationsDeleteCall struct {
	s    *Service
	name string
	opt_ map[string]interface{}
}

// Delete: Deletes a long-running operation. This method indicates that
// the client is no longer interested in the operation result. It does
// not cancel the operation. If the server doesn't support this method,
// it returns `google.rpc.Code.UNIMPLEMENTED`.
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "{+name}")
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
	//   "description": "Deletes a long-running operation. This method indicates that the client is no longer interested in the operation result. It does not cancel the operation. If the server doesn't support this method, it returns `google.rpc.Code.UNIMPLEMENTED`.",
	//   "httpMethod": "DELETE",
	//   "id": "bigtableclusteradmin.operations.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{+name}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "bigtableclusteradmin.operations.get":

type OperationsGetCall struct {
	s    *Service
	name string
	opt_ map[string]interface{}
}

// Get: Gets the latest state of a long-running operation. Clients can
// use this method to poll the operation result at intervals as
// recommended by the API service.
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "{+name}")
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
	//   "description": "Gets the latest state of a long-running operation. Clients can use this method to poll the operation result at intervals as recommended by the API service.",
	//   "httpMethod": "GET",
	//   "id": "bigtableclusteradmin.operations.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{+name}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "bigtableclusteradmin.operations.list":

type OperationsListCall struct {
	s    *Service
	name string
	opt_ map[string]interface{}
}

// List: Lists operations that match the specified filter in the
// request. If the server doesn't support this method, it returns
// `UNIMPLEMENTED`. NOTE: the `name` binding below allows API services
// to override the binding to use different resource name schemes, such
// as `users/*/operations`.
func (r *OperationsService) List(name string) *OperationsListCall {
	c := &OperationsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.name = name
	return c
}

// Filter sets the optional parameter "filter":
func (c *OperationsListCall) Filter(filter string) *OperationsListCall {
	c.opt_["filter"] = filter
	return c
}

// PageSize sets the optional parameter "pageSize":
func (c *OperationsListCall) PageSize(pageSize int64) *OperationsListCall {
	c.opt_["pageSize"] = pageSize
	return c
}

// PageToken sets the optional parameter "pageToken":
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "{name}")
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
	//   "description": "Lists operations that match the specified filter in the request. If the server doesn't support this method, it returns `UNIMPLEMENTED`. NOTE: the `name` binding below allows API services to override the binding to use different resource name schemes, such as `users/*/operations`.",
	//   "httpMethod": "GET",
	//   "id": "bigtableclusteradmin.operations.list",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "filter": {
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "name": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "{name}",
	//   "response": {
	//     "$ref": "ListOperationsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "bigtableclusteradmin.projects.aggregated.clusters.list":

type ProjectsAggregatedClustersListCall struct {
	s    *Service
	name string
	opt_ map[string]interface{}
}

// List: Lists all clusters in the given project, along with any zones
// for which cluster information could not be retrieved.
func (r *ProjectsAggregatedClustersService) List(name string) *ProjectsAggregatedClustersListCall {
	c := &ProjectsAggregatedClustersListCall{s: r.s, opt_: make(map[string]interface{})}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsAggregatedClustersListCall) Fields(s ...googleapi.Field) *ProjectsAggregatedClustersListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsAggregatedClustersListCall) Do() (*ListClustersResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{+name}/aggregated/clusters")
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
	var ret *ListClustersResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists all clusters in the given project, along with any zones for which cluster information could not be retrieved.",
	//   "httpMethod": "GET",
	//   "id": "bigtableclusteradmin.projects.aggregated.clusters.list",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{+name}/aggregated/clusters",
	//   "response": {
	//     "$ref": "ListClustersResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "bigtableclusteradmin.projects.zones.list":

type ProjectsZonesListCall struct {
	s    *Service
	name string
	opt_ map[string]interface{}
}

// List: Lists the supported zones for the given project.
func (r *ProjectsZonesService) List(name string) *ProjectsZonesListCall {
	c := &ProjectsZonesListCall{s: r.s, opt_: make(map[string]interface{})}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsZonesListCall) Fields(s ...googleapi.Field) *ProjectsZonesListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsZonesListCall) Do() (*ListZonesResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{+name}/zones")
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
	var ret *ListZonesResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists the supported zones for the given project.",
	//   "httpMethod": "GET",
	//   "id": "bigtableclusteradmin.projects.zones.list",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{+name}/zones",
	//   "response": {
	//     "$ref": "ListZonesResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "bigtableclusteradmin.projects.zones.clusters.create":

type ProjectsZonesClustersCreateCall struct {
	s                    *Service
	name                 string
	createclusterrequest *CreateClusterRequest
	opt_                 map[string]interface{}
}

// Create: Creates a cluster and begins preparing it to begin serving.
// The returned cluster embeds as its "current_operation" a long-running
// operation which can be used to track the progress of turning up the
// new cluster. Immediately upon completion of this request: * The
// cluster will be readable via the API, with all requested attributes
// but no allocated resources. Until completion of the embedded
// operation: * Cancelling the operation will render the cluster
// immediately unreadable via the API. * All other attempts to modify or
// delete the cluster will be rejected. Upon completion of the embedded
// operation: * Billing for all successfully-allocated resources will
// begin (some types may have lower than the requested levels). * New
// tables can be created in the cluster. * The cluster's allocated
// resource levels will be readable via the API. The embedded
// operation's "metadata" field type is
// [CreateClusterMetadata][google.bigtable.admin.cluster.v1.CreateCluster
// Metadata] The embedded operation's "response" field type is
// [Cluster][google.bigtable.admin.cluster.v1.Cluster], if successful.
func (r *ProjectsZonesClustersService) Create(name string, createclusterrequest *CreateClusterRequest) *ProjectsZonesClustersCreateCall {
	c := &ProjectsZonesClustersCreateCall{s: r.s, opt_: make(map[string]interface{})}
	c.name = name
	c.createclusterrequest = createclusterrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsZonesClustersCreateCall) Fields(s ...googleapi.Field) *ProjectsZonesClustersCreateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsZonesClustersCreateCall) Do() (*Cluster, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.createclusterrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{+name}/clusters")
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
	var ret *Cluster
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a cluster and begins preparing it to begin serving. The returned cluster embeds as its \"current_operation\" a long-running operation which can be used to track the progress of turning up the new cluster. Immediately upon completion of this request: * The cluster will be readable via the API, with all requested attributes but no allocated resources. Until completion of the embedded operation: * Cancelling the operation will render the cluster immediately unreadable via the API. * All other attempts to modify or delete the cluster will be rejected. Upon completion of the embedded operation: * Billing for all successfully-allocated resources will begin (some types may have lower than the requested levels). * New tables can be created in the cluster. * The cluster's allocated resource levels will be readable via the API. The embedded operation's \"metadata\" field type is [CreateClusterMetadata][google.bigtable.admin.cluster.v1.CreateClusterMetadata] The embedded operation's \"response\" field type is [Cluster][google.bigtable.admin.cluster.v1.Cluster], if successful.",
	//   "httpMethod": "POST",
	//   "id": "bigtableclusteradmin.projects.zones.clusters.create",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{+name}/clusters",
	//   "request": {
	//     "$ref": "CreateClusterRequest"
	//   },
	//   "response": {
	//     "$ref": "Cluster"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "bigtableclusteradmin.projects.zones.clusters.delete":

type ProjectsZonesClustersDeleteCall struct {
	s    *Service
	name string
	opt_ map[string]interface{}
}

// Delete: Marks a cluster and all of its tables for permanent deletion
// in 7 days. Immediately upon completion of the request: * Billing will
// cease for all of the cluster's reserved resources. * The cluster's
// "delete_time" field will be set 7 days in the future. Soon afterward:
// * All tables within the cluster will become unavailable. Prior to the
// cluster's "delete_time": * The cluster can be recovered with a call
// to UndeleteCluster. * All other attempts to modify or delete the
// cluster will be rejected. At the cluster's "delete_time": * The
// cluster and *all of its tables* will immediately and irrevocably
// disappear from the API, and their data will be permanently deleted.
func (r *ProjectsZonesClustersService) Delete(name string) *ProjectsZonesClustersDeleteCall {
	c := &ProjectsZonesClustersDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsZonesClustersDeleteCall) Fields(s ...googleapi.Field) *ProjectsZonesClustersDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsZonesClustersDeleteCall) Do() (*Empty, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{+name}")
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
	//   "description": "Marks a cluster and all of its tables for permanent deletion in 7 days. Immediately upon completion of the request: * Billing will cease for all of the cluster's reserved resources. * The cluster's \"delete_time\" field will be set 7 days in the future. Soon afterward: * All tables within the cluster will become unavailable. Prior to the cluster's \"delete_time\": * The cluster can be recovered with a call to UndeleteCluster. * All other attempts to modify or delete the cluster will be rejected. At the cluster's \"delete_time\": * The cluster and *all of its tables* will immediately and irrevocably disappear from the API, and their data will be permanently deleted.",
	//   "httpMethod": "DELETE",
	//   "id": "bigtableclusteradmin.projects.zones.clusters.delete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{+name}",
	//   "response": {
	//     "$ref": "Empty"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "bigtableclusteradmin.projects.zones.clusters.get":

type ProjectsZonesClustersGetCall struct {
	s    *Service
	name string
	opt_ map[string]interface{}
}

// Get: Gets information about a particular cluster.
func (r *ProjectsZonesClustersService) Get(name string) *ProjectsZonesClustersGetCall {
	c := &ProjectsZonesClustersGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsZonesClustersGetCall) Fields(s ...googleapi.Field) *ProjectsZonesClustersGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsZonesClustersGetCall) Do() (*Cluster, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{+name}")
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
	var ret *Cluster
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets information about a particular cluster.",
	//   "httpMethod": "GET",
	//   "id": "bigtableclusteradmin.projects.zones.clusters.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{+name}",
	//   "response": {
	//     "$ref": "Cluster"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/cloud-platform.read-only"
	//   ]
	// }

}

// method id "bigtableclusteradmin.projects.zones.clusters.patch":

type ProjectsZonesClustersPatchCall struct {
	s       *Service
	name    string
	cluster *Cluster
	opt_    map[string]interface{}
}

// Patch: Updates a cluster, and begins allocating or releasing
// resources as requested. The returned cluster embeds as its
// "current_operation" a long-running operation which can be used to
// track the progress of updating the cluster. Immediately upon
// completion of this request: * For resource types where a decrease in
// the cluster's allocation has been requested, billing will be based on
// the newly-requested level. Until completion of the embedded
// operation: * Cancelling the operation will set its metadata's
// "cancelled_at_time", and begin restoring resources to their
// pre-request values. The operation is guaranteed to succeed at undoing
// all resource changes, after which point it will terminate with a
// CANCELLED status. * All other attempts to modify or delete the
// cluster will be rejected. * Reading the cluster via the API will
// continue to give the pre-request resource levels. Upon completion of
// the embedded operation: * Billing will begin for all
// successfully-allocated resources (some types may have lower than the
// requested levels). * All newly-reserved resources will be available
// for serving the cluster's tables. * The cluster's new resource levels
// will be readable via the API.
// [UpdateClusterMetadata][google.bigtable.admin.cluster.v1.UpdateCluster
// Metadata] The embedded operation's "response" field type is
// [Cluster][google.bigtable.admin.cluster.v1.Cluster], if successful.
// This method supports patch semantics.
func (r *ProjectsZonesClustersService) Patch(name string, cluster *Cluster) *ProjectsZonesClustersPatchCall {
	c := &ProjectsZonesClustersPatchCall{s: r.s, opt_: make(map[string]interface{})}
	c.name = name
	c.cluster = cluster
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsZonesClustersPatchCall) Fields(s ...googleapi.Field) *ProjectsZonesClustersPatchCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsZonesClustersPatchCall) Do() (*Cluster, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "{+name}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
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
	var ret *Cluster
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates a cluster, and begins allocating or releasing resources as requested. The returned cluster embeds as its \"current_operation\" a long-running operation which can be used to track the progress of updating the cluster. Immediately upon completion of this request: * For resource types where a decrease in the cluster's allocation has been requested, billing will be based on the newly-requested level. Until completion of the embedded operation: * Cancelling the operation will set its metadata's \"cancelled_at_time\", and begin restoring resources to their pre-request values. The operation is guaranteed to succeed at undoing all resource changes, after which point it will terminate with a CANCELLED status. * All other attempts to modify or delete the cluster will be rejected. * Reading the cluster via the API will continue to give the pre-request resource levels. Upon completion of the embedded operation: * Billing will begin for all successfully-allocated resources (some types may have lower than the requested levels). * All newly-reserved resources will be available for serving the cluster's tables. * The cluster's new resource levels will be readable via the API. [UpdateClusterMetadata][google.bigtable.admin.cluster.v1.UpdateClusterMetadata] The embedded operation's \"response\" field type is [Cluster][google.bigtable.admin.cluster.v1.Cluster], if successful. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "bigtableclusteradmin.projects.zones.clusters.patch",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{+name}",
	//   "request": {
	//     "$ref": "Cluster"
	//   },
	//   "response": {
	//     "$ref": "Cluster"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "bigtableclusteradmin.projects.zones.clusters.undelete":

type ProjectsZonesClustersUndeleteCall struct {
	s    *Service
	name string
	opt_ map[string]interface{}
}

// Undelete: Cancels the scheduled deletion of an cluster and begins
// preparing it to resume serving. The returned operation will also be
// embedded as the cluster's "current_operation". Immediately upon
// completion of this request: * The cluster's "delete_time" field will
// be unset, protecting it from automatic deletion. Until completion of
// the returned operation: * The operation cannot be cancelled. Upon
// completion of the returned operation: * Billing for the cluster's
// resources will resume. * All tables within the cluster will be
// available.
// [UndeleteClusterMetadata][google.bigtable.admin.cluster.v1.UndeleteClu
// sterMetadata] The embedded operation's "response" field type is
// [Cluster][google.bigtable.admin.cluster.v1.Cluster], if successful.
func (r *ProjectsZonesClustersService) Undelete(name string) *ProjectsZonesClustersUndeleteCall {
	c := &ProjectsZonesClustersUndeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsZonesClustersUndeleteCall) Fields(s ...googleapi.Field) *ProjectsZonesClustersUndeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsZonesClustersUndeleteCall) Do() (*Operation, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{+name}:undelete")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
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
	//   "description": "Cancels the scheduled deletion of an cluster and begins preparing it to resume serving. The returned operation will also be embedded as the cluster's \"current_operation\". Immediately upon completion of this request: * The cluster's \"delete_time\" field will be unset, protecting it from automatic deletion. Until completion of the returned operation: * The operation cannot be cancelled. Upon completion of the returned operation: * Billing for the cluster's resources will resume. * All tables within the cluster will be available. [UndeleteClusterMetadata][google.bigtable.admin.cluster.v1.UndeleteClusterMetadata] The embedded operation's \"response\" field type is [Cluster][google.bigtable.admin.cluster.v1.Cluster], if successful.",
	//   "httpMethod": "POST",
	//   "id": "bigtableclusteradmin.projects.zones.clusters.undelete",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{+name}:undelete",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "bigtableclusteradmin.projects.zones.clusters.update":

type ProjectsZonesClustersUpdateCall struct {
	s       *Service
	name    string
	cluster *Cluster
	opt_    map[string]interface{}
}

// Update: Updates a cluster, and begins allocating or releasing
// resources as requested. The returned cluster embeds as its
// "current_operation" a long-running operation which can be used to
// track the progress of updating the cluster. Immediately upon
// completion of this request: * For resource types where a decrease in
// the cluster's allocation has been requested, billing will be based on
// the newly-requested level. Until completion of the embedded
// operation: * Cancelling the operation will set its metadata's
// "cancelled_at_time", and begin restoring resources to their
// pre-request values. The operation is guaranteed to succeed at undoing
// all resource changes, after which point it will terminate with a
// CANCELLED status. * All other attempts to modify or delete the
// cluster will be rejected. * Reading the cluster via the API will
// continue to give the pre-request resource levels. Upon completion of
// the embedded operation: * Billing will begin for all
// successfully-allocated resources (some types may have lower than the
// requested levels). * All newly-reserved resources will be available
// for serving the cluster's tables. * The cluster's new resource levels
// will be readable via the API.
// [UpdateClusterMetadata][google.bigtable.admin.cluster.v1.UpdateCluster
// Metadata] The embedded operation's "response" field type is
// [Cluster][google.bigtable.admin.cluster.v1.Cluster], if successful.
func (r *ProjectsZonesClustersService) Update(name string, cluster *Cluster) *ProjectsZonesClustersUpdateCall {
	c := &ProjectsZonesClustersUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.name = name
	c.cluster = cluster
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ProjectsZonesClustersUpdateCall) Fields(s ...googleapi.Field) *ProjectsZonesClustersUpdateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ProjectsZonesClustersUpdateCall) Do() (*Cluster, error) {
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
	urls := googleapi.ResolveRelative(c.s.BasePath, "{+name}")
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
	var ret *Cluster
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates a cluster, and begins allocating or releasing resources as requested. The returned cluster embeds as its \"current_operation\" a long-running operation which can be used to track the progress of updating the cluster. Immediately upon completion of this request: * For resource types where a decrease in the cluster's allocation has been requested, billing will be based on the newly-requested level. Until completion of the embedded operation: * Cancelling the operation will set its metadata's \"cancelled_at_time\", and begin restoring resources to their pre-request values. The operation is guaranteed to succeed at undoing all resource changes, after which point it will terminate with a CANCELLED status. * All other attempts to modify or delete the cluster will be rejected. * Reading the cluster via the API will continue to give the pre-request resource levels. Upon completion of the embedded operation: * Billing will begin for all successfully-allocated resources (some types may have lower than the requested levels). * All newly-reserved resources will be available for serving the cluster's tables. * The cluster's new resource levels will be readable via the API. [UpdateClusterMetadata][google.bigtable.admin.cluster.v1.UpdateClusterMetadata] The embedded operation's \"response\" field type is [Cluster][google.bigtable.admin.cluster.v1.Cluster], if successful.",
	//   "httpMethod": "PUT",
	//   "id": "bigtableclusteradmin.projects.zones.clusters.update",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{+name}",
	//   "request": {
	//     "$ref": "Cluster"
	//   },
	//   "response": {
	//     "$ref": "Cluster"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}
