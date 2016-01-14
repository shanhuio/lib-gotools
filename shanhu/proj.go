package shanhu

import (
	"html/template"

	"e8vm.io/tools/repodb"
)

func repoFromPath(path string) string {
	switch path {
	case "/", "/e8vm/e8vm":
		return "e8vm.io/e8vm"
	case "/e8vm/tools":
		return "e8vm.io/tools"
	}
	return "e8vm.io/e8vm"
}

func projDat(db *repodb.RepoDB, c *context, user, path string) (
	interface{}, error,
) {
	repo := repoFromPath(path)

	b, err := db.LatestBuild(repo)
	if err != nil {
		return nil, err
	}

	type d struct {
		Repo string
		User string
		Proj template.JS
	}

	return &d{
		Repo: repo,
		User: user,
		Proj: template.JS(string(b.Struct)),
	}, nil
}
