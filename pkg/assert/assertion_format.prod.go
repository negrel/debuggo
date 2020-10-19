// +build !assert

package assert

import (
	http "net/http"
	url "net/url"
	time "time"
)

func Conditionf( // Conditionf uses a Comparison to assert a complex condition.
_ Comparison, _ string, _ ...interface{}) {
}

func Containsf( // Containsf asserts that the specified string, list(array, slice...) or map contains the
// specified substring or element.
//
//    assert.Containsf(t, "Hello World", "World", "error message %s", "formatted")
//    assert.Containsf(t, ["Hello", "World"], "World", "error message %s", "formatted")
//    assert.Containsf(t, {"Hello": "World"}, "Hello", "error message %s", "formatted")
_ interface{}, _ interface{}, _ string, _ ...interface{}) {
}

func DirExistsf( // DirExistsf checks whether a directory exists in the given path. It also fails
// if the path is a file rather a directory or there is an error checking whether it exists.
_ string, _ string, _ ...interface{}) {
}

func ElementsMatchf( // ElementsMatchf asserts that the specified listA(array, slice...) is equal to specified
// listB(array, slice...) ignoring the order of the elements. If there are duplicate elements,
// the number of appearances of each of them in both lists should match.
//
// assert.ElementsMatchf(t, [1, 3, 2, 3], [1, 3, 3, 2], "error message %s", "formatted")
_ interface{}, _ interface{}, _ string, _ ...interface{}) {
}

func Emptyf( // Emptyf asserts that the specified object is empty.  I.e. nil, "", false, 0 or either
// a slice or a channel with len == 0.
//
//  assert.Emptyf(t, obj, "error message %s", "formatted")
_ interface{}, _ string, _ ...interface{}) {
}

func Equalf( // Equalf asserts that two objects are equal.
//
//    assert.Equalf(t, 123, 123, "error message %s", "formatted")
//
// Pointer variable equality is determined based on the equality of the
// referenced values (as opposed to the memory addresses). Function equality
// cannot be determined and will always fail.
_ interface{}, _ interface{}, _ string, _ ...interface{}) {
}

func EqualErrorf( // EqualErrorf asserts that a function returned an error (i.e. not `nil`)
// and that it is equal to the provided error.
//
//   actualObj, err := SomeFunction()
//   assert.EqualErrorf(t, err,  expectedErrorString, "error message %s", "formatted")
_ error, _ string, _ string, _ ...interface{}) {
}

func EqualValuesf( // EqualValuesf asserts that two objects are equal or convertable to the same types
// and equal.
//
//    assert.EqualValuesf(t, uint32(123), int32(123), "error message %s", "formatted")
_ interface{}, _ interface{}, _ string, _ ...interface{}) {
}

func Errorf( // Errorf asserts that a function returned an error (i.e. not `nil`).
//
//   actualObj, err := SomeFunction()
//   if assert.Errorf(t, err, "error message %s", "formatted") {
// 	   assert.Equal(t, expectedErrorf, err)
//   }
_ error, _ string, _ ...interface{}) {
}

func ErrorAsf( // ErrorAsf asserts that at least one of the errors in err's chain matches target, and if so, sets target to that error value.
// This is a wrapper for errors.As.
_ error, _ interface{}, _ string, _ ...interface{}) {
}

func ErrorIsf( // ErrorIsf asserts that at least one of the errors in err's chain matches target.
// This is a wrapper for errors.Is.
_ error, _ error, _ string, _ ...interface{}) {
}

func Eventuallyf( // Eventuallyf asserts that given condition will be met in waitFor time,
// periodically checking target function each tick.
//
//    assert.Eventuallyf(t, func() bool { return true; }, time.Second, 10*time.Millisecond, "error message %s", "formatted")
_ func(), _ time.Duration, _ time.Duration, _ string, _ ...interface{}) {
}

func Exactlyf( // Exactlyf asserts that two objects are equal in value and type.
//
//    assert.Exactlyf(t, int32(123), int64(123), "error message %s", "formatted")
_ interface{}, _ interface{}, _ string, _ ...interface{}) {
}

