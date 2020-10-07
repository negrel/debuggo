package main

import (
	"go/ast"
)

func removeBodyExportedFunc(node ast.Node) (recursive bool) {
	recursive = true

	funcDecl, isFuncDecl := node.(*ast.FuncDecl)
	if !isFuncDecl || !ast.IsExported(funcDecl.Name.Name) {
		return
	}

	funcDecl.Body.List = []ast.Stmt{}

	return
}
