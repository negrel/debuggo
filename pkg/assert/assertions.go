// +build assert

package assert

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"os"
	"reflect"
	"regexp"
	"runtime"
	"runtime/debug"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"

	"github.com/davecgh/go-spew/spew"
	"github.com/pmezard/go-difflib/difflib"
	yaml "gopkg.in/yaml.v3"
)

//go:generate sh -c "cd ../_codegen && go build && cd - && ../_codegen/_codegen -output-package=assert -template=assertion_format.go.tmpl"

// TestingT is an interface wrapper around *testing.T
type TestingT interface {
	Errorf(format string, args ...interface{})
}

// ComparisonAssertionFunc is a common function prototype when comparing two values.  Can be useful
// for table driven tests.
type ComparisonAssertionFunc func(TestingT, interface{}, interface{}, ...interface{}) bool

// ValueAssertionFunc is a common function prototype when validating a single value.  Can be useful
// for table driven tests.
type ValueAssertionFunc func(TestingT, interface{}, ...interface{}) bool

// BoolAssertionFunc is a common function prototype when validating a bool value.  Can be useful
// for table driven tests.
type BoolAssertionFunc func(TestingT, bool, ...interface{}) bool

// ErrorAssertionFunc is a common function prototype when validating an error value.  Can be useful
// for table driven tests.
type ErrorAssertionFunc func(TestingT, error, ...interface{}) bool

// Comparison is a custom function that returns true on success and false on failure
type Comparison func() (success bool)

/*
	Helper functions
*/

// ObjectsAreEqual determines if two objects are considered equal.
//
// This function does no assertion of any kind.
func ObjectsAreEqual(expected, actual interface{}) { debuggoGen_ObjectsAreEqual(expected, actual) }

// ObjectsAreEqualValues gets whether two objects are equal, or if their
// values are equal.
func ObjectsAreEqualValues(expected, actual interface{}) {
	debuggoGen_ObjectsAreEqualValues(expected, actual)
}

// Attempt comparison after type conversion

/* CallerInfo is necessary because the assert functions use the testing object
internally, causing it to print the file:line of the assert method, rather than where
the problem actually occurred in calling code.*/

// CallerInfo returns an array of strings containing the file and line number
// of each stack frame leading from the current test to the assert call that
// failed.
func CallerInfo() { debuggoGen_CallerInfo() }

// The breaks below failed to terminate the loop, and we ran off the
// end of the call stack.

// This is a huge edge case, but it will panic if this is the case, see #180

// testing.tRunner is the standard library function that calls
// tests. Subtests are called directly by tRunner, without going through
// the Test/Benchmark/Example function that contains the t.Run calls, so
// with subtests we should break when we hit tRunner, without adding it
// to the list of callers.

// Drop the package

// Stolen from the `go test` tool.
// isTest tells whether name looks like a test (or benchmark, according to prefix).
// It is a Test (say) if there is a character after Test that is not a lower-case letter.
// We don't want TesticularCancer.
func isTest(name, prefix string) bool {
	if !strings.HasPrefix(name, prefix) {
		return false
	}
	if len(name) == len(prefix) { // "Test" is ok
		return true
	}
	r, _ := utf8.DecodeRuneInString(name[len(prefix):])
	return !unicode.IsLower(r)
}

func messageFromMsgAndArgs(msgAndArgs ...interface{}) string {
	if len(msgAndArgs) == 0 || msgAndArgs == nil {
		return ""
	}
	if len(msgAndArgs) == 1 {
		msg := msgAndArgs[0]
		if msgAsStr, ok := msg.(string); ok {
			return msgAsStr
		}
		return fmt.Sprintf("%+v", msg)
	}
	if len(msgAndArgs) > 1 {
		return fmt.Sprintf(msgAndArgs[0].(string), msgAndArgs[1:]...)
	}
	return ""
}

// Aligns the provided message so that all lines after the first line start at the same location as the first line.
// Assumes that the first line starts at the correct location (after carriage return, tab, label, spacer and tab).
// The longestLabelLen parameter specifies the length of the longest label in the output (required becaues this is the
// basis on which the alignment occurs).
func indentMessageLines(message string, longestLabelLen int) string {
	outBuf := new(bytes.Buffer)

	for i, scanner := 0, bufio.NewScanner(strings.NewReader(message)); scanner.Scan(); i++ {
		// no need to align first line because it starts at the correct location (after the label)
		if i != 0 {
			// append alignLen+1 spaces to align with "{{longestLabel}}:" before adding tab
			outBuf.WriteString("\n\t" + strings.Repeat(" ", longestLabelLen+1) + "\t")
		}
		outBuf.WriteString(scanner.Text())
	}

	return outBuf.String()
}

type failNower interface {
	FailNow()
}

// FailNow fails test
func FailNow(failureMessage string, msgAndArgs ...interface{}) {
	debuggoGen_FailNow(failureMessage, msgAndArgs)
}

// We cannot extend TestingT with FailNow() and
// maintain backwards compatibility, so we fallback
// to panicking when FailNow is not available in
// TestingT.
// See issue #263

// Fail reports a failure through
func Fail(failureMessage string, msgAndArgs ...interface{}) {
	debuggoGen_Fail(failureMessage, msgAndArgs)
}

// Add test name if the Go version supports it

type labeledContent struct {
	label   string
	content string
}

// labeledOutput returns a string consisting of the provided labeledContent. Each labeled output is appended in the following manner:
//
//   \t{{label}}:{{align_spaces}}\t{{content}}\n
//
// The initial carriage return is required to undo/erase any padding added by testing.T.Errorf. The "\t{{label}}:" is for the label.
// If a label is shorter than the longest label provided, padding spaces are added to make all the labels match in length. Once this
// alignment is achieved, "\t{{content}}\n" is added for the output.
//
// If the content of the labeledOutput contains line breaks, the subsequent lines are aligned so that they start at the same location as the first line.
func labeledOutput(content ...labeledContent) string {
	longestLabel := 0
	for _, v := range content {
		if len(v.label) > longestLabel {
			longestLabel = len(v.label)
		}
	}
	var output string
	for _, v := range content {
		output += "\t" + v.label + ":" + strings.Repeat(" ", longestLabel-len(v.label)) + "\t" + indentMessageLines(v.content, longestLabel) + "\n"
	}
	return output
}

// Implements asserts that an object is implemented by the specified interface.
//
//    assert.Implements(t, (*MyInterface)(nil), new(MyObject))
func Implements(interfaceObject interface{}, object interface{}, msgAndArgs ...interface{}) {
	debuggoGen_Implements(interfaceObject, object, msgAndArgs)
}

// IsType asserts that the specified objects are of the same type.
func IsType(expectedType interface{}, object interface{}, msgAndArgs ...interface{}) {
	debuggoGen_IsType(expectedType, object, msgAndArgs)
}

// Equal asserts that two objects are equal.
//
//    assert.Equal(t, 123, 123)
//
// Pointer variable equality is determined based on the equality of the
// referenced values (as opposed to the memory addresses). Function equality
// cannot be determined and will always fail.
func Equal(expected, actual interface{}, msgAndArgs ...interface{}) {
	debuggoGen_Equal(expected, actual, msgAndArgs)
}

// validateEqualArgs checks whether provided arguments can be safely used in the
// Equal/NotEqual functions.
func validateEqualArgs(expected, actual interface{}) error {
	if expected == nil && actual == nil {
		return nil
	}

	if isFunction(expected) || isFunction(actual) {
		return errors.New("cannot take func type as argument")
	}
	return nil
}

// Same asserts that two pointers reference the same object.
//
//    assert.Same(t, ptr1, ptr2)
//
// Both arguments must be pointer variables. Pointer variable sameness is
// determined based on the equality of both type and value.
func Same(expected, actual interface{}, msgAndArgs ...interface{}) {
	debuggoGen_Same(expected, actual, msgAndArgs)
}

