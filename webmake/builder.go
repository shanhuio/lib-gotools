package webmake

import (
	"io"
	"os"
	"path/filepath"
	"strings"
)

type builder struct {
	p   *Project
	dir *dir
}

func newBuilder(p *Project) (*builder, error) {
	dir, err := newDir(p.In)
	if err != nil {
		return nil, err
	}

	if err := os.MkdirAll(p.Out, 0755); err != nil {
		return nil, err
	}

	return &builder{p: p, dir: dir}, nil
}

func (b *builder) fin(f string) string {
	return filepath.Join(b.p.In, f)
}

func (b *builder) fout(f string) string {
	return filepath.Join(b.p.Out, f)
}

func (b *builder) foutAs(f, ext string) string {
	extNow := filepath.Ext(f)
	return filepath.Join(b.p.Out, strings.TrimSuffix(f, extNow)+ext)
}

func (b *builder) createFout(f string) (*os.File, error) {
	return os.Create(b.fout(f))
}

func (b *builder) copyFile(f string) error {
	fin, err := os.Open(b.fin(f))
	if err != nil {
		return err
	}
	defer fin.Close()

	fout, err := b.createFout(f)
	if err != nil {
		return err
	}
	defer fout.Close()

	_, err = io.Copy(fout, fin)
	if err != nil {
		return err
	}

	if err := fin.Close(); err != nil {
		return err
	}
	if err := fout.Close(); err != nil {
		return err
	}
	return nil
}
