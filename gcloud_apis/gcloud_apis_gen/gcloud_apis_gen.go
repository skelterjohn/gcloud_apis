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
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"

	"github.com/googlecloudplatform/gcloud/gcloud_apis/gcloud_apis_gen/discovery"
)

func capitalize(s string) string {
	return strings.ToUpper(s[:1]) + s[1:]
}

func serviceMethodName(methodId string) string {
	tokens := strings.Split(methodId, ".")
	return capitalize(tokens[len(tokens)-1])
}

func funcNameForMethod(m discovery.Method, v string) (name string) {
	parts := strings.Split(m.Id, ".")
	name = fmt.Sprintf("%s_%s_", capitalize(parts[0]), v)
	for _, part := range parts[1:] {
		name += capitalize(part)
	}
	return
}

func commandForMethod(method discovery.Method, doc *discovery.Discovery) string {
	api_then_method := strings.SplitN(method.Id, ".", 2)
	return api_then_method[0] + ":" + doc.Version + "." + api_then_method[1]
}

func commandForPreferredMethod(method discovery.Method, doc *discovery.Discovery) string {
	return method.Id
}

type QueryFuncInfo struct {
	FuncName  string
	Type      string
	Parameter discovery.Parameter
}

type PositionalParam struct {
	ParamName string
	QueryName string
	Type      string
}

type InfoForMethod struct {
	FuncName          string
	ServiceName       string
	ServiceMethodName string
	Method            discovery.Method
	QueryFuncs        map[string]QueryFuncInfo
	PositionalParams  []PositionalParam
	HasFlags          bool
	HasQuery          bool
}

type InfosForMethod []InfoForMethod

func (is InfosForMethod) Len() int           { return len(is) }
func (is InfosForMethod) Swap(i, j int)      { is[i], is[j] = is[j], is[i] }
func (is InfosForMethod) Less(i, j int) bool { return is[i].FuncName < is[j].FuncName }

