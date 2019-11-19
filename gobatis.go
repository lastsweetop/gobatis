package main

import (
	"flag"
	"gobatis/generate"
	"log"
)

var (
	trimprefix  = flag.String("trimprefix", "", "trim the `prefix` from the generated constant names")
	linecomment = flag.Bool("linecomment", false, "use line comment text as printed text when present")
)

func main() {
	flag.Parse()
	log.Println("gobatis")

	// We accept either one directory or a list of files. Which do we have?
	args := flag.Args()
	if len(args) == 0 {
		// Default: process whole package in current directory.
		args = []string{"."}
	}

	g := generate.Generator{
		TrimPrefix:  *trimprefix,
		LineComment: *linecomment,
	}
	g.ParsePackage(args, nil)

	g.Generate()

}
