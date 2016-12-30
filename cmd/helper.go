package cmd

import (
	"log"

	"github.com/schreibe72/docker-registry-client/registry"
)

var hub *registry.Registry

func connect() {
	if hub == nil {
		var err error
		logger := registry.Log
		if !Verbose {
			logger = registry.Quiet
		}
		hub, err = registry.New(Server, "", "", logger)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// func main() {
// 	log.SetOutput(ioutil.Discard)
// 	url := "https://docker.sueddeutsche.de/"
// 	username := "" // anonymous
// 	password := "" // anonymous
// 	hub, err := registry.New(url, username, password)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	repositories, err := hub.Repositories()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	for _, repo := range repositories {
// 		fmt.Printf("%s\n", repo)
// 		tags, err := hub.Tags(repo)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		for _, tag := range tags {
// 			digest, err := hub.ManifestDigest(repo, tag)
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 			fmt.Printf("       %s:%s %s\n", repo, tag, digest)
// 		}
// 	}
// }
