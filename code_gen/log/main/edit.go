package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	"github.com/negrel/asttk/pkg/inspector"
	"github.com/negrel/asttk/pkg/parse"
)

var logLevelsName = []string{
	"Panic",
	"Fatal",
	"Error",
	"Warn",
	"Info",
	"Debug",
	"Trace",
}

func editFile(file *parse.GoFile, logLevel int) {
	// Start the inspection/edition of the AST
	filter := getFilter(logLevelsName[logLevel])
	inspector.New(
		removeFuncBody(filter),
	).Inspect(file.AST())

	// Write the edited AST into the buffer
	buf := &bytes.Buffer{}
	err := file.Fprint(buf)
	if err != nil {
		log.Fatal(err)
	}

	// Write buffer to the disk
	buildTag := strings.ToLower(logLevelsName[logLevel])
	fileName := filepath.Join("pkg", "log", addSuffix(file.Name(), "."+buildTag))
	err = ioutil.WriteFile(
		fileName,
		append([]byte(fmt.Sprintf("// +build %v\n\n", strings.ToLower(buildTag))), buf.Bytes()...),
		0755,
	)
	if err != nil {
		log.Fatal(err)
	}
}

// ------------
// editor hooks
// ------------

type filter func(name string) bool

func getFilter(level string) filter {
	i := -1
	for _, logLevel := range logLevelsName {
		i++
		if logLevel == level {
			break
		}
	}

	return func(name string) bool {
		for j, logLevel := range logLevelsName {
			if j == i+1 {
				break
			}

			if strings.HasPrefix(name, logLevel) {
				return true
			}
		}

		return false
	}
}

func removeFuncBody(filter filter) inspector.Inspector {
	return func(node ast.Node) (recursive bool) {
		recursive = true

		funcDecl, isFuncDecl := node.(*ast.FuncDecl)
		if !isFuncDecl {
			return
		}

		isNeeded := filter(funcDecl.Name.Name)
		if !isNeeded {
			funcDecl.Body.List = []ast.Stmt{}
		}

		return false
	}
}
