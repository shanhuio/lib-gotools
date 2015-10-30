package main

import (
	"io"
)

type errWriter struct {
	w io.Writer
	e error
}

func (w *errWriter) Write(buf []byte) (int, error) {
	if w.e != nil {
		return 0, w.e
	}

	n, e := w.w.Write(buf)
	if e != nil {
		w.e = e
	}
	return n, e
}
