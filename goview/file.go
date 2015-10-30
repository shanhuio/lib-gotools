package goview

import (
	"go/scanner"
	"go/token"
	"io/ioutil"
)

type file struct {
	f    *token.File // the loaded file
	path string      // the path to the file on disk
}

func newFile(f *token.File, p string) *file {
	ret := new(file)
	ret.f = f
	ret.path = p

	return ret
}

func (f *file) toks() ([]byte, []*tok, scanner.ErrorList, error) {
	bs, e := ioutil.ReadFile(f.path)
	if e != nil {
		return nil, nil, nil, e
	}

	var errs scanner.ErrorList
	s := new(scanner.Scanner)
	s.Init(f.f, bs, errs.Add, scanner.ScanComments)

	var ret []*tok
	for {
		p, t, lit := s.Scan()
		ret = append(ret, &tok{pos: p, tok: t, lit: lit})
		if t == token.EOF {
			break
		}
	}

	return bs, ret, errs, nil
}
