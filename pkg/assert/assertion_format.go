// +build assert

/*
* CODE GENERATED AUTOMATICALLY WITH github.com/stretchr/testify/_codegen
* THIS FILE MUST NOT BE EDITED BY HAND
 */

package assert

import (
	http "net/http"
	url "net/url"
	time "time"
)

// Conditionf uses a Comparison to assert a complex condition.
func Conditionf(comp Comparison, msg string, args ...interface{}) {
	debuggoGen_Conditionf(comp, msg, args)
}

// Containsf asserts that the specified string, list(array, slice...) or map contains the
// specified substring or element.
//
//    assert.Containsf(t, "Hello World", "World", "error message %s", "formatted")
//    assert.Containsf(t, ["Hello", "World"], "World", "error message %s", "formatted")
//    assert.Containsf(t, {"Hello": "World"}, "Hello", "error message %s", "formatted")
func Containsf(s interface{}, contains interface{}, msg string, args ...interface{}) {
	debuggoGen_Containsf(s, contains, msg, args)
}

// DirExistsf checks whether a directory exists in the given path. It also fails
// if the path is a file rather a directory or there is an error checking whether it exists.
func DirExistsf(path string, msg string, args ...interface{}) { debuggoGen_DirExistsf(path, msg, args) }

// ElementsMatchf asserts that the specified listA(array, slice...) is equal to specified
// listB(array, slice...) ignoring the order of the elements. If there are duplicate elements,
// the number of appearances of each of them in both lists should match.
//
// assert.ElementsMatchf(t, [1, 3, 2, 3], [1, 3, 3, 2], "error message %s", "formatted")
func ElementsMatchf(listA interface{}, listB interface{}, msg string, args ...interface{}) {
	debuggoGen_ElementsMatchf(listA, listB, msg, args)
}

// Emptyf asserts that the specified object is empty.  I.e. nil, "", false, 0 or either
// a slice or a channel with len == 0.
//
//  assert.Emptyf(t, obj, "error message %s", "formatted")
func Emptyf(object interface{}, msg string, args ...interface{}) {
	debuggoGen_Emptyf(object, msg, args)
}

// Equalf asserts that two objects are equal.
//
//    assert.Equalf(t, 123, 123, "error message %s", "formatted")
//
// Pointer variable equality is determined based on the equality of the
// referenced values (as opposed to the memory addresses). Function equality
// cannot be determined and will always fail.
func Equalf(expected interface{}, actual interface{}, msg string, args ...interface{}) {
	debuggoGen_Equalf(expected, actual, msg, args)
}

// EqualErrorf asserts that a function returned an error (i.e. not `nil`)
// and that it is equal to the provided error.
//
//   actualObj, err := SomeFunction()
//   assert.EqualErrorf(t, err,  expectedErrorString, "error message %s", "formatted")
func EqualErrorf(theError error, errString string, msg string, args ...interface{}) {
	debuggoGen_EqualErrorf(theError, errString, msg, args)
}

// EqualValuesf asserts that two objects are equal or convertable to the same types
// and equal.
//
//    assert.EqualValuesf(t, uint32(123), int32(123), "error message %s", "formatted")
func EqualValuesf(expected interface{}, actual interface{}, msg string, args ...interface{}) {
	debuggoGen_EqualValuesf(expected, actual, msg, args)
}

// Errorf asserts that a function returned an error (i.e. not `nil`).
//
//   actualObj, err := SomeFunction()
//   if assert.Errorf(t, err, "error message %s", "formatted") {
// 	   assert.Equal(t, expectedErrorf, err)
//   }
func Errorf(err error, msg string, args ...interface{}) { debuggoGen_Errorf(err, msg, args) }

// ErrorAsf asserts that at least one of the errors in err's chain matches target, and if so, sets target to that error value.
// This is a wrapper for errors.As.
func ErrorAsf(err error, target interface{}, msg string, args ...interface{}) {
	debuggoGen_ErrorAsf(err, target, msg, args)
}

// ErrorIsf asserts that at least one of the errors in err's chain matches target.
// This is a wrapper for errors.Is.
func ErrorIsf(err error, target error, msg string, args ...interface{}) {
	debuggoGen_ErrorIsf(err, target, msg, args)
}

// Eventuallyf asserts that given condition will be met in waitFor time,
// periodically checking target function each tick.
//
//    assert.Eventuallyf(t, func() bool { return true; }, time.Second, 10*time.Millisecond, "error message %s", "formatted")
func Eventuallyf(condition func() bool, waitFor time.Duration, tick time.Duration, msg string, args ...interface{}) {
	debuggoGen_Eventuallyf(condition, waitFor, tick, msg, args)
}

