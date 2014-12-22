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
	"github.com/GoogleCloudPlatform/gcloud/gcloud_apis/discovery_docs"
	"strings"
	"testing"
)

var simpleDoc = `
{
  "discoveryVersion": "v1"
}
`

func TestSimpleLoad(t *testing.T) {
	doc, err := LoadDiscovery(strings.NewReader(simpleDoc))
	if err != nil {
		t.Errorf(err.Error())
	}
	if doc.DiscoveryVersion != "v1" {
		t.Errorf("did not get version, expected %q, got %q", "v1", doc.DiscoveryVersion)
	}
}

func TestLoadAutoscalerV1beta2(t *testing.T) {
	doc, err := LoadDiscovery(strings.NewReader(discovery_docs.AutoscalerV1beta2))
	if err != nil {
		t.Errorf(err.Error())
	}
	if doc.DiscoveryVersion != "v1" {
		t.Errorf("did not get version, expected %q, got %q", "v1", doc.DiscoveryVersion)
	}
}
