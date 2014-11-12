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

package commands_util

import (
	"errors"
	"fmt"
	"strings"
)

var (
	NotFlagError = errors.New("args don't start with a flag")
)

/*
Get a flag-value pair from the next one or two arguments, and return the args
that are left over. One arg is consumed if it is of the form "--flag=value",
two if it is of the form "--flag value".
*/
func extractFlagValue(args []string) (remainingArgs []string, flag, value string, err error) {
	if len(args) == 0 {
		panic("empty args")
	}
	if args[0] == "--" {
		panic("trying to extract flag from remainder")
	}
	if !strings.HasPrefix(args[0], "--") {
		err = NotFlagError
		return
	}
	first_arg := args[0][2:]
	if equals_index := strings.Index(first_arg, "="); equals_index != -1 {
		flag = first_arg[:equals_index]
		value = first_arg[equals_index+1:]
		remainingArgs = args[1:]
		return
	}
	if len(args) < 2 {
		err = fmt.Errorf("unspecified flag value: %q", args[0])
		return
	}
	flag = first_arg
	value = args[1]
	remainingArgs = args[2:]
	return
}

/*
Get all flag-value pairs from this args list, and return a new slice with the
args that are left over.
*/
func ExtractFlagValues(args []string) (remainingArgs []string, values map[string]string, err error) {
	values = map[string]string{}
	for len(args) != 0 {
		if args[0] == "--" {
			// no more processing flags, the remaining args are positional now.
			remainingArgs = append(remainingArgs, args[1:]...)
			return
		}

		var leftOver []string
		var f, v string
		leftOver, f, v, err = extractFlagValue(args)
		if err == NotFlagError {
			remainingArgs = append(remainingArgs, args[0])
			args = args[1:]
			err = nil
			continue
		}
		if err != nil {
			return
		}

		values[f] = v
		args = leftOver
	}

	return
}
