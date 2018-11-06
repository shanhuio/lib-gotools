package gocheck

import (
	"go/build"

	"github.com/h8liu/gcimporter"
	"shanhu.io/smlvm/dagvis"
)

// DepGraph returns the dependency graph for files in a package.
func DepGraph(
	ctx *build.Context, path string, alias *gcimporter.AliasMap,
) (*dagvis.Graph, error) {
	c, err := newCheckerPath(ctx, path, alias)
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
func DepGraphPkg(
	ctx *build.Context, pkg *build.Package, alias *gcimporter.AliasMap,
) (*dagvis.Graph, error) {
	c := newChecker(ctx, pkg, alias)
	files, err := c.listFiles()
	if err != nil {
		return nil, err
	}
	return c.depGraph(files)
}
