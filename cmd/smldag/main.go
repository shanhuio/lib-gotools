package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"shanhu.io/smlvm/dagvis"

	"shanhu.io/tools/godep"
	"shanhu.io/tools/goload"
)

func ne(e error) {
	if e != nil {
		fmt.Fprintln(os.Stderr, e)
		os.Exit(-1)
	}
}

func saveLayout(g *dagvis.Graph, f string) {
	m, e := dagvis.LayoutJSON(g)
	ne(e)

	e = ioutil.WriteFile(f, m, 0644)
	ne(e)
}

func saveRevLayout(g *dagvis.Graph, f string) {
	m, e := dagvis.RevLayoutJSON(g)
	ne(e)

	e = ioutil.WriteFile(f, m, 0644)
	ne(e)
}

func repoDep(repo string) (*dagvis.Graph, error) {
	if repo == "-" {
		return godep.StdDep()
	}

	pkgs, err := goload.ListPkgs(repo)
	if err != nil {
		return nil, err
	}
	g, err := godep.PkgDep(pkgs)
	if err != nil {
		return nil, err
	}

	repoSlash := repo + "/"
	return g.Rename(func(name string) (string, error) {
		if name == repo {
			return "~", nil
		}
		return strings.TrimPrefix(name, repoSlash), nil
	})
}

func main() {
	repo := flag.String("repo", "", "repository to generate the dependency map")
	out := flag.String("out", "godag.json", "output JSON file")
	flag.Parse()

	g, e := repoDep(*repo)
	ne(e)
	saveLayout(g, *out)
}
