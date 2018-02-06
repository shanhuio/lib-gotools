package gocheck

import (
	"shanhu.io/smlvm/dagvis"
)

// DepGraph returns the dependency graph for files in a package.
func DepGraph(path string) (*dagvis.Graph, error) {
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
