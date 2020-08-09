package debug

// AssertTrue function assert that the given boolean is true.
func AssertTrue(isTrue bool, panicMsg string) {
	if !isTrue {
		panic(panicMsg)
	}
}
