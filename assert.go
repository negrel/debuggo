// +build !debugo !debugo-assert

package debugo


// Assert the given bool and panic if false.
// assertion are executed if the debug mode
// is enabled.
func Assert(_ bool, _ string) {}

// AssertF the given function and panic if return
// false. Assertion are executed if the debug mode
// is enabled.
func AssertF(_ func() (bool, string)) {}