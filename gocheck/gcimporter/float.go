package gcimporter

import (
	"go/constant"
	"go/token"
)

func readFloat(p *reader) constant.Value {
	sign := readInt(p)
	if sign == 0 {
		return constant.MakeInt64(0)
	}

	exp := readInt(p)
	mant := []byte(readString(p)) // big endian

	// remove leading 0's if any
	for len(mant) > 0 && mant[0] == 0 {
		mant = mant[1:]
	}

	// convert to little endian
	// TODO(gri) go/constant should have a more direct conversion function
	//           (e.g., once it supports a big.Float based implementation)
	for i, j := 0, len(mant)-1; i < j; i, j = i+1, j-1 {
		mant[i], mant[j] = mant[j], mant[i]
	}

	// adjust exponent (constant.MakeFromBytes creates an integer value,
	// but mant represents the mantissa bits such that 0.5 <= mant < 1.0)
	exp -= len(mant) << 3
	if len(mant) > 0 {
		for msd := mant[len(mant)-1]; msd&0x80 == 0; msd <<= 1 {
			exp++
		}
	}

	x := constant.MakeFromBytes(mant)
	switch {
	case exp < 0:
		d := constant.Shift(constant.MakeInt64(1), token.SHL, uint(-exp))
		x = constant.BinaryOp(x, token.QUO, d)
	case exp > 0:
		x = constant.Shift(x, token.SHL, uint(exp))
	}

	if sign < 0 {
		x = constant.UnaryOp(token.SUB, x, 0)
	}
	return x
}
