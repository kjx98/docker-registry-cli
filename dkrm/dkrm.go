package main

import (
	"fmt"
	"github.com/kjx98/docker-registry-client/registry"
	"os"
)

func main() {
	hub, err := registry.New("http://kjx27", "", "")
	if err != nil {
		panic(err)
	}
	repos, err := hub.Repositories()
	if err != nil {
		panic("list repos fail")
	}
	if len(repos) == 0 {
		return
	}
	fmt.Println("reposities:")
	for _, ss := range repos {
		fmt.Println("\t", ss)
	}
	if len(os.Args) == 1 {
		fmt.Println("no repo to remove")
		return
	}
	fmt.Printf("Try rm %s repo\n", os.Args[1])
	digest, err := hub.ManifestDigest(os.Args[1], "latest")
	if err != nil {
		fmt.Printf("Can't get digest for %s\n", os.Args[1])
		return
	}
	err = hub.DeleteManifest(os.Args[1], digest)
	fmt.Println("result: ", err)
	return
}
