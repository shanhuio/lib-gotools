package gorepo

import (
	"encoding/json"
	"io"
	"path/filepath"

	"e8vm.io/e8vm/dagvis"
	"e8vm.io/tools/godep"
	"e8vm.io/tools/goload"
)

type repo struct {
	path string
	pkgs []string
	prog *goload.Program
}

func newRepo(path string) (*repo, error) {
	pkgs, err := goload.ListPkgs(path)
	if err != nil {
		return nil, err
	}

	prog, err := goload.Pkgs(pkgs)
	if err != nil {
		return nil, err
	}

	return &repo{
		path: path,
		pkgs: pkgs,
		prog: prog,
	}, nil
}

func (r *repo) trimPath(path string) (string, error) {
	return filepath.Rel(r.path, path)
}

func (r *repo) pkgDeps() (interface{}, error) {
	g, err := godep.PkgDep(r.pkgs)
	if err != nil {
		return nil, err
	}

	g, err = g.Rename(r.trimPath)
	if err != nil {
		return nil, err
	}
	m, err := dagvis.Layout(g)
	if err != nil {
		return nil, err
	}
	return dagvis.JSONMap(m), nil
}

func (r *repo) fileDeps() (map[string]interface{}, []error) {
	deps := godep.FileDepLoaded(r.prog)

	var errs []error
	ret := make(map[string]interface{})
	for pkg, dep := range deps {
		m, err := dagvis.Layout(dep)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		p, err := r.trimPath(pkg)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		ret[p] = dagvis.JSONMap(m)
	}
	return ret, errs
}

// Build builds a repo and writes the output via writer w.
func Build(path string, w io.Writer) []error {
	r, err := newRepo(path)
	if err != nil {
		return []error{err}
	}

	var deps struct {
		Pkgs  interface{}
		Files map[string]interface{}
	}

	deps.Pkgs, err = r.pkgDeps()
	if err != nil {
		return []error{err}
	}

	var errs []error
	deps.Files, errs = r.fileDeps()
	if err != nil {
		return errs
	}

	enc := json.NewEncoder(w)
	if err := enc.Encode(deps); err != nil {
		return []error{err}
	}
	return nil
}
