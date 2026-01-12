package main

import (
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {
	var myH myHandler
	h := NoSurf(&myH)

	// test if return of NoSurf is
	switch v := h.(type) {
	case http.Handler:
		// do nothing, test passed
	default:
		t.Errorf("type is not http.Handler, but is %T", v)
	}
}
