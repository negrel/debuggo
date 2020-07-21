package main

import (
	debuggo "github.com/negrel/debuggo/internal/generator"
	"github.com/urfave/cli"
)

// generate command
var generate = cli.Command{
	Name:        "generate",
	ShortName:   "gen",
	Usage:       "Generate debugging package.",
	UsageText:   "debuggo generate ",
	Description: "Generate an optimized version (for production) of the given debugging package.",
	Flags: []cli.Flag{
		cli.StringSliceFlag{
			Name:      "src",
			Usage:     "path to the source package.",
			Required:  true,
			TakesFile: true,
		},
		cli.StringSliceFlag{
			Name:      "tags",
			Usage:     "build flags is a list of command-line flags to be passed through to the build system's query tool.",
			Required:  false,
			TakesFile: true,
		},
		cli.StringFlag{
			Name:      "output",
			Usage:     "output directory",
			Value:     "../",
			TakesFile: true,
		},
	},
	Action: func(ctx *cli.Context) {
		opt := debuggo.Options{
			PkgPattern: ctx.StringSlice("src"),
			PkgTags:    ctx.StringSlice("tags"),
			Output:     ctx.String("output"),
		}

		debuggo.Generate(opt)
	},
}
