package main

import (
	"flag"
	"go/build"
	"log"
	"net/http"
	"path/filepath"
)

func pkgSrcDir() (string, error) {
	const pkgPath = "e8vm.io/tools/cmd/e8doc"
	pkg, e := build.Import(pkgPath, "", build.FindOnly)
	if e != nil {
		return "", e
	}

	return pkg.Dir, nil
}

func pkgSubDir(sub string) (string, error) {
	p, e := pkgSrcDir()
	if e != nil {
		return "", e
	}
	return filepath.Join(p, sub), nil
}

var (
	addr    = flag.String("addr", "localhost:8888", "server listen address")
	pkgPath = flag.String("pkg", "e8vm.io/e8vm", "package path to plot")
)

func main() {
	wwwRoot, e := pkgSubDir("www")
	fatale(e)

	flag.Parse()

	proj := newProject("/p", *pkgPath)
	hook := newHook(proj)

	http.Handle("/p/", proj)
	http.Handle("/h/", hook)
	http.Handle("/", http.FileServer(http.Dir(wwwRoot)))

	log.Printf("source loaded. serving at %s.", *addr)
	for {
		fatale(http.ListenAndServe(*addr, nil))
	}
}
