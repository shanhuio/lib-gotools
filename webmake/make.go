package webmake

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
)

func buildHTML(p *Project, dir *dir, f string) error {
	inPath := filepath.Join(p.In, f)
	outPath := filepath.Join(p.Out, f)
	fmt.Println(inPath)

	bs, err := ioutil.ReadFile(inPath)
	if err != nil {
		return err
	}
	// TODO: parse for body, styles and scripts

	if dir.template == nil {
		// no template just copy the contents
		return ioutil.WriteFile(outPath, bs, 0644)
	}

	fout, err := os.Create(outPath)
	if err != nil {
		return err
	}

	page := &Page{
		Body: template.HTML(string(bs)),
	}
	if err := page.Render(fout, dir.template); err != nil {
		fout.Close()
		return err
	}
	if err := fout.Close(); err != nil {
		return err
	}
	return nil
}

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
