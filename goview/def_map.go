package goview

import (
	"go/token"
)

type defMap struct {
	m map[token.Pos]token.Pos
}

func (m *defMap) add(use, defined token.Pos) {
	old := m.m[use]
	if old != defined {
		panic("inconsistent ref")
	}

	m.m[use] = defined
}

func (m *defMap) def(use token.Pos) token.Pos {
	return m.m[use]
}
