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
	"github.com/negrel/debuggo/internal/editor"
	"golang.org/x/tools/go/packages"
)

type astEditorHookTest struct {
	srcPkg *packages.Package
	outPkg *packages.Package
}

func newAstEditorHookTests(hook string) (astEditorTests []astEditorHookTest) {
	srcPkgs, err := getAllPackagesIn("./_test/editor_option/" + hook + "/success/src")
	if err != nil {
		log.Fatal(err)
	}
	outPkgs, err := getAllPackagesIn("./_test/editor_option/" + hook + "/success/result")
	if err != nil {
		log.Fatal(err)
	}

	for i, srcPkg := range srcPkgs {
		outPkg := outPkgs[i]

		astEditorTests = append(astEditorTests, astEditorHookTest{
			srcPkg,
			outPkg,
		})
	}

	return astEditorTests
}

func newPanicAstEditorHookTests(path string) (astEditorTests []astEditorHookTest) {
	srcPkgs, err := getAllPackagesIn("./_test/editor/" + path + "/panic/")
	if err != nil {
		log.Fatal(err)
	}

	for _, srcPkg := range srcPkgs {
		astEditorTests = append(astEditorTests, astEditorHookTest{
			srcPkg,
			nil,
		})
	}

	return astEditorTests
}

func (t astEditorHookTest) shouldPanic() bool {
	return strings.Contains(t.srcPkg.PkgPath, "/panic/")
}

func getFunctionName(i interface{}) string {
	fullname := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	s := strings.Split(fullname, ".")

	return s[len(s)-1]
}

func fileToString(pkg *packages.Package, file *ast.File) string {
	buf := &bytes.Buffer{}
	err := printer.Fprint(buf, pkg.Fset, file)
	if err != nil {
		panic(err)
	}

	return buf.String()
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

func TestAllAstEditorHooks(t *testing.T) {
	for _, option := range allAstEditorOptions {
		optionName := getFunctionName(option)
		astEditorTests := append(newAstEditorHookTests(optionName), newPanicAstEditorHookTests(optionName)...)
		editor := editor.newAstEditor(option)

		for _, test := range astEditorTests {
			err := testAstEditorHook(editor, test)
			if err != nil {
				t.Errorf("%v option failed a test:\n%v", optionName, err)
			}
		}
	}
}

func testAstEditorHook(editor FileEditor, test astEditorHookTest) error {
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

		editor.edit(srcFile)

		got := fileToString(test.srcPkg, srcFile)

		outFile := test.outPkg.Syntax[i]
		expected := fileToString(test.outPkg, outFile)

		if diff := cmp.Diff(expected, got); diff != "" {
			return fmt.Errorf("editor doesn't generate the expected result:\n%s", diff)
		}
	}

	return nil
}