// Exactlyf asserts that two objects are equal in value and type.
//
//    assert.Exactlyf(t, int32(123), int64(123), "error message %s", "formatted")
func Exactlyf(expected interface{}, actual interface{}, msg string, args ...interface{}) {
	debuggoGen_Exactlyf(expected, actual, msg, args)
}

// Failf reports a failure through
func Failf(failureMessage string, msg string, args ...interface{}) {
	debuggoGen_Failf(failureMessage, msg, args)
}

// FailNowf fails test
func FailNowf(failureMessage string, msg string, args ...interface{}) {
	debuggoGen_FailNowf(failureMessage, msg, args)
}

// Falsef asserts that the specified value is false.
//
//    assert.Falsef(t, myBool, "error message %s", "formatted")
func Falsef(value bool, msg string, args ...interface{}) { debuggoGen_Falsef(value, msg, args) }

// FileExistsf checks whether a file exists in the given path. It also fails if
// the path points to a directory or there is an error when trying to check the file.
func FileExistsf(path string, msg string, args ...interface{}) {
	debuggoGen_FileExistsf(path, msg, args)
}

// Greaterf asserts that the first element is greater than the second
//
//    assert.Greaterf(t, 2, 1, "error message %s", "formatted")
//    assert.Greaterf(t, float64(2), float64(1), "error message %s", "formatted")
//    assert.Greaterf(t, "b", "a", "error message %s", "formatted")
func Greaterf(e1 interface{}, e2 interface{}, msg string, args ...interface{}) {
	debuggoGen_Greaterf(e1, e2, msg, args)
}

// GreaterOrEqualf asserts that the first element is greater than or equal to the second
//
//    assert.GreaterOrEqualf(t, 2, 1, "error message %s", "formatted")
//    assert.GreaterOrEqualf(t, 2, 2, "error message %s", "formatted")
//    assert.GreaterOrEqualf(t, "b", "a", "error message %s", "formatted")
//    assert.GreaterOrEqualf(t, "b", "b", "error message %s", "formatted")
func GreaterOrEqualf(e1 interface{}, e2 interface{}, msg string, args ...interface{}) {
	debuggoGen_GreaterOrEqualf(e1, e2, msg, args)
}

// HTTPBodyContainsf asserts that a specified handler returns a
// body that contains a string.
//
//  assert.HTTPBodyContainsf(t, myHandler, "GET", "www.google.com", nil, "I'm Feeling Lucky", "error message %s", "formatted")
//
// Returns whether the assertion was successful (true) or not (false).
func HTTPBodyContainsf(handler http.HandlerFunc, method string, url string, values url.Values, str interface{}, msg string, args ...interface{}) {
	debuggoGen_HTTPBodyContainsf(handler, method, url, values, str, msg, args)
}

// HTTPBodyNotContainsf asserts that a specified handler returns a
// body that does not contain a string.
//
//  assert.HTTPBodyNotContainsf(t, myHandler, "GET", "www.google.com", nil, "I'm Feeling Lucky", "error message %s", "formatted")
//
// Returns whether the assertion was successful (true) or not (false).
func HTTPBodyNotContainsf(handler http.HandlerFunc, method string, url string, values url.Values, str interface{}, msg string, args ...interface{}) {
	debuggoGen_HTTPBodyNotContainsf(handler, method, url, values, str, msg, args)
}

// HTTPErrorf asserts that a specified handler returns an error status code.
//
//  assert.HTTPErrorf(t, myHandler, "POST", "/a/b/c", url.Values{"a": []string{"b", "c"}}
//
// Returns whether the assertion was successful (true) or not (false).
func HTTPErrorf(handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...interface{}) {
	debuggoGen_HTTPErrorf(handler, method, url, values, msg, args)
}

// HTTPRedirectf asserts that a specified handler returns a redirect status code.
//
//  assert.HTTPRedirectf(t, myHandler, "GET", "/a/b/c", url.Values{"a": []string{"b", "c"}}
//
// Returns whether the assertion was successful (true) or not (false).
func HTTPRedirectf(handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...interface{}) {
	debuggoGen_HTTPRedirectf(handler, method, url, values, msg, args)
}

// HTTPStatusCodef asserts that a specified handler returns a specified status code.
//
//  assert.HTTPStatusCodef(t, myHandler, "GET", "/notImplemented", nil, 501, "error message %s", "formatted")
//
// Returns whether the assertion was successful (true) or not (false).
func HTTPStatusCodef(handler http.HandlerFunc, method string, url string, values url.Values, statuscode int, msg string, args ...interface{}) {
	debuggoGen_HTTPStatusCodef(handler, method, url, values, statuscode, msg, args)
}

// HTTPSuccessf asserts that a specified handler returns a success status code.
//
//  assert.HTTPSuccessf(t, myHandler, "POST", "http://www.google.com", nil, "error message %s", "formatted")
//
// Returns whether the assertion was successful (true) or not (false).
func HTTPSuccessf(handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...interface{}) {
	debuggoGen_HTTPSuccessf(handler, method, url, values, msg, args)
}

