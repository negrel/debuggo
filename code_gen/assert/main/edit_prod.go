package main

import (
	"go/ast"
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

func removeUnexportedFunc(node ast.Node) bool {
	file, isFile := node.(*ast.File)
	if !isFile {
		return true
	}

	i := -1
	for i != len(file.Decls)-1 {
		i++
		decl := file.Decls[i]

		funcDecl, isFuncDecl := decl.(*ast.FuncDecl)
		if !isFuncDecl {
			continue
		}

		if ast.IsExported(funcDecl.Name.Name) {
			continue
		}

		file.Decls = append(file.Decls[:i], file.Decls[i+1:]...)
		i--
	}
	return false
}
