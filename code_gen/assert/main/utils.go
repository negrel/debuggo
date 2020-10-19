package main

import (
	"go/ast"
	"path/filepath"
)

func min(a int, b int) int {
	if a < b {
		return a
	}

	return b
}

func addSuffix(filename, suffix string) string {
	ext := filepath.Ext(filename)
	extLen := len(ext)
	nameLen := len(filename)

	return filename[:nameLen-extLen] + suffix + ext
}

func extractArguments(fl *ast.FieldList) []ast.Expr {
	result := make([]ast.Expr, 0, len(fl.List))

	for _, field := range fl.List {
		for _, name := range field.Names {
			result = append(result, &ast.Ident{
				Name: name.Name,
			})
		}
	}

	return result
}