// NotSame asserts that two pointers do not reference the same object.
//
//    assert.NotSame(t, ptr1, ptr2)
//
// Both arguments must be pointer variables. Pointer variable sameness is
// determined based on the equality of both type and value.
func NotSame(expected, actual interface{}, msgAndArgs ...interface{}) {
	debuggoGen_NotSame(expected, actual, msgAndArgs)
}

// samePointers compares two generic interface objects and returns whether
// they point to the same object
func samePointers(first, second interface{}) bool {
	firstPtr, secondPtr := reflect.ValueOf(first), reflect.ValueOf(second)
	if firstPtr.Kind() != reflect.Ptr || secondPtr.Kind() != reflect.Ptr {
		return false
	}

	firstType, secondType := reflect.TypeOf(first), reflect.TypeOf(second)
	if firstType != secondType {
		return false
	}

	// compare pointer addresses
	return first == second
}

// formatUnequalValues takes two values of arbitrary types and returns string
// representations appropriate to be presented to the user.
//
// If the values are not of like type, the returned strings will be prefixed
// with the type name, and the value will be enclosed in parenthesis similar
// to a type conversion in the Go grammar.
func formatUnequalValues(expected, actual interface{}) (e string, a string) {
	if reflect.TypeOf(expected) != reflect.TypeOf(actual) {
		return fmt.Sprintf("%T(%s)", expected, truncatingFormat(expected)),
			fmt.Sprintf("%T(%s)", actual, truncatingFormat(actual))
	}
	switch expected.(type) {
	case time.Duration:
		return fmt.Sprintf("%v", expected), fmt.Sprintf("%v", actual)
	}
	return truncatingFormat(expected), truncatingFormat(actual)
}

// truncatingFormat formats the data and truncates it if it's too long.
//
// This helps keep formatted error messages lines from exceeding the
// bufio.MaxScanTokenSize max line length that the go testing framework imposes.
func truncatingFormat(data interface{}) string {
	value := fmt.Sprintf("%#v", data)
	max := bufio.MaxScanTokenSize - 100 // Give us some space the type info too if needed.
	if len(value) > max {
		value = value[0:max] + "<... truncated>"
	}
	return value
}

// EqualValues asserts that two objects are equal or convertable to the same types
// and equal.
//
//    assert.EqualValues(t, uint32(123), int32(123))
func EqualValues(expected, actual interface{}, msgAndArgs ...interface{}) {
	debuggoGen_EqualValues(expected, actual, msgAndArgs)
}

// Exactly asserts that two objects are equal in value and type.
//
//    assert.Exactly(t, int32(123), int64(123))
func Exactly(expected, actual interface{}, msgAndArgs ...interface{}) {
	debuggoGen_Exactly(expected, actual, msgAndArgs)
}

// NotNil asserts that the specified object is not nil.
//
//    assert.NotNil(t, err)
func NotNil(object interface{}, msgAndArgs ...interface{}) { debuggoGen_NotNil(object, msgAndArgs) }

// containsKind checks if a specified kind in the slice of kinds.
func containsKind(kinds []reflect.Kind, kind reflect.Kind) bool {
	for i := 0; i < len(kinds); i++ {
		if kind == kinds[i] {
			return true
		}
	}

	return false
}

// isNil checks if a specified object is nil or not, without Failing.
func isNil(object interface{}) bool {
	if object == nil {
		return true
	}

	value := reflect.ValueOf(object)
	kind := value.Kind()
	isNilableKind := containsKind(
		[]reflect.Kind{
			reflect.Chan, reflect.Func,
			reflect.Interface, reflect.Map,
			reflect.Ptr, reflect.Slice},
		kind)

	if isNilableKind && value.IsNil() {
		return true
	}

	return false
}

// Nil asserts that the specified object is nil.
//
//    assert.Nil(t, err)
func Nil(object interface{}, msgAndArgs ...interface{}) { debuggoGen_Nil(object, msgAndArgs) }

// isEmpty gets whether the specified object is considered empty or not.
func isEmpty(object interface{}) bool {

	// get nil case out of the way
	if object == nil {
		return true
	}

	objValue := reflect.ValueOf(object)

	switch objValue.Kind() {
	// collection types are empty when they have no element
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice:
		return objValue.Len() == 0
		// pointers are empty if nil or if the value they point to is empty
	case reflect.Ptr:
		if objValue.IsNil() {
			return true
		}
		deref := objValue.Elem().Interface()
		return isEmpty(deref)
		// for all other types, compare against the zero value
	default:
		zero := reflect.Zero(objValue.Type())
		return reflect.DeepEqual(object, zero.Interface())
	}
}

// Empty asserts that the specified object is empty.  I.e. nil, "", false, 0 or either
// a slice or a channel with len == 0.
//
//  assert.Empty(t, obj)
func Empty(object interface{}, msgAndArgs ...interface{}) { debuggoGen_Empty(object, msgAndArgs) }

// NotEmpty asserts that the specified object is NOT empty.  I.e. not nil, "", false, 0 or either
// a slice or a channel with len == 0.
//
//  if assert.NotEmpty(t, obj) {
//    assert.Equal(t, "two", obj[1])
//  }
func NotEmpty(object interface{}, msgAndArgs ...interface{}) { debuggoGen_NotEmpty(object, msgAndArgs) }

// getLen try to get length of object.
// return (false, 0) if impossible.
func getLen(x interface{}) (ok bool, length int) {
	v := reflect.ValueOf(x)
	defer func() {
		if e := recover(); e != nil {
			ok = false
		}
	}()
	return true, v.Len()
}

// Len asserts that the specified object has specific length.
// Len also fails if the object has a type that len() not accept.
//
//    assert.Len(t, mySlice, 3)
func Len(object interface{}, length int, msgAndArgs ...interface{}) {
	debuggoGen_Len(object, length, msgAndArgs)
}

// True asserts that the specified value is true.
//
//    assert.True(t, myBool)
func True(value bool, msgAndArgs ...interface{}) { debuggoGen_True(value, msgAndArgs) }

// False asserts that the specified value is false.
//
//    assert.False(t, myBool)
func False(value bool, msgAndArgs ...interface{}) { debuggoGen_False(value, msgAndArgs) }

// NotEqual asserts that the specified values are NOT equal.
//
//    assert.NotEqual(t, obj1, obj2)
//
// Pointer variable equality is determined based on the equality of the
// referenced values (as opposed to the memory addresses).
func NotEqual(expected, actual interface{}, msgAndArgs ...interface{}) {
	debuggoGen_NotEqual(expected, actual, msgAndArgs)
}

// NotEqualValues asserts that two objects are not equal even when converted to the same type
//
//    assert.NotEqualValues(t, obj1, obj2)
func NotEqualValues(expected, actual interface{}, msgAndArgs ...interface{}) {
	debuggoGen_NotEqualValues(expected, actual, msgAndArgs)
}

// containsElement try loop over the list check if the list includes the element.
// return (false, false) if impossible.
// return (true, false) if element was not found.
// return (true, true) if element was found.
func includeElement(list interface{}, element interface{}) (ok, found bool) {

	listValue := reflect.ValueOf(list)
	listKind := reflect.TypeOf(list).Kind()
	defer func() {
		if e := recover(); e != nil {
			ok = false
			found = false
		}
	}()

	if listKind == reflect.String {
		elementValue := reflect.ValueOf(element)
		return true, strings.Contains(listValue.String(), elementValue.String())
	}

	if listKind == reflect.Map {
		mapKeys := listValue.MapKeys()
		for i := 0; i < len(mapKeys); i++ {
			if debuggoGen_ObjectsAreEqual(mapKeys[i].Interface(), element) {
				return true, true
			}
		}
		return true, false
	}

	for i := 0; i < listValue.Len(); i++ {
		if debuggoGen_ObjectsAreEqual(listValue.Index(i).Interface(), element) {
			return true, true
		}
	}
	return true, false

}

