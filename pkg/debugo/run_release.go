// +build !debugo

package debugo

// Run the given function.
//
// The function will be removed by the compiler
// for production the debugo build tag is used.
func Run(_ func()) {}