// Implementsf asserts that an object is implemented by the specified interface.
//
//    assert.Implementsf(t, (*MyInterface)(nil), new(MyObject), "error message %s", "formatted")
func Implementsf(interfaceObject interface{}, object interface{}, msg string, args ...interface{}) {
	debuggoGen_Implementsf(interfaceObject, object, msg, args)
}

// InDeltaf asserts that the two numerals are within delta of each other.
//
// 	 assert.InDeltaf(t, math.Pi, 22/7.0, 0.01, "error message %s", "formatted")
func InDeltaf(expected interface{}, actual interface{}, delta float64, msg string, args ...interface{}) {
	debuggoGen_InDeltaf(expected, actual, delta, msg, args)
}

// InDeltaMapValuesf is the same as InDelta, but it compares all values between two maps. Both maps must have exactly the same keys.
func InDeltaMapValuesf(expected interface{}, actual interface{}, delta float64, msg string, args ...interface{}) {
	debuggoGen_InDeltaMapValuesf(expected, actual, delta, msg, args)
}

// InDeltaSlicef is the same as InDelta, except it compares two slices.
func InDeltaSlicef(expected interface{}, actual interface{}, delta float64, msg string, args ...interface{}) {
	debuggoGen_InDeltaSlicef(expected, actual, delta, msg, args)
}

// InEpsilonf asserts that expected and actual have a relative error less than epsilon
func InEpsilonf(expected interface{}, actual interface{}, epsilon float64, msg string, args ...interface{}) {
	debuggoGen_InEpsilonf(expected, actual, epsilon, msg, args)
}

// InEpsilonSlicef is the same as InEpsilon, except it compares each value from two slices.
func InEpsilonSlicef(expected interface{}, actual interface{}, epsilon float64, msg string, args ...interface{}) {
	debuggoGen_InEpsilonSlicef(expected, actual, epsilon, msg, args)
}

// IsDecreasingf asserts that the collection is decreasing
//
//    assert.IsDecreasingf(t, []int{2, 1, 0}, "error message %s", "formatted")
//    assert.IsDecreasingf(t, []float{2, 1}, "error message %s", "formatted")
//    assert.IsDecreasingf(t, []string{"b", "a"}, "error message %s", "formatted")
func IsDecreasingf(object interface{}, msg string, args ...interface{}) {
	debuggoGen_IsDecreasingf(object, msg, args)
}

// IsIncreasingf asserts that the collection is increasing
//
//    assert.IsIncreasingf(t, []int{1, 2, 3}, "error message %s", "formatted")
//    assert.IsIncreasingf(t, []float{1, 2}, "error message %s", "formatted")
//    assert.IsIncreasingf(t, []string{"a", "b"}, "error message %s", "formatted")
func IsIncreasingf(object interface{}, msg string, args ...interface{}) {
	debuggoGen_IsIncreasingf(object, msg, args)
}

// IsNonDecreasingf asserts that the collection is not decreasing
//
//    assert.IsNonDecreasingf(t, []int{1, 1, 2}, "error message %s", "formatted")
//    assert.IsNonDecreasingf(t, []float{1, 2}, "error message %s", "formatted")
//    assert.IsNonDecreasingf(t, []string{"a", "b"}, "error message %s", "formatted")
func IsNonDecreasingf(object interface{}, msg string, args ...interface{}) {
	debuggoGen_IsNonDecreasingf(object, msg, args)
}

// IsNonIncreasingf asserts that the collection is not increasing
//
//    assert.IsNonIncreasingf(t, []int{2, 1, 1}, "error message %s", "formatted")
//    assert.IsNonIncreasingf(t, []float{2, 1}, "error message %s", "formatted")
//    assert.IsNonIncreasingf(t, []string{"b", "a"}, "error message %s", "formatted")
func IsNonIncreasingf(object interface{}, msg string, args ...interface{}) {
	debuggoGen_IsNonIncreasingf(object, msg, args)
}

// IsTypef asserts that the specified objects are of the same type.
func IsTypef(expectedType interface{}, object interface{}, msg string, args ...interface{}) {
	debuggoGen_IsTypef(expectedType, object, msg, args)
}

// JSONEqf asserts that two JSON strings are equivalent.
//
//  assert.JSONEqf(t, `{"hello": "world", "foo": "bar"}`, `{"foo": "bar", "hello": "world"}`, "error message %s", "formatted")
func JSONEqf(expected string, actual string, msg string, args ...interface{}) {
	debuggoGen_JSONEqf(expected, actual, msg, args)
}

// Lenf asserts that the specified object has specific length.
// Lenf also fails if the object has a type that len() not accept.
//
//    assert.Lenf(t, mySlice, 3, "error message %s", "formatted")
func Lenf(object interface{}, length int, msg string, args ...interface{}) {
	debuggoGen_Lenf(object, length, msg, args)
}

