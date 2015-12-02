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

package discovery

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func LoadDiscovery(r io.Reader) (doc *Discovery, err error) {
	doc = &Discovery{}
	dec := json.NewDecoder(r)
	err = dec.Decode(doc)
	return
}

func LoadDiscoveriesFromDirectory(dir string) (docs map[string]*Discovery, err error) {
	docs = make(map[string]*Discovery)

	dir_file, err := os.Open(dir)
	if err != nil {
		return
	}
	names, err := dir_file.Readdirnames(-1)
	if err != nil {
		return
	}
	for _, name := range names {
		if !strings.HasSuffix(name, ".json") {
			continue
		}
		var disc_file *os.File
		disc_file, err = os.Open(filepath.Join(dir, name))
		if err != nil {
			return
		}
		var doc *Discovery
		doc, err = LoadDiscovery(disc_file)
		if err != nil {
			return
		}
		docs[fmt.Sprintf("%s_%s", doc.Name, doc.Version)] = doc
	}
	return
}

type Discovery struct {
	Kind             string
	DiscoveryVersion string
	Id               string
	Name             string
	Version          string
	Revision         string
	Title            string
	Description      string
	Icons            struct {
		X16 string
		X32 string
	}
	DocumentationLink string
	Labels            []string
	Protocol          string // must be "rest"
	BaseUrl           string
	BasePath          string
	RootUrl           string
	ServicePath       string
	BatchPath         string // must be "batch"
	Parameters        map[string]Parameter
	Auth              struct {
		Oauth2 struct {
			Scopes map[string]struct {
				Description string
			}
		}
	}
	Features  []string
	Schemas   map[string]Parameter
	Methods   map[string]Method
	Resources map[string]Resource
}

type Resource struct {
	Methods map[string]Method
	Resources map[string]Resource
}

type JsonSchema interface{}

type Method struct {
	Id             string
	Path           string
    FlatPath       string
	HttpMethod     string
	Parameters     map[string]Parameter
	ParameterOrder []string
	Request        struct {
		Ref string `json:"$ref"`
	}
	Response struct {
		Ref string `json:"$ref"`
	}
	Scopes                []string
	SupportsMediaDownload bool
	SupportsMediaUpload   bool
	MediaUpload           struct {
		Accept    []string
		MaxSize   string
		Protocols struct {
			Simple struct {
				Multipart bool // must be true
				Path      string
			}
			Resumable struct {
				Multipart bool // must be true
				Path      string
			}
		}
	}
	SupportsSubscription bool
}

type Parameter struct {
	Id                   string
	Type                 string
	Ref                  string `json:"$ref"`
	Description          string
	Default              string
	Required             bool
	Format               string
	Pattern              string
	Minimum              string
	Maximum              string
	Enum                 []string
	EnumDescriptions     []string
	Repeated             bool
	Location             string
	Properties           map[string]JsonSchema
	AdditionalProperties JsonSchema
	Items                JsonSchema
	Annotations          struct {
		Required []string
	}
}
