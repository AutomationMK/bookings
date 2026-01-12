package main

import (
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {
	var myH myHandler
	h := NoSurf(&myH)

	// test if return is http.Handler
	switch v := h.(type) {
	case http.Handler:
		// do nothing, test passed
	default:
		t.Errorf("type is not http.Handler, but is %T", v)
	}
}

func TestSessionLoad(t *testing.T) {
	var myH myHandler
	h := SessionLoad(&myH)

	// test if return is http.Handler
	switch v := h.(type) {
	case http.Handler:
		// do nothing, test passed
	default:
		t.Errorf("type is not http.Handler, but is %T", v)
	}
}
