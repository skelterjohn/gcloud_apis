package main

import (
	"flag"
	"fmt"
	"github.com/googlecloudplatform/gcloud/gcloud_apis/gcloud_apis_gen/discovery"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func funcNameForMethod(m discovery.Method, v string) (name string) {
	parts := strings.Split(m.Id, ".")
	for _, part := range parts {
		name += strings.ToUpper(part[:1]) + part[1:]
	}
	name += "_" + v
	return
}

/*
Write the source code for the commands. This includes a single "commands.go",
which makes a mapping from method id to command func, and a file for every
discovery doc in the form of "API_VERSION.go".
*/
func writeCommandsSource(commands_dir string, docs map[string]*discovery.Discovery) (err error) {
	commandsFile, err := os.Create(filepath.Join(commands_dir, "commands.go"))
	if err != nil {
		log.Fatal(err)
	}

	methodFuncs := map[string]string{}
	methods := map[string]discovery.Method{}
	for _, doc := range docs {
		docMethods := map[string]discovery.Method{}
		for _, resource := range doc.Resources {
			for _, method := range resource.Methods {
				funcName := funcNameForMethod(method, doc.Version)
				methodFuncs[method.Id+"#"+doc.Version] = funcName
				methods[funcName] = method
				docMethods[funcName] = method
			}
		}

		apiCommandsPath := filepath.Join(commands_dir, fmt.Sprintf("%s_%s.go", doc.Name, doc.Version))
		var apiCommandsFile *os.File
		apiCommandsFile, err = os.Create(apiCommandsPath)

		commandsSourceTemplate.ExecuteTemplate(apiCommandsFile, "api_commands", docMethods)
	}
	_ = commandsFile
	commandsSourceTemplate.ExecuteTemplate(commandsFile, "all_commands", methodFuncs)
	return
}
func main() {

	var discovery_dir, client_dir, commands_dir string

	fs := flag.NewFlagSet("gcloud_apis_gen", flag.ExitOnError)
	fs.StringVar(&discovery_dir, "discovery-dir", "", "directory to find discovery docs")
	fs.StringVar(&client_dir, "clients-dir", "", "directory to write client source")
	fs.StringVar(&commands_dir, "commands-dir", "", "directory to write command source")
	err := fs.Parse(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	docs, err := discovery.LoadDiscoveriesFromDirectory(discovery_dir)
	if err != nil {
		log.Fatal(err)
	}

	err = writeCommandsSource(commands_dir, docs)
	if err != nil {
		log.Fatal(err)
	}

	// Ensure that all the generated code is property go fmt ed.
	cmd := exec.Command("go", "fmt", ".")
	cmd.Dir = commands_dir
	err = cmd.Run()
	if err != nil {
		log.Printf("while formatting in %s: %v", cmd.Dir, err)
	}
	cmd = exec.Command("go", "fmt", ".")
	cmd.Dir = client_dir
	err = cmd.Run()
	if err != nil {
		log.Printf("while formatting in %s: %v", cmd.Dir, err)
	}
}
