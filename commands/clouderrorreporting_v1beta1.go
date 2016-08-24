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

	api_client "github.com/skelterjohn/gcloud_apis/clients/clouderrorreporting/v1beta1"
	"github.com/skelterjohn/gcloud_apis/commands_util"
)

// some generated code may not use all imported libs
var _ = fmt.Println
var _ = io.Copy
var _ = os.Stdin
var _ = strings.Split

func Clouderrorreporting_v1beta1_ProjectsDeleteEvents(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectName"))

		if len(pathParams) != 0 {
			if strings.Contains("v1beta1/{+projectName}/events", "+") {
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
	service := api_client.NewProjectsService(api_service)

	// Only positional arguments should remain in args.
	if len(args) != 1 {
		usageFunc()
	}

	expectedParams := []string{
		"projectName",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectName, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}

	call := service.DeleteEvents(param_projectName)

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

func Clouderrorreporting_v1beta1_ProjectsEventsList(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectName"))

		if len(pathParams) != 0 {
			if strings.Contains("v1beta1/{+projectName}/events", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [--groupId=VALUE]"

		usageBits += " [--pageSize=VALUE]"

		usageBits += " [--pageToken=VALUE]"

		usageBits += " [--serviceFilter.service=VALUE]"

		usageBits += " [--serviceFilter.version=VALUE]"

		usageBits += " [--timeRange.period=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsEventsService(api_service)

	queryParamNames := map[string]bool{
		"groupId":               false,
		"pageSize":              false,
		"pageToken":             false,
		"serviceFilter.service": false,
		"serviceFilter.version": false,
		"timeRange.period":      false,
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
		"projectName",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectName, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}

	call := service.List(param_projectName)

	// Set query parameters.
	if value, ok := flagValues["groupId"]; ok {
		query_groupId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.GroupId(query_groupId)
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
	if value, ok := flagValues["serviceFilter_service"]; ok {
		query_serviceFilter_service, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.ServiceFilterService(query_serviceFilter_service)
	}
	if value, ok := flagValues["serviceFilter_version"]; ok {
		query_serviceFilter_version, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.ServiceFilterVersion(query_serviceFilter_version)
	}
	if value, ok := flagValues["timeRange_period"]; ok {
		query_timeRange_period, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.TimeRangePeriod(query_timeRange_period)
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

func Clouderrorreporting_v1beta1_ProjectsEventsReport(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectName"))

		if len(pathParams) != 0 {
			if strings.Contains("v1beta1/{+projectName}/events:report", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		commands_util.PrintRequestExample(&api_client.ReportedErrorEvent{})

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsEventsService(api_service)

	args, flagValues, err := commands_util.ExtractFlagValues(args)
	if err != nil {
		return err
	}

	// Only positional arguments should remain in args.
	if len(args) == 0 || len(args) > 2 {
		usageFunc()
	}

	request := &api_client.ReportedErrorEvent{}
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
		"projectName",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectName, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}

	call := service.Report(param_projectName,
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

func Clouderrorreporting_v1beta1_ProjectsGroupStatsList(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectName"))

		if len(pathParams) != 0 {
			if strings.Contains("v1beta1/{+projectName}/groupStats", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [--alignment=VALUE]"

		usageBits += " [--alignmentTime=VALUE]"

		usageBits += " [--groupId=VALUE]"

		usageBits += " [--order=VALUE]"

		usageBits += " [--pageSize=VALUE]"

		usageBits += " [--pageToken=VALUE]"

		usageBits += " [--serviceFilter.service=VALUE]"

		usageBits += " [--serviceFilter.version=VALUE]"

		usageBits += " [--timeRange.period=VALUE]"

		usageBits += " [--timedCountDuration=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsGroupStatsService(api_service)

	queryParamNames := map[string]bool{
		"alignment":             false,
		"alignmentTime":         false,
		"groupId":               false,
		"order":                 false,
		"pageSize":              false,
		"pageToken":             false,
		"serviceFilter.service": false,
		"serviceFilter.version": false,
		"timeRange.period":      false,
		"timedCountDuration":    false,
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
		"projectName",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectName, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}

	call := service.List(param_projectName)

	// Set query parameters.
	if value, ok := flagValues["alignment"]; ok {
		query_alignment, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.Alignment(query_alignment)
	}
	if value, ok := flagValues["alignmentTime"]; ok {
		query_alignmentTime, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.AlignmentTime(query_alignmentTime)
	}
	if value, ok := flagValues["groupId"]; ok {
		query_groupId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.GroupId(query_groupId)
	}
	if value, ok := flagValues["order"]; ok {
		query_order, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.Order(query_order)
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
	if value, ok := flagValues["serviceFilter_service"]; ok {
		query_serviceFilter_service, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.ServiceFilterService(query_serviceFilter_service)
	}
	if value, ok := flagValues["serviceFilter_version"]; ok {
		query_serviceFilter_version, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.ServiceFilterVersion(query_serviceFilter_version)
	}
	if value, ok := flagValues["timeRange_period"]; ok {
		query_timeRange_period, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.TimeRangePeriod(query_timeRange_period)
	}
	if value, ok := flagValues["timedCountDuration"]; ok {
		query_timedCountDuration, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.TimedCountDuration(query_timedCountDuration)
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

func Clouderrorreporting_v1beta1_ProjectsGroupsGet(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("groupName"))

		if len(pathParams) != 0 {
			if strings.Contains("v1beta1/{+groupName}", "+") {
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
	service := api_client.NewProjectsGroupsService(api_service)

	// Only positional arguments should remain in args.
	if len(args) != 1 {
		usageFunc()
	}

	expectedParams := []string{
		"groupName",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_groupName, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}

	call := service.Get(param_groupName)

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

func Clouderrorreporting_v1beta1_ProjectsGroupsUpdate(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("name"))

		if len(pathParams) != 0 {
			if strings.Contains("v1beta1/{+name}", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		commands_util.PrintRequestExample(&api_client.ErrorGroup{})

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsGroupsService(api_service)

	args, flagValues, err := commands_util.ExtractFlagValues(args)
	if err != nil {
		return err
	}

	// Only positional arguments should remain in args.
	if len(args) == 0 || len(args) > 2 {
		usageFunc()
	}

	request := &api_client.ErrorGroup{}
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
		"name",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_name, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}

	call := service.Update(param_name,
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