func Failf( // Failf reports a failure through
_ string, _ string, _ ...interface{}) {
}

func FailNowf( // FailNowf fails test
_ string, _ string, _ ...interface{}) {
}

func Falsef( // Falsef asserts that the specified value is false.
//
//    assert.Falsef(t, myBool, "error message %s", "formatted")
_ bool, _ string, _ ...interface{}) {
}

func FileExistsf( // FileExistsf checks whether a file exists in the given path. It also fails if
// the path points to a directory or there is an error when trying to check the file.
_ string, _ string, _ ...interface{}) {
}

func Greaterf( // Greaterf asserts that the first element is greater than the second
//
//    assert.Greaterf(t, 2, 1, "error message %s", "formatted")
//    assert.Greaterf(t, float64(2), float64(1), "error message %s", "formatted")
//    assert.Greaterf(t, "b", "a", "error message %s", "formatted")
_ interface{}, _ interface{}, _ string, _ ...interface{}) {
}

func GreaterOrEqualf( // GreaterOrEqualf asserts that the first element is greater than or equal to the second
//
//    assert.GreaterOrEqualf(t, 2, 1, "error message %s", "formatted")
//    assert.GreaterOrEqualf(t, 2, 2, "error message %s", "formatted")
//    assert.GreaterOrEqualf(t, "b", "a", "error message %s", "formatted")
//    assert.GreaterOrEqualf(t, "b", "b", "error message %s", "formatted")
_ interface{}, _ interface{}, _ string, _ ...interface{}) {
}

func HTTPBodyContainsf( // HTTPBodyContainsf asserts that a specified handler returns a
// body that contains a string.
//
//  assert.HTTPBodyContainsf(t, myHandler, "GET", "www.google.com", nil, "I'm Feeling Lucky", "error message %s", "formatted")
//
// Returns whether the assertion was successful (true) or not (false).
_ http.HandlerFunc, _ string, _ string, _ url.Values, _ interface{}, _ string, _ ...interface{}) {
}

func HTTPBodyNotContainsf( // HTTPBodyNotContainsf asserts that a specified handler returns a
// body that does not contain a string.
//
//  assert.HTTPBodyNotContainsf(t, myHandler, "GET", "www.google.com", nil, "I'm Feeling Lucky", "error message %s", "formatted")
//
// Returns whether the assertion was successful (true) or not (false).
_ http.HandlerFunc, _ string, _ string, _ url.Values, _ interface{}, _ string, _ ...interface{}) {
}

func HTTPErrorf( // HTTPErrorf asserts that a specified handler returns an error status code.
//
//  assert.HTTPErrorf(t, myHandler, "POST", "/a/b/c", url.Values{"a": []string{"b", "c"}}
//
// Returns whether the assertion was successful (true) or not (false).
_ http.HandlerFunc, _ string, _ string, _ url.Values, _ string, _ ...interface{}) {
}

func HTTPRedirectf( // HTTPRedirectf asserts that a specified handler returns a redirect status code.
//
//  assert.HTTPRedirectf(t, myHandler, "GET", "/a/b/c", url.Values{"a": []string{"b", "c"}}
//
// Returns whether the assertion was successful (true) or not (false).
_ http.HandlerFunc, _ string, _ string, _ url.Values, _ string, _ ...interface{}) {
}

func HTTPStatusCodef( // HTTPStatusCodef asserts that a specified handler returns a specified status code.
//
//  assert.HTTPStatusCodef(t, myHandler, "GET", "/notImplemented", nil, 501, "error message %s", "formatted")
//
// Returns whether the assertion was successful (true) or not (false).
_ http.HandlerFunc, _ string, _ string, _ url.Values, _ int, _ string, _ ...interface{}) {
}

func HTTPSuccessf( // HTTPSuccessf asserts that a specified handler returns a success status code.
//
//  assert.HTTPSuccessf(t, myHandler, "POST", "http://www.google.com", nil, "error message %s", "formatted")
//
// Returns whether the assertion was successful (true) or not (false).
_ http.HandlerFunc, _ string, _ string, _ url.Values, _ string, _ ...interface{}) {
}

