package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type postData struct {
	key   string
	value string
}

var theTests = []struct {
	name               string
	url                string
	method             string
	params             []postData
	expectedStatusCode int
}{
	{"home", "/", "GET", []postData{}, http.StatusOK},
	{"about", "/about", "GET", []postData{}, http.StatusOK},
	{"contact", "/contact", "GET", []postData{}, http.StatusOK},
	{"rooms", "/rooms", "GET", []postData{}, http.StatusOK},
	{"deluxe-room", "/rooms/deluxe-room", "GET", []postData{}, http.StatusOK},
	{"premium-suite", "/rooms/premium-suite", "GET", []postData{}, http.StatusOK},
	{"search-availibility", "/search-availability", "GET", []postData{}, http.StatusOK},
	{"reservation", "/make-reservation", "GET", []postData{}, http.StatusOK},
	{"post-search-availibility", "/search-availability", "POST", []postData{
		{key: "arrive_date", value: "1/12/2026"},
		{key: "departure_date", value: "1/14/2026"},
	}, http.StatusOK},
	{"post-search-availibility-json", "/search-availability-json", "POST", []postData{
		{key: "arrive_date", value: "1/12/2026"},
		{key: "departure_date", value: "1/14/2026"},
	}, http.StatusOK},
	{"post-reservation", "/make-reservation", "POST", []postData{
		{key: "first_name", value: "John"},
		{key: "last_name", value: "smith"},
		{key: "email", value: "johnSmith@ex.com"},
		{key: "phone", value: "123-456-7890"},
	}, http.StatusOK},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, e := range theTests {
		if e.method == "GET" {
			resp, err := ts.Client().Get(ts.URL + e.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s, expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
		} else {
			values := url.Values{}
			for _, x := range e.params {
				values.Add(x.key, x.value)
			}
			resp, err := ts.Client().PostForm(ts.URL+e.url, values)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s, expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
		}
	}
}
