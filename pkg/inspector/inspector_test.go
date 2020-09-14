package inspector

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"

	"github.com/stretchr/testify/assert"
)

type counter struct {
	value int
}

func (c *counter) count() int {
	return c.value
}

func (c *counter) topLevelDecl(n ast.Node) bool {
	if _, isDecl := n.(ast.Decl); isDecl {
		c.value++
		return false
	}

	return true
}

func (c *counter) importsCounter(n ast.Node) bool {
	if _, isImport := n.(*ast.ImportSpec); isImport {
		c.value++
		return false
	}

	return true
}

func TestEditor_Inspect(t *testing.T) {
	src := `
package main

import "fmt"

func main() {
	// Should not be inspected by funcParamsCounter
	greetWorld := func() {
		fmt.Println("Hello world")
	}

	greet("world")
	greetWorld()
}

func greet(name string) {
	fmt.Println("Hello", name)
}
`
	fs := token.NewFileSet()
	file, err := parser.ParseFile(fs, "", src, parser.AllErrors)
	assert.Nil(t, err)

	topLevelDeclCounter := new(counter)
	importsCounter := new(counter)
	editor := New(importsCounter.importsCounter, topLevelDeclCounter.topLevelDecl)

	editor.Inspect(file)

	assert.Equal(t, 3, topLevelDeclCounter.count())
	assert.Equal(t, 1, importsCounter.count())
}
