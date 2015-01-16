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

	api_client "github.com/GoogleCloudPlatform/gcloud/gcloud_apis/clients/replicapoolupdater/v1beta1"
	"github.com/GoogleCloudPlatform/gcloud/gcloud_apis/commands_util"
)

// some generated code may not use all imported libs
var _ = fmt.Println
var _ = io.Copy
var _ = os.Stdin

func Replicapoolupdater_v1beta1_UpdatesCancel(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewUpdatesService(api_service)

	// Only positional arguments should remain in args.
	if len(args) != 1 {
		usageFunc()
	}

	expectedParams := []string{
		"project",
		"zone",
		"instanceGroupManager",
		"updateHandle",
	}
	paramValues := strings.Split(args[0], "/")
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_project := paramValues[0]
	param_zone := paramValues[1]
	param_instanceGroupManager := paramValues[2]
	param_updateHandle := paramValues[3]

	call := service.Cancel(param_project, param_zone, param_instanceGroupManager, param_updateHandle)

	err = call.Do()
	if err != nil {
		return err
	}

	return nil
}

func Replicapoolupdater_v1beta1_UpdatesGet(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewUpdatesService(api_service)

	// Only positional arguments should remain in args.
	if len(args) != 1 {
		usageFunc()
	}

	expectedParams := []string{
		"project",
		"zone",
		"instanceGroupManager",
		"updateHandle",
	}
	paramValues := strings.Split(args[0], "/")
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_project := paramValues[0]
	param_zone := paramValues[1]
	param_instanceGroupManager := paramValues[2]
	param_updateHandle := paramValues[3]

	call := service.Get(param_project, param_zone, param_instanceGroupManager, param_updateHandle)

	var response *api_client.Update
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

func Replicapoolupdater_v1beta1_UpdatesInsert(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewUpdatesService(api_service)

	args, flagValues, err := commands_util.ExtractFlagValues(args)
	if err != nil {
		return err
	}

	// Only positional arguments should remain in args.
	if len(args) == 0 || len(args) > 2 {
		usageFunc()
	}

	request := &api_client.Update{}
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
	paramValues := strings.Split(args[0], "/")
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_project := paramValues[0]
	param_zone := paramValues[1]
	param_instanceGroupManager := paramValues[2]

	call := service.Insert(param_project, param_zone, param_instanceGroupManager,
		request,
	)

	var response *api_client.InsertResponse
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

func Replicapoolupdater_v1beta1_UpdatesList(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)

		usageBits += " [--maxResults=VALUE]"

		usageBits += " [--pageToken=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewUpdatesService(api_service)

	queryParamNames := map[string]bool{
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
		"instanceGroupManager",
	}
	paramValues := strings.Split(args[0], "/")
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_project := paramValues[0]
	param_zone := paramValues[1]
	param_instanceGroupManager := paramValues[2]

	call := service.List(param_project, param_zone, param_instanceGroupManager)

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

	var response *api_client.UpdateList
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

func Replicapoolupdater_v1beta1_UpdatesPause(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewUpdatesService(api_service)

	// Only positional arguments should remain in args.
	if len(args) != 1 {
		usageFunc()
	}

	expectedParams := []string{
		"project",
		"zone",
		"instanceGroupManager",
		"updateHandle",
	}
	paramValues := strings.Split(args[0], "/")
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_project := paramValues[0]
	param_zone := paramValues[1]
	param_instanceGroupManager := paramValues[2]
	param_updateHandle := paramValues[3]

	call := service.Pause(param_project, param_zone, param_instanceGroupManager, param_updateHandle)

	err = call.Do()
	if err != nil {
		return err
	}

	return nil
}

func Replicapoolupdater_v1beta1_UpdatesRollback(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewUpdatesService(api_service)

	// Only positional arguments should remain in args.
	if len(args) != 1 {
		usageFunc()
	}

	expectedParams := []string{
		"project",
		"zone",
		"instanceGroupManager",
		"updateHandle",
	}
	paramValues := strings.Split(args[0], "/")
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_project := paramValues[0]
	param_zone := paramValues[1]
	param_instanceGroupManager := paramValues[2]
	param_updateHandle := paramValues[3]

	call := service.Rollback(param_project, param_zone, param_instanceGroupManager, param_updateHandle)

	err = call.Do()
	if err != nil {
		return err
	}

	return nil
}

func Replicapoolupdater_v1beta1_UpdatesRollforward(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewUpdatesService(api_service)

	// Only positional arguments should remain in args.
	if len(args) != 1 {
		usageFunc()
	}

	expectedParams := []string{
		"project",
		"zone",
		"instanceGroupManager",
		"updateHandle",
	}
	paramValues := strings.Split(args[0], "/")
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_project := paramValues[0]
	param_zone := paramValues[1]
	param_instanceGroupManager := paramValues[2]
	param_updateHandle := paramValues[3]

	call := service.Rollforward(param_project, param_zone, param_instanceGroupManager, param_updateHandle)

	err = call.Do()
	if err != nil {
		return err
	}

	return nil
}
