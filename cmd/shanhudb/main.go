package main

import (
	"flag"
	"log"

	"e8vm.io/tools/gorepo"
	"e8vm.io/tools/repodb"
)

var (
	initDB   = flag.Bool("init", false, "init db")
	dbPath   = flag.String("db", "shanhu.db", "path to the database")
	addGoPkg = flag.String("go", "", "add go package into database")
)

func main() {
	flag.Parse()

	switch {
	case *initDB:
		err := repodb.Create(*dbPath)
		if err != nil {
			log.Fatal(err)
		}
	case *addGoPkg != "":
		b, errs := gorepo.Build(*addGoPkg)
		for _, err := range errs {
			log.Print(err)
		}
		if len(errs) != 0 {
			log.Fatal("(build failed")
		}

		db, err := repodb.Open(*dbPath)
		if err != nil {
			log.Fatal(err)
		}
		if err := db.Add(b); err != nil {
			log.Fatal(err)
		}
		log.Printf("%s : %s added into %s", b.Name, b.Build, *dbPath)

	default:
		log.Print("(did nothing)")
	}
}
