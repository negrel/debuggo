// +build assert

package assert

import (
	http "net/http"
	url "net/url"
	time "time"
)

func // Conditionf uses a Comparison to assert a complex condition.
Conditionf(comp Comparison, msg string, args ...interface{}) {
	debuggoGen_Conditionf(comp, msg, args)
}

func // Containsf asserts that the specified string, list(array, slice...) or map contains the
// specified substring or element.
//
//    assert.Containsf(t, "Hello World", "World", "error message %s", "formatted")
//    assert.Containsf(t, ["Hello", "World"], "World", "error message %s", "formatted")
//    assert.Containsf(t, {"Hello": "World"}, "Hello", "error message %s", "formatted")
Containsf(s interface{}, contains interface{}, msg string, args ...interface{}) {
	debuggoGen_Containsf(s, contains, msg, args)
}

func // DirExistsf checks whether a directory exists in the given path. It also fails
// if the path is a file rather a directory or there is an error checking whether it exists.
DirExistsf(path string, msg string, args ...interface{}) {
	debuggoGen_DirExistsf(path, msg, args)
}

func // ElementsMatchf asserts that the specified listA(array, slice...) is equal to specified
// listB(array, slice...) ignoring the order of the elements. If there are duplicate elements,
// the number of appearances of each of them in both lists should match.
//
// assert.ElementsMatchf(t, [1, 3, 2, 3], [1, 3, 3, 2], "error message %s", "formatted")
ElementsMatchf(listA interface{}, listB interface{}, msg string, args ...interface{}) {
	debuggoGen_ElementsMatchf(listA, listB, msg, args)
}

func // Emptyf asserts that the specified object is empty.  I.e. nil, "", false, 0 or either
// a slice or a channel with len == 0.
//
//  assert.Emptyf(t, obj, "error message %s", "formatted")
Emptyf(object interface{}, msg string, args ...interface{}) {
	debuggoGen_Emptyf(object, msg, args)
}

func // Equalf asserts that two objects are equal.
//
//    assert.Equalf(t, 123, 123, "error message %s", "formatted")
//
// Pointer variable equality is determined based on the equality of the
// referenced values (as opposed to the memory addresses). Function equality
// cannot be determined and will always fail.
Equalf(expected interface{}, actual interface{}, msg string, args ...interface{}) {
	debuggoGen_Equalf(expected, actual, msg, args)
}

func // EqualErrorf asserts that a function returned an error (i.e. not `nil`)
// and that it is equal to the provided error.
//
//   actualObj, err := SomeFunction()
//   assert.EqualErrorf(t, err,  expectedErrorString, "error message %s", "formatted")
EqualErrorf(theError error, errString string, msg string, args ...interface{}) {
	debuggoGen_EqualErrorf(theError, errString, msg, args)
}

func // EqualValuesf asserts that two objects are equal or convertable to the same types
// and equal.
//
//    assert.EqualValuesf(t, uint32(123), int32(123), "error message %s", "formatted")
EqualValuesf(expected interface{}, actual interface{}, msg string, args ...interface{}) {
	debuggoGen_EqualValuesf(expected, actual, msg, args)
}

func // Errorf asserts that a function returned an error (i.e. not `nil`).
//
//   actualObj, err := SomeFunction()
//   if assert.Errorf(t, err, "error message %s", "formatted") {
// 	   assert.Equal(t, expectedErrorf, err)
//   }
Errorf(err error, msg string, args ...interface{}) {
	debuggoGen_Errorf(err, msg, args)
}

func // ErrorAsf asserts that at least one of the errors in err's chain matches target, and if so, sets target to that error value.
// This is a wrapper for errors.As.
ErrorAsf(err error, target interface{}, msg string, args ...interface{}) {
	debuggoGen_ErrorAsf(err, target, msg, args)
}

func // ErrorIsf asserts that at least one of the errors in err's chain matches target.
// This is a wrapper for errors.Is.
ErrorIsf(err error, target error, msg string, args ...interface{}) {
	debuggoGen_ErrorIsf(err, target, msg, args)
}

