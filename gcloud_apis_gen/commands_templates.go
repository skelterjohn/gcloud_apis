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

package main

import (
	"text/template"
)

var commandsSourceTemplate = template.New("commands")

func init() {
	commandsFormats := []string{
		commandsFormat,
		apiCommandsFormat,
		methodFormat,
	}
	for _, format := range commandsFormats {
		commandsSourceTemplate = template.Must(commandsSourceTemplate.Parse(format))
	}
}

var commandsFormat = `{{define "all_commands"}}// generated code: commands.go
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
  "net/http"
)

type Context struct {
  Client *http.Client
  InvocationMethod string
}

type Command func(context Context, args ...string) (err error)

var AllMethods = map[string]Command{
{{range $id, $func := .}}  "{{$id}}": {{$func}},
{{end}}}
{{end}}`

var apiCommandsFormat = `{{define "api_commands"}}// generated code: api commands
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

  api_client "{{.ClientImportPath}}"
  "github.com/GoogleCloudPlatform/gcloud/gcloud_apis/commands_util"
)

// some generated code may not use all imported libs
var _ = fmt.Println
var _ = io.Copy
var _ = os.Stdin
var _ = strings.Split

{{range .Methods}}{{template "api_method" .}}{{end}}

{{end}}`

var methodFormat = `{{define "api_method"}}
func {{.FuncName}}(context Context, args ...string) error {

  usageFunc := func() {
    usageBits := fmt.Sprintf("gcloud_apis %s", context.InvocationMethod)
    {{if .Method.Request.Ref}}
    usageBits += " [REQUEST_FILE|-] [--REQUEST_KEY=VALUE]*"
    {{end}}
    {{if .HasQuery}}
    {{ range $k, $p := .Method.Parameters}}{{  if eq $p.Location "query"}}
    {{   if $p.Required}}
    usageBits += " --{{$k}}=VALUE"
    {{   else}}
    usageBits += " [--{{$k}}=VALUE]"
    {{   end}}
    {{  end}}{{ end}}
    {{end}}
    fmt.Fprintf(os.Stderr, "Usage:\n\t%s\n", usageBits)
    os.Exit(1)
  }

  api_service, err := api_client.New(context.Client)
  if err != nil {
    return err
  }
  service := api_client.New{{.ServiceName}}Service(api_service)

  {{if .HasQuery}}queryParamNames := map[string]bool{
    {{ range $k, $p := .Method.Parameters}}{{  if eq $p.Location "query"}}"{{$k}}": {{$p.Required}},
    {{  end}}{{ end}}
  }{{end}}

  {{if .HasFlags}}args, flagValues, err := commands_util.ExtractFlagValues(args)
  if err != nil {
    return err
  }
  {{end}}
  {{if .HasQuery}}for k, r := range queryParamNames {
    if _, ok := flagValues[k]; r && !ok {
      return fmt.Errorf("missing required flag %q", "--"+k)
    }
  }{{end}}

  // Only positional arguments should remain in args.
  {{if .Method.Request.Ref}}if len(args) == 0 || len(args) > 2 {
    usageFunc()
  }

  request := &api_client.{{.Method.Request.Ref}}{}
  if len(args) == 2 {
    err = commands_util.PopulateRequestFromFilename(&request, args[1])
    if err != nil {
      return err
    }
  }

  {{ if .Method.SupportsMediaUpload}}
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
  {{ end}}

  {{ if .HasQuery}}
  // Any flags that aren't query parameters are applied to the request.
  keyValues := map[string]string{}
  for k, v := range flagValues {
    if _, ok := queryParamNames[k]; !ok {
      keyValues[k] = v
    }
  }
  {{ else}}
  keyValues := flagValues
  {{ end}}
  err = commands_util.OverwriteRequestWithValues(&request, keyValues)
  if err != nil {
    return err
  }{{else}}if len(args) != 1 {
    usageFunc()
  }
  {{end}}

  expectedParams := []string{
    {{range .PositionalParams}}{{ if .QueryName}}{{ else}}"{{.ParamName}}",{{ end}}
    {{end}}}
  paramValues := commands_util.SplitParamValues(args[0])
  if len(paramValues) != len(expectedParams) {
    return commands_util.ErrForWrongParams(expectedParams, paramValues, args)
  }

  {{range $i, $v := .PositionalParams}}{{/*
    */}}{{if .QueryName}}
    param_{{.ParamName}}, err := commands_util.ConvertValue_{{.Type}}(flagValues["{{.QueryName}}"])
    if err != nil {
      return err
    }{{/*
    */}}{{else}}
    param_{{.ParamName}}, err := commands_util.ConvertValue_{{.Type}}(paramValues[{{$i}}])
    if err != nil {
      return err
    }{{/*
    */}}{{end}}{{/*
  */}}{{end}}

  call := service.{{.ServiceMethodName}}({{/*
  */}}{{range $i, $v := .PositionalParams}}{{/*
    */}}param_{{.ParamName}},{{/*
  */}}{{end}}{{/*
  */}}{{if .Method.Request.Ref}}
    request,{{/*
  */}}{{end}}
  )
  
  {{if .QueryFuncs}}// Set query parameters.
  {{ range $name, $querydata := .QueryFuncs}}if value, ok := flagValues["{{$name}}"]; ok {
      query_{{$name}}, err := commands_util.ConvertValue_{{$querydata.Type}}(value)
      if err != nil {
        return err
      }
      call.{{$querydata.FuncName}}(query_{{$name}})
    }
  {{ end}}
  {{end}}

  {{if .Method.SupportsMediaUpload}}
  if media != nil {
    call.Media(media)
  }
  {{end}}

  {{if .Method.Response.Ref}}
  var response *api_client.{{.Method.Response.Ref}}
  response, {{end}}err = call.Do()
  if err != nil {
    return err
  }

  {{if .Method.Response.Ref}}err = commands_util.PrintResponse(response)
  if err != nil {
    return err
  }{{end}}

  return nil
}
{{end}}`
