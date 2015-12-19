package webmake

import (
	"os/exec"
)

func buildLess(b *builder, f string) error {
	fin := b.fin(f)
	fout := b.foutAs(f, ".css")
	cmd := exec.Command("lessc", fin, fout)
	return cmd.Run()
}
