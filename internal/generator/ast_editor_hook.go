package generator

import (
	"fmt"
	"go/ast"
	"go/token"
	"strings"
)

var allAstEditorOptions = []astEditorOption{
	removePkgLevelFuncBodyOption,
	removeUnusedImportsOption,
	removeUnattachedCommentsOption,
}

func removePkgLevelFuncBodyOption(e *astEditor) {
	e.beforeEditHooks = append(e.beforeEditHooks, removePkgLevelFuncBody)
}

func removePkgLevelFuncBody(file *ast.File) {
	for _, d := range file.Decls {
		decl, ok := d.(*ast.FuncDecl)
		if !ok {
			continue
		}

		if isExported := ast.IsExported(decl.Name.Name); isExported {
			if returnSomething := decl.Type.Results != nil; returnSomething {
				panic(fmt.Sprintf("exported function '%v' shouldn't return any value", decl.Name.Name))
			}

			decl.Body.List = []ast.Stmt{}
		}
	}
}

type unusedImportsRemover struct {
	allImports      []*ast.ImportSpec
	requiredImports []ast.Spec
}

func removeUnusedImportsOption(e *astEditor) {
	hook := &unusedImportsRemover{
		allImports:      make([]*ast.ImportSpec, 0),
		requiredImports: make([]ast.Spec, 0),
	}

	e.beforeEditHooks = append(e.beforeEditHooks, hook.beforeEditHook)
	e.nodeHooks = append(e.nodeHooks, hook.nodeHook)
	e.afterEditHooks = append(e.afterEditHooks, hook.afterEditHook)
}

func (r *unusedImportsRemover) beforeEditHook(file *ast.File) {
	r.allImports = file.Imports
}

func (r *unusedImportsRemover) nodeHook(n ast.Node) (recursive bool) {
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

func (r *unusedImportsRemover) afterEditHook(file *ast.File) {
	for _, d := range file.Decls {
		decl, ok := d.(*ast.GenDecl)
		if !ok {
			continue
		}

		if decl.Tok != token.IMPORT {
			continue
		}

		decl.Specs = r.requiredImports
		break
	}
}

func removeUnattachedCommentsOption(e *astEditor) {
	e.afterEditHooks = append(e.afterEditHooks, removeUnattachedComments)
}

func removeUnattachedComments(file *ast.File) {
	file.Comments = nil
}
