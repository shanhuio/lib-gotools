package goview

import (
	"log"
	"path/filepath"

	"e8vm.io/tools/goload"
)

// Files generates the rendered files for a loaded program
func Files(prog *goload.Program) (map[string][]byte, []error) {
	fset := prog.Fset

	ret := make(map[string][]byte)
	var errs []error

	for _, p := range prog.Pkgs {
		pinfo := prog.Imported[p]
		for _, astf := range pinfo.Files {
			tokFile := fset.File(astf.Pos())
			name := tokFile.Name()
			view := newFile(tokFile, name)

			bs, e := fileHTML(view, nil) //TODO: add def map
			if e != nil {
				log.Println(e)
				errs = append(errs, e)
				continue
			}

			base := filepath.Base(name)
			path := filepath.Join(p, base)
			ret[path] = bs
		}
	}

	return ret, errs
}
