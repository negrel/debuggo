package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "🔺 Debuggo"
	app.Usage = "Optimized debugging package generator."
	app.Description = `Debuggo generate an optimized debugging package for minimal
	 performance / size overhead in production.`
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
	app.ExitErrHandler = func(context *cli.Context, err error) {
		_, _ = fmt.Fprint(os.Stderr, err)
	}

	_ = app.Run(os.Args)
}
