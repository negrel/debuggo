package pkg_with_subpkg

import (
	"fmt"
)

// Greet all the given person name.
func Greet(names ...string) {
	for _, name := range names {
		fmt.Println("Hello,", name)
	}
}
