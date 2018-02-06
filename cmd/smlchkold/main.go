package main

import (
	"flag"
	"fmt"
	"go/scanner"
	"go/token"
	"io/ioutil"
	"os"
	"strings"

	"shanhu.io/smlvm/dagvis"
	"shanhu.io/smlvm/lexing"
	"shanhu.io/smlvm/textbox"

	"shanhu.io/tools/godep"
	"shanhu.io/tools/goload"
)

func errExit(err error) {
	if err == nil {
		return
	}
	fmt.Fprintln(os.Stderr, err)
	os.Exit(-1)
}

func errsExit(errs []*lexing.Error) {
	if len(errs) == 0 {
		return
	}
	for _, err := range errs {
		fmt.Fprintln(os.Stderr, err)
	}
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
			fin, err := os.Open(name)
			if lexing.LogError(errs, err) {
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

func validLineCommentContent(s string) bool {
	if s == "" {
		return true
	}
	if strings.HasPrefix(s, " ") {
		return true
	}
	if strings.HasPrefix(s, "\t") {
		return true
	}
	return false
}

func toLexingPos(p token.Position) *lexing.Pos {
	return &lexing.Pos{
		File: p.Filename,
		Line: p.Line,
		Col:  p.Column,
	}
}

func tokenPos(fset *token.FileSet, pos token.Pos) *lexing.Pos {
	return toLexingPos(fset.Position(pos))
}

func checkLineComment(prog *goload.Program) []*lexing.Error {
	errs := lexing.NewErrorList()

	fset := prog.Fset
	errHandler := func(pos token.Position, msg string) {
		errs.Errorf(toLexingPos(pos), "%s", msg)
	}
	for _, p := range prog.Pkgs {
		pinfo := prog.Imported[p]
		for _, astf := range pinfo.Files {
			tokFile := fset.File(astf.Pos())
			s := new(scanner.Scanner)

			bs, err := ioutil.ReadFile(tokFile.Name())
			if lexing.LogError(errs, err) {
				continue
			}

			s.Init(tokFile, bs, errHandler, scanner.ScanComments)
			for {
				pos, tok, lit := s.Scan()
				if tok == token.EOF {
					break
				}
				if tok == token.COMMENT && strings.HasPrefix(lit, "//") {
					if !validLineCommentContent(lit[2:]) {
						errs.Errorf(
							tokenPos(fset, pos),
							"please add a space to comment %q", lit,
						)
					}
				}
			}
		}
	}
	return errs.Errs()
}

func main() {
	path := flag.String("path", "shanhu.io/smlvm", "go pkg path to check")
	textHeight := flag.Int("height", 300, "maximum height for a single file")
	textWidth := flag.Int("width", 80, "maximum width for a single file")

	flag.Parse()

	pkgs, err := goload.ListPkgs(*path)
	errExit(err)

	prog, err := goload.Pkgs(pkgs)
	errExit(err)

	fileDeps := godep.FileDepLoaded(prog)
	for p, g := range fileDeps {
		if _, err := dagvis.IsDAG(g); err != nil {
			errExit(fmt.Errorf("%s: %s", p, err))
		}
	}

	errs := checkRectLoaded(prog, *textHeight, *textWidth)
	errsExit(errs)

	errs = checkLineComment(prog)
	errsExit(errs)
}
