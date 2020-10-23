package main

import (
	"log"
	"path/filepath"

	"github.com/negrel/asttk/pkg/parse"
)

func main() {
	pkg, err := parse.Package(filepath.Join("code_gen", "log", "log"), false)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range pkg.Files {
		if file.Name() != "exported.go" {
			err = file.WriteFile(
				filepath.Join("pkg", "log", file.Name()),
			)
			if err != nil {
				log.Fatal()
			}
			continue
		}

		logLevel := len(logLevelsName)
		for logLevel != 0 {
			logLevel--
			editFile(file, logLevel)
		}

		editProdFile(file)
	}
}
