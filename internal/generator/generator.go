package generator

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"golang.org/x/tools/go/packages"
)

// Generate start the build process of debugging package.
func Generate(opt Options) {
	gen := generator{
		outputDir: opt.Output,
	}
	pkgs := gen.parsePackages(opt.PkgPattern, opt.PkgTags)

	// Generate in a separate goroutine the
	// debugging packages.
	var wg sync.WaitGroup
	for _, pkg := range pkgs {
		wg.Add(1)
		go func(pkg *Package) {
			gen.generate(pkg)
			wg.Done()
		}(pkg)
	}

	wg.Wait()

	fmt.Println()

	for _, err := range gen.errors {
		fmt.Fprintln(os.Stderr, err)
	}
}

// generator is responsible for handling the generation
// process of the packages.
type generator struct {
	errors    []error
	outputDir string
}

// parsePackage find package and return an array of Package.
func (gen *generator) parsePackages(patterns []string, tags []string) []*Package {
	cnf := &packages.Config{
		Mode:       packages.LoadSyntax,
		Tests:      false,
		BuildFlags: []string{fmt.Sprintf("-tags=%s", strings.Join(tags, " "))},
	}

	pkgs, err := packages.Load(cnf, patterns...)
	if err != nil {
		log.Fatal(err)
	}

	// Package object to return
	rPkgs := make([]*Package, 0, len(pkgs))

	for _, pkg := range pkgs {
		// We don't generate the package if it contain any error.
		for _, pkgError := range pkg.Errors {
			gen.errors = append(gen.errors, errors.New(pkgError.Error()))
		}

		rPkgs = append(rPkgs, NewPackage(pkg))
	}

	return rPkgs
}

// Generate the given debugging package.
func (gen *generator) generate(pkg *Package) {
	var modes = [2]bool{false, true}

	// Creating OutputDir
	pkgPath := ""
	if filepath.IsAbs(gen.outputDir) {
		pkgPath = gen.outputDir
	} else {
		var err error
		pkgPath, err = filepath.Abs(gen.outputDir)
		if err != nil {
			gen.errors = append(gen.errors, err)
			return
		}
	}
	pkgPath = fmt.Sprintf("%v/%v", pkgPath, pkg.name)
	err := os.MkdirAll(pkgPath, os.ModePerm)
	if err != nil {
		gen.errors = append(gen.errors)
		return
	}

	for _, file := range pkg.files {
		for _, release := range modes {
			content, err := file.generateContent(release)
			if err != nil {
				gen.errors = append(gen.errors, err)
				continue
			}

			filename := file.name
			if !release {
				filename = strings.Join(strings.Split(filename, ".go"), "")
				filename += ".debug.go"
			}

			path := fmt.Sprintf("%v/%v", pkgPath, filename)

			err = ioutil.WriteFile(path, content, 0644)
			if err != nil {
				gen.errors = append(gen.errors, err)
				continue
			}
		}
	}
}
