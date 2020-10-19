// +build !assert

package assert

import (
	"time"

	"github.com/davecgh/go-spew/spew"
)

type TestingT interface {
	Errorf( // TestingT is an interface wrapper around *testing.T
		_ string, _ ...interface{})
}

type ComparisonAssertionFunc func(TestingT, interface{}, interface{}, ...interface{})

type ValueAssertionFunc func(TestingT, interface{}, ...interface{})

type BoolAssertionFunc func(TestingT, bool, ...interface{})

type ErrorAssertionFunc func(TestingT, error, ...interface{})

type Comparison func()

func ObjectsAreEqual( // ComparisonAssertionFunc is a common function prototype when comparing two values.  Can be useful
// for table driven tests.
// ObjectsAreEqual determines if two objects are considered equal.
//
// This function does no assertion of any kind.
	_, _ interface{}) {
}

func ObjectsAreEqualValues( // ObjectsAreEqualValues gets whether two objects are equal, or if their
// values are equal.
	_, _ interface{}) {
}

func CallerInfo() { // CallerInfo returns an array of strings containing the file and line number
	// of each stack frame leading from the current test to the assert call that
	// failed.
}

type failNower interface{ FailNow() }

func FailNow( // FailNow fails test
	_ string, _ ...interface{}) {
}

func Fail( // Fail reports a failure through
	_ string, _ ...interface{}) {
}

type labeledContent struct {
	label   string
	content string
}

func Implements( // Implements asserts that an object is implemented by the specified interface.
//
//    assert.Implements(t, (*MyInterface)(nil), new(MyObject))
	_ interface{}, _ interface{}, _ ...interface{}) {
}

func IsType( // IsType asserts that the specified objects are of the same type.
	_ interface{}, _ interface{}, _ ...interface{}) {
}

func Equal( // Equal asserts that two objects are equal.
//
//    assert.Equal(t, 123, 123)
//
// Pointer variable equality is determined based on the equality of the
// referenced values (as opposed to the memory addresses). Function equality
// cannot be determined and will always fail.
	_, _ interface{}, _ ...interface{}) {
}

func Same( // Same asserts that two pointers reference the same object.
//
//    assert.Same(t, ptr1, ptr2)
//
// Both arguments must be pointer variables. Pointer variable sameness is
// determined based on the equality of both type and value.
	_, _ interface{}, _ ...interface{}) {
}

func NotSame( // NotSame asserts that two pointers do not reference the same object.
//
//    assert.NotSame(t, ptr1, ptr2)
//
// Both arguments must be pointer variables. Pointer variable sameness is
// determined based on the equality of both type and value.
	_, _ interface{}, _ ...interface{}) {
}

func EqualValues( // EqualValues asserts that two objects are equal or convertable to the same types
// and equal.
//
//    assert.EqualValues(t, uint32(123), int32(123))
	_, _ interface{}, _ ...interface{}) {
}

func Exactly( // Exactly asserts that two objects are equal in value and type.
//
//    assert.Exactly(t, int32(123), int64(123))
	_, _ interface{}, _ ...interface{}) {
}

func NotNil( // NotNil asserts that the specified object is not nil.
//
//    assert.NotNil(t, err)
	_ interface{}, _ ...interface{}) {
}

func Nil( // Nil asserts that the specified object is nil.
//
//    assert.Nil(t, err)
	_ interface{}, _ ...interface{}) {
}

func Empty( // Empty asserts that the specified object is empty.  I.e. nil, "", false, 0 or either
// a slice or a channel with len == 0.
//
//  assert.Empty(t, obj)
	_ interface{}, _ ...interface{}) {
}

func NotEmpty( // NotEmpty asserts that the specified object is NOT empty.  I.e. not nil, "", false, 0 or either
// a slice or a channel with len == 0.
//
//  if assert.NotEmpty(t, obj) {
//    assert.Equal(t, "two", obj[1])
//  }
	_ interface{}, _ ...interface{}) {
}

func Len( // Len asserts that the specified object has specific length.
// Len also fails if the object has a type that len() not accept.
//
//    assert.Len(t, mySlice, 3)
	_ interface{}, _ int, _ ...interface{}) {
}

func True( // True asserts that the specified value is true.
//
//    assert.True(t, myBool)
	_ bool, _ ...interface{}) {
}

func False( // False asserts that the specified value is false.
//
//    assert.False(t, myBool)
	_ bool, _ ...interface{}) {
}

