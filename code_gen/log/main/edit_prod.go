package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/token"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	"github.com/negrel/asttk/pkg/inspector"
	"github.com/negrel/asttk/pkg/parse"
	"github.com/negrel/asttk/pkg/utils"
)

func prodBuildTags() string {
	buildTags := strings.Join(logLevelsName,
		",!",
	)

	return "!" + strings.ToLower(buildTags[:])
}

func editProdFile(file *parse.GoFile) {
	findUnusedImports, removeUnusedImports := utils.RemoveUnusedImports()
	inspector.New(
		removeAllFuncBody,
		removeUnexportedDecls,
		findUnusedImports,
	).Inspect(file.AST())
	removeUnusedImports(file.AST())

	// Write the edited AST into the buffer
	buf := &bytes.Buffer{}
	err := file.Fprint(buf)
	if err != nil {
		log.Fatal(err)
	}

	// Write buffer to the disk
	fileName := filepath.Join("pkg", "log", file.Name())
	err = ioutil.WriteFile(
		fileName,
		append([]byte(fmt.Sprintf("// +build %v\n\n", prodBuildTags())), buf.Bytes()...),
		0755,
	)
	if err != nil {
		log.Fatal(err)
	}
}

// ------------
// editor hooks
// ------------

func removeAllFuncBody(node ast.Node) (recursive bool) {
	recursive = true

	funcDecl, isFuncDecl := node.(*ast.FuncDecl)
	if !isFuncDecl {
		return
	}

	funcDecl.Body.List = []ast.Stmt{}

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
