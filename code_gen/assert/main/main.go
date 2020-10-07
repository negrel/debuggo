package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/negrel/asttk/pkg/edit"
	"github.com/negrel/asttk/pkg/inspector"
	"github.com/negrel/asttk/pkg/parse"
)

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

	prod := inspector.New(
		edit.ApplyOnTopDecl(removeBodyExportedFunc),
	)

	for _, file := range pkg.Files() {
		editor.Inspect(file.AST())

		path := filepath.Join("pkg", "assert", addSuffix(file.Name(), ".debug"))
		buildTag := []byte("// +build debuggo-assert\n\n")
		err = ioutil.WriteFile(path, append(buildTag, file.Byte()...), os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}

		// Production version
		prod.Inspect(file.AST())

		path = filepath.Join("pkg", "assert", addSuffix(file.Name(), ".prod"))
		buildTag = []byte("// +build !debuggo-assert\n\n")
		err = ioutil.WriteFile(path, append(buildTag, file.Byte()...), os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}
}
