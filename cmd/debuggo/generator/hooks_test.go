package generator

import (
	"bytes"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/negrel/debuggo/pkg/inspector"
)

type hookTest struct {
	src string
	out string
}

var unusedImportsRemoverTests = []hookTest{
	{
		src: `package main

import (
	"fmt"
	// image is imported but not used
	"image"
)

type date struct {
	dd, mm, yy string
}

func main() {
	image := date{
		dd:	"01",
		mm:	"01",
		yy:	"1970",
	}

	fmt.Println("Hello world", image.dd)
}
`,
		out: `package main

import (
	"fmt"
)

type date struct {
	dd, mm, yy string
}

func main() {
	image := date{
		dd:	"01",
		mm:	"01",
		yy:	"1970",
	}

	fmt.Println("Hello world", image.dd)
}
`,
	},
}

func TestUnusedImportsRemover(t *testing.T) {
	uir := new(unusedImportsRemover)
	chief := inspector.New(uir.inspect)

	for _, test := range unusedImportsRemoverTests {
		fset := token.NewFileSet()
		file, err := parser.ParseFile(fset, "unusedImportsRemover", test.src, parser.AllErrors)
		if err != nil {
			t.Fatal(err)
		}

		chief.Inspect(file)
		uir.removeImports(file)

		result := bytes.NewBuffer([]byte{})
		err = printer.Fprint(result, fset, file)
		if err != nil {
			log.Fatal(err)
		}

		if diff := cmp.Diff(test.out, result.String()); diff != "" {
			t.Fatal(diff)
		}
	}
}
