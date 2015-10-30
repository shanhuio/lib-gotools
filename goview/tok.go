package goview

import (
	"go/token"
)

type tok struct {
	pos token.Pos
	tok token.Token
	lit string
}

func (t *tok) Lit() string {
	if t.tok.IsOperator() {
		return t.tok.String()
	}
	return t.lit
}
