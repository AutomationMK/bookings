package handlers

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/AutomationMK/bookings/internal/models"
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
	//{"home", "/", "GET", []postData{}, http.StatusOK},
	//{"about", "/about", "GET", []postData{}, http.StatusOK},
	//{"contact", "/contact", "GET", []postData{}, http.StatusOK},
	//{"rooms", "/rooms", "GET", []postData{}, http.StatusOK},
	//{"deluxe-room", "/rooms/deluxe-room", "GET", []postData{}, http.StatusOK},
	//{"premium-suite", "/rooms/premium-suite", "GET", []postData{}, http.StatusOK},
	//{"search-availibility", "/search-availability", "GET", []postData{}, http.StatusOK},
	//{"reservation", "/make-reservation", "GET", []postData{}, http.StatusOK},
	//{"post-search-availibility", "/search-availability", "POST", []postData{
	//	{key: "arrive_date", value: "1/12/2026"},
	//	{key: "departure_date", value: "1/14/2026"},
	//}, http.StatusOK},
	//{"post-search-availibility-json", "/search-availability-json", "POST", []postData{
	//	{key: "arrive_date", value: "1/12/2026"},
	//	{key: "departure_date", value: "1/14/2026"},
	//}, http.StatusOK},
	//{"post-reservation", "/make-reservation", "POST", []postData{
	//	{key: "first_name", value: "John"},
	//	{key: "last_name", value: "smith"},
	//	{key: "email", value: "johnSmith@ex.com"},
	//	{key: "phone", value: "123-456-7890"},
	//}, http.StatusOK},
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

func TestRepository_Reserve(t *testing.T) {
	reservation := models.Reservation{
		RoomID: 8,
		Room: models.Room{
			ID:       8,
			RoomName: "Premium Suite",
		},
	}

	req, _ := http.NewRequest("GET", "/make-reservation", nil)
	ctx := getCtx(req)
	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()
	session.Put(ctx, "reservation", reservation)

	handler := http.HandlerFunc(Repo.Reserve)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Reservation handler returned http code %d instead of %d", rr.Code, http.StatusOK)
	}
}

func getCtx(req *http.Request) context.Context {
	ctx, err := session.Load(req.Context(), req.Header.Get("X-Session"))
	if err != nil {
		log.Println(err)
	}
	return ctx
}
