package test

import (
	"testing"

	"github.com/negrel/debuggo/internal/generator"
)

const (
	dataDir    string = "./_data/src"
	debuggoDir string = "./result"
	resultDir  string = "./_data_result"
)

func TestDataSet(t *testing.T) {
	gen, err := generator.New(
		generator.SrcDir(dataDir),
		generator.OutputDir(debuggoDir),
	)
	if err != nil {
		t.Fatal(err)
	}

	gen.Start()
	if err := gen.Error(); err != nil {
		t.Fatal(err)
	}
}
