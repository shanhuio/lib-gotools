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
		Repo     string
		RepoUser string
		Dirs     []string
		User     string
		Commit   string
		Proj     template.JS
	}

	commit := b.Build
	if len(commit) > 7 {
		commit = commit[:7]
	}

	return &d{
		Repo:     repo,
		RepoUser: repoUser,
		Dirs:     dirs,
		User:     user,
		Commit:   commit,
		Proj:     template.JS(string(b.Struct)),
	}, nil
}
