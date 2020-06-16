// +build !debugo

package debugo

// Print is a shortcut for debug.Run(func() { fmt.Print() })
func Print(a ...interface{}) {}

// Printf is a shortcut for debug.Run(func() { fmt.Printf() })
func Printf(format string, a ...interface{}) {}

// Println is a shortcut for debug.Run(func() { fmt.Println() })
func Println(a ...interface{}) {}
