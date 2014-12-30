// generated code: api commands
/*
Copyright 2014 Google Inc. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package commands

import (
	"fmt"
	"io"
	"os"
	"strings"

	api_client "github.com/GoogleCloudPlatform/gcloud/gcloud_apis/clients/developerprojects/v1"
	"github.com/GoogleCloudPlatform/gcloud/gcloud_apis/commands_util"
)

// some generated code may not use all imported libs
var _ = fmt.Println
var _ = io.Copy
var _ = os.Stdin

func Developerprojects_v1_ProjectsCreate(context Context, args ...string) error {
	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsService(api_service)

	queryParamNames := map[string]bool{
		"appengineStorageLocation": false,
		"createAppengineProject":   false,
	}

	args, flagValues, err := commands_util.ExtractFlagValues(args)
	if err != nil {
		return err
	}

	for k, r := range queryParamNames {
		if _, ok := flagValues[k]; r && !ok {
			return fmt.Errorf("missing required flag %q", "--"+k)
		}
	}

	// Only positional arguments should remain in args.
	if len(args) == 0 || len(args) > 2 {
		commands_util.UsageForMethod(context.InvocationMethod, "POST")
	}

	request := &api_client.Project{}
	if len(args) == 2 {
		err = commands_util.PopulateRequestFromFilename(&request, args[1])
		if err != nil {
			return err
		}
	}

	// Any flags that aren't query parameters are applied to the request.
	keyValues := map[string]string{}
	for k, v := range flagValues {
		if _, ok := queryParamNames[k]; !ok {
			keyValues[k] = v
		}
	}

	err = commands_util.OverwriteRequestWithValues(&request, keyValues)
	if err != nil {
		return err
	}

	expectedParams := []string{}
	paramValues := strings.Split(args[0], "/")
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	call := service.Create(
		request,
	)

	// Set query parameters.
	if value, ok := flagValues["appengineStorageLocation"]; ok {
		query_appengineStorageLocation, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.AppengineStorageLocation(query_appengineStorageLocation)
	}
	if value, ok := flagValues["createAppengineProject"]; ok {
		query_createAppengineProject, err := commands_util.ConvertValue_bool(value)
		if err != nil {
			return err
		}
		call.CreateAppengineProject(query_createAppengineProject)
	}

	var response *api_client.Project
	response, err = call.Do()
	if err != nil {
		return err
	}

	err = commands_util.PrintResponse(response)
	if err != nil {
		return err
	}

	return nil
}

func Developerprojects_v1_ProjectsDelete(context Context, args ...string) error {
	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsService(api_service)

	// Only positional arguments should remain in args.
	if len(args) != 1 {
		commands_util.UsageForMethod(context.InvocationMethod, "DELETE")
	}

	expectedParams := []string{
		"projectId",
	}
	paramValues := strings.Split(args[0], "/")
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectId := paramValues[0]

	call := service.Delete(param_projectId)

	err = call.Do()
	if err != nil {
		return err
	}

	return nil
}

func Developerprojects_v1_ProjectsGet(context Context, args ...string) error {
	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsService(api_service)

	// Only positional arguments should remain in args.
	if len(args) != 1 {
		commands_util.UsageForMethod(context.InvocationMethod, "GET")
	}

	expectedParams := []string{
		"projectId",
	}
	paramValues := strings.Split(args[0], "/")
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectId := paramValues[0]

	call := service.Get(param_projectId)

	var response *api_client.Project
	response, err = call.Do()
	if err != nil {
		return err
	}

	err = commands_util.PrintResponse(response)
	if err != nil {
		return err
	}

	return nil
}

func Developerprojects_v1_ProjectsList(context Context, args ...string) error {
	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsService(api_service)

	queryParamNames := map[string]bool{
		"maxResults": false,
		"pageToken":  false,
		"query":      false,
	}

	args, flagValues, err := commands_util.ExtractFlagValues(args)
	if err != nil {
		return err
	}

	for k, r := range queryParamNames {
		if _, ok := flagValues[k]; r && !ok {
			return fmt.Errorf("missing required flag %q", "--"+k)
		}
	}

	// Only positional arguments should remain in args.
	if len(args) != 1 {
		commands_util.UsageForMethod(context.InvocationMethod, "GET")
	}

	expectedParams := []string{}
	paramValues := strings.Split(args[0], "/")
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	call := service.List()

	// Set query parameters.
	if value, ok := flagValues["maxResults"]; ok {
		query_maxResults, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.MaxResults(query_maxResults)
	}
	if value, ok := flagValues["pageToken"]; ok {
		query_pageToken, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.PageToken(query_pageToken)
	}
	if value, ok := flagValues["query"]; ok {
		query_query, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.Query(query_query)
	}

	var response *api_client.ListProjectsResponse
	response, err = call.Do()
	if err != nil {
		return err
	}

	err = commands_util.PrintResponse(response)
	if err != nil {
		return err
	}

	return nil
}

func Developerprojects_v1_ProjectsPatch(context Context, args ...string) error {
	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsService(api_service)

	args, flagValues, err := commands_util.ExtractFlagValues(args)
	if err != nil {
		return err
	}

	// Only positional arguments should remain in args.
	if len(args) == 0 || len(args) > 2 {
		commands_util.UsageForMethod(context.InvocationMethod, "PATCH")
	}

	request := &api_client.Project{}
	if len(args) == 2 {
		err = commands_util.PopulateRequestFromFilename(&request, args[1])
		if err != nil {
			return err
		}
	}

	keyValues := flagValues

	err = commands_util.OverwriteRequestWithValues(&request, keyValues)
	if err != nil {
		return err
	}

	expectedParams := []string{
		"projectId",
	}
	paramValues := strings.Split(args[0], "/")
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectId := paramValues[0]

	call := service.Patch(param_projectId,
		request,
	)

	var response *api_client.Project
	response, err = call.Do()
	if err != nil {
		return err
	}

	err = commands_util.PrintResponse(response)
	if err != nil {
		return err
	}

	return nil
}

func Developerprojects_v1_ProjectsUndelete(context Context, args ...string) error {
	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsService(api_service)

	// Only positional arguments should remain in args.
	if len(args) != 1 {
		commands_util.UsageForMethod(context.InvocationMethod, "POST")
	}

	expectedParams := []string{
		"projectId",
	}
	paramValues := strings.Split(args[0], "/")
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectId := paramValues[0]

	call := service.Undelete(param_projectId)

	err = call.Do()
	if err != nil {
		return err
	}

	return nil
}

func Developerprojects_v1_ProjectsUpdate(context Context, args ...string) error {
	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsService(api_service)

	args, flagValues, err := commands_util.ExtractFlagValues(args)
	if err != nil {
		return err
	}

	// Only positional arguments should remain in args.
	if len(args) == 0 || len(args) > 2 {
		commands_util.UsageForMethod(context.InvocationMethod, "PUT")
	}

	request := &api_client.Project{}
	if len(args) == 2 {
		err = commands_util.PopulateRequestFromFilename(&request, args[1])
		if err != nil {
			return err
		}
	}

	keyValues := flagValues

	err = commands_util.OverwriteRequestWithValues(&request, keyValues)
	if err != nil {
		return err
	}

	expectedParams := []string{
		"projectId",
	}
	paramValues := strings.Split(args[0], "/")
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectId := paramValues[0]

	call := service.Update(param_projectId,
		request,
	)

	var response *api_client.Project
	response, err = call.Do()
	if err != nil {
		return err
	}

	err = commands_util.PrintResponse(response)
	if err != nil {
		return err
	}

	return nil
}