// Contains asserts that the specified string, list(array, slice...) or map contains the
// specified substring or element.
//
//    assert.Contains(t, "Hello World", "World")
//    assert.Contains(t, ["Hello", "World"], "World")
//    assert.Contains(t, {"Hello": "World"}, "Hello")
func Contains(s, contains interface{}, msgAndArgs ...interface{}) {
	debuggoGen_Contains(s, contains, msgAndArgs)
}

// NotContains asserts that the specified string, list(array, slice...) or map does NOT contain the
// specified substring or element.
//
//    assert.NotContains(t, "Hello World", "Earth")
//    assert.NotContains(t, ["Hello", "World"], "Earth")
//    assert.NotContains(t, {"Hello": "World"}, "Earth")
func NotContains(s, contains interface{}, msgAndArgs ...interface{}) {
	debuggoGen_NotContains(s, contains, msgAndArgs)
}

// Subset asserts that the specified list(array, slice...) contains all
// elements given in the specified subset(array, slice...).
//
//    assert.Subset(t, [1, 2, 3], [1, 2], "But [1, 2, 3] does contain [1, 2]")
func Subset(list, subset interface{}, msgAndArgs ...interface{}) {
	debuggoGen_Subset(list, subset, msgAndArgs)
}

// we consider nil to be equal to the nil set

// NotSubset asserts that the specified list(array, slice...) contains not all
// elements given in the specified subset(array, slice...).
//
//    assert.NotSubset(t, [1, 3, 4], [1, 2], "But [1, 3, 4] does not contain [1, 2]")
func NotSubset(list, subset interface{}, msgAndArgs ...interface{}) {
	debuggoGen_NotSubset(list, subset, msgAndArgs)
}

// ElementsMatch asserts that the specified listA(array, slice...) is equal to specified
// listB(array, slice...) ignoring the order of the elements. If there are duplicate elements,
// the number of appearances of each of them in both lists should match.
//
// assert.ElementsMatch(t, [1, 3, 2, 3], [1, 3, 3, 2])
func ElementsMatch(listA, listB interface{}, msgAndArgs ...interface{}) {
	debuggoGen_ElementsMatch(listA, listB, msgAndArgs)
}

// isList checks that the provided value is array or slice.
func isList(list interface{}, msgAndArgs ...interface{}) (ok bool) {
	kind := reflect.TypeOf(list).Kind()
	if kind != reflect.Array && kind != reflect.Slice {
		return debuggoGen_Fail(fmt.Sprintf("%q has an unsupported type %s, expecting array or slice", list, kind),
			msgAndArgs...)
	}
	return true
}

// diffLists diffs two arrays/slices and returns slices of elements that are only in A and only in B.
// If some element is present multiple times, each instance is counted separately (e.g. if something is 2x in A and
// 5x in B, it will be 0x in extraA and 3x in extraB). The order of items in both lists is ignored.
func diffLists(listA, listB interface{}) (extraA, extraB []interface{}) {
	aValue := reflect.ValueOf(listA)
	bValue := reflect.ValueOf(listB)

	aLen := aValue.Len()
	bLen := bValue.Len()

	// Mark indexes in bValue that we already used
	visited := make([]bool, bLen)
	for i := 0; i < aLen; i++ {
		element := aValue.Index(i).Interface()
		found := false
		for j := 0; j < bLen; j++ {
			if visited[j] {
				continue
			}
			if debuggoGen_ObjectsAreEqual(bValue.Index(j).Interface(), element) {
				visited[j] = true
				found = true
				break
			}
		}
		if !found {
			extraA = append(extraA, element)
		}
	}

	for j := 0; j < bLen; j++ {
		if visited[j] {
			continue
		}
		extraB = append(extraB, bValue.Index(j).Interface())
	}

	return
}

func formatListDiff(listA, listB interface{}, extraA, extraB []interface{}) string {
	var msg bytes.Buffer

	msg.WriteString("elements differ")
	if len(extraA) > 0 {
		msg.WriteString("\n\nextra elements in list A:\n")
		msg.WriteString(spewConfig.Sdump(extraA))
	}
	if len(extraB) > 0 {
		msg.WriteString("\n\nextra elements in list B:\n")
		msg.WriteString(spewConfig.Sdump(extraB))
	}
	msg.WriteString("\n\nlistA:\n")
	msg.WriteString(spewConfig.Sdump(listA))
	msg.WriteString("\n\nlistB:\n")
	msg.WriteString(spewConfig.Sdump(listB))

	return msg.String()
}

// Condition uses a Comparison to assert a complex condition.
func Condition(comp Comparison, msgAndArgs ...interface{}) { debuggoGen_Condition(comp, msgAndArgs) }

// PanicTestFunc defines a func that should be passed to the assert.Panics and assert.NotPanics
// methods, and represents a simple func that takes no arguments, and returns nothing.
type PanicTestFunc func()

// didPanic returns true if the function passed to it panics. Otherwise, it returns false.
func didPanic(f PanicTestFunc) (bool, interface{}, string) {

	didPanic := false
	var message interface{}
	var stack string
	func() {

		defer func() {
			if message = recover(); message != nil {
				didPanic = true
				stack = string(debug.Stack())
			}
		}()

		// call the target function
		f()

	}()

	return didPanic, message, stack

}

// Panics asserts that the code inside the specified PanicTestFunc panics.
//
//   assert.Panics(t, func(){ GoCrazy() })
func Panics(f PanicTestFunc, msgAndArgs ...interface{}) { debuggoGen_Panics(f, msgAndArgs) }

// PanicsWithValue asserts that the code inside the specified PanicTestFunc panics, and that
// the recovered panic value equals the expected panic value.
//
//   assert.PanicsWithValue(t, "crazy error", func(){ GoCrazy() })
func PanicsWithValue(expected interface{}, f PanicTestFunc, msgAndArgs ...interface{}) {
	debuggoGen_PanicsWithValue(expected, f, msgAndArgs)
}

// PanicsWithError asserts that the code inside the specified PanicTestFunc
// panics, and that the recovered panic value is an error that satisfies the
// EqualError comparison.
//
//   assert.PanicsWithError(t, "crazy error", func(){ GoCrazy() })
func PanicsWithError(errString string, f PanicTestFunc, msgAndArgs ...interface{}) {
	debuggoGen_PanicsWithError(errString, f, msgAndArgs)
}

// NotPanics asserts that the code inside the specified PanicTestFunc does NOT panic.
//
//   assert.NotPanics(t, func(){ RemainCalm() })
func NotPanics(f PanicTestFunc, msgAndArgs ...interface{}) { debuggoGen_NotPanics(f, msgAndArgs) }

// WithinDuration asserts that the two times are within duration delta of each other.
//
//   assert.WithinDuration(t, time.Now(), time.Now(), 10*time.Second)
func WithinDuration(expected, actual time.Time, delta time.Duration, msgAndArgs ...interface{}) {
	debuggoGen_WithinDuration(expected, actual, delta, msgAndArgs)
}

func toFloat(x interface{}) (float64, bool) {
	var xf float64
	xok := true

	switch xn := x.(type) {
	case uint:
		xf = float64(xn)
	case uint8:
		xf = float64(xn)
	case uint16:
		xf = float64(xn)
	case uint32:
		xf = float64(xn)
	case uint64:
		xf = float64(xn)
	case int:
		xf = float64(xn)
	case int8:
		xf = float64(xn)
	case int16:
		xf = float64(xn)
	case int32:
		xf = float64(xn)
	case int64:
		xf = float64(xn)
	case float32:
		xf = float64(xn)
	case float64:
		xf = xn
	case time.Duration:
		xf = float64(xn)
	default:
		xok = false
	}

	return xf, xok
}

