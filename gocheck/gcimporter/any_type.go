package gcimporter

import (
	"go/types"
)

type anyType struct{}

func (t anyType) Underlying() types.Type { return t }
func (t anyType) String() string         { return "any" }
