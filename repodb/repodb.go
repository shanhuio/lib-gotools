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

// Open opens a repo database.
func Open(path string) (*RepoDB, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	return &RepoDB{db: db}, nil
}

func (db *RepoDB) x(q string, args ...interface{}) (sql.Result, error) {
	res, err := db.db.Exec(q, args...)
	return res, qerr(q, err)
}

func (db *RepoDB) q1(q string, args ...interface{}) *sql.Row {
	return db.db.QueryRow(q, args...)
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
func (db *RepoDB) Add(b *Build) error {
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
			b.Build, f, jsonBytes(content),
		)
		if err != nil {
			return err
		}
	}

	_, err = db.x(`
		insert or replace into latest_builds (repo, build)
		values (?, ?);`,
		b.Name, b.Build,
	)
	if err != nil {
		return err
	}

	return nil
}

// LatestBuild replies the latest build of a repo
func (db *RepoDB) LatestBuild(repo string) (*LatestBuild, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	q := `select builds.build, builds.lang, builds.struct
		from builds, latest_builds
		where latest_builds.repo = ?
		and latest_builds.build = builds.build;`

	row := db.q1(q, repo)

	var build, lang string
	var structure []byte
	err := row.Scan(&build, &lang, &structure)
	if err != nil {
		return nil, qerr(q, err)
	}

	return &LatestBuild{
		Repo:   repo,
		Build:  build,
		Lang:   lang,
		Struct: structure,
	}, nil
}

// LatestFile replies the particular file in the lastest build of a repo.
func (db *RepoDB) LatestFile(repo, path string) (*LatestFile, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	q := `select files.content, files.build from
		builds, latest_builds
		where latest_builds.repo = ?
		and latest_builds.build = files.build
		and files.file = ?`
	row := db.q1(q, repo, path)
	var build string
	var ret []byte
	err := row.Scan(&ret, &build)
	if err != nil {
		return nil, qerr(q, err)
	}

	return &LatestFile{
		Repo:    repo,
		Path:    path,
		Build:   build,
		Content: ret,
	}, nil
}
