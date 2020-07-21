package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "🔺 Debuggo"
	app.Usage = "Optimized debugging package generator."
	app.Description = `Debuggo generate an optimized debugging package for minimal
	 performance / size overhead in production.
`
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
	app.EnableBashCompletion = true

	// opt := debuggo.Options{
	// 	PkgPattern: []string{"../../debuggo.example/helloworld/internal/debuggo/person"},
	// 	PkgTags:    []string{},
	// 	Output:     "../",
	// 	OutputDir:  "debug",
	// }

	// debuggo.Generate(opt)

	app.Run(os.Args)
}
