// +build debug

package debugo

func init() {
	Assert = assert
	AssertF = assertF
}

// Assert the given bool and panic if false.
// assertion are executed if the debug mode
// is enabled.
func assert(ok bool, msg string) {
	if !ok {
		panic(msg)
	}
}

// AssertF the given function and panic if return
// false. Assertion are executed if the debug mode
// is enabled.
func assertF(ok func() bool, msg string) {
	if !ok() {
		panic(msg)
	}
}
