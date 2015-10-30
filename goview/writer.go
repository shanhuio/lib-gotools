package goview

import (
	"bytes"
	"errors"
	"fmt"
	"go/scanner"
	"go/token"
	"html"
	"io"
)

func runeHTML(r rune) string {
	if r == '\t' {
		return "&nbsp;&nbsp;&nbsp;&nbsp;"
	} else if r == ' ' {
		return "&nbsp;"
	} else if r == '\n' {
		return "<br>\n"
	}
	return html.EscapeString(string(r))
}

func tokClass(tok *tok, def token.Pos) string {
	lit := tok.lit
	t := tok.tok

	switch {
	case t == token.IDENT:
		if _, found := builtInFuncMap[lit]; found {
			return "bfunc"
		} else if _, found := builtInTypeMap[lit]; found {
			return "btype"
		}
		return "ident"
	case t == token.COMMENT:
		return "cm"
	case t >= token.INT && t <= token.IMAG:
		return "num"
	case t == token.CHAR || t == token.STRING:
		return "str"
	case t.IsOperator():
		return "op"
	case t.IsKeyword():
		return "kw"
	}

	return "na"
}

type writer struct {
	f *file

	toks []*tok            // tokens
	bs   []byte            // the bytes
	errs scanner.ErrorList // parsing errors
	in   *bytes.Reader     // reader
	out  *bytes.Buffer     // output buffer
	off  int               // reader offset

	base int // base offset in the file set
	size int
	line int
}

func fileHTML(f *file, dmap *defMap) ([]byte, error) {
	w := newWriter()
	return w.file(f, dmap)
}

func newWriter() *writer {
	ret := new(writer)
	ret.out = new(bytes.Buffer)

	return ret
}

func (w *writer) writeRune(ch rune) {
	fmt.Fprint(w.out, runeHTML(ch))
	if ch == '\n' {
		w.line++
	}
}

func (w *writer) initFile(f *file) error {
	var e error

	w.f = f
	w.bs, w.toks, w.errs, e = f.toks()
	if e != nil {
		return e
	}

	w.base = f.f.Base()
	w.size = f.f.Size()
	if w.size != len(w.bs) {
		return errors.New("inconsistent file size")
	}
	w.in = bytes.NewReader(w.bs)
	w.line = 0

	return nil
}

func (w *writer) writeNonToken(til int) {
	for w.off < til {
		ch, n, e := w.in.ReadRune()
		if e == io.EOF {
			panic("unexpected eof")
		}

		w.off += n
		w.writeRune(ch)
	}
}

func (w *writer) readMatchingToken(t *tok) string {
	lit := t.Lit()

	nb := len([]byte(lit))
	buf := make([]byte, nb)
	_, e := w.in.Read(buf)
	if e != nil {
		panic(e)
	}
	if lit != string(buf) {
		panic("lit mismatch")
	}

	w.off += nb

	return lit
}

func (w *writer) writeToken(lit, class string, pos, def token.Pos) {
	if class == "" {
		class = "token"
	} else {
		class = class + " token"
	}

	fmt.Fprintf(w.out, `<span class="%s" id="t%d">`,
		class, int(pos),
	)
	for _, ch := range lit {
		w.writeRune(ch)
	}
	fmt.Fprintf(w.out, `</span>`)
}

func (w *writer) generate(dmap *defMap) {
	w.off = 0

	for _, t := range w.toks {
		if t.tok == token.EOF {
			break // end of file
		}

		if t.lit == "\n" && t.tok == token.SEMICOLON {
			// implicit semicolon
			continue
		}

		toff := int(t.pos) - w.base
		if toff < 0 || toff >= w.size {
			panic("invalid file token offset")
		}

		w.writeNonToken(toff) // often white spaces
		if w.off != toff {
			panic("rune not aligned")
		}

		lit := w.readMatchingToken(t)

		var def token.Pos
		if dmap != nil {
			def = dmap.def(t.pos)
		}

		class := tokClass(t, def)
		w.writeToken(lit, class, t.pos, def)
	}
}

func (w *writer) file(f *file, dmap *defMap) ([]byte, error) {
	e := w.initFile(f)
	if e != nil {
		return nil, e
	}

	w.generate(dmap)

	ret := w.out.Bytes()
	return ret, nil
}
