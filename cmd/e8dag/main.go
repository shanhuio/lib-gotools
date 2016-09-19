package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"shanhu.io/smlvm/dagvis"
	"shanhu.io/tools/godep"
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

func main() {
	std, e := godep.StdDep()
	ne(e)
	saveLayout(std, "dags/gostd.json")

	pkgs := []string{
		"e8vm.io/tools/dagvis",
		"e8vm.io/tools/godep",
		"e8vm.io/tools/goview",

		"e8vm.io/e8vm/arch8",
		"e8vm.io/e8vm/asm8",
		"e8vm.io/e8vm/asm8/ast",
		"e8vm.io/e8vm/asm8/parse",
		"e8vm.io/e8vm/build8",
		"e8vm.io/e8vm/link8",
		"e8vm.io/e8vm/dasm8",
		"e8vm.io/e8vm/lex8",
		"e8vm.io/e8vm/cmd/e8",
	}

	dags, e := godep.FileDep(pkgs)
	ne(e)
	for p, g := range dags {
		p = strings.TrimPrefix(p, "e8vm.io/")
		p = strings.Replace(p, "/", "_", -1)
		p = "dags/" + p + ".json"

		saveLayout(g, p)
	}
}