func // Eventuallyf asserts that given condition will be met in waitFor time,
// periodically checking target function each tick.
//
//    assert.Eventuallyf(t, func() bool { return true; }, time.Second, 10*time.Millisecond, "error message %s", "formatted")
Eventuallyf(condition func() bool, waitFor time.Duration, tick time.Duration, msg string, args ...interface{}) {
	debuggoGen_Eventuallyf(condition, waitFor, tick, msg, args)
}

func // Exactlyf asserts that two objects are equal in value and type.
//
//    assert.Exactlyf(t, int32(123), int64(123), "error message %s", "formatted")
Exactlyf(expected interface{}, actual interface{}, msg string, args ...interface{}) {
	debuggoGen_Exactlyf(expected, actual, msg, args)
}

func // Failf reports a failure through
Failf(failureMessage string, msg string, args ...interface{}) {
	debuggoGen_Failf(failureMessage, msg, args)
}

func // FailNowf fails test
FailNowf(failureMessage string, msg string, args ...interface{}) {
	debuggoGen_FailNowf(failureMessage, msg, args)
}

func // Falsef asserts that the specified value is false.
//
//    assert.Falsef(t, myBool, "error message %s", "formatted")
Falsef(value bool, msg string, args ...interface{}) {
	debuggoGen_Falsef(value, msg, args)
}

func // FileExistsf checks whether a file exists in the given path. It also fails if
// the path points to a directory or there is an error when trying to check the file.
FileExistsf(path string, msg string, args ...interface{}) {
	debuggoGen_FileExistsf(path, msg, args)
}

func // Greaterf asserts that the first element is greater than the second
//
//    assert.Greaterf(t, 2, 1, "error message %s", "formatted")
//    assert.Greaterf(t, float64(2), float64(1), "error message %s", "formatted")
//    assert.Greaterf(t, "b", "a", "error message %s", "formatted")
Greaterf(e1 interface{}, e2 interface{}, msg string, args ...interface{}) {
	debuggoGen_Greaterf(e1, e2, msg, args)
}

func // GreaterOrEqualf asserts that the first element is greater than or equal to the second
//
//    assert.GreaterOrEqualf(t, 2, 1, "error message %s", "formatted")
//    assert.GreaterOrEqualf(t, 2, 2, "error message %s", "formatted")
//    assert.GreaterOrEqualf(t, "b", "a", "error message %s", "formatted")
//    assert.GreaterOrEqualf(t, "b", "b", "error message %s", "formatted")
GreaterOrEqualf(e1 interface{}, e2 interface{}, msg string, args ...interface{}) {
	debuggoGen_GreaterOrEqualf(e1, e2, msg, args)
}

func // HTTPBodyContainsf asserts that a specified handler returns a
// body that contains a string.
//
//  assert.HTTPBodyContainsf(t, myHandler, "GET", "www.google.com", nil, "I'm Feeling Lucky", "error message %s", "formatted")
//
// Returns whether the assertion was successful (true) or not (false).
HTTPBodyContainsf(handler http.HandlerFunc, method string, url string, values url.Values, str interface{}, msg string, args ...interface{}) {
	debuggoGen_HTTPBodyContainsf(handler, method, url, values, str, msg, args)
}

func // HTTPBodyNotContainsf asserts that a specified handler returns a
// body that does not contain a string.
//
//  assert.HTTPBodyNotContainsf(t, myHandler, "GET", "www.google.com", nil, "I'm Feeling Lucky", "error message %s", "formatted")
//
// Returns whether the assertion was successful (true) or not (false).
HTTPBodyNotContainsf(handler http.HandlerFunc, method string, url string, values url.Values, str interface{}, msg string, args ...interface{}) {
	debuggoGen_HTTPBodyNotContainsf(handler, method, url, values, str, msg, args)
}

func // HTTPErrorf asserts that a specified handler returns an error status code.
//
//  assert.HTTPErrorf(t, myHandler, "POST", "/a/b/c", url.Values{"a": []string{"b", "c"}}
//
// Returns whether the assertion was successful (true) or not (false).
HTTPErrorf(handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...interface{}) {
	debuggoGen_HTTPErrorf(handler, method, url, values, msg, args)
}

