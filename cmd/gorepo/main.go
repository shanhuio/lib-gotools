package main

import (
	"flag"
	"log"
	"os"

	"e8vm.io/tools/gorepo"
)

var (
	projPath = flag.String("proj", "e8vm.io/e8vm", "the project path")
)

func main() {
	flag.Parse()

	hash, err := gorepo.GitCommit(*projPath)
	if err != nil {
		log.Print(err)
	}
	log.Print(hash)

	errs := gorepo.Build(*projPath, os.Stdout)
	for _, err := range errs {
		log.Print(err)
	}
}
