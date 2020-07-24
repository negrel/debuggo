package debuggo

import "fmt"

const pkg = "debuggo"

// Assert the given condition and panic if false.
func Assert(ok bool, panicMsg string) {
	if !ok {
		panic(fmt.Sprintf("[%v] - %v", pkg, panicMsg))
	}
}
