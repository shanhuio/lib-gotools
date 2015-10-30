package main

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"sync"
)

type memWeb struct {
	pages map[string][]byte

	repath func(p string) string

	lock sync.RWMutex
}

func newMemWeb() *memWeb {
	ret := new(memWeb)
	ret.pages = make(map[string][]byte)
	return ret
}

func (mw *memWeb) clear() {
	// create a new map, and defer the actual cleaning to the
	// garbage collector
	mw.pages = make(map[string][]byte)
}

func (mw *memWeb) setPage(p string, content []byte) {
	p = path.Clean(p)

	if content == nil {
		delete(mw.pages, p)
	} else {
		mw.pages[p] = content
	}
}

var _ http.Handler = new(memWeb)

// ServeHTTP makes memWeb an http handler.
func (mw *memWeb) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	p := req.URL.Path
	if mw.repath != nil {
		p = mw.repath(p) // transforms the path
	}

	mw.lock.RLock()
	defer mw.lock.RUnlock()

	ew := &errWriter{w: w}
	bs, found := mw.pages[p]
	if !found {
		w.WriteHeader(404)
		fmt.Fprintf(ew, "page at %q not exists", p)
	} else {
		ew.Write(bs)
	}

	if ew.e != nil {
		log.Print(ew.e)
	}
}
