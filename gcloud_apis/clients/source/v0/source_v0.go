// Package source provides access to the .
//
// Usage example:
//
//   import "code.google.com/p/google-api-go-client/source/v0"
//   ...
//   sourceService, err := source.New(oauthHttpClient)
package source

import (
	"bytes"
	"code.google.com/p/google-api-go-client/googleapi"
	"encoding/json"
	"errors"
	"fmt"
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

const apiId = "source:v0"
const apiName = "source"
const apiVersion = "v0"
const basePath = "https://www.googleapis.com/source/v0/"

// OAuth2 scopes used by this API.
const (
	// Manage your user profile and projects on Project Hosting
	ProjecthostingScope = "https://www.googleapis.com/auth/projecthosting"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Repos = NewReposService(s)
	return s, nil
}

type Service struct {
	client   *http.Client
	BasePath string // API endpoint base URL

	Repos *ReposService
}

func NewReposService(s *Service) *ReposService {
	rs := &ReposService{s: s}
	return rs
}

type ReposService struct {
	s *Service
}

type ListReposResponse struct {
	// Repos: Repos belonging to the project.
	Repos []*Repo `json:"repos,omitempty"`
}

type Repo struct {
	// CloneUrl: URL where this repo can be cloned.
	CloneUrl string `json:"cloneUrl,omitempty"`

	// RepoId: ID of the Repo.
	RepoId string `json:"repoId,omitempty"`

	// RepoName: User-defined name of the repo (or 'default')
	RepoName string `json:"repoName,omitempty"`
}

// method id "source.repos.list":

type ReposListCall struct {
	s         *Service
	projectId string
	opt_      map[string]interface{}
}

// List: List Repos belonging to a project.
func (r *ReposService) List(projectId string) *ReposListCall {
	c := &ReposListCall{s: r.s, opt_: make(map[string]interface{})}
	c.projectId = projectId
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ReposListCall) Fields(s ...googleapi.Field) *ReposListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *ReposListCall) Do() (*ListReposResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "{projectId}")
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
	//   "description": "List Repos belonging to a project.",
	//   "httpMethod": "GET",
	//   "id": "source.repos.list",
	//   "parameterOrder": [
	//     "projectId"
	//   ],
	//   "parameters": {
	//     "projectId": {
	//       "description": "ID of the project for which to list Repos. Examples: user-chosen-project-id, yellow-banana-33, dyspeptic-wombat-87",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "{projectId}",
	//   "response": {
	//     "$ref": "ListReposResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/projecthosting"
	//   ]
	// }

}