// InDelta asserts that the two numerals are within delta of each other.
//
// 	 assert.InDelta(t, math.Pi, 22/7.0, 0.01)
func InDelta(expected, actual interface{}, delta float64, msgAndArgs ...interface{}) {
	debuggoGen_InDelta(expected, actual, delta, msgAndArgs)
}

// InDeltaSlice is the same as InDelta, except it compares two slices.
func InDeltaSlice(expected, actual interface{}, delta float64, msgAndArgs ...interface{}) {
	debuggoGen_InDeltaSlice(expected, actual, delta, msgAndArgs)
}

// InDeltaMapValues is the same as InDelta, but it compares all values between two maps. Both maps must have exactly the same keys.
func InDeltaMapValues(expected, actual interface{}, delta float64, msgAndArgs ...interface{}) {
	debuggoGen_InDeltaMapValues(expected, actual, delta, msgAndArgs)
}

func calcRelativeError(expected, actual interface{}) (float64, error) {
	af, aok := toFloat(expected)
	if !aok {
		return 0, fmt.Errorf("expected value %q cannot be converted to float", expected)
	}
	if math.IsNaN(af) {
		return 0, errors.New("expected value must not be NaN")
	}
	if af == 0 {
		return 0, fmt.Errorf("expected value must have a value other than zero to calculate the relative error")
	}
	bf, bok := toFloat(actual)
	if !bok {
		return 0, fmt.Errorf("actual value %q cannot be converted to float", actual)
	}
	if math.IsNaN(bf) {
		return 0, errors.New("actual value must not be NaN")
	}

	return math.Abs(af-bf) / math.Abs(af), nil
}

// InEpsilon asserts that expected and actual have a relative error less than epsilon
func InEpsilon(expected, actual interface{}, epsilon float64, msgAndArgs ...interface{}) {
	debuggoGen_InEpsilon(expected, actual, epsilon, msgAndArgs)
}

// InEpsilonSlice is the same as InEpsilon, except it compares each value from two slices.
func InEpsilonSlice(expected, actual interface{}, epsilon float64, msgAndArgs ...interface{}) {
	debuggoGen_InEpsilonSlice(expected, actual, epsilon, msgAndArgs)
}

/*
	Errors
*/

// NoError asserts that a function returned no error (i.e. `nil`).
//
//   actualObj, err := SomeFunction()
//   if assert.NoError(t, err) {
//	   assert.Equal(t, expectedObj, actualObj)
//   }
func NoError(err error, msgAndArgs ...interface{}) { debuggoGen_NoError(err, msgAndArgs) }

// Error asserts that a function returned an error (i.e. not `nil`).
//
//   actualObj, err := SomeFunction()
//   if assert.Error(t, err) {
//	   assert.Equal(t, expectedError, err)
//   }
func Error(err error, msgAndArgs ...interface{}) { debuggoGen_Error(err, msgAndArgs) }

// EqualError asserts that a function returned an error (i.e. not `nil`)
// and that it is equal to the provided error.
//
//   actualObj, err := SomeFunction()
//   assert.EqualError(t, err,  expectedErrorString)
func EqualError(theError error, errString string, msgAndArgs ...interface{}) {
	debuggoGen_EqualError(theError, errString, msgAndArgs)
}

// don't need to use deep equals here, we know they are both strings

// matchRegexp return true if a specified regexp matches a string.
func matchRegexp(rx interface{}, str interface{}) bool {

	var r *regexp.Regexp
	if rr, ok := rx.(*regexp.Regexp); ok {
		r = rr
	} else {
		r = regexp.MustCompile(fmt.Sprint(rx))
	}

	return (r.FindStringIndex(fmt.Sprint(str)) != nil)

}

// Regexp asserts that a specified regexp matches a string.
//
//  assert.Regexp(t, regexp.MustCompile("start"), "it's starting")
//  assert.Regexp(t, "start...$", "it's not starting")
func Regexp(rx interface{}, str interface{}, msgAndArgs ...interface{}) {
	debuggoGen_Regexp(rx, str, msgAndArgs)
}

// NotRegexp asserts that a specified regexp does not match a string.
//
//  assert.NotRegexp(t, regexp.MustCompile("starts"), "it's starting")
//  assert.NotRegexp(t, "^start", "it's not starting")
func NotRegexp(rx interface{}, str interface{}, msgAndArgs ...interface{}) {
	debuggoGen_NotRegexp(rx, str, msgAndArgs)
}

// Zero asserts that i is the zero value for its type.
func Zero(i interface{}, msgAndArgs ...interface{}) { debuggoGen_Zero(i, msgAndArgs) }

// NotZero asserts that i is not the zero value for its type.
func NotZero(i interface{}, msgAndArgs ...interface{}) { debuggoGen_NotZero(i, msgAndArgs) }

// FileExists checks whether a file exists in the given path. It also fails if
// the path points to a directory or there is an error when trying to check the file.
func FileExists(path string, msgAndArgs ...interface{}) { debuggoGen_FileExists(path, msgAndArgs) }

// NoFileExists checks whether a file does not exist in a given path. It fails
// if the path points to an existing _file_ only.
func NoFileExists(path string, msgAndArgs ...interface{}) { debuggoGen_NoFileExists(path, msgAndArgs) }

// DirExists checks whether a directory exists in the given path. It also fails
// if the path is a file rather a directory or there is an error checking whether it exists.
func DirExists(path string, msgAndArgs ...interface{}) { debuggoGen_DirExists(path, msgAndArgs) }

// NoDirExists checks whether a directory does not exist in the given path.
// It fails if the path points to an existing _directory_ only.
func NoDirExists(path string, msgAndArgs ...interface{}) { debuggoGen_NoDirExists(path, msgAndArgs) }

// JSONEq asserts that two JSON strings are equivalent.
//
//  assert.JSONEq(t, `{"hello": "world", "foo": "bar"}`, `{"foo": "bar", "hello": "world"}`)
func JSONEq(expected string, actual string, msgAndArgs ...interface{}) {
	debuggoGen_JSONEq(expected, actual, msgAndArgs)
}

// YAMLEq asserts that two YAML strings are equivalent.
func YAMLEq(expected string, actual string, msgAndArgs ...interface{}) {
	debuggoGen_YAMLEq(expected, actual, msgAndArgs)
}

func typeAndKind(v interface{}) (reflect.Type, reflect.Kind) {
	t := reflect.TypeOf(v)
	k := t.Kind()

	if k == reflect.Ptr {
		t = t.Elem()
		k = t.Kind()
	}
	return t, k
}

// diff returns a diff of both values as long as both are of the same type and
// are a struct, map, slice, array or string. Otherwise it returns an empty string.
func diff(expected interface{}, actual interface{}) string {
	if expected == nil || actual == nil {
		return ""
	}

	et, ek := typeAndKind(expected)
	at, _ := typeAndKind(actual)

	if et != at {
		return ""
	}

	if ek != reflect.Struct && ek != reflect.Map && ek != reflect.Slice && ek != reflect.Array && ek != reflect.String {
		return ""
	}

	var e, a string
	if et != reflect.TypeOf("") {
		e = spewConfig.Sdump(expected)
		a = spewConfig.Sdump(actual)
	} else {
		e = reflect.ValueOf(expected).String()
		a = reflect.ValueOf(actual).String()
	}

	diff, _ := difflib.GetUnifiedDiffString(difflib.UnifiedDiff{
		A:        difflib.SplitLines(e),
		B:        difflib.SplitLines(a),
		FromFile: "Expected",
		FromDate: "",
		ToFile:   "Actual",
		ToDate:   "",
		Context:  1,
	})

	return "\n\nDiff:\n" + diff
}