func Implementsf( // Implementsf asserts that an object is implemented by the specified interface.
//
//    assert.Implementsf(t, (*MyInterface)(nil), new(MyObject), "error message %s", "formatted")
_ interface{}, _ interface{}, _ string, _ ...interface{}) {
}

func InDeltaf( // InDeltaf asserts that the two numerals are within delta of each other.
//
// 	 assert.InDeltaf(t, math.Pi, 22/7.0, 0.01, "error message %s", "formatted")
_ interface{}, _ interface{}, _ float64, _ string, _ ...interface{}) {
}

func InDeltaMapValuesf( // InDeltaMapValuesf is the same as InDelta, but it compares all values between two maps. Both maps must have exactly the same keys.
_ interface{}, _ interface{}, _ float64, _ string, _ ...interface{}) {
}

func InDeltaSlicef( // InDeltaSlicef is the same as InDelta, except it compares two slices.
_ interface{}, _ interface{}, _ float64, _ string, _ ...interface{}) {
}

func InEpsilonf( // InEpsilonf asserts that expected and actual have a relative error less than epsilon
_ interface{}, _ interface{}, _ float64, _ string, _ ...interface{}) {
}

func InEpsilonSlicef( // InEpsilonSlicef is the same as InEpsilon, except it compares each value from two slices.
_ interface{}, _ interface{}, _ float64, _ string, _ ...interface{}) {
}

func IsDecreasingf( // IsDecreasingf asserts that the collection is decreasing
//
//    assert.IsDecreasingf(t, []int{2, 1, 0}, "error message %s", "formatted")
//    assert.IsDecreasingf(t, []float{2, 1}, "error message %s", "formatted")
//    assert.IsDecreasingf(t, []string{"b", "a"}, "error message %s", "formatted")
_ interface{}, _ string, _ ...interface{}) {
}

func IsIncreasingf( // IsIncreasingf asserts that the collection is increasing
//
//    assert.IsIncreasingf(t, []int{1, 2, 3}, "error message %s", "formatted")
//    assert.IsIncreasingf(t, []float{1, 2}, "error message %s", "formatted")
//    assert.IsIncreasingf(t, []string{"a", "b"}, "error message %s", "formatted")
_ interface{}, _ string, _ ...interface{}) {
}

func IsNonDecreasingf( // IsNonDecreasingf asserts that the collection is not decreasing
//
//    assert.IsNonDecreasingf(t, []int{1, 1, 2}, "error message %s", "formatted")
//    assert.IsNonDecreasingf(t, []float{1, 2}, "error message %s", "formatted")
//    assert.IsNonDecreasingf(t, []string{"a", "b"}, "error message %s", "formatted")
_ interface{}, _ string, _ ...interface{}) {
}

func IsNonIncreasingf( // IsNonIncreasingf asserts that the collection is not increasing
//
//    assert.IsNonIncreasingf(t, []int{2, 1, 1}, "error message %s", "formatted")
//    assert.IsNonIncreasingf(t, []float{2, 1}, "error message %s", "formatted")
//    assert.IsNonIncreasingf(t, []string{"b", "a"}, "error message %s", "formatted")
_ interface{}, _ string, _ ...interface{}) {
}

func IsTypef( // IsTypef asserts that the specified objects are of the same type.
_ interface{}, _ interface{}, _ string, _ ...interface{}) {
}

func JSONEqf( // JSONEqf asserts that two JSON strings are equivalent.
//
//  assert.JSONEqf(t, `{"hello": "world", "foo": "bar"}`, `{"foo": "bar", "hello": "world"}`, "error message %s", "formatted")
_ string, _ string, _ string, _ ...interface{}) {
}

func Lenf( // Lenf asserts that the specified object has specific length.
// Lenf also fails if the object has a type that len() not accept.
//
//    assert.Lenf(t, mySlice, 3, "error message %s", "formatted")
_ interface{}, _ int, _ string, _ ...interface{}) {
}

func Lessf( // Lessf asserts that the first element is less than the second
//
//    assert.Lessf(t, 1, 2, "error message %s", "formatted")
//    assert.Lessf(t, float64(1), float64(2), "error message %s", "formatted")
//    assert.Lessf(t, "a", "b", "error message %s", "formatted")
_ interface{}, _ interface{}, _ string, _ ...interface{}) {
}

