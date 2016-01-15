package gorepo

import (
	"fmt"
	"path/filepath"

	"e8vm.io/e8vm/dagvis"
	"e8vm.io/tools/godep"
	"e8vm.io/tools/goload"
	"e8vm.io/tools/goview"
	"e8vm.io/tools/repodb"
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

func (r *repo) files() (map[string]interface{}, []error) {
	files, errs := goview.Files(r.prog)
	if len(errs) > 0 {
		return nil, errs
	}

	ret := make(map[string]interface{})
	for path, f := range files {
		ret[path] = f
	}
	return ret, nil
}

// Build builds a repo into a repo build.
func Build(path string) (*repodb.Build, []error) {
	// this will also check if it is in a git repo
	buildHash, err := GitCommitHash(path)
	if err != nil {
		return nil, []error{err}
	}

	r, err := newRepo(path)
	if err != nil {
		return nil, []error{err}
	}

	var deps struct {
		Pkgs  interface{}
		Files map[string]interface{}
	}

	deps.Pkgs, err = r.pkgDeps()
	if err != nil {
		return nil, []error{err}
	}

	var errs []error
	deps.Files, errs = r.fileDeps()
	if len(errs) > 0 {
		return nil, errs
	}

	files, errs := r.files()
	if len(errs) > 0 {
		return nil, errs
	}

	return &repodb.Build{
		Name:   path,
		Build:  buildHash,
		Lang:   "go",
		Struct: deps,
		Files:  files,
	}, nil
}

// UpdateDB updates the repository in the database.
func UpdateDB(db *repodb.RepoDB, path string) []error {
	if err := GitPull(path); err != nil {
		return []error{fmt.Errorf("git pull: %s", err)}
	}

	b, errs := Build(path)
	if len(errs) > 0 {
		return errs
	}

	err := db.Add(b)
	if err != nil {
		return []error{err}
	}
	return nil
}
