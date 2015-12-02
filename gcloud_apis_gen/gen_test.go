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
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestGeneratedCodeCompiles(t *testing.T) {
	workspace := os.TempDir()
	commands_tmp := filepath.Join(workspace, "commands")
	clients_tmp := filepath.Join(workspace, "clients")
	err := os.MkdirAll(commands_tmp, 0777)
	if err != nil {
		t.Error(err)
	}
	err = os.MkdirAll(clients_tmp, 0777)
	if err != nil {
		t.Error(err)
	}
	err = exec.Command(
		"gcloud_apis_gen",
		"--discovery-dir", "../discovery_docs/",
		"--clients-dir", clients_tmp,
		"--commands-dir", commands_tmp).Run()
	if err != nil {
		t.Error(err)
	}
	cmd := exec.Command(
		"go", "build", ".")
	cmd.Dir = commands_tmp
	err = cmd.Run()
	if err != nil {
		t.Error(err)
	}
}
