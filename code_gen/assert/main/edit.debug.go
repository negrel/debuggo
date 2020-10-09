package main

import (
	"go/ast"
	"go/token"
	"strings"
)

type debugInspector struct {
	decls []ast.Decl
}

func newDebugInspector() *debugInspector {
	return &debugInspector{
		decls: make([]ast.Decl, 0),
	}
}

func (d *debugInspector) scanImportsDecl(node ast.Node) (recursive bool) {
	recursive = true

	if _, isFile := node.(*ast.File); isFile {
		d.decls = make([]ast.Decl, 0)
	}

	decl, isGenDecl := node.(*ast.GenDecl)
	if !isGenDecl {
		return
	}

	if decl.Tok != token.IMPORT {
		return
	}

	d.decls = append(d.decls, decl)

	return
}

func (d *debugInspector) scanExportedFuncDecl(node ast.Node) (recursive bool) {
	recursive = true

	if _, isFile := node.(*ast.File); isFile {
		d.decls = make([]ast.Decl, 0)
	}

	funcDecl, isFuncDecl := node.(*ast.FuncDecl)
	if !isFuncDecl {
		return
	}

	if !strings.HasPrefix(funcDecl.Name.Name, exportedFuncNamePrefix) {
		return
	}

	replaceFuncDecl := &ast.FuncDecl{
		Name: ast.NewIdent(funcDecl.Name.Name[len(exportedFuncNamePrefix):]),
		Doc:  funcDecl.Doc,
		Type: &ast.FuncType{
			Params: funcDecl.Type.Params,
			Results: &ast.FieldList{
				List: []*ast.Field{},
			},
		},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				&ast.ExprStmt{
					X: &ast.CallExpr{
						Fun:  funcDecl.Name,
						Args: extractArguments(funcDecl.Type.Params),
					},
				},
			},
		},
	}

	d.decls = append(d.decls, replaceFuncDecl)

	return
}

func extractArguments(fl *ast.FieldList) []ast.Expr {
	result := make([]ast.Expr, 0, len(fl.List))

	for _, field := range fl.List {
		for _, name := range field.Names {
			result = append(result, &ast.Ident{
				Name: name.Name,
			})
		}
	}

	return result
}
