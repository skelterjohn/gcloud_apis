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

	api_client "github.com/skelterjohn/gcloud_apis/clients/storage/v1"
	"github.com/skelterjohn/gcloud_apis/commands_util"
)

// some generated code may not use all imported libs
var _ = fmt.Println
var _ = io.Copy
var _ = os.Stdin
var _ = strings.Split

func Storage_v1_BucketAccessControlsDelete(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("bucket"))
		pathParams = append(pathParams, commands_util.AngrySnakes("entity"))

		if len(pathParams) != 0 {
			if strings.Contains("b/{bucket}/acl/{entity}", "+") {
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
	service := api_client.NewBucketAccessControlsService(api_service)

	// Only positional arguments should remain in args.
	if len(args) != 1 {
		usageFunc()
	}

	expectedParams := []string{
		"bucket",
		"entity",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_bucket, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_entity, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}

	call := service.Delete(param_bucket, param_entity)

	if err := call.Do(); err != nil {
		return err
	}

	return nil
}

func Storage_v1_BucketAccessControlsGet(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("bucket"))
		pathParams = append(pathParams, commands_util.AngrySnakes("entity"))

		if len(pathParams) != 0 {
			if strings.Contains("b/{bucket}/acl/{entity}", "+") {
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
	service := api_client.NewBucketAccessControlsService(api_service)

	// Only positional arguments should remain in args.
	if len(args) != 1 {
		usageFunc()
	}

	expectedParams := []string{
		"bucket",
		"entity",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_bucket, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_entity, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}

	call := service.Get(param_bucket, param_entity)

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

func Storage_v1_BucketAccessControlsInsert(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("bucket"))

		if len(pathParams) != 0 {
			if strings.Contains("b/{bucket}/acl", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		commands_util.PrintRequestExample(&api_client.BucketAccessControl{})

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewBucketAccessControlsService(api_service)

	args, flagValues, err := commands_util.ExtractFlagValues(args)
	if err != nil {
		return err
	}

	// Only positional arguments should remain in args.
	if len(args) == 0 || len(args) > 2 {
		usageFunc()
	}

	request := &api_client.BucketAccessControl{}
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
		"bucket",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_bucket, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}

	call := service.Insert(param_bucket,
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

func Storage_v1_BucketAccessControlsList(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("bucket"))

		if len(pathParams) != 0 {
			if strings.Contains("b/{bucket}/acl", "+") {
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
	service := api_client.NewBucketAccessControlsService(api_service)

	// Only positional arguments should remain in args.
	if len(args) != 1 {
		usageFunc()
	}

	expectedParams := []string{
		"bucket",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_bucket, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}

	call := service.List(param_bucket)

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

func Storage_v1_BucketAccessControlsPatch(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("bucket"))
		pathParams = append(pathParams, commands_util.AngrySnakes("entity"))

		if len(pathParams) != 0 {
			if strings.Contains("b/{bucket}/acl/{entity}", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		commands_util.PrintRequestExample(&api_client.BucketAccessControl{})

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewBucketAccessControlsService(api_service)

	args, flagValues, err := commands_util.ExtractFlagValues(args)
	if err != nil {
		return err
	}

	// Only positional arguments should remain in args.
	if len(args) == 0 || len(args) > 2 {
		usageFunc()
	}

	request := &api_client.BucketAccessControl{}
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
		"bucket",
		"entity",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_bucket, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_entity, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}

	call := service.Patch(param_bucket, param_entity,
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

func Storage_v1_BucketAccessControlsUpdate(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("bucket"))
		pathParams = append(pathParams, commands_util.AngrySnakes("entity"))

		if len(pathParams) != 0 {
			if strings.Contains("b/{bucket}/acl/{entity}", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		commands_util.PrintRequestExample(&api_client.BucketAccessControl{})

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewBucketAccessControlsService(api_service)

	args, flagValues, err := commands_util.ExtractFlagValues(args)
	if err != nil {
		return err
	}

	// Only positional arguments should remain in args.
	if len(args) == 0 || len(args) > 2 {
		usageFunc()
	}

	request := &api_client.BucketAccessControl{}
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
		"bucket",
		"entity",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_bucket, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_entity, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}

	call := service.Update(param_bucket, param_entity,
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

func Storage_v1_BucketsDelete(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("bucket"))

		if len(pathParams) != 0 {
			if strings.Contains("b/{bucket}", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [--ifMetagenerationMatch=VALUE]"

		usageBits += " [--ifMetagenerationNotMatch=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewBucketsService(api_service)

	queryParamNames := map[string]bool{
		"ifMetagenerationMatch":    false,
		"ifMetagenerationNotMatch": false,
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
		"bucket",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_bucket, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}

	call := service.Delete(param_bucket)

	// Set query parameters.
	if value, ok := flagValues["ifMetagenerationMatch"]; ok {
		query_ifMetagenerationMatch, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.IfMetagenerationMatch(query_ifMetagenerationMatch)
	}
	if value, ok := flagValues["ifMetagenerationNotMatch"]; ok {
		query_ifMetagenerationNotMatch, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.IfMetagenerationNotMatch(query_ifMetagenerationNotMatch)
	}

	if err := call.Do(); err != nil {
		return err
	}

	return nil
}

func Storage_v1_BucketsGet(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("bucket"))

		if len(pathParams) != 0 {
			if strings.Contains("b/{bucket}", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [--ifMetagenerationMatch=VALUE]"

		usageBits += " [--ifMetagenerationNotMatch=VALUE]"

		usageBits += " [--projection=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewBucketsService(api_service)

	queryParamNames := map[string]bool{
		"ifMetagenerationMatch":    false,
		"ifMetagenerationNotMatch": false,
		"projection":               false,
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
		"bucket",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_bucket, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}

	call := service.Get(param_bucket)

	// Set query parameters.
	if value, ok := flagValues["ifMetagenerationMatch"]; ok {
		query_ifMetagenerationMatch, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.IfMetagenerationMatch(query_ifMetagenerationMatch)
	}
	if value, ok := flagValues["ifMetagenerationNotMatch"]; ok {
		query_ifMetagenerationNotMatch, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.IfMetagenerationNotMatch(query_ifMetagenerationNotMatch)
	}
	if value, ok := flagValues["projection"]; ok {
		query_projection, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.Projection(query_projection)
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

func Storage_v1_BucketsGetIamPolicy(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("bucket"))

		if len(pathParams) != 0 {
			if strings.Contains("b/{bucket}/iam", "+") {
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
	service := api_client.NewBucketsService(api_service)

	// Only positional arguments should remain in args.
	if len(args) != 1 {
		usageFunc()
	}

	expectedParams := []string{
		"bucket",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_bucket, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}

	call := service.GetIamPolicy(param_bucket)

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

func Storage_v1_BucketsInsert(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("project"))

		if len(pathParams) != 0 {
			if strings.Contains("b", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		usageBits += " [--predefinedAcl=VALUE]"

		usageBits += " [--predefinedDefaultObjectAcl=VALUE]"

		usageBits += " --project=VALUE"

		usageBits += " [--projection=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		commands_util.PrintRequestExample(&api_client.Bucket{})

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewBucketsService(api_service)

	queryParamNames := map[string]bool{
		"predefinedAcl":              false,
		"predefinedDefaultObjectAcl": false,
		"project":                    true,
		"projection":                 false,
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

	request := &api_client.Bucket{}
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
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_project, err := commands_util.ConvertValue_string(flagValues["project"])
	if err != nil {
		return err
	}

	call := service.Insert(param_project,
		request,
	)

	// Set query parameters.
	if value, ok := flagValues["predefinedAcl"]; ok {
		query_predefinedAcl, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.PredefinedAcl(query_predefinedAcl)
	}
	if value, ok := flagValues["predefinedDefaultObjectAcl"]; ok {
		query_predefinedDefaultObjectAcl, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.PredefinedDefaultObjectAcl(query_predefinedDefaultObjectAcl)
	}
	if value, ok := flagValues["projection"]; ok {
		query_projection, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.Projection(query_projection)
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

func Storage_v1_BucketsList(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("project"))

		if len(pathParams) != 0 {
			if strings.Contains("b", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [--maxResults=VALUE]"

		usageBits += " [--pageToken=VALUE]"

		usageBits += " [--prefix=VALUE]"

		usageBits += " --project=VALUE"

		usageBits += " [--projection=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewBucketsService(api_service)

	queryParamNames := map[string]bool{
		"maxResults": false,
		"pageToken":  false,
		"prefix":     false,
		"project":    true,
		"projection": false,
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

	param_project, err := commands_util.ConvertValue_string(flagValues["project"])
	if err != nil {
		return err
	}

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
	if value, ok := flagValues["prefix"]; ok {
		query_prefix, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.Prefix(query_prefix)
	}
	if value, ok := flagValues["projection"]; ok {
		query_projection, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.Projection(query_projection)
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

func Storage_v1_BucketsPatch(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("bucket"))

		if len(pathParams) != 0 {
			if strings.Contains("b/{bucket}", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		usageBits += " [--ifMetagenerationMatch=VALUE]"

		usageBits += " [--ifMetagenerationNotMatch=VALUE]"

		usageBits += " [--predefinedAcl=VALUE]"

		usageBits += " [--predefinedDefaultObjectAcl=VALUE]"

		usageBits += " [--projection=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		commands_util.PrintRequestExample(&api_client.Bucket{})

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewBucketsService(api_service)

	queryParamNames := map[string]bool{
		"ifMetagenerationMatch":      false,
		"ifMetagenerationNotMatch":   false,
		"predefinedAcl":              false,
		"predefinedDefaultObjectAcl": false,
		"projection":                 false,
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

	request := &api_client.Bucket{}
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
		"bucket",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_bucket, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}

	call := service.Patch(param_bucket,
		request,
	)

	// Set query parameters.
	if value, ok := flagValues["ifMetagenerationMatch"]; ok {
		query_ifMetagenerationMatch, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.IfMetagenerationMatch(query_ifMetagenerationMatch)
	}
	if value, ok := flagValues["ifMetagenerationNotMatch"]; ok {
		query_ifMetagenerationNotMatch, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.IfMetagenerationNotMatch(query_ifMetagenerationNotMatch)
	}
	if value, ok := flagValues["predefinedAcl"]; ok {
		query_predefinedAcl, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.PredefinedAcl(query_predefinedAcl)
	}
	if value, ok := flagValues["predefinedDefaultObjectAcl"]; ok {
		query_predefinedDefaultObjectAcl, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.PredefinedDefaultObjectAcl(query_predefinedDefaultObjectAcl)
	}
	if value, ok := flagValues["projection"]; ok {
		query_projection, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.Projection(query_projection)
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

func Storage_v1_BucketsSetIamPolicy(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("bucket"))

		if len(pathParams) != 0 {
			if strings.Contains("b/{bucket}/iam", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		commands_util.PrintRequestExample(&api_client.Policy{})

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewBucketsService(api_service)

	args, flagValues, err := commands_util.ExtractFlagValues(args)
	if err != nil {
		return err
	}

	// Only positional arguments should remain in args.
	if len(args) == 0 || len(args) > 2 {
		usageFunc()
	}

	request := &api_client.Policy{}
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
		"bucket",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_bucket, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}

	call := service.SetIamPolicy(param_bucket,
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

func Storage_v1_BucketsTestIamPermissions(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("bucket"))
		pathParams = append(pathParams, commands_util.AngrySnakes("permissions"))

		if len(pathParams) != 0 {
			if strings.Contains("b/{bucket}/iam/testPermissions", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " --permissions=VALUE"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewBucketsService(api_service)

	queryParamNames := map[string]bool{
		"permissions": true,
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
		"bucket",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_bucket, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_permissions, err := commands_util.ConvertValue_strings(flagValues["permissions"])
	if err != nil {
		return err
	}

	call := service.TestIamPermissions(param_bucket, param_permissions)

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

func Storage_v1_BucketsUpdate(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("bucket"))

		if len(pathParams) != 0 {
			if strings.Contains("b/{bucket}", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		usageBits += " [--ifMetagenerationMatch=VALUE]"

		usageBits += " [--ifMetagenerationNotMatch=VALUE]"

		usageBits += " [--predefinedAcl=VALUE]"

		usageBits += " [--predefinedDefaultObjectAcl=VALUE]"

		usageBits += " [--projection=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		commands_util.PrintRequestExample(&api_client.Bucket{})

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewBucketsService(api_service)

	queryParamNames := map[string]bool{
		"ifMetagenerationMatch":      false,
		"ifMetagenerationNotMatch":   false,
		"predefinedAcl":              false,
		"predefinedDefaultObjectAcl": false,
		"projection":                 false,
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

	request := &api_client.Bucket{}
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
		"bucket",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_bucket, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}

	call := service.Update(param_bucket,
		request,
	)

	// Set query parameters.
	if value, ok := flagValues["ifMetagenerationMatch"]; ok {
		query_ifMetagenerationMatch, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.IfMetagenerationMatch(query_ifMetagenerationMatch)
	}
	if value, ok := flagValues["ifMetagenerationNotMatch"]; ok {
		query_ifMetagenerationNotMatch, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.IfMetagenerationNotMatch(query_ifMetagenerationNotMatch)
	}
	if value, ok := flagValues["predefinedAcl"]; ok {
		query_predefinedAcl, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.PredefinedAcl(query_predefinedAcl)
	}
	if value, ok := flagValues["predefinedDefaultObjectAcl"]; ok {
		query_predefinedDefaultObjectAcl, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.PredefinedDefaultObjectAcl(query_predefinedDefaultObjectAcl)
	}
	if value, ok := flagValues["projection"]; ok {
		query_projection, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.Projection(query_projection)
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

func Storage_v1_ChannelsStop(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string

		if len(pathParams) != 0 {
			if strings.Contains("channels/stop", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		commands_util.PrintRequestExample(&api_client.Channel{})

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewChannelsService(api_service)

	args, flagValues, err := commands_util.ExtractFlagValues(args)
	if err != nil {
		return err
	}

	// Only positional arguments should remain in args.
	if len(args) == 0 || len(args) > 2 {
		usageFunc()
	}

	request := &api_client.Channel{}
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

	call := service.Stop(
		request,
	)

	if err := call.Do(); err != nil {
		return err
	}

	return nil
}

func Storage_v1_DefaultObjectAccessControlsDelete(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("bucket"))
		pathParams = append(pathParams, commands_util.AngrySnakes("entity"))

		if len(pathParams) != 0 {
			if strings.Contains("b/{bucket}/defaultObjectAcl/{entity}", "+") {
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
	service := api_client.NewDefaultObjectAccessControlsService(api_service)

	// Only positional arguments should remain in args.
	if len(args) != 1 {
		usageFunc()
	}

	expectedParams := []string{
		"bucket",
		"entity",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_bucket, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_entity, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}

	call := service.Delete(param_bucket, param_entity)

	if err := call.Do(); err != nil {
		return err
	}

	return nil
}

func Storage_v1_DefaultObjectAccessControlsGet(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("bucket"))
		pathParams = append(pathParams, commands_util.AngrySnakes("entity"))

		if len(pathParams) != 0 {
			if strings.Contains("b/{bucket}/defaultObjectAcl/{entity}", "+") {
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
	service := api_client.NewDefaultObjectAccessControlsService(api_service)

	// Only positional arguments should remain in args.
	if len(args) != 1 {
		usageFunc()
	}

	expectedParams := []string{
		"bucket",
		"entity",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_bucket, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_entity, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}

	call := service.Get(param_bucket, param_entity)

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

func Storage_v1_DefaultObjectAccessControlsInsert(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("bucket"))

		if len(pathParams) != 0 {
			if strings.Contains("b/{bucket}/defaultObjectAcl", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		commands_util.PrintRequestExample(&api_client.ObjectAccessControl{})

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewDefaultObjectAccessControlsService(api_service)

	args, flagValues, err := commands_util.ExtractFlagValues(args)
	if err != nil {
		return err
	}

	// Only positional arguments should remain in args.
	if len(args) == 0 || len(args) > 2 {
		usageFunc()
	}

	request := &api_client.ObjectAccessControl{}
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
		"bucket",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_bucket, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}

	call := service.Insert(param_bucket,
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

func Storage_v1_DefaultObjectAccessControlsList(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("bucket"))

		if len(pathParams) != 0 {
			if strings.Contains("b/{bucket}/defaultObjectAcl", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [--ifMetagenerationMatch=VALUE]"

		usageBits += " [--ifMetagenerationNotMatch=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewDefaultObjectAccessControlsService(api_service)

	queryParamNames := map[string]bool{
		"ifMetagenerationMatch":    false,
		"ifMetagenerationNotMatch": false,
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
		"bucket",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_bucket, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}

	call := service.List(param_bucket)

	// Set query parameters.
	if value, ok := flagValues["ifMetagenerationMatch"]; ok {
		query_ifMetagenerationMatch, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.IfMetagenerationMatch(query_ifMetagenerationMatch)
	}
	if value, ok := flagValues["ifMetagenerationNotMatch"]; ok {
		query_ifMetagenerationNotMatch, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.IfMetagenerationNotMatch(query_ifMetagenerationNotMatch)
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

func Storage_v1_DefaultObjectAccessControlsPatch(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("bucket"))
		pathParams = append(pathParams, commands_util.AngrySnakes("entity"))

		if len(pathParams) != 0 {
			if strings.Contains("b/{bucket}/defaultObjectAcl/{entity}", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		commands_util.PrintRequestExample(&api_client.ObjectAccessControl{})

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewDefaultObjectAccessControlsService(api_service)

	args, flagValues, err := commands_util.ExtractFlagValues(args)
	if err != nil {
		return err
	}

	// Only positional arguments should remain in args.
	if len(args) == 0 || len(args) > 2 {
		usageFunc()
	}

	request := &api_client.ObjectAccessControl{}
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
		"bucket",
		"entity",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_bucket, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_entity, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}

	call := service.Patch(param_bucket, param_entity,
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

func Storage_v1_DefaultObjectAccessControlsUpdate(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("bucket"))
		pathParams = append(pathParams, commands_util.AngrySnakes("entity"))

		if len(pathParams) != 0 {
			if strings.Contains("b/{bucket}/defaultObjectAcl/{entity}", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		commands_util.PrintRequestExample(&api_client.ObjectAccessControl{})

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewDefaultObjectAccessControlsService(api_service)

	args, flagValues, err := commands_util.ExtractFlagValues(args)
	if err != nil {
		return err
	}

	// Only positional arguments should remain in args.
	if len(args) == 0 || len(args) > 2 {
		usageFunc()
	}

	request := &api_client.ObjectAccessControl{}
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
		"bucket",
		"entity",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_bucket, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_entity, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}

	call := service.Update(param_bucket, param_entity,
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

func Storage_v1_NotificationsDelete(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("bucket"))
		pathParams = append(pathParams, commands_util.AngrySnakes("notification"))

		if len(pathParams) != 0 {
			if strings.Contains("b/{bucket}/notificationConfigs/{notification}", "+") {
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
	service := api_client.NewNotificationsService(api_service)

	// Only positional arguments should remain in args.
	if len(args) != 1 {
		usageFunc()
	}

	expectedParams := []string{
		"bucket",
		"notification",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_bucket, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_notification, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}

	call := service.Delete(param_bucket, param_notification)

	if err := call.Do(); err != nil {
		return err
	}

	return nil
}

func Storage_v1_NotificationsGet(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("bucket"))
		pathParams = append(pathParams, commands_util.AngrySnakes("notification"))

		if len(pathParams) != 0 {
			if strings.Contains("b/{bucket}/notificationConfigs/{notification}", "+") {
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
	service := api_client.NewNotificationsService(api_service)

	// Only positional arguments should remain in args.
	if len(args) != 1 {
		usageFunc()
	}

	expectedParams := []string{
		"bucket",
		"notification",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_bucket, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_notification, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}

	call := service.Get(param_bucket, param_notification)

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

func Storage_v1_NotificationsInsert(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("bucket"))

		if len(pathParams) != 0 {
			if strings.Contains("b/{bucket}/notificationConfigs", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		commands_util.PrintRequestExample(&api_client.Notification{})

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewNotificationsService(api_service)

	args, flagValues, err := commands_util.ExtractFlagValues(args)
	if err != nil {
		return err
	}

	// Only positional arguments should remain in args.
	if len(args) == 0 || len(args) > 2 {
		usageFunc()
	}

	request := &api_client.Notification{}
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
		"bucket",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_bucket, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}

	call := service.Insert(param_bucket,
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

func Storage_v1_NotificationsList(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("bucket"))

		if len(pathParams) != 0 {
			if strings.Contains("b/{bucket}/notificationConfigs", "+") {
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
	service := api_client.NewNotificationsService(api_service)

	// Only positional arguments should remain in args.
	if len(args) != 1 {
		usageFunc()
	}

	expectedParams := []string{
		"bucket",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_bucket, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}

	call := service.List(param_bucket)

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

func Storage_v1_ObjectAccessControlsDelete(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("bucket"))
		pathParams = append(pathParams, commands_util.AngrySnakes("object"))
		pathParams = append(pathParams, commands_util.AngrySnakes("entity"))

		if len(pathParams) != 0 {
			if strings.Contains("b/{bucket}/o/{object}/acl/{entity}", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [--generation=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewObjectAccessControlsService(api_service)

	queryParamNames := map[string]bool{
		"generation": false,
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
		"bucket",
		"object",
		"entity",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_bucket, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_object, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}
	param_entity, err := commands_util.ConvertValue_string(paramValues[2])
	if err != nil {
		return err
	}

	call := service.Delete(param_bucket, param_object, param_entity)

	// Set query parameters.
	if value, ok := flagValues["generation"]; ok {
		query_generation, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.Generation(query_generation)
	}

	if err := call.Do(); err != nil {
		return err
	}

	return nil
}

func Storage_v1_ObjectAccessControlsGet(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("bucket"))
		pathParams = append(pathParams, commands_util.AngrySnakes("object"))
		pathParams = append(pathParams, commands_util.AngrySnakes("entity"))

		if len(pathParams) != 0 {
			if strings.Contains("b/{bucket}/o/{object}/acl/{entity}", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [--generation=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewObjectAccessControlsService(api_service)

	queryParamNames := map[string]bool{
		"generation": false,
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
		"bucket",
		"object",
		"entity",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_bucket, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_object, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}
	param_entity, err := commands_util.ConvertValue_string(paramValues[2])
	if err != nil {
		return err
	}

	call := service.Get(param_bucket, param_object, param_entity)

	// Set query parameters.
	if value, ok := flagValues["generation"]; ok {
		query_generation, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.Generation(query_generation)
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

func Storage_v1_ObjectAccessControlsInsert(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("bucket"))
		pathParams = append(pathParams, commands_util.AngrySnakes("object"))

		if len(pathParams) != 0 {
			if strings.Contains("b/{bucket}/o/{object}/acl", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		usageBits += " [--generation=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		commands_util.PrintRequestExample(&api_client.ObjectAccessControl{})

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewObjectAccessControlsService(api_service)

	queryParamNames := map[string]bool{
		"generation": false,
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

	request := &api_client.ObjectAccessControl{}
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
		"bucket",
		"object",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_bucket, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_object, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}

	call := service.Insert(param_bucket, param_object,
		request,
	)

	// Set query parameters.
	if value, ok := flagValues["generation"]; ok {
		query_generation, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.Generation(query_generation)
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

func Storage_v1_ObjectAccessControlsList(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("bucket"))
		pathParams = append(pathParams, commands_util.AngrySnakes("object"))

		if len(pathParams) != 0 {
			if strings.Contains("b/{bucket}/o/{object}/acl", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [--generation=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewObjectAccessControlsService(api_service)

	queryParamNames := map[string]bool{
		"generation": false,
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
		"bucket",
		"object",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_bucket, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_object, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}

	call := service.List(param_bucket, param_object)

	// Set query parameters.
	if value, ok := flagValues["generation"]; ok {
		query_generation, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.Generation(query_generation)
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

func Storage_v1_ObjectAccessControlsPatch(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("bucket"))
		pathParams = append(pathParams, commands_util.AngrySnakes("object"))
		pathParams = append(pathParams, commands_util.AngrySnakes("entity"))

		if len(pathParams) != 0 {
			if strings.Contains("b/{bucket}/o/{object}/acl/{entity}", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		usageBits += " [--generation=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		commands_util.PrintRequestExample(&api_client.ObjectAccessControl{})

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewObjectAccessControlsService(api_service)

	queryParamNames := map[string]bool{
		"generation": false,
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

	request := &api_client.ObjectAccessControl{}
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
		"bucket",
		"object",
		"entity",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_bucket, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_object, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}
	param_entity, err := commands_util.ConvertValue_string(paramValues[2])
	if err != nil {
		return err
	}

	call := service.Patch(param_bucket, param_object, param_entity,
		request,
	)

	// Set query parameters.
	if value, ok := flagValues["generation"]; ok {
		query_generation, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.Generation(query_generation)
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

func Storage_v1_ObjectAccessControlsUpdate(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("bucket"))
		pathParams = append(pathParams, commands_util.AngrySnakes("object"))
		pathParams = append(pathParams, commands_util.AngrySnakes("entity"))

		if len(pathParams) != 0 {
			if strings.Contains("b/{bucket}/o/{object}/acl/{entity}", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		usageBits += " [--generation=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		commands_util.PrintRequestExample(&api_client.ObjectAccessControl{})

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewObjectAccessControlsService(api_service)

	queryParamNames := map[string]bool{
		"generation": false,
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

	request := &api_client.ObjectAccessControl{}
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
		"bucket",
		"object",
		"entity",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_bucket, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_object, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}
	param_entity, err := commands_util.ConvertValue_string(paramValues[2])
	if err != nil {
		return err
	}

	call := service.Update(param_bucket, param_object, param_entity,
		request,
	)

	// Set query parameters.
	if value, ok := flagValues["generation"]; ok {
		query_generation, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.Generation(query_generation)
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

func Storage_v1_ObjectsCompose(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("destinationBucket"))
		pathParams = append(pathParams, commands_util.AngrySnakes("destinationObject"))

		if len(pathParams) != 0 {
			if strings.Contains("b/{destinationBucket}/o/{destinationObject}/compose", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		usageBits += " [--destinationPredefinedAcl=VALUE]"

		usageBits += " [--ifGenerationMatch=VALUE]"

		usageBits += " [--ifMetagenerationMatch=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		commands_util.PrintRequestExample(&api_client.ComposeRequest{})

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewObjectsService(api_service)

	queryParamNames := map[string]bool{
		"destinationPredefinedAcl": false,
		"ifGenerationMatch":        false,
		"ifMetagenerationMatch":    false,
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

	request := &api_client.ComposeRequest{}
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
		"destinationBucket",
		"destinationObject",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_destinationBucket, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_destinationObject, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}

	call := service.Compose(param_destinationBucket, param_destinationObject,
		request,
	)

	// Set query parameters.
	if value, ok := flagValues["destinationPredefinedAcl"]; ok {
		query_destinationPredefinedAcl, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.DestinationPredefinedAcl(query_destinationPredefinedAcl)
	}
	if value, ok := flagValues["ifGenerationMatch"]; ok {
		query_ifGenerationMatch, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.IfGenerationMatch(query_ifGenerationMatch)
	}
	if value, ok := flagValues["ifMetagenerationMatch"]; ok {
		query_ifMetagenerationMatch, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.IfMetagenerationMatch(query_ifMetagenerationMatch)
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

func Storage_v1_ObjectsCopy(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("sourceBucket"))
		pathParams = append(pathParams, commands_util.AngrySnakes("sourceObject"))
		pathParams = append(pathParams, commands_util.AngrySnakes("destinationBucket"))
		pathParams = append(pathParams, commands_util.AngrySnakes("destinationObject"))

		if len(pathParams) != 0 {
			if strings.Contains("b/{sourceBucket}/o/{sourceObject}/copyTo/b/{destinationBucket}/o/{destinationObject}", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		usageBits += " [--destinationPredefinedAcl=VALUE]"

		usageBits += " [--ifGenerationMatch=VALUE]"

		usageBits += " [--ifGenerationNotMatch=VALUE]"

		usageBits += " [--ifMetagenerationMatch=VALUE]"

		usageBits += " [--ifMetagenerationNotMatch=VALUE]"

		usageBits += " [--ifSourceGenerationMatch=VALUE]"

		usageBits += " [--ifSourceGenerationNotMatch=VALUE]"

		usageBits += " [--ifSourceMetagenerationMatch=VALUE]"

		usageBits += " [--ifSourceMetagenerationNotMatch=VALUE]"

		usageBits += " [--projection=VALUE]"

		usageBits += " [--sourceGeneration=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		commands_util.PrintRequestExample(&api_client.Object{})

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewObjectsService(api_service)

	queryParamNames := map[string]bool{
		"destinationPredefinedAcl":       false,
		"ifGenerationMatch":              false,
		"ifGenerationNotMatch":           false,
		"ifMetagenerationMatch":          false,
		"ifMetagenerationNotMatch":       false,
		"ifSourceGenerationMatch":        false,
		"ifSourceGenerationNotMatch":     false,
		"ifSourceMetagenerationMatch":    false,
		"ifSourceMetagenerationNotMatch": false,
		"projection":                     false,
		"sourceGeneration":               false,
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

	request := &api_client.Object{}
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
		"sourceBucket",
		"sourceObject",
		"destinationBucket",
		"destinationObject",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_sourceBucket, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_sourceObject, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}
	param_destinationBucket, err := commands_util.ConvertValue_string(paramValues[2])
	if err != nil {
		return err
	}
	param_destinationObject, err := commands_util.ConvertValue_string(paramValues[3])
	if err != nil {
		return err
	}

	call := service.Copy(param_sourceBucket, param_sourceObject, param_destinationBucket, param_destinationObject,
		request,
	)

	// Set query parameters.
	if value, ok := flagValues["destinationPredefinedAcl"]; ok {
		query_destinationPredefinedAcl, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.DestinationPredefinedAcl(query_destinationPredefinedAcl)
	}
	if value, ok := flagValues["ifGenerationMatch"]; ok {
		query_ifGenerationMatch, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.IfGenerationMatch(query_ifGenerationMatch)
	}
	if value, ok := flagValues["ifGenerationNotMatch"]; ok {
		query_ifGenerationNotMatch, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.IfGenerationNotMatch(query_ifGenerationNotMatch)
	}
	if value, ok := flagValues["ifMetagenerationMatch"]; ok {
		query_ifMetagenerationMatch, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.IfMetagenerationMatch(query_ifMetagenerationMatch)
	}
	if value, ok := flagValues["ifMetagenerationNotMatch"]; ok {
		query_ifMetagenerationNotMatch, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.IfMetagenerationNotMatch(query_ifMetagenerationNotMatch)
	}
	if value, ok := flagValues["ifSourceGenerationMatch"]; ok {
		query_ifSourceGenerationMatch, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.IfSourceGenerationMatch(query_ifSourceGenerationMatch)
	}
	if value, ok := flagValues["ifSourceGenerationNotMatch"]; ok {
		query_ifSourceGenerationNotMatch, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.IfSourceGenerationNotMatch(query_ifSourceGenerationNotMatch)
	}
	if value, ok := flagValues["ifSourceMetagenerationMatch"]; ok {
		query_ifSourceMetagenerationMatch, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.IfSourceMetagenerationMatch(query_ifSourceMetagenerationMatch)
	}
	if value, ok := flagValues["ifSourceMetagenerationNotMatch"]; ok {
		query_ifSourceMetagenerationNotMatch, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.IfSourceMetagenerationNotMatch(query_ifSourceMetagenerationNotMatch)
	}
	if value, ok := flagValues["projection"]; ok {
		query_projection, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.Projection(query_projection)
	}
	if value, ok := flagValues["sourceGeneration"]; ok {
		query_sourceGeneration, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.SourceGeneration(query_sourceGeneration)
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

func Storage_v1_ObjectsDelete(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("bucket"))
		pathParams = append(pathParams, commands_util.AngrySnakes("object"))

		if len(pathParams) != 0 {
			if strings.Contains("b/{bucket}/o/{object}", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [--generation=VALUE]"

		usageBits += " [--ifGenerationMatch=VALUE]"

		usageBits += " [--ifGenerationNotMatch=VALUE]"

		usageBits += " [--ifMetagenerationMatch=VALUE]"

		usageBits += " [--ifMetagenerationNotMatch=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewObjectsService(api_service)

	queryParamNames := map[string]bool{
		"generation":               false,
		"ifGenerationMatch":        false,
		"ifGenerationNotMatch":     false,
		"ifMetagenerationMatch":    false,
		"ifMetagenerationNotMatch": false,
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
		"bucket",
		"object",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_bucket, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_object, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}

	call := service.Delete(param_bucket, param_object)

	// Set query parameters.
	if value, ok := flagValues["generation"]; ok {
		query_generation, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.Generation(query_generation)
	}
	if value, ok := flagValues["ifGenerationMatch"]; ok {
		query_ifGenerationMatch, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.IfGenerationMatch(query_ifGenerationMatch)
	}
	if value, ok := flagValues["ifGenerationNotMatch"]; ok {
		query_ifGenerationNotMatch, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.IfGenerationNotMatch(query_ifGenerationNotMatch)
	}
	if value, ok := flagValues["ifMetagenerationMatch"]; ok {
		query_ifMetagenerationMatch, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.IfMetagenerationMatch(query_ifMetagenerationMatch)
	}
	if value, ok := flagValues["ifMetagenerationNotMatch"]; ok {
		query_ifMetagenerationNotMatch, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.IfMetagenerationNotMatch(query_ifMetagenerationNotMatch)
	}

	if err := call.Do(); err != nil {
		return err
	}

	return nil
}

func Storage_v1_ObjectsGet(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("bucket"))
		pathParams = append(pathParams, commands_util.AngrySnakes("object"))

		if len(pathParams) != 0 {
			if strings.Contains("b/{bucket}/o/{object}", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [--generation=VALUE]"

		usageBits += " [--ifGenerationMatch=VALUE]"

		usageBits += " [--ifGenerationNotMatch=VALUE]"

		usageBits += " [--ifMetagenerationMatch=VALUE]"

		usageBits += " [--ifMetagenerationNotMatch=VALUE]"

		usageBits += " [--projection=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewObjectsService(api_service)

	queryParamNames := map[string]bool{
		"generation":               false,
		"ifGenerationMatch":        false,
		"ifGenerationNotMatch":     false,
		"ifMetagenerationMatch":    false,
		"ifMetagenerationNotMatch": false,
		"projection":               false,
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
		"bucket",
		"object",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_bucket, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_object, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}

	call := service.Get(param_bucket, param_object)

	// Set query parameters.
	if value, ok := flagValues["generation"]; ok {
		query_generation, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.Generation(query_generation)
	}
	if value, ok := flagValues["ifGenerationMatch"]; ok {
		query_ifGenerationMatch, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.IfGenerationMatch(query_ifGenerationMatch)
	}
	if value, ok := flagValues["ifGenerationNotMatch"]; ok {
		query_ifGenerationNotMatch, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.IfGenerationNotMatch(query_ifGenerationNotMatch)
	}
	if value, ok := flagValues["ifMetagenerationMatch"]; ok {
		query_ifMetagenerationMatch, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.IfMetagenerationMatch(query_ifMetagenerationMatch)
	}
	if value, ok := flagValues["ifMetagenerationNotMatch"]; ok {
		query_ifMetagenerationNotMatch, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.IfMetagenerationNotMatch(query_ifMetagenerationNotMatch)
	}
	if value, ok := flagValues["projection"]; ok {
		query_projection, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.Projection(query_projection)
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

func Storage_v1_ObjectsGetIamPolicy(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("bucket"))
		pathParams = append(pathParams, commands_util.AngrySnakes("object"))

		if len(pathParams) != 0 {
			if strings.Contains("b/{bucket}/o/{object}/iam", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [--generation=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewObjectsService(api_service)

	queryParamNames := map[string]bool{
		"generation": false,
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
		"bucket",
		"object",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_bucket, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_object, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}

	call := service.GetIamPolicy(param_bucket, param_object)

	// Set query parameters.
	if value, ok := flagValues["generation"]; ok {
		query_generation, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.Generation(query_generation)
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

func Storage_v1_ObjectsInsert(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("bucket"))

		if len(pathParams) != 0 {
			if strings.Contains("b/{bucket}/o", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		usageBits += " [--contentEncoding=VALUE]"

		usageBits += " [--ifGenerationMatch=VALUE]"

		usageBits += " [--ifGenerationNotMatch=VALUE]"

		usageBits += " [--ifMetagenerationMatch=VALUE]"

		usageBits += " [--ifMetagenerationNotMatch=VALUE]"

		usageBits += " [--name=VALUE]"

		usageBits += " [--predefinedAcl=VALUE]"

		usageBits += " [--projection=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		commands_util.PrintRequestExample(&api_client.Object{})

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewObjectsService(api_service)

	queryParamNames := map[string]bool{
		"contentEncoding":          false,
		"ifGenerationMatch":        false,
		"ifGenerationNotMatch":     false,
		"ifMetagenerationMatch":    false,
		"ifMetagenerationNotMatch": false,
		"name":          false,
		"predefinedAcl": false,
		"projection":    false,
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

	request := &api_client.Object{}
	if len(args) == 2 {
		err = commands_util.PopulateRequestFromFilename(&request, args[1])
		if err != nil {
			return err
		}
	}

	var media io.Reader
	if len(args) == 3 {
		if args[2] == "-" {
			media = os.Stdin
		} else {
			media, err = os.Open(args[2])
			if err != nil {
				return err
			}
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
		"bucket",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_bucket, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}

	call := service.Insert(param_bucket,
		request,
	)

	// Set query parameters.
	if value, ok := flagValues["contentEncoding"]; ok {
		query_contentEncoding, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.ContentEncoding(query_contentEncoding)
	}
	if value, ok := flagValues["ifGenerationMatch"]; ok {
		query_ifGenerationMatch, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.IfGenerationMatch(query_ifGenerationMatch)
	}
	if value, ok := flagValues["ifGenerationNotMatch"]; ok {
		query_ifGenerationNotMatch, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.IfGenerationNotMatch(query_ifGenerationNotMatch)
	}
	if value, ok := flagValues["ifMetagenerationMatch"]; ok {
		query_ifMetagenerationMatch, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.IfMetagenerationMatch(query_ifMetagenerationMatch)
	}
	if value, ok := flagValues["ifMetagenerationNotMatch"]; ok {
		query_ifMetagenerationNotMatch, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.IfMetagenerationNotMatch(query_ifMetagenerationNotMatch)
	}
	if value, ok := flagValues["name"]; ok {
		query_name, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.Name(query_name)
	}
	if value, ok := flagValues["predefinedAcl"]; ok {
		query_predefinedAcl, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.PredefinedAcl(query_predefinedAcl)
	}
	if value, ok := flagValues["projection"]; ok {
		query_projection, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.Projection(query_projection)
	}

	if media != nil {
		call.Media(media)
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

func Storage_v1_ObjectsList(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("bucket"))

		if len(pathParams) != 0 {
			if strings.Contains("b/{bucket}/o", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [--delimiter=VALUE]"

		usageBits += " [--maxResults=VALUE]"

		usageBits += " [--pageToken=VALUE]"

		usageBits += " [--prefix=VALUE]"

		usageBits += " [--projection=VALUE]"

		usageBits += " [--versions=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewObjectsService(api_service)

	queryParamNames := map[string]bool{
		"delimiter":  false,
		"maxResults": false,
		"pageToken":  false,
		"prefix":     false,
		"projection": false,
		"versions":   false,
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
		"bucket",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_bucket, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}

	call := service.List(param_bucket)

	// Set query parameters.
	if value, ok := flagValues["delimiter"]; ok {
		query_delimiter, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.Delimiter(query_delimiter)
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
	if value, ok := flagValues["prefix"]; ok {
		query_prefix, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.Prefix(query_prefix)
	}
	if value, ok := flagValues["projection"]; ok {
		query_projection, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.Projection(query_projection)
	}
	if value, ok := flagValues["versions"]; ok {
		query_versions, err := commands_util.ConvertValue_bool(value)
		if err != nil {
			return err
		}
		call.Versions(query_versions)
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

func Storage_v1_ObjectsPatch(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("bucket"))
		pathParams = append(pathParams, commands_util.AngrySnakes("object"))

		if len(pathParams) != 0 {
			if strings.Contains("b/{bucket}/o/{object}", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		usageBits += " [--generation=VALUE]"

		usageBits += " [--ifGenerationMatch=VALUE]"

		usageBits += " [--ifGenerationNotMatch=VALUE]"

		usageBits += " [--ifMetagenerationMatch=VALUE]"

		usageBits += " [--ifMetagenerationNotMatch=VALUE]"

		usageBits += " [--predefinedAcl=VALUE]"

		usageBits += " [--projection=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		commands_util.PrintRequestExample(&api_client.Object{})

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewObjectsService(api_service)

	queryParamNames := map[string]bool{
		"generation":               false,
		"ifGenerationMatch":        false,
		"ifGenerationNotMatch":     false,
		"ifMetagenerationMatch":    false,
		"ifMetagenerationNotMatch": false,
		"predefinedAcl":            false,
		"projection":               false,
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

	request := &api_client.Object{}
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
		"bucket",
		"object",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_bucket, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_object, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}

	call := service.Patch(param_bucket, param_object,
		request,
	)

	// Set query parameters.
	if value, ok := flagValues["generation"]; ok {
		query_generation, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.Generation(query_generation)
	}
	if value, ok := flagValues["ifGenerationMatch"]; ok {
		query_ifGenerationMatch, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.IfGenerationMatch(query_ifGenerationMatch)
	}
	if value, ok := flagValues["ifGenerationNotMatch"]; ok {
		query_ifGenerationNotMatch, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.IfGenerationNotMatch(query_ifGenerationNotMatch)
	}
	if value, ok := flagValues["ifMetagenerationMatch"]; ok {
		query_ifMetagenerationMatch, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.IfMetagenerationMatch(query_ifMetagenerationMatch)
	}
	if value, ok := flagValues["ifMetagenerationNotMatch"]; ok {
		query_ifMetagenerationNotMatch, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.IfMetagenerationNotMatch(query_ifMetagenerationNotMatch)
	}
	if value, ok := flagValues["predefinedAcl"]; ok {
		query_predefinedAcl, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.PredefinedAcl(query_predefinedAcl)
	}
	if value, ok := flagValues["projection"]; ok {
		query_projection, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.Projection(query_projection)
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

func Storage_v1_ObjectsRewrite(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("sourceBucket"))
		pathParams = append(pathParams, commands_util.AngrySnakes("sourceObject"))
		pathParams = append(pathParams, commands_util.AngrySnakes("destinationBucket"))
		pathParams = append(pathParams, commands_util.AngrySnakes("destinationObject"))

		if len(pathParams) != 0 {
			if strings.Contains("b/{sourceBucket}/o/{sourceObject}/rewriteTo/b/{destinationBucket}/o/{destinationObject}", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		usageBits += " [--destinationPredefinedAcl=VALUE]"

		usageBits += " [--ifGenerationMatch=VALUE]"

		usageBits += " [--ifGenerationNotMatch=VALUE]"

		usageBits += " [--ifMetagenerationMatch=VALUE]"

		usageBits += " [--ifMetagenerationNotMatch=VALUE]"

		usageBits += " [--ifSourceGenerationMatch=VALUE]"

		usageBits += " [--ifSourceGenerationNotMatch=VALUE]"

		usageBits += " [--ifSourceMetagenerationMatch=VALUE]"

		usageBits += " [--ifSourceMetagenerationNotMatch=VALUE]"

		usageBits += " [--maxBytesRewrittenPerCall=VALUE]"

		usageBits += " [--projection=VALUE]"

		usageBits += " [--rewriteToken=VALUE]"

		usageBits += " [--sourceGeneration=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		commands_util.PrintRequestExample(&api_client.Object{})

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewObjectsService(api_service)

	queryParamNames := map[string]bool{
		"destinationPredefinedAcl":       false,
		"ifGenerationMatch":              false,
		"ifGenerationNotMatch":           false,
		"ifMetagenerationMatch":          false,
		"ifMetagenerationNotMatch":       false,
		"ifSourceGenerationMatch":        false,
		"ifSourceGenerationNotMatch":     false,
		"ifSourceMetagenerationMatch":    false,
		"ifSourceMetagenerationNotMatch": false,
		"maxBytesRewrittenPerCall":       false,
		"projection":                     false,
		"rewriteToken":                   false,
		"sourceGeneration":               false,
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

	request := &api_client.Object{}
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
		"sourceBucket",
		"sourceObject",
		"destinationBucket",
		"destinationObject",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_sourceBucket, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_sourceObject, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}
	param_destinationBucket, err := commands_util.ConvertValue_string(paramValues[2])
	if err != nil {
		return err
	}
	param_destinationObject, err := commands_util.ConvertValue_string(paramValues[3])
	if err != nil {
		return err
	}

	call := service.Rewrite(param_sourceBucket, param_sourceObject, param_destinationBucket, param_destinationObject,
		request,
	)

	// Set query parameters.
	if value, ok := flagValues["destinationPredefinedAcl"]; ok {
		query_destinationPredefinedAcl, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.DestinationPredefinedAcl(query_destinationPredefinedAcl)
	}
	if value, ok := flagValues["ifGenerationMatch"]; ok {
		query_ifGenerationMatch, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.IfGenerationMatch(query_ifGenerationMatch)
	}
	if value, ok := flagValues["ifGenerationNotMatch"]; ok {
		query_ifGenerationNotMatch, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.IfGenerationNotMatch(query_ifGenerationNotMatch)
	}
	if value, ok := flagValues["ifMetagenerationMatch"]; ok {
		query_ifMetagenerationMatch, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.IfMetagenerationMatch(query_ifMetagenerationMatch)
	}
	if value, ok := flagValues["ifMetagenerationNotMatch"]; ok {
		query_ifMetagenerationNotMatch, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.IfMetagenerationNotMatch(query_ifMetagenerationNotMatch)
	}
	if value, ok := flagValues["ifSourceGenerationMatch"]; ok {
		query_ifSourceGenerationMatch, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.IfSourceGenerationMatch(query_ifSourceGenerationMatch)
	}
	if value, ok := flagValues["ifSourceGenerationNotMatch"]; ok {
		query_ifSourceGenerationNotMatch, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.IfSourceGenerationNotMatch(query_ifSourceGenerationNotMatch)
	}
	if value, ok := flagValues["ifSourceMetagenerationMatch"]; ok {
		query_ifSourceMetagenerationMatch, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.IfSourceMetagenerationMatch(query_ifSourceMetagenerationMatch)
	}
	if value, ok := flagValues["ifSourceMetagenerationNotMatch"]; ok {
		query_ifSourceMetagenerationNotMatch, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.IfSourceMetagenerationNotMatch(query_ifSourceMetagenerationNotMatch)
	}
	if value, ok := flagValues["maxBytesRewrittenPerCall"]; ok {
		query_maxBytesRewrittenPerCall, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.MaxBytesRewrittenPerCall(query_maxBytesRewrittenPerCall)
	}
	if value, ok := flagValues["projection"]; ok {
		query_projection, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.Projection(query_projection)
	}
	if value, ok := flagValues["rewriteToken"]; ok {
		query_rewriteToken, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.RewriteToken(query_rewriteToken)
	}
	if value, ok := flagValues["sourceGeneration"]; ok {
		query_sourceGeneration, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.SourceGeneration(query_sourceGeneration)
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

func Storage_v1_ObjectsSetIamPolicy(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("bucket"))
		pathParams = append(pathParams, commands_util.AngrySnakes("object"))

		if len(pathParams) != 0 {
			if strings.Contains("b/{bucket}/o/{object}/iam", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		usageBits += " [--generation=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		commands_util.PrintRequestExample(&api_client.Policy{})

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewObjectsService(api_service)

	queryParamNames := map[string]bool{
		"generation": false,
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

	request := &api_client.Policy{}
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
		"bucket",
		"object",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_bucket, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_object, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}

	call := service.SetIamPolicy(param_bucket, param_object,
		request,
	)

	// Set query parameters.
	if value, ok := flagValues["generation"]; ok {
		query_generation, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.Generation(query_generation)
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

func Storage_v1_ObjectsTestIamPermissions(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("bucket"))
		pathParams = append(pathParams, commands_util.AngrySnakes("object"))
		pathParams = append(pathParams, commands_util.AngrySnakes("permissions"))

		if len(pathParams) != 0 {
			if strings.Contains("b/{bucket}/o/{object}/iam/testPermissions", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [--generation=VALUE]"

		usageBits += " --permissions=VALUE"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewObjectsService(api_service)

	queryParamNames := map[string]bool{
		"generation":  false,
		"permissions": true,
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
		"bucket",
		"object",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_bucket, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_object, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}
	param_permissions, err := commands_util.ConvertValue_strings(flagValues["permissions"])
	if err != nil {
		return err
	}

	call := service.TestIamPermissions(param_bucket, param_object, param_permissions)

	// Set query parameters.
	if value, ok := flagValues["generation"]; ok {
		query_generation, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.Generation(query_generation)
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

func Storage_v1_ObjectsUpdate(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("bucket"))
		pathParams = append(pathParams, commands_util.AngrySnakes("object"))

		if len(pathParams) != 0 {
			if strings.Contains("b/{bucket}/o/{object}", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		usageBits += " [--generation=VALUE]"

		usageBits += " [--ifGenerationMatch=VALUE]"

		usageBits += " [--ifGenerationNotMatch=VALUE]"

		usageBits += " [--ifMetagenerationMatch=VALUE]"

		usageBits += " [--ifMetagenerationNotMatch=VALUE]"

		usageBits += " [--predefinedAcl=VALUE]"

		usageBits += " [--projection=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		commands_util.PrintRequestExample(&api_client.Object{})

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewObjectsService(api_service)

	queryParamNames := map[string]bool{
		"generation":               false,
		"ifGenerationMatch":        false,
		"ifGenerationNotMatch":     false,
		"ifMetagenerationMatch":    false,
		"ifMetagenerationNotMatch": false,
		"predefinedAcl":            false,
		"projection":               false,
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

	request := &api_client.Object{}
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
		"bucket",
		"object",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_bucket, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_object, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}

	call := service.Update(param_bucket, param_object,
		request,
	)

	// Set query parameters.
	if value, ok := flagValues["generation"]; ok {
		query_generation, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.Generation(query_generation)
	}
	if value, ok := flagValues["ifGenerationMatch"]; ok {
		query_ifGenerationMatch, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.IfGenerationMatch(query_ifGenerationMatch)
	}
	if value, ok := flagValues["ifGenerationNotMatch"]; ok {
		query_ifGenerationNotMatch, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.IfGenerationNotMatch(query_ifGenerationNotMatch)
	}
	if value, ok := flagValues["ifMetagenerationMatch"]; ok {
		query_ifMetagenerationMatch, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.IfMetagenerationMatch(query_ifMetagenerationMatch)
	}
	if value, ok := flagValues["ifMetagenerationNotMatch"]; ok {
		query_ifMetagenerationNotMatch, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.IfMetagenerationNotMatch(query_ifMetagenerationNotMatch)
	}
	if value, ok := flagValues["predefinedAcl"]; ok {
		query_predefinedAcl, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.PredefinedAcl(query_predefinedAcl)
	}
	if value, ok := flagValues["projection"]; ok {
		query_projection, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.Projection(query_projection)
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

func Storage_v1_ObjectsWatchAll(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("bucket"))

		if len(pathParams) != 0 {
			if strings.Contains("b/{bucket}/o/watch", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		usageBits += " [--delimiter=VALUE]"

		usageBits += " [--maxResults=VALUE]"

		usageBits += " [--pageToken=VALUE]"

		usageBits += " [--prefix=VALUE]"

		usageBits += " [--projection=VALUE]"

		usageBits += " [--versions=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		commands_util.PrintRequestExample(&api_client.Channel{})

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewObjectsService(api_service)

	queryParamNames := map[string]bool{
		"delimiter":  false,
		"maxResults": false,
		"pageToken":  false,
		"prefix":     false,
		"projection": false,
		"versions":   false,
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

	request := &api_client.Channel{}
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
		"bucket",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_bucket, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}

	call := service.WatchAll(param_bucket,
		request,
	)

	// Set query parameters.
	if value, ok := flagValues["delimiter"]; ok {
		query_delimiter, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.Delimiter(query_delimiter)
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
	if value, ok := flagValues["prefix"]; ok {
		query_prefix, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.Prefix(query_prefix)
	}
	if value, ok := flagValues["projection"]; ok {
		query_projection, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.Projection(query_projection)
	}
	if value, ok := flagValues["versions"]; ok {
		query_versions, err := commands_util.ConvertValue_bool(value)
		if err != nil {
			return err
		}
		call.Versions(query_versions)
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

func Storage_v1_ProjectsServiceAccountGet(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectId"))

		if len(pathParams) != 0 {
			if strings.Contains("projects/{projectId}/serviceAccount", "+") {
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
	service := api_client.NewProjectsServiceAccountService(api_service)

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

	call := service.Get(param_projectId)

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
