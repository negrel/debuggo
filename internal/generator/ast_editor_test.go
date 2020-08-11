package generator

import (
	"bytes"
	"go/printer"
	"log"
	"testing"

	"github.com/google/go-cmp/cmp"
	"golang.org/x/tools/go/packages"
)

type astEditorTest struct {
	srcPkg *packages.Package
	outPkg *packages.Package
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

	for i, srcPkg := range srcPkgs {
		outPkg := outPkgs[i]

		if srcPkg.Name != outPkg.Name {
			log.Fatalf("Invalid test package name are different\n%v != %v", srcPkg.Name, outPkg.Name)
		}

		astEditorTests = append(astEditorTests, astEditorTest{
			srcPkg,
			outPkg,
		})
	}

	return astEditorTests
}

func TestPkgLevelFuncBodyRemover(t *testing.T) {
	editor := newAstEditor(removePkgLevelFuncBodyOption)

	for _, test := range newAstEditorTests("func_body") {
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

func TestUnusedImportsRemover(t *testing.T) {
	editor := newAstEditor(removeUnusedImportsOption)

	for _, test := range newAstEditorTests("unused_imports") {
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
