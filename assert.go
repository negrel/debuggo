// +build debug

package debugo

func init() {
	Assert = assert
	AssertF = assertF
}

// Assert the given bool and panic if false.
// assertion are executed if the debug mode
// is enabled.
func assert(ok bool, err string) {
	if !ok {
		panic(err)
	}
}

// AssertF the given function and panic if return
// false. Assertion are executed if the debug mode
// is enabled.
func assertF(fn func() (bool, string)) {
	if ok, err := fn(); !ok {
		panic(err)
	}
}
