package gcimporter

import (
	"strings"
)

func readPath(p *reader) string {
	if p.debugFormat {
		readMarker(p, 'p')
	}
	// if the path was seen before, i is its index (>= 0)
	// (the empty string is at index 0)
	i := readRawInt64(p)
	if i >= 0 {
		return p.pathList[i]
	}
	// otherwise, i is the negative path length (< 0)
	a := make([]string, -i)
	for n := range a {
		a[n] = readString(p)
	}
	s := strings.Join(a, "/")
	p.pathList = append(p.pathList, s)
	return s
}
