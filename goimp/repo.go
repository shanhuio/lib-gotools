package goimp

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// Repo is a Golang repository that this handler will handle.
type Repo struct {
	ImportRoot string
	VCS        string
	VCSRoot    string
}

func host(path string) string {
	i := strings.Index(path, "/")
	if i < 0 {
		return path
	}
	return path[:i]
}

// NewGitRepo creates a new git repository for import redirection.
func NewGitRepo(path, repoAddr string) *Repo {
	return &Repo{
		ImportRoot: path,
		VCS:        "git",
		VCSRoot:    repoAddr,
	}
}

// MetaLine returns the HTML meta line that needs to be included in the
// header of the page.
func (r *Repo) MetaLine() string {
	return fmt.Sprintf(
		`<meta name="go-import" content="%s %s %s">`,
		r.ImportRoot, r.VCS, r.VCSRoot,
	)
}

func (r *Repo) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := strings.TrimSuffix(host(r.ImportRoot)+req.URL.Path, "/")

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

	log.Println(d.ImportRoot, d.Suffix)

	buf := new(bytes.Buffer)
	if err := tmpl.Execute(buf, d); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write(buf.Bytes())
}
