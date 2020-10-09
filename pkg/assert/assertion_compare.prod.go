// +build !assert

package assert

import ()

func // Greater asserts that the first element is greater than the second
//
//    assert.Greater(t, 2, 1)
//    assert.Greater(t, float64(2), float64(1))
//    assert.Greater(t, "b", "a")
Greater(e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) {
}

func // GreaterOrEqual asserts that the first element is greater than or equal to the second
//
//    assert.GreaterOrEqual(t, 2, 1)
//    assert.GreaterOrEqual(t, 2, 2)
//    assert.GreaterOrEqual(t, "b", "a")
//    assert.GreaterOrEqual(t, "b", "b")
GreaterOrEqual(e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) {
}

func // Less asserts that the first element is less than the second
//
//    assert.Less(t, 1, 2)
//    assert.Less(t, float64(1), float64(2))
//    assert.Less(t, "a", "b")
Less(e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) {
}

func // LessOrEqual asserts that the first element is less than or equal to the second
//
//    assert.LessOrEqual(t, 1, 2)
//    assert.LessOrEqual(t, 2, 2)
//    assert.LessOrEqual(t, "a", "b")
//    assert.LessOrEqual(t, "b", "b")
LessOrEqual(e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) {
}
