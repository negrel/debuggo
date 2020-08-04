package generator

import (
	"go/ast"
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
	}
}

type removeUnusedImportsHook struct {
	allImports      []*ast.ImportSpec
	requiredImports []*ast.ImportSpec
}

func removeUnusedImportsOption(e *astEditor) {
	hook := &removeUnusedImportsHook{
		allImports:      make([]*ast.ImportSpec, 0, 8),
		requiredImports: make([]*ast.ImportSpec, 0, 8),
	}

	e.beforeEditHooks = append(e.beforeEditHooks, hook.beforeParseHook)
	e.nodeHooks = append(e.nodeHooks, hook.nodeHook)
	e.afterEditHooks = append(e.afterEditHooks, hook.afterParseHook)
}

func (r *removeUnusedImportsHook) beforeParseHook(file *ast.File) {
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

func (r *removeUnusedImportsHook) afterParseHook(file *ast.File) {
	file.Imports = r.requiredImports
}
