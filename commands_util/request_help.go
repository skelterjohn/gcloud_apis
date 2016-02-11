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
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

func PrintRequestExample(req interface{}) {
	fmt.Fprintln(os.Stderr, "Request template (note that output-only fields are still listed):")
	populateObject(reflect.ValueOf(req))
	data, err := json.MarshalIndent(req, "\t", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "\tError generating JSON: %v", err)
		return
	}
	fmt.Fprintf(os.Stderr, "\t%s\n", data)
}

func populateObject(v reflect.Value) {
	switch v.Type().Kind() {
	case reflect.Ptr:
		populatePointer(v)
	case reflect.Struct:
		populateStruct(v)
	case reflect.Slice:
		populateSlice(v)
	case reflect.String:
		populateString(v)
	default:
		// zero is fine
	}
	// also other primitives
}

func populatePointer(pv reflect.Value) {
	if pv.IsNil() {
		pv.Set(reflect.New(pv.Type().Elem()))
	}
	populateObject(pv.Elem())
}

func populateStruct(sv reflect.Value) {
	for i := 0; i < sv.NumField(); i++ {
		populateObject(sv.Field(i))
	}
}

func populateSlice(sv reflect.Value) {
	sv.Set(reflect.MakeSlice(sv.Type(), 1, 1))
	populateObject(sv.Index(0))

}

func populateString(sv reflect.Value) {
	sv.SetString("STRING")
}
