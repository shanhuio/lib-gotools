package webmake

import (
	"html/template"
	"io"
)

// Page is the structure for a page.
type Page struct {
	Title   string
	Body    template.HTML
	Styles  []string
	Scripts []string
}

// Render renders a page with a template into a writer.
func (p *Page) Render(w io.Writer, t *Template) error {
	tmpl, err := template.New("index").Parse(t.Index)
	if err != nil {
		return err
	}

	var dat struct {
		*Page
		Header template.HTML
		Footer template.HTML
	}
	dat.Page = p
	dat.Header = template.HTML(t.Header)
	dat.Footer = template.HTML(t.Footer)

	return tmpl.Execute(w, dat)
}
