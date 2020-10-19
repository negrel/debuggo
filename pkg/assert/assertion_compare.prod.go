// +build !assert

package assert

import (
	"reflect"
)

type CompareType int

const (
	compareLess	CompareType	= iota - 1
	compareEqual
	compareGreater
)

var (
	intType		= reflect.TypeOf(int(1))
	int8Type	= reflect.TypeOf(int8(1))
	int16Type	= reflect.TypeOf(int16(1))
	int32Type	= reflect.TypeOf(int32(1))
	int64Type	= reflect.TypeOf(int64(1))
	uintType	= reflect.TypeOf(uint(1))
	uint8Type	= reflect.TypeOf(uint8(1))
	uint16Type	= reflect.TypeOf(uint16(1))
	uint32Type	= reflect.TypeOf(uint32(1))
	uint64Type	= reflect.TypeOf(uint64(1))
	float32Type	= reflect.TypeOf(float32(1))
	float64Type	= reflect.TypeOf(float64(1))
	stringType	= reflect.TypeOf("")
)

func Greater( // Greater asserts that the first element is greater than the second
//
//    assert.Greater(t, 2, 1)
//    assert.Greater(t, float64(2), float64(1))
//    assert.Greater(t, "b", "a")
_ interface{}, _ interface{}, _ ...interface{}) {
}

func GreaterOrEqual( // GreaterOrEqual asserts that the first element is greater than or equal to the second
//
//    assert.GreaterOrEqual(t, 2, 1)
//    assert.GreaterOrEqual(t, 2, 2)
//    assert.GreaterOrEqual(t, "b", "a")
//    assert.GreaterOrEqual(t, "b", "b")
_ interface{}, _ interface{}, _ ...interface{}) {
}

func Less( // Less asserts that the first element is less than the second
//
//    assert.Less(t, 1, 2)
//    assert.Less(t, float64(1), float64(2))
//    assert.Less(t, "a", "b")
_ interface{}, _ interface{}, _ ...interface{}) {
}

func LessOrEqual( // LessOrEqual asserts that the first element is less than or equal to the second
//
//    assert.LessOrEqual(t, 1, 2)
//    assert.LessOrEqual(t, 2, 2)
//    assert.LessOrEqual(t, "a", "b")
//    assert.LessOrEqual(t, "b", "b")
_ interface{}, _ interface{}, _ ...interface{}) {
}
