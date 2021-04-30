package gocheck

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"sort"

	"shanhu.io/smlvm/dagvis"
)

type depGraphInput struct {
	fset  *token.FileSet
	files []*ast.File
	info  *types.Info
	pkg   *types.Package
}

func pkgDepGraph(in *depGraphInput) (*dagvis.Graph, error) {
	info, typesPkg := in.info, in.pkg

	depsMap := make(map[token.Pos]map[token.Pos]bool)
	for _, f := range in.files {
		depsMap[filePos(in.fset, f.Pos())] = make(map[token.Pos]bool)
	}

	for use, obj := range info.Uses {
		if obj.Pkg() != typesPkg {
			continue // ignore inter-pkg refs
		}

		fused := filePos(in.fset, use.NamePos)
		fdef := filePos(in.fset, obj.Pos())

		if fused == fdef {
			continue
		}

		if _, found := depsMap[fdef]; !found {
			path := typesPkg.Path()
			panic(fmt.Errorf("%s not found in %s", use.Name, path))
		}
		depsMap[fdef][fused] = true
	}

	ret := make(map[string][]string)
	for f, deps := range depsMap {
		var lst []string
		for dep := range deps {
			lst = append(lst, trimBase(filename(in.fset, dep)))
		}
		sort.Strings(lst)
		ret[trimBase(filename(in.fset, f))] = lst
	}
	return &dagvis.Graph{Nodes: ret}, nil
}
