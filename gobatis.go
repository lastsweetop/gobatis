package main

import (
	"flag"
	"gobatis/generate"
	"log"
	"strings"
)

var (
	typeNames = flag.String("type", "", "comma-separated list of type names; must be set")
	trimprefix  = flag.String("trimprefix", "", "trim the `prefix` from the generated constant names")
	linecomment = flag.Bool("linecomment", false, "use line comment text as printed text when present")
)

func main() {
	flag.Parse()
	log.Println("gobatis")

	types := strings.Split(*typeNames, ",")
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

	// Run generate for each type.
	for _, typeName := range types {
		g.Generate(typeName)
	}

}
