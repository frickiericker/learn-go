package main

import (
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func Test_checkStatusCode_200(t *testing.T) {
	var res http.Response
	res.Status = "200 OK"
	res.StatusCode = 200
	err := checkStatusCode(&res)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func Test_checkStatusCode_404(t *testing.T) {
	var res http.Response
	res.Status = "404 Not Found"
	res.StatusCode = 404
	err := checkStatusCode(&res)
	if err == nil {
		t.Errorf("404 error not caught")
	}
}

// For mocking the body of HTTP response
type stringReadCloser struct {
	*strings.Reader
}

func (r *stringReadCloser) Close() error {
	return nil
}

func newStringReadCloser(s string) *stringReadCloser {
	return &stringReadCloser{strings.NewReader(s)}
}

func Test_parseJsonResponse_array(t *testing.T) {
	var res http.Response
	res.Status = "200 OK"
	res.StatusCode = 200
	res.Proto = "HTTP/1.0"
	res.ProtoMajor = 1
	res.ProtoMinor = 0
	res.Body = newStringReadCloser(`[1, 2, 3, 4]`)

	var actual []int
	err := parseJsonResponse(&res, &actual)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	expected := []int{1, 2, 3, 4}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("unexpected result: %v", actual)
	}
}

func Test_parseJsonResponse_simpleMap(t *testing.T) {
	var res http.Response
	res.Status = "200 OK"
	res.StatusCode = 200
	res.Proto = "HTTP/1.0"
	res.ProtoMajor = 1
	res.ProtoMinor = 0
	res.Body = newStringReadCloser(`{"apple": 10, "banana": 20}`)

	var actual map[string]int
	err := parseJsonResponse(&res, &actual)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	expected := map[string]int{"apple": 10, "banana": 20}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("unexpected result: %v", actual)
	}
}
