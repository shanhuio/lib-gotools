package gcimporter

import (
	"go/types"
)

// Tags. Must be < 0.
const (
	// Objects
	packageTag = -(iota + 1)
	constTag
	typeTag
	varTag
	funcTag
	endTag

	// Types
	namedTag
	arrayTag
	sliceTag
	dddTag
	structTag
	pointerTag
	signatureTag
	interfaceTag
	mapTag
	chanTag

	// Values
	falseTag
	trueTag
	int64Tag
	floatTag
	fractionTag // not used by gc
	complexTag
	stringTag
	nilTag     // only used by gc (appears in exported inlined function bodies)
	unknownTag // not used by gc (only appears in packages with errors)

	// Type aliases
	aliasTag
)

// objTag returns the tag value for each object kind.
func objTag(obj types.Object) int {
	switch obj.(type) {
	case *types.Const:
		return constTag
	case *types.TypeName:
		return typeTag
	case *types.Var:
		return varTag
	case *types.Func:
		return funcTag
	default:
		errorf("unexpected object: %v (%T)", obj, obj) // panics
		panic("unreachable")
	}
}

func sameObj(a, b types.Object) bool {
	// Because unnamed types are not canonicalized, we cannot simply compare
	// types for (pointer) identity.  Ideally we'd check equality of constant
	// values as well, but this is good enough.
	return objTag(a) == objTag(b) && types.Identical(a.Type(), b.Type())
}

func readTagOrIndex(p *reader) int {
	if p.debugFormat {
		readMarker(p, 't')
	}
	return int(readRawInt64(p))
}
