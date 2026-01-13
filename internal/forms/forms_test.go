package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestNew(t *testing.T) {
	var form any
	r := httptest.NewRequest("POST", "/some-route", nil)
	form = New(r.PostForm)

	// check if return is *Form
	switch v := form.(type) {
	case *Form:
		// do nothing test passed
	default:
		t.Errorf("type is not *Form, type is %T", v)
	}
}

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/some-route", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when should have been valid")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/some-route", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("form shows valid when required fields missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "a")
	postedData.Add("c", "a")

	r, _ = http.NewRequest("POST", "/some-route", nil)

	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("shows does not have required fields when it should")
	}
}
