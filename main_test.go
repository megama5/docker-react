package main

import (
	"net/http"
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

	if rw.buf != "Hello from Go on Elastic Beanstalk\n" {
		t.Fatal(rw.buf, "!=", "Hello from Go on Elastic Beanstalk")
	}
}