func isFunction(arg interface{}) bool {
	if arg == nil {
		return false
	}
	return reflect.TypeOf(arg).Kind() == reflect.Func
}

var spewConfig = spew.ConfigState{
	Indent:                  " ",
	DisablePointerAddresses: true,
	DisableCapacities:       true,
	SortKeys:                true,
	DisableMethods:          true,
}

type tHelper interface {
	Helper()
}

// Eventually asserts that given condition will be met in waitFor time,
// periodically checking target function each tick.
//
//    assert.Eventually(t, func() bool { return true; }, time.Second, 10*time.Millisecond)
func Eventually(condition func() bool, waitFor time.Duration, tick time.Duration, msgAndArgs ...interface{}) {
	debuggoGen_Eventually(condition, waitFor, tick, msgAndArgs)
}

// Never asserts that the given condition doesn't satisfy in waitFor time,
// periodically checking the target function each tick.
//
//    assert.Never(t, func() bool { return false; }, time.Second, 10*time.Millisecond)
func Never(condition func() bool, waitFor time.Duration, tick time.Duration, msgAndArgs ...interface{}) {
	debuggoGen_Never(condition, waitFor, tick, msgAndArgs)
}

// ErrorIs asserts that at least one of the errors in err's chain matches target.
// This is a wrapper for errors.Is.
func ErrorIs(err, target error, msgAndArgs ...interface{}) {
	debuggoGen_ErrorIs(err, target, msgAndArgs)
}

// NotErrorIs asserts that at none of the errors in err's chain matches target.
// This is a wrapper for errors.Is.
func NotErrorIs(err, target error, msgAndArgs ...interface{}) {
	debuggoGen_NotErrorIs(err, target, msgAndArgs)
}

// ErrorAs asserts that at least one of the errors in err's chain matches target, and if so, sets target to that error value.
// This is a wrapper for errors.As.
func ErrorAs(err error, target interface{}, msgAndArgs ...interface{}) {
	debuggoGen_ErrorAs(err, target, msgAndArgs)
}

func buildErrorChainString(err error) string {
	if err == nil {
		return ""
	}

	e := errors.Unwrap(err)
	chain := fmt.Sprintf("%q", err.Error())
	for e != nil {
		chain += fmt.Sprintf("\n\t%q", e.Error())
		e = errors.Unwrap(e)
	}
	return chain
}

func debuggoGen_ObjectsAreEqual(expected, actual interface{}) bool {
	if expected == nil || actual == nil {
		return expected == actual
	}

	exp, ok := expected.([]byte)
	if !ok {
		return reflect.DeepEqual(expected, actual)
	}

	act, ok := actual.([]byte)
	if !ok {
		return false
	}
	if exp == nil || act == nil {
		return exp == nil && act == nil
	}
	return bytes.Equal(exp, act)
}

func debuggoGen_ObjectsAreEqualValues(expected, actual interface{}) bool {
	if debuggoGen_ObjectsAreEqual(expected, actual) {
		return true
	}

	actualType := reflect.TypeOf(actual)
	if actualType == nil {
		return false
	}
	expectedValue := reflect.ValueOf(expected)
	if expectedValue.IsValid() && expectedValue.Type().ConvertibleTo(actualType) {

		return reflect.DeepEqual(expectedValue.Convert(actualType).Interface(), actual)
	}

	return false
}

func debuggoGen_CallerInfo() []string {

	var pc uintptr
	var ok bool
	var file string
	var line int
	var name string

	callers := []string{}
	for i := 0; ; i++ {
		pc, file, line, ok = runtime.Caller(i)
		if !ok {

			break
		}

		if file == "<autogenerated>" {
			break
		}

		f := runtime.FuncForPC(pc)
		if f == nil {
			break
		}
		name = f.Name()

		if name == "testing.tRunner" {
			break
		}

		parts := strings.Split(file, "/")
		file = parts[len(parts)-1]
		if len(parts) > 1 {
			dir := parts[len(parts)-2]
			if (dir != "assert" && dir != "mock" && dir != "require") || file == "mock_test.go" {
				callers = append(callers, fmt.Sprintf("%s:%d", file, line))
			}
		}

		segments := strings.Split(name, ".")
		name = segments[len(segments)-1]
		if isTest(name, "Test") ||
			isTest(name, "Benchmark") ||
			isTest(name, "Example") {
			break
		}
	}

	return callers
}

func debuggoGen_FailNow(failureMessage string, msgAndArgs ...interface{}) bool {
	debuggoGen_Fail(failureMessage, msgAndArgs...)

	return false
}

func debuggoGen_Fail(failureMessage string, msgAndArgs ...interface{}) bool {

	content := []labeledContent{
		{"Error Trace", strings.Join(debuggoGen_CallerInfo(), "\n\t\t\t")},
		{"Error", failureMessage},
	}

	message := messageFromMsgAndArgs(msgAndArgs...)
	if len(message) > 0 {
		content = append(content, labeledContent{"Messages", message})
	}
	panic(fmt.Sprint("\n%s", ""+labeledOutput(content...)))

	return false
}

func debuggoGen_Implements(interfaceObject interface{}, object interface{}, msgAndArgs ...interface{}) bool {

	interfaceType := reflect.TypeOf(interfaceObject).Elem()

	if object == nil {
		return debuggoGen_Fail(fmt.Sprintf("Cannot check if nil implements %v", interfaceType), msgAndArgs...)
	}
	if !reflect.TypeOf(object).Implements(interfaceType) {
		return debuggoGen_Fail(fmt.Sprintf("%T must implement %v", object, interfaceType), msgAndArgs...)
	}

	return true
}

func debuggoGen_IsType(expectedType interface{}, object interface{}, msgAndArgs ...interface{}) bool {

	if !debuggoGen_ObjectsAreEqual(reflect.TypeOf(object), reflect.TypeOf(expectedType)) {
		return debuggoGen_Fail(fmt.Sprintf("Object expected to be of type %v, but was %v", reflect.TypeOf(expectedType), reflect.TypeOf(object)), msgAndArgs...)
	}

	return true
}

func debuggoGen_Equal(expected, actual interface{}, msgAndArgs ...interface{}) bool {

	if err := validateEqualArgs(expected, actual); err != nil {
		return debuggoGen_Fail(fmt.Sprintf("Invalid operation: %#v == %#v (%s)",
			expected, actual, err), msgAndArgs...)
	}

	if !debuggoGen_ObjectsAreEqual(expected, actual) {
		diff := diff(expected, actual)
		expected, actual = formatUnequalValues(expected, actual)
		return debuggoGen_Fail(fmt.Sprintf("Not equal: \n"+
			"expected: %s\n"+
			"actual  : %s%s", expected, actual, diff), msgAndArgs...)
	}

	return true

}

func debuggoGen_Same(expected, actual interface{}, msgAndArgs ...interface{}) bool {

	if !samePointers(expected, actual) {
		return debuggoGen_Fail(fmt.Sprintf("Not same: \n"+
			"expected: %p %#v\n"+
			"actual  : %p %#v", expected, expected, actual, actual), msgAndArgs...)
	}

	return true
}

func debuggoGen_NotSame(expected, actual interface{}, msgAndArgs ...interface{}) bool {

	if samePointers(expected, actual) {
		return debuggoGen_Fail(fmt.Sprintf(
			"Expected and actual point to the same object: %p %#v",
			expected, expected), msgAndArgs...)
	}
	return true
}

