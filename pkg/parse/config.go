package parse

import (
	"golang.org/x/tools/go/packages"
)

// Config is the config used to load file/packages.
// Config.Dir is overwritten before loading files.
var Config = packages.Config{
	Mode: packages.NeedName | packages.NeedSyntax |
		packages.NeedImports | packages.NeedCompiledGoFiles |
		packages.NeedFiles | packages.NeedTypes |
		packages.NeedDeps,

	Tests:      false,
	BuildFlags: []string{},
}
