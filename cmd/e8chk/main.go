package main

import (
	"flag"
	"fmt"
	"os"

	"e8vm.io/e8vm/dagvis"
	"e8vm.io/e8vm/textbox"
	"e8vm.io/tools/godep"
	"e8vm.io/tools/goload"
)

func errExit(e error) {
	if e == nil {
		return
	}
	fmt.Fprintln(os.Stderr, e)
	os.Exit(-1)
}

func main() {
	path := flag.String("path", "e8vm.io/e8vm", "go pkg path to check")
	textHeight := flag.Int("height", 300, "maximum height for a single file")
	textWidth := flag.Int("width", 78, "maximum width for a single file")

	flag.Parse()

	pkgs, e := goload.ListPkgs(*path)
	errExit(e)

	prog, e := goload.Pkgs(pkgs)
	errExit(e)

	es := textbox.CheckRectLoaded(prog, *textHeight, *textWidth)
	if es != nil {
		for _, e := range es {
			fmt.Fprintln(os.Stderr, e)
		}
		os.Exit(-1)
	}

	fileDeps := godep.FileDepLoaded(prog)

	for p, g := range fileDeps {
		_, e = dagvis.IsDAG(g)
		if e != nil {
			errExit(fmt.Errorf("%s: %s", p, e))
		}
	}
}
