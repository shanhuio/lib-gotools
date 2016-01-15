package repodb

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3" // for sqlite3
)

// Create creates the database with the schema.
func Create(f string) error {
	db, err := sql.Open("sqlite3", f)
	if err != nil {
		return err
	}

	q := func(s string) {
		if err != nil {
			return
		}
		_, e := db.Exec(s)
		if e != nil {
			err = e
		}
	}

	q(` create table builds (
			repo text not null,
			build text primary key not null,
			lang text not null,
			struct blob not null
		);
	`)

	q(` create table latest_builds (
			repo text primary key not null,
			build text not null
		);
	`)

	q(` create table files (
			build text not null,
			file text not null,
			content blob not null
		);
	`)

	q(`create index file_index on files (build, file);`)

	return err
}