func debuggoGen_EqualValues(expected, actual interface{}, msgAndArgs ...interface{}) bool {

	if !debuggoGen_ObjectsAreEqualValues(expected, actual) {
		diff := diff(expected, actual)
		expected, actual = formatUnequalValues(expected, actual)
		return debuggoGen_Fail(fmt.Sprintf("Not equal: \n"+
			"expected: %s\n"+
			"actual  : %s%s", expected, actual, diff), msgAndArgs...)
	}

	return true

}

func debuggoGen_Exactly(expected, actual interface{}, msgAndArgs ...interface{}) bool {

	aType := reflect.TypeOf(expected)
	bType := reflect.TypeOf(actual)

	if aType != bType {
		return debuggoGen_Fail(fmt.Sprintf("Types expected to match exactly\n\t%v != %v", aType, bType), msgAndArgs...)
	}

	return debuggoGen_Equal(expected, actual, msgAndArgs...)

}

func debuggoGen_NotNil(object interface{}, msgAndArgs ...interface{}) bool {
	if !isNil(object) {
		return true
	}

	return debuggoGen_Fail("Expected value not to be nil.", msgAndArgs...)
}

func debuggoGen_Nil(object interface{}, msgAndArgs ...interface{}) bool {
	if isNil(object) {
		return true
	}

	return debuggoGen_Fail(fmt.Sprintf("Expected nil, but got: %#v", object), msgAndArgs...)
}

func debuggoGen_Empty(object interface{}, msgAndArgs ...interface{}) bool {
	pass := isEmpty(object)
	if !pass {
		debuggoGen_Fail(fmt.Sprintf("Should be empty, but was %v", object), msgAndArgs...)
	}

	return pass

}

func debuggoGen_NotEmpty(object interface{}, msgAndArgs ...interface{}) bool {
	pass := !isEmpty(object)
	if !pass {
		debuggoGen_Fail(fmt.Sprintf("Should NOT be empty, but was %v", object), msgAndArgs...)
	}

	return pass

}

func debuggoGen_Len(object interface{}, length int, msgAndArgs ...interface{}) bool {

	ok, l := getLen(object)
	if !ok {
		return debuggoGen_Fail(fmt.Sprintf("\"%s\" could not be applied builtin len()", object), msgAndArgs...)
	}

	if l != length {
		return debuggoGen_Fail(fmt.Sprintf("\"%s\" should have %d item(s), but has %d", object, length, l), msgAndArgs...)
	}
	return true
}

func debuggoGen_True(value bool, msgAndArgs ...interface{}) bool {
	if !value {

		return debuggoGen_Fail("Should be true", msgAndArgs...)
	}

	return true

}

func debuggoGen_False(value bool, msgAndArgs ...interface{}) bool {
	if value {

		return debuggoGen_Fail("Should be false", msgAndArgs...)
	}

	return true

}

func debuggoGen_NotEqual(expected, actual interface{}, msgAndArgs ...interface{}) bool {

	if err := validateEqualArgs(expected, actual); err != nil {
		return debuggoGen_Fail(fmt.Sprintf("Invalid operation: %#v != %#v (%s)",
			expected, actual, err), msgAndArgs...)
	}

	if debuggoGen_ObjectsAreEqual(expected, actual) {
		return debuggoGen_Fail(fmt.Sprintf("Should not be: %#v\n", actual), msgAndArgs...)
	}

	return true

}

func debuggoGen_NotEqualValues(expected, actual interface{}, msgAndArgs ...interface{}) bool {

	if debuggoGen_ObjectsAreEqualValues(expected, actual) {
		return debuggoGen_Fail(fmt.Sprintf("Should not be: %#v\n", actual), msgAndArgs...)
	}

	return true
}

func debuggoGen_Contains(s, contains interface{}, msgAndArgs ...interface{}) bool {

	ok, found := includeElement(s, contains)
	if !ok {
		return debuggoGen_Fail(fmt.Sprintf("%#v could not be applied builtin len()", s), msgAndArgs...)
	}
	if !found {
		return debuggoGen_Fail(fmt.Sprintf("%#v does not contain %#v", s, contains), msgAndArgs...)
	}

	return true

}

func debuggoGen_NotContains(s, contains interface{}, msgAndArgs ...interface{}) bool {

	ok, found := includeElement(s, contains)
	if !ok {
		return debuggoGen_Fail(fmt.Sprintf("\"%s\" could not be applied builtin len()", s), msgAndArgs...)
	}
	if found {
		return debuggoGen_Fail(fmt.Sprintf("\"%s\" should not contain \"%s\"", s, contains), msgAndArgs...)
	}

	return true

}

func debuggoGen_Subset(list, subset interface{}, msgAndArgs ...interface{}) (ok bool) {

	if subset == nil {
		return true
	}

	subsetValue := reflect.ValueOf(subset)
	defer func() {
		if e := recover(); e != nil {
			ok = false
		}
	}()

	listKind := reflect.TypeOf(list).Kind()
	subsetKind := reflect.TypeOf(subset).Kind()

	if listKind != reflect.Array && listKind != reflect.Slice {
		return debuggoGen_Fail(fmt.Sprintf("%q has an unsupported type %s", list, listKind), msgAndArgs...)
	}

	if subsetKind != reflect.Array && subsetKind != reflect.Slice {
		return debuggoGen_Fail(fmt.Sprintf("%q has an unsupported type %s", subset, subsetKind), msgAndArgs...)
	}

	for i := 0; i < subsetValue.Len(); i++ {
		element := subsetValue.Index(i).Interface()
		ok, found := includeElement(list, element)
		if !ok {
			return debuggoGen_Fail(fmt.Sprintf("\"%s\" could not be applied builtin len()", list), msgAndArgs...)
		}
		if !found {
			return debuggoGen_Fail(fmt.Sprintf("\"%s\" does not contain \"%s\"", list, element), msgAndArgs...)
		}
	}

	return true
}

func debuggoGen_NotSubset(list, subset interface{}, msgAndArgs ...interface{}) (ok bool) {

	if subset == nil {
		return debuggoGen_Fail(fmt.Sprintf("nil is the empty set which is a subset of every set"), msgAndArgs...)
	}

	subsetValue := reflect.ValueOf(subset)
	defer func() {
		if e := recover(); e != nil {
			ok = false
		}
	}()

	listKind := reflect.TypeOf(list).Kind()
	subsetKind := reflect.TypeOf(subset).Kind()

	if listKind != reflect.Array && listKind != reflect.Slice {
		return debuggoGen_Fail(fmt.Sprintf("%q has an unsupported type %s", list, listKind), msgAndArgs...)
	}

	if subsetKind != reflect.Array && subsetKind != reflect.Slice {
		return debuggoGen_Fail(fmt.Sprintf("%q has an unsupported type %s", subset, subsetKind), msgAndArgs...)
	}

	for i := 0; i < subsetValue.Len(); i++ {
		element := subsetValue.Index(i).Interface()
		ok, found := includeElement(list, element)
		if !ok {
			return debuggoGen_Fail(fmt.Sprintf("\"%s\" could not be applied builtin len()", list), msgAndArgs...)
		}
		if !found {
			return true
		}
	}

	return debuggoGen_Fail(fmt.Sprintf("%q is a subset of %q", subset, list), msgAndArgs...)
}

func debuggoGen_ElementsMatch(listA, listB interface{}, msgAndArgs ...interface{}) (ok bool) {

	if isEmpty(listA) && isEmpty(listB) {
		return true
	}

	if !isList(listA, msgAndArgs...) || !isList(listB, msgAndArgs...) {
		return false
	}

	extraA, extraB := diffLists(listA, listB)

	if len(extraA) == 0 && len(extraB) == 0 {
		return true
	}

	return debuggoGen_Fail(formatListDiff(listA, listB, extraA, extraB), msgAndArgs...)
}

