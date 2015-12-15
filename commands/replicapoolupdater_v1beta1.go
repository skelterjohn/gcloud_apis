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

	api_client "github.com/skelterjohn/gcloud_apis/clients/replicapoolupdater/v1beta1"
	"github.com/skelterjohn/gcloud_apis/commands_util"
)

// some generated code may not use all imported libs
var _ = fmt.Println
var _ = io.Copy
var _ = os.Stdin
var _ = strings.Split

func Replicapoolupdater_v1beta1_RollingUpdatesCancel(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("project"))
		pathParams = append(pathParams, commands_util.AngrySnakes("zone"))
		pathParams = append(pathParams, commands_util.AngrySnakes("rollingUpdate"))

		if len(pathParams) != 0 {
			if strings.Contains("{project}/zones/{zone}/rollingUpdates/{rollingUpdate}/cancel", "+") {
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
	service := api_client.NewRollingUpdatesService(api_service)

	// Only positional arguments should remain in args.
	if len(args) != 1 {
		usageFunc()
	}

	expectedParams := []string{
		"project",
		"zone",
		"rollingUpdate",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_project, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_zone, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}
	param_rollingUpdate, err := commands_util.ConvertValue_string(paramValues[2])
	if err != nil {
		return err
	}

	call := service.Cancel(param_project, param_zone, param_rollingUpdate)

	var response *api_client.Operation
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

func Replicapoolupdater_v1beta1_RollingUpdatesGet(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("project"))
		pathParams = append(pathParams, commands_util.AngrySnakes("zone"))
		pathParams = append(pathParams, commands_util.AngrySnakes("rollingUpdate"))

		if len(pathParams) != 0 {
			if strings.Contains("{project}/zones/{zone}/rollingUpdates/{rollingUpdate}", "+") {
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
	service := api_client.NewRollingUpdatesService(api_service)

	// Only positional arguments should remain in args.
	if len(args) != 1 {
		usageFunc()
	}

	expectedParams := []string{
		"project",
		"zone",
		"rollingUpdate",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_project, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_zone, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}
	param_rollingUpdate, err := commands_util.ConvertValue_string(paramValues[2])
	if err != nil {
		return err
	}

	call := service.Get(param_project, param_zone, param_rollingUpdate)

	var response *api_client.RollingUpdate
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

func Replicapoolupdater_v1beta1_RollingUpdatesInsert(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("project"))
		pathParams = append(pathParams, commands_util.AngrySnakes("zone"))

		if len(pathParams) != 0 {
			if strings.Contains("{project}/zones/{zone}/rollingUpdates", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewRollingUpdatesService(api_service)

	args, flagValues, err := commands_util.ExtractFlagValues(args)
	if err != nil {
		return err
	}

	// Only positional arguments should remain in args.
	if len(args) == 0 || len(args) > 2 {
		usageFunc()
	}

	request := &api_client.RollingUpdate{}
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
		"project",
		"zone",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_project, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_zone, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}

	call := service.Insert(param_project, param_zone,
		request,
	)

	var response *api_client.Operation
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

func Replicapoolupdater_v1beta1_RollingUpdatesList(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("project"))
		pathParams = append(pathParams, commands_util.AngrySnakes("zone"))

		if len(pathParams) != 0 {
			if strings.Contains("{project}/zones/{zone}/rollingUpdates", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [--filter=VALUE]"

		usageBits += " [--maxResults=VALUE]"

		usageBits += " [--pageToken=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewRollingUpdatesService(api_service)

	queryParamNames := map[string]bool{
		"filter":     false,
		"maxResults": false,
		"pageToken":  false,
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
		"project",
		"zone",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_project, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_zone, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}

	call := service.List(param_project, param_zone)

	// Set query parameters.
	if value, ok := flagValues["filter"]; ok {
		query_filter, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.Filter(query_filter)
	}
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

	var response *api_client.RollingUpdateList
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

func Replicapoolupdater_v1beta1_RollingUpdatesListInstanceUpdates(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("project"))
		pathParams = append(pathParams, commands_util.AngrySnakes("zone"))
		pathParams = append(pathParams, commands_util.AngrySnakes("rollingUpdate"))

		if len(pathParams) != 0 {
			if strings.Contains("{project}/zones/{zone}/rollingUpdates/{rollingUpdate}/instanceUpdates", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [--filter=VALUE]"

		usageBits += " [--maxResults=VALUE]"

		usageBits += " [--pageToken=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewRollingUpdatesService(api_service)

	queryParamNames := map[string]bool{
		"filter":     false,
		"maxResults": false,
		"pageToken":  false,
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
		"project",
		"zone",
		"rollingUpdate",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_project, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_zone, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}
	param_rollingUpdate, err := commands_util.ConvertValue_string(paramValues[2])
	if err != nil {
		return err
	}

	call := service.ListInstanceUpdates(param_project, param_zone, param_rollingUpdate)

	// Set query parameters.
	if value, ok := flagValues["filter"]; ok {
		query_filter, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.Filter(query_filter)
	}
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

	var response *api_client.InstanceUpdateList
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

func Replicapoolupdater_v1beta1_RollingUpdatesPause(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("project"))
		pathParams = append(pathParams, commands_util.AngrySnakes("zone"))
		pathParams = append(pathParams, commands_util.AngrySnakes("rollingUpdate"))

		if len(pathParams) != 0 {
			if strings.Contains("{project}/zones/{zone}/rollingUpdates/{rollingUpdate}/pause", "+") {
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
	service := api_client.NewRollingUpdatesService(api_service)

	// Only positional arguments should remain in args.
	if len(args) != 1 {
		usageFunc()
	}

	expectedParams := []string{
		"project",
		"zone",
		"rollingUpdate",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_project, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_zone, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}
	param_rollingUpdate, err := commands_util.ConvertValue_string(paramValues[2])
	if err != nil {
		return err
	}

	call := service.Pause(param_project, param_zone, param_rollingUpdate)

	var response *api_client.Operation
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

func Replicapoolupdater_v1beta1_RollingUpdatesResume(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("project"))
		pathParams = append(pathParams, commands_util.AngrySnakes("zone"))
		pathParams = append(pathParams, commands_util.AngrySnakes("rollingUpdate"))

		if len(pathParams) != 0 {
			if strings.Contains("{project}/zones/{zone}/rollingUpdates/{rollingUpdate}/resume", "+") {
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
	service := api_client.NewRollingUpdatesService(api_service)

	// Only positional arguments should remain in args.
	if len(args) != 1 {
		usageFunc()
	}

	expectedParams := []string{
		"project",
		"zone",
		"rollingUpdate",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_project, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_zone, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}
	param_rollingUpdate, err := commands_util.ConvertValue_string(paramValues[2])
	if err != nil {
		return err
	}

	call := service.Resume(param_project, param_zone, param_rollingUpdate)

	var response *api_client.Operation
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

func Replicapoolupdater_v1beta1_RollingUpdatesRollback(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("project"))
		pathParams = append(pathParams, commands_util.AngrySnakes("zone"))
		pathParams = append(pathParams, commands_util.AngrySnakes("rollingUpdate"))

		if len(pathParams) != 0 {
			if strings.Contains("{project}/zones/{zone}/rollingUpdates/{rollingUpdate}/rollback", "+") {
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
	service := api_client.NewRollingUpdatesService(api_service)

	// Only positional arguments should remain in args.
	if len(args) != 1 {
		usageFunc()
	}

	expectedParams := []string{
		"project",
		"zone",
		"rollingUpdate",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_project, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_zone, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}
	param_rollingUpdate, err := commands_util.ConvertValue_string(paramValues[2])
	if err != nil {
		return err
	}

	call := service.Rollback(param_project, param_zone, param_rollingUpdate)

	var response *api_client.Operation
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

func Replicapoolupdater_v1beta1_ZoneOperationsGet(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("project"))
		pathParams = append(pathParams, commands_util.AngrySnakes("zone"))
		pathParams = append(pathParams, commands_util.AngrySnakes("operation"))

		if len(pathParams) != 0 {
			if strings.Contains("{project}/zones/{zone}/operations/{operation}", "+") {
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
	service := api_client.NewZoneOperationsService(api_service)

	// Only positional arguments should remain in args.
	if len(args) != 1 {
		usageFunc()
	}

	expectedParams := []string{
		"project",
		"zone",
		"operation",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_project, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_zone, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}
	param_operation, err := commands_util.ConvertValue_string(paramValues[2])
	if err != nil {
		return err
	}

	call := service.Get(param_project, param_zone, param_operation)

	var response *api_client.Operation
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

func Replicapoolupdater_v1beta1_ZoneOperationsList(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("project"))
		pathParams = append(pathParams, commands_util.AngrySnakes("zone"))

		if len(pathParams) != 0 {
			if strings.Contains("{project}/zones/{zone}/operations", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [--filter=VALUE]"

		usageBits += " [--maxResults=VALUE]"

		usageBits += " [--pageToken=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewZoneOperationsService(api_service)

	queryParamNames := map[string]bool{
		"filter":     false,
		"maxResults": false,
		"pageToken":  false,
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
		"project",
		"zone",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_project, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_zone, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}

	call := service.List(param_project, param_zone)

	// Set query parameters.
	if value, ok := flagValues["filter"]; ok {
		query_filter, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.Filter(query_filter)
	}
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

	var response *api_client.OperationList
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
