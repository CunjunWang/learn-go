package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var tests = []struct {
	h       appHandler
	code    int
	message string
}{
	{errPanic, 500, "Internal Server Error"},
	{errUserError, 400, "user error"},
	{errNotFound, 404, "Not Found"},
	{errNoPermission, 403, "Forbidden"},
	{errUnknown, 500, "Internal Server Error"},
	{noErr, 200, "no error"},
}

func errPanic(writer http.ResponseWriter, r *http.Request) error {
	panic(123)
}

type testingUserError string

func (e testingUserError) Error() string {
	return e.Message()
}

func (e testingUserError) Message() string {
	return string(e)
}

func errUserError(writer http.ResponseWriter, r *http.Request) error {
	return testingUserError("user error")
}

func errNotFound(writer http.ResponseWriter, r *http.Request) error {
	return os.ErrNotExist
}

func errNoPermission(writer http.ResponseWriter, r *http.Request) error {
	return os.ErrPermission
}

func errUnknown(writer http.ResponseWriter, r *http.Request) error {
	return errors.New("unknown error")
}

func noErr(writer http.ResponseWriter, r *http.Request) error {
	fmt.Fprintln(writer, "no error")
	return nil
}

func TestErrWrapper(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.h)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "https://www.baidu.com", nil)
		f(response, request)
		verifyResponse(response.Result(), tt.code, tt.message, t)
	}
}

func TestErrWrapperInServer(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.h)
		server := httptest.NewServer(http.HandlerFunc(f))
		response, _ := http.Get(server.URL)
		verifyResponse(response, tt.code, tt.message, t)
	}
}

func verifyResponse(response *http.Response,
	expectedCode int, expectedMessage string, t *testing.T) {
	b, _ := ioutil.ReadAll(response.Body)
	body := strings.Trim(string(b), "\n")
	if response.StatusCode != expectedCode || body != expectedMessage {
		t.Errorf("Expected (%d, %s); got (%d, %s)",
			expectedCode, expectedMessage, response.StatusCode, body)
	}
}
