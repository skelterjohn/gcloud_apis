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

	api_client "github.com/skelterjohn/gcloud_apis/clients/logging/v2beta1"
	"github.com/skelterjohn/gcloud_apis/commands_util"
)

// some generated code may not use all imported libs
var _ = fmt.Println
var _ = io.Copy
var _ = os.Stdin
var _ = strings.Split

func Logging_v2beta1_BillingAccountsLogsDelete(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("billingAccountsId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("logsId"))

		if len(pathParams) != 0 {
			if strings.Contains("v2beta1/billingAccounts/{billingAccountsId}/logs/{logsId}", "+") {
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
	service := api_client.NewBillingAccountsLogsService(api_service)

	// Only positional arguments should remain in args.
	if len(args) != 1 {
		usageFunc()
	}

	expectedParams := []string{
		"billingAccountsId",
		"logsId",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_billingAccountsId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_logsId, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}

	call := service.Delete(param_billingAccountsId, param_logsId)

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

func Logging_v2beta1_BillingAccountsLogsList(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("billingAccountsId"))

		if len(pathParams) != 0 {
			if strings.Contains("v2beta1/billingAccounts/{billingAccountsId}/logs", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [--pageSize=VALUE]"

		usageBits += " [--pageToken=VALUE]"

		usageBits += " [--resourceIndexPrefix=VALUE]"

		usageBits += " [--resourceType=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewBillingAccountsLogsService(api_service)

	queryParamNames := map[string]bool{
		"pageSize":            false,
		"pageToken":           false,
		"resourceIndexPrefix": false,
		"resourceType":        false,
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
		"billingAccountsId",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_billingAccountsId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}

	call := service.List(param_billingAccountsId)

	// Set query parameters.
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
	if value, ok := flagValues["resourceIndexPrefix"]; ok {
		query_resourceIndexPrefix, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.ResourceIndexPrefix(query_resourceIndexPrefix)
	}
	if value, ok := flagValues["resourceType"]; ok {
		query_resourceType, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.ResourceType(query_resourceType)
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

func Logging_v2beta1_BillingAccountsResourceKeysList(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("billingAccountsId"))

		if len(pathParams) != 0 {
			if strings.Contains("v2beta1/billingAccounts/{billingAccountsId}/resourceKeys", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [--pageSize=VALUE]"

		usageBits += " [--pageToken=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewBillingAccountsResourceKeysService(api_service)

	queryParamNames := map[string]bool{
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

	expectedParams := []string{
		"billingAccountsId",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_billingAccountsId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}

	call := service.List(param_billingAccountsId)

	// Set query parameters.
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

func Logging_v2beta1_BillingAccountsResourceTypesValuesList(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("billingAccountsId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("resourceTypesId"))

		if len(pathParams) != 0 {
			if strings.Contains("v2beta1/billingAccounts/{billingAccountsId}/resourceTypes/{resourceTypesId}/values", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [--depth=VALUE]"

		usageBits += " [--indexPrefix=VALUE]"

		usageBits += " [--pageSize=VALUE]"

		usageBits += " [--pageToken=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewBillingAccountsResourceTypesValuesService(api_service)

	queryParamNames := map[string]bool{
		"depth":       false,
		"indexPrefix": false,
		"pageSize":    false,
		"pageToken":   false,
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
		"billingAccountsId",
		"resourceTypesId",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_billingAccountsId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_resourceTypesId, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}

	call := service.List(param_billingAccountsId, param_resourceTypesId)

	// Set query parameters.
	if value, ok := flagValues["depth"]; ok {
		query_depth, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.Depth(query_depth)
	}
	if value, ok := flagValues["indexPrefix"]; ok {
		query_indexPrefix, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.IndexPrefix(query_indexPrefix)
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

func Logging_v2beta1_EntriesList(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string

		if len(pathParams) != 0 {
			if strings.Contains("v2beta1/entries:list", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		commands_util.PrintRequestExample(&api_client.ListLogEntriesRequest{})

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewEntriesService(api_service)

	args, flagValues, err := commands_util.ExtractFlagValues(args)
	if err != nil {
		return err
	}

	// Only positional arguments should remain in args.
	if len(args) == 0 || len(args) > 2 {
		usageFunc()
	}

	request := &api_client.ListLogEntriesRequest{}
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

	expectedParams := []string{}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	call := service.List(
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

func Logging_v2beta1_EntriesRead(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string

		if len(pathParams) != 0 {
			if strings.Contains("v2beta1/entries:read", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		commands_util.PrintRequestExample(&api_client.ReadLogEntriesRequest{})

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewEntriesService(api_service)

	args, flagValues, err := commands_util.ExtractFlagValues(args)
	if err != nil {
		return err
	}

	// Only positional arguments should remain in args.
	if len(args) == 0 || len(args) > 2 {
		usageFunc()
	}

	request := &api_client.ReadLogEntriesRequest{}
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

	expectedParams := []string{}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	call := service.Read(
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

func Logging_v2beta1_EntriesWrite(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string

		if len(pathParams) != 0 {
			if strings.Contains("v2beta1/entries:write", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		commands_util.PrintRequestExample(&api_client.WriteLogEntriesRequest{})

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewEntriesService(api_service)

	args, flagValues, err := commands_util.ExtractFlagValues(args)
	if err != nil {
		return err
	}

	// Only positional arguments should remain in args.
	if len(args) == 0 || len(args) > 2 {
		usageFunc()
	}

	request := &api_client.WriteLogEntriesRequest{}
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

	expectedParams := []string{}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	call := service.Write(
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

func Logging_v2beta1_GetLogsUsage(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string

		if len(pathParams) != 0 {
			if strings.Contains("v2beta1/logs_usage", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [--endTime=VALUE]"

		usageBits += " [--resourceName=VALUE]"

		usageBits += " [--resourceTier=VALUE]"

		usageBits += " [--startTime=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewV2beta1Service(api_service)

	queryParamNames := map[string]bool{
		"endTime":      false,
		"resourceName": false,
		"resourceTier": false,
		"startTime":    false,
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

	call := service.GetLogsUsage()

	// Set query parameters.
	if value, ok := flagValues["endTime"]; ok {
		query_endTime, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.EndTime(query_endTime)
	}
	if value, ok := flagValues["resourceName"]; ok {
		query_resourceName, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.ResourceName(query_resourceName)
	}
	if value, ok := flagValues["resourceTier"]; ok {
		query_resourceTier, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.ResourceTier(query_resourceTier)
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

func Logging_v2beta1_MonitoredResourceDescriptorsList(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string

		if len(pathParams) != 0 {
			if strings.Contains("v2beta1/monitoredResourceDescriptors", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [--pageSize=VALUE]"

		usageBits += " [--pageToken=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewMonitoredResourceDescriptorsService(api_service)

	queryParamNames := map[string]bool{
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

func Logging_v2beta1_OrganizationsLogsDelete(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("organizationsId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("logsId"))

		if len(pathParams) != 0 {
			if strings.Contains("v2beta1/organizations/{organizationsId}/logs/{logsId}", "+") {
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
	service := api_client.NewOrganizationsLogsService(api_service)

	// Only positional arguments should remain in args.
	if len(args) != 1 {
		usageFunc()
	}

	expectedParams := []string{
		"organizationsId",
		"logsId",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_organizationsId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_logsId, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}

	call := service.Delete(param_organizationsId, param_logsId)

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

func Logging_v2beta1_OrganizationsLogsList(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("organizationsId"))

		if len(pathParams) != 0 {
			if strings.Contains("v2beta1/organizations/{organizationsId}/logs", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [--pageSize=VALUE]"

		usageBits += " [--pageToken=VALUE]"

		usageBits += " [--resourceIndexPrefix=VALUE]"

		usageBits += " [--resourceType=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewOrganizationsLogsService(api_service)

	queryParamNames := map[string]bool{
		"pageSize":            false,
		"pageToken":           false,
		"resourceIndexPrefix": false,
		"resourceType":        false,
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
		"organizationsId",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_organizationsId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}

	call := service.List(param_organizationsId)

	// Set query parameters.
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
	if value, ok := flagValues["resourceIndexPrefix"]; ok {
		query_resourceIndexPrefix, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.ResourceIndexPrefix(query_resourceIndexPrefix)
	}
	if value, ok := flagValues["resourceType"]; ok {
		query_resourceType, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.ResourceType(query_resourceType)
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

func Logging_v2beta1_OrganizationsResourceKeysList(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("organizationsId"))

		if len(pathParams) != 0 {
			if strings.Contains("v2beta1/organizations/{organizationsId}/resourceKeys", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [--pageSize=VALUE]"

		usageBits += " [--pageToken=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewOrganizationsResourceKeysService(api_service)

	queryParamNames := map[string]bool{
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

	expectedParams := []string{
		"organizationsId",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_organizationsId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}

	call := service.List(param_organizationsId)

	// Set query parameters.
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

func Logging_v2beta1_OrganizationsResourceTypesValuesList(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("organizationsId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("resourceTypesId"))

		if len(pathParams) != 0 {
			if strings.Contains("v2beta1/organizations/{organizationsId}/resourceTypes/{resourceTypesId}/values", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [--depth=VALUE]"

		usageBits += " [--indexPrefix=VALUE]"

		usageBits += " [--pageSize=VALUE]"

		usageBits += " [--pageToken=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewOrganizationsResourceTypesValuesService(api_service)

	queryParamNames := map[string]bool{
		"depth":       false,
		"indexPrefix": false,
		"pageSize":    false,
		"pageToken":   false,
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
		"organizationsId",
		"resourceTypesId",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_organizationsId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_resourceTypesId, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}

	call := service.List(param_organizationsId, param_resourceTypesId)

	// Set query parameters.
	if value, ok := flagValues["depth"]; ok {
		query_depth, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.Depth(query_depth)
	}
	if value, ok := flagValues["indexPrefix"]; ok {
		query_indexPrefix, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.IndexPrefix(query_indexPrefix)
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

func Logging_v2beta1_ProjectsLogsDelete(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectsId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("logsId"))

		if len(pathParams) != 0 {
			if strings.Contains("v2beta1/projects/{projectsId}/logs/{logsId}", "+") {
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
	service := api_client.NewProjectsLogsService(api_service)

	// Only positional arguments should remain in args.
	if len(args) != 1 {
		usageFunc()
	}

	expectedParams := []string{
		"projectsId",
		"logsId",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectsId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_logsId, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}

	call := service.Delete(param_projectsId, param_logsId)

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

func Logging_v2beta1_ProjectsLogsList(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectsId"))

		if len(pathParams) != 0 {
			if strings.Contains("v2beta1/projects/{projectsId}/logs", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [--pageSize=VALUE]"

		usageBits += " [--pageToken=VALUE]"

		usageBits += " [--resourceIndexPrefix=VALUE]"

		usageBits += " [--resourceType=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsLogsService(api_service)

	queryParamNames := map[string]bool{
		"pageSize":            false,
		"pageToken":           false,
		"resourceIndexPrefix": false,
		"resourceType":        false,
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
	if value, ok := flagValues["resourceIndexPrefix"]; ok {
		query_resourceIndexPrefix, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.ResourceIndexPrefix(query_resourceIndexPrefix)
	}
	if value, ok := flagValues["resourceType"]; ok {
		query_resourceType, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.ResourceType(query_resourceType)
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

func Logging_v2beta1_ProjectsMetricsCreate(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectsId"))

		if len(pathParams) != 0 {
			if strings.Contains("v2beta1/projects/{projectsId}/metrics", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		commands_util.PrintRequestExample(&api_client.LogMetric{})

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsMetricsService(api_service)

	args, flagValues, err := commands_util.ExtractFlagValues(args)
	if err != nil {
		return err
	}

	// Only positional arguments should remain in args.
	if len(args) == 0 || len(args) > 2 {
		usageFunc()
	}

	request := &api_client.LogMetric{}
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

func Logging_v2beta1_ProjectsMetricsDelete(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectsId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("metricsId"))

		if len(pathParams) != 0 {
			if strings.Contains("v2beta1/projects/{projectsId}/metrics/{metricsId}", "+") {
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
	service := api_client.NewProjectsMetricsService(api_service)

	// Only positional arguments should remain in args.
	if len(args) != 1 {
		usageFunc()
	}

	expectedParams := []string{
		"projectsId",
		"metricsId",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectsId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_metricsId, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}

	call := service.Delete(param_projectsId, param_metricsId)

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

func Logging_v2beta1_ProjectsMetricsGet(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectsId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("metricsId"))

		if len(pathParams) != 0 {
			if strings.Contains("v2beta1/projects/{projectsId}/metrics/{metricsId}", "+") {
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
	service := api_client.NewProjectsMetricsService(api_service)

	// Only positional arguments should remain in args.
	if len(args) != 1 {
		usageFunc()
	}

	expectedParams := []string{
		"projectsId",
		"metricsId",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectsId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_metricsId, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}

	call := service.Get(param_projectsId, param_metricsId)

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

func Logging_v2beta1_ProjectsMetricsList(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectsId"))

		if len(pathParams) != 0 {
			if strings.Contains("v2beta1/projects/{projectsId}/metrics", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [--pageSize=VALUE]"

		usageBits += " [--pageToken=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsMetricsService(api_service)

	queryParamNames := map[string]bool{
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

func Logging_v2beta1_ProjectsMetricsUpdate(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectsId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("metricsId"))

		if len(pathParams) != 0 {
			if strings.Contains("v2beta1/projects/{projectsId}/metrics/{metricsId}", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		commands_util.PrintRequestExample(&api_client.LogMetric{})

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsMetricsService(api_service)

	args, flagValues, err := commands_util.ExtractFlagValues(args)
	if err != nil {
		return err
	}

	// Only positional arguments should remain in args.
	if len(args) == 0 || len(args) > 2 {
		usageFunc()
	}

	request := &api_client.LogMetric{}
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
		"metricsId",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectsId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_metricsId, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}

	call := service.Update(param_projectsId, param_metricsId,
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

func Logging_v2beta1_ProjectsResourceKeysList(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectsId"))

		if len(pathParams) != 0 {
			if strings.Contains("v2beta1/projects/{projectsId}/resourceKeys", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [--pageSize=VALUE]"

		usageBits += " [--pageToken=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsResourceKeysService(api_service)

	queryParamNames := map[string]bool{
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

func Logging_v2beta1_ProjectsResourceTypesValuesList(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectsId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("resourceTypesId"))

		if len(pathParams) != 0 {
			if strings.Contains("v2beta1/projects/{projectsId}/resourceTypes/{resourceTypesId}/values", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [--depth=VALUE]"

		usageBits += " [--indexPrefix=VALUE]"

		usageBits += " [--pageSize=VALUE]"

		usageBits += " [--pageToken=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsResourceTypesValuesService(api_service)

	queryParamNames := map[string]bool{
		"depth":       false,
		"indexPrefix": false,
		"pageSize":    false,
		"pageToken":   false,
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
		"resourceTypesId",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectsId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_resourceTypesId, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}

	call := service.List(param_projectsId, param_resourceTypesId)

	// Set query parameters.
	if value, ok := flagValues["depth"]; ok {
		query_depth, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.Depth(query_depth)
	}
	if value, ok := flagValues["indexPrefix"]; ok {
		query_indexPrefix, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.IndexPrefix(query_indexPrefix)
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

func Logging_v2beta1_ProjectsSinksCreate(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectsId"))

		if len(pathParams) != 0 {
			if strings.Contains("v2beta1/projects/{projectsId}/sinks", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		commands_util.PrintRequestExample(&api_client.LogSink{})

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsSinksService(api_service)

	args, flagValues, err := commands_util.ExtractFlagValues(args)
	if err != nil {
		return err
	}

	// Only positional arguments should remain in args.
	if len(args) == 0 || len(args) > 2 {
		usageFunc()
	}

	request := &api_client.LogSink{}
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

func Logging_v2beta1_ProjectsSinksDelete(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectsId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("sinksId"))

		if len(pathParams) != 0 {
			if strings.Contains("v2beta1/projects/{projectsId}/sinks/{sinksId}", "+") {
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
	service := api_client.NewProjectsSinksService(api_service)

	// Only positional arguments should remain in args.
	if len(args) != 1 {
		usageFunc()
	}

	expectedParams := []string{
		"projectsId",
		"sinksId",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectsId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_sinksId, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}

	call := service.Delete(param_projectsId, param_sinksId)

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

func Logging_v2beta1_ProjectsSinksGet(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectsId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("sinksId"))

		if len(pathParams) != 0 {
			if strings.Contains("v2beta1/projects/{projectsId}/sinks/{sinksId}", "+") {
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
	service := api_client.NewProjectsSinksService(api_service)

	// Only positional arguments should remain in args.
	if len(args) != 1 {
		usageFunc()
	}

	expectedParams := []string{
		"projectsId",
		"sinksId",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectsId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_sinksId, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}

	call := service.Get(param_projectsId, param_sinksId)

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

func Logging_v2beta1_ProjectsSinksList(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectsId"))

		if len(pathParams) != 0 {
			if strings.Contains("v2beta1/projects/{projectsId}/sinks", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [--pageSize=VALUE]"

		usageBits += " [--pageToken=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsSinksService(api_service)

	queryParamNames := map[string]bool{
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

func Logging_v2beta1_ProjectsSinksUpdate(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectsId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("sinksId"))

		if len(pathParams) != 0 {
			if strings.Contains("v2beta1/projects/{projectsId}/sinks/{sinksId}", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		commands_util.PrintRequestExample(&api_client.LogSink{})

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsSinksService(api_service)

	args, flagValues, err := commands_util.ExtractFlagValues(args)
	if err != nil {
		return err
	}

	// Only positional arguments should remain in args.
	if len(args) == 0 || len(args) > 2 {
		usageFunc()
	}

	request := &api_client.LogSink{}
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
		"sinksId",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectsId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_sinksId, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}

	call := service.Update(param_projectsId, param_sinksId,
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
