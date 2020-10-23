package main

import (
	"bytes"
	"go/ast"
	"go/printer"
	"go/token"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/negrel/asttk/pkg/inspector"
	"github.com/negrel/asttk/pkg/parse"
	"github.com/negrel/asttk/pkg/utils"
)

func init() {
	var findUnusedImports inspector.Inspector
	findUnusedImports, removeUnusedImports = utils.RemoveUnusedImports()

	prodEditor = inspector.New(
		removeUnexportedDecls,
		renameFuncParams,
		removeFuncResult,
		removeFuncBody,
		findUnusedImports,
	)
}

var prodEditor *inspector.Lead
var removeUnusedImports func(file *ast.File)

func editProdFile(file *parse.GoFile) {
	// Start the inspection/edition of the AST
	prodFile := &ast.File{
		Name:    file.AST().Name,
		Decls:   file.AST().Decls,
		Imports: file.AST().Imports,
	}
	prodEditor.Inspect(prodFile)
	removeUnusedImports(prodFile)

	// Write the edited AST into the buffer
	buf := &bytes.Buffer{}
	err := printer.Fprint(buf, token.NewFileSet(), prodFile)
	if err != nil {
		log.Fatal(err)
	}

	// Write the buffer to the disk
	prodFileName := addSuffix(file.Name(), ".prod")
	err = ioutil.WriteFile(
		filepath.Join("pkg", "assert", prodFileName),
		append([]byte("// +build !assert\n\n"), buf.Bytes()...),
		0755,
	)
	if err != nil {
		log.Fatal(err)
	}
}

// -----------------
// prod editor hooks
// -----------------

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