func // HTTPRedirectf asserts that a specified handler returns a redirect status code.
//
//  assert.HTTPRedirectf(t, myHandler, "GET", "/a/b/c", url.Values{"a": []string{"b", "c"}}
//
// Returns whether the assertion was successful (true) or not (false).
HTTPRedirectf(handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...interface{}) {
	debuggoGen_HTTPRedirectf(handler, method, url, values, msg, args)
}

func // HTTPStatusCodef asserts that a specified handler returns a specified status code.
//
//  assert.HTTPStatusCodef(t, myHandler, "GET", "/notImplemented", nil, 501, "error message %s", "formatted")
//
// Returns whether the assertion was successful (true) or not (false).
HTTPStatusCodef(handler http.HandlerFunc, method string, url string, values url.Values, statuscode int, msg string, args ...interface{}) {
	debuggoGen_HTTPStatusCodef(handler, method, url, values, statuscode, msg, args)
}

func // HTTPSuccessf asserts that a specified handler returns a success status code.
//
//  assert.HTTPSuccessf(t, myHandler, "POST", "http://www.google.com", nil, "error message %s", "formatted")
//
// Returns whether the assertion was successful (true) or not (false).
HTTPSuccessf(handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...interface{}) {
	debuggoGen_HTTPSuccessf(handler, method, url, values, msg, args)
}

func // Implementsf asserts that an object is implemented by the specified interface.
//
//    assert.Implementsf(t, (*MyInterface)(nil), new(MyObject), "error message %s", "formatted")
Implementsf(interfaceObject interface{}, object interface{}, msg string, args ...interface{}) {
	debuggoGen_Implementsf(interfaceObject, object, msg, args)
}

func // InDeltaf asserts that the two numerals are within delta of each other.
//
// 	 assert.InDeltaf(t, math.Pi, 22/7.0, 0.01, "error message %s", "formatted")
InDeltaf(expected interface{}, actual interface{}, delta float64, msg string, args ...interface{}) {
	debuggoGen_InDeltaf(expected, actual, delta, msg, args)
}

func // InDeltaMapValuesf is the same as InDelta, but it compares all values between two maps. Both maps must have exactly the same keys.
InDeltaMapValuesf(expected interface{}, actual interface{}, delta float64, msg string, args ...interface{}) {
	debuggoGen_InDeltaMapValuesf(expected, actual, delta, msg, args)
}

func // InDeltaSlicef is the same as InDelta, except it compares two slices.
InDeltaSlicef(expected interface{}, actual interface{}, delta float64, msg string, args ...interface{}) {
	debuggoGen_InDeltaSlicef(expected, actual, delta, msg, args)
}

func // InEpsilonf asserts that expected and actual have a relative error less than epsilon
InEpsilonf(expected interface{}, actual interface{}, epsilon float64, msg string, args ...interface{}) {
	debuggoGen_InEpsilonf(expected, actual, epsilon, msg, args)
}

func // InEpsilonSlicef is the same as InEpsilon, except it compares each value from two slices.
InEpsilonSlicef(expected interface{}, actual interface{}, epsilon float64, msg string, args ...interface{}) {
	debuggoGen_InEpsilonSlicef(expected, actual, epsilon, msg, args)
}

func // IsDecreasingf asserts that the collection is decreasing
//
//    assert.IsDecreasingf(t, []int{2, 1, 0}, "error message %s", "formatted")
//    assert.IsDecreasingf(t, []float{2, 1}, "error message %s", "formatted")
//    assert.IsDecreasingf(t, []string{"b", "a"}, "error message %s", "formatted")
IsDecreasingf(object interface{}, msg string, args ...interface{}) {
	debuggoGen_IsDecreasingf(object, msg, args)
}

func // IsIncreasingf asserts that the collection is increasing
//
//    assert.IsIncreasingf(t, []int{1, 2, 3}, "error message %s", "formatted")
//    assert.IsIncreasingf(t, []float{1, 2}, "error message %s", "formatted")
//    assert.IsIncreasingf(t, []string{"a", "b"}, "error message %s", "formatted")
IsIncreasingf(object interface{}, msg string, args ...interface{}) {
	debuggoGen_IsIncreasingf(object, msg, args)
}

