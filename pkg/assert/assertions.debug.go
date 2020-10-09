// +build assert

package assert

import (
	"time"
)

func // ObjectsAreEqual determines if two objects are considered equal.
//
// This function does no assertion of any kind.
ObjectsAreEqual(expected, actual interface{}) {
	debuggoGen_ObjectsAreEqual(expected, actual)
}

func // ObjectsAreEqualValues gets whether two objects are equal, or if their
// values are equal.
ObjectsAreEqualValues(expected, actual interface{}) {
	debuggoGen_ObjectsAreEqualValues(expected, actual)
}

func // CallerInfo returns an array of strings containing the file and line number
// of each stack frame leading from the current test to the assert call that
// failed.
CallerInfo() {
	debuggoGen_CallerInfo()
}

func // FailNow fails test
FailNow(failureMessage string, msgAndArgs ...interface{}) {
	debuggoGen_FailNow(failureMessage, msgAndArgs)
}

func // Fail reports a failure through
Fail(failureMessage string, msgAndArgs ...interface{}) {
	debuggoGen_Fail(failureMessage, msgAndArgs)
}

func // Implements asserts that an object is implemented by the specified interface.
//
//    assert.Implements(t, (*MyInterface)(nil), new(MyObject))
Implements(interfaceObject interface{}, object interface{}, msgAndArgs ...interface{}) {
	debuggoGen_Implements(interfaceObject, object, msgAndArgs)
}

func // IsType asserts that the specified objects are of the same type.
IsType(expectedType interface{}, object interface{}, msgAndArgs ...interface{}) {
	debuggoGen_IsType(expectedType, object, msgAndArgs)
}

func // Equal asserts that two objects are equal.
//
//    assert.Equal(t, 123, 123)
//
// Pointer variable equality is determined based on the equality of the
// referenced values (as opposed to the memory addresses). Function equality
// cannot be determined and will always fail.
Equal(expected, actual interface{}, msgAndArgs ...interface{}) {
	debuggoGen_Equal(expected, actual, msgAndArgs)
}

func // Same asserts that two pointers reference the same object.
//
//    assert.Same(t, ptr1, ptr2)
//
// Both arguments must be pointer variables. Pointer variable sameness is
// determined based on the equality of both type and value.
Same(expected, actual interface{}, msgAndArgs ...interface{}) {
	debuggoGen_Same(expected, actual, msgAndArgs)
}

func // NotSame asserts that two pointers do not reference the same object.
//
//    assert.NotSame(t, ptr1, ptr2)
//
// Both arguments must be pointer variables. Pointer variable sameness is
// determined based on the equality of both type and value.
NotSame(expected, actual interface{}, msgAndArgs ...interface{}) {
	debuggoGen_NotSame(expected, actual, msgAndArgs)
}

func // EqualValues asserts that two objects are equal or convertable to the same types
// and equal.
//
//    assert.EqualValues(t, uint32(123), int32(123))
EqualValues(expected, actual interface{}, msgAndArgs ...interface{}) {
	debuggoGen_EqualValues(expected, actual, msgAndArgs)
}

func // Exactly asserts that two objects are equal in value and type.
//
//    assert.Exactly(t, int32(123), int64(123))
Exactly(expected, actual interface{}, msgAndArgs ...interface{}) {
	debuggoGen_Exactly(expected, actual, msgAndArgs)
}

func // NotNil asserts that the specified object is not nil.
//
//    assert.NotNil(t, err)
NotNil(object interface{}, msgAndArgs ...interface{}) {
	debuggoGen_NotNil(object, msgAndArgs)
}

func // Nil asserts that the specified object is nil.
//
//    assert.Nil(t, err)
Nil(object interface{}, msgAndArgs ...interface{}) {
	debuggoGen_Nil(object, msgAndArgs)
}

func // Empty asserts that the specified object is empty.  I.e. nil, "", false, 0 or either
// a slice or a channel with len == 0.
//
//  assert.Empty(t, obj)
Empty(object interface{}, msgAndArgs ...interface{}) {
	debuggoGen_Empty(object, msgAndArgs)
}

