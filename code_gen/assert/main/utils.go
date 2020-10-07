package main

import (
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
