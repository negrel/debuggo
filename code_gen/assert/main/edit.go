package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/token"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/negrel/asttk/pkg/inspector"
	"github.com/negrel/asttk/pkg/parse"
	"github.com/negrel/asttk/pkg/utils"
)

func init() {
	editor = inspector.New(
		renameFuncWrapper(),
		removeTestingTInFuncDecl,
		removeTestingTInFuncCall,
		replaceTErrorfWithPanic,
		removeTTypeAssert,
	)
}

// editor is an ast inspector that generate the debug files.
var editor *inspector.Lead

func editFile(file *parse.GoFile) {
	// Start the inspection/edition of the AST
	editor.Inspect(file.AST())

	// Write the edited AST into the buffer
	buf := &bytes.Buffer{}
	err := file.Fprint(buf)
	if err != nil {
		log.Fatal(err)
	}

	// Write the buffer to the disk
	err = ioutil.WriteFile(
		filepath.Join("pkg", "assert", file.Name()),
		append([]byte("// +build assert\n\n"), buf.Bytes()...),
		0755,
	)
	if err != nil {
		log.Fatal(err)
	}
}

// ------------
// editor hooks
// ------------

func removeTestingTInFuncDecl(node ast.Node) bool {
	funcDecl, isFuncDecl := node.(*ast.FuncDecl)
	if !isFuncDecl {
		return true
	}

	fnType := funcDecl.Type
	i := -1
	for i != len(fnType.Params.List)-1 {
		i++
		param := fnType.Params.List[i]

		if fmt.Sprint(param.Type) != "TestingT" {
			continue
		}

		fnType.Params.List = append(fnType.Params.List[:i], fnType.Params.List[i+1:]...)
		i--
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

	i := -1
	for i != len(block.List)-1 {
		i++

		stmt := block.List[i]
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

		block.List = append(block.List[:i], block.List[i+1:]...)
		i--
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