/*
Write the source code for the commands. This includes a single "commands.go",
which makes a mapping from method id to command func, and a file for every
discovery doc in the form of "API_VERSION.go".
*/
func writeCommandsSource(commands_dir string, docs map[string]*discovery.Discovery) error {
	commandsFile, err := os.Create(filepath.Join(commands_dir, "commands.go"))
	if err != nil {
		log.Fatal(err)
	}

	methodFuncs := map[string]string{}
	methods := map[string]discovery.Method{}

	for _, doc := range docs {

		preferredVersion, ok := preferredVersions[doc.Name]
		if !ok {
			panic("no preferred version for " + doc.Name)
		}
		isPreferred := preferredVersion == doc.Version

		type InfoForFile struct {
			ClientImportPath string
			Methods          []InfoForMethod
		}
		var info InfoForFile
		info.ClientImportPath = fmt.Sprintf(
			"github.com/googlecloudplatform/gcloud/gcloud_apis/clients/%s/%s",
			doc.Name, doc.Version)

		var addResource func(serviceName string, resource discovery.Resource)
		addResource = func(serviceName string, resource discovery.Resource) {
			for _, method := range resource.Methods {

				funcName := funcNameForMethod(method, doc.Version)

				methodFuncs[commandForMethod(method, doc)] = funcName
				if isPreferred {
					methodFuncs[commandForPreferredMethod(method, doc)] = funcName
				}

				methods[funcName] = method
				methodInfo := InfoForMethod{
					FuncName:          funcName,
					ServiceName:       capitalize(serviceName),
					ServiceMethodName: serviceMethodName(method.Id),
					Method:            method,
					HasFlags:          method.Request.Ref != "",
				}
				positionalParams := map[string]PositionalParam{}
				methodInfo.QueryFuncs = make(map[string]QueryFuncInfo)
				for paramName, paramData := range method.Parameters {
					var paramType string
					switch paramData.Type {
					case "string":
						switch paramData.Format {
						case "uint64":
							paramType = "uint64"
						case "int64":
							paramType = "int64"
						case "uint32":
							paramType = "uint32"
						case "int32":
							paramType = "int32"
						default:
							paramType = "string"
						}
					case "integer":
						paramType = "int64"
					case "number":
						paramType = "float64"
					case "boolean":
						paramType = "bool"
					default:
						panic("unknown parameter type: " + paramData.Type)
					}
					switch paramData.Location {
					case "query":
						methodInfo.HasFlags = true
						methodInfo.HasQuery = true
						if paramData.Required {
							positionalParams[paramName] = PositionalParam{
								ParamName: paramName,
								QueryName: paramName,
								Type:      paramType,
							}
						} else {
							methodInfo.QueryFuncs[paramName] = QueryFuncInfo{
								FuncName:  capitalize(paramName),
								Parameter: paramData,
								Type:      paramType,
							}
						}
					case "path":
						positionalParams[paramName] = PositionalParam{
							ParamName: paramName,
							Type:      paramType,
						}
					default:
						panic(fmt.Sprintf("unexpected param location: %q", paramData.Location))
					}
				}
				for _, name := range method.ParameterOrder {
					methodInfo.PositionalParams = append(
						methodInfo.PositionalParams,
						positionalParams[name],
					)
				}

				info.Methods = append(info.Methods, methodInfo)
			}
			for subServiceName, subResource := range resource.Resources {
				addResource(serviceName+capitalize(subServiceName), subResource)
			}
		}

		for serviceName, resource := range doc.Resources {
			addResource(serviceName, resource)
		}

		sort.Sort(InfosForMethod(info.Methods))

		apiCommandsPath := filepath.Join(commands_dir, fmt.Sprintf("%s_%s.go", doc.Name, doc.Version))
		var apiCommandsFile *os.File
		apiCommandsFile, err = os.Create(apiCommandsPath)

		commandsSourceTemplate.ExecuteTemplate(apiCommandsFile, "api_commands", info)
	}
	commandsSourceTemplate.ExecuteTemplate(commandsFile, "all_commands", methodFuncs)

	// Ensure that all the generated code is property go fmt ed.
	cmd := exec.Command("go", "fmt", ".")
	cmd.Dir = commands_dir

	if err := cmd.Run(); err != nil {
		// this err shadows intentionally, only print it to log.
		log.Printf("while formatting in %s: %v", cmd.Dir, err)
	}

	return err
}

func writeClientsSource(discovery_dir, clients_dir string, docs map[string]*discovery.Discovery) error {
	for api, doc := range docs {
		discovery_file := filepath.Join(discovery_dir, api+".json")
		client_source_dir := filepath.Join(clients_dir, doc.Name, doc.Version)
		client_source_file := filepath.Join(client_source_dir, api+".go")

		clients_stat, err := os.Stat(clients_dir)
		if err != nil {
			return fmt.Errorf("directory %q does not exist", clients_dir)
		}

		err = os.MkdirAll(client_source_dir, clients_stat.Mode())
		if err != nil {
			return fmt.Errorf("could not create directory %q", client_source_dir)
		}
		cmd := exec.Command(
			"google-api-go-generator",
			"-api_json_file", discovery_file,
			"-output", client_source_file)
		err = cmd.Run()
		if err != nil {
			return fmt.Errorf("could not generate client: %s", err.Error())
		}
	}
	return nil
}

func main() {
	var discovery_dir, clients_dir, commands_dir string

	fs := flag.NewFlagSet("gcloud_apis_gen", flag.ExitOnError)
	fs.StringVar(&discovery_dir, "discovery-dir", "", "directory to find discovery docs")
	fs.StringVar(&clients_dir, "clients-dir", "", "directory to write client source")
	fs.StringVar(&commands_dir, "commands-dir", "", "directory to write command source")
	err := fs.Parse(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	docs, err := discovery.LoadDiscoveriesFromDirectory(discovery_dir)
	if err != nil {
		log.Fatal(err)
	}

	err = writeCommandsSource(commands_dir, docs)
	if err != nil {
		log.Fatal(err)
	}

	err = writeClientsSource(discovery_dir, clients_dir, docs)
	if err != nil {
		log.Fatal(err)
	}
}
