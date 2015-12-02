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
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"

	"github.com/skelterjohn/overwrite"
)

func PopulateRequestFromFilename(requestObject interface{}, filename string) error {
	var r io.Reader
	var err error
	if filename == "-" {
		r = os.Stdin
	} else {
		r, err = os.Open(filename)
		if err != nil {
			return err
		}
	}
	return json.NewDecoder(r).Decode(requestObject)
}

func OverwriteRequestWithValues(requestObject interface{}, keyValues map[string]string) error {
	for key, val := range keyValues {
		err := overwrite.Overwrite(requestObject, key, val)
		if err != nil {
			return fmt.Errorf("for key \"--%s\" while evaluating %s: %s", key, val, err)
		}
	}
	return nil
}

func ErrForWrongParams(expectedParams, paramValues, args []string) error {
	value := "value"
	if len(expectedParams) > 1 {
		value += "s"
	}
	expectedArgument := strings.ToUpper(strings.Join(expectedParams, "/"))
	return fmt.Errorf(
		"expected %d %s for %s, but %q has %d.",
		len(expectedParams), value, expectedArgument, args[0], len(paramValues))
}

func ConvertValue_string(s string) (string, error) {
	return s, nil
}

func ConvertValue_int64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

func ConvertValue_uint64(s string) (uint64, error) {
	return strconv.ParseUint(s, 10, 64)
}

func ConvertValue_int32(s string) (int32, error) {
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return int32(v), err
}

func ConvertValue_uint32(s string) (uint32, error) {
	v, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint32(v), err
}

func ConvertValue_float64(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

func ConvertValue_bool(s string) (bool, error) {
	return strconv.ParseBool(s)
}

func PrintResponse(response interface{}) error {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(response); err != nil {
		return err
	}
	var indentedBuf bytes.Buffer
	if err := json.Indent(&indentedBuf, buf.Bytes(), "", "  "); err != nil {
		return err
	}
	_, err := fmt.Println(&indentedBuf)
	return err
}

/*
Get each of the tokens from a joined arg. Tokens are considered to be joined
by whatever the first character is. If the first character is in [a-zA-Z0-9],
a slash is assumed.
*/
func SplitParamValues(arg string) (params []string) {
	if len(arg) == 0 {
		return []string{}
	}
	sep := "/"
	firstLetter := rune(arg[0])
	if !(unicode.IsDigit(firstLetter) || unicode.IsLetter(firstLetter)) {
		sep = arg[0:1]
		arg = arg[1:]
	}
	return strings.Split(arg, sep)
}
