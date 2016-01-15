package goview

import (
	"log"
	"path/filepath"

	"e8vm.io/tools/goload"
)

// File saves the rendering of a file.
type File struct {
	HTML  []byte `json:"c"`
	Nline int    `json:"n"`
}

// Files generates the rendered files for a loaded program
func Files(prog *goload.Program) (map[string]*File, []error) {
	fset := prog.Fset

	ret := make(map[string]*File)
	var errs []error

	for _, p := range prog.Pkgs {
		pinfo := prog.Imported[p]
		for _, astf := range pinfo.Files {
			tokFile := fset.File(astf.Pos())
			name := tokFile.Name()
			view := newFile(tokFile, name)

			bs, nline, e := fileHTML(view, nil) // TODO: add def map
			if e != nil {
				log.Println(e)
				errs = append(errs, e)
				continue
			}

			base := filepath.Base(name)
			path := filepath.Join(p, base)
			ret[path] = &File{bs, nline}
		}
	}

	return ret, errs
}