func // NotEmpty asserts that the specified object is NOT empty.  I.e. not nil, "", false, 0 or either
// a slice or a channel with len == 0.
//
//  if assert.NotEmpty(t, obj) {
//    assert.Equal(t, "two", obj[1])
//  }
NotEmpty(object interface{}, msgAndArgs ...interface{}) {
	debuggoGen_NotEmpty(object, msgAndArgs)
}

func // Len asserts that the specified object has specific length.
// Len also fails if the object has a type that len() not accept.
//
//    assert.Len(t, mySlice, 3)
Len(object interface{}, length int, msgAndArgs ...interface{}) {
	debuggoGen_Len(object, length, msgAndArgs)
}

func // True asserts that the specified value is true.
//
//    assert.True(t, myBool)
True(value bool, msgAndArgs ...interface{}) {
	debuggoGen_True(value, msgAndArgs)
}

func // False asserts that the specified value is false.
//
//    assert.False(t, myBool)
False(value bool, msgAndArgs ...interface{}) {
	debuggoGen_False(value, msgAndArgs)
}

func // NotEqual asserts that the specified values are NOT equal.
//
//    assert.NotEqual(t, obj1, obj2)
//
// Pointer variable equality is determined based on the equality of the
// referenced values (as opposed to the memory addresses).
NotEqual(expected, actual interface{}, msgAndArgs ...interface{}) {
	debuggoGen_NotEqual(expected, actual, msgAndArgs)
}

func // NotEqualValues asserts that two objects are not equal even when converted to the same type
//
//    assert.NotEqualValues(t, obj1, obj2)
NotEqualValues(expected, actual interface{}, msgAndArgs ...interface{}) {
	debuggoGen_NotEqualValues(expected, actual, msgAndArgs)
}

func // Contains asserts that the specified string, list(array, slice...) or map contains the
// specified substring or element.
//
//    assert.Contains(t, "Hello World", "World")
//    assert.Contains(t, ["Hello", "World"], "World")
//    assert.Contains(t, {"Hello": "World"}, "Hello")
Contains(s, contains interface{}, msgAndArgs ...interface{}) {
	debuggoGen_Contains(s, contains, msgAndArgs)
}

func // NotContains asserts that the specified string, list(array, slice...) or map does NOT contain the
// specified substring or element.
//
//    assert.NotContains(t, "Hello World", "Earth")
//    assert.NotContains(t, ["Hello", "World"], "Earth")
//    assert.NotContains(t, {"Hello": "World"}, "Earth")
NotContains(s, contains interface{}, msgAndArgs ...interface{}) {
	debuggoGen_NotContains(s, contains, msgAndArgs)
}

func // Subset asserts that the specified list(array, slice...) contains all
// elements given in the specified subset(array, slice...).
//
//    assert.Subset(t, [1, 2, 3], [1, 2], "But [1, 2, 3] does contain [1, 2]")
Subset(list, subset interface{}, msgAndArgs ...interface{}) {
	debuggoGen_Subset(list, subset, msgAndArgs)
}

func // NotSubset asserts that the specified list(array, slice...) contains not all
// elements given in the specified subset(array, slice...).
//
//    assert.NotSubset(t, [1, 3, 4], [1, 2], "But [1, 3, 4] does not contain [1, 2]")
NotSubset(list, subset interface{}, msgAndArgs ...interface{}) {
	debuggoGen_NotSubset(list, subset, msgAndArgs)
}

func // ElementsMatch asserts that the specified listA(array, slice...) is equal to specified
// listB(array, slice...) ignoring the order of the elements. If there are duplicate elements,
// the number of appearances of each of them in both lists should match.
//
// assert.ElementsMatch(t, [1, 3, 2, 3], [1, 3, 3, 2])
ElementsMatch(listA, listB interface{}, msgAndArgs ...interface{}) {
	debuggoGen_ElementsMatch(listA, listB, msgAndArgs)
}

func // Condition uses a Comparison to assert a complex condition.
Condition(comp Comparison, msgAndArgs ...interface{}) {
	debuggoGen_Condition(comp, msgAndArgs)
}

func // Panics asserts that the code inside the specified PanicTestFunc panics.
//
//   assert.Panics(t, func(){ GoCrazy() })
Panics(f PanicTestFunc, msgAndArgs ...interface{}) {
	debuggoGen_Panics(f, msgAndArgs)
}

