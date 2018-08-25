package gcimporter

import (
	"go/types"
)

// A dddSlice is a types.Type representing ...T parameters.
// It only appears for parameter types and does not escape
// the importer.
type dddSlice struct {
	elem types.Type
}

func (t *dddSlice) Underlying() types.Type { return t }
func (t *dddSlice) String() string         { return "..." + t.elem.String() }
