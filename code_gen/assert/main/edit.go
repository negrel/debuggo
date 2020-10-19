package main

import (
	"fmt"
	"go/ast"
	"go/token"

	"github.com/negrel/asttk/pkg/inspector"
	"github.com/negrel/asttk/pkg/utils"
)

func removeTestingTInFuncDecl(node ast.Node) bool {
	funcDecl, isFuncDecl := node.(*ast.FuncDecl)
	if !isFuncDecl {
		return true
	}

	fnType := funcDecl.Type
	for i, param := range fnType.Params.List {
		if fmt.Sprint(param.Type) != "TestingT" {
			continue
		}

		length := min(i+1, len(fnType.Params.List))
		fnType.Params.List = append(fnType.Params.List[:i], fnType.Params.List[length:]...)
	}

	return false
}

func removeTestingTInFuncCall(node ast.Node) (recursive bool) {
	recursive = true

	callExpr, isCallExpr := node.(*ast.CallExpr)
	if !isCallExpr {
		return
	}

	for i, arg := range callExpr.Args {
		ident, isIdent := arg.(*ast.Ident)
		if !isIdent {
			continue
		}

		if ident.Name != "t" {
			continue
		}
		callExpr.Args = append(callExpr.Args[:i], callExpr.Args[i+1:]...)
	}

	return
}

const exportedFuncNamePrefix = "debuggoGen_"

func funcFilter(name string) (replaceName string, ok bool) {
	if !ast.IsExported(name) {
		return "", false
	}

	replaceName = exportedFuncNamePrefix + name

	return replaceName, true
}

func renameFuncWrapper() inspector.Inspector {
	_, renameFuncCall := utils.RenameFunc(funcFilter)
	return inspector.Lieutenant(replaceExportedFunc, renameFuncCall)
}

func replaceExportedFunc(node ast.Node) bool {

	file, isFile := node.(*ast.File)
	if !isFile {
		return true
	}

	for i, decl := range file.Decls {
		funcDecl, isFuncDecl := decl.(*ast.FuncDecl)
		if !isFuncDecl {
			continue
		}

		newName, ok := funcFilter(funcDecl.Name.Name)
		if !ok || newName == "" {
			continue
		}

		file.Decls[i] = &ast.FuncDecl{
			Name: funcDecl.Name,
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

		funcDecl.Name = ast.NewIdent(newName)
		file.Decls = append(file.Decls, funcDecl)
	}
	return false
}

func removeTTypeAssert(node ast.Node) (recursive bool) {
	recursive = true

	block, isBlockStmt := node.(*ast.BlockStmt)
	if !isBlockStmt {
		return
	}

	for i, stmt := range block.List {
		ifStmt, isIfStmt := stmt.(*ast.IfStmt)
		if !isIfStmt {
			continue
		}

		assignStmt, isAssignStmt := ifStmt.Init.(*ast.AssignStmt)
		if !isAssignStmt {
			continue
		}

		typeAssertExpr, isTypeAssertExpr := assignStmt.Rhs[0].(*ast.TypeAssertExpr)
		if !isTypeAssertExpr {
			return
		}

		ident, isIdent := typeAssertExpr.X.(*ast.Ident)
		if !isIdent || ident.Name != "t" {
			continue
		}

		index := min(i+1, len(block.List)-1)
		block.List = append(block.List[:i], block.List[index:]...)
	}

	return
}

func replaceTErrorfWithPanic(node ast.Node) (recursive bool) {
	recursive = true

	callExpr, isCallExpr := node.(*ast.CallExpr)
	if !isCallExpr {
		return
	}

	selector, isSelectorExpr := callExpr.Fun.(*ast.SelectorExpr)
	if !isSelectorExpr {
		return
	}

	if fmt.Sprint(selector.X) != "t" ||
		fmt.Sprint(selector.Sel) != "Errorf" {
		return
	}

	callExpr.Fun = ast.NewIdent("panic")

	arg := &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X:   ast.NewIdent("fmt"),
			Sel: ast.NewIdent("Sprint"),
		},
		Args: callExpr.Args,
	}

	callExpr.Args = []ast.Expr{arg}
	callExpr.Ellipsis = token.NoPos

	return
}