// Lessf asserts that the first element is less than the second
//
//    assert.Lessf(t, 1, 2, "error message %s", "formatted")
//    assert.Lessf(t, float64(1), float64(2), "error message %s", "formatted")
//    assert.Lessf(t, "a", "b", "error message %s", "formatted")
func Lessf(e1 interface{}, e2 interface{}, msg string, args ...interface{}) {
	debuggoGen_Lessf(e1, e2, msg, args)
}

// LessOrEqualf asserts that the first element is less than or equal to the second
//
//    assert.LessOrEqualf(t, 1, 2, "error message %s", "formatted")
//    assert.LessOrEqualf(t, 2, 2, "error message %s", "formatted")
//    assert.LessOrEqualf(t, "a", "b", "error message %s", "formatted")
//    assert.LessOrEqualf(t, "b", "b", "error message %s", "formatted")
func LessOrEqualf(e1 interface{}, e2 interface{}, msg string, args ...interface{}) {
	debuggoGen_LessOrEqualf(e1, e2, msg, args)
}

// Neverf asserts that the given condition doesn't satisfy in waitFor time,
// periodically checking the target function each tick.
//
//    assert.Neverf(t, func() bool { return false; }, time.Second, 10*time.Millisecond, "error message %s", "formatted")
func Neverf(condition func() bool, waitFor time.Duration, tick time.Duration, msg string, args ...interface{}) {
	debuggoGen_Neverf(condition, waitFor, tick, msg, args)
}

// Nilf asserts that the specified object is nil.
//
//    assert.Nilf(t, err, "error message %s", "formatted")
func Nilf(object interface{}, msg string, args ...interface{}) { debuggoGen_Nilf(object, msg, args) }

// NoDirExistsf checks whether a directory does not exist in the given path.
// It fails if the path points to an existing _directory_ only.
func NoDirExistsf(path string, msg string, args ...interface{}) {
	debuggoGen_NoDirExistsf(path, msg, args)
}

// NoErrorf asserts that a function returned no error (i.e. `nil`).
//
//   actualObj, err := SomeFunction()
//   if assert.NoErrorf(t, err, "error message %s", "formatted") {
// 	   assert.Equal(t, expectedObj, actualObj)
//   }
func NoErrorf(err error, msg string, args ...interface{}) { debuggoGen_NoErrorf(err, msg, args) }

// NoFileExistsf checks whether a file does not exist in a given path. It fails
// if the path points to an existing _file_ only.
func NoFileExistsf(path string, msg string, args ...interface{}) {
	debuggoGen_NoFileExistsf(path, msg, args)
}

// NotContainsf asserts that the specified string, list(array, slice...) or map does NOT contain the
// specified substring or element.
//
//    assert.NotContainsf(t, "Hello World", "Earth", "error message %s", "formatted")
//    assert.NotContainsf(t, ["Hello", "World"], "Earth", "error message %s", "formatted")
//    assert.NotContainsf(t, {"Hello": "World"}, "Earth", "error message %s", "formatted")
func NotContainsf(s interface{}, contains interface{}, msg string, args ...interface{}) {
	debuggoGen_NotContainsf(s, contains, msg, args)
}

// NotEmptyf asserts that the specified object is NOT empty.  I.e. not nil, "", false, 0 or either
// a slice or a channel with len == 0.
//
//  if assert.NotEmptyf(t, obj, "error message %s", "formatted") {
//    assert.Equal(t, "two", obj[1])
//  }
func NotEmptyf(object interface{}, msg string, args ...interface{}) {
	debuggoGen_NotEmptyf(object, msg, args)
}

// NotEqualf asserts that the specified values are NOT equal.
//
//    assert.NotEqualf(t, obj1, obj2, "error message %s", "formatted")
//
// Pointer variable equality is determined based on the equality of the
// referenced values (as opposed to the memory addresses).
func NotEqualf(expected interface{}, actual interface{}, msg string, args ...interface{}) {
	debuggoGen_NotEqualf(expected, actual, msg, args)
}

// NotEqualValuesf asserts that two objects are not equal even when converted to the same type
//
//    assert.NotEqualValuesf(t, obj1, obj2, "error message %s", "formatted")
func NotEqualValuesf(expected interface{}, actual interface{}, msg string, args ...interface{}) {
	debuggoGen_NotEqualValuesf(expected, actual, msg, args)
}

// NotErrorIsf asserts that at none of the errors in err's chain matches target.
// This is a wrapper for errors.Is.
func NotErrorIsf(err error, target error, msg string, args ...interface{}) {
	debuggoGen_NotErrorIsf(err, target, msg, args)
}

