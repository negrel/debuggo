package generator

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/printer"
	"log"
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"golang.org/x/tools/go/packages"
)

type astEditorTest struct {
	srcPkg *packages.Package
	outPkg *packages.Package
}

func newAstEditorTests(path string) (astEditorTests []astEditorTest) {
	srcPkgs, err := getAllPackagesIn("./_test/editor/" + path + "/success/src")
	if err != nil {
		log.Fatal(err)
	}
	outPkgs, err := getAllPackagesIn("./_test/editor/" + path + "/success/result")
	if err != nil {
		log.Fatal(err)
	}

	for i, srcPkg := range srcPkgs {
		outPkg := outPkgs[i]

		astEditorTests = append(astEditorTests, astEditorTest{
			srcPkg,
			outPkg,
		})
	}

	return astEditorTests
}

func newPanicAstEditorTests(path string) (astEditorTests []astEditorTest) {
	srcPkgs, err := getAllPackagesIn("./_test/editor/" + path + "/panic/")
	if err != nil {
		log.Fatal(err)
	}

	for _, srcPkg := range srcPkgs {
		astEditorTests = append(astEditorTests, astEditorTest{
			srcPkg,
			nil,
		})
	}

	return astEditorTests
}

func (t astEditorTest) shouldPanic() bool {
	return strings.Contains(t.srcPkg.PkgPath, "/panic/")
}

func getFunctionName(i interface{}) string {
	fullname := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	s := strings.Split(fullname, ".")

	return s[len(s)-1]
}

func fileToString(pkg *packages.Package, file *ast.File) (string, error) {
	buf := &bytes.Buffer{}
	err := printer.Fprint(buf, pkg.Fset, file)

	return buf.String(), err
}

func assertPanic(f func()) (panicked bool) {
	defer func() {
		if r := recover(); r != nil {
			panicked = true
		}
	}()
	f()

	return
}

func TestAllAstEditorOptions(t *testing.T) {
	for _, option := range allAstEditorOptions {
		testAstEditorOption(t, option)
	}
}

func testAstEditorOption(t *testing.T, option astEditorOption) {
	editor := newAstEditor(option)
	optionName := getFunctionName(option)

	astEditorTests := append(newAstEditorTests(optionName), newPanicAstEditorTests(optionName)...)
	for _, test := range astEditorTests {
		err := testAstEditor(editor, test)
		if err != nil {
			t.Errorf("%v option failed a test:\n%v", optionName, err)
		}
	}
}

func testAstEditor(editor FileEditor, test astEditorTest) error {
	shouldPanic := test.shouldPanic()

	for i, srcFile := range test.srcPkg.Syntax {
		if shouldPanic {
			notPanicked := !assertPanic(func() {
				editor.edit(srcFile)
			})

			if notPanicked {
				return fmt.Errorf("%v/%v didn't panic as expected", test.srcPkg.Name, srcFile.Name.Name)
			}

			continue
		}

		outFile := test.outPkg.Syntax[i]

		editor.edit(srcFile)

		got, err := fileToString(test.srcPkg, srcFile)
		if err != nil {
			return err
		}

		expected, err := fileToString(test.outPkg, outFile)
		if err != nil {
			return err
		}

		if diff := cmp.Diff(expected, got); diff != "" {
			return fmt.Errorf("editor doesn't generate the expected result:\n%s", diff)
		}
	}

	return nil
}
