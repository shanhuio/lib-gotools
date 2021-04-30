package gocheck

import (
	"go/ast"
	"go/build"
	"go/parser"
	"go/token"
	"go/types"
	"path/filepath"
	"strings"

	"shanhu.io/gcimporter"
	"shanhu.io/smlvm/dagvis"
	"shanhu.io/smlvm/lexing"
)

type checker struct {
	path string

	ctx      *build.Context
	buildPkg *build.Package

	fset  *token.FileSet
	alias *gcimporter.AliasMap
}

func trimSuffix(name string) string {
	return strings.TrimSuffix(name, ".go")
}

func newCheckerPath(
	ctx *build.Context, path string, alias *gcimporter.AliasMap,
) (*checker, error) {
	if alias != nil {
		path = alias.Map(path)
	}
	pkg, err := ctx.Import(path, "", 0)
	if err != nil {
		return nil, err
	}

	return newChecker(ctx, pkg, alias), nil
}

func newChecker(
	ctx *build.Context, pkg *build.Package, alias *gcimporter.AliasMap,
) *checker {
	fset := token.NewFileSet()
	return &checker{
		ctx:      ctx,
		path:     pkg.ImportPath,
		buildPkg: pkg,
		fset:     fset,
		alias:    alias,
	}
}

func (c *checker) listFiles() ([]*ast.File, error) {
	var srcFiles []string
	srcFiles = append(srcFiles, c.buildPkg.GoFiles...)
	srcFiles = append(srcFiles, c.buildPkg.CgoFiles...)

	var files []*ast.File
	for _, baseName := range srcFiles {
		filename := filepath.Join(c.buildPkg.Dir, baseName)
		f, err := parser.ParseFile(c.fset, filename, nil, 0)
		if err != nil {
			return nil, err
		}
		files = append(files, f)
	}

	return files, nil
}

func (c *checker) typesCheck(files []*ast.File) (
	*types.Info, *types.Package, error,
) {
	config := &types.Config{
		Importer:    gcimporter.New(c.ctx, c.alias),
		FakeImportC: true,
	}
	info := &types.Info{
		Uses: make(map[*ast.Ident]types.Object),
	}

	typesPkg, err := config.Check(c.path, c.fset, files, info)
	if err != nil {
		return nil, nil, err
	}
	return info, typesPkg, nil
}

func (c *checker) depGraph(files []*ast.File) (*dagvis.Graph, error) {
	info, typesPkg, err := c.typesCheck(files)
	if err != nil {
		return nil, err
	}

	in := &depGraphInput{
		fset:  c.fset,
		files: files,
		info:  info,
		pkg:   typesPkg,
	}
	return pkgDepGraph(in)
}

// CheckAll checks everything for a package.
func CheckAll(path string, h, w int) []*lexing.Error {
	c, err := newCheckerPath(&build.Default, path, nil)
	if err != nil {
		return lexing.SingleErr(err)
	}

	files, err := c.listFiles()
	if err != nil {
		return lexing.SingleErr(err)
	}

	g, err := c.depGraph(files)
	if err != nil {
		return lexing.SingleErr(err)
	}

	if err := dagvis.CheckDAG(g); err != nil {
		return lexing.SingleErr(err)
	}

	names := listFileNames(c.fset, files)
	if errs := CheckRect(names, h, w); errs != nil {
		return errs
	}

	return CheckLineComment(c.fset, files)
}
