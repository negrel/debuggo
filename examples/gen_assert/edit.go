package main

import (
	"fmt"
	"go/ast"
	"go/token"
)

func removeTestingTInFuncDecl(node ast.Node) bool {
	fnType, isFuncType := node.(*ast.FuncType)
	if !isFuncType {
		return true
	}

	for i, param := range fnType.Params.List {
		if fmt.Sprint(param.Type) != "TestingT" {
			continue
		}

		fnType.Params.List = append(fnType.Params.List[:i], fnType.Params.List[i+1:]...)
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

func removeTTypeAssert(node ast.Node) (recursive bool) {
	recursive = true

	block, isBlockStmt := node.(*ast.BlockStmt)
	if !isBlockStmt {
		return
	}

	for i, stmt := range block.List {
		ifStmt, isIfStmt := stmt.(*ast.IfStmt)
		if !isIfStmt {
			return
		}

		assignStmt, isAssignStmt := ifStmt.Init.(*ast.AssignStmt)
		if !isAssignStmt {
			return
		}

		typeAssertExpr, isTypeAssertExpr := assignStmt.Rhs[0].(*ast.TypeAssertExpr)
		if !isTypeAssertExpr {
			return
		}

		ident, isIdent := typeAssertExpr.X.(*ast.Ident)
		if !isIdent || ident.Name != "t" {
			return
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
