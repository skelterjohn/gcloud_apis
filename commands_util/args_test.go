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
	"reflect"
	"testing"
)

func considerCase(t *testing.T, foo func() map[string]interface{}, input interface{}, expected map[string]interface{}) {
	errf := func(k string, v, e interface{}) {
		t.Errorf("for %q, got %s %q, expected %q", input, k, v, e)
	}

	var result map[string]interface{}

	if expectedPanicMsg, ok := expected["panic"]; ok {
		defer func() {
			panicMsg := recover()
			if panicMsg == nil {
				errf("panic", panicMsg, expectedPanicMsg)
				return
			}
			if !reflect.DeepEqual(panicMsg, expectedPanicMsg) {
				errf("panic", panicMsg, expectedPanicMsg)
				return
			}
		}()
	}
	result = foo()
	if len(result) != len(expected) {
		for key, expectedValue := range expected {
			value, ok := result[key]
			if !ok {
				errf(key, value, expectedValue)
			}
		}
	}
	for key, value := range result {
		expectedValue, ok := expected[key]
		if !ok || !reflect.DeepEqual(value, expectedValue) {
			errf(key, value, expectedValue)
		}
	}
}

func TestExtractFlagValue(t *testing.T) {

	type Case struct {
		Args     []string
		Expected map[string]interface{}
	}

	cases := []Case{
		{
			Args: []string{"--foo", "bar", "baz"},
			Expected: map[string]interface{}{
				"flag":  "foo",
				"value": "bar",
				"args":  []string{"baz"},
			},
		},
		{
			Args: []string{"--foo=bar", "baz"},
			Expected: map[string]interface{}{
				"flag":  "foo",
				"value": "bar",
				"args":  []string{"baz"},
			},
		},
		{
			Args: []string{"--foo"},
			Expected: map[string]interface{}{
				"err": "unspecified flag value: \"--foo\"",
			},
		},
		{
			Args: []string{"x", "--foo"},
			Expected: map[string]interface{}{
				"err": "args don't start with a flag",
			},
		},
		{
			Args: []string{},
			Expected: map[string]interface{}{
				"panic": "empty args",
			},
		},
	}
	for _, fvt := range cases {
		foo := func() map[string]interface{} {
			args, flag, value, err := extractFlagValue(fvt.Args)
			result := map[string]interface{}{}
			if len(args) != 0 {
				result["args"] = args
			}
			if flag != "" {
				result["flag"] = flag
			}
			if value != "" {
				result["value"] = value
			}
			if err != nil {
				result["err"] = err.Error()
			}
			return result
		}
		considerCase(t, foo, fvt.Args, fvt.Expected)
	}
}

func TestExtractFlagValues(t *testing.T) {
	type TestCase struct {
		Args           []string
		Expected       map[string]interface{}
		ExpectedValues map[string]string
		ExpectedArgs   []string
		ExpectedErr    error
		ExpectedPanic  string
	}

	cases := []TestCase{
		{
			Args: []string{"--foo", "bar", "baz"},
			Expected: map[string]interface{}{
				"values": map[string]string{
					"foo": "bar",
				},
				"args": []string{"baz"},
			},
		},
		{
			Args: []string{"baz", "--foo", "bar"},
			Expected: map[string]interface{}{
				"values": map[string]string{
					"foo": "bar",
				},
				"args": []string{"baz"},
			},
		},
		{
			Args: []string{"--foo", "bar", "--baz", "qux"},
			Expected: map[string]interface{}{
				"values": map[string]string{
					"foo": "bar",
					"baz": "qux",
				},
			},
		},
		{
			Args: []string{"--foo", "bar", "--", "--baz", "qux"},
			Expected: map[string]interface{}{
				"values": map[string]string{
					"foo": "bar",
				},
				"args": []string{"--baz", "qux"},
			},
		},
		{
			Args: []string{"--foo", "bar", "--"},
			Expected: map[string]interface{}{
				"values": map[string]string{
					"foo": "bar",
				},
			},
		},
		{
			Args: []string{"x", "--foo", "bar", "y", "--", "--baz", "qux"},
			Expected: map[string]interface{}{
				"values": map[string]string{
					"foo": "bar",
				},
				"args": []string{"x", "y", "--baz", "qux"},
			},
		},
	}

	for _, fvt := range cases {
		foo := func() map[string]interface{} {
			args, values, err := ExtractFlagValues(fvt.Args)
			result := map[string]interface{}{}
			if len(args) != 0 {
				result["args"] = args
			}
			if len(values) != 0 {
				result["values"] = values
			}
			if err != nil {
				result["err"] = err.Error()
			}
			return result
		}
		considerCase(t, foo, fvt.Args, fvt.Expected)
	}
}
