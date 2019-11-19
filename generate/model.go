package generate

import (
	"bytes"
	"go/ast"
	"go/types"
	"gobatis/sqlparser"
)

// Generator holds the state of the analysis. Primarily used to buffer
// the output for format.Source.
type Generator struct {
	Buf bytes.Buffer // Accumulated output.
	Pkg *Package     // Package we are scanning.

	TrimPrefix  string
	LineComment bool
}

// Value represents a declared constant.
type Mapper struct {
	File *File
	Name string
	Func []Func
}

type Func struct {
	Name    string
	Results []Result
	Param   Param
	Tag     string
	Sql     *sqlparser.SqlSynx
}
type Param struct {
	Name string
}

type Result struct {
	Name    string
	IsArray bool
}

// File holds a single parsed file and associated data.
type File struct {
	imp  []string
	pkg  *Package  // Package to which this file belongs.
	file *ast.File // Parsed AST.
	// These fields are reset for each type being generated.
	typeName string   // Name of the constant type.
	mappers  []Mapper // Accumulator for constant values of that type.

	trimPrefix  string
	lineComment bool
}

type Package struct {
	Name  string
	defs  map[*ast.Ident]types.Object
	files []*File
}
