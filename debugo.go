package debugo

// Assert check that the given bool is true
// else it will panic.
var Assert func(bool, string) 

// AssertF check that the given function
// return true else it will panic.
var AssertF func(func() bool, string)
