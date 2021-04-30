package gocheck

import (
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/packages"
	"shanhu.io/misc/errcode"
	"shanhu.io/smlvm/dagvis"
	"shanhu.io/smlvm/lexing"
)

type modChecker struct {
	path string
	pkg  *packages.Package
	fset *token.FileSet
}

func newModChecker(pkg *packages.Package) *modChecker {
	return &modChecker{
		path: pkg.PkgPath,
		pkg:  pkg,
		fset: pkg.Fset,
	}
}

func (c *modChecker) depGraph(files []*ast.File) (*dagvis.Graph, error) {
	in := &depGraphInput{
		fset:  c.fset,
		files: files,
		info:  c.pkg.TypesInfo,
		pkg:   c.pkg.Types,
	}
	return pkgDepGraph(in)
}

func ModCheckAll(dir, pkg string, h, w int) []*lexing.Error {
	var loadMode packages.LoadMode
	for _, m := range []packages.LoadMode{
		packages.NeedTypes,
		packages.NeedFiles,
		packages.NeedTypesInfo,
		packages.NeedModule,
		packages.NeedSyntax,
	} {
		loadMode |= m
	}

	fset := token.NewFileSet()
	config := &packages.Config{
		Mode: loadMode,
		Dir:  dir,
		Fset: fset,
	}
	pkgs, err := packages.Load(config, pkg)
	if err != nil {
		return lexing.SingleErr(err)
	}

	if len(pkgs) != 1 {
		err := errcode.Internalf("got %d packages", len(pkgs))
		return lexing.SingleErr(err)
	}

	c := newModChecker(pkgs[0])

	files := c.pkg.Syntax

	g, err := c.depGraph(files)
	if err != nil {
		return lexing.SingleErr(err)
	}

	if err := dagvis.CheckDAG(g); err != nil {
		return lexing.SingleErr(err)
	}

	names := listFileNames(fset, files)
	if errs := CheckRect(names, h, w); errs != nil {
		return errs
	}

	return CheckLineComment(c.fset, files)
}
