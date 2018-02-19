package gcimporter

import (
	"go/types"
)

func readPkg(p *reader) *types.Package {
	// if the package was seen before, i is its index (>= 0)
	i := readTagOrIndex(p)
	if i >= 0 {
		return p.pkgList[i]
	}

	// otherwise, i is the package tag (< 0)
	if i != packageTag {
		errorf("unexpected package tag %d version %d", i, p.version)
	}

	// read package data
	name := readString(p)
	var path string
	if p.version >= 5 {
		path = readPath(p)
	} else {
		path = readString(p)
	}

	// we should never see an empty package name
	if name == "" {
		errorf("empty package name in import")
	}

	// an empty path denotes the package we are currently importing;
	// it must be the first package we see
	if (path == "") != (len(p.pkgList) == 0) {
		errorf("package path %q for pkg index %d", path, len(p.pkgList))
	}

	// if the package was imported before, use that one; otherwise create a new
	// one
	if path == "" {
		path = p.importpath
	}
	pkg := p.imports[path]
	if pkg == nil {
		pkg = types.NewPackage(path, name)
		p.imports[path] = pkg
	} else if pkg.Name() != name {
		errorf(
			"conflicting names %s and %s for package %q",
			pkg.Name(), name, path,
		)
	}
	p.pkgList = append(p.pkgList, pkg)

	return pkg
}

func readQualifiedName(p *reader) (pkg *types.Package, name string) {
	name = readString(p)
	pkg = readPkg(p)
	return
}
