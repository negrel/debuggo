package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

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

	for _, file := range pkg.Files() {
		editor.Inspect(file.AST())

		path := filepath.Join("pkg", "assert", file.Name())
		buildTag := []byte("// +build debuggo-assert\n\n")
		err = ioutil.WriteFile(path, append(buildTag, file.Byte()...), os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func min(a int, b int) int {
	if a < b {
		return a
	}

	return b
}
