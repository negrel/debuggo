package debugo

// Assert check that the given bool is true
// else it will panic.
var Assert = func(_ bool, _ string) {}

// AssertF check that the given function
// return true else it will panic.
var AssertF = func(_ func() (bool, string)) {}
