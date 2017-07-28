package godep

import (
	"fmt"
	"sort"
	"strings"

	"go/build"
	"golang.org/x/tools/go/buildutil"

	"shanhu.io/p/shanhu.io/smlvm/dagvis"
)

func skipPkg(p string) bool {
	if p == "C" {
		return true
	}
	if p == "lib9" {
		return true
	}
	if strings.HasPrefix(p, "lib9/") {
		return true
	}
	if p == "cmd" {
		return true
	}
	if strings.HasPrefix(p, "cmd/") {
		return true
	}
	if strings.HasPrefix(p, "vendor/") {
		return true
	}
	return false
}

// ListStdPkgs returns a list of all packages in
// the golang standard library
func ListStdPkgs() []string {
	c := build.Default
	c.CgoEnabled = false
	c.GOPATH = ""

	pkgs := buildutil.AllPackages(&c)

	var ret []string
	for _, p := range pkgs {
		if skipPkg(p) {
			continue
		}

		_, err := build.Import(p, "", 0)
		if err != nil {
			continue // ignore error packges
		}

		ret = append(ret, p)
	}
	return ret
}

type pkgDep struct {
	pkgs []string

	pkgSet map[string]struct{}
}

func (d *pkgDep) imports(p string) ([]string, error) {
	pkg, err := build.Import(p, "", 0) // parse the package
	if err != nil {
		return nil, err
	}

	var ret []string
	for _, imp := range pkg.Imports {
		if _, b := d.pkgSet[imp]; !b {
			continue // filter out pkgs not in the set
		}
		ret = append(ret, imp)
	}
	sort.Strings(ret)

	return ret, nil
}

func (d *pkgDep) build() (*dagvis.Graph, error) {
	d.pkgSet = make(map[string]struct{})
	for _, p := range d.pkgs {
		d.pkgSet[p] = struct{}{}
	}

	ret := make(map[string][]string)

	for _, p := range d.pkgs {
		imps, e := d.imports(p)
		if e != nil {
			return nil, e
		}
		ret[p] = imps
	}

	e := d.check(ret)
	if e != nil {
		return nil, e
	}

	g := &dagvis.Graph{Nodes: ret}
	return g.Reverse(), nil
}

func (d *pkgDep) check(g map[string][]string) error {
	for p, subs := range g {
		for _, sub := range subs {
			if _, b := g[sub]; b {
				continue
			}
			return fmt.Errorf("pkg %q missing (for %q)", sub, p)
		}
	}

	return nil
}

// PkgDep retunrs the dependency graph for the particular packages.
func PkgDep(pkgs []string) (*dagvis.Graph, error) {
	d := &pkgDep{pkgs: pkgs}
	return d.build()
}

// StdDep returns the dependency graph for Go std library.
func StdDep() (*dagvis.Graph, error) {
	pkgs := ListStdPkgs()
	return PkgDep(pkgs)
}
