// +build !debugo !debugo-assert

package debugo


// Assert the given bool and panic if false.
//
// You must compile using the 'debugo' or
//'debugo-assert' build flag, otherwise
// Assert will be removed by the compiler.
func Assert(_ bool, _ string) {}
