package main

import (
	"github.com/negrel/debuggo/internal/generator"
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
		cli.StringFlag{
			Name:      "src-dir",
			Usage:     "path to the directory that contains source packages.",
			TakesFile: true,
			Required:  true,
		},
		cli.StringFlag{
			Name:      "src-file",
			Usage:     "path to a Go file.",
			TakesFile: true,
			Required:  true,
		},
		cli.StringFlag{
			Name:      "cmn-dir",
			Usage:     "common directory",
			TakesFile: true,
		},
		cli.StringSliceFlag{
			Name:      "tags",
			Usage:     "build flags is a list of command-line flags to be passed through to the build system's query tool.",
			TakesFile: true,
		},
		cli.StringFlag{
			Name:      "out-dir",
			Usage:     "output directory",
			Value:     "../debug",
			TakesFile: true,
		},
	},
	Action: func(ctx *cli.Context) error {
		srcDir := ctx.String("src-dir")
		outDir := ctx.String("out-dir")
		commonDir := ctx.String("cmn-dir")

		options := make([]generator.Option, 0, 3)
		options = append(options, generator.SrcDir(srcDir), generator.OutputDir(outDir))
		if commonDir != "" {
			options = append(options, generator.CommonDir(commonDir))
		}

		gen, err := generator.New(options...)
		if err != nil {
			return err
		}

		gen.Start()

		return gen.Error()
	},
}
