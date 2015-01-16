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

	api_client "github.com/GoogleCloudPlatform/gcloud/gcloud_apis/clients/dns/v1beta1"
	"github.com/GoogleCloudPlatform/gcloud/gcloud_apis/commands_util"
)

// some generated code may not use all imported libs
var _ = fmt.Println
var _ = io.Copy
var _ = os.Stdin

func Dns_v1beta1_ChangesCreate(context Context, args ...string) error {

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
	service := api_client.NewChangesService(api_service)

	args, flagValues, err := commands_util.ExtractFlagValues(args)
	if err != nil {
		return err
	}

	// Only positional arguments should remain in args.
	if len(args) == 0 || len(args) > 2 {
		usageFunc()
	}

	request := &api_client.Change{}
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
		"managedZone",
	}
	paramValues := strings.Split(args[0], "/")
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_project := paramValues[0]
	param_managedZone := paramValues[1]

	call := service.Create(param_project, param_managedZone,
		request,
	)

	var response *api_client.Change
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

func Dns_v1beta1_ChangesGet(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewChangesService(api_service)

	// Only positional arguments should remain in args.
	if len(args) != 1 {
		usageFunc()
	}

	expectedParams := []string{
		"project",
		"managedZone",
		"changeId",
	}
	paramValues := strings.Split(args[0], "/")
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_project := paramValues[0]
	param_managedZone := paramValues[1]
	param_changeId := paramValues[2]

	call := service.Get(param_project, param_managedZone, param_changeId)

	var response *api_client.Change
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

func Dns_v1beta1_ChangesList(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)

		usageBits += " [--maxResults=VALUE]"

		usageBits += " [--pageToken=VALUE]"

		usageBits += " [--sortBy=VALUE]"

		usageBits += " [--sortOrder=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewChangesService(api_service)

	queryParamNames := map[string]bool{
		"maxResults": false,
		"pageToken":  false,
		"sortBy":     false,
		"sortOrder":  false,
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
		"managedZone",
	}
	paramValues := strings.Split(args[0], "/")
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_project := paramValues[0]
	param_managedZone := paramValues[1]

	call := service.List(param_project, param_managedZone)

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
	if value, ok := flagValues["sortBy"]; ok {
		query_sortBy, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SortBy(query_sortBy)
	}
	if value, ok := flagValues["sortOrder"]; ok {
		query_sortOrder, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SortOrder(query_sortOrder)
	}

	var response *api_client.ChangesListResponse
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

func Dns_v1beta1_ManagedZonesCreate(context Context, args ...string) error {

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
	service := api_client.NewManagedZonesService(api_service)

	args, flagValues, err := commands_util.ExtractFlagValues(args)
	if err != nil {
		return err
	}

	// Only positional arguments should remain in args.
	if len(args) == 0 || len(args) > 2 {
		usageFunc()
	}

	request := &api_client.ManagedZone{}
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
	}
	paramValues := strings.Split(args[0], "/")
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_project := paramValues[0]

	call := service.Create(param_project,
		request,
	)

	var response *api_client.ManagedZone
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

func Dns_v1beta1_ManagedZonesDelete(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewManagedZonesService(api_service)

	// Only positional arguments should remain in args.
	if len(args) != 1 {
		usageFunc()
	}

	expectedParams := []string{
		"project",
		"managedZone",
	}
	paramValues := strings.Split(args[0], "/")
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_project := paramValues[0]
	param_managedZone := paramValues[1]

	call := service.Delete(param_project, param_managedZone)

	err = call.Do()
	if err != nil {
		return err
	}

	return nil
}

func Dns_v1beta1_ManagedZonesGet(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewManagedZonesService(api_service)

	// Only positional arguments should remain in args.
	if len(args) != 1 {
		usageFunc()
	}

	expectedParams := []string{
		"project",
		"managedZone",
	}
	paramValues := strings.Split(args[0], "/")
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_project := paramValues[0]
	param_managedZone := paramValues[1]

	call := service.Get(param_project, param_managedZone)

	var response *api_client.ManagedZone
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

func Dns_v1beta1_ManagedZonesList(context Context, args ...string) error {

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
	service := api_client.NewManagedZonesService(api_service)

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
	}
	paramValues := strings.Split(args[0], "/")
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_project := paramValues[0]

	call := service.List(param_project)

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

	var response *api_client.ManagedZonesListResponse
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

func Dns_v1beta1_ProjectsGet(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)

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
		"project",
	}
	paramValues := strings.Split(args[0], "/")
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_project := paramValues[0]

	call := service.Get(param_project)

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

func Dns_v1beta1_ResourceRecordSetsList(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)

		usageBits += " [--maxResults=VALUE]"

		usageBits += " [--name=VALUE]"

		usageBits += " [--pageToken=VALUE]"

		usageBits += " [--type=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewResourceRecordSetsService(api_service)

	queryParamNames := map[string]bool{
		"maxResults": false,
		"name":       false,
		"pageToken":  false,
		"type":       false,
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
		"managedZone",
	}
	paramValues := strings.Split(args[0], "/")
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_project := paramValues[0]
	param_managedZone := paramValues[1]

	call := service.List(param_project, param_managedZone)

	// Set query parameters.
	if value, ok := flagValues["maxResults"]; ok {
		query_maxResults, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.MaxResults(query_maxResults)
	}
	if value, ok := flagValues["name"]; ok {
		query_name, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.Name(query_name)
	}
	if value, ok := flagValues["pageToken"]; ok {
		query_pageToken, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.PageToken(query_pageToken)
	}
	if value, ok := flagValues["type"]; ok {
		query_type, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.Type(query_type)
	}

	var response *api_client.ResourceRecordSetsListResponse
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
