package gocheck

import (
	"fmt"
	"go/ast"
	"go/build"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"path/filepath"
	"sort"
	"strings"

	"shanhu.io/smlvm/dagvis"
	"shanhu.io/smlvm/lexing"
)

type checker struct {
	path     string
	buildPkg *build.Package
	fset     *token.FileSet
}

func trimSuffix(name string) string {
	return strings.TrimSuffix(name, ".go")
}

func newChecker(path string) (*checker, error) {
	pkg, err := build.Import(path, "", 0)
	if err != nil {
		return nil, err
	}

	var srcFiles []string
	srcFiles = append(srcFiles, pkg.GoFiles...)
	srcFiles = append(srcFiles, pkg.CgoFiles...)

	fset := token.NewFileSet()
	return &checker{
		path:     path,
		buildPkg: pkg,
		fset:     fset,
	}, nil
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
		Importer: importer.Default(),
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

	depsMap := make(map[token.Pos]map[token.Pos]bool)
	for _, f := range files {
		depsMap[filePos(c.fset, f.Pos())] = make(map[token.Pos]bool)
	}

	for use, obj := range info.Uses {
		if obj.Pkg() != typesPkg {
			continue // ignore inter-pkg refs
		}

		fused := filePos(c.fset, use.NamePos)
		fdef := filePos(c.fset, obj.Pos())

		if fused == fdef {
			continue
		}

		if _, found := depsMap[fdef]; !found {
			panic(fmt.Errorf("%s not found in %s", use.Name, c.path))
		}
		depsMap[fdef][fused] = true
	}

	ret := make(map[string][]string)
	for f, deps := range depsMap {
		var lst []string
		for dep := range deps {
			lst = append(lst, trimBase(filename(c.fset, dep)))
		}
		sort.Strings(lst)
		ret[trimBase(filename(c.fset, f))] = lst
	}
	return &dagvis.Graph{Nodes: ret}, nil
}

func (c *checker) checkRect(files []*ast.File, h, w int) []*lexing.Error {
	var names []string
	for _, f := range files {
		tokFile := c.fset.File(f.Pos())
		name := tokFile.Name()
		names = append(names, name)
	}

	return CheckRect(names, h, w)
}

// CheckAll checks everything for a package.
func CheckAll(path string, h, w int) []*lexing.Error {
	c, err := newChecker(path)
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

	if _, err := dagvis.IsDAG(g); err != nil {
		return lexing.SingleErr(err)
	}

	if errs := c.checkRect(files, h, w); errs != nil {
		return errs
	}

	return CheckLineComment(c.fset, files)
}