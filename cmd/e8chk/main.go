package main

import (
	"flag"
	"fmt"
	"os"

	"shanhu.io/smlvm/dagvis"
	"shanhu.io/smlvm/lexing"
	"shanhu.io/smlvm/textbox"
	"shanhu.io/tools/godep"
	"shanhu.io/tools/goload"
)

func errExit(e error) {
	if e == nil {
		return
	}
	fmt.Fprintln(os.Stderr, e)
	os.Exit(-1)
}

func checkRectLoaded(prog *goload.Program, h, w int) []*lexing.Error {
	errs := lexing.NewErrorList()

	fset := prog.Fset
	for _, p := range prog.Pkgs {
		pinfo := prog.Imported[p]
		for _, astf := range pinfo.Files {
			tokFile := fset.File(astf.Pos())
			name := tokFile.Name()
			fin, e := os.Open(name)
			if lexing.LogError(errs, e) {
				continue
			}

			errs.AddAll(textbox.CheckRect(name, fin, h, w))
			if lexing.LogError(errs, fin.Close()) {
				continue
			}
		}
	}

	return errs.Errs()
}

func main() {
	path := flag.String("path", "e8vm.io/e8vm", "go pkg path to check")
	textHeight := flag.Int("height", 300, "maximum height for a single file")
	textWidth := flag.Int("width", 80, "maximum width for a single file")

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