func LessOrEqualf( // LessOrEqualf asserts that the first element is less than or equal to the second
//
//    assert.LessOrEqualf(t, 1, 2, "error message %s", "formatted")
//    assert.LessOrEqualf(t, 2, 2, "error message %s", "formatted")
//    assert.LessOrEqualf(t, "a", "b", "error message %s", "formatted")
//    assert.LessOrEqualf(t, "b", "b", "error message %s", "formatted")
_ interface{}, _ interface{}, _ string, _ ...interface{}) {
}

func Neverf( // Neverf asserts that the given condition doesn't satisfy in waitFor time,
// periodically checking the target function each tick.
//
//    assert.Neverf(t, func() bool { return false; }, time.Second, 10*time.Millisecond, "error message %s", "formatted")
_ func(), _ time.Duration, _ time.Duration, _ string, _ ...interface{}) {
}

func Nilf( // Nilf asserts that the specified object is nil.
//
//    assert.Nilf(t, err, "error message %s", "formatted")
_ interface{}, _ string, _ ...interface{}) {
}

func NoDirExistsf( // NoDirExistsf checks whether a directory does not exist in the given path.
// It fails if the path points to an existing _directory_ only.
_ string, _ string, _ ...interface{}) {
}

func NoErrorf( // NoErrorf asserts that a function returned no error (i.e. `nil`).
//
//   actualObj, err := SomeFunction()
//   if assert.NoErrorf(t, err, "error message %s", "formatted") {
// 	   assert.Equal(t, expectedObj, actualObj)
//   }
_ error, _ string, _ ...interface{}) {
}

func NoFileExistsf( // NoFileExistsf checks whether a file does not exist in a given path. It fails
// if the path points to an existing _file_ only.
_ string, _ string, _ ...interface{}) {
}

func NotContainsf( // NotContainsf asserts that the specified string, list(array, slice...) or map does NOT contain the
// specified substring or element.
//
//    assert.NotContainsf(t, "Hello World", "Earth", "error message %s", "formatted")
//    assert.NotContainsf(t, ["Hello", "World"], "Earth", "error message %s", "formatted")
//    assert.NotContainsf(t, {"Hello": "World"}, "Earth", "error message %s", "formatted")
_ interface{}, _ interface{}, _ string, _ ...interface{}) {
}

func NotEmptyf( // NotEmptyf asserts that the specified object is NOT empty.  I.e. not nil, "", false, 0 or either
// a slice or a channel with len == 0.
//
//  if assert.NotEmptyf(t, obj, "error message %s", "formatted") {
//    assert.Equal(t, "two", obj[1])
//  }
_ interface{}, _ string, _ ...interface{}) {
}

func NotEqualf( // NotEqualf asserts that the specified values are NOT equal.
//
//    assert.NotEqualf(t, obj1, obj2, "error message %s", "formatted")
//
// Pointer variable equality is determined based on the equality of the
// referenced values (as opposed to the memory addresses).
_ interface{}, _ interface{}, _ string, _ ...interface{}) {
}

func NotEqualValuesf( // NotEqualValuesf asserts that two objects are not equal even when converted to the same type
//
//    assert.NotEqualValuesf(t, obj1, obj2, "error message %s", "formatted")
_ interface{}, _ interface{}, _ string, _ ...interface{}) {
}

func NotErrorIsf( // NotErrorIsf asserts that at none of the errors in err's chain matches target.
// This is a wrapper for errors.Is.
_ error, _ error, _ string, _ ...interface{}) {
}

func NotNilf( // NotNilf asserts that the specified object is not nil.
//
//    assert.NotNilf(t, err, "error message %s", "formatted")
_ interface{}, _ string, _ ...interface{}) {
}

func NotPanicsf( // NotPanicsf asserts that the code inside the specified PanicTestFunc does NOT panic.
//
//   assert.NotPanicsf(t, func(){ RemainCalm() }, "error message %s", "formatted")
_ PanicTestFunc, _ string, _ ...interface{}) {
}