func // IsNonDecreasingf asserts that the collection is not decreasing
//
//    assert.IsNonDecreasingf(t, []int{1, 1, 2}, "error message %s", "formatted")
//    assert.IsNonDecreasingf(t, []float{1, 2}, "error message %s", "formatted")
//    assert.IsNonDecreasingf(t, []string{"a", "b"}, "error message %s", "formatted")
IsNonDecreasingf(object interface{}, msg string, args ...interface{}) {
	debuggoGen_IsNonDecreasingf(object, msg, args)
}

func // IsNonIncreasingf asserts that the collection is not increasing
//
//    assert.IsNonIncreasingf(t, []int{2, 1, 1}, "error message %s", "formatted")
//    assert.IsNonIncreasingf(t, []float{2, 1}, "error message %s", "formatted")
//    assert.IsNonIncreasingf(t, []string{"b", "a"}, "error message %s", "formatted")
IsNonIncreasingf(object interface{}, msg string, args ...interface{}) {
	debuggoGen_IsNonIncreasingf(object, msg, args)
}

func // IsTypef asserts that the specified objects are of the same type.
IsTypef(expectedType interface{}, object interface{}, msg string, args ...interface{}) {
	debuggoGen_IsTypef(expectedType, object, msg, args)
}

func // JSONEqf asserts that two JSON strings are equivalent.
//
//  assert.JSONEqf(t, `{"hello": "world", "foo": "bar"}`, `{"foo": "bar", "hello": "world"}`, "error message %s", "formatted")
JSONEqf(expected string, actual string, msg string, args ...interface{}) {
	debuggoGen_JSONEqf(expected, actual, msg, args)
}

func // Lenf asserts that the specified object has specific length.
// Lenf also fails if the object has a type that len() not accept.
//
//    assert.Lenf(t, mySlice, 3, "error message %s", "formatted")
Lenf(object interface{}, length int, msg string, args ...interface{}) {
	debuggoGen_Lenf(object, length, msg, args)
}

func // Lessf asserts that the first element is less than the second
//
//    assert.Lessf(t, 1, 2, "error message %s", "formatted")
//    assert.Lessf(t, float64(1), float64(2), "error message %s", "formatted")
//    assert.Lessf(t, "a", "b", "error message %s", "formatted")
Lessf(e1 interface{}, e2 interface{}, msg string, args ...interface{}) {
	debuggoGen_Lessf(e1, e2, msg, args)
}

func // LessOrEqualf asserts that the first element is less than or equal to the second
//
//    assert.LessOrEqualf(t, 1, 2, "error message %s", "formatted")
//    assert.LessOrEqualf(t, 2, 2, "error message %s", "formatted")
//    assert.LessOrEqualf(t, "a", "b", "error message %s", "formatted")
//    assert.LessOrEqualf(t, "b", "b", "error message %s", "formatted")
LessOrEqualf(e1 interface{}, e2 interface{}, msg string, args ...interface{}) {
	debuggoGen_LessOrEqualf(e1, e2, msg, args)
}

func // Neverf asserts that the given condition doesn't satisfy in waitFor time,
// periodically checking the target function each tick.
//
//    assert.Neverf(t, func() bool { return false; }, time.Second, 10*time.Millisecond, "error message %s", "formatted")
Neverf(condition func() bool, waitFor time.Duration, tick time.Duration, msg string, args ...interface{}) {
	debuggoGen_Neverf(condition, waitFor, tick, msg, args)
}

func // Nilf asserts that the specified object is nil.
//
//    assert.Nilf(t, err, "error message %s", "formatted")
Nilf(object interface{}, msg string, args ...interface{}) {
	debuggoGen_Nilf(object, msg, args)
}

func // NoDirExistsf checks whether a directory does not exist in the given path.
// It fails if the path points to an existing _directory_ only.
NoDirExistsf(path string, msg string, args ...interface{}) {
	debuggoGen_NoDirExistsf(path, msg, args)
}

