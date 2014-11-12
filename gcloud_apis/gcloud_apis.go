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
	"log"
	"os"

	"github.com/googlecloudplatform/gcloud/gcloud_apis/commands"
)

func Usage() {
	log.Fatal("usage issue")
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
			log.Fatal(err)
		}
	default:
		methodToRun := os.Args[1]
		methodFunc, ok := commands.AllMethods[methodToRun]
		if !ok {
			log.Fatalf("unknown method %q. To list methods, run '$ gcloud_apis list'.", methodToRun)
		}
		client, err := getAuthenticatedClient()
		if err != nil {
			log.Fatal(err)
		}
		context := commands.Context{
			InvocationMethod: methodToRun,
			Client:           client,
		}

		err = methodFunc(context, os.Args[2:]...)
		if err != nil {
			log.Fatal(err)
		}
	}

}
