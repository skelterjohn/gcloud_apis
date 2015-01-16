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

//go:generate ./gen.bash

package main

import (
	"fmt"
	"os"

	"github.com/GoogleCloudPlatform/gcloud/gcloud_apis/commands"
)

func Usage() {
	fmt.Fprintf(os.Stderr, `Usage:
	gcloud_apis list [METHOD|PARTIAL_METHOD]
	gcloud_apis METHOD [PARAM[/PARAM]*] [--KEY=VALUE]* [REQUEST_BODY]
`)
	os.Exit(1)
}

func problem(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(2)
}

func main() {
	if len(os.Args) == 1 {
		Usage()
	}

	switch os.Args[1] {
	case "list":
		if len(os.Args) > 3 {
			Usage()
		}
		prefix := ""
		if len(os.Args) == 3 {
			prefix = os.Args[2]
		}
		err := ListMethods(prefix)
		if err != nil {
			problem(err)
		}
	default:
		methodToRun := os.Args[1]
		methodFunc, ok := commands.AllMethods[methodToRun]
		if !ok {
			fmt.Fprintf(os.Stderr, "Unknown method %q. To list methods, run '$ gcloud_apis list'.\n", methodToRun)
			os.Exit(1)
		}
		client, err := getAuthenticatedClient()
		if err != nil {
			problem(err)
		}
		context := commands.Context{
			InvocationMethod: methodToRun,
			Client:           client,
		}

		err = methodFunc(context, os.Args[2:]...)
		if err != nil {
			problem(err)
		}
	}

}
