package generator

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/printer"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"golang.org/x/tools/go/packages"
)

// Generator is responsible for handling the generation
// process of the packages.
type Generator struct {
	packages   []*packages.Package
	commonPkgs []*packages.Package
	errors     []error
	outputDir  string
}

// New return a new debugging package generator.
func New(opts ...Option) (*Generator, error) {
	gen := &Generator{
		errors: make([]error, 0, 8),
	}

	for _, opt := range opts {
		err := opt(gen)
		if err != nil {
			return gen, err
		}
	}

	return gen, nil
}

func (gen *Generator) Error() error {
	if len(gen.errors) == 0 {
		return nil
	}

	err := fmt.Sprintf("%v error(s) occured during code generation :", len(gen.errors))
	for _, e := range gen.errors {
		err = fmt.Sprintf("%v\n	%v", err, e)
	}

	return fmt.Errorf(err)
}

// Start the generation process.
func (gen *Generator) Start() {
	gen.generatePackages()
}

func (gen *Generator) generatePackages() {
	var wg sync.WaitGroup
	wg.Add(len(gen.packages))

	for _, pkg := range gen.packages {
		go func(pkg *packages.Package) {
			gen.generateSinglePkg(pkg)
			wg.Done()
		}(pkg)
	}

	wg.Wait()
}

func (gen *Generator) generateSinglePkg(pkg *packages.Package) {
	err := mkdirAll(filepath.Join(gen.outputDir, pkg.Name))
	gen.reportError(err)

	filesHeader := gen.generateFilesPrefixFor(pkg)
	gen.addCommonFilesTo(pkg)

	for index, file := range pkg.Syntax {
		_, filename := filepath.Split(pkg.GoFiles[index])

		// Debug & Production
		for i := 0; i < 2; i++ {
			isProduction := i == 1

			if isProduction {
				filename = debugFileNameFor(filename)
				gen.editFile(file)
			}

			buf := bytes.NewBuffer(filesHeader)
			_ = printer.Fprint(buf, pkg.Fset, file)

			outputPath := filepath.Join(gen.outputDir, pkg.Name, filename)
			_ = ioutil.WriteFile(outputPath, buf.Bytes(), os.ModePerm)
		}
	}
}

func (gen *Generator) editFile(file *ast.File) {
	editor := newAstEditor(
		removePkgLevelFuncBodyOption,
		removeUnusedImportsOption,
	)

	editor.edit(file)
}

// TODO Add support for common subpackage.
func (gen *Generator) addCommonFilesTo(pkg *packages.Package) {
	for _, cmnPkg := range gen.commonPkgs {
		pkg.Syntax = append(pkg.Syntax, cmnPkg.Syntax...)
		pkg.GoFiles = append(pkg.GoFiles, cmnPkg.GoFiles...)
	}
}

func (gen *Generator) reportError(err error) {
	if err == nil {
		return
	}

	gen.errors = append(gen.errors, err)
}

func (gen *Generator) generateFilesPrefixFor(pkg *packages.Package) []byte {
	buf := &bytes.Buffer{}
	buf.WriteString("// Code generated by Debuggo. DO NOT EDIT.\n\n")
	buf.WriteString("// +build !" + pkg.Name + "\n\n")

	return buf.Bytes()
}

func mkdirAll(path string) error {
	if filepath.Ext(path) != "" {
		path, _ = filepath.Split(path)
	}

	return os.MkdirAll(path, os.ModePerm)
}

func debugFileNameFor(originalName string) string {
	filename := strings.Join(strings.Split(originalName, ".go"), "")
	filename += ".debug.go"

	return filename
}
