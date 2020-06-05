// +build debugo debugo-assert

package debugo

// Assert the given bool and panic if false.
// assertion are executed if the debug mode
// is enabled.
func Assert(ok bool, err string) {
	if !ok {
		panic(err)
	}
}

