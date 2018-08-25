package gcimporter

import (
	"go/types"
)

func readObj(p *reader, tag int) {
	switch tag {
	case constTag:
		pos := readPos(p)
		pkg, name := readQualifiedName(p)
		typ := readType(p, nil)
		val := readValue(p)
		declare(types.NewConst(pos, pkg, name, typ, val))

	case aliasTag:
		// TODO(gri) verify type alias hookup is correct
		pos := readPos(p)
		pkg, name := readQualifiedName(p)
		typ := readType(p, nil)
		declare(types.NewTypeName(pos, pkg, name, typ))

	case typeTag:
		readType(p, nil)

	case varTag:
		pos := readPos(p)
		pkg, name := readQualifiedName(p)
		typ := readType(p, nil)
		declare(types.NewVar(pos, pkg, name, typ))

	case funcTag:
		pos := readPos(p)
		pkg, name := readQualifiedName(p)
		params, isddd := readParamList(p)
		result, _ := readParamList(p)
		sig := types.NewSignature(nil, params, result, isddd)
		declare(types.NewFunc(pos, pkg, name, sig))

	default:
		errorf("unexpected object tag %d", tag)
	}
}
