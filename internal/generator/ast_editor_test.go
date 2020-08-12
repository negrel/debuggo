package generator

import (
	"bytes"
	"go/ast"
	"go/printer"
	"log"
	"testing"

	"github.com/google/go-cmp/cmp"
	"golang.org/x/tools/go/packages"
)

type astEditorTest struct {
	srcPkg      *packages.Package
	outPkg      *packages.Package
	shouldPanic bool
}

func newAstEditorTests(path string) (astEditorTests []astEditorTest) {
	srcPkgs, err := getAllPackagesIn("./_test/editor/" + path + "/src")
	if err != nil {
		log.Fatal(err)
	}
	outPkgs, err := getAllPackagesIn("./_test/editor/" + path + "/result")
	if err != nil {
		log.Fatal(err)
	}

	for _, srcPkg := range srcPkgs {
		outPkg := findPkgIn(srcPkg.Name, outPkgs)

		astEditorTests = append(astEditorTests, astEditorTest{
			srcPkg,
			outPkg,
			outPkg == nil,
		})
	}

	return astEditorTests
}

func findPkgIn(name string, pkgList []*packages.Package) *packages.Package {
	for _, pkg := range pkgList {
		if pkg.Name == name {
			return pkg
		}
	}

	return nil
}

func fileToString(pkg *packages.Package, file *ast.File) (string, error) {
	buf := &bytes.Buffer{}
	err := printer.Fprint(buf, pkg.Fset, file)

	return buf.String(), err
}

func assertPanic(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	f()
}

func TestPkgLevelFuncBodyRemover(t *testing.T) {
	editor := newAstEditor(removePkgLevelFuncBodyOption)

	for _, test := range newAstEditorTests("func_body") {
		for i, srcFile := range test.srcPkg.Syntax {
			if test.shouldPanic {
				assertPanic(t, func() {
					editor.edit(srcFile)
				})
				continue

			} else {
				editor.edit(srcFile)
			}

			outFile := test.outPkg.Syntax[i]

			got, err := fileToString(test.srcPkg, srcFile)
			if err != nil {
				t.Fatal(err)
			}

			expected, err := fileToString(test.outPkg, outFile)
			if err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(expected, got); diff != "" {
				t.Errorf("editor doesn't generate the expected result:\n%s", diff)
			}
		}
	}
}

func TestUnusedImportsRemover(t *testing.T) {
	editor := newAstEditor(removeUnusedImportsOption)

	for _, test := range newAstEditorTests("unused_imports") {
		for i, srcFile := range test.srcPkg.Syntax {
			outFile := test.outPkg.Syntax[i]

			if test.shouldPanic {
				assertPanic(t, func() {
					editor.edit(srcFile)
				})

				continue
			} else {
				editor.edit(srcFile)
			}

			got, err := fileToString(test.srcPkg, srcFile)
			if err != nil {
				t.Fatal(err)
			}

			expected, err := fileToString(test.outPkg, outFile)
			if err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(expected, got); diff != "" {
				t.Errorf("editor doesn't generate the expected result:\n%s", diff)
			}
		}
	}
}

func TestUnattachedCommentsRemover(t *testing.T) {
	editor := newAstEditor(removeUnattachedCommentsOption)

	for _, test := range newAstEditorTests("unattached_comments") {
		for i, srcFile := range test.srcPkg.Syntax {
			editor.edit(srcFile)
			outFile := test.outPkg.Syntax[i]

			got := &bytes.Buffer{}
			err := printer.Fprint(got, test.srcPkg.Fset, srcFile)
			if err != nil {
				t.Fatal(err)
			}

			expected := &bytes.Buffer{}
			err = printer.Fprint(expected, test.outPkg.Fset, outFile)
			if err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(expected.String(), got.String()); diff != "" {
				t.Errorf("editor doesn't generate the expected result:\n%s", diff)
			}
		}
	}
}
