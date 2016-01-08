package main

import (
	"flag"
	"log"
	"net/http"

	"e8vm.io/tools/shanhu"
)

var (
	config = flag.String("config", "shanhu.config", "path to config file")
	addr   = flag.String("addr", "localhost:3355", "address to serve on")
)

func ne(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	flag.Parse()
	c, err := shanhu.LoadConfig(*config)
	ne(err)
	h, err := shanhu.NewHandler(c)
	ne(err)

	http.Handle("/", h)
	log.Printf("server at %s", *addr)
	err = http.ListenAndServe(*addr, nil)
	ne(err)
}
