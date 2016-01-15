package shanhu

import (
	"html/template"
	"path"

	"e8vm.io/tools/repodb"
)

func fileDat(db *repodb.RepoDB, user, p string) (
	interface{}, error,
) {
	repoUser, dirs := pathSplit(p)
	repo := path.Join(repoUser, dirs[0])

	f, err := db.LatestFile(repo, p)
	if err != nil {
		return nil, err
	}

	type d struct {
		User     string
		Repo     string
		RepoUser string
		Dirs     []string
		Commit   string
		File     template.JS
	}

	return &d{
		User:     user,
		Repo:     repo,
		RepoUser: repoUser,
		Dirs:     dirs,
		Commit:   shortCommit(f.Build),
		File:     template.JS(string(f.Content)),
	}, nil
}
