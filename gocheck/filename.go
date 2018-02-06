package gocheck

import (
	"go/token"
	"path/filepath"
	"strings"
)

func filePos(fset *token.FileSet, p token.Pos) token.Pos {
	return fset.File(p).Pos(0)
}

func filename(fset *token.FileSet, p token.Pos) string {
	return fset.Position(p).Filename
}

func trimBase(name string) string {
	base := filepath.Base(name)
	return strings.TrimSuffix(base, ".go")
}
