// Copyright (C) 2022  Shanhu Tech Inc.
//
// This program is free software: you can redistribute it and/or modify it
// under the terms of the GNU Affero General Public License as published by the
// Free Software Foundation, either version 3 of the License, or (at your
// option) any later version.
//
// This program is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
// FITNESS FOR A PARTICULAR PURPOSE.  See the GNU Affero General Public License
// for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package smake // import "shanhu.io/gotools/smake"

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"shanhu.io/misc/errcode"
)

func workDir() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	abs, err := filepath.Abs(wd)
	if err != nil {
		return "", err
	}
	return filepath.EvalSymlinks(abs)
}

func usingGoMod() bool {
	v, ok := os.LookupEnv("GO111MODULE")
	if !ok {
		return true
	}
	return strings.ToLower(v) != "off"
}

func run(dir string) error {
	mod := usingGoMod()

	if dir == "" {
		wd, err := workDir()
		if err != nil {
			return err
		}

		if mod {
			root, err := findGoModuleRoot(wd)
			if err != nil {
				return errcode.Annotate(err, "find module root")
			}
			wd = root
		}

		if err := os.Chdir(wd); err != nil {
			return err
		}

		dir = wd
	}

	gopath, err := absGOPATH()
	if err != nil {
		return err
	}

	c := newContext(gopath, dir, mod)
	return smake(c)
}

// Main is the entry point for smake.
func Main() {
	workDir := flag.String("dir", "", "work directory")
	flag.Parse()

	if err := run(*workDir); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
