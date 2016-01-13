package gorepo

import (
	"go/build"
	"os/exec"
	"strings"
)

// GitCommitHash gets the head commit of the master branch.
func GitCommitHash(path string) (string, error) {
	pkg, err := build.Import(path, "", build.FindOnly)
	if err != nil {
		return "", err
	}
	srcPath := pkg.Dir

	cmd := exec.Command("git", "show", "-s", "--format=%H")
	cmd.Dir = srcPath

	hash, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(hash)), nil
}