func debuggoGen_Condition(comp Comparison, msgAndArgs ...interface{}) bool {

	result := comp()
	if !result {
		debuggoGen_Fail("Condition failed!", msgAndArgs...)
	}
	return result
}

func debuggoGen_Panics(f PanicTestFunc, msgAndArgs ...interface{}) bool {

	if funcDidPanic, panicValue, _ := didPanic(f); !funcDidPanic {
		return debuggoGen_Fail(fmt.Sprintf("func %#v should panic\n\tPanic value:\t%#v", f, panicValue), msgAndArgs...)
	}

	return true
}

func debuggoGen_PanicsWithValue(expected interface{}, f PanicTestFunc, msgAndArgs ...interface{}) bool {

	funcDidPanic, panicValue, panickedStack := didPanic(f)
	if !funcDidPanic {
		return debuggoGen_Fail(fmt.Sprintf("func %#v should panic\n\tPanic value:\t%#v", f, panicValue), msgAndArgs...)
	}
	if panicValue != expected {
		return debuggoGen_Fail(fmt.Sprintf("func %#v should panic with value:\t%#v\n\tPanic value:\t%#v\n\tPanic stack:\t%s", f, expected, panicValue, panickedStack), msgAndArgs...)
	}

	return true
}

func debuggoGen_PanicsWithError(errString string, f PanicTestFunc, msgAndArgs ...interface{}) bool {

	funcDidPanic, panicValue, panickedStack := didPanic(f)
	if !funcDidPanic {
		return debuggoGen_Fail(fmt.Sprintf("func %#v should panic\n\tPanic value:\t%#v", f, panicValue), msgAndArgs...)
	}
	panicErr, ok := panicValue.(error)
	if !ok || panicErr.Error() != errString {
		return debuggoGen_Fail(fmt.Sprintf("func %#v should panic with error message:\t%#v\n\tPanic value:\t%#v\n\tPanic stack:\t%s", f, errString, panicValue, panickedStack), msgAndArgs...)
	}

	return true
}

func debuggoGen_NotPanics(f PanicTestFunc, msgAndArgs ...interface{}) bool {

	if funcDidPanic, panicValue, panickedStack := didPanic(f); funcDidPanic {
		return debuggoGen_Fail(fmt.Sprintf("func %#v should not panic\n\tPanic value:\t%v\n\tPanic stack:\t%s", f, panicValue, panickedStack), msgAndArgs...)
	}

	return true
}

func debuggoGen_WithinDuration(expected, actual time.Time, delta time.Duration, msgAndArgs ...interface{}) bool {

	dt := expected.Sub(actual)
	if dt < -delta || dt > delta {
		return debuggoGen_Fail(fmt.Sprintf("Max difference between %v and %v allowed is %v, but difference was %v", expected, actual, delta, dt), msgAndArgs...)
	}

	return true
}

func debuggoGen_InDelta(expected, actual interface{}, delta float64, msgAndArgs ...interface{}) bool {

	af, aok := toFloat(expected)
	bf, bok := toFloat(actual)

	if !aok || !bok {
		return debuggoGen_Fail(fmt.Sprintf("Parameters must be numerical"), msgAndArgs...)
	}

	if math.IsNaN(af) {
		return debuggoGen_Fail(fmt.Sprintf("Expected must not be NaN"), msgAndArgs...)
	}

	if math.IsNaN(bf) {
		return debuggoGen_Fail(fmt.Sprintf("Expected %v with delta %v, but was NaN", expected, delta), msgAndArgs...)
	}

	dt := af - bf
	if dt < -delta || dt > delta {
		return debuggoGen_Fail(fmt.Sprintf("Max difference between %v and %v allowed is %v, but difference was %v", expected, actual, delta, dt), msgAndArgs...)
	}

	return true
}

func debuggoGen_InDeltaSlice(expected, actual interface{}, delta float64, msgAndArgs ...interface{}) bool {

	if expected == nil || actual == nil ||
		reflect.TypeOf(actual).Kind() != reflect.Slice ||
		reflect.TypeOf(expected).Kind() != reflect.Slice {
		return debuggoGen_Fail(fmt.Sprintf("Parameters must be slice"), msgAndArgs...)
	}

	actualSlice := reflect.ValueOf(actual)
	expectedSlice := reflect.ValueOf(expected)

	for i := 0; i < actualSlice.Len(); i++ {
		result := debuggoGen_InDelta(actualSlice.Index(i).Interface(), expectedSlice.Index(i).Interface(), delta, msgAndArgs...)
		if !result {
			return result
		}
	}

	return true
}

func debuggoGen_InDeltaMapValues(expected, actual interface{}, delta float64, msgAndArgs ...interface{}) bool {

	if expected == nil || actual == nil ||
		reflect.TypeOf(actual).Kind() != reflect.Map ||
		reflect.TypeOf(expected).Kind() != reflect.Map {
		return debuggoGen_Fail("Arguments must be maps", msgAndArgs...)
	}

	expectedMap := reflect.ValueOf(expected)
	actualMap := reflect.ValueOf(actual)

	if expectedMap.Len() != actualMap.Len() {
		return debuggoGen_Fail("Arguments must have the same number of keys", msgAndArgs...)
	}

	for _, k := range expectedMap.MapKeys() {
		ev := expectedMap.MapIndex(k)
		av := actualMap.MapIndex(k)

		if !ev.IsValid() {
			return debuggoGen_Fail(fmt.Sprintf("missing key %q in expected map", k), msgAndArgs...)
		}

		if !av.IsValid() {
			return debuggoGen_Fail(fmt.Sprintf("missing key %q in actual map", k), msgAndArgs...)
		}

		if !debuggoGen_InDelta(

			ev.Interface(),
			av.Interface(),
			delta,
			msgAndArgs...,
		) {
			return false
		}
	}

	return true
}

func debuggoGen_InEpsilon(expected, actual interface{}, epsilon float64, msgAndArgs ...interface{}) bool {

	if math.IsNaN(epsilon) {
		return debuggoGen_Fail("epsilon must not be NaN")
	}
	actualEpsilon, err := calcRelativeError(expected, actual)
	if err != nil {
		return debuggoGen_Fail(err.Error(), msgAndArgs...)
	}
	if actualEpsilon > epsilon {
		return debuggoGen_Fail(fmt.Sprintf("Relative error is too high: %#v (expected)\n"+
			"        < %#v (actual)", epsilon, actualEpsilon), msgAndArgs...)
	}

	return true
}

func debuggoGen_InEpsilonSlice(expected, actual interface{}, epsilon float64, msgAndArgs ...interface{}) bool {

	if expected == nil || actual == nil ||
		reflect.TypeOf(actual).Kind() != reflect.Slice ||
		reflect.TypeOf(expected).Kind() != reflect.Slice {
		return debuggoGen_Fail(fmt.Sprintf("Parameters must be slice"), msgAndArgs...)
	}

	actualSlice := reflect.ValueOf(actual)
	expectedSlice := reflect.ValueOf(expected)

	for i := 0; i < actualSlice.Len(); i++ {
		result := debuggoGen_InEpsilon(actualSlice.Index(i).Interface(), expectedSlice.Index(i).Interface(), epsilon)
		if !result {
			return result
		}
	}

	return true
}

func debuggoGen_NoError(err error, msgAndArgs ...interface{}) bool {
	if err != nil {

		return debuggoGen_Fail(fmt.Sprintf("Received unexpected error:\n%+v", err), msgAndArgs...)
	}

	return true
}

func debuggoGen_Error(err error, msgAndArgs ...interface{}) bool {
	if err == nil {

		return debuggoGen_Fail("An error is expected but got nil.", msgAndArgs...)
	}

	return true
}

