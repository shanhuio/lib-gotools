package webmake

import (
	"html/template"
	"io/ioutil"
)

func buildHTML(b *builder, f string) error {
	tmpl := b.dir.template
	if tmpl == nil {
		return b.copyFile(f)
	}

	fin := b.fin(f)
	bs, err := ioutil.ReadFile(fin)
	if err != nil {
		return err
	}
	fout, err := b.createFout(f)
	if err != nil {
		return err
	}
	defer fout.Close()

	page := &Page{
		Body: template.HTML(string(bs)),
	}
	if err := page.Render(fout, tmpl); err != nil {
		return err
	}
	if err := fout.Close(); err != nil {
		return err
	}
	return nil
}