// NotNilf asserts that the specified object is not nil.
//
//    assert.NotNilf(t, err, "error message %s", "formatted")
func NotNilf(object interface{}, msg string, args ...interface{}) {
	debuggoGen_NotNilf(object, msg, args)
}

// NotPanicsf asserts that the code inside the specified PanicTestFunc does NOT panic.
//
//   assert.NotPanicsf(t, func(){ RemainCalm() }, "error message %s", "formatted")
func NotPanicsf(f PanicTestFunc, msg string, args ...interface{}) {
	debuggoGen_NotPanicsf(f, msg, args)
}

// NotRegexpf asserts that a specified regexp does not match a string.
//
//  assert.NotRegexpf(t, regexp.MustCompile("starts"), "it's starting", "error message %s", "formatted")
//  assert.NotRegexpf(t, "^start", "it's not starting", "error message %s", "formatted")
func NotRegexpf(rx interface{}, str interface{}, msg string, args ...interface{}) {
	debuggoGen_NotRegexpf(rx, str, msg, args)
}

// NotSamef asserts that two pointers do not reference the same object.
//
//    assert.NotSamef(t, ptr1, ptr2, "error message %s", "formatted")
//
// Both arguments must be pointer variables. Pointer variable sameness is
// determined based on the equality of both type and value.
func NotSamef(expected interface{}, actual interface{}, msg string, args ...interface{}) {
	debuggoGen_NotSamef(expected, actual, msg, args)
}

// NotSubsetf asserts that the specified list(array, slice...) contains not all
// elements given in the specified subset(array, slice...).
//
//    assert.NotSubsetf(t, [1, 3, 4], [1, 2], "But [1, 3, 4] does not contain [1, 2]", "error message %s", "formatted")
func NotSubsetf(list interface{}, subset interface{}, msg string, args ...interface{}) {
	debuggoGen_NotSubsetf(list, subset, msg, args)
}

// NotZerof asserts that i is not the zero value for its type.
func NotZerof(i interface{}, msg string, args ...interface{}) { debuggoGen_NotZerof(i, msg, args) }

// Panicsf asserts that the code inside the specified PanicTestFunc panics.
//
//   assert.Panicsf(t, func(){ GoCrazy() }, "error message %s", "formatted")
func Panicsf(f PanicTestFunc, msg string, args ...interface{}) { debuggoGen_Panicsf(f, msg, args) }

// PanicsWithErrorf asserts that the code inside the specified PanicTestFunc
// panics, and that the recovered panic value is an error that satisfies the
// EqualError comparison.
//
//   assert.PanicsWithErrorf(t, "crazy error", func(){ GoCrazy() }, "error message %s", "formatted")
func PanicsWithErrorf(errString string, f PanicTestFunc, msg string, args ...interface{}) {
	debuggoGen_PanicsWithErrorf(errString, f, msg, args)
}

// PanicsWithValuef asserts that the code inside the specified PanicTestFunc panics, and that
// the recovered panic value equals the expected panic value.
//
//   assert.PanicsWithValuef(t, "crazy error", func(){ GoCrazy() }, "error message %s", "formatted")
func PanicsWithValuef(expected interface{}, f PanicTestFunc, msg string, args ...interface{}) {
	debuggoGen_PanicsWithValuef(expected, f, msg, args)
}

// Regexpf asserts that a specified regexp matches a string.
//
//  assert.Regexpf(t, regexp.MustCompile("start"), "it's starting", "error message %s", "formatted")
//  assert.Regexpf(t, "start...$", "it's not starting", "error message %s", "formatted")
func Regexpf(rx interface{}, str interface{}, msg string, args ...interface{}) {
	debuggoGen_Regexpf(rx, str, msg, args)
}

// Samef asserts that two pointers reference the same object.
//
//    assert.Samef(t, ptr1, ptr2, "error message %s", "formatted")
//
// Both arguments must be pointer variables. Pointer variable sameness is
// determined based on the equality of both type and value.
func Samef(expected interface{}, actual interface{}, msg string, args ...interface{}) {
	debuggoGen_Samef(expected, actual, msg, args)
}

// Subsetf asserts that the specified list(array, slice...) contains all
// elements given in the specified subset(array, slice...).
//
//    assert.Subsetf(t, [1, 2, 3], [1, 2], "But [1, 2, 3] does contain [1, 2]", "error message %s", "formatted")
func Subsetf(list interface{}, subset interface{}, msg string, args ...interface{}) {
	debuggoGen_Subsetf(list, subset, msg, args)
}

// Truef asserts that the specified value is true.
//
//    assert.Truef(t, myBool, "error message %s", "formatted")
func Truef(value bool, msg string, args ...interface{}) { debuggoGen_Truef(value, msg, args) }

