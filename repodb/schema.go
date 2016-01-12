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
			repo text,
			build text primary key,
			lang text,
			struct text
		);
	`)

	q(` create table latest_builds (
			repo text primary key,
			build text
		);
	`)

	q(` create table files (
			repo text,
			build text primary key,
			file text,
			content text
		);
	`)

	return err
}
