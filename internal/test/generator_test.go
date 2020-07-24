package test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/negrel/debuggo/internal/generator"
)

const (
	dataDir    string = "./_data"
	debuggoDir string = "./result"
	resultDir  string = "./_data_result"
)

func TestDataSet(t *testing.T) {
	// Test Data packages
	pkgsInfos, err := ioutil.ReadDir(dataDir)
	if err != nil {
		t.Fatal(err)
	}

	pkgs := func() []string {
		pkgNames := make([]string, len(pkgsInfos))

		for i, info := range pkgsInfos {
			pkgNames[i] = dataDir + "/" + info.Name()
		}

		return pkgNames
	}()

	opt := generator.Options{
		PkgPattern: pkgs,
		Output:     debuggoDir,
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
	for pkgName := range result {
		debuggoPkg := debuggo[pkgName]
		resultPkg := result[pkgName]

		for filename := range resultPkg.fset {
			debuggoFile, ok := debuggoPkg.fset[filename]
			if !ok {
				t.Fatalf("%v is missing in %v/%v", filename, debuggoDir, pkgName)
			}
			resultFile := resultPkg.fset[filename]

			if !bytes.Equal(debuggoFile, resultFile) {
				t.Fatalf("%v/%v/%v is not equal to %v/%v/%v", debuggoDir, pkgName, filename, resultDir, pkgName, filename)
			}
		}
	}
}