// WithinDurationf asserts that the two times are within duration delta of each other.
//
//   assert.WithinDurationf(t, time.Now(), time.Now(), 10*time.Second, "error message %s", "formatted")
func WithinDurationf(expected time.Time, actual time.Time, delta time.Duration, msg string, args ...interface{}) {
	debuggoGen_WithinDurationf(expected, actual, delta, msg, args)
}

// YAMLEqf asserts that two YAML strings are equivalent.
func YAMLEqf(expected string, actual string, msg string, args ...interface{}) {
	debuggoGen_YAMLEqf(expected, actual, msg, args)
}

// Zerof asserts that i is the zero value for its type.
func Zerof(i interface{}, msg string, args ...interface{}) { debuggoGen_Zerof(i, msg, args) }

func debuggoGen_Conditionf(comp Comparison, msg string, args ...interface{}) bool {

	return debuggoGen_Condition(comp, append([]interface{}{msg}, args...)...)
}

func debuggoGen_Containsf(s interface{}, contains interface{}, msg string, args ...interface{}) bool {

	return debuggoGen_Contains(s, contains, append([]interface{}{msg}, args...)...)
}

func debuggoGen_DirExistsf(path string, msg string, args ...interface{}) bool {

	return debuggoGen_DirExists(path, append([]interface{}{msg}, args...)...)
}

func debuggoGen_ElementsMatchf(listA interface{}, listB interface{}, msg string, args ...interface{}) bool {

	return debuggoGen_ElementsMatch(listA, listB, append([]interface{}{msg}, args...)...)
}

func debuggoGen_Emptyf(object interface{}, msg string, args ...interface{}) bool {

	return debuggoGen_Empty(object, append([]interface{}{msg}, args...)...)
}

func debuggoGen_Equalf(expected interface{}, actual interface{}, msg string, args ...interface{}) bool {

	return debuggoGen_Equal(expected, actual, append([]interface{}{msg}, args...)...)
}

func debuggoGen_EqualErrorf(theError error, errString string, msg string, args ...interface{}) bool {

	return debuggoGen_EqualError(theError, errString, append([]interface{}{msg}, args...)...)
}

func debuggoGen_EqualValuesf(expected interface{}, actual interface{}, msg string, args ...interface{}) bool {

	return debuggoGen_EqualValues(expected, actual, append([]interface{}{msg}, args...)...)
}

func debuggoGen_Errorf(err error, msg string, args ...interface{}) bool {

	return debuggoGen_Error(err, append([]interface{}{msg}, args...)...)
}

func debuggoGen_ErrorAsf(err error, target interface{}, msg string, args ...interface{}) bool {

	return debuggoGen_ErrorAs(err, target, append([]interface{}{msg}, args...)...)
}

func debuggoGen_ErrorIsf(err error, target error, msg string, args ...interface{}) bool {

	return debuggoGen_ErrorIs(err, target, append([]interface{}{msg}, args...)...)
}

func debuggoGen_Eventuallyf(condition func() bool, waitFor time.Duration, tick time.Duration, msg string, args ...interface{}) bool {

	return debuggoGen_Eventually(condition, waitFor, tick, append([]interface{}{msg}, args...)...)
}

func debuggoGen_Exactlyf(expected interface{}, actual interface{}, msg string, args ...interface{}) bool {

	return debuggoGen_Exactly(expected, actual, append([]interface{}{msg}, args...)...)
}

func debuggoGen_Failf(failureMessage string, msg string, args ...interface{}) bool {

	return debuggoGen_Fail(failureMessage, append([]interface{}{msg}, args...)...)
}

func debuggoGen_FailNowf(failureMessage string, msg string, args ...interface{}) bool {

	return debuggoGen_FailNow(failureMessage, append([]interface{}{msg}, args...)...)
}

func debuggoGen_Falsef(value bool, msg string, args ...interface{}) bool {

	return debuggoGen_False(value, append([]interface{}{msg}, args...)...)
}

func debuggoGen_FileExistsf(path string, msg string, args ...interface{}) bool {

	return debuggoGen_FileExists(path, append([]interface{}{msg}, args...)...)
}

func debuggoGen_Greaterf(e1 interface{}, e2 interface{}, msg string, args ...interface{}) bool {

	return debuggoGen_Greater(e1, e2, append([]interface{}{msg}, args...)...)
}

func debuggoGen_GreaterOrEqualf(e1 interface{}, e2 interface{}, msg string, args ...interface{}) bool {

	return debuggoGen_GreaterOrEqual(e1, e2, append([]interface{}{msg}, args...)...)
}

func debuggoGen_HTTPBodyContainsf(handler http.HandlerFunc, method string, url string, values url.Values, str interface{}, msg string, args ...interface{}) bool {

	return debuggoGen_HTTPBodyContains(handler, method, url, values, str, append([]interface{}{msg}, args...)...)
}

