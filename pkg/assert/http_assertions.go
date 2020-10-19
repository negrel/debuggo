// +build assert

package assert

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
)

// httpCode is a helper that returns HTTP code of the response. It returns -1 and
// an error if building a new request fails.
func httpCode(handler http.HandlerFunc, method, url string, values url.Values) (int, error) {
	w := httptest.NewRecorder()
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return -1, err
	}
	req.URL.RawQuery = values.Encode()
	handler(w, req)
	return w.Code, nil
}

// HTTPSuccess asserts that a specified handler returns a success status code.
//
//  assert.HTTPSuccess(t, myHandler, "POST", "http://www.google.com", nil)
//
// Returns whether the assertion was successful (true) or not (false).
func HTTPSuccess(handler http.HandlerFunc, method, url string, values url.Values, msgAndArgs ...interface{}) {
	debuggoGen_HTTPSuccess(handler, method, url, values, msgAndArgs)
}

// HTTPRedirect asserts that a specified handler returns a redirect status code.
//
//  assert.HTTPRedirect(t, myHandler, "GET", "/a/b/c", url.Values{"a": []string{"b", "c"}}
//
// Returns whether the assertion was successful (true) or not (false).
func HTTPRedirect(handler http.HandlerFunc, method, url string, values url.Values, msgAndArgs ...interface{}) {
	debuggoGen_HTTPRedirect(handler, method, url, values, msgAndArgs)
}

// HTTPError asserts that a specified handler returns an error status code.
//
//  assert.HTTPError(t, myHandler, "POST", "/a/b/c", url.Values{"a": []string{"b", "c"}}
//
// Returns whether the assertion was successful (true) or not (false).
func HTTPError(handler http.HandlerFunc, method, url string, values url.Values, msgAndArgs ...interface{}) {
	debuggoGen_HTTPError(handler, method, url, values, msgAndArgs)
}

// HTTPStatusCode asserts that a specified handler returns a specified status code.
//
//  assert.HTTPStatusCode(t, myHandler, "GET", "/notImplemented", nil, 501)
//
// Returns whether the assertion was successful (true) or not (false).
func HTTPStatusCode(handler http.HandlerFunc, method, url string, values url.Values, statuscode int, msgAndArgs ...interface{}) {
	debuggoGen_HTTPStatusCode(handler, method, url, values, statuscode, msgAndArgs)
}

// HTTPBody is a helper that returns HTTP body of the response. It returns
// empty string if building a new request fails.
func HTTPBody(handler http.HandlerFunc, method, url string, values url.Values) {
	debuggoGen_HTTPBody(handler, method, url, values)
}

// HTTPBodyContains asserts that a specified handler returns a
// body that contains a string.
//
//  assert.HTTPBodyContains(t, myHandler, "GET", "www.google.com", nil, "I'm Feeling Lucky")
//
// Returns whether the assertion was successful (true) or not (false).
func HTTPBodyContains(handler http.HandlerFunc, method, url string, values url.Values, str interface{}, msgAndArgs ...interface{}) {
	debuggoGen_HTTPBodyContains(handler, method, url, values, str, msgAndArgs)
}

// HTTPBodyNotContains asserts that a specified handler returns a
// body that does not contain a string.
//
//  assert.HTTPBodyNotContains(t, myHandler, "GET", "www.google.com", nil, "I'm Feeling Lucky")
//
// Returns whether the assertion was successful (true) or not (false).
func HTTPBodyNotContains(handler http.HandlerFunc, method, url string, values url.Values, str interface{}, msgAndArgs ...interface{}) {
	debuggoGen_HTTPBodyNotContains(handler, method, url, values, str, msgAndArgs)
}

func debuggoGen_HTTPSuccess(handler http.HandlerFunc, method, url string, values url.Values, msgAndArgs ...interface{}) bool {

	code, err := httpCode(handler, method, url, values)
	if err != nil {
		debuggoGen_Fail(fmt.Sprintf("Failed to build test request, got error: %s", err))
	}

	isSuccessCode := code >= http.StatusOK && code <= http.StatusPartialContent
	if !isSuccessCode {
		debuggoGen_Fail(fmt.Sprintf("Expected HTTP success status code for %q but received %d", url+"?"+values.Encode(), code))
	}

	return isSuccessCode
}

func debuggoGen_HTTPRedirect(handler http.HandlerFunc, method, url string, values url.Values, msgAndArgs ...interface{}) bool {

	code, err := httpCode(handler, method, url, values)
	if err != nil {
		debuggoGen_Fail(fmt.Sprintf("Failed to build test request, got error: %s", err))
	}

	isRedirectCode := code >= http.StatusMultipleChoices && code <= http.StatusTemporaryRedirect
	if !isRedirectCode {
		debuggoGen_Fail(fmt.Sprintf("Expected HTTP redirect status code for %q but received %d", url+"?"+values.Encode(), code))
	}

	return isRedirectCode
}

func debuggoGen_HTTPError(handler http.HandlerFunc, method, url string, values url.Values, msgAndArgs ...interface{}) bool {

	code, err := httpCode(handler, method, url, values)
	if err != nil {
		debuggoGen_Fail(fmt.Sprintf("Failed to build test request, got error: %s", err))
	}

	isErrorCode := code >= http.StatusBadRequest
	if !isErrorCode {
		debuggoGen_Fail(fmt.Sprintf("Expected HTTP error status code for %q but received %d", url+"?"+values.Encode(), code))
	}

	return isErrorCode
}

func debuggoGen_HTTPStatusCode(handler http.HandlerFunc, method, url string, values url.Values, statuscode int, msgAndArgs ...interface{}) bool {

	code, err := httpCode(handler, method, url, values)
	if err != nil {
		debuggoGen_Fail(fmt.Sprintf("Failed to build test request, got error: %s", err))
	}

	successful := code == statuscode
	if !successful {
		debuggoGen_Fail(fmt.Sprintf("Expected HTTP status code %d for %q but received %d", statuscode, url+"?"+values.Encode(), code))
	}

	return successful
}

func debuggoGen_HTTPBody(handler http.HandlerFunc, method, url string, values url.Values) string {
	w := httptest.NewRecorder()
	req, err := http.NewRequest(method, url+"?"+values.Encode(), nil)
	if err != nil {
		return ""
	}
	handler(w, req)
	return w.Body.String()
}

func debuggoGen_HTTPBodyContains(handler http.HandlerFunc, method, url string, values url.Values, str interface{}, msgAndArgs ...interface{}) bool {

	body := debuggoGen_HTTPBody(handler, method, url, values)

	contains := strings.Contains(body, fmt.Sprint(str))
	if !contains {
		debuggoGen_Fail(fmt.Sprintf("Expected response body for \"%s\" to contain \"%s\" but found \"%s\"", url+"?"+values.Encode(), str, body))
	}

	return contains
}

func debuggoGen_HTTPBodyNotContains(handler http.HandlerFunc, method, url string, values url.Values, str interface{}, msgAndArgs ...interface{}) bool {

	body := debuggoGen_HTTPBody(handler, method, url, values)

	contains := strings.Contains(body, fmt.Sprint(str))
	if contains {
		debuggoGen_Fail(fmt.Sprintf("Expected response body for \"%s\" to NOT contain \"%s\" but found \"%s\"", url+"?"+values.Encode(), str, body))
	}

	return !contains
}
