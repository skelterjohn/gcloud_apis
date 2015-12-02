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

	api_client "github.com/skelterjohn/gcloud_apis/clients/source/v1"
	"github.com/skelterjohn/gcloud_apis/commands_util"
)

// some generated code may not use all imported libs
var _ = fmt.Println
var _ = io.Copy
var _ = os.Stdin
var _ = strings.Split

func Source_v1_ListChangedFiles(context Context, args ...string) error {

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
	service := api_client.NewV1Service(api_service)

	args, flagValues, err := commands_util.ExtractFlagValues(args)
	if err != nil {
		return err
	}

	// Only positional arguments should remain in args.
	if len(args) == 0 || len(args) > 2 {
		usageFunc()
	}

	request := &api_client.ListChangedFilesRequest{}
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

	call := service.ListChangedFiles(
		request,
	)

	var response *api_client.ListChangedFilesResponse
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

func Source_v1_ProjectsReposAliasesCreate(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		usageBits += " [--uid=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsReposAliasesService(api_service)

	queryParamNames := map[string]bool{
		"uid": false,
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

	request := &api_client.Alias{}
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
		"repoName",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_repoName, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}

	call := service.Create(param_projectId, param_repoName,
		request,
	)

	// Set query parameters.
	if value, ok := flagValues["uid"]; ok {
		query_uid, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.Uid(query_uid)
	}

	var response *api_client.Alias
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

func Source_v1_ProjectsReposAliasesDelete(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)

		usageBits += " [--revisionId=VALUE]"

		usageBits += " [--uid=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsReposAliasesService(api_service)

	queryParamNames := map[string]bool{
		"revisionId": false,
		"uid":        false,
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
		"repoName",
		"kind",
		"name",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_repoName, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}
	param_kind, err := commands_util.ConvertValue_string(paramValues[2])
	if err != nil {
		return err
	}
	param_name, err := commands_util.ConvertValue_string(paramValues[3])
	if err != nil {
		return err
	}

	call := service.Delete(param_projectId, param_repoName, param_kind, param_name)

	// Set query parameters.
	if value, ok := flagValues["revisionId"]; ok {
		query_revisionId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.RevisionId(query_revisionId)
	}
	if value, ok := flagValues["uid"]; ok {
		query_uid, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.Uid(query_uid)
	}

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

func Source_v1_ProjectsReposAliasesFilesGet(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)

		usageBits += " [--aliasName=VALUE]"

		usageBits += " [--cloudWorkspace.snapshotId=VALUE]"

		usageBits += " [--cloudWorkspace.workspaceId.name=VALUE]"

		usageBits += " [--cloudWorkspace.workspaceId.repoId.projectRepoId.projectId=VALUE]"

		usageBits += " [--cloudWorkspace.workspaceId.repoId.projectRepoId.repoName=VALUE]"

		usageBits += " [--cloudWorkspace.workspaceId.repoId.uid=VALUE]"

		usageBits += " [--gerrit.aliasContext.kind=VALUE]"

		usageBits += " [--gerrit.aliasContext.name=VALUE]"

		usageBits += " [--gerrit.aliasName=VALUE]"

		usageBits += " [--gerrit.gerritProject=VALUE]"

		usageBits += " [--gerrit.hostUri=VALUE]"

		usageBits += " [--gerrit.revisionId=VALUE]"

		usageBits += " [--git.revisionId=VALUE]"

		usageBits += " [--git.url=VALUE]"

		usageBits += " [--pageSize=VALUE]"

		usageBits += " [--pageToken=VALUE]"

		usageBits += " [--revisionId=VALUE]"

		usageBits += " [--startPosition=VALUE]"

		usageBits += " [--uid=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsReposAliasesFilesService(api_service)

	queryParamNames := map[string]bool{
		"aliasName":                                                 false,
		"cloudWorkspace.snapshotId":                                 false,
		"cloudWorkspace.workspaceId.name":                           false,
		"cloudWorkspace.workspaceId.repoId.projectRepoId.projectId": false,
		"cloudWorkspace.workspaceId.repoId.projectRepoId.repoName":  false,
		"cloudWorkspace.workspaceId.repoId.uid":                     false,
		"gerrit.aliasContext.kind":                                  false,
		"gerrit.aliasContext.name":                                  false,
		"gerrit.aliasName":                                          false,
		"gerrit.gerritProject":                                      false,
		"gerrit.hostUri":                                            false,
		"gerrit.revisionId":                                         false,
		"git.revisionId":                                            false,
		"git.url":                                                   false,
		"pageSize":                                                  false,
		"pageToken":                                                 false,
		"revisionId":                                                false,
		"startPosition":                                             false,
		"uid":                                                       false,
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
		"repoName",
		"kind",
		"name",
		"path",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_repoName, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}
	param_kind, err := commands_util.ConvertValue_string(paramValues[2])
	if err != nil {
		return err
	}
	param_name, err := commands_util.ConvertValue_string(paramValues[3])
	if err != nil {
		return err
	}
	param_path, err := commands_util.ConvertValue_string(paramValues[4])
	if err != nil {
		return err
	}

	call := service.Get(param_projectId, param_repoName, param_kind, param_name, param_path)

	// Set query parameters.
	if value, ok := flagValues["aliasName"]; ok {
		query_aliasName, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.AliasName(query_aliasName)
	}
	if value, ok := flagValues["cloudWorkspace_snapshotId"]; ok {
		query_cloudWorkspace_snapshotId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.CloudWorkspaceSnapshotId(query_cloudWorkspace_snapshotId)
	}
	if value, ok := flagValues["cloudWorkspace_workspaceId_name"]; ok {
		query_cloudWorkspace_workspaceId_name, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.CloudWorkspaceWorkspaceIdName(query_cloudWorkspace_workspaceId_name)
	}
	if value, ok := flagValues["cloudWorkspace_workspaceId_repoId_projectRepoId_projectId"]; ok {
		query_cloudWorkspace_workspaceId_repoId_projectRepoId_projectId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.CloudWorkspaceWorkspaceIdRepoIdProjectRepoIdProjectId(query_cloudWorkspace_workspaceId_repoId_projectRepoId_projectId)
	}
	if value, ok := flagValues["cloudWorkspace_workspaceId_repoId_projectRepoId_repoName"]; ok {
		query_cloudWorkspace_workspaceId_repoId_projectRepoId_repoName, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.CloudWorkspaceWorkspaceIdRepoIdProjectRepoIdRepoName(query_cloudWorkspace_workspaceId_repoId_projectRepoId_repoName)
	}
	if value, ok := flagValues["cloudWorkspace_workspaceId_repoId_uid"]; ok {
		query_cloudWorkspace_workspaceId_repoId_uid, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.CloudWorkspaceWorkspaceIdRepoIdUid(query_cloudWorkspace_workspaceId_repoId_uid)
	}
	if value, ok := flagValues["gerrit_aliasContext_kind"]; ok {
		query_gerrit_aliasContext_kind, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.GerritAliasContextKind(query_gerrit_aliasContext_kind)
	}
	if value, ok := flagValues["gerrit_aliasContext_name"]; ok {
		query_gerrit_aliasContext_name, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.GerritAliasContextName(query_gerrit_aliasContext_name)
	}
	if value, ok := flagValues["gerrit_aliasName"]; ok {
		query_gerrit_aliasName, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.GerritAliasName(query_gerrit_aliasName)
	}
	if value, ok := flagValues["gerrit_gerritProject"]; ok {
		query_gerrit_gerritProject, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.GerritGerritProject(query_gerrit_gerritProject)
	}
	if value, ok := flagValues["gerrit_hostUri"]; ok {
		query_gerrit_hostUri, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.GerritHostUri(query_gerrit_hostUri)
	}
	if value, ok := flagValues["gerrit_revisionId"]; ok {
		query_gerrit_revisionId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.GerritRevisionId(query_gerrit_revisionId)
	}
	if value, ok := flagValues["git_revisionId"]; ok {
		query_git_revisionId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.GitRevisionId(query_git_revisionId)
	}
	if value, ok := flagValues["git_url"]; ok {
		query_git_url, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.GitUrl(query_git_url)
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
	if value, ok := flagValues["revisionId"]; ok {
		query_revisionId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.RevisionId(query_revisionId)
	}
	if value, ok := flagValues["startPosition"]; ok {
		query_startPosition, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.StartPosition(query_startPosition)
	}
	if value, ok := flagValues["uid"]; ok {
		query_uid, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.Uid(query_uid)
	}

	var response *api_client.ReadResponse
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

func Source_v1_ProjectsReposAliasesGet(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)

		usageBits += " [--uid=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsReposAliasesService(api_service)

	queryParamNames := map[string]bool{
		"uid": false,
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
		"repoName",
		"kind",
		"name",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_repoName, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}
	param_kind, err := commands_util.ConvertValue_string(paramValues[2])
	if err != nil {
		return err
	}
	param_name, err := commands_util.ConvertValue_string(paramValues[3])
	if err != nil {
		return err
	}

	call := service.Get(param_projectId, param_repoName, param_kind, param_name)

	// Set query parameters.
	if value, ok := flagValues["uid"]; ok {
		query_uid, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.Uid(query_uid)
	}

	var response *api_client.Alias
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

func Source_v1_ProjectsReposAliasesList(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)

		usageBits += " [--kind=VALUE]"

		usageBits += " [--pageSize=VALUE]"

		usageBits += " [--pageToken=VALUE]"

		usageBits += " [--uid=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsReposAliasesService(api_service)

	queryParamNames := map[string]bool{
		"kind":      false,
		"pageSize":  false,
		"pageToken": false,
		"uid":       false,
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
		"repoName",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_repoName, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}

	call := service.List(param_projectId, param_repoName)

	// Set query parameters.
	if value, ok := flagValues["kind"]; ok {
		query_kind, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.Kind(query_kind)
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
	if value, ok := flagValues["uid"]; ok {
		query_uid, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.Uid(query_uid)
	}

	var response *api_client.ListAliasesResponse
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

func Source_v1_ProjectsReposAliasesListFiles(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)

		usageBits += " [--aliasName=VALUE]"

		usageBits += " [--cloudWorkspace.snapshotId=VALUE]"

		usageBits += " [--cloudWorkspace.workspaceId.name=VALUE]"

		usageBits += " [--cloudWorkspace.workspaceId.repoId.projectRepoId.projectId=VALUE]"

		usageBits += " [--cloudWorkspace.workspaceId.repoId.projectRepoId.repoName=VALUE]"

		usageBits += " [--cloudWorkspace.workspaceId.repoId.uid=VALUE]"

		usageBits += " [--gerrit.aliasContext.kind=VALUE]"

		usageBits += " [--gerrit.aliasContext.name=VALUE]"

		usageBits += " [--gerrit.aliasName=VALUE]"

		usageBits += " [--gerrit.gerritProject=VALUE]"

		usageBits += " [--gerrit.hostUri=VALUE]"

		usageBits += " [--gerrit.revisionId=VALUE]"

		usageBits += " [--git.revisionId=VALUE]"

		usageBits += " [--git.url=VALUE]"

		usageBits += " [--pageSize=VALUE]"

		usageBits += " [--pageToken=VALUE]"

		usageBits += " [--revisionId=VALUE]"

		usageBits += " [--uid=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsReposAliasesService(api_service)

	queryParamNames := map[string]bool{
		"aliasName":                                                 false,
		"cloudWorkspace.snapshotId":                                 false,
		"cloudWorkspace.workspaceId.name":                           false,
		"cloudWorkspace.workspaceId.repoId.projectRepoId.projectId": false,
		"cloudWorkspace.workspaceId.repoId.projectRepoId.repoName":  false,
		"cloudWorkspace.workspaceId.repoId.uid":                     false,
		"gerrit.aliasContext.kind":                                  false,
		"gerrit.aliasContext.name":                                  false,
		"gerrit.aliasName":                                          false,
		"gerrit.gerritProject":                                      false,
		"gerrit.hostUri":                                            false,
		"gerrit.revisionId":                                         false,
		"git.revisionId":                                            false,
		"git.url":                                                   false,
		"pageSize":                                                  false,
		"pageToken":                                                 false,
		"revisionId":                                                false,
		"uid":                                                       false,
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
		"repoName",
		"kind",
		"name",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_repoName, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}
	param_kind, err := commands_util.ConvertValue_string(paramValues[2])
	if err != nil {
		return err
	}
	param_name, err := commands_util.ConvertValue_string(paramValues[3])
	if err != nil {
		return err
	}

	call := service.ListFiles(param_projectId, param_repoName, param_kind, param_name)

	// Set query parameters.
	if value, ok := flagValues["aliasName"]; ok {
		query_aliasName, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.AliasName(query_aliasName)
	}
	if value, ok := flagValues["cloudWorkspace_snapshotId"]; ok {
		query_cloudWorkspace_snapshotId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.CloudWorkspaceSnapshotId(query_cloudWorkspace_snapshotId)
	}
	if value, ok := flagValues["cloudWorkspace_workspaceId_name"]; ok {
		query_cloudWorkspace_workspaceId_name, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.CloudWorkspaceWorkspaceIdName(query_cloudWorkspace_workspaceId_name)
	}
	if value, ok := flagValues["cloudWorkspace_workspaceId_repoId_projectRepoId_projectId"]; ok {
		query_cloudWorkspace_workspaceId_repoId_projectRepoId_projectId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.CloudWorkspaceWorkspaceIdRepoIdProjectRepoIdProjectId(query_cloudWorkspace_workspaceId_repoId_projectRepoId_projectId)
	}
	if value, ok := flagValues["cloudWorkspace_workspaceId_repoId_projectRepoId_repoName"]; ok {
		query_cloudWorkspace_workspaceId_repoId_projectRepoId_repoName, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.CloudWorkspaceWorkspaceIdRepoIdProjectRepoIdRepoName(query_cloudWorkspace_workspaceId_repoId_projectRepoId_repoName)
	}
	if value, ok := flagValues["cloudWorkspace_workspaceId_repoId_uid"]; ok {
		query_cloudWorkspace_workspaceId_repoId_uid, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.CloudWorkspaceWorkspaceIdRepoIdUid(query_cloudWorkspace_workspaceId_repoId_uid)
	}
	if value, ok := flagValues["gerrit_aliasContext_kind"]; ok {
		query_gerrit_aliasContext_kind, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.GerritAliasContextKind(query_gerrit_aliasContext_kind)
	}
	if value, ok := flagValues["gerrit_aliasContext_name"]; ok {
		query_gerrit_aliasContext_name, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.GerritAliasContextName(query_gerrit_aliasContext_name)
	}
	if value, ok := flagValues["gerrit_aliasName"]; ok {
		query_gerrit_aliasName, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.GerritAliasName(query_gerrit_aliasName)
	}
	if value, ok := flagValues["gerrit_gerritProject"]; ok {
		query_gerrit_gerritProject, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.GerritGerritProject(query_gerrit_gerritProject)
	}
	if value, ok := flagValues["gerrit_hostUri"]; ok {
		query_gerrit_hostUri, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.GerritHostUri(query_gerrit_hostUri)
	}
	if value, ok := flagValues["gerrit_revisionId"]; ok {
		query_gerrit_revisionId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.GerritRevisionId(query_gerrit_revisionId)
	}
	if value, ok := flagValues["git_revisionId"]; ok {
		query_git_revisionId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.GitRevisionId(query_git_revisionId)
	}
	if value, ok := flagValues["git_url"]; ok {
		query_git_url, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.GitUrl(query_git_url)
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
	if value, ok := flagValues["revisionId"]; ok {
		query_revisionId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.RevisionId(query_revisionId)
	}
	if value, ok := flagValues["uid"]; ok {
		query_uid, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.Uid(query_uid)
	}

	var response *api_client.ListFilesResponse
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

func Source_v1_ProjectsReposAliasesUpdate(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		usageBits += " [--oldRevisionId=VALUE]"

		usageBits += " [--uid=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsReposAliasesService(api_service)

	queryParamNames := map[string]bool{
		"oldRevisionId": false,
		"uid":           false,
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

	request := &api_client.Alias{}
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
		"repoName",
		"aliasesId",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_repoName, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}
	param_aliasesId, err := commands_util.ConvertValue_string(paramValues[2])
	if err != nil {
		return err
	}

	call := service.Update(param_projectId, param_repoName, param_aliasesId,
		request,
	)

	// Set query parameters.
	if value, ok := flagValues["oldRevisionId"]; ok {
		query_oldRevisionId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.OldRevisionId(query_oldRevisionId)
	}
	if value, ok := flagValues["uid"]; ok {
		query_uid, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.Uid(query_uid)
	}

	var response *api_client.Alias
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

func Source_v1_ProjectsReposCreate(context Context, args ...string) error {

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
	service := api_client.NewProjectsReposService(api_service)

	args, flagValues, err := commands_util.ExtractFlagValues(args)
	if err != nil {
		return err
	}

	// Only positional arguments should remain in args.
	if len(args) == 0 || len(args) > 2 {
		usageFunc()
	}

	request := &api_client.Repo{}
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

	var response *api_client.Repo
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

func Source_v1_ProjectsReposDelete(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)

		usageBits += " [--uid=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsReposService(api_service)

	queryParamNames := map[string]bool{
		"uid": false,
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
		"repoName",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_repoName, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}

	call := service.Delete(param_projectId, param_repoName)

	// Set query parameters.
	if value, ok := flagValues["uid"]; ok {
		query_uid, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.Uid(query_uid)
	}

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

func Source_v1_ProjectsReposFilesReadFromWorkspaceOrAlias(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)

		usageBits += " [--alias=VALUE]"

		usageBits += " [--pageSize=VALUE]"

		usageBits += " [--pageToken=VALUE]"

		usageBits += " [--startPosition=VALUE]"

		usageBits += " [--uid=VALUE]"

		usageBits += " [--workspaceName=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsReposFilesService(api_service)

	queryParamNames := map[string]bool{
		"alias":         false,
		"pageSize":      false,
		"pageToken":     false,
		"startPosition": false,
		"uid":           false,
		"workspaceName": false,
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
		"repoName",
		"path",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_repoName, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}
	param_path, err := commands_util.ConvertValue_string(paramValues[2])
	if err != nil {
		return err
	}

	call := service.ReadFromWorkspaceOrAlias(param_projectId, param_repoName, param_path)

	// Set query parameters.
	if value, ok := flagValues["alias"]; ok {
		query_alias, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.Alias(query_alias)
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
	if value, ok := flagValues["startPosition"]; ok {
		query_startPosition, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.StartPosition(query_startPosition)
	}
	if value, ok := flagValues["uid"]; ok {
		query_uid, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.Uid(query_uid)
	}
	if value, ok := flagValues["workspaceName"]; ok {
		query_workspaceName, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.WorkspaceName(query_workspaceName)
	}

	var response *api_client.ReadResponse
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

func Source_v1_ProjectsReposGet(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)

		usageBits += " [--uid=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsReposService(api_service)

	queryParamNames := map[string]bool{
		"uid": false,
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
		"repoName",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_repoName, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}

	call := service.Get(param_projectId, param_repoName)

	// Set query parameters.
	if value, ok := flagValues["uid"]; ok {
		query_uid, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.Uid(query_uid)
	}

	var response *api_client.Repo
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

func Source_v1_ProjectsReposList(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsReposService(api_service)

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

	var response *api_client.ListReposResponse
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

func Source_v1_ProjectsReposMerge(context Context, args ...string) error {

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
	service := api_client.NewProjectsReposService(api_service)

	args, flagValues, err := commands_util.ExtractFlagValues(args)
	if err != nil {
		return err
	}

	// Only positional arguments should remain in args.
	if len(args) == 0 || len(args) > 2 {
		usageFunc()
	}

	request := &api_client.MergeRequest{}
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
		"repoName",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_repoName, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}

	call := service.Merge(param_projectId, param_repoName,
		request,
	)

	var response *api_client.Workspace
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

func Source_v1_ProjectsReposRevisionsFilesGet(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)

		usageBits += " [--aliasContext.kind=VALUE]"

		usageBits += " [--aliasContext.name=VALUE]"

		usageBits += " [--aliasName=VALUE]"

		usageBits += " [--cloudWorkspace.snapshotId=VALUE]"

		usageBits += " [--cloudWorkspace.workspaceId.name=VALUE]"

		usageBits += " [--cloudWorkspace.workspaceId.repoId.projectRepoId.projectId=VALUE]"

		usageBits += " [--cloudWorkspace.workspaceId.repoId.projectRepoId.repoName=VALUE]"

		usageBits += " [--cloudWorkspace.workspaceId.repoId.uid=VALUE]"

		usageBits += " [--gerrit.aliasContext.kind=VALUE]"

		usageBits += " [--gerrit.aliasContext.name=VALUE]"

		usageBits += " [--gerrit.aliasName=VALUE]"

		usageBits += " [--gerrit.gerritProject=VALUE]"

		usageBits += " [--gerrit.hostUri=VALUE]"

		usageBits += " [--gerrit.revisionId=VALUE]"

		usageBits += " [--git.revisionId=VALUE]"

		usageBits += " [--git.url=VALUE]"

		usageBits += " [--pageSize=VALUE]"

		usageBits += " [--pageToken=VALUE]"

		usageBits += " [--startPosition=VALUE]"

		usageBits += " [--uid=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsReposRevisionsFilesService(api_service)

	queryParamNames := map[string]bool{
		"aliasContext.kind":                                         false,
		"aliasContext.name":                                         false,
		"aliasName":                                                 false,
		"cloudWorkspace.snapshotId":                                 false,
		"cloudWorkspace.workspaceId.name":                           false,
		"cloudWorkspace.workspaceId.repoId.projectRepoId.projectId": false,
		"cloudWorkspace.workspaceId.repoId.projectRepoId.repoName":  false,
		"cloudWorkspace.workspaceId.repoId.uid":                     false,
		"gerrit.aliasContext.kind":                                  false,
		"gerrit.aliasContext.name":                                  false,
		"gerrit.aliasName":                                          false,
		"gerrit.gerritProject":                                      false,
		"gerrit.hostUri":                                            false,
		"gerrit.revisionId":                                         false,
		"git.revisionId":                                            false,
		"git.url":                                                   false,
		"pageSize":                                                  false,
		"pageToken":                                                 false,
		"startPosition":                                             false,
		"uid":                                                       false,
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
		"repoName",
		"revisionId",
		"path",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_repoName, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}
	param_revisionId, err := commands_util.ConvertValue_string(paramValues[2])
	if err != nil {
		return err
	}
	param_path, err := commands_util.ConvertValue_string(paramValues[3])
	if err != nil {
		return err
	}

	call := service.Get(param_projectId, param_repoName, param_revisionId, param_path)

	// Set query parameters.
	if value, ok := flagValues["aliasContext_kind"]; ok {
		query_aliasContext_kind, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.AliasContextKind(query_aliasContext_kind)
	}
	if value, ok := flagValues["aliasContext_name"]; ok {
		query_aliasContext_name, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.AliasContextName(query_aliasContext_name)
	}
	if value, ok := flagValues["aliasName"]; ok {
		query_aliasName, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.AliasName(query_aliasName)
	}
	if value, ok := flagValues["cloudWorkspace_snapshotId"]; ok {
		query_cloudWorkspace_snapshotId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.CloudWorkspaceSnapshotId(query_cloudWorkspace_snapshotId)
	}
	if value, ok := flagValues["cloudWorkspace_workspaceId_name"]; ok {
		query_cloudWorkspace_workspaceId_name, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.CloudWorkspaceWorkspaceIdName(query_cloudWorkspace_workspaceId_name)
	}
	if value, ok := flagValues["cloudWorkspace_workspaceId_repoId_projectRepoId_projectId"]; ok {
		query_cloudWorkspace_workspaceId_repoId_projectRepoId_projectId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.CloudWorkspaceWorkspaceIdRepoIdProjectRepoIdProjectId(query_cloudWorkspace_workspaceId_repoId_projectRepoId_projectId)
	}
	if value, ok := flagValues["cloudWorkspace_workspaceId_repoId_projectRepoId_repoName"]; ok {
		query_cloudWorkspace_workspaceId_repoId_projectRepoId_repoName, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.CloudWorkspaceWorkspaceIdRepoIdProjectRepoIdRepoName(query_cloudWorkspace_workspaceId_repoId_projectRepoId_repoName)
	}
	if value, ok := flagValues["cloudWorkspace_workspaceId_repoId_uid"]; ok {
		query_cloudWorkspace_workspaceId_repoId_uid, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.CloudWorkspaceWorkspaceIdRepoIdUid(query_cloudWorkspace_workspaceId_repoId_uid)
	}
	if value, ok := flagValues["gerrit_aliasContext_kind"]; ok {
		query_gerrit_aliasContext_kind, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.GerritAliasContextKind(query_gerrit_aliasContext_kind)
	}
	if value, ok := flagValues["gerrit_aliasContext_name"]; ok {
		query_gerrit_aliasContext_name, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.GerritAliasContextName(query_gerrit_aliasContext_name)
	}
	if value, ok := flagValues["gerrit_aliasName"]; ok {
		query_gerrit_aliasName, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.GerritAliasName(query_gerrit_aliasName)
	}
	if value, ok := flagValues["gerrit_gerritProject"]; ok {
		query_gerrit_gerritProject, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.GerritGerritProject(query_gerrit_gerritProject)
	}
	if value, ok := flagValues["gerrit_hostUri"]; ok {
		query_gerrit_hostUri, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.GerritHostUri(query_gerrit_hostUri)
	}
	if value, ok := flagValues["gerrit_revisionId"]; ok {
		query_gerrit_revisionId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.GerritRevisionId(query_gerrit_revisionId)
	}
	if value, ok := flagValues["git_revisionId"]; ok {
		query_git_revisionId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.GitRevisionId(query_git_revisionId)
	}
	if value, ok := flagValues["git_url"]; ok {
		query_git_url, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.GitUrl(query_git_url)
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
	if value, ok := flagValues["startPosition"]; ok {
		query_startPosition, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.StartPosition(query_startPosition)
	}
	if value, ok := flagValues["uid"]; ok {
		query_uid, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.Uid(query_uid)
	}

	var response *api_client.ReadResponse
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

func Source_v1_ProjectsReposRevisionsGet(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)

		usageBits += " [--uid=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsReposRevisionsService(api_service)

	queryParamNames := map[string]bool{
		"uid": false,
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
		"repoName",
		"revisionId",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_repoName, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}
	param_revisionId, err := commands_util.ConvertValue_string(paramValues[2])
	if err != nil {
		return err
	}

	call := service.Get(param_projectId, param_repoName, param_revisionId)

	// Set query parameters.
	if value, ok := flagValues["uid"]; ok {
		query_uid, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.Uid(query_uid)
	}

	var response *api_client.Revision
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

func Source_v1_ProjectsReposRevisionsGetBatchGet(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)

		usageBits += " [--revisionIds=VALUE]"

		usageBits += " [--uid=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsReposRevisionsService(api_service)

	queryParamNames := map[string]bool{
		"revisionIds": false,
		"uid":         false,
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
		"repoName",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_repoName, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}

	call := service.GetBatchGet(param_projectId, param_repoName)

	// Set query parameters.
	if value, ok := flagValues["revisionIds"]; ok {
		query_revisionIds, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.RevisionIds(query_revisionIds)
	}
	if value, ok := flagValues["uid"]; ok {
		query_uid, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.Uid(query_uid)
	}

	var response *api_client.GetRevisionsResponse
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

func Source_v1_ProjectsReposRevisionsList(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)

		usageBits += " [--ends=VALUE]"

		usageBits += " [--pageSize=VALUE]"

		usageBits += " [--pageToken=VALUE]"

		usageBits += " [--path=VALUE]"

		usageBits += " [--starts=VALUE]"

		usageBits += " [--uid=VALUE]"

		usageBits += " [--walkDirection=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsReposRevisionsService(api_service)

	queryParamNames := map[string]bool{
		"ends":          false,
		"pageSize":      false,
		"pageToken":     false,
		"path":          false,
		"starts":        false,
		"uid":           false,
		"walkDirection": false,
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
		"repoName",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_repoName, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}

	call := service.List(param_projectId, param_repoName)

	// Set query parameters.
	if value, ok := flagValues["ends"]; ok {
		query_ends, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.Ends(query_ends)
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
	if value, ok := flagValues["path"]; ok {
		query_path, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.Path(query_path)
	}
	if value, ok := flagValues["starts"]; ok {
		query_starts, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.Starts(query_starts)
	}
	if value, ok := flagValues["uid"]; ok {
		query_uid, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.Uid(query_uid)
	}
	if value, ok := flagValues["walkDirection"]; ok {
		query_walkDirection, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.WalkDirection(query_walkDirection)
	}

	var response *api_client.ListRevisionsResponse
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

func Source_v1_ProjectsReposRevisionsListFiles(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)

		usageBits += " [--aliasContext.kind=VALUE]"

		usageBits += " [--aliasContext.name=VALUE]"

		usageBits += " [--aliasName=VALUE]"

		usageBits += " [--cloudWorkspace.snapshotId=VALUE]"

		usageBits += " [--cloudWorkspace.workspaceId.name=VALUE]"

		usageBits += " [--cloudWorkspace.workspaceId.repoId.projectRepoId.projectId=VALUE]"

		usageBits += " [--cloudWorkspace.workspaceId.repoId.projectRepoId.repoName=VALUE]"

		usageBits += " [--cloudWorkspace.workspaceId.repoId.uid=VALUE]"

		usageBits += " [--gerrit.aliasContext.kind=VALUE]"

		usageBits += " [--gerrit.aliasContext.name=VALUE]"

		usageBits += " [--gerrit.aliasName=VALUE]"

		usageBits += " [--gerrit.gerritProject=VALUE]"

		usageBits += " [--gerrit.hostUri=VALUE]"

		usageBits += " [--gerrit.revisionId=VALUE]"

		usageBits += " [--git.revisionId=VALUE]"

		usageBits += " [--git.url=VALUE]"

		usageBits += " [--pageSize=VALUE]"

		usageBits += " [--pageToken=VALUE]"

		usageBits += " [--uid=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsReposRevisionsService(api_service)

	queryParamNames := map[string]bool{
		"aliasContext.kind":                                         false,
		"aliasContext.name":                                         false,
		"aliasName":                                                 false,
		"cloudWorkspace.snapshotId":                                 false,
		"cloudWorkspace.workspaceId.name":                           false,
		"cloudWorkspace.workspaceId.repoId.projectRepoId.projectId": false,
		"cloudWorkspace.workspaceId.repoId.projectRepoId.repoName":  false,
		"cloudWorkspace.workspaceId.repoId.uid":                     false,
		"gerrit.aliasContext.kind":                                  false,
		"gerrit.aliasContext.name":                                  false,
		"gerrit.aliasName":                                          false,
		"gerrit.gerritProject":                                      false,
		"gerrit.hostUri":                                            false,
		"gerrit.revisionId":                                         false,
		"git.revisionId":                                            false,
		"git.url":                                                   false,
		"pageSize":                                                  false,
		"pageToken":                                                 false,
		"uid":                                                       false,
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
		"repoName",
		"revisionId",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_repoName, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}
	param_revisionId, err := commands_util.ConvertValue_string(paramValues[2])
	if err != nil {
		return err
	}

	call := service.ListFiles(param_projectId, param_repoName, param_revisionId)

	// Set query parameters.
	if value, ok := flagValues["aliasContext_kind"]; ok {
		query_aliasContext_kind, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.AliasContextKind(query_aliasContext_kind)
	}
	if value, ok := flagValues["aliasContext_name"]; ok {
		query_aliasContext_name, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.AliasContextName(query_aliasContext_name)
	}
	if value, ok := flagValues["aliasName"]; ok {
		query_aliasName, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.AliasName(query_aliasName)
	}
	if value, ok := flagValues["cloudWorkspace_snapshotId"]; ok {
		query_cloudWorkspace_snapshotId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.CloudWorkspaceSnapshotId(query_cloudWorkspace_snapshotId)
	}
	if value, ok := flagValues["cloudWorkspace_workspaceId_name"]; ok {
		query_cloudWorkspace_workspaceId_name, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.CloudWorkspaceWorkspaceIdName(query_cloudWorkspace_workspaceId_name)
	}
	if value, ok := flagValues["cloudWorkspace_workspaceId_repoId_projectRepoId_projectId"]; ok {
		query_cloudWorkspace_workspaceId_repoId_projectRepoId_projectId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.CloudWorkspaceWorkspaceIdRepoIdProjectRepoIdProjectId(query_cloudWorkspace_workspaceId_repoId_projectRepoId_projectId)
	}
	if value, ok := flagValues["cloudWorkspace_workspaceId_repoId_projectRepoId_repoName"]; ok {
		query_cloudWorkspace_workspaceId_repoId_projectRepoId_repoName, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.CloudWorkspaceWorkspaceIdRepoIdProjectRepoIdRepoName(query_cloudWorkspace_workspaceId_repoId_projectRepoId_repoName)
	}
	if value, ok := flagValues["cloudWorkspace_workspaceId_repoId_uid"]; ok {
		query_cloudWorkspace_workspaceId_repoId_uid, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.CloudWorkspaceWorkspaceIdRepoIdUid(query_cloudWorkspace_workspaceId_repoId_uid)
	}
	if value, ok := flagValues["gerrit_aliasContext_kind"]; ok {
		query_gerrit_aliasContext_kind, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.GerritAliasContextKind(query_gerrit_aliasContext_kind)
	}
	if value, ok := flagValues["gerrit_aliasContext_name"]; ok {
		query_gerrit_aliasContext_name, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.GerritAliasContextName(query_gerrit_aliasContext_name)
	}
	if value, ok := flagValues["gerrit_aliasName"]; ok {
		query_gerrit_aliasName, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.GerritAliasName(query_gerrit_aliasName)
	}
	if value, ok := flagValues["gerrit_gerritProject"]; ok {
		query_gerrit_gerritProject, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.GerritGerritProject(query_gerrit_gerritProject)
	}
	if value, ok := flagValues["gerrit_hostUri"]; ok {
		query_gerrit_hostUri, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.GerritHostUri(query_gerrit_hostUri)
	}
	if value, ok := flagValues["gerrit_revisionId"]; ok {
		query_gerrit_revisionId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.GerritRevisionId(query_gerrit_revisionId)
	}
	if value, ok := flagValues["git_revisionId"]; ok {
		query_git_revisionId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.GitRevisionId(query_git_revisionId)
	}
	if value, ok := flagValues["git_url"]; ok {
		query_git_url, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.GitUrl(query_git_url)
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
	if value, ok := flagValues["uid"]; ok {
		query_uid, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.Uid(query_uid)
	}

	var response *api_client.ListFilesResponse
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

func Source_v1_ProjectsReposWorkspacesCommitWorkspace(context Context, args ...string) error {

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
	service := api_client.NewProjectsReposWorkspacesService(api_service)

	args, flagValues, err := commands_util.ExtractFlagValues(args)
	if err != nil {
		return err
	}

	// Only positional arguments should remain in args.
	if len(args) == 0 || len(args) > 2 {
		usageFunc()
	}

	request := &api_client.CommitWorkspaceRequest{}
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
		"repoName",
		"name",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_repoName, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}
	param_name, err := commands_util.ConvertValue_string(paramValues[2])
	if err != nil {
		return err
	}

	call := service.CommitWorkspace(param_projectId, param_repoName, param_name,
		request,
	)

	var response *api_client.Workspace
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

func Source_v1_ProjectsReposWorkspacesCreate(context Context, args ...string) error {

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
	service := api_client.NewProjectsReposWorkspacesService(api_service)

	args, flagValues, err := commands_util.ExtractFlagValues(args)
	if err != nil {
		return err
	}

	// Only positional arguments should remain in args.
	if len(args) == 0 || len(args) > 2 {
		usageFunc()
	}

	request := &api_client.CreateWorkspaceRequest{}
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
		"repoName",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_repoName, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}

	call := service.Create(param_projectId, param_repoName,
		request,
	)

	var response *api_client.Workspace
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

func Source_v1_ProjectsReposWorkspacesDelete(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)

		usageBits += " [--currentSnapshotId=VALUE]"

		usageBits += " [--uid=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsReposWorkspacesService(api_service)

	queryParamNames := map[string]bool{
		"currentSnapshotId": false,
		"uid":               false,
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
		"repoName",
		"name",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_repoName, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}
	param_name, err := commands_util.ConvertValue_string(paramValues[2])
	if err != nil {
		return err
	}

	call := service.Delete(param_projectId, param_repoName, param_name)

	// Set query parameters.
	if value, ok := flagValues["currentSnapshotId"]; ok {
		query_currentSnapshotId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.CurrentSnapshotId(query_currentSnapshotId)
	}
	if value, ok := flagValues["uid"]; ok {
		query_uid, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.Uid(query_uid)
	}

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

func Source_v1_ProjectsReposWorkspacesFilesGet(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)

		usageBits += " [--cloudRepo.aliasContext.kind=VALUE]"

		usageBits += " [--cloudRepo.aliasContext.name=VALUE]"

		usageBits += " [--cloudRepo.aliasName=VALUE]"

		usageBits += " [--cloudRepo.repoId.projectRepoId.projectId=VALUE]"

		usageBits += " [--cloudRepo.repoId.projectRepoId.repoName=VALUE]"

		usageBits += " [--cloudRepo.repoId.uid=VALUE]"

		usageBits += " [--cloudRepo.revisionId=VALUE]"

		usageBits += " [--gerrit.aliasContext.kind=VALUE]"

		usageBits += " [--gerrit.aliasContext.name=VALUE]"

		usageBits += " [--gerrit.aliasName=VALUE]"

		usageBits += " [--gerrit.gerritProject=VALUE]"

		usageBits += " [--gerrit.hostUri=VALUE]"

		usageBits += " [--gerrit.revisionId=VALUE]"

		usageBits += " [--git.revisionId=VALUE]"

		usageBits += " [--git.url=VALUE]"

		usageBits += " [--pageSize=VALUE]"

		usageBits += " [--pageToken=VALUE]"

		usageBits += " [--snapshotId=VALUE]"

		usageBits += " [--startPosition=VALUE]"

		usageBits += " [--uid=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsReposWorkspacesFilesService(api_service)

	queryParamNames := map[string]bool{
		"cloudRepo.aliasContext.kind":              false,
		"cloudRepo.aliasContext.name":              false,
		"cloudRepo.aliasName":                      false,
		"cloudRepo.repoId.projectRepoId.projectId": false,
		"cloudRepo.repoId.projectRepoId.repoName":  false,
		"cloudRepo.repoId.uid":                     false,
		"cloudRepo.revisionId":                     false,
		"gerrit.aliasContext.kind":                 false,
		"gerrit.aliasContext.name":                 false,
		"gerrit.aliasName":                         false,
		"gerrit.gerritProject":                     false,
		"gerrit.hostUri":                           false,
		"gerrit.revisionId":                        false,
		"git.revisionId":                           false,
		"git.url":                                  false,
		"pageSize":                                 false,
		"pageToken":                                false,
		"snapshotId":                               false,
		"startPosition":                            false,
		"uid":                                      false,
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
		"repoName",
		"name",
		"path",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_repoName, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}
	param_name, err := commands_util.ConvertValue_string(paramValues[2])
	if err != nil {
		return err
	}
	param_path, err := commands_util.ConvertValue_string(paramValues[3])
	if err != nil {
		return err
	}

	call := service.Get(param_projectId, param_repoName, param_name, param_path)

	// Set query parameters.
	if value, ok := flagValues["cloudRepo_aliasContext_kind"]; ok {
		query_cloudRepo_aliasContext_kind, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.CloudRepoAliasContextKind(query_cloudRepo_aliasContext_kind)
	}
	if value, ok := flagValues["cloudRepo_aliasContext_name"]; ok {
		query_cloudRepo_aliasContext_name, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.CloudRepoAliasContextName(query_cloudRepo_aliasContext_name)
	}
	if value, ok := flagValues["cloudRepo_aliasName"]; ok {
		query_cloudRepo_aliasName, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.CloudRepoAliasName(query_cloudRepo_aliasName)
	}
	if value, ok := flagValues["cloudRepo_repoId_projectRepoId_projectId"]; ok {
		query_cloudRepo_repoId_projectRepoId_projectId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.CloudRepoRepoIdProjectRepoIdProjectId(query_cloudRepo_repoId_projectRepoId_projectId)
	}
	if value, ok := flagValues["cloudRepo_repoId_projectRepoId_repoName"]; ok {
		query_cloudRepo_repoId_projectRepoId_repoName, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.CloudRepoRepoIdProjectRepoIdRepoName(query_cloudRepo_repoId_projectRepoId_repoName)
	}
	if value, ok := flagValues["cloudRepo_repoId_uid"]; ok {
		query_cloudRepo_repoId_uid, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.CloudRepoRepoIdUid(query_cloudRepo_repoId_uid)
	}
	if value, ok := flagValues["cloudRepo_revisionId"]; ok {
		query_cloudRepo_revisionId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.CloudRepoRevisionId(query_cloudRepo_revisionId)
	}
	if value, ok := flagValues["gerrit_aliasContext_kind"]; ok {
		query_gerrit_aliasContext_kind, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.GerritAliasContextKind(query_gerrit_aliasContext_kind)
	}
	if value, ok := flagValues["gerrit_aliasContext_name"]; ok {
		query_gerrit_aliasContext_name, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.GerritAliasContextName(query_gerrit_aliasContext_name)
	}
	if value, ok := flagValues["gerrit_aliasName"]; ok {
		query_gerrit_aliasName, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.GerritAliasName(query_gerrit_aliasName)
	}
	if value, ok := flagValues["gerrit_gerritProject"]; ok {
		query_gerrit_gerritProject, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.GerritGerritProject(query_gerrit_gerritProject)
	}
	if value, ok := flagValues["gerrit_hostUri"]; ok {
		query_gerrit_hostUri, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.GerritHostUri(query_gerrit_hostUri)
	}
	if value, ok := flagValues["gerrit_revisionId"]; ok {
		query_gerrit_revisionId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.GerritRevisionId(query_gerrit_revisionId)
	}
	if value, ok := flagValues["git_revisionId"]; ok {
		query_git_revisionId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.GitRevisionId(query_git_revisionId)
	}
	if value, ok := flagValues["git_url"]; ok {
		query_git_url, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.GitUrl(query_git_url)
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
	if value, ok := flagValues["snapshotId"]; ok {
		query_snapshotId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SnapshotId(query_snapshotId)
	}
	if value, ok := flagValues["startPosition"]; ok {
		query_startPosition, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.StartPosition(query_startPosition)
	}
	if value, ok := flagValues["uid"]; ok {
		query_uid, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.Uid(query_uid)
	}

	var response *api_client.ReadResponse
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

func Source_v1_ProjectsReposWorkspacesGet(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)

		usageBits += " [--uid=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsReposWorkspacesService(api_service)

	queryParamNames := map[string]bool{
		"uid": false,
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
		"repoName",
		"name",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_repoName, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}
	param_name, err := commands_util.ConvertValue_string(paramValues[2])
	if err != nil {
		return err
	}

	call := service.Get(param_projectId, param_repoName, param_name)

	// Set query parameters.
	if value, ok := flagValues["uid"]; ok {
		query_uid, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.Uid(query_uid)
	}

	var response *api_client.Workspace
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

func Source_v1_ProjectsReposWorkspacesList(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)

		usageBits += " [--uid=VALUE]"

		usageBits += " [--view=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsReposWorkspacesService(api_service)

	queryParamNames := map[string]bool{
		"uid":  false,
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
		"repoName",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_repoName, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}

	call := service.List(param_projectId, param_repoName)

	// Set query parameters.
	if value, ok := flagValues["uid"]; ok {
		query_uid, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.Uid(query_uid)
	}
	if value, ok := flagValues["view"]; ok {
		query_view, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.View(query_view)
	}

	var response *api_client.ListWorkspacesResponse
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

func Source_v1_ProjectsReposWorkspacesListFiles(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)

		usageBits += " [--cloudRepo.aliasContext.kind=VALUE]"

		usageBits += " [--cloudRepo.aliasContext.name=VALUE]"

		usageBits += " [--cloudRepo.aliasName=VALUE]"

		usageBits += " [--cloudRepo.repoId.projectRepoId.projectId=VALUE]"

		usageBits += " [--cloudRepo.repoId.projectRepoId.repoName=VALUE]"

		usageBits += " [--cloudRepo.repoId.uid=VALUE]"

		usageBits += " [--cloudRepo.revisionId=VALUE]"

		usageBits += " [--gerrit.aliasContext.kind=VALUE]"

		usageBits += " [--gerrit.aliasContext.name=VALUE]"

		usageBits += " [--gerrit.aliasName=VALUE]"

		usageBits += " [--gerrit.gerritProject=VALUE]"

		usageBits += " [--gerrit.hostUri=VALUE]"

		usageBits += " [--gerrit.revisionId=VALUE]"

		usageBits += " [--git.revisionId=VALUE]"

		usageBits += " [--git.url=VALUE]"

		usageBits += " [--pageSize=VALUE]"

		usageBits += " [--pageToken=VALUE]"

		usageBits += " [--snapshotId=VALUE]"

		usageBits += " [--uid=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsReposWorkspacesService(api_service)

	queryParamNames := map[string]bool{
		"cloudRepo.aliasContext.kind":              false,
		"cloudRepo.aliasContext.name":              false,
		"cloudRepo.aliasName":                      false,
		"cloudRepo.repoId.projectRepoId.projectId": false,
		"cloudRepo.repoId.projectRepoId.repoName":  false,
		"cloudRepo.repoId.uid":                     false,
		"cloudRepo.revisionId":                     false,
		"gerrit.aliasContext.kind":                 false,
		"gerrit.aliasContext.name":                 false,
		"gerrit.aliasName":                         false,
		"gerrit.gerritProject":                     false,
		"gerrit.hostUri":                           false,
		"gerrit.revisionId":                        false,
		"git.revisionId":                           false,
		"git.url":                                  false,
		"pageSize":                                 false,
		"pageToken":                                false,
		"snapshotId":                               false,
		"uid":                                      false,
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
		"repoName",
		"name",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_repoName, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}
	param_name, err := commands_util.ConvertValue_string(paramValues[2])
	if err != nil {
		return err
	}

	call := service.ListFiles(param_projectId, param_repoName, param_name)

	// Set query parameters.
	if value, ok := flagValues["cloudRepo_aliasContext_kind"]; ok {
		query_cloudRepo_aliasContext_kind, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.CloudRepoAliasContextKind(query_cloudRepo_aliasContext_kind)
	}
	if value, ok := flagValues["cloudRepo_aliasContext_name"]; ok {
		query_cloudRepo_aliasContext_name, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.CloudRepoAliasContextName(query_cloudRepo_aliasContext_name)
	}
	if value, ok := flagValues["cloudRepo_aliasName"]; ok {
		query_cloudRepo_aliasName, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.CloudRepoAliasName(query_cloudRepo_aliasName)
	}
	if value, ok := flagValues["cloudRepo_repoId_projectRepoId_projectId"]; ok {
		query_cloudRepo_repoId_projectRepoId_projectId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.CloudRepoRepoIdProjectRepoIdProjectId(query_cloudRepo_repoId_projectRepoId_projectId)
	}
	if value, ok := flagValues["cloudRepo_repoId_projectRepoId_repoName"]; ok {
		query_cloudRepo_repoId_projectRepoId_repoName, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.CloudRepoRepoIdProjectRepoIdRepoName(query_cloudRepo_repoId_projectRepoId_repoName)
	}
	if value, ok := flagValues["cloudRepo_repoId_uid"]; ok {
		query_cloudRepo_repoId_uid, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.CloudRepoRepoIdUid(query_cloudRepo_repoId_uid)
	}
	if value, ok := flagValues["cloudRepo_revisionId"]; ok {
		query_cloudRepo_revisionId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.CloudRepoRevisionId(query_cloudRepo_revisionId)
	}
	if value, ok := flagValues["gerrit_aliasContext_kind"]; ok {
		query_gerrit_aliasContext_kind, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.GerritAliasContextKind(query_gerrit_aliasContext_kind)
	}
	if value, ok := flagValues["gerrit_aliasContext_name"]; ok {
		query_gerrit_aliasContext_name, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.GerritAliasContextName(query_gerrit_aliasContext_name)
	}
	if value, ok := flagValues["gerrit_aliasName"]; ok {
		query_gerrit_aliasName, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.GerritAliasName(query_gerrit_aliasName)
	}
	if value, ok := flagValues["gerrit_gerritProject"]; ok {
		query_gerrit_gerritProject, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.GerritGerritProject(query_gerrit_gerritProject)
	}
	if value, ok := flagValues["gerrit_hostUri"]; ok {
		query_gerrit_hostUri, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.GerritHostUri(query_gerrit_hostUri)
	}
	if value, ok := flagValues["gerrit_revisionId"]; ok {
		query_gerrit_revisionId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.GerritRevisionId(query_gerrit_revisionId)
	}
	if value, ok := flagValues["git_revisionId"]; ok {
		query_git_revisionId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.GitRevisionId(query_git_revisionId)
	}
	if value, ok := flagValues["git_url"]; ok {
		query_git_url, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.GitUrl(query_git_url)
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
	if value, ok := flagValues["snapshotId"]; ok {
		query_snapshotId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SnapshotId(query_snapshotId)
	}
	if value, ok := flagValues["uid"]; ok {
		query_uid, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.Uid(query_uid)
	}

	var response *api_client.ListFilesResponse
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

func Source_v1_ProjectsReposWorkspacesModifyWorkspace(context Context, args ...string) error {

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
	service := api_client.NewProjectsReposWorkspacesService(api_service)

	args, flagValues, err := commands_util.ExtractFlagValues(args)
	if err != nil {
		return err
	}

	// Only positional arguments should remain in args.
	if len(args) == 0 || len(args) > 2 {
		usageFunc()
	}

	request := &api_client.ModifyWorkspaceRequest{}
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
		"repoName",
		"name",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_repoName, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}
	param_name, err := commands_util.ConvertValue_string(paramValues[2])
	if err != nil {
		return err
	}

	call := service.ModifyWorkspace(param_projectId, param_repoName, param_name,
		request,
	)

	var response *api_client.Workspace
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

func Source_v1_ProjectsReposWorkspacesRefreshWorkspace(context Context, args ...string) error {

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
	service := api_client.NewProjectsReposWorkspacesService(api_service)

	args, flagValues, err := commands_util.ExtractFlagValues(args)
	if err != nil {
		return err
	}

	// Only positional arguments should remain in args.
	if len(args) == 0 || len(args) > 2 {
		usageFunc()
	}

	request := &api_client.RefreshWorkspaceRequest{}
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
		"repoName",
		"name",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_repoName, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}
	param_name, err := commands_util.ConvertValue_string(paramValues[2])
	if err != nil {
		return err
	}

	call := service.RefreshWorkspace(param_projectId, param_repoName, param_name,
		request,
	)

	var response *api_client.Workspace
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

func Source_v1_ProjectsReposWorkspacesResolveFiles(context Context, args ...string) error {

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
	service := api_client.NewProjectsReposWorkspacesService(api_service)

	args, flagValues, err := commands_util.ExtractFlagValues(args)
	if err != nil {
		return err
	}

	// Only positional arguments should remain in args.
	if len(args) == 0 || len(args) > 2 {
		usageFunc()
	}

	request := &api_client.ResolveFilesRequest{}
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
		"repoName",
		"name",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_repoName, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}
	param_name, err := commands_util.ConvertValue_string(paramValues[2])
	if err != nil {
		return err
	}

	call := service.ResolveFiles(param_projectId, param_repoName, param_name,
		request,
	)

	var response *api_client.Workspace
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

func Source_v1_ProjectsReposWorkspacesRevertRefresh(context Context, args ...string) error {

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
	service := api_client.NewProjectsReposWorkspacesService(api_service)

	args, flagValues, err := commands_util.ExtractFlagValues(args)
	if err != nil {
		return err
	}

	// Only positional arguments should remain in args.
	if len(args) == 0 || len(args) > 2 {
		usageFunc()
	}

	request := &api_client.RevertRefreshRequest{}
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
		"repoName",
		"name",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_repoName, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}
	param_name, err := commands_util.ConvertValue_string(paramValues[2])
	if err != nil {
		return err
	}

	call := service.RevertRefresh(param_projectId, param_repoName, param_name,
		request,
	)

	var response *api_client.Workspace
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

func Source_v1_ProjectsReposWorkspacesSnapshotsFilesGet(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)

		usageBits += " [--cloudRepo.aliasContext.kind=VALUE]"

		usageBits += " [--cloudRepo.aliasContext.name=VALUE]"

		usageBits += " [--cloudRepo.aliasName=VALUE]"

		usageBits += " [--cloudRepo.repoId.projectRepoId.projectId=VALUE]"

		usageBits += " [--cloudRepo.repoId.projectRepoId.repoName=VALUE]"

		usageBits += " [--cloudRepo.repoId.uid=VALUE]"

		usageBits += " [--cloudRepo.revisionId=VALUE]"

		usageBits += " [--gerrit.aliasContext.kind=VALUE]"

		usageBits += " [--gerrit.aliasContext.name=VALUE]"

		usageBits += " [--gerrit.aliasName=VALUE]"

		usageBits += " [--gerrit.gerritProject=VALUE]"

		usageBits += " [--gerrit.hostUri=VALUE]"

		usageBits += " [--gerrit.revisionId=VALUE]"

		usageBits += " [--git.revisionId=VALUE]"

		usageBits += " [--git.url=VALUE]"

		usageBits += " [--pageSize=VALUE]"

		usageBits += " [--pageToken=VALUE]"

		usageBits += " [--startPosition=VALUE]"

		usageBits += " [--uid=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsReposWorkspacesSnapshotsFilesService(api_service)

	queryParamNames := map[string]bool{
		"cloudRepo.aliasContext.kind":              false,
		"cloudRepo.aliasContext.name":              false,
		"cloudRepo.aliasName":                      false,
		"cloudRepo.repoId.projectRepoId.projectId": false,
		"cloudRepo.repoId.projectRepoId.repoName":  false,
		"cloudRepo.repoId.uid":                     false,
		"cloudRepo.revisionId":                     false,
		"gerrit.aliasContext.kind":                 false,
		"gerrit.aliasContext.name":                 false,
		"gerrit.aliasName":                         false,
		"gerrit.gerritProject":                     false,
		"gerrit.hostUri":                           false,
		"gerrit.revisionId":                        false,
		"git.revisionId":                           false,
		"git.url":                                  false,
		"pageSize":                                 false,
		"pageToken":                                false,
		"startPosition":                            false,
		"uid":                                      false,
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
		"repoName",
		"name",
		"snapshotId",
		"path",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_repoName, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}
	param_name, err := commands_util.ConvertValue_string(paramValues[2])
	if err != nil {
		return err
	}
	param_snapshotId, err := commands_util.ConvertValue_string(paramValues[3])
	if err != nil {
		return err
	}
	param_path, err := commands_util.ConvertValue_string(paramValues[4])
	if err != nil {
		return err
	}

	call := service.Get(param_projectId, param_repoName, param_name, param_snapshotId, param_path)

	// Set query parameters.
	if value, ok := flagValues["cloudRepo_aliasContext_kind"]; ok {
		query_cloudRepo_aliasContext_kind, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.CloudRepoAliasContextKind(query_cloudRepo_aliasContext_kind)
	}
	if value, ok := flagValues["cloudRepo_aliasContext_name"]; ok {
		query_cloudRepo_aliasContext_name, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.CloudRepoAliasContextName(query_cloudRepo_aliasContext_name)
	}
	if value, ok := flagValues["cloudRepo_aliasName"]; ok {
		query_cloudRepo_aliasName, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.CloudRepoAliasName(query_cloudRepo_aliasName)
	}
	if value, ok := flagValues["cloudRepo_repoId_projectRepoId_projectId"]; ok {
		query_cloudRepo_repoId_projectRepoId_projectId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.CloudRepoRepoIdProjectRepoIdProjectId(query_cloudRepo_repoId_projectRepoId_projectId)
	}
	if value, ok := flagValues["cloudRepo_repoId_projectRepoId_repoName"]; ok {
		query_cloudRepo_repoId_projectRepoId_repoName, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.CloudRepoRepoIdProjectRepoIdRepoName(query_cloudRepo_repoId_projectRepoId_repoName)
	}
	if value, ok := flagValues["cloudRepo_repoId_uid"]; ok {
		query_cloudRepo_repoId_uid, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.CloudRepoRepoIdUid(query_cloudRepo_repoId_uid)
	}
	if value, ok := flagValues["cloudRepo_revisionId"]; ok {
		query_cloudRepo_revisionId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.CloudRepoRevisionId(query_cloudRepo_revisionId)
	}
	if value, ok := flagValues["gerrit_aliasContext_kind"]; ok {
		query_gerrit_aliasContext_kind, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.GerritAliasContextKind(query_gerrit_aliasContext_kind)
	}
	if value, ok := flagValues["gerrit_aliasContext_name"]; ok {
		query_gerrit_aliasContext_name, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.GerritAliasContextName(query_gerrit_aliasContext_name)
	}
	if value, ok := flagValues["gerrit_aliasName"]; ok {
		query_gerrit_aliasName, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.GerritAliasName(query_gerrit_aliasName)
	}
	if value, ok := flagValues["gerrit_gerritProject"]; ok {
		query_gerrit_gerritProject, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.GerritGerritProject(query_gerrit_gerritProject)
	}
	if value, ok := flagValues["gerrit_hostUri"]; ok {
		query_gerrit_hostUri, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.GerritHostUri(query_gerrit_hostUri)
	}
	if value, ok := flagValues["gerrit_revisionId"]; ok {
		query_gerrit_revisionId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.GerritRevisionId(query_gerrit_revisionId)
	}
	if value, ok := flagValues["git_revisionId"]; ok {
		query_git_revisionId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.GitRevisionId(query_git_revisionId)
	}
	if value, ok := flagValues["git_url"]; ok {
		query_git_url, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.GitUrl(query_git_url)
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
	if value, ok := flagValues["startPosition"]; ok {
		query_startPosition, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.StartPosition(query_startPosition)
	}
	if value, ok := flagValues["uid"]; ok {
		query_uid, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.Uid(query_uid)
	}

	var response *api_client.ReadResponse
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

func Source_v1_ProjectsReposWorkspacesSnapshotsGet(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)

		usageBits += " [--uid=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsReposWorkspacesSnapshotsService(api_service)

	queryParamNames := map[string]bool{
		"uid": false,
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
		"repoName",
		"name",
		"snapshotId",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_repoName, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}
	param_name, err := commands_util.ConvertValue_string(paramValues[2])
	if err != nil {
		return err
	}
	param_snapshotId, err := commands_util.ConvertValue_string(paramValues[3])
	if err != nil {
		return err
	}

	call := service.Get(param_projectId, param_repoName, param_name, param_snapshotId)

	// Set query parameters.
	if value, ok := flagValues["uid"]; ok {
		query_uid, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.Uid(query_uid)
	}

	var response *api_client.Snapshot
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

func Source_v1_ProjectsReposWorkspacesSnapshotsList(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)

		usageBits += " [--pageSize=VALUE]"

		usageBits += " [--pageToken=VALUE]"

		usageBits += " [--uid=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsReposWorkspacesSnapshotsService(api_service)

	queryParamNames := map[string]bool{
		"pageSize":  false,
		"pageToken": false,
		"uid":       false,
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
		"repoName",
		"name",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_repoName, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}
	param_name, err := commands_util.ConvertValue_string(paramValues[2])
	if err != nil {
		return err
	}

	call := service.List(param_projectId, param_repoName, param_name)

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
	if value, ok := flagValues["uid"]; ok {
		query_uid, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.Uid(query_uid)
	}

	var response *api_client.ListSnapshotsResponse
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

func Source_v1_ProjectsReposWorkspacesSnapshotsListFiles(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)

		usageBits += " [--cloudRepo.aliasContext.kind=VALUE]"

		usageBits += " [--cloudRepo.aliasContext.name=VALUE]"

		usageBits += " [--cloudRepo.aliasName=VALUE]"

		usageBits += " [--cloudRepo.repoId.projectRepoId.projectId=VALUE]"

		usageBits += " [--cloudRepo.repoId.projectRepoId.repoName=VALUE]"

		usageBits += " [--cloudRepo.repoId.uid=VALUE]"

		usageBits += " [--cloudRepo.revisionId=VALUE]"

		usageBits += " [--gerrit.aliasContext.kind=VALUE]"

		usageBits += " [--gerrit.aliasContext.name=VALUE]"

		usageBits += " [--gerrit.aliasName=VALUE]"

		usageBits += " [--gerrit.gerritProject=VALUE]"

		usageBits += " [--gerrit.hostUri=VALUE]"

		usageBits += " [--gerrit.revisionId=VALUE]"

		usageBits += " [--git.revisionId=VALUE]"

		usageBits += " [--git.url=VALUE]"

		usageBits += " [--pageSize=VALUE]"

		usageBits += " [--pageToken=VALUE]"

		usageBits += " [--uid=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsReposWorkspacesSnapshotsService(api_service)

	queryParamNames := map[string]bool{
		"cloudRepo.aliasContext.kind":              false,
		"cloudRepo.aliasContext.name":              false,
		"cloudRepo.aliasName":                      false,
		"cloudRepo.repoId.projectRepoId.projectId": false,
		"cloudRepo.repoId.projectRepoId.repoName":  false,
		"cloudRepo.repoId.uid":                     false,
		"cloudRepo.revisionId":                     false,
		"gerrit.aliasContext.kind":                 false,
		"gerrit.aliasContext.name":                 false,
		"gerrit.aliasName":                         false,
		"gerrit.gerritProject":                     false,
		"gerrit.hostUri":                           false,
		"gerrit.revisionId":                        false,
		"git.revisionId":                           false,
		"git.url":                                  false,
		"pageSize":                                 false,
		"pageToken":                                false,
		"uid":                                      false,
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
		"repoName",
		"name",
		"snapshotId",
	}
	paramValues := commands_util.SplitParamValues(args[0])
	if len(paramValues) != len(expectedParams) {
		return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
	}

	param_projectId, err := commands_util.ConvertValue_string(paramValues[0])
	if err != nil {
		return err
	}
	param_repoName, err := commands_util.ConvertValue_string(paramValues[1])
	if err != nil {
		return err
	}
	param_name, err := commands_util.ConvertValue_string(paramValues[2])
	if err != nil {
		return err
	}
	param_snapshotId, err := commands_util.ConvertValue_string(paramValues[3])
	if err != nil {
		return err
	}

	call := service.ListFiles(param_projectId, param_repoName, param_name, param_snapshotId)

	// Set query parameters.
	if value, ok := flagValues["cloudRepo_aliasContext_kind"]; ok {
		query_cloudRepo_aliasContext_kind, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.CloudRepoAliasContextKind(query_cloudRepo_aliasContext_kind)
	}
	if value, ok := flagValues["cloudRepo_aliasContext_name"]; ok {
		query_cloudRepo_aliasContext_name, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.CloudRepoAliasContextName(query_cloudRepo_aliasContext_name)
	}
	if value, ok := flagValues["cloudRepo_aliasName"]; ok {
		query_cloudRepo_aliasName, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.CloudRepoAliasName(query_cloudRepo_aliasName)
	}
	if value, ok := flagValues["cloudRepo_repoId_projectRepoId_projectId"]; ok {
		query_cloudRepo_repoId_projectRepoId_projectId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.CloudRepoRepoIdProjectRepoIdProjectId(query_cloudRepo_repoId_projectRepoId_projectId)
	}
	if value, ok := flagValues["cloudRepo_repoId_projectRepoId_repoName"]; ok {
		query_cloudRepo_repoId_projectRepoId_repoName, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.CloudRepoRepoIdProjectRepoIdRepoName(query_cloudRepo_repoId_projectRepoId_repoName)
	}
	if value, ok := flagValues["cloudRepo_repoId_uid"]; ok {
		query_cloudRepo_repoId_uid, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.CloudRepoRepoIdUid(query_cloudRepo_repoId_uid)
	}
	if value, ok := flagValues["cloudRepo_revisionId"]; ok {
		query_cloudRepo_revisionId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.CloudRepoRevisionId(query_cloudRepo_revisionId)
	}
	if value, ok := flagValues["gerrit_aliasContext_kind"]; ok {
		query_gerrit_aliasContext_kind, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.GerritAliasContextKind(query_gerrit_aliasContext_kind)
	}
	if value, ok := flagValues["gerrit_aliasContext_name"]; ok {
		query_gerrit_aliasContext_name, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.GerritAliasContextName(query_gerrit_aliasContext_name)
	}
	if value, ok := flagValues["gerrit_aliasName"]; ok {
		query_gerrit_aliasName, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.GerritAliasName(query_gerrit_aliasName)
	}
	if value, ok := flagValues["gerrit_gerritProject"]; ok {
		query_gerrit_gerritProject, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.GerritGerritProject(query_gerrit_gerritProject)
	}
	if value, ok := flagValues["gerrit_hostUri"]; ok {
		query_gerrit_hostUri, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.GerritHostUri(query_gerrit_hostUri)
	}
	if value, ok := flagValues["gerrit_revisionId"]; ok {
		query_gerrit_revisionId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.GerritRevisionId(query_gerrit_revisionId)
	}
	if value, ok := flagValues["git_revisionId"]; ok {
		query_git_revisionId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.GitRevisionId(query_git_revisionId)
	}
	if value, ok := flagValues["git_url"]; ok {
		query_git_url, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.GitUrl(query_git_url)
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
	if value, ok := flagValues["uid"]; ok {
		query_uid, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.Uid(query_uid)
	}

	var response *api_client.ListFilesResponse
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
