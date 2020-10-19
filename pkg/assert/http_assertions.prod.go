// +build !assert

package assert

import (
	"net/http"
	"net/url"
)

func HTTPSuccess( // HTTPSuccess asserts that a specified handler returns a success status code.
//
//  assert.HTTPSuccess(t, myHandler, "POST", "http://www.google.com", nil)
//
// Returns whether the assertion was successful (true) or not (false).
_ http.HandlerFunc, _, _ string, _ url.Values, _ ...interface{}) {
}

func HTTPRedirect( // HTTPRedirect asserts that a specified handler returns a redirect status code.
//
//  assert.HTTPRedirect(t, myHandler, "GET", "/a/b/c", url.Values{"a": []string{"b", "c"}}
//
// Returns whether the assertion was successful (true) or not (false).
_ http.HandlerFunc, _, _ string, _ url.Values, _ ...interface{}) {
}

func HTTPError( // HTTPError asserts that a specified handler returns an error status code.
//
//  assert.HTTPError(t, myHandler, "POST", "/a/b/c", url.Values{"a": []string{"b", "c"}}
//
// Returns whether the assertion was successful (true) or not (false).
_ http.HandlerFunc, _, _ string, _ url.Values, _ ...interface{}) {
}

func HTTPStatusCode( // HTTPStatusCode asserts that a specified handler returns a specified status code.
//
//  assert.HTTPStatusCode(t, myHandler, "GET", "/notImplemented", nil, 501)
//
// Returns whether the assertion was successful (true) or not (false).
_ http.HandlerFunc, _, _ string, _ url.Values, _ int, _ ...interface{}) {
}

func HTTPBody( // HTTPBody is a helper that returns HTTP body of the response. It returns
// empty string if building a new request fails.
_ http.HandlerFunc, _, _ string, _ url.Values) {
}

func HTTPBodyContains( // HTTPBodyContains asserts that a specified handler returns a
// body that contains a string.
//
//  assert.HTTPBodyContains(t, myHandler, "GET", "www.google.com", nil, "I'm Feeling Lucky")
//
// Returns whether the assertion was successful (true) or not (false).
_ http.HandlerFunc, _, _ string, _ url.Values, _ interface{}, _ ...interface{}) {
}

func HTTPBodyNotContains( // HTTPBodyNotContains asserts that a specified handler returns a
// body that does not contain a string.
//
//  assert.HTTPBodyNotContains(t, myHandler, "GET", "www.google.com", nil, "I'm Feeling Lucky")
//
// Returns whether the assertion was successful (true) or not (false).
_ http.HandlerFunc, _, _ string, _ url.Values, _ interface{}, _ ...interface{}) {
}
