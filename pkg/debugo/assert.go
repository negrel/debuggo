// +build debugo

package debugo

// Assert the given bool and panic if false.
func Assert(ok bool, err string) {
	if !ok {
		panic(err)
	}
}
