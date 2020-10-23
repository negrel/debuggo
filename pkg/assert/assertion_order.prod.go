// +build !assert

package assert

import ()

func IsIncreasing( // IsIncreasing asserts that the collection is increasing
	//
	//    assert.IsIncreasing(t, []int{1, 2, 3})
	//    assert.IsIncreasing(t, []float{1, 2})
	//    assert.IsIncreasing(t, []string{"a", "b"})
	_ interface{}, _ ...interface{}) {
}

func IsNonIncreasing( // IsNonIncreasing asserts that the collection is not increasing
	//
	//    assert.IsNonIncreasing(t, []int{2, 1, 1})
	//    assert.IsNonIncreasing(t, []float{2, 1})
	//    assert.IsNonIncreasing(t, []string{"b", "a"})
	_ interface{}, _ ...interface{}) {
}

func IsDecreasing( // IsDecreasing asserts that the collection is decreasing
	//
	//    assert.IsDecreasing(t, []int{2, 1, 0})
	//    assert.IsDecreasing(t, []float{2, 1})
	//    assert.IsDecreasing(t, []string{"b", "a"})
	_ interface{}, _ ...interface{}) {
}

func IsNonDecreasing( // IsNonDecreasing asserts that the collection is not decreasing
	//
	//    assert.IsNonDecreasing(t, []int{1, 1, 2})
	//    assert.IsNonDecreasing(t, []float{1, 2})
	//    assert.IsNonDecreasing(t, []string{"a", "b"})
	_ interface{}, _ ...interface{}) {
}