func NotEqual( // NotEqual asserts that the specified values are NOT equal.
//
//    assert.NotEqual(t, obj1, obj2)
//
// Pointer variable equality is determined based on the equality of the
// referenced values (as opposed to the memory addresses).
	_, _ interface{}, _ ...interface{}) {
}

func NotEqualValues( // NotEqualValues asserts that two objects are not equal even when converted to the same type
//
//    assert.NotEqualValues(t, obj1, obj2)
	_, _ interface{}, _ ...interface{}) {
}

func Contains( // Contains asserts that the specified string, list(array, slice...) or map contains the
// specified substring or element.
//
//    assert.Contains(t, "Hello World", "World")
//    assert.Contains(t, ["Hello", "World"], "World")
//    assert.Contains(t, {"Hello": "World"}, "Hello")
	_, _ interface{}, _ ...interface{}) {
}

func NotContains( // NotContains asserts that the specified string, list(array, slice...) or map does NOT contain the
// specified substring or element.
//
//    assert.NotContains(t, "Hello World", "Earth")
//    assert.NotContains(t, ["Hello", "World"], "Earth")
//    assert.NotContains(t, {"Hello": "World"}, "Earth")
	_, _ interface{}, _ ...interface{}) {
}

func Subset( // Subset asserts that the specified list(array, slice...) contains all
// elements given in the specified subset(array, slice...).
//
//    assert.Subset(t, [1, 2, 3], [1, 2], "But [1, 2, 3] does contain [1, 2]")
	_, _ interface{}, _ ...interface{}) {
}

func NotSubset( // NotSubset asserts that the specified list(array, slice...) contains not all
// elements given in the specified subset(array, slice...).
//
//    assert.NotSubset(t, [1, 3, 4], [1, 2], "But [1, 3, 4] does not contain [1, 2]")
	_, _ interface{}, _ ...interface{}) {
}

func ElementsMatch( // ElementsMatch asserts that the specified listA(array, slice...) is equal to specified
// listB(array, slice...) ignoring the order of the elements. If there are duplicate elements,
// the number of appearances of each of them in both lists should match.
//
// assert.ElementsMatch(t, [1, 3, 2, 3], [1, 3, 3, 2])
	_, _ interface{}, _ ...interface{}) {
}

func Condition( // Condition uses a Comparison to assert a complex condition.
	_ Comparison, _ ...interface{}) {
}

type PanicTestFunc func()

func Panics( // PanicTestFunc defines a func that should be passed to the assert.Panics and assert.NotPanics
// methods, and represents a simple func that takes no arguments, and returns nothing.
// Panics asserts that the code inside the specified PanicTestFunc panics.
//
//   assert.Panics(t, func(){ GoCrazy() })
	_ PanicTestFunc, _ ...interface{}) {
}

func PanicsWithValue( // PanicsWithValue asserts that the code inside the specified PanicTestFunc panics, and that
// the recovered panic value equals the expected panic value.
//
//   assert.PanicsWithValue(t, "crazy error", func(){ GoCrazy() })
	_ interface{}, _ PanicTestFunc, _ ...interface{}) {
}

func PanicsWithError( // PanicsWithError asserts that the code inside the specified PanicTestFunc
// panics, and that the recovered panic value is an error that satisfies the
// EqualError comparison.
//
//   assert.PanicsWithError(t, "crazy error", func(){ GoCrazy() })
	_ string, _ PanicTestFunc, _ ...interface{}) {
}

func NotPanics( // NotPanics asserts that the code inside the specified PanicTestFunc does NOT panic.
//
//   assert.NotPanics(t, func(){ RemainCalm() })
	_ PanicTestFunc, _ ...interface{}) {
}

func WithinDuration( // WithinDuration asserts that the two times are within duration delta of each other.
//
//   assert.WithinDuration(t, time.Now(), time.Now(), 10*time.Second)
	_, _ time.Time, _ time.Duration, _ ...interface{}) {
}

func InDelta( // InDelta asserts that the two numerals are within delta of each other.
//
// 	 assert.InDelta(t, math.Pi, 22/7.0, 0.01)
	_, _ interface{}, _ float64, _ ...interface{}) {
}

func InDeltaSlice( // InDeltaSlice is the same as InDelta, except it compares two slices.
	_, _ interface{}, _ float64, _ ...interface{}) {
}

func InDeltaMapValues( // InDeltaMapValues is the same as InDelta, but it compares all values between two maps. Both maps must have exactly the same keys.
	_, _ interface{}, _ float64, _ ...interface{}) {
}

