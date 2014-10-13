package main

import (
	"flag"
	"fmt"
	"github.com/googlecloudplatform/gcloud/gcloud_apis/gcloud_apis_gen/discovery"
	"log"
	"os"
)

func main() {

	fs := flag.NewFlagSet("gcloud_apis_gen", flag.ExitOnError)
	discovery_dir := fs.String("discovery", "", "directory containing discovery docs")
	err := fs.Parse(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	docs, err := discovery.LoadDiscoveriesFromDirectory(*discovery_dir)
	if err != nil {
		log.Fatal(err)
	}
	for name := range docs {
		fmt.Println(name)
	}
}
