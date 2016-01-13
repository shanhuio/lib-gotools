package main

import (
	"encoding/json"
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

	hash, err := gorepo.GitCommitHash(*projPath)
	if err != nil {
		log.Print(err)
	}
	log.Print(hash)

	b, errs := gorepo.Build(*projPath)
	for _, err := range errs {
		log.Print(err)
	}
	if len(errs) > 0 {
		log.Fatal("building failed")
	}

	enc := json.NewEncoder(os.Stdout)
	if err := enc.Encode(b.Struct); err != nil {
		log.Fatal(err)
	}
}
