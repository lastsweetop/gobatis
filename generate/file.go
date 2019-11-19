package generate

import (
	"go/ast"
	"go/token"
	"gobatis/sqlparser"
	"log"
	"reflect"
	"strings"
)

// genDecl processes one declaration clause.
func (f *File) genDecl(node ast.Node) bool {
	decl, ok := node.(*ast.GenDecl)
	if ok && decl.Tok == token.IMPORT {
		for _, spec := range decl.Specs {
			ispec := spec.(*ast.ImportSpec)
			f.imp = append(f.imp, ispec.Path.Value)
		}
		return false
	}

	if ok && decl.Tok == token.TYPE {
		for _, spec := range decl.Specs {
			tspec := spec.(*ast.TypeSpec)
			if strings.Index(tspec.Name.Name, "Mapper") == len(tspec.Name.Name)-6 {
				continue
			}

			v := Mapper{
				Name: tspec.Name.Name + "Mapper",
				Func: make([]Func, 0),
				File: f,
			}
			var param Param
			sspec := tspec.Type.(*ast.StructType)
			for _, field := range sspec.Fields.List {
				ft := field.Type.(*ast.FuncType)

				for _, p := range ft.Params.List {
					s, isStar := p.Type.(*ast.StarExpr)
					if isStar {
						x := s.X.(*ast.SelectorExpr)
						param = Param{
							Name: x.X.(*ast.Ident).Name + "." + x.Sel.Name,
						}
					}
				}

				result := make([]Result, 0)
				for _, r := range ft.Results.List {
					a, isArray := r.Type.(*ast.ArrayType)
					s, isStar := r.Type.(*ast.StarExpr)
					if isArray {
						elt := a.Elt.(*ast.SelectorExpr)
						result = append(result, Result{
							Name:    elt.X.(*ast.Ident).Name + "." + elt.Sel.Name,
							IsArray: true,
						})
					}
					if isStar {
						x := s.X.(*ast.SelectorExpr)
						result = append(result, Result{
							Name:    x.X.(*ast.Ident).Name + "." + x.Sel.Name,
							IsArray: false,
						})
					}
				}
				tag := reflect.StructTag(field.Tag.Value[1 : len(field.Tag.Value)-1])
				sql := sqlparser.Parser(tag.Get("batis")[:len(tag.Get("batis"))])

				v.Func = append(v.Func, Func{
					Name:    field.Names[0].Name,
					Results: result,
					Tag:     tag.Get("batis"),
					Sql:     sql,
					Param:   param,
				})
			}

			f.mappers = append(f.mappers, v)
			log.Println(spec)
		}
		return false
	}
	return true
}
