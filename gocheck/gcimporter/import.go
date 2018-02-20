// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package gcimporter implements Import for gc-generated object files.
package gcimporter

import (
	"bufio"
	"fmt"
	"go/build"
	"go/types"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// debugging/development support
const debug = false

var pkgExts = [...]string{".a", ".o"}

func findPkgContext(
	ctx *build.Context, path, srcDir string,
) (filename, id string) {
	if path == "" {
		return
	}

	var noext string
	switch {
	default:
		// "x" -> "$GOPATH/pkg/$GOOS_$GOARCH/x.ext", "x"
		// Don't require the source files to be present.
		if abs, err := filepath.Abs(srcDir); err == nil { // see issue 14282
			srcDir = abs
		}
		bp, err := ctx.Import(path, srcDir, build.FindOnly|build.AllowBinary)
		if err != nil {
			panic(err)
		}
		if bp.PkgObj == "" {
			id = path // make sure we have an id to print in error message
			return
		}
		noext = strings.TrimSuffix(bp.PkgObj, ".a")
		id = bp.ImportPath

	case build.IsLocalImport(path):
		// "./x" -> "/this/directory/x.ext", "/this/directory/x"
		noext = filepath.Join(srcDir, path)
		id = noext

	case filepath.IsAbs(path):
		// for completeness only - go/build.Import
		// does not support absolute imports
		// "/x" -> "/x.ext", "/x"
		noext = path
		id = path
	}

	if false { // for debugging
		if path != id {
			fmt.Printf("%s -> %s\n", path, id)
		}
	}

	// try extensions
	for _, ext := range pkgExts {
		filename = noext + ext
		if f, err := os.Stat(filename); err == nil && !f.IsDir() {
			return
		}
	}

	filename = "" // not found
	return
}

// FindPkg returns the filename and unique package id for an import
// path based on package information provided by build.Import (using
// the build.Default build.Context). A relative srcDir is interpreted
// relative to the current working directory.
// If no file was found, an empty filename is returned.
func FindPkg(path, srcDir string) (filename, id string) {
	return findPkgContext(&build.Default, path, srcDir)
}

// Import imports a gc-generated package given its import path and srcDir, adds
// the corresponding package object to the packages map, and returns the object.
// The packages map must contain all packages already imported.
func Import(
	packages map[string]*types.Package, path, srcDir string,
) (pkg *types.Package, err error) {
	return importContext(&build.Default, packages, path, srcDir)
}

func importContext(
	ctx *build.Context, packages map[string]*types.Package,
	path, srcDir string,
) (pkg *types.Package, err error) {
	filename, id := findPkgContext(ctx, path, srcDir)
	if filename == "" {
		if path == "unsafe" {
			return types.Unsafe, nil
		}
		return nil, fmt.Errorf("can't find import: %q", id)
	}

	// no need to re-import if the package was imported completely before
	if pkg = packages[id]; pkg != nil && pkg.Complete() {
		return
	}

	// open file
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			// add file name to error
			err = fmt.Errorf("%s: %v", filename, err)
		}
	}()
	var rc io.ReadCloser = f
	defer func() {
		rc.Close()
	}()

	var hdr string
	buf := bufio.NewReader(rc)
	if hdr, err = FindExportData(buf); err != nil {
		return
	}

	switch hdr {
	case "$$\n":
		err = fmt.Errorf(
			"import %q: old export format no longer "+
				"supported (recompile library)", path,
		)
	case "$$B\n":
		var data []byte
		data, err = ioutil.ReadAll(buf)
		if err == nil {
			// TODO(gri): allow clients of go/importer to provide a FileSet.
			// Or, define a new standard go/types/gcexportdata package.
			pkg, err = bimport(packages, data, id)
			return
		}
	default:
		err = fmt.Errorf("unknown export data header: %q", hdr)
	}

	return
}
