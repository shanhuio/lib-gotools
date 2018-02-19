package gcimporter

import (
	"go/token"
	"sync"
)

var (
	fakeLines     []int
	fakeLinesOnce sync.Once
)

const deltaNewFile = -64 // see cmd/compile/internal/gc/bexport.go

func readPos(p *reader) token.Pos {
	if !p.posInfoFormat {
		return token.NoPos
	}

	file := p.prevFile
	line := p.prevLine
	delta := readInt(p)
	line += delta
	if p.version >= 5 {
		if delta == deltaNewFile {
			if n := readInt(p); n >= 0 {
				// file changed
				file = readPath(p)
				line = n
			}
		}
	} else {
		if delta == 0 {
			if n := readInt(p); n >= 0 {
				// file changed
				file = p.prevFile[:n] + readString(p)
				line = readInt(p)
			}
		}
	}
	p.prevFile = file
	p.prevLine = line

	// Synthesize a token.Pos

	// Since we don't know the set of needed file positions, we
	// reserve maxlines positions per file.
	const maxlines = 64 * 1024
	f := p.files[file]
	if f == nil {
		f = p.fset.AddFile(file, -1, maxlines)
		p.files[file] = f
		// Allocate the fake linebreak indices on first use.
		// TODO(adonovan): opt: save ~512KB using a more complex scheme?
		fakeLinesOnce.Do(func() {
			fakeLines = make([]int, maxlines)
			for i := range fakeLines {
				fakeLines[i] = i
			}
		})
		f.SetLines(fakeLines)
	}

	if line > maxlines {
		line = 1
	}

	// Treat the file as if it contained only newlines
	// and column=1: use the line number as the offset.
	return f.Pos(line - 1)
}
