package pkg_with_subpkg

import (
	"fmt"

	"github.com/negrel/debuggo/pkg/parse/_data/pkg/pkg_with_subpkg/log"
)

// Greet all the given person name.
func Greet(names ...string) {
	for _, name := range names {
		log.Print(name)
		fmt.Println("Hello,", name)
	}
}
