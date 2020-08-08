package generator

import (
	"fmt"
	"go/ast"
	"go/token"
	"strings"
)

func removePkgLevelFuncBodyOption(e *astEditor) {
	e.beforeEditHooks = append(e.beforeEditHooks, removePkgLevelFuncBody)
}

func removePkgLevelFuncBody(file *ast.File) {
	for _, d := range file.Decls {
		decl, ok := d.(*ast.FuncDecl)
		if !ok {
			continue
		}

		decl.Body.List = []ast.Stmt{}

		if isExported, returnSomething := ast.IsExported(decl.Name.Name), decl.Type.Results != nil; isExported && returnSomething {
			panic(fmt.Sprintf("exported function '%v' shouldn't return any value", decl.Name.Name))
		}
	}
}

type removeUnusedImportsHook struct {
	allImports      []*ast.ImportSpec
	requiredImports []ast.Spec
}

func removeUnusedImportsOption(e *astEditor) {
	hook := &removeUnusedImportsHook{
		allImports:      make([]*ast.ImportSpec, 0),
		requiredImports: make([]ast.Spec, 0),
	}

	e.beforeEditHooks = append(e.beforeEditHooks, hook.beforeEditHook)
	e.nodeHooks = append(e.nodeHooks, hook.nodeHook)
	e.afterEditHooks = append(e.afterEditHooks, hook.afterEditHook)
}

func (r *removeUnusedImportsHook) beforeEditHook(file *ast.File) {
	r.allImports = file.Imports
}

func (r *removeUnusedImportsHook) nodeHook(n ast.Node) (recursive bool) {
	recursive = true

	ident, ok := n.(*ast.Ident)
	if !ok {
		return
	}

	for _, _import := range r.allImports {
		// Storing package identifier (name or last folder name in path)
		var name string
		if identifier := _import.Name; identifier != nil {
			name = identifier.Name
		} else {
			slice := strings.Split(_import.Path.Value, "/")
			name = slice[len(slice)-1]
			name = strings.Trim(name, "\"")
		}

		if name == ident.Name {
			r.requiredImports = append(r.requiredImports, _import)
		}
	}

	return
}

func (r *removeUnusedImportsHook) afterEditHook(file *ast.File) {
	for _, d := range file.Decls {
		decl, ok := d.(*ast.GenDecl)
		if !ok {
			continue
		}

		if decl.Tok != token.IMPORT {
			continue
		}

		decl.Specs = r.requiredImports
	}
}

func sanitizeOption(e *astEditor) {
	e.afterEditHooks = append(e.afterEditHooks, sanitizeHook)
}

func sanitizeHook(file *ast.File) {
	ast.FilterFile(file, func(decl string) bool {
		return true
	})
}
