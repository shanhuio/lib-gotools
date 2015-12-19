package webmake

import (
	"os"
	"path/filepath"
	"strings"
)

func listAll(dir string) (
	files map[string]os.FileInfo,
	dirs map[string]os.FileInfo,
	err error,
) {
	files = make(map[string]os.FileInfo)
	dirs = make(map[string]os.FileInfo)

	walk := func(path string, info os.FileInfo, err error) error {
		base := filepath.Base(path)
		if base != "." && (strings.HasPrefix(base, "_") ||
			strings.HasPrefix(base, ".")) {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		if info.IsDir() {
			dirs[path] = info
		} else {
			files[path] = info
		}
		return nil
	}

	err = filepath.Walk(dir, walk)
	return files, dirs, err
}