func InEpsilon( // InEpsilon asserts that expected and actual have a relative error less than epsilon
	_, _ interface{}, _ float64, _ ...interface{}) {
}

func InEpsilonSlice( // InEpsilonSlice is the same as InEpsilon, except it compares each value from two slices.
	_, _ interface{}, _ float64, _ ...interface{}) {
}

func NoError( // NoError asserts that a function returned no error (i.e. `nil`).
//
//   actualObj, err := SomeFunction()
//   if assert.NoError(t, err) {
//	   assert.Equal(t, expectedObj, actualObj)
//   }
	_ error, _ ...interface{}) {
}

func Error( // Error asserts that a function returned an error (i.e. not `nil`).
//
//   actualObj, err := SomeFunction()
//   if assert.Error(t, err) {
//	   assert.Equal(t, expectedError, err)
//   }
	_ error, _ ...interface{}) {
}

func EqualError( // EqualError asserts that a function returned an error (i.e. not `nil`)
// and that it is equal to the provided error.
//
//   actualObj, err := SomeFunction()
//   assert.EqualError(t, err,  expectedErrorString)
	_ error, _ string, _ ...interface{}) {
}

func Regexp( // Regexp asserts that a specified regexp matches a string.
//
//  assert.Regexp(t, regexp.MustCompile("start"), "it's starting")
//  assert.Regexp(t, "start...$", "it's not starting")
	_ interface{}, _ interface{}, _ ...interface{}) {
}

func NotRegexp( // NotRegexp asserts that a specified regexp does not match a string.
//
//  assert.NotRegexp(t, regexp.MustCompile("starts"), "it's starting")
//  assert.NotRegexp(t, "^start", "it's not starting")
	_ interface{}, _ interface{}, _ ...interface{}) {
}

func Zero( // Zero asserts that i is the zero value for its type.
	_ interface{}, _ ...interface{}) {
}

func NotZero( // NotZero asserts that i is not the zero value for its type.
	_ interface{}, _ ...interface{}) {
}

func FileExists( // FileExists checks whether a file exists in the given path. It also fails if
// the path points to a directory or there is an error when trying to check the file.
	_ string, _ ...interface{}) {
}

func NoFileExists( // NoFileExists checks whether a file does not exist in a given path. It fails
// if the path points to an existing _file_ only.
	_ string, _ ...interface{}) {
}

func DirExists( // DirExists checks whether a directory exists in the given path. It also fails
// if the path is a file rather a directory or there is an error checking whether it exists.
	_ string, _ ...interface{}) {
}

func NoDirExists( // NoDirExists checks whether a directory does not exist in the given path.
// It fails if the path points to an existing _directory_ only.
	_ string, _ ...interface{}) {
}

func JSONEq( // JSONEq asserts that two JSON strings are equivalent.
//
//  assert.JSONEq(t, `{"hello": "world", "foo": "bar"}`, `{"foo": "bar", "hello": "world"}`)
	_ string, _ string, _ ...interface{}) {
}

func YAMLEq( // YAMLEq asserts that two YAML strings are equivalent.
	_ string, _ string, _ ...interface{}) {
}

var spewConfig = spew.ConfigState{Indent: " ", DisablePointerAddresses: true, DisableCapacities: true, SortKeys: true, DisableMethods: true}

type tHelper interface{ Helper() }

func Eventually( // Eventually asserts that given condition will be met in waitFor time,
// periodically checking target function each tick.
//
//    assert.Eventually(t, func() bool { return true; }, time.Second, 10*time.Millisecond)
	_ func(), _ time.Duration, _ time.Duration, _ ...interface{}) {
}

func Never( // Never asserts that the given condition doesn't satisfy in waitFor time,
// periodically checking the target function each tick.
//
//    assert.Never(t, func() bool { return false; }, time.Second, 10*time.Millisecond)
	_ func(), _ time.Duration, _ time.Duration, _ ...interface{}) {
}

func ErrorIs( // ErrorIs asserts that at least one of the errors in err's chain matches target.
// This is a wrapper for errors.Is.
	_, _ error, _ ...interface{}) {
}

func NotErrorIs( // NotErrorIs asserts that at none of the errors in err's chain matches target.
// This is a wrapper for errors.Is.
	_, _ error, _ ...interface{}) {
}

func ErrorAs( // ErrorAs asserts that at least one of the errors in err's chain matches target, and if so, sets target to that error value.
// This is a wrapper for errors.As.
	_ error, _ interface{}, _ ...interface{}) {
}
