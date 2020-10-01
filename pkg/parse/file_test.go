package parse

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFile_NewFile_EmptyPath(t *testing.T) {
	_, err := File("")
	assert.NotNil(t, err)
}

func TestFile_NewFile_InvalidPathSyntax(t *testing.T) {
	goFile, err := File("??|\\/file.go")
	assert.NotNil(t, err, err)
	assert.Nil(t, goFile)
}

func TestFile_NewFile_NotGoFile(t *testing.T) {
	filePath := filepath.Join("..", "..", "..", "README.md")
	_, err := File(filePath)
	assert.NotNil(t, err)
}

func TestFile_NewFile(t *testing.T) {
	filePath := filepath.Join(".", "_data", "file", "greet.go")
	filePath, _ = filepath.Abs(filePath)

	goFile, err := File(filePath)
	assert.Nil(t, err, err)

	assert.Equal(t, filePath, goFile.Path())
	assert.Equal(t, filepath.Dir(filePath), goFile.Dir())
	assert.Equal(t, filepath.Base(filePath), goFile.Name())
	assert.NotNil(t, goFile.AST())
}
