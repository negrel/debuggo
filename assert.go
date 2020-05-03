package debug

// DEBUG mode enabled
const DEBUG = true

// Assert the given bool and panic if false.
// assertion are executed if the debug mode
// is enabled.
func Assert(ok bool, msg string) {
	if DEBUG {
		if !ok {
			panic(msg)
		}
	}
}

// AssertF the given function and panic if return
// false. Assertion are executed if the debug mode
// is enabled.
func AssertF(ok func() bool, msg string) {
	if DEBUG {
		if !ok() {
			panic(msg)
		}
	}
}
