package gcimporter

import (
	"fmt"
	"go/token"
	"go/types"
	"sort"
	"strconv"
	"strings"
)

type byPath []*types.Package

func (a byPath) Len() int           { return len(a) }
func (a byPath) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byPath) Less(i, j int) bool { return a[i].Path() < a[j].Path() }

// bimport imports a package from the serialized package data and returns the
// number of bytes consumed and a reference to the package.  If the export data
// version is not recognized or the format is otherwise compromised, an error
// is returned.
func bimport(
	imports map[string]*types.Package, data []byte, path string,
) (pkg *types.Package, err error) {
	fset := token.NewFileSet()
	// catch panics and return them as errors
	defer func() {
		if e := recover(); e != nil {
			// The package (filename) causing the problem is added to this
			// error by a wrapper in the caller (Import in gcimporter.go).
			// Return a (possibly nil or incomplete) package unchanged (see
			// #16088).
			err = fmt.Errorf(
				"cannot import, possibly version skew (%v) "+
					"- reinstall package", e,
			)
		}
	}()

	p := &reader{
		imports:    imports,
		data:       data,
		importpath: path,
		version:    -1,           // unknown version
		strList:    []string{""}, // empty string is mapped to 0
		pathList:   []string{""}, // empty string is mapped to 0
		fset:       fset,
		files:      make(map[string]*token.File),
	}

	// read version info
	var versionstr string
	if b := readRawByte(p); b == 'c' || b == 'd' {
		// Go1.7 encoding; first byte encodes low-level
		// encoding format (compact vs debug).
		// For backward-compatibility only (avoid problems with
		// old installed packages). Newly compiled packages use
		// the extensible format string.
		// TODO(gri) Remove this support eventually; after Go1.8.
		if b == 'd' {
			p.debugFormat = true
		}
		p.trackAllTypes = readRawByte(p) == 'a'
		p.posInfoFormat = readInt(p) != 0
		versionstr = readString(p)
		if versionstr == "v1" {
			p.version = 0
		}
	} else {
		// Go1.8 extensible encoding
		// read version string and extract version number (ignore anything
		// after the version number)
		versionstr = readRawStringln(p, b)
		s := strings.SplitN(versionstr, " ", 3)
		if len(s) >= 2 && s[0] == "version" {
			if v, err := strconv.Atoi(s[1]); err == nil && v > 0 {
				p.version = v
			}
		}
	}

	// read version specific flags - extend as necessary
	switch p.version {
	// case 6:
	// 	...
	//	fallthrough
	case 5, 4, 3, 2, 1:
		p.debugFormat = readRawStringln(p, readRawByte(p)) == "debug"
		p.trackAllTypes = readInt(p) != 0
		p.posInfoFormat = readInt(p) != 0
	case 0:
		// Go1.7 encoding format - nothing to do here
	default:
		errorf("unknown export format version %d (%q)", p.version, versionstr)
	}

	// --- generic export data ---

	// populate typList with predeclared "known" types
	p.typList = append(p.typList, predeclared...)

	// read package data
	pkg = readPkg(p)

	// read objects of phase 1 only (see cmd/compile/internal/gc/bexport.go)
	objcount := 0
	for {
		tag := readTagOrIndex(p)
		if tag == endTag {
			break
		}
		readObj(p, tag)
		objcount++
	}

	// self-verification
	if count := readInt(p); count != objcount {
		errorf("got %d objects; want %d", objcount, count)
	}

	// ignore compiler-specific import data

	// complete interfaces
	// TODO(gri) re-investigate if we still need to do this in a delayed fashion
	for _, typ := range p.interfaceList {
		typ.Complete()
	}

	// record all referenced packages as imports
	list := append(([]*types.Package)(nil), p.pkgList[1:]...)
	sort.Sort(byPath(list))
	pkg.SetImports(list)

	// package was imported completely and without errors
	pkg.MarkComplete()

	return pkg, nil
}
