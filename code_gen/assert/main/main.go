package main

import (
	"log"
	"path/filepath"

	"github.com/negrel/asttk/pkg/parse"
)

func main() {
	pkg, err := parse.Package(
		filepath.Join("code_gen", "assert", "testify", "assert"),
		false,
	)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range pkg.Files {
		if isBlackListed(file.Name()) {
			continue
		}

		editFile(file)
		editProdFile(file)
	}
}
