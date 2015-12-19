package webmake

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// Template is a page template.
type Template struct {
	Index  string
	Header string
	Footer string
}

// NewTemplate creates a new template from a folder.
func NewTemplate(dir string) (*Template, error) {
	open := func(s string) (string, error) {
		f := filepath.Join(dir, s)
		bs, err := ioutil.ReadFile(f)
		if os.IsNotExist(err) {
			return "", nil
		}
		if err != nil {
			return "", err
		}
		return string(bs), nil
	}

	var err error
	ret := new(Template)
	for f, s := range map[string]*string{
		"index.html":  &ret.Index,
		"header.html": &ret.Header,
		"footer.html": &ret.Footer,
	} {
		*s, err = open(f)
		if err != nil {
			return nil, err
		}
	}
	return ret, nil
}
