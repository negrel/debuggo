package main

import (
	"fmt"
	"io/ioutil"

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
		cli.StringFlag{
			Name:      "src-dir",
			Usage:     "path to the directory that contains source packages.",
			TakesFile: true,
		},
		cli.StringSliceFlag{
			Name:      "src",
			Usage:     "path to the source package.",
			TakesFile: true,
		},
		cli.StringSliceFlag{
			Name:      "tags",
			Usage:     "build flags is a list of command-line flags to be passed through to the build system's query tool.",
			TakesFile: true,
		},
		cli.StringFlag{
			Name:      "output",
			Usage:     "output directory",
			Value:     "../",
			TakesFile: true,
		},
	},
	Action: func(ctx *cli.Context) error {
		srcDir := ctx.String("src-dir")
		infos, err := ioutil.ReadDir(srcDir)
		if err != nil {
			return err
		}

		srcPkgs := func() []string {
			n := make([]string, len(infos))

			for i, info := range infos {
				n[i] = srcDir + "/" + info.Name()
			}

			return n
		}()

		opt := debuggo.Options{
			PkgPattern: append(srcPkgs, ctx.StringSlice("src")...),
			PkgTags:    ctx.StringSlice("tags"),
			Output:     ctx.String("output"),
		}

		errs := debuggo.Generate(opt)
		e := ""

		for _, err := range errs {
			e += err.Error() + "\n"
		}

		return fmt.Errorf(e)
	},
}
