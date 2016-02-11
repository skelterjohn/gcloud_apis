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

	api_client "github.com/skelterjohn/gcloud_apis/clients/replicapool/v1beta2"
	"github.com/skelterjohn/gcloud_apis/commands_util"
)

// some generated code may not use all imported libs
var _ = fmt.Println
var _ = io.Copy
var _ = os.Stdin
var _ = strings.Split

func Replicapool_v1beta2_InstanceGroupManagersAbandonInstances(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("project"))
		pathParams = append(pathParams, commands_util.AngrySnakes("zone"))
		pathParams = append(pathParams, commands_util.AngrySnakes("instanceGroupManager"))

		if len(pathParams) != 0 {
			if strings.Contains("{project}/zones/{zone}/instanceGroupManagers/{instanceGroupManager}/abandonInstances", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		commands_util.PrintRequestExample(&api_client.InstanceGroupManagersAbandonInstancesRequest{})

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewInstanceGroupManagersService(api_service)

	args, flagValues, err := commands_util.ExtractFlagValues(args)
	if err != nil {
		return err
	}

	// Only positional arguments should remain in args.
	if len(args) == 0 || len(args) > 2 {
		usageFunc()
	}

	request := &api_client.InstanceGroupManagersAbandonInstancesRequest{}
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
		"instanceGroupManager",
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
	param_instanceGroupManager, err := commands_util.ConvertValue_string(paramValues[2])
	if err != nil {
		return err
	}

	call := service.AbandonInstances(param_project, param_zone, param_instanceGroupManager,
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

func Replicapool_v1beta2_InstanceGroupManagersDelete(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("project"))
		pathParams = append(pathParams, commands_util.AngrySnakes("zone"))
		pathParams = append(pathParams, commands_util.AngrySnakes("instanceGroupManager"))

		if len(pathParams) != 0 {
			if strings.Contains("{project}/zones/{zone}/instanceGroupManagers/{instanceGroupManager}", "+") {
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
	service := api_client.NewInstanceGroupManagersService(api_service)

	// Only positional arguments should remain in args.
	if len(args) != 1 {
		usageFunc()
	}

	expectedParams := []string{
		"project",
		"zone",
		"instanceGroupManager",
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
	param_instanceGroupManager, err := commands_util.ConvertValue_string(paramValues[2])
	if err != nil {
		return err
	}

	call := service.Delete(param_project, param_zone, param_instanceGroupManager)

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

func Replicapool_v1beta2_InstanceGroupManagersDeleteInstances(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("project"))
		pathParams = append(pathParams, commands_util.AngrySnakes("zone"))
		pathParams = append(pathParams, commands_util.AngrySnakes("instanceGroupManager"))

		if len(pathParams) != 0 {
			if strings.Contains("{project}/zones/{zone}/instanceGroupManagers/{instanceGroupManager}/deleteInstances", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		commands_util.PrintRequestExample(&api_client.InstanceGroupManagersDeleteInstancesRequest{})

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewInstanceGroupManagersService(api_service)

	args, flagValues, err := commands_util.ExtractFlagValues(args)
	if err != nil {
		return err
	}

	// Only positional arguments should remain in args.
	if len(args) == 0 || len(args) > 2 {
		usageFunc()
	}

	request := &api_client.InstanceGroupManagersDeleteInstancesRequest{}
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
		"instanceGroupManager",
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
	param_instanceGroupManager, err := commands_util.ConvertValue_string(paramValues[2])
	if err != nil {
		return err
	}

	call := service.DeleteInstances(param_project, param_zone, param_instanceGroupManager,
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

func Replicapool_v1beta2_InstanceGroupManagersGet(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("project"))
		pathParams = append(pathParams, commands_util.AngrySnakes("zone"))
		pathParams = append(pathParams, commands_util.AngrySnakes("instanceGroupManager"))

		if len(pathParams) != 0 {
			if strings.Contains("{project}/zones/{zone}/instanceGroupManagers/{instanceGroupManager}", "+") {
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
	service := api_client.NewInstanceGroupManagersService(api_service)

	// Only positional arguments should remain in args.
	if len(args) != 1 {
		usageFunc()
	}

	expectedParams := []string{
		"project",
		"zone",
		"instanceGroupManager",
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
	param_instanceGroupManager, err := commands_util.ConvertValue_string(paramValues[2])
	if err != nil {
		return err
	}

	call := service.Get(param_project, param_zone, param_instanceGroupManager)

	var response *api_client.InstanceGroupManager
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

func Replicapool_v1beta2_InstanceGroupManagersInsert(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("project"))
		pathParams = append(pathParams, commands_util.AngrySnakes("zone"))
		pathParams = append(pathParams, commands_util.AngrySnakes("size"))

		if len(pathParams) != 0 {
			if strings.Contains("{project}/zones/{zone}/instanceGroupManagers", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		usageBits += " --size=VALUE"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		commands_util.PrintRequestExample(&api_client.InstanceGroupManager{})

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewInstanceGroupManagersService(api_service)

	queryParamNames := map[string]bool{
		"size": true,
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

	request := &api_client.InstanceGroupManager{}
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
	param_size, err := commands_util.ConvertValue_int64(flagValues["size"])
	if err != nil {
		return err
	}

	call := service.Insert(param_project, param_zone, param_size,
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

func Replicapool_v1beta2_InstanceGroupManagersList(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("project"))
		pathParams = append(pathParams, commands_util.AngrySnakes("zone"))

		if len(pathParams) != 0 {
			if strings.Contains("{project}/zones/{zone}/instanceGroupManagers", "+") {
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
	service := api_client.NewInstanceGroupManagersService(api_service)

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

	var response *api_client.InstanceGroupManagerList
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

func Replicapool_v1beta2_InstanceGroupManagersRecreateInstances(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("project"))
		pathParams = append(pathParams, commands_util.AngrySnakes("zone"))
		pathParams = append(pathParams, commands_util.AngrySnakes("instanceGroupManager"))

		if len(pathParams) != 0 {
			if strings.Contains("{project}/zones/{zone}/instanceGroupManagers/{instanceGroupManager}/recreateInstances", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		commands_util.PrintRequestExample(&api_client.InstanceGroupManagersRecreateInstancesRequest{})

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewInstanceGroupManagersService(api_service)

	args, flagValues, err := commands_util.ExtractFlagValues(args)
	if err != nil {
		return err
	}

	// Only positional arguments should remain in args.
	if len(args) == 0 || len(args) > 2 {
		usageFunc()
	}

	request := &api_client.InstanceGroupManagersRecreateInstancesRequest{}
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
		"instanceGroupManager",
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
	param_instanceGroupManager, err := commands_util.ConvertValue_string(paramValues[2])
	if err != nil {
		return err
	}

	call := service.RecreateInstances(param_project, param_zone, param_instanceGroupManager,
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

func Replicapool_v1beta2_InstanceGroupManagersResize(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("project"))
		pathParams = append(pathParams, commands_util.AngrySnakes("zone"))
		pathParams = append(pathParams, commands_util.AngrySnakes("instanceGroupManager"))
		pathParams = append(pathParams, commands_util.AngrySnakes("size"))

		if len(pathParams) != 0 {
			if strings.Contains("{project}/zones/{zone}/instanceGroupManagers/{instanceGroupManager}/resize", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " --size=VALUE"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewInstanceGroupManagersService(api_service)

	queryParamNames := map[string]bool{
		"size": true,
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
		"instanceGroupManager",
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
	param_instanceGroupManager, err := commands_util.ConvertValue_string(paramValues[2])
	if err != nil {
		return err
	}
	param_size, err := commands_util.ConvertValue_int64(flagValues["size"])
	if err != nil {
		return err
	}

	call := service.Resize(param_project, param_zone, param_instanceGroupManager, param_size)

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

func Replicapool_v1beta2_InstanceGroupManagersSetAutoHealingPolicy(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("project"))
		pathParams = append(pathParams, commands_util.AngrySnakes("zone"))
		pathParams = append(pathParams, commands_util.AngrySnakes("instanceGroupManager"))

		if len(pathParams) != 0 {
			if strings.Contains("{project}/zones/{zone}/instanceGroupManagers/{instanceGroupManager}/setAutoHealingPolicy", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		commands_util.PrintRequestExample(&api_client.InstanceGroupManagersSetAutoHealingPolicyRequest{})

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewInstanceGroupManagersService(api_service)

	args, flagValues, err := commands_util.ExtractFlagValues(args)
	if err != nil {
		return err
	}

	// Only positional arguments should remain in args.
	if len(args) == 0 || len(args) > 2 {
		usageFunc()
	}

	request := &api_client.InstanceGroupManagersSetAutoHealingPolicyRequest{}
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
		"instanceGroupManager",
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
	param_instanceGroupManager, err := commands_util.ConvertValue_string(paramValues[2])
	if err != nil {
		return err
	}

	call := service.SetAutoHealingPolicy(param_project, param_zone, param_instanceGroupManager,
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

func Replicapool_v1beta2_InstanceGroupManagersSetInstanceTemplate(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("project"))
		pathParams = append(pathParams, commands_util.AngrySnakes("zone"))
		pathParams = append(pathParams, commands_util.AngrySnakes("instanceGroupManager"))

		if len(pathParams) != 0 {
			if strings.Contains("{project}/zones/{zone}/instanceGroupManagers/{instanceGroupManager}/setInstanceTemplate", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		commands_util.PrintRequestExample(&api_client.InstanceGroupManagersSetInstanceTemplateRequest{})

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewInstanceGroupManagersService(api_service)

	args, flagValues, err := commands_util.ExtractFlagValues(args)
	if err != nil {
		return err
	}

	// Only positional arguments should remain in args.
	if len(args) == 0 || len(args) > 2 {
		usageFunc()
	}

	request := &api_client.InstanceGroupManagersSetInstanceTemplateRequest{}
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
		"instanceGroupManager",
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
	param_instanceGroupManager, err := commands_util.ConvertValue_string(paramValues[2])
	if err != nil {
		return err
	}

	call := service.SetInstanceTemplate(param_project, param_zone, param_instanceGroupManager,
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

func Replicapool_v1beta2_InstanceGroupManagersSetTargetPools(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("project"))
		pathParams = append(pathParams, commands_util.AngrySnakes("zone"))
		pathParams = append(pathParams, commands_util.AngrySnakes("instanceGroupManager"))

		if len(pathParams) != 0 {
			if strings.Contains("{project}/zones/{zone}/instanceGroupManagers/{instanceGroupManager}/setTargetPools", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		commands_util.PrintRequestExample(&api_client.InstanceGroupManagersSetTargetPoolsRequest{})

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewInstanceGroupManagersService(api_service)

	args, flagValues, err := commands_util.ExtractFlagValues(args)
	if err != nil {
		return err
	}

	// Only positional arguments should remain in args.
	if len(args) == 0 || len(args) > 2 {
		usageFunc()
	}

	request := &api_client.InstanceGroupManagersSetTargetPoolsRequest{}
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
		"instanceGroupManager",
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
	param_instanceGroupManager, err := commands_util.ConvertValue_string(paramValues[2])
	if err != nil {
		return err
	}

	call := service.SetTargetPools(param_project, param_zone, param_instanceGroupManager,
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

func Replicapool_v1beta2_ZoneOperationsGet(context Context, args ...string) error {

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

func Replicapool_v1beta2_ZoneOperationsList(context Context, args ...string) error {

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