func // NoErrorf asserts that a function returned no error (i.e. `nil`).
//
//   actualObj, err := SomeFunction()
//   if assert.NoErrorf(t, err, "error message %s", "formatted") {
// 	   assert.Equal(t, expectedObj, actualObj)
//   }
NoErrorf(err error, msg string, args ...interface{}) {
	debuggoGen_NoErrorf(err, msg, args)
}

func // NoFileExistsf checks whether a file does not exist in a given path. It fails
// if the path points to an existing _file_ only.
NoFileExistsf(path string, msg string, args ...interface{}) {
	debuggoGen_NoFileExistsf(path, msg, args)
}

func // NotContainsf asserts that the specified string, list(array, slice...) or map does NOT contain the
// specified substring or element.
//
//    assert.NotContainsf(t, "Hello World", "Earth", "error message %s", "formatted")
//    assert.NotContainsf(t, ["Hello", "World"], "Earth", "error message %s", "formatted")
//    assert.NotContainsf(t, {"Hello": "World"}, "Earth", "error message %s", "formatted")
NotContainsf(s interface{}, contains interface{}, msg string, args ...interface{}) {
	debuggoGen_NotContainsf(s, contains, msg, args)
}

func // NotEmptyf asserts that the specified object is NOT empty.  I.e. not nil, "", false, 0 or either
// a slice or a channel with len == 0.
//
//  if assert.NotEmptyf(t, obj, "error message %s", "formatted") {
//    assert.Equal(t, "two", obj[1])
//  }
NotEmptyf(object interface{}, msg string, args ...interface{}) {
	debuggoGen_NotEmptyf(object, msg, args)
}

func // NotEqualf asserts that the specified values are NOT equal.
//
//    assert.NotEqualf(t, obj1, obj2, "error message %s", "formatted")
//
// Pointer variable equality is determined based on the equality of the
// referenced values (as opposed to the memory addresses).
NotEqualf(expected interface{}, actual interface{}, msg string, args ...interface{}) {
	debuggoGen_NotEqualf(expected, actual, msg, args)
}

func // NotEqualValuesf asserts that two objects are not equal even when converted to the same type
//
//    assert.NotEqualValuesf(t, obj1, obj2, "error message %s", "formatted")
NotEqualValuesf(expected interface{}, actual interface{}, msg string, args ...interface{}) {
	debuggoGen_NotEqualValuesf(expected, actual, msg, args)
}

func // NotErrorIsf asserts that at none of the errors in err's chain matches target.
// This is a wrapper for errors.Is.
NotErrorIsf(err error, target error, msg string, args ...interface{}) {
	debuggoGen_NotErrorIsf(err, target, msg, args)
}

func // NotNilf asserts that the specified object is not nil.
//
//    assert.NotNilf(t, err, "error message %s", "formatted")
NotNilf(object interface{}, msg string, args ...interface{}) {
	debuggoGen_NotNilf(object, msg, args)
}

func // NotPanicsf asserts that the code inside the specified PanicTestFunc does NOT panic.
//
//   assert.NotPanicsf(t, func(){ RemainCalm() }, "error message %s", "formatted")
NotPanicsf(f PanicTestFunc, msg string, args ...interface{}) {
	debuggoGen_NotPanicsf(f, msg, args)
}

func // NotRegexpf asserts that a specified regexp does not match a string.
//
//  assert.NotRegexpf(t, regexp.MustCompile("starts"), "it's starting", "error message %s", "formatted")
//  assert.NotRegexpf(t, "^start", "it's not starting", "error message %s", "formatted")
NotRegexpf(rx interface{}, str interface{}, msg string, args ...interface{}) {
	debuggoGen_NotRegexpf(rx, str, msg, args)
}

func // NotSamef asserts that two pointers do not reference the same object.
//
//    assert.NotSamef(t, ptr1, ptr2, "error message %s", "formatted")
//
// Both arguments must be pointer variables. Pointer variable sameness is
// determined based on the equality of both type and value.
NotSamef(expected interface{}, actual interface{}, msg string, args ...interface{}) {
	debuggoGen_NotSamef(expected, actual, msg, args)
}

