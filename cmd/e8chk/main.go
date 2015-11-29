package main

import (
	"flag"
	"fmt"
	"os"

	"e8vm.io/e8vm/dagvis"
	"e8vm.io/e8vm/lex8"
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

func checkRectLoaded(prog *goload.Program, h, w int) []*lex8.Error {
	errs := lex8.NewErrorList()

	fset := prog.Fset
	for _, p := range prog.Pkgs {
		pinfo := prog.Imported[p]
		for _, astf := range pinfo.Files {
			tokFile := fset.File(astf.Pos())
			name := tokFile.Name()
			fin, e := os.Open(name)
			if lex8.LogError(errs, e) {
				continue
			}

			errs.AddAll(textbox.CheckRect(name, fin, h, w))
			if lex8.LogError(errs, fin.Close()) {
				continue
			}
		}
	}

	return errs.Errs()
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

	es := checkRectLoaded(prog, *textHeight, *textWidth)
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
