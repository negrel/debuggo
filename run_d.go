// +build debugo

package debugo

// Run the given function.
func Run(fn func()) {
	fn()
}