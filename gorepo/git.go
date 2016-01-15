package gorepo

import (
	"go/build"
	"os/exec"
	"strings"
)

func repoSrcPath(path string) (string, error) {
	pkg, err := build.Import(path, "", build.FindOnly)
	if err != nil {
		return "", err
	}
	return pkg.Dir, nil
}

// GitCommitHash gets the head commit of the master branch.
func GitCommitHash(path string) (string, error) {
	srcPath, err := repoSrcPath(path)
	if err != nil {
		return "", err
	}

	cmd := exec.Command("git", "show", "-s", "--format=%H")
	cmd.Dir = srcPath

	hash, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(hash)), nil
}

// GitPull runs "git pull" in the repository
func GitPull(path string) error {
	srcPath, err := repoSrcPath(path)
	if err != nil {
		return err
	}

	cmd := exec.Command("git", "pull", "origin", "master")
	cmd.Dir = srcPath
	return cmd.Run()
}
