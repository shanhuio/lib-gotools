package webmake

import (
	"fmt"
)

func buildAll(b *builder) error {
	var err error

	o := func(fset *fileSet, fn func(b *builder, f string) error) error {
		if err != nil {
			return err
		}

		for _, f := range fset.names {
			fin := b.fin(f)
			fmt.Println(fin)
			if err = fn(b, f); err != nil {
				return err
			}
		}
		return nil
	}

	dir := b.dir
	o(dir.htmlFiles, buildHTML)
	o(dir.lessFiles, buildLess)
	return err
}

// Build builds everything under the given directory
func Build(p *Project) error {
	b, err := newBuilder(p)
	if err != nil {
		return err
	}

	return buildAll(b)
}
