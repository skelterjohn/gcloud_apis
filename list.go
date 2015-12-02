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
	"fmt"
	"sort"
	"strings"

	"github.com/GoogleCloudPlatform/gcloud/gcloud_apis/commands"
)

func ListMethods(args []string) error {
	fullNames := false
	for i, arg := range args {
		if arg == "--full" {
			fullNames = true
			args = append(args[:i], args[i+1:]...)
			break
		}
	}

	if len(args) > 1 {
		Usage()
	}
	prefix := ""
	if len(args) == 1 {
		prefix = args[0]
	}
	return listMethods(prefix, fullNames)
}

func listMethods(prefix string, fullNames bool) error {
	listedThings := map[string]bool{}
	if prefix == "" {
		for methodName := range commands.AllMethods {
			if fullNames {
				listedThings[methodName] = true
			} else {
				methodPrefix := strings.SplitN(methodName, ".", 2)[0]
				listedThings[methodPrefix] = true
			}
		}
	} else {
		for methodName := range commands.AllMethods {
			if methodName == prefix {
				listedThings[methodName] = true
				break
			}
			if strings.HasPrefix(methodName, prefix+".") {
				firstUntypedDot := strings.Index(methodName[len(prefix)+1:], ".")
				if !fullNames && firstUntypedDot != -1 {
					listedThings[methodName[:firstUntypedDot+len(prefix)+1]] = true
				} else {
					listedThings[methodName] = true
				}
			}
		}
	}

	if len(listedThings) == 0 {
		return fmt.Errorf("no listing for %q", prefix)
	}

	var listedThingSlice []string
	for thing := range listedThings {
		listedThingSlice = append(listedThingSlice, thing)
	}

	sort.Sort(sort.StringSlice(listedThingSlice))
	for _, thing := range listedThingSlice {
		fmt.Println(thing)
	}

	return nil
}
