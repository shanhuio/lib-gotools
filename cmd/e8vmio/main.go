package main

import (
	"flag"
	"log"
	"net/http"

	"e8vm.io/tools/goimp"
)

var (
	addr = flag.String("addr", "localhost:3000", "address to listen")
	dom  = flag.String("domain", "e8vm.io", "the domain")
)

func main() {
	flag.Parse()

	route := func(sub string) string { return "/" + sub + "/" }
	path := func(sub string) string { return *dom + "/" + sub }

	for sub, repo := range map[string]string{
		"e8vm":   "https://github.com/e8vm/e8vm",
		"tools":  "https://github.com/e8vm/tools",
		"shanhu": "https://github.com/e8vm/shanhu",
	} {
		http.Handle(route(sub), goimp.NewGitRepo(path(sub), repo))
	}

	log.Printf("serve at %s", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
