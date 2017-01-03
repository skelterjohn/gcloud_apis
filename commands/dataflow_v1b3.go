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

	api_client "github.com/skelterjohn/gcloud_apis/clients/dataflow/v1b3"
	"github.com/skelterjohn/gcloud_apis/commands_util"
)

// some generated code may not use all imported libs
var _ = fmt.Println
var _ = io.Copy
var _ = os.Stdin
var _ = strings.Split

func Dataflow_v1b3_ProjectsJobsCreate(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectId"))

		if len(pathParams) != 0 {
			if strings.Contains("v1b3/projects/{projectId}/jobs", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		usageBits += " [--location=VALUE]"

		usageBits += " [--replaceJobId=VALUE]"

		usageBits += " [--view=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		commands_util.PrintRequestExample(&api_client.Job{})

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsJobsService(api_service)

	queryParamNames := map[string]bool{
		"location":     false,
		"replaceJobId": false,
		"view":         false,
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

	request := &api_client.Job{}
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
		"projectId",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}

	call := service.Create(param_projectId,
		request,
	)

	// Set query parameters.
	if value, ok := flagValues["location"]; ok {
		query_location, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.Location(query_location)
	}
	if value, ok := flagValues["replaceJobId"]; ok {
		query_replaceJobId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.ReplaceJobId(query_replaceJobId)
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

func Dataflow_v1b3_ProjectsJobsDebugGetConfig(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("jobId"))

		if len(pathParams) != 0 {
			if strings.Contains("v1b3/projects/{projectId}/jobs/{jobId}/debug/getConfig", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		commands_util.PrintRequestExample(&api_client.GetDebugConfigRequest{})

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsJobsDebugService(api_service)

	args, flagValues, err := commands_util.ExtractFlagValues(args)
	if err != nil {
		return err
	}

	// Only positional arguments should remain in args.
	if len(args) == 0 || len(args) > 2 {
		usageFunc()
	}

	request := &api_client.GetDebugConfigRequest{}
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
		"jobId",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_jobId, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}

	call := service.GetConfig(param_projectId, param_jobId,
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

func Dataflow_v1b3_ProjectsJobsDebugSendCapture(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("jobId"))

		if len(pathParams) != 0 {
			if strings.Contains("v1b3/projects/{projectId}/jobs/{jobId}/debug/sendCapture", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		commands_util.PrintRequestExample(&api_client.SendDebugCaptureRequest{})

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsJobsDebugService(api_service)

	args, flagValues, err := commands_util.ExtractFlagValues(args)
	if err != nil {
		return err
	}

	// Only positional arguments should remain in args.
	if len(args) == 0 || len(args) > 2 {
		usageFunc()
	}

	request := &api_client.SendDebugCaptureRequest{}
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
		"jobId",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_jobId, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}

	call := service.SendCapture(param_projectId, param_jobId,
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

func Dataflow_v1b3_ProjectsJobsGet(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("jobId"))

		if len(pathParams) != 0 {
			if strings.Contains("v1b3/projects/{projectId}/jobs/{jobId}", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [--location=VALUE]"

		usageBits += " [--view=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsJobsService(api_service)

	queryParamNames := map[string]bool{
		"location": false,
		"view":     false,
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
		"projectId",
		"jobId",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_jobId, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}

	call := service.Get(param_projectId, param_jobId)

	// Set query parameters.
	if value, ok := flagValues["location"]; ok {
		query_location, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.Location(query_location)
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

func Dataflow_v1b3_ProjectsJobsGetMetrics(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("jobId"))

		if len(pathParams) != 0 {
			if strings.Contains("v1b3/projects/{projectId}/jobs/{jobId}/metrics", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [--location=VALUE]"

		usageBits += " [--startTime=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsJobsService(api_service)

	queryParamNames := map[string]bool{
		"location":  false,
		"startTime": false,
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
		"projectId",
		"jobId",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_jobId, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}

	call := service.GetMetrics(param_projectId, param_jobId)

	// Set query parameters.
	if value, ok := flagValues["location"]; ok {
		query_location, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.Location(query_location)
	}
	if value, ok := flagValues["startTime"]; ok {
		query_startTime, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.StartTime(query_startTime)
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

func Dataflow_v1b3_ProjectsJobsList(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectId"))

		if len(pathParams) != 0 {
			if strings.Contains("v1b3/projects/{projectId}/jobs", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [--filter=VALUE]"

		usageBits += " [--location=VALUE]"

		usageBits += " [--pageSize=VALUE]"

		usageBits += " [--pageToken=VALUE]"

		usageBits += " [--view=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsJobsService(api_service)

	queryParamNames := map[string]bool{
		"filter":    false,
		"location":  false,
		"pageSize":  false,
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
		"projectId",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}

	call := service.List(param_projectId)

	// Set query parameters.
	if value, ok := flagValues["filter"]; ok {
		query_filter, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.Filter(query_filter)
	}
	if value, ok := flagValues["location"]; ok {
		query_location, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.Location(query_location)
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

func Dataflow_v1b3_ProjectsJobsMessagesList(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("jobId"))

		if len(pathParams) != 0 {
			if strings.Contains("v1b3/projects/{projectId}/jobs/{jobId}/messages", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [--endTime=VALUE]"

		usageBits += " [--location=VALUE]"

		usageBits += " [--minimumImportance=VALUE]"

		usageBits += " [--pageSize=VALUE]"

		usageBits += " [--pageToken=VALUE]"

		usageBits += " [--startTime=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsJobsMessagesService(api_service)

	queryParamNames := map[string]bool{
		"endTime":           false,
		"location":          false,
		"minimumImportance": false,
		"pageSize":          false,
		"pageToken":         false,
		"startTime":         false,
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
		"projectId",
		"jobId",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_jobId, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}

	call := service.List(param_projectId, param_jobId)

	// Set query parameters.
	if value, ok := flagValues["endTime"]; ok {
		query_endTime, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.EndTime(query_endTime)
	}
	if value, ok := flagValues["location"]; ok {
		query_location, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.Location(query_location)
	}
	if value, ok := flagValues["minimumImportance"]; ok {
		query_minimumImportance, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.MinimumImportance(query_minimumImportance)
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
	if value, ok := flagValues["startTime"]; ok {
		query_startTime, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.StartTime(query_startTime)
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

func Dataflow_v1b3_ProjectsJobsUpdate(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("jobId"))

		if len(pathParams) != 0 {
			if strings.Contains("v1b3/projects/{projectId}/jobs/{jobId}", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		usageBits += " [--location=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		commands_util.PrintRequestExample(&api_client.Job{})

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsJobsService(api_service)

	queryParamNames := map[string]bool{
		"location": false,
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

	request := &api_client.Job{}
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
		"projectId",
		"jobId",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_jobId, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}

	call := service.Update(param_projectId, param_jobId,
		request,
	)

	// Set query parameters.
	if value, ok := flagValues["location"]; ok {
		query_location, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.Location(query_location)
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

func Dataflow_v1b3_ProjectsJobsWorkItemsLease(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("jobId"))

		if len(pathParams) != 0 {
			if strings.Contains("v1b3/projects/{projectId}/jobs/{jobId}/workItems:lease", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		commands_util.PrintRequestExample(&api_client.LeaseWorkItemRequest{})

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsJobsWorkItemsService(api_service)

	args, flagValues, err := commands_util.ExtractFlagValues(args)
	if err != nil {
		return err
	}

	// Only positional arguments should remain in args.
	if len(args) == 0 || len(args) > 2 {
		usageFunc()
	}

	request := &api_client.LeaseWorkItemRequest{}
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
		"jobId",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_jobId, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}

	call := service.Lease(param_projectId, param_jobId,
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

func Dataflow_v1b3_ProjectsJobsWorkItemsReportStatus(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("jobId"))

		if len(pathParams) != 0 {
			if strings.Contains("v1b3/projects/{projectId}/jobs/{jobId}/workItems:reportStatus", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		commands_util.PrintRequestExample(&api_client.ReportWorkItemStatusRequest{})

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsJobsWorkItemsService(api_service)

	args, flagValues, err := commands_util.ExtractFlagValues(args)
	if err != nil {
		return err
	}

	// Only positional arguments should remain in args.
	if len(args) == 0 || len(args) > 2 {
		usageFunc()
	}

	request := &api_client.ReportWorkItemStatusRequest{}
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
		"jobId",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_jobId, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}

	call := service.ReportStatus(param_projectId, param_jobId,
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

func Dataflow_v1b3_ProjectsLocationsJobsCreate(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("location"))

		if len(pathParams) != 0 {
			if strings.Contains("v1b3/projects/{projectId}/locations/{location}/jobs", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		usageBits += " [--replaceJobId=VALUE]"

		usageBits += " [--view=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		commands_util.PrintRequestExample(&api_client.Job{})

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsLocationsJobsService(api_service)

	queryParamNames := map[string]bool{
		"replaceJobId": false,
		"view":         false,
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

	request := &api_client.Job{}
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
		"projectId",
		"location",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_location, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}

	call := service.Create(param_projectId, param_location,
		request,
	)

	// Set query parameters.
	if value, ok := flagValues["replaceJobId"]; ok {
		query_replaceJobId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.ReplaceJobId(query_replaceJobId)
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

func Dataflow_v1b3_ProjectsLocationsJobsGet(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("location"))
		pathParams = append(pathParams, commands_util.AngrySnakes("jobId"))

		if len(pathParams) != 0 {
			if strings.Contains("v1b3/projects/{projectId}/locations/{location}/jobs/{jobId}", "+") {
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
	service := api_client.NewProjectsLocationsJobsService(api_service)

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
		"projectId",
		"location",
		"jobId",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_location, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}
	param_jobId, err := commands_util.ConvertValue_string(paramValues[2])
	if err != nil {
		return err
	}

	call := service.Get(param_projectId, param_location, param_jobId)

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

func Dataflow_v1b3_ProjectsLocationsJobsGetMetrics(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("location"))
		pathParams = append(pathParams, commands_util.AngrySnakes("jobId"))

		if len(pathParams) != 0 {
			if strings.Contains("v1b3/projects/{projectId}/locations/{location}/jobs/{jobId}/metrics", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [--startTime=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsLocationsJobsService(api_service)

	queryParamNames := map[string]bool{
		"startTime": false,
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
		"projectId",
		"location",
		"jobId",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_location, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}
	param_jobId, err := commands_util.ConvertValue_string(paramValues[2])
	if err != nil {
		return err
	}

	call := service.GetMetrics(param_projectId, param_location, param_jobId)

	// Set query parameters.
	if value, ok := flagValues["startTime"]; ok {
		query_startTime, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.StartTime(query_startTime)
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

func Dataflow_v1b3_ProjectsLocationsJobsList(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("location"))

		if len(pathParams) != 0 {
			if strings.Contains("v1b3/projects/{projectId}/locations/{location}/jobs", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [--filter=VALUE]"

		usageBits += " [--pageSize=VALUE]"

		usageBits += " [--pageToken=VALUE]"

		usageBits += " [--view=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsLocationsJobsService(api_service)

	queryParamNames := map[string]bool{
		"filter":    false,
		"pageSize":  false,
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
		"projectId",
		"location",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_location, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}

	call := service.List(param_projectId, param_location)

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

func Dataflow_v1b3_ProjectsLocationsJobsMessagesList(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("location"))
		pathParams = append(pathParams, commands_util.AngrySnakes("jobId"))

		if len(pathParams) != 0 {
			if strings.Contains("v1b3/projects/{projectId}/locations/{location}/jobs/{jobId}/messages", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [--endTime=VALUE]"

		usageBits += " [--minimumImportance=VALUE]"

		usageBits += " [--pageSize=VALUE]"

		usageBits += " [--pageToken=VALUE]"

		usageBits += " [--startTime=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsLocationsJobsMessagesService(api_service)

	queryParamNames := map[string]bool{
		"endTime":           false,
		"minimumImportance": false,
		"pageSize":          false,
		"pageToken":         false,
		"startTime":         false,
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
		"projectId",
		"location",
		"jobId",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_location, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}
	param_jobId, err := commands_util.ConvertValue_string(paramValues[2])
	if err != nil {
		return err
	}

	call := service.List(param_projectId, param_location, param_jobId)

	// Set query parameters.
	if value, ok := flagValues["endTime"]; ok {
		query_endTime, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.EndTime(query_endTime)
	}
	if value, ok := flagValues["minimumImportance"]; ok {
		query_minimumImportance, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.MinimumImportance(query_minimumImportance)
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
	if value, ok := flagValues["startTime"]; ok {
		query_startTime, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.StartTime(query_startTime)
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

func Dataflow_v1b3_ProjectsLocationsJobsUpdate(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("location"))
		pathParams = append(pathParams, commands_util.AngrySnakes("jobId"))

		if len(pathParams) != 0 {
			if strings.Contains("v1b3/projects/{projectId}/locations/{location}/jobs/{jobId}", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		commands_util.PrintRequestExample(&api_client.Job{})

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsLocationsJobsService(api_service)

	args, flagValues, err := commands_util.ExtractFlagValues(args)
	if err != nil {
		return err
	}

	// Only positional arguments should remain in args.
	if len(args) == 0 || len(args) > 2 {
		usageFunc()
	}

	request := &api_client.Job{}
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
		"location",
		"jobId",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_location, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}
	param_jobId, err := commands_util.ConvertValue_string(paramValues[2])
	if err != nil {
		return err
	}

	call := service.Update(param_projectId, param_location, param_jobId,
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

func Dataflow_v1b3_ProjectsLocationsJobsWorkItemsLease(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("location"))
		pathParams = append(pathParams, commands_util.AngrySnakes("jobId"))

		if len(pathParams) != 0 {
			if strings.Contains("v1b3/projects/{projectId}/locations/{location}/jobs/{jobId}/workItems:lease", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		commands_util.PrintRequestExample(&api_client.LeaseWorkItemRequest{})

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsLocationsJobsWorkItemsService(api_service)

	args, flagValues, err := commands_util.ExtractFlagValues(args)
	if err != nil {
		return err
	}

	// Only positional arguments should remain in args.
	if len(args) == 0 || len(args) > 2 {
		usageFunc()
	}

	request := &api_client.LeaseWorkItemRequest{}
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
		"location",
		"jobId",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_location, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}
	param_jobId, err := commands_util.ConvertValue_string(paramValues[2])
	if err != nil {
		return err
	}

	call := service.Lease(param_projectId, param_location, param_jobId,
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

func Dataflow_v1b3_ProjectsLocationsJobsWorkItemsReportStatus(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("location"))
		pathParams = append(pathParams, commands_util.AngrySnakes("jobId"))

		if len(pathParams) != 0 {
			if strings.Contains("v1b3/projects/{projectId}/locations/{location}/jobs/{jobId}/workItems:reportStatus", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		commands_util.PrintRequestExample(&api_client.ReportWorkItemStatusRequest{})

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsLocationsJobsWorkItemsService(api_service)

	args, flagValues, err := commands_util.ExtractFlagValues(args)
	if err != nil {
		return err
	}

	// Only positional arguments should remain in args.
	if len(args) == 0 || len(args) > 2 {
		usageFunc()
	}

	request := &api_client.ReportWorkItemStatusRequest{}
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
		"location",
		"jobId",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_location, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}
	param_jobId, err := commands_util.ConvertValue_string(paramValues[2])
	if err != nil {
		return err
	}

	call := service.ReportStatus(param_projectId, param_location, param_jobId,
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

func Dataflow_v1b3_ProjectsTemplatesCreate(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectId"))

		if len(pathParams) != 0 {
			if strings.Contains("v1b3/projects/{projectId}/templates", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		commands_util.PrintRequestExample(&api_client.CreateJobFromTemplateRequest{})

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsTemplatesService(api_service)

	args, flagValues, err := commands_util.ExtractFlagValues(args)
	if err != nil {
		return err
	}

	// Only positional arguments should remain in args.
	if len(args) == 0 || len(args) > 2 {
		usageFunc()
	}

	request := &api_client.CreateJobFromTemplateRequest{}
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
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}

	call := service.Create(param_projectId,
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

func Dataflow_v1b3_ProjectsWorkerMessages(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectId"))

		if len(pathParams) != 0 {
			if strings.Contains("v1b3/projects/{projectId}/WorkerMessages", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		commands_util.PrintRequestExample(&api_client.SendWorkerMessagesRequest{})

		os.Exit(1)
	}

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
		usageFunc()
	}

	request := &api_client.SendWorkerMessagesRequest{}
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
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}

	call := service.WorkerMessages(param_projectId,
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
