package main

import (
	"net/http"
	"strings"
	"testing"
)

type MockedResponseWriter struct {
	http.ResponseWriter
	buf string
}

func (w *MockedResponseWriter) Write(b []byte) (int, error) {
	w.buf = string(b)
	return len(b), nil
}

func TestMainFunc(t *testing.T) {
	rw := &MockedResponseWriter{}
	response(rw, nil)

	if !strings.Contains(rw.buf, "Hello from Go on Elastic Beanstalk") {
		t.Fatal(rw.buf, "!=", "Hello from Go on Elastic Beanstalk")
	}
}
