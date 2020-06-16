// +build !debugo

package debugo

// Assert the given bool and panic if false.
//
// The function will be removed by the compiler
// for production the debugo build tag is used.
func Assert(_ bool, _ string) {}
