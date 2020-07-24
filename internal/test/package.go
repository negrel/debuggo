package test

import (
	"fmt"
	"io/ioutil"
	"os"
)

type pkg struct {
	name    string
	fset    map[string][]byte
	subPkgs []*pkg
}

func newPkg(path string) (*pkg, error) {
	result := &pkg{
		fset: make(map[string][]byte, 8),
	}

	// Checking pkg directory
	srcFiles, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, srcFile := range srcFiles {
		filePath := fmt.Sprintf("%v/%v", path, srcFile.Name())
		if stat, err := os.Stat(filePath); err != nil {
			return nil, err
		} else if stat.IsDir() {
			subPkg, err := newPkg(filePath)
			if err != nil {
				return nil, err
			}

			result.subPkgs = append(result.subPkgs, subPkg)
			continue
		}

		content, err := ioutil.ReadFile(filePath)
		if err != nil {
			return nil, err
		}

		result.fset[srcFile.Name()] = content
	}

	return result, nil
}
