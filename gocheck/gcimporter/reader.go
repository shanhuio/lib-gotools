package gcimporter

import (
	"go/token"
	"go/types"
)

type reader struct {
	imports    map[string]*types.Package
	data       []byte
	importpath string
	buf        []byte // for reading strings
	version    int    // export format version

	// object lists
	strList       []string           // in order of appearance
	pathList      []string           // in order of appearance
	pkgList       []*types.Package   // in order of appearance
	typList       []types.Type       // in order of appearance
	interfaceList []*types.Interface // for delayed completion only
	trackAllTypes bool

	// position encoding
	posInfoFormat bool
	prevFile      string
	prevLine      int
	fset          *token.FileSet
	files         map[string]*token.File

	// debugging support
	debugFormat bool
	read        int // bytes read
}

func (p *reader) record(t types.Type) { p.typList = append(p.typList, t) }
