package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/negrel/debuggo/pkg/inspector"
	"github.com/negrel/debuggo/pkg/parse"
)

func main() {
	pkg, err := parse.Package(
		filepath.Join("examples", "gen_assert", "testify", "assert"),
		//filepath.Join("pkg", "assert"),
		false,
	)
	if err != nil {
		log.Fatal(err)
	}

	editor := inspector.New(
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

//func main() {
//	file, err := parse.File(
//		filepath.Join("examples", "gen_assert", "testify", "assert", "assertions.go"),
//		//filepath.Join("pkg", "assert", "assertions.go"),
//	)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	editor := inspector.New(
//		removeTestingTInFuncDecl,
//		removeTestingTInFuncCall,
//		replaceTErrorfWithPanic,
//		removeTTypeAssert,
//	)
//
//	editor.Inspect(file.AST())
//
//	path := filepath.Join("pkg", "assert", "assertions.go")
//	err = ioutil.WriteFile(path, file.Byte(), os.ModePerm)
//	if err != nil {
//		log.Fatal(err)
//	}
//}
