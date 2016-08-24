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

	api_client "github.com/skelterjohn/gcloud_apis/clients/bigtableadmin/v2"
	"github.com/skelterjohn/gcloud_apis/commands_util"
)

// some generated code may not use all imported libs
var _ = fmt.Println
var _ = io.Copy
var _ = os.Stdin
var _ = strings.Split

func Bigtableadmin_v2_OperationsCancel(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("operationsId"))

		if len(pathParams) != 0 {
			if strings.Contains("v2/operations/{operationsId}:cancel", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewOperationsService(api_service)

	// Only positional arguments should remain in args.
	if len(args) != 1 {
		usageFunc()
	}

	expectedParams := []string{
		"operationsId",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_operationsId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}

	call := service.Cancel(param_operationsId)

	response, err := call.Do()
	if err != nil {
		return err
	}

	err = commands_util.PrintResponse(response)
	if err != nil {
		return err
	}

	return nil
}

func Bigtableadmin_v2_OperationsDelete(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("operationsId"))

		if len(pathParams) != 0 {
			if strings.Contains("v2/operations/{operationsId}", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewOperationsService(api_service)

	// Only positional arguments should remain in args.
	if len(args) != 1 {
		usageFunc()
	}

	expectedParams := []string{
		"operationsId",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_operationsId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}

	call := service.Delete(param_operationsId)

	response, err := call.Do()
	if err != nil {
		return err
	}

	err = commands_util.PrintResponse(response)
	if err != nil {
		return err
	}

	return nil
}

func Bigtableadmin_v2_OperationsGet(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("operationsId"))

		if len(pathParams) != 0 {
			if strings.Contains("v2/operations/{operationsId}", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewOperationsService(api_service)

	// Only positional arguments should remain in args.
	if len(args) != 1 {
		usageFunc()
	}

	expectedParams := []string{
		"operationsId",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_operationsId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}

	call := service.Get(param_operationsId)

	response, err := call.Do()
	if err != nil {
		return err
	}

	err = commands_util.PrintResponse(response)
	if err != nil {
		return err
	}

	return nil
}

func Bigtableadmin_v2_OperationsList(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string

		if len(pathParams) != 0 {
			if strings.Contains("v2/operations", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [--filter=VALUE]"

		usageBits += " [--pageSize=VALUE]"

		usageBits += " [--pageToken=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewOperationsService(api_service)

	queryParamNames := map[string]bool{
		"filter":    false,
		"pageSize":  false,
		"pageToken": false,
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
		usageFunc()
	}

	expectedParams := []string{}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	call := service.List()

	// Set query parameters.
	if value, ok := flagValues["filter"]; ok {
		query_filter, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.Filter(query_filter)
	}
	if value, ok := flagValues["pageSize"]; ok {
		query_pageSize, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.PageSize(query_pageSize)
	}
	if value, ok := flagValues["pageToken"]; ok {
		query_pageToken, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.PageToken(query_pageToken)
	}

	response, err := call.Do()
	if err != nil {
		return err
	}

	err = commands_util.PrintResponse(response)
	if err != nil {
		return err
	}

	return nil
}

func Bigtableadmin_v2_ProjectsInstancesClustersCreate(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectsId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("instancesId"))

		if len(pathParams) != 0 {
			if strings.Contains("v2/projects/{projectsId}/instances/{instancesId}/clusters", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		usageBits += " [--clusterId=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		commands_util.PrintRequestExample(&api_client.Cluster{})

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsInstancesClustersService(api_service)

	queryParamNames := map[string]bool{
		"clusterId": false,
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
		usageFunc()
	}

	request := &api_client.Cluster{}
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

	expectedParams := []string{
		"projectsId",
		"instancesId",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectsId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_instancesId, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}

	call := service.Create(param_projectsId, param_instancesId,
		request,
	)

	// Set query parameters.
	if value, ok := flagValues["clusterId"]; ok {
		query_clusterId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.ClusterId(query_clusterId)
	}

	response, err := call.Do()
	if err != nil {
		return err
	}

	err = commands_util.PrintResponse(response)
	if err != nil {
		return err
	}

	return nil
}

func Bigtableadmin_v2_ProjectsInstancesClustersDelete(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectsId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("instancesId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("clustersId"))

		if len(pathParams) != 0 {
			if strings.Contains("v2/projects/{projectsId}/instances/{instancesId}/clusters/{clustersId}", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsInstancesClustersService(api_service)

	// Only positional arguments should remain in args.
	if len(args) != 1 {
		usageFunc()
	}

	expectedParams := []string{
		"projectsId",
		"instancesId",
		"clustersId",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectsId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_instancesId, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}
	param_clustersId, err := commands_util.ConvertValue_string(paramValues[2])
	if err != nil {
		return err
	}

	call := service.Delete(param_projectsId, param_instancesId, param_clustersId)

	response, err := call.Do()
	if err != nil {
		return err
	}

	err = commands_util.PrintResponse(response)
	if err != nil {
		return err
	}

	return nil
}

func Bigtableadmin_v2_ProjectsInstancesClustersGet(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectsId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("instancesId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("clustersId"))

		if len(pathParams) != 0 {
			if strings.Contains("v2/projects/{projectsId}/instances/{instancesId}/clusters/{clustersId}", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsInstancesClustersService(api_service)

	// Only positional arguments should remain in args.
	if len(args) != 1 {
		usageFunc()
	}

	expectedParams := []string{
		"projectsId",
		"instancesId",
		"clustersId",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectsId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_instancesId, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}
	param_clustersId, err := commands_util.ConvertValue_string(paramValues[2])
	if err != nil {
		return err
	}

	call := service.Get(param_projectsId, param_instancesId, param_clustersId)

	response, err := call.Do()
	if err != nil {
		return err
	}

	err = commands_util.PrintResponse(response)
	if err != nil {
		return err
	}

	return nil
}

func Bigtableadmin_v2_ProjectsInstancesClustersList(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectsId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("instancesId"))

		if len(pathParams) != 0 {
			if strings.Contains("v2/projects/{projectsId}/instances/{instancesId}/clusters", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [--pageToken=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsInstancesClustersService(api_service)

	queryParamNames := map[string]bool{
		"pageToken": false,
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
		usageFunc()
	}

	expectedParams := []string{
		"projectsId",
		"instancesId",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectsId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_instancesId, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}

	call := service.List(param_projectsId, param_instancesId)

	// Set query parameters.
	if value, ok := flagValues["pageToken"]; ok {
		query_pageToken, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.PageToken(query_pageToken)
	}

	response, err := call.Do()
	if err != nil {
		return err
	}

	err = commands_util.PrintResponse(response)
	if err != nil {
		return err
	}

	return nil
}

func Bigtableadmin_v2_ProjectsInstancesClustersUpdate(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectsId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("instancesId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("clustersId"))

		if len(pathParams) != 0 {
			if strings.Contains("v2/projects/{projectsId}/instances/{instancesId}/clusters/{clustersId}", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		commands_util.PrintRequestExample(&api_client.Cluster{})

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsInstancesClustersService(api_service)

	args, flagValues, err := commands_util.ExtractFlagValues(args)
	if err != nil {
		return err
	}

	// Only positional arguments should remain in args.
	if len(args) == 0 || len(args) > 2 {
		usageFunc()
	}

	request := &api_client.Cluster{}
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
		"projectsId",
		"instancesId",
		"clustersId",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectsId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_instancesId, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}
	param_clustersId, err := commands_util.ConvertValue_string(paramValues[2])
	if err != nil {
		return err
	}

	call := service.Update(param_projectsId, param_instancesId, param_clustersId,
		request,
	)

	response, err := call.Do()
	if err != nil {
		return err
	}

	err = commands_util.PrintResponse(response)
	if err != nil {
		return err
	}

	return nil
}

func Bigtableadmin_v2_ProjectsInstancesCreate(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectsId"))

		if len(pathParams) != 0 {
			if strings.Contains("v2/projects/{projectsId}/instances", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		commands_util.PrintRequestExample(&api_client.CreateInstanceRequest{})

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsInstancesService(api_service)

	args, flagValues, err := commands_util.ExtractFlagValues(args)
	if err != nil {
		return err
	}

	// Only positional arguments should remain in args.
	if len(args) == 0 || len(args) > 2 {
		usageFunc()
	}

	request := &api_client.CreateInstanceRequest{}
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
		"projectsId",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectsId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}

	call := service.Create(param_projectsId,
		request,
	)

	response, err := call.Do()
	if err != nil {
		return err
	}

	err = commands_util.PrintResponse(response)
	if err != nil {
		return err
	}

	return nil
}

func Bigtableadmin_v2_ProjectsInstancesDelete(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectsId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("instancesId"))

		if len(pathParams) != 0 {
			if strings.Contains("v2/projects/{projectsId}/instances/{instancesId}", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsInstancesService(api_service)

	// Only positional arguments should remain in args.
	if len(args) != 1 {
		usageFunc()
	}

	expectedParams := []string{
		"projectsId",
		"instancesId",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectsId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_instancesId, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}

	call := service.Delete(param_projectsId, param_instancesId)

	response, err := call.Do()
	if err != nil {
		return err
	}

	err = commands_util.PrintResponse(response)
	if err != nil {
		return err
	}

	return nil
}

func Bigtableadmin_v2_ProjectsInstancesGet(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectsId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("instancesId"))

		if len(pathParams) != 0 {
			if strings.Contains("v2/projects/{projectsId}/instances/{instancesId}", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsInstancesService(api_service)

	// Only positional arguments should remain in args.
	if len(args) != 1 {
		usageFunc()
	}

	expectedParams := []string{
		"projectsId",
		"instancesId",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectsId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_instancesId, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}

	call := service.Get(param_projectsId, param_instancesId)

	response, err := call.Do()
	if err != nil {
		return err
	}

	err = commands_util.PrintResponse(response)
	if err != nil {
		return err
	}

	return nil
}

func Bigtableadmin_v2_ProjectsInstancesList(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectsId"))

		if len(pathParams) != 0 {
			if strings.Contains("v2/projects/{projectsId}/instances", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [--pageToken=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsInstancesService(api_service)

	queryParamNames := map[string]bool{
		"pageToken": false,
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
		usageFunc()
	}

	expectedParams := []string{
		"projectsId",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectsId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}

	call := service.List(param_projectsId)

	// Set query parameters.
	if value, ok := flagValues["pageToken"]; ok {
		query_pageToken, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.PageToken(query_pageToken)
	}

	response, err := call.Do()
	if err != nil {
		return err
	}

	err = commands_util.PrintResponse(response)
	if err != nil {
		return err
	}

	return nil
}

func Bigtableadmin_v2_ProjectsInstancesTablesCreate(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectsId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("instancesId"))

		if len(pathParams) != 0 {
			if strings.Contains("v2/projects/{projectsId}/instances/{instancesId}/tables", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		commands_util.PrintRequestExample(&api_client.CreateTableRequest{})

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsInstancesTablesService(api_service)

	args, flagValues, err := commands_util.ExtractFlagValues(args)
	if err != nil {
		return err
	}

	// Only positional arguments should remain in args.
	if len(args) == 0 || len(args) > 2 {
		usageFunc()
	}

	request := &api_client.CreateTableRequest{}
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
		"projectsId",
		"instancesId",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectsId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_instancesId, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}

	call := service.Create(param_projectsId, param_instancesId,
		request,
	)

	response, err := call.Do()
	if err != nil {
		return err
	}

	err = commands_util.PrintResponse(response)
	if err != nil {
		return err
	}

	return nil
}

func Bigtableadmin_v2_ProjectsInstancesTablesDelete(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectsId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("instancesId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("tablesId"))

		if len(pathParams) != 0 {
			if strings.Contains("v2/projects/{projectsId}/instances/{instancesId}/tables/{tablesId}", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsInstancesTablesService(api_service)

	// Only positional arguments should remain in args.
	if len(args) != 1 {
		usageFunc()
	}

	expectedParams := []string{
		"projectsId",
		"instancesId",
		"tablesId",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectsId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_instancesId, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}
	param_tablesId, err := commands_util.ConvertValue_string(paramValues[2])
	if err != nil {
		return err
	}

	call := service.Delete(param_projectsId, param_instancesId, param_tablesId)

	response, err := call.Do()
	if err != nil {
		return err
	}

	err = commands_util.PrintResponse(response)
	if err != nil {
		return err
	}

	return nil
}

func Bigtableadmin_v2_ProjectsInstancesTablesDropRowRange(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectsId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("instancesId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("tablesId"))

		if len(pathParams) != 0 {
			if strings.Contains("v2/projects/{projectsId}/instances/{instancesId}/tables/{tablesId}:dropRowRange", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		commands_util.PrintRequestExample(&api_client.DropRowRangeRequest{})

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsInstancesTablesService(api_service)

	args, flagValues, err := commands_util.ExtractFlagValues(args)
	if err != nil {
		return err
	}

	// Only positional arguments should remain in args.
	if len(args) == 0 || len(args) > 2 {
		usageFunc()
	}

	request := &api_client.DropRowRangeRequest{}
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
		"projectsId",
		"instancesId",
		"tablesId",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectsId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_instancesId, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}
	param_tablesId, err := commands_util.ConvertValue_string(paramValues[2])
	if err != nil {
		return err
	}

	call := service.DropRowRange(param_projectsId, param_instancesId, param_tablesId,
		request,
	)

	response, err := call.Do()
	if err != nil {
		return err
	}

	err = commands_util.PrintResponse(response)
	if err != nil {
		return err
	}

	return nil
}

func Bigtableadmin_v2_ProjectsInstancesTablesGet(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectsId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("instancesId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("tablesId"))

		if len(pathParams) != 0 {
			if strings.Contains("v2/projects/{projectsId}/instances/{instancesId}/tables/{tablesId}", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [--view=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsInstancesTablesService(api_service)

	queryParamNames := map[string]bool{
		"view": false,
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
		usageFunc()
	}

	expectedParams := []string{
		"projectsId",
		"instancesId",
		"tablesId",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectsId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_instancesId, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}
	param_tablesId, err := commands_util.ConvertValue_string(paramValues[2])
	if err != nil {
		return err
	}

	call := service.Get(param_projectsId, param_instancesId, param_tablesId)

	// Set query parameters.
	if value, ok := flagValues["view"]; ok {
		query_view, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.View(query_view)
	}

	response, err := call.Do()
	if err != nil {
		return err
	}

	err = commands_util.PrintResponse(response)
	if err != nil {
		return err
	}

	return nil
}

func Bigtableadmin_v2_ProjectsInstancesTablesList(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectsId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("instancesId"))

		if len(pathParams) != 0 {
			if strings.Contains("v2/projects/{projectsId}/instances/{instancesId}/tables", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [--pageToken=VALUE]"

		usageBits += " [--view=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsInstancesTablesService(api_service)

	queryParamNames := map[string]bool{
		"pageToken": false,
		"view":      false,
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
		usageFunc()
	}

	expectedParams := []string{
		"projectsId",
		"instancesId",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectsId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_instancesId, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}

	call := service.List(param_projectsId, param_instancesId)

	// Set query parameters.
	if value, ok := flagValues["pageToken"]; ok {
		query_pageToken, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.PageToken(query_pageToken)
	}
	if value, ok := flagValues["view"]; ok {
		query_view, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.View(query_view)
	}

	response, err := call.Do()
	if err != nil {
		return err
	}

	err = commands_util.PrintResponse(response)
	if err != nil {
		return err
	}

	return nil
}

func Bigtableadmin_v2_ProjectsInstancesTablesModifyColumnFamilies(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectsId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("instancesId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("tablesId"))

		if len(pathParams) != 0 {
			if strings.Contains("v2/projects/{projectsId}/instances/{instancesId}/tables/{tablesId}:modifyColumnFamilies", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		commands_util.PrintRequestExample(&api_client.ModifyColumnFamiliesRequest{})

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsInstancesTablesService(api_service)

	args, flagValues, err := commands_util.ExtractFlagValues(args)
	if err != nil {
		return err
	}

	// Only positional arguments should remain in args.
	if len(args) == 0 || len(args) > 2 {
		usageFunc()
	}

	request := &api_client.ModifyColumnFamiliesRequest{}
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
		"projectsId",
		"instancesId",
		"tablesId",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectsId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_instancesId, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}
	param_tablesId, err := commands_util.ConvertValue_string(paramValues[2])
	if err != nil {
		return err
	}

	call := service.ModifyColumnFamilies(param_projectsId, param_instancesId, param_tablesId,
		request,
	)

	response, err := call.Do()
	if err != nil {
		return err
	}

	err = commands_util.PrintResponse(response)
	if err != nil {
		return err
	}

	return nil
}

func Bigtableadmin_v2_ProjectsInstancesUpdate(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectsId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("instancesId"))

		if len(pathParams) != 0 {
			if strings.Contains("v2/projects/{projectsId}/instances/{instancesId}", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		commands_util.PrintRequestExample(&api_client.Instance{})

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsInstancesService(api_service)

	args, flagValues, err := commands_util.ExtractFlagValues(args)
	if err != nil {
		return err
	}

	// Only positional arguments should remain in args.
	if len(args) == 0 || len(args) > 2 {
		usageFunc()
	}

	request := &api_client.Instance{}
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
		"projectsId",
		"instancesId",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectsId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_instancesId, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}

	call := service.Update(param_projectsId, param_instancesId,
		request,
	)

	response, err := call.Do()
	if err != nil {
		return err
	}

	err = commands_util.PrintResponse(response)
	if err != nil {
		return err
	}

	return nil
}
