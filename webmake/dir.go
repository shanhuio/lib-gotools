package webmake

import (
	"os"
	"path/filepath"
	"strings"
)

type dir struct {
	path     string
	allFiles *fileSet

	htmlFiles   *fileSet
	lessFiles   *fileSet
	cssFiles    *fileSet
	mdFiles     *fileSet
	coffeeFiles *fileSet
	jsFiles     *fileSet
	goFiles     *fileSet
	pngFiles    *fileSet

	extMap map[string]*fileSet

	template *Template
}

func typeOf(p, ext string) bool {
	return strings.HasSuffix(p, "."+ext)
}

func (d *dir) addFile(name string, info os.FileInfo) bool {
	ext := filepath.Ext(name)
	fs := d.extMap[ext]
	if fs == nil {
		return false
	}
	fs.add(info)
	d.allFiles.add(info)
	return true
}

func newEmptyDir(path string) *dir {
	ret := &dir{
		path: path,

		allFiles: newFileSet(),

		htmlFiles:   newFileSet(),
		lessFiles:   newFileSet(),
		cssFiles:    newFileSet(),
		mdFiles:     newFileSet(),
		coffeeFiles: newFileSet(),
		jsFiles:     newFileSet(),
		goFiles:     newFileSet(),
		pngFiles:    newFileSet(),
	}
	ret.extMap = map[string]*fileSet{
		".html":     ret.htmlFiles,
		".htm":      ret.htmlFiles,
		".less":     ret.lessFiles,
		".css":      ret.cssFiles,
		".md":       ret.mdFiles,
		".markdown": ret.mdFiles,
		".coffee":   ret.coffeeFiles,
		".js":       ret.jsFiles,
		".go":       ret.goFiles,
		".png":      ret.pngFiles,
	}
	return ret
}

func newDir(path string) (*dir, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	infos, err := f.Readdir(0)
	if err != nil {
		return nil, err
	}

	fs := newFileSet()
	for _, info := range infos {
		fs.add(info)
	}
	fs.sortNames()

	ret := newEmptyDir(path)
	for _, name := range fs.names {
		info := fs.infoMap[name]
		if info.IsDir() {
			if name == "template" {
				t, err := NewTemplate(filepath.Join(path, name))
				if err != nil {
					return nil, err
				}
				ret.template = t
				continue
			}

			// TODO: load sub dir
		}

		ret.addFile(name, info)
	}

	return ret, nil
}
