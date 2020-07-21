package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "🔺 Debuggo"
	app.HelpName = "debuggo"
	app.Version = "1.0.0"
	app.HideVersion = true
	app.Authors = []cli.Author{
		{
			Name:  "Alexandre Negrel",
			Email: "negrel.dev@protonmail.com",
		},
	}
	app.Commands = []cli.Command{
		generate,
	}

	// opt := debuggo.Options{
	// 	PkgPattern: []string{"../../debuggo.example/helloworld/internal/debuggo/person"},
	// 	PkgTags:    []string{},
	// 	Output:     "../",
	// 	OutputDir:  "debug",
	// }

	// debuggo.Generate(opt)

	app.Run(os.Args)
}