func debuggoGen_HTTPBodyNotContainsf(handler http.HandlerFunc, method string, url string, values url.Values, str interface{}, msg string, args ...interface{}) bool {

	return debuggoGen_HTTPBodyNotContains(handler, method, url, values, str, append([]interface{}{msg}, args...)...)
}

func debuggoGen_HTTPErrorf(handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...interface{}) bool {

	return debuggoGen_HTTPError(handler, method, url, values, append([]interface{}{msg}, args...)...)
}

func debuggoGen_HTTPRedirectf(handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...interface{}) bool {

	return debuggoGen_HTTPRedirect(handler, method, url, values, append([]interface{}{msg}, args...)...)
}

func debuggoGen_HTTPStatusCodef(handler http.HandlerFunc, method string, url string, values url.Values, statuscode int, msg string, args ...interface{}) bool {

	return debuggoGen_HTTPStatusCode(handler, method, url, values, statuscode, append([]interface{}{msg}, args...)...)
}

func debuggoGen_HTTPSuccessf(handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...interface{}) bool {

	return debuggoGen_HTTPSuccess(handler, method, url, values, append([]interface{}{msg}, args...)...)
}

func debuggoGen_Implementsf(interfaceObject interface{}, object interface{}, msg string, args ...interface{}) bool {

	return debuggoGen_Implements(interfaceObject, object, append([]interface{}{msg}, args...)...)
}

func debuggoGen_InDeltaf(expected interface{}, actual interface{}, delta float64, msg string, args ...interface{}) bool {

	return debuggoGen_InDelta(expected, actual, delta, append([]interface{}{msg}, args...)...)
}

func debuggoGen_InDeltaMapValuesf(expected interface{}, actual interface{}, delta float64, msg string, args ...interface{}) bool {

	return debuggoGen_InDeltaMapValues(expected, actual, delta, append([]interface{}{msg}, args...)...)
}

func debuggoGen_InDeltaSlicef(expected interface{}, actual interface{}, delta float64, msg string, args ...interface{}) bool {

	return debuggoGen_InDeltaSlice(expected, actual, delta, append([]interface{}{msg}, args...)...)
}

func debuggoGen_InEpsilonf(expected interface{}, actual interface{}, epsilon float64, msg string, args ...interface{}) bool {

	return debuggoGen_InEpsilon(expected, actual, epsilon, append([]interface{}{msg}, args...)...)
}

func debuggoGen_InEpsilonSlicef(expected interface{}, actual interface{}, epsilon float64, msg string, args ...interface{}) bool {

	return debuggoGen_InEpsilonSlice(expected, actual, epsilon, append([]interface{}{msg}, args...)...)
}

func debuggoGen_IsDecreasingf(object interface{}, msg string, args ...interface{}) bool {

	return debuggoGen_IsDecreasing(object, append([]interface{}{msg}, args...)...)
}

func debuggoGen_IsIncreasingf(object interface{}, msg string, args ...interface{}) bool {

	return debuggoGen_IsIncreasing(object, append([]interface{}{msg}, args...)...)
}

func debuggoGen_IsNonDecreasingf(object interface{}, msg string, args ...interface{}) bool {

	return debuggoGen_IsNonDecreasing(object, append([]interface{}{msg}, args...)...)
}

func debuggoGen_IsNonIncreasingf(object interface{}, msg string, args ...interface{}) bool {

	return debuggoGen_IsNonIncreasing(object, append([]interface{}{msg}, args...)...)
}

func debuggoGen_IsTypef(expectedType interface{}, object interface{}, msg string, args ...interface{}) bool {

	return debuggoGen_IsType(expectedType, object, append([]interface{}{msg}, args...)...)
}

func debuggoGen_JSONEqf(expected string, actual string, msg string, args ...interface{}) bool {

	return debuggoGen_JSONEq(expected, actual, append([]interface{}{msg}, args...)...)
}

func debuggoGen_Lenf(object interface{}, length int, msg string, args ...interface{}) bool {

	return debuggoGen_Len(object, length, append([]interface{}{msg}, args...)...)
}

func debuggoGen_Lessf(e1 interface{}, e2 interface{}, msg string, args ...interface{}) bool {

	return debuggoGen_Less(e1, e2, append([]interface{}{msg}, args...)...)
}

func debuggoGen_LessOrEqualf(e1 interface{}, e2 interface{}, msg string, args ...interface{}) bool {

	return debuggoGen_LessOrEqual(e1, e2, append([]interface{}{msg}, args...)...)
}

func debuggoGen_Neverf(condition func() bool, waitFor time.Duration, tick time.Duration, msg string, args ...interface{}) bool {

	return debuggoGen_Never(condition, waitFor, tick, append([]interface{}{msg}, args...)...)
}

func debuggoGen_Nilf(object interface{}, msg string, args ...interface{}) bool {

	return debuggoGen_Nil(object, append([]interface{}{msg}, args...)...)
}