func debuggoGen_EqualError(theError error, errString string, msgAndArgs ...interface{}) bool {

	if !debuggoGen_Error(theError, msgAndArgs...) {
		return false
	}
	expected := errString
	actual := theError.Error()

	if expected != actual {
		return debuggoGen_Fail(fmt.Sprintf("Error message not equal:\n"+
			"expected: %q\n"+
			"actual  : %q", expected, actual), msgAndArgs...)
	}
	return true
}

func debuggoGen_Regexp(rx interface{}, str interface{}, msgAndArgs ...interface{}) bool {

	match := matchRegexp(rx, str)

	if !match {
		debuggoGen_Fail(fmt.Sprintf("Expect \"%v\" to match \"%v\"", str, rx), msgAndArgs...)
	}

	return match
}

func debuggoGen_NotRegexp(rx interface{}, str interface{}, msgAndArgs ...interface{}) bool {

	match := matchRegexp(rx, str)

	if match {
		debuggoGen_Fail(fmt.Sprintf("Expect \"%v\" to NOT match \"%v\"", str, rx), msgAndArgs...)
	}

	return !match

}

func debuggoGen_Zero(i interface{}, msgAndArgs ...interface{}) bool {

	if i != nil && !reflect.DeepEqual(i, reflect.Zero(reflect.TypeOf(i)).Interface()) {
		return debuggoGen_Fail(fmt.Sprintf("Should be zero, but was %v", i), msgAndArgs...)
	}
	return true
}

func debuggoGen_NotZero(i interface{}, msgAndArgs ...interface{}) bool {

	if i == nil || reflect.DeepEqual(i, reflect.Zero(reflect.TypeOf(i)).Interface()) {
		return debuggoGen_Fail(fmt.Sprintf("Should not be zero, but was %v", i), msgAndArgs...)
	}
	return true
}

func debuggoGen_FileExists(path string, msgAndArgs ...interface{}) bool {

	info, err := os.Lstat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return debuggoGen_Fail(fmt.Sprintf("unable to find file %q", path), msgAndArgs...)
		}
		return debuggoGen_Fail(fmt.Sprintf("error when running os.Lstat(%q): %s", path, err), msgAndArgs...)
	}
	if info.IsDir() {
		return debuggoGen_Fail(fmt.Sprintf("%q is a directory", path), msgAndArgs...)
	}
	return true
}

func debuggoGen_NoFileExists(path string, msgAndArgs ...interface{}) bool {

	info, err := os.Lstat(path)
	if err != nil {
		return true
	}
	if info.IsDir() {
		return true
	}
	return debuggoGen_Fail(fmt.Sprintf("file %q exists", path), msgAndArgs...)
}

func debuggoGen_DirExists(path string, msgAndArgs ...interface{}) bool {

	info, err := os.Lstat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return debuggoGen_Fail(fmt.Sprintf("unable to find file %q", path), msgAndArgs...)
		}
		return debuggoGen_Fail(fmt.Sprintf("error when running os.Lstat(%q): %s", path, err), msgAndArgs...)
	}
	if !info.IsDir() {
		return debuggoGen_Fail(fmt.Sprintf("%q is a file", path), msgAndArgs...)
	}
	return true
}

func debuggoGen_NoDirExists(path string, msgAndArgs ...interface{}) bool {

	info, err := os.Lstat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return true
		}
		return true
	}
	if !info.IsDir() {
		return true
	}
	return debuggoGen_Fail(fmt.Sprintf("directory %q exists", path), msgAndArgs...)
}

func debuggoGen_JSONEq(expected string, actual string, msgAndArgs ...interface{}) bool {

	var expectedJSONAsInterface, actualJSONAsInterface interface{}

	if err := json.Unmarshal([]byte(expected), &expectedJSONAsInterface); err != nil {
		return debuggoGen_Fail(fmt.Sprintf("Expected value ('%s') is not valid json.\nJSON parsing error: '%s'", expected, err.Error()), msgAndArgs...)
	}

	if err := json.Unmarshal([]byte(actual), &actualJSONAsInterface); err != nil {
		return debuggoGen_Fail(fmt.Sprintf("Input ('%s') needs to be valid json.\nJSON parsing error: '%s'", actual, err.Error()), msgAndArgs...)
	}

	return debuggoGen_Equal(expectedJSONAsInterface, actualJSONAsInterface, msgAndArgs...)
}

func debuggoGen_YAMLEq(expected string, actual string, msgAndArgs ...interface{}) bool {

	var expectedYAMLAsInterface, actualYAMLAsInterface interface{}

	if err := yaml.Unmarshal([]byte(expected), &expectedYAMLAsInterface); err != nil {
		return debuggoGen_Fail(fmt.Sprintf("Expected value ('%s') is not valid yaml.\nYAML parsing error: '%s'", expected, err.Error()), msgAndArgs...)
	}

	if err := yaml.Unmarshal([]byte(actual), &actualYAMLAsInterface); err != nil {
		return debuggoGen_Fail(fmt.Sprintf("Input ('%s') needs to be valid yaml.\nYAML error: '%s'", actual, err.Error()), msgAndArgs...)
	}

	return debuggoGen_Equal(expectedYAMLAsInterface, actualYAMLAsInterface, msgAndArgs...)
}

func debuggoGen_Eventually(condition func() bool, waitFor time.Duration, tick time.Duration, msgAndArgs ...interface{}) bool {

	ch := make(chan bool, 1)

	timer := time.NewTimer(waitFor)
	defer timer.Stop()

	ticker := time.NewTicker(tick)
	defer ticker.Stop()

	for tick := ticker.C; ; {
		select {
		case <-timer.C:
			return debuggoGen_Fail("Condition never satisfied", msgAndArgs...)
		case <-tick:
			tick = nil
			go func() { ch <- condition() }()
		case v := <-ch:
			if v {
				return true
			}
			tick = ticker.C
		}
	}
}

func debuggoGen_Never(condition func() bool, waitFor time.Duration, tick time.Duration, msgAndArgs ...interface{}) bool {

	ch := make(chan bool, 1)

	timer := time.NewTimer(waitFor)
	defer timer.Stop()

	ticker := time.NewTicker(tick)
	defer ticker.Stop()

	for tick := ticker.C; ; {
		select {
		case <-timer.C:
			return true
		case <-tick:
			tick = nil
			go func() { ch <- condition() }()
		case v := <-ch:
			if v {
				return debuggoGen_Fail("Condition satisfied", msgAndArgs...)
			}
			tick = ticker.C
		}
	}
}

func debuggoGen_ErrorIs(err, target error, msgAndArgs ...interface{}) bool {

	if errors.Is(err, target) {
		return true
	}

	var expectedText string
	if target != nil {
		expectedText = target.Error()
	}

	chain := buildErrorChainString(err)

	return debuggoGen_Fail(fmt.Sprintf("Target error should be in err chain:\n"+
		"expected: %q\n"+
		"in chain: %s", expectedText, chain,
	), msgAndArgs...)
}

func debuggoGen_NotErrorIs(err, target error, msgAndArgs ...interface{}) bool {

	if !errors.Is(err, target) {
		return true
	}

	var expectedText string
	if target != nil {
		expectedText = target.Error()
	}

	chain := buildErrorChainString(err)

	return debuggoGen_Fail(fmt.Sprintf("Target error should not be in err chain:\n"+
		"found: %q\n"+
		"in chain: %s", expectedText, chain,
	), msgAndArgs...)
}

func debuggoGen_ErrorAs(err error, target interface{}, msgAndArgs ...interface{}) bool {

	if errors.As(err, target) {
		return true
	}

	chain := buildErrorChainString(err)

	return debuggoGen_Fail(fmt.Sprintf("Should be in error chain:\n"+
		"expected: %q\n"+
		"in chain: %s", target, chain,
	), msgAndArgs...)
}
