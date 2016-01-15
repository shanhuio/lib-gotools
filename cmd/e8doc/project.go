package main

import (
	"encoding/json"
	"log"
	"path/filepath"

	"e8vm.io/e8vm/dagvis"
	"e8vm.io/tools/godep"
	"e8vm.io/tools/goload"
	"e8vm.io/tools/goview"
)

type project struct {
	path string

	*memWeb

	errs []error // errors
}

func newProject(prefix, path string) *project {
	ret := new(project)
	ret.memWeb = newMemWeb()
	ret.memWeb.repath = func(p string) string {
		ret, e := filepath.Rel(prefix, p)
		if e != nil {
			return p
		}
		return ret
	}

	ret.path = path

	ret.build()

	return ret
}

func (p *project) jsonObj(obj interface{}) []byte {
	type Obj struct {
		Nerror int         `json:"nerror"`
		Obj    interface{} `json:"data"`
	}

	bs, e := json.Marshal(&Obj{Nerror: len(p.errs), Obj: obj})
	fatale(e)

	return bs
}

func (p *project) build() {
	p.lock.Lock()
	defer p.lock.Unlock()
	p.clear()

	// clear the saving slots
	pkgs, e := goload.ListPkgs(p.path)
	if p.err(e) {
		log.Printf("error: %s", e)
		return
	}
	pkgDep := p.buildPkgDep(pkgs)

	var fileDeps map[string]interface{}
	var files map[string]*goview.File

	prog, e := goload.Pkgs(pkgs)
	if !p.err(e) {
		fileDeps = p.buildFileDeps(prog)

		var es []error
		files, es = goview.Files(prog)
		p.err(es...)
	}

	p.setPage("INDEX", p.jsonObj(pkgDep))
	for pkg, dep := range fileDeps {
		pkg, e = p.trimPath(pkg)
		if p.err(e) {
			continue
		}
		p.setPage(pkg, p.jsonObj(dep))
	}

	for path, f := range files {
		path, e := p.trimPath(path)
		if p.err(e) {
			continue
		}
		p.setPage(path, f.HTML)
	}

	for _, e := range p.errs {
		log.Printf("error: %s", e)
	}
}

func (p *project) trimPath(path string) (string, error) {
	return filepath.Rel(p.path, path)
}

func (p *project) err(es ...error) bool {
	ret := false
	for _, e := range es {
		if e == nil {
			continue
		}
		ret = true
		p.errs = append(p.errs, e)
	}
	return ret
}

func (p *project) buildPkgDep(pkgs []string) interface{} {
	g, e := godep.PkgDep(pkgs)
	if p.err(e) {
		return nil
	}

	g, e = g.Rename(p.trimPath)
	if p.err(e) {
		return nil
	}

	m, e := dagvis.Layout(g)
	if p.err(e) {
		return nil
	}

	return dagvis.JSONMap(m)
}

func (p *project) buildFileDeps(prog *goload.Program) map[string]interface{} {
	pkgDeps := godep.FileDepLoaded(prog)

	ret := make(map[string]interface{})
	for pkg, g := range pkgDeps {
		m, e := dagvis.Layout(g)
		if p.err(e) {
			continue
		}
		ret[pkg] = dagvis.JSONMap(m)
	}

	return ret
}
