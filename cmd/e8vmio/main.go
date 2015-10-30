package main

import (
	"log"
	"net/http"

	"e8vm.io/tools/goimp"
)

func main() {
	for path, repo := range map[string]string{
		"e8vm.io/e8vm":  "https://github.com/e8vm/e8vm",
		"e8vm.io/tools": "https://github.com/e8vm/tools",
	} {
		http.Handle(path, goimp.NewGitRepo(path, repo))
	}

	log.Fatal(http.ListenAndServe(":3000", nil))
}
