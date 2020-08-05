package invalid

// Add function return the sum of all
// the given arguments.
func Add(a int, b ...int) int {
	result := a

	for _, n := range b {
		result += n
	}

	return result
}