func // PanicsWithValue asserts that the code inside the specified PanicTestFunc panics, and that
// the recovered panic value equals the expected panic value.
//
//   assert.PanicsWithValue(t, "crazy error", func(){ GoCrazy() })
PanicsWithValue(expected interface{}, f PanicTestFunc, msgAndArgs ...interface{}) {
	debuggoGen_PanicsWithValue(expected, f, msgAndArgs)
}

func // PanicsWithError asserts that the code inside the specified PanicTestFunc
// panics, and that the recovered panic value is an error that satisfies the
// EqualError comparison.
//
//   assert.PanicsWithError(t, "crazy error", func(){ GoCrazy() })
PanicsWithError(errString string, f PanicTestFunc, msgAndArgs ...interface{}) {
	debuggoGen_PanicsWithError(errString, f, msgAndArgs)
}

func // NotPanics asserts that the code inside the specified PanicTestFunc does NOT panic.
//
//   assert.NotPanics(t, func(){ RemainCalm() })
NotPanics(f PanicTestFunc, msgAndArgs ...interface{}) {
	debuggoGen_NotPanics(f, msgAndArgs)
}

func // WithinDuration asserts that the two times are within duration delta of each other.
//
//   assert.WithinDuration(t, time.Now(), time.Now(), 10*time.Second)
WithinDuration(expected, actual time.Time, delta time.Duration, msgAndArgs ...interface{}) {
	debuggoGen_WithinDuration(expected, actual, delta, msgAndArgs)
}

func // InDelta asserts that the two numerals are within delta of each other.
//
// 	 assert.InDelta(t, math.Pi, 22/7.0, 0.01)
InDelta(expected, actual interface{}, delta float64, msgAndArgs ...interface{}) {
	debuggoGen_InDelta(expected, actual, delta, msgAndArgs)
}

func // InDeltaSlice is the same as InDelta, except it compares two slices.
InDeltaSlice(expected, actual interface{}, delta float64, msgAndArgs ...interface{}) {
	debuggoGen_InDeltaSlice(expected, actual, delta, msgAndArgs)
}

func // InDeltaMapValues is the same as InDelta, but it compares all values between two maps. Both maps must have exactly the same keys.
InDeltaMapValues(expected, actual interface{}, delta float64, msgAndArgs ...interface{}) {
	debuggoGen_InDeltaMapValues(expected, actual, delta, msgAndArgs)
}

func // InEpsilon asserts that expected and actual have a relative error less than epsilon
InEpsilon(expected, actual interface{}, epsilon float64, msgAndArgs ...interface{}) {
	debuggoGen_InEpsilon(expected, actual, epsilon, msgAndArgs)
}

func // InEpsilonSlice is the same as InEpsilon, except it compares each value from two slices.
InEpsilonSlice(expected, actual interface{}, epsilon float64, msgAndArgs ...interface{}) {
	debuggoGen_InEpsilonSlice(expected, actual, epsilon, msgAndArgs)
}

func // NoError asserts that a function returned no error (i.e. `nil`).
//
//   actualObj, err := SomeFunction()
//   if assert.NoError(t, err) {
//	   assert.Equal(t, expectedObj, actualObj)
//   }
NoError(err error, msgAndArgs ...interface{}) {
	debuggoGen_NoError(err, msgAndArgs)
}

func // Error asserts that a function returned an error (i.e. not `nil`).
//
//   actualObj, err := SomeFunction()
//   if assert.Error(t, err) {
//	   assert.Equal(t, expectedError, err)
//   }
Error(err error, msgAndArgs ...interface{}) {
	debuggoGen_Error(err, msgAndArgs)
}

func // EqualError asserts that a function returned an error (i.e. not `nil`)
// and that it is equal to the provided error.
//
//   actualObj, err := SomeFunction()
//   assert.EqualError(t, err,  expectedErrorString)
EqualError(theError error, errString string, msgAndArgs ...interface{}) {
	debuggoGen_EqualError(theError, errString, msgAndArgs)
}

