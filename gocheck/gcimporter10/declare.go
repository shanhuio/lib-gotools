package gcimporter

import (
	"go/types"
)

func declare(obj types.Object) {
	pkg := obj.Pkg()
	if alt := pkg.Scope().Insert(obj); alt != nil {
		// This can only trigger if we import a (non-type) object a second
		// time.  Excluding type aliases, this cannot happen because 1) we only
		// import a package once; and 2) we ignore compiler-specific export
		// data which may contain functions whose inlined function bodies refer
		// to other functions that were already imported.  However, type
		// aliases require reexporting the original type, so we need to allow
		// it (see also the comment in cmd/compile/internal/gc/bimport.go,
		// method importer.obj, switch case importing functions).
		//
		// TODO(gri) review/update this comment once the gc compiler handles
		// type aliases.
		if !sameObj(obj, alt) {
			errorf(
				"inconsistent import:\n\t%v\npreviously imported as:\n\t%v\n",
				obj, alt,
			)
		}
	}
}
