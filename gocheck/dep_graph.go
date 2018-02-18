package gocheck

import (
	"go/build"

	"shanhu.io/smlvm/dagvis"
)

// DepGraph returns the dependency graph for files in a package.
func DepGraph(path string) (*dagvis.Graph, error) {
	c, err := newCheckerPath(path)
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
func DepGraphPkg(pkg *build.Package) (*dagvis.Graph, error) {
	c := newChecker(pkg)
	files, err := c.listFiles()
	if err != nil {
		return nil, err
	}
	return c.depGraph(files)
}
