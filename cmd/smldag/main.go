package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"shanhu.io/p/shanhu.io/smlvm/dagvis"

	"shanhu.io/p/shanhu.io/tools/godep"
	"shanhu.io/p/shanhu.io/tools/goload"
)

func ne(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}
}

func saveLayoutBytes(bs []byte, f string) {
	if strings.HasSuffix(f, ".js") {
		out, err := os.Create(f)
		ne(err)
		defer out.Close()

		_, err = io.WriteString(out, "var dag = ")
		ne(err)

		_, err = out.Write(bs)
		ne(err)

		_, err = io.WriteString(out, ";")
		ne(err)

		ne(out.Close())
		return
	}

	ne(ioutil.WriteFile(f, bs, 0644))
}

func saveLayout(g *dagvis.Graph, f string) {
	m, err := dagvis.LayoutJSON(g)
	ne(err)
	saveLayoutBytes(m, f)
}

func repoDep(repo string) (*dagvis.Graph, error) {
	if repo == "" {
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
