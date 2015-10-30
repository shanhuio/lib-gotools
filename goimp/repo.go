package goimp

import (
	"bytes"
	"net/http"
	"strings"
)

// Repo is a Golang repository that this handler will handle.
type Repo struct {
	ImportRoot string
	VCS        string
	VCSRoot    string
}

// NewGitRepo creates a new git repository for import redirection.
func NewGitRepo(path, repoAddr string) *Repo {
	return &Repo{
		ImportRoot: path,
		VCS:        "git",
		VCSRoot:    repoAddr,
	}
}

func (r *Repo) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := strings.TrimSuffix(req.Host+req.URL.Path, "/")
	if !strings.HasPrefix(path, r.ImportRoot) {
		http.NotFound(w, req)
		return
	}

	d := &data{
		ImportRoot: r.ImportRoot,
		VCS:        r.VCS,
		VCSRoot:    r.VCSRoot,
		Suffix:     strings.TrimSuffix(path, r.ImportRoot),
	}

	buf := new(bytes.Buffer)
	if err := tmpl.Execute(buf, d); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write(buf.Bytes())
}
