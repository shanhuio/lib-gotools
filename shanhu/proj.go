package shanhu

import (
	"html/template"

	"e8vm.io/tools/repodb"
)

func repoFromPath(path string) string {
	switch path {
	case "/", "/e8vm":
		return "e8vm.io/e8vm"
	case "/tools":
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

	var dat struct {
		User string
		Proj template.JS
	}
	dat.User = user
	dat.Proj = template.JS(string(b.Struct))

	return &dat, nil
}
