package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
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

func TestForm_Has(t *testing.T) {
	// add dummy post data to test
	postedData := url.Values{}
	postedData.Add("a", "test")

	// create a test post request with encoded post data
	r := httptest.NewRequest(http.MethodPost, "/submit", strings.NewReader(postedData.Encode()))

	// set request header for form handling
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// populate r.Form and r.PostForm with url.Values
	r.ParseForm()

	// create *Form to use for testing
	form := New(r.PostForm)

	// test both cases of if a field does or does not exist
	if !form.Has("a", r) {
		t.Error("shows field not existing but it should")
	}
	if form.Has("b", r) {
		t.Error("shows field existing but in should not")
	}
}

func TestForm_MinLength(t *testing.T) {
	// add dummy post data to test
	postedData := url.Values{}
	postValue := "test"
	postedData.Add("a", postValue)

	// create a test post request with encoded post data
	r := httptest.NewRequest(http.MethodPost, "/submit", strings.NewReader(postedData.Encode()))

	// set request header for form handling
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// populate r.Form and r.PostForm with url.Values
	r.ParseForm()

	// create *Form to use for testing
	form := New(r.PostForm)

	if !form.MinLength("a", len(postValue), r) {
		t.Error("has length exactly the size of minimum legth but did not pass")
	}
	if !form.MinLength("a", len(postValue)-1, r) {
		t.Error("has length larger than minimum length but did not pass")
	}
	if form.MinLength("a", len(postValue)+1, r) {
		t.Error("has length smaller than minimum length but still passed")
	}
}
