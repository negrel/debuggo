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

var skip = map[string]struct{}{
	"assertion_forward.go":  {},
	"forward_assertions.go": {},
	"doc.go":                {},
	"errors.go":             {},
}

func main() {
	pkg, err := parse.Package(
		filepath.Join("code_gen", "assert", "testify", "assert"),
		false,
	)
	if err != nil {
		log.Fatal(err)
	}

	editor := inspector.New(
		renameFuncWrapper(),
		removeTestingTInFuncDecl,
		removeTestingTInFuncCall,
		replaceTErrorfWithPanic,
		removeTTypeAssert,
	)

	findUnusedImports, removeUnusedImports := utils.RemoveUnusedImports()
	prodEditor := inspector.New(
		removeUnexportedDecls,
		renameFuncParams,
		removeFuncResult,
		removeFuncBody,
		renameFuncParams,
		findUnusedImports,
	)

	outputDir := filepath.Join("pkg", "assert")
	for _, file := range pkg.Files {
		if _, ok := skip[file.Name()]; ok {
			continue
		}

		editor.Inspect(file.AST())
		buf := &bytes.Buffer{}
		err = file.Fprint(buf)
		if err != nil {
			log.Fatal(err)
		}
		err = ioutil.WriteFile(
			filepath.Join(outputDir, file.Name()),
			append([]byte("// +build assert\n\n"), buf.Bytes()...),
			0755,
		)
		if err != nil {
			log.Fatal(err)
		}

		// Prod file
		prodFileName := addSuffix(file.Name(), ".prod")
		prodFile := &ast.File{
			Name:    ast.NewIdent(pkg.Name()),
			Decls:   file.AST().Decls,
			Imports: file.AST().Imports,
		}
		prodEditor.Inspect(prodFile)
		removeUnusedImports(prodFile)

		buf = &bytes.Buffer{}
		err = printer.Fprint(buf, token.NewFileSet(), prodFile)
		if err != nil {
			log.Fatal(err)
		}

		err = ioutil.WriteFile(
			filepath.Join(outputDir, prodFileName),
			append([]byte("// +build !assert\n\n"), buf.Bytes()...),
			0755,
		)
		if err != nil {
			log.Fatal(err)
		}
	}
}
