package test

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/negrel/debuggo/internal/generator"
)

const (
	debuggoDir string = "./result"
	resultDir  string = "./_data_result"
)

func TestDataSet(t *testing.T) {
	// Test Data packages

	opt := generator.Options{
		PkgPattern: []string{
			"./_data/test1",
		},
		Output: debuggoDir,
	}

	generator.Generate(opt)

	// Comparing expected result and debuggo result.
	debuggo := make(map[string]*pkg, 8)
	result := make(map[string]*pkg, 8)

	for i := 0; i < 2; i++ {
		dirs := [2]string{
			debuggoDir,
			resultDir,
		}

		// Setup pkg maps
		func() {
			dir, err := os.Open(dirs[i])
			if err != nil {
				t.Fatal(err)
			}

			pkgNames, err := dir.Readdirnames(9999)
			if err != nil {
				t.Fatal(err)
			}

			// Comparing result and expected result.
			for _, pkgName := range pkgNames {
				path := fmt.Sprintf("%v/%v", dirs[i], pkgName)
				pkg, err := newPkg(path)
				if err != nil {
					t.Fatal(err)
				}

				if dirs[i] == debuggoDir {
					debuggo[pkgName] = pkg
				} else {
					result[pkgName] = pkg
				}
			}
		}()
	}

	// Comparing result and expected result
	for name := range debuggo {
		debuggoPkg := debuggo[name]
		resultPkg := result[name]

		for filename := range debuggoPkg.fset {
			debuggoFile := debuggoPkg.fset[filename]
			resultFile := resultPkg.fset[filename]

			if !bytes.Equal(debuggoFile, resultFile) {
				t.Fatalf("%v/%v/%v is not equal to %v/%v/%v", debuggoDir, name, filename, resultDir, name, filename)
			}
		}
	}
}
