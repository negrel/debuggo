package test

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/negrel/debuggo/internal/generator"

	"github.com/google/go-cmp/cmp"
)

const (
	dataDir   string = "./_data/src"
	outputDir string = "./output"
	resultDir string = "./_data_result"
)

type testPkg struct {
	path    string
	goFiles map[string][]byte
}

func compareExpectedAndGeneratedFiles() error {
	generatedPkgs, err := generator.GetAllPackages(outputDir)
	if err != nil {
		return err
	}

	expectedPkgs, err := generator.GetAllPackages(resultDir)
	if err != nil {
		return err
	}

	for i := range generatedPkgs {
		generatedPkg := generatedPkgs[i]
		expectedPkg := expectedPkgs[i]

		for j := range generatedPkg.GoFiles {
			generatedFilePath := generatedPkg.GoFiles[j]
			expectedFilePath := expectedPkg.GoFiles[j]

			generatedFile, err := ioutil.ReadFile(generatedFilePath)
			if err != nil {
				return err
			}

			expectedFile, err := ioutil.ReadFile(expectedFilePath)
			if err != nil {
				return nil
			}

			if diff := cmp.Diff(generatedFile, expectedFile); diff != "" {
				header := fmt.Sprintf("\nGENERATED FILE: %v\nEXPECTED FILE: %v", generatedFilePath, expectedFilePath)
				return fmt.Errorf("%v\nErrors generated file is different from the expected result:\n%s", header, diff)
			}
		}
	}

	return nil
}

func TestDataSet(t *testing.T) {
	err := os.RemoveAll(outputDir)
	if err != nil {
		t.Fatal(err)
	}

	gen, err := generator.New(
		generator.SrcDir(dataDir),
		generator.OutputDir(outputDir),
	)
	if err != nil {
		t.Fatal(err)
	}

	gen.Start()
	if err := gen.Error(); err != nil {
		t.Fatal(err)
	}

	err = compareExpectedAndGeneratedFiles()
	if err != nil {
		t.Fatal(err)
	}
}
