package gcimporter

import (
	"go/constant"
	"go/token"
)

func readValue(p *reader) constant.Value {
	switch tag := readTagOrIndex(p); tag {
	case falseTag:
		return constant.MakeBool(false)
	case trueTag:
		return constant.MakeBool(true)
	case int64Tag:
		return constant.MakeInt64(readInt64(p))
	case floatTag:
		return readFloat(p)
	case complexTag:
		re := readFloat(p)
		im := readFloat(p)
		return constant.BinaryOp(re, token.ADD, constant.MakeImag(im))
	case stringTag:
		return constant.MakeString(readString(p))
	case unknownTag:
		return constant.MakeUnknown()
	default:
		errorf("unexpected value tag %d", tag) // panics
		panic("unreachable")
	}
}
