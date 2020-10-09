package main

import (
	"go/ast"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/negrel/asttk/pkg/edit"
	"github.com/negrel/asttk/pkg/inspector"
	"github.com/negrel/asttk/pkg/parse"
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

	di := newDebugInspector()
	editor := inspector.New(
		renameFuncWrapper(),
		removeTestingTInFuncDecl,
		removeTestingTInFuncCall,
		di.scanExportedFuncDecl,
		di.scanImportsDecl,
		replaceTErrorfWithPanic,
		removeTTypeAssert,
	)

	findUnusedImports, removeUnusedImports := edit.RemoveUnusedImports()
	prodEditor := inspector.New(
		findUnusedImports,
		removeBodyExportedFunc,
	)

	outputDir := filepath.Join("pkg", "assert")
	for _, file := range pkg.Files {
		if _, ok := skip[file.Name()]; ok {
			continue
		}

		editor.Inspect(file.AST())
		err = file.WriteFile(filepath.Join(outputDir, file.Name()))
		if err != nil {
			log.Fatal(err)
		}

		// Debug file
		debugFile := parse.NewGoFile(addSuffix(file.Name(), ".debug"), &ast.File{
			Name:    ast.NewIdent(pkg.Name()),
			Decls:   di.decls,
			Imports: file.AST().Imports,
		})
		ast.Inspect(debugFile.AST(), findUnusedImports)
		removeUnusedImports(debugFile.AST())

		data, err := debugFile.Bytes()
		if err != nil {
			log.Fatal(err)
		}
		err = ioutil.WriteFile(
			filepath.Join(outputDir, debugFile.Name()),
			append([]byte("// +build assert\n\n"), data...),
			0755,
		)
		if err != nil {
			log.Fatal(err)
		}

		// Prod file
		prodFile := parse.NewGoFile(addSuffix(file.Name(), ".prod"), debugFile.AST())
		prodEditor.Inspect(prodFile.AST())
		removeUnusedImports(debugFile.AST())

		data, err = prodFile.Bytes()
		if err != nil {
			log.Fatal(err)
		}
		err = ioutil.WriteFile(
			filepath.Join(outputDir, prodFile.Name()),
			append([]byte("// +build !assert\n\n"), data...),
			0755,
		)
		if err != nil {
			log.Fatal(err)
		}
	}
}
