package gcimporter

import (
	"go/types"
)

func readFieldName(p *reader, parent *types.Package) (
	pkg *types.Package, name string, alias bool,
) {
	name = readString(p)
	pkg = parent
	if pkg == nil {
		// use the imported package instead
		pkg = p.pkgList[0]
	}
	if p.version == 0 && name == "_" {
		// version 0 didn't export a package for _ fields
		return
	}
	switch name {
	case "":
		// 1) field name matches base type name and is exported: nothing to do
	case "?":
		// 2) field name matches base type name and is not exported: need
		// package
		name = ""
		pkg = readPkg(p)
	case "@":
		// 3) field name doesn't match type name (alias)
		name = readString(p)
		alias = true
		fallthrough
	default:
		if !exported(name) {
			pkg = readPkg(p)
		}
	}
	return
}
