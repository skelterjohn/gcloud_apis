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
		var pathParams []string

		if len(pathParams) != 0 {
			if strings.Contains("v1:listChangedFiles", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		commands_util.PrintRequestExample(&api_client.ListChangedFilesRequest{})

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

func Source_v1_ProjectsReposAliasesCreate(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("repoName"))

		if len(pathParams) != 0 {
			if strings.Contains("v1/projects/{projectId}/repos/{repoName}/aliases", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		usageBits += " [--repoId.uid=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		commands_util.PrintRequestExample(&api_client.Alias{})

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsReposAliasesService(api_service)

	queryParamNames := map[string]bool{
		"repoId.uid": false,
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
	if value, ok := flagValues["repoId_uid"]; ok {
		query_repoId_uid, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.RepoIdUid(query_repoId_uid)
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

func Source_v1_ProjectsReposAliasesDelete(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("repoName"))
		pathParams = append(pathParams, commands_util.AngrySnakes("kind"))
		pathParams = append(pathParams, commands_util.AngrySnakes("name"))

		if len(pathParams) != 0 {
			if strings.Contains("v1/projects/{projectId}/repos/{repoName}/aliases/{kind}/{name}", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [--repoId.uid=VALUE]"

		usageBits += " [--revisionId=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsReposAliasesService(api_service)

	queryParamNames := map[string]bool{
		"repoId.uid": false,
		"revisionId": false,
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
	if value, ok := flagValues["repoId_uid"]; ok {
		query_repoId_uid, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.RepoIdUid(query_repoId_uid)
	}
	if value, ok := flagValues["revisionId"]; ok {
		query_revisionId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.RevisionId(query_revisionId)
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

func Source_v1_ProjectsReposAliasesFilesGet(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("repoName"))
		pathParams = append(pathParams, commands_util.AngrySnakes("kind"))
		pathParams = append(pathParams, commands_util.AngrySnakes("name"))
		pathParams = append(pathParams, commands_util.AngrySnakes("path"))

		if len(pathParams) != 0 {
			if strings.Contains("v1/projects/{projectId}/repos/{repoName}/aliases/{kind}/{name}/files/{+path}", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [--pageSize=VALUE]"

		usageBits += " [--pageToken=VALUE]"

		usageBits += " [--sourceContext.cloudRepo.aliasName=VALUE]"

		usageBits += " [--sourceContext.cloudRepo.repoId.uid=VALUE]"

		usageBits += " [--sourceContext.cloudRepo.revisionId=VALUE]"

		usageBits += " [--sourceContext.cloudWorkspace.snapshotId=VALUE]"

		usageBits += " [--sourceContext.cloudWorkspace.workspaceId.name=VALUE]"

		usageBits += " [--sourceContext.cloudWorkspace.workspaceId.repoId.projectRepoId.projectId=VALUE]"

		usageBits += " [--sourceContext.cloudWorkspace.workspaceId.repoId.projectRepoId.repoName=VALUE]"

		usageBits += " [--sourceContext.cloudWorkspace.workspaceId.repoId.uid=VALUE]"

		usageBits += " [--sourceContext.gerrit.aliasContext.kind=VALUE]"

		usageBits += " [--sourceContext.gerrit.aliasContext.name=VALUE]"

		usageBits += " [--sourceContext.gerrit.aliasName=VALUE]"

		usageBits += " [--sourceContext.gerrit.gerritProject=VALUE]"

		usageBits += " [--sourceContext.gerrit.hostUri=VALUE]"

		usageBits += " [--sourceContext.gerrit.revisionId=VALUE]"

		usageBits += " [--sourceContext.git.revisionId=VALUE]"

		usageBits += " [--sourceContext.git.url=VALUE]"

		usageBits += " [--startPosition=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsReposAliasesFilesService(api_service)

	queryParamNames := map[string]bool{
		"pageSize":                                                                false,
		"pageToken":                                                               false,
		"sourceContext.cloudRepo.aliasName":                                       false,
		"sourceContext.cloudRepo.repoId.uid":                                      false,
		"sourceContext.cloudRepo.revisionId":                                      false,
		"sourceContext.cloudWorkspace.snapshotId":                                 false,
		"sourceContext.cloudWorkspace.workspaceId.name":                           false,
		"sourceContext.cloudWorkspace.workspaceId.repoId.projectRepoId.projectId": false,
		"sourceContext.cloudWorkspace.workspaceId.repoId.projectRepoId.repoName":  false,
		"sourceContext.cloudWorkspace.workspaceId.repoId.uid":                     false,
		"sourceContext.gerrit.aliasContext.kind":                                  false,
		"sourceContext.gerrit.aliasContext.name":                                  false,
		"sourceContext.gerrit.aliasName":                                          false,
		"sourceContext.gerrit.gerritProject":                                      false,
		"sourceContext.gerrit.hostUri":                                            false,
		"sourceContext.gerrit.revisionId":                                         false,
		"sourceContext.git.revisionId":                                            false,
		"sourceContext.git.url":                                                   false,
		"startPosition":                                                           false,
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
	if value, ok := flagValues["sourceContext_cloudRepo_aliasName"]; ok {
		query_sourceContext_cloudRepo_aliasName, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextCloudRepoAliasName(query_sourceContext_cloudRepo_aliasName)
	}
	if value, ok := flagValues["sourceContext_cloudRepo_repoId_uid"]; ok {
		query_sourceContext_cloudRepo_repoId_uid, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextCloudRepoRepoIdUid(query_sourceContext_cloudRepo_repoId_uid)
	}
	if value, ok := flagValues["sourceContext_cloudRepo_revisionId"]; ok {
		query_sourceContext_cloudRepo_revisionId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextCloudRepoRevisionId(query_sourceContext_cloudRepo_revisionId)
	}
	if value, ok := flagValues["sourceContext_cloudWorkspace_snapshotId"]; ok {
		query_sourceContext_cloudWorkspace_snapshotId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextCloudWorkspaceSnapshotId(query_sourceContext_cloudWorkspace_snapshotId)
	}
	if value, ok := flagValues["sourceContext_cloudWorkspace_workspaceId_name"]; ok {
		query_sourceContext_cloudWorkspace_workspaceId_name, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextCloudWorkspaceWorkspaceIdName(query_sourceContext_cloudWorkspace_workspaceId_name)
	}
	if value, ok := flagValues["sourceContext_cloudWorkspace_workspaceId_repoId_projectRepoId_projectId"]; ok {
		query_sourceContext_cloudWorkspace_workspaceId_repoId_projectRepoId_projectId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextCloudWorkspaceWorkspaceIdRepoIdProjectRepoIdProjectId(query_sourceContext_cloudWorkspace_workspaceId_repoId_projectRepoId_projectId)
	}
	if value, ok := flagValues["sourceContext_cloudWorkspace_workspaceId_repoId_projectRepoId_repoName"]; ok {
		query_sourceContext_cloudWorkspace_workspaceId_repoId_projectRepoId_repoName, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextCloudWorkspaceWorkspaceIdRepoIdProjectRepoIdRepoName(query_sourceContext_cloudWorkspace_workspaceId_repoId_projectRepoId_repoName)
	}
	if value, ok := flagValues["sourceContext_cloudWorkspace_workspaceId_repoId_uid"]; ok {
		query_sourceContext_cloudWorkspace_workspaceId_repoId_uid, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextCloudWorkspaceWorkspaceIdRepoIdUid(query_sourceContext_cloudWorkspace_workspaceId_repoId_uid)
	}
	if value, ok := flagValues["sourceContext_gerrit_aliasContext_kind"]; ok {
		query_sourceContext_gerrit_aliasContext_kind, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextGerritAliasContextKind(query_sourceContext_gerrit_aliasContext_kind)
	}
	if value, ok := flagValues["sourceContext_gerrit_aliasContext_name"]; ok {
		query_sourceContext_gerrit_aliasContext_name, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextGerritAliasContextName(query_sourceContext_gerrit_aliasContext_name)
	}
	if value, ok := flagValues["sourceContext_gerrit_aliasName"]; ok {
		query_sourceContext_gerrit_aliasName, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextGerritAliasName(query_sourceContext_gerrit_aliasName)
	}
	if value, ok := flagValues["sourceContext_gerrit_gerritProject"]; ok {
		query_sourceContext_gerrit_gerritProject, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextGerritGerritProject(query_sourceContext_gerrit_gerritProject)
	}
	if value, ok := flagValues["sourceContext_gerrit_hostUri"]; ok {
		query_sourceContext_gerrit_hostUri, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextGerritHostUri(query_sourceContext_gerrit_hostUri)
	}
	if value, ok := flagValues["sourceContext_gerrit_revisionId"]; ok {
		query_sourceContext_gerrit_revisionId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextGerritRevisionId(query_sourceContext_gerrit_revisionId)
	}
	if value, ok := flagValues["sourceContext_git_revisionId"]; ok {
		query_sourceContext_git_revisionId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextGitRevisionId(query_sourceContext_git_revisionId)
	}
	if value, ok := flagValues["sourceContext_git_url"]; ok {
		query_sourceContext_git_url, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextGitUrl(query_sourceContext_git_url)
	}
	if value, ok := flagValues["startPosition"]; ok {
		query_startPosition, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.StartPosition(query_startPosition)
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

func Source_v1_ProjectsReposAliasesGet(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("repoName"))
		pathParams = append(pathParams, commands_util.AngrySnakes("kind"))
		pathParams = append(pathParams, commands_util.AngrySnakes("name"))

		if len(pathParams) != 0 {
			if strings.Contains("v1/projects/{projectId}/repos/{repoName}/aliases/{kind}/{name}", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [--repoId.uid=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsReposAliasesService(api_service)

	queryParamNames := map[string]bool{
		"repoId.uid": false,
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
	if value, ok := flagValues["repoId_uid"]; ok {
		query_repoId_uid, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.RepoIdUid(query_repoId_uid)
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

func Source_v1_ProjectsReposAliasesList(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("repoName"))

		if len(pathParams) != 0 {
			if strings.Contains("v1/projects/{projectId}/repos/{repoName}/aliases", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [--kind=VALUE]"

		usageBits += " [--pageSize=VALUE]"

		usageBits += " [--pageToken=VALUE]"

		usageBits += " [--repoId.uid=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsReposAliasesService(api_service)

	queryParamNames := map[string]bool{
		"kind":       false,
		"pageSize":   false,
		"pageToken":  false,
		"repoId.uid": false,
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
	if value, ok := flagValues["repoId_uid"]; ok {
		query_repoId_uid, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.RepoIdUid(query_repoId_uid)
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

func Source_v1_ProjectsReposAliasesListFiles(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("repoName"))
		pathParams = append(pathParams, commands_util.AngrySnakes("kind"))
		pathParams = append(pathParams, commands_util.AngrySnakes("name"))

		if len(pathParams) != 0 {
			if strings.Contains("v1/projects/{projectId}/repos/{repoName}/aliases/{kind}/{name}:listFiles", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [--pageSize=VALUE]"

		usageBits += " [--pageToken=VALUE]"

		usageBits += " [--sourceContext.cloudRepo.aliasName=VALUE]"

		usageBits += " [--sourceContext.cloudRepo.repoId.uid=VALUE]"

		usageBits += " [--sourceContext.cloudRepo.revisionId=VALUE]"

		usageBits += " [--sourceContext.cloudWorkspace.snapshotId=VALUE]"

		usageBits += " [--sourceContext.cloudWorkspace.workspaceId.name=VALUE]"

		usageBits += " [--sourceContext.cloudWorkspace.workspaceId.repoId.projectRepoId.projectId=VALUE]"

		usageBits += " [--sourceContext.cloudWorkspace.workspaceId.repoId.projectRepoId.repoName=VALUE]"

		usageBits += " [--sourceContext.cloudWorkspace.workspaceId.repoId.uid=VALUE]"

		usageBits += " [--sourceContext.gerrit.aliasContext.kind=VALUE]"

		usageBits += " [--sourceContext.gerrit.aliasContext.name=VALUE]"

		usageBits += " [--sourceContext.gerrit.aliasName=VALUE]"

		usageBits += " [--sourceContext.gerrit.gerritProject=VALUE]"

		usageBits += " [--sourceContext.gerrit.hostUri=VALUE]"

		usageBits += " [--sourceContext.gerrit.revisionId=VALUE]"

		usageBits += " [--sourceContext.git.revisionId=VALUE]"

		usageBits += " [--sourceContext.git.url=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsReposAliasesService(api_service)

	queryParamNames := map[string]bool{
		"pageSize":                                                                false,
		"pageToken":                                                               false,
		"sourceContext.cloudRepo.aliasName":                                       false,
		"sourceContext.cloudRepo.repoId.uid":                                      false,
		"sourceContext.cloudRepo.revisionId":                                      false,
		"sourceContext.cloudWorkspace.snapshotId":                                 false,
		"sourceContext.cloudWorkspace.workspaceId.name":                           false,
		"sourceContext.cloudWorkspace.workspaceId.repoId.projectRepoId.projectId": false,
		"sourceContext.cloudWorkspace.workspaceId.repoId.projectRepoId.repoName":  false,
		"sourceContext.cloudWorkspace.workspaceId.repoId.uid":                     false,
		"sourceContext.gerrit.aliasContext.kind":                                  false,
		"sourceContext.gerrit.aliasContext.name":                                  false,
		"sourceContext.gerrit.aliasName":                                          false,
		"sourceContext.gerrit.gerritProject":                                      false,
		"sourceContext.gerrit.hostUri":                                            false,
		"sourceContext.gerrit.revisionId":                                         false,
		"sourceContext.git.revisionId":                                            false,
		"sourceContext.git.url":                                                   false,
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
	if value, ok := flagValues["sourceContext_cloudRepo_aliasName"]; ok {
		query_sourceContext_cloudRepo_aliasName, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextCloudRepoAliasName(query_sourceContext_cloudRepo_aliasName)
	}
	if value, ok := flagValues["sourceContext_cloudRepo_repoId_uid"]; ok {
		query_sourceContext_cloudRepo_repoId_uid, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextCloudRepoRepoIdUid(query_sourceContext_cloudRepo_repoId_uid)
	}
	if value, ok := flagValues["sourceContext_cloudRepo_revisionId"]; ok {
		query_sourceContext_cloudRepo_revisionId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextCloudRepoRevisionId(query_sourceContext_cloudRepo_revisionId)
	}
	if value, ok := flagValues["sourceContext_cloudWorkspace_snapshotId"]; ok {
		query_sourceContext_cloudWorkspace_snapshotId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextCloudWorkspaceSnapshotId(query_sourceContext_cloudWorkspace_snapshotId)
	}
	if value, ok := flagValues["sourceContext_cloudWorkspace_workspaceId_name"]; ok {
		query_sourceContext_cloudWorkspace_workspaceId_name, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextCloudWorkspaceWorkspaceIdName(query_sourceContext_cloudWorkspace_workspaceId_name)
	}
	if value, ok := flagValues["sourceContext_cloudWorkspace_workspaceId_repoId_projectRepoId_projectId"]; ok {
		query_sourceContext_cloudWorkspace_workspaceId_repoId_projectRepoId_projectId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextCloudWorkspaceWorkspaceIdRepoIdProjectRepoIdProjectId(query_sourceContext_cloudWorkspace_workspaceId_repoId_projectRepoId_projectId)
	}
	if value, ok := flagValues["sourceContext_cloudWorkspace_workspaceId_repoId_projectRepoId_repoName"]; ok {
		query_sourceContext_cloudWorkspace_workspaceId_repoId_projectRepoId_repoName, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextCloudWorkspaceWorkspaceIdRepoIdProjectRepoIdRepoName(query_sourceContext_cloudWorkspace_workspaceId_repoId_projectRepoId_repoName)
	}
	if value, ok := flagValues["sourceContext_cloudWorkspace_workspaceId_repoId_uid"]; ok {
		query_sourceContext_cloudWorkspace_workspaceId_repoId_uid, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextCloudWorkspaceWorkspaceIdRepoIdUid(query_sourceContext_cloudWorkspace_workspaceId_repoId_uid)
	}
	if value, ok := flagValues["sourceContext_gerrit_aliasContext_kind"]; ok {
		query_sourceContext_gerrit_aliasContext_kind, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextGerritAliasContextKind(query_sourceContext_gerrit_aliasContext_kind)
	}
	if value, ok := flagValues["sourceContext_gerrit_aliasContext_name"]; ok {
		query_sourceContext_gerrit_aliasContext_name, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextGerritAliasContextName(query_sourceContext_gerrit_aliasContext_name)
	}
	if value, ok := flagValues["sourceContext_gerrit_aliasName"]; ok {
		query_sourceContext_gerrit_aliasName, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextGerritAliasName(query_sourceContext_gerrit_aliasName)
	}
	if value, ok := flagValues["sourceContext_gerrit_gerritProject"]; ok {
		query_sourceContext_gerrit_gerritProject, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextGerritGerritProject(query_sourceContext_gerrit_gerritProject)
	}
	if value, ok := flagValues["sourceContext_gerrit_hostUri"]; ok {
		query_sourceContext_gerrit_hostUri, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextGerritHostUri(query_sourceContext_gerrit_hostUri)
	}
	if value, ok := flagValues["sourceContext_gerrit_revisionId"]; ok {
		query_sourceContext_gerrit_revisionId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextGerritRevisionId(query_sourceContext_gerrit_revisionId)
	}
	if value, ok := flagValues["sourceContext_git_revisionId"]; ok {
		query_sourceContext_git_revisionId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextGitRevisionId(query_sourceContext_git_revisionId)
	}
	if value, ok := flagValues["sourceContext_git_url"]; ok {
		query_sourceContext_git_url, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextGitUrl(query_sourceContext_git_url)
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

func Source_v1_ProjectsReposAliasesUpdate(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("repoName"))
		pathParams = append(pathParams, commands_util.AngrySnakes("aliasesId"))

		if len(pathParams) != 0 {
			if strings.Contains("v1/projects/{projectId}/repos/{repoName}/aliases/{aliasesId}", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		usageBits += " [--oldRevisionId=VALUE]"

		usageBits += " [--repoId.uid=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		commands_util.PrintRequestExample(&api_client.Alias{})

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsReposAliasesService(api_service)

	queryParamNames := map[string]bool{
		"oldRevisionId": false,
		"repoId.uid":    false,
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
	if value, ok := flagValues["repoId_uid"]; ok {
		query_repoId_uid, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.RepoIdUid(query_repoId_uid)
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

func Source_v1_ProjectsReposCreate(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectId"))

		if len(pathParams) != 0 {
			if strings.Contains("v1/projects/{projectId}/repos", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		commands_util.PrintRequestExample(&api_client.Repo{})

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

func Source_v1_ProjectsReposDelete(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("repoName"))

		if len(pathParams) != 0 {
			if strings.Contains("v1/projects/{projectId}/repos/{repoName}", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [--repoId.uid=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsReposService(api_service)

	queryParamNames := map[string]bool{
		"repoId.uid": false,
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
	if value, ok := flagValues["repoId_uid"]; ok {
		query_repoId_uid, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.RepoIdUid(query_repoId_uid)
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

func Source_v1_ProjectsReposFilesReadFromWorkspaceOrAlias(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("repoName"))
		pathParams = append(pathParams, commands_util.AngrySnakes("path"))

		if len(pathParams) != 0 {
			if strings.Contains("v1/projects/{projectId}/repos/{repoName}/files/{+path}:readFromWorkspaceOrAlias", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [--alias=VALUE]"

		usageBits += " [--pageSize=VALUE]"

		usageBits += " [--pageToken=VALUE]"

		usageBits += " [--repoId.uid=VALUE]"

		usageBits += " [--startPosition=VALUE]"

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
		"repoId.uid":    false,
		"startPosition": false,
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
	if value, ok := flagValues["repoId_uid"]; ok {
		query_repoId_uid, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.RepoIdUid(query_repoId_uid)
	}
	if value, ok := flagValues["startPosition"]; ok {
		query_startPosition, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.StartPosition(query_startPosition)
	}
	if value, ok := flagValues["workspaceName"]; ok {
		query_workspaceName, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.WorkspaceName(query_workspaceName)
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

func Source_v1_ProjectsReposGet(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("repoName"))

		if len(pathParams) != 0 {
			if strings.Contains("v1/projects/{projectId}/repos/{repoName}", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [--repoId.uid=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsReposService(api_service)

	queryParamNames := map[string]bool{
		"repoId.uid": false,
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
	if value, ok := flagValues["repoId_uid"]; ok {
		query_repoId_uid, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.RepoIdUid(query_repoId_uid)
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

func Source_v1_ProjectsReposList(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectId"))

		if len(pathParams) != 0 {
			if strings.Contains("v1/projects/{projectId}/repos", "+") {
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

func Source_v1_ProjectsReposMerge(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("repoName"))

		if len(pathParams) != 0 {
			if strings.Contains("v1/projects/{projectId}/repos/{repoName}:merge", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		commands_util.PrintRequestExample(&api_client.MergeRequest{})

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

func Source_v1_ProjectsReposRevisionsFilesGet(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("repoName"))
		pathParams = append(pathParams, commands_util.AngrySnakes("revisionId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("path"))

		if len(pathParams) != 0 {
			if strings.Contains("v1/projects/{projectId}/repos/{repoName}/revisions/{revisionId}/files/{+path}", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [--pageSize=VALUE]"

		usageBits += " [--pageToken=VALUE]"

		usageBits += " [--sourceContext.cloudRepo.aliasContext.kind=VALUE]"

		usageBits += " [--sourceContext.cloudRepo.aliasContext.name=VALUE]"

		usageBits += " [--sourceContext.cloudRepo.aliasName=VALUE]"

		usageBits += " [--sourceContext.cloudRepo.repoId.uid=VALUE]"

		usageBits += " [--sourceContext.cloudWorkspace.snapshotId=VALUE]"

		usageBits += " [--sourceContext.cloudWorkspace.workspaceId.name=VALUE]"

		usageBits += " [--sourceContext.cloudWorkspace.workspaceId.repoId.projectRepoId.projectId=VALUE]"

		usageBits += " [--sourceContext.cloudWorkspace.workspaceId.repoId.projectRepoId.repoName=VALUE]"

		usageBits += " [--sourceContext.cloudWorkspace.workspaceId.repoId.uid=VALUE]"

		usageBits += " [--sourceContext.gerrit.aliasContext.kind=VALUE]"

		usageBits += " [--sourceContext.gerrit.aliasContext.name=VALUE]"

		usageBits += " [--sourceContext.gerrit.aliasName=VALUE]"

		usageBits += " [--sourceContext.gerrit.gerritProject=VALUE]"

		usageBits += " [--sourceContext.gerrit.hostUri=VALUE]"

		usageBits += " [--sourceContext.gerrit.revisionId=VALUE]"

		usageBits += " [--sourceContext.git.revisionId=VALUE]"

		usageBits += " [--sourceContext.git.url=VALUE]"

		usageBits += " [--startPosition=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsReposRevisionsFilesService(api_service)

	queryParamNames := map[string]bool{
		"pageSize":                                                                false,
		"pageToken":                                                               false,
		"sourceContext.cloudRepo.aliasContext.kind":                               false,
		"sourceContext.cloudRepo.aliasContext.name":                               false,
		"sourceContext.cloudRepo.aliasName":                                       false,
		"sourceContext.cloudRepo.repoId.uid":                                      false,
		"sourceContext.cloudWorkspace.snapshotId":                                 false,
		"sourceContext.cloudWorkspace.workspaceId.name":                           false,
		"sourceContext.cloudWorkspace.workspaceId.repoId.projectRepoId.projectId": false,
		"sourceContext.cloudWorkspace.workspaceId.repoId.projectRepoId.repoName":  false,
		"sourceContext.cloudWorkspace.workspaceId.repoId.uid":                     false,
		"sourceContext.gerrit.aliasContext.kind":                                  false,
		"sourceContext.gerrit.aliasContext.name":                                  false,
		"sourceContext.gerrit.aliasName":                                          false,
		"sourceContext.gerrit.gerritProject":                                      false,
		"sourceContext.gerrit.hostUri":                                            false,
		"sourceContext.gerrit.revisionId":                                         false,
		"sourceContext.git.revisionId":                                            false,
		"sourceContext.git.url":                                                   false,
		"startPosition":                                                           false,
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
	if value, ok := flagValues["sourceContext_cloudRepo_aliasContext_kind"]; ok {
		query_sourceContext_cloudRepo_aliasContext_kind, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextCloudRepoAliasContextKind(query_sourceContext_cloudRepo_aliasContext_kind)
	}
	if value, ok := flagValues["sourceContext_cloudRepo_aliasContext_name"]; ok {
		query_sourceContext_cloudRepo_aliasContext_name, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextCloudRepoAliasContextName(query_sourceContext_cloudRepo_aliasContext_name)
	}
	if value, ok := flagValues["sourceContext_cloudRepo_aliasName"]; ok {
		query_sourceContext_cloudRepo_aliasName, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextCloudRepoAliasName(query_sourceContext_cloudRepo_aliasName)
	}
	if value, ok := flagValues["sourceContext_cloudRepo_repoId_uid"]; ok {
		query_sourceContext_cloudRepo_repoId_uid, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextCloudRepoRepoIdUid(query_sourceContext_cloudRepo_repoId_uid)
	}
	if value, ok := flagValues["sourceContext_cloudWorkspace_snapshotId"]; ok {
		query_sourceContext_cloudWorkspace_snapshotId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextCloudWorkspaceSnapshotId(query_sourceContext_cloudWorkspace_snapshotId)
	}
	if value, ok := flagValues["sourceContext_cloudWorkspace_workspaceId_name"]; ok {
		query_sourceContext_cloudWorkspace_workspaceId_name, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextCloudWorkspaceWorkspaceIdName(query_sourceContext_cloudWorkspace_workspaceId_name)
	}
	if value, ok := flagValues["sourceContext_cloudWorkspace_workspaceId_repoId_projectRepoId_projectId"]; ok {
		query_sourceContext_cloudWorkspace_workspaceId_repoId_projectRepoId_projectId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextCloudWorkspaceWorkspaceIdRepoIdProjectRepoIdProjectId(query_sourceContext_cloudWorkspace_workspaceId_repoId_projectRepoId_projectId)
	}
	if value, ok := flagValues["sourceContext_cloudWorkspace_workspaceId_repoId_projectRepoId_repoName"]; ok {
		query_sourceContext_cloudWorkspace_workspaceId_repoId_projectRepoId_repoName, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextCloudWorkspaceWorkspaceIdRepoIdProjectRepoIdRepoName(query_sourceContext_cloudWorkspace_workspaceId_repoId_projectRepoId_repoName)
	}
	if value, ok := flagValues["sourceContext_cloudWorkspace_workspaceId_repoId_uid"]; ok {
		query_sourceContext_cloudWorkspace_workspaceId_repoId_uid, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextCloudWorkspaceWorkspaceIdRepoIdUid(query_sourceContext_cloudWorkspace_workspaceId_repoId_uid)
	}
	if value, ok := flagValues["sourceContext_gerrit_aliasContext_kind"]; ok {
		query_sourceContext_gerrit_aliasContext_kind, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextGerritAliasContextKind(query_sourceContext_gerrit_aliasContext_kind)
	}
	if value, ok := flagValues["sourceContext_gerrit_aliasContext_name"]; ok {
		query_sourceContext_gerrit_aliasContext_name, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextGerritAliasContextName(query_sourceContext_gerrit_aliasContext_name)
	}
	if value, ok := flagValues["sourceContext_gerrit_aliasName"]; ok {
		query_sourceContext_gerrit_aliasName, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextGerritAliasName(query_sourceContext_gerrit_aliasName)
	}
	if value, ok := flagValues["sourceContext_gerrit_gerritProject"]; ok {
		query_sourceContext_gerrit_gerritProject, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextGerritGerritProject(query_sourceContext_gerrit_gerritProject)
	}
	if value, ok := flagValues["sourceContext_gerrit_hostUri"]; ok {
		query_sourceContext_gerrit_hostUri, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextGerritHostUri(query_sourceContext_gerrit_hostUri)
	}
	if value, ok := flagValues["sourceContext_gerrit_revisionId"]; ok {
		query_sourceContext_gerrit_revisionId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextGerritRevisionId(query_sourceContext_gerrit_revisionId)
	}
	if value, ok := flagValues["sourceContext_git_revisionId"]; ok {
		query_sourceContext_git_revisionId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextGitRevisionId(query_sourceContext_git_revisionId)
	}
	if value, ok := flagValues["sourceContext_git_url"]; ok {
		query_sourceContext_git_url, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextGitUrl(query_sourceContext_git_url)
	}
	if value, ok := flagValues["startPosition"]; ok {
		query_startPosition, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.StartPosition(query_startPosition)
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

func Source_v1_ProjectsReposRevisionsGet(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("repoName"))
		pathParams = append(pathParams, commands_util.AngrySnakes("revisionId"))

		if len(pathParams) != 0 {
			if strings.Contains("v1/projects/{projectId}/repos/{repoName}/revisions/{revisionId}", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [--repoId.uid=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsReposRevisionsService(api_service)

	queryParamNames := map[string]bool{
		"repoId.uid": false,
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
	if value, ok := flagValues["repoId_uid"]; ok {
		query_repoId_uid, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.RepoIdUid(query_repoId_uid)
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

func Source_v1_ProjectsReposRevisionsGetBatchGet(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("repoName"))

		if len(pathParams) != 0 {
			if strings.Contains("v1/projects/{projectId}/repos/{repoName}/revisions:batchGet", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [--repoId.uid=VALUE]"

		usageBits += " [--revisionIds=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsReposRevisionsService(api_service)

	queryParamNames := map[string]bool{
		"repoId.uid":  false,
		"revisionIds": false,
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
	if value, ok := flagValues["repoId_uid"]; ok {
		query_repoId_uid, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.RepoIdUid(query_repoId_uid)
	}
	if value, ok := flagValues["revisionIds"]; ok {
		query_revisionIds, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.RevisionIds(query_revisionIds)
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

func Source_v1_ProjectsReposRevisionsList(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("repoName"))

		if len(pathParams) != 0 {
			if strings.Contains("v1/projects/{projectId}/repos/{repoName}/revisions", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [--ends=VALUE]"

		usageBits += " [--pageSize=VALUE]"

		usageBits += " [--pageToken=VALUE]"

		usageBits += " [--path=VALUE]"

		usageBits += " [--repoId.uid=VALUE]"

		usageBits += " [--starts=VALUE]"

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
		"repoId.uid":    false,
		"starts":        false,
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
	if value, ok := flagValues["repoId_uid"]; ok {
		query_repoId_uid, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.RepoIdUid(query_repoId_uid)
	}
	if value, ok := flagValues["starts"]; ok {
		query_starts, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.Starts(query_starts)
	}
	if value, ok := flagValues["walkDirection"]; ok {
		query_walkDirection, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.WalkDirection(query_walkDirection)
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

func Source_v1_ProjectsReposRevisionsListFiles(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("repoName"))
		pathParams = append(pathParams, commands_util.AngrySnakes("revisionId"))

		if len(pathParams) != 0 {
			if strings.Contains("v1/projects/{projectId}/repos/{repoName}/revisions/{revisionId}:listFiles", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [--pageSize=VALUE]"

		usageBits += " [--pageToken=VALUE]"

		usageBits += " [--sourceContext.cloudRepo.aliasContext.kind=VALUE]"

		usageBits += " [--sourceContext.cloudRepo.aliasContext.name=VALUE]"

		usageBits += " [--sourceContext.cloudRepo.aliasName=VALUE]"

		usageBits += " [--sourceContext.cloudRepo.repoId.uid=VALUE]"

		usageBits += " [--sourceContext.cloudWorkspace.snapshotId=VALUE]"

		usageBits += " [--sourceContext.cloudWorkspace.workspaceId.name=VALUE]"

		usageBits += " [--sourceContext.cloudWorkspace.workspaceId.repoId.projectRepoId.projectId=VALUE]"

		usageBits += " [--sourceContext.cloudWorkspace.workspaceId.repoId.projectRepoId.repoName=VALUE]"

		usageBits += " [--sourceContext.cloudWorkspace.workspaceId.repoId.uid=VALUE]"

		usageBits += " [--sourceContext.gerrit.aliasContext.kind=VALUE]"

		usageBits += " [--sourceContext.gerrit.aliasContext.name=VALUE]"

		usageBits += " [--sourceContext.gerrit.aliasName=VALUE]"

		usageBits += " [--sourceContext.gerrit.gerritProject=VALUE]"

		usageBits += " [--sourceContext.gerrit.hostUri=VALUE]"

		usageBits += " [--sourceContext.gerrit.revisionId=VALUE]"

		usageBits += " [--sourceContext.git.revisionId=VALUE]"

		usageBits += " [--sourceContext.git.url=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsReposRevisionsService(api_service)

	queryParamNames := map[string]bool{
		"pageSize":                                                                false,
		"pageToken":                                                               false,
		"sourceContext.cloudRepo.aliasContext.kind":                               false,
		"sourceContext.cloudRepo.aliasContext.name":                               false,
		"sourceContext.cloudRepo.aliasName":                                       false,
		"sourceContext.cloudRepo.repoId.uid":                                      false,
		"sourceContext.cloudWorkspace.snapshotId":                                 false,
		"sourceContext.cloudWorkspace.workspaceId.name":                           false,
		"sourceContext.cloudWorkspace.workspaceId.repoId.projectRepoId.projectId": false,
		"sourceContext.cloudWorkspace.workspaceId.repoId.projectRepoId.repoName":  false,
		"sourceContext.cloudWorkspace.workspaceId.repoId.uid":                     false,
		"sourceContext.gerrit.aliasContext.kind":                                  false,
		"sourceContext.gerrit.aliasContext.name":                                  false,
		"sourceContext.gerrit.aliasName":                                          false,
		"sourceContext.gerrit.gerritProject":                                      false,
		"sourceContext.gerrit.hostUri":                                            false,
		"sourceContext.gerrit.revisionId":                                         false,
		"sourceContext.git.revisionId":                                            false,
		"sourceContext.git.url":                                                   false,
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
	if value, ok := flagValues["sourceContext_cloudRepo_aliasContext_kind"]; ok {
		query_sourceContext_cloudRepo_aliasContext_kind, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextCloudRepoAliasContextKind(query_sourceContext_cloudRepo_aliasContext_kind)
	}
	if value, ok := flagValues["sourceContext_cloudRepo_aliasContext_name"]; ok {
		query_sourceContext_cloudRepo_aliasContext_name, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextCloudRepoAliasContextName(query_sourceContext_cloudRepo_aliasContext_name)
	}
	if value, ok := flagValues["sourceContext_cloudRepo_aliasName"]; ok {
		query_sourceContext_cloudRepo_aliasName, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextCloudRepoAliasName(query_sourceContext_cloudRepo_aliasName)
	}
	if value, ok := flagValues["sourceContext_cloudRepo_repoId_uid"]; ok {
		query_sourceContext_cloudRepo_repoId_uid, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextCloudRepoRepoIdUid(query_sourceContext_cloudRepo_repoId_uid)
	}
	if value, ok := flagValues["sourceContext_cloudWorkspace_snapshotId"]; ok {
		query_sourceContext_cloudWorkspace_snapshotId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextCloudWorkspaceSnapshotId(query_sourceContext_cloudWorkspace_snapshotId)
	}
	if value, ok := flagValues["sourceContext_cloudWorkspace_workspaceId_name"]; ok {
		query_sourceContext_cloudWorkspace_workspaceId_name, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextCloudWorkspaceWorkspaceIdName(query_sourceContext_cloudWorkspace_workspaceId_name)
	}
	if value, ok := flagValues["sourceContext_cloudWorkspace_workspaceId_repoId_projectRepoId_projectId"]; ok {
		query_sourceContext_cloudWorkspace_workspaceId_repoId_projectRepoId_projectId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextCloudWorkspaceWorkspaceIdRepoIdProjectRepoIdProjectId(query_sourceContext_cloudWorkspace_workspaceId_repoId_projectRepoId_projectId)
	}
	if value, ok := flagValues["sourceContext_cloudWorkspace_workspaceId_repoId_projectRepoId_repoName"]; ok {
		query_sourceContext_cloudWorkspace_workspaceId_repoId_projectRepoId_repoName, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextCloudWorkspaceWorkspaceIdRepoIdProjectRepoIdRepoName(query_sourceContext_cloudWorkspace_workspaceId_repoId_projectRepoId_repoName)
	}
	if value, ok := flagValues["sourceContext_cloudWorkspace_workspaceId_repoId_uid"]; ok {
		query_sourceContext_cloudWorkspace_workspaceId_repoId_uid, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextCloudWorkspaceWorkspaceIdRepoIdUid(query_sourceContext_cloudWorkspace_workspaceId_repoId_uid)
	}
	if value, ok := flagValues["sourceContext_gerrit_aliasContext_kind"]; ok {
		query_sourceContext_gerrit_aliasContext_kind, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextGerritAliasContextKind(query_sourceContext_gerrit_aliasContext_kind)
	}
	if value, ok := flagValues["sourceContext_gerrit_aliasContext_name"]; ok {
		query_sourceContext_gerrit_aliasContext_name, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextGerritAliasContextName(query_sourceContext_gerrit_aliasContext_name)
	}
	if value, ok := flagValues["sourceContext_gerrit_aliasName"]; ok {
		query_sourceContext_gerrit_aliasName, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextGerritAliasName(query_sourceContext_gerrit_aliasName)
	}
	if value, ok := flagValues["sourceContext_gerrit_gerritProject"]; ok {
		query_sourceContext_gerrit_gerritProject, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextGerritGerritProject(query_sourceContext_gerrit_gerritProject)
	}
	if value, ok := flagValues["sourceContext_gerrit_hostUri"]; ok {
		query_sourceContext_gerrit_hostUri, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextGerritHostUri(query_sourceContext_gerrit_hostUri)
	}
	if value, ok := flagValues["sourceContext_gerrit_revisionId"]; ok {
		query_sourceContext_gerrit_revisionId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextGerritRevisionId(query_sourceContext_gerrit_revisionId)
	}
	if value, ok := flagValues["sourceContext_git_revisionId"]; ok {
		query_sourceContext_git_revisionId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextGitRevisionId(query_sourceContext_git_revisionId)
	}
	if value, ok := flagValues["sourceContext_git_url"]; ok {
		query_sourceContext_git_url, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextGitUrl(query_sourceContext_git_url)
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

func Source_v1_ProjectsReposUpdate(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("repoName"))

		if len(pathParams) != 0 {
			if strings.Contains("v1/projects/{projectId}/repos/{repoName}", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		commands_util.PrintRequestExample(&api_client.UpdateRepoRequest{})

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

	request := &api_client.UpdateRepoRequest{}
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

	call := service.Update(param_projectId, param_repoName,
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

func Source_v1_ProjectsReposWorkspacesCommitWorkspace(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("repoName"))
		pathParams = append(pathParams, commands_util.AngrySnakes("name"))

		if len(pathParams) != 0 {
			if strings.Contains("v1/projects/{projectId}/repos/{repoName}/workspaces/{name}:commitWorkspace", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		commands_util.PrintRequestExample(&api_client.CommitWorkspaceRequest{})

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

func Source_v1_ProjectsReposWorkspacesCreate(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("repoName"))

		if len(pathParams) != 0 {
			if strings.Contains("v1/projects/{projectId}/repos/{repoName}/workspaces", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		commands_util.PrintRequestExample(&api_client.CreateWorkspaceRequest{})

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

func Source_v1_ProjectsReposWorkspacesDelete(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("repoName"))
		pathParams = append(pathParams, commands_util.AngrySnakes("name"))

		if len(pathParams) != 0 {
			if strings.Contains("v1/projects/{projectId}/repos/{repoName}/workspaces/{name}", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [--currentSnapshotId=VALUE]"

		usageBits += " [--workspaceId.repoId.uid=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsReposWorkspacesService(api_service)

	queryParamNames := map[string]bool{
		"currentSnapshotId":      false,
		"workspaceId.repoId.uid": false,
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
	if value, ok := flagValues["workspaceId_repoId_uid"]; ok {
		query_workspaceId_repoId_uid, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.WorkspaceIdRepoIdUid(query_workspaceId_repoId_uid)
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

func Source_v1_ProjectsReposWorkspacesFilesGet(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("repoName"))
		pathParams = append(pathParams, commands_util.AngrySnakes("name"))
		pathParams = append(pathParams, commands_util.AngrySnakes("path"))

		if len(pathParams) != 0 {
			if strings.Contains("v1/projects/{projectId}/repos/{repoName}/workspaces/{name}/files/{+path}", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [--pageSize=VALUE]"

		usageBits += " [--pageToken=VALUE]"

		usageBits += " [--sourceContext.cloudRepo.aliasContext.kind=VALUE]"

		usageBits += " [--sourceContext.cloudRepo.aliasContext.name=VALUE]"

		usageBits += " [--sourceContext.cloudRepo.aliasName=VALUE]"

		usageBits += " [--sourceContext.cloudRepo.repoId.projectRepoId.projectId=VALUE]"

		usageBits += " [--sourceContext.cloudRepo.repoId.projectRepoId.repoName=VALUE]"

		usageBits += " [--sourceContext.cloudRepo.repoId.uid=VALUE]"

		usageBits += " [--sourceContext.cloudRepo.revisionId=VALUE]"

		usageBits += " [--sourceContext.cloudWorkspace.snapshotId=VALUE]"

		usageBits += " [--sourceContext.cloudWorkspace.workspaceId.repoId.uid=VALUE]"

		usageBits += " [--sourceContext.gerrit.aliasContext.kind=VALUE]"

		usageBits += " [--sourceContext.gerrit.aliasContext.name=VALUE]"

		usageBits += " [--sourceContext.gerrit.aliasName=VALUE]"

		usageBits += " [--sourceContext.gerrit.gerritProject=VALUE]"

		usageBits += " [--sourceContext.gerrit.hostUri=VALUE]"

		usageBits += " [--sourceContext.gerrit.revisionId=VALUE]"

		usageBits += " [--sourceContext.git.revisionId=VALUE]"

		usageBits += " [--sourceContext.git.url=VALUE]"

		usageBits += " [--startPosition=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsReposWorkspacesFilesService(api_service)

	queryParamNames := map[string]bool{
		"pageSize":                                               false,
		"pageToken":                                              false,
		"sourceContext.cloudRepo.aliasContext.kind":              false,
		"sourceContext.cloudRepo.aliasContext.name":              false,
		"sourceContext.cloudRepo.aliasName":                      false,
		"sourceContext.cloudRepo.repoId.projectRepoId.projectId": false,
		"sourceContext.cloudRepo.repoId.projectRepoId.repoName":  false,
		"sourceContext.cloudRepo.repoId.uid":                     false,
		"sourceContext.cloudRepo.revisionId":                     false,
		"sourceContext.cloudWorkspace.snapshotId":                false,
		"sourceContext.cloudWorkspace.workspaceId.repoId.uid":    false,
		"sourceContext.gerrit.aliasContext.kind":                 false,
		"sourceContext.gerrit.aliasContext.name":                 false,
		"sourceContext.gerrit.aliasName":                         false,
		"sourceContext.gerrit.gerritProject":                     false,
		"sourceContext.gerrit.hostUri":                           false,
		"sourceContext.gerrit.revisionId":                        false,
		"sourceContext.git.revisionId":                           false,
		"sourceContext.git.url":                                  false,
		"startPosition":                                          false,
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
	if value, ok := flagValues["sourceContext_cloudRepo_aliasContext_kind"]; ok {
		query_sourceContext_cloudRepo_aliasContext_kind, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextCloudRepoAliasContextKind(query_sourceContext_cloudRepo_aliasContext_kind)
	}
	if value, ok := flagValues["sourceContext_cloudRepo_aliasContext_name"]; ok {
		query_sourceContext_cloudRepo_aliasContext_name, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextCloudRepoAliasContextName(query_sourceContext_cloudRepo_aliasContext_name)
	}
	if value, ok := flagValues["sourceContext_cloudRepo_aliasName"]; ok {
		query_sourceContext_cloudRepo_aliasName, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextCloudRepoAliasName(query_sourceContext_cloudRepo_aliasName)
	}
	if value, ok := flagValues["sourceContext_cloudRepo_repoId_projectRepoId_projectId"]; ok {
		query_sourceContext_cloudRepo_repoId_projectRepoId_projectId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextCloudRepoRepoIdProjectRepoIdProjectId(query_sourceContext_cloudRepo_repoId_projectRepoId_projectId)
	}
	if value, ok := flagValues["sourceContext_cloudRepo_repoId_projectRepoId_repoName"]; ok {
		query_sourceContext_cloudRepo_repoId_projectRepoId_repoName, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextCloudRepoRepoIdProjectRepoIdRepoName(query_sourceContext_cloudRepo_repoId_projectRepoId_repoName)
	}
	if value, ok := flagValues["sourceContext_cloudRepo_repoId_uid"]; ok {
		query_sourceContext_cloudRepo_repoId_uid, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextCloudRepoRepoIdUid(query_sourceContext_cloudRepo_repoId_uid)
	}
	if value, ok := flagValues["sourceContext_cloudRepo_revisionId"]; ok {
		query_sourceContext_cloudRepo_revisionId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextCloudRepoRevisionId(query_sourceContext_cloudRepo_revisionId)
	}
	if value, ok := flagValues["sourceContext_cloudWorkspace_snapshotId"]; ok {
		query_sourceContext_cloudWorkspace_snapshotId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextCloudWorkspaceSnapshotId(query_sourceContext_cloudWorkspace_snapshotId)
	}
	if value, ok := flagValues["sourceContext_cloudWorkspace_workspaceId_repoId_uid"]; ok {
		query_sourceContext_cloudWorkspace_workspaceId_repoId_uid, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextCloudWorkspaceWorkspaceIdRepoIdUid(query_sourceContext_cloudWorkspace_workspaceId_repoId_uid)
	}
	if value, ok := flagValues["sourceContext_gerrit_aliasContext_kind"]; ok {
		query_sourceContext_gerrit_aliasContext_kind, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextGerritAliasContextKind(query_sourceContext_gerrit_aliasContext_kind)
	}
	if value, ok := flagValues["sourceContext_gerrit_aliasContext_name"]; ok {
		query_sourceContext_gerrit_aliasContext_name, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextGerritAliasContextName(query_sourceContext_gerrit_aliasContext_name)
	}
	if value, ok := flagValues["sourceContext_gerrit_aliasName"]; ok {
		query_sourceContext_gerrit_aliasName, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextGerritAliasName(query_sourceContext_gerrit_aliasName)
	}
	if value, ok := flagValues["sourceContext_gerrit_gerritProject"]; ok {
		query_sourceContext_gerrit_gerritProject, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextGerritGerritProject(query_sourceContext_gerrit_gerritProject)
	}
	if value, ok := flagValues["sourceContext_gerrit_hostUri"]; ok {
		query_sourceContext_gerrit_hostUri, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextGerritHostUri(query_sourceContext_gerrit_hostUri)
	}
	if value, ok := flagValues["sourceContext_gerrit_revisionId"]; ok {
		query_sourceContext_gerrit_revisionId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextGerritRevisionId(query_sourceContext_gerrit_revisionId)
	}
	if value, ok := flagValues["sourceContext_git_revisionId"]; ok {
		query_sourceContext_git_revisionId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextGitRevisionId(query_sourceContext_git_revisionId)
	}
	if value, ok := flagValues["sourceContext_git_url"]; ok {
		query_sourceContext_git_url, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextGitUrl(query_sourceContext_git_url)
	}
	if value, ok := flagValues["startPosition"]; ok {
		query_startPosition, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.StartPosition(query_startPosition)
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

func Source_v1_ProjectsReposWorkspacesGet(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("repoName"))
		pathParams = append(pathParams, commands_util.AngrySnakes("name"))

		if len(pathParams) != 0 {
			if strings.Contains("v1/projects/{projectId}/repos/{repoName}/workspaces/{name}", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [--workspaceId.repoId.uid=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsReposWorkspacesService(api_service)

	queryParamNames := map[string]bool{
		"workspaceId.repoId.uid": false,
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
	if value, ok := flagValues["workspaceId_repoId_uid"]; ok {
		query_workspaceId_repoId_uid, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.WorkspaceIdRepoIdUid(query_workspaceId_repoId_uid)
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

func Source_v1_ProjectsReposWorkspacesList(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("repoName"))

		if len(pathParams) != 0 {
			if strings.Contains("v1/projects/{projectId}/repos/{repoName}/workspaces", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [--repoId.uid=VALUE]"

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
		"repoId.uid": false,
		"view":       false,
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
	if value, ok := flagValues["repoId_uid"]; ok {
		query_repoId_uid, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.RepoIdUid(query_repoId_uid)
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

func Source_v1_ProjectsReposWorkspacesListFiles(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("repoName"))
		pathParams = append(pathParams, commands_util.AngrySnakes("name"))

		if len(pathParams) != 0 {
			if strings.Contains("v1/projects/{projectId}/repos/{repoName}/workspaces/{name}:listFiles", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [--pageSize=VALUE]"

		usageBits += " [--pageToken=VALUE]"

		usageBits += " [--sourceContext.cloudRepo.aliasContext.kind=VALUE]"

		usageBits += " [--sourceContext.cloudRepo.aliasContext.name=VALUE]"

		usageBits += " [--sourceContext.cloudRepo.aliasName=VALUE]"

		usageBits += " [--sourceContext.cloudRepo.repoId.projectRepoId.projectId=VALUE]"

		usageBits += " [--sourceContext.cloudRepo.repoId.projectRepoId.repoName=VALUE]"

		usageBits += " [--sourceContext.cloudRepo.repoId.uid=VALUE]"

		usageBits += " [--sourceContext.cloudRepo.revisionId=VALUE]"

		usageBits += " [--sourceContext.cloudWorkspace.snapshotId=VALUE]"

		usageBits += " [--sourceContext.cloudWorkspace.workspaceId.repoId.uid=VALUE]"

		usageBits += " [--sourceContext.gerrit.aliasContext.kind=VALUE]"

		usageBits += " [--sourceContext.gerrit.aliasContext.name=VALUE]"

		usageBits += " [--sourceContext.gerrit.aliasName=VALUE]"

		usageBits += " [--sourceContext.gerrit.gerritProject=VALUE]"

		usageBits += " [--sourceContext.gerrit.hostUri=VALUE]"

		usageBits += " [--sourceContext.gerrit.revisionId=VALUE]"

		usageBits += " [--sourceContext.git.revisionId=VALUE]"

		usageBits += " [--sourceContext.git.url=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsReposWorkspacesService(api_service)

	queryParamNames := map[string]bool{
		"pageSize":                                               false,
		"pageToken":                                              false,
		"sourceContext.cloudRepo.aliasContext.kind":              false,
		"sourceContext.cloudRepo.aliasContext.name":              false,
		"sourceContext.cloudRepo.aliasName":                      false,
		"sourceContext.cloudRepo.repoId.projectRepoId.projectId": false,
		"sourceContext.cloudRepo.repoId.projectRepoId.repoName":  false,
		"sourceContext.cloudRepo.repoId.uid":                     false,
		"sourceContext.cloudRepo.revisionId":                     false,
		"sourceContext.cloudWorkspace.snapshotId":                false,
		"sourceContext.cloudWorkspace.workspaceId.repoId.uid":    false,
		"sourceContext.gerrit.aliasContext.kind":                 false,
		"sourceContext.gerrit.aliasContext.name":                 false,
		"sourceContext.gerrit.aliasName":                         false,
		"sourceContext.gerrit.gerritProject":                     false,
		"sourceContext.gerrit.hostUri":                           false,
		"sourceContext.gerrit.revisionId":                        false,
		"sourceContext.git.revisionId":                           false,
		"sourceContext.git.url":                                  false,
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
	if value, ok := flagValues["sourceContext_cloudRepo_aliasContext_kind"]; ok {
		query_sourceContext_cloudRepo_aliasContext_kind, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextCloudRepoAliasContextKind(query_sourceContext_cloudRepo_aliasContext_kind)
	}
	if value, ok := flagValues["sourceContext_cloudRepo_aliasContext_name"]; ok {
		query_sourceContext_cloudRepo_aliasContext_name, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextCloudRepoAliasContextName(query_sourceContext_cloudRepo_aliasContext_name)
	}
	if value, ok := flagValues["sourceContext_cloudRepo_aliasName"]; ok {
		query_sourceContext_cloudRepo_aliasName, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextCloudRepoAliasName(query_sourceContext_cloudRepo_aliasName)
	}
	if value, ok := flagValues["sourceContext_cloudRepo_repoId_projectRepoId_projectId"]; ok {
		query_sourceContext_cloudRepo_repoId_projectRepoId_projectId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextCloudRepoRepoIdProjectRepoIdProjectId(query_sourceContext_cloudRepo_repoId_projectRepoId_projectId)
	}
	if value, ok := flagValues["sourceContext_cloudRepo_repoId_projectRepoId_repoName"]; ok {
		query_sourceContext_cloudRepo_repoId_projectRepoId_repoName, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextCloudRepoRepoIdProjectRepoIdRepoName(query_sourceContext_cloudRepo_repoId_projectRepoId_repoName)
	}
	if value, ok := flagValues["sourceContext_cloudRepo_repoId_uid"]; ok {
		query_sourceContext_cloudRepo_repoId_uid, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextCloudRepoRepoIdUid(query_sourceContext_cloudRepo_repoId_uid)
	}
	if value, ok := flagValues["sourceContext_cloudRepo_revisionId"]; ok {
		query_sourceContext_cloudRepo_revisionId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextCloudRepoRevisionId(query_sourceContext_cloudRepo_revisionId)
	}
	if value, ok := flagValues["sourceContext_cloudWorkspace_snapshotId"]; ok {
		query_sourceContext_cloudWorkspace_snapshotId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextCloudWorkspaceSnapshotId(query_sourceContext_cloudWorkspace_snapshotId)
	}
	if value, ok := flagValues["sourceContext_cloudWorkspace_workspaceId_repoId_uid"]; ok {
		query_sourceContext_cloudWorkspace_workspaceId_repoId_uid, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextCloudWorkspaceWorkspaceIdRepoIdUid(query_sourceContext_cloudWorkspace_workspaceId_repoId_uid)
	}
	if value, ok := flagValues["sourceContext_gerrit_aliasContext_kind"]; ok {
		query_sourceContext_gerrit_aliasContext_kind, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextGerritAliasContextKind(query_sourceContext_gerrit_aliasContext_kind)
	}
	if value, ok := flagValues["sourceContext_gerrit_aliasContext_name"]; ok {
		query_sourceContext_gerrit_aliasContext_name, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextGerritAliasContextName(query_sourceContext_gerrit_aliasContext_name)
	}
	if value, ok := flagValues["sourceContext_gerrit_aliasName"]; ok {
		query_sourceContext_gerrit_aliasName, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextGerritAliasName(query_sourceContext_gerrit_aliasName)
	}
	if value, ok := flagValues["sourceContext_gerrit_gerritProject"]; ok {
		query_sourceContext_gerrit_gerritProject, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextGerritGerritProject(query_sourceContext_gerrit_gerritProject)
	}
	if value, ok := flagValues["sourceContext_gerrit_hostUri"]; ok {
		query_sourceContext_gerrit_hostUri, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextGerritHostUri(query_sourceContext_gerrit_hostUri)
	}
	if value, ok := flagValues["sourceContext_gerrit_revisionId"]; ok {
		query_sourceContext_gerrit_revisionId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextGerritRevisionId(query_sourceContext_gerrit_revisionId)
	}
	if value, ok := flagValues["sourceContext_git_revisionId"]; ok {
		query_sourceContext_git_revisionId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextGitRevisionId(query_sourceContext_git_revisionId)
	}
	if value, ok := flagValues["sourceContext_git_url"]; ok {
		query_sourceContext_git_url, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextGitUrl(query_sourceContext_git_url)
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

func Source_v1_ProjectsReposWorkspacesModifyWorkspace(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("repoName"))
		pathParams = append(pathParams, commands_util.AngrySnakes("name"))

		if len(pathParams) != 0 {
			if strings.Contains("v1/projects/{projectId}/repos/{repoName}/workspaces/{name}:modifyWorkspace", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		commands_util.PrintRequestExample(&api_client.ModifyWorkspaceRequest{})

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

func Source_v1_ProjectsReposWorkspacesRefreshWorkspace(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("repoName"))
		pathParams = append(pathParams, commands_util.AngrySnakes("name"))

		if len(pathParams) != 0 {
			if strings.Contains("v1/projects/{projectId}/repos/{repoName}/workspaces/{name}:refreshWorkspace", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		commands_util.PrintRequestExample(&api_client.RefreshWorkspaceRequest{})

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

func Source_v1_ProjectsReposWorkspacesResolveFiles(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("repoName"))
		pathParams = append(pathParams, commands_util.AngrySnakes("name"))

		if len(pathParams) != 0 {
			if strings.Contains("v1/projects/{projectId}/repos/{repoName}/workspaces/{name}:resolveFiles", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		commands_util.PrintRequestExample(&api_client.ResolveFilesRequest{})

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

func Source_v1_ProjectsReposWorkspacesRevertRefresh(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("repoName"))
		pathParams = append(pathParams, commands_util.AngrySnakes("name"))

		if len(pathParams) != 0 {
			if strings.Contains("v1/projects/{projectId}/repos/{repoName}/workspaces/{name}:revertRefresh", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
		commands_util.PrintRequestExample(&api_client.RevertRefreshRequest{})

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

func Source_v1_ProjectsReposWorkspacesSnapshotsFilesGet(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("repoName"))
		pathParams = append(pathParams, commands_util.AngrySnakes("name"))
		pathParams = append(pathParams, commands_util.AngrySnakes("snapshotId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("path"))

		if len(pathParams) != 0 {
			if strings.Contains("v1/projects/{projectId}/repos/{repoName}/workspaces/{name}/snapshots/{snapshotId}/files/{+path}", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [--pageSize=VALUE]"

		usageBits += " [--pageToken=VALUE]"

		usageBits += " [--sourceContext.cloudRepo.aliasContext.kind=VALUE]"

		usageBits += " [--sourceContext.cloudRepo.aliasContext.name=VALUE]"

		usageBits += " [--sourceContext.cloudRepo.aliasName=VALUE]"

		usageBits += " [--sourceContext.cloudRepo.repoId.projectRepoId.projectId=VALUE]"

		usageBits += " [--sourceContext.cloudRepo.repoId.projectRepoId.repoName=VALUE]"

		usageBits += " [--sourceContext.cloudRepo.repoId.uid=VALUE]"

		usageBits += " [--sourceContext.cloudRepo.revisionId=VALUE]"

		usageBits += " [--sourceContext.cloudWorkspace.workspaceId.repoId.uid=VALUE]"

		usageBits += " [--sourceContext.gerrit.aliasContext.kind=VALUE]"

		usageBits += " [--sourceContext.gerrit.aliasContext.name=VALUE]"

		usageBits += " [--sourceContext.gerrit.aliasName=VALUE]"

		usageBits += " [--sourceContext.gerrit.gerritProject=VALUE]"

		usageBits += " [--sourceContext.gerrit.hostUri=VALUE]"

		usageBits += " [--sourceContext.gerrit.revisionId=VALUE]"

		usageBits += " [--sourceContext.git.revisionId=VALUE]"

		usageBits += " [--sourceContext.git.url=VALUE]"

		usageBits += " [--startPosition=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsReposWorkspacesSnapshotsFilesService(api_service)

	queryParamNames := map[string]bool{
		"pageSize":                                               false,
		"pageToken":                                              false,
		"sourceContext.cloudRepo.aliasContext.kind":              false,
		"sourceContext.cloudRepo.aliasContext.name":              false,
		"sourceContext.cloudRepo.aliasName":                      false,
		"sourceContext.cloudRepo.repoId.projectRepoId.projectId": false,
		"sourceContext.cloudRepo.repoId.projectRepoId.repoName":  false,
		"sourceContext.cloudRepo.repoId.uid":                     false,
		"sourceContext.cloudRepo.revisionId":                     false,
		"sourceContext.cloudWorkspace.workspaceId.repoId.uid":    false,
		"sourceContext.gerrit.aliasContext.kind":                 false,
		"sourceContext.gerrit.aliasContext.name":                 false,
		"sourceContext.gerrit.aliasName":                         false,
		"sourceContext.gerrit.gerritProject":                     false,
		"sourceContext.gerrit.hostUri":                           false,
		"sourceContext.gerrit.revisionId":                        false,
		"sourceContext.git.revisionId":                           false,
		"sourceContext.git.url":                                  false,
		"startPosition":                                          false,
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
	if value, ok := flagValues["sourceContext_cloudRepo_aliasContext_kind"]; ok {
		query_sourceContext_cloudRepo_aliasContext_kind, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextCloudRepoAliasContextKind(query_sourceContext_cloudRepo_aliasContext_kind)
	}
	if value, ok := flagValues["sourceContext_cloudRepo_aliasContext_name"]; ok {
		query_sourceContext_cloudRepo_aliasContext_name, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextCloudRepoAliasContextName(query_sourceContext_cloudRepo_aliasContext_name)
	}
	if value, ok := flagValues["sourceContext_cloudRepo_aliasName"]; ok {
		query_sourceContext_cloudRepo_aliasName, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextCloudRepoAliasName(query_sourceContext_cloudRepo_aliasName)
	}
	if value, ok := flagValues["sourceContext_cloudRepo_repoId_projectRepoId_projectId"]; ok {
		query_sourceContext_cloudRepo_repoId_projectRepoId_projectId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextCloudRepoRepoIdProjectRepoIdProjectId(query_sourceContext_cloudRepo_repoId_projectRepoId_projectId)
	}
	if value, ok := flagValues["sourceContext_cloudRepo_repoId_projectRepoId_repoName"]; ok {
		query_sourceContext_cloudRepo_repoId_projectRepoId_repoName, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextCloudRepoRepoIdProjectRepoIdRepoName(query_sourceContext_cloudRepo_repoId_projectRepoId_repoName)
	}
	if value, ok := flagValues["sourceContext_cloudRepo_repoId_uid"]; ok {
		query_sourceContext_cloudRepo_repoId_uid, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextCloudRepoRepoIdUid(query_sourceContext_cloudRepo_repoId_uid)
	}
	if value, ok := flagValues["sourceContext_cloudRepo_revisionId"]; ok {
		query_sourceContext_cloudRepo_revisionId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextCloudRepoRevisionId(query_sourceContext_cloudRepo_revisionId)
	}
	if value, ok := flagValues["sourceContext_cloudWorkspace_workspaceId_repoId_uid"]; ok {
		query_sourceContext_cloudWorkspace_workspaceId_repoId_uid, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextCloudWorkspaceWorkspaceIdRepoIdUid(query_sourceContext_cloudWorkspace_workspaceId_repoId_uid)
	}
	if value, ok := flagValues["sourceContext_gerrit_aliasContext_kind"]; ok {
		query_sourceContext_gerrit_aliasContext_kind, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextGerritAliasContextKind(query_sourceContext_gerrit_aliasContext_kind)
	}
	if value, ok := flagValues["sourceContext_gerrit_aliasContext_name"]; ok {
		query_sourceContext_gerrit_aliasContext_name, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextGerritAliasContextName(query_sourceContext_gerrit_aliasContext_name)
	}
	if value, ok := flagValues["sourceContext_gerrit_aliasName"]; ok {
		query_sourceContext_gerrit_aliasName, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextGerritAliasName(query_sourceContext_gerrit_aliasName)
	}
	if value, ok := flagValues["sourceContext_gerrit_gerritProject"]; ok {
		query_sourceContext_gerrit_gerritProject, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextGerritGerritProject(query_sourceContext_gerrit_gerritProject)
	}
	if value, ok := flagValues["sourceContext_gerrit_hostUri"]; ok {
		query_sourceContext_gerrit_hostUri, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextGerritHostUri(query_sourceContext_gerrit_hostUri)
	}
	if value, ok := flagValues["sourceContext_gerrit_revisionId"]; ok {
		query_sourceContext_gerrit_revisionId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextGerritRevisionId(query_sourceContext_gerrit_revisionId)
	}
	if value, ok := flagValues["sourceContext_git_revisionId"]; ok {
		query_sourceContext_git_revisionId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextGitRevisionId(query_sourceContext_git_revisionId)
	}
	if value, ok := flagValues["sourceContext_git_url"]; ok {
		query_sourceContext_git_url, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextGitUrl(query_sourceContext_git_url)
	}
	if value, ok := flagValues["startPosition"]; ok {
		query_startPosition, err := commands_util.ConvertValue_int64(value)
		if err != nil {
			return err
		}
		call.StartPosition(query_startPosition)
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

func Source_v1_ProjectsReposWorkspacesSnapshotsGet(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("repoName"))
		pathParams = append(pathParams, commands_util.AngrySnakes("name"))
		pathParams = append(pathParams, commands_util.AngrySnakes("snapshotId"))

		if len(pathParams) != 0 {
			if strings.Contains("v1/projects/{projectId}/repos/{repoName}/workspaces/{name}/snapshots/{snapshotId}", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [--workspaceId.repoId.uid=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsReposWorkspacesSnapshotsService(api_service)

	queryParamNames := map[string]bool{
		"workspaceId.repoId.uid": false,
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
	if value, ok := flagValues["workspaceId_repoId_uid"]; ok {
		query_workspaceId_repoId_uid, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.WorkspaceIdRepoIdUid(query_workspaceId_repoId_uid)
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

func Source_v1_ProjectsReposWorkspacesSnapshotsList(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("repoName"))
		pathParams = append(pathParams, commands_util.AngrySnakes("name"))

		if len(pathParams) != 0 {
			if strings.Contains("v1/projects/{projectId}/repos/{repoName}/workspaces/{name}/snapshots", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [--pageSize=VALUE]"

		usageBits += " [--pageToken=VALUE]"

		usageBits += " [--workspaceId.repoId.uid=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsReposWorkspacesSnapshotsService(api_service)

	queryParamNames := map[string]bool{
		"pageSize":               false,
		"pageToken":              false,
		"workspaceId.repoId.uid": false,
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
	if value, ok := flagValues["workspaceId_repoId_uid"]; ok {
		query_workspaceId_repoId_uid, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.WorkspaceIdRepoIdUid(query_workspaceId_repoId_uid)
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

func Source_v1_ProjectsReposWorkspacesSnapshotsListFiles(context Context, args ...string) error {

	usageFunc := func() {
		usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
		var pathParams []string
		pathParams = append(pathParams, commands_util.AngrySnakes("projectId"))
		pathParams = append(pathParams, commands_util.AngrySnakes("repoName"))
		pathParams = append(pathParams, commands_util.AngrySnakes("name"))
		pathParams = append(pathParams, commands_util.AngrySnakes("snapshotId"))

		if len(pathParams) != 0 {
			if strings.Contains("v1/projects/{projectId}/repos/{repoName}/workspaces/{name}/snapshots/{snapshotId}:listFiles", "+") {
				usageBits += " @" + strings.Join(pathParams, "@")
			} else {
				usageBits += " " + strings.Join(pathParams, "/")
			}
		}

		usageBits += " [--pageSize=VALUE]"

		usageBits += " [--pageToken=VALUE]"

		usageBits += " [--sourceContext.cloudRepo.aliasContext.kind=VALUE]"

		usageBits += " [--sourceContext.cloudRepo.aliasContext.name=VALUE]"

		usageBits += " [--sourceContext.cloudRepo.aliasName=VALUE]"

		usageBits += " [--sourceContext.cloudRepo.repoId.projectRepoId.projectId=VALUE]"

		usageBits += " [--sourceContext.cloudRepo.repoId.projectRepoId.repoName=VALUE]"

		usageBits += " [--sourceContext.cloudRepo.repoId.uid=VALUE]"

		usageBits += " [--sourceContext.cloudRepo.revisionId=VALUE]"

		usageBits += " [--sourceContext.cloudWorkspace.workspaceId.repoId.uid=VALUE]"

		usageBits += " [--sourceContext.gerrit.aliasContext.kind=VALUE]"

		usageBits += " [--sourceContext.gerrit.aliasContext.name=VALUE]"

		usageBits += " [--sourceContext.gerrit.aliasName=VALUE]"

		usageBits += " [--sourceContext.gerrit.gerritProject=VALUE]"

		usageBits += " [--sourceContext.gerrit.hostUri=VALUE]"

		usageBits += " [--sourceContext.gerrit.revisionId=VALUE]"

		usageBits += " [--sourceContext.git.revisionId=VALUE]"

		usageBits += " [--sourceContext.git.url=VALUE]"

		fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)

		os.Exit(1)
	}

	api_service, err := api_client.New(context.Client)
	if err != nil {
		return err
	}
	service := api_client.NewProjectsReposWorkspacesSnapshotsService(api_service)

	queryParamNames := map[string]bool{
		"pageSize":                                               false,
		"pageToken":                                              false,
		"sourceContext.cloudRepo.aliasContext.kind":              false,
		"sourceContext.cloudRepo.aliasContext.name":              false,
		"sourceContext.cloudRepo.aliasName":                      false,
		"sourceContext.cloudRepo.repoId.projectRepoId.projectId": false,
		"sourceContext.cloudRepo.repoId.projectRepoId.repoName":  false,
		"sourceContext.cloudRepo.repoId.uid":                     false,
		"sourceContext.cloudRepo.revisionId":                     false,
		"sourceContext.cloudWorkspace.workspaceId.repoId.uid":    false,
		"sourceContext.gerrit.aliasContext.kind":                 false,
		"sourceContext.gerrit.aliasContext.name":                 false,
		"sourceContext.gerrit.aliasName":                         false,
		"sourceContext.gerrit.gerritProject":                     false,
		"sourceContext.gerrit.hostUri":                           false,
		"sourceContext.gerrit.revisionId":                        false,
		"sourceContext.git.revisionId":                           false,
		"sourceContext.git.url":                                  false,
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
	if value, ok := flagValues["sourceContext_cloudRepo_aliasContext_kind"]; ok {
		query_sourceContext_cloudRepo_aliasContext_kind, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextCloudRepoAliasContextKind(query_sourceContext_cloudRepo_aliasContext_kind)
	}
	if value, ok := flagValues["sourceContext_cloudRepo_aliasContext_name"]; ok {
		query_sourceContext_cloudRepo_aliasContext_name, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextCloudRepoAliasContextName(query_sourceContext_cloudRepo_aliasContext_name)
	}
	if value, ok := flagValues["sourceContext_cloudRepo_aliasName"]; ok {
		query_sourceContext_cloudRepo_aliasName, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextCloudRepoAliasName(query_sourceContext_cloudRepo_aliasName)
	}
	if value, ok := flagValues["sourceContext_cloudRepo_repoId_projectRepoId_projectId"]; ok {
		query_sourceContext_cloudRepo_repoId_projectRepoId_projectId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextCloudRepoRepoIdProjectRepoIdProjectId(query_sourceContext_cloudRepo_repoId_projectRepoId_projectId)
	}
	if value, ok := flagValues["sourceContext_cloudRepo_repoId_projectRepoId_repoName"]; ok {
		query_sourceContext_cloudRepo_repoId_projectRepoId_repoName, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextCloudRepoRepoIdProjectRepoIdRepoName(query_sourceContext_cloudRepo_repoId_projectRepoId_repoName)
	}
	if value, ok := flagValues["sourceContext_cloudRepo_repoId_uid"]; ok {
		query_sourceContext_cloudRepo_repoId_uid, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextCloudRepoRepoIdUid(query_sourceContext_cloudRepo_repoId_uid)
	}
	if value, ok := flagValues["sourceContext_cloudRepo_revisionId"]; ok {
		query_sourceContext_cloudRepo_revisionId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextCloudRepoRevisionId(query_sourceContext_cloudRepo_revisionId)
	}
	if value, ok := flagValues["sourceContext_cloudWorkspace_workspaceId_repoId_uid"]; ok {
		query_sourceContext_cloudWorkspace_workspaceId_repoId_uid, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextCloudWorkspaceWorkspaceIdRepoIdUid(query_sourceContext_cloudWorkspace_workspaceId_repoId_uid)
	}
	if value, ok := flagValues["sourceContext_gerrit_aliasContext_kind"]; ok {
		query_sourceContext_gerrit_aliasContext_kind, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextGerritAliasContextKind(query_sourceContext_gerrit_aliasContext_kind)
	}
	if value, ok := flagValues["sourceContext_gerrit_aliasContext_name"]; ok {
		query_sourceContext_gerrit_aliasContext_name, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextGerritAliasContextName(query_sourceContext_gerrit_aliasContext_name)
	}
	if value, ok := flagValues["sourceContext_gerrit_aliasName"]; ok {
		query_sourceContext_gerrit_aliasName, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextGerritAliasName(query_sourceContext_gerrit_aliasName)
	}
	if value, ok := flagValues["sourceContext_gerrit_gerritProject"]; ok {
		query_sourceContext_gerrit_gerritProject, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextGerritGerritProject(query_sourceContext_gerrit_gerritProject)
	}
	if value, ok := flagValues["sourceContext_gerrit_hostUri"]; ok {
		query_sourceContext_gerrit_hostUri, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextGerritHostUri(query_sourceContext_gerrit_hostUri)
	}
	if value, ok := flagValues["sourceContext_gerrit_revisionId"]; ok {
		query_sourceContext_gerrit_revisionId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextGerritRevisionId(query_sourceContext_gerrit_revisionId)
	}
	if value, ok := flagValues["sourceContext_git_revisionId"]; ok {
		query_sourceContext_git_revisionId, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextGitRevisionId(query_sourceContext_git_revisionId)
	}
	if value, ok := flagValues["sourceContext_git_url"]; ok {
		query_sourceContext_git_url, err := commands_util.ConvertValue_string(value)
		if err != nil {
			return err
		}
		call.SourceContextGitUrl(query_sourceContext_git_url)
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
