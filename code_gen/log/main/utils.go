package main

import (
	"path/filepath"
)

func addSuffix(filename, suffix string) string {
	ext := filepath.Ext(filename)
	extLen := len(ext)
	nameLen := len(filename)

	return filename[:nameLen-extLen] + suffix + ext
}
