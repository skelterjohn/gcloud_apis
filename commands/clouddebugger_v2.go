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

	api_client "github.com/skelterjohn/gcloud_apis/clients/clouddebugger/v2"
	"github.com/skelterjohn/gcloud_apis/commands_util"
)

// some generated code may not use all imported libs
var _ = fmt.Println
var _ = io.Copy
var _ = os.Stdin
var _ = strings.Split

func Clouddebugger_v2_ControllerDebuggeesBreakpointsList(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("debuggeeId"))

		if len(pathParams) != 0 {
			if strings.Contains("v2/controller/debuggees/{debuggeeId}/breakpoints", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [--successOnTimeout=VALUE]"

		usageBits += " [--waitToken=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewControllerDebuggeesBreakpointsService(api_service)

	queryParamNames := map[string]bool{
		"successOnTimeout": false,
		"waitToken":        false,
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
		"debuggeeId",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_debuggeeId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}

	call := service.List(param_debuggeeId)

	// Set query parameters.
	if value, ok := flagValues["successOnTimeout"]; ok {
		query_successOnTimeout, err := commands_util.ConvertValue_bool(value)
		if err != nil {
			return err
		}
		call.SuccessOnTimeout(query_successOnTimeout)
	}
	if value, ok := flagValues["waitToken"]; ok {
		query_waitToken, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.WaitToken(query_waitToken)
	}

	var response *api_client.ListActiveBreakpointsResponse
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

func Clouddebugger_v2_ControllerDebuggeesBreakpointsUpdate(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("debuggeeId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("id"))

		if len(pathParams) != 0 {
			if strings.Contains("v2/controller/debuggees/{debuggeeId}/breakpoints/{id}", "+") {
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
	service := api_client.NewControllerDebuggeesBreakpointsService(api_service)

	args, flagValues, err := commands_util.ExtractFlagValues(args)
	if err != nil {
		return err
	}

	// Only positional arguments should remain in args.
	if len(args) == 0 || len(args) > 2 {
		usageFunc()
	}

	request := &api_client.UpdateActiveBreakpointRequest{}
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
		"debuggeeId",
		"id",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_debuggeeId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_id, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}

	call := service.Update(param_debuggeeId, param_id,
		request,
	)

	var response *api_client.UpdateActiveBreakpointResponse
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

func Clouddebugger_v2_ControllerDebuggeesRegister(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string

		if len(pathParams) != 0 {
			if strings.Contains("v2/controller/debuggees/register", "+") {
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
	service := api_client.NewControllerDebuggeesService(api_service)

	args, flagValues, err := commands_util.ExtractFlagValues(args)
	if err != nil {
		return err
	}

	// Only positional arguments should remain in args.
	if len(args) == 0 || len(args) > 2 {
		usageFunc()
	}

	request := &api_client.RegisterDebuggeeRequest{}
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

	call := service.Register(
		request,
	)

	var response *api_client.RegisterDebuggeeResponse
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

func Clouddebugger_v2_DebuggerDebuggeesBreakpointsDelete(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("debuggeeId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("breakpointId"))

		if len(pathParams) != 0 {
			if strings.Contains("v2/debugger/debuggees/{debuggeeId}/breakpoints/{breakpointId}", "+") {
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
	service := api_client.NewDebuggerDebuggeesBreakpointsService(api_service)

	// Only positional arguments should remain in args.
	if len(args) != 1 {
		usageFunc()
	}

	expectedParams := []string{
		"debuggeeId",
		"breakpointId",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_debuggeeId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_breakpointId, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}

	call := service.Delete(param_debuggeeId, param_breakpointId)

	var response *api_client.Empty
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

func Clouddebugger_v2_DebuggerDebuggeesBreakpointsGet(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("debuggeeId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("breakpointId"))

		if len(pathParams) != 0 {
			if strings.Contains("v2/debugger/debuggees/{debuggeeId}/breakpoints/{breakpointId}", "+") {
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
	service := api_client.NewDebuggerDebuggeesBreakpointsService(api_service)

	// Only positional arguments should remain in args.
	if len(args) != 1 {
		usageFunc()
	}

	expectedParams := []string{
		"debuggeeId",
		"breakpointId",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_debuggeeId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_breakpointId, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}

	call := service.Get(param_debuggeeId, param_breakpointId)

	var response *api_client.GetBreakpointResponse
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

func Clouddebugger_v2_DebuggerDebuggeesBreakpointsList(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("debuggeeId"))

		if len(pathParams) != 0 {
			if strings.Contains("v2/debugger/debuggees/{debuggeeId}/breakpoints", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [--action.value=VALUE]"

		usageBits += " [--includeAllUsers=VALUE]"

		usageBits += " [--includeInactive=VALUE]"

		usageBits += " [--stripResults=VALUE]"

		usageBits += " [--waitToken=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewDebuggerDebuggeesBreakpointsService(api_service)

	queryParamNames := map[string]bool{
		"action.value":    false,
		"includeAllUsers": false,
		"includeInactive": false,
		"stripResults":    false,
		"waitToken":       false,
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
		"debuggeeId",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_debuggeeId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}

	call := service.List(param_debuggeeId)

	// Set query parameters.
	if value, ok := flagValues["action_value"]; ok {
		query_action_value, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.ActionValue(query_action_value)
	}
	if value, ok := flagValues["includeAllUsers"]; ok {
		query_includeAllUsers, err := commands_util.ConvertValue_bool(value)
		if err != nil {
			return err
		}
		call.IncludeAllUsers(query_includeAllUsers)
	}
	if value, ok := flagValues["includeInactive"]; ok {
		query_includeInactive, err := commands_util.ConvertValue_bool(value)
		if err != nil {
			return err
		}
		call.IncludeInactive(query_includeInactive)
	}
	if value, ok := flagValues["stripResults"]; ok {
		query_stripResults, err := commands_util.ConvertValue_bool(value)
		if err != nil {
			return err
		}
		call.StripResults(query_stripResults)
	}
	if value, ok := flagValues["waitToken"]; ok {
		query_waitToken, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.WaitToken(query_waitToken)
	}

	var response *api_client.ListBreakpointsResponse
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

func Clouddebugger_v2_DebuggerDebuggeesBreakpointsSet(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("debuggeeId"))

		if len(pathParams) != 0 {
			if strings.Contains("v2/debugger/debuggees/{debuggeeId}/breakpoints/set", "+") {
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
	service := api_client.NewDebuggerDebuggeesBreakpointsService(api_service)

	args, flagValues, err := commands_util.ExtractFlagValues(args)
	if err != nil {
		return err
	}

	// Only positional arguments should remain in args.
	if len(args) == 0 || len(args) > 2 {
		usageFunc()
	}

	request := &api_client.Breakpoint{}
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
		"debuggeeId",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_debuggeeId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}

	call := service.Set(param_debuggeeId,
		request,
	)

	var response *api_client.SetBreakpointResponse
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

func Clouddebugger_v2_DebuggerDebuggeesList(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string

		if len(pathParams) != 0 {
			if strings.Contains("v2/debugger/debuggees", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [--includeInactive=VALUE]"

		usageBits += " [--project=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewDebuggerDebuggeesService(api_service)

	queryParamNames := map[string]bool{
		"includeInactive": false,
		"project":         false,
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
	if value, ok := flagValues["includeInactive"]; ok {
		query_includeInactive, err := commands_util.ConvertValue_bool(value)
		if err != nil {
			return err
		}
		call.IncludeInactive(query_includeInactive)
	}
	if value, ok := flagValues["project"]; ok {
		query_project, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.Project(query_project)
	}

	var response *api_client.ListDebuggeesResponse
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
