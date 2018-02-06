package godep

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

	depsMap := make(map[string]map[string]bool)
	for _, f := range files {
		name := filename(c.fset, f.Pos())
		depsMap[name] = make(map[string]bool)
	}

	for use, obj := range info.Uses {
		if obj.Pkg() != typesPkg {
			continue // ignore inter-pkg refs
		}

		fused := filename(c.fset, use.NamePos)
		fdef := filename(c.fset, obj.Pos())

		if fused == fdef {
			continue
		}

		if _, found := depsMap[fdef]; !found {
			panic(fmt.Errorf("%s not found in %s", fdef, c.path))
		}
		depsMap[fdef][fused] = true
	}

	ret := make(map[string][]string)
	for f, deps := range depsMap {
		var lst []string
		for dep := range deps {
			lst = append(lst, dep)
		}
		sort.Strings(lst)

		ret[f] = lst
	}
	return &dagvis.Graph{Nodes: ret}, nil
}

// FileDep2 returns the dependency graph for files in a package.
func FileDep2(path string) (*dagvis.Graph, error) {
	c, err := newChecker(path)
	if err != nil {
		return nil, err
	}

	files, err := c.listFiles()
	if err != nil {
		return nil, err
	}

	return c.depGraph(files)
}
