package generator

import (
	"go/token"
	"go/types"
	"path/filepath"

	"golang.org/x/tools/go/packages"
)

// Package define a parsed package
type Package struct {
	path  string
	name  string
	files []*file
	fset  *token.FileSet
	scope *types.Scope
}

// NewPackage returns a new Package for the given package path and name
func NewPackage(pkg *packages.Package) *Package {
	// The package object to return
	rPkg := &Package{
		scope: pkg.Types.Scope(),
		name:  pkg.Name,
		path:  pkg.PkgPath,
		files: make([]*file, len(pkg.GoFiles)),
		fset:  pkg.Fset,
	}

	// The file of the package.
	for i, fileAST := range pkg.Syntax {

		// Computing file name.
		dir, filename := filepath.Split(pkg.GoFiles[i])

		rPkg.files[i] = &file{
			pkg:  rPkg,
			ast:  fileAST,
			name: filename,
			dir:  dir,
		}
	}

	return rPkg
}
