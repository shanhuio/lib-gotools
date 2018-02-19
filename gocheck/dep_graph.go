package gocheck

import (
	"go/build"

	"shanhu.io/smlvm/dagvis"
)

// DepGraph returns the dependency graph for files in a package.
func DepGraph(ctx *build.Context, path string) (*dagvis.Graph, error) {
	c, err := newCheckerPath(ctx, path)
	if err != nil {
		return nil, err
	}

	files, err := c.listFiles()
	if err != nil {
		return nil, err
	}

	return c.depGraph(files)
}

// DepGraphPkg returns the dependency graph for files in a loaded package.
func DepGraphPkg(ctx *build.Context, pkg *build.Package) (
	*dagvis.Graph, error,
) {
	c := newChecker(ctx, pkg)
	files, err := c.listFiles()
	if err != nil {
		return nil, err
	}
	return c.depGraph(files)
}
