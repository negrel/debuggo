package parse

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPkg_NewPkg_EmptyPath(t *testing.T) {
	_, err := Package("", false)
	assert.NotNil(t, err)
}

func TestPkg_NewPkg_InvalidPathSyntax(t *testing.T) {
	_, err := Package("??|\\/file.go", false)
	assert.NotNil(t, err, err)
}

func TestPkg_NewPkg_NotDirectory(t *testing.T) {
	filePath := filepath.Join("..", "..", "..", "README.md")
	_, err := Package(filePath, false)
	assert.NotNil(t, err)
}

func TestPkg_NewPkg_NoSubPkg(t *testing.T) {
	dir := filepath.Join(".", "_data", "pkg", "pkg_with_subpkg")

	pkg, err := Package(dir, false)
	assert.Nil(t, err, err)

	assert.NotNil(t, pkg)
	assert.Len(t, pkg.SubPkgs(), 0)
	assert.Len(t, pkg.Files(), 1)
}

func TestPkg_NewPkg_WithSubPkg(t *testing.T) {
	dir := filepath.Join(".", "_data", "pkg", "pkg_with_subpkg")

	pkg, err := Package(dir, true)
	assert.Nil(t, err, err)

	assert.NotNil(t, pkg)
	assert.Len(t, pkg.SubPkgs(), 1)
	assert.Len(t, pkg.Files(), 1)

	subPkg := pkg.SubPkgs()[0]
	subPkgPath := filepath.Join(dir, subPkg.Name())
	subPkgPath, err = filepath.Abs(subPkgPath)
	assert.Nil(t, err)
	assert.Equal(t, subPkg.Path(), subPkgPath)
}
