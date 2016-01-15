package shanhu

import (
	"html/template"

	"e8vm.io/tools/repodb"
)

func repoFromPath(path string) string {
	if path == "" {
		return "e8vm.io/e8vm"
	}
	return path
}

func projDat(db *repodb.RepoDB, user, path string) (
	interface{}, error,
) {
	repo := repoFromPath(path)

	b, err := db.LatestBuild(repo)
	if err != nil {
		return nil, err
	}

	repoUser, dirs := pathSplit(repo)

	type d struct {
		User     string
		Repo     string
		RepoUser string
		Dirs     []string
		Commit   string
		Proj     template.JS
	}

	return &d{
		User:     user,
		Repo:     repo,
		RepoUser: repoUser,
		Dirs:     dirs,
		Commit:   shortCommit(b.Build),
		Proj:     template.JS(string(b.Struct)),
	}, nil
}
