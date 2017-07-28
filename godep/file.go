package godep

import (
	"fmt"
	"go/token"
	"go/types"
	"path/filepath"
	"sort"
	"strings"

	"golang.org/x/tools/go/loader"

	"shanhu.io/p/shanhu.io/smlvm/dagvis"
	"shanhu.io/p/shanhu.io/tools/goload"
)

// filename returns a file node name for a file
func filename(fset *token.FileSet, p token.Pos) string {
	ret := filepath.Base(fset.Position(p).Filename)
	sufs := []string{".go", "_test.go"}

	for _, suf := range sufs {
		if strings.HasSuffix(ret, suf) {
			return strings.TrimSuffix(ret, suf)
		}
	}

	return ret
}

type fileDep struct {
	fset  *token.FileSet
	pkg   *types.Package
	pinfo *loader.PackageInfo

	deps map[string]map[string]struct{}
}

func (d *fileDep) addDep(fdef, fused string) {
	if fdef == fused {
		return
	}

	if _, found := d.deps[fdef]; !found {
		panic(fmt.Errorf("%s %s %s", d.pkg.Path(), fdef, fused))
	}
	d.deps[fdef][fused] = struct{}{}
}

// types analyzes file depdency using information generated
// by package types
func (d *fileDep) types() {
	for use, obj := range d.pinfo.Uses {
		pack := obj.Pkg()
		if pack != d.pkg {
			continue // ignore inter-pkg refs
		}

		fused := filename(d.fset, use.NamePos)
		fdef := filename(d.fset, obj.Pos())
		d.addDep(fdef, fused)
	}
}

// listFiles list all the files in a package
func listFiles(fset *token.FileSet, pinfo *loader.PackageInfo) []string {
	files := make(map[string]struct{})
	for _, f := range pinfo.Files {
		files[filename(fset, f.Package)] = struct{}{}
	}

	ret := make([]string, 0, len(files))
	for f := range files {
		ret = append(ret, f)
	}
	sort.Strings(ret)

	return ret
}

// build makes the file dependency graph
func (d *fileDep) build() map[string][]string {
	files := listFiles(d.fset, d.pinfo)

	// init the map
	d.deps = make(map[string]map[string]struct{})
	for _, f := range files {
		d.deps[f] = make(map[string]struct{})
	}

	d.types()

	// make it a list
	ret := make(map[string][]string)
	for f, deps := range d.deps {
		lst := make([]string, 0, len(deps))
		for dep := range deps {
			lst = append(lst, dep)
		}
		sort.Strings(lst)

		ret[f] = lst
	}

	return ret
}

// FileDepLoaded returns the dependency graph for files in a loaded
// program.
func FileDepLoaded(prog *goload.Program) map[string]*dagvis.Graph {
	fset := prog.Fset
	ret := make(map[string]*dagvis.Graph)

	for _, p := range prog.Pkgs {
		pinfo := prog.Imported[p]
		if pinfo == nil {
			panic("pinfo: " + p)
		}

		tpkg := pinfo.Pkg
		fdep := &fileDep{
			fset:  fset,
			pkg:   tpkg,
			pinfo: pinfo,
		}
		g := fdep.build()

		ret[p] = &dagvis.Graph{Nodes: g}
	}

	return ret
}

// FileDep returns the dependency graph for files in a set of packages.
func FileDep(pkgs []string) (map[string]*dagvis.Graph, error) {
	prog, e := goload.Pkgs(pkgs)
	if e != nil {
		return nil, e
	}

	return FileDepLoaded(prog), nil
}