func NotRegexpf( // NotRegexpf asserts that a specified regexp does not match a string.
//
//  assert.NotRegexpf(t, regexp.MustCompile("starts"), "it's starting", "error message %s", "formatted")
//  assert.NotRegexpf(t, "^start", "it's not starting", "error message %s", "formatted")
_ interface{}, _ interface{}, _ string, _ ...interface{}) {
}

func NotSamef( // NotSamef asserts that two pointers do not reference the same object.
//
//    assert.NotSamef(t, ptr1, ptr2, "error message %s", "formatted")
//
// Both arguments must be pointer variables. Pointer variable sameness is
// determined based on the equality of both type and value.
_ interface{}, _ interface{}, _ string, _ ...interface{}) {
}

func NotSubsetf( // NotSubsetf asserts that the specified list(array, slice...) contains not all
// elements given in the specified subset(array, slice...).
//
//    assert.NotSubsetf(t, [1, 3, 4], [1, 2], "But [1, 3, 4] does not contain [1, 2]", "error message %s", "formatted")
_ interface{}, _ interface{}, _ string, _ ...interface{}) {
}

func NotZerof( // NotZerof asserts that i is not the zero value for its type.
_ interface{}, _ string, _ ...interface{}) {
}

func Panicsf( // Panicsf asserts that the code inside the specified PanicTestFunc panics.
//
//   assert.Panicsf(t, func(){ GoCrazy() }, "error message %s", "formatted")
_ PanicTestFunc, _ string, _ ...interface{}) {
}

func PanicsWithErrorf( // PanicsWithErrorf asserts that the code inside the specified PanicTestFunc
// panics, and that the recovered panic value is an error that satisfies the
// EqualError comparison.
//
//   assert.PanicsWithErrorf(t, "crazy error", func(){ GoCrazy() }, "error message %s", "formatted")
_ string, _ PanicTestFunc, _ string, _ ...interface{}) {
}

func PanicsWithValuef( // PanicsWithValuef asserts that the code inside the specified PanicTestFunc panics, and that
// the recovered panic value equals the expected panic value.
//
//   assert.PanicsWithValuef(t, "crazy error", func(){ GoCrazy() }, "error message %s", "formatted")
_ interface{}, _ PanicTestFunc, _ string, _ ...interface{}) {
}

func Regexpf( // Regexpf asserts that a specified regexp matches a string.
//
//  assert.Regexpf(t, regexp.MustCompile("start"), "it's starting", "error message %s", "formatted")
//  assert.Regexpf(t, "start...$", "it's not starting", "error message %s", "formatted")
_ interface{}, _ interface{}, _ string, _ ...interface{}) {
}

func Samef( // Samef asserts that two pointers reference the same object.
//
//    assert.Samef(t, ptr1, ptr2, "error message %s", "formatted")
//
// Both arguments must be pointer variables. Pointer variable sameness is
// determined based on the equality of both type and value.
_ interface{}, _ interface{}, _ string, _ ...interface{}) {
}

func Subsetf( // Subsetf asserts that the specified list(array, slice...) contains all
// elements given in the specified subset(array, slice...).
//
//    assert.Subsetf(t, [1, 2, 3], [1, 2], "But [1, 2, 3] does contain [1, 2]", "error message %s", "formatted")
_ interface{}, _ interface{}, _ string, _ ...interface{}) {
}

func Truef( // Truef asserts that the specified value is true.
//
//    assert.Truef(t, myBool, "error message %s", "formatted")
_ bool, _ string, _ ...interface{}) {
}

func WithinDurationf( // WithinDurationf asserts that the two times are within duration delta of each other.
//
//   assert.WithinDurationf(t, time.Now(), time.Now(), 10*time.Second, "error message %s", "formatted")
_ time.Time, _ time.Time, _ time.Duration, _ string, _ ...interface{}) {
}

func YAMLEqf( // YAMLEqf asserts that two YAML strings are equivalent.
_ string, _ string, _ string, _ ...interface{}) {
}

func Zerof( // Zerof asserts that i is the zero value for its type.
_ interface{}, _ string, _ ...interface{}) {
}
