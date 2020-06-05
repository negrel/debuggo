// +build !debugo

package debugo

// Run the given function.
//
// You must compile using the 'debugo' or
//'debugo-assert' build flag, otherwise
// Run will be removed by the compiler.
func Run(_ func()) {}