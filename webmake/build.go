package webmake

import (
	"os"
)

// Build builds everything under the given directory
func Build(p *Project) error {
	dir, err := newDir(p.In)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(p.Out, 0755); err != nil {
		return err
	}

	for _, f := range dir.htmlFiles.names {
		if err := buildHTML(p, dir, f); err != nil {
			return err
		}
	}
	return nil
}
