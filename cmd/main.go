package main

import "github.com/negrel/debuggo/internal/debuggo"

func main() {
	opt := debuggo.Options{
		PkgPattern: []string{"github.com/negrel/debuggo/examples/helloworld/internal/debuggo/person"},
		PkgTags:    []string{},
		Output:     "../",
		OutputDir:  "debug",
	}

	debuggo.Generate(opt)
}
