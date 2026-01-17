package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/AutomationMK/bookings/internal/config"
	"github.com/AutomationMK/bookings/internal/driver"
	"github.com/AutomationMK/bookings/internal/forms"
	"github.com/AutomationMK/bookings/internal/helpers"
	"github.com/AutomationMK/bookings/internal/models"
	"github.com/AutomationMK/bookings/internal/render"
	"github.com/AutomationMK/bookings/internal/repository"
	"github.com/AutomationMK/bookings/internal/repository/dbrepo"
)

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
}

// Repo the repository used by the handlers
var Repo *Repository

// NewRepo creates a new Repository
func NewRepo(a *config.AppConfig, db *driver.DB) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewPostgresRepo(db.SQL, a),
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// send the data to the template
	render.Template(w, r, "about.page.tmpl", &models.TemplateData{})
}

// Reserve handles the make-reservation page
func (m *Repository) Reserve(w http.ResponseWriter, r *http.Request) {
	var emptyReservation models.Reservation
	data := make(map[string]any)
	data["reservation"] = emptyReservation

	render.Template(w, r, "reservation.page.tmpl", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}

// PostReserve handles the posting of a reservation form
func (m *Repository) PostReserve(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	// parse date data
	ad := r.Form.Get("arrival_date")
	dd := r.Form.Get("departure_date")
	layout := "1/2/2006"
	arrivalDate, err := time.Parse(layout, ad)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	departureDate, err := time.Parse(layout, dd)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	// convert room_id to integer
	roomID, err := strconv.Atoi(r.Form.Get("room_id"))
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	reservation := models.Reservation{
		FirstName:     r.Form.Get("first_name"),
		LastName:      r.Form.Get("last_name"),
		Email:         r.Form.Get("email"),
		Phone:         r.Form.Get("phone"),
		ArrivalDate:   arrivalDate,
		DepartureDate: departureDate,
		RoomID:        roomID,
	}

	form := forms.New(r.PostForm)

	form.Required("first_name", "last_name", "email", "arrival_date", "departure_date", "room_id")
	form.MinLength("first_name", 3)
	form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]any)
		data["reservation"] = reservation

		render.Template(w, r, "reservation.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}

	newReservationID, err := m.DB.InsertReservation(reservation)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	restriction := models.RoomRestriction{
		ArrivalDate:   arrivalDate,
		DepartureDate: departureDate,
		RoomID:        roomID,
		ReservationID: newReservationID,
		RestrictionID: 1,
	}

	err = m.DB.InsertRoomRestriction(restriction)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	m.App.Session.Put(r.Context(), "reservation", reservation)

	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)
}

func (m *Repository) ReservationSummary(w http.ResponseWriter, r *http.Request) {
	reservation, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		m.App.ErrorLog.Println("cannot get item from sessio")
		m.App.Session.Put(r.Context(), "error", "Can't get reservation from session")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	m.App.Session.Remove(r.Context(), "reservation")
	data := make(map[string]any)
	data["reservation"] = reservation

	render.Template(w, r, "reservation-summary.page.tmpl", &models.TemplateData{
		Data: data,
	})
}

// Availability handles the search-availability page
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "search-availability.page.tmpl", &models.TemplateData{})
}

// PostAvailability handles post data from search-availability page
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	// parse date data
	ad := r.Form.Get("arrival_date")
	dd := r.Form.Get("departure_date")
	layout := "1/2/2006"
	arrivalDate, err := time.Parse(layout, ad)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	departureDate, err := time.Parse(layout, dd)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	rooms, err := m.DB.SearchAvailabilityForAllRooms(arrivalDate, departureDate)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	for _, room := range rooms {
		m.App.InfoLog.Println("Room:", room.ID, room.RoomName)
	}

	w.Write([]byte(fmt.Sprintf("start date is %s and end date is %s", ad, dd)))
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

// AvailabilityJSON handles request for availability and sends JSON response
func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK:      true,
		Message: "Available!",
	}

	out, err := json.MarshalIndent(resp, "", "     ")
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

// Contact handles the contact page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "contact.page.tmpl", &models.TemplateData{})
}

// Rooms handles the Rooms page
func (m *Repository) Rooms(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "rooms.page.tmpl", &models.TemplateData{})
}

// Deluxe handles the Deluxe room page
func (m *Repository) Deluxe(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "deluxe-room.page.tmpl", &models.TemplateData{})
}

// Premium handles the Premium suite page
func (m *Repository) Premium(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "premium-suite.page.tmpl", &models.TemplateData{})
}
