package main

import (
	"go/ast"
	"go/token"
)

func removeFuncBody(node ast.Node) (recursive bool) {
	recursive = true

	funcDecl, isFuncDecl := node.(*ast.FuncDecl)
	if !isFuncDecl {
		return
	}
	funcDecl.Body.List = []ast.Stmt{}

	return
}

func removeFuncResult(node ast.Node) bool {
	funcType, isFuncType := node.(*ast.FuncType)
	if !isFuncType {
		return true
	}

	funcType.Results = &ast.FieldList{List: []*ast.Field{}}

	return false
}

func removeUnexportedDecls(node ast.Node) bool {
	file, isFile := node.(*ast.File)
	if !isFile {
		return true
	}

	i := -1
	for i != len(file.Decls)-1 {
		i++
		decl := file.Decls[i]

		switch d := decl.(type) {
		case *ast.FuncDecl:
			if ast.IsExported(d.Name.Name) {
				continue
			}
		case *ast.GenDecl:
			if needGenDecl(d) {
				continue
			}
		}

		file.Decls = append(file.Decls[:i], file.Decls[i+1:]...)
		i--
	}
	return false
}

func needGenDecl(d *ast.GenDecl) bool {
	if d.Tok == token.IMPORT {
		return true
	}

	i := -1
	for _, spec := range d.Specs {
		i++

		switch s := spec.(type) {
		case *ast.TypeSpec:
			if ast.IsExported(s.Name.Name) {
				return true
			}

		case *ast.ValueSpec:
			j := -1
			for j != len(s.Names)-1 {
				j++
				name := s.Names[j]

				if ast.IsExported(name.Name) {
					continue
				}

				s.Names = append(s.Names[:j], s.Names[j+1:]...)
				j--
			}

			return len(s.Names) != 0
		}
	}

	return false
}

func renameFuncParams(node ast.Node) (recursive bool) {
	recursive = true

	funcType, isFuncType := node.(*ast.FuncType)
	if !isFuncType {
		return
	}

	noName := ast.NewIdent("_")
	for _, params := range funcType.Params.List {
		for i := range params.Names {
			params.Names[i] = noName
		}
	}

	return false
}
