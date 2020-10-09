// +build assert

package assert

import ()

func // IsIncreasing asserts that the collection is increasing
//
//    assert.IsIncreasing(t, []int{1, 2, 3})
//    assert.IsIncreasing(t, []float{1, 2})
//    assert.IsIncreasing(t, []string{"a", "b"})
IsIncreasing(object interface{}, msgAndArgs ...interface{}) {
	debuggoGen_IsIncreasing(object, msgAndArgs)
}

func // IsNonIncreasing asserts that the collection is not increasing
//
//    assert.IsNonIncreasing(t, []int{2, 1, 1})
//    assert.IsNonIncreasing(t, []float{2, 1})
//    assert.IsNonIncreasing(t, []string{"b", "a"})
IsNonIncreasing(object interface{}, msgAndArgs ...interface{}) {
	debuggoGen_IsNonIncreasing(object, msgAndArgs)
}

func // IsDecreasing asserts that the collection is decreasing
//
//    assert.IsDecreasing(t, []int{2, 1, 0})
//    assert.IsDecreasing(t, []float{2, 1})
//    assert.IsDecreasing(t, []string{"b", "a"})
IsDecreasing(object interface{}, msgAndArgs ...interface{}) {
	debuggoGen_IsDecreasing(object, msgAndArgs)
}

func // IsNonDecreasing asserts that the collection is not decreasing
//
//    assert.IsNonDecreasing(t, []int{1, 1, 2})
//    assert.IsNonDecreasing(t, []float{1, 2})
//    assert.IsNonDecreasing(t, []string{"a", "b"})
IsNonDecreasing(object interface{}, msgAndArgs ...interface{}) {
	debuggoGen_IsNonDecreasing(object, msgAndArgs)
}
