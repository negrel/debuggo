package main

import (
	"fmt"
	"go/ast"
	"log"
	"path/filepath"
	"reflect"

	"github.com/negrel/asttk/pkg/inspector"
	"github.com/negrel/asttk/pkg/parse"
)

func main() {
	pkg, err := parse.Package(filepath.Join("code_gen", "log", "log"), false)
	if err != nil {
		log.Fatal(err)
	}

	// renamer := inspector.Lieutenant(
	// _, renameFuncCall := utils.RenameFunc(exportedFuncFilter)
	// )
	editor := inspector.New(
		// renameFuncCall,
		func(node ast.Node) (recursive bool) {
			recursive = true

			if node == nil {
				return
			}

			comment, isComment := node.(*ast.Comment)
			if !isComment {
				return
			}
			fmt.Println(comment)
			*comment = ast.Comment{}

			nodeValue := reflect.ValueOf(node)
			nodeValue = reflect.Indirect(nodeValue)

			funcDecl, isFuncDecl := node.(*ast.File)
			if !isFuncDecl {
				return
			}

			log.Println("func decl", funcDecl.Comments)

			doc := nodeValue.FieldByName("Doc")
			comments := nodeValue.FieldByName("Comments")

			log.Println(nodeValue, doc, comments)

			return
		},
	)

	for _, file := range pkg.Files {
		if file.Name() != "logger.go" {
			continue
		}

		editor.Inspect(file.AST())

		data, err := file.Bytes()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(data))
	}
}
