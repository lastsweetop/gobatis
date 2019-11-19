package generate

import (
	"fmt"
	"go/ast"
	"go/format"
	"go/token"
	"gobatis/utils"
	"golang.org/x/tools/go/packages"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func (g *Generator) Printf(format string, args ...interface{}) {
	fmt.Fprintf(&g.Buf, format, args...)
}

// format returns the gofmt-ed contents of the Generator's buffer.
func (g *Generator) Format() []byte {
	src, err := format.Source(g.Buf.Bytes())
	if err != nil {
		// Should never happen, but can arise when developing this code.
		// The user can compile the output to see the error.
		log.Printf("warning: internal error: invalid Go generated: %s", err)
		log.Printf("warning: compile the package to analyze the error")
		return g.Buf.Bytes()
	}
	return src
}

// parsePackage analyzes the single package constructed from the patterns and tags.
// parsePackage exits if there is an error.
func (g *Generator) ParsePackage(patterns []string, tags []string) {
	cfg := &packages.Config{
		Mode: packages.LoadSyntax,
		// TODO: Need to think about constants in test files. Maybe write type_string_test.go
		// in a separate pass? For later.
		Tests:      false,
		BuildFlags: []string{fmt.Sprintf("-tags=%s", strings.Join(tags, " "))},
	}
	pkgs, err := packages.Load(cfg, patterns...)
	if err != nil {
		log.Fatal(err)
	}
	if len(pkgs) != 1 {
		log.Fatalf("error: %d packages found", len(pkgs))
	}
	g.AddPackage(pkgs[0])
}

// addPackage adds a type checked Package and its syntax files to the generator.
func (g *Generator) AddPackage(pkg *packages.Package) {
	g.Pkg = &Package{
		Name:  pkg.Name,
		defs:  pkg.TypesInfo.Defs,
		files: make([]*File, len(pkg.Syntax)),
	}

	for i, file := range pkg.Syntax {
		g.Pkg.files[i] = &File{
			file:        file,
			imp:         make([]string, 0),
			pkg:         g.Pkg,
			trimPrefix:  g.TrimPrefix,
			lineComment: g.LineComment,
		}
	}
}

func (g *Generator) Generate() {

	mappers := make([]Mapper, 0, 100)
	for _, file := range g.Pkg.files {
		file.mappers = nil
		if file.file != nil {
			ast.Inspect(file.file, file.genDecl)
			fset := token.NewFileSet()
			ast.Print(fset, file.file)
			mappers = append(mappers, file.mappers...)

		}
	}

	for _, m := range mappers {
		g.Buf.Reset()
		g.Printf("// Code generated by \"gobatis %s\"; DO NOT EDIT.\n", strings.Join(os.Args[1:], " "))
		g.Printf("package %s", g.Pkg.Name)
		g.Printf("\n\n")
		for _, im := range m.File.imp {
			g.Printf("import %s\n", im)
		}

		g.Printf("\n")
		g.Printf("type %s struct {\n", m.Name)
		g.Printf("}\n\n")

		for _, f := range m.Func {
			log.Println(f.Sql.Action)
			switch f.Sql.Action {
			case "select":
				g.Select(m, f)
				break
			case "delete":
				g.Delete(m, f)
				break
			}
		}

		src := g.Format()
		dir := filepath.Dir(".")
		baseName := fmt.Sprintf("%s.go", utils.LowerFirstWord(m.Name))
		outputName := filepath.Join(dir, baseName)
		err := ioutil.WriteFile(outputName, src, 0644)
		if err != nil {
			log.Fatalf("writing output: %s", err)
		}
	}

}

func (g *Generator) Select(m Mapper, f Func) {
	g.Printf("func (this *%s) %s(", m.Name, f.Name)
	log.Println(f.Param)
	if f.Param.Name != "" {
		g.Printf("param *%s", f.Param.Name)
	}
	g.Printf(")")
	if len(f.Results) == 1 {
		switch f.Results[0].Type {
		case "array":
			g.Printf("[]")
			break
		case "star":
			g.Printf("*")
			break
		}
		g.Printf(f.Results[0].Name)
	}
	g.Printf("{\n")
	g.Printf(`rows, err := db.Query("%s"`, f.Tag)

	if f.Param.Name != "" {
		for _, p := range f.Sql.Params {
			g.Printf(",param.%s", p)
		}
	}
	g.Printf(")\n")
	g.Printf(`if err != nil {
								log.Println(err.Error())
								return nil
								}
								`)
	g.Printf(`defer rows.Close()`)
	g.Printf("\n")

	if f.Results[0].Type == "array" {
		g.Printf(`results:=make([]%s,0)`, f.Results[0].Name)
	}
	g.Printf("\n")

	if f.Results[0].Type == "array" {
		g.Printf("for ")
	} else {
		g.Printf("if ")
	}
	g.Printf(`rows.Next() {
								temp:=%s{}
								`, f.Results[0].Name)
	g.Printf("rows.Scan(")

	for i, s := range f.Sql.Fields {
		g.Printf("&temp.%s", s)
		if i != len(f.Sql.Fields)-1 {
			g.Printf(",")
		}

	}

	g.Printf(")")
	g.Printf("\n")
	if f.Results[0].Type == "array" {
		g.Printf(`results = append(results, temp)`)
	} else {
		g.Printf("return &temp")
	}
	g.Printf("\n}\n")
	if f.Results[0].Type == "array" {
		g.Printf("return results")
	} else {
		g.Printf("return nil")
	}
	g.Printf("}")
	g.Printf("\n\n")
}

func (g *Generator) Delete(m Mapper, f Func) {
	g.Printf("func (this *%s) %s(", m.Name, f.Name)
	if f.Param.Name != "" {
		g.Printf("param *%s", f.Param.Name)
	}
	g.Printf(")")
	if len(f.Results) == 1 {
		g.Printf(f.Results[0].Name)
	}
	g.Printf("{\n")
	g.Printf(`stmt, err := db.Prepare("%s")
					if err != nil {
						return err
					}
					defer stmt.Close()
					`, f.Tag)
	g.Printf("res, err := stmt.Exec(")
	if f.Param.Name != "" {
		for i, p := range f.Sql.Params {
			g.Printf("param.%s", p)
			if i != len(f.Sql.Params)-1 {
				g.Printf(",")
			}
		}
	}
	g.Printf(")\n")

	g.Printf(`if err != nil {
						log.Println(err.Error())
						return err
					}
					return nil
					`)
	g.Printf("}")
	g.Printf("\n\n")
}
