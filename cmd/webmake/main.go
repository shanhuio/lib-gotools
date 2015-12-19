package main

import (
	"flag"
	"log"
	"os"

	"e8vm.io/tools/webmake"
)

var (
	inDir    = flag.String("in", ".", "input directory")
	outDir   = flag.String("out", "_", "output directory")
	doReadme = flag.Bool("readme", false, "also process readme.md")
)

func test1() {
	t, err := webmake.NewTemplate("template")
	if err != nil {
		log.Fatal(err)
	}

	p := &webmake.Page{
		Title:   "A sample page",
		Styles:  []string{"style.css"},
		Scripts: []string{"jquery.js"},
	}
	err = p.Render(os.Stdout, t)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	p := &webmake.Project{
		In:       *inDir,
		Out:      *outDir,
		DoReadme: *doReadme,
	}

	if err := webmake.Build(p); err != nil {
		log.Fatal(err)
	}
}