func // Regexp asserts that a specified regexp matches a string.
//
//  assert.Regexp(t, regexp.MustCompile("start"), "it's starting")
//  assert.Regexp(t, "start...$", "it's not starting")
Regexp(rx interface{}, str interface{}, msgAndArgs ...interface{}) {
	debuggoGen_Regexp(rx, str, msgAndArgs)
}

func // NotRegexp asserts that a specified regexp does not match a string.
//
//  assert.NotRegexp(t, regexp.MustCompile("starts"), "it's starting")
//  assert.NotRegexp(t, "^start", "it's not starting")
NotRegexp(rx interface{}, str interface{}, msgAndArgs ...interface{}) {
	debuggoGen_NotRegexp(rx, str, msgAndArgs)
}

func // Zero asserts that i is the zero value for its type.
Zero(i interface{}, msgAndArgs ...interface{}) {
	debuggoGen_Zero(i, msgAndArgs)
}

func // NotZero asserts that i is not the zero value for its type.
NotZero(i interface{}, msgAndArgs ...interface{}) {
	debuggoGen_NotZero(i, msgAndArgs)
}

func // FileExists checks whether a file exists in the given path. It also fails if
// the path points to a directory or there is an error when trying to check the file.
FileExists(path string, msgAndArgs ...interface{}) {
	debuggoGen_FileExists(path, msgAndArgs)
}

func // NoFileExists checks whether a file does not exist in a given path. It fails
// if the path points to an existing _file_ only.
NoFileExists(path string, msgAndArgs ...interface{}) {
	debuggoGen_NoFileExists(path, msgAndArgs)
}

func // DirExists checks whether a directory exists in the given path. It also fails
// if the path is a file rather a directory or there is an error checking whether it exists.
DirExists(path string, msgAndArgs ...interface{}) {
	debuggoGen_DirExists(path, msgAndArgs)
}

func // NoDirExists checks whether a directory does not exist in the given path.
// It fails if the path points to an existing _directory_ only.
NoDirExists(path string, msgAndArgs ...interface{}) {
	debuggoGen_NoDirExists(path, msgAndArgs)
}

func // JSONEq asserts that two JSON strings are equivalent.
//
//  assert.JSONEq(t, `{"hello": "world", "foo": "bar"}`, `{"foo": "bar", "hello": "world"}`)
JSONEq(expected string, actual string, msgAndArgs ...interface{}) {
	debuggoGen_JSONEq(expected, actual, msgAndArgs)
}

func // YAMLEq asserts that two YAML strings are equivalent.
YAMLEq(expected string, actual string, msgAndArgs ...interface{}) {
	debuggoGen_YAMLEq(expected, actual, msgAndArgs)
}

func // Eventually asserts that given condition will be met in waitFor time,
// periodically checking target function each tick.
//
//    assert.Eventually(t, func() bool { return true; }, time.Second, 10*time.Millisecond)
Eventually(condition func() bool, waitFor time.Duration, tick time.Duration, msgAndArgs ...interface{}) {
	debuggoGen_Eventually(condition, waitFor, tick, msgAndArgs)
}

func // Never asserts that the given condition doesn't satisfy in waitFor time,
// periodically checking the target function each tick.
//
//    assert.Never(t, func() bool { return false; }, time.Second, 10*time.Millisecond)
Never(condition func() bool, waitFor time.Duration, tick time.Duration, msgAndArgs ...interface{}) {
	debuggoGen_Never(condition, waitFor, tick, msgAndArgs)
}

func // ErrorIs asserts that at least one of the errors in err's chain matches target.
// This is a wrapper for errors.Is.
ErrorIs(err, target error, msgAndArgs ...interface{}) {
	debuggoGen_ErrorIs(err, target, msgAndArgs)
}

func // NotErrorIs asserts that at none of the errors in err's chain matches target.
// This is a wrapper for errors.Is.
NotErrorIs(err, target error, msgAndArgs ...interface{}) {
	debuggoGen_NotErrorIs(err, target, msgAndArgs)
}

func // ErrorAs asserts that at least one of the errors in err's chain matches target, and if so, sets target to that error value.
// This is a wrapper for errors.As.
ErrorAs(err error, target interface{}, msgAndArgs ...interface{}) {
	debuggoGen_ErrorAs(err, target, msgAndArgs)
}
