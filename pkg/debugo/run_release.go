// +build !debugo

package debugo

// Run the given function.
// 
// The function will be removed by the compiler
// for production if you use the release build tag.
func Run(_ func()) {}