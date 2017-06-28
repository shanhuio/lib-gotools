package goload

import (
	"go/build"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func isNoGoError(e error) bool {
	if e == nil {
		return false
	}
	_, hit := e.(*build.NoGoError)
	return hit
}

// ListPkgs list all packages under a package path.
func ListPkgs(p string) ([]string, error) {
	pkg, e := build.Import(p, "", build.FindOnly)
	if e != nil && !isNoGoError(e) {
		return nil, e
	}

	var ret []string

	e = filepath.Walk(pkg.Dir, func(
		p string,
		info os.FileInfo,
		err error,
	) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			return nil
		}

		path, e := filepath.Rel(pkg.SrcRoot, p)
		if e != nil {
			return e
		}
		base := filepath.Base(path)

		if strings.HasPrefix(base, "_") || strings.HasPrefix(base, ".") {
			return filepath.SkipDir
		}
		if base == "vendor" {
			return filepath.SkipDir
		}

		pkg, e := build.Import(path, "", 0) // check if it is a package
		if e != nil {
			if isNoGoError(e) {
				return nil
			}
			return e
		}

		ret = append(ret, pkg.ImportPath)
		return nil
	})

	if e != nil {
		return nil, e
	}

	sort.Strings(ret)
	return ret, nil
}