func // NotSubsetf asserts that the specified list(array, slice...) contains not all
// elements given in the specified subset(array, slice...).
//
//    assert.NotSubsetf(t, [1, 3, 4], [1, 2], "But [1, 3, 4] does not contain [1, 2]", "error message %s", "formatted")
NotSubsetf(list interface{}, subset interface{}, msg string, args ...interface{}) {
	debuggoGen_NotSubsetf(list, subset, msg, args)
}

func // NotZerof asserts that i is not the zero value for its type.
NotZerof(i interface{}, msg string, args ...interface{}) {
	debuggoGen_NotZerof(i, msg, args)
}

func // Panicsf asserts that the code inside the specified PanicTestFunc panics.
//
//   assert.Panicsf(t, func(){ GoCrazy() }, "error message %s", "formatted")
Panicsf(f PanicTestFunc, msg string, args ...interface{}) {
	debuggoGen_Panicsf(f, msg, args)
}

func // PanicsWithErrorf asserts that the code inside the specified PanicTestFunc
// panics, and that the recovered panic value is an error that satisfies the
// EqualError comparison.
//
//   assert.PanicsWithErrorf(t, "crazy error", func(){ GoCrazy() }, "error message %s", "formatted")
PanicsWithErrorf(errString string, f PanicTestFunc, msg string, args ...interface{}) {
	debuggoGen_PanicsWithErrorf(errString, f, msg, args)
}

func // PanicsWithValuef asserts that the code inside the specified PanicTestFunc panics, and that
// the recovered panic value equals the expected panic value.
//
//   assert.PanicsWithValuef(t, "crazy error", func(){ GoCrazy() }, "error message %s", "formatted")
PanicsWithValuef(expected interface{}, f PanicTestFunc, msg string, args ...interface{}) {
	debuggoGen_PanicsWithValuef(expected, f, msg, args)
}

func // Regexpf asserts that a specified regexp matches a string.
//
//  assert.Regexpf(t, regexp.MustCompile("start"), "it's starting", "error message %s", "formatted")
//  assert.Regexpf(t, "start...$", "it's not starting", "error message %s", "formatted")
Regexpf(rx interface{}, str interface{}, msg string, args ...interface{}) {
	debuggoGen_Regexpf(rx, str, msg, args)
}

func // Samef asserts that two pointers reference the same object.
//
//    assert.Samef(t, ptr1, ptr2, "error message %s", "formatted")
//
// Both arguments must be pointer variables. Pointer variable sameness is
// determined based on the equality of both type and value.
Samef(expected interface{}, actual interface{}, msg string, args ...interface{}) {
	debuggoGen_Samef(expected, actual, msg, args)
}

func // Subsetf asserts that the specified list(array, slice...) contains all
// elements given in the specified subset(array, slice...).
//
//    assert.Subsetf(t, [1, 2, 3], [1, 2], "But [1, 2, 3] does contain [1, 2]", "error message %s", "formatted")
Subsetf(list interface{}, subset interface{}, msg string, args ...interface{}) {
	debuggoGen_Subsetf(list, subset, msg, args)
}

func // Truef asserts that the specified value is true.
//
//    assert.Truef(t, myBool, "error message %s", "formatted")
Truef(value bool, msg string, args ...interface{}) {
	debuggoGen_Truef(value, msg, args)
}

func // WithinDurationf asserts that the two times are within duration delta of each other.
//
//   assert.WithinDurationf(t, time.Now(), time.Now(), 10*time.Second, "error message %s", "formatted")
WithinDurationf(expected time.Time, actual time.Time, delta time.Duration, msg string, args ...interface{}) {
	debuggoGen_WithinDurationf(expected, actual, delta, msg, args)
}

func // YAMLEqf asserts that two YAML strings are equivalent.
YAMLEqf(expected string, actual string, msg string, args ...interface{}) {
	debuggoGen_YAMLEqf(expected, actual, msg, args)
}

func // Zerof asserts that i is the zero value for its type.
Zerof(i interface{}, msg string, args ...interface{}) {
	debuggoGen_Zerof(i, msg, args)
}
