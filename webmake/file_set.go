package webmake

import (
	"os"
	"sort"
)

type fileSet struct {
	names   []string
	infoMap map[string]os.FileInfo
}

func newFileSet() *fileSet {
	return &fileSet{
		infoMap: make(map[string]os.FileInfo),
	}
}

func (fs *fileSet) sortNames() {
	sort.Strings(fs.names)
}

func (fs *fileSet) add(info os.FileInfo) {
	name := info.Name()
	if _, found := fs.infoMap[name]; found {
		panic("name conflict")
	}

	fs.names = append(fs.names, info.Name())
	fs.infoMap[name] = info
}
