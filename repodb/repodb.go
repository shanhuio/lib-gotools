package repodb

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"sync"

	_ "github.com/mattn/go-sqlite3" // for sqlite3
)

// RepoDB is the database that stores the repository data.
type RepoDB struct {
	db *sql.DB
	mu sync.Mutex
}

func (db *RepoDB) x(q string, args ...interface{}) (sql.Result, error) {
	res, err := db.db.Exec(q, args...)
	return res, qerr(q, err)
}

func jsonBytes(v interface{}) []byte {
	bs, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return bs
}

func validBuildName(name string) bool {
	for _, c := range name {
		if c >= 'a' && c <= 'z' {
			continue
		}
		if c >= '0' && c <= '9' {
			continue
		}
		return false
	}
	return true
}

// Add adds a build of a repo into the database. It should use a unique hash
// for the build field to unqiuely identify the build.
func (db *RepoDB) Add(b *RepoBuild) error {
	if !validBuildName(b.Build) {
		return fmt.Errorf("invalid build name: %q", b.Build)
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	_, err := db.x(`
		insert into builds
		(repo, build, lang, struct)
		values (?, ?, ?, ?)`,
		b.Name, b.Build, b.Lang, jsonBytes(b.Struct),
	)
	if err != nil {
		return err
	}

	for f, content := range b.Files {
		_, err := db.x(`
			insert into files
			(build, file, content)
			values (?, ?, ?, ?)`,
			b.Build, fmt.Sprintf("%s:%s", b.Build, f),
			jsonBytes(content),
		)
		if err != nil {
			return err
		}
	}

	_, err = db.x(`
		insert into latest_builds (repo, build)
		values (?, ?);`,
		b.Name, b.Build,
	)
	if err != nil {
		return err
	}

	return nil
}