func debuggoGen_NoDirExistsf(path string, msg string, args ...interface{}) bool {

	return debuggoGen_NoDirExists(path, append([]interface{}{msg}, args...)...)
}

func debuggoGen_NoErrorf(err error, msg string, args ...interface{}) bool {

	return debuggoGen_NoError(err, append([]interface{}{msg}, args...)...)
}

func debuggoGen_NoFileExistsf(path string, msg string, args ...interface{}) bool {

	return debuggoGen_NoFileExists(path, append([]interface{}{msg}, args...)...)
}

func debuggoGen_NotContainsf(s interface{}, contains interface{}, msg string, args ...interface{}) bool {

	return debuggoGen_NotContains(s, contains, append([]interface{}{msg}, args...)...)
}

func debuggoGen_NotEmptyf(object interface{}, msg string, args ...interface{}) bool {

	return debuggoGen_NotEmpty(object, append([]interface{}{msg}, args...)...)
}

func debuggoGen_NotEqualf(expected interface{}, actual interface{}, msg string, args ...interface{}) bool {

	return debuggoGen_NotEqual(expected, actual, append([]interface{}{msg}, args...)...)
}

func debuggoGen_NotEqualValuesf(expected interface{}, actual interface{}, msg string, args ...interface{}) bool {

	return debuggoGen_NotEqualValues(expected, actual, append([]interface{}{msg}, args...)...)
}

func debuggoGen_NotErrorIsf(err error, target error, msg string, args ...interface{}) bool {

	return debuggoGen_NotErrorIs(err, target, append([]interface{}{msg}, args...)...)
}

func debuggoGen_NotNilf(object interface{}, msg string, args ...interface{}) bool {

	return debuggoGen_NotNil(object, append([]interface{}{msg}, args...)...)
}

func debuggoGen_NotPanicsf(f PanicTestFunc, msg string, args ...interface{}) bool {

	return debuggoGen_NotPanics(f, append([]interface{}{msg}, args...)...)
}

func debuggoGen_NotRegexpf(rx interface{}, str interface{}, msg string, args ...interface{}) bool {

	return debuggoGen_NotRegexp(rx, str, append([]interface{}{msg}, args...)...)
}

func debuggoGen_NotSamef(expected interface{}, actual interface{}, msg string, args ...interface{}) bool {

	return debuggoGen_NotSame(expected, actual, append([]interface{}{msg}, args...)...)
}

func debuggoGen_NotSubsetf(list interface{}, subset interface{}, msg string, args ...interface{}) bool {

	return debuggoGen_NotSubset(list, subset, append([]interface{}{msg}, args...)...)
}

func debuggoGen_NotZerof(i interface{}, msg string, args ...interface{}) bool {

	return debuggoGen_NotZero(i, append([]interface{}{msg}, args...)...)
}

func debuggoGen_Panicsf(f PanicTestFunc, msg string, args ...interface{}) bool {

	return debuggoGen_Panics(f, append([]interface{}{msg}, args...)...)
}

func debuggoGen_PanicsWithErrorf(errString string, f PanicTestFunc, msg string, args ...interface{}) bool {

	return debuggoGen_PanicsWithError(errString, f, append([]interface{}{msg}, args...)...)
}

func debuggoGen_PanicsWithValuef(expected interface{}, f PanicTestFunc, msg string, args ...interface{}) bool {

	return debuggoGen_PanicsWithValue(expected, f, append([]interface{}{msg}, args...)...)
}

func debuggoGen_Regexpf(rx interface{}, str interface{}, msg string, args ...interface{}) bool {

	return debuggoGen_Regexp(rx, str, append([]interface{}{msg}, args...)...)
}

func debuggoGen_Samef(expected interface{}, actual interface{}, msg string, args ...interface{}) bool {

	return debuggoGen_Same(expected, actual, append([]interface{}{msg}, args...)...)
}

func debuggoGen_Subsetf(list interface{}, subset interface{}, msg string, args ...interface{}) bool {

	return debuggoGen_Subset(list, subset, append([]interface{}{msg}, args...)...)
}

func debuggoGen_Truef(value bool, msg string, args ...interface{}) bool {

	return debuggoGen_True(value, append([]interface{}{msg}, args...)...)
}

func debuggoGen_WithinDurationf(expected time.Time, actual time.Time, delta time.Duration, msg string, args ...interface{}) bool {

	return debuggoGen_WithinDuration(expected, actual, delta, append([]interface{}{msg}, args...)...)
}

func debuggoGen_YAMLEqf(expected string, actual string, msg string, args ...interface{}) bool {

	return debuggoGen_YAMLEq(expected, actual, append([]interface{}{msg}, args...)...)
}

func debuggoGen_Zerof(i interface{}, msg string, args ...interface{}) bool {

	return debuggoGen_Zero(i, append([]interface{}{msg}, args...)...)
}
